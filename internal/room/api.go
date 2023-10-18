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
		return
	}

	response.Ok(c)
}

func (r RoomApi) GetRoomList(c *gin.Context) {
	rooms, err := r.Service.GetRoomList(c)
	if err != nil {
		response.FailWithError(err, c)
		return
	}

	response.OkWithData(rooms, c)
}

func (r RoomApi) UpdateRoom(c *gin.Context) {
	err := r.Service.UpdateRoom(c)
	if err != nil {
		response.FailWithError(err, c)
		return
	}

	response.Ok(c)
}
