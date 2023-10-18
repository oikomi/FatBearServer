package dev

import "github.com/oikomi/FatBearServer/pkg/model"

type Dev struct {
	model.BaseModel
	DevName  string `json:"dev_name"  gorm:"column:dev_name;type:varchar(255);not null;index" `
	Cmd      string `json:"cmd"  gorm:"column:cmd;type:varchar(255);not null;index" `
	SendUser string `json:"send_user"  gorm:"column:send_user;type:varchar(255);not null;index" `
}

type DevInfo struct {
	model.BaseModel
	DevName   string `json:"dev_name"  gorm:"column:dev_name;type:varchar(255);not null;index" `
	Password  string `json:"password" gorm:"column:password;type:varchar(255);not null"`
	ModelName string `json:"model_name"  gorm:"column:model_name;type:varchar(255);not null;index" `
}

type Order struct {
	model.BaseModel
	DevName   string `json:"dev_name"  gorm:"column:dev_name;type:varchar(255);not null;index" `
	ModelName string `json:"model_name"  gorm:"column:model_name;type:varchar(255);not null;index" `
	SendUser  string `json:"send_user"  gorm:"column:send_user;type:varchar(255);not null;index" `
	Vibration string `json:"vibration"  gorm:"column:vibration;type:varchar(255);not null" `
	Duration  int    `json:"duration"  gorm:"column:duration;type:int;not null" `
	Token     int    `json:"token"  gorm:"column:token;type:int;not null" `
	Status    int    `json:"status"  gorm:"column:status;type:int;not null" `
}
