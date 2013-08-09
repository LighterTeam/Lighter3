/**
 * Created with IntelliJ IDEA.
 * User: Administrator
 * Date: 13-8-7
 * Time: 下午4:17
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

var Pool_GateWayServer [TSCommon.Max_Pool_GateWayServer]*CGateWay

func GoGameServer() {
	tcp := new(TSTCP)
	tcp.Create_Server(":"+TSUtil.ToString(TSCommon.GateWayAdaptServer_Port),
		func () {
			fmt.Println("网关适配服启动成功! Port:",TSUtil.ToString(TSCommon.GateWayAdaptServer_Port))
		},
		func (conn net.Conn) uint64 {
			UUID := adaptUUID_GW()
			Pool_GateWayServer[UUID] = &CGateWay{conn, UUID}
			fmt.Println("GateWay客户端上线! 获取动态UUID:", UUID)
			return UUID
		},
		func (conn net.Conn, sBuffer string, UUID uint64) {
			fmt.Println("GateWay客户端 Buffer", sBuffer, " UUID:", UUID)
		},
		func (conn net.Conn, UUID uint64) {
			fmt.Println("GateWay客户端下线!", UUID)
			Pool_GateWayServer[UUID] = nil
		},
	)
}

func adaptUUID_GW() uint64 {
	var id uint64
	for i := 0; i < len(Pool_GateWayServer); i++ {
		if Pool_GateWayServer[i] == nil {
			id = uint64(i)
			break
		}
	}
	return id
}
