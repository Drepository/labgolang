package main

import "fmt"

type ContaCorrente struct {
	Titular        string
	Agencia, Conta int
	Saldo          float64
}

func (c *ContaCorrente) sacar(valorDoSaque float64) {
	c.Saldo -= valorDoSaque
}

func main() {

	contaDoGuilherme := ContaCorrente{Titular: "Guilherme", Agencia: 123, Conta: 22222, Saldo: 105.51}

	contaDaBruna := ContaCorrente{"Bruna", 222, 111222, 200}

	contaDoPatrik := ContaCorrente{Titular: "Patrik", Saldo: 2000}

	contaDoMrMarzio := ContaCorrente{"Marzio", 0, 0, 5000}

	fmt.Println(contaDoGuilherme)
	fmt.Println(contaDaBruna)
	fmt.Println(contaDoPatrik)
	fmt.Println(contaDoMrMarzio)

	var contaDaCris *ContaCorrente
	contaDaCris = new(ContaCorrente)
	contaDaCris.Titular = "Cris"

	fmt.Println(contaDaCris)

	fmt.Println(ContaCorrente{})

	contaDaBruna.sacar(150)

	fmt.Println(contaDaBruna)
}
