package user

import (
	"github.com/gin-gonic/gin"
	"github.com/oikomi/FatBearServer/pkg/router"
)

func InitRouter(g *gin.RouterGroup) {
	r := router.NewRouter(g.Group("user"))
	a := NewUserApi()
	r.BindApi("", a)
	r.BindPost("login", a.Login)

	r.BindPost("addToken", a.AddToken)

	r.BindGet("getToken", a.GetToken)

	r.BindGet("info", a.GetUser)

}
