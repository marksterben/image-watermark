package domain

type SingleResponse[T any] struct {
	Data T `json:"data"`
}

type MultipleResponse[T any] struct {
	Data []T `json:"data"`
}
