package is

import (
	"github.com/go-zoox/core-utils/strings"
	"github.com/go-zoox/zoox"
)

// Apt check if the client is using apt
func Apt(ctx *zoox.Context) (ok bool) {
	// apt
	//
	//	{
	//	  "Accept": [
	//	    "*/*"
	//	  ],
	//	  "Accept-Encoding": [
	//	    "gzip, deflate, br"
	//	  ],
	//	  "User-Agent": [
	//	    "Debian APT-HTTP/1.3 (2.2.3)"
	//	  ],
	//	  "X-Request-Id": [
	//	    "14987760a9f2/nTkGwZGh8l-000023"
	//	  ]
	//	}
	return strings.Contains(ctx.UserAgent(), "Debian APT-HTTP/")
}
