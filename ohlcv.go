package ticker

import (
	"log"
	"time"
)

//TODO
// time: minuten den börjar räkna från
// open: första pris
// high: max
// low: min
// close: sista pris
// volume: sista dayVolume - första dayVolume

type Ohlcv struct {
	Symbol string
	Open   int
	High   int
	Low    int
	Close  int
	Volume int
	Time   time.Time
}

func (m Ohlcv) processTick(s Store, t Tick) {
	log.Println("ohlcv.go")
	s.Push(t)
}
