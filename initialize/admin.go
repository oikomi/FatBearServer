package initialize

import (
	"github.com/gin-gonic/gin"
	"github.com/oikomi/FatBearServer/config"
	"github.com/oikomi/FatBearServer/internal/app"
	"github.com/oikomi/FatBearServer/pkg/admin"
)

func Admin(r *gin.Engine) {
	if !config.GVA_CONFIG.Admin.Enable {
		return
	}
	admin.Init(r, nil)
	app.Admin()
}
