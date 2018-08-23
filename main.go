package main

import (
	"flag"

	"github.com/lon9/ww/client"
	"github.com/lon9/ww/server"
)

func main() {
	var (
		mode string
	)
	flag.StringVar(&mode, "m", "c", "Mode")
	flag.Parse()
	if mode == "c" {
		client := client.NewClient()
		client.Run()
	} else if mode == "s" {
		server := server.NewTestServer()
		server.Run()
	}
}
