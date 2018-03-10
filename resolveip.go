package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println(os.Stderr, "Usage : %s hostname \n", os.Args[0])
		os.Exit(1)
	}

	name := os.Args[1]

	addr, err := net.ResolveIPAddr("ip", name)
	if err != nil {
		fmt.Println(os.Stderr, "Resolution error")
		os.Exit(2)
	}
	fmt.Println("Resolved address is ", addr)
	os.Exit(0)
}
