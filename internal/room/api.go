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

func (r RoomApi) Heartbeat(c *gin.Context) {
	err := r.Service.Heartbeat(c)
	if err != nil {
		response.FailWithError(err, c)
		return
	}

	response.Ok(c)
}

// @Description	CreateRoom
// @Accept			json
// @Produce		json
// @Param			super_token	header	string			false	"Authentication header"
// @Param			createRoom	body	CreateRoomReq	true	"create room"
// @Router			/api/v1/room/create [post]
func (r RoomApi) CreateRoom(c *gin.Context) {
	err := r.Service.CreateRoom(c)
	if err != nil {
		response.FailWithError(err, c)
		return
	}

	response.Ok(c)
}

// @Description	GetRoomList
// @Accept			json
// @Produce		json
// @Param			super_token	header	string	false	"Authentication header"
// @Success		200			{array}	Room
// @Router			/api/v1/room/list [get]
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

// @Description	updateRoomStatus
// @Accept			json
// @Produce		json
// @Param			super_token			header	string				false	"Authentication header"
// @Param			updateRoomStatus	body	updateRoomStatusReq	true	"updateRoomStatus"
// @Router			/api/v1/room/updateRoomStatus [post]
func (r RoomApi) updateRoomStatus(c *gin.Context) {
	err := r.Service.updateRoomStatus(c)
	if err != nil {
		response.FailWithError(err, c)
		return
	}

	response.Ok(c)
}

// @Description	GetRoomMsg
// @Accept			json
// @Produce		json
// @Param			super_token	header	string	false	"Authentication header"
// @Success		200			{array}	RoomMsg
// @Router			/api/v1/room/msg [get]
func (r RoomApi) GetRoomMsg(c *gin.Context) {
	msgs, err := r.Service.GetRoomMsg(c)
	if err != nil {
		response.FailWithError(err, c)
		return
	}

	response.OkWithData(msgs, c)
}

// @Description	GetRoomMsg
// @Accept			json
// @Produce		json
// @Param			super_token		header	string			false	"Authentication header"
// @Param			SendRoomMsgReq	body	SendRoomMsgReq	true	"send room msg"
// @Router			/api/v1/room/msg [post]
func (r RoomApi) SendRoomMsg(c *gin.Context) {
	err := r.Service.SendRoomMsg(c)
	if err != nil {
		response.FailWithError(err, c)
		return
	}

	response.Ok(c)
}

// room msg

type RoomMsgApi struct {
	api.Api
	Service RoomMsgService
}

func NewRoomMsgApi() RoomMsgApi {
	var msg RoomMsg
	s := NewRoomMsgService(msg)
	baseApi := api.NewApi[RoomMsg](s)
	return RoomMsgApi{Api: baseApi, Service: s}
}

func (r RoomMsgApi) GetRoomMsg(c *gin.Context) {
	msgs, err := r.Service.GetRoomMsg(c)
	if err != nil {
		response.FailWithError(err, c)
		return
	}

	response.OkWithData(msgs, c)
}

func (r RoomMsgApi) SendRoomMsg(c *gin.Context) {
	err := r.Service.SendRoomMsg(c)
	if err != nil {
		response.FailWithError(err, c)
		return
	}

	response.Ok(c)
}
