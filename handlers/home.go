package handlers

import (
	"go_fullstack/views"
	"net/http"
)

func (h *Handler) HandleHome(w http.ResponseWriter, r *http.Request) {
	views.Home().Render(r.Context(), w)
}
