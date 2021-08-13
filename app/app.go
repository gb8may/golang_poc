package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

func Catch() *cli.App {
	app := cli.NewApp()
	app.Name = "NetCatcher"
	app.Usage = "Get Ip and server names from websites."

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "github.com",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Get IP address from website",
			Flags:  flags,
			Action: getIp,
		},
		{
			Name:   "servers",
			Usage:  "Get server name from website",
			Flags:  flags,
			Action: getServer,
		},
	}

	return app
}

func getIp(c *cli.Context) {
	host := c.String("host")

	ip, erro := net.LookupIP(host)
	if erro != nil {
		log.Fatal(erro)
	}

	for _, ip := range ip {
		fmt.Println(ip)
	}
}

func getServer(c *cli.Context) {
	host := c.String("host")

	server, erro := net.LookupNS(host)
	if erro != nil {
		log.Fatal(erro)
	}

	for _, server := range server {
		fmt.Println(server.Host)
	}
}
