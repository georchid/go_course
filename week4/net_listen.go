package main

import (
	"bufio"
	"fmt"
	"net"
)

func handleConnction(conn net.Conn) {
	name := conn.RemoteAddr().String()
	fmt.Printf("%+v connected\n", name)
	conn.Write([]byte("Hello " + name + "\n\r"))
	defer conn.Close()

	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		text := scanner.Text()
		if text == "Exit" {
			conn.Write([]byte("Bye\n\r"))
			fmt.Println(name, "disconnected")
			break
		} else if text != "" {
			fmt.Println(name, "enter", text)
			conn.Write([]byte("You enter " + text + "\n\r"))
		}
	}
}

func main() {
	listener, err := net.Listen("tcp", ":8081")
	if err != nil {
		panic(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		go handleConnction(conn)
	}
}
