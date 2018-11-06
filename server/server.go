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

var conSql redis.Conn

func init(){
	
	/*if err!=nil{
		fmt.Printf("Redis can't reach!%v\n",err)
	}*/
		
}

func main() {
	conSql,_:=redis.Dial("tcp","127.0.0.1:6379")
	defer conSql.Close()
	serverPort := "8000"
	if len(os.Args) == 2 {
		serverPort = os.Args[1]
	}
	ln, err := net.Listen("tcp", ":"+serverPort)
	if err != nil {
		fmt.Printf("占用端口失败！\n")
	}
	//fmt.Printf("sql-be %v\n",conSql)
	for {
		conn, err := ln.Accept()
		//fmt.Printf("New client come in!\n")
		if err != nil {
			fmt.Printf("%v！\n", err)
		}
		go handleConnection(conn,conSql)
	}
}

func handleConnection(conn net.Conn,conSql redis.Conn) {
	var newClient client
	newClient.userAddr = conn.RemoteAddr().String()
	//fmt.Printf("New client come in!Address is %v\n", newClient.userAddr)

	inputMsg := bufio.NewScanner(conn)
	for inputMsg.Scan() {
		if newClient.online == false {
			/*newClient.userName = inputMsg.Text()
			tempPass,err:=conSql.Do("get",newClient.userName)
			if err!=nil{
				fmt.Println("Redis can't reach!")
			}else if tempPass==""{
				_,err:=conSql.Do("set",client.userName,client.userPass)
			}else if tempPass!=client.userPass{
				fmt.Printf("Password for %v not correct!\n",client.userName)
				return
			}
			}
			fmt.Printf("%v enter the room\n", newClient.userName)*/
			msg:=inputMsg.Text()
			code:=handleMsg([]byte(msg),&newClient,conn,conSql)
			if code==2{
				_, _= fmt.Fprintf(conn, strconv.Itoa(2)+"\n")
			}else if code==0{
				_, _= fmt.Fprintf(conn, strconv.Itoa(0)+"\n")
			}

		} else {
			msg:=inputMsg.Text()
			code:=handleMsg([]byte(msg),&newClient,conn,conSql)
			if code==3{
				_, _ = fmt.Fprintf(conn, strconv.Itoa(3)+"\n")
			}else if code==0{
				_,_ = fmt.Fprintf(conn, strconv.Itoa(0)+"\n")
			}else if code==2{
				_, _= fmt.Fprintf(conn, strconv.Itoa(2)+"\n")
			}
		}

	}
	fmt.Printf("%v leave the room.\n",newClient.userName)
	newClient.online=false
	defer conn.Close()
}

//消息格式：code+msg，其中code：0表示成功，1表示数据库链接问题，2表示密码错误，3表示用户未上线
func handleMsg(msg []byte, newclient *client,conn net.Conn,conSql redis.Conn) int {
	if len(msg)==0{
		return 0
	}
	switch msg[0]{
	case []byte("0")[0]:
		newclient.userName=string(msg[1:])
		return 0
	case []byte("1")[0]:
		tempPass,err:=redis.String(conSql.Do("get",newclient.userName))
		//fmt.Printf("%T",tempPass)
		newclient.userPass=string(msg[1:])
			if err!=nil{
				fmt.Println("Redis can't reach!")
				return 1
			}else if tempPass==""{
				_,err=conSql.Do("set",newclient.userName,newclient.userPass)
				return 0
			}else if tempPass!=newclient.userPass{
				fmt.Printf("inputPass is %v,saved is%v\n",newclient.userPass,tempPass)
				fmt.Printf("Password for %v not correct!\n",newclient.userName)
				return 2
			}else{
				newclient.online=true
				fmt.Printf("%v enter the room\n", newclient.userName)
				return 0
			}
		case []byte("2")[0]:
			if newclient.online!=true{
				fmt.Printf("User(%v) is not online.\n",newclient.userName)
				return 3
			}
			fmt.Printf("%v:%v\n", newclient.userName, string(msg[1:]))
			fmt.Fprintf(conn,"%v:%v\n",newclient.userName,string(msg[1:]))
			return 0
	}
	return 0
}