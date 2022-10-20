package main

import "fmt"

func main() {
	//bool, string, numeric
	/*
		var a bool = true
		var a string = "SotecMovil"
		var a uint8 = 100
		var a byte = 100
		var a rune = 100
		var a rune = '|'
		var a float64 = 23.56
	*/
	var a uint8 = 100
	var b uint16 = 23000
	//se castea la variable con unit16(numero)
	c := uint16(a) + b
	_ = 234
	var _ string = "test"
	var d bool

	fmt.Printf("Tipo:%T, Valor:%v \n", c, c)
	fmt.Printf("Tipo:%T, Valor:%v \n", d, d)
}
