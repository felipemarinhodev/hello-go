package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const delay = 5
const monitoring = 3

func main()  {
	showIntro()

	registerLog("site-false", false)
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
	fmt.Println("")
	return commandRead
}

func startMonitoring() {
	fmt.Println("Monitorando...")

	sites := readSitesFile();

	for i := 0; i < monitoring; i++ {
		for i, site := range sites {
			fmt.Println("Testando site", i, ":", site)
			testSite(site)
		}

		time.Sleep(delay * time.Second)
		fmt.Println("")
	}
	fmt.Println("")
}

func testSite(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
		registerLog(site, true)
		} else {
			fmt.Println("Site:", site, "está com problema. Status code:", resp.StatusCode)
			registerLog(site, true)
	}
}

func readSitesFile() []string {
	var sites []string
	file, err := os.Open("sites.txt")
	if err != nil {
		fmt.Println("Ocorreu um erro", err)
	}

	reader := bufio.NewReader(file)
	for  {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		if err == io.EOF {
			break
		}
		fmt.Println("Endereços localizados", line)
		sites = append(sites, line)
	}
	fmt.Println(sites)
	file.Close()
	return sites
}

func registerLog(site string, status bool) {

	file, err := os.OpenFile("log.txt", os.O_RDWR| os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("registerLog err", err)
	}
	file.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - online: "+ strconv.FormatBool(status) + "\n")
	file.Close()
}
