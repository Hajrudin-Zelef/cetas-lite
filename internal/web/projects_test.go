package web

import (
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"cetas-lite/internal/store"
	"cetas-lite/internal/vault"
	"cetas-lite/internal/workspace"
)

// testServerWithProjects crée un serveur de test avec un gestionnaire de
// projets branché.
func testServerWithProjects(t *testing.T) (*Server, string) {
	t.Helper()
	t.Setenv("CETAS_LITE_VAULT_PASSWORD", "mot-de-passe-test")
	s := newTestServer(t, true)
	token := registerAndLogin(t, s.Handler(), "sam")
	s.SetWorkspaceManager(workspace.New(t.TempDir(), s.st, vaultLazyTest{s.st}))
	return s, token
}

type vaultLazyTest struct{ st *store.Store }

func (v vaultLazyTest) Encrypt(plain, aad []byte) ([]byte, error) {
	vv, err := vault.Open(v.st)
	if err != nil {
		return nil, err
	}
	return vv.Encrypt(plain, aad)
}

func (v vaultLazyTest) Decrypt(data, aad []byte) ([]byte, error) {
	vv, err := vault.Open(v.st)
	if err != nil {
		return nil, err
	}
	return vv.Decrypt(data, aad)
}

func TestProjectsCRUD(t *testing.T) {
	s, token := testServerWithProjects(t)
	h := s.Handler()

	rec := doJSON(t, h, http.MethodPost, "/api/projects", token, map[string]any{"name": "Demo", "mode": "local"})
	if rec.Code != http.StatusCreated {
		t.Fatalf("create = %d: %s", rec.Code, rec.Body.String())
	}
	created := decode(t, rec)
	id, _ := created["id"].(string)
	if id == "" {
		t.Fatal("id manquant")
	}

	rec = doJSON(t, h, http.MethodGet, "/api/projects", token, nil)
	if rec.Code != http.StatusOK {
		t.Fatalf("list = %d", rec.Code)
	}
	list := decode(t, rec)["projects"].([]any)
	if len(list) != 1 {
		t.Fatalf("projets: %d", len(list))
	}

	// Projet actif.
	rec = doJSON(t, h, http.MethodPut, "/api/projects/active", token, map[string]any{"id": id})
	if rec.Code != http.StatusOK {
		t.Fatalf("active put = %d", rec.Code)
	}
	rec = doJSON(t, h, http.MethodGet, "/api/projects/active", token, nil)
	if active := decode(t, rec)["active"].(string); active != id {
		t.Fatalf("active = %q", active)
	}

	// Arborescence vide.
	rec = doJSON(t, h, http.MethodGet, "/api/projects/"+id+"/tree", token, nil)
	if rec.Code != http.StatusOK {
		t.Fatalf("tree = %d", rec.Code)
	}

	// Suppression.
	rec = doJSON(t, h, http.MethodDelete, "/api/projects/"+id, token, nil)
	if rec.Code != http.StatusOK {
		t.Fatalf("delete = %d", rec.Code)
	}
	rec = doJSON(t, h, http.MethodGet, "/api/projects", token, nil)
	if n := len(decode(t, rec)["projects"].([]any)); n != 0 {
		t.Fatalf("projets restants: %d", n)
	}
	// Le projet actif est réinitialisé.
	rec = doJSON(t, h, http.MethodGet, "/api/projects/active", token, nil)
	if active := decode(t, rec)["active"].(string); active != "" {
		t.Fatalf("active devrait être vide: %q", active)
	}
}

func TestProjectUploadAndFile(t *testing.T) {
	s, token := testServerWithProjects(t)
	h := s.Handler()

	rec := doJSON(t, h, http.MethodPost, "/api/projects", token, map[string]any{"name": "Up", "mode": "local"})
	id := decode(t, rec)["id"].(string)

	// Upload multipart avec chemins relatifs.
	var b strings.Builder
	w := newMultipartWriter(&b)
	addMultipartFile(t, w, "files", "a.txt", "hello")
	addMultipartFile(t, w, "files", "b.txt", "x")
	addMultipartFile(t, w, "files", "evil.txt", "nope")
	// La stdlib Go ne conserve que le nom de base des fichiers multipart :
	// la structure des dossiers passe par le champ "paths" (JSON).
	pf, err := w.w.CreateFormField("paths")
	if err != nil {
		t.Fatal(err)
	}
	_, _ = pf.Write([]byte(`["src/a.txt","b.txt","../../evil.txt"]`))
	w.close()
	req := multipartRequest(t, "/api/projects/"+id+"/upload", &b, w.boundaryStr())
	req.Header.Set("Authorization", "Bearer "+token)
	rec2 := serveRequest(t, h, req)
	if rec2.Code != http.StatusOK {
		t.Fatalf("upload = %d: %s", rec2.Code, rec2.Body.String())
	}
	if n := int(decode(t, rec2)["count"].(float64)); n != 2 {
		t.Fatalf("count = %d (le chemin ../../evil.txt doit être rejeté)", n)
	}
	// Aucune évasion : le fichier reste dans le workspace.
	rec = doJSON(t, h, http.MethodGet, "/api/projects/"+id+"/file?path=..%2Fevil.txt", token, nil)
	if rec.Code == http.StatusOK {
		t.Fatal("lecture hors workspace acceptée")
	}

	// Lecture du fichier.
	rec = doJSON(t, h, http.MethodGet, "/api/projects/"+id+"/file?path=src/a.txt", token, nil)
	if rec.Code != http.StatusOK {
		t.Fatalf("file = %d", rec.Code)
	}
	if content := decode(t, rec)["content"].(string); content != "hello" {
		t.Fatalf("content = %q", content)
	}

	// Arborescence : le fichier apparaît.
	rec = doJSON(t, h, http.MethodGet, "/api/projects/"+id+"/tree", token, nil)
	body := rec.Body.String()
	if !strings.Contains(body, "a.txt") {
		t.Fatalf("tree sans a.txt: %s", body)
	}

	// Suppression du fichier.
	rec = doJSON(t, h, http.MethodDelete, "/api/projects/"+id+"/file?path=src/a.txt", token, nil)
	if rec.Code != http.StatusOK {
		t.Fatalf("file delete = %d", rec.Code)
	}
}

func TestConnectorsGitHub(t *testing.T) {
	s, token := testServerWithProjects(t)
	h := s.Handler()

	// Non connecté au départ.
	rec := doJSON(t, h, http.MethodGet, "/api/connectors", token, nil)
	if rec.Code != http.StatusOK {
		t.Fatalf("connectors = %d", rec.Code)
	}
	gh := decode(t, rec)["github"].(map[string]any)
	if gh["connected"] != false {
		t.Fatal("github devrait être déconnecté")
	}

	// Token vide refusé.
	rec = doJSON(t, h, http.MethodPut, "/api/connectors/github", token, map[string]any{"token": ""})
	if rec.Code != http.StatusBadRequest {
		t.Fatalf("token vide = %d", rec.Code)
	}

	// Déconnexion idempotente.
	rec = doJSON(t, h, http.MethodDelete, "/api/connectors/github", token, nil)
	if rec.Code != http.StatusOK {
		t.Fatalf("disconnect = %d", rec.Code)
	}
}

// --- helpers multipart ---

type mpWriter struct {
	w        *multipart.Writer
	buf      *strings.Builder
	boundary string
}

func newMultipartWriter(buf *strings.Builder) *mpWriter {
	w := multipart.NewWriter(buf)
	return &mpWriter{w: w, buf: buf, boundary: w.Boundary()}
}

func (m *mpWriter) boundaryStr() string { return m.boundary }

func addMultipartFile(t *testing.T, m *mpWriter, field, filename, content string) {
	t.Helper()
	fw, err := m.w.CreateFormFile(field, filename)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := io.WriteString(fw, content); err != nil {
		t.Fatal(err)
	}
}

func (m *mpWriter) close() { _ = m.w.Close() }

func multipartRequest(t *testing.T, path string, body *strings.Builder, boundary string) *http.Request {
	t.Helper()
	req := httptest.NewRequest(http.MethodPost, path, strings.NewReader(body.String()))
	req.Header.Set("Content-Type", "multipart/form-data; boundary="+boundary)
	return req
}

func serveRequest(t *testing.T, h http.Handler, req *http.Request) *httptest.ResponseRecorder {
	t.Helper()
	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, req)
	return rec
}
