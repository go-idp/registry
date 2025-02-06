package is

import (
	"github.com/go-zoox/zoox"
)

// Node check if the client is using node
func Node(ctx *zoox.Context) (ok bool) {
	// {
	// 	"Accept": [
	// 		"*/*"
	// 	],
	// 	"Accept-Encoding": [
	// 		"br, gzip, deflate"
	// 	],
	// 	"Accept-Language": [
	// 		"*"
	// 	],
	// 	"Connection": [
	// 		"close"
	// 	],
	// 	"Sec-Fetch-Mode": [
	// 		"cors"
	// 	],
	// 	"User-Agent": [
	// 		"node"
	// 	],
	// 	"X-Request-Id": [
	// 		"bd0d1acc674f/wGZ2JoFc32-099691"
	// 	]
	// }
	return ctx.UserAgent() == "node"
}
