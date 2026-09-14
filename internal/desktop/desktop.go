package desktop

import (
	"context"
	"net"
	"net/http"
	"time"
)

// ServeLocal demarre le handler HTTP sur 127.0.0.1 avec un port ephemere.
// Retourne l'adresse d'ecoute (hote:port) et une fonction d'arret propre.
func ServeLocal(h http.Handler) (string, func(context.Context) error, error) {
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		return "", nil, err
	}
	srv := &http.Server{
		Handler:           h,
		ReadHeaderTimeout: 10 * time.Second,
	}
	go func() {
		_ = srv.Serve(ln)
	}()
	return ln.Addr().String(), srv.Shutdown, nil
}
