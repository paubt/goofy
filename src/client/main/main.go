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
	fmt.Println("Connecting to " + connType + "server" + connHost + ":" + connPort)
	//try to make a connection to the specified connection
	conn, err := net.Dial(connType, connHost+":"+connPort)
	if err != nil {
		fmt.Println("Error connecting:", err.Error())
		os.Exit(1)
	}

	for {
		// what to send?
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Text to send: ")
		text, _ := reader.ReadString('\n')
		// send to server
		fmt.Fprintf(conn, text+"\n")
		// wait for reply
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Message from server: " + message)
	}
}
