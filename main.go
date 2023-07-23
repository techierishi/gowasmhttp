package main

import (
	"syscall/js"

	utl "github.com/techierishi/wasmlua/util"
)

func runLua(this js.Value, inputs []js.Value) interface{} {
	lr := utl.LuaRunner{}
	lr.RunLuaFunc(inputs[0].String())

	return []interface{}{1, "two"}
}

func main() {
	c := make(chan struct{}, 0)
	js.Global().Set("runlua", js.FuncOf(runLua))
	<-c
}
