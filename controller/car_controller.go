package controller

import (
	"fmt"
	"gin_test/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CarController struct{}

// Create action: POST /cars
func (pc CarController) AddCar(c *gin.Context) {
	var s service.CarService
	p, err := s.CreateModel(c)

	if err != nil {
		c.AbortWithStatus(400)
		fmt.Println(err)
	} else {
		c.JSON(201, p)
	}
}

// Index action: GET /cars
func (pc CarController) ShowAllCar(c *gin.Context) {
	var s service.CarService

	p, err := s.GetAll()

	if limit := c.Query("limit"); limit != "" {
		p, err = s.GetSum(limit)
	}

	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, p)
	}
}

// Show action: GET /cars/:limit
func (pc CarController) ShowLimitCar(c *gin.Context) {
	limit := c.Params.ByName("limit")
	var s service.CarService
	p, err := s.GetSum(limit)

	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, p)
	}
}

// Show action: GET /cars/:id
func (pc CarController) ShowCar(c *gin.Context) {
	carId := c.Params.ByName("id")
	var s service.CarService
	var as service.AuctionService
	var strerr error
	p, err := s.GetByID(carId)

	q, _ := as.GetByCarID(strconv.FormatUint(uint64(p.ID), 10))

	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		p.BodyType = bt[p.BodyType]
		p.GasOil = nr[p.GasOil]
		p.Tv = tv[p.Tv]
		p.Audio = ad[p.Audio]
		p.Visual = vi[p.Visual]
		p.HeadLight = hl[p.HeadLight]
		p.RepairHistory = rh[p.RepairHistory]
		p.RecyclingConsignment = rc[p.RecyclingConsignment]
		p.LeagalMaintenance = lm[p.LeagalMaintenance]
		p.Warranty = wr[p.Warranty]
		p.StartTime = q.StartTime
		p.StartPrice, strerr = strconv.ParseFloat(strconv.FormatFloat(p.PurchasePrice*1.1*0.0001, 'f', 1, 64), 64)
		if strerr != nil {
			c.AbortWithStatus(404)
			fmt.Println(err)
			return
		}
		c.JSON(200, p)
	}
}

// Show action: GET /cars/:id
func (pc CarController) ShowCarD(c *gin.Context) {
	carId := c.Params.ByName("id")
	var s service.CarService
	p, err := s.GetByID(carId)

	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, p)
	}
}

// Update action: PUT /cars/:id
func (pc CarController) UpdateCar(c *gin.Context) {
	carId := c.Params.ByName("id")
	var s service.CarService
	p, err := s.UpdateByID(carId, c)

	if err != nil {
		c.AbortWithStatus(400)
		fmt.Println(err)
	} else {
		c.JSON(200, p)
	}
}

// Delete action: DELETE /cars/:id
func (pc CarController) DeleteCar(c *gin.Context) {
	carId := c.Params.ByName("id")
	var s service.CarService

	if err := s.DeleteByID(carId); err != nil {
		c.AbortWithStatus(403)
		fmt.Println(err)
	} else {
		c.JSON(204, gin.H{"id #" + carId: "deleted"})
	}
}
