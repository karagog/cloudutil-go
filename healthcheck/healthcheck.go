// Package healthcheck implements an http handler that responds with a
// healthcheck that tells you if the service has started or not.
package healthcheck

import (
	"net/http"
	"sync"
)

const (
	OKStatus       = "OK"
	StartingStatus = "STARTING"
)

var (
	mutex  sync.Mutex
	status string = StartingStatus // guarded by mutex
)

// Register registers a healthcheck handler on "/healthcheck".
func Register(m *http.ServeMux) {
	m.HandleFunc("/healthcheck", handle)
}

func handle(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	data := status
	mutex.Unlock()

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(data))
}

// SetOK marks the current health OK, which tells clients that we're ready to serve.
func SetOK() {
	mutex.Lock()
	status = OKStatus
	mutex.Unlock()
}
