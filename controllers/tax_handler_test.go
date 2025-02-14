package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestTaxHandler checks different scenarios for the TaxHandler()
func TestTaxHandler(t *testing.T) {
	tests := []struct {
		name           string
		queryString    string
		expectedStatus int
	}{

		{
			name:           "Missing income parameter",
			queryString:    "?year=2019",
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "Invalid income value",
			queryString:    "?income=invalid&year=2019",
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "Negative income value",
			queryString:    "?income=-5000&year=2019",
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "Missing year parameter",
			queryString:    "?income=60000",
			expectedStatus: http.StatusBadRequest,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			req, err := http.NewRequest("GET", "/tax"+test.queryString, nil)
			if err != nil {
				t.Fatalf("Could not create request: %v", err)
			}

			// Create a response recorder to capture the response
			res := httptest.NewRecorder()
			handler := http.HandlerFunc(TaxHandler)

			// Call the handler
			handler.ServeHTTP(res, req)

			if status := res.Code; status != test.expectedStatus {
				t.Errorf("Expected status code %d, got %d", test.expectedStatus, status)
			}
		})
	}
}
