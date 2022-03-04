package databaseClient

import (
	"context"
	"fmt"
	"github.com/supperdoggy/alert/alert_proto"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"io"
)

type IDatabaseClient interface {
	ListenGetAlertStream(ch chan *alert_proto.Alert)
}

type databaseClient struct {
	logger *zap.Logger

	client alert_proto.AlertDatabaseServiceClient
}

func NewDatabaseClient(url, port string, l *zap.Logger, option ...grpc.DialOption) (IDatabaseClient, error) {
	urlport := fmt.Sprintf("%s:%s", url, port)
	conn, err := grpc.Dial(urlport, option...)
	if err != nil {
		l.Error("error dialing with alert database client", zap.Error(err), zap.Any("url", url))
		return nil, err
	}

	client := alert_proto.NewAlertDatabaseServiceClient(conn)
	return &databaseClient{
		logger: l,
		client: client,
	}, nil

}

func (d *databaseClient) ListenGetAlertStream(ch chan *alert_proto.Alert) {
	ctx := context.Background()
	stream, err := d.getAlertStream(&ctx)
	if err != nil {
		d.logger.Fatal("error getting first GetAlertStream", zap.Error(err))
	}

	for {
		msg, err := stream.Recv()
		// if stream closes - get new stream
		if err == io.EOF {
			stream, err = d.getAlertStream(&ctx)
			if err != nil {
				d.logger.Error("error setting new stream", zap.Error(err))
			}
			continue
		}
		// if got other error then something is wrong
		if err != nil {
			d.logger.Fatal("got unexpected error while getting GetAlertStream", zap.Error(err))
		}
		ch <- msg.GetAlert()
	}
}

func (d *databaseClient) getAlertStream(ctx *context.Context) (alert_proto.AlertDatabaseService_GetAlertStreamClient, error) {
	stream, err := d.client.GetAlertStream(*ctx, &alert_proto.GetAlertStreamRequest{})
	return stream, err
}
