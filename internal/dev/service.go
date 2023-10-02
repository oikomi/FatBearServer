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
		DevName: req.DevName,
		Cmd:     req.Cmd,
	}
	mapper := model.NewMapper[Dev](dev, nil)
	err = mapper.Insert(&dev)
	if err != nil {
		config.GVA_LOG.Info("insert failed")
		return errors.Errorf("insert dev cmd failed: %s, %s", req.DevName, req.Cmd)
	}

	return nil
}

func (s DevService) GetCmd(c *gin.Context) (Dev, error) {
	var req GetCmdReq
	err := c.ShouldBind(&req)
	if err != nil {
		config.GVA_LOG.Error("get cmd failed", zap.Error(err))
		return Dev{}, err
	}

	dev := Dev{
		DevName: req.DevName,
	}
	mapper := model.NewMapper[Dev](dev, nil)
	cmd, err := mapper.SelectOne()
	if err != nil {
		return cmd, errors.Errorf("find dev failed: %s", req.DevName)
	}

	return cmd, nil
}
