package main

import (
	"fmt"
	"time"
)

// fungsi delay di golang
// bersifat blocking
func main() {
	fmt.Println("start")
	time.Sleep(time.Second * 3) // delay selama 3 detik
	fmt.Println("after 3 second")

	var timer = time.NewTimer(4 * time.Second) //time.NewTimer() mengembalikan objek bertipe *time.Timer yang memiliki property C yang bertipe channel.
	fmt.Println("start new timer")
	<-timer.C
	fmt.Println("finish new timer")
	var ttt string = string(104)
	fmt.Println(ttt)
}
