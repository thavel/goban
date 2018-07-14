package main

import (
	"os"

	cli "gopkg.in/urfave/cli.v1"

	"github.com/thavel/goban/cmd"
)

const appVersion = "0.0.0"

func main() {
	app := cli.NewApp()
	app.Name = "Goban"
	app.Usage = "Ready-to-go management tool for small enterprises"
	app.Version = appVersion
	app.Commands = []cli.Command{
		cmd.Server,
	}
	app.Flags = append(app.Flags, []cli.Flag{}...)
	app.Run(os.Args)
}
