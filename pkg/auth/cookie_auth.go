package auth

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/oikomi/FatBearServer/config"
	"github.com/oikomi/FatBearServer/pkg/response"
	"github.com/oikomi/FatBearServer/utils"
)

func CookieAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		if strings.Contains(c.Request.RequestURI, "api/v1/user/login") {
			c.Next()
			return
		}
		token, err := c.Cookie("token")
		if err != nil {
			config.GVA_LOG.Error(err.Error())
			response.FailWithStatusCode(401, "未登录或非法访问", c)
			// c.Redirect(http.StatusFound, "/login")
			return
		}
		if token == "" {
			config.GVA_LOG.Error("token is empty")
			response.FailWithStatusCode(401, "未登录或非法访问", c)
			// c.Redirect(http.StatusFound, "/login")
			return
		}

		j := utils.NewJWT()
		claims, err := j.ParseToken(token)
		if err != nil {
			config.GVA_LOG.Error(err.Error())
			response.FailWithStatusCode(401, "未登录或非法访问", c)
			// c.Redirect(http.StatusFound, "/login")
			return
		}
		c.Set("user", claims.Username)
		c.Next()
	}
}
