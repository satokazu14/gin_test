package server

import (
	"gin_test/controller"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Init() {
	r := router()
	r.Run(":3000")
}

func router() *gin.Engine {
	r := gin.Default()

	// ここからCorsの設定
	r.Use(cors.New(cors.Config{
		// アクセスを許可したいアクセス元
		AllowOrigins: []string{
			"*",
		},
		// アクセスを許可したいHTTPメソッド(以下の例だとPUTやDELETEはアクセスできません)
		AllowMethods: []string{
			"POST",
			"GET",
			"OPTIONS",
			"PUT",
			"DELETE",
		},
		// 許可したいHTTPリクエストヘッダ
		AllowHeaders: []string{
			"Access-Control-Allow-Credentials",
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"Authorization",
		},
		// cookieなどの情報を必要とするかどうか
		AllowCredentials: true,
		// preflightリクエストの結果をキャッシュする時間
		MaxAge: 24 * time.Hour,
	}))

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

	car := r.Group("/cars")
	{
		ctrl := controller.CarController{}
		car.GET("", ctrl.ShowAllCar)
		car.GET("/:id", ctrl.ShowCar)
		car.POST("", ctrl.AddCar)
		car.PUT("/:id", ctrl.UpdateCar)
		car.DELETE("/:id", ctrl.DeleteCar)
	}

	top := r.Group("/top")
	{
		ctrl := controller.AuctionController{}
		top.GET("", ctrl.Top)
		top.GET("/all", ctrl.TopAll)
	}

	bid := r.Group("/bids")
	{
		ctrl := controller.BidController{}
		bid.GET("", ctrl.Index)
		bid.GET("/auction/:id", ctrl.AuctionId)
		bid.GET("/user/:id", ctrl.UserId)
		bid.GET("/car/:id", ctrl.CarId)
		bid.POST("", ctrl.Create)
		bid.GET("/result/:id", ctrl.Result)
	}

	return r
}
