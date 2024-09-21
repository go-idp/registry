package is

import (
	"github.com/go-zoox/core-utils/strings"
	"github.com/go-zoox/zoox"
)

// Cnpm check if the client is using cnpm
//
//	{
//	  "Accept": [
//	    "application/json"
//	  ],
//	  "Connection": [
//	    "keep-alive"
//	  ],
//	  "User-Agent": [
//	    "npminstall/6.5.1 node-urllib/3.0.0 Node.js/16.20.2 (OS X; arm64)"
//	  ],
//	  "X-Request-Id": [
//	    "14987760a9f2/nTkGwZGh8l-000022"
//	  ]
//	}
//
//	{
//		"Connection": [
//			"close"
//		],
//		"Remoteip": [
//			"222.71.122.104"
//		],
//		"User-Agent": [
//			"node-urllib/3.27.0 Node.js/20.9.0 (OS X; arm64)"
//		],
//		"X-Request-Id": [
//			"91e641ce1926/do9mLDeeMk-022983"
//		]
//	}
func Cnpm(ctx *zoox.Context) (ok bool) {
	if ok := strings.Contains(ctx.UserAgent(), "npminstall/"); ok {
		return true
	}

	if ok := strings.Contains(ctx.UserAgent(), "node-urllib/"); ok {
		return true
	}

	return false
}
