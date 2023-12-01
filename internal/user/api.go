package user

import (
	"github.com/gin-gonic/gin"
	"github.com/oikomi/FatBearServer/pkg/api"
	"github.com/oikomi/FatBearServer/pkg/auth"
	"github.com/oikomi/FatBearServer/pkg/response"
)

type UserApi struct {
	api.Api
	Service UserService
}

func NewUserApi() UserApi {
	var u auth.BaseUser
	s := NewUserService(u)
	baseApi := api.NewApi[auth.BaseUser](s)
	return UserApi{Api: baseApi, Service: s}
}

// @Description	user login
// @Accept			json
// @Produce		json
// @Param			super_token	header	string		false	"Authentication header"
// @Param			LoginReq	body	LoginReq	true	"user login"
// @Router			/api/v1/user/login [post]
func (u UserApi) Login(c *gin.Context) {
	token, err := u.Service.Login(c)
	if err != nil {
		response.FailWithError(err, c)
		return
	}

	response.OkWithData(token, c)
}

// @Description	user add token
// @Accept			json
// @Produce		json
// @Param			super_token	header	string		false	"Authentication header"
// @Param			AddTokenReq	body	AddTokenReq	true	"AddToken"
// @Router			/api/v1/user/addToken [post]
func (u UserApi) AddToken(c *gin.Context) {
	err := u.Service.AddToken(c)
	if err != nil {
		response.FailWithError(err, c)
		return
	}

	response.Ok(c)
}

// @Description	user add token
// @Accept			json
// @Produce		json
// @Param			super_token	header		string		false	"Authentication header"
// @Param			name	query	string	true	"用户名称"
// @Success		200			{integer}	string		"token"
// @Router			/api/v1/user/getToken [get]
func (u UserApi) GetToken(c *gin.Context) {
	token, err := u.Service.GetToken(c)
	if err != nil {
		response.FailWithError(err, c)
		return
	}

	response.OkWithData(token, c)
}
