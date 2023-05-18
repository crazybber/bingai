package main

import (
	"adams549659584/go-proxy-bingai/api"
	"adams549659584/go-proxy-bingai/common"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	http.HandleFunc("/web/", webStatic)

	http.HandleFunc("/sydney/ChatHub", api.ChatHub)

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

func webStatic(w http.ResponseWriter, r *http.Request) {
	if _, ok := WEB_PATH_MAP[r.URL.Path]; ok || r.URL.Path == common.PROXY_WEB_PREFIX_PATH {
		http.StripPrefix(common.PROXY_WEB_PREFIX_PATH, http.FileServer(GetWebFS())).ServeHTTP(w, r)
	} else {
		common.NewSingleHostReverseProxy(common.BING_URL).ServeHTTP(w, r)
	}
}
