package main

//el paquete flag, nos ayuda a programar con la lína de comandos
import (
	"flag"
	"fmt"
	"os"
)

func main() {
	//filte pattern (ordena todas las entradas)
	//flagPattern := flag.String("p", "", flag.Arg(0))
	//flagAll := flag.Bool("a", false, "all file including hide files")
	//flagNumberRecords := flag.Int("n", 0, "number of records")

	//order flags (ordenan la salida)
	//hasOrderByTime := flag.Bool("t", false, "sort by time, oldest first")
	//hasOrderBySize := flag.Bool("s", false, "sort by file size, smallest first")
	//hasOrderByReverse := flag.Bool("r", false, "reverse order while sorting")

	//fmt.Println("pattern: ", *flagPattern)
	//fmt.Println("all: ", *flagAll)
	//fmt.Println("flagNumberRecords: ", *flagNumberRecords)
	//fmt.Println("hasOrderByTime: ", *hasOrderByTime)
	//fmt.Println("hasOrderBySize: ", *hasOrderBySize)
	//fmt.Println("hasOrderByReverse: ", *hasOrderByReverse)

	flag.Parse() //--> Ejecuta el análisis gramatical, es como si ejecutara los flags
	path := flag.Arg(0)
	if path == "" {
		path = "."
	}
	dirs, err := os.ReadDir(path)
	if err != nil {
		panic(err)
	}

	for _, dir := range dirs {
		fmt.Println(dir.Name())
	}
	/*
		fmt.Println(flag.Args()) args regresa un slice de strings, de todos los argentos después  escribir go run main.go argumento1 argumento2.. etc. Args toma como argumentos el string separado por espacios
		fmt.Println(flag.Arg(0)) arg le el string del argumento que está en la posición cero porque solo indicamos en el ()
	*/

}
