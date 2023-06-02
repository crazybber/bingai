package api

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSysConf(t *testing.T) {
	req, err := http.NewRequest("GET", "/sysconf", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(SysConf)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := `{"IsSysCK":true,"IsAuth":false}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}
