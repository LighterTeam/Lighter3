// GWProject
package main

import "time"

func main() {
	go GoClientServer()
	go GoGameAdaptClient()
	for {
		time.Sleep(1000)
	}
}
