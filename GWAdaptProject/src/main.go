/**
 * Created with IntelliJ IDEA.
 * User: Administrator
 * Date: 13-8-7
 * Time: 下午2:58
 * To change this template use File | Settings | File Templates.
 */
package main

import (
	"time"
)

func main() {
	go GoAdaptServer()
	go GoClientServer()

	for {
		time.Sleep(1000)
	}
}
