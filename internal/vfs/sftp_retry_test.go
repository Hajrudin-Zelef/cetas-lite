package vfs

// Tests de la fiabilité SFTP (phase 4) : détection d'erreur réseau
// transitoire (F6.3), boucle de réessai avec backoff, sonde temporisée (F6.2).

import (
	"context"
	"crypto/ed25519"
	"crypto/rand"
	"errors"
	"fmt"
	"io"
	"net"
	"testing"
	"time"

	"golang.org/x/crypto/ssh"
)

func TestIsTransientNetErr(t *testing.T) {
	transient := []error{
		&net.OpError{Op: "read", Err: errors.New("connection reset by peer")},
		&net.DNSError{Err: "timeout", IsTimeout: true},
		io.ErrUnexpectedEOF,
		errors.New("connection reset by peer"),
		errors.New("broken pipe"),
		errors.New("ssh: connection closed"),
		errors.New("use of closed network connection"),
		errors.New("read: connection timed out"),
		fmt.Errorf("stat: %w", io.ErrUnexpectedEOF),
		fmt.Errorf("dial: %w", &net.OpError{Op: "dial", Err: errors.New("no route to host")}),
	}
	for _, err := range transient {
		if !isTransientNetErr(err) {
			t.Errorf("isTransientNetErr(%v) = false, attendu true", err)
		}
	}

	notTransient := []error{
		nil,
		ErrNotFound,
		ErrOutsideRoot,
		fmt.Errorf("lecture: %w", ErrNotFound),
		context.Canceled,
		context.DeadlineExceeded,
		errors.New("permission denied"),
		errors.New("delai depasse"),
		errors.New("clef hote inconnue: SHA256:abc"),
		errors.New("open notes_eof.txt: permission denied"),
	}
	for _, err := range notTransient {
		if isTransientNetErr(err) {
			t.Errorf("isTransientNetErr(%v) = true, attendu false", err)
		}
	}
}

func TestRetryTransient(t *testing.T) {
	boom := errors.New("connection reset by peer")

	t.Run("succes immediat", func(t *testing.T) {
		attempts, resets := 0, 0
		err := retryTransient(context.Background(), func() { resets++ }, func() error {
			attempts++
			return nil
		})
		if err != nil || attempts != 1 || resets != 0 {
			t.Errorf("attendus (nil, 1, 0), obtenus (%v, %d, %d)", err, attempts, resets)
		}
	})

	t.Run("transitoire puis succes", func(t *testing.T) {
		attempts, resets := 0, 0
		err := retryTransient(context.Background(), func() { resets++ }, func() error {
			attempts++
			if attempts < 3 {
				return boom
			}
			return nil
		})
		if err != nil || attempts != 3 || resets != 2 {
			t.Errorf("attendus (nil, 3, 2), obtenus (%v, %d, %d)", err, attempts, resets)
		}
	})

	t.Run("toujours transitoire", func(t *testing.T) {
		attempts, resets := 0, 0
		err := retryTransient(context.Background(), func() { resets++ }, func() error {
			attempts++
			return boom
		})
		if !errors.Is(err, boom) || attempts != 3 || resets != 2 {
			t.Errorf("attendus (boom, 3, 2), obtenus (%v, %d, %d)", err, attempts, resets)
		}
	})

	t.Run("erreur metier sans reessai", func(t *testing.T) {
		attempts, resets := 0, 0
		err := retryTransient(context.Background(), func() { resets++ }, func() error {
			attempts++
			return ErrNotFound
		})
		if !errors.Is(err, ErrNotFound) || attempts != 1 || resets != 0 {
			t.Errorf("attendus (ErrNotFound, 1, 0), obtenus (%v, %d, %d)", err, attempts, resets)
		}
	})

	t.Run("contexte annule", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		cancel()
		attempts := 0
		err := retryTransient(ctx, func() {}, func() error {
			attempts++
			return boom
		})
		if !errors.Is(err, context.Canceled) || attempts != 1 {
			t.Errorf("attendus (canceled, 1), obtenus (%v, %d)", err, attempts)
		}
	})

	t.Run("backoff effectif", func(t *testing.T) {
		start := time.Now()
		_ = retryTransient(context.Background(), func() {}, func() error { return boom })
		// 2 réessais : 200ms + 400ms = 600ms minimum (marge large).
		if elapsed := time.Since(start); elapsed < 500*time.Millisecond {
			t.Errorf("backoff trop court : %v", elapsed)
		}
	})
}

func TestNeedProbe(t *testing.T) {
	now := time.Now()
	if needProbe(now, now) {
		t.Error("sonde immédiate après contrôle : attendu false")
	}
	if needProbe(now.Add(-29*time.Second), now) {
		t.Error("sonde à 29s d'inactivité : attendu false")
	}
	if !needProbe(now.Add(-31*time.Second), now) {
		t.Error("sonde à 31s d'inactivité : attendu true")
	}
	if !needProbe(time.Time{}, now) {
		t.Error("sonde sans historique : attendu true")
	}
}

// Tests phase 5 : TOFU avant authentification (F3) et watchdog (F4).

func testHostKey(t *testing.T) ssh.PublicKey {
	t.Helper()
	pub, _, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		t.Fatal(err)
	}
	k, err := ssh.NewPublicKey(pub)
	if err != nil {
		t.Fatal(err)
	}
	return k
}

func TestCheckHostKey(t *testing.T) {
	key := testHostKey(t)
	other := testHostKey(t)

	// Clé inconnue : ErrUnknownHostKey avec clé et empreinte.
	err := checkHostKey(nil, key)
	var unknown *ErrUnknownHostKey
	if !errors.As(err, &unknown) {
		t.Fatalf("attendu ErrUnknownHostKey, obtenu %v", err)
	}
	if len(unknown.Key) == 0 || unknown.Fingerprint == "" {
		t.Fatal("ErrUnknownHostKey doit transporter la cle et son empreinte")
	}
	if !keysEqual(unknown.Key, key.Marshal()) {
		t.Fatal("la cle transportee doit etre la cle presentee")
	}

	// Clé connue et identique : OK.
	if err := checkHostKey(key.Marshal(), key); err != nil {
		t.Fatalf("cle connue : attendu nil, obtenu %v", err)
	}

	// Clé changée : erreur dure, pas ErrUnknownHostKey.
	err = checkHostKey(other.Marshal(), key)
	if err == nil {
		t.Fatal("cle changee : erreur attendue")
	}
	if errors.As(err, &unknown) {
		t.Fatalf("cle changee : ne doit pas etre ErrUnknownHostKey (%v)", err)
	}
}

func TestRunWatchedTimeout(t *testing.T) {
	old := sftpOpTimeout
	sftpOpTimeout = 30 * time.Millisecond
	defer func() { sftpOpTimeout = old }()

	s := &SFTPFS{}
	release := make(chan struct{})
	done := make(chan error, 1)
	go func() {
		done <- s.runWatched(context.Background(), func() error {
			<-release // simule un appel SFTP figé, débloqué par resetConn
			return errors.New("debloque")
		})
	}()
	// Le watchdog (30 ms) doit frapper bien avant qu'on libère l'op (200 ms).
	time.Sleep(200 * time.Millisecond)
	close(release)
	select {
	case err := <-done:
		if !errors.Is(err, errSFTPOpTimeout) {
			t.Fatalf("attendu errSFTPOpTimeout, obtenu %v", err)
		}
	case <-time.After(5 * time.Second):
		t.Fatal("runWatched n'est jamais revenu")
	}
}

func TestRunWatchedCtxCancel(t *testing.T) {
	old := sftpOpTimeout
	sftpOpTimeout = time.Hour // seul le ctx doit frapper
	defer func() { sftpOpTimeout = old }()

	s := &SFTPFS{}
	ctx, cancel := context.WithCancel(context.Background())
	release := make(chan struct{})
	done := make(chan error, 1)
	go func() {
		done <- s.runWatched(ctx, func() error {
			<-release
			return nil
		})
	}()
	time.Sleep(50 * time.Millisecond)
	cancel()
	close(release)
	select {
	case err := <-done:
		if !errors.Is(err, context.Canceled) {
			t.Fatalf("attendu context.Canceled, obtenu %v", err)
		}
	case <-time.After(5 * time.Second):
		t.Fatal("runWatched n'est jamais revenu")
	}
}

func TestOpTimeoutNotTransient(t *testing.T) {
	if isTransientNetErr(errSFTPOpTimeout) {
		t.Fatal("errSFTPOpTimeout ne doit jamais etre reessaye")
	}
	if isTransientNetErr(context.DeadlineExceeded) {
		t.Fatal("context.DeadlineExceeded ne doit jamais etre reessaye")
	}
}
