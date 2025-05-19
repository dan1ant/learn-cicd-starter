package auth

import (
	"errors"
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	cases := []struct {
		headers map[string][]string
		want    string
		err     error
	}{
		{
			headers: http.Header{
				"Authorization": []string{"ApiKey cb0851a1-7893-47ab-a87d-6fee3e843429"},
			},
			want: "cb0851a1-7893-47ab-a87d-6fee3e843429",
		},
		{
			headers: http.Header{},
			err:     ErrNoAuthHeaderIncluded,
		},
		{
			headers: http.Header{
				"Authorization": []string{"ApiKey"},
			},
			err: errors.New("malformed authorization header"),
		},
	}

	for i, c := range cases {
		if got, err := GetAPIKey(c.headers); err != nil && !errors.Is(err, c.err) {
			t.Errorf("Test %d failed: got [%v] want [%v]", i+1, err, c.err)
		} else if got != c.want {
			t.Errorf("Test %d failed: got [%s] want [%s]", i+1, got, c.want)
		}
	}
}
