package main

import "fmt"

func main() {
	//Operadores aritméticos: (), *, /, %, +, -
	//Operadores de asignación: =, +=, -=, /=, %=
	var a = 4 + 2*3
	var b = 10
	b += 2
	fmt.Println(a)
	fmt.Println(b)
	//Declaraciones post-incremento y post-decremento: ++, --
	//No son una expresión, sino una declaración
	//Forma correcta
	var c = 3
	c--
	//Forma incorrecta
	/*
		var c int = 5
		c= c++
	*/
	fmt.Println(c)
	//Operadores de Comparación: >, <, ==, !=, >=.<=. Respuesta booleana
	fmt.Println(b != a)
	//Operadores Lógicos: &&,||
	var age = 30
	fmt.Println("Es adulto?:", age >= 18 && age < 60)
	fmt.Println("Es niño o viejo?:", age < 18 || age >= 60)
	//Unario : ! (negación)
	fmt.Println(!(4 == 4))
}
