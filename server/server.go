package server

import (
	"gin_test/controller"

	"github.com/gin-gonic/gin"
)

func Init() {
	r := router()
	r.Run(":3000")
}

func router() *gin.Engine {
	r := gin.Default()

	u := r.Group("/users")
	{
		ctrl := controller.UserController{}
		u.GET("", ctrl.Index)
		u.GET("/:id", ctrl.Show)
		u.POST("", ctrl.Create)
		u.PUT("/:id", ctrl.Update)
		u.DELETE("/:id", ctrl.Delete)
	}

	au := r.Group("/auctions")
	{
		ctrl := controller.AuctionController{}
		au.GET("", ctrl.Index)
		au.GET("/:id", ctrl.Show)
		au.POST("", ctrl.Create)
		au.PUT("/:id", ctrl.Update)
		au.DELETE("/:id", ctrl.Delete)
	}

	return r
}
