package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"github.com/oikomi/FatBearServer/config"
	docs "github.com/oikomi/FatBearServer/docs"
	"github.com/oikomi/FatBearServer/initialize"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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

//	@title			Swagger Example API
//	@version		1.0
//	@description	This is a sample server Petstore server.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

// @host	127.0.0.1:8080
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

	HealthGroup := router.Group("")
	{
		// 健康监测
		HealthGroup.GET("/health", HealthCheck)
	}

	// 初始化Admin
	initialize.Admin(router)

	// initialize.Swagger(router)

	// docs.SwaggerInfo.Host = "127.0.0.1:8080"
	// docs.SwaggerInfo.Host = "120.55.60.98:8080"
	docs.SwaggerInfo.Host = "54.146.155.159:8080"
	// url := ginSwagger.URL("http://127.0.0.1:8080/swagger/doc.json") // The url pointing to API definition
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	runServer(router)
}

// @Summary		HealthCheck
// @Description	HealthCheck
// @Accept			json
// @Produce		json
// @Success		200	{string}	string	"ok"
// @Router			/health [get]
func HealthCheck(g *gin.Context) {
	g.JSON(http.StatusOK, "ok")
}
