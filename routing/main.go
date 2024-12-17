package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	for _, packet := range incomingPackets {
		for _, route := range routeTable {
			r := packet.destinationAddr.Mask(net.IPMask(route.subnetMask))
			log.Println(r.String())
			if r.Equal(route.network) {
				fmt.Printf("For destination %s, nexthop is %s, network is %s\n", packet.destinationAddr, route.nextHop, route.network)
				break
			}
		}
	}
}
