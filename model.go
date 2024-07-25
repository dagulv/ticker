package ticker

import "time"

type Company struct {
	Symbol string `json:"symbol"`
	Name   string `json:"name"`
}

type SubscribeMessage struct {
	Subscribe []string `json:"subscribe"`
}

type OHLCV struct {
	Symbol string
	Open   int
	High   int
	Low    int
	Close  int
	Volume int
	Time   time.Time
}

type Tick struct {
	Time      time.Time `json:"time"`
	Symbol    string    `json:"symbol"`
	Price     float32   `json:"price"`
	DayVolume int64     `json:"dayVolume"`
}
