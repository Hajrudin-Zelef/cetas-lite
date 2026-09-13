package auth

import (
	"errors"
	"path/filepath"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"cetas-lite/internal/store"
)

func newManager(t *testing.T) *Manager {
	t.Helper()
	st, err := store.Open(filepath.Join(t.TempDir(), "auth.db"))
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = st.Close() })
	m, err := New(st)
	if err != nil {
		t.Fatal(err)
	}
	return m
}

func TestHashVerify(t *testing.T) {
	h, err := HashPassword("motdepasse")
	if err != nil {
		t.Fatal(err)
	}
	if !VerifyPassword("motdepasse", h) {
		t.Fatal("le bon mot de passe doit etre valide")
	}
	if VerifyPassword("mauvais", h) {
		t.Fatal("un mauvais mot de passe doit etre refuse")
	}
	if VerifyPassword("motdepasse", "sha256:abc") {
		t.Fatal("un hash inconnu doit etre refuse")
	}
}

func TestRegisterAndAuthenticate(t *testing.T) {
	m := newManager(t)
	if err := m.Register("sam", "motdepasse"); err != nil {
		t.Fatal(err)
	}
	if err := m.Register("sam", "autre-chose"); !errors.Is(err, ErrUserExists) {
		t.Fatalf("doublon = %v", err)
	}
	if err := m.Register("ab", "motdepasse"); !errors.Is(err, ErrInvalidUsername) {
		t.Fatalf("username court = %v", err)
	}
	if err := m.Register("bobby", "court"); !errors.Is(err, ErrInvalidPassword) {
		t.Fatalf("mot de passe court = %v", err)
	}
	u, err := m.Authenticate("sam", "motdepasse")
	if err != nil {
		t.Fatal(err)
	}
	if u.Role != "user" {
		t.Fatalf("role = %q", u.Role)
	}
	if _, err := m.Authenticate("sam", "mauvais"); !errors.Is(err, ErrInvalidCredentials) {
		t.Fatalf("mauvais mdp = %v", err)
	}
	if _, err := m.Authenticate("inconnu", "motdepasse"); !errors.Is(err, ErrInvalidCredentials) {
		t.Fatalf("inconnu = %v", err)
	}
}

func TestTokenRoundTrip(t *testing.T) {
	m := newManager(t)
	if err := m.Register("sam", "motdepasse"); err != nil {
		t.Fatal(err)
	}
	u, _ := m.Authenticate("sam", "motdepasse")
	tok, err := m.Token(u)
	if err != nil {
		t.Fatal(err)
	}
	claims, err := m.Parse(tok)
	if err != nil {
		t.Fatal(err)
	}
	if claims.Username != "sam" || claims.Role != "user" {
		t.Fatalf("claims = %+v", claims)
	}
}

func TestTokenTampered(t *testing.T) {
	m := newManager(t)
	if err := m.Register("sam", "motdepasse"); err != nil {
		t.Fatal(err)
	}
	u, _ := m.Authenticate("sam", "motdepasse")
	tok, _ := m.Token(u)
	if _, err := m.Parse(tok + "x"); !errors.Is(err, ErrInvalidToken) {
		t.Fatalf("token altere = %v", err)
	}
	if _, err := m.Parse("pas-un-token"); !errors.Is(err, ErrInvalidToken) {
		t.Fatalf("token invalide = %v", err)
	}
}

func TestTokenExpired(t *testing.T) {
	m := newManager(t)
	claims := jwt.MapClaims{
		"sub": "sam",
		"exp": time.Now().Add(-time.Hour).Unix(),
	}
	tok, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(m.secret)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := m.Parse(tok); !errors.Is(err, ErrInvalidToken) {
		t.Fatalf("token expire = %v", err)
	}
}
