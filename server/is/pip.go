package is

import (
	"github.com/go-zoox/core-utils/strings"
	"github.com/go-zoox/zoox"
)

// Pip is the interface for the pip command
func Pip(ctx *zoox.Context) (ok bool) {
	// pip
	//	{
	//	  "Accept": [
	//	    "*/*"
	//	  ],
	//	  "Connection": [
	//	    "keep-alive"
	//	  ],
	//	  "Host": [
	//	    "pypi.org"
	//	  ],
	//	  "User-Agent": [
	//	    "pip/21.2.4"
	//	  ]
	//	}
	return strings.Contains(ctx.UserAgent(), "pip/")
}
