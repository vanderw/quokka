package main

import (
	"fmt"

	lua "github.com/yuin/gopher-lua"
)

var exports = map[string]lua.LGFunction{
	"myfunc": myFunc,
}

func myFunc(L *lua.LState) int {
	fmt.Println("Im in myFunc()")
	return 0
}

func Loader(L *lua.LState) int {
	mod := L.SetFuncs(L.NewTable(), exports)
	L.SetField(mod, "name", lua.LString("value"))
	L.Push(mod)
	return 1
}
