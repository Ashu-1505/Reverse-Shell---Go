package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

const (
	SERVER_HOST = "192.168.56.1"
	SERVER_PORT = "9988"
	SERVER_TYPE = "tcp"
)

func main() {
	fmt.Println("Server Running...")
	server, err := net.Listen(SERVER_TYPE, SERVER_HOST+":"+SERVER_PORT)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	defer server.Close()
	fmt.Println("Listening on " + SERVER_HOST + ":" + SERVER_PORT)
	fmt.Println("Waiting for client...")
	for {
		connection, err := server.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		fmt.Println("client connected")
		go StartShell(connection)
	}
}

func StartShell(conn net.Conn) {
	defer conn.Close()

	fmt.Println("Shell Started with, ", conn.RemoteAddr())

	for {
		fmt.Println(">>")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		err := scanner.Err()
		if err != nil {
			log.Fatal(err)
		}
		tex := scanner.Text()
		if tex == "quit" {
			conn.Write([]byte(tex + "\n"))
			return
		}
		conn.Write([]byte(tex + "\n"))

		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}
		if message == "quit" {
			fmt.Print("Over")
			conn.Close()
		}
		fmt.Print("<---: " + message)

	}
}
