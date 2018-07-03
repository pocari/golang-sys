package main

import (
	"compress/gzip"
	"encoding/json"
	"io"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Enoding", "gzip")
	w.Header().Set("Content-Type", "application/json")

	source := map[string]string{
		"Hello": "World",
	}
	bytes, err := json.Marshal(source)

	if err != nil {
		panic("error encode json")
	}

	gzipw := gzip.NewWriter(w)
	defer gzipw.Close() // ここでgzipのフッタが書き込まれるらしい

	wrapper := io.MultiWriter(os.Stdout, gzipw)
	wrapper.Write(bytes)
	gzipw.Flush()
}

func Q2_3() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
