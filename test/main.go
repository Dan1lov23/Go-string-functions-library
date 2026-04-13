package main

import (
	"fmt"
	"func/strFunc"
)

func main() {
	fmt.Println(strFunc.ReverseStr("golang"))
	fmt.Println(strFunc.AddElement("golang", 3, "i"))
	fmt.Println(strFunc.DeleteElement("golang", 2))
	fmt.Println(strFunc.ReplaceElement("golang", 5, "d"))
	fmt.Println(strFunc.CheckSubStringInclude("golang", "go"))
}

