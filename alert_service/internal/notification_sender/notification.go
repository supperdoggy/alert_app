package notification_sender

import (
	"github.com/pkg/errors"
	"github.com/supperdoggy/alert/alert_proto"
	"github.com/supperdoggy/alert/alert_service/internal/databaseClient"
	"github.com/supperdoggy/alert/alerts_structs"
	"go.uber.org/zap"
	"io"
	"sync"
)

type ISender interface {
	Start() error
	AddNewClientStream(stream alert_proto.NotificationService_GetAlertsServer, userID, token string, doneSignal chan struct{})
}

type sender struct {
	logger *zap.Logger
	listener databaseClient.IDatabaseClient

	streamConnections []alerts_structs.ClientGetAlertStream
	mut sync.Mutex
}

func NewSender(l *zap.Logger, listener databaseClient.IDatabaseClient) ISender {
	return &sender{
		logger:            l,
		listener:          listener,
		streamConnections: make([]alerts_structs.ClientGetAlertStream, 0),
	}
}

func (s *sender) Start() error {
	ch := make(chan *alert_proto.Alert)
	go s.listener.ListenGetAlertStream(ch)

	for {
		alert, ok := <- ch
		if !ok {
			return errors.New("channel closed")
		}

		s.mut.Lock()
		// if we have no connections then we just loose msg for now
		for k, v := range s.streamConnections {
			msg := alert_proto.GetAlertsResponse{Alert: alert}

			err := v.Stream.Send(&msg)
			if err == io.EOF {
				// hope it works man
				// removes the stream from pull of streams
				s.streamConnections = append(s.streamConnections[:k], s.streamConnections[k+1:]...)
				v.CloseSignal <- struct{}{}
				continue
			}

			if err != nil {
				s.logger.Error("error sending response", zap.Error(err))
				s.streamConnections = append(s.streamConnections[:k], s.streamConnections[k+1:]...)
				v.CloseSignal <- struct{}{}
				continue
			}
		}
		s.mut.Unlock()
	}
}

func (s *sender) AddNewClientStream(stream alert_proto.NotificationService_GetAlertsServer, userID, token string, closeSignal chan struct{}) {
	s.mut.Lock()
	s.streamConnections = append(s.streamConnections, alerts_structs.ClientGetAlertStream{
		Stream: stream, UserID: userID, Token: token, CloseSignal: closeSignal,
	})
	s.mut.Unlock()
}
