package main

import (
	"context"
	"fmt"
	cfg "github.com/supperdoggy/alert/alert_client/internal/config"
	"github.com/supperdoggy/alert/alert_client/internal/db_client"
	"github.com/supperdoggy/alert/alert_client/internal/notification_client"
	"github.com/supperdoggy/alert/alert_proto"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"time"
)

func main() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err.Error())
	}

	config := cfg.GetConfig()
	db, err := db_client.NewAlertDBServiceClient(config.DBServiceURL, config.DBServicePort, logger, grpc.WithInsecure())
	if err != nil {
		logger.Fatal("error connecting to db service client", zap.Error(err), zap.Any("config", config))
	}

	// create 100 connections to listen to the notifications
	for i:=0; i<100; i++ {
		logger.Info("waiting for msg", zap.Any("i", i))
		go func() {
			numstr := fmt.Sprintf("%v", i)
			client, err := notification_client.NewNotificationClient(config.NotificationServiceURL, config.NotificationServicePort, logger, grpc.WithInsecure())
			if err != nil {
				logger.Error("error creating notification client", zap.Error(err), zap.Any("i", numstr))
				return
			}
			stream, err := client.GetAlerts(context.Background(),
				&alert_proto.GetAlertsRequest{Token: numstr, UserId: numstr})
			if err != nil {
				logger.Error("error calling Getalerts", zap.Error(err), zap.Any("i", numstr))
				return
			}
			for {
				msg, err := stream.Recv()
				if err != nil {
					logger.Error("error recv msg", zap.Error(err), zap.Any("i", numstr))
					return
				}
				logger.Info("got msg", zap.Any("msg", msg), zap.Any("i", numstr))
			}
		}()
	}

	// send alert once in 5 secs
	go func() {
		t := 0
		for {
			id := fmt.Sprintf("%v", t)
			_, err := db.NewAlert(context.Background(), &alert_proto.NewAlertRequest{Alert: &alert_proto.Alert{
				Tag:                "tag"+id,
				Address:            "address"+id,
				Title:              "title"+id,
				Body:               "body"+id,
				ExpirationDatetime: time.Now().Unix(),
			}})
			if err != nil {
				logger.Error("error calling NewAlert", zap.Error(err))
			}
			t++
			time.Sleep(5 * time.Second)
		}
	}()

	time.Sleep(1 * time.Hour)
}

func createNewTestAlert(db db_client.IDatabaseServiceClient, l *zap.Logger) {

}
