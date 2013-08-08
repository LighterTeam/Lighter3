/**
 * Created with IntelliJ IDEA.
 * User: Administrator
 * Date: 13-8-7
 * Time: 下午4:17
 * To change this template use File | Settings | File Templates.
 */
package main

import (
	"TSTCP"
	"fmt"
	"net"
)

func FunServerInit() {
	fmt.Println("网关适配服启动成功!")
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

func GoAdaptServer() {
	tcp := new(TSTCP.TSTCP)
	tcp.Create_Server("localhost:9188", FunServerInit, FunConnectNew, FunReceiveBuffer, FunConnectClose)
}
