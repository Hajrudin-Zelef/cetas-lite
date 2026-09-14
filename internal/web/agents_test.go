package web

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func doAuthed(t *testing.T, method, url, tok string, body any) (int, map[string]any) {
	t.Helper()
	var reader io.Reader
	if body != nil {
		raw, err := json.Marshal(body)
		if err != nil {
			t.Fatal(err)
		}
		reader = bytes.NewReader(raw)
	}
	req, err := http.NewRequest(method, url, reader)
	if err != nil {
		t.Fatal(err)
	}
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	if tok != "" {
		req.Header.Set("Authorization", "Bearer "+tok)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	var out map[string]any
	_ = json.NewDecoder(resp.Body).Decode(&out)
	return resp.StatusCode, out
}

func registerUser(t *testing.T, base, username, password string) {
	t.Helper()
	body, _ := json.Marshal(map[string]string{"username": username, "password": password})
	resp, err := http.Post(base+"/api/auth/register", "application/json", bytes.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}
	resp.Body.Close()
}

func loginUser(t *testing.T, base, username, password string) string {
	t.Helper()
	body, _ := json.Marshal(map[string]string{"username": username, "password": password})
	resp, err := http.Post(base+"/api/auth/login", "application/json", bytes.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	var out map[string]any
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		t.Fatal(err)
	}
	tok, _ := out["token"].(string)
	if tok == "" {
		t.Fatal("token vide")
	}
	return tok
}

func TestAgentsCreateListDelete(t *testing.T) {
	s := newTestServerWith(t, &fakeProvider{content: "bonjour"})
	ts := httptest.NewServer(s.Handler())
	defer ts.Close()
	tok := tokenFor(t, ts.URL)

	// Famille non-agent refusee.
	code, _ := doAuthed(t, http.MethodPost, ts.URL+"/api/agents", tok,
		map[string]any{"family": "samagent-nano", "mode": "free", "message": "x"})
	if code == http.StatusOK {
		t.Fatal("une famille non-agent devrait etre refusee")
	}

	code, st := doAuthed(t, http.MethodPost, ts.URL+"/api/agents", tok,
		map[string]any{"family": "code", "mode": "standard", "message": "ecris un test"})
	if code != http.StatusOK {
		t.Fatalf("create status = %d (%v)", code, st)
	}
	id, _ := st["id"].(string)
	if id == "" {
		t.Fatal("id manquant")
	}

	// Liste.
	code, listOut := doAuthed(t, http.MethodGet, ts.URL+"/api/agents", tok, nil)
	if code != http.StatusOK {
		t.Fatalf("list status = %d", code)
	}
	agents, _ := listOut["agents"].([]any)
	if len(agents) != 1 || agents[0].(map[string]any)["id"] != id {
		t.Fatalf("liste inattendue: %v", listOut)
	}

	// Etat.
	code, _ = doAuthed(t, http.MethodGet, ts.URL+"/api/agents/"+id+"/state", tok, nil)
	if code != http.StatusOK {
		t.Fatalf("state status = %d", code)
	}

	// Message de suivi (tour probablement termine : fake provider rapide).
	// On accepte 200 (tour relance) ou 409 (tour encore en cours).
	code, _ = doAuthed(t, http.MethodPost, ts.URL+"/api/agents/"+id+"/message", tok,
		map[string]any{"message": "suite"})
	if code != http.StatusOK && code != http.StatusConflict {
		t.Fatalf("message status = %d", code)
	}

	// Stop.
	code, _ = doAuthed(t, http.MethodPost, ts.URL+"/api/agents/"+id+"/stop", tok, nil)
	if code != http.StatusOK {
		t.Fatalf("stop status = %d", code)
	}

	// Suppression.
	code, _ = doAuthed(t, http.MethodDelete, ts.URL+"/api/agents/"+id, tok, nil)
	if code != http.StatusOK {
		t.Fatalf("delete status = %d", code)
	}

	// Message sur un agent supprime -> erreur.
	code, _ = doAuthed(t, http.MethodPost, ts.URL+"/api/agents/"+id+"/message", tok,
		map[string]any{"message": "x"})
	if code == http.StatusOK {
		t.Fatal("message sur agent supprime devrait echouer")
	}

	// Etat sur un agent supprime -> 404.
	code, _ = doAuthed(t, http.MethodGet, ts.URL+"/api/agents/"+id+"/state", tok, nil)
	if code != http.StatusNotFound {
		t.Fatalf("state apres delete: attendu 404, recu %d", code)
	}
}

func TestAgentsIsolationBetweenUsers(t *testing.T) {
	s := newTestServerWith(t, &fakeProvider{content: "x"})
	ts := httptest.NewServer(s.Handler())
	defer ts.Close()
	tok := tokenFor(t, ts.URL)

	code, st := doAuthed(t, http.MethodPost, ts.URL+"/api/agents", tok,
		map[string]any{"family": "code", "mode": "standard", "message": "x"})
	if code != http.StatusOK {
		t.Fatalf("create status = %d", code)
	}
	id := st["id"].(string)

	registerUser(t, ts.URL, "bob", "motdepasse")
	tokBob := loginUser(t, ts.URL, "bob", "motdepasse")

	code, _ = doAuthed(t, http.MethodGet, ts.URL+"/api/agents/"+id+"/state", tokBob, nil)
	if code != http.StatusNotFound {
		t.Fatalf("bob ne devrait pas voir l'agent de sam (status=%d)", code)
	}
	code, listOut := doAuthed(t, http.MethodGet, ts.URL+"/api/agents", tokBob, nil)
	if code != http.StatusOK {
		t.Fatalf("list bob status = %d", code)
	}
	if agents, _ := listOut["agents"].([]any); len(agents) != 0 {
		t.Fatalf("bob devrait avoir 0 agent, recu %v", agents)
	}
}
