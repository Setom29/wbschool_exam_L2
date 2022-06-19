package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"regexp"
	"syscall"
	"time"
)

/*
=== Утилита telnet ===

Реализовать примитивный telnet клиент:
Примеры вызовов:
go-telnet --timeout=10s host port
go-telnet mysite.ru 8080
go-telnet --timeout=3s 1.1.1.1 123

Программа должна подключаться к указанному хосту (ip или доменное имя) и порту по протоколу TCP.
После подключения STDIN программы должен записываться в сокет, а данные полученные и сокета должны выводиться в STDOUT
Опционально в программу можно передать таймаут на подключение к серверу (через аргумент --timeout, по умолчанию 10s).

При нажатии Ctrl+D программа должна закрывать сокет и завершаться. Если сокет закрывается со стороны сервера, программа должна также завершаться.
При подключении к несуществующему сервер, программа должна завершаться через timeout.
*/

type TelnetClient struct {
	timeout time.Duration
	host    string
	port    string
}

func newTelenetClient() (*TelnetClient, error) {
	timeout := flag.String("timeout", "10s", "Timeout to connect to the server.")
	flag.Parse()
	if match, _ := regexp.MatchString(`^[0-9]+s$`, *timeout); !match {
		return nil, errors.New(fmt.Sprintf("Invalid timeout value: %s", *timeout))
	}
	t, _ := time.ParseDuration(*timeout)
	if len(flag.Args()) != 2 {
		return nil, errors.New("Wrong args amount.")
	}
	return &TelnetClient{timeout: t, host: flag.Args()[0], port: flag.Args()[1]}, nil
}

func readFromSocket(conn net.Conn, errChan chan error) {
	input := make([]byte, 1024)
	for {
		n, err := conn.Read(input)
		if err != nil {
			errChan <- fmt.Errorf("remoute server stopped: %v", err)
			return
		}
		fmt.Println(string(input[:n]))
	}
}

// writeToSocket - read from stdin and write to conn
func writeToSocket(conn net.Conn, errChan chan error) {
	reader := bufio.NewReader(os.Stdin)
	for {
		text, err := reader.ReadBytes('\n')
		if err != nil {
			errChan <- err
			return
		}
		// remove "\n"
		text = text[:len(text)-1]

		_, err = conn.Write(text)
		if err != nil {
			errChan <- err
			return
		}
	}
}

func main() {
	tClient, err := newTelenetClient()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(*tClient)

	conn, err := net.DialTimeout("tcp", tClient.host+":"+tClient.port, tClient.timeout)
	if err != nil {
		fmt.Println("Failed to establish a connection with the server.", err)
	}
	fmt.Println("The connection is established")

	defer conn.Close()

	// handle signal
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	errChan := make(chan error)

	go readFromSocket(conn, errChan)
	go writeToSocket(conn, errChan)

	select {
	case s := <-sigs:
		fmt.Println("\nConnection stopped by signal:", s)
	case e := <-errChan:
		fmt.Println("Connection stopped by", e)
	}
}
