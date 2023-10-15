package initialize

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/oikomi/FatBearServer/config"
	"github.com/oikomi/FatBearServer/internal/app"
	"github.com/oikomi/FatBearServer/internal/dev"
	"github.com/oikomi/FatBearServer/internal/room"
	"github.com/oikomi/FatBearServer/internal/user"
	"github.com/oikomi/FatBearServer/middleware"
	"github.com/oikomi/FatBearServer/pkg/auth"
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

	store := cookie.NewStore([]byte("token"))
	store.Options(sessions.Options{
		MaxAge: 60,
	})
	Router.Use(sessions.Sessions("token", store))

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"*"}
	Router.Use(cors.New(corsConfig))

	Router.Use(auth.CookieAuth())

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
	user.InitRouter(ApiGroup)

	return Router
}
