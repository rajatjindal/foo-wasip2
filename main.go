package main

import (
	"fmt"

	"github.com/rajatjindal/foo-wasip2/foo-namespace/pkg/foo"
)

func init() {
	// inboundhttp.HandleFn(func(req inboundhttp.Request) *inboundhttp.Response {
	// 	return &inboundhttp.Response{
	// 		Status: 201,
	// 	}
	// })

	// foo.Greet()
	// _ = inboundhttp.HandleRequest(httptypes.Request{})
	// opt := keyvalue.StoreOpen("default")
	// if opt.IsErr() {
	// 	panic(opt.Err())
	// }

	// store := opt.OK()
	// d := []byte("aa")
	// store.Set("hello", cm.ToList(d))
	// res := store.Get("hello")
	// if res.IsErr() {
	// 	panic(res.Err())
	// }

	// fmt.Println(string(res.OK().Some().Slice()))
}

func main() {
	f := &foo.Impl{}
	fmt.Println(f.Greet())
}
