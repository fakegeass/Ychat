package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

type client struct {
	userName string
	msg      string
}

func main() {
	serverAddr := os.Args[1]
	conn, err := net.Dial("tcp", serverAddr+":8000")
	if err != nil {
		fmt.Printf("Please check the address of the server(%v)\n", serverAddr)
		os.Exit(-1)
	}
	fmt.Printf("Please enter your name:")
	var userName string
	newClient := &client{userName, ""}
	input := bufio.NewScanner(os.Stdin)
	if input.Scan() {
		newClient.userName = input.Text()+"\n"
		fmt.Printf("Welcome  %v\n",newClient.userName)
	}
	_, err = fmt.Fprintf(conn, newClient.userName)
	if err != nil {
		fmt.Printf("Name send error!\n")
		os.Exit(2)
	}
	fmt.Println("Send name done!")
	inputMsg := bufio.NewScanner(os.Stdin)
		for inputMsg.Scan() {
			newClient.msg = inputMsg.Text()+"\n"
			_, err = fmt.Fprintf(conn, newClient.msg)
		if err != nil {
			fmt.Printf("Send error! Message is:%v\n", newClient.msg)
			os.Exit(2)
		}
		
	}

}
