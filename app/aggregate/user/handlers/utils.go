package handlers

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func getUserIDFromURL(r *http.Request) (uint64, error) {
	const userIDURLKey = "id"

	return strconv.ParseUint(chi.URLParam(r, userIDURLKey), 10, 64)
}
