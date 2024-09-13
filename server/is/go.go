package is

import (
	"github.com/go-zoox/core-utils/strings"
	"github.com/go-zoox/zoox"
)

// Go checks if the client is using go
func Go(ctx *zoox.Context) (ok bool) {
	// go
	//
	//	{
	//	  "Accept": [
	//	    "*/*"
	//	  ],
	//	  "Accept-Encoding": [
	//	    "identity"
	//	  ],
	//	  "User-Agent": [
	//	    "Go-http-client/1.1"
	//	  ],
	//	  "X-Request-Id": [
	//	    "14987760a9f2/nTkGwZGh8l-000022"
	//	  ]
	//	}
	return strings.Contains(ctx.UserAgent(), "Go-http-client/")
}
