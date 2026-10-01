---
id: collect-261001-rattrapage/rattrapage/go-guide-7
title: "Guide Go (Golang) — du script sysadmin au service de production"
domain: rattrapage
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: ["agent", "arr", "decode"]
source: docs/RAG/collect-261001-rattrapage/go_guide.md
source_anchor: ""
source_lines: [1647, 1927]
sha256: 05141ccae5fdd275243776b480ef0fd34b58fa15228051affef6c4cd4bc7bad9
---

# Guide Go (Golang) — du script sysadmin au service de production

func creerHote(w http.ResponseWriter, r *http.Request) {
	var h Hote
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields() // rejette les champs inconnus : config stricte
	if err := dec.Decode(&h); err != nil {
		http.Error(w, "JSON invalide: "+err.Error(), http.StatusBadRequest)
		return
	}
	if h.Nom == "" || h.IP == "" {
		http.Error(w, "nom et ip requis", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"id": h.Nom, "statut": "créé"})
}
```

## 50. Logs : `log` puis `slog` (structuré, ≥ 1.21)

```go
// Simple (stdlib historique)
log.Println("démarrage")
log.Printf("hôte %s injoignable: %v", h, err)
log.Fatal(err) // affiche + os.Exit(1) — à réserver à main/init !

// Structuré avec slog (Go 1.21+) : indispensable en prod/supervision
import "log/slog"

logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
	Level: slog.LevelInfo,
}))
slog.SetDefault(logger)

slog.Info("sonde terminée", "hote", h, "latence_ms", 12, "ok", true)
slog.Warn("retry", "tentative", 2, "hote", h)
slog.Error("échec", "err", err)
// Sortie : {"time":"...","level":"INFO","msg":"sonde terminée","hote":"web-01",...}
```

Niveaux : `Debug < Info < Warn < Error`. En prod : JSON vers stdout,
collecté par l'agent de logs (pas de fichiers de logs gérés à la main :
principe 12-factor, et `logrotate` en plan B).

## 51. Exécuter des commandes externes : `os/exec`

Le pont entre Go et l'existant shell.

```go
// Simple : capturer la sortie
out, err := exec.Command("uptime").Output()
if err != nil {
	if ee, ok := err.(*exec.ExitError); ok {
		fmt.Println("code:", ee.ExitCode(), "stderr:", string(ee.Stderr))
	}
}

// Avec contexte (timeout qui TUE le processus)
ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
defer cancel()
cmd := exec.CommandContext(ctx, "rsync", "-av", "/src/", "/dst/")
cmd.Stdout = os.Stdout // streamer au lieu de bufferiser
cmd.Stderr = os.Stderr
if err := cmd.Run(); err != nil {
	log.Fatalf("rsync: %v", err)
}

// Entrée stdin + shell complet (avec prudence : injection !)
cmd = exec.Command("sh", "-c", "df -h | awk 'NR>1 {print $5, $6}'")
// JAMAIS de concaténation de variable utilisateur dans "sh -c" sans validation.

// Vérifier la présence d'un binaire
if _, err := exec.LookPath("dig"); err != nil {
	log.Fatal("dig introuvable dans le PATH")
}
```

Sécurité : préférez `exec.Command("dig", "+short", hote)` (args séparés, pas de shell)
à `sh -c` avec interpolation. Voir section 66.

## 52. Signaux OS : arrêt propre (SIGTERM/SIGINT)

```go
func attendreSignal() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	sig := <-sigs
	fmt.Println("\nsignal reçu:", sig)
	// puis : cancel(), srv.Shutdown(), wg.Wait()...
}

// Version avec context : la plus propre
func contexteSignal() (context.Context, context.CancelFunc) {
	return signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
}

func main() {
	ctx, stop := contexteSignal()
	defer stop()
	// ... serveurs, workers utilisent ctx ...
	<-ctx.Done() // bloqué jusqu'au Ctrl+C / SIGTERM (systemd !)
	fmt.Println("arrêt demandé, nettoyage...")
}
```

Sous systemd, `SIGTERM` est le signal d'arrêt par défaut : un binaire Go qui
l'intercepte s'arrête proprement (flush, close, déréférencement). Voir section 72.

## 53. CLI simple avec `flag` (stdlib)

```go
package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	var (
		hote    = flag.String("hote", "localhost", "hôte à sonder")
		port    = flag.Int("port", 22, "port TCP")
		timeout = flag.Duration("timeout", 5*time.Second, "délai max")
		verbeux = flag.Bool("v", false, "mode verbeux")
	)
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options]\nSonde TCP d'un hôte.\n\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()

	if flag.NArg() > 0 {
		fmt.Println("args positionnels:", flag.Args())
	}
	fmt.Printf("sonde %s:%d (timeout %s, verbeux=%v)\n", *hote, *port, *timeout, *verbeux)
}
```

`flag` suffit pour 80 % des outils internes. Pour des sous-commandes (`monoutil
hotes list`), passez à Cobra (section 54).

## 54. CLI avancée : Cobra + Viper

```bash
go get github.com/spf13/cobra@latest
go get github.com/spf13/viper@latest
```

```go
package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "monoutil",
	Short: "Boîte à outils infra",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// Viper : config fichier + env + flags, par priorité
		viper.SetConfigName("monoutil")
		viper.AddConfigPath("/etc/monoutil")
		viper.AddConfigPath("$HOME/.monoutil")
		viper.SetEnvPrefix("MONOUTIL")
		viper.AutomaticEnv()
		_ = viper.ReadInConfig() // absent = ok, les flags/env suffisent
	},
}

var sondeCmd = &cobra.Command{
	Use:   "sonde [hote]",
	Short: "Sonde TCP",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		port, _ := cmd.Flags().GetInt("port")
		timeout := viper.GetDuration("timeout") // flag > env > config > défaut
		return sonder(args[0], port, timeout)
	},
}

func main() {
	sondeCmd.Flags().Int("port", 22, "port TCP")
	sondeCmd.Flags().Duration("timeout", 5_000_000_000, "délai max") // 5s
	_ = viper.BindPFlag("timeout", sondeCmd.Flags().Lookup("timeout"))
	rootCmd.AddCommand(sondeCmd)
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
```

Ordre de priorité Viper : **flag > variable d'environnement > fichier config > défaut**.
Idéal : comportement surchargeable sans recompiler.

## 55. Compilation croisée : un binaire par OS/arch

```bash
# Lister les cibles supportées
go tool dist list | grep -E 'linux|windows|darwin' | head

# Compiler pour Windows depuis Linux (CGO désactivé = 100% statique)
GOOS=windows GOARCH=amd64 go build -o monoutil.exe .

# Linux ARM64 (Raspberry Pi, serveurs ARM)
GOOS=linux GOARCH=arm64 go build -o monoutil-arm64 .

# macOS Apple Silicon
GOOS=darwin GOARCH=arm64 go build -o monoutil-mac .

# Binaire allégé : sans symboles de debug ni table des symboles
go build -ldflags="-s -w" -o monoutil .
# -s : strip symbol table | -w : strip DWARF => -20 à -30% de taille

# Injecter version/date à la compilation
go build -ldflags="-s -w -X main.version=$(git describe --tags) -X main.date=$(date -u +%FT%TZ)" .
```

```go
var (
	version = "dev"     // surchargé par -X main.version=...
	date    = "inconnue"
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "--version" {
		fmt.Printf("monoutil %s (%s)\n", version, date)
		return
	}
	// ...
}
```

> Avec `CGO_ENABLED=0` (défaut si aucun import C), le binaire est **statique** :
> déposable tel quel sur n'importe quelle machine du même OS/arch, y compris
> dans une image `scratch` Docker ou sur un équipement minimal.

## 56. Build tags : code conditionnel par plateforme

```go
//go:build linux
// +build linux   // ligne legacy, garder les deux pour < 1.17 (rare)

package main

import "syscall"

func limites() syscall.Rlimit { /* spécifique Linux */ return syscall.Rlimit{} }
```

```go
//go:build windows

package main

func limites() struct{} { /* stub Windows */ return struct{}{} }
```

Fichier `limites_linux.go` / `limites_windows.go` : le compilateur choisit selon
`GOOS`. Utile pour le code vraiment spécifique (syscall, chemins). Pour le reste,
`runtime.GOOS` suffit :

```go
if runtime.GOOS == "windows" {
	// ...
}
```

## 57. Tests natifs : `testing`, table-driven

Fichier `sonde_test.go` (même package, suffixe `_test.go`) :

```go
package main

import "testing"

