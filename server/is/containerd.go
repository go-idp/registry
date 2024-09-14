package is

import (
	"github.com/go-zoox/core-utils/strings"
	"github.com/go-zoox/zoox"
)

// Containerd check if the client is using containerd
//
//	{
//	   "headers": {
//	     "Accept": [
//	       "application/vnd.docker.distribution.manifest.v2+json, application/vnd.docker.distribution.manifest.list.v2+json, application/vnd.oci.image.manifest.v1+json, application/vnd.oci.image.index.v1+json, */*"
//	     ],
//	     "Connection": [
//	       "close"
//	     ],
//	     "User-Agent": [
//	       "containerd/1.4.0+unknown"
//	     ],
//	     "X-Request-Id": [
//	       "aee97e84b6bb/gL9ZD34gkv-000001"
//	     ]
//	   },
//	   "method": "HEAD",
//	   "path": "/v2/idp/openjdk/manifests/v11-1"
//	 }
func Containerd(ctx *zoox.Context) (ok bool) {
	return strings.StartsWith(ctx.UserAgent(), "containerd/")
}
