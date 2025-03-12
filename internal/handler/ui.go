package handler

import (
	"net/http"

	mainpage "github.com/koyo-os/murocami/internal/view/page"
)

func (h *Handler) MainPage(w http.ResponseWriter, r *http.Request) {
	err := mainpage.Mainpage().Render(r.Context(), w)
	if err != nil{
		h.Logger.Errorf("error render page: %v",err)
		http.Error(w, "cant render page: " + err.Error(), http.StatusInternalServerError)
		return
	}
}