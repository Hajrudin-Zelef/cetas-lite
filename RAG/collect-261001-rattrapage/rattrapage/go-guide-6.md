---
id: collect-261001-rattrapage/rattrapage/go-guide-6
title: "Guide Go (Golang) — du script sysadmin au service de production"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["arr"]
source: docs/RAG/collect-261001-rattrapage/go_guide.md
source_anchor: ""
source_lines: [1357, 1646]
sha256: 043d59a343246bcd6e56a6d9ea6522802156d668ab084e1168aab00ece80baa2
---

# Guide Go (Golang) — du script sysadmin au service de production

// --- Pool : recycler des objets (buffers) ---
var bufPool = sync.Pool{New: func() any { return make([]byte, 4096) }}
buf := bufPool.Get().([]byte)
defer bufPool.Put(buf)
```

> Depuis **Go 1.22**, la variable de boucle `for _, h := range` est **recréée à
> chaque itération** : `go func() { sonder(h) }()` est sûr. Avant 1.22, il fallait
> passer `h` en paramètre (comme ci-dessus) — gardez l'habitude, elle reste correcte.

Checklist concurrence :
- [ ] `go vet` ne signale pas de copie de `sync.Mutex` (ne jamais copier un Mutex !)
- [ ] Chaque `Lock` a son `Unlock` (defer de préférence)
- [ ] Pas d'accès concurrent à une map sans protection (`go run -race`, section 63)
- [ ] `WaitGroup.Add` **avant** `go`, `Done` en `defer` dans la goroutine

## 42. `context` : annulation et deadlines propres

```go
import "context"

// Timeout : toute l'opération (et ses sous-appels) meurt après 5s
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel() // LIBÈRE les ressources — ne jamais oublier

req, _ := http.NewRequestWithContext(ctx, "GET", url, nil)
resp, err := http.DefaultClient.Do(req)
if err != nil {
	// ctx.DeadlineExceeded ou context.Canceled
}

// Annulation manuelle
ctx2, cancel2 := context.WithCancel(context.Background())
go surveiller(ctx2)
// ... plus tard :
cancel2()

// Valeurs (avec parcimonie : request-scoped uniquement, jamais de secrets !)
ctx3 := context.WithValue(ctx, cleRequeteID, "abc-123")
```

Règles :
- `ctx` est **toujours le premier paramètre** : `func Faire(ctx context.Context, ...)`.
- Ne stockez jamais un `Context` dans un struct ; propagez-le en paramètre.
- `defer cancel()` même si vous pensez ne pas en avoir besoin (fuite de timer sinon).
- `context.Background()` dans `main`/`init`/tests ; `r.Context()` dans les handlers HTTP.
---

## 43. Pattern : worker pool (N workers, file de tâches)

Le classique du sysadmin : sonder 500 hôtes avec 20 connexions parallèles max.

```go
func workerPool(hotes []string, nbWorkers int) {
	taches := make(chan string)
	var wg sync.WaitGroup

	// Lancer les workers
	for i := 0; i < nbWorkers; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for h := range taches { // jusqu'à fermeture du channel
				if err := sonder(h); err != nil {
					fmt.Printf("worker %d: %s: %v\n", id, h, err)
				}
			}
		}(i)
	}

	// Distribuer
	for _, h := range hotes {
		taches <- h
	}
	close(taches) // signale la fin aux workers
	wg.Wait()
}
```

Variante avec résultats et erreurs via `errgroup` : voir section 46.

## 44. Pattern : fan-in / fan-out

- **Fan-out** : distribuer du travail à plusieurs goroutines (le worker pool ci-dessus).
- **Fan-in** : fusionner plusieurs channels en un seul.

```go
// Fan-in : multiplexe plusieurs sources vers un channel unique
func fusionner(cs ...<-chan Resultat) <-chan Resultat {
	sortie := make(chan Resultat)
	var wg sync.WaitGroup
	for _, c := range cs {
		wg.Add(1)
		go func(c <-chan Resultat) {
			defer wg.Done()
			for r := range c {
				sortie <- r
			}
		}(c)
	}
	go func() {
		wg.Wait()
		close(sortie) // fermer quand toutes les sources sont épuisées
	}()
	return sortie
}

// Usage : 3 sondes parallèles, résultats unifiés
c1 := sonderAsync("web-01")
c2 := sonderAsync("db-01")
c3 := sonderAsync("dns-01")
for r := range fusionner(c1, c2, c3) {
	fmt.Println(r.Hote, r.Ok)
}
```

## 45. Pattern : pipeline à étapes

```go
// Chaque étape : channel en entrée, channel en sortie, fermeture en cascade
func listerFichiers(racine string) <-chan string {
	sortie := make(chan string)
	go func() {
		defer close(sortie)
		filepath.WalkDir(racine, func(p string, d fs.DirEntry, err error) error {
			if err == nil && !d.IsDir() {
				sortie <- p
			}
			return nil
		})
	}()
	return sortie
}

func filtrerLogs(entree <-chan string) <-chan string {
	sortie := make(chan string)
	go func() {
		defer close(sortie)
		for p := range entree {
			if strings.HasSuffix(p, ".log") {
				sortie <- p
			}
		}
	}()
	return sortie
}

// Assemblage lisible
for p := range filtrerLogs(listerFichiers("/var/log")) {
	fmt.Println(p)
}
```

Ajoutez un `ctx` en paramètre de chaque étape pour une annulation propre
(`select` sur `ctx.Done()` dans les boucles d'envoi).

## 46. `errgroup` : goroutines + erreurs + annulation

`golang.org/x/sync/errgroup` — le chaînon manquant de la stdlib.

```go
import "golang.org/x/sync/errgroup"

func controlerTous(ctx context.Context, hotes []string) error {
	g, ctx := errgroup.WithContext(ctx)
	g.SetLimit(20) // max 20 goroutines simultanées (sémaphore intégré)

	for _, h := range hotes {
		h := h
		g.Go(func() error {
			// Si une goroutine échoue, ctx est annulé pour les autres
			return sonderAvecContexte(ctx, h)
		})
	}
	return g.Wait() // première erreur rencontrée (ou nil)
}
```

C'est le pattern à utiliser par défaut pour du travail parallèle avec gestion
d'erreur : plus lisible que WaitGroup + channel d'erreurs bricolé.

## 47. Serveur HTTP avec `net/http` (stdlib, ≥ 1.22)

Depuis Go 1.22, le routeur stdlib gère les **méthodes et paramètres** : fini
l'excuse du framework pour une API interne.

```go
package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func main() {
	mux := http.NewServeMux()

	// Routes avec méthode + paramètre {nom}
	mux.HandleFunc("GET /sante", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"statut": "ok"})
	})
	mux.HandleFunc("GET /hotes/{nom}", func(w http.ResponseWriter, r *http.Request) {
		nom := r.PathValue("nom") // Go 1.22+
		fmt.Fprintf(w, "hôte: %s", nom)
	})
	mux.HandleFunc("POST /hotes", creerHote)

	srv := &http.Server{
		Addr:         ":8080",
		Handler:      avecLogs(mux), // middleware
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Arrêt gracieux (voir section 55)
	go func() {
		log.Println("écoute sur :8080")
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()
	attendreSignal()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	srv.Shutdown(ctx)
}

// Middleware : log + request ID
func avecLogs(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		debut := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("%s %s %s", r.Method, r.URL.Path, time.Since(debut))
	})
}
```

## 48. Client HTTP : timeouts obligatoires

Le client par défaut **n'a pas de timeout** : une API pendue = goroutine pendue.

```go
client := &http.Client{
	Timeout: 10 * time.Second, // timeout TOTAL de la requête
}

// GET simple
resp, err := client.Get("http://10.0.0.53:8080/sante")
if err != nil {
	log.Fatal(err)
}
defer resp.Body.Close() // OBLIGATOIRE : sinon fuite de connexions

if resp.StatusCode != http.StatusOK {
	log.Fatalf("statut inattendu: %s", resp.Status)
}
corps, _ := io.ReadAll(resp.Body)

// POST JSON
payload, _ := json.Marshal(map[string]string{"nom": "web-02"})
req, _ := http.NewRequest("POST", url, bytes.NewReader(payload))
req.Header.Set("Content-Type", "application/json")
req.Header.Set("Authorization", "Bearer "+os.Getenv("API_TOKEN"))
resp, err = client.Do(req)
```

Réglages fins (pool de connexions, TLS) via `http.Transport` :

```go
tr := &http.Transport{
	MaxIdleConns:        50,
	MaxIdleConnsPerHost: 10,
	IdleConnTimeout:     90 * time.Second,
	TLSClientConfig:     &tls.Config{MinVersion: tls.VersionTLS12},
}
client := &http.Client{Timeout: 10 * time.Second, Transport: tr}
```

## 49. JSON sur HTTP : API complète minimale

```go
type Hote struct {
	Nom string `json:"nom"`
	IP  string `json:"ip"`
}

