package controller

import (
	"fmt"
	"gin_test/service"

	"github.com/gin-gonic/gin"
)

type AuctionController struct{}

var bt = map[string]string{
	"sd": "セダン",
	"cp": "クーペ",
	"op": "オープンカー",
	"sw": "ステーションワゴン",
	"mv": "ミニバン",
	"su": "SUV・クロカン",
	"pt": "ピックアップ",
	"hb": "ハッチバック",
	"cc": "コンパクトカー",
	"kj": "軽自動車",
	"hv": "ハイブリット",
	"cm": "キャンピングカー",
	"fs": "福祉車両",
	"vn": "商用車・バン",
	"tr": "トラック",
	"at": "その他",
}

var nr = map[string]string{
	"gr": "レギュラー",
	"gh": "ハイオク",
	"di": "軽油",
	"mt": "メタノール",
	"cn": "天然ガス",
	"lp": "プロパンガス",
}

var tv = map[string]string{
	"0": "無",
	"1": "ワンセグ",
	"2": "フルセグ",
}

var ad = map[string]string{
	"0": "無",
	"1": "CD",
	"2": "CDチェンジャー",
}

var vi = map[string]string{
	"0": "無",
	"1": "DVD",
	"2": "BD",
}

var hl = map[string]string{
	"hg": "ハロゲン",
	"hi": "HID",
	"le": "LED",
	"ls": "レーザー",
}

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
		for i := 0; i < len(ac); i++ {
			a[0].Cars[i].BodyType = bt[a[0].Cars[i].BodyType]
			a[0].Cars[i].GasOil = nr[a[0].Cars[i].GasOil]
			a[0].Cars[i].Tv = tv[a[0].Cars[i].Tv]
			a[0].Cars[i].Audio = ad[a[0].Cars[i].Audio]
			a[0].Cars[i].Visual = vi[a[0].Cars[i].Visual]
			a[0].Cars[i].HeadLight = hl[a[0].Cars[i].HeadLight]
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
