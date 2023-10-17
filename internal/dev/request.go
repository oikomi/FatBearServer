package dev

import "go.uber.org/zap/zapcore"

type Request struct {
	Name string `json:"name" form:"name"`
}

type SendCmdReq struct {
	DevName  string `json:"dev_name" `
	Cmd      string `json:"cmd" `
	SendUser string `json:"send_user" `
}

type GetCmdReq struct {
	DevName string `json:"dev_name" form:"dev_name"`
}

type DevLoginReq struct {
	DevName  string `json:"dev_name" `
	Password      string `json:"password" `
}


// MarshalLogObject implements zapcore.ObjectMarshaler.
func (o GetCmdReq) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddString("dev_name", o.DevName)
	return nil
}
