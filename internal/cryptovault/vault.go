package cryptovault

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"fmt"
	"io"

	"golang.org/x/crypto/scrypt"
)

const (
	scryptN = 32768
	scryptR = 8
	scryptP = 1
	keyLen  = 32
	saltLen = 16
)

var ErrDecrypt = errors.New("dechiffrement impossible")

type Vault struct {
	aead cipher.AEAD
}

func DeriveKey(password string, salt []byte) ([]byte, error) {
	return scrypt.Key([]byte(password), salt, scryptN, scryptR, scryptP, keyLen)
}

func NewSalt() ([]byte, error) {
	s := make([]byte, saltLen)
	if _, err := io.ReadFull(rand.Reader, s); err != nil {
		return nil, fmt.Errorf("generation du sel: %w", err)
	}
	return s, nil
}

func New(password string, salt []byte) (*Vault, error) {
	if len(salt) == 0 {
		return nil, errors.New("sel vide")
	}
	key, err := DeriveKey(password, salt)
	if err != nil {
		return nil, fmt.Errorf("derivation de cle: %w", err)
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("chiffrement AES: %w", err)
	}
	aead, err := cipher.NewGCM(block)
	if err != nil {
		return nil, fmt.Errorf("mode GCM: %w", err)
	}
	return &Vault{aead: aead}, nil
}

func (v *Vault) Encrypt(plain, aad []byte) ([]byte, error) {
	iv := make([]byte, v.aead.NonceSize())
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, fmt.Errorf("generation du nonce: %w", err)
	}
	return v.aead.Seal(iv, iv, plain, aad), nil
}

func (v *Vault) Decrypt(data, aad []byte) ([]byte, error) {
	ns := v.aead.NonceSize()
	if len(data) < ns {
		return nil, ErrDecrypt
	}
	iv, ct := data[:ns], data[ns:]
	out, err := v.aead.Open(nil, iv, ct, aad)
	if err != nil {
		return nil, ErrDecrypt
	}
	return out, nil
}
