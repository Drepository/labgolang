package main

import (
	"bufio"
	"fmt"
	"os"
)

const NOMEARQUIVO = "sites.txt"

func main() {
	quebraPadrao()
	mostarCabecalho()

	for {
		quebraPadrao()
		mostraOpcoes()

		quebraPadrao()
		comando := leituraDeComando()

		opcaoLida(comando)
	}

}

func leituraDeComando() int {
	comandoLido := 0
	fmt.Print("Informe opção:")
	fmt.Scan(&comandoLido)
	fmt.Println("Comando lido:", comandoLido)
	return comandoLido
}

func opcaoLida(comandoLido int) {
	switch comandoLido {
	case 1:
		monitoraEnderecos()
	case 2:
		carregaLog()
	case 0:
		finalizaSistema()
	default:
		opcaoInvalida()
	}
}

func monitoraEnderecos() {
	quebraPadrao()
	fmt.Println("Monitorando endereços...")

	sites := carregaSitesDoArquivo()
	for i, site := range sites {
		fmt.Println(i, site)
	}

	fmt.Println("Testando Site...")
	for i, site := range sites {
		testaSite()
	}
}

func testaSite() {

}

func carregarquivo(nameFile string) []string {

	arquivo, err := os.Open(nameFile)

	leitor := bufio.NewScanner(arquivo)
	leitor.Split(bufio.ScanLines)
	var linhas []string

	for leitor.Scan() {
		linhas = append(linhas, leitor.Text())
	}

	arquivo.Close()
	if err != nil {
		fmt.Println("Ocorreu uma falha:", err)
	}

	arquivo.Close()

	return linhas
}

func carregaSitesDoArquivo() []string {

	site := carregarquivo(NOMEARQUIVO)
	return site
}

func carregaLog() {
	quebraPadrao()
	fmt.Println("Carregando Logs....")
	logs := carregarquivo("logs.txt")
	for i, log := range logs {
		fmt.Println(i, log)
	}
}

func finalizaSistema() {
	quebraPadrao()
	fmt.Println("Programa Finalizado..")
	os.Exit(0)
}

func opcaoInvalida() {
	quebraPadrao()
	fmt.Println("Opcao informada é invalida!!")
	os.Exit(-1)
}

func quebraPadrao() {
	quebraLinha(1)
}

func quebraLinha(quantidade int) {
	const maxquebra = 10

	if quantidade > maxquebra {
		quantidade = maxquebra
	} else if quantidade <= 0 {
		quantidade = 1
	}

	for i := 0; i <= quantidade; i++ {
		fmt.Println("")
	}
}

func mostarCabecalho() {
	versao := 1.1
	desenvolvedor := "Felipe"
	fmt.Println("Versao do Software:", versao)
	fmt.Println("Olá", desenvolvedor, "seja bem vindo!")
}

func mostraOpcoes() {
	fmt.Println("Menu Principal")
	fmt.Println("1 - Monitorar Site")
	fmt.Println("2 - Carregar Log")
	fmt.Println("0 - Sair")
}
