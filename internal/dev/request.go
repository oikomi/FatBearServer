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
	DevName   string `json:"dev_name"  example:"设备名称" format:"string"`
	Password  string `json:"password" example:"设备密码" format:"string"`
	ModelName string `json:"model_name" example:"设备主播账号名称" format:"string"`
}

type DevLogoutReq struct {
	DevName   string `json:"dev_name"  example:"设备名称" format:"string"`
	ModelName string `json:"model_name" example:"设备主播账号名称" format:"string"`
}


type OrderReq struct {
	DevName   string `json:"dev_name" form:"dev_name" example:"设备名称" format:"string"`
	ModelName string `json:"model_name" form:"model_name" example:"主播账号名称" format:"string"`
	SendUser  string `json:"send_user" form:"send_user" example:"观众名称" format:"string"`
	Vibration string `json:"vibration" form:"vibration" example:"震动强度, 比如 Medium" format:"string"`
	Duration  int    `json:"duration" form:"duration" example:"10" format:"int"`
	Token     int    `json:"token" form:"token" example:"50" format:"int"`
}

type OrderListReq struct {
	SendUser string `form:"send_user" json:"send_user"`
}

type GetUserByDevReq struct {
	DevName string `form:"dev_name" json:"dev_name"`
}

type SetReq struct {
	ModelName string `json:"model_name" `
	Vibration string `json:"vibration" `
	Duration  int    `json:"duration" `
	Token     int    `json:"token" `
}

type AddSetReq struct {
	ModelName string `json:"model_name" `
	Vibration string `json:"vibration" `
	Duration  int    `json:"duration" `
	Token     int    `json:"token" `
}

type DelSetReq struct {
	Id int64 `json:"id" `
}

// MarshalLogObject implements zapcore.ObjectMarshaler.
func (o GetCmdReq) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddString("dev_name", o.DevName)
	return nil
}
