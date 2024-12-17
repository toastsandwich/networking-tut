package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Println("usage: showip <ip>")
		return
	}

	host := os.Args[1]
	addrs, err := net.LookupIP(host)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("ip address(es) for", host, ":")
	for _, addr := range addrs {
		if addr.To4() != nil {
			fmt.Printf("\tipV4: %s\n", addr.String())
		} else {
			fmt.Printf("\tipV6: %s\n", addr.String())
		}
	}
}
