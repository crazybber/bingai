package api

import (
	"crazybber/go-proxy-bingai/common"
	"net/http"
	"net/http/httptest"
	"testing"
)

func init() {

	//init a WEB_PATH_MAP
	WEB_PATH_MAP = map[string]bool{
		"web/index.html":             true,
		"web/sw.js":                  true,
		"web/registerSW.js":          true,
		"web/js/bing/chat/config.js": true,
	}
}

// 1. 测试请求一个在WEB_PATH_MAP中不存在的路径，期望返回http.StatusNotFound。
func TestWebStaticStatusNotFound(t *testing.T) {
	// Test case 1
	req1, err := http.NewRequest("GET", "/notfound", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr1 := httptest.NewRecorder()
	handler1 := http.HandlerFunc(WebStatic)
	handler1.ServeHTTP(rr1, req1)
	if status := rr1.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusNotFound)
	}
}

// 2. 测试请求以common.PROXY_WEB_PREFIX_PATH开头的路径，期望返回http.StatusOK。
func TestWebStaticPROXY_WEB_PREFIX_PAT(t *testing.T) {
	// Test case 2
	req2, err := http.NewRequest("GET", common.PROXY_WEB_PREFIX_PATH, nil)
	if err != nil {
		t.Fatal(err)
	}
	rr2 := httptest.NewRecorder()
	handler2 := http.HandlerFunc(WebStatic)
	handler2.ServeHTTP(rr2, req2)
	if status := rr2.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

// 3. 测试请求一个在WEB_PATH_MAP中存在的路径，期望返回http.StatusOK。
func TestWebStaticWEB_PATH_MAP(t *testing.T) {

	// Test case 3
	req3, err := http.NewRequest("GET", "/web/index.html", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr3 := httptest.NewRecorder()
	handler3 := http.HandlerFunc(WebStatic)
	handler3.ServeHTTP(rr3, req3)
	if status := rr3.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
