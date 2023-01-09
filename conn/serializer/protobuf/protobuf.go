package protobuf

import (
	"google.golang.org/protobuf/proto"
	"quokka.io/quokka"
)

type Serializer struct{}

func NewSerializer() *Serializer {
	return &Serializer{}
}

func (s *Serializer) Marshal(v interface{}) ([]byte, error) {
	pb, ok := v.(proto.Message)
	if !ok {
		return nil, quokka.ErrWrongValueType
	}
	return proto.Marshal(pb)
}

func (s *Serializer) Unmarshal(data []byte, v interface{}) error {
	pb, ok := v.(proto.Message)
	if !ok {
		return quokka.ErrWrongValueType
	}
	return proto.Unmarshal(data, pb)
}

func (s *Serializer) Name() string {
	return "protobuf"
}
