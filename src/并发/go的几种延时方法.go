package main

import "time"

func main() {
	// case 1
	time.Sleep(time.Second)
	//case 2
	<-time.After(time.Second) //return a channel
	//case 3
	timer := time.NewTimer(time.Second)
	<-timer.C //可读取

}
