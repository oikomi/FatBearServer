package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"github.com/oikomi/FatBearServer/config"
	"github.com/oikomi/FatBearServer/initialize"
	"go.uber.org/zap"
)

type server interface {
	ListenAndServe() error
}

func runTlsServer(r *gin.Engine) {
	err := http.ListenAndServeTLS(":443",
		"/Users/harold/godev/src/github.com/oikomi/FatBearServer/api.missvib.com.pem",
		"/Users/harold/godev/src/github.com/oikomi/FatBearServer/api.missvib.com.key", nil)
	if err != nil {
		config.GVA_LOG.Error(err.Error())
	}

}

func runServer(r *gin.Engine) {
	address := ":8080"
	s := initServer(address, r)

	// go runTlsServer(r)

	if err := s.ListenAndServe(); err != nil {
		config.GVA_LOG.Error(err.Error())
	}
}

func initServer(address string, router *gin.Engine) server {
	s := endless.NewServer(address, router)
	s.ReadHeaderTimeout = 20 * time.Second
	s.WriteTimeout = 20 * time.Second
	s.MaxHeaderBytes = 1 << 20
	return s
}

func main() {
	// 初始化配置
	config.GVA_VP = initialize.Viper()

	// 初始化日志
	config.GVA_LOG = initialize.Zap()
	zap.ReplaceGlobals(config.GVA_LOG)

	// 初始化数据库
	config.GVA_DB = initialize.Gorm() // gorm连接数据库
	if config.GVA_DB != nil {
		// 程序结束前关闭数据库链接
		db, _ := config.GVA_DB.DB()
		defer db.Close()
	} else {
		fmt.Println("数据库启动失败...")
		return
	}

	// 初始化Router
	router := initialize.Routers()
	if router == nil {
		return
	}

	// 初始化Admin
	initialize.Admin(router)

	runServer(router)
}
