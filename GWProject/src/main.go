// GWProject
package main

import "time"

func main() {
	go GoGWClient()
	for {
		time.Sleep(1000)
	}
}
