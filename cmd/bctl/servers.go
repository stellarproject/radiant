package main

import (
	"fmt"

	"github.com/codegangsta/cli"
	"github.com/ehazlett/blackbird"
)

var serversCommand = cli.Command{
	Name:  "servers",
	Usage: "manage servers",
	Subcommands: []cli.Command{
		addServerCommand,
		removeServerCommand,
		listServersCommand,
		reloadCommand,
	},
}

var addServerCommand = cli.Command{
	Name:   "add",
	Usage:  "add server",
	Action: addServer,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Usage: "host name of server",
		},
		cli.StringSliceFlag{
			Name:  "upstream",
			Usage: "server upstreams",
			Value: &cli.StringSlice{},
		},
		cli.DurationFlag{
			Name:  "timeouts",
			Usage: "server timeouts",
		},
		cli.BoolFlag{
			Name:  "tls",
			Usage: "enable tls",
		},
	},
}

func addServer(ctx *cli.Context) error {
	client, err := getClient(ctx)
	if err != nil {
		return err
	}
	defer client.Close()

	host := ctx.String("host")
	upstreams := ctx.StringSlice("upstream")
	opts := []blackbird.AddOpts{
		blackbird.WithUpstreams(upstreams...),
		blackbird.WithTimeouts(ctx.Duration("timeout")),
	}
	if ctx.Bool("tls") {
		opts = append(opts, blackbird.WithTLS)
	}

	if err := client.AddServer(host, opts...); err != nil {
		return err
	}

	return nil
}

var removeServerCommand = cli.Command{
	Name:      "remove",
	Usage:     "remove server",
	Action:    removeServer,
	ArgsUsage: "[HOST]",
}

func removeServer(ctx *cli.Context) error {
	host := ctx.Args().First()
	if host == "" {
		return fmt.Errorf("you must specify a host")
	}

	client, err := getClient(ctx)
	if err != nil {
		return err
	}
	defer client.Close()

	if err := client.RemoveServer(host); err != nil {
		return err
	}

	return nil
}

var listServersCommand = cli.Command{
	Name:   "list",
	Usage:  "list servers",
	Action: listServers,
}

func listServers(ctx *cli.Context) error {
	client, err := getClient(ctx)
	if err != nil {
		return err
	}
	defer client.Close()

	servers, err := client.Servers()
	if err != nil {
		return err
	}

	fmt.Println(servers)

	return nil
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
