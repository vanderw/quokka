package main

import (
	"fmt"

	"github.com/hjson/hjson-go/v4"
)

func main() {

	type Sample struct {
		Rate  int
		Array []string
	}

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

	var dat map[string]interface{}
	if err := hjson.Unmarshal(sampleText, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)
	rate := dat["rate"].(float64)
	array := dat["array"].([]interface{})
	str1 := array[0].(string)

	fmt.Println(rate, str1)

	// Parse to struct
	var sample Sample
	hjson.Unmarshal(sampleText, &sample)
	fmt.Println(sample)

	var sa SampleAlias
	hjson.Unmarshal(sampleText, &sa)
	fmt.Println(sa)
}
