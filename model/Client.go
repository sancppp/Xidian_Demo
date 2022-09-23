package model

import "net"

type Client struct {
	Conn     net.Conn
	UserName string
}
