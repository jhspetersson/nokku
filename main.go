package main

import (
	"errors"
	"fmt"
	"log"
	"net"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func main() {
	args := os.Args

	if len(args) < 3 {
		fmt.Println("Port-knocking client.")
		fmt.Println("https://github.com/jhspetersson/nokku")
		fmt.Println("Usage: nokku HOSTNAME[:PORT] [PROTO] [pause] [HOSTNAME2:]PORT2 [pause] [[HOSTNAME3:]PORT3 [[HOSTNAME4:]PORT4 [...]]]")
		os.Exit(1)
	}

	var host, port string
	var proto = "tcp"

	// delay between connection tries, 1 second by default
	var delay = 1

	// pause duration, 1 second by default
	var pause = 1

	// starting index to parse command sequence from program arguments
	var idx = 1

	if ok, paramIndex := locateParam(args, "-d", "--delay"); ok {
		if parsed, err := strconv.Atoi(args[paramIndex+1]); err != nil {
			delay = parsed
			idx += 2
		} else {
			log.Fatalf("Error parsing delay: %s", err)
		}
	}

	if ok, paramIndex := locateParam(args, "-p", "--pause"); ok {
		if parsed, err := strconv.Atoi(args[paramIndex+1]); err != nil {
			pause = parsed
			idx += 2
		} else {
			log.Fatalf("Error parsing pause: %s", err)
		}
	}

	for _, arg := range args[idx:] {
		if isPause(arg) {
			fmt.Println(".....")
			time.Sleep(time.Duration(pause) * time.Second)
			continue
		}

		parsedHost, parsedPort, parsedProto, err := parseArg(arg)

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

		if parsedProto != "" {
			proto = parsedProto
		}

		if host != "" && port != "" {
			addr := host + ":" + port

			fmt.Println("Knocking to " + addr + " (" + proto + ")")

			conn, err := net.Dial(proto, addr)
			if err == nil {
				_ = conn.Close()
			}

			time.Sleep(time.Duration(delay) * time.Second)
		}
	}

	fmt.Println("Sequence completed!")
}

func locateParam(args []string, params ...string) (ok bool, idx int) {
	for i, arg := range args {
		for _, param := range params {
			if strings.ToLower(param) == strings.ToLower(arg) {
				return true, i
			}
		}
	}

	return false, 0
}

func parseArg(arg string) (host, port, proto string, err error) {
	if strings.ToLower(arg) == "tcp" {
		proto = "tcp"
	} else if strings.ToLower(arg) == "udp" {
		proto = "udp"
	} else if strings.Contains(arg, ":") {
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

func isPause(arg string) bool {
	return strings.ToLower(arg) == "pause"
}
