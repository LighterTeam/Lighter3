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
	. "TSTCP"
	"TSUtil"
	"fmt"
	"net"
)

var tcp *TSTCP

func FunClientInit(conn net.Conn) {
	fmt.Println("客户端连接成功!")
}

func FunReceiveBuffer(conn net.Conn, sBuffer string) {

}

func FunConnectClose(conn net.Conn) {
	fmt.Println("客户端断开连接!")
}

func GoGWClient() {
	tcp = new(TSTCP)
	tcp.Create_Client(TSCommon.GameAdaptServer_IP+":"+TSUtil.ToString(TSCommon.GameAdaptServer_Port), FunClientInit, FunReceiveBuffer, FunConnectClose)
}
