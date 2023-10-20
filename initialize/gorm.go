package initialize

import (
	"os"
	"os/user"

	"github.com/oikomi/FatBearServer/config"
	"github.com/oikomi/FatBearServer/internal/app"
	"github.com/oikomi/FatBearServer/internal/dev"
	"github.com/oikomi/FatBearServer/internal/room"
	"github.com/oikomi/FatBearServer/pkg/auth"
	"gorm.io/gorm"
)

func Gorm() *gorm.DB {
	switch config.GVA_CONFIG.DbType {
	case "mysql":
		return GormMysql()
	default:
		return GormMysql()
	}
}

// RegisterTables 注册数据库表专用
func RegisterTables(db *gorm.DB) {
	err := db.Set("gorm:table_options", "CHARSET=utf8mb4").AutoMigrate(
		// 系统模块表
		auth.BaseUser{},
		app.App{}, // app表注册
		room.Room{},
		room.RoomMsg{},
		dev.Dev{},
		dev.DevInfo{},
		user.User{},
		dev.Order{},
		dev.DevSetting{},
	)
	if err != nil {
		os.Exit(0)
	}
}
