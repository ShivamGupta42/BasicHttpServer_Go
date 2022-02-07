package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

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
	t.Run("returns shivam's score", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/players/shivam", nil)
		response := httptest.NewRecorder()

		PlayerServer(response, request)

		got := response.Body.String()
		want := "10"
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

}
