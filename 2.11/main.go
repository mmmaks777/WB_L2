package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"time"
)

func main() {
	timeout := flag.Duration("timeout", 10*time.Second, "Connection timeout")
	flag.Parse()

	args := flag.Args()
	if len(args) != 2 {
		fmt.Println("Usage: go-telnet [--timeout=10s] host port")
		os.Exit(1)
	}

	host := args[0]
	port := args[1]

	address := net.JoinHostPort(host, port)
	conn, err := net.DialTimeout("tcp", address, *timeout)
	if err != nil {
		fmt.Println("Connection error:", err)
		os.Exit(1)
	}
	defer conn.Close()

	done := make(chan struct{})

	go func() {
		_, err = io.Copy(os.Stdout, conn)
		if err != nil && !errors.Is(err, net.ErrClosed) {
			fmt.Println("Error reading from socket:", err)
		}
		done <- struct{}{}
	}()

	go func() {
		reader := bufio.NewReader(os.Stdin)
		for {
			input, err := reader.ReadString('\n')
			if err != nil {
				if errors.Is(err, io.EOF) {
					break
				}
				fmt.Println("Error reading from STDIN:", err)
			}
			_, err = conn.Write([]byte(input))
			if err != nil {
				fmt.Println("Error writing to socket: ", err)
			}
		}
		done <- struct{}{}
	}()

	fmt.Printf("Connected to %s:%s\n", host, port)
	<-done
	close(done)
	fmt.Println("Connection closed")
}
