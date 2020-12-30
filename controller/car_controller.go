package user

import (
	"fmt"
	car "gin_test/service"
	user "gin_test/service"

	"github.com/gin-gonic/gin"
)

type CarController struct{}

// Create action: POST /cars
func (pc Controller) AddCar(c *gin.Context) {
	var s car.Service
	p, err := s.CreateModel(c)

	if err != nil {
		c.AbortWithStatus(400)
		fmt.Println(err)
	} else {
		c.JSON(201, p)
	}
}

// Index action: GET /cars
func (pc Controller) ShowAllCar(c *gin.Context) {
	var s car.Service
	p, err := s.GetAll()

	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, p)
	}
}

// Show action: GET /cars/:limit
func (pc Controller) ShowLimitCar(c *gin.Context) {
	limit := c.Params.ByName("limit")
	var s car.Service
	p, err := s.GetSum(limit)

	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, p)
	}
}

// Show action: GET /cars/:id
func (pc Controller) ShowCar(c *gin.Context) {
	carId := c.Params.ByName("id")
	var s car.Service
	p, err := s.GetByID(carId)

	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, p)
	}
}

// Update action: PUT /cars/:id
func (pc Controller) UpdateCar(c *gin.Context) {
	carId := c.Params.ByName("id")
	var s user.Service
	p, err := s.UpdateByID(carId, c)

	if err != nil {
		c.AbortWithStatus(400)
		fmt.Println(err)
	} else {
		c.JSON(200, p)
	}
}

// Delete action: DELETE /cars/:id
func (pc Controller) DeleteCar(c *gin.Context) {
	carId := c.Params.ByName("id")
	var s user.Service

	if err := s.DeleteByID(carId); err != nil {
		c.AbortWithStatus(403)
		fmt.Println(err)
	} else {
		c.JSON(204, gin.H{"id #" + carId: "deleted"})
	}
}
