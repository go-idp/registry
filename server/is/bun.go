package is

import (
	"github.com/go-zoox/core-utils/strings"
	"github.com/go-zoox/zoox"
)

// Bun check if the client is using bun
//
//	{
//		"Accept": [
//			"*/*"
//		],
//		"Accept-Encoding": [
//			"br, gzip, deflate"
//		],
//		"Connection": [
//			"close"
//		],
//		"User-Agent": [
//			"Bun/0.5.9"
//		],
//		"X-Request-Id": [
//			"be9e9db2f422/rU1byNY6LG-005769"
//		]
//	}
func Bun(ctx *zoox.Context) (ok bool) {
	return strings.StartsWith(ctx.UserAgent(), "Bun/")
}
