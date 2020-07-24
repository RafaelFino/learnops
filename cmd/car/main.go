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

	errs := viagem.TestaViagem()

	if len(errs) > 0 {
		panic("deu ruim!")
	}
}
