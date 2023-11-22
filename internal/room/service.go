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

func (s RoomService) checkRoomExist(name string) (bool, error) {
	w := model.NewWrapper()
	w.Eq("room_name", name)

	m := Room{}
	mapper := model.NewMapper[Room](m, w)
	findRoom, err := mapper.SelectOne()
	if err == nil && findRoom.ID > 0 {
		config.GVA_LOG.Info("Room exist : ", zap.String("room", name))
		return true, nil
	} else {
		return false, err
	}
}

func (s RoomService) Heartbeat(c *gin.Context) error {
	var req HeartbeatReq
	err := c.ShouldBind(&req)
	if err != nil {
		config.GVA_LOG.Error("Bind Heartbeat failed", zap.Error(err))
		return err
	}

	// isExist, err := checkRoomExist(req.Name)
	// if err != nil {

	// }

	return nil
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
		config.GVA_LOG.Error("Room already exist", zap.String("room", req.Name))
		return errors.Errorf("Room already exist: %s", req.Name)
	}
	room := Room{
		RoomName:   req.Name,
		Creator:    req.Creator,
		RoomStatus: int(CREATE),
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

func (s RoomService) GetRoomMsg(c *gin.Context) ([]RoomMsg, error) {
	var req GetRoomMsgReq
	err := c.ShouldBind(&req)
	if err != nil {
		config.GVA_LOG.Error("Bind GetRoomMsgReq failed", zap.Error(err))
		return nil, err
	}
	w := model.NewWrapper()
	if req.Name != "" {
		w.Eq("room_name", req.Name)
	}

	m := RoomMsg{}
	mapper := model.NewMapper[RoomMsg](m, w)
	msgs, err := mapper.Select()
	if err != nil {
		config.GVA_LOG.Error("select room failed", zap.String("room", req.Name))
		return nil, errors.Errorf("select room failed: %s", req.Name)
	}

	return *msgs, nil
}

func (s RoomService) SendRoomMsg(c *gin.Context) error {
	var req SendRoomMsgReq
	err := c.ShouldBind(&req)
	if err != nil {
		config.GVA_LOG.Error("Send msg failed", zap.Error(err))
		return err
	}

	msg := RoomMsg{
		RoomName: req.RoomName,
		SendUser: req.SendUser,
		Msg:      req.Msg,
	}
	mapper := model.NewMapper[RoomMsg](msg, nil)
	err = mapper.Insert(&msg)
	if err != nil {
		config.GVA_LOG.Info("insert failed")
		return errors.Errorf("insert msg failed: %s, %s, %s", req.RoomName, req.SendUser, req.Msg)
	}

	return nil
}

// room msg
type RoomMsgService struct {
	service.Service[RoomMsg]
}

func NewRoomMsgService(msg RoomMsg) RoomMsgService {
	return RoomMsgService{service.NewBaseService[RoomMsg](msg)}
}

func (s RoomMsgService) GetRoomMsg(c *gin.Context) ([]RoomMsg, error) {

	var req GetRoomMsgReq
	err := c.ShouldBind(&req)
	if err != nil {
		config.GVA_LOG.Error("Bind GetRoomMsgReq failed", zap.Error(err))
		return nil, err
	}
	w := model.NewWrapper()
	if req.Name != "" {
		w.Eq("room_name", req.Name)
	}

	m := RoomMsg{}
	mapper := model.NewMapper[RoomMsg](m, w)
	msgs, err := mapper.Select()
	if err != nil {
		config.GVA_LOG.Error("select room failed", zap.String("room", req.Name))
		return nil, errors.Errorf("select room failed: %s", req.Name)
	}

	return *msgs, nil
}

func (s RoomMsgService) SendRoomMsg(c *gin.Context) error {
	var req SendRoomMsgReq
	err := c.ShouldBind(&req)
	if err != nil {
		config.GVA_LOG.Error("Send msg failed", zap.Error(err))
		return err
	}

	msg := RoomMsg{
		RoomName: req.RoomName,
		SendUser: req.SendUser,
		Msg:      req.Msg,
	}
	mapper := model.NewMapper[RoomMsg](msg, nil)
	err = mapper.Insert(&msg)
	if err != nil {
		config.GVA_LOG.Info("insert failed")
		return errors.Errorf("insert msg failed: %s, %s, %s", req.RoomName, req.SendUser, req.Msg)
	}

	return nil
}
