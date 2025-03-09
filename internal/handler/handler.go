package handler

import "net/http"

func WebHookHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method must ve post", http.StatusMethodNotAllowed)
	}
}