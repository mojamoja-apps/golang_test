package main

import (
	"fmt"
)

func main() {
	pl("1")
	defer pl("2") // ここが遅延実行される
	pl("3")
}

func pl(str string) {
	fmt.Println(str)
}
