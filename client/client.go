package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

const (
	connectionHost   = "localhost"
	connectionSocket = "127.0.0.1:8081"
	connectionType   = "tcp"
)

func establishConnection()bool{
	conn, err := net.Dial(connectionType, connectionSocket)
	if err != nil {
		fmt.Printf("\nError in connecting to network : %v", err)
		return false
	}

	questionSet, err := bufio.NewReader(conn).ReadString('\r')
	if err != nil {
		fmt.Printf("\nError in reading from connection : %v\n", err)
		return false
	}
	fmt.Printf("\n Question Set : \n %+v \n", questionSet)
	handleRequest(conn)
	return true
}

func handleRequest(conn net.Conn) {
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Printf("Client Request - ")
		text, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("\nError in reading input ... : \n", err)
			return
		}
		fmt.Fprintf(conn, text+"\n")
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Printf("\nError in reading from connection : %v\n", err)
			return
		}
		fmt.Printf("Server Response - " + message)
	}
}

func main() {
	establishConnection()
}
