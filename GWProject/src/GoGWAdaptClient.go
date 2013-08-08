/**
 * Created with IntelliJ IDEA.
 * User: Administrator
 * Date: 13-8-8
 * Time: 上午9:52
 * To change this template use File | Settings | File Templates.
 */
package main

import (
	"TSCommon"
	"TSUtil"
	. "TSTCP"
	"fmt"
	"net"
)

var tcp *TSTCP

func FunClientInit(conn net.Conn) {
	fmt.Println("客户端连接成功!")
	for i := 0; i < 10; i++ {
		tcp.SendBuffer("客户端断开连接")
	}
}

func FunReceiveBuffer(conn net.Conn, sBuffer string) {

}

func FunConnectClose(conn net.Conn) {
	fmt.Println("客户端断开连接!")
}

func GoGWAdaptClient() {
	tcp = new(TSTCP)
	localhost := TSCommon.GateWayAdaptServer_IP + ":" + TSUtil.ToString(TSCommon.GateWayAdaptServer_Port);
	tcp.Create_Client(localhost, FunClientInit, FunReceiveBuffer, FunConnectClose)
}
