package room

type Request struct {
	Name string `json:"name" form:"name"`
}

type CreateRoomReq struct {
	Name    string `json:"name" `
	Creator string `json:"creator" `
}

type GetRoomListReq struct {
	Name    string `json:"name" `
}

