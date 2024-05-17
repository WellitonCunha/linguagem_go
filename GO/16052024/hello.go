package main

import (
	"fmt"
	"reflect"
)

func main() {
	var nome string = "Welliton"
	idade := 32
	var altura float32 = 1.70

	fmt.Println("Meu nome é:", reflect.TypeOf(nome))
	fmt.Println("Minha idade é do tipo:", reflect.TypeOf(idade))
	fmt.Println("Minha altura é do tipo:", reflect.TypeOf(altura))
}
