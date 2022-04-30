package main

import (
	"banking/timezone"
	"fmt"
	"time"
)

func main() {
	//app.Start()
	timezone.Start()
	loc, _ := time.LoadLocation("Asia/Shanghai")
	fmt.Println(time.Now().In(loc))

	loc, _ = time.LoadLocation("America/New_York")
	fmt.Println(time.Now().In(loc))

}
