package room

import (
	"github.com/gin-gonic/gin"
	"github.com/oikomi/FatBearServer/pkg/router"
)

func InitRouter(g *gin.RouterGroup) {
	r := router.NewRouter(g.Group("room"))
	a := NewRoomApi()
	r.BindApi("", a)
	r.BindGet("create", a.CreateRoom) //
	r.BindGet("hello", a.Hello)       //
}
