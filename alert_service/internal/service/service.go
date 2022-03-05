package service

import (
	"github.com/supperdoggy/alert/alert_proto"
	"github.com/supperdoggy/alert/alert_service/internal/notification_sender"
	"go.uber.org/zap"
	"time"
)

// TODO ADD NEW CONNECTION AND SERVE THEM

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
	// IVE MADE SOME SHIT THAT WONT WORK IF NOT SLEEP GOROUTINE
	(*s.sender).AddNewClientStream(stream, req.UserId, req.Token)
	time.Sleep(1*time.Hour)
	return nil
}
