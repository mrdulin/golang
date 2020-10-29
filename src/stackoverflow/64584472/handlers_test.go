package handlers

import (
  "testing"
  "bytes"
  "io/ioutil"
  "net/http/httptest"
)


func TestGetRouterOk(t *testing.T) {
  assertResponseBody := func(t *testing.T, s *httptest.Server, expectedBody string) {
		resp, err := s.Client().Get(s.URL+"/health")
		if err != nil {
			t.Fatalf("unexpected error getting from server: %v", err)
		}
		if resp.StatusCode != 200 {
			t.Fatalf("expected a status code of 200, got %v", resp.StatusCode)
		}
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Fatalf("unexpected error reading body: %v", err)
		}
		if !bytes.Equal(body, []byte(expectedBody)) {
			t.Fatalf("response should be ok, was: %q", string(body))
		}
  }
  
	var subject IHandlerProvider = HandlerProvider{}
	router := subject.GetRouter()
  s := httptest.NewServer(router)
  defer s.Close()
  assertResponseBody(t, s, "ok")
}
