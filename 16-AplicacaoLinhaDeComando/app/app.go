package app

import (
	"fmt"
	"github.com/urfave/cli"
	"log"
	"net"
)

// Generate Return application CLI exec
func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "Application CLI"
	app.Usage = "Search Server IPs Name"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "github.com",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Search IPs",
			Flags:  flags,
			Action: searchIps,
		},
		{
			Name:   "servers",
			Usage:  "Search Servers",
			Flags:  flags,
			Action: searchServers,
		},
	}

	return app
}

func searchIps(c *cli.Context) {
	host := c.String("host")

	// net
	ips, erro := net.LookupIP(host)

	if erro != nil {
		log.Fatal(erro)
	}

	for _, ip := range ips {
		fmt.Println("IP =====> ", ip)
	}
}

func searchServers(c *cli.Context) {
	host := c.String("host")

	// net
	servers, erro := net.LookupNS(host) // NS => name server

	if erro != nil {
		log.Fatal(erro)
	}

	for _, server := range servers {
		fmt.Println("Server =====> ", server.Host)
	}
}
