package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"github.com/gomodule/redigo/redis"
)

type client struct {
	userName string
	userAddr string
	msg      string
	online   bool
	userPass string
}

const debug bool = false

func init(){
	
	/*if err!=nil{
		fmt.Printf("Redis can't reach!%v\n",err)
	}*/
		
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
	allClient:=make(map[net.Conn]bool) 
	for {
		conn, err := ln.Accept()
		allClient[conn]=true
		//fmt.Printf("New client come in!\n")
		if err != nil {
			fmt.Printf("%v！\n", err)
		}
		go handleConnection(conn,&allClient)
	}
}

//消息格式：code+msg，其中code：0表示成功，1表示数据库链接问题，2表示密码错误，3表示用户未上线，4表示用户重复上线
//5表示服务器端错误
func handleConnection(conn net.Conn,allClient *map[net.Conn]bool) {
	var newClient client
	newClient.userAddr = conn.RemoteAddr().String()
	//fmt.Printf("New client come in!Address is %v\n", newClient.userAddr)
	inputMsg := bufio.NewScanner(conn)
	for inputMsg.Scan() {
		msg:=inputMsg.Text()
	if msg==""{
		fmt.Fprintf(conn, strconv.Itoa(0)+"\n")
	}else{
	if debug{
		fmt.Printf("Received msg: %#v",msg)
	}
	switch msg[0]{
	case []byte("0")[0]:
		if newClient.online==true{
			fmt.Fprintf(conn, strconv.Itoa(4)+"\n")
			
		}
		newClient.userName=string(msg[1:])
		fmt.Fprintf(conn, strconv.Itoa(0)+"\n")
		
	case []byte("1")[0]:
		if newClient.online==true{
			fmt.Fprintf(conn, strconv.Itoa(4)+"\n")
			
		}
		conSql,err:=redis.Dial("tcp","127.0.0.1:6379")
		defer conSql.Close()
		tempPass,_:=redis.String(conSql.Do("get",newClient.userName))
		//fmt.Printf("%T",tempPass)
		newClient.userPass=string(msg[1:])
			if err!=nil{
				fmt.Println("Redis can't reach!")
				fmt.Fprintf(conn, strconv.Itoa(5)+"\n")
				 
			}else if tempPass==""{
				_,err=conSql.Do("set",newClient.userName,newClient.userPass)
				newClient.online=true
				fmt.Fprintf(conn, strconv.Itoa(0)+"\n")
				 
			}else if tempPass!=newClient.userPass{
				fmt.Printf("inputPass is %v,saved is %v\n",newClient.userPass,tempPass)
				fmt.Printf("Password for %v not correct!\n",newClient.userName)
				fmt.Fprintf(conn, strconv.Itoa(2)+"\n")
				 
			}else{
				newClient.online=true
				fmt.Printf("%v enter the room\n", newClient.userName)
				fmt.Fprintf(conn, strconv.Itoa(0)+"\n")
				
			}
		case []byte("2")[0]:
			if newClient.online!=true{
				fmt.Printf("User(%v) is not online.\n",newClient.userName)
				fmt.Fprintf(conn, strconv.Itoa(3)+"\n")
				 
			}
			fmt.Printf("%v:%v\n", newClient.userName, string(msg[1:]))
			for connTemp:=range *allClient{
				fmt.Fprintf(connTemp,"0%v:%v\n",newClient.userName,string(msg[1:]))
			}
			
	}}
}
	fmt.Printf("%v leave the room.\n",newClient.userName)
	newClient.online=false
	defer conn.Close()
	return 
}