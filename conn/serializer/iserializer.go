package serializer

type (
	Marshaler interface {
		Marshal(interface{}) ([]byte, error)
	}

	Unmarshaler interface {
		Unmarshal([]byte, interface{}) error
	}

	ISerializer interface {
		Marshaler
		Unmarshaler
		Name() string
	}
)
