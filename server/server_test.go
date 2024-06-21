package server

import (
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"leart.com/art/cypher" // Импортируем пакет с функцией if_decod
)

func TestDecodeHandler(t *testing.T) {
	validEncodedText := "[5 #][5 -_]-[5 #]" // Замените на реальную закодированную строку

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

func TestDecodeHandlerInvalidInput(t *testing.T) {
	invalidEncodedText := "" // Пустая строка, которая не может быть декодирована

	// Проверяем, что пустая строка не является декодированным сообщением
	if !cypher.If_Decod(invalidEncodedText) { // Используем функцию из пакета cypher
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
