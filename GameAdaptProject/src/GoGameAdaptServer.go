package main

import (
	"TSCommon"
	. "TSTCP"
	"TSUtil"
	"fmt"
	"net"
)

func FunServerInit() {
	fmt.Println("游戏适配服启动成功!")
}

func FunConnectNew(conn net.Conn) {
	fmt.Println("客户端上线!")
}

func FunReceiveBuffer(conn net.Conn, sBuffer string) {
	fmt.Println("xx===============================:", sBuffer)
}

func FunConnectClose(conn net.Conn) {
	fmt.Println("客户端下线!")
}

func GoGWServer() {
	tcp := new(TSTCP)
	tcp.Create_Server(TSCommon.GameAdaptServer_IP+":"+TSUtil.ToString(TSCommon.GameAdaptServer_Port),
		FunServerInit, FunConnectNew, FunReceiveBuffer, FunConnectClose)
}
