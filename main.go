package main

import (
	"fmt"

	"myproject/greet"
	"myproject/utils"
)

func main() {
	fmt.Println(utils.Add(2, 3))
	fmt.Println(greet.Hello("Ali"))
}
