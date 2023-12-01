package dev

import (
	"github.com/gin-gonic/gin"
	"github.com/oikomi/FatBearServer/pkg/router"
)

func InitRouter(g *gin.RouterGroup) {
	r := router.NewRouter(g.Group("dev"))
	a := NewDevApi()
	r.BindApi("", a)
	r.BindPost("cmd", a.SendCmd)
	r.BindGet("cmd", a.GetCmd)

	r.BindPost("login", a.Login)

	r.BindGet("order", a.OrderList)
	r.BindPost("order", a.Order)

	// dev order
	devOrderRouter := router.NewRouter(g.Group("devorder"))
	devOrderApi := NewDevOrderApi()
	devOrderRouter.BindApi("", devOrderApi)
	devOrderRouter.BindPost("order", devOrderApi.Order)
	devOrderRouter.BindGet("order", devOrderApi.OrderList)

	r.BindGet("set", a.Set)
	r.BindPost("set", a.Set)

}
