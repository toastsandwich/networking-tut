package main

import (
	"encoding/binary"
	"fmt"
	"log"
	"net"
	"syscall"
)

func main() {
	// Open a raw socket (IPv4 and TCP)
	fd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_RAW, syscall.IPPROTO_TCP)
	if err != nil {
		log.Fatalf("Error opening raw socket: %v\n", err)
	}
	defer syscall.Close(fd)

	// Make sure to run the program with root privileges
	fmt.Println("Listening for TCP traffic (port filtering)...")

	// Loop to continuously capture packets
	for {
		packet := make([]byte, 65536)
		_, _, err := syscall.Recvfrom(fd, packet, 0)
		if err != nil {
			log.Fatalf("Error receiving packet: %v\n", err)
		}

		// Parse and print packet details
		handlePacket(packet)
	}
}

// Handle the incoming packet
func handlePacket(packet []byte) {
	// Extract the IP header (20 bytes)
	ipHeader := packet[0:20]
	srcIP := net.IP(ipHeader[12:16])
	dstIP := net.IP(ipHeader[16:20])

	// Extract the TCP header (the IP header length is 20 bytes, so TCP starts from byte 20)
	tcpHeader := packet[20:40]
	srcPort := binary.BigEndian.Uint16(tcpHeader[0:2])
	dstPort := binary.BigEndian.Uint16(tcpHeader[2:4])

	// Only monitor specific ports (e.g., port 80 for HTTP)
	if dstPort == 80 || dstPort == 443 {
		fmt.Printf("Packet from %s:%d to %s:%d\n", srcIP, srcPort, dstIP, dstPort)
	}
}
