package main

import (
	"bufio"
	"os"

	lua "github.com/yuin/gopher-lua"
	"github.com/yuin/gopher-lua/parse"
)

func CompileLua(filepath string) (*lua.FunctionProto, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	reader := bufio.NewReader(file)
	chunk, err := parse.Parse(reader, filepath)
	if err != nil {
		return nil, err
	}

	proto, err := lua.Compile(chunk, filepath)
	if err != nil {
		return nil, err
	}
	return proto, nil
}

func DoCompiledFile(L *lua.LState, proto *lua.FunctionProto) error {
	lfunc := L.NewFunctionFromProto(proto)
	L.Push(lfunc)
	return L.PCall(0, lua.MultRet, nil)
}

// share code between lua states
func exampleUsage() {
	codeToShare, _ := CompileLua("test.lua")
	a := lua.NewState()
	b := lua.NewState()
	c := lua.NewState()
	DoCompiledFile(a, codeToShare)
	DoCompiledFile(b, codeToShare)
	DoCompiledFile(c, codeToShare)
}
