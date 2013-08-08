/**
 * Created with IntelliJ IDEA.
 * User: Administrator
 * Date: 13-8-8
 * Time: 上午9:52
 * To change this template use File | Settings | File Templates.
 */
package main

import (
	"TSTCP"
	"fmt"
	"net"
)

var tstcp *TSTCP.TSTCP

func FunClientInit(conn net.Conn) {
	fmt.Println("客户端连接成功!")
	for i := 0; i < 10000; i++ {
		tstcp.SendBuffer("客户端断开连接")
	}
}

func FunReceiveBuffer(conn net.Conn, sBuffer string) {

}

func FunConnectClose(conn net.Conn) {
	fmt.Println("客户端断开连接!")
}

func GoGWClient() {
	tstcp = new(TSTCP.TSTCP)
	tstcp.Create_Client("localhost:9188", FunClientInit, FunReceiveBuffer, FunConnectClose)
}
