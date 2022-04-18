module github.com/EDDYCJY/go-gin-example

go 1.16

require (
	github.com/gin-gonic/gin v1.7.7
	github.com/go-ini/ini v1.66.4
)

require (
	github.com/alecthomas/template v0.0.0-20190718012654-fb15b899a751
	github.com/astaxie/beego v1.12.3
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/go-openapi/swag v0.21.1 // indirect
	github.com/go-playground/validator/v10 v10.10.1 // indirect
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/gomodule/redigo v2.0.0+incompatible
	github.com/jinzhu/gorm v1.9.16
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/robfig/cron v1.2.0
	github.com/shiena/ansicolor v0.0.0-20200904210342-c7312218db18 // indirect
	github.com/swaggo/gin-swagger v1.2.0
	github.com/swaggo/swag v1.8.0
	github.com/ugorji/go v1.2.7 // indirect
	github.com/unknwon/com v1.0.1
	golang.org/x/crypto v0.0.0-20220315160706-3147a52a75dd // indirect
	golang.org/x/net v0.0.0-20220225172249-27dd8689420f // indirect
	golang.org/x/sys v0.0.0-20220319134239-a9b59b0215f8 // indirect
	golang.org/x/tools v0.1.10 // indirect
	google.golang.org/protobuf v1.27.1 // indirect
)

replace (
	github.com/EDDYCJY/go-gin-example/conf => ./go-gin-example/pkg/conf
	github.com/EDDYCJY/go-gin-example/docs => ./go-gin-example/docs
	github.com/EDDYCJY/go-gin-example/middleware => ./go-gin-example/middleware
	github.com/EDDYCJY/go-gin-example/models => ./go-gin-example/models
	github.com/EDDYCJY/go-gin-example/pkg/e => ./go-gin-example/pkg/e
	github.com/EDDYCJY/go-gin-example/pkg/logging => ./go-gin-example/pkg/logging
	github.com/EDDYCJY/go-gin-example/pkg/setting => ./go-gin-example/pkg/setting
	github.com/EDDYCJY/go-gin-example/pkg/util => ./go-gin-example/pkg/util
	github.com/EDDYCJY/go-gin-example/routers => ./go-gin-example/routers
	github.com/EDDYCJY/go-gin-example/setting => ./go-gin-example/setting
)
