package main

import (
	"fmt"
	"net/http"
	"os"
)

func main()  {
	showIntro()

	for {
		showMenu()

		command := readCommand()

		switch command {
			case 1:
				startMonitoring()
			case 2:
				fmt.Println("Exibindo Logs...")
			case 0:
				fmt.Println("Saindo do programa...")
				os.Exit(0)
			default:
				fmt.Println("Comando não suportado.")
				os.Exit(-1)
		}
	}
}

func showMenu() {
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do programa.")
}

func showIntro() {
	name := "Felipe"
	version := 1.1
	fmt.Println("olá sr.", name)
	fmt.Println("este programa está na versão", version)
}

func readCommand() int {
	var commandRead int;
	fmt.Scan(&commandRead)

	fmt.Println("O comando escolhido foi", commandRead)
	return commandRead
}

func startMonitoring() {
	fmt.Println("Monitorando...")

	sites := []string {
		"https://random-status-code.herokuapp.com",
		"https://www.alura.com.br",
		"https://www.casadocodigo.com.br",
		"https://www.caelum.com.br",
	}

	for _, site := range sites {
		// fmt.Println(sites[i])

		resp, _ := http.Get(site)

		if resp.StatusCode == 200 {
			fmt.Println("Site:", site, "foi carregado com sucesso!")
			} else {
				fmt.Println("Site:", site, "está com problema. Status code:", resp.StatusCode)
			}
		}
}

