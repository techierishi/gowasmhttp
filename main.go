package main

import (
	"context"
	"fmt"
	"syscall/js"
	"time"

	fetch "marwan.io/wasm-fetch"
)

func runLua(this js.Value, inputs []js.Value) interface{} {
	// lr := utl.LuaRunner{}
	// stdOut, _ := lr.RunLuaFunc(inputs[0].String())
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, _ := fetch.Fetch("https://reqres.in/api/users?page=2", &fetch.Opts{
		Method: fetch.MethodGet,
		Signal: ctx,
	})

	fmt.Println(string(res.Body))
	return res.Body
}

func main() {
	c := make(chan struct{}, 0)
	js.Global().Set("runlua", js.FuncOf(runLua))
	<-c
}
