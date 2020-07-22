package main

import (
	"log"

	"learnops/internal/car"
)

var viagem *car.Viagem

func main() {
	defer func() {
		if err := recover(); err != nil {
			log.Println("panic occurred:", err)
			car.Imprime(viagem)
		}
	}()

	viagem = car.CriaViagem()

	err := viagem.Carro.Acelera(120)

	if err != nil {
		panic(err)
	}

	car.Imprime(viagem)
}
