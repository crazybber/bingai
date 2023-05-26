package main

import (
	"crazybber/go-proxy-bingai/api"
	"crazybber/go-proxy-bingai/api/helper"
	"crazybber/go-proxy-bingai/common"
	"log"
	"net/http"
	"os"
	"time"
)

var (
	version   string = "latest"
	buildDate string = "latest"
	commitId  string = "crazybber"
)

func main() {

	//return version + "-" + buildDate + "-" + commitId when visiting /ver url
	http.HandleFunc("/ver/", func(w http.ResponseWriter, r *http.Request) {
		// Concatenate variables into response string
		response := version + " " + buildDate + " " + commitId
		w.Write([]byte(response))

	})

	http.HandleFunc("/sysconf", api.SysConf)

	http.HandleFunc("/sydney/", api.Sydney)

	http.HandleFunc("/web/", webStatic)

	http.HandleFunc("/", api.Index)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	addr := ":" + port

	log.Println("Starting BingAI Proxy At " + addr)

	srv := &http.Server{
		Addr:         addr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}

// Serve static pages
func webStatic(w http.ResponseWriter, r *http.Request) {
	if _, ok := WEB_PATH_MAP[r.URL.Path]; ok || r.URL.Path == common.PROXY_WEB_PREFIX_PATH {
		http.StripPrefix(common.PROXY_WEB_PREFIX_PATH, http.FileServer(GetWebFS())).ServeHTTP(w, r)
	} else {
		if !helper.CheckAuth(r) {
			helper.UnauthorizedResult(w)
			return
		}
		common.NewSingleHostReverseProxy(common.BING_URL).ServeHTTP(w, r)
	}

}
