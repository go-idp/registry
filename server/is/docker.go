package is

import (
	"github.com/go-zoox/core-utils/strings"
	"github.com/go-zoox/zoox"
)

// Docker check if the client is using docker
//
// Example:
//
//	{
//	  "Accept": [
//	    "application/vnd.oci.image.index.v1+json",
//	    "application/vnd.docker.distribution.manifest.v1+prettyjws",
//	    "application/json",
//	    "application/vnd.oci.image.manifest.v1+json",
//	    "application/vnd.docker.distribution.manifest.v2+json",
//	    "application/vnd.docker.distribution.manifest.list.v2+json"
//	  ],
//	  "Accept-Encoding": [
//	    "gzip"
//	  ],
//	  "Connection": [
//	    "close"
//	  ],
//	  "User-Agent": [
//	    "docker/24.0.6 go/go1.20.7 git-commit/1a79695 kernel/6.4.16-linuxkit os/linux arch/arm64 UpstreamClient(Docker-Client/24.0.6 \\(darwin\\))"
//	  ],
//	  "X-Request-Id": [
//	    "14987760a9f2/nTkGwZGh8l-000018"
//	  ]
//	}
func Docker(ctx *zoox.Context) (ok bool) {
	return strings.Contains(ctx.UserAgent(), "docker/")
}
