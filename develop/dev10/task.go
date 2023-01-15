package main

import (
	"flag"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"time"
)

type connection struct {
	host    string
	port    string
	timeout time.Duration
}

var Conn connection

func main() {
	var t string
	flag.StringVar(&t, "timeout", "10s", "timeout for connection to server")
	flag.Parse()

	host := flag.Arg(0)
	port := flag.Arg(1)
	processDuration(t)
	fmt.Println(host, port, Conn.timeout)
}

func createConnection() connection {
	Conn := connection{
		host:    flag.Arg(0),
		port:    flag.Arg(1),
		timeout: time.Second,
	}
	return Conn
}
func processDuration(t string) {
	exp, err := regexp.Compile(`^\d+[smh]$`)
	if err != nil {
		log.Fatal("error handling timeout")
	}
	if !exp.MatchString(t) {
		log.Println("invalid timeout format. Usage: task [-timeout=10s] <host> <port>")
		Conn.timeout = 10 * time.Second
	} else {
		length, _ := strconv.Atoi(t[:len(t)-1])
		Conn.timeout = time.Duration(length) * time.Second
		switch t[len(t)-1] {
		case 'm':
			Conn.timeout *= 60
		case 'h':
			Conn.timeout *= 3600
		}
	}
}
