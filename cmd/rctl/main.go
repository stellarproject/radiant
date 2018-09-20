package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"github.com/sirupsen/logrus"
	"github.com/stellarproject/radiant"
	"github.com/stellarproject/radiant/version"
)

func main() {
	app := cli.NewApp()
	app.Name = "bctl"
	app.Version = version.BuildVersion()
	app.Author = "@stellarproject"
	app.Email = ""
	app.Usage = version.Description
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "debug, D",
			Usage: "Enable debug logging",
		},
		cli.StringFlag{
			Name:  "addr, a",
			Usage: "radiant grpc address",
			Value: "unix:///run/radiant.sock",
		},
	}
	app.Commands = []cli.Command{
		serversCommand,
		reloadCommand,
		configCommand,
	}
	app.Before = func(ctx *cli.Context) error {
		if ctx.Bool("debug") {
			logrus.SetLevel(logrus.DebugLevel)
		}

		return nil
	}

	if err := app.Run(os.Args); err != nil {
		logrus.Fatal(err)
	}
}

func getClient(ctx *cli.Context) (*radiant.Client, error) {
	return radiant.NewClient(ctx.GlobalString("addr"))
}

var reloadCommand = cli.Command{
	Name:   "reload",
	Usage:  "reload proxy service",
	Action: reload,
}

func reload(ctx *cli.Context) error {
	client, err := getClient(ctx)
	if err != nil {
		return err
	}
	defer client.Close()
	return client.Reload()
}

var configCommand = cli.Command{
	Name:   "config",
	Usage:  "get current proxy config",
	Action: config,
}

func config(ctx *cli.Context) error {
	client, err := getClient(ctx)
	if err != nil {
		return err
	}
	defer client.Close()
	data, err := client.Config()
	if err != nil {
		return err
	}
	fmt.Println(string(data))
	return nil
}
