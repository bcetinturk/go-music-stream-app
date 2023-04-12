package main

import (
	"net/http"

	"mediaservice/handlers"
)

func streamAudio(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		file, fileInfo, err := handlers.GetAudio(r.URL.Path)
		defer file.Close()
		if err != nil {
			switch e := err.(type) {
			case *handlers.RequestError:
				http.Error(w, e.Error(), e.StatusCode)
			}
		}

		w.Header().Set("Content-Type", "audio/mpeg")
		http.ServeContent(w, r, "", fileInfo.ModTime(), file)
		return
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

func main() {
	http.HandleFunc("/audio/", streamAudio)

	http.ListenAndServe(":8080", nil)
}
