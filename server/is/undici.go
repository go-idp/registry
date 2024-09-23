package is

import (
	"github.com/go-zoox/core-utils/strings"
	"github.com/go-zoox/zoox"
)

// Undici check if the client is using undici
//
//	{
//		"Accept": [
//			"*/*"
//		],
//		"Accept-Encoding": [
//			"br, gzip, deflate"
//		],
//		"Accept-Language": [
//			"*"
//		],
//		"Connection": [
//			"close"
//		],
//		"Sec-Fetch-Mode": [
//			"cors"
//		],
//		"User-Agent": [
//			"undici"
//		],
//		"X-Request-Id": [
//			"be9e9db2f422/rU1byNY6LG-005692"
//		]
//	}
func Undici(ctx *zoox.Context) (ok bool) {
	return strings.StartsWith(ctx.UserAgent(), "undici")
}
