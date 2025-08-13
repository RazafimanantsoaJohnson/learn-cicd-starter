package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	// req, _ := http.NewRequest("GET", "https://google.com", nil)
	cases := []http.Header{
		http.Header{"Authorization": []string{"ApiKey Bomboclaat"}},
		http.Header{"Authorization": []string{"ApiKey Ambilamaa"}},
	}

	for c := range cases {
		_, err := GetAPIKey(cases[c])
		if err != nil {
			t.Errorf("An error occured: %v", err)
		}
	}
}
