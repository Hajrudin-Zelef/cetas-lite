package vfs

// Tests de la fiabilité SFTP (phase 4) : détection d'erreur réseau
// transitoire (F6.3), boucle de réessai avec backoff, sonde temporisée (F6.2).

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net"
	"testing"
	"time"
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
