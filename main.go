package ticker

// import (
// 	"encoding/base64"
// 	"log"
// 	"net/url"
// 	"os"
// 	"os/signal"
// 	"time"

// 	tt "github.com/dagulv/ticker/ticker-template"
// 	"github.com/gorilla/websocket"
// 	"google.golang.org/protobuf/proto"
// )

// type SubscribeMessage struct {
// 	Subscribe []string `json:"subscribe"`
// }

// func main() {
// 	// ctx, cancel := context.WithCancel(context.Background())
// 	// defer cancel()

// 	interrupt := make(chan os.Signal, 1)
// 	signal.Notify(interrupt, os.Interrupt)

// 	if err := start(interrupt); err != nil {
// 		panic(err)
// 	}
// }

// func start(interrupt chan os.Signal) (err error) {
// 	subscribeSymbols := []string{"BTC-USD", "^GSPC", "^DJI", "^IXIC", "^RUT", "CL=F", "GC=F", "SWED-A.ST", "NFLX", "ETH-USD", "ADA-USD", "ACHR"}
// 	u := url.URL{Scheme: "wss", Host: "streamer.finance.yahoo.com", Path: ""}
// 	log.Printf("connecting to %s...", u.String())

// 	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
// 	if err != nil {
// 		return
// 	}
// 	defer c.Close()

// 	subscribeMessage := SubscribeMessage{
// 		Subscribe: subscribeSymbols,
// 	}

// 	if err = c.WriteJSON(subscribeMessage); err != nil {
// 		log.Println(err)
// 		return
// 	}

// 	done := make(chan struct{})

// 	go func() {
// 		defer close(done)
// 		for {
// 			_, message, err := c.ReadMessage()
// 			if err != nil {
// 				log.Println("read:", err)
// 				return
// 			}

// 			b := make([]byte, base64.StdEncoding.DecodedLen(len(message)))
// 			n, err := base64.StdEncoding.Decode(b, message)

// 			if err != nil {
// 				return
// 			}

// 			decodedMessage := b[:n]

// 			tickData := &tt.Ticker{}
// 			if err = proto.Unmarshal(decodedMessage, tickData); err != nil {
// 				return
// 			}

// 			log.Println(tickData)
// 		}
// 	}()

// 	for {
// 		select {
// 		case <-done:
// 			return
// 		case <-interrupt:
// 			log.Println("interrupt")

// 			if err = c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "")); err != nil {
// 				return
// 			}

// 			select {
// 			case <-done:
// 			case <-time.After(time.Second):
// 			}
// 			return
// 		}
// 	}
// }
