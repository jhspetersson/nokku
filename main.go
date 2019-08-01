package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	args := os.Args

	if len(args) < 3 {
		fmt.Println("Port-knocking client.")
		fmt.Println("https://github.com/jhspetersson/nokku")
		fmt.Println("Usage: nokku HOSTNAME PORT1 [PORT2 [PORT3 [...]]]")
		os.Exit(1)
	}

	host := args[1]

	for _, port := range args[2:] {
		addr := host + ":" + port
		fmt.Println("Knocking to " + addr)
		conn, err := net.Dial("tcp", addr)
		if err == nil {
			_ = conn.Close()
		}
	}

	fmt.Println("Sequence completed!")
}
