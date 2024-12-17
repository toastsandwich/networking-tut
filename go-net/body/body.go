package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	args := os.Args
	if len(args) != 3 {
		fmt.Println("usage: body <host> <port>")
		return
	}

	ip := os.Args[1]
	port := os.Args[2]

	conn, err := net.Dial("tcp", ip+":"+port)
	if err != nil {
		log.Fatal(err)
	}
	for {
		buf := make([]byte, 4096)
		input := bufio.NewReader(os.Stdin)
		bytes, err := input.ReadBytes('\n')
		if err != nil {
			log.Println(err)
			continue
		}
		_, err = conn.Write(bytes)
		if err != nil {
			log.Println(err)
			continue
		}
		n, err := conn.Read(buf)
		if err != nil {
			log.Println(err)
			continue
		}
		msg := string(buf[:n])
		fmt.Println(msg)
	}
}
