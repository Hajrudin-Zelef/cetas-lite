---
id: collect-261001-rattrapage/rattrapage/go-guide-1
title: "Guide Go (Golang) — du script sysadmin au service de production"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["agents", "attention"]
source: docs/RAG/collect-261001-rattrapage/go_guide.md
source_anchor: ""
source_lines: [1, 235]
sha256: b929c0b926bb9299c0c3271eaa98b4ad75f43766317a9c5e61a94fde09aacd1b
---

# Guide Go (Golang) — du script sysadmin au service de production

> Pour Zelef — chef de service systèmes & énergies, sysadmin.
> Angle : pratique, outillage, CLI, automation, déploiement simple.
> Go 1.21+ (certaines syntaxes notées `≥ 1.22`, `≥ 1.23`, etc. le cas échéant).
> Ce guide est dense : tutoriel + référence. Tout le code a été relu pour une syntaxe correcte.

**Plan rapide** : installation → syntaxe → types → structs/interfaces → erreurs →
modules → fichiers → concurrence (goroutines/channels) → réseau/HTTP →
tests → CLI → déploiement → cas pratiques sysadmin → pense-bête → glossaire → quiz.

---

## 1. Pourquoi Go pour un sysadmin ?

| Besoin sysadmin | Réponse Go |
|---|---|
| Scripts d'automation fiables | Binaire compilé, erreurs explicites, pas d'interpréteur à installer |
| Outils CLI distribués aux équipes | Un seul binaire statique, compilation croisée native |
| Supervision / agents / exporters | Concurrence native (goroutines), faible empreinte mémoire |
| Remplacer des scripts shell fragiles | Typage statique, gestion d'erreurs obligatoire, tests intégrés |
| Services internes (API, webhooks) | `net/http` en bibliothèque standard, démarrage en ms |

Go ne remplace pas bash pour un one-liner. Il remplace les scripts de 200+ lignes
qui deviennent critiques, et les outils Python qui exigent un runtime sur chaque machine.

Points d'attention honnêtes :
- Gestion d'erreurs verbeuse (c'est un choix assumé du langage).
- Pas de génériques avant Go 1.18 ; écosystème plus petit que Python pour la data science.
- Le modèle de concurrence est simple mais exige de la discipline (voir sections 37–44).

## 2. Installation sur Linux (Debian/Ubuntu)

Méthode recommandée : archive officielle (les paquets distro sont souvent en retard).

```bash
# 1. Télécharger (vérifier la dernière version sur https://go.dev/dl/)
GO_VERSION="1.23.4"
wget "https://go.dev/dl/go${GO_VERSION}.linux-amd64.tar.gz"

# 2. (Optionnel) vérifier la somme SHA256 publiée sur go.dev/dl
# sha256sum go1.23.4.linux-amd64.tar.gz

# 3. Installer dans /usr/local (supprimer l'ancienne version d'abord)
sudo rm -rf /usr/local/go
sudo tar -C /usr/local -xzf "go${GO_VERSION}.linux-amd64.tar.gz"

# 4. Ajouter au PATH (persistant)
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
# Pour que "go install" dépose les binaires dans le PATH aussi :
echo 'export PATH=$PATH:$HOME/go/bin' >> ~/.bashrc
source ~/.bashrc

# 5. Vérifier
go version        # go version go1.23.4 linux/amd64
go env GOPATH GOROOT GOCACHE GOENV
```

> Sous Windows/macOS : installeur sur https://go.dev/dl/ ; sous Debian `apt install golang`
> existe mais fournit souvent une version ancienne — à éviter pour un usage sérieux.

Checklist post-installation :
- [ ] `go version` affiche la version attendue
- [ ] `$HOME/go/bin` est dans le PATH
- [ ] `go env GOPROXY` répond (proxy de modules, défaut `https://proxy.golang.org,direct`)
- [ ] Un `hello world` compile (section 3)

## 3. Premier programme : le rituel `hello`

```bash
mkdir -p ~/go/src/hello && cd ~/go/src/hello
go mod init hello          # crée go.mod (voir section 5)
```

`main.go` :
```go
package main

import "fmt"

func main() {
	fmt.Println("Hello, infra !")
}
```

```bash
go run .        # compile + exécute sans écrire de binaire
go build -o hello .   # produit le binaire ./hello
./hello
```

Règles à retenir dès le premier fichier :
- Le point d'entrée est **toujours** `func main()` dans `package main`.
- Un programme = un ou plusieurs fichiers `.go` du même package dans un dossier.
- `go run .` compile en mémoire ; `go build` produit un binaire réel.

## 4. La toolchain `go` : les 10 commandes à connaître

| Commande | Rôle |
|---|---|
| `go run .` | Compile et exécute (dev rapide) |
| `go build -o app .` | Produit le binaire |
| `go vet ./...` | Analyse statique (bugs probables) |
| `gofmt -l .` | Liste les fichiers mal formatés |
| `go test ./...` | Lance les tests |
| `go mod tidy` | Nettoie `go.mod`/`go.sum` |
| `go install tool@version` | Installe un outil (`≥ 1.16`, hors module courant) |
| `go list -m all` | Liste les dépendances |
| `go env -w GOPROXY=...` | Écrit la config persistante (`go env GOENV` pour le fichier) |
| `go doc fmt.Println` | Doc d'un symbole, sans quitter le terminal |

Workflow minimal sain avant chaque commit :
```bash
gofmt -l . && go vet ./... && go test ./...
```

## 5. Modules Go : `go mod`, le gestionnaire de dépendances

Depuis Go 1.16, les modules sont le mode par défaut. Oubliez `GOPATH/src` historique.

```bash
go mod init infra/outils          # nom = chemin du module (souvent URL du repo)
go get github.com/spf13/cobra@latest   # ajoute une dépendance
go mod tidy                      # ajoute ce qui manque, retire l'inutile
```

`go.mod` typique :
```go
module infra/outils

go 1.23

require github.com/spf13/cobra v1.8.1

require (
	github.com/spf13/pflag v1.0.5 // indirect
	golang.org/x/sys v0.20.0 // indirect
)
```

- `go.sum` = empreintes cryptographiques des dépendances (à versionner).
- Versionner `go.mod` **et** `go.sum` dans git. Ne jamais éditer `go.sum` à la main.
- `go mod vendor` : copie les dépendances dans `vendor/` (builds offline / air-gap).

Tableau des commandes modules :

| Commande | Effet |
|---|---|
| `go mod init <nom>` | Crée le module |
| `go get pkg@version` | Ajoute/met à jour (`@latest`, `@v1.2.3`, `@commit`) |
| `go get -u ./...` | Met à jour toutes les déps (avec prudence) |
| `go mod tidy` | Synchronise go.mod/go.sum avec le code |
| `go mod verify` | Vérifie l'intégrité du cache |
| `go mod why pkg` | Explique pourquoi une dépe
...[truncated 10263 chars]
---

---

## 6. Variables : déclaration, `:=`, portée

```go
var nom string = "web-01"   // forme longue
var port = 22               // type inféré
hote := "db-01"             // := déclare + infère (UNIQUEMENT dans une fonction)
var a, b int = 1, 2         // plusieurs
var (
	ip      string = "10.0.0.80"
	actif   bool   = true
	retries int           // 0 par défaut (zéro value, voir section 10)
)

const Timeout = 5           // constante : voir section 7

// Portée : minuscule = package, MAJUSCULE = exporté (voir section 30)
```

Règles :
- `:=` exige **au moins une variable nouvelle** à gauche : `x := 1; x, y := 2, 3` est légal.
- Variable déclarée et **jamais utilisée** = erreur de compilation (le compilateur est strict).
- Préférez `:=` en local, `var` au niveau package ou quand le zéro value suffit.

## 7. Constantes et `iota` : les énumérations Go

```go
const (
	SeuilCritique = 90
	SeuilAlerte   = 70
	Version       = "1.4.0"
)

// iota : compteur auto-incrémenté dans un bloc const
const (
	Lundi = iota // 0
	Mardi        // 1 (iota implicite)
	Mercredi     // 2
	_            // 3 : _ saute une valeur
	Vendredi     // 4
)

// Énumération de statuts typée (idiome)
type Statut int

const (
	StatutInconnu Statut = iota
	StatutOK
	StatutKO
	StatutCritique
)

func (s Statut) String() string {
	return []string{"inconnu", "ok", "ko", "critique"}[s]
}
```

Les constantes sont **non typées** par défaut (`const x = 42` s'adapte au contexte),
sauf si on précise le type. `iota` recommence à 0 dans chaque bloc `const`.

## 8. Types numériques : choisir la bonne taille

| Type | Taille | Usage |
|---|---|---|
| `int` / `uint` | 32 ou 64 bits selon plateforme | **Défaut** : boucles, index, compteurs |
| `int8`…`int64`, `uint8`…`uint64` | Fixe | Protocoles, binaire, perf mémoire |
| `float32` / `float64` | 32 / 64 bits | `float64` par défaut (mesures, ratios) |
| `byte` | alias `uint8` | Octets bruts |
| `rune` | alias `int32` | Point de code Unicode |
| `uintptr` | Pointeur non sûr | Rare : interfaçage bas niveau |

```go
var n int = 42
var f float64 = 3.14
var b byte = 'A'      // 65 : les quotes simples = rune/octet, pas string !
var s string = "A"    // doubles quotes = string

