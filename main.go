package main

import (
	"errors"
	"fmt"
	"log"
	"net"
	"os"
	"regexp"
	"strings"
)

func main() {
	args := os.Args

	if len(args) < 3 {
		fmt.Println("Port-knocking client.")
		fmt.Println("https://github.com/jhspetersson/nokku")
		fmt.Println("Usage: nokku HOSTNAME[:PORT] [HOSTNAME2:]PORT2 [[HOSTNAME3:]PORT3 [[HOSTNAME4:]PORT4 [...]]]")
		os.Exit(1)
	}

	var host, port string

	for _, arg := range args[1:] {
		parsedHost, parsedPort, err := parseArg(arg)

		if err != nil {
			log.Fatalf("Error: %s", err)
		}

		if parsedHost != "" {
			host = parsedHost
		}

		if parsedPort != "" {
			if host == "" {
				log.Fatal("set hostname first")
			}

			port = parsedPort
		}

		if host != "" && port != "" {
			addr := host + ":" + port

			fmt.Println("Knocking to " + addr)

			conn, err := net.Dial("tcp", addr)
			if err == nil {
				_ = conn.Close()
			}
		}
	}

	fmt.Println("Sequence completed!")
}

func parseArg(arg string) (host, port string, err error) {
	if strings.Contains(arg, ":") {
		parts := strings.Split(arg, ":")
		host = parts[0]
		port = parts[1]
	} else if strings.Contains(arg, ".") {
		host = arg
	} else if ok, err := regexp.MatchString("^\\d+$", arg); ok && err == nil {
		port = arg
	} else {
		err = errors.New("could not parse argument")
	}

	return
}
