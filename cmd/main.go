package main

import (
	"os"

	"github.com/gg-tools/markdown-server/internal/route"
)

var (
	bindAddr     = env("BIND_ADDR", ":80")
	staticRoot   = env("STATIC_ROOT", "./res/static/")
	markdownRoot = env("MARKDOWN_ROOT", "./res/markdowns/")
)

func main() {
	route.Serve(bindAddr, staticRoot, markdownRoot)
}

func env(name string, defaultValue string) string {
	if v := os.Getenv(name); v != "" {
		return v
	} else {
		return defaultValue
	}
}
