package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Returns the command line application ready to execute
func Create() *cli.App {
	app := cli.NewApp()
	app.Name = "Command Line Application"
	app.Usage = "Search IPs and server names"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "youtube.com.br",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Search IPs",
			Flags:  flags,
			Action: searchIp,
		},
		{
			Name:   "servers",
			Usage:  "Search servers",
			Flags:  flags,
			Action: searchServer,
		},
	}

	return app
}

func searchIp(c *cli.Context) {
	host := c.String("host")

	ips, ipsError := net.LookupIP(host)
	if ipsError != nil {
		log.Fatal(ipsError)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func searchServer(c *cli.Context) {
	host := c.String("host")

	servers, serverError := net.LookupNS(host)
	if serverError != nil {
		log.Fatal(serverError)
	}

	for _, server := range servers {
		fmt.Println(server.Host)
	}
}
