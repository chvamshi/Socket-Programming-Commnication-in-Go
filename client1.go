package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	// resolving the TCP address and error handling
	tcpAddr, err := net.ResolveTCPAddr("tcp4", ":1200")
	checkError(err)
	// connecting to the server and error handling
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)
	// reader := bufio.NewReader(os.Stdin)

	reader := bufio.NewReader(os.Stdin)
	// sending a message to indicate that the client wants to receive
	fmt.Print("Text to send: ")
	text, _ := reader.ReadString('\n')
	// send to server
	fmt.Fprintf(conn, text+"\n")
	// wait for reply
	// to continously receive messages from the server buffer
	for {
		message, _ := bufio.NewReader(conn).ReadString('\n')
		message = string(message)
		message = strings.TrimSpace(message)
		fmt.Println("Message from server: " + message)
		if message == "exit" {
			break
		}
	}
	os.Exit(0)
}

// small function to handle the errors
func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
