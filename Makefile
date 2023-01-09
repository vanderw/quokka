
all: luatest tomltest hjsontest

luatest: cmd/luatest/*.go
	go build ./cmd/luatest/...

tomltest: cmd/tomltest/*.go
	go build ./cmd/tomltest/...

hjsontest: cmd/hjsontest/*.go
	go build ./cmd/hjsontest/...