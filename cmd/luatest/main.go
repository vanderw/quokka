package main

import (
	"fmt"

	lua "github.com/yuin/gopher-lua"
)

// calling this from lua
func Double(L *lua.LState) int {
	lv := L.ToInt(1)
	lv2 := L.ToString(2)
	fmt.Printf("args:%v, %v\n", lv, lv2)
	L.Push(lua.LNumber(lv * 2))
	L.Push(lua.LString(lv2))
	return 2
}

func main() {
	L := lua.NewState()
	defer L.Close()

	L.PreloadModule("mymodule", Loader)

	L.SetGlobal("double", L.NewFunction(Double))

	registerPersonType(L)

	// execute lua string
	if err := L.DoString(`print("Hello")`); err != nil {
		panic(err)
	}

	// execute lua file
	if err := L.DoFile("main.lua"); err != nil {
		panic(err)
	}

	// calling lua from go
	if err := L.CallByParam(lua.P{
		Fn:      L.GetGlobal("l_double"),
		NRet:    1,
		Protect: true,
	}, lua.LNumber(10)); err != nil {
		panic(err)
	}
	ret := L.Get(-1) // returned value
	L.Pop(1)         // remove received value
	fmt.Println("returned from lua:", ret)
}
