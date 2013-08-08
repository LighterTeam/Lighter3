/**
 * Created with IntelliJ IDEA.
 * User: Administrator
 * Date: 13-8-7
 * Time: 下午4:17
 * To change this template use File | Settings | File Templates.
 */
package main

import (
	"net"
	"fmt"
	"bytes"
	"encoding/binary"
	"time"
)

func ConnectNew(conn net.Conn) {

}

func ReceiveBuffer(conn net.Conn, sBuffer string) {

}

func SendBuffer(conn net.Conn, sBuffer string) {
	buf := make([]byte, len(sBuffer) + 4);
	copy(buf[:4],[]byte(len(sBuffer)))
	copy(buf[4:],[]byte(sBuffer))
	conn.Write(buf);
}

func GoAdaptServer() {
	listener, err := net.Listen("tcp", "localhost:9188")
	if err != nil {
		fmt.Println("Tcp listen at port 9188 failed, err", err.Error())
		return
	}
	for {
		var conn net.Conn
		conn, err = listener.Accept()
		if err != nil {
			fmt.Println("Accept failed, err", err.Error())
			return
		}
		fmt.Println("Accept connect.")
		ConnectNew(conn);
		go tcpHandler(conn)
	}
}
func tcpHandler(conn net.Conn) {
	var cache []byte = make([]byte, 32)
	var testLen uint32;
	buf := bytes.NewBuffer(make([]byte, 0, 1024));

	for {
		size, err := conn.Read(cache)
		if err != nil {
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
				len := make([]byte, 4)
				_, err = buf.Read(len)
				testLen = binary.BigEndian.Uint32(len)
			}

			if int(testLen) > buf.Len() || testLen == 0 {
				break
			}

			data := make([]byte, testLen)
			_, err = buf.Read(data)

			buf := data[4:]
			fmt.Printf("content %v bytes: %v\n", testLen, string(buf))
			ReceiveBuffer(conn, string(buf))
		}

		time.Sleep(1000)
	}
}


