package server

import (
	"io/ioutil"
	"net/http"
	"testing"
)

func TestHealthCheck(t *testing.T) {
	startTestServer(t)

	healthCheckAPI := TestAddr + "/health"
	resp, err := http.Get(healthCheckAPI)
	if err != nil {
		t.Errorf("get health check api error: %v\n", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("read health check response body error: %v\n", err)
	}
	expected := `{"message":"health"}`
	if string(expected) != string(body) {
		t.Errorf("response body error, expected is %v but got %v\n", expected, body)
	}
}
