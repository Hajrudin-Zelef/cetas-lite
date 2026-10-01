---
id: collect-261001-rattrapage/rattrapage/go-guide-9
title: "Guide Go (Golang) — du script sysadmin au service de production"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["agent"]
source: docs/RAG/collect-261001-rattrapage/go_guide.md
source_anchor: ""
source_lines: [2142, 2450]
sha256: c8cc00c7d0864fef6cea377af0c19b3f451715082aaa399be795bbfc31a872e4
---

# Guide Go (Golang) — du script sysadmin au service de production

| # | Erreur | Correct |
|---|---|---|
| 1 | Oublier de réaffecter `append` : `append(s, v)` seul | `s = append(s, v)` |
| 2 | Écrire dans une map `nil` | `m := make(map[string]int)` |
| 3 | Ignorer `err` : `f, _ := os.Open(p)` | Toujours tester `err` |
| 4 | Oublier `defer resp.Body.Close()` | Le mettre juste après l'erreur testée |
| 5 | Copier un `sync.Mutex` (struct contenant un mutex passé par valeur) | Receveurs/paramètres pointeurs |
| 6 | Variable de boucle capturée (avant Go 1.22) | Passer en paramètre de la closure |
| 7 | `go` routine sans synchronisation, `main` qui quitte | `sync.WaitGroup` / `errgroup` |
| 8 | Comparer avec `==` des slices ou attendre un ordre de map | `slices.Equal`, trier les clés |
| 9 | Division entière surprise : `7/3 == 2` | Convertir en `float64` d'abord |
| 10 | `panic` pour une erreur normale | Retourner `error` |
| 11 | `select` sans `default`/`timeout` qui bloque pour toujours | Toujours un cas de sortie |
| 12 | Secrets/URLs en dur, `InsecureSkipVerify: true` | Env, fichiers, config (section 62) |

Bonus — le piège du slice partagé :
```go
a := []int{1, 2, 3, 4}
b := a[1:3]  // b partage le tableau de a
b[0] = 99
fmt.Println(a) // [1 99 3 4] : surprise si on l'oublie !
```

## 65. Cas pratique 1 — Sonde TCP parallèle (script sysadmin)

L'outil "premier jour" : 200 hôtes, 20 workers, timeout, sortie triée.

```go
// sonde/main.go
package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"os"
	"sort"
	"sync"
	"time"

	"golang.org/x/sync/errgroup"
)

type Resultat struct {
	Hote string
	Ok   bool
	Ms   int64
	Err  string
}

func sonder(ctx context.Context, hote string, port int, timeout time.Duration) Resultat {
	debut := time.Now()
	d := net.Dialer{Timeout: timeout}
	conn, err := d.DialContext(ctx, "tcp", fmt.Sprintf("%s:%d", hote, port))
	r := Resultat{Hote: hote, Ms: time.Since(debut).Milliseconds()}
	if err != nil {
		r.Err = err.Error()
		return r
	}
	conn.Close()
	r.Ok = true
	return r
}

func main() {
	port := flag.Int("port", 22, "port TCP")
	workers := flag.Int("w", 20, "parallélisme")
	timeout := flag.Duration("timeout", 3*time.Second, "timeout par hôte")
	flag.Parse()
	hotes := flag.Args()
	if len(hotes) == 0 {
		fmt.Fprintln(os.Stderr, "usage: sonde [-port=22] [-w=20] hote1 hote2 ...")
		os.Exit(2)
	}

	ctx := context.Background()
	g, ctx := errgroup.WithContext(ctx)
	g.SetLimit(*workers)
	resultats := make([]Resultat, len(hotes))
	for i, h := range hotes {
		i, h := i, h
		g.Go(func() error {
			resultats[i] = sonder(ctx, h, *port, *timeout)
			return nil // on ne veut pas annuler les autres sur un échec
		})
	}
	_ = g.Wait()

	sort.Slice(resultats, func(i, j int) bool { return resultats[i].Hote < resultats[j].Hote })
	ko := 0
	for _, r := range resultats {
		statut := "OK"
		if !r.Ok {
			statut = "KO"
			ko++
		}
		fmt.Printf("%-20s %-3s %4d ms %s\n", r.Hote, statut, r.Ms, r.Err)
	}
	fmt.Printf("--- %d/%d joignables ---\n", len(resultats)-ko, len(resultats))
	if ko > 0 {
		os.Exit(1) // code de sortie exploitable par Nagios/Zabbix/scripts
	}
}
```

```bash
go mod init sonde && go get golang.org/x/sync@latest
go build -o sonde .
./sonde -port=443 -w=30 web-01 web-02 db-01
echo $?  # 0 si tout OK, 1 sinon
```

## 66. Cas pratique 2 — Rotation de logs / ménage disque

```go
// menage/main.go : supprime les .log de plus de N jours, avec dry-run
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

func main() {
	racine := flag.String("dir", "/var/log/monapp", "dossier à nettoyer")
	jours := flag.Int("jours", 30, "âge max en jours")
	dryRun := flag.Bool("dry-run", true, "ne rien supprimer (défaut: true)")
	flag.Parse()

	limite := time.Now().AddDate(0, 0, -*jours)
	var supprimes, conserves int64
	var octets int64

	err := filepath.WalkDir(*racine, func(p string, d os.DirEntry, err error) error {
		if err != nil {
			log.Printf("accès %s: %v", p, err)
			return nil // continuer malgré tout
		}
		if d.IsDir() || filepath.Ext(p) != ".log" {
			return nil
		}
		info, err := d.Info()
		if err != nil {
			return nil
		}
		if info.ModTime().Before(limite) {
			octets += info.Size()
			if *dryRun {
				fmt.Println("[dry-run] supprimerait:", p)
			} else {
				if err := os.Remove(p); err != nil {
					log.Printf("suppression %s: %v", p, err)
				} else {
					fmt.Println("supprimé:", p)
				}
			}
			supprimes++
		} else {
			conserves++
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("supprimés: %d, conservés: %d, %.1f Mo libérés\n",
		supprimes, conserves, float64(octets)/1e6)
}
```

```bash
./menage -dir=/var/log -jours=90            # dry-run par défaut : sûr
./menage -dir=/var/log -jours=90 -dry-run=false  # action réelle
```

> Le `dry-run` à `true` par défaut est un garde-fou : tout script destructeur
> devrait exiger un flag explicite pour agir.

## 67. Cas pratique 3 — Exporter Prometheus (agent de supervision)

Un exporter minimal : Go excelle ici (binaire unique, faible empreinte).

```go
// exporter/main.go
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"
	"time"
)

var demarrage = time.Now()

func metriques(w http.ResponseWriter, r *http.Request) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Fprintf(w, `# HELP uptime_secondes Durée de fonctionnement
# TYPE uptime_secondes counter
uptime_secondes %d
# HELP goroutines Nombre de goroutines
# TYPE goroutines gauge
goroutines %d
# HELP memoire_octets Mémoire allouée
# TYPE memoire_octets gauge
memoire_octets %d
`, int64(time.Since(demarrage).Seconds()), runtime.NumGoroutine(), m.Alloc)
}

func main() {
	http.HandleFunc("GET /metrics", metriques)
	http.HandleFunc("GET /sante", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})
	addr := ":9100"
	log.Println("exporter sur", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err)
	}
}
```

```bash
go build -o exporter . && ./exporter &
curl localhost:9100/metrics
```

Pour un vrai exporter : `github.com/prometheus/client_golang` (compteurs,
histogrammes, labels). Le format texte ci-dessus suffit pour un besoin ponctuel
et se scrape tel quel par Prometheus.

## 68. Cas pratique 4 — Webhook de déploiement / redémarrage

```go
// webhook/main.go : reçoit un POST signé, exécute une action (ex: reload service)
package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"time"
)

func verifierSignature(secret, corps []byte, sigHex string) bool {
	mac := hmac.New(sha256.New, secret)
	mac.Write(corps)
	attendu := hex.EncodeToString(mac.Sum(nil))
	return hmac.Equal([]byte(attendu), []byte(sigHex))
}

func main() {
	secret := os.Getenv("WEBHOOK_SECRET")
	if secret == "" {
		log.Fatal("WEBHOOK_SECRET requis")
	}
	http.HandleFunc("POST /redeploy", func(w http.ResponseWriter, r *http.Request) {
		r.Body = http.MaxBytesReader(w, r.Body, 1<<20)
		corps, _ := io.ReadAll(r.Body)
		if !verifierSignature([]byte(secret), corps, r.Header.Get("X-Signature")) {
			http.Error(w, "signature invalide", http.StatusUnauthorized)
			return
		}
		// Exécution asynchrone : répondre vite, agir ensuite
		go func() {
			cmd := exec.Command("systemctl", "reload", "monapp")
			if out, err := cmd.CombinedOutput(); err != nil {
				log.Printf("reload: %v: %s", err, out)
			} else {
				log.Println("service rechargé")
			}
		}()
		w.WriteHeader(http.StatusAccepted)
	})
	srv := &http.Server{Addr: "127.0.0.1:8080", ReadTimeout: 5 * time.Second}
	log.Fatal(srv.ListenAndServe())
}
```

Sécurité du webhook : signature HMAC, bind local + reverse proxy, taille limitée,
pas d'args issus du corps dans `exec` (commande figée ici).

## 69. Cas pratique 5 — Inventaire JSON → rapport

```go
// rapport/main.go : lit inventaire.json, sort un rapport texte trié
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sort"
)

