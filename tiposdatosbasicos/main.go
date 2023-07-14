package main

//Estos son los paquetes de go
import "fmt"

func main() {
	//bool, string, numeric
	//En string simpre llevan comilla doble porque una comilla significa otra cosa en GO
	//numeric se subdivide en varios tipos de datos numéricos
	//el dato uint es igual al tipo de arquitectura del S.O ej. mi arquitectura es de 64 y se escribo uint solo, será igual a uint64
	//lo mismo ocurre si usolo solo int
	/*
		var a bool = true
		var a string = "SotecMovil"
		var a uint8 = 100
		var a byte = 100
		var a rune = 100
		//la letra que coloques dentro de las comillas con rune, es para que nos muestre el unicoe del caracter que escribimos
		var a rune = '|'
		var a float64 = 23.56
	*/
	var a uint8 = 100
	var b uint16 = 23000

	//no puedo operar los tipos de datos int 16 con un int 32, tengo que castear para hacerlo
	//se castea la variable con unit16(numero)
	c := uint16(a) + b

	//Identificador blank = "_" es para usarlo cuando no nececito imprimir una salida y que go marque un error
	_ = 234
	var _ string = "test"
	var d bool
	var entero int = 555

	//la función printf, es para imprimir con formato, porque quiero saber el tipo de la variable y su valor
	//En el siguiente ejemplo se usan los verbos
	//Los valores son todos los símboloes que usan el prefijo "%" + una letra que significa algo
	//Lo puedo ver en la documentación de GO los valores de la función printf
	//%T = es el tipo de dato
	//%v= es el valor de la variable por default
	fmt.Printf("Tipo:%T, Valor:%v \n", c, c)
	fmt.Printf("Tipo:%T, Valor:%v \n", d, d)
	fmt.Printf("Tipo:%T, Valor:%v \n", entero, entero)

	//Valores ceros
	//Es cuando declaro una variable y no le asigno un valor
	//los valores ceros, cambian dependiendo el tipo de dato
	var z string // para los string el valor cero es vacio
	var q byte   //para cualquier entero el valor cero siempre sera cero
	var w bool   //para el bool el valor cero es false
	//el verbo %q es para que imprima el valor entre comillas de los string
	fmt.Printf("Tipo: %T, Valor %q \n", z, z)
	fmt.Printf("Tipo: %T, Valor %v \n", q, q)
	fmt.Printf("Tipo: %T, Valor %v \n", w, w)
}
