package server

import (
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"leart.com/art/cypher" // Import the package with the If_Decod function
)

// TestDecodeHandler tests the decodeHandler with valid input
func TestDecodeHandler(t *testing.T) {
	validEncodedText := "[5 #][5 -_]-[5 #]" // Replace with actual encoded string

	req, err := http.NewRequest("POST", "/decode", strings.NewReader("text="+validEncodedText))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(decodeHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusAccepted {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusAccepted)
	}
}

// TestIndexHandler tests the indexHandler
func TestIndexHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(indexHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected, _ := os.ReadFile("server/templates/index.html")
	if !strings.Contains(rr.Body.String(), string(expected)) {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

// TestDecodeHandlerInvalidInput tests the decodeHandler with invalid input
func TestDecodeHandlerInvalidInput(t *testing.T) {
	invalidEncodedText := "" // Empty string which cannot be decoded

	// Ensure the empty string is not a decoded message
	if !cypher.If_Decod(invalidEncodedText) { // Use the function from the cypher package
		t.Fatalf("Expected invalidEncodedText to be undecoded, but it is already decoded")
	}

	req, err := http.NewRequest("POST", "/decode", strings.NewReader("text="+invalidEncodedText))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(decodeHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("DecodeHandler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}
}

// TestDecodeEndpointInvalidInput tests the decodeHandler with invalid input and endpoint
func TestDecodeEndpointInvalidInput(t *testing.T) {
	invalidEncodedText := ""
	expectedStatusCode := http.StatusBadRequest

	req, err := http.NewRequest("POST", "/decoder", strings.NewReader("text="+invalidEncodedText))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(decodeHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != expectedStatusCode {
		t.Errorf("DecodeHandler returned wrong status code: got %v want %v", status, expectedStatusCode)
	}
}
