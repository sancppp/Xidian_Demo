package main

import (
	"tianzhenxiongProject/mysql"
	"tianzhenxiongProject/server"
)

func main() {
	mysql.Default()
	server.Start()
}
