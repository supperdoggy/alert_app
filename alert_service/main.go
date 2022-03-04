package main

import (
	cfg "github.com/supperdoggy/alert/alert_service/internal/config"
	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err.Error())
	}

	config := cfg.GetConfig()


}
