package room

import "github.com/oikomi/FatBearServer/pkg/model"

type Room struct {
	model.BaseModel
	RoomName   string `json:"room_name"  gorm:"column:room_name;type:varchar(255);not null" `
	RoomStatus string `json:"room_status"  gorm:"column:room_status;type:int;not null" `
}
