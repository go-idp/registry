package is

import (
	"github.com/go-zoox/core-utils/strings"
	"github.com/go-zoox/zoox"
)

// TypeInstaller checks if the client is using a type installer
//
//	Repo: https://github.com/nfour/types-installer
//
//	{
//	   "headers": {
//	     "Accept": [
//	       "*/*"
//	     ],
//	     "Accept-Encoding": [
//	       "gzip,deflate"
//	     ],
//	     "Connection": [
//	       "close"
//	     ],
//	     "Content-Encoding": [
//	       "gzip"
//	     ],
//	     "Content-Length": [
//	       "185"
//	     ],
//	     "Content-Type": [
//	       "application/json"
//	     ],
//	     "Npm-Auth-Type": [
//	       "web"
//	     ],
//	     "Npm-Command": [
//	       "install"
//	     ],
//	     "User-Agent": [
//	       "typesInstaller/5.5.4"
//	     ],
//	     "X-Request-Id": [
//	       "4b5860a3e5d9/LKroIteYXu-000854"
//	     ]
//	   },
//	   "method": "POST",
//	   "path": "/-/npm/v1/security/audits/quick"
//	}
func TypeInstaller(ctx *zoox.Context) (ok bool) {
	return strings.StartsWith(ctx.UserAgent(), "typesInstaller/")
}
