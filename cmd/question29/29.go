package main

import (
	"fmt"
	"math"
	"math/big"
)

/*
Написать программу, которая перемножает, делит, складывает, вычитает 2
числовых переменных a,b, значение которые > 2^20
*/

func main() {
	var (
		a int64 = -math.MaxInt64
		b int64 = math.MaxInt64
	)

	bigA := big.NewInt(a)
	bigB := big.NewInt(b)

	result := new(big.Int)

	fmt.Println("mult: ", result.Mul(bigA, bigB))
	fmt.Println("div: ", result.Div(bigA, bigB))
	fmt.Println("sub: ", result.Sub(bigA, bigB))
	fmt.Println("add: ", result.Add(bigA, bigB))
}
