package main

import (
	"fmt"
	"github.com/DowneyL/the-way-to-gin/pkg/setting"
	"github.com/DowneyL/the-way-to-gin/routers"
	"net/http"
)

func main() {
	engine := routers.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        engine,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	_ = s.ListenAndServe()
}
