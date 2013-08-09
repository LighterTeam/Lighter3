package main

import "net"

type CGateWay struct {
	Conn net.Conn
	UUID uint64
}


