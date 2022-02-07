package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

type StubPlayerStore struct {
	scores map[string]string
}

func (s *StubPlayerStore) GetPlayerScore(name string) string {
	return s.scores[name]
}

func TestPlayerServer(t *testing.T) {
	/* First Test
	mockHttpRequest, _ := http.NewRequest(http.MethodGet, "/players/Shivam", nil)
	responseSpy := httptest.NewRecorder()

	PlayerServer(responseSpy, mockHttpRequest)
	got := responseSpy.Body.String()
	want := "20"

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
	*/

	// Second Test
	//t.Run("returns shivam's score", func(t *testing.T) {
	//	request, _ := http.NewRequest(http.MethodGet, "/players/shivam", nil)
	//	response := httptest.NewRecorder()
	//
	//	server := PlayerServer{}
	//	server.ServeHTTP(response, request)
	//	got := response.Body.String()
	//	want := "10"
	//	if got != want {
	//		t.Errorf("got %q, want %q", got, want)
	//	}
	//})

	//Third test
	t.Run("returns shivam's score", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/players/shivam", nil)
		response := httptest.NewRecorder()

		server := PlayerServer{&StubPlayerStore{map[string]string{"shivam": "10", "anu": "20"}}}
		server.ServeHTTP(response, request)
		got := response.Body.String()
		want := "10"
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

}
