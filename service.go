package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println(os.Stderr, "Usage : %s networktype services \n", os.Args[0])
		os.Exit(1)
	}

	networkType := os.Args[1]
	services := os.Args[2]

	port, err := net.LookupPort(networkType, services)
	if err != nil {
		fmt.Println("Error: ", err.Error())
		os.Exit(2)
	}
	fmt.Println("Service port ", port)
	os.Exit(0)
}
