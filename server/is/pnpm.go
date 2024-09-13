package is

import (
	"github.com/go-zoox/core-utils/strings"
	"github.com/go-zoox/zoox"
)

// Pnpm check if the client is using pnpm
func Pnpm(ctx *zoox.Context) (ok bool) {
	//	pnpm
	//	{
	//	  "Accept": [
	//	    "application/vnd.npm.install-v1+json; q=1.0, application/json; q=0.8, */*"
	//	  ],
	//	  "Accept-Encoding": [
	//	    "gzip,deflate,br"
	//	  ],
	//	  "User-Agent": [
	//	    "pnpm/8.9.0 npm/? node/v16.20.2 darwin arm64"
	//	  ],
	//	  "X-Request-Id": [
	//	    "14987760a9f2/nTkGwZGh8l-000025"
	//	  ]
	//	}
	return strings.Contains(ctx.UserAgent(), "pnpm/")
}
