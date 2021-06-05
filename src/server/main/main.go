package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

//specify connection
const (
	connHost = "localhost"
	connPort = "8080"
	connType = "tcp"
)

func main() {
	//print connection details
	fmt.Println("Starting " + connType + " server on " + connHost + ":" + connPort)
	//create listener
	ln, err := net.Listen(connType, connHost+":"+connPort)
	//handle errors by printing and sys exit
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	//close listener at the end
	defer ln.Close()

	conn, _ := ln.Accept()
	for {
		// get message, output
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Message Received:", message)
	}
}
