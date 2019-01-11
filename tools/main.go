package main

import (
	"interview/tools/pprofuse"
	"net/http"
	_ "net/http/pprof"
)

func main() {

	go pprofuse.Use()

	http.ListenAndServe("0.0.0.0:6060", nil)

}
