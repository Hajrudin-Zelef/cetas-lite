package auth

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/scrypt"

	"cetas-lite/internal/store"
)

const (
	hashN       = 16384
	hashR       = 8
	hashP       = 1
	hashKeyLen  = 32
	hashSaltLen = 16
	tokenTTL    = 24 * time.Hour
	minUsername = 3
	minPassword = 8
)

var (
	ErrUserExists         = errors.New("cet utilisateur existe deja")
	ErrInvalidCredentials = errors.New("identifiants incorrects")
	ErrInvalidUsername    = errors.New("nom d'utilisateur trop court (3 caracteres minimum)")
	ErrInvalidPassword    = errors.New("mot de passe trop court (8 caracteres minimum)")
	ErrInvalidToken       = errors.New("token invalide ou expire")
)

type Claims struct {
	Username string
	Role     string
}

type Manager struct {
	st        *store.Store
	secret    []byte
	dummyHash string
}

func New(st *store.Store) (*Manager, error) {
	secret, err := st.JWTSecret()
	if err != nil {
		return nil, err
	}
	dummy, err := HashPassword("dummy-password-for-timing")
	if err != nil {
		return nil, err
	}
	return &Manager{st: st, secret: secret, dummyHash: dummy}, nil
}

func HashPassword(password string) (string, error) {
	salt := make([]byte, hashSaltLen)
	if _, err := io.ReadFull(rand.Reader, salt); err != nil {
		return "", fmt.Errorf("generation du sel: %w", err)
	}
	dk, err := scrypt.Key([]byte(password), salt, hashN, hashR, hashP, hashKeyLen)
	if err != nil {
		return "", fmt.Errorf("derivation scrypt: %w", err)
	}
	return fmt.Sprintf("scrypt$%d$%d$%d$%s$%s", hashN, hashR, hashP,
		hex.EncodeToString(salt), hex.EncodeToString(dk)), nil
}

func VerifyPassword(password, encoded string) bool {
	parts := strings.Split(encoded, "$")
	if len(parts) != 6 || parts[0] != "scrypt" {
		return false
	}
	n, errN := strconv.Atoi(parts[1])
	r, errR := strconv.Atoi(parts[2])
	p, errP := strconv.Atoi(parts[3])
	if errN != nil || errR != nil || errP != nil {
		return false
	}
	salt, err := hex.DecodeString(parts[4])
	if err != nil {
		return false
	}
	want, err := hex.DecodeString(parts[5])
	if err != nil {
		return false
	}
	got, err := scrypt.Key([]byte(password), salt, n, r, p, len(want))
	if err != nil {
		return false
	}
	return subtle.ConstantTimeCompare(got, want) == 1
}

func (m *Manager) Register(username, password string) error {
	username = strings.TrimSpace(username)
	if len([]rune(username)) < minUsername {
		return ErrInvalidUsername
	}
	if len([]rune(password)) < minPassword {
		return ErrInvalidPassword
	}
	if _, ok := m.st.GetUser(username); ok {
		return ErrUserExists
	}
	h, err := HashPassword(password)
	if err != nil {
		return err
	}
	return m.st.PutUser(&store.User{
		Username:     username,
		PasswordHash: h,
		Role:         "user",
		Created:      time.Now().UTC().Format(time.RFC3339),
	})
}

func (m *Manager) Authenticate(username, password string) (*store.User, error) {
	u, ok := m.st.GetUser(strings.TrimSpace(username))
	if !ok {
		VerifyPassword(password, m.dummyHash)
		return nil, ErrInvalidCredentials
	}
	if !VerifyPassword(password, u.PasswordHash) {
		return nil, ErrInvalidCredentials
	}
	return u, nil
}

func (m *Manager) Token(u *store.User) (string, error) {
	now := time.Now()
	claims := jwt.MapClaims{
		"sub":  u.Username,
		"role": u.Role,
		"iat":  now.Unix(),
		"exp":  now.Add(tokenTTL).Unix(),
	}
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(m.secret)
}

func (m *Manager) Parse(token string) (*Claims, error) {
	parsed, err := jwt.Parse(token, func(t *jwt.Token) (any, error) {
		if t.Method != jwt.SigningMethodHS256 {
			return nil, ErrInvalidToken
		}
		return m.secret, nil
	})
	if err != nil || !parsed.Valid {
		return nil, ErrInvalidToken
	}
	mc, ok := parsed.Claims.(jwt.MapClaims)
	if !ok {
		return nil, ErrInvalidToken
	}
	sub, _ := mc["sub"].(string)
	role, _ := mc["role"].(string)
	if sub == "" {
		return nil, ErrInvalidToken
	}
	return &Claims{Username: sub, Role: role}, nil
}
