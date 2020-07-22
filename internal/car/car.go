package car

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"
)

type RodaPosicao string

const (
	PosicaoDianteiraEsquerda RodaPosicao = "dianteira-esquerda"
	PosicaoDianteiraDireita  RodaPosicao = "dianteira-direita"
	PosicaoTraseiraEsquerda  RodaPosicao = "traseira-esquerda"
	PosicaoTraseiraDireita   RodaPosicao = "traseira-direita"
	PosicaoEstepe            RodaPosicao = "estepe"
	PosicaoNenhum            RodaPosicao = "nenhum"
)

type Viagem struct {
	Carro   *Carro
	Estrada *Via
}

type Faixa int

type Via struct {
	QuantidadeDeFaixas int
	FaixaAtual         Faixa
	VelocidadeMaxima   int
	VelocidadeMinima   int
}

type Carro struct {
	Rodas             map[RodaPosicao]*Roda
	Estepe            *Roda
	Velocidade        int
	Estado            EstadoCarro
	CintoMotorista    bool
	EstadoSeta        EstadoSeta
	FreioDeMao        bool
	EstadoPiscaAlerta bool
	PortaMotorita     EstadoPorta
	PortaMalas        EstadoPorta
	PosicaoSuspenso   RodaPosicao
	Macaco            *MacacoAutomotivo
	ChaveDeRoda       *FerramentaDeRoda
	Chave             *ChaveDoCarro
	DentroDoCarro     bool
	Seta              EstadoSeta
	CarLog            []string
}

type EstadoPorta string

const (
	PortaAberta  EstadoPorta = "aberta"
	PortaFechada EstadoPorta = "fechada"
)

type Roda struct {
	Estado    EstadoRoda
	Parafusos int
}

type EstadoCarro string

const (
	CarroLigado    EstadoCarro = "ligado"
	CarroDesligado             = "desligado"
)

type EstadoRoda string

const (
	RodaOk           EstadoRoda = "ok"
	RodaComProblemas            = "problema"
)

type Parafuso bool

type ChaveDoCarro struct {
}

type MacacoAutomotivo struct {
}

type FerramentaDeRoda struct {
}

type EstadoSeta string

const (
	SetaDesligada EstadoSeta = "desligada"
	SetaEsquerda  EstadoSeta = "esquerda"
	SetaDireita   EstadoSeta = "direita"
)

func Imprime(obj interface{}) {
	raw, err := json.MarshalIndent(obj, "", "\t")

	if err != nil {
		panic(err)
	}

	fmt.Println(string(raw))
}

func CriaViagem() *Viagem {
	carro := &Carro{
		Velocidade: 80,
		FreioDeMao: false,
		Estado:     CarroLigado,
		Estepe: &Roda{
			Estado:    RodaOk,
			Parafusos: 0,
		},

		CintoMotorista:    true,
		EstadoSeta:        SetaDesligada,
		PortaMotorita:     PortaFechada,
		PortaMalas:        PortaFechada,
		PosicaoSuspenso:   PosicaoNenhum,
		Macaco:            &MacacoAutomotivo{},
		ChaveDeRoda:       &FerramentaDeRoda{},
		Chave:             &ChaveDoCarro{},
		EstadoPiscaAlerta: false,
		DentroDoCarro:     true,
		Seta:              SetaDesligada,
	}

	carro.Rodas = map[RodaPosicao]*Roda{
		PosicaoDianteiraEsquerda: &Roda{
			Estado:    RodaOk,
			Parafusos: 4,
		},
		PosicaoDianteiraDireita: &Roda{

			Estado:    RodaOk,
			Parafusos: 4,
		},
		PosicaoTraseiraEsquerda: &Roda{
			Estado:    RodaOk,
			Parafusos: 4,
		},
		PosicaoTraseiraDireita: &Roda{
			Estado:    RodaOk,
			Parafusos: 4,
		},
	}

	geradorDeNumeroAleatorio := rand.New(rand.NewSource(time.Now().UnixNano()))

	switch geradorDeNumeroAleatorio.Intn(4) {
	case 0:
		carro.Rodas[PosicaoDianteiraEsquerda].Estado = RodaComProblemas
	case 1:
		carro.Rodas[PosicaoDianteiraDireita].Estado = RodaComProblemas
	case 2:
		carro.Rodas[PosicaoTraseiraEsquerda].Estado = RodaComProblemas
	case 3:
		carro.Rodas[PosicaoTraseiraDireita].Estado = RodaComProblemas
	}

	qtd := 1 + geradorDeNumeroAleatorio.Intn(2)
	estrada := &Via{
		QuantidadeDeFaixas: qtd,
		FaixaAtual:         Faixa(geradorDeNumeroAleatorio.Intn(qtd)),
		VelocidadeMaxima:   80 + geradorDeNumeroAleatorio.Intn(4)*10,
		VelocidadeMinima:   80,
	}

	return &Viagem{
		Carro:   carro,
		Estrada: estrada,
	}
}

func (v *Viagem) TestaViagem() []error {
	var ret []string

	if v.Carro.Velocidade == 0 {
		ret = append(ret, "Vc está parado, assim não vai chegar em lugar algum")
	}

	if v.Carro.Velocidade < v.Estrada.VelocidadeMinima {
		ret = append(ret, "está tudo bem? sua velocidade está abaixo do permitido")
	}

	if v.Carro.Velocidade < v.Estrada.VelocidadeMinima {
		ret = append(ret, "tá com pressa? sua velocidade está acima do permitido")
	}

	if len(ret) > 0 {
		return v.Carro.LogErr("TestaViagem", ret)
	}

	v.Carro.Log("TestaViagem", "ok")
	return nil
}

func (v *Viagem) EhSeguroParar() bool {
	return v.Estrada.FaixaAtual == 0
}

func (v *Viagem) MudaFaixa(novaFaixa Faixa) []error {
	var ret []string

	if v.Carro.Velocidade == 0 {
		ret = append(ret, "vc está parado, assim não vai chegar em lugar algum")
	}

	if novaFaixa < 0 || int(novaFaixa) > v.Estrada.QuantidadeDeFaixas {
		ret = append(ret, "essa faixa não existe")
	}

	if novaFaixa > v.Estrada.FaixaAtual && (novaFaixa != v.Estrada.FaixaAtual+1) {
		ret = append(ret, "vc não pode pular mais de uma faixa por vez")
	}

	if novaFaixa < v.Estrada.FaixaAtual && (novaFaixa != v.Estrada.FaixaAtual-1) {
		ret = append(ret, "vc não pode pular mais de uma faixa por vez")
	}

	if len(ret) > 0 {
		return v.Carro.LogErr("MudaFaixa", ret)
	} else {
		v.Estrada.FaixaAtual = novaFaixa
	}

	v.Carro.Log("MudaFaixa", "ok")
	return nil
}

func (c *Carro) LogErr(source string, errs []string) []error {
	var ret []error
	for i, e := range errs {
		ret = append(ret, fmt.Errorf("[err: %d] %s", i, e))
		c.CarLog = append(c.CarLog, fmt.Sprintf("[%s] [%s] Error: %s", time.Now().Format(time.RFC3339), source, e))
	}

	return ret
}

func (c *Carro) Log(source string, msg string) {
	c.CarLog = append(c.CarLog, fmt.Sprintf("[%s] [%s] %s", time.Now().Format(time.RFC3339), source, msg))
}

func (c *Carro) Liga(chave *ChaveDoCarro) []error {
	var ret []string

	c.Chave = chave

	if c.Estado == CarroLigado {
		ret = append(ret, "o carro já está ligado")
	}

	if c.Velocidade > 0 {
		ret = append(ret, "o carro já está andando, você tem certeza de que está certo do que está fazendo?")
	}

	if c.Chave == nil {
		ret = append(ret, "chave não está no contato, me entregue uma chave válida!")
	}

	if c.EstadoPiscaAlerta {
		ret = append(ret, "o pisca alerta está ligado")
	}

	if len(c.Rodas) != 4 {
		ret = append(ret, "Alguma roda está faltando")
	}

	if c.PosicaoSuspenso != PosicaoNenhum {
		ret = append(ret, "Seu carro está suspenso, preste atenção!")
	}

	for k, r := range c.Rodas {
		if r.Estado != RodaOk {
			ret = append(ret, fmt.Sprintf("a roda %s não está ok!", k))
		}

		if r.Parafusos < 4 {
			ret = append(ret, fmt.Sprintf("a roda %s não está com a quantidade correta de parafusos!", k))
		}
	}

	if c.PortaMotorita == PortaAberta {
		ret = append(ret, "a porta do motorista está aberta")
	}

	if c.PortaMalas == PortaAberta {
		ret = append(ret, "a porta do porta-malas está aberta")
	}

	if !c.FreioDeMao {
		ret = append(ret, "o freio de mão não está puxado, isso é perigoso!")
	}

	if !c.DentroDoCarro {
		ret = append(ret, "Não é carrinho de controle remoto, vc precisa estar dentro dele para isso")
	}

	if c.Macaco == nil {
		ret = append(ret, "Cadêo o seu macaco?")
	}

	if c.ChaveDeRoda == nil {
		ret = append(ret, "Cadêo a sua chave de rodas?")
	}

	if len(ret) > 0 {
		return c.LogErr("Liga", ret)
	} else {
		c.Estado = CarroLigado
	}

	c.Log("Liga", "ok")
	return nil
}

func (c *Carro) Desliga() (*ChaveDoCarro, []error) {
	var ret []string
	var chave *ChaveDoCarro

	if c.Chave == nil {
		ret = append(ret, "chave não está no contato, talvez vc a tenha perdido")
	} else {
		chave = c.Chave
		c.Chave = nil
	}

	if c.Estado == CarroDesligado {
		ret = append(ret, "o carro já está deligado")
	}

	if c.Velocidade > 0 {
		ret = append(ret, "o carro está andando, vai dar merda!")
	}

	if len(c.Rodas) != 4 {
		ret = append(ret, "Alguma roda está faltando, você bebeu antes de dirigir?")
	}

	if c.PosicaoSuspenso != PosicaoNenhum {
		ret = append(ret, "Seu carro está suspenso, vc está bem?")
	}

	for k, r := range c.Rodas {
		if r.Parafusos < 4 {
			ret = append(ret, fmt.Sprintf("a roda %s não está com a quantidade correta de parafusos! Vc deixou aquele seu primo mecânico mexer de novo?", k))
		}
	}

	if c.PortaMotorita == PortaAberta {
		ret = append(ret, "a porta do motorista está aberta, pega leve que isso aqui não é circo")
	}

	if c.PortaMalas == PortaAberta {
		ret = append(ret, "a porta do porta-malas está aberta, está levando algo muito grande no carro?")
	}

	if !c.FreioDeMao {
		ret = append(ret, "o freio de mão não está puxado, não é assim que se para com segurança")
	}

	if !c.DentroDoCarro {
		ret = append(ret, "Não é carrinho de controle remoto, vc precisa estar dentro dele para isso")
	}

	if len(ret) > 0 {
		return chave, c.LogErr("Desliga", ret)
	}

	c.Log("Desliga", "ok")
	return chave, nil
}

func (c *Carro) Freia() (int, []error) {
	var ret []string

	if c.Estado == CarroDesligado {
		ret = append(ret, "o carro já está deligado")
	}

	if len(c.Rodas) != 4 {
		ret = append(ret, "Alguma roda está faltando, você bebeu antes de dirigir?")
	}

	if c.PosicaoSuspenso != PosicaoNenhum {
		ret = append(ret, "Seu carro está suspenso, vc está bem?")
	}

	for k, r := range c.Rodas {
		if r.Parafusos < 4 {
			ret = append(ret, fmt.Sprintf("a roda %s não está com a quantidade correta de parafusos! Vc deixou aquele seu primo mecânico mexer de novo?", k))
		}
	}

	if c.PortaMotorita == PortaAberta {
		ret = append(ret, "a porta do motorista está aberta, pega leve que isso aqui não é circo")
	}

	if c.PortaMalas == PortaAberta {
		ret = append(ret, "a porta do porta-malas está aberta, está levando algo muito grande no carro?")
	}

	if c.FreioDeMao {
		ret = append(ret, "o freio de mão está puxado, não é assim que se para com segurança")
	}

	if !c.DentroDoCarro {
		ret = append(ret, "Não é carrinho de controle remoto, vc precisa estar dentro dele para isso")
	}

	if len(ret) > 0 {
		return c.Velocidade, c.LogErr("Freia", ret)
	} else {
		if c.Velocidade > 0 {
			c.Velocidade -= 10
		}

		if c.Velocidade < 0 {
			c.Velocidade = 0
		}
	}

	c.Log("Freia", "ok")
	return c.Velocidade, nil
}

func (c *Carro) Acelera(novaVelocidade int) []error {
	var ret []string

	if c.Estado != CarroLigado {
		ret = append(ret, "o carro não está ligado")
	}

	if c.EstadoPiscaAlerta {
		ret = append(ret, "o pisca alerta está ligado, você quer mesmo acelerar?")
	}

	if len(c.Rodas) != 4 {
		ret = append(ret, "Alguma roda está faltando, vc não deveria dirigir")
	}

	if c.PosicaoSuspenso != PosicaoNenhum {
		ret = append(ret, "Seu carro está suspenso, preste atenção!")
	}

	for k, r := range c.Rodas {
		if r.Estado != RodaOk {
			ret = append(ret, fmt.Sprintf("a roda %s não está ok!", k))
		}

		if r.Parafusos < 4 {
			ret = append(ret, fmt.Sprintf("a roda %s não está com a quantidade correta de parafusos!", k))
		}
	}

	if c.PortaMotorita == PortaAberta {
		ret = append(ret, "a porta do motorista está aberta")
	}

	if c.PortaMalas == PortaAberta {
		ret = append(ret, "a porta do porta-malas está aberta, anda levando algo muito grande lá atrás?")
	}

	if c.FreioDeMao {
		ret = append(ret, "o freio de mão está puxado, vc sabe oq vc está fazendo?")
	}

	if c.Velocidade > novaVelocidade {
		ret = append(ret, "Se você quer diminuir, é melhor pisar no freio")
	}

	if !c.DentroDoCarro {
		ret = append(ret, "Não é carrinho de controle remoto, vc precisa estar dentro dele para isso")
	}

	if c.Macaco == nil {
		ret = append(ret, "Cadêo o seu macaco?")
	}

	if c.ChaveDeRoda == nil {
		ret = append(ret, "Cadêo a sua chave de rodas?")
	}

	if len(ret) > 0 {
		return c.LogErr("Acelera", ret)
	} else {
		c.Velocidade = novaVelocidade
	}

	c.Log("Acelera", "ok")
	return nil
}

func (c *Carro) AcionaFreioDeMao(aciona bool) []error {
	var ret []string

	if c.Estado == CarroLigado && aciona {
		ret = append(ret, "o carro está ligado, você tem certeza?")
	}

	if c.Velocidade > 0 && aciona {
		ret = append(ret, "o carro está andando, vc vai fazer bosta")
	}

	for k, r := range c.Rodas {
		if r.Parafusos < 4 {
			ret = append(ret, fmt.Sprintf("a roda %s não está com a quantidade correta de parafusos! Vc deixou aquele seu primo mecânico mexer de novo?", k))
		}
	}

	if len(c.Rodas) != 4 && !aciona {
		ret = append(ret, "Alguma roda está faltando, vc quer soltar isso?")
	}

	if c.PosicaoSuspenso != PosicaoNenhum && !aciona {
		ret = append(ret, "Seu carro está suspenso, preste atenção!")
	}

	if !c.DentroDoCarro {
		ret = append(ret, "Não é carrinho de controle remoto, vc precisa estar dentro dele para isso")
	}

	if len(ret) > 0 {
		return c.LogErr("AcionaFreioDeMao", ret)
	} else {
		c.FreioDeMao = aciona
	}

	c.Log("AcionaFreioDeMao", "ok")
	return nil
}

func (c *Carro) SaiDoCarro() []error {
	var ret []string

	if c.Estado == CarroLigado {
		ret = append(ret, "o carro está ligado, você tem certeza?")
	}

	if c.Velocidade > 0 {
		ret = append(ret, "o carro está andando, vc vai fazer bosta")
	}

	if c.PosicaoSuspenso != PosicaoNenhum {
		ret = append(ret, "Seu carro está suspenso, preste atenção!")
	}

	if c.PortaMotorita != PortaAberta {
		ret = append(ret, "Sua porta está fechada, vai sair pela janela?")
	}

	if len(ret) > 0 {
		return c.LogErr("SaiDoCarro", ret)
	} else {
		c.DentroDoCarro = false
	}

	c.Log("SaiDoCarro", "ok")
	return nil
}

func (c *Carro) EntraNoCarro() []error {
	var ret []string

	if c.Velocidade > 0 {
		ret = append(ret, "o carro está andando, vc vai fazer bosta")
	}

	if c.PortaMotorita != PortaAberta {
		ret = append(ret, "Sua porta está fechada, vai entrar pela janela?")
	}

	if len(ret) > 0 {
		return c.LogErr("EntraNoCarro", ret)
	} else {
		c.DentroDoCarro = true
	}

	c.Log("EntraNoCarro", "ok")
	return nil
}

func (c *Carro) LigaSeta(novoEstado EstadoSeta) []error {
	var ret []string

	if !c.DentroDoCarro {
		ret = append(ret, "Não é carrinho de controle remoto, vc precisa estar dentro dele para isso")
	}

	if len(ret) > 0 {
		return c.LogErr("LigaSeta", ret)
	} else {
		c.Seta = novoEstado
	}

	c.Log("LigaSeta", "ok")
	return nil
}

func (c *Carro) MudaEstadoPortaMotorista(novoEstado EstadoPorta) []error {
	var ret []string

	if c.Estado == CarroLigado {
		ret = append(ret, "o carro está ligado, você tem certeza?")
	}

	if c.Velocidade > 0 {
		ret = append(ret, "o carro está andando, vc vai fazer bosta")
	}

	if !c.FreioDeMao {
		ret = append(ret, "o carro não está com o freio de mão acionado, tá cagando de novo")
	}

	for k, r := range c.Rodas {
		if r.Parafusos < 4 {
			ret = append(ret, fmt.Sprintf("a roda %s não está com a quantidade correta de parafusos! Vc deixou aquele seu primo mecânico mexer de novo?", k))
		}
	}

	if len(ret) > 0 {
		return c.LogErr("MudaEstadoPortaMotorista", ret)
	} else {
		c.PortaMotorita = novoEstado
	}

	c.Log("MudaEstadoPortaMotorista", "ok")
	return nil
}

func (c *Carro) MudaEstadoPortaMalas(chave *ChaveDoCarro, novoEstado EstadoPorta) []error {
	var ret []string

	if chave == nil {
		ret = append(ret, "Esse aqui só abre com a chave")
	}

	if c.Chave != nil {
		ret = append(ret, "A chave não está no contato ainda?")
	}

	if c.Estado == CarroLigado {
		ret = append(ret, "o carro está ligado, você tem certeza?")
	}

	if c.Velocidade > 0 {
		ret = append(ret, "o carro está andando, vc vai fazer bosta")
	}

	if !c.FreioDeMao {
		ret = append(ret, "o carro não está com o freio de mão acionado, tá cagando de novo")
	}

	for k, r := range c.Rodas {
		if r.Parafusos < 4 {
			ret = append(ret, fmt.Sprintf("a roda %s não está com a quantidade correta de parafusos! Vc deixou aquele seu primo mecânico mexer de novo?", k))
		}
	}

	if c.DentroDoCarro {
		ret = append(ret, "Esse é um carro simples, vc precisa estar fora do carro para fazer isso")
	}

	if len(ret) > 0 {
		return c.LogErr("MudaEstadoPortaMalas", ret)
	} else {
		c.PortaMalas = novoEstado
	}

	c.Log("MudaEstadoPortaMalas", "ok")
	return nil
}

func (c *Carro) PegaMacaco() (*MacacoAutomotivo, []error) {
	var ret []string
	var macaco *MacacoAutomotivo

	if c.Estado == CarroLigado {
		ret = append(ret, "o carro está ligado, você tem certeza?")
	}

	if c.Velocidade > 0 {
		ret = append(ret, "o carro está andando, vc vai fazer bosta")
	}

	if !c.FreioDeMao {
		ret = append(ret, "o carro não está com o freio de mão acionado, tá cagando de novo")
	}

	if c.PortaMalas != PortaAberta {
		ret = append(ret, "o porta malas não está aberto, quer quebrar o vidro?")
	}

	if c.DentroDoCarro {
		ret = append(ret, "Esse é um carro simples, vc precisa estar fora do carro para fazer isso")
	}

	if len(ret) > 0 {
		return macaco, c.LogErr("PegaMacaco", ret)
	} else {
		macaco = c.Macaco
		c.Macaco = nil
	}

	c.Log("PegaMacaco", "ok")
	return macaco, nil
}

func (c *Carro) GuardaMacaco(macaco *MacacoAutomotivo) []error {
	var ret []string

	if macaco == nil {
		ret = append(ret, "Cadê o macaco? Tá doido?")
	}

	if c.Estado == CarroLigado {
		ret = append(ret, "o carro está ligado, você tem certeza?")
	}

	if c.Velocidade > 0 {
		ret = append(ret, "o carro está andando, vc vai fazer bosta")
	}

	if !c.FreioDeMao {
		ret = append(ret, "o carro não está com o freio de mão acionado, tá cagando de novo")
	}

	if c.PortaMalas != PortaAberta {
		ret = append(ret, "o porta malas não está aberto, quer quebrar o vidro?")
	}

	if c.DentroDoCarro {
		ret = append(ret, "Esse é um carro simples, vc precisa estar fora do carro para fazer isso")
	}

	if len(ret) > 0 {
		return c.LogErr("GuardaMacaco", ret)
	} else {
		c.Macaco = macaco
	}

	c.Log("GuardaMacaco", "ok")
	return nil
}

func (c *Carro) PegaChaveDeRodas() (*FerramentaDeRoda, []error) {
	var ret []string
	var chave *FerramentaDeRoda

	if c.Estado == CarroLigado {
		ret = append(ret, "o carro está ligado, você tem certeza?")
	}

	if c.Velocidade > 0 {
		ret = append(ret, "o carro está andando, vc vai fazer bosta")
	}

	if !c.FreioDeMao {
		ret = append(ret, "o carro não está com o freio de mão acionado, tá cagando de novo")
	}

	if c.PortaMalas != PortaAberta {
		ret = append(ret, "o porta malas não está aberto, quer quebrar o vidro?")
	}

	if c.DentroDoCarro {
		ret = append(ret, "Esse é um carro simples, vc precisa estar fora do carro para fazer isso")
	}

	if len(ret) > 0 {
		return chave, c.LogErr("PegaChaveDeRodas", ret)
	} else {
		chave = c.ChaveDeRoda
		c.ChaveDeRoda = nil
	}

	c.Log("PegaChaveDeRodas", "ok")
	return chave, nil
}

func (c *Carro) GuardaChaveDeRodas(chave *FerramentaDeRoda) []error {
	var ret []string

	if chave == nil {
		ret = append(ret, "Cadê a chave? Tá doido?")
	}

	if c.Estado == CarroLigado {
		ret = append(ret, "o carro está ligado, você tem certeza?")
	}

	if c.Velocidade > 0 {
		ret = append(ret, "o carro está andando, vc vai fazer bosta")
	}

	if !c.FreioDeMao {
		ret = append(ret, "o carro não está com o freio de mão acionado, tá cagando de novo")
	}

	if c.PortaMalas != PortaAberta {
		ret = append(ret, "o porta malas não está aberto, quer quebrar o vidro?")
	}

	if c.DentroDoCarro {
		ret = append(ret, "Esse é um carro simples, vc precisa estar fora do carro para fazer isso")
	}

	if len(ret) > 0 {
		return c.LogErr("GuardaChaveDeRodas", ret)
	} else {
		c.ChaveDeRoda = chave
	}

	c.Log("GuardaChaveDeRodas", "ok")
	return nil
}

func (c *Carro) TiraParafusos(posicao RodaPosicao, chave *FerramentaDeRoda) (int, []error) {
	var ret []string
	var parafusos int

	if chave == nil {
		ret = append(ret, "Cadê a chave? Tá doido?")
	}

	if c.Rodas[posicao].Parafusos == 0 {
		ret = append(ret, "A roda não tem parafusos, vc já estava morto")
	}

	if c.Velocidade > 0 {
		ret = append(ret, "o carro está andando, vc vai fazer bosta")
	}

	if !c.FreioDeMao {
		ret = append(ret, "o carro não está com o freio de mão acionado, tá cagando de novo")
	}

	if c.DentroDoCarro {
		ret = append(ret, "Esse é um carro simples, vc precisa estar fora do carro para fazer isso")
	}

	if c.PosicaoSuspenso != posicao {
		ret = append(ret, "O carro ainda não está suspenso nessa posição, pode até dar certo, mas prefiro não arriscar")
	}

	if r, found := c.Rodas[posicao]; found {
		parafusos = r.Parafusos
		r.Parafusos = 0
	} else {
		ret = append(ret, "não achei uma roda nessa posição")
	}

	if len(ret) > 0 {
		return parafusos, c.LogErr("TiraParafusos", ret)
	}

	c.Log("TiraParafusos", "ok")
	return parafusos, nil
}

func (c *Carro) ColocaParafusos(posicao RodaPosicao, chave *FerramentaDeRoda, parafusos int) []error {
	var ret []string

	if chave == nil {
		ret = append(ret, "Cadê a chave? Tá doido?")
	}

	if parafusos == 0 {
		ret = append(ret, "Preciso dos parafusos")
	}

	if c.Velocidade > 0 {
		ret = append(ret, "o carro está andando, vc vai fazer bosta")
	}

	if !c.FreioDeMao {
		ret = append(ret, "o carro não está com o freio de mão acionado, tá cagando de novo")
	}

	if c.DentroDoCarro {
		ret = append(ret, "Esse é um carro simples, vc precisa estar fora do carro para fazer isso")
	}

	if c.PosicaoSuspenso != posicao {
		ret = append(ret, "O carro ainda não está suspenso nessa posição, pode até dar certo, mas prefiro não arriscar")
	}

	if r, found := c.Rodas[posicao]; found {
		r.Parafusos = parafusos
	} else {
		ret = append(ret, "não achei uma roda nessa posição")
	}

	if len(ret) > 0 {
		return c.LogErr("ColocaParafusos", ret)
	}

	c.Log("ColocaParafusos", "ok")
	return nil
}

func (c *Carro) TiraRoda(posicao RodaPosicao) (*Roda, []error) {
	var ret []string
	var roda *Roda

	if c.Velocidade > 0 {
		ret = append(ret, "o carro está andando, vc vai fazer bosta")
	}

	if !c.FreioDeMao {
		ret = append(ret, "o carro não está com o freio de mão acionado, tá cagando de novo")
	}

	if c.DentroDoCarro {
		ret = append(ret, "Esse é um carro simples, vc precisa estar fora do carro para fazer isso")
	}

	if c.PosicaoSuspenso != posicao {
		ret = append(ret, "O carro ainda não está suspenso nessa posição, pode até dar certo, mas prefiro não arriscar")
	}

	if r, found := c.Rodas[posicao]; found {
		if r == nil {
			ret = append(ret, "não tem nenhnuma roda aqui")
		} else {
			if r.Parafusos > 0 {
				ret = append(ret, "A roda está presa com parafusos")
			} else {
				roda = r
				delete(c.Rodas, posicao)
			}
		}
	} else {
		ret = append(ret, "não achei uma roda nessa posição")
	}

	if len(ret) > 0 {
		return roda, c.LogErr("TiraRoda", ret)
	}

	c.Log("TiraRoda", "ok")
	return roda, nil
}

func (c *Carro) ColocaRoda(roda *Roda, posicao RodaPosicao) []error {
	var ret []string

	if c.Velocidade > 0 {
		ret = append(ret, "o carro está andando, vc vai fazer bosta")
	}

	if !c.FreioDeMao {
		ret = append(ret, "o carro não está com o freio de mão acionado, tá cagando de novo")
	}

	if c.DentroDoCarro {
		ret = append(ret, "Esse é um carro simples, vc precisa estar fora do carro para fazer isso")
	}

	if c.PosicaoSuspenso != posicao {
		ret = append(ret, "O carro ainda não está suspenso nessa posição, pode até dar certo, mas prefiro não arriscar")
	}

	if _, found := c.Rodas[posicao]; found {
		ret = append(ret, "Já existe uma roda aqui, vc tá bem?")
	} else {
		c.Rodas[posicao] = roda
	}

	if len(ret) > 0 {
		return c.LogErr("ColocaRoda", ret)
	}

	c.Log("ColocaRoda", "ok")
	return nil
}

func (c *Carro) ColocaCintoDeSegurancaMotorista() []error {
	var ret []string

	if c.Velocidade > 0 {
		ret = append(ret, "você deveria ter feito isso com o carro parado")
	}

	if !c.DentroDoCarro {
		ret = append(ret, "Não dá para colocar o cinto com vc fora do carro")
	}

	if c.CintoMotorista == true {
		ret = append(ret, "Ow retardado, você já está de cinto")
	}

	if len(ret) > 0 {
		return c.LogErr("ColocaCintoDeSegurancaMotorista", ret)
	} else {
		c.CintoMotorista = true
	}

	c.Log("ColocaCintoDeSegurancaMotorista", "ok")
	return nil
}

func (c *Carro) RetiraCintoDeSegurancaMotorista() []error {
	var ret []string

	if c.Velocidade > 0 {
		ret = append(ret, "você não deveria fazer isso com o carro em movimento")
	}

	if !c.DentroDoCarro {
		ret = append(ret, "Não dá para tirar o cinto com vc fora do carro")
	}

	if c.CintoMotorista == false {
		ret = append(ret, "Ow retardado, você já está sem o cinto")
	}

	if len(ret) > 0 {
		return c.LogErr("RetiraCintoDeSegurancaMotorista", ret)
	} else {
		c.CintoMotorista = false
	}

	c.Log("RetiraCintoDeSegurancaMotorista", "ok")
	return nil
}

func (c *Carro) MudaEstadoPiscaAlerta(ligado bool) []error {
	var ret []string

	if !c.DentroDoCarro {
		ret = append(ret, "Não é carrinho de controle remoto, vc precisa estar dentro dele para isso")
	}

	if c.EstadoPiscaAlerta == ligado {
		ret = append(ret, "O pisca alerta já está nesse estado")
	}

	if len(ret) > 0 {
		return c.LogErr("MudaEstadoPiscaAlerta", ret)
	} else {
		c.EstadoPiscaAlerta = ligado
	}

	c.Log("MudaEstadoPiscaAlerta", "ok")
	return nil
}

func (c *Carro) UsaMacaco(macaco *MacacoAutomotivo, posicao RodaPosicao) []error {
	var ret []string

	if macaco == nil {
		ret = append(ret, "Vc é o super homem? cadê o macado?")
	}

	if c.Macaco != nil {
		ret = append(ret, "Vc é ainda não pegou o macado?")
	}

	if c.Velocidade > 0 {
		ret = append(ret, "o carro está andando, vc vai fazer bosta")
	}

	if !c.FreioDeMao {
		ret = append(ret, "o carro não está com o freio de mão acionado, tá cagando de novo")
	}

	if c.DentroDoCarro {
		ret = append(ret, "Esse é um carro simples, vc precisa estar fora do carro para fazer isso")
	}

	if c.PosicaoSuspenso == posicao {
		ret = append(ret, "O carro já está suspenso nessa posição, vc sabe mesmo oq está fazendo?")
	}

	if len(ret) > 0 {
		return c.LogErr("UsaMacaco", ret)
	} else {
		c.PosicaoSuspenso = posicao
	}

	c.Log("UsaMacaco", "ok")
	return nil
}

func (c *Carro) RetiraMacaco(macaco *MacacoAutomotivo, posicao RodaPosicao) []error {
	var ret []string

	if macaco == nil {
		ret = append(ret, "Vc é o super homem? cadê o macado?")
	}

	if c.Macaco != nil {
		ret = append(ret, "Vc é ainda não pegou o macado?")
	}

	if c.Velocidade > 0 {
		ret = append(ret, "o carro está andando, vc vai fazer bosta")
	}

	if !c.FreioDeMao {
		ret = append(ret, "o carro não está com o freio de mão acionado, tá cagando de novo")
	}

	if c.DentroDoCarro {
		ret = append(ret, "Esse é um carro simples, vc precisa estar fora do carro para fazer isso")
	}

	if len(ret) > 0 {
		return c.LogErr("RetiraMacaco", ret)
	} else {
		c.PosicaoSuspenso = PosicaoNenhum
	}

	c.Log("RetiraMacaco", "ok")
	return nil
}
