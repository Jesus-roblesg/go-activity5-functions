package main

import "fmt"
import "math/rand"

func mayorArray(args ...int) int { // args es un slice de los parámetros recibidos
	mayor := 0
	for _, v := range args { // iterar sobre los parámetros como un slice
		if mayor < v {
			mayor = v
		}
	}

	return mayor
}

func switchAB(a *int, b *int) {
	aux := *a
	*a = *b // *x dereferenciación/acceso a la dirección de memoria
	*b = aux
}

func generadorPares(n int64) int64 {
	// func (r *Rand) Int63() int64

	random := rand.Int63n(n)
	if random%2 != 0 {
		random = random * 2
	}

	return random
}

func main() {
	
	a := 1
	b := 2
	
	switchAB(&a, &b)
	fmt.Println(mayorArray(4, 6, 12, 8))
	fmt.Println(b)
	fmt.Println(a)
	fmt.Println(generadorPares(1000))
	
}