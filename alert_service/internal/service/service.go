package service

import (
	"context"
	"github.com/supperdoggy/alert/alert_proto"
	"github.com/supperdoggy/alert/alert_service/internal/databaseClient"
	"github.com/supperdoggy/alert/alert_service/internal/notification_sender"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type IService interface {
	GetAlerts(*alert_proto.GetAlertsRequest, alert_proto.NotificationService_GetAlertsServer) error
	GetAllActiveAlerts(ctx context.Context, in *alert_proto.GetAllActiveAlertsRequest) (*alert_proto.GetAllActiveAlertsResponse, error)
}

type service struct {
	logger *zap.Logger

	// just to add to the pull
	// maybe rewrite it when have some time
	sender *notification_sender.ISender

	db databaseClient.IDatabaseClient
}

func NewService(l *zap.Logger, sender *notification_sender.ISender, db databaseClient.IDatabaseClient) IService {
	return &service{
		logger: l,
		sender: sender,
		db: db,
	}
}

func (s *service) GetAlerts(req *alert_proto.GetAlertsRequest, stream alert_proto.NotificationService_GetAlertsServer) error {
	closeSignal := make(chan struct{})

	(*s.sender).AddNewClientStream(stream, req.UserId, req.Token, closeSignal)

	_ = <- closeSignal
	return nil
}

func (s *service) GetAllActiveAlerts(ctx context.Context, in *alert_proto.GetAllActiveAlertsRequest) (*alert_proto.GetAllActiveAlertsResponse, error) {
	if in == nil || in.GetUserId() == "" || in.GetToken() == "" {
		s.logger.Error("error fill all request fields", zap.Any("in", in))
		return nil, status.Error(codes.InvalidArgument, "fill all request fields")
	}

	response, err := s.db.GetAllActiveAlerts(context.Background(), in)
	if err != nil {
		code := status.Code(err)
		if code == codes.InvalidArgument {
			return nil, status.Error(code, "invalid argument given")
		}

		if code == codes.NotFound {
			return nil, err
		}

		s.logger.Error("got internal error from db service", zap.Error(err))
		return nil, status.Error(code, "internal db service error")
	}

	return response, nil
}

