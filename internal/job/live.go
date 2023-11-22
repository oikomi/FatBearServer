package job

import (
	"time"

	"github.com/oikomi/FatBearServer/config"
	"github.com/oikomi/FatBearServer/internal/room"
	"github.com/oikomi/FatBearServer/pkg/model"
	"github.com/panjf2000/ants/v2"
	"go.uber.org/zap"
)

func liveScan() {

	w := model.NewWrapper()
	// w.Eq("room_name", name)

	m := room.Room{}
	mapper := model.NewMapper[room.Room](m, w)
	findRooms, err := mapper.Select()
	if err != nil {
		config.GVA_LOG.Error("live scan failed", zap.Error(err))
	}

	if len(*findRooms) == 0 {
		return
	}

	now := time.Now()
	timeDur := time.Duration(60) * time.Second
	// nowAddTime := now.Add(timeDur)

	for _, v := range *findRooms {
		if now.After(v.UpdatedAt.Time().Add(timeDur)) {
			err = config.GVA_DB.Debug().Model(&room.Room{}).Where("room_name=?", v.RoomName).Update("room_status", 2).Error
			if err != nil {
				config.GVA_LOG.Error("update live room status failed", zap.Error(err))
			}
		}
	}

}

func LiveJob() {

	defer ants.Release()

	p, _ := ants.NewPoolWithFunc(10, func(i interface{}) {
		liveScan()
	})
	defer p.Release()
}
