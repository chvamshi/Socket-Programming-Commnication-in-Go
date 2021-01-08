# Socket-Programming-Commnication-in-Go
Message oriented communication is an alternative to remote procedure calls and remote invocations. Message oriented communication is used in the cases like when it cannot be assumed that the receiving side is executing at the time request is issued and the nature in RPC i.e., a client is blocked until its request been processed may need to replaced.
•	Many distributed systems and applications are built directly on top of the simple message oriented-model offered by the transport layer i.e., transport level sockets.
•	The communication using sockets is done by using the following model.



The idea of my application is to transport the messages from one client to another client using sockets.


It starts with client 1 sending data to server and the server store the messages in a buffer and when the client 2 connects with the server it sends the messages stored in the buffer making it a persistent type of communication. 
The application is implemented using golang and made use of goroutines as much as possible.



Execution:
•	The client that wants to send data will send “sender” as its first message, which indicates server to store the messages in the buffer
•	The client that wants to receive the data will send “receiver” as its first message indicating the server to send the messages stored in the buffer
•	Both clients and server are connected through the server’s socket which is “127.0.0.1:1200”
•	The port number can be changed.
•	At 1st the server should be executed, followed the by the client that wants to send the data
•	In this example, the “client.go” is the sender program and hence it has to send the 1st message as “sender” and “client1.go” is the receiver program and it has to send the 1st message as “receiver”
