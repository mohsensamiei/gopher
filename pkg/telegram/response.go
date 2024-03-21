package telegram

type Response[T any] struct {
	OK     bool `json:"ok"`
	Result T    `json:"result"`
}
