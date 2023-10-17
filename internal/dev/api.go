package dev

import (
	"github.com/gin-gonic/gin"
	"github.com/oikomi/FatBearServer/pkg/api"
	"github.com/oikomi/FatBearServer/pkg/response"
)

type DevApi struct {
	api.Api
	Service DevService
}

func NewDevApi() DevApi {
	var d Dev
	s := NewDevService(d)
	baseApi := api.NewApi[Dev](s)
	return DevApi{Api: baseApi, Service: s}
}

func (r DevApi) SendCmd(c *gin.Context) {
	err := r.Service.SendCmd(c)
	if err != nil {
		response.FailWithError(err, c)
		return
	}

	response.Ok(c)
}

func (r DevApi) GetCmd(c *gin.Context) {
	resp, err := r.Service.GetCmd(c)
	if err != nil {
		response.FailWithError(err, c)
		return
	}

	response.OkWithData(resp, c)
}


func (r DevApi) Login(c *gin.Context) {
	err := r.Service.Login(c)
	if err != nil {
		response.FailWithError(err, c)
		return
	}

	response.Ok(c)
}

