package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"time"
)

type Carro struct {
	Rodas          []Roda
	Estepe         Roda
	Velocidade     int
	Freio          bool
	Estado         EstadoCarro
	CintoMotorista bool
	PortaMalas     Bagageiro
}

type RodaPosicao int

const (
	DianteiraEsquerda = 0
	DianteiraDireita  = 1
	TraseiraEsquerda  = 2
	TraseiraDireita   = 3
	Estepe            = 4
)

type Roda struct {
	Posicao   int
	Estado    EstadoRoda
	Parafusos []bool
}

type EstadoCarro int

const (
	Ligado EstadoCarro = iota
	Desligado
)

type EstadoRoda int

const (
	Ok EstadoRoda = iota
	ComProblemas
)

type Bagageiro struct {
	Trancado bool
	Estepe   Roda
	Macaco   MacacoAutomotivo
	Chave    FerramentaDeRoda
}

type MacacoAutomotivo struct {
}

type FerramentaDeRoda struct {
}

func ImprimeCarro(carro *Carro) {
	raw, err := json.MarshalIndent(carro, "", "\t")

	if err != nil {
		panic(err)
	}

	fmt.Println(string(raw))
}

func CriaCarroComProblema() *Carro {
	carro := &Carro{
		Velocidade: 80,
		Freio:      false,
		Estado:     Ligado,
		Rodas:      make([]Roda, 4),
		Estepe: Roda{
			Estado:    Ok,
			Posicao:   Estepe,
			Parafusos: []bool{},
		},
	}

	for i := 0; i < 4; i++ {
		carro.Rodas[i] = Roda{
			Posicao:   i,
			Estado:    Ok,
			Parafusos: []bool{true, true, true, true},
		}
	}

	geradorDeNumeroAleatorio := rand.New(rand.NewSource(time.Now().UnixNano()))

	carro.Rodas[geradorDeNumeroAleatorio.Intn(4)].Estado = ComProblemas

	return carro
}

func SegueViagem(carro *Carro) {
	for _, roda := range carro.Rodas {
		if roda.Posicao != Estepe {
			if roda.Estado != Ok {
				panic(fmt.Errorf("Vc morreu! Sua roda %d está com problemas!", roda.Posicao))
			}

			if len(roda.Parafusos) < 3 {
				panic(fmt.Errorf("Vc morreu! Sua roda %d não está com os parafusos!", roda.Posicao))
			}
		}
	}

	if !carro.CintoMotorista {
		panic(fmt.Errorf("Vc está sem cinto de segurança!"))
	}

	if carro.Velocidade != 0 {
		panic(fmt.Errorf("O carro não está parado, não sei o que fazer!"))
	}

	if !carro.Freio {
		panic(fmt.Errorf("Seu carro não está com o freio acionado, não sei o que fazer!"))
	}
}

var carro *Carro

func main() {
	defer func() {
		if err := recover(); err != nil {

			log.Println("panic occurred:", err)
		}

		ImprimeCarro(carro)
	}()

	carro = CriaCarroComProblema()

	SegueViagem(carro)
}
