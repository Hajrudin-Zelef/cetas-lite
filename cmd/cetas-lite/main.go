package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"cetas-lite/internal/auth"
	"cetas-lite/internal/config"
	"cetas-lite/internal/cryptovault"
	"cetas-lite/internal/store"
	"cetas-lite/internal/web"
)

var version = "dev"

func usage() {
	fmt.Fprint(os.Stderr, `cetas-lite — assistant IA auto-heberge

Usage:
  cetas-lite serve                 demarre le serveur HTTP
  cetas-lite version               affiche la version
  cetas-lite keys list             liste les providers configures
  cetas-lite keys set <p> [valeur] chiffre et enregistre une cle (valeur sinon sur stdin)
  cetas-lite keys delete <p>       supprime une cle

Variables:
  CETAS_LITE_HOME                 repertoire de donnees (defaut: config/cetas-lite)
  CETAS_LITE_ADDR                 adresse d'ecoute (defaut: 127.0.0.1:8787)
  CETAS_LITE_REGISTRATION_OPEN    autoriser l'inscription (defaut: true)
  CETAS_LITE_VAULT_PASSWORD       mot de passe du coffre (requis pour keys)
`)
}

func main() {
	if len(os.Args) < 2 {
		usage()
		os.Exit(2)
	}
	var err error
	switch os.Args[1] {
	case "serve":
		err = runServe()
	case "version":
		fmt.Println("cetas-lite", version)
	case "keys":
		err = runKeys(os.Args[2:])
	case "help", "-h", "--help":
		usage()
	default:
		usage()
		os.Exit(2)
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, "erreur:", err)
		os.Exit(1)
	}
}

func runServe() error {
	cfg, err := config.Load()
	if err != nil {
		return err
	}
	st, err := store.Open(cfg.DBPath)
	if err != nil {
		return err
	}
	defer func() { _ = st.Close() }()

	authMgr, err := auth.New(st)
	if err != nil {
		return err
	}

	srv := web.New(cfg, st, authMgr, version)
	httpSrv := &http.Server{
		Addr:              cfg.Addr,
		Handler:           srv.Handler(),
		ReadHeaderTimeout: 10 * time.Second,
	}

	errCh := make(chan error, 1)
	go func() {
		slog.Info("cetas-lite demarre", "addr", cfg.Addr, "home", cfg.Home, "version", version)
		if err := httpSrv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			errCh <- err
		}
	}()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	select {
	case err := <-errCh:
		return err
	case <-ctx.Done():
		slog.Info("arret en cours")
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		return httpSrv.Shutdown(shutdownCtx)
	}
}

func openStore() (*config.Config, *store.Store, error) {
	cfg, err := config.Load()
	if err != nil {
		return nil, nil, err
	}
	st, err := store.Open(cfg.DBPath)
	if err != nil {
		return nil, nil, err
	}
	return cfg, st, nil
}

func runKeys(args []string) error {
	if len(args) < 1 {
		return errors.New("usage: cetas-lite keys set|list|delete [provider] [valeur]")
	}
	_, st, err := openStore()
	if err != nil {
		return err
	}
	defer func() { _ = st.Close() }()

	switch args[0] {
	case "list":
		names, err := st.ListSecrets()
		if err != nil {
			return err
		}
		if len(names) == 0 {
			fmt.Println("(aucune cle)")
			return nil
		}
		for _, n := range names {
			fmt.Println(n)
		}
		return nil
	case "delete":
		if len(args) < 2 {
			return errors.New("provider requis")
		}
		if err := st.DeleteSecret(args[1]); err != nil {
			return err
		}
		fmt.Println("cle supprimee:", args[1])
		return nil
	case "set":
		if len(args) < 2 {
			return errors.New("provider requis")
		}
		provider := args[1]
		value := ""
		if len(args) >= 3 {
			value = args[2]
		} else {
			raw, err := io.ReadAll(os.Stdin)
			if err != nil {
				return err
			}
			value = strings.TrimSpace(string(raw))
		}
		if value == "" {
			return errors.New("valeur vide")
		}
		vault, err := openVault(st)
		if err != nil {
			return err
		}
		ct, err := vault.Encrypt([]byte(value), []byte(provider))
		if err != nil {
			return err
		}
		if err := st.PutSecret(provider, ct); err != nil {
			return err
		}
		fmt.Println("cle enregistree:", provider)
		return nil
	default:
		return fmt.Errorf("action inconnue: %s", args[0])
	}
}

func openVault(st *store.Store) (*cryptovault.Vault, error) {
	password := os.Getenv("CETAS_LITE_VAULT_PASSWORD")
	if password == "" {
		return nil, errors.New("CETAS_LITE_VAULT_PASSWORD requis")
	}
	salt, ok := st.GetMeta("vault_salt")
	if !ok {
		var err error
		salt, err = cryptovault.NewSalt()
		if err != nil {
			return nil, err
		}
		if err := st.PutMeta("vault_salt", salt); err != nil {
			return nil, err
		}
	}
	return cryptovault.New(password, salt)
}
