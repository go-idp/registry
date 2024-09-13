package is

import (
	"github.com/go-zoox/core-utils/strings"
	"github.com/go-zoox/zoox"
)

// Wget check if the client is using wget
func Wget(ctx *zoox.Context) (ok bool) {
	// wget
	//
	//	{
	//	  "Accept": [
	//	    "*/*"
	//	  ],
	//	  "Accept-Encoding": [
	//	    "identity"
	//	  ],
	//	  "User-Agent": [
	//	    "Wget/1.20.3 (linux-gnu)"
	//	  ],
	//	  "X-Request-Id": [
	//	    "14987760a9f2/nTkGwZGh8l-000021"
	//	  ]
	//	}
	return strings.StartsWith(ctx.UserAgent(), "Wget/")
}
