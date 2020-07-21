package main

import (
	"log"

	"learnops/internal/car"
)

var carro *car.Carro

func main() {
	defer func() {
		if err := recover(); err != nil {
			log.Println("panic occurred:", err)
			car.Imprime(carro)
		}
	}()

	carro = car.CriaCarroComProblema()

	err := carro.Acelera(120)

	if err != nil {
		panic(err)
	}

	car.Imprime(carro)
}
