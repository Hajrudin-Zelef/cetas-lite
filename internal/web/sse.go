package web

import (
	"encoding/json"
	"net/http"
	"sync"
	"time"
)

func sseStart(w http.ResponseWriter) (*sync.Mutex, func(), bool) {
	flusher, ok := w.(http.Flusher)
	if !ok {
		writeError(w, http.StatusInternalServerError, "streaming non supporte")
		return nil, nil, false
	}
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache, no-transform")
	w.Header().Set("X-Accel-Buffering", "no")
	w.Header().Set("Connection", "keep-alive")
	w.WriteHeader(http.StatusOK)

	mu := &sync.Mutex{}
	done := make(chan struct{})
	go func() {
		t := time.NewTicker(4 * time.Second)
		defer t.Stop()
		for {
			select {
			case <-done:
				return
			case <-t.C:
				mu.Lock()
				_, err := w.Write([]byte(": ping\n\n"))
				flusher.Flush()
				mu.Unlock()
				if err != nil {
					return
				}
			}
		}
	}()
	stop := func() { close(done) }
	return mu, stop, true
}

func sseEmitter(w http.ResponseWriter, mu *sync.Mutex) func(map[string]any) bool {
	flusher, _ := w.(http.Flusher)
	return func(obj map[string]any) bool {
		b, err := json.Marshal(obj)
		if err != nil {
			return true
		}
		mu.Lock()
		defer mu.Unlock()
		if _, err := w.Write([]byte("data: " + string(b) + "\n\n")); err != nil {
			return false
		}
		if flusher != nil {
			flusher.Flush()
		}
		return true
	}
}
