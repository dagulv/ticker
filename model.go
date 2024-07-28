package ticker

type Company struct {
	Symbol string `json:"symbol"`
	Name   string `json:"name"`
}

type SubscribeMessage struct {
	Subscribe []string `json:"subscribe"`
}
