package room

type Request struct {
	Name string `json:"name" form:"name"`
}

type CreateRoomReq struct {
	Name    string `json:"name" example:"房间名称" format:"string"`
	Creator string `json:"creator" example:"房间创建者，就是主播账号名称" format:"string"`
}

type updateRoomStatusReq struct {
	Name   string `json:"name" example:"房间名称" format:"string"`
	Status int    `json:"status" example:"1" format:"int"`
}

type HeartbeatReq struct {
	Name string `json:"name" `
}

type GetRoomListReq struct {
	Name string `json:"name" `
}

type GetRoomMsgReq struct {
	Name string `form:"name" `
}

type SendRoomMsgReq struct {
	RoomName string `json:"name" `
	SendUser string `json:"send_user" `
	Msg      string `json:"msg" `
}
