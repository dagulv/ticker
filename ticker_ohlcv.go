package ticker

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/rs/xid"
)

var (
	client         = &http.Client{Timeout: 10 * time.Second}
	avanzaEndpoint = "https://www.avanza.se/_api/price-chart/stock"
	timePeriod     = "today"
)

func (t Ticker[T, I]) FetchHistoric(ctx context.Context) {
	ticker := time.NewTicker(time.Second)
	var doneForToday bool
	done := make(chan bool)
	go func() {
		defer close(done)
		for {
			select {
			case <-ctx.Done():
				ticker.Stop()
				return
			case <-ticker.C:
				h, _, _ := time.Now().Clock()
				if h > 19 || h < 7 {
					if doneForToday {
						continue
					}

					for _, id := range t.identifier {
						t.fetchData(id)
						time.Sleep(time.Second)
					}
					doneForToday = true
				} else if doneForToday {
					doneForToday = false
				}
			}
		}
	}()
}

func (t Ticker[T, I]) fetchData(id I) {
	var data AvanzaHistory

	resp, err := client.Get(fmt.Sprintf("%s/%d?timePeriod=%s", avanzaEndpoint, id, timePeriod))

	if err != nil {
		return
	}

	defer resp.Body.Close()

	if err = json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return
	}
	log.Println(data.Ohlcv)
	for _, single := range data.Ohlcv {
		t.store.Push(Ohlcv{
			Id:     any(id).(xid.ID),
			Open:   single.Open,
			High:   single.High,
			Low:    single.Low,
			Close:  single.Close,
			Volume: single.TotalVolumeTraded,
			Time:   time.Unix(int64(single.Timestamp)/1000, 0).UTC(),
		})
	}
}
