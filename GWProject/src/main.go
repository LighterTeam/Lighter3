// GWProject
package main

import "time"

func main() {
	go GoGWAdaptClient()
	go GoGameAdaptClient()
	go GoClientServer()
	go GoGameServer()

	for {
		time.Sleep(1000)
	}
}
