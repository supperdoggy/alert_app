package service

import (
	"context"
	db2 "github.com/supperdoggy/alert/alert_db/internal/db"
	"github.com/supperdoggy/alert/alert_proto"
	"github.com/supperdoggy/alert/alerts_structs"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type IService interface {
	NewAlert(context.Context, *alert_proto.NewAlertRequest) (*alert_proto.NewAlertResponse, error)
	GetAlertStream(*alert_proto.GetAlertStreamRequest, alert_proto.AlertDatabaseService_GetAlertStreamServer) error
}

type service struct {
	l *zap.Logger
	db db2.IDB
}

var (
	alertChan = make(chan *alert_proto.Alert)
)

func NewService(l *zap.Logger, db db2.IDB) IService {
	return &service{
		l:  l,
		db: db,
	}
}

func (s *service) NewAlert(ctx context.Context, request *alert_proto.NewAlertRequest) (*alert_proto.NewAlertResponse, error) {
	if request.GetAlert() == nil {
		return nil, status.Error(codes.InvalidArgument, "nil alert")
	}

	a := alerts_structs.AlertDB{
		Title: request.GetAlert().GetTitle(),
		Body: request.GetAlert().GetBody(),
		Tag: request.GetAlert().GetTag(),
		Address: request.GetAlert().GetAddress(),
		ExpirationDateTime: request.GetAlert().GetExpirationDatetime(),
	}

	code, err := s.db.NewAlert(a)
	if err != nil {
		return nil, status.Error(code, err.Error())
	}

	// can block the stream god dammit
	// its just to make it work
	alertChan <- request.GetAlert()
	resp := alert_proto.NewAlertResponse{Alert: request.GetAlert()}

	return &resp, nil
}

func (s *service) GetAlertStream(req *alert_proto.GetAlertStreamRequest, stream alert_proto.AlertDatabaseService_GetAlertStreamServer) error {
	s.l.Info("got alert stream")
	select {
	case a := <- alertChan:
		resp := alert_proto.GetAlertStreamResponse{Alert: a}
		if err := stream.Send(&resp); err != nil {
			return status.Error(codes.Canceled, "closed stream")
		}
	}
	return nil
}
