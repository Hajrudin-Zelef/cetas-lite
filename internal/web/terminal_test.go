package web

import (
	"encoding/base64"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestTerminalCreateInputDelete(t *testing.T) {
	s := newTestServerWith(t, &fakeProvider{content: "x"})
	ts := httptest.NewServer(s.Handler())
	defer ts.Close()
	tok := tokenFor(t, ts.URL)

	// Creation.
	code, out := doAuthed(t, http.MethodPost, ts.URL+"/api/terminal", tok, map[string]any{})
	if code != http.StatusOK {
		t.Fatalf("create status = %d (%v)", code, out)
	}
	id, _ := out["id"].(string)
	if id == "" {
		t.Fatal("id manquant")
	}

	// Liste.
	code, listOut := doAuthed(t, http.MethodGet, ts.URL+"/api/terminal", tok, nil)
	if code != http.StatusOK {
		t.Fatalf("list status = %d", code)
	}
	sess, _ := listOut["sessions"].([]any)
	if len(sess) != 1 {
		t.Fatalf("sessions inattendues: %v", listOut)
	}

	// Entree.
	code, _ = doAuthed(t, http.MethodPost, ts.URL+"/api/terminal/"+id+"/input", tok,
		map[string]any{"data": "echo ok\n"})
	if code != http.StatusOK {
		t.Fatalf("input status = %d", code)
	}

	// Resize.
	code, _ = doAuthed(t, http.MethodPost, ts.URL+"/api/terminal/"+id+"/resize", tok,
		map[string]any{"cols": 100, "rows": 30})
	if code != http.StatusOK {
		t.Fatalf("resize status = %d", code)
	}

	// Flux SSE : on lit jusqu'a voir la sortie.
	req, _ := http.NewRequest(http.MethodGet, ts.URL+"/api/terminal/"+id+"/stream", nil)
	req.Header.Set("Authorization", "Bearer "+tok)
	client := &http.Client{Timeout: 15 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("stream status = %d", resp.StatusCode)
	}
	buf := make([]byte, 0, 4096)
	tmp := make([]byte, 1024)
	deadline := time.Now().Add(10 * time.Second)
	found := false
	for time.Now().Before(deadline) && !found {
		n, rerr := resp.Body.Read(tmp)
		if n > 0 {
			buf = append(buf, tmp[:n]...)
			// Extraire les payloads data:.
			for _, line := range strings.Split(string(buf), "\n") {
				line = strings.TrimSpace(line)
				if !strings.HasPrefix(line, "data: ") {
					continue
				}
				var ev map[string]any
				if err := json.Unmarshal([]byte(strings.TrimPrefix(line, "data: ")), &ev); err != nil {
					continue
				}
				if b64, ok := ev["output"].(string); ok {
					raw, _ := base64.StdEncoding.DecodeString(b64)
					if strings.Contains(string(raw), "ok") {
						found = true
						break
					}
				}
			}
		}
		if rerr != nil {
			break
		}
	}
	if !found {
		t.Fatalf("sortie 'ok' non vue dans le flux (recu %d octets)", len(buf))
	}

	// Suppression.
	code, _ = doAuthed(t, http.MethodDelete, ts.URL+"/api/terminal/"+id, tok, nil)
	if code != http.StatusOK {
		t.Fatalf("delete status = %d", code)
	}
	code, listOut = doAuthed(t, http.MethodGet, ts.URL+"/api/terminal", tok, nil)
	if sess, _ := listOut["sessions"].([]any); len(sess) != 0 {
		t.Fatal("la session devrait etre supprimee")
	}
}

func TestTerminalCwdForbidden(t *testing.T) {
	s := newTestServerWith(t, &fakeProvider{content: "x"})
	ts := httptest.NewServer(s.Handler())
	defer ts.Close()
	tok := tokenFor(t, ts.URL)

	code, _ := doAuthed(t, http.MethodPost, ts.URL+"/api/terminal", tok,
		map[string]any{"cwd": "/etc"})
	if code != http.StatusForbidden && code != http.StatusBadRequest {
		t.Fatalf("cwd hors racines: attendu 403/400, recu %d", code)
	}
	code, _ = doAuthed(t, http.MethodGet, ts.URL+"/api/terminal/xyz/stream", tok, nil)
	if code != http.StatusNotFound {
		t.Fatalf("stream inconnu: attendu 404, recu %d", code)
	}
}
