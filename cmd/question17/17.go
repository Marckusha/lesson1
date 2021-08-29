package main

import (
	"fmt"
)

/*
Написать программу, которая в рантайме способна определить тип переменной
— int, string, bool, channel из переменной типа interface{}.
*/

func varType(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Println("int", v)
	case string:
		fmt.Println("string", v)
	case bool:
		fmt.Println("bool", v)
	case chan interface{}:
		fmt.Println("channel", v)
	default:
		fmt.Println("not found")
	}
}

func main() {
	varType(123)
	varType("hello world")
	varType(true)
	varType(make(chan interface{}))
	varType(make(chan string))
}
