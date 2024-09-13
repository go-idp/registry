package is

import (
	"github.com/go-zoox/core-utils/strings"
	"github.com/go-zoox/zoox"
)

// Yarn check if the client is using yarn
func Yarn(ctx *zoox.Context) (ok bool) {
	// yarn
	//
	//	{
	//	  "Accept": [
	//	    "application/vnd.npm.install-v1+json; q=1.0, application/json; q=0.8, */*"
	//	  ],
	//	  "Accept-Encoding": [
	//	    "gzip, deflate"
	//	  ],
	//	  "User-Agent": [
	//	    "yarn/1.22.19 npm/? node/v16.20.2 darwin arm64"
	//	  ],
	//	  "X-Request-Id": [
	//	    "14987760a9f2/nTkGwZGh8l-000023"
	//	  ]
	//	}
	return strings.Contains(ctx.UserAgent(), "yarn/")
}
