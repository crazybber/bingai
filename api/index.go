package api

import (
	"crazybber/go-proxy-bingai/common"
	"crazybber/go-proxy-bingai/helper"
	"net/http"
	"strings"
)

func Index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		http.Redirect(w, r, common.PROXY_WEB_PAGE_PATH, http.StatusFound)
		return
	}
	if strings.HasPrefix("/turing", r.URL.Path) {
		if !helper.CheckAuth(r) {
			helper.UnauthorizedResult(w)
			return
		}
	}
	common.NewSingleHostReverseProxy(common.BING_URL).ServeHTTP(w, r)

}
