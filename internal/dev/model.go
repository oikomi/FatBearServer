package dev

import "github.com/oikomi/FatBearServer/pkg/model"

type Dev struct {
	model.BaseModel
	DevName  string `json:"dev_name"  gorm:"column:dev_name;type:varchar(255);not null;index" `
	Cmd      string `json:"cmd"  gorm:"column:cmd;type:varchar(255);not null;index" `
	SendUser string `json:"send_user"  gorm:"column:send_user;type:varchar(255);not null;index" `
}
