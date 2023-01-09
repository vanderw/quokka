package aba

type Aba interface {
	Unmarshal(data []byte, v any) error
	// Sometimes we need to translate a struct to a string
	// but later writing the string back into a configure file is not recommended.
	// Because this will result in losting all the comments information.
	Marshal(v any) ([]byte, error)
}
