package main

import (
	"fmt"
	"time"
)

func main() {
	arr := [...]int{0, 11, 222, 333, 44444}
	for key := range arr {
		fmt.Println(arr[key])
		time.Sleep(1 * time.Second)
	}
}
