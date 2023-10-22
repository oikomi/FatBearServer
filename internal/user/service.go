package user

import (
	"github.com/gin-gonic/gin"
	"github.com/oikomi/FatBearServer/config"
	"github.com/oikomi/FatBearServer/pkg/auth"
	"github.com/oikomi/FatBearServer/pkg/model"
	"github.com/oikomi/FatBearServer/pkg/service"
	"github.com/oikomi/FatBearServer/utils"
	"github.com/pkg/errors"
	"go.uber.org/zap"
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

func (s UserService) Login(c *gin.Context) (string, error) {
	// var req LoginReq
	// err := c.ShouldBind(&req)
	// if err != nil {
	// 	config.GVA_LOG.Error("Login failed", zap.Error(err))
	// 	return err
	// }

	user := auth.BaseUser{}

	token, err := user.Login(c)
	if err != nil {
		config.GVA_LOG.Info("login failed")
		return "", errors.Errorf("login failed")
	}

	return token, nil
}

func (s UserService) AddToken(c *gin.Context) error {
	var req AddTokenReq
	err := c.ShouldBind(&req)
	if err != nil {
		config.GVA_LOG.Error("Bind AddTokenReq failed", zap.Error(err))
		return err
	}

	config.GVA_LOG.Info("add token req , name is %s", zap.String("name", req.Name))

	user := auth.BaseUser{
		Name: req.Name,
	}

	w := model.NewWrapper()
	w.Eq("name", req.Name)

	mapper := model.NewMapper[auth.BaseUser](user, w)
	storeUser, err := mapper.SelectOne()
	if err != nil {
		config.GVA_LOG.Error("user not exist", zap.String("name", req.Name))
		return errors.Errorf("user not exist: %s", req.Name)
	}

	w2 := model.NewWrapper()
	w2.Eq("name", req.Name)

	mapper = model.NewMapper[auth.BaseUser](auth.BaseUser{
		Name: req.Name,
	}, w2)
	err = mapper.Update("token", storeUser.Token+req.Token)
	if err != nil {
		return errors.Errorf("update user token failed: %s", req.Name)
	}

	return nil
}

func (s UserService) GetToken(c *gin.Context) (int, error) {
	var req GetTokenReq
	err := c.ShouldBind(&req)
	if err != nil {
		config.GVA_LOG.Error("Bind GetTokenReq failed", zap.Error(err))
		return 0, err
	}

	config.GVA_LOG.Info("get token req , name is", zap.String("name", req.Name))

	user := auth.BaseUser{
		Name: req.Name,
	}

	w := model.NewWrapper()
	w.Eq("name", req.Name)

	mapper := model.NewMapper[auth.BaseUser](user, w)
	storeUser, err := mapper.SelectOne()
	if err != nil {
		config.GVA_LOG.Error("user not exist", zap.String("name", req.Name))
		return 0, errors.Errorf("user not exist: %s", req.Name)
	}

	return storeUser.Token, nil
}
