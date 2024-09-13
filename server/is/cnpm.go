package is

import (
	"github.com/go-zoox/core-utils/strings"
	"github.com/go-zoox/zoox"
)

// Cnpm check if the client is using cnpm
func Cnpm(ctx *zoox.Context) (ok bool) {
	// cnpm
	//	{
	//	  "Accept": [
	//	    "application/json"
	//	  ],
	//	  "Connection": [
	//	    "keep-alive"
	//	  ],
	//	  "User-Agent": [
	//	    "npminstall/6.5.1 node-urllib/3.0.0 Node.js/16.20.2 (OS X; arm64)"
	//	  ],
	//	  "X-Request-Id": [
	//	    "14987760a9f2/nTkGwZGh8l-000022"
	//	  ]
	//	}
	return strings.Contains(ctx.UserAgent(), "npminstall/")
}
