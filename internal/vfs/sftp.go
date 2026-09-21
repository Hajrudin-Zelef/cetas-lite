package vfs

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net"
	"os"
	"path"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

// SFTPConfig décrit une connexion SFTP vers un dossier distant.
type SFTPConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	// PrivateKey : clé privée OpenSSH (optionnelle, prioritaire sur password).
	PrivateKey    []byte
	KeyPassphrase string
	// RemotePath : dossier distant servant de racine (ex. /home/u/projet).
	RemotePath string
	// HostKey : clé publique hôte connue (TOFU). Vide au premier contact :
	// la clé présentée est alors renvoyée via ErrUnknownHostKey pour
	// validation explicite avant stockage.
	HostKey []byte
}

// ErrUnknownHostKey est renvoyé au premier contact avec un serveur : il
// transporte l'empreinte de la clé présentée pour validation utilisateur.
type ErrUnknownHostKey struct {
	Fingerprint string
	Key         []byte
}

func (e *ErrUnknownHostKey) Error() string {
	return "cle hote inconnue: " + e.Fingerprint
}

// sshDialTimeout : délai max d'établissement de la connexion SSH
// (TCP + handshake). Appliqué même quand le ctx appelant n'a pas de
// deadline propre (F4).
const sshDialTimeout = 10 * time.Second

// checkHostKey vérifie la clé hôte présentée AVANT toute authentification
// (F3) : appelée par HostKeyCallback pendant le handshake SSH, donc avant
// l'envoi de tout secret (mot de passe notamment). Comportement inchangé
// côté appelant : clé inconnue -> ErrUnknownHostKey (clé à faire valider
// puis stocker), clé changée -> erreur dure.
func checkHostKey(known []byte, key ssh.PublicKey) error {
	presented := key.Marshal()
	if len(known) > 0 && !keysEqual(known, presented) {
		return errors.New("la cle hote du serveur a change (attaque possible) — reverifiez la connexion")
	}
	if len(known) == 0 {
		return &ErrUnknownHostKey{Fingerprint: ssh.FingerprintSHA256(key), Key: presented}
	}
	return nil
}

// dialSSH établit la connexion SSH : vérification TOFU de la clé hôte
// avant authentification (F3), établissement annulable et borné dans le
// temps même sans deadline sur ctx (F4).
func dialSSH(ctx context.Context, cfg SFTPConfig) (*ssh.Client, error) {
	if cfg.Host == "" || cfg.User == "" {
		return nil, errors.New("hote ou utilisateur manquant")
	}
	port := cfg.Port
	if port <= 0 {
		port = 22
	}
	var auths []ssh.AuthMethod
	if len(cfg.PrivateKey) > 0 {
		var signer ssh.Signer
		var err error
		if cfg.KeyPassphrase != "" {
			signer, err = ssh.ParsePrivateKeyWithPassphrase(cfg.PrivateKey, []byte(cfg.KeyPassphrase))
		} else {
			signer, err = ssh.ParsePrivateKey(cfg.PrivateKey)
		}
		if err != nil {
			return nil, fmt.Errorf("cle privee invalide: %w", err)
		}
		auths = append(auths, ssh.PublicKeys(signer))
	}
	if cfg.Password != "" {
		auths = append(auths, ssh.Password(cfg.Password))
	}
	if len(auths) == 0 {
		return nil, errors.New("aucune methode d'authentification")
	}

	sshCfg := &ssh.ClientConfig{
		User: cfg.User,
		Auth: auths,
		// F3 : la vérification TOFU a lieu ici, pendant le handshake,
		// avant que l'authentification n'envoie le moindre secret.
		HostKeyCallback: func(_ string, _ net.Addr, key ssh.PublicKey) error {
			return checkHostKey(cfg.HostKey, key)
		},
	}
	addr := fmt.Sprintf("%s:%d", cfg.Host, port)
	// F4 : l'établissement est borné (sshDialTimeout) même si le ctx
	// appelant n'a pas de deadline ; une annulation reste prioritaire.
	dialCtx := ctx
	cancel := func() {}
	if _, ok := ctx.Deadline(); !ok {
		dialCtx, cancel = context.WithTimeout(ctx, sshDialTimeout)
	}
	defer cancel()
	dialer := &net.Dialer{Timeout: sshDialTimeout}
	conn, err := dialer.DialContext(dialCtx, "tcp", addr)
	if err != nil {
		return nil, fmt.Errorf("connexion SSH: %w", err)
	}
	// NewClientConn ne prend pas de ctx : le handshake tourne en goroutine
	// et la fermeture de conn débloque un handshake figé (F4).
	type dialRes struct {
		sc  ssh.Conn
		ch  <-chan ssh.NewChannel
		req <-chan *ssh.Request
		err error
	}
	done := make(chan dialRes, 1)
	go func() {
		sc, ch, req, err := ssh.NewClientConn(conn, addr, sshCfg)
		done <- dialRes{sc, ch, req, err}
	}()
	select {
	case <-dialCtx.Done():
		conn.Close()
		<-done
		return nil, dialCtx.Err()
	case r := <-done:
		if r.err != nil {
			conn.Close()
			var unknown *ErrUnknownHostKey
			if errors.As(r.err, &unknown) {
				return nil, unknown
			}
			return nil, fmt.Errorf("connexion SSH: %w", r.err)
		}
		return ssh.NewClient(r.sc, r.ch, r.req), nil
	}
}

func keysEqual(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// SFTPFS est l'implémentation distante de FS (dossier via SFTP).
type SFTPFS struct {
	cfg  SFTPConfig
	name string

	mu     sync.Mutex
	sshCl  *ssh.Client
	sftpCl *sftp.Client
	// lastProbe : dernier contrôle de santé réussi (F6.2).
	lastProbe time.Time
}

func NewSFTP(cfg SFTPConfig, name string) *SFTPFS {
	if name == "" {
		name = cfg.User + "@" + cfg.Host + ":" + cfg.RemotePath
	}
	return &SFTPFS{cfg: cfg, name: name}
}

func (s *SFTPFS) Name() string { return s.name }
func (s *SFTPFS) Remote() bool { return true }

// client retourne un client SFTP connecté (reconnexion paresseuse).
func (s *SFTPFS) client(ctx context.Context) (*sftp.Client, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.sftpCl != nil {
		// F6.2 : sonde temporisée — on ne sonde la connexion que si elle
		// est inactive depuis sftpProbeInterval, pas à chaque opération.
		// Sous le seuil, une connexion morte sera récupérée par le
		// réessai F6.3 de l'opération.
		if !needProbe(s.lastProbe, time.Now()) {
			return s.sftpCl, nil
		}
		if _, err := s.sftpCl.Stat(s.cfg.RemotePath); err == nil {
			s.lastProbe = time.Now()
			return s.sftpCl, nil
		}
		s.closeLocked()
	}
	sshCl, err := dialSSH(ctx, s.cfg)
	if err != nil {
		return nil, err
	}
	cl, err := sftp.NewClient(sshCl)
	if err != nil {
		sshCl.Close()
		return nil, fmt.Errorf("session SFTP: %w", err)
	}
	// La racine distante doit exister.
	if _, err := cl.Stat(s.cfg.RemotePath); err != nil {
		cl.Close()
		sshCl.Close()
		return nil, fmt.Errorf("dossier distant introuvable: %s", s.cfg.RemotePath)
	}
	s.sshCl = sshCl
	s.sftpCl = cl
	s.lastProbe = time.Now()
	return cl, nil
}

func (s *SFTPFS) closeLocked() {
	if s.sftpCl != nil {
		s.sftpCl.Close()
		s.sftpCl = nil
	}
	if s.sshCl != nil {
		s.sshCl.Close()
		s.sshCl = nil
	}
}

func (s *SFTPFS) Close() error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.closeLocked()
	return nil
}

// rpath construit le chemin distant absolu depuis un relatif.
func (s *SFTPFS) rpath(rel string) (string, error) {
	clean := ""
	if strings.TrimSpace(rel) != "" {
		var err error
		clean, err = cleanRel(rel)
		if err != nil {
			return "", err
		}
	}
	base := path.Clean("/" + strings.TrimPrefix(s.cfg.RemotePath, "/"))
	if clean == "" {
		return base, nil
	}
	joined := path.Join(base, path.Clean("/"+clean))
	if joined != base && !strings.HasPrefix(joined, base+"/") {
		return "", ErrOutsideRoot
	}
	return joined, nil
}

func (s *SFTPFS) Resolve(rel string) (string, error) {
	if strings.TrimSpace(rel) == "" {
		return "", nil
	}
	return cleanRel(rel)
}

func (s *SFTPFS) ReadFile(ctx context.Context, rel string) ([]byte, error) {
	var out []byte
	err := s.doRetry(ctx, func(cl *sftp.Client) error {
		rp, err := s.rpath(rel)
		if err != nil {
			return err
		}
		f, err := cl.Open(rp)
		if err != nil {
			return toFSerr(err)
		}
		defer f.Close()
		out, err = io.ReadAll(io.LimitReader(f, 64<<20))
		return err
	})
	return out, err
}

func (s *SFTPFS) WriteFile(ctx context.Context, rel string, data []byte, perm os.FileMode) error {
	return s.doRetry(ctx, func(cl *sftp.Client) error {
		rp, err := s.rpath(rel)
		if err != nil {
			return err
		}
		// Crée les dossiers parents.
		if dir := path.Dir(rp); dir != s.cfg.RemotePath {
			if err := cl.MkdirAll(dir); err != nil {
				return err
			}
		}
		f, err := cl.OpenFile(rp, os.O_WRONLY|os.O_CREATE|os.O_TRUNC)
		if err != nil {
			return toFSerr(err)
		}
		defer f.Close()
		if _, err := f.Write(data); err != nil {
			return err
		}
		return nil
	})
}

func (s *SFTPFS) MkdirAll(ctx context.Context, rel string) error {
	return s.doRetry(ctx, func(cl *sftp.Client) error {
		rp, err := s.rpath(rel)
		if err != nil {
			return err
		}
		return cl.MkdirAll(rp)
	})
}

func (s *SFTPFS) Remove(ctx context.Context, rel string) error {
	return s.doRetry(ctx, func(cl *sftp.Client) error {
		rp, err := s.rpath(rel)
		if err != nil {
			return err
		}
		return cl.Remove(rp)
	})
}

func (s *SFTPFS) Rename(ctx context.Context, oldrel, newrel string) error {
	return s.doRetry(ctx, func(cl *sftp.Client) error {
		oldp, err := s.rpath(oldrel)
		if err != nil {
			return err
		}
		newp, err := s.rpath(newrel)
		if err != nil {
			return err
		}
		return cl.Rename(oldp, newp)
	})
}

func toEntrySFTP(rel string, fi os.FileInfo) Entry {
	e := Entry{Path: rel, IsDir: fi.IsDir(), Size: fi.Size(), ModTime: fi.ModTime()}
	if e.IsDir && !strings.HasSuffix(e.Path, "/") {
		e.Path += "/"
	}
	return e
}

func (s *SFTPFS) Stat(ctx context.Context, rel string) (Entry, error) {
	var out Entry
	err := s.doRetry(ctx, func(cl *sftp.Client) error {
		rp, err := s.rpath(rel)
		if err != nil {
			return err
		}
		fi, err := cl.Stat(rp)
		if err != nil {
			return toFSerr(err)
		}
		clean, _ := cleanRel(rel)
		if clean == "" {
			out = Entry{Path: "", IsDir: true}
			return nil
		}
		out = toEntrySFTP(clean, fi)
		return nil
	})
	return out, err
}

func (s *SFTPFS) ReadDir(ctx context.Context, rel string) ([]Entry, error) {
	var out []Entry
	err := s.doRetry(ctx, func(cl *sftp.Client) error {
		rp, err := s.rpath(rel)
		if err != nil {
			return err
		}
		fis, err := cl.ReadDir(rp)
		if err != nil {
			return toFSerr(err)
		}
		clean, _ := cleanRel(rel)
		out = make([]Entry, 0, len(fis))
		for _, fi := range fis {
			if SkipEntry(fi.Name()) {
				continue
			}
			rp2 := fi.Name()
			if clean != "" {
				rp2 = clean + "/" + fi.Name()
			}
			out = append(out, toEntrySFTP(rp2, fi))
		}
		sort.Slice(out, func(i, j int) bool {
			if out[i].IsDir != out[j].IsDir {
				return out[i].IsDir
			}
			return out[i].Path < out[j].Path
		})
		return nil
	})
	return out, err
}

func (s *SFTPFS) Walk(ctx context.Context, fn func(Entry) error) error {
	var rec func(rel string) error
	rec = func(rel string) error {
		if ctx.Err() != nil {
			return ctx.Err()
		}
		ents, err := s.ReadDir(ctx, rel)
		if err != nil {
			return nil // dossier illisible : on continue
		}
		for _, e := range ents {
			if err := fn(e); err != nil {
				return err
			}
			if e.IsDir {
				sub := strings.TrimSuffix(e.Path, "/")
				if err := rec(sub); err != nil {
					return err
				}
			}
		}
		return nil
	}
	return rec("")
}

// Exec exécute une commande sur le serveur distant via SSH.
func (s *SFTPFS) Exec(ctx context.Context, name string, args []string, env map[string]string, timeout time.Duration) (string, error) {
	if timeout <= 0 {
		timeout = 60 * time.Second
	}
	s.mu.Lock()
	sshCl := s.sshCl
	s.mu.Unlock()
	if sshCl == nil {
		if _, err := s.client(ctx); err != nil {
			return "", err
		}
		s.mu.Lock()
		sshCl = s.sshCl
		s.mu.Unlock()
	}
	cctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()
	sess, err := sshCl.NewSession()
	if err != nil && isTransientNetErr(err) {
		// La connexion SSH semblait établie mais est morte : on la
		// rétablit une fois puis on recrée la session. La commande n'a
		// pas démarré : aucun risque de double exécution (F6.3).
		s.resetConn()
		if _, cerr := s.client(ctx); cerr != nil {
			return "", err
		}
		s.mu.Lock()
		sshCl = s.sshCl
		s.mu.Unlock()
		sess, err = sshCl.NewSession()
	}
	if err != nil {
		return "", err
	}
	defer sess.Close()
	// Commande confinée au dossier du projet distant.
	remote := s.cfg.RemotePath
	quoted := make([]string, 0, len(args)+1)
	quoted = append(quoted, name)
	quoted = append(quoted, args...)
	inner := shJoin(quoted)
	if len(env) > 0 {
		assigns := make([]string, 0, len(env))
		for k, v := range env {
			assigns = append(assigns, k+"="+shQuote(v))
		}
		sort.Strings(assigns)
		inner = "env " + strings.Join(assigns, " ") + " " + inner
	}
	cmd := "cd " + shQuote(remote) + " && " + inner
	done := make(chan struct{})
	var out []byte
	var runErr error
	go func() {
		defer close(done)
		out, runErr = sess.CombinedOutput(cmd)
	}()
	select {
	case <-cctx.Done():
		sess.Signal(ssh.SIGKILL)
		<-done
		return string(out), errors.New("delai depasse")
	case <-done:
		return string(out), runErr
	}
}

func shQuote(s string) string {
	return "'" + strings.ReplaceAll(s, "'", "'\\''") + "'"
}

func shJoin(args []string) string {
	qs := make([]string, len(args))
	for i, a := range args {
		qs[i] = shQuote(a)
	}
	return strings.Join(qs, " ")
}

func toFSerr(err error) error {
	if err == nil {
		return nil
	}
	if os.IsNotExist(err) {
		return ErrNotFound
	}
	return err
}
