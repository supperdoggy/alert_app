package db_client

import (
	"context"
	"fmt"
	"github.com/supperdoggy/alert/alert_proto"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type IDatabaseServiceClient interface {
	NewAlert(ctx context.Context, in *alert_proto.NewAlertRequest, opts ...grpc.CallOption) (*alert_proto.NewAlertResponse, error)
}

type dbClient struct {
	logger *zap.Logger
	client alert_proto.AlertDatabaseServiceClient
}

func NewAlertDBServiceClient(url, port string, l *zap.Logger, option ...grpc.DialOption) (IDatabaseServiceClient, error) {
	urlport := fmt.Sprintf("%s:%s", url, port)
	conn, err := grpc.Dial(urlport, option...)
	if err != nil {
		l.Error("error dialing with alert database client", zap.Error(err), zap.Any("url", url))
		return nil, err
	}

	client := alert_proto.NewAlertDatabaseServiceClient(conn)
	return &dbClient{
		logger: l,
		client: client,
	}, nil
}

func (d *dbClient) NewAlert(ctx context.Context, in *alert_proto.NewAlertRequest, opts ...grpc.CallOption) (*alert_proto.NewAlertResponse, error) {
	resp, err := d.client.NewAlert(ctx, in, opts...)
	if err != nil {
		d.logger.Error("error NewAlert", zap.Error(err), zap.Any("in", in))
	}
	return resp, err
}
