package message

import (
	"errors"
	"fmt"
	"strings"
	"sync"
)

type Type byte

const (
	MsgRequest  Type = 0x00
	MsgNotify   Type = 0x01
	MsgResponse Type = 0x02
	MsgPush     Type = 0x03
)

// map type_code -> type_string
var types = map[Type]string{
	MsgRequest:  "Request",
	MsgNotify:   "Notify",
	MsgResponse: "Response",
	MsgPush:     "Push",
}

func (t Type) String() string {
	return types[t]
}

func (t Type) routable() bool {
	return t == MsgRequest || t == MsgNotify || t == MsgPush
}

func (t Type) invalid() bool {
	return t < MsgRequest || t > MsgPush
}

var (
	routesCodesMutex = sync.RWMutex{}
	routes           = make(map[string]uint16)
	codes            = make(map[uint16]string)
)

var (
	ErrWrongMessageType  = errors.New("wrong message type")
	ErrInvalidMessage    = errors.New("invalid message")
	ErrRouteInfoNotFound = errors.New("route info not found in dictionary")
)

type Message struct {
	Type       Type
	Id         uint
	Route      string
	Data       []byte
	compressed bool
	encrypted  bool
	Err        bool
}

func New(err ...bool) *Message {
	m := &Message{}
	if len(err) > 0 {
		m.Err = err[0]
	}
	return m
}

func (m Message) String() string {
	return fmt.Sprintf("Type: %s, Id: %d, Route: %s, Compressed: %t, Encrypted: %t, Error: %t, Data: %v, BodyLength: %d",
		m.Type.String(), m.Id, m.Route, m.compressed, m.encrypted, m.Err, m.Data, len(m.Data))
}

func SetDictionary(dict map[string]uint16) error {
	if dict == nil {
		return nil
	}
	routesCodesMutex.Lock()
	defer routesCodesMutex.Unlock()

	for route, code := range dict {
		r := strings.TrimSpace(route)

		if _, ok := routes[r]; ok {
			return fmt.Errorf("duplicated route(route: %s, code: %d)", r, code)
		}

		if _, ok := codes[code]; ok {
			return fmt.Errorf("duplicated route(route: %s, code: %d)", r, code)
		}

		routes[r] = code
		codes[code] = r
	} // for

	return nil
}

func GetDictionary() map[string]uint16 {
	routesCodesMutex.RLock()
	defer routesCodesMutex.RUnlock()

	dict := make(map[string]uint16)
	for k, v := range routes {
		dict[k] = v
	}
	return dict
}
