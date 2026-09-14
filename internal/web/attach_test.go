package web

import (
	"bytes"
	"encoding/json"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"testing"

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

func TestAttachRejectsUnsupportedAndImages(t *testing.T) {
	s, tok := newAttachServer(t)
	h := s.Handler()
	if rec := uploadAttachment(t, h, tok, "doc.docx", []byte("x")); rec.Code != http.StatusBadRequest {
		t.Fatalf("docx status = %d", rec.Code)
	}
	if rec := uploadAttachment(t, h, tok, "photo.png", []byte("x")); rec.Code != http.StatusBadRequest {
		t.Fatalf("image status = %d", rec.Code)
	}
}

func TestAttachUnauthorized(t *testing.T) {
	s, _ := newAttachServer(t)
	rec := uploadAttachment(t, s.Handler(), "", "a.md", []byte("x"))
	if rec.Code != http.StatusUnauthorized {
		t.Fatalf("status = %d", rec.Code)
	}
}
