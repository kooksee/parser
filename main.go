package main

import (
	"github.com/kooksee/parser/internal"
	"os"
)

func main() {
	var port = "8559"
	var debug = true

	if e := os.Getenv("port"); e != "" {
		port = e
	}

	if e := os.Getenv("debug"); e != "" {
		debug = e == "true"
	}

	if err := internal.Start(debug, port); err != nil {
		panic(err.Error())
	}
}
