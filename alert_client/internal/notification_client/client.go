package notification_client

import (
	"context"
	"fmt"
	"github.com/supperdoggy/alert/alert_proto"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type INotificationClient interface {
	GetAlerts(ctx context.Context, in *alert_proto.GetAlertsRequest, opts ...grpc.CallOption) (alert_proto.NotificationService_GetAlertsClient, error)
}

type notificationClient struct {
	logger *zap.Logger
	client alert_proto.NotificationServiceClient
}

func NewNotificationClient(url, port string, l *zap.Logger, option ...grpc.DialOption) (INotificationClient, error) {
	urlport := fmt.Sprintf("%s:%s", url, port)
	conn, err := grpc.Dial(urlport, option...)
	if err != nil {
		l.Error("error dialing with alert database client", zap.Error(err), zap.Any("url", url))
		return nil, err
	}

	client := alert_proto.NewNotificationServiceClient(conn)
	return &notificationClient{
		logger: l,
		client: client,
	}, nil

}

func (n *notificationClient) GetAlerts(ctx context.Context, in *alert_proto.GetAlertsRequest, opts ...grpc.CallOption) (alert_proto.NotificationService_GetAlertsClient, error) {
	res, err := n.client.GetAlerts(ctx, in, opts...)
	if err != nil {
		n.logger.Error("error getting alerts", zap.Error(err), zap.Any("in", in))
	}
	return res, err
}
