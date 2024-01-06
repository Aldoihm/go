package main

import "fmt"

type Product[T uint | string] struct {
	ID          T
	Description string
	Price       float64
}

func main() {
	product1 := Product[uint]{ID: 1, Description: "ZAPATOS", Price: 30}
	product2 := Product[string]{ID: "e049c10d-4503-4825-9783-dbffc17f2583", Description: "CAMISA", Price: 40}
	fmt.Println(product1)
	fmt.Println(product2)
}
