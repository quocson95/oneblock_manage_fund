package main

import (
	"fmt"
	"oneblock_manage_fund/config"
	"oneblock_manage_fund/controllers"
	"oneblock_manage_fund/models"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func init() {
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.TimeKey = "time"
	config.EncoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder
	config.DisableCaller = false
	config.DisableStacktrace = true
	logger, _ := config.Build(zap.AddCaller())
	zap.ReplaceGlobals(logger)
}

func main() {
	// load config
	{
		zap.L().Info("read config")
		configFile := "config.yaml"
		f, err := os.Open(configFile)
		if err != nil {
			panic(err)
		}
		if err = config.Parse(f); err != nil {
			panic(err)
		}
	}
	// init db
	{
		dsn := config.Get().DB.DSN
		zap.L().With(zap.String("dsn", dsn)).Info("connect db")
		err := models.DefaulDB.Init(dsn, &models.User{})
		if err != nil {
			panic(err)
		}
	}

	r := gin.Default()
	// cors
	{
		r.Use(cors.New(cors.Config{
			AllowOrigins:     []string{"*"},
			AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "HEAD"},
			AllowHeaders:     []string{"*"},
			ExposeHeaders:    []string{"*"},
			AllowCredentials: true,
			AllowOriginFunc: func(origin string) bool {
				return true
			},
			MaxAge: 12 * time.Hour,
		}))
	}
	r.GET("/api/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	new(controllers.Account).Router(r.Group("/api/account"))
	oauthController := &controllers.Oath2Controller{}
	oauthController.Init(config.Get().Google.ID, config.Get().Google.Secret, config.Get().Google.CallbackSSO)
	oauthController.Router(r.Group("/api/auth"))
	new(controllers.LoginController).Router(r.Group("/auth"))
	new(controllers.BinanceController).Router(r.Group("/binance"))
	// add defaul user
	// go createDefaultUser()
	r.Run(fmt.Sprintf(":%d", config.Get().Be.Port)) // listen and serve on 0.0.0.0:8080
}

func createDefaultUser() {
	emails := []string{
		"dangquocson1995@gmail.com",
		"nghuuloc512@gmail.com",
		"nguyentrungbmt17@gmail.com",
		"haotran1689@gmail.com",
		"dtoan.bui@gmail.com",
	}
	for _, email := range emails {
		u := models.User{
			UserName:     email,
			Email:        email,
			Role:         models.RoleUserAdmin,
			UsdtInWallet: 0,
		}
		if err := u.Create(); err != nil {
			zap.L().With(zap.Error(err)).With(zap.String("email", email)).Error("add user failed")
		}
	}
}
