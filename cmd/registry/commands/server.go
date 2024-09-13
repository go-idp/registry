package commands

import (
	"github.com/go-idp/registry/server"
	"github.com/go-zoox/cli"
	"github.com/go-zoox/config"
)

func RegisterServer(app *cli.MultipleProgram) {
	app.Register("server", &cli.Command{
		Name:  "server",
		Usage: "Run the registry server",
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:    "port",
				Usage:   "server port",
				Aliases: []string{"p"},
				EnvVars: []string{"PORT"},
				Value:   8080,
			},
			&cli.StringFlag{
				Name:    "config",
				Usage:   "config file",
				Aliases: []string{"c"},
				EnvVars: []string{"CONFIG"},
				Value:   "/etc/registry/config.yml",
				// Required: true,
			},
		},
		Action: func(ctx *cli.Context) error {
			configFile := ctx.String("config")

			cfg := &server.Config{}
			// load config file
			err := config.Load(cfg, &config.LoadOptions{
				FilePath: configFile,
			})
			if err != nil {
				return err
			}

			if ctx.Int("port") != 0 {
				cfg.Port = ctx.Int("port")
			}

			return server.New().Run(cfg)
		},
	})
}
