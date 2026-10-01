---
id: collect-261001-rattrapage/rattrapage/go-guide-3
title: "Guide Go (Golang) — du script sysadmin au service de production"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["arr", "embedding"]
source: docs/RAG/collect-261001-rattrapage/go_guide.md
source_anchor: ""
source_lines: [513, 786]
sha256: 25629b106a9a6d7bc1e795e4ecd002e06a2580f7a1dbada0f803905acdfb3699
---

# Guide Go (Golang) — du script sysadmin au service de production

Points critiques :
- **Les maps ne sont pas thread-safe** : accès concurrents = `sync.Map` ou mutex (section 41).
- L'ordre d'itération est volontairement aléatoire (ne jamais en dépendre).
- Une map `nil` se lit sans paniquer (`v, ok := m[k]` → zéro + false) mais
  l'écriture panique : initialisez avec `make` ou un littéral.
- Clés : tout type **comparable** (`string`, `int`, struct sans slice...). Pas de slice en clé.

Exemple sysadmin : compter des occurrences (lignes de log par code) :
```go
codes := map[string]int{}
for _, ligne := range lignes {
	code := extraireCode(ligne)
	codes[code]++ // zéro implicite si absent : pas besoin de tester !
}
```

## 17. Chaînes, runes et bytes : l'UTF-8 sans douleur

```go
s := "héllo"                  // string = suite d'OCTETS, immuable, UTF-8 par convention
fmt.Println(len(s))           // 6 octets (é = 2 octets), pas 5 caractères !

for i, r := range s {         // range décode l'UTF-8 : r est une rune
	fmt.Printf("pos octet %d : %c (U+%04X)\n", i, r, r)
}

runes := []rune(s)            // conversion en runes pour compter les caractères
fmt.Println(len(runes))       // 5

// Concaténation en boucle : strings.Builder (pas += répété)
var b strings.Builder
for _, h := range hotes {
	b.WriteString(h)
	b.WriteByte('\n')
}
texte := b.String()

// Manipulations courantes
strings.Contains(s, "llo")
strings.HasPrefix(s, "hé")
strings.Split("a,b,c", ",")   // []string{"a","b","c"}
strings.Join([]string{"a","b"}, ",")
strings.TrimSpace("  x  ")
strings.ToUpper(s)
strings.ReplaceAll(s, "l", "L")
strconv.Itoa(42)              // int -> string
n, err := strconv.Atoi("42")  // string -> int
```

Règles d'or :
- Une `string` est immuable : toute "modification" crée une nouvelle chaîne.
- `len(s)` = octets. Pour des caractères : `len([]rune(s))` ou `utf8.RuneCountInString(s).
- En boucle intensive, `strings.Builder` bat la concaténation `+=` (allocations).

## 18. Pointeurs : juste ce qu'il faut

```go
x := 42
p := &x          // p pointe vers x
fmt.Println(*p)  // 42 : déréférencement
*p = 100
fmt.Println(x)   // 100

// Pointeur nil : le déréférencer = panic
var q *int
fmt.Println(q == nil) // true

// Passage par pointeur pour modifier ou éviter une grosse copie
func reset(c *Compteur) { c.valeur = 0 }
```

Ce que vous devez savoir, pas plus :
- `&v` prend l'adresse, `*p` accède à la valeur pointée.
- Go n'a **pas** d'arithmétique de pointeurs (contrairement à C).
- Les slices et maps sont déjà des références : pas besoin de `*[]int` en paramètre.
- Receveurs de méthodes : pointeur `*T` si la méthode modifie le struct (section 26).

## 19. Fonctions : valeurs de retour multiples, nommées

```go
// Deux valeurs de retour : l'idiome (résultat, erreur)
func diviser(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division par zéro")
	}
	return a / b, nil
}

// Retours nommés (utiles avec defer, voir section 22)
func ping(hote string) (latence time.Duration, err error) {
	latence = mesurer(hote) // assigné aux variables nommées
	return                 // "return nu" : renvoie latence, err
}

// Nombre variable d'arguments
func somme(nombres ...int) int {
	total := 0
	for _, n := range nombres {
		total += n
	}
	return total
}
somme(1, 2, 3)
somme(slice...) // déplier un slice

// Fonction anonyme + closure (capture les variables environnantes)
compteur := func() func() int {
	n := 0
	return func() int {
		n++
		return n
	}
}()
fmt.Println(compteur(), compteur()) // 1 2
```

## 20. `defer` : le nettoyage garanti

`func` différée : exécutée à la **fin** de la fonction appelante, en LIFO.

```go
func traiter(fichier string) error {
	f, err := os.Open(fichier)
	if err != nil {
		return err
	}
	defer f.Close() // garanti, même en cas de return anticipé ou panic

	// ... lecture ...
	return nil
}

// Ordre LIFO
func demo() {
	defer fmt.Println("1er différé, exécuté en dernier")
	defer fmt.Println("2e différé, exécuté en premier")
	fmt.Println("corps")
}
// Affiche : corps / 2e différé... / 1er différé...
```

Usages sysadmin typiques : `defer f.Close()`, `defer mu.Unlock()`,
`defer resp.Body.Close()` (HTTP client — **obligatoire** pour réutiliser les connexions).

> Astuce : les arguments de `defer` sont évalués **immédiatement**.
> `defer fmt.Println(i)` affiche la valeur de `i` au moment du `defer`, pas à la fin.

## 21. `panic` et `recover` : l'exception qui n'en est pas une

```go
func critiquer(n int) {
	if n < 0 {
		panic("négatif interdit") // arrête la goroutine (et le programme si non récupéré)
	}
}

func securise() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("récupéré :", r)
		}
	}()
	critiquer(-1)
	fmt.Println("jamais atteint")
}
```

Doctrine Go :
- `panic` = bug du programmeur ou état irrécupérable (index hors limites, map nil en écriture).
- **Jamais** pour une erreur d'exécution normale (fichier absent, réseau coupé) → `error`.
- `recover` ne sert que dans un `defer`, typiquement aux frontières (serveur HTTP,
  goroutine worker) pour éviter qu'un panic ne tue tout le processus.
- La bibliothèque standard panique rarement ; votre code applicatif ne devrait
  quasiment jamais appeler `panic` directement.
---

## 22. Structs : vos objets métier

```go
type Serveur struct {
	Nom     string   // champ EXPORTÉ (majuscule) : visible hors package
	IP      string
	Port    int
	Tags    []string
	enLigne bool     // champ non exporté : privé au package
}

// Littéraux
s1 := Serveur{Nom: "dns-01", IP: "10.0.0.53", Port: 53} // nommé : recommandé
s2 := Serveur{"dns-01", "10.0.0.53", 53, nil, true}      // positionnel : fragile, à éviter

// Champ par champ
var s3 Serveur
s3.Nom = "ntp-01"

// Struct anonyme (config ponctuelle)
cfg := struct {
	Timeout int
	Retry   int
}{Timeout: 5, Retry: 3}
```

## 23. Méthodes : fonctions attachées à un type

```go
// Receveur valeur : ne modifie pas l'original (ok pour petits structs / lecture)
func (s Serveur) Adresse() string {
	return fmt.Sprintf("%s:%d", s.IP, s.Port)
}

// Receveur pointeur : modifie l'original, évite la copie
func (s *Serveur) MarquerEnLigne() {
	s.enLigne = true
}

s := Serveur{Nom: "dns-01", IP: "10.0.0.53", Port: 53}
fmt.Println(s.Adresse())
s.MarquerEnLigne() // Go prend &s automatiquement
```

Règle de cohérence : si **une** méthode d'un type utilise un receveur pointeur,
utilisez le pointeur pour **toutes** les méthodes du type (recommandation `go vet`
implicite via les linters).

## 24. Interfaces : le polymorphisme à la Go

Petites interfaces, définies **côté consommateur**. L'idiome : `io.Reader`, `io.Writer`.

```go
// Une interface = un contrat de méthodes
type Verifiable interface {
	Verifier() error
}

// Serveur IMPLÉMENTE Verifiable implicitement (pas de "implements")
func (s Serveur) Verifier() error {
	if s.IP == "" {
		return errors.New("IP manquante")
	}
	return nil
}

func controler(v Verifiable) error {
	return v.Verifier() // marche avec TOUT type ayant Verifier() error
}
```

Points clés :
- Implémentation **implicite** : aucun mot-clé, le compilateur vérifie.
- Préférez les petites interfaces (1-2 méthodes). `interface{}` / `any` = absence de contrat.
- `any` est un alias de `interface{}` (Go 1.18+). À utiliser avec parcimonie.
- L'interface vide accepte tout : utile pour le JSON générique, dangereuse partout ailleurs.

## 25. Composition (embedding) : l'alternative à l'héritage

Go n'a pas d'héritage. On compose :

```go
type Base struct {
	Nom string
}

func (b Base) Decrire() string { return "hôte: " + b.Nom }

type ServeurWeb struct {
	Base          // embarqué : ses champs/méthodes sont "promus"
	Port    int
}

w := ServeurWeb{Base: Base{Nom: "web-01"}, Port: 8080}
fmt.Println(w.Nom)      // promu : w.Base.Nom
fmt.Println(w.Decrire()) // méthode promue
```

