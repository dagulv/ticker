package ticker

import (
	"context"
	"encoding/base64"
	"log"
	"net/url"
	"time"

	tt "github.com/dagulv/ticker/ticker-template"
	"github.com/gorilla/websocket"
	"google.golang.org/protobuf/proto"
)

var yfEndpoint = url.URL{Scheme: "wss", Host: "streamer.finance.yahoo.com", Path: ""}

func (t *Ticker[T, I]) StartWebsocket(ctx context.Context) (err error) {
	log.Printf("connecting to %s...", yfEndpoint.String())

	done := make(chan struct{})

	var conn *websocket.Conn

	if conn, err = connect(yfEndpoint.String(), t.identifier); err != nil {
		return
	}

	defer conn.Close()

	go func() {
		defer close(done)
		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				log.Println(err)

				// Reconnecting in case of 1006 TODO
				conn.Close()
				if conn, err = connect(yfEndpoint.String(), t.identifier); err != nil {
					return
				}

				continue
			}

			b := make([]byte, base64.StdEncoding.DecodedLen(len(message)))
			n, err := base64.StdEncoding.Decode(b, message)

			if err != nil {
				return
			}

			decodedMessage := b[:n]

			tickTemplateData := &tt.Ticker{}
			if err = proto.Unmarshal(decodedMessage, tickTemplateData); err != nil {
				return
			}

			tickData := Tick{
				Time:      time.Unix(tickTemplateData.Time/1000, 0).UTC(),
				Symbol:    tickTemplateData.Id,
				Price:     tickTemplateData.Price,
				DayVolume: tickTemplateData.DayVolume,
			}

			t.processTick(tickData, t.data)
		}
	}()

	for {
		select {
		case <-done:
			return
		case <-ctx.Done():
			log.Println("interrupt")

			if err = conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "")); err != nil {
				return
			}

			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return
		}
	}
}

func connect[I any](endpoint string, subscribeSymbols []I) (conn *websocket.Conn, err error) {
	conn, _, err = websocket.DefaultDialer.Dial(endpoint, nil)

	if err != nil {
		return
	}

	subscribeMessage := SubscribeMessage[I]{
		Subscribe: subscribeSymbols,
	}

	if err = conn.WriteJSON(subscribeMessage); err != nil {
		log.Println(err)
		return
	}

	return
}
