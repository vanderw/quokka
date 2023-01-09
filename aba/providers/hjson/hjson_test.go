package hjson

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHjson(t *testing.T) {
	type SampleAlias struct {
		Rett    int      `json:"rate"`
		Ashtray []string `json:"array"`
	}

	sampleText := []byte(`
	{
        # specify rate in requests/second
        rate: 1000
        array:
        [
            foo
            bar
        ]
    }`)

	parser := New()
	sa := SampleAlias{}
	err := parser.Unmarshal(sampleText, &sa)
	assert.Nil(t, err)
	t.Log(sa)

	res, err := parser.Marshal(&sa)
	assert.Nil(t, err)
	t.Log(string(res))
}
