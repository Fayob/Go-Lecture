package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main()  {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Panic(err)
		}

		io.WriteString(conn, "\nHello from TCP server\n")
		io.WriteString(conn, "\nCool\n")
		fmt.Println(conn, "How is your day?")
		fmt.Printf("%v Well, I hope", conn)

		conn.Close()
	}
}