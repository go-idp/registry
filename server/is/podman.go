package is

import (
	"github.com/go-zoox/core-utils/strings"
	"github.com/go-zoox/zoox"
)

// Podman checks if the client is using podman
//
// API: /v2
//
//	{
//		"Accept-Encoding": [
//			"gzip"
//		],
//		"Connection": [
//			"close"
//		],
//		"Docker-Distribution-Api-Version": [
//			"registry/2.0"
//		],
//		"User-Agent": [
//			"containers/5.28.0 (github.com/containers/image)"
//		],
//		"X-Request-Id": [
//			"91e641ce1926/do9mLDeeMk-000021"
//		]
//	}
//
// API: /v1/_ping
//
//	{
//		"Accept-Encoding": [
//			"gzip"
//		],
//		"Connection": [
//			"close"
//		],
//		"Docker-Distribution-Api-Version": [
//			"registry/2.0"
//		],
//		"User-Agent": [
//			"containers/5.28.0 (github.com/containers/image)"
//		],
//		"X-Request-Id": [
//			"91e641ce1926/do9mLDeeMk-000022"
//		]
//	}
func Podman(ctx *zoox.Context) (ok bool) {
	return strings.StartsWith(ctx.UserAgent(), "containers/")
}
