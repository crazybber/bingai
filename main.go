package main

import (
	"crazybber/go-proxy-bingai/api"
	"embed"
	"log"
	"net/http"
	"os"
	"time"
)

var (
	//go:embed web/*
	webFS     embed.FS
	version   string = "latest"
	buildDate string = "latest"
	commitId  string = "crazybber"
)

// init to bundle the static files
func init() {
	api.WEB_FS = webFS
	api.InitStaticPages()
}

func main() {

	//return version + "-" + buildDate + "-" + commitId when visiting /ver url
	http.HandleFunc("/ver/", func(w http.ResponseWriter, r *http.Request) {
		// Concatenate variables into response string
		response := version + " " + buildDate + " " + commitId
		w.Write([]byte(response))

	})

	http.HandleFunc("/sysconf/", api.SysConf)

	http.HandleFunc("/sydney/", api.Sydney)

	http.HandleFunc("/web/", api.WebStatic)

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
