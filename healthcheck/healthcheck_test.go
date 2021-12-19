package healthcheck

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthCheck(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(handle))
	defer ts.Close()

	url := fmt.Sprintf("http://%s/%s", ts.Listener.Addr().String(), "healthcheck")
	resp, err := http.Get(url)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	if got, want := string(data), StartingStatus; got != want {
		t.Fatalf("Got %v, want %v", got, want)
	}

	// Now mark it OK and check again.
	SetOK()
	resp, err = http.Get(url)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	data, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	if got, want := string(data), OKStatus; got != want {
		t.Fatalf("Got %v, want %v", got, want)
	}
}
