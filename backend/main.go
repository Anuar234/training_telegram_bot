package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

type Video struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	YoutubeID   string `json:"youtubeId"`
	Description string `json:"description,omitempty"`
}

var videos = []Video{
	{"1", "Intro: как пользоваться тренажёром", "dQw4w9WgXcQ", "Короткое вводное видео"},
	{"2", "Тренировка 1 — Разминка", "VIDEO_ID_2", ""},
}

func videoHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(videos)
}

func main() {
	http.HandleFunc("/api/videos", videoHandler)

	dist := "./frontend/dist"
	fs := http.FileServer(http.Dir(dist))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		p := filepath.Join(dist, r.URL.Path)
		if _, err := os.Stat(p); os.IsNotExist(err) {
			http.ServeFile(w, r, filepath.Join(dist, "index.html"))
			return
		}
		fs.ServeHTTP(w, r)
	})

	port := "8080"

	log.Printf("Server listening on : %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}