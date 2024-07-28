package ticker

import (
	"log"
)

type Method interface {
	processTick(Store, Tick)
}

type Store interface {
	Push(Method)
}

func New(m Method, s Store, symbols ...string) Ticker {
	return Ticker{
		data:    m,
		symbols: symbols,
		store:   s,
	}
}

type Ticker struct {
	data    Method
	symbols []string
	store   Store
}

func (t Ticker) processTick(tick Tick, m Method) {
	log.Println("ticker.go")

	m.processTick(t.store, tick)
}

// package ticker

// import "log"

// type MethodType interface {
// 	Tick | Ohlcv
// }

// type Method interface {
// 	MethodType
// 	processTick(Tick)
// }

// type Store[T MethodType] interface {
// 	Push(T)
// }

// func New[T Method](s Store[T], symbols ...string) Ticker[T] {
// 	return Ticker[T]{
// 		symbols: symbols,
// 		store:   s,
// 	}
// }

// type Ticker[T Method] struct {
// 	data    T
// 	symbols []string
// 	store   Store[T]
// }

// func (t Ticker[T]) processTick(tick Tick, m T) {
// 	log.Println("ticker.go")

// 	m.processTick()
// 	// t.store.Push(m.processTick)
// }
