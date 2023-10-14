package user

import (
	"github.com/gin-gonic/gin"
	"github.com/oikomi/FatBearServer/config"
	"github.com/oikomi/FatBearServer/pkg/auth"
	"github.com/oikomi/FatBearServer/pkg/model"
	"github.com/oikomi/FatBearServer/pkg/service"
	"github.com/oikomi/FatBearServer/utils"
	"github.com/pkg/errors"
)

type UserService struct {
	service.Service[auth.BaseUser]
}

func NewUserService(u auth.BaseUser) UserService {
	return UserService{service.NewBaseService[auth.BaseUser](u)}
}

func (s UserService) MakeMapper(c *gin.Context) model.Mapper[auth.BaseUser] {
	var r Request
	err := c.ShouldBindQuery(&r)
	utils.CheckError(err)
	w := model.NewWrapper()
	w.Like("name", r.Name)
	m := model.NewMapper[auth.BaseUser](auth.BaseUser{}, w)
	return m
}

func (s UserService) MakeResponse(val model.Model) any {
	u := val.(auth.BaseUser)
	res := Response{
		UserName: u.Name,
	}
	return res
}

func (s UserService) Login(c *gin.Context) error {
	// var req LoginReq
	// err := c.ShouldBind(&req)
	// if err != nil {
	// 	config.GVA_LOG.Error("Login failed", zap.Error(err))
	// 	return err
	// }

	user := auth.BaseUser{}

	err := user.Login(c)
	if err != nil {
		config.GVA_LOG.Info("login failed")
		return errors.Errorf("login failed")
	}

	return nil
}
