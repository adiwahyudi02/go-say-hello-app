package main

import (
	"fmt"

	gosayhello "github.com/adiwahyudi02/go-say-hello-module/v2"
)

func main() {
	hello := gosayhello.SayHello("Adi Wahyudi")
	fmt.Println(hello)
}
