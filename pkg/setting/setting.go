package setting

import (
	"github.com/go-ini/ini"
	"log"
	"time"
)

var (
	Cfg *ini.File

	// Base
	RunMode string

	// App
	PageSize  int
	JwtSecret string

	// Server
	HTTPPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
)

func init() {
	var err error
	configPath := "conf/app.ini"
	Cfg, err = ini.Load(configPath)
	if err != nil {
		log.Fatalf("Fail to parse '%s': %v", configPath, err)
	}

	loadBase()
	loadApp()
	loadServer()
}

func loadBase() {
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
}

func loadApp() {
	section, err := Cfg.GetSection("app")
	if err != nil {
		log.Fatalf("Fail to get section 'app': %v", err)
	}

	PageSize = section.Key("PAGE_SIZE").MustInt(10)
	JwtSecret = section.Key("APP_SECRET").MustString("!@)*#)!@U#@*!@!)")
}

func loadServer() {
	section, err := Cfg.GetSection("server")
	if err != nil {
		log.Fatalf("Fail to get section 'server': %v", err)
	}

	HTTPPort = section.Key("HTTP_PORT").MustInt(9999)
	ReadTimeout = time.Duration(section.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	WriteTimeout = time.Duration(section.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
}
