---
id: collect-261001-rattrapage/rattrapage/go-guide-5
title: "Guide Go (Golang) — du script sysadmin au service de production"
domain: rattrapage
role: reference
task: reference
actors: []
dates: ["2006-01-02", "2026-09-26"]
keywords: ["attention", "decode"]
source: docs/RAG/collect-261001-rattrapage/go_guide.md
source_anchor: ""
source_lines: [1060, 1356]
sha256: 4c393306880f951f216e696e231f276c3713d990f67a2573ebc50177ec0331a2
---

# Guide Go (Golang) — du script sysadmin au service de production

```go
filepath.Join("opt", "monoutil", "data") // "opt/monoutil/data" (séparateur OS)
filepath.Dir("/a/b/c.txt")              // "/a/b"
filepath.Base("/a/b/c.txt")             // "c.txt"
filepath.Ext("archive.tar.gz")          // ".gz"
filepath.Abs("relatif")                 // chemin absolu

// Lister avec motif
fichiers, _ := filepath.Glob("/etc/*.conf")

// Parcours récursif
err := filepath.WalkDir("/opt/data", func(chemin string, d fs.DirEntry, err error) error {
	if err != nil {
		return err
	}
	if !d.IsDir() && strings.HasSuffix(chemin, ".log") {
		fmt.Println(chemin)
	}
	return nil
})
```

## 35. `encoding/json` : le format d'échange sysadmin

```go
type Hote struct {
	Nom     string   `json:"nom"`
	IP      string   `json:"ip"`
	Tags    []string `json:"tags,omitempty"` // omitempty : omis si vide
	Secret  string   `json:"-"`              // jamais sérialisé
	EnLigne bool     `json:"en_ligne"`
}

// Struct -> JSON
h := Hote{Nom: "web-01", IP: "10.0.0.80", EnLigne: true}
data, err := json.MarshalIndent(h, "", "  ")
fmt.Println(string(data))

// JSON -> struct (champs non exportés ignorés silencieusement !)
var h2 Hote
if err := json.Unmarshal(data, &h2); err != nil {
	log.Fatal(err)
}

// Décoder un flux (fichier, réponse HTTP)
f, _ := os.Open("hotes.json")
defer f.Close()
var liste []Hote
dec := json.NewDecoder(f)
if err := dec.Decode(&liste); err != nil {
	log.Fatal(err)
}

// JSON générique quand la structure est inconnue
var gen map[string]any
json.Unmarshal([]byte(`{"a":1,"b":"x"}`), &gen)
fmt.Println(gen["a"].(float64)) // les nombres deviennent float64 !
```

Pièges JSON :
- Champs **non exportés** (minuscule) = ignorés silencieusement dans les deux sens.
- Nombres → `float64` en décodage générique : pour de gros entiers, utilisez
  `json.Number` ou décodez vers une struct typée.
- Tags `json:"-"` pour exclure les secrets des dumps (section 66).

## 36. CSV, YAML, TOML : les autres formats de l'inventaire

```go
// CSV : encoding/csv (stdlib)
f, _ := os.Open("inventaire.csv")
r := csv.NewReader(f)
r.Comma = ';' // CSV "français"
lignes, err := r.ReadAll()

w := csv.NewWriter(os.Stdout)
w.Write([]string{"nom", "ip"})
w.Write([]string{"web-01", "10.0.0.80"})
w.Flush()

// YAML : pas en stdlib — gopkg.in/yaml.v3 (go get)
// TOML : pas en stdlib — github.com/pelletier/go-toml/v2 ou BurntSushi/toml
```

```yaml
# Exemple inventaire.yaml (lu avec yaml.v3)
hotes:
  - nom: web-01
    ip: 10.0.0.80
    roles: [web, prod]
```

```go
// Lecture YAML (après go get gopkg.in/yaml.v3)
import "gopkg.in/yaml.v3"

type Inventaire struct {
	Hotes []Hote `yaml:"hotes"`
}
data, _ := os.ReadFile("inventaire.yaml")
var inv Inventaire
yaml.Unmarshal(data, &inv) // tags yaml: si les noms diffèrent
```

Choix de format : JSON pour les API, YAML pour la config humaine, TOML pour la
config simple, CSV pour l'échange tableur. Viper (section 52) abstrait tout ça.

## 37. Le temps : `time`, durées, timeouts

```go
maintenant := time.Now()
fmt.Println(maintenant.Format("2006-01-02 15:04:05")) // format = date de référence !
fmt.Println(maintenant.Format(time.RFC3339))          // 2026-09-26T23:34:27Z

// La date de référence Go : 02/01/2006 15:04:05 (1,2,3,4,5,6,7)
t, err := time.Parse("2006-01-02", "2026-09-26")

// Durées
d, _ := time.ParseDuration("1h30m")
time.Sleep(500 * time.Millisecond)
deadline := time.Now().Add(5 * time.Second)

// Mesurer
debut := time.Now()
traiter()
fmt.Println("durée:", time.Since(debut))

// Ticker : action périodique (supervision)
ticker := time.NewTicker(60 * time.Second)
defer ticker.Stop()
for {
	select {
	case t := <-ticker.C:
		collecterMetriques(t)
	case <-ctx.Done():
		return
	}
}

// Fuseaux
paris, _ := time.LoadLocation("Europe/Paris")
fmt.Println(maintenant.In(paris).Format("15:04 MST"))
```

> Le layout `2006-01-02 15:04:05` surprend : ce n'est pas un motif `YYYY-MM-DD`,
> c'est **la date de référence** `Mon Jan 2 15:04:05 MST 2006` (1 2 3 4 5 6 7).
> On écrit l'heure qu'il est "le 2 janvier 2006 à 15h04:05".

## 38. Goroutines : la concurrence en un mot-clé

```go
func sonder(hote string) {
	// ... ping / tcp dial ...
}

func main() {
	hotes := []string{"web-01", "db-01", "dns-01"}

	// Lancer une goroutine : préfixer par 'go'
	for _, h := range hotes {
		go sonder(h) // s'exécute en parallèle, ne bloque pas
	}

	time.Sleep(2 * time.Second) // ATTENTION : bricolage, voir WaitGroup section 41
	fmt.Println("fini (peut-être)")
}
```

Réalités :
- Une goroutine coûte ~8 Ko de pile initiale (vs ~1 Mo pour un thread OS).
  On en lance des **milliers** sans sourciller.
- `go f()` ne retourne rien et ne propage pas d'erreur : communiquez via channels (39).
- Le `main` qui se termine **tue** toutes les goroutines : synchronisez (WaitGroup).
- Ordre d'exécution non garanti : ne jamais supposer un ordre entre goroutines.

## 39. Channels : le tuyau typé entre goroutines

```go
// Channel non bufferisé : l'envoi bloque jusqu'à réception (rendez-vous)
ch := make(chan string)
go func() {
	ch <- "résultat" // bloque jusqu'à ce que main reçoive
}()
msg := <-ch // reçoit
fmt.Println(msg)

// Channel bufferisé : l'envoi bloque seulement si plein
file := make(chan int, 10)
file <- 1 // ne bloque pas (place libre)

// Fermer : signale "plus de données" (seul l'ÉMETTEUR ferme)
go func() {
	for i := 0; i < 3; i++ {
		ch2 <- i
	}
	close(ch2)
}()
for v := range ch2 { // boucle jusqu'à fermeture
	fmt.Println(v)
}

// Tester la fermeture
v, ok := <-ch2 // ok == false si fermé ET vide
```

| Type | Création | Sémantique |
|---|---|---|
| Non bufferisé | `make(chan T)` | Rendez-vous : envoi ↔ réception synchrones |
| Bufferisé | `make(chan T, n)` | File de n places, découplage partiel |
| Lecture seule / écriture seule | `chan<- T`, `<-chan T` | Dans les signatures : exprime l'intention |

Règles : on ne ferme que côté émetteur ; envoyer sur un channel fermé = panic ;
recevoir sur un channel `nil` bloque **pour toujours** (utile dans `select`, section 40).

## 40. `select` : l'aiguillage des channels

```go
select {
case msg := <-ch1:
	fmt.Println("ch1:", msg)
case ch2 <- "ping":
	fmt.Println("envoyé sur ch2")
case <-time.After(2 * time.Second): // timeout !
	fmt.Println("timeout")
default: // optionnel : ne bloque pas si aucun cas prêt
	fmt.Println("rien de prêt")
}
```

Motifs indispensables :

```go
// 1. Timeout sur opération
select {
case res := <-resultat:
	traiter(res)
case <-time.After(5 * time.Second):
	return errors.New("timeout après 5s")
}

// 2. Annulation via context (voir section 42)
select {
case <-ctx.Done():
	return ctx.Err()
case v := <-ch:
	fmt.Println(v)
}

// 3. Non-bloquant : envoyer seulement si un lecteur attend
select {
case ch <- v:
default:
	// file pleine / personne n'écoute : on jette ou on compte
}
```

`select` choisit **au hasard** si plusieurs cas sont prêts simultanément :
ne jamais en dépendre pour une priorité (ou ordonner les tests manuellement).

## 41. `sync` : Mutex, WaitGroup, Once

```go
import "sync"

// --- WaitGroup : attendre des goroutines ---
var wg sync.WaitGroup
for _, h := range hotes {
	wg.Add(1)
	go func(hote string) {
		defer wg.Done() // ou wg.Add(-1)
		sonder(hote)
	}(h) // passer h en paramètre : évite la capture de la variable de boucle
}
wg.Wait() // bloque jusqu'à ce que le compteur retombe à 0

// --- Mutex : protéger une ressource partagée ---
var mu sync.Mutex
compteurs := map[string]int{}
go func() {
	mu.Lock()
	compteurs["ok"]++
	mu.Unlock() // ou defer mu.Unlock() juste après Lock
}()

// RWMutex : verrou lecture/écriture (lectures concurrentes ok)
var rwmu sync.RWMutex
rwmu.RLock(); v := cache[k]; rwmu.RUnlock()

// --- Once : initialisation unique thread-safe ---
var (
	cfg     Config
	cfgOnce sync.Once
)
func getConfig() Config {
	cfgOnce.Do(func() { cfg = chargerConfig() })
	return cfg
}

