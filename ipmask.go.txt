package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println(os.Stderr, "Usage : %s dotted ip address \n", os.Args[0])
		os.Exit(1)
	}

	dotaddrs := os.Args[1]
	addrs := net.ParseIP(dotaddrs)
	if addrs == nil {
		fmt.Println("Invalid address")
		os.Exit(1)
	}
	mask := addrs.DefaultMask()
	network := addrs.Mask(mask)
	ones, bits := mask.Size()

	fmt.Println(" Address is ", addrs.String(), "\n",
		"Default mask length is ", bits, "\n",
		"Leading ones count ", ones, "\n",
		"Mask in hex ", mask.String(), "\n",
		"Network is ", network.String())
	os.Exit(0)
}
