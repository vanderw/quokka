package message

type Encoder interface {
	IsCompressionEnabled() bool
	IsEncryptionEnabled() bool
	Encode(message Message) ([]byte, error)
}
type MessageEncoder struct {
	CompressionEnabled bool
	CryptionEnabled    bool
}

func NewMessageEncoder(compressionEnabled, cryptionEnabled bool) *MessageEncoder {
	me := &MessageEncoder{
		CompressionEnabled: compressionEnabled,
		CryptionEnabled:    cryptionEnabled,
	}
	return me
}
