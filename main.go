package main

import (
	"net/http"

	incominghandler "github.com/rajatjindal/foo-wasip2/wasi/http/incoming-handler"
)

func init() {
	incominghandler.CustomHandle(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("from inside my handler"))
	})
}

func main() {}
