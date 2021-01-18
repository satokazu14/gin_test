package controller

import (
	"fmt"
	"gin_test/service"

	"github.com/gin-gonic/gin"
)

type SfbidController struct{}

func (sc SfbidController) ShowAll(c *gin.Context) {
	var s service.SfbidService

	p, err := s.GetAll()

	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, p)
	}
}

func (sc SfbidController) State(c *gin.Context) {
	id := c.Params.ByName("id")
	state := c.Params.ByName("state")
	var s service.SfbidService

	s.UpdateState(id, state)
}
