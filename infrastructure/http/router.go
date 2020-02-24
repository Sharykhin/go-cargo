package http

import (
	"net/http"
	"github.com/gorilla/mux"
)

func router() http.Handler {
	r := mux.NewRouter()

	return r
}
