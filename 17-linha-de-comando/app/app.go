package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicação de Linha de Comando"
	app.Usage = "Busca IPs e Nomes de Servidores na internet"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "devbook.com.br",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:  "ip",
			Usage: "Busca de IPs de endereços na internet",
			Flags: flags,
			Action: searchIps,
		},
		{
			Name:   "server",
			Usage:  "Busca o nome dos servidores na internet",
			Flags:  flags,
			Action: searchIpsServes,
		},
	}

	return app
}

func searchIps(c *cli.Context) {
	host := c.String("host")

	ips, err := net.LookupIP(host) // IP Number
	if err != nil {
		log.Fatal(err)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func searchIpsServes(c *cli.Context) {
	host := c.String("host")

	servers, err := net.LookupNS(host) // Server Name
	if err != nil {
		log.Fatal(err)
	}

	for _, server := range servers {
		fmt.Println(server.Host)
	}
}
