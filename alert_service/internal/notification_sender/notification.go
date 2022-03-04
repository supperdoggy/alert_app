package notification_sender

import (
	"github.com/pkg/errors"
	"github.com/supperdoggy/alert/alert_proto"
	"github.com/supperdoggy/alert/alert_service/internal/databaseClient"
	"go.uber.org/zap"
	"io"
	"sync"
)

type ISender interface {
	Start() error
	AddNewClientStream(stream alert_proto.AlertDatabaseService_GetAlertStreamServer)
}

type sender struct {
	logger *zap.Logger
	listener databaseClient.IDatabaseClient

	streamConnections []alert_proto.AlertDatabaseService_GetAlertStreamServer
	mut sync.Mutex
}

func NewSender(l *zap.Logger, listener databaseClient.IDatabaseClient) ISender {
	return &sender{
		logger:            l,
		listener:          listener,
		streamConnections: make([]alert_proto.AlertDatabaseService_GetAlertStreamServer, 0),
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
			msg := alert_proto.GetAlertStreamResponse{Alert: alert}
			err := v.Send(&msg)
			if err == io.EOF {
				// hope it works man
				s.streamConnections = append(s.streamConnections[:k], s.streamConnections[k+1:]...)
				continue
			}

			if err != nil {
				s.logger.Error("error sending response", zap.Error(err))
			}

		}
		s.mut.Unlock()
	}
}

func (s *sender) AddNewClientStream(stream alert_proto.AlertDatabaseService_GetAlertStreamServer) {
	s.mut.Lock()
	s.streamConnections = append(s.streamConnections, stream)
	s.mut.Unlock()
}
