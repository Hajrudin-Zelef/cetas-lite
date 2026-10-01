---
id: collect-261001-rattrapage/rattrapage/go-guide-4
title: "Guide Go (Golang) — du script sysadmin au service de production"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/go_guide.md
source_anchor: ""
source_lines: [787, 1059]
sha256: b543ba2b9ec46779aad9710ce83d47ce1e902943f707b49c73477fe77cbb9d6a
---

# Guide Go (Golang) — du script sysadmin au service de production

C'est de la **délégation**, pas de l'héritage : pas de polymorphisme implicite
ni de redéfinition virtuelle. Pour le polymorphisme, voir les interfaces (24).

## 26. Assertions de type et `type switch`

```go
var v any = "hello"

// Assertion avec test (ne panique pas)
s, ok := v.(string)
if ok {
	fmt.Println("c'est une string:", s)
}

// Assertion directe : panique si le type ne correspond pas (à éviter hors certitude)
s2 := v.(string)

// Type switch : aiguillage propre
switch t := v.(type) {
case string:
	fmt.Println("string de longueur", len(t))
case int:
	fmt.Println("int:", t)
case nil:
	fmt.Println("nil")
default:
	fmt.Printf("type inattendu %T\n", t)
}
```

## 27. Génériques (Go 1.18+) : quand s'en servir

```go
// Fonction générique : contrainte via ~
func Min[T int | int64 | float64](a, b T) T {
	if a < b {
		return a
	}
	return b
}

// Contraintes du paquet cmp / slices (~1.21)
import "cmp"
func Max2[T cmp.Ordered](a, b T) T { ... }

// Type générique
type Pile[T any] struct {
	elements []T
}

func (p *Pile[T]) Push(v T) { p.elements = append(p.elements, v) }
func (p *Pile[T]) Pop() (T, bool) {
	if len(p.elements) == 0 {
		var zero T
		return zero, false
	}
	v := p.elements[len(p.elements)-1]
	p.elements = p.elements[:len(p.elements)-1]
	return v, true
}
```

Doctrine : les génériques servent aux **conteneurs et algorithmes** (slices, maps,
sets, files). Pour la logique métier, les interfaces restent l'outil principal.
N'abusez pas : une fonction générique là où `any` + assertion suffisait complique la lecture.

## 28. Gestion d'erreurs explicite : la philosophie Go

Pas d'exceptions. Une fonction qui peut échouer renvoie `error` en **dernière** position.

```go
contenu, err := os.ReadFile("/etc/hostname")
if err != nil {
	// Traiter : logguer, enrichir, propager, ou abandonner — mais JAMAIS ignorer silencieusement
	return fmt.Errorf("lecture hostname: %w", err)
}
fmt.Println(string(contenu))
```

Le motif `if err != nil` représente une part visible du code Go : c'est normal,
c'est documenté, c'est le prix de l'explicite. Les erreurs sont des **valeurs**,
on les inspecte, on les enveloppe, on les compare.

| Réflexe | Exemple |
|---|---|
| Propager en enrichissant | `return fmt.Errorf("connexion %s: %w", hote, err)` |
| Valeur sentinelle | `var ErrTimeout = errors.New("timeout")` puis `errors.Is(err, ErrTimeout)` |
| Erreur typée | `var e *MonErreur; errors.As(err, &e)` |
| Abandonner (main, init) | `log.Fatalf("...: %v", err)` |

## 29. Créer, envelopper, inspecter les erreurs

```go
import (
	"errors"
	"fmt"
)

// 1. Erreur simple
err := errors.New("disque plein")

// 2. Erreur formatée
err = fmt.Errorf("écriture %s: %d octets manquants", chemin, n)

// 3. Envelopper (wrap) avec %w : conserve la chaîne pour errors.Is/As
if err != nil {
	return fmt.Errorf("sauvegarde journalière: %w", err)
}

// 4. Sentinelle + test
var ErrFichierAbsent = errors.New("fichier absent")

func charger(p string) error {
	if _, err := os.Stat(p); err != nil {
		return fmt.Errorf("%w: %s", ErrFichierAbsent, p)
	}
	return nil
}

if errors.Is(charger("/x"), ErrFichierAbsent) {
	fmt.Println("créer le fichier par défaut")
}

// 5. Erreur personnalisée avec contexte
type ErreurReseau struct {
	Hote string
	Op   string
	Err  error
}

func (e *ErreurReseau) Error() string {
	return fmt.Sprintf("%s %s: %v", e.Op, e.Hote, e.Err)
}
func (e *ErreurReseau) Unwrap() error { return e.Err } // dépliage pour Is/As

// 6. Extraire un type précis
var er *ErreurReseau
if errors.As(err, &er) {
	fmt.Println("hôte fautif:", er.Hote)
}
```

Checklist erreurs :
- [ ] Chaque `err != nil` est traité (jamais `_ =` silencieux hors cas justifié et commenté)
- [ ] `%w` (pas `%v`) quand on veut garder la chaîne d'inspection
- [ ] Messages en minuscules, sans ponctuation finale (convention)
- [ ] Pas de `panic` pour une erreur d'exécution (section 21)

## 30. Packages : organisation et visibilité

```
monoutil/
├── go.mod
├── main.go            # package main : le binaire
├── cmd/
│   └── monoutil/
│       └── main.go    # variante : main dans cmd/ (convention CLI)
├── internal/
│   └── inventaire/    # importable UNIQUEMENT par ce module
│       └── inventaire.go  # package inventaire
└── pkg/               # (convention) code réutilisable hors module
    └── reseau/
        └── ping.go    # package reseau
```

Règles :
- **1 dossier = 1 package**. Le nom du package = nom du dossier (sauf `main`).
- Exporté = commence par une **majuscule** (`Serveur`, `Verifier`). Non exporté = minuscule.
- `internal/` : le compilateur **interdit** l'import depuis l'extérieur du module parent.
- Évitez les imports cycliques : le compilateur les refuse. Restructurez (extrayez un package commun).
- Un fichier peut importer un package sous alias : `import inv "infra/outils/internal/inventaire"`.
- Import pour effet de bord seul : `import _ "net/http/pprof"` (init du package).
---

## 31. Entrées-sorties console : `fmt` au quotidien

```go
fmt.Println("texte", 42, true)          // espaces auto, \n final
fmt.Printf("hôte=%s ip=%s charge=%.1f%%\n", nom, ip, charge)
fmt.Sprintf("...")                       // renvoie la string au lieu d'afficher
fmt.Fprintf(os.Stderr, "erreur: %v\n", err) // vers un writer quelconque

// Verbes essentiels
// %v valeur brute | %+v struct avec noms de champs | %#v syntaxe Go
// %T type | %d entier | %s string | %q string quotée | %x hexadécimal
// %f flottant | %.2f 2 décimales | %t booléen | %p pointeur
type S struct{ Nom string; Port int }
fmt.Printf("%+v\n", S{"dns", 53}) // {Nom:dns Port:53}

// Lecture clavier
var nom string
fmt.Print("Nom d'hôte : ")
fmt.Scanln(&nom)
```

## 32. Fichiers avec `os` : lire, écrire, infos

```go
// Lire tout un (petit) fichier
data, err := os.ReadFile("/etc/hosts")
if err != nil {
	log.Fatal(err)
}

// Écrire (crée ou tronque, droits 0644)
err = os.WriteFile("/tmp/test.txt", []byte("hello\n"), 0644)

// Ajouter à la fin : ouvrir avec flags
f, err := os.OpenFile("/var/log/monoutil.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
if err != nil {
	log.Fatal(err)
}
defer f.Close()
f.WriteString("ligne de log\n")

// Infos et existence
info, err := os.Stat("/etc/hostname")
if err != nil {
	if os.IsNotExist(err) {
		fmt.Println("absent")
	}
} else {
	fmt.Println(info.Size(), info.Mode(), info.ModTime())
}

// Créer / renommer / supprimer
os.MkdirAll("/opt/monoutil/data", 0755) // parents inclus
os.Rename("/tmp/a", "/tmp/b")
os.Remove("/tmp/b")          // fichier vide ou dossier vide
os.RemoveAll("/tmp/cache")   // récursif : dangereux, à manier avec précaution

// Variables d'environnement
os.Getenv("HOME")            // "" si absente
os.LookupEnv("MODE")         // (valeur, présente?)
os.Setenv("MODE", "prod")
```

> Droits : `0644`/`0755` en littéral octal (`0` préfixe). Depuis Go 1.13 on peut
> aussi écrire `0o644`. Préférez `0o644` dans le code neuf.

## 33. `bufio` : traiter les gros fichiers ligne par ligne

Ne jamais `os.ReadFile` un log de 2 Go. Scanner ligne par ligne :

```go
f, err := os.Open("/var/log/syslog")
if err != nil {
	log.Fatal(err)
}
defer f.Close()

sc := bufio.NewScanner(f)
sc.Buffer(make([]byte, 1024*1024), 1024*1024) // lignes > 64 Ko si besoin
n := 0
for sc.Scan() {
	ligne := sc.Text()
	if strings.Contains(ligne, "error") {
		n++
	}
}
if err := sc.Err(); err != nil {
	log.Fatalf("scan: %v", err)
}
fmt.Println("lignes d'erreur:", n)

// Écriture bufferisée (flush !)
w := bufio.NewWriter(f2)
w.WriteString("...")
w.Flush() // sans Flush, rien n'est écrit
```

## 34. `path/filepath` : chemins portables

