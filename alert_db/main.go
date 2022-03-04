package main

import (
	"fmt"
	cfg "github.com/supperdoggy/alert/alert_db/internal/config"
	db2 "github.com/supperdoggy/alert/alert_db/internal/db"
	service2 "github.com/supperdoggy/alert/alert_db/internal/service"
	"github.com/supperdoggy/alert/alert_proto"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

func main() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err.Error())
	}

	config := cfg.GetConfig()
	db, err := db2.NewDB(logger, config.Url, config.DBName, config.AlertsCollectionName)
	if err != nil {
		logger.Fatal("error connecting to db", zap.Error(err), zap.Any("config", config))
	}

	service := service2.NewService(logger, db)

	urlport := fmt.Sprintf("%s:%s", config.Url, config.Port)
	listener, err := net.Listen("tcp", urlport)
	if err != nil {
		logger.Fatal("error creating listener", zap.Any("config", config), zap.Error(err))
	}

	grpcServer := grpc.NewServer()
	alert_proto.RegisterAlertDatabaseServiceServer(grpcServer, service)
	if !config.IsProd {
		reflection.Register(grpcServer)
	}

	logger.Info("starting server on", zap.Any("listener", urlport))

	if err := grpcServer.Serve(listener); err != nil {
		logger.Fatal("fatal error serving db server", zap.Error(err))
	}
}
