package main

import "fmt"

func main() {
	var a = 4 + 2*3
	var b = 10
	b += 2
	fmt.Println(a)
	fmt.Println(b)
	var c = 3
	c--
	fmt.Println(c)
	//Operadores de Comparación: >, <, ==, !=, >=.<=
	fmt.Println(b != a)
	//Operadores Lógicos: &&,||
	var age = 30
	fmt.Println("age:", age >= 18 && age < 60)
	fmt.Println("niño o viejo:", age < 18 || age >= 60)
	//Unario : !
	fmt.Println(!(4 != 4))
}
