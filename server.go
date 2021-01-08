package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

//local buffer
var messages []string

func main() {
	//lister object and error handling
	l, err := net.Listen("tcp", ":1200")
	if err != nil {
		log.Fatal(err)
	}
	//closing the listener at the end
	defer l.Close()
	//Infinite loop to handle multiple clients at once
	for {
		//connection object and error handling
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		//goroutine for each client to increase concureent excecution
		go func(c net.Conn) {
			//reading the 1st message from the connected client
			message, err := bufio.NewReader(conn).ReadString('\n')
			if err != nil {
				log.Printf("Error: %+v", err.Error())
				return
			}
			//casting the received message into string and removing any trailing spaces
			message = string(message)
			message = strings.TrimSpace(message)
			fmt.Fprintf(conn, "hey...\n")
			//log.Println("Message:", message, len(message))

			//checking if the connected cleint wants to send data or receive data
			if message == "sender" {
				//receiving messages from the sedning client until it stopss
				for {
					message, err := bufio.NewReader(conn).ReadString('\n')
					if err != nil {
						log.Printf("Error: %+v", err.Error())
						return
					}
					message = string(message)
					message = strings.TrimSpace(message)
					//if the client wants to stop sending messages then it sends "exit" as a message
					if message != "exit" {
						messages = append(messages, message)
						fmt.Fprintf(conn, "hey...\n")
						log.Println("Message:", message, len(message))
					} else {
						fmt.Fprintf(conn, "exit")
						break
					}
				}
			} else if message == "receiver" { //if the connected client wants receive the data then it sends "receiver" as the 1st message
				//fmt.Println(messages)
				//sending the buffered messages to the client
				for _, val := range messages {
					fmt.Println(val)
					fmt.Fprintf(conn, val+"\n")
				}
				fmt.Fprintf(conn, "exit")
			}
		}(conn)
	}
}
