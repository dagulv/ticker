package ticker

import (
	"context"
	"time"
)

type Ticker interface {
	ExposeTick(context.Context, Tick) error
}

type Tick struct {
	Time      time.Time `json:"time"`
	Symbol    string    `json:"symbol"`
	Price     float32   `json:"price"`
	DayVolume int64     `json:"dayVolume"`
}

type Company struct {
	Symbol string `json:"symbol"`
	Name   string `json:"name"`
}

type SubscribeMessage struct {
	Subscribe []string `json:"subscribe"`
}
