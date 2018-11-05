package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

type client struct {
	userName string
	userAddr string
	msg      string
	online   bool
}

func main() {
	serverPort := "8000"
	if len(os.Args) == 2 {
		serverPort = os.Args[1]
	}
	ln, err := net.Listen("tcp", ":"+serverPort)
	if err != nil {
		fmt.Printf("占用端口失败！\n")
	}
	for {
		conn, err := ln.Accept()
		//fmt.Printf("New client come in!\n")
		if err != nil {
			fmt.Printf("%v！\n", err)
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	var newClient client
	newClient.userAddr = conn.RemoteAddr().String()
	//fmt.Printf("New client come in!Address is %v\n", newClient.userAddr)

	inputMsg := bufio.NewScanner(conn)
	for inputMsg.Scan() {
		if newClient.online == false {
			newClient.userName = inputMsg.Text()
			fmt.Printf("%v enter the room\n", newClient.userName)
			newClient.online = true
		} else {
			newClient.msg = inputMsg.Text()
			fmt.Printf("%v:%v\n", newClient.userName, newClient.msg)
		}

	}
	fmt.Printf("%v leave the room.\n",newClient.userName)
	defer conn.Close()
}
