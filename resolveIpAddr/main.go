package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	args := os.Args
	if len(args) != 3 {
		exit("usage: resolveIP <protocol> <name>")
	}
	protocol := args[1]
	name := args[2]

	addr, err := net.ResolveIPAddr(protocol, name)
	if err != nil {
		exit(err.Error())
	}

	fmt.Println(name,":", addr.String())
}

func exit(mssg string) {
	fmt.Println(mssg)
	os.Exit(1)
}
