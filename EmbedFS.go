package main

import (
	"embed"
	"io/fs"
	"strings"
)

type EmbedFS struct {
	f embed.FS
}

func (embed *EmbedFS) Open(name string) (fs.File, error) {
	if strings.HasSuffix(name, "/") {
		name += "index.html"
	}
	return embed.f.Open(name)
}