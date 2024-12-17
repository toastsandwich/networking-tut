package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		exit("usage: stick <ip>")
	}

	addr := net.ParseIP(args[1])
	if addr == nil {
		exit("invalid ip")
	}

	fmt.Println(addr.String())
}

func exit(mssg string) {
	fmt.Println(mssg)
	os.Exit(1)
}
