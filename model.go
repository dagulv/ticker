package ticker

type Company struct {
	Symbol string `json:"symbol"`
	Name   string `json:"name"`
}

type SubscribeMessage[I any] struct {
	Subscribe []I `json:"subscribe"`
}
