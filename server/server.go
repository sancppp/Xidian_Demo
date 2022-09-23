package server

import (
	"fmt"
	"net"
	"tianzhenxiongProject/model"
	"time"
)

var ConnMap = make(map[string]model.Client)

func Start() {
	listen, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Printf("server listen error:%v", err)
		return
	}
	fmt.Println("Server Start Success")
	// 使用 defer 在运行结束后优雅的关闭
	defer listen.Close()

	for {
		// 当接收到连接请求时
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("conn fail ...")
			continue
		}
		fmt.Println(conn.RemoteAddr(), "connect successed!")

		// handle 为每一个客户端开单独的协程进行业务操作
		go handle(conn)
	}

}

func SendAll(msg string, userName string) {
	for name, client := range ConnMap {
		if name == userName {
			//不发送给自己
			continue
		}
		sendMessage := fmt.Sprintf("%v [%s]: %v\n", time.Now().Format("2006-01-02 15:04:05"), userName, msg)
		_, err := client.Conn.Write([]byte(sendMessage))
		if err != nil {
			fmt.Println("client Conn Error")
			delete(ConnMap, name)
			continue
		}
	}
}
