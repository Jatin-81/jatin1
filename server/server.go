package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

const (
	connectionHost = "localhost"
	connectionPort = "8081"
	connectionType = "tcp"
)

// func establishConnection() {

// }

func handleRequest(conn net.Conn) {
	conn.Write([]byte(fmt.Sprintf("%+v \r", Messages)))
	for {
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Printf("\nConnection Closed : %v \n%v\n", conn.RemoteAddr(), err)
			return
		}
		fmt.Printf("\nMessage got form connection %v :\n%v", conn.RemoteAddr(), message)
		newMessage, ok := Messages[strings.Trim(message, "\r \n")]
		if !ok {
			newMessage = "Could not recognize the message"
		}
		conn.Write([]byte(newMessage + "\n"))
		fmt.Printf("Response Sent :\n%v", newMessage)
	}
}

func main() {
	listener, err := net.Listen(connectionType, connectionHost+":"+connectionPort)
	if err != nil {
		fmt.Printf("\n Error in Listening to server !! \n %v", err)
		os.Exit(1)
	}
	fmt.Printf("Server Listening on %v:%v\n", connectionHost, listener.Addr())
	defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("\n Error accepting: \n", err.Error())
			os.Exit(1)
		}
		fmt.Printf("\nListening on %v", conn.RemoteAddr())
		defer conn.Close()
		go handleRequest(conn)
	}
}
