//go:build windows

package desktop

import (
	"errors"

	webview2 "github.com/jchv/go-webview2"
)

// Open ouvre une fenetre native WebView2 (Edge) sur l'URL donnee.
// Bloque jusqu'a la fermeture de la fenetre.
func Open(url, title string, width, height int, debug bool) error {
	w := webview2.NewWithOptions(webview2.WebViewOptions{
		Debug:     debug,
		AutoFocus: true,
		WindowOptions: webview2.WindowOptions{
			Title:  title,
			Width:  uint(width),
			Height: uint(height),
			Center: true,
		},
	})
	if w == nil {
		return errors.New("impossible de creer la fenetre : runtime WebView2 indisponible (Windows 10/11 l'inclut par defaut)")
	}
	defer w.Destroy()
	w.Navigate(url)
	w.Run()
	return nil
}
