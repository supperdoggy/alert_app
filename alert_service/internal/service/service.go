package service

import (
	"github.com/supperdoggy/alert/alert_proto"
	"github.com/supperdoggy/alert/alert_service/internal/notification_sender"
	"go.uber.org/zap"
)

type IService interface {
	GetAlerts(*alert_proto.GetAlertsRequest, alert_proto.NotificationService_GetAlertsServer) error

}

type service struct {
	logger *zap.Logger

	// just to add to the pull
	// maybe rewrite it when have some time
	sender *notification_sender.ISender
}

func NewService(l *zap.Logger, sender *notification_sender.ISender) IService {
	return &service{
		logger: l,
		sender: sender,
	}
}

func (s *service) GetAlerts(req *alert_proto.GetAlertsRequest, stream alert_proto.NotificationService_GetAlertsServer) error {
	closeSignal := make(chan struct{})

	(*s.sender).AddNewClientStream(stream, req.UserId, req.Token, closeSignal)

	_ = <- closeSignal
	return nil
}
