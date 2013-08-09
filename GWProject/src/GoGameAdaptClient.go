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
	_ "fmt"
	"net"
)

//func FunClientInit(conn net.Conn) {
//	fmt.Println("客户端连接成功!")
//}

//func FunReceiveBuffer(conn net.Conn, sBuffer string, UUID uint64) {

//}

//func FunConnectClose(conn net.Conn, UUID uint64) {
//	fmt.Println("客户端断开连接!")
//}

func GoGameAdaptClient() {
	tcp := new(TSTCP)
	localhost := TSCommon.GameAdaptServer_IP + ":" + TSUtil.ToString(TSCommon.GameAdaptServer_Port)
	tcp.Create_Client(localhost,
		func(conn net.Conn) {

		},
		func(conn net.Conn, sBuffer string, UUID uint64) {

		},
		func(conn net.Conn, UUID uint64) {

		})
}
