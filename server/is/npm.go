package is

import (
	"github.com/go-zoox/core-utils/strings"
	"github.com/go-zoox/zoox"
)

// Npm check if the client is using npm
func Npm(ctx *zoox.Context) (ok bool) {
	// npm
	//	// npm
	//	{
	//	  "Accept": [
	//	    "application/vnd.npm.install-v1+json; q=1.0, application/json; q=0.8, */*"
	//	  ],
	//	  "Accept-Encoding": [
	//	    "gzip,deflate"
	//	  ],
	//	  "Connection": [
	//	    "close"
	//	  ],
	//	  "Npm-Auth-Type": [
	//	    "web"
	//	  ],
	//	  "Npm-Command": [
	//	    "install"
	//	  ],
	//	  "Pacote-Pkg-Id": [
	//	    "registry:@doreamonjs/application"
	//	  ],
	//	  "Pacote-Req-Type": [
	//	    "packument"
	//	  ],
	//	  "Pacote-Version": [
	//	    "15.1.1"
	//	  ],
	//	  "User-Agent": [
	//	    "npm/9.6.4 node/v16.20.2 darwin arm64 workspaces/false"
	//	  ],
	//	  "X-Request-Id": [
	//	    "14987760a9f2/nTkGwZGh8l-000020"
	//	  ]
	//	}
	return strings.Contains(ctx.UserAgent(), "npm/")
}
