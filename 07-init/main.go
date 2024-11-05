package main

import (
	"fmt"
	"strconv"
)

var ix = 0

func init() {
	ix++
	fmt.Println("init関数" + strconv.Itoa(ix))
}
func main() {
	ix++
	fmt.Println("main関数" + strconv.Itoa(ix))
}
