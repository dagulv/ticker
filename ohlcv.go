package ticker

import (
	"log"
	"time"

	"github.com/rs/xid"
)

//TODO
// time: minuten den börjar räkna från
// open: första pris
// high: max
// low: min
// close: sista pris
// volume: sista dayVolume - första dayVolume

type Ohlcv struct {
	Id     xid.ID
	Open   float64
	High   float64
	Low    float64
	Close  float64
	Volume int
	Time   time.Time
}

type AvanzaHistory struct {
	Ohlcv []AvanzaOhlcv `json:"ohlc"`
}

type AvanzaOhlcv struct {
	Timestamp         int     `json:"timestamp"`
	Open              float64 `json:"open"`
	Close             float64 `json:"close"`
	Low               float64 `json:"low"`
	High              float64 `json:"high"`
	TotalVolumeTraded int     `json:"totalVolumeTraded"`
}

func (m Ohlcv) processTick(s Store, t Tick) {
	log.Println("ohlcv.go")
	s.Push(t)
}
