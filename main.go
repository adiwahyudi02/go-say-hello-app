package main

import (
	"fmt"

	gosayhello "github.com/adiwahyudi02/go-say-hello-module"
)

func main() {
	hello := gosayhello.SayHello()
	fmt.Println(hello)
}
