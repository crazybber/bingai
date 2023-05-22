package main

import (
	"embed"
	"io/fs"
	"net/http"
	"os"
	"path/filepath"
)

var (
	//go:embed web/*
	webFS embed.FS
	// is debug or not
	IS_DEBUG_MODE bool
	//path mapping
	WEB_PATH_MAP = make(map[string]bool)
)

func init() {
	IS_DEBUG_MODE = os.Getenv("Go_Proxy_BingAI_Debug") == "true"

	var err error
	if IS_DEBUG_MODE {
		err = initWebPathMapByDir()
	} else {
		err = initWebPathMapByFS()
	}
	if err != nil {
		panic(err)
	}
}

func initWebPathMapByDir() error {
	err := filepath.WalkDir("web", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			WEB_PATH_MAP["/"+path] = true
		}
		return nil
	})
	return err
}

func initWebPathMapByFS() error {
	err := fs.WalkDir(webFS, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			WEB_PATH_MAP["/"+path] = true
		}
		return nil
	})
	return err
}

func GetWebFS() http.FileSystem {
	if IS_DEBUG_MODE {
		return http.Dir("web")
	} else {
		return http.FS(webFS)
	}
}
