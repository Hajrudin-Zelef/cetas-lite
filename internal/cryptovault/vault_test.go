package cryptovault

import (
	"bytes"
	"errors"
	"testing"
)

func TestRoundTrip(t *testing.T) {
	salt, err := NewSalt()
	if err != nil {
		t.Fatal(err)
	}
	v, err := New("mot-de-passe", salt)
	if err != nil {
		t.Fatal(err)
	}
	plain := []byte("sk-secret-key")
	aad := []byte("deepseek")
	ct, err := v.Encrypt(plain, aad)
	if err != nil {
		t.Fatal(err)
	}
	if bytes.Contains(ct, plain) {
		t.Fatal("le texte en clair apparait dans le chiffre")
	}
	got, err := v.Decrypt(ct, aad)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(got, plain) {
		t.Fatalf("dechiffre = %q, want %q", got, plain)
	}
}

func TestDecryptWrongAAD(t *testing.T) {
	salt, _ := NewSalt()
	v, _ := New("pw", salt)
	ct, _ := v.Encrypt([]byte("secret"), []byte("deepseek"))
	if _, err := v.Decrypt(ct, []byte("openrouter")); !errors.Is(err, ErrDecrypt) {
		t.Fatalf("AAD different doit echouer, got %v", err)
	}
}

func TestDecryptWrongPassword(t *testing.T) {
	salt, _ := NewSalt()
	v1, _ := New("bon", salt)
	v2, _ := New("mauvais", salt)
	ct, _ := v1.Encrypt([]byte("secret"), nil)
	if _, err := v2.Decrypt(ct, nil); !errors.Is(err, ErrDecrypt) {
		t.Fatalf("mauvais mot de passe doit echouer, got %v", err)
	}
}

func TestDecryptTruncated(t *testing.T) {
	salt, _ := NewSalt()
	v, _ := New("pw", salt)
	if _, err := v.Decrypt([]byte("court"), nil); !errors.Is(err, ErrDecrypt) {
		t.Fatalf("donnees tronquees doivent echouer, got %v", err)
	}
}
