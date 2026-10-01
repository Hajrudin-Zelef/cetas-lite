---
id: collect-261001-rattrapage/rattrapage/go-guide-10
title: "Guide Go (Golang) — du script sysadmin au service de production"
domain: rattrapage
role: reference
task: reference
actors: []
dates: ["2006-01-02"]
keywords: ["arr"]
source: docs/RAG/collect-261001-rattrapage/go_guide.md
source_anchor: ""
source_lines: [2451, 2655]
sha256: ac8fe596b9c8de1e4f7c1efb517197118e2c9f5b08d5e76eb4e38117d765543f
---

# Guide Go (Golang) — du script sysadmin au service de production

type Hote struct {
	Nom     string   `json:"nom"`
	IP      string   `json:"ip"`
	Roles   []string `json:"roles"`
	EnLigne bool     `json:"en_ligne"`
}

func main() {
	data, err := os.ReadFile("inventaire.json")
	if err != nil {
		log.Fatal(err)
	}
	var hotes []Hote
	if err := json.Unmarshal(data, &hotes); err != nil {
		log.Fatal(err)
	}
	sort.Slice(hotes, func(i, j int) bool { return hotes[i].Nom < hotes[j].Nom })

	parRole := map[string]int{}
	hs := 0
	for _, h := range hotes {
		if !h.EnLigne {
			hs++
		}
		for _, r := range h.Roles {
			parRole[r]++
		}
	}
	fmt.Printf("Hôtes: %d (hors-ligne: %d)\n", len(hotes), hs)
	fmt.Println("Par rôle:")
	roles := make([]string, 0, len(parRole))
	for r := range parRole {
		roles = append(roles, r)
	}
	sort.Strings(roles)
	for _, r := range roles {
		fmt.Printf("  %-10s %d\n", r, parRole[r])
	}
}
```

## 70. Déploiement : binaire statique, systemd, Docker

**Systemd** — `/etc/systemd/system/monoutil.service` :
```ini
[Unit]
Description=Mon outil Go
After=network.target

[Service]
Type=simple
User=monoutil
Environment=MONOUTIL_TIMEOUT=10s
ExecStart=/opt/monoutil/monoutil sonde web-01
Restart=on-failure
RestartSec=5

[Install]
WantedBy=multi-user.target
```
```bash
sudo systemctl daemon-reload && sudo systemctl enable --now monoutil
```

**Docker multi-stage** (binaire statique → image ~10 Mo) :
```dockerfile
FROM golang:1.23-alpine AS build
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o /monoutil .

FROM scratch
COPY --from=build /monoutil /monoutil
# Si HTTPS : COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
ENTRYPOINT ["/monoutil"]
```

Checklist mise en prod :
- [ ] Build reproductible : tag git → `-X main.version=...`
- [ ] `govulncheck ./...` propre
- [ ] Config par env/fichier (jamais recompilé par environnement)
- [ ] Logs JSON sur stdout, niveau réglable (`slog.Level`)
- [ ] Healthcheck (`/sante`) + arrêt gracieux SIGTERM (section 52)
- [ ] Utilisateur dédié, pas root ; droits fichiers minimaux

## 71. Pense-bête de poche

```
go run .                          exécuter
go build -ldflags="-s -w" -o app . compiler léger
go vet ./... && gofmt -l .         hygiène
go test -race -cover ./...         tests stricts
go mod tidy                        nettoyer les déps

var x int = 0      x := 0          déclaration (:= en fonction)
const Pi = 3.14                    constante
s := []int{1,2}                    slice
m := map[string]int{}              map (make ! jamais nil en écriture)
s = append(s, 3)                   TOUJOURS réaffecter
for i, v := range s                boucle
if err != nil { return ... }       gestion d'erreur
defer f.Close()                    nettoyage garanti
go f()                             goroutine
ch := make(chan int, 10)           channel bufferisé
select { case v := <-ch: ... }     multiplexage
ctx, cancel := context.WithTimeout  annulation
defer cancel()
wg.Add(1); go func(){defer wg.Done()}(); wg.Wait()

Serveur{IP: "x"}                   struct (champs nommés)
func (s *Serveur) M()              méthode (pointeur si mutation)
type I interface{ M() error }      interface (implicite)
v.(string) / switch v := v.(type)  assertion / type switch
fmt.Errorf("...: %w", err)         wrap d'erreur
errors.Is / errors.As             inspection
json.Marshal / json.Unmarshal     JSON (champs exportés !)
time.Parse("2006-01-02", s)       date de référence !
GOOS=windows go build             compilation croisée
```

## 72. Glossaire

| Terme | Sens |
|---|---|
| Binaire statique | Exécutable sans dépendance externe (libc incluse via `CGO_ENABLED=0`) |
| Channel | File typée de communication entre goroutines |
| Closure | Fonction capturant les variables de son environnement |
| `defer` | Appel reporté à la fin de la fonction (LIFO) |
| errgroup | Groupe de goroutines avec propagation d'erreur/annulation (`x/sync`) |
| go.mod / go.sum | Manifeste du module / empreintes des dépendances |
| GOROOT / GOPATH / GOCACHE | Install Go / espace de travail / cache de build |
| Goroutine | Fonction s'exécutant en parallèle, légère (~8 Ko) |
| Interface | Contrat de méthodes, implémentation implicite |
| `nil` | Valeur zéro des pointeurs, slices, maps, channels, interfaces |
| Pointeur | Adresse mémoire (`&v` / `*p`), pas d'arithmétique en Go |
| Race (data race) | Accès concurrents non synchronisés — détecté par `-race` |
| Receveur | Le `(s *Serveur)` devant une méthode |
| Rune | Point de code Unicode (alias `int32`) |
| Slice | Vue (ptr+len+cap) sur un tableau sous-jacent |
| Struct | Type composite à champs nommés |
| Vendoring | Copie des déps dans `vendor/` (`go mod vendor`) |
| Vet | Analyse statique de code suspect (`go vet`) |
| Wrap (erreur) | Enrichir en conservant la cause (`%w`, `Unwrap`) |

## 73. Quiz : 10 questions + réponses

**Q1.** Que vaut `fmt.Println(7/3)` et pourquoi ?
> `2` : division entière entre deux `int`, la partie décimale est tronquée.

**Q2.** Quelle est l'erreur dans `append(s, 42)` utilisé seul ?
> `append` peut réallouer : le résultat doit être réaffecté — `s = append(s, 42)`.

**Q3.** Pourquoi ce code panique-t-il : `var m map[string]int; m["a"] = 1` ?
> La map est `nil` : lecture ok, mais l'écriture sur une map nil panique. Initialiser avec `make` ou un littéral.

**Q4.** `defer` : dans quel ordre s'exécutent plusieurs appels différés ?
> LIFO : le dernier `defer` enregistré s'exécute en premier.

**Q5.** Un `main` qui lance `go traiter()` puis se termine immédiatement : que se passe-t-il ?
> La fin de `main` tue toutes les goroutines : `traiter` peut ne jamais s'exécuter. Synchroniser avec `WaitGroup`/`errgroup`.

**Q6.** Comment propager proprement une erreur en ajoutant du contexte ?
> `return fmt.Errorf("contexte: %w", err)` — `%w` conserve la chaîne pour `errors.Is`/`errors.As`.

**Q7.** Channel non bufferisé : que se passe-t-il à `ch <- v` sans récepteur ?
> L'envoi bloque jusqu'à ce qu'un récepteur soit prêt (rendez-vous).

**Q8.** À quoi sert `context.WithTimeout` + `defer cancel()` ?
> Borner la durée d'une opération et libérer les ressources associées (timer, goroutines).

**Q9.** Commande pour compiler pour Windows ARM64 depuis Linux ?
> `GOOS=windows GOARCH=arm64 go build -o app.exe .`

**Q10.** Pourquoi `resp.Body.Close()` est-il obligatoire après un `client.Do` ?
> Sinon la connexion n'est pas rendue au pool : fuite de sockets/fds, puis épuisement.

## 74. Pour aller plus loin

**Documentation officielle**
- https://go.dev/doc/ — tutoriels et guides
- https://pkg.go.dev/ — référence des paquets
- *Effective Go* (go.dev/doc/effective_go) et *Go Wiki*

**Livres**
- *The Go Programming Language* (Donovan & Kernighan) — la référence
- *Concurrency in Go* (Cox-Buday) — goroutines/channels en profondeur
- *Go in Action* / *100 Go Mistakes* (Teiva Harsanyi) — les pièges

**Outils à adopter**
- `golangci-lint` — méta-linter (vet + dizaines de règles)
- `govulncheck` — vulnérabilités des dépendances
- `delve` — debugueur, `go tool pprof` — profiling
- Cobra/Viper (CLI), `client_golang` (Prometheus), Testify (assertions)

**Feuille de route proposée (4 semaines, 30 min/jour)**
1. Semaine 1 : sections 1–21 — réécrire 2 scripts shell en Go
2. Semaine 2 : sections 22–42 — structs, erreurs, fichiers, JSON
3. Semaine 3 : sections 38–46 — concurrence : la sonde parallèle sur le parc réel
4. Semaine 4 : sections 47–56, 70 — API interne + binaire déployé via systemd

> Dernier conseil : Go récompense la simplicité explicite. Quand vous hésitez
> entre une abstraction élégante et du code direct, choisissez le code direct —
> c'est l'esprit du langage, et votre vous-dans-6-mois vous remerciera.
