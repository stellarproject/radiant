package main

import (
	"os"

	"github.com/codegangsta/cli"
	"github.com/ehazlett/blackbird"
	"github.com/ehazlett/blackbird/version"
	log "github.com/sirupsen/logrus"
)

func main() {
	app := cli.NewApp()
	app.Name = "bctl"
	app.Version = version.BuildVersion()
	app.Author = "@ehazlett"
	app.Email = ""
	app.Usage = version.Description
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "debug, D",
			Usage: "Enable debug logging",
		},
		cli.StringFlag{
			Name:  "addr, a",
			Usage: "blackbird grpc address",
			Value: "127.0.0.1:9000",
		},
	}
	app.Commands = []cli.Command{
		serversCommand,
	}
	app.Before = func(ctx *cli.Context) error {
		if ctx.Bool("debug") {
			log.SetLevel(log.DebugLevel)
		}

		return nil
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func getClient(ctx *cli.Context) (*blackbird.Client, error) {
	return blackbird.NewClient(ctx.GlobalString("addr"))
}
