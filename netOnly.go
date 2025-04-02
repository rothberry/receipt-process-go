package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func netOnly() {
	fmt.Println("Startin")

	ln, err := net.Listen("tcp", ":8000")
	fmt.Println("ln", ln)
	if err != nil {
		fmt.Println("Notat working", err)
		os.Exit(1)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Seconda not working", err)
			os.Exit(2)
		}
		fmt.Println("conn", conn)
		io.WriteString(os.Stdout, "It worked?")
		// go handleConnection(conn)
	}

}
