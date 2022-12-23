// HttpFileServer project main.go
package main

import (
	"flag"

	"github.com/kolonse/KolonseWeb"
	"github.com/kolonse/KolonseWeb/middleWare/StaticDir"
)

var dir = flag.String("d", ".", "-d : static file dir")
var port = flag.Int("p", 54321, "-p : port")

func main() {
	flag.Parse()
	KolonseWeb.DefaultApp.Use(StaticDir.NewMiddleWare(*dir))
	KolonseWeb.DefaultApp.Listen("0.0.0.0", *port)
}
