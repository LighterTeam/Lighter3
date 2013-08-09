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

func GoGWAdaptClient() {
	tcp := new(TSTCP)
	localhost := TSCommon.GateWayAdaptServer_IP + ":" + TSUtil.ToString(TSCommon.GateWayAdaptServer_Port)
	tcp.Create_Client(localhost,
		func(conn net.Conn) {
			fmt.Println("GW客户端连接成功!")
		},
		func(conn net.Conn, sBuffer string, UUID uint64) {

		},
		func(conn net.Conn, UUID uint64) {
			fmt.Println("GW客户端断开连接!")
		})
}
