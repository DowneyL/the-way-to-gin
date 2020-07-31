package main

import (
	"context"
	"fmt"
	"github.com/DowneyL/the-way-to-gin/models"
	"github.com/DowneyL/the-way-to-gin/pkg/logging"
	"github.com/DowneyL/the-way-to-gin/pkg/setting"
	"github.com/DowneyL/the-way-to-gin/routers"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	setting.Setup()
	models.Setup()
	logging.Setup()

	engine := routers.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.ServerSetting.HttpPort),
		Handler:        engine,
		ReadTimeout:    setting.ServerSetting.ReadTimeout,
		WriteTimeout:   setting.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := s.ListenAndServe(); err != nil {
			logging.InfoPrintf("Listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	logging.Info("Shutdown Server ...")

	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()
	if err := s.Shutdown(ctx); err != nil {
		logging.InfoPrintf("Server Shutdown", err)
	}

	logging.Info("Server exiting")
}
