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
