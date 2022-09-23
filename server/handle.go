package server

import (
	"fmt"
	"net"
	"strconv"
	"strings"
	"tianzhenxiongProject/model"
)

const (
	CreateStudent = iota + 1
	GetStudent
	ModifyStudent
	DeleteStudent
	SendMessage
	Quit
)

func handle(conn net.Conn) {
	//有一个来自客户端的连接
	name := conn.RemoteAddr().String()
	ConnMap[name] = model.Client{
		Conn:     conn,
		UserName: name,
	}
	defer delete(ConnMap, name)
	conn.Write([]byte("Login Success,your username is " + name + "\n"))
	SendAll("login", name)
	for {
		data := make([]byte, 255)
		messageLen, err := conn.Read(data)
		if messageLen <= 1 || err != nil {
			continue
		}
		fmt.Println(string(data[0 : messageLen-1]))
		msgStr := strings.Split(string(data[0:messageLen-1]), "$ ")
		fmt.Println(msgStr)
		var clientMessage Message
		if clientMessage.Operation, err = strconv.Atoi(msgStr[0]); err != nil {
			fmt.Println("Operation error!!" + err.Error())
			continue
		}
		clientMessage.Conn = conn
		clientMessage.Name = name
		clientMessage.Body = msgStr[1]
		switch clientMessage.Operation {
		case CreateStudent:
			clientMessage.CreateStudent()
		case GetStudent:
			clientMessage.GetStudent()
		case ModifyStudent:
			clientMessage.ModifyStudent()
		case DeleteStudent:
			clientMessage.DeleteStudent()
		case SendMessage:
			clientMessage.SendMessage()
		case Quit:
			clientMessage.Quit()
			conn.Close()
			break
		default:
			fmt.Println("Operation error!!")
		}
	}

}
