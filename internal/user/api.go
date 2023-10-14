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

func (u UserApi) Login(c *gin.Context) {
	err := u.Service.Login(c)
	if err != nil {
		response.FailWithError(err, c)
		return
	}

	response.Ok(c)
}