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
}
