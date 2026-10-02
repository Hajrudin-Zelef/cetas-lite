package web

import (
	"bytes"
	"encoding/json"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"unicode/utf8"

	"cetas-lite/internal/attach"
)

func multipartFile(t *testing.T, field, name string, data []byte) (*bytes.Buffer, string) {
	t.Helper()
	var body bytes.Buffer
	mw := multipart.NewWriter(&body)
	fw, err := mw.CreateFormFile(field, name)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := fw.Write(data); err != nil {
		t.Fatal(err)
	}
	if err := mw.Close(); err != nil {
		t.Fatal(err)
	}
	return &body, mw.FormDataContentType()
}

func uploadAttachment(t *testing.T, h http.Handler, token, name string, data []byte) *httptest.ResponseRecorder {
	t.Helper()
	body, ct := multipartFile(t, "file", name, data)
	req := httptest.NewRequest(http.MethodPost, "/api/chat/attach", body)
	req.Header.Set("Content-Type", ct)
	req.Header.Set("Authorization", "Bearer "+token)
	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, req)
	return rec
}

func newAttachServer(t *testing.T) (*Server, string) {
	t.Helper()
	s := newTestServerWith(t, &fakeProvider{content: "ok"})
	s.engine.SetAttachments(attach.New(t.TempDir(), 20<<20))
	ts := httptest.NewServer(s.Handler())
	t.Cleanup(ts.Close)
	return s, tokenFor(t, ts.URL)
}

func TestAttachUploadGetDelete(t *testing.T) {
	s, tok := newAttachServer(t)
	h := s.Handler()

	rec := uploadAttachment(t, h, tok, "cours.md", []byte("contenu du cours"))
	if rec.Code != http.StatusOK {
		t.Fatalf("upload status = %d (%s)", rec.Code, rec.Body.String())
	}
	var a map[string]any
	if err := json.Unmarshal(rec.Body.Bytes(), &a); err != nil {
		t.Fatal(err)
	}
	id, _ := a["id"].(string)
	if id == "" || a["kind"] != "text" {
		t.Fatalf("reponse upload = %v", a)
	}

	req := httptest.NewRequest(http.MethodGet, "/api/chat/attach/"+id, nil)
	req.Header.Set("Authorization", "Bearer "+tok)
	got := httptest.NewRecorder()
	h.ServeHTTP(got, req)
	if got.Code != http.StatusOK || got.Body.String() != "contenu du cours" {
		t.Fatalf("get status = %d body = %q", got.Code, got.Body.String())
	}

	req = httptest.NewRequest(http.MethodDelete, "/api/chat/attach/"+id, nil)
	req.Header.Set("Authorization", "Bearer "+tok)
	del := httptest.NewRecorder()
	h.ServeHTTP(del, req)
	if del.Code != http.StatusOK {
		t.Fatalf("delete status = %d", del.Code)
	}
}

func TestAttachRejectsUnsupportedAcceptsImages(t *testing.T) {
	s, tok := newAttachServer(t)
	h := s.Handler()
	if rec := uploadAttachment(t, h, tok, "doc.docx", []byte("x")); rec.Code != http.StatusBadRequest {
		t.Fatalf("docx status = %d", rec.Code)
	}
	rec := uploadAttachment(t, h, tok, "photo.png", []byte("\x89PNG"))
	if rec.Code != http.StatusOK {
		t.Fatalf("image status = %d (%s)", rec.Code, rec.Body.String())
	}
	var a map[string]any
	_ = json.Unmarshal(rec.Body.Bytes(), &a)
	if a["kind"] != "image" {
		t.Fatalf("kind image attendu: %v", a)
	}
}

func TestAttachUnauthorized(t *testing.T) {
	s, _ := newAttachServer(t)
	rec := uploadAttachment(t, s.Handler(), "", "a.md", []byte("x"))
	if rec.Code != http.StatusUnauthorized {
		t.Fatalf("status = %d", rec.Code)
	}
}

func getSnippet(t *testing.T, h http.Handler, tok, id string) *httptest.ResponseRecorder {
	t.Helper()
	req := httptest.NewRequest(http.MethodGet, "/api/chat/attach/"+id+"?snippet=1", nil)
	req.Header.Set("Authorization", "Bearer "+tok)
	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, req)
	return rec
}

func TestAttachSnippet(t *testing.T) {
	s, tok := newAttachServer(t)
	h := s.Handler()

	rec := uploadAttachment(t, h, tok, "notes.md", []byte("ligne un\nligne deux"))
	if rec.Code != http.StatusOK {
		t.Fatalf("upload status = %d (%s)", rec.Code, rec.Body.String())
	}
	var a map[string]any
	_ = json.Unmarshal(rec.Body.Bytes(), &a)
	id, _ := a["id"].(string)

	got := getSnippet(t, h, tok, id)
	if got.Code != http.StatusOK {
		t.Fatalf("snippet status = %d (%s)", got.Code, got.Body.String())
	}
	if ct := got.Header().Get("Content-Type"); ct != "text/plain; charset=utf-8" {
		t.Fatalf("content-type = %q", ct)
	}
	if got.Body.String() != "ligne un\nligne deux" {
		t.Fatalf("snippet = %q", got.Body.String())
	}
}

func TestAttachSnippetTruncateUTF8(t *testing.T) {
	s, tok := newAttachServer(t)
	h := s.Handler()

	// 700 runes avec des multi-octets (é/€) : coupe à 600 runes + ellipse.
	long := strings.Repeat("é€x", 234) // 702 runes
	rec := uploadAttachment(t, h, tok, "long.txt", []byte(long))
	var a map[string]any
	_ = json.Unmarshal(rec.Body.Bytes(), &a)
	id, _ := a["id"].(string)

	got := getSnippet(t, h, tok, id)
	body := got.Body.String()
	r := []rune(body)
	if len(r) != 601 || r[600] != '…' {
		t.Fatalf("snippet tronqué attendu à 601 runes + …, got %d runes", len(r))
	}
	if !utf8.ValidString(body) {
		t.Fatal("snippet invalide UTF-8")
	}
}

func TestAttachSnippetImageRejected(t *testing.T) {
	s, tok := newAttachServer(t)
	h := s.Handler()
	rec := uploadAttachment(t, h, tok, "photo.png", []byte("\x89PNG"))
	var a map[string]any
	_ = json.Unmarshal(rec.Body.Bytes(), &a)
	id, _ := a["id"].(string)
	if got := getSnippet(t, h, tok, id); got.Code != http.StatusBadRequest {
		t.Fatalf("image snippet status = %d", got.Code)
	}
}

func getText(t *testing.T, h http.Handler, tok, id string) *httptest.ResponseRecorder {
	t.Helper()
	req := httptest.NewRequest(http.MethodGet, "/api/chat/attach/"+id+"?text=1", nil)
	req.Header.Set("Authorization", "Bearer "+tok)
	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, req)
	return rec
}

func TestAttachTextFull(t *testing.T) {
	s, tok := newAttachServer(t)
	h := s.Handler()
	// 700 runes : le snippet coupe a 600, ?text=1 renvoie tout.
	long := strings.Repeat("é", 700)
	rec := uploadAttachment(t, h, tok, "full.txt", []byte(long))
	var a map[string]any
	_ = json.Unmarshal(rec.Body.Bytes(), &a)
	id, _ := a["id"].(string)

	got := getText(t, h, tok, id)
	if got.Code != http.StatusOK {
		t.Fatalf("text status = %d (%s)", got.Code, got.Body.String())
	}
	if got.Body.String() != long {
		t.Fatalf("texte complet attendu (%d runes), got %d", len([]rune(long)), len([]rune(got.Body.String())))
	}
	if !utf8.ValidString(got.Body.String()) {
		t.Fatal("texte invalide UTF-8")
	}
}

func TestAttachTextImageRejected(t *testing.T) {
	s, tok := newAttachServer(t)
	h := s.Handler()
	rec := uploadAttachment(t, h, tok, "photo.png", []byte("\x89PNG"))
	var a map[string]any
	_ = json.Unmarshal(rec.Body.Bytes(), &a)
	id, _ := a["id"].(string)
	if got := getText(t, h, tok, id); got.Code != http.StatusBadRequest {
		t.Fatalf("image text status = %d", got.Code)
	}
}

func TestAttachTextEmptyNoContent(t *testing.T) {
	s, tok := newAttachServer(t)
	h := s.Handler()
	rec := uploadAttachment(t, h, tok, "vide.txt", []byte("   \n  "))
	var a map[string]any
	_ = json.Unmarshal(rec.Body.Bytes(), &a)
	id, _ := a["id"].(string)
	if got := getText(t, h, tok, id); got.Code != http.StatusNoContent {
		t.Fatalf("text vide: status = %d (attendu 204)", got.Code)
	}
}

func TestAttachSnippetAuthAndNotFound(t *testing.T) {
	s, _ := newAttachServer(t)
	h := s.Handler()
	req := httptest.NewRequest(http.MethodGet, "/api/chat/attach/abc?snippet=1", nil)
	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, req)
	if rec.Code != http.StatusUnauthorized {
		t.Fatalf("sans token: status = %d", rec.Code)
	}
	req = httptest.NewRequest(http.MethodGet, "/api/chat/attach/nope?snippet=1", nil)
	req.Header.Set("Authorization", "Bearer x")
	rec = httptest.NewRecorder()
	h.ServeHTTP(rec, req)
	if rec.Code != http.StatusUnauthorized && rec.Code != http.StatusNotFound {
		t.Fatalf("id inconnu: status = %d", rec.Code)
	}
}
