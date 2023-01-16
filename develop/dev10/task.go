package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
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

	Conn.host = flag.Arg(0)
	Conn.port = flag.Arg(1)
	processDuration(t)

	go startServer()

	var conn net.Conn
	var err error
	var end = time.Now().Add(Conn.timeout)

	for time.Now().Before(end) {
		conn, err = net.Dial("tcp", Conn.host+":"+Conn.port)
		if err == nil {
			break
		}
		time.Sleep(time.Millisecond * 33)
	}
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	log.Println("connected")

	// Чтение из сокета
	go func() {
		r := bufio.NewScanner(conn)
		for r.Scan() {
			fmt.Println(r.Text())
		}
	}()

	// Запись ввода в сокет
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		_, err := fmt.Fprintln(conn, s.Text())

		// При закрытии сокета сервером заканчиваем работу
		if err != nil {
			log.Println(err)
			return
		}
	}
}

// Обработка флага таймаута
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

// Запуск сервера
func startServer() {
	ln, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Println(err)
	}
	for {
		conn, err := ln.Accept()
		log.Println("connected to server")
		if err != nil {
			log.Println(err)
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	log.Println("handling connection", conn)
	r := bufio.NewScanner(conn)
	for r.Scan() {
		_, err := fmt.Fprintln(conn, "my answer:", r.Text())
		if err != nil {
			log.Println(err)
			return
		}
	}

}
