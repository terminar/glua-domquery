package domquery

import (
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
	lua "github.com/yuin/gopher-lua"
)

func Ldocument_FromString(L *lua.LState) int {
	data := L.CheckString(1)
	buf := strings.NewReader(data)
	doc, err := goquery.NewDocumentFromReader(buf)

	if err != nil {
		L.Push(lua.LNil)
		L.Push(lua.LString(err.Error()))
		return 2
	}

	ud := L.NewUserData()
	ud.Value = doc.Selection
	L.SetMetatable(ud, L.GetTypeMetatable("selection"))

	L.Push(ud)
	L.Push(lua.LNil)
	return 2
}

func Ldocument_FromFile(L *lua.LState) int {
	fname := L.CheckString(1)
	buf, ferr := os.Open(fname)
	if ferr != nil {
		L.Push(lua.LNil)
		L.Push(lua.LString(ferr.Error()))
		return 2
	}
	doc, err := goquery.NewDocumentFromReader(buf)

	if err != nil {
		L.Push(lua.LNil)
		L.Push(lua.LString(err.Error()))
		return 2
	}

	ud := L.NewUserData()
	ud.Value = doc.Selection
	L.SetMetatable(ud, L.GetTypeMetatable("selection"))

	L.Push(ud)
	L.Push(lua.LNil)
	return 2
}
