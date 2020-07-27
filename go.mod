module github.com/DowneyL/the-way-to-gin

go 1.14

require (
	github.com/astaxie/beego v1.12.2
	github.com/gin-gonic/gin v1.6.3
	github.com/go-ini/ini v1.57.0
	github.com/go-playground/validator/v10 v10.3.0 // indirect
	github.com/go-sql-driver/mysql v1.5.0
	github.com/golang/protobuf v1.4.2 // indirect
	github.com/jinzhu/gorm v1.9.15
	github.com/json-iterator/go v1.1.10 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.1 // indirect
	github.com/smartystreets/goconvey v1.6.4 // indirect
	github.com/unknwon/com v1.0.1
	golang.org/x/sys v0.0.0-20200724161237-0e2f3a69832c // indirect
	google.golang.org/protobuf v1.25.0 // indirect
	gopkg.in/ini.v1 v1.57.0 // indirect
	gopkg.in/yaml.v2 v2.3.0 // indirect
)

replace (
	github.com/DowneyL/the-way-to-gin/conf => ./conf
	github.com/DowneyL/the-way-to-gin/middleware => ./middleware
	github.com/DowneyL/the-way-to-gin/models => ./models
	github.com/DowneyL/the-way-to-gin/pkg/e => ./pkg/e
	github.com/DowneyL/the-way-to-gin/pkg/setting => ./pkg/setting
	github.com/DowneyL/the-way-to-gin/pkg/util => ./pkg/util
	github.com/DowneyL/the-way-to-gin/routers => ./routers
)
