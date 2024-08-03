package ticker

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/robfig/cron/v3"
)

var (
	client = &http.Client{
		Timeout: 10 * time.Second,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				MinVersion: tls.VersionTLS13,
			},
		},
	}
	avanzaEndpoint = "https://www.avanza.se/_api/price-chart/stock"
	timePeriod     = "today"
)

func (t Ticker[T, I]) HistoricJob(ctx context.Context) {
	c := cron.New()
	job := func() {
		for i := range t.identifier {
			t.fetchData(i)
		}
	}
	c.AddFunc("0 23 * * 1-5", job)

	c.Start()
}

func (t Ticker[T, I]) fetchData(index int) {
	var data AvanzaHistory

	url := avanzaEndpoint + "/" + strconv.Itoa(any(t.identifier[index]).(int)) + "?timePeriod=" + timePeriod

	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		return
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "ticker")

	resp, err := client.Do(req)

	if err != nil {
		return
	}

	defer resp.Body.Close()

	if err = json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return
	}

	for _, single := range data.Ohlcv {
		t.store.Push(Ohlcv{
			Id:     t.ids[index],
			Open:   single.Open,
			High:   single.High,
			Low:    single.Low,
			Close:  single.Close,
			Volume: single.TotalVolumeTraded,
			Time:   time.Unix(int64(single.Timestamp)/1000, 0).UTC(),
		})
	}
}
