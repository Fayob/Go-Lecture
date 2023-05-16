package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func handle(conn net.Conn)  {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		fmt.Printf("I heard you say %s\n", ln)
	}
	defer conn.Close()

	// It got here?
	fmt.Println("Code go here?")
}

func main()  {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}

	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Panic(err)
		}
		go handle(conn)
	}
}