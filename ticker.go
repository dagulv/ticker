package ticker

import (
	"log"

	"github.com/rs/xid"
)

type Method interface {
	processTick(Store, Tick)
}

type Store interface {
	Push(Method)
}

func New[T Method, I any](s Store, ids []xid.ID, identifier []I) Ticker[T, I] {
	return Ticker[T, I]{
		ids:        ids,
		identifier: identifier,
		store:      s,
	}
}

type Ticker[T Method, I any] struct {
	data       T
	ids        []xid.ID
	identifier []I
	store      Store
}

func (t Ticker[T, I]) processTick(tick Tick, m Method) {
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
