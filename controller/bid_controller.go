package controller

import (
	"fmt"
	"gin_test/service"

	"github.com/gin-gonic/gin"
)

type BidController struct{}

func (bc BidController) Index(c *gin.Context) {
	var bid service.BidService
	p, err := bid.GetAll()

	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, p)
	}
}

func (bc BidController) AuctionId(c *gin.Context) {
	id := c.Params.ByName("id")
	var bid service.BidService
	p, err := bid.GetByAuctionID(id)

	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, p)
	}
}

func (bc BidController) Result(c *gin.Context) {
	id := c.Params.ByName("id")
	var bid service.BidService
	p, err := bid.GetResult(id)
	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, p)
	}
}

func (bc BidController) UserId(c *gin.Context) {
	id := c.Params.ByName("id")
	var bid service.BidService
	p, err := bid.GetByUserID(id)

	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, p)
	}
}

func (bc BidController) CarId(c *gin.Context) {
	id := c.Params.ByName("id")
	var bid service.BidService
	p, err := bid.GetByCarID(id)

	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, p)
	}
}

func (bc BidController) Create(c *gin.Context) {
	var bid service.BidService
	p, err := bid.CreateBid(c)

	if err != nil {
		c.AbortWithStatus(400)
		fmt.Println(err)
	} else {
		c.JSON(201, p)
	}
}
