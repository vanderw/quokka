package main

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

func main() {
	blob := `
		contacts = [
			"Donald Duck <donald@duckburg.com>",
			"Scrooge McDuck <scrooge@duckburg.com>",
		]
	`
	var contacts struct {
		Contacts []string
	}

	meta, err := toml.Decode(blob, &contacts)
	fmt.Println(err)
	fmt.Println(meta)
}
