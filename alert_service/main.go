package main

import (
	"fmt"
	"github.com/supperdoggy/alert/alert_proto"
	cfg "github.com/supperdoggy/alert/alert_service/internal/config"
	"github.com/supperdoggy/alert/alert_service/internal/databaseClient"
	"github.com/supperdoggy/alert/alert_service/internal/notification_sender"
	service2 "github.com/supperdoggy/alert/alert_service/internal/service"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
	"runtime"
	"time"
)

func main() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err.Error())
	}

	config := cfg.GetConfig()
	dbclient, err := databaseClient.NewDatabaseClient(config.DBServiceURL, config.DBServicePort, logger, grpc.WithInsecure())
	notificationClient := notification_sender.NewSender(logger, dbclient)
	service := service2.NewService(logger, &notificationClient, dbclient)

	urlport := fmt.Sprintf("%s:%s", config.Url, config.Port)
	listener, err := net.Listen("tcp", urlport)
	if err != nil {
		logger.Fatal("error creating listener", zap.Any("config", config), zap.Error(err))
	}

	grpcServer := grpc.NewServer()
	alert_proto.RegisterNotificationServiceServer(grpcServer, service)

	go DisplayNumberOfGoroutines(*logger)

	logger.Info("starting server on", zap.Any("listener", urlport))
	go func() {
		err := notificationClient.Start()
		if err != nil {
			logger.Error("got error starting sending notification")
		}
	} ()

	if err := grpcServer.Serve(listener); err != nil {
		logger.Fatal("fatal error serving db server", zap.Error(err))
	}
}

// DisplayNumberOfGoroutines - once in a minute displays amount of routines
func DisplayNumberOfGoroutines(l zap.Logger) {
	for {
		l.Info("number of routines", zap.Any("routines", runtime.NumGoroutine()))
		time.Sleep(1 * time.Minute)
	}
}
