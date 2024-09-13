package is

import (
	"github.com/go-zoox/core-utils/strings"
	"github.com/go-zoox/zoox"
)

// Apk checks if the client is using apk
func Apk(ctx *zoox.Context) (ok bool) {
	// apk
	//
	//	{
	//	  "Accept": [
	//	    "*/*"
	//	  ],
	//	  "Accept-Encoding": [
	//	    "identity"
	//	  ],
	//	  "User-Agent": [
	//	    "Alpine apk-tools 2.10.5"
	//	  ],
	//	  "X-Request-Id": [
	//	    "14987760a9f2/nTkGwZGh8l-000021"
	//	  ]
	//	}
	return strings.Contains(ctx.UserAgent(), "Alpine apk-tools")
}
