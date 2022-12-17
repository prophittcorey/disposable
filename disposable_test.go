package disposable

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDomains(t *testing.T) {
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "10minutemail.ru\nfoobar.com")
	}))

	defer svr.Close()

	Sources = map[string][]byte{
		svr.URL: []byte{},
	}

	ds := Domains()

	if len(ds) != 2 {
		t.Fatalf("failed to find two domains; got %d", len(ds))
	}
}

func TestCheck(t *testing.T) {
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "10minutemail.ru\nfoobar.com")
	}))

	defer svr.Close()

	Sources = map[string][]byte{
		svr.URL: []byte{},
	}

	if disposable, err := Check("10minutemail.ru"); err != nil || !disposable {
		t.Fatalf("failed to identify an disposable domain; 10minutemail")
	}

	if disposable, err := Check("foobar.com"); err != nil || !disposable {
		t.Fatalf("failed to identify a disposable domain; got foobar")
	}
}
