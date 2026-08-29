package frontend

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
)

//go:embed dist/*
var FS embed.FS

func Handler() http.Handler {
	fsTemp, err := fs.Sub(FS, "dist")
	if err != nil {
		log.Fatal(err)
	}
	return http.FileServer(http.FS(fsTemp))
}
