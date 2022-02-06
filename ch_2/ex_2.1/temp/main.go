package main

import (
	"fmt"

	"temp/tempconv"
)

func main() {
	var c tempconv.Celsius = 100
	fmt.Println(c)
	fmt.Println(tempconv.AbsoluteZeroC)
	fmt.Println(tempconv.CToK(c))
}
