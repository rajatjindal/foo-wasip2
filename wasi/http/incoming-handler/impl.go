package incominghandler

import (
	"fmt"
	"net/http"
	"os"

	"github.com/rajatjindal/foo-wasip2/wasi/http/types"
)

// func init() {
// 	instance = &Impl{}
// }

// type Impl struct{}

// func (i *Impl) HandleRequest(req Request) Response {
// 	instance.
// }

// //export foo-namespace:pkg/foo@2.0.0#greet
// func GreetExported() *string {
// 	x := instance.Greet()
// 	return &x
// }

var handler = defaultHandler
var defaultHandler = func(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(os.Stderr, "http handler undefined")
}

func CustomHandle(fn func(w http.ResponseWriter, r *http.Request)) {
	handler = fn
}

//export wasi:http/incoming-handler@0.2.0#handle
func HandleExported(req *IncomingRequest, res2 *ResponseOutparam) {
	var res = types.NewOutgoingResponse(types.NewFields())
	// var body []byte
	// if req.Body.Some() != nil {
	// 	// body = []byte(req.Body.Some())
	// }

	var method string
	if req.Method() == types.MethodGet() {
		method = "GET"
	}

	pwq := req.PathWithQuery()
	pwqp := &pwq
	s := pwqp.Some()
	fmt.Println("url is ", *s)
	innerRequest, err := http.NewRequest(method, *s, nil)
	if err != nil {
		res.SetStatusCode(http.StatusInternalServerError)
		// return res
		return
	}

	// for k, v := range req.Headers.(cm.List).Data() {

	// }

	w := &innerResponse{}

	handler(w, innerRequest)

	res.SetStatusCode(types.StatusCode(w.status))

	// result := cm.OK[types.OutgoingResponse](res)
	// types.ResponseOutparamSet(cm.ResourceNone, cm.ErrResult[types.OutgoingResponse, types.ErrorCode]{})
}

type innerResponse struct {
	status int
	body   []byte
	header http.Header
}

func (i *innerResponse) Header() http.Header {
	i.header = http.Header{}
	return i.header
}

func (i *innerResponse) Write(b []byte) (int, error) {
	i.body = b
	return 0, nil
}

func (i *innerResponse) WriteHeader(statusCode int) {
	i.status = statusCode
}
