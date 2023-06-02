package api

import (
	"crazybber/go-proxy-bingai/common"
	"crazybber/go-proxy-bingai/helper"
	"embed"
	"io/fs"
	"net/http"
	"path/filepath"
)

var (
	//path mapping
	WEB_PATH_MAP = make(map[string]bool)
	//point to main webfs
	WEB_FS embed.FS
)

func InitStaticPages() {
	var err error
	if common.IS_DEBUG_MODE {
		err = initWebPathMapByDir()
	} else {
		err = initWebPathMapBywebFS()
	}
	if err != nil {
		panic(err)
	}
}

// Serve static pages
func WebStatic(w http.ResponseWriter, r *http.Request) {
	if _, ok := WEB_PATH_MAP[r.URL.Path]; ok || r.URL.Path == common.PROXY_WEB_PREFIX_PATH {
		http.FileServer(GetWebFS()).ServeHTTP(w, r)
	} else {
		if !helper.CheckAuth(r) {
			helper.UnauthorizedResult(w)
			return
		}
		common.NewSingleHostReverseProxy(common.BING_URL).ServeHTTP(w, r)
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

func initWebPathMapBywebFS() error {
	err := fs.WalkDir(WEB_FS, ".", func(path string, d fs.DirEntry, err error) error {
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
	if common.IS_DEBUG_MODE {
		return http.Dir("web")
	} else {
		return http.FS(WEB_FS)
	}
}
