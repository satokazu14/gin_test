package controller

import (
	"fmt"
	"gin_test/config"
	"gin_test/entity"
	"gin_test/service"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
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

var rh = map[string]string{
	"0": "無",
	"1": "有",
}

var rc = map[string]string{
	"0": "込",
	"1": "別",
}

var lm = map[string]string{
	"0": "無",
	"1": "付",
}

var wr = map[string]string{
	"0": "無",
	"1": "付",
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
	var strerr error
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
			a[0].Cars[i].RepairHistory = rh[a[0].Cars[i].RepairHistory]
			a[0].Cars[i].RecyclingConsignment = rc[a[0].Cars[i].RecyclingConsignment]
			a[0].Cars[i].LeagalMaintenance = lm[a[0].Cars[i].LeagalMaintenance]
			a[0].Cars[i].Warranty = wr[a[0].Cars[i].Warranty]
			a[0].Cars[i].StartTime = ac[i].StartTime
			a[0].Cars[i].StartPrice, strerr = strconv.ParseFloat(strconv.FormatFloat(a[0].Cars[i].PurchasePrice*1.1*0.0001, 'f', 1, 64), 64)
			if strerr != nil {
				c.AbortWithStatus(404)
				fmt.Println(err)
				return
			}
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

var rooms = make(map[string]*config.Hub)
var upgrader = websocket.Upgrader{}

func (pc AuctionController) AuctionRoom(c *gin.Context) {
	chat := c.Param("id")

	// channelが存在する場合
	_, ok := rooms[chat]
	if !ok {
		rooms[chat] = config.NewHub()
	}
	room := rooms[chat]

	go func() {
		for {
			// メッセージ受け取り
			bid := <-room.Broadcast
			fmt.Println(bid)
			// クライアントの数だけループ
			for client := range room.Clients {
				// 書き込む
				err := client.WriteJSON(bid)
				if err != nil {
					log.Printf("error occurred while writing message to client: %v", err)
					client.Close()
					delete(room.Clients, client)
				}
			}
		}
	}()
	// websocket の状態を更新
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	websocket, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Fatal("error upgrading request to a websocket::", err)
	}
	// websocket を閉じる
	defer websocket.Close()

	room.Clients[websocket] = true

	for {
		var bid entity.Bid
		var bidservice service.BidService
		// メッセージ読み込み
		err := websocket.ReadJSON(&bid)
		fmt.Println(bid)
		if err != nil {
			log.Printf("error occurred while reading message: %v", err)
			delete(room.Clients, websocket)
			break
		}
		bidservice.AuctionBid(bid)
		// メッセージを受け取る
		room.Broadcast <- bid
	}
}
