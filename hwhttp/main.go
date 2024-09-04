package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)
func loggingMiddleware(n http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Получен запрос: Method=%s, Path=%s", r.Method, r.URL.Path)
		n.ServeHTTP(w, r)
	})
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, World!")
	})

	http.HandleFunc("/greet", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello!")
	})

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		if name == "" {
			name = "Killer"
		}
		fmt.Fprintf(w, "Hello, %s!", name)
	})


	http.HandleFunc("/json", func(w http.ResponseWriter, r *http.Request) {
		response := map[string]string{"message": "Hello, JSON!"}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	})

	http.HandleFunc("/post", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Неверный метод запроса", http.StatusMethodNotAllowed)
			return
		}
		fmt.Fprintln(w, "Получен POST-запрос")
	})

	http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Невозможно прочитать тело", http.StatusBadRequest)
			return
		}
		fmt.Fprintln(w, string(body))
	})

	http.HandleFunc("/form", func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			http.Error(w, "Невозможно разобрать форму", http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(r.Form)
	})


	http.HandleFunc("/redirect", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/greet", http.StatusFound)
	})

	http.HandleFunc("/headers", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("X-Custom-Header", "CustomHeaderValue")
		fmt.Fprintln(w, `{"message": "Headers set!"}`)
	})

	http.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, "404 Not Found")
	})

	http.HandleFunc("/upload", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}
		file, _, err := r.FormFile("file")
		if err != nil {
			http.Error(w, "Unable to retrieve file", http.StatusBadRequest)
			return
		}
		defer file.Close()

		out, err := os.Create("uploaded_file")
		if err != nil {
			http.Error(w, "Unable to create file", http.StatusInternalServerError)
			return
		}
		defer out.Close()

		_, err = io.Copy(out, file)
		if err != nil {
			http.Error(w, "Unable to save file", http.StatusInternalServerError)
			return
		}

		fmt.Fprintln(w, "File uploaded successfully")
	})


	http.HandleFunc("/set-cookie", func(w http.ResponseWriter, r *http.Request) {
		expiration := time.Now().Add(365 * 24 * time.Hour)
		cookie := http.Cookie{Name: "username", Value: "GoUser", Expires: expiration}
		http.SetCookie(w, &cookie)
		fmt.Fprintln(w, "Cookie set!")
	})

	http.HandleFunc("/get-cookie", func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("username")
		if err != nil {
			http.Error(w, "No cookie found", http.StatusNotFound)
			return
		}
		fmt.Fprintf(w, "Cookie: %s=%s", cookie.Name, cookie.Value)
	})

	loggedRouter := loggingMiddleware(http.DefaultServeMux)

	
	fmt.Println("Starting server on :8080")
	err := http.ListenAndServe(":8080", loggedRouter)
	if err != nil {
		log.Fatal(err)
	}
}
