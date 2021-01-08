package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	// if len(os.Args) != 2 {
	// 	fmt.Fprintf(os.Stderr, "Usage: %s host:port ", os.Args[0])
	// 	os.Exit(1)
	// }
	// service := os.Args[1]
	// resolving the tcp address and error handling
	tcpAddr, err := net.ResolveTCPAddr("tcp4", ":1200")
	checkError(err)
	// connecting to the tcp server
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)
	// reader := bufio.NewReader(os.Stdin)
	// For sending as many messages as client wishes
	for {
		reader := bufio.NewReader(os.Stdin)
		// sedning a message to indicate that the client wants to send meesages
		fmt.Print("Text to send: ")
		text, _ := reader.ReadString('\n')
		// send messages to the server
		fmt.Fprintf(conn, text+"\n")
		// wait for reply
		message, _ := bufio.NewReader(conn).ReadString('\n')
		//result, _ := ioutil.ReadAll(conn)
		// if the meesage sent is 'exit' then server stops and again sends 'exit' which stops the client
		if message == "exit" {
			break
		}
		fmt.Print("Message from server: " + message)
	}
	os.Exit(0)
}

// a small function to handle errors
func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
