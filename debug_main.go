//+build !release

package main

import (
	"fmt"
	"net/http"
	_ "orderfood/docs"
)

var isReleaseMode bool

func initServer() {
	isReleaseMode = false
}

func run(s *http.Server) {
	fmt.Printf("Listen: %s\n", s.Addr)
	if err := s.ListenAndServe(); err != nil {
		panic(err)
	}
}
