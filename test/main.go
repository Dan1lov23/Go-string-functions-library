package main

import (
	"fmt"
	"func/strFunc"
)

func main() {
	fmt.Println(strFunc.ReverseStr("golang"))
	fmt.Println(strFunc.AddElementInString("golang", 3, "i"))
	fmt.Println(strFunc.DeleteElementInString("golang", 2))
}
