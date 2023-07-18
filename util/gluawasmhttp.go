package utl

import (
	"fmt"

	lua "github.com/yuin/gopher-lua"
	fetch "marwan.io/wasm-fetch"
)

type httpModule struct {
}

type empty struct{}

func NewHttpModule() *httpModule {
	return NewHttpModuleWithDo()
}

func NewHttpModuleWithDo() *httpModule {

	return &httpModule{}
}

func (h *httpModule) Loader(L *lua.LState) int {
	mod := L.SetFuncs(L.NewTable(), map[string]lua.LGFunction{
		"get":     h.get,
		"delete":  h.delete,
		"head":    h.head,
		"patch":   h.patch,
		"post":    h.post,
		"put":     h.put,
		"request": h.request,
	})
	registerHttpResponseType(mod, L)
	L.Push(mod)
	return 1
}

func (h *httpModule) get(L *lua.LState) int {
	return h.doRequestAndPush(L, "get", L.ToString(1), L.ToTable(2))
}

func (h *httpModule) delete(L *lua.LState) int {
	return h.doRequestAndPush(L, "delete", L.ToString(1), L.ToTable(2))
}

func (h *httpModule) head(L *lua.LState) int {
	return h.doRequestAndPush(L, "head", L.ToString(1), L.ToTable(2))
}

func (h *httpModule) patch(L *lua.LState) int {
	return h.doRequestAndPush(L, "patch", L.ToString(1), L.ToTable(2))
}

func (h *httpModule) post(L *lua.LState) int {
	return h.doRequestAndPush(L, "post", L.ToString(1), L.ToTable(2))
}

func (h *httpModule) put(L *lua.LState) int {
	return h.doRequestAndPush(L, "put", L.ToString(1), L.ToTable(2))
}

func (h *httpModule) request(L *lua.LState) int {
	return h.doRequestAndPush(L, L.ToString(1), L.ToString(2), L.ToTable(3))
}

func (h *httpModule) doRequest(L *lua.LState, method string, url string, options *lua.LTable) (*lua.LUserData, error) {

	ctx := L.Context()

	res, err := fetch.Fetch(url, &fetch.Opts{
		Method: method,
		Signal: ctx,
	})
	if err != nil {
		return nil, err
	}

	fmt.Println(string(res.Body))
	return newHttpResponse(res, nil, 0, L), nil
}

func (h *httpModule) doRequestAndPush(L *lua.LState, method string, url string, options *lua.LTable) int {
	response, err := h.doRequest(L, method, url, options)

	if err != nil {
		L.Push(lua.LNil)
		L.Push(lua.LString(fmt.Sprintf("%s", err)))
		return 2
	}

	L.Push(response)
	return 1
}

func toTable(v lua.LValue) *lua.LTable {
	if lv, ok := v.(*lua.LTable); ok {
		return lv
	}
	return nil
}
