package dev

type Request struct {
	Name string `json:"name" form:"name"`
}

type SendCmdReq struct {
	DevName string `json:"dev_name" `
	Cmd     string `json:"cmd" `
}

type GetCmdReq struct {
	DevName string `json:"dev_name" `
}
