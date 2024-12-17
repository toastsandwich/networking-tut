package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("usage: mynetmask <bit>")
		os.Exit(1)
	}

	bits := args[1]

	b := ToInt(bits)
	if b > 32 {
		fmt.Println("usage: mynetmask <bit>")
		fmt.Println("bit less than 32 is required")
		os.Exit(1)
	}

	one := "1"
	zero := "0"
	binaryNetMask := strings.Repeat(one, int(b)) + strings.Repeat(zero, 32-int(b))

	fmt.Println("net mask in binary:", binaryNetMask)

	netMask := make([]string, 0)
	for i := 0; i <= 24; i += 8 {
		nm := string(binaryNetMask[i : i+8])
		netMask = append(netMask, fmt.Sprintf("%d", BinToInt(nm)))
	}
	fmt.Println("net mask:", strings.Join(netMask, "."))

	hosts := 255 - ToInt(netMask[3])
	fmt.Printf("%d host addresses given (1 is provider), total: %d\n", hosts-1, hosts)
}

func ToInt(str string) int {
	i, err := strconv.ParseInt(str, 10, 32)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return int(i)
}

func BinToInt(str string) int {
	i, err := strconv.ParseInt(str, 2, 32)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return int(i)
}
