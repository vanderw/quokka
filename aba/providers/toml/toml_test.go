package toml

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToml(t *testing.T) {
	blob := `
	contacts = [
		"Donald Duck <donald@duckburg.com>",
		"Scrooge McDuck <scrooge@duckburg.com>",
	]
`
	var contacts struct {
		Contacts []string `toml:"contacts"`
	}

	parser := New()
	err := parser.Unmarshal([]byte(blob), &contacts)
	assert.Nil(t, err)
	t.Log(contacts)

	res, err := parser.Marshal(&contacts)
	assert.Nil(t, err)
	t.Log(string(res))
}
