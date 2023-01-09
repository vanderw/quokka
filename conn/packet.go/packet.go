package packet

import (
	"errors"
	"fmt"
)

type Type byte

const (
	_                    Type = iota
	PktHandshake              = 0x01
	PktHandshakeAck           = 0x02
	PktHandshakeHeatbeat      = 0x03
	PktData                   = 0x04
	PktKick                   = 0x05
	PktMax                    = 0x06
)

var (
	ErrWrongPacketType     = errors.New("wrong packet type")
	ErrInvalidPacketHeader = errors.New("invalid packet header")
)

type Packet struct {
	Type   Type
	Length int
	Data   []byte
}

func New() *Packet {
	return &Packet{}
}

func (p *Packet) String() string {
	return fmt.Sprintf("Type: %d, Length: %d, Data: %s",
		p.Type, p.Length, string(p.Data))
}
