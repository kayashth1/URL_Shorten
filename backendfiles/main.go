package main

import (
    "crypto/md5"
    "encoding/hex"
    "encoding/json"
    "errors"
    "fmt"
    "net/http"
    "time"

    "github.com/rs/cors"
)

type URL struct {
    ID           string    `json:"id"`
    OriginalURL  string    `json:"original_url"`
    ShortURL     string    `json:"short_url"`
    CreationDate time.Time `json:"creation_date"`
}

var urlDB = make(map[string]URL)

func generateShortURL(originalURL string) string {
    hasher := md5.New()
    hasher.Write([]byte(originalURL))
    hash := hex.EncodeToString(hasher.Sum(nil))
    return hash[:8]
}

func createURL(originalURL string) string {
    shortURL := generateShortURL(originalURL)
    id := shortURL
    urlDB[id] = URL{
        ID:           id,
        OriginalURL:  originalURL,
        ShortURL:     shortURL,
        CreationDate: time.Now(),
    }
    return shortURL
}

func getURL(id string) (URL, error) {
    url, ok := urlDB[id]
    if !ok {
        return URL{}, errors.New("URL not found")
    }
    return url, nil
}

func ShortURLHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodOptions {
        w.WriteHeader(http.StatusOK)
        return
    }

    var data struct {
        URL string `json:"url"`
    }
    err := json.NewDecoder(r.Body).Decode(&data)
    if err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    shortURL := createURL(data.URL)
    response := struct {
        ShortURL string `json:"short_url"`
    }{ShortURL: shortURL}

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}

func redirectURLHandler(w http.ResponseWriter, r *http.Request) {
    id := r.URL.Path[len("/redirect/"):]


    url, err := getURL(id)
	fmt.Print(url.OriginalURL);
    if err != nil {
        http.Error(w, "Invalid request", http.StatusNotFound)
        return
    }


    http.Redirect(w, r, url.OriginalURL, http.StatusFound)
}

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/shorten", ShortURLHandler)
    mux.HandleFunc("/redirect/", redirectURLHandler) 

    handler := cors.Default().Handler(mux)

    fmt.Println("Starting server on port 3000...")
    http.ListenAndServe(":3000", handler)
}
