package main

import "net"

type Route struct {
	subnetMask net.IP
	network    net.IP
	nextHop    net.IP
}

type RoutingTable []Route

type Packet struct {
	sourceAddr      net.IP
	destinationAddr net.IP
	data            []byte
}

type IncomingPackets []Packet

var (
	routeTable = RoutingTable{
		Route{
			subnetMask: net.IPv4(255, 255, 255, 0),
			network:    net.IPv4(192, 168, 1, 0),
			nextHop:    net.IPv4(10, 10, 10, 1),
		},
		Route{
			subnetMask: net.IPv4(255, 255, 255, 240),
			network:    net.IPv4(192, 168, 1, 128),
			nextHop:    net.IPv4(10, 10, 10, 2),
		},
		Route{
			subnetMask: net.IPv4(255, 255, 255, 0),
			network:    net.IPv4(192, 168, 2, 0),
			nextHop:    net.IPv4(10, 10, 10, 3),
		},
		Route{
			subnetMask: net.IPv4(0, 0, 0, 0),
			network:    net.IPv4(0, 0, 0, 0),
			nextHop:    net.IPv4(10, 10, 10, 10),
		},
	}

	incomingPackets = IncomingPackets{
		Packet{
			sourceAddr:      net.IPv4(192, 168, 2, 10),
			destinationAddr: net.IPv4(192, 168, 1, 20),
			data:            []byte("Hello, this is packet 1"),
		},
		Packet{
			sourceAddr:      net.IPv4(10, 0, 0, 5),
			destinationAddr: net.IPv4(192, 168, 2, 15),
			data:            []byte("Packet 2 with different data"),
		},
		Packet{
			sourceAddr:      net.IPv4(172, 16, 0, 3),
			destinationAddr: net.IPv4(172, 16, 1, 8),
			data:            []byte("This is packet 3"),
		},
		Packet{
			sourceAddr:      net.IPv4(192, 168, 3, 25),
			destinationAddr: net.IPv4(192, 168, 4, 30),
			data:            []byte("Final packet in the list"),
		},
		Packet{
			sourceAddr:      net.IPv4(1, 2, 3, 4),
			destinationAddr: net.IPv4(9, 8, 7, 6),
			data:            []byte("Unkown, send to 0.0.0.0"),
		},
	}
)
