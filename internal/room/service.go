package room

import (
	"github.com/gin-gonic/gin"
	"github.com/oikomi/FatBearServer/config"
	"github.com/oikomi/FatBearServer/pkg/model"
	"github.com/oikomi/FatBearServer/pkg/service"
	"github.com/oikomi/FatBearServer/utils"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type RoomService struct {
	service.Service[Room]
}

func NewRoomService(r Room) RoomService {
	return RoomService{service.NewBaseService[Room](r)}
}

func (s RoomService) MakeMapper(c *gin.Context) model.Mapper[Room] {
	var r Request
	err := c.ShouldBindQuery(&r)
	utils.CheckError(err)
	w := model.NewWrapper()
	w.Like("name", r.Name)
	m := model.NewMapper[Room](Room{}, w)
	return m
}

func (s RoomService) MakeResponse(val model.Model) any {
	r := val.(Room)
	res := Response{
		Name: r.RoomName,
	}
	return res
}

func (s RoomService) CreateRoom(c *gin.Context) error {
	var req CreateRoomReq
	err := c.ShouldBind(&req)
	if err != nil {
		config.GVA_LOG.Error("Bind CreateRoomReq failed", zap.Error(err))
		return err
	}

	m := Room{RoomName: req.Name}
	mapper := model.NewMapper[Room](m, nil)
	_, err = mapper.SelectOne()
	if err == nil {
		config.GVA_LOG.Error("Room exist", zap.String("room", req.Name))
		return errors.Errorf("Room exist: %s", req.Name)
	}
	room := Room{
		RoomName:   req.Name,
		Creator:    req.Creator,
		RoomStatus: int(created),
	}
	err = mapper.Insert(&room)
	if err != nil {
		config.GVA_LOG.Info("insert failed")
		return errors.Errorf("insert room failed: %s", req.Name)
	}

	return nil
}

func (s RoomService) Hello() string {
	return "Hello World"
}
