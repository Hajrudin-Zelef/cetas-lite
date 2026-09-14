// Package terminal : terminal integre (PTY) multi-session expose en SSE.
//
// Chaque session lance un shell dans un repertoire de travail autorise
// (workspace utilisateur ou racine des worktrees). La sortie est diffusee
// aux abonnes via des canaux ; le transport reseau (SSE) est gere par la
// couche web. Sur Unix un vrai PTY est utilise (creack/pty), sous Windows
// le shell tourne avec des pipes rediriges.
package terminal

import (
	"errors"
	"io"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

// MaxSessionsPerUser borne le nombre de terminaux simultanes par utilisateur.
const MaxSessionsPerUser = 6

// backlogMax conserve les derniers octets emis pour les (re)connexions.
const backlogMax = 32 << 10

var (
	ErrNotFound   = errors.New("session introuvable")
	ErrTooMany    = errors.New("trop de sessions terminal ouvertes")
	ErrBadCwd     = errors.New("repertoire de travail non autorise")
	ErrBadPayload = errors.New("donnees invalides")
)

var idCounter int64

// Session represente un shell interactif en cours d'execution.
type Session struct {
	ID      string
	User    string
	Cwd     string
	Created time.Time

	stdin  io.WriteCloser
	resize func(cols, rows int) error
	kill   func() error

	mu       sync.Mutex
	subs     map[chan []byte]struct{}
	backlog  []byte
	exited   chan struct{}
	exitErr  error
	exitOnce sync.Once
}

// Subscribe retourne un canal recevant les octets de sortie (copies
// immuables), le backlog recent, et un canal ferme a la fin du shell.
func (s *Session) Subscribe() (ch chan []byte, backlog []byte, exited <-chan struct{}) {
	ch = make(chan []byte, 64)
	s.mu.Lock()
	s.subs[ch] = struct{}{}
	backlog = append([]byte(nil), s.backlog...)
	exited = s.exited
	s.mu.Unlock()
	return ch, backlog, exited
}

// Unsubscribe detache un abonne.
func (s *Session) Unsubscribe(ch chan []byte) {
	s.mu.Lock()
	delete(s.subs, ch)
	s.mu.Unlock()
}

func (s *Session) broadcast(p []byte) {
	s.mu.Lock()
	if len(p) > 0 {
		s.backlog = append(s.backlog, p...)
		if len(s.backlog) > backlogMax {
			s.backlog = append([]byte(nil), s.backlog[len(s.backlog)-backlogMax:]...)
		}
	}
	for ch := range s.subs {
		select {
		case ch <- append([]byte(nil), p...):
		default:
			// Abonne trop lent : on jette le chunk plutot que de bloquer
			// la lecture du PTY.
		}
	}
	s.mu.Unlock()
}

func (s *Session) finish(err error) {
	s.exitOnce.Do(func() {
		s.mu.Lock()
		s.exitErr = err
		for ch := range s.subs {
			close(ch)
		}
		s.subs = map[chan []byte]struct{}{}
		close(s.exited)
		s.mu.Unlock()
	})
}

// ExitErr retourne l'erreur de fin du shell (nil si sortie normale).
func (s *Session) ExitErr() error {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.exitErr
}

// WriteInput ecrit des octets sur l'entree standard du shell.
func (s *Session) WriteInput(p []byte) error {
	if len(p) == 0 {
		return nil
	}
	if len(p) > 1<<20 {
		return ErrBadPayload
	}
	_, err := s.stdin.Write(p)
	return err
}

// Resize redimensionne le PTY (vrai redimensionnement sur Unix et Windows/ConPTY).
func (s *Session) Resize(cols, rows int) error {
	if cols < 10 {
		cols = 10
	}
	if rows < 5 {
		rows = 5
	}
	if cols > 500 {
		cols = 500
	}
	if rows > 200 {
		rows = 200
	}
	return s.resize(cols, rows)
}

// Manager gere les sessions terminal de tous les utilisateurs.
type Manager struct {
	mu       sync.Mutex
	sessions map[string]*Session
	roots    []string
}

// NewManager cree un gestionnaire ; roots sont les racines autorisees pour
// le repertoire de travail (resolues en chemins absolus).
func NewManager(roots ...string) *Manager {
	abs := make([]string, 0, len(roots))
	for _, r := range roots {
		if r = strings.TrimSpace(r); r == "" {
			continue
		}
		if a, err := filepath.Abs(r); err == nil {
			abs = append(abs, a)
		}
	}
	return &Manager{sessions: map[string]*Session{}, roots: abs}
}

// AllowedRoots retourne les racines autorisees (copie).
func (m *Manager) AllowedRoots() []string {
	m.mu.Lock()
	defer m.mu.Unlock()
	return append([]string(nil), m.roots...)
}

// validCwd verifie que cwd est sous l'une des racines autorisees.
func (m *Manager) validCwd(cwd string) (string, bool) {
	abs, err := filepath.Abs(cwd)
	if err != nil {
		return "", false
	}
	real, err := filepath.EvalSymlinks(abs)
	if err != nil {
		real = abs
	}
	for _, r := range m.roots {
		if real == r || strings.HasPrefix(real, r+string(filepath.Separator)) {
			return real, true
		}
	}
	return "", false
}

func newID() string {
	n := atomic.AddInt64(&idCounter, 1)
	return time.Now().UTC().Format("20060102150405") + "-" + strconv.FormatInt(n, 10)
}

// Create demarre un shell dans cwd (doit etre sous une racine autorisee).
func (m *Manager) Create(user, cwd string) (*Session, error) {
	user = strings.TrimSpace(user)
	if user == "" {
		return nil, errors.New("utilisateur requis")
	}
	real, ok := m.validCwd(cwd)
	if !ok {
		return nil, ErrBadCwd
	}
	m.mu.Lock()
	count := 0
	for _, s := range m.sessions {
		if s.User == user {
			count++
		}
	}
	if count >= MaxSessionsPerUser {
		m.mu.Unlock()
		return nil, ErrTooMany
	}
	m.mu.Unlock()

	stdin, stdout, resize, kill, wait, err := startShell(real)
	if err != nil {
		return nil, err
	}
	s := &Session{
		ID:      newID(),
		User:    user,
		Cwd:     real,
		Created: time.Now(),
		stdin:   stdin,
		resize:  resize,
		kill:    kill,
		subs:    map[chan []byte]struct{}{},
		exited:  make(chan struct{}),
	}
	_ = s.Resize(120, 30)

	// Double verification sous verrou : deux creations concurrentes ne
	// doivent pas depasser la limite (la premiere verification a eu lieu
	// avant startShell, sans verrou).
	m.mu.Lock()
	cur := 0
	for _, other := range m.sessions {
		if other.User == user {
			cur++
		}
	}
	if cur >= MaxSessionsPerUser {
		m.mu.Unlock()
		_ = kill()
		return nil, ErrTooMany
	}
	m.sessions[s.ID] = s
	m.mu.Unlock()

	go func() {
		buf := make([]byte, 4096)
		for {
			n, rerr := stdout.Read(buf)
			if n > 0 {
				s.broadcast(buf[:n])
			}
			if rerr != nil {
				break
			}
		}
		s.finish(wait())
		m.mu.Lock()
		delete(m.sessions, s.ID)
		m.mu.Unlock()
	}()
	return s, nil
}

// Get retourne la session si elle appartient a l'utilisateur.
func (m *Manager) Get(user, id string) (*Session, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	s, ok := m.sessions[id]
	if !ok || s.User != user {
		return nil, false
	}
	return s, true
}

// List retourne les sessions de l'utilisateur.
func (m *Manager) List(user string) []*Session {
	m.mu.Lock()
	defer m.mu.Unlock()
	var out []*Session
	for _, s := range m.sessions {
		if s.User == user {
			out = append(out, s)
		}
	}
	return out
}

// Delete tue la session et la retire.
func (m *Manager) Delete(user, id string) bool {
	s, ok := m.Get(user, id)
	if !ok {
		return false
	}
	_ = s.kill()
	s.finish(errors.New("session fermee"))
	m.mu.Lock()
	delete(m.sessions, id)
	m.mu.Unlock()
	return true
}

// Close tue toutes les sessions (arret du serveur).
func (m *Manager) Close() {
	m.mu.Lock()
	sess := make([]*Session, 0, len(m.sessions))
	for _, s := range m.sessions {
		sess = append(sess, s)
	}
	m.mu.Unlock()
	for _, s := range sess {
		_ = s.kill()
		s.finish(errors.New("arret du serveur"))
	}
}
