package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Gerar fai fazer o jabacule
func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "App de linha de cmd"
	app.Usage = "Busca o IP"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "mauricionofre.com",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Busca a parada",
			Flags:  flags,
			Action: BuscarIps,
		},
		{
			Name:   "servidores",
			Usage:  "busca server",
			Flags:  flags,
			Action: buscarServidores,
		},
	}
	return app
}

func BuscarIps(c *cli.Context) {
	host := c.String("host")

	ips, erro := net.LookupIP(host)
	if erro != nil {
		log.Fatal(erro)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func buscarServidores(c *cli.Context) {
	host := c.String("host")

	servidores, erro := net.LookupNS(host)
	if erro != nil {
		log.Fatal(erro)
	}
	for _, servidor := range servidores {
		fmt.Println(servidor.Host)
	}
}
