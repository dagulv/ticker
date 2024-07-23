package ticker

import "context"

type Ticker interface {
	Add(method, Store)
	Start(context.Context) error
}

type Store interface {
	Store(context.Context, method)
}

type method interface {
	processTick(Tick)
	Add(...string)
	GetData() any
	GetSymbols() []string
	newMethod(Store) method
	Store()
}

type ticker struct {
	methods []method
}

func New() Ticker {
	return &ticker{
		methods: make([]method, 0),
	}
}

func (t *ticker) Add(m method, s Store) {

	t.methods = append(t.methods, m.newMethod(s))
}
