package is

import (
	"github.com/go-zoox/core-utils/strings"
	"github.com/go-zoox/zoox"
)

// Git checks if the client is using git
//
//	{
//		"Accept": [
//			"*/*"
//		],
//		"Accept-Encoding": [
//			"deflate, gzip"
//		],
//		"Accept-Language": [
//			"en-US, *;q=0.9"
//		],
//		"Cache-Control": [
//			"no-cache"
//		],
//		"Connection": [
//			"close"
//		],
//		"Git-Protocol": [
//			"version=2"
//		],
//		"Pragma": [
//			"no-cache"
//		],
//		"User-Agent": [
//			"git/2.40.0"
//		],
//		"X-Request-Id": [
//			"febcc805d9a3/8gk319QOph-000009"
//		]
//	}
func Git(ctx *zoox.Context) (ok bool) {
	return strings.StartsWith(ctx.UserAgent(), "git/")
}
