package room

import "github.com/oikomi/FatBearServer/pkg/model"

type Status int

const (
	CREATE Status = iota
	LIVING
	EXIT
)

type Room struct {
	model.BaseModel
	RoomName   string `json:"room_name"  gorm:"column:room_name;type:varchar(255);not null;index;unique" example:"房间名称" format:"string"`
	Creator    string `json:"creator"  gorm:"column:creator;type:varchar(255);not null"  example:"房间创建者，就是主播账号名称" format:"string"`
	RoomStatus int    `json:"room_status"  gorm:"column:room_status;type:int;not null"  example:"0" format:"int"`
}

type RoomMsg struct {
	model.BaseModel
	RoomName string `json:"room_name"  gorm:"column:room_name;type:varchar(255);not null" `
	SendUser string `json:"send_user"  gorm:"column:send_user;type:varchar(255);not null" `
	Msg      string `json:"msg"  gorm:"column:msg;type:varchar(255);not null" `
}
