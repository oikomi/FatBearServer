package app

import (
	"github.com/oikomi/FatBearServer/pkg/admin"
)

func Admin() {
	var a App
	admin.New(a, "app", "应用")
}
