package toml

import (
	"bytes"

	bstoml "github.com/BurntSushi/toml"
)

type Toml struct{}

func New() *Toml {
	return &Toml{}
}

func (t Toml) Unmarshal(data []byte, v any) error {
	_, err := bstoml.Decode(string(data), v) // _ present `metadata`, i ignore it here
	return err
}

func (t Toml) Marshal(v any) ([]byte, error) {
	var buf bytes.Buffer
	err := bstoml.NewEncoder(&buf).Encode(v)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
