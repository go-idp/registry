package is

import (
	"github.com/go-zoox/core-utils/strings"
	"github.com/go-zoox/zoox"
)

// Curl checks if the client is using curl
func Curl(ctx *zoox.Context) (ok bool) {
	// curl
	//
	//	{
	//	  "Accept": [
	//	    "*/*"
	//	  ],
	//	  "Accept-Encoding": [
	//	    "gzip"
	//	  ],
	//	  "User-Agent": [
	//	    "curl/7.64.1"
	//	  ],
	//	  "X-Request-Id": [
	//	    "14987760a9f2/nTkGwZGh8l-000020"
	//	  ]
	//	}
	return strings.StartsWith(ctx.UserAgent(), "curl/")
}
