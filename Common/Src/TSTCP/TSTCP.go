package TSTCP

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"time"
	"TSUtil"
	"strings"
	"strconv"
)

type ConnectNew func(conn net.Conn) uint64
type ReceiveBuffer func(conn net.Conn, sBuffer string, UUID uint64)
type ServerInit func()
type ClientInit func(conn net.Conn)
type ConnectClose func(conn net.Conn, UUID uint64)

type TSTCP struct {
	conn net.Conn
}

func (this *TSTCP) SendBuffer(sBuffer string) {
	if this.conn == nil {
		return
	}
	buf := make([]byte, len(sBuffer)+4)
	binary.BigEndian.PutUint32(buf[:4], uint32(len(sBuffer)))
	copy(buf[4:], []byte(sBuffer))

	fmt.Println("Write:",buf)

	this.conn.Write(buf)
}

func (this *TSTCP) Create_Server(webpath string, init ServerInit, funCN ConnectNew, funRB ReceiveBuffer, funCC ConnectClose) {
	listener, err := net.Listen("tcp", webpath)
	if err != nil {
		fmt.Println("Tcp listen at port 9188 failed, err", err.Error())
		return
	}
	init()
	for {
		this.conn, err = listener.Accept()
		if err != nil {
			fmt.Println("Accept failed, err", err.Error())
			return
		}
		fmt.Println("Accept connect.")
		UUID := funCN(this.conn)
		this.SendBuffer("_RegistUUID," + TSUtil.ToString(int(UUID)));
		go tcpHandler(this.conn, funRB, funCC, UUID)
	}
}

func (this *TSTCP) Create_Client(webpath string, init ClientInit, funRB ReceiveBuffer, funCC ConnectClose) {
	conn, err := net.Dial("tcp", webpath)
	this.conn = conn
	if err != nil {
		fmt.Println("没有服务器连接:" + webpath)
		return
	}
	init(this.conn)
	go tcpHandler(this.conn, funRB, funCC, 0)
}

func tcpHandler(conn net.Conn, funRB ReceiveBuffer, funCC ConnectClose, UUID uint64) {
	var cache []byte = make([]byte, 32)
	var testLen uint32
	buf := bytes.NewBuffer(make([]byte, 0, 64))

	for {
		size, err := conn.Read(cache)
		if err != nil {
			funCC(conn, UUID)
			fmt.Printf("Read error, %v\n", err.Error())
			return
		}
		fmt.Printf("Read %v bytes: %v.\n", size, cache)

		buf.Write(cache[:size])
		fmt.Printf("Buffer: %v\n", buf.Bytes())

		for {
			if buf.Len() == 0 {
				testLen = 0
				break
			}

			if testLen == 0 {
				if buf.Len() < 4 {
					break
				}
				len := make([]byte, 4)
				_, err = buf.Read(len)
				testLen = binary.BigEndian.Uint32(len)
			}

			if int(testLen) > buf.Len() || testLen == 0 {
				break
			}

			data := make([]byte, testLen)
			_, err = buf.Read(data)

			fmt.Printf("content %v bytes: %v\n", testLen, string(data))

			ss := string(data)
			ssList := strings.Split(ss,",");
			fmt.Println("__RegistUUID : ", ssList[0])
			if ssList[0] == "__RegistUUID" {
				value,_ := strconv.Atoi(ssList[1]);
				UUID = uint64(value)
				fmt.Println("__RegistUUID : ", UUID)
			} else {
				funRB(conn, ss, UUID)
			}
			testLen = 0
		}
		time.Sleep(1000)
	}
}

