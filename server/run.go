package server

import (
	"net/http"

	"github.com/go-zoox/core-utils/fmt"
	"github.com/go-zoox/proxy"

	"github.com/go-idp/registry"
	"github.com/go-idp/registry/server/is"
	"github.com/go-zoox/logger"
	"github.com/go-zoox/zoox"
	"github.com/go-zoox/zoox/defaults"
)

func (s *server) Run(cfg *Config) error {
	logger.Infof("server config: %+v", cfg)

	app := defaults.Default()

	app.SetBanner(fmt.Sprintf(`
               _    __                    _     __          
  ___ ____    (_)__/ /__    _______ ___ _(_)__ / /_______ __
 / _ '/ _ \  / / _  / _ \  / __/ -_) _ '/ (_-</ __/ __/ // /
 \_, /\___/ /_/\_,_/ .__/ /_/  \__/\_, /_/___/\__/_/  \_, / 
/___/             /_/             /___/              /___/  %s

Registry Service for Go IDP

____________________________________O/_______
                                    O\
`, registry.Version))

	app.Use(func(ctx *zoox.Context) {
		s.requests.Inc(1)
		ctx.Next()
	})

	app.Use(func(ctx *zoox.Context) {
		// check which package manager

		proxyToRegistry := func(name string, registryX Registry) {
			if registryX.Server == "" {
				panic(fmt.Sprintf("[package manager: %s] registry is not set", name))
			}

			ctx.Logger.Infof("[package manager: %s][registry: %s] %s %s...", name, registryX.Server, ctx.Method, ctx.Path)
			ctx.Proxy(registryX.Server, &proxy.SingleHostConfig{
				RequestHeaders: registryX.Headers,
				OnRequest: func(req *http.Request) error {
					req.Header.Set("X-Registry-Client", fmt.Sprintf("go-idp-registry/%s", registry.Version))
					return nil
				},
			})
		}

		switch {
		case is.Yarn(ctx):
			proxyToRegistry("yarn", cfg.Registries.Npm)
			return

		case is.Pnpm(ctx):
			proxyToRegistry("pnpm", cfg.Registries.Npm)
			return

		case is.Cnpm(ctx):
			proxyToRegistry("cnpm", cfg.Registries.Npm)
			return

		case is.Npm(ctx):
			proxyToRegistry("npm", cfg.Registries.Npm)
			return

		case is.Docker(ctx):
			proxyToRegistry("docker", cfg.Registries.Docker)
			return

		case is.Maven(ctx):
			proxyToRegistry("maven", cfg.Registries.Maven)
			return

		case is.Go(ctx):
			proxyToRegistry("go", cfg.Registries.Go)
			return

		case is.Pip(ctx):
			proxyToRegistry("pip", cfg.Registries.Python)
			return

		case is.Apt(ctx):
			proxyToRegistry("apt", cfg.Registries.Apt)
			return

		case is.Apk(ctx):
			proxyToRegistry("apk", cfg.Registries.Apk)
			return

		case is.Yum(ctx):
			proxyToRegistry("yum", cfg.Registries.Yum)
			return

		case is.Chrome(ctx):
			s.Info()(ctx)
			return

		default:
			fmt.PrintJSON("unknown client", zoox.H{
				"method":  ctx.Method,
				"path":    ctx.Path,
				"headers": ctx.Request.Header,
			})

			ctx.JSON(http.StatusBadGateway, zoox.H{
				"code":    400,
				"message": "unknown client",
				"detail": zoox.H{
					"method":  ctx.Method,
					"path":    ctx.Path,
					"headers": ctx.Request.Header,
				},
			})
		}
	})

	// app.Get("/@@/info", s.Info())

	return app.Run(fmt.Sprintf(":%d", cfg.Port))
}
