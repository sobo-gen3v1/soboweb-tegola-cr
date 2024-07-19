package ui

import (
	"embed"
	"io/fs"
	"net/http"
	"strings"

	"github.com/sobo-gen3v1/soboweb-tegola-cr/internal/log"
)

//Embed UI dist Folder recursively
//go:embed dist/*
var dist embed.FS

const (
	distDir       = "dist"
	assetsDir     = "assets"
	indexJSPrefix = "index-"
)

func GetDistFileSystem() http.FileSystem {
	distFS := GetDistFS()
	return http.FS(distFS)
}

func GetDistFS() fs.FS {
	distFS, err := fs.Sub(dist, distDir)
	if err != nil {
		log.Fatal(err)
	}
	return distFS
}

func Version() string {
	// read the assets/ directory so we can attempt to find the built .js file
	files, err := fs.ReadDir(GetDistFS(), assetsDir)
	if err != nil {
		return ""
	}

	for _, entry := range files {
		name := entry.Name()
		if strings.HasPrefix(name, indexJSPrefix) && strings.HasSuffix(name, ".js") {
			// expect it to be index-{version}.js
			return name[len(indexJSPrefix) : len(name)-3]
		}
	}

	return ""
}
