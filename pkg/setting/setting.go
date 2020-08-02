package setting

import (
	"github.com/go-ini/ini"
	"log"
	"time"
)

type App struct {
	JwtSecret         string
	PageSize          int
	RuntimeRootPath   string
	ImagePrefixUrl    string
	ImageSavePath     string
	ImageMaxSize      int
	ImageAllowExtends []string
	LogSavePath       string
	LogSaveName       string
	LogFileExtend     string
	TimeFormat        string
}

type Server struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type Database struct {
	Type        string
	User        string
	Password    string
	Host        string
	Name        string
	TablePrefix string
}

type Redis struct {
	Host        string
	Password    string
	MaxIdle     int
	MaxActive   int
	IdleTimeout time.Duration
}

var (
	Cfg             *ini.File
	AppSetting      = &App{}
	ServerSetting   = &Server{}
	DatabaseSetting = &Database{}
	RedisSetting    = &Redis{}
)

func Setup() {
	Cfg, err := ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("Faile to parse 'conf/app.ini': %v", err)
	}

	if err := Cfg.Section("app").MapTo(AppSetting); err != nil {
		log.Fatalf("Cfg.MapTo AppSetting err: %v\n", err)
	} else {
		AppSetting.ImageMaxSize = AppSetting.ImageMaxSize * 1024 * 1024
	}

	if err := Cfg.Section("server").MapTo(ServerSetting); err != nil {
		log.Fatalf("Cfg.MapTo ServerSetting err: %v\n", err)
	} else {
		ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
		ServerSetting.WriteTimeout = ServerSetting.WriteTimeout * time.Second
	}

	if err := Cfg.Section("database").MapTo(DatabaseSetting); err != nil {
		log.Fatalf("Cfg.MapTo DatabaseSetting err: %v\n", err)
	}

	if err := Cfg.Section("redis").MapTo(RedisSetting); err != nil {
		log.Fatalf("Cfg.MapTo RedisSetting err: %v\n", err)
	}
}
