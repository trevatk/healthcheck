// Package port exposed endpoints
package port

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"go.uber.org/zap"
)

// HTTPServer exposed endpoints
type HTTPServer struct {
	log *zap.SugaredLogger
}

// NewHTTPServer create new http server instance
func NewHTTPServer(logger *zap.Logger) *HTTPServer {
	return &HTTPServer{log: logger.Named("http server").Sugar()}
}

// NewRouter chi router implementation of http handler
func NewRouter(httpServer *HTTPServer) *chi.Mux {

	r := chi.NewRouter()

	r.Get("/health", httpServer.health)

	return r
}

func (h *HTTPServer) health(w http.ResponseWriter, _ *http.Request) {

	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("OK"))
	if err != nil {
		h.log.Errorf("error encoding health check response %v", err)
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
	}
}
