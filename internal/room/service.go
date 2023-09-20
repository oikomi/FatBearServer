package room

import (
	"github.com/gin-gonic/gin"
	"github.com/oikomi/FatBearServer/pkg/model"
	"github.com/oikomi/FatBearServer/pkg/service"
	"github.com/oikomi/FatBearServer/utils"
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

func (s RoomService) CreateRoom() string {
	

	return "Hello World"
}

func (s RoomService) Hello() string {
	return "Hello World"
}
