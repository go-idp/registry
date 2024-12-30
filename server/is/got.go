package is

import (
	"github.com/go-zoox/core-utils/strings"
	"github.com/go-zoox/zoox"
)

// Got check if the client is using got
func Got(ctx *zoox.Context) (ok bool) {
	// got
	// {
	// 	"Accept-Encoding": [
	// 		"gzip, deflate, br"
	// 	],
	// 	"Connection": [
	// 		"close"
	// 	],
	// 	"User-Agent": [
	// 		"got (https://github.com/sindresorhus/got)"
	// 	]
	// }

	return strings.Contains(ctx.UserAgent(), "sindresorhus/got")
}
