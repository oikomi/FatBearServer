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

	w := model.NewWrapper()
	w.Eq("room_name", req.Name)

	m := Room{}
	mapper := model.NewMapper[Room](m, w)
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
		config.GVA_LOG.Error("insert failed")
		return errors.Errorf("insert room failed: %s", req.Name)
	}

	return nil
}

func (s RoomService) GetRoomList(c *gin.Context) ([]Room, error) {

	var req GetRoomListReq
	err := c.ShouldBind(&req)
	if err != nil {
		config.GVA_LOG.Error("Bind GetRoomListReq failed", zap.Error(err))
		return nil, err
	}
	w := model.NewWrapper()
	if req.Name != "" {
		w.Eq("room_name", req.Name)
	}

	m := Room{}
	mapper := model.NewMapper[Room](m, w)
	rooms, err := mapper.Select()
	if err != nil {
		config.GVA_LOG.Error("select room failed", zap.String("room", req.Name))
		return nil, errors.Errorf("select room failed: %s", req.Name)
	}

	return *rooms, nil
}

func (s RoomService) UpdateRoom(c *gin.Context) error {

	return nil
}
