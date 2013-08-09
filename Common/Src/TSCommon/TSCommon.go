package TSCommon

import (
	"fmt"
	"runtime"
)

func init () {
	fmt.Println("TSCommon init");
	runtime.GOMAXPROCS(runtime.NumCPU())	
}

var GateWayAdaptServer_IP string = "127.0.0.1"	// 网关适配服务器IP
var GateWayAdaptServer_Port int = 9000			// 网关适配服务器Port

var GameAdaptServer_IP string = "127.0.0.1"		// 游戏适配服务器IP
var GameAdaptServer_Port int = 9001				// 游戏适配服务器Port

var GateWayServer_IP string = "127.0.0.1"		// 网关服务器IP
var GateWayServer_Port_Min int = 10000			// 网关服务器Port 动态区间 10000-19999
var GateWayServer_Port_Max int = 19999

var GameServer_IP string = "127.0.0.1"			// 游戏服务器IP
var GameServer_Port_Min int = 20000				// 游戏服务器Port 动态区间 20000-29999
var GameServer_Port_Max int = 29999

var DBServer_IP string = "127.0.0.1"					// DB前端IP
var DBServer_Port int = 9002							// DB前端Port

var LoginServer_IP string = "127.0.0.1"				// 登陆服IP
var LoginServer_Port int = 9003						// 登陆服Port

var ChatServer_IP string = "127.0.0.1"				// 聊天服IP
var ChatServer_Port int = 9004						// 聊天服Port

const Max_Pool_GateWayServer int = 100					//网关池上限


