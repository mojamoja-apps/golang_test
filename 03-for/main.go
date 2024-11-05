package main

import (
	"fmt"
	"time"
)

func main() {
	for ix := 0; ix < 5; ix++ {
		fmt.Println(ix)
		time.Sleep(1 * time.Second)
	}
}
