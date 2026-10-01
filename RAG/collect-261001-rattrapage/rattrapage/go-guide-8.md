---
id: collect-261001-rattrapage/rattrapage/go-guide-8
title: "Guide Go (Golang) — du script sysadmin au service de production"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["benchmarks"]
source: docs/RAG/collect-261001-rattrapage/go_guide.md
source_anchor: ""
source_lines: [1928, 2141]
sha256: 9055e76cfa4126118e7a3bbadf8960abf6f84866f83c900510badc1664d5ea74
---

# Guide Go (Golang) — du script sysadmin au service de production

// Test table-driven : LE standard Go
func TestAdresse(t *testing.T) {
	tests := []struct {
		nom     string
		ip      string
		port    int
		attendu string
	}{
		{"basique", "10.0.0.53", 53, "10.0.0.53:53"},
		{"port zéro", "10.0.0.1", 0, "10.0.0.1:0"},
	}
	for _, tt := range tests {
		t.Run(tt.nom, func(t *testing.T) {
			s := Serveur{IP: tt.ip, Port: tt.port}
			if got := s.Adresse(); got != tt.attendu {
				t.Errorf("Adresse() = %q, attendu %q", got, tt.attendu)
			}
		})
	}
}

func TestDiviserParZero(t *testing.T) {
	if _, err := diviser(1, 0); err == nil {
		t.Fatal("erreur attendue pour b=0")
	}
}
```

```bash
go test ./...              # tout le module
go test -run TestAdresse . # un seul test
go test -v ./...           # verbeux
go test -race ./...        # détecteur de data races (CI obligatoire)
go test -cover ./...       # couverture
go test -coverprofile=c.out ./... && go tool cover -html=c.out
```

## 58. Benchmarks et exemples

```go
func BenchmarkConcat(b *testing.B) {
	hotes := []string{"a", "b", "c", "d"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var sb strings.Builder
		for _, h := range hotes {
			sb.WriteString(h)
		}
		_ = sb.String()
	}
}

// Exemple : apparaît dans la doc (go doc) et est TESTÉ
func ExampleServeur_Adresse() {
	s := Serveur{IP: "10.0.0.53", Port: 53}
	fmt.Println(s.Adresse())
	// Output: 10.0.0.53:53
}
```

```bash
go test -bench=. -benchmem ./...   # -benchmem : allocations par opération
```

Sortie typique :
```
BenchmarkConcat-8   5000000   240 ns/op   32 B/op   2 allocs/op
```
Objectif perf : réduire `B/op` et `allocs/op` d'abord (moins d'allocations =
moins de travail pour le GC).

## 59. Fuzzing natif (Go 1.18+)

```go
func FuzzParseLigne(f *testing.F) {
	// Graines : cas connus
	f.Add("web-01,10.0.0.80,up")
	f.Add("malformé")
	f.Fuzz(func(t *testing.T, ligne string) {
		h, err := parseLigne(ligne)
		if err == nil && h.IP == "" {
			t.Fatalf("IP vide sans erreur pour %q", ligne)
		}
		// Invariant : ne doit jamais paniquer
	})
}
```

```bash
go test -fuzz=FuzzParseLigne -fuzztime=30s .
```

Le fuzzer génère des entrées aléatoires pour faire paniquer votre parser :
excellent pour tout ce qui lit des données externes (logs, CSV, réponses d'API).

## 60. Debug : `delve`, `pprof`, traces

```bash
# Delve (debugueur) : go install github.com/go-delve/delve/cmd/dlv@latest
dlv debug .              # lance avec debugueur
# (dlv) break main.go:42 | (dlv) continue | (dlv) print var | (dlv) step

# Profiling intégré : exposer pprof sur un port interne
import _ "net/http/pprof"
go func() { http.ListenAndServe("127.0.0.1:6060", nil) }()

# Capturer 30s de profil CPU puis analyser
go tool pprof http://127.0.0.1:6060/debug/pprof/profile?seconds=30
# (pprof) top10 | (pprof) list maFonction | (pprof) web (graphique)

# Profil mémoire
go tool pprof http://127.0.0.1:6060/debug/pprof/heap

# Trace d'exécution (goroutines, GC, syscalls)
curl -o trace.out http://127.0.0.1:6060/debug/pprof/trace?seconds=5
go tool trace trace.out
```

> N'exposez `pprof` que sur `127.0.0.1` ou un réseau admin : il révèle les
> chemins, une partie de la mémoire et permet de geler le processus.

Réflexes debug sans outils :
- `fmt.Printf("%+v\n", ...)` ciblé (retirer après).
- `go run -race` pour les data races (le bug de concurrence n°1).
- `GODEBUG=gctrace=1` pour voir l'activité du ramasse-miettes.
---

## 61. Performance : les 20 % qui comptent

1. **Mesurez d'abord** : `pprof` avant toute optimisation (section 60).
2. **Allocations** : `strings.Builder`, `make` avec capacité, `sync.Pool` pour les buffers chauds.
3. **Évitez les copies** : receveurs pointeurs, `[]byte` plutôt que `string` en boucle chaude.
4. **I/O > CPU** : en sysadmin, le goulot est le réseau/disque → concurrence (workers), pas micro-optims.
5. **GC** : `GOGC=100` par défaut ; baisser `GOGC` = GC plus fréquent = moins de pics mémoire.

```go
// Mauvais : concaténation en boucle => O(n²) allocations
var s string
for _, h := range hotes {
	s += h + "\n"
}

// Bon : Builder, une seule allocation finale (amortie)
var b strings.Builder
b.Grow(estimOctets) // si estimable
for _, h := range hotes {
	b.WriteString(h)
	b.WriteByte('\n')
}
s = b.String()
```

Tableau des leviers par coût/bénéfice :

| Levier | Coût | Gain typique |
|---|---|---|
| Bufferiser les I/O (`bufio`) | Faible | Énorme sur gros fichiers |
| Worker pool (paralléliser I/O) | Moyen | x10–x50 sur sondes réseau |
| Réduire les allocations | Moyen | -30 % CPU sur chemins chauds |
| `sync.Pool` | Moyen | Visible > 10k obj/s |
| Assembleur / unsafe | Élevé | Rarement justifié : ne pas y toucher |

## 62. Sécurité : checklist du binaire Go

- [ ] **Aucun secret dans le code** : `os.Getenv`, fichiers, Vault — jamais en dur, jamais dans `go.mod`.
- [ ] `go vet ./...` propre ; dépendances à jour (`go list -m -u all`, `go get -u` ciblé).
- [ ] Vulnérabilités connues : `go install golang.org/x/vuln/cmd/govulncheck@latest` puis `govulncheck ./...`.
- [ ] TLS : `MinVersion: tls.VersionTLS12`, vérifier les certificats (ne jamais `InsecureSkipVerify: true` en prod).
- [ ] `os/exec` : args séparés, jamais `sh -c` avec interpolation non validée (section 51).
- [ ] HTTP : timeouts partout (client **et** serveur, sections 47–48), limites de taille (`http.MaxBytesReader`).
- [ ] Fichiers : droits minimaux (`0o600` pour secrets), `MkdirAll` avec `0o750`/`0o700`.
- [ ] `pprof`/debug : bind `127.0.0.1` uniquement (section 60).
- [ ] Binaire : `CGO_ENABLED=0`, build depuis un commit taggé, SBOM si exigé (`go list -m all`).
- [ ] Entrées externes : valider (taille, charset, chemins — pas de `../` vers `/etc`).

```go
// Limiter la taille d'un corps HTTP (anti-DoS mémoire)
r.Body = http.MaxBytesReader(w, r.Body, 1<<20) // 1 Mo max

// Empêcher la traversée de chemin sur un paramètre fichier
nom := r.PathValue("nom")
if strings.Contains(nom, "..") || strings.ContainsRune(nom, '/') {
	http.Error(w, "nom invalide", http.StatusBadRequest)
	return
}
chemin := filepath.Join("/srv/data", filepath.Clean(nom))
```

## 63. Bonnes pratiques & style (Effective Go condensé)

1. **Formatage non négociable** : `gofmt` (ou `goimports`). Zéro débat de style.
2. **Noms courts** en portée courte (`i`, `s`, `err`), explicites en portée large.
3. **Commentaires** : toute fonction/méthode exportée a un commentaire commençant par son nom.
4. **Erreurs** : traitées, enveloppées avec `%w`, messages en minuscules (section 29).
5. **Interfaces petites**, définies côté usage (section 24).
6. **Pas de variables globales mutables** partagées sans synchronisation.
7. **Acceptez des interfaces, retournez des structs** (concret en sortie, abstrait en entrée).
8. **`context.Context` en premier paramètre**, jamais dans un struct (section 42).
9. **Évitez `init()`** sauf enregistrement de plugins/drivers ; préférez des constructeurs explicites `New...`.
10. **Un package = une responsabilité** ; `internal/` pour ce qui ne doit pas fuir.
11. **Tests** : table-driven, `-race` en CI, couverture sur le code critique (pas de course au %).
12. **Logs structurés** (`slog`) plutôt que `fmt.Println` en prod.

Exemple de doc conforme :
```go
// Sonder teste la connectivité TCP vers hote:port avec le timeout donné.
// Il renvoie nil si la connexion réussit avant l'expiration du délai.
func Sonder(ctx context.Context, hote string, port int, timeout time.Duration) error {
	// ...
}
```

## 64. Erreurs classiques des débutants (12 à ne plus faire)

