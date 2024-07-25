package ticker

type ohlcv struct {
	data    OHLCV
	symbols []string
	store   Store
}

func (m ohlcv) processTick(t Tick) {
	m.Store()
}

func (m ohlcv) Add(symbols ...string) {
	m.symbols = append(m.symbols, symbols...)
}

func (m ohlcv) GetData() any {
	return m.data
}

func (m ohlcv) GetSymbols() []string {
	return m.symbols
}

func (m ohlcv) newMethod(s Store) method {
	return &tick{
		store: s,
	}
}

func (m ohlcv) Store() {
	m.store.Push(nil, m)
}
