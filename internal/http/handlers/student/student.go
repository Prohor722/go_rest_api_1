package student

import (
	"log/slog"
	"net/http"
)

func New() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request){
		slog.Info("creating a student")
		w.Write([]byte("welcome to students api..."))
	}
}