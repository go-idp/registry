package is

import (
	"github.com/go-zoox/core-utils/strings"
	"github.com/go-zoox/zoox"
)

// Chrome checks if the client is using chrome
func Chrome(ctx *zoox.Context) (ok bool) {
	// chrome
	//
	//	{
	//	  "Accept": [
	//	    "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8"
	//	  ],
	//	  "Accept-Encoding": [
	//	    "gzip, deflate, sdch"
	//	  ],
	//	  "Accept-Language": [
	//	    "en-US,en;q=0.8"
	//	  ],
	//	  "User-Agent": [
	//	    "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/51.0.2704.106 Safari/537.36"
	//	  ],
	//	  "X-Request-Id": [
	//	    "14987760a9f2/nTkGwZGh8l-000023"
	//	  ]
	//	}
	return strings.Contains(ctx.UserAgent(), "Chrome/")
}
