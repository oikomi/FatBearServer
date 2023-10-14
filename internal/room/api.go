package room

import (
	"github.com/gin-gonic/gin"
	"github.com/oikomi/FatBearServer/pkg/api"
	"github.com/oikomi/FatBearServer/pkg/response"
)

type RoomApi struct {
	api.Api
	Service RoomService
}

func NewRoomApi() RoomApi {
	var r Room
	s := NewRoomService(r)
	baseApi := api.NewApi[Room](s)
	return RoomApi{Api: baseApi, Service: s}
}

func (r RoomApi) CreateRoom(c *gin.Context) {
	err := r.Service.CreateRoom(c)
	if err != nil {
		response.FailWithError(err, c)
	}

	response.Ok(c)
}
