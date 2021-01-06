package controller

import (
	"fmt"
	"gin_test/service"

	"github.com/gin-gonic/gin"
)

type AuctionController struct{}

func (pc AuctionController) Index(c *gin.Context) {
	var auction service.AuctionService
	p, err := auction.GetAll()

	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, p)
	}
}

func (pc AuctionController) Top(c *gin.Context) {
	var auction service.AuctionService
	a, ac, err := auction.GetTopInfo()

	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		fmt.Println(ac[0])
		for i := 0; i < len(ac); i++ {
			a[0].Cars[i].StartTime = ac[i].StartTime
		}
		c.JSON(200, a)
	}
}

func (pc AuctionController) TopAll(c *gin.Context) {
	var auction service.AuctionService
	a, err := auction.GetTopAllInfo()

	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, a)
	}
}

func (pc AuctionController) Create(c *gin.Context) {
	var auction service.AuctionService
	p, err := auction.CreateModel(c)

	if err != nil {
		c.AbortWithStatus(400)
		fmt.Println(err)
	} else {
		c.JSON(201, p)
	}
}

func (pc AuctionController) Show(c *gin.Context) {
	id := c.Params.ByName("id")
	var auction service.AuctionService
	p, err := auction.GetByID(id)

	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, p)
	}
}

func (pc AuctionController) Update(c *gin.Context) {
	id := c.Params.ByName("id")
	var auction service.AuctionService
	p, err := auction.UpdateByID(id, c)

	if err != nil {
		c.AbortWithStatus(400)
		fmt.Println(err)
	} else {
		c.JSON(200, p)
	}
}

// Delete action: DELETE /auctions/:id
func (pc AuctionController) Delete(c *gin.Context) {
	id := c.Params.ByName("id")
	var auction service.AuctionService

	if err := auction.DeleteByID(id); err != nil {
		c.AbortWithStatus(403)
		fmt.Println(err)
	} else {
		c.JSON(204, gin.H{"id #" + id: "deleted"})
	}
}