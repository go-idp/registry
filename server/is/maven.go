package is

import (
	"github.com/go-zoox/core-utils/strings"
	"github.com/go-zoox/zoox"
)

// Maven is the maven registry
func Maven(ctx *zoox.Context) (ok bool) {
	// maven
	//	{
	//	  "Accept": [
	//	    "text/html, image/gif, image/jpeg, *; q=.2, */*; q=.2"
	//	  ],
	//	  "Connection": [
	//	    "keep-alive"
	//	  ],
	//	  "Host": [
	//	    "repo.maven.apache.org"
	//	  ],
	//	  "User-Agent": [
	//	    "Apache-Maven/3.8.4 (Java 11.0.12; Mac OS X 11.6.2)"
	//	  ]
	//	}
	return strings.Contains(ctx.Header().Get("User-Agent"), "Apache-Maven/")
}
