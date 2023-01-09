package hjson

import "github.com/hjson/hjson-go/v4"

type Hjson struct{}

func New() *Hjson {
	return &Hjson{}
}

func (hj Hjson) Unmarshal(data []byte, v any) error {
	return hjson.Unmarshal(data, v)
}

func (hj Hjson) Marshal(v any) ([]byte, error) {
	return hjson.Marshal(v)
}
