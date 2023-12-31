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

// func SetHeader() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
// 		c.Next()
// 	}
// }

func Routers() *gin.Engine {
	if err := utils.Translator("zh"); err != nil {
		config.GVA_LOG.Error(err.Error())
		return nil
	}

	Router := gin.Default()

	//gin.SetMode(gin.DebugMode)

	// Router.Use(SetHeader())

	// corsConfig := cors.DefaultConfig()
	// corsConfig.AllowOrigins = []string{"http://localhost:8080", "http://localhost:5173", "http://127.0.0.1:8080"}
	// // corsConfig.AllowAllOrigins = true
	// corsConfig.AllowHeaders =  []string{"*"}
	// corsConfig.AllowCredentials = true
	// corsConfig.AllowMethods =  []string{"*"}
	// Router.Use(cors.New(corsConfig))

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"}
	corsConfig.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Cookie", "token", "Token"}
	corsConfig.AllowOrigins = []string{"http://localhost:8080", "http://127.0.0.1:8080", "http://localhost:5173"}
	corsConfig.AllowCredentials = true
	Router.Use(cors.New(corsConfig))

	store := cookie.NewStore([]byte("token"))
	store.Options(sessions.Options{
		SameSite: http.SameSiteNoneMode,
	})
	Router.Use(sessions.Sessions("token", store))

	Router.Use(auth.CookieAuth())

	Router.Use(middleware.Recovery())
	Router.Use(middleware.Logger())

	ApiGroup := Router.Group("api/v1")
	app.InitRouter(ApiGroup)
	room.InitRouter(ApiGroup)

	dev.InitRouter(ApiGroup)
	user.InitRouter(ApiGroup)

	return Router
}
