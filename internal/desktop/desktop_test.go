package desktop

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"runtime"
	"strings"
	"testing"
	"time"
)

func TestServeLocal(t *testing.T) {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "bonjour-bureau")
	})
	addr, shutdown, err := ServeLocal(h)
	if err != nil {
		t.Fatalf("ServeLocal: %v", err)
	}
	if !strings.HasPrefix(addr, "127.0.0.1:") {
		t.Fatalf("adresse inattendue: %q (attendu 127.0.0.1:<port>)", addr)
	}
	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get("http://" + addr + "/")
	if err != nil {
		t.Fatalf("GET: %v", err)
	}
	body, _ := io.ReadAll(resp.Body)
	resp.Body.Close()
	if string(body) != "bonjour-bureau" {
		t.Fatalf("corps inattendu: %q", body)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := shutdown(ctx); err != nil {
		t.Fatalf("shutdown: %v", err)
	}
}

func TestOpenNonWindows(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("fenetre native non testable en CI")
	}
	// Sur les plateformes non-Windows, Open doit echouer proprement.
	if err := Open("http://127.0.0.1:1/", "Test", 800, 600, false); err == nil {
		t.Fatal("Open aurait du echouer hors Windows")
	}
}
