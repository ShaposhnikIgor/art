package server

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"runtime"

	"leart.com/art/cypher"
)

var tmpl *template.Template

// Initialize the template
func init() {
	_, filename, _, _ := runtime.Caller(0)
	path := filepath.Join(filepath.Dir(filename), "templates", "index.html")
	tmpl = template.Must(template.ParseFiles(path))
}

// NoCacheMiddleware ensures no caching for responses
func NoCacheMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "no-store, no-cache, must-revalidate, max-age=0")
		w.Header().Set("Pragma", "no-cache")
		w.Header().Set("Expires", "0")
		next.ServeHTTP(w, r)
	})
}

// NewHandler sets up the HTTP handlers and routes
func NewHandler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/decode", decodeHandler)
	mux.HandleFunc("/encode", encodeHandler)
	staticHandler := http.StripPrefix("/static/", http.FileServer(http.Dir("server/static")))
	mux.Handle("/static/", NoCacheMiddleware(staticHandler))

	return mux
}

// indexHandler handles requests to the root URL
func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		log.Printf("Method not allowed: %s", r.Method)
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	// Set HTTP status to 200 before sending the response body
	w.WriteHeader(http.StatusOK)
	if err := tmpl.Execute(w, nil); err != nil {
		logError(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}

// decodeHandler handles requests to decode text
func decodeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		log.Printf("Method not allowed: %s", r.Method)
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	encodedText := r.FormValue("text")
	if encodedText == "" {
		log.Printf("Bad request: empty text")
		http.Error(w, "Bad request: empty text", http.StatusBadRequest)
		return
	}
	// Validate the encoded text
	if cypher.If_Decod(encodedText) || !cypher.IsBalanced([]byte(encodedText)) {
		log.Printf("Bad request: invalid encoded text")
		http.Error(w, "Bad request: invalid encoded text", http.StatusBadRequest)
		return
	}
	// Attempt to decode the text
	decodedText, err := cypher.Decod_Art(encodedText)
	if err != nil {
		log.Printf("Error decoding text: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Encoded text is valid, return HTTP 202
	w.WriteHeader(http.StatusAccepted)
	if err := tmpl.Execute(w, map[string]string{"Result": decodedText}); err != nil {
		logError(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}

// encodeHandler handles requests to encode text
func encodeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		log.Printf("Method not allowed: %s", r.Method)
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	message := r.FormValue("text")
	if message == "" {
		log.Printf("Bad request: empty text")
		http.Error(w, "Bad request: empty text", http.StatusBadRequest)
		return
	}
	encodedText := cypher.Encod_Art(message)
	if err := tmpl.Execute(w, map[string]string{"Result": encodedText}); err != nil {
		logError(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}

// logError logs the error with the file and line number
func logError(err error) {
	_, file, line, _ := runtime.Caller(1)
	log.Printf("%s:%d: %v", file, line, err)
}
