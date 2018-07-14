package cmd

import (
	"fmt"
	"os"

	"github.com/valyala/fasthttp"
	cli "gopkg.in/urfave/cli.v1"
	"gopkg.in/urfave/cli.v1/altsrc"

	"github.com/thavel/goban/models"
	"github.com/thavel/goban/pkg/database"
	"github.com/thavel/goban/pkg/logger"
	"github.com/thavel/goban/routes"
)

// Server starts goban server
var (
	flags = []cli.Flag{
		cli.StringFlag{Name: "config, c", Value: "./config.yml"},
		altsrc.NewIntFlag(cli.IntFlag{Name: "server.port", Value: 8080}),
		altsrc.NewStringFlag(cli.StringFlag{Name: "server.logs", Value: "info"}),
		altsrc.NewStringFlag(cli.StringFlag{Name: "persistence.host"}),
		altsrc.NewIntFlag(cli.IntFlag{Name: "persistence.port", Value: 3306}),
		altsrc.NewStringFlag(cli.StringFlag{Name: "persistence.database"}),
		altsrc.NewStringFlag(cli.StringFlag{Name: "persistence.username"}),
		altsrc.NewStringFlag(cli.StringFlag{Name: "persistence.password"}),
	}
	Server = cli.Command{
		Name:   "server",
		Action: runServer,
		Flags:  flags,
		Before: before,
	}
)

func before(c *cli.Context) error {
	configFile := c.String("config")
	if _, err := os.Stat(configFile); err == nil {
		return altsrc.InitInputSourceWithContext(
			flags, altsrc.NewYamlSourceFromFlagFunc("config"),
		)(c)
	}
	logger.Infof("No config file found, use defaults")
	return nil
}

func runServer(c *cli.Context) error {
	// Logs
	logger.SetLevel(c.String("server.logs"))

	// Setup database
	dbConfig := database.Config{
		Host:     c.String("persistence.host"),
		Port:     c.Int("persistence.port"),
		Database: c.String("persistence.database"),
		Username: c.String("persistence.username"),
		Password: c.String("persistence.password"),
	}
	err := database.Setup(dbConfig, models.Tables, models.FKeys...)
	if err != nil {
		logger.Errorf("Fail to connect to database: %v", err)
		return err
	}
	defer database.Close()

	// Setup API
	server := &fasthttp.Server{
		Name:    "goban",
		Handler: routes.Handler(),
	}
	uri := fmt.Sprintf(":%d", c.Int("server.port"))
	logger.Infof("HTTP server listening on %s", uri)
	err = server.ListenAndServe(uri)
	if err != nil {
		logger.Fatalf("HTTP server failure: %v", err)
		return err
	}

	return nil
}
