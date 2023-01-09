package main

import lua "github.com/yuin/gopher-lua"

type Person struct {
	Name string
}

func newPerson(L *lua.LState) int {
	person := &Person{L.CheckString(1)}
	ud := L.NewUserData()
	ud.Value = person
	L.SetMetatable(ud, L.GetTypeMetatable(luaPersonTypeName))
	L.Push(ud)
	return 1
}

func checkPerson(L *lua.LState) *Person {
	ud := L.CheckUserData(1)
	if v, ok := ud.Value.(*Person); ok {
		return v
	}
	L.ArgError(1, "person expected")
	return nil
}

func personGetSetName(L *lua.LState) int {
	p := checkPerson(L)
	if L.GetTop() == 2 {
		p.Name = L.CheckString(2)
		return 0
	}
	L.Push(lua.LString(p.Name))
	return 1
}

var personMethods = map[string]lua.LGFunction{
	"name": personGetSetName,
}

const luaPersonTypeName = "person"

func registerPersonType(L *lua.LState) {
	mt := L.NewTypeMetatable(luaPersonTypeName)
	L.SetGlobal(luaPersonTypeName, mt)
	L.SetField(mt, "new", L.NewFunction(newPerson))
	L.SetField(mt, "__index", L.SetFuncs(L.NewTable(), personMethods))
}
