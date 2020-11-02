package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandlersStatus(t *testing.T) {
	handler := Handler()

	vs := []struct {
		method         string
		path           string
		expectedStatus int
	}{
		{
			"GET",
			"/count",
			http.StatusOK,
		},
		{
			"GET",
			"/",
			http.StatusNotFound,
		},
		{
			"GET",
			"/counter",
			http.StatusNotFound,
		},
		{
			"GET",
			"/count/",
			http.StatusNotFound,
		},
	}
	for _, v := range vs {
		recorder := httptest.NewRecorder()
		request, _ := http.NewRequest(v.method, v.path, nil)
		handler.ServeHTTP(recorder, request)
		actualStatus := recorder.Code
		if actualStatus != v.expectedStatus {
			t.Errorf("Calling %s:%s returned status %d. Expected: %d.", v.method, v.path, actualStatus, v.expectedStatus)
		}
	}

}
