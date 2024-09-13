package is

import (
	"github.com/go-zoox/core-utils/strings"
	"github.com/go-zoox/zoox"
)

// Yum check if the client is using yum
func Yum(ctx *zoox.Context) (ok bool) {
	// yum
	//
	//	{
	//	  "Accept": [
	//	    "*/*"
	//	  ],
	//	  "Accept-Encoding": [
	//	    "identity"
	//	  ],
	//	  "User-Agent": [
	//	    "urlgrabber/3.10 yum/3.4.3"
	//	  ],
	//	  "X-Request-Id": [
	//	    "14987760a9f2/nTkGwZGh8l-000020"
	//	  ]
	//	}
	return strings.Contains(ctx.UserAgent(), "yum/")
}
