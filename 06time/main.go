package main

import (
	"fmt"
	"time"
)

func main() {

	presentTime := time.Now()
	fmt.Println(presentTime)
	fmt.Println(presentTime.Format("01-05-2022 15:04:14 Monday"))
}
