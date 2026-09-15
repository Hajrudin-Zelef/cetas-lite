package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"runtime"
	"strings"
	"syscall"
	"time"

	"cetas-lite/internal/alias"
	"cetas-lite/internal/attach"
	"cetas-lite/internal/auth"
	"cetas-lite/internal/backup"
	"cetas-lite/internal/chat"
	"cetas-lite/internal/config"
	"cetas-lite/internal/customtools"
	"cetas-lite/internal/desktop"
	"cetas-lite/internal/local"
	"cetas-lite/internal/mcp"
	"cetas-lite/internal/memory"
	"cetas-lite/internal/modelcaps"
	"cetas-lite/internal/plugins"
	"cetas-lite/internal/provider"
	"cetas-lite/internal/search"
	"cetas-lite/internal/store"
	"cetas-lite/internal/terminal"
	"cetas-lite/internal/vault"
	"cetas-lite/internal/web"
	"cetas-lite/internal/workspace"
	"cetas-lite/internal/worktree"
)

var version = "dev"

func usage() {
	fmt.Fprint(os.Stderr, `cetas-lite — assistant IA auto-heberge

Usage:
  cetas-lite serve                 demarre le serveur HTTP
  cetas-lite desktop               application bureau (fenetre native, Windows)
  cetas-lite version               affiche la version
  cetas-lite keys list             liste les providers configures
  cetas-lite keys set <p> [valeur] chiffre et enregistre une cle (valeur sinon sur stdin)
  cetas-lite keys delete <p>       supprime une cle
  cetas-lite mcp                   liste les serveurs MCP et leurs outils
  cetas-lite tools                 liste les outils personnalises (tools.json)
  cetas-lite backup <fichier>      sauvegarde la base (+ mcp.json) dans un tar.gz (serveur arrete)
  cetas-lite restore <fichier>     restaure la base depuis un tar.gz (serveur arrete)

Variables:
  CETAS_LITE_HOME                 repertoire de donnees (defaut: config/cetas-lite)
  CETAS_LITE_ADDR                 adresse d'ecoute (defaut: 127.0.0.1:8787)
  CETAS_LITE_REGISTRATION_OPEN    inscription: non defini = 1er compte puis ferme ; true = ouvert ; false = ferme
  CETAS_LITE_TRUST_PROXY          faire confiance a X-Forwarded-For (derriere nginx) (defaut: false)
  CETAS_LITE_SANDBOX              isolation Bash/RunScript: none|auto|bwrap (defaut: none)
  CETAS_LITE_VAULT_PASSWORD       mot de passe du coffre (requis pour keys)
`)
}

func main() {
	if len(os.Args) < 2 {
		// Double-clic sous Windows : ouvrir directement l'application bureau.
		if runtime.GOOS == "windows" {
			if err := runDesktop(); err != nil {
				fmt.Fprintln(os.Stderr, "erreur:", err)
				os.Exit(1)
			}
			return
		}
		usage()
		os.Exit(2)
	}
	var err error
	switch os.Args[1] {
	case "serve":
		err = runServe()
	case "desktop":
		err = runDesktop()
	case "version":
		fmt.Println("cetas-lite", version)
	case "keys":
		err = runKeys(os.Args[2:])
	case "mcp":
		err = runMCP(os.Args[2:])
	case "tools":
		err = runTools(os.Args[2:])
	case "backup":
		err = runBackup(os.Args[2:])
	case "restore":
		err = runRestore(os.Args[2:])
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

// app regroupe le handler HTTP et le nettoyage, commun a `serve` et `desktop`.
type app struct {
	cfg     *config.Config
	handler http.Handler
	cleanup func()
}

func buildApp() (*app, error) {
	cfg, err := config.Load()
	if err != nil {
		return nil, err
	}
	st, err := store.Open(cfg.DBPath)
	if err != nil {
		return nil, err
	}

	authMgr, err := auth.New(st)
	if err != nil {
		_ = st.Close()
		return nil, err
	}

	client := provider.NewHTTPClient()
	localURLs := localURLsFromEnv()
	// Les URL definies depuis l'interface (Configuration -> API et Modeles,
	// meta "local_url_<id>") priment sur les variables d'environnement.
	for _, id := range provider.LocalEngines {
		if raw, ok := st.GetMeta("local_url_" + id); ok && len(raw) > 0 {
			localURLs[id] = string(raw)
		}
	}
	keys := loadProviderKeys(st)
	registry := provider.Build(keys, localURLs, client)
	discover := local.New(local.DefaultEngines(localURLs), client)
	families := loadFamilies(st)
	engine := chat.NewEngine(registry, families, st, discover, cfg.WorkspaceDir)
	engine.SetAllowScript(cfg.AllowScript)
	iso, err := chat.ResolveIsolation(cfg.Sandbox, cfg.WorkspaceDir)
	if err != nil {
		_ = st.Close()
		return nil, err
	}
	engine.SetIsolation(iso)
	if iso == chat.IsolationBwrap {
		slog.Info("isolation bwrap active")
	}
	engine.SetSearcher(search.NewWithConfig(web.LoadSearchConfig(st, keys), client))
	engine.SetMemory(memory.New(cfg.MemoryDir))
	attachStore := attach.New(filepath.Join(cfg.Home, "uploads"), 20<<20)
	engine.SetAttachments(attachStore)
	// Les pièces jointes ne sont plus supprimées à l'envoi (le tour agent
	// les lit de façon asynchrone, la régénération peut les relire) : on
	// purge ici les orphelins de plus de 7 jours à chaque démarrage.
	if n := attachStore.CleanOlderThan(7 * 24 * time.Hour); n > 0 {
		slog.Info("pieces jointes orphelines purgees", "count", n)
	}
	engine.SetCapabilities(modelcaps.Load(st))
	engine.SetMarexPath(filepath.Join(cfg.Home, "MAREX.md"))
	customManager, err := customtools.NewManager(cfg.ToolsPath, client)
	if err != nil {
		_ = st.Close()
		return nil, err
	}
	engine.SetCustom(customManager)
	pluginManager := plugins.NewManager(cfg.PluginsDir, client)
	if err := pluginManager.Load(); err != nil {
		_ = st.Close()
		return nil, err
	}
	engine.SetPlugins(pluginManager)
	if infos := pluginManager.Plugins(); len(infos) > 0 {
		slog.Info("plugins charges", "count", len(infos), "dir", cfg.PluginsDir)
	}
	for _, perr := range pluginManager.Errors() {
		slog.Warn("plugin ignore", "erreur", perr)
	}
	mcpManager, err := mcp.NewManager(cfg.MCPPath, version)
	if err != nil {
		_ = st.Close()
		return nil, err
	}
	engine.SetMCP(mcpManager)
	if mcpManager.Configured() {
		slog.Info("mcp configure", "path", cfg.MCPPath)
	}
	wtMgr, err := worktree.NewManager(filepath.Join(cfg.Home, "worktrees"))
	if err != nil {
		_ = st.Close()
		return nil, err
	}
	engine.SetWorktreeManager(wtMgr)
	// Projets de l'agent : upload local ou dossier distant SFTP.
	wsMgr := workspace.New(cfg.Home, st, vaultLazy{st: st})
	engine.SetWorkspaceManager(wsMgr)
	termMgr := terminal.NewManager(cfg.WorkspaceDir, wtMgr.Root())

	srv := web.New(cfg, st, authMgr, engine, termMgr, registry, client, version)
	srv.SetWorkspaceManager(wsMgr)
	return &app{
		cfg:     cfg,
		handler: srv.Handler(),
		cleanup: func() {
			wsMgr.CloseAll()
			termMgr.Close()
			mcpManager.Close()
			_ = st.Close()
		},
	}, nil
}

// vaultLazy ouvre le coffre chiffré à la demande pour le gestionnaire de
// projets (les secrets SFTP sont chiffrés avant stockage).
type vaultLazy struct{ st *store.Store }

func (v vaultLazy) Encrypt(plain, aad []byte) ([]byte, error) {
	vv, err := vault.Open(v.st)
	if err != nil {
		return nil, err
	}
	return vv.Encrypt(plain, aad)
}

func (v vaultLazy) Decrypt(data, aad []byte) ([]byte, error) {
	vv, err := vault.Open(v.st)
	if err != nil {
		return nil, err
	}
	return vv.Decrypt(data, aad)
}

func runServe() error {
	a, err := buildApp()
	if err != nil {
		return err
	}
	defer a.cleanup()

	httpSrv := &http.Server{
		Addr:              a.cfg.Addr,
		Handler:           a.handler,
		ReadHeaderTimeout: 10 * time.Second,
	}

	errCh := make(chan error, 1)
	go func() {
		slog.Info("cetas-lite demarre", "addr", a.cfg.Addr, "home", a.cfg.Home, "version", version)
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

// runDesktop demarre le serveur en local (127.0.0.1, port ephemere) puis ouvre
// une fenetre native (WebView2 sous Windows). La fermeture de la fenetre arrete
// proprement le serveur.
func runDesktop() error {
	a, err := buildApp()
	if err != nil {
		return err
	}
	defer a.cleanup()

	// En mode GUI Windows il n'y a pas de console : journaliser aussi dans un fichier.
	if f, ferr := os.OpenFile(filepath.Join(a.cfg.Home, "desktop.log"), os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600); ferr == nil {
		slog.SetDefault(slog.New(slog.NewTextHandler(io.MultiWriter(os.Stderr, f), nil)))
		defer func() { _ = f.Close() }()
	}

	addr, shutdown, err := desktop.ServeLocal(a.handler)
	if err != nil {
		return err
	}
	url := "http://" + addr + "/"
	slog.Info("cetas-lite bureau", "url", url, "home", a.cfg.Home, "version", version)

	debug := os.Getenv("CETAS_LITE_DEBUG") == "true"
	if err := desktop.Open(url, "Cetas", 1280, 800, debug); err != nil {
		return err
	}
	slog.Info("fenetre fermee, arret en cours")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	return shutdown(ctx)
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
		vault, err := vault.Open(st)
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

func runMCP(args []string) error {
	cfg, err := config.Load()
	if err != nil {
		return err
	}
	mgr, err := mcp.NewManager(cfg.MCPPath, version)
	if err != nil {
		return err
	}
	defer mgr.Close()
	if !mgr.Configured() {
		fmt.Println("(aucun serveur mcp configure dans " + cfg.MCPPath + ")")
		return nil
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	tools := mgr.Tools(ctx)
	for _, s := range mgr.Servers() {
		status := "hors ligne"
		if s.Connected {
			status = "connecte"
		}
		line := fmt.Sprintf("%s\t%s\t%s\t%d outil(s)", s.Name, s.Transport, status, s.Tools)
		if s.Error != "" {
			line += "  [" + s.Error + "]"
		}
		fmt.Println(line)
		for _, t := range tools {
			if t.Server == s.Name {
				fmt.Printf("  - %s  (%s)\n", t.Exposed, t.Name)
			}
		}
	}
	return nil
}

func runBackup(args []string) error {
	if len(args) < 1 {
		return errors.New("usage: cetas-lite backup <fichier.tar.gz>")
	}
	cfg, err := unlockedConfig()
	if err != nil {
		return err
	}
	if err := backup.Create(cfg.DBPath, cfg.MCPPath, args[0]); err != nil {
		return err
	}
	fmt.Println("sauvegarde ecrite:", args[0])
	return nil
}

func runRestore(args []string) error {
	if len(args) < 1 {
		return errors.New("usage: cetas-lite restore <fichier.tar.gz>")
	}
	cfg, err := unlockedConfig()
	if err != nil {
		return err
	}
	if err := backup.Restore(args[0], cfg.Home); err != nil {
		return err
	}
	fmt.Println("restauration effectuee dans", cfg.Home)
	return nil
}

func unlockedConfig() (*config.Config, error) {
	cfg, st, err := openStore()
	if err == nil {
		_ = st.Close()
		return cfg, nil
	}
	if errors.Is(err, store.ErrLocked) {
		return nil, errors.New("base verrouillee: arretez le serveur")
	}
	return config.Load()
}

func runTools(args []string) error {
	cfg, err := config.Load()
	if err != nil {
		return err
	}
	mgr, err := customtools.NewManager(cfg.ToolsPath, nil)
	if err != nil {
		return err
	}
	if !mgr.Configured() {
		fmt.Println("(aucun outil personnalise dans " + cfg.ToolsPath + ")")
		return nil
	}
	for _, d := range mgr.Defs() {
		fmt.Printf("%s\t%s\n", d.Name, d.Description)
	}
	return nil
}

func localURLsFromEnv() map[string]string {
	return map[string]string{
		"llamacpp": strings.TrimSpace(os.Getenv("CETAS_LITE_LLAMACPP_URL")),
		"ollama":   strings.TrimSpace(os.Getenv("CETAS_LITE_OLLAMA_URL")),
		"lmstudio": strings.TrimSpace(os.Getenv("CETAS_LITE_LMSTUDIO_URL")),
	}
}

func loadProviderKeys(st *store.Store) map[string]string {
	keys := map[string]string{}
	names, err := st.ListSecrets()
	if err != nil || len(names) == 0 {
		return keys
	}
	vault, err := vault.Open(st)
	if err != nil {
		slog.Warn("cles chiffrees ignorees (coffre indisponible)", "raison", err)
		return keys
	}
	for _, name := range names {
		ct, ok := st.GetSecret(name)
		if !ok {
			continue
		}
		pt, err := vault.Decrypt(ct, []byte(name))
		if err != nil {
			continue
		}
		keys[name] = string(pt)
	}
	return keys
}

func loadFamilies(st *store.Store) []alias.Family {
	fams := alias.Defaults()
	raw, ok := st.GetMeta("aliases")
	if !ok {
		return fams
	}
	var ov alias.Overrides
	if err := json.Unmarshal(raw, &ov); err != nil {
		return fams
	}
	return alias.Apply(fams, ov)
}
