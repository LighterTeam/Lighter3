// GWProject
package main

import "time"

func main() {
	go GoGameAdaptServer()
	for {
		time.Sleep(1000)
	}
}
