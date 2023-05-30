package api

import (
	"crazybber/go-proxy-bingai/common"
	"crazybber/go-proxy-bingai/helper"
	"net/http"
)

func Sydney(w http.ResponseWriter, r *http.Request) {
	if !helper.CheckAuth(r) {
		helper.UnauthorizedResult(w)
		return
	}
	common.NewSingleHostReverseProxy(common.BING_SYDNEY_URL).ServeHTTP(w, r)
}
