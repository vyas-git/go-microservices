package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, rq *http.Request) {
		d, err := ioutil.ReadAll(rq.Body)
		if err != nil {
			http.Error(rw, "Bad request", 400)
			return
		}
		fmt.Fprintf(rw, "Hello %s", d)
	})

	http.ListenAndServe(":9090", nil)
}
