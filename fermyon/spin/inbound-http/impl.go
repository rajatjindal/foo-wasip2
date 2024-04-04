package inboundhttp

import (
	"bytes"
	"fmt"
	"net/http"
	"os"

	httptypes "github.com/rajatjindal/foo-wasip2/fermyon/spin/http-types"
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

func Handle(fn func(w http.ResponseWriter, r *http.Request)) {
	handler = fn
}

//export fermyon:spin/inbound-http@2.0.0#handle-request
func HandleExported(req *Request) *Response {
	var res = &Response{}
	var body []byte
	if req.Body.Some() != nil {
		// body = []byte(req.Body.Some())
	}

	var method string
	if req.Method == 0 {
		method = "GET"
	}

	innerRequest, err := http.NewRequest(method, string(req.URI), bytes.NewReader(body))
	if err != nil {
		res.Status = http.StatusInternalServerError
		return res
	}

	// for k, v := range req.Headers.(cm.List).Data() {

	// }

	w := &innerResponse{}

	handler(w, innerRequest)

	res.Status = httptypes.HTTPStatus(w.status)

	return res
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
