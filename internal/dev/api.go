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

// @Description	DevLogin
// @Accept			json
// @Produce		json
// @Param			super_token	header	string		false	"Authentication header"
// @Param			DevLoginReq	body	DevLoginReq	true	"dev login"
// @Router			/api/v1/dev/login [post]
func (r DevApi) Login(c *gin.Context) {
	err := r.Service.Login(c)
	if err != nil {
		response.FailWithError(err, c)
		return
	}

	response.Ok(c)
}

// @Description	OrderList
// @Accept			json
// @Produce		json
// @Param			super_token	header	string	false	"Authentication header"
// @Param			send_user	query	string	true	"model 名字"
// @Success		200			{array}	Order
// @Router			/api/v1/dev/order [get]
func (r DevApi) OrderList(c *gin.Context) {
	orders, err := r.Service.OrderList(c)
	if err != nil {
		response.FailWithError(err, c)
		return
	}

	response.OkWithData(orders, c)
}

// @Description	send Order
// @Accept			json
// @Produce		json
// @Param			super_token	header	string		false	"Authentication header"
// @Param			OrderReq	body	OrderReq	true	"send order"
// @Router			/api/v1/dev/order [post]
func (r DevApi) Order(c *gin.Context) {
	err := r.Service.Order(c)
	if err != nil {
		response.FailWithError(err, c)
		return
	}

	response.Ok(c)
}

// @Description	getUserByDev
// @Accept			json
// @Produce		json
// @Param			super_token	header	string	false	"Authentication header"
// @Param			dev_name	query	string	true	"dev名字"
// @Success		200			{object}	auth.BaseUser	"user"
// @Router			/api/v1/dev/getUserByDev [get]
func (r DevApi) getUserByDev(c *gin.Context) {
	user, err := r.Service.getUserByDev(c)
	if err != nil {
		response.FailWithError(err, c)
		return
	}

	response.OkWithData(user, c)
}

// order api

type DevOrderApi struct {
	api.Api
	Service DevOrderService
}

func NewDevOrderApi() DevOrderApi {
	var o Order
	s := NewDevOrderService(o)
	baseApi := api.NewApi[Order](s)
	return DevOrderApi{Api: baseApi, Service: s}
}

func (r DevOrderApi) Order(c *gin.Context) {
	err := r.Service.Order(c)
	if err != nil {
		response.FailWithError(err, c)
		return
	}

	response.Ok(c)
}

func (r DevOrderApi) OrderList(c *gin.Context) {
	orders, err := r.Service.OrderList(c)
	if err != nil {
		response.FailWithError(err, c)
		return
	}

	response.OkWithData(orders, c)
}

// @Description	set
// @Accept			json
// @Produce		json
// @Param			super_token	header	string	false	"Authentication header"
// @Param			model_name	query	string	true	"model 名字"
// @Success		200			{array}	DevSetting
// @Router			/api/v1/dev/set [get]
func (r DevApi) Set(c *gin.Context) {
	sets, err := r.Service.Set(c)
	if err != nil {
		response.FailWithError(err, c)
		return
	}

	response.OkWithData(sets, c)
}

// @Description	add set
// @Accept			json
// @Produce		json
// @Param			super_token	header	string		false	"Authentication header"
// @Param			AddSetReq	body	AddSetReq	true	"add set"
// @Router			/api/v1/dev/set [post]
func (r DevApi) AddSet(c *gin.Context) {
	err := r.Service.AddSet(c)
	if err != nil {
		response.FailWithError(err, c)
		return
	}

	response.Ok(c)
}

// @Description	del set
// @Accept			json
// @Produce		json
// @Param			super_token	header	string		false	"Authentication header"
// @Param			DelSetReq	body	DelSetReq	true	"del set"
// @Router			/api/v1/dev/delSet [post]
func (r DevApi) DelSet(c *gin.Context) {
	err := r.Service.DelSet(c)
	if err != nil {
		response.FailWithError(err, c)
		return
	}

	response.Ok(c)
}
