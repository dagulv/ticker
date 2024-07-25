package ticker

type tick struct {
	data    Tick
	symbols []string
	store   Store
}

func (m tick) processTick(t Tick) {
	m.Store()
}

func (m tick) Add(symbols ...string) {
	m.symbols = append(m.symbols, symbols...)
}

func (m tick) GetData() any {
	return m.data
}

func (m tick) GetSymbols() []string {
	return m.symbols
}

func (m tick) newMethod(s Store) method {
	return &tick{
		store: s,
	}
}

func (m tick) Store() {
	m.store.Push(nil, m)
}
