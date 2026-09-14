//go:build !windows

package desktop

import "errors"

// Open n'est supporte que sous Windows (WebView2).
func Open(url, title string, width, height int, debug bool) error {
	return errors.New("le mode bureau n'est disponible que sous Windows")
}
