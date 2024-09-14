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
//
// // buildkit (docker buildx)
// {                                                                              [3/2218]
//
//		"headers": {
//		  "Accept": [
//		    "application/vnd.docker.distribution.manifest.v2+json, application/vnd.docker.distribution.manifest.list.v2+json, application/vnd.oci.image.manifest.v1+json, application/vnd.oci.image.index.v1+json, */*"
//		  ],
//		  "Connection": [
//		    "close"
//		  ],
//		  "Traceparent": [
//		    "00-e6c58f2b63707ac86153ead04283d8e0-749297dfcf6b0a62-01"
//		  ],
//		  "User-Agent": [
//		    "buildkit/v0.12"
//		  ],
//		  "X-Request-Id": [
//		    "3b9f38c61c5f/eXYwXysyvl-000372"
//		  ]
//		},
//		"method": "HEAD",
//		"path": "/v2/whatwewant/builder-ta-node/manifests/v16-1"
//	}
func Docker(ctx *zoox.Context) (ok bool) {
	// docker build
	if ok := strings.Contains(ctx.UserAgent(), "docker/"); ok {
		return true
	}

	// docker buildx build => buildkit
	if ok := strings.StartsWith(ctx.UserAgent(), "buildkit/"); ok {
		return true
	}

	return false
}
