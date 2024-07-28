package ticker

import (
	"log"
	"time"
)

type Tick struct {
	Time      time.Time `json:"time"`
	Symbol    string    `json:"symbol"`
	Price     float32   `json:"price"`
	DayVolume int64     `json:"dayVolume"`
}

func (m Tick) processTick(s Store, t Tick) {
	log.Println("tick.go")
	s.Push(t)
}
