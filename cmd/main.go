package main

import (
	"log"
	"net/http"
)

func main () {
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("ok"))
	})

	log.Println("Listening on port :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err);
	}
};

