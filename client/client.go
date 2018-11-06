package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
)

type client struct {
	userName string
	userPass string 
	msg      string
}

func main() {
	serverAddr := "127.0.0.1"
	if len(os.Args) == 2 {
		serverAddr = os.Args[1]
	}
	conn, err := net.Dial("tcp", serverAddr+":8000")
	defer conn.Close()
	if err != nil {
		fmt.Printf("Please check the address of the server(%v)\n", serverAddr)
		os.Exit(-1)
	}
	fmt.Printf("Please enter your name:")
	var userName string
	newClient := &client{userName, "",""}
	input := bufio.NewScanner(os.Stdin)
	//msg:=bufio.NewScanner(conn)

	//消息格式：code+msg，其中code：0表示用户名，1表示密码，2表示聊天消息
	if input.Scan() {
		newClient.userName = input.Text()
	}
	_, err = fmt.Fprintf(conn, strconv.Itoa(0)+newClient.userName+"\n")
	if err != nil {
		fmt.Printf("Name send error!\n")
		os.Exit(-1)
	}
	fmt.Printf("Please enter your password:")
	if input.Scan() {
		newClient.userPass = input.Text()
	}
	_, err = fmt.Fprintf(conn, strconv.Itoa(1)+newClient.userName+"\n")
	if err != nil {
		fmt.Printf("Password send error!\n")
		os.Exit(-1)
	}
	go handleMsgFromServer(conn)
	/*if msg.Scan(){
		code:=msg.Text()
		if code[0]==[]byte("2")[0]{
			fmt.Println("Wrong password!")
			os.Exit(2)
		}
	}*/
	//fmt.Println("Send name done!")
	inputMsg := bufio.NewScanner(os.Stdin)
	for inputMsg.Scan() {
		newClient.msg = inputMsg.Text() + "\n"
		_, err = fmt.Fprintf(conn, strconv.Itoa(2)+newClient.msg+"\n")
		
		if err != nil {
			fmt.Printf("Send error! Message is:%v\n", newClient.msg)
			os.Exit(-1)
		}

	}

}

func handleMsgFromServer(conn net.Conn){
	input:=bufio.NewScanner(conn)
	for input.Scan(){
		code:=input.Text()
		if code[0]==[]byte("3")[0]{
			fmt.Println("Please logon in first!")
		}else if code[0]==[]byte("0")[0]{
			continue
		}else if code[0]==[]byte("2")[0]{
			fmt.Println("Wrong password!")
			os.Exit(2)
		}
			fmt.Println(code)
		
	}
}