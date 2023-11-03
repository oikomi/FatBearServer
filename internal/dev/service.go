package dev

import (
	"github.com/gin-gonic/gin"
	"github.com/oikomi/FatBearServer/config"
	"github.com/oikomi/FatBearServer/pkg/model"
	"github.com/oikomi/FatBearServer/pkg/service"
	"github.com/oikomi/FatBearServer/utils"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type DevService struct {
	service.Service[Dev]
}

func NewDevService(r Dev) DevService {
	return DevService{service.NewBaseService[Dev](r)}
}

func (s DevService) MakeMapper(c *gin.Context) model.Mapper[Dev] {
	var r Request
	err := c.ShouldBindQuery(&r)
	utils.CheckError(err)
	w := model.NewWrapper()
	w.Like("name", r.Name)
	m := model.NewMapper[Dev](Dev{}, w)
	return m
}

func (s DevService) MakeResponse(val model.Model) any {
	d := val.(Dev)
	res := Response{
		DevName: d.DevName,
	}
	return res
}

func (s DevService) SendCmd(c *gin.Context) error {
	var req SendCmdReq
	err := c.ShouldBind(&req)
	if err != nil {
		config.GVA_LOG.Error("Semd cmd failed", zap.Error(err))
		return err
	}

	dev := Dev{
		DevName:  req.DevName,
		Cmd:      req.Cmd,
		SendUser: req.SendUser,
	}
	mapper := model.NewMapper[Dev](dev, nil)
	err = mapper.Insert(&dev)
	if err != nil {
		config.GVA_LOG.Info("insert failed")
		return errors.Errorf("insert dev cmd failed: %s, %s", req.DevName, req.Cmd)
	}

	return nil
}

func (s DevService) GetCmd(c *gin.Context) (*Dev, error) {
	var req GetCmdReq
	err := c.ShouldBind(&req)
	if err != nil {
		config.GVA_LOG.Error("get cmd failed", zap.Error(err))
		return nil, err
	}

	config.GVA_LOG.Info("get cmd", zap.Object("req", req))

	w := model.NewWrapper()
	w.Eq("dev_name", req.DevName)

	mapper := model.NewMapper[Dev](Dev{}, w)
	cmds, err := mapper.Select()
	if err != nil {
		return nil, errors.Errorf("find dev failed: %s", req.DevName)
	}
	if len(*cmds) >= 1 {
		return &((*cmds)[0]), nil
	}

	return nil, nil
}

func (s DevService) Login(c *gin.Context) error {
	var req DevLoginReq

	err := c.ShouldBind(&req)
	if err != nil {
		config.GVA_LOG.Error("dev login failed", zap.Error(err))
		return err
	}

	w := model.NewWrapper()
	w.Eq("dev_name", req.DevName)

	mapper := model.NewMapper[DevInfo](DevInfo{}, w)
	dev, err := mapper.SelectOne()
	if err != nil {
		return errors.Errorf("find dev failed: %s", req.DevName)
	}
	if dev.ModelName != req.ModelName {
		config.GVA_LOG.Info("model name not the same")
		mapper := model.NewMapper[DevInfo](DevInfo{DevName: req.DevName}, w)
		err = mapper.Update("model_name", req.ModelName)
		if err != nil {
			return errors.Errorf("update dev failed: %s", req.DevName)
		}
	}

	return nil
}

func (s DevService) OrderList(c *gin.Context) ([]Order, error) {
	var req OrderListReq
	err := c.ShouldBind(&req)
	if err != nil {
		config.GVA_LOG.Error("get order failed", zap.Error(err))
		return nil, err
	}

	w := model.NewWrapper()
	if req.SendUser != "" {
		w.Eq("model_name", "host").EqF("status", 0)
	}

	mapper := model.NewMapper[Order](Order{}, w)
	orders, err := mapper.Select()
	if err != nil {
		config.GVA_LOG.Error("select order failed")
		return nil, errors.Errorf("select order failed: %s", err)
	}

	// w = model.NewWrapper()
	// w.Eq("model_name", req.SendUser)
	// w.EqF("status", 0)
	
	err = config.GVA_DB.Debug().Model(&Order{}).Where("model_name=?", req.SendUser).Where("status=?", 0).Update("status", 1).Error
	if err != nil {
		return nil, errors.Errorf("update order status failed: %s", req.SendUser)
	}

	return *orders, nil
}

func (s DevService) Order(c *gin.Context) error {
	var req OrderReq
	err := c.ShouldBind(&req)
	if err != nil {
		config.GVA_LOG.Error("do order failed", zap.Error(err))
		return err
	}

	order := Order{
		DevName:   req.DevName,
		ModelName: req.ModelName,
		SendUser:  req.SendUser,
		Vibration: req.Vibration,
		Duration:  req.Duration,
		Token:     req.Token,
	}
	mapper := model.NewMapper[Order](order, nil)
	err = mapper.Insert(&order)
	if err != nil {
		config.GVA_LOG.Error("insert failed")
		return errors.Errorf("insert order failed: %s", req.DevName)
	}

	return nil
}


// order
type DevOrderService struct {
	service.Service[Order]
}

func NewDevOrderService(r Order) DevOrderService {
	return DevOrderService{service.NewBaseService[Order](r)}
}

func (s DevOrderService) Order(c *gin.Context) error {
	var req OrderReq
	err := c.ShouldBind(&req)
	if err != nil {
		config.GVA_LOG.Error("do order failed", zap.Error(err))
		return err
	}

	order := Order{
		DevName:   req.DevName,
		ModelName: req.ModelName,
		SendUser:  req.SendUser,
		Vibration: req.Vibration,
		Duration:  req.Duration,
		Token:     req.Token,
	}
	mapper := model.NewMapper[Order](order, nil)
	err = mapper.Insert(&order)
	if err != nil {
		config.GVA_LOG.Error("insert failed")
		return errors.Errorf("insert order failed: %s", req.DevName)
	}

	return nil
}

func (s DevOrderService) OrderList(c *gin.Context) ([]Order, error) {
	var req OrderListReq
	err := c.ShouldBind(&req)
	if err != nil {
		config.GVA_LOG.Error("get order failed", zap.Error(err))
		return nil, err
	}

	w := model.NewWrapper()
	if req.SendUser != "" {
		w.Eq("model_name", "host").EqF("status", 0)
	}

	mapper := model.NewMapper[Order](Order{}, w)
	orders, err := mapper.Select()
	if err != nil {
		config.GVA_LOG.Error("select order failed")
		return nil, errors.Errorf("select order failed: %s", err)
	}

	// w = model.NewWrapper()
	// w.Eq("model_name", req.SendUser)
	// w.EqF("status", 0)
	
	// err = config.GVA_DB.Debug().Model(&Order{}).Where("model_name=?", req.SendUser).Where("status=?", 0).Update("status", 1).Error
	// if err != nil {
	// 	return nil, errors.Errorf("update order status failed: %s", req.SendUser)
	// }

	return *orders, nil
}

func (s DevService) Set(c *gin.Context) ([]DevSetting, error) {
	var req SetReq
	err := c.ShouldBind(&req)
	if err != nil {
		config.GVA_LOG.Error("do set failed", zap.Error(err))
		return nil, err
	}

	w := model.NewWrapper()
	if req.ModelName != "" {
		w.Eq("model_name", req.ModelName)
	}

	set := DevSetting{
		ModelName: req.ModelName,
	}
	mapper := model.NewMapper[DevSetting](set, w)
	sets, err := mapper.Select()
	if err != nil {
		config.GVA_LOG.Error("select dev set failed")
		return nil, errors.Errorf("select dev set failed: %s", err)
	}

	return *sets, nil
}
