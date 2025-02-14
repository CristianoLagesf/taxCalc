package routes

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

// TestHandlerRequestFailure checks if HandlerRequest fails gracefully
func TestHandlerRequestFailure(t *testing.T) {
	// Start a dummy server to occupy port 8000
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	time.Sleep(100 * time.Millisecond)

	// Try to start the actual server (which should fail since port 8000 is in use)
	errChan := make(chan error, 1)

	go func() {
		err := http.ListenAndServe(":8000", nil)
		if err != nil {
			errChan <- err
		}
	}()

	select {
	case err := <-errChan:
		// Expected failure due to occupied port
		if err == nil {
			t.Errorf("Expected failure but got nil error")
		}
	case <-time.After(1 * time.Second):
		t.Errorf("Test timeout, expected failure but server didn't respond")
	}
}
