package main

import (
	"flag"
	"fmt"
	"github.com/ofux/docilemonkey/docilemonkey"
	"log"
	"net/http"
)

func main() {
	var port = *flag.Int("port", 8080, "server port")

	http.HandleFunc("/", docilemonkey.Handler)

	log.Print(`====================================================
Docile Monkey server will respond to any relative URL.

Query params:
	s: will be the HTTP status code of the response (ex: 200, 404, 500, etc.)
	t: time to wait before responding (ex: 10s, 200ms, etc.)
	b: will be the body of the response
	bb: if parameter 'b' has no value and if bb=1 then the body of the request (if any) will be sent back in the response

Example of request:
	GET http://localhost:8080/foo/bar?s=201&t=500ms&b={"hello":"world"}

`)
	log.Printf("Listening on port %s...\n", port)
	log.Println(`====================================================`)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		log.Fatal(err)
	}
}
