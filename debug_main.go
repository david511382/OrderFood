//+build !release

package main

import _ "orderfood/docs"

var isReleaseMode bool

func initServer() {
	isReleaseMode = false
}

func getAddr() string {
	return cfg.Domain()
}

func close() {

}
