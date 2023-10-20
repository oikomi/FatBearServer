package room

import "github.com/oikomi/FatBearServer/pkg/model"

type Status int

const (
	created Status = iota
)

type Room struct {
	model.BaseModel
	RoomName   string `json:"room_name"  gorm:"column:room_name;type:varchar(255);not null;index;unique" `
	Creator    string `json:"creator"  gorm:"column:creator;type:varchar(255);not null" `
	RoomStatus int    `json:"room_status"  gorm:"column:room_status;type:int;not null" `
}

type RoomMsg struct {
	model.BaseModel
	RoomName string `json:"room_name"  gorm:"column:room_name;type:varchar(255);not null" `
	SendUser string `json:"send_user"  gorm:"column:send_user;type:varchar(255);not null" `
	Msg      string `json:"msg"  gorm:"column:msg;type:varchar(255);not null" `
}
