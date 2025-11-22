package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// HandleSearch Aqui vamos bater no google api
func HandleSearch(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Bateu aqui")
	defer r.Body.Close()

	// Le a string enviada
	b, _ := io.ReadAll(r.Body)

	fmt.Println(string(b))

	query := strings.TrimSpace(string(b))
	googleURL := "https://www.googleapis.com/books/v1/volumes?q=" + url.QueryEscape(query)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	gReq, _ := http.NewRequestWithContext(ctx, http.MethodGet, googleURL, nil)
	resp, _ := http.DefaultClient.Do(gReq)
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(body)
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}
