package main

import (
	"fmt"
	"time"
)

func main() {
	for range time.Tick(time.Millisecond * 1000) {
		fmt.Println(time.Now().Format("FileTime：2006,01,02,15:04:05"))
	}
}
