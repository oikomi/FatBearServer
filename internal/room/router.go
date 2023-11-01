package room

import (
	"github.com/gin-gonic/gin"
	"github.com/oikomi/FatBearServer/pkg/router"
)

func InitRouter(g *gin.RouterGroup) {
	r := router.NewRouter(g.Group("room"))
	a := NewRoomApi()
	r.BindApi("", a)
	r.BindPost("create", a.CreateRoom)
	r.BindGet("list", a.GetRoomList)
	r.BindPost("update", a.UpdateRoom)
	r.BindGet("msg", a.GetRoomMsg)


	msgRouter := router.NewRouter(g.Group("roommsg"))
	rooMsgApi := NewRoomMsgApi()
	msgRouter.BindApi("", rooMsgApi)
	msgRouter.BindGet("msg", rooMsgApi.GetRoomMsg)
	msgRouter.BindPost("msg", rooMsgApi.SendRoomMsg)

}
