package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		exit("usage: lookup <name>")
	}

	name := os.Args[1]

	cname, err := net.LookupCNAME(name)
	if err != nil {
		exit(err.Error())
	}
	fmt.Println(cname)

	m, err := net.LookupMX(name)
	if err != nil {
		exit(err.Error())
	}
	mx(m)

	txt, err := net.LookupTXT(name)
	if err != nil {
		exit(err.Error())
	}
	fmt.Println(txt)

	addrs, err := net.LookupHost(cname)
	if err != nil {
		exit(err.Error())
	}

	for _, addr := range addrs {
		fmt.Println(addr)
	}
}

func exit(mssg string) {
	fmt.Println(mssg)
	// os.Exit(1)
}

func mx(MX []*net.MX) {
	for _, mx := range MX {
		fmt.Println(*mx)
	}
}
