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

	carro.Acelera(120)

	velocidade, err := carro.Freia()

	for velocidade > 0 {
		velocidade, err = carro.Freia()
		if err != nil {
			panic(err)
		}
		log.Printf("nova velocidade: %d", velocidade)
	}

	car.Imprime(carro)
}
