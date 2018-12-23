package main

import (
	"fmt"
	"net/http"

	"github.com/rerost/unstable/uhttp"
)

func main() {
	http.Handle("/", uhttp.WithUnstable(handler))
	if err := http.ListenAndServe(":3000", nil); err != nil {
		fmt.Println(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Sample"))
}
