package initialize

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oikomi/FatBearServer/config"
	"github.com/oikomi/FatBearServer/internal/app"
	"github.com/oikomi/FatBearServer/internal/dev"
	"github.com/oikomi/FatBearServer/internal/room"
	"github.com/oikomi/FatBearServer/middleware"
	"github.com/oikomi/FatBearServer/utils"
)

func HealthCheck(g *gin.Context) {
	g.JSON(http.StatusOK, "ok...")
}

func Routers() *gin.Engine {

	if err := utils.Translator("zh"); err != nil {
		config.GVA_LOG.Error(err.Error())
		return nil
	}

	Router := gin.Default()
	//gin.SetMode(gin.DebugMode)

	Router.Use(middleware.Recovery())
	Router.Use(middleware.Logger())
	HealthGroup := Router.Group("")
	{
		// 健康监测
		HealthGroup.GET("/health", HealthCheck)
	}

	ApiGroup := Router.Group("api/v1")
	app.InitRouter(ApiGroup)
	room.InitRouter(ApiGroup)

	dev.InitRouter(ApiGroup)

	return Router
}
