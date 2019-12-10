package domquery

import (
	lua "github.com/yuin/gopher-lua"
)

//var hnode *goquery.Selection
//hnode.html

var VERSION string = "1.0.1"
var LNAME string = "domquery"

var API_Lselection = map[string]lua.LGFunction{
	// General
	"last":  Lselection_Last,
	"first": Lselection_First,
	"slice": Lselection_Slice,
	//"get":           Lselection_Get,
	"index":         Lselection_Index,
	"indexSelector": Lselection_IndexSelector,
	"html":          Lselection_Html,
	"children":      Lselection_Children,
	"parent":        Lselection_Parent,
	"parents":       Lselection_Parents,
	"each":          Lselection_Each,
	"eachBreak":     Lselection_EachBreak,

	// Filter
	"filter":       Lselection_Filter,
	"find":         Lselection_Find,
	"notEq":        Lselection_Not,
	"has":          Lselection_Has,
	"hasClass":     Lselection_HasClass,
	"is":           Lselection_Is,
	"hasSelection": Lselection_Filter,
	//"hasNodes":     Lselection_Filter,
	"closest": Lselection_Closest,

	// Manipulation
	"attr":            Lselection_Attr,
	"attrOr":          Lselection_AttrOr,
	"addClass":        Lselection_AddClass,
	"append":          Lselection_Append,
	"appendSelection": Lselection_AppendSelection,
	"appendHtml":      Lselection_AppendHtml,
	//"appendNodes":          Lselection_AppendNodes,
	"after":          Lselection_After,
	"afterSelection": Lselection_AfterSelection,
	"afterHtml":      Lselection_AfterHtml,
	//"afterNodes":           Lselection_AfterNodes,
	"before":          Lselection_Before,
	"beforeSelection": Lselection_BeforeSelection,
	"beforeHtml":      Lselection_BeforeHtml,
	//"beforeNodes":          Lselection_BeforeNodes,
	"clone":            Lselection_Clone,
	"empty":            Lselection_Empty,
	"prepend":          Lselection_Prepend,
	"prependSelection": Lselection_PrependSelection,
	"prependHtml":      Lselection_PrependHtml,
	//"prependNodes":         Lselection_PrependNodes,
	"remove":               Lselection_Remove,
	"removeFiltered":       Lselection_RemoveFiltered,
	"replaceWith":          Lselection_ReplaceWith,
	"replaceWithSelection": Lselection_ReplaceWithSelection,
	"replaceWithHtml":      Lselection_ReplaceWithHtml,
	//"replaceWithNodes":     Lselection_ReplaceWithNodes,
	"setAttr":     Lselection_SetAttr,
	"setHtml":     Lselection_SetHtml,
	"setText":     Lselection_SetText,
	"toggleClass": Lselection_ToggleClass,
	"text":        Lselection_Text,
	//"unwrap":               Lselection_Unwrap,
	//"wrap":                 Lselection_Wrap,
	//"wrapSelection":        Lselection_WrapSelection,
	//"wrapHtml":             Lselection_WrapHtml,
	//"wrapNode":             Lselection_WrapNode,
	//"WrapAll":              Lselection_WrapAll,
}

var API_Ldocumment = map[string]lua.LGFunction{
	"fromFile":   Ldocument_FromFile,
	"fromString": Ldocument_FromString,
}

var API_Lnode = map[string]lua.LGFunction{
	/*"attr":   Lnode_AttrValue,
	"data":   Lnode_AttrData,
	"parent": Lnode_Parent,
	"value":  Lnode_Value,
	//"nextSibling":  Lnode_NextSibling,
	//"prevSibling":  Lnode_PrevSibling,
	"firstChild":   Lnode_firstChild,
	"lastChild":    Lnode_LastChild,
	"appendChild":  Lnode_AppendChild,
	"insertBefore": Lnode_InsertBefore,*/
}

func Preload(L *lua.LState) {
	L.PreloadModule(LNAME, Loader)
}

// Loader is the module loader function.
func Loader(L *lua.LState) int {
	t := L.NewTable()

	luaSelection := L.NewTypeMetatable("selection")
	L.SetField(luaSelection, "__index", L.SetFuncs(L.NewTable(), API_Lselection))
	t.RawSetH(lua.LString("selection"), luaSelection)

	t.RawSetH(lua.LString("__version__"), lua.LString(VERSION))

	L.SetFuncs(t, API_Ldocumment)
	L.Push(t)

	return 1
}
