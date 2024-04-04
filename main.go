package main

import (
	"net/http"

	inboundhttp "github.com/rajatjindal/foo-wasip2/fermyon/spin/inbound-http"
)

func init() {
	inboundhttp.Handle(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("from inside my handler"))
	})
}

func main() {}
