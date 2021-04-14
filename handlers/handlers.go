package handlers

import (
	"net/http"
	"strings"
	"testt/db"
	"time"
)

type httpHandler struct {
	httpClient *http.Client
	Store      db.Datastore
}

// NewHandler returns a new http handler
func NewHandler(store db.Datastore) *httpHandler {
	return &httpHandler{
		Store: store,
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (h *httpHandler) HandleShortUrl(w http.ResponseWriter, r *http.Request) {
	code := strings.Replace(r.URL.RequestURI(), "/", "", -1)
	if code == "graphiql" {
		return
	}
	// search for code in db
	urlMap, err := h.Store.Get(code)
	if err != nil {
		http.Error(w, "invalid short url", http.StatusNotFound)
		return
	}

	http.Redirect(w, r, urlMap.URL, 301)
	return
}
