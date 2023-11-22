package room

type Request struct {
	Name string `json:"name" form:"name"`
}

type CreateRoomReq struct {
	Name    string `json:"name" `
	Creator string `json:"creator" `
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
