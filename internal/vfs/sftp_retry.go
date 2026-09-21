package vfs

import (
	"context"
	"errors"
	"io"
	"net"
	"strings"
	"time"

	"github.com/pkg/sftp"
)

// Fiabilité SFTP — phase 4.
//
// F6.2 : sonde de santé temporisée. La connexion n'est plus sondée à chaque
// opération (un aller-retour Stat par appel) : la sonde ne s'exécute que si
// la connexion est inactive depuis sftpProbeInterval. Les rafales
// d'opérations de l'agent ne paient plus ce coût ; une connexion morte
// restée sous le seuil est de toute façon récupérée par le réessai F6.3.
//
// F6.3 : réessais sur erreur réseau transitoire. Chaque opération fichier
// est réessayée (backoff exponentiel) quand l'échec ressemble à un incident
// réseau (connexion coupée, reset, timeout). Les réécritures sont
// idempotentes (ouverture O_TRUNC + réécriture complète), les lectures
// aussi. Les erreurs métier (introuvable, hors racine) et le contexte
// annulé ne sont jamais réessayés.

// sftpProbeInterval : inactivité au-delà de laquelle la sonde de santé
// s'exécute avant réutilisation de la connexion (F6.2). Variable pour
// les tests.
var sftpProbeInterval = 30 * time.Second

// needProbe indique s'il faut sonder la connexion existante (F6.2).
func needProbe(last, now time.Time) bool {
	return now.Sub(last) >= sftpProbeInterval
}

// sftpMaxRetries : réessais après l'échec initial (1 essai + 2 réessais).
const sftpMaxRetries = 2

// sftpRetryBackoff : délai avant le premier réessai, doublé ensuite.
const sftpRetryBackoff = 200 * time.Millisecond

// transientSubstrings : fragments de messages typiques d'un incident
// réseau (connexion coupée, reset, timeout). Complète la détection par
// types (net.Error) pour les erreurs déjà « aplaties » en chaîne.
var transientSubstrings = []string{
	"connection reset by peer",
	"broken pipe",
	"use of closed network connection",
	"connection closed",
	"unexpected eof",
	"connection timed out",
	"i/o timeout",
	"no route to host",
	"network is unreachable",
	"connection refused",
}

// isTransientNetErr : l'erreur ressemble-t-elle à un incident réseau
// transitoire justifiant un réessai (F6.3) ? Les erreurs métier
// (introuvable, hors racine) et le contexte annulé ne sont jamais
// transitoires : elles sont retournées immédiatement, sans réessai.
func isTransientNetErr(err error) bool {
	if err == nil {
		return false
	}
	if errors.Is(err, ErrNotFound) ||
		errors.Is(err, ErrOutsideRoot) ||
		errors.Is(err, context.Canceled) ||
		errors.Is(err, context.DeadlineExceeded) {
		return false
	}
	var nerr net.Error
	if errors.As(err, &nerr) {
		return true
	}
	var operr *net.OpError
	if errors.As(err, &operr) {
		return true
	}
	if errors.Is(err, io.ErrUnexpectedEOF) {
		return true
	}
	msg := strings.ToLower(err.Error())
	for _, sub := range transientSubstrings {
		if strings.Contains(msg, sub) {
			return true
		}
	}
	return false
}

// retryTransient exécute attempt ; si l'erreur est réseau-transitoire,
// reset() invalide la connexion puis l'opération est réessayée avec un
// backoff exponentiel, jusqu'à sftpMaxRetries réessais (F6.3).
func retryTransient(ctx context.Context, reset func(), attempt func() error) error {
	backoff := sftpRetryBackoff
	var err error
	for i := 0; ; i++ {
		err = attempt()
		if err == nil || !isTransientNetErr(err) || i >= sftpMaxRetries {
			return err
		}
		reset()
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(backoff):
		}
		backoff *= 2
	}
}

// doRetry exécute op avec un client SFTP : en cas d'erreur réseau
// transitoire, la connexion est invalidée puis rétablie paresseusement,
// et l'opération est réessayée (F6.3).
func (s *SFTPFS) doRetry(ctx context.Context, op func(cl *sftp.Client) error) error {
	return retryTransient(ctx, s.resetConn, func() error {
		cl, err := s.client(ctx)
		if err != nil {
			return err
		}
		return op(cl)
	})
}

// resetConn invalide la connexion (entre deux tentatives).
func (s *SFTPFS) resetConn() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.closeLocked()
}
