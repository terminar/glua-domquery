package domquery

import (
	//"github.com/PuerkitoBio/goquery"
	lua "github.com/yuin/gopher-lua"
)

func Lnode_Attr(L *lua.LState) int {
	/*sel := check_Lnode(L, 1)
	qry := L.CheckString(2)
	L.Push(Lselection_New(L, sel.Filter(qry)))*/
	return 1
}
