---
id: collect-261001-rattrapage/rattrapage/go-guide-2
title: "Guide Go (Golang) — du script sysadmin au service de production"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["benchmarks"]
source: docs/RAG/collect-261001-rattrapage/go_guide.md
source_anchor: ""
source_lines: [236, 512]
sha256: ea26b07b8e3f4908ef32be60ff119ac25557b0cd9c74bd24c7696b94f3e80267
---

# Guide Go (Golang) — du script sysadmin au service de production

// Conversions EXPLICITES obligatoires (pas de promotion implicite)
var i int = 7
var f2 float64 = float64(i) / 2 // 3.5
// var f3 float64 = i / 2       // ERREUR : type mismatch
```

> Pas de conversion implicite en Go : `int` + `int64` ne compile pas. C'est
> verbeux mais ça élimine toute une classe de bugs silencieux.

## 9. Booléens et chaînes : l'essentiel

```go
var ok bool = true
// Les conditions exigent un booléen : pas de 0/1, pas de pointeur nil
// if port { ... }        // ERREUR : port n'est pas un bool
// if 1 { ... }           // ERREUR

var s string = "hello"
s2 := `chemin brut C:\logs\app.log` // backticks : string brute, sans échappement
multi := `ligne1
ligne2` // multi-lignes autorisées

// Concaténation simple (hors boucle chaude)
msg := "hôte " + hote + " KO"
```

## 10. Zéro values : le "constructeur gratuit"

Chaque type a une valeur zéro utilisable immédiatement — pas de constructeur obligatoire.

| Type | Zéro value |
|---|---|
| `int`, `float` | `0` |
| `string` | `""` |
| `bool` | `false` |
| Pointeur, slice, map, channel, fonction, interface | `nil` |

```go
var s Serveur        // struct zéro : champs à leurs zéro values
fmt.Println(s.Port)  // 0 — utilisable sans initialisation

var compteurs map[string]int // nil : LECTURE ok (retourne zéro), ÉCRITURE = panic
n := compteurs["x"]          // 0, pas de panic
// compteurs["x"] = 1        // PANIC : initialiser avec make d'abord

var liste []int      // slice nil : append fonctionne !
liste = append(liste, 1) // OK

// Idiome : profiter du zéro value pour les options
type Options struct {
	Timeout time.Duration // 0 = à interpréter comme défaut dans le code
	Retries int           // 0 = 1 tentative par défaut, à documenter
}
```

---

## 11. Opérateurs

| Famille | Opérateurs |
|---|---|
| Arithmétiques | `+ - * / %` (division entière si les deux opérandes sont entiers) |
| Affectation | `= += -= *= /= %= &= \|= ^= <<= >>=` |
| Comparaison | `== != < <= > >=` |
| Logiques | `&&` (et, court-circuit) `\|\|` (ou, court-circuit) `!` (non) |
| Bits | `&` (et) `\|` (ou) `^` (ou exclusif / NOT unaire) `&^` (ET-NON : clear) `<<` `>>` |

```go
a, b := 7, 3
fmt.Println(a / b)   // 2  (division entière !)
fmt.Println(a % b)   // 1

x := 0b1100           // littéral binaire (Go 1.13+)
fmt.Printf("%b %d\n", x&^0b0100, x) // clear du bit 2

// Pas d'opérateur ternaire en Go ! On écrit :
max := a
if b > a {
	max = b
}
```

Piège classique : `a / b` avec deux `int` tronque. Pour un quotient décimal,
convertir d'abord : `float64(a) / float64(b)`.

## 12. Contrôle : `if`, `switch`

```go
// if avec initialisation (portée limitée au bloc) — idiome très courant
if err := chargerConfig(); err != nil {
	log.Fatalf("config: %v", err)
}

if charge := cpuCharge(); charge > 90 {
	alert("CPU critique")
} else if charge > 70 {
	alert("CPU élevé")
}

// switch sans expression = switch true
switch {
case charge > 90:
	fmt.Println("critique")
case charge > 70:
	fmt.Println("élevé")
default:
	fmt.Println("ok")
}

// switch sur valeur : pas de 'break' nécessaire (pas de fallthrough implicite)
switch jour {
case "sam", "dim": // plusieurs valeurs par case
	fmt.Println("week-end")
case "lun":
	fmt.Println("lundi")
	fallthrough // explicite si vraiment voulu (rare)
default:
	fmt.Println("semaine")
}

// switch de type (voir section 29)
switch v := x.(type) {
case string:
	fmt.Println("string:", v)
case int:
	fmt.Println("int:", v)
}
```

## 13. Boucles : `for` fait tout (pas de `while`)

```go
// Style while
i := 0
for i < 5 {
	fmt.Println(i)
	i++
}

// Classique
for i := 0; i < 5; i++ {
	fmt.Println(i)
}

// Boucle infinie (serveurs, workers)
for {
	travail()
}

// Parcours avec range (index + valeur)
services := []string{"dns", "dhcp", "ntp"}
for i, s := range services {
	fmt.Printf("%d: %s\n", i, s)
}
for _, s := range services { // _ ignore l'index
	fmt.Println(s)
}

// Parcours d'une map (ordre ALÉATOIRE à chaque exécution)
for nom, ip := range hotes {
	fmt.Println(nom, ip)
}

// break / continue avec label (sortir de boucles imbriquées)
boucle:
	for _, h := range hotes {
		for _, p := range ports {
			if injoignable(h, p) {
				break boucle
			}
		}
	}
```

> `for range` sur un entier (`for i := range 10`) : possible depuis **Go 1.22**.
> Avant, écrire `for i := 0; i < 10; i++`.

## 14. Tableaux : taille fixe, rarement utilisés directement

```go
var codes [4]int              // [0 0 0 0], len = 4
jours := [7]string{"lun", "mar", "mer", "jeu", "ven", "sam", "dim"}
fmt.Println(len(jours))       // 7

// Les tableaux se copient par VALEUR à l'affectation et au passage en paramètre
a := [2]int{1, 2}
b := a
b[0] = 99
fmt.Println(a[0]) // 1 : a inchangé
```

En pratique on utilise des **slices** (section 15), pas des tableaux.
Retenez juste : un tableau `[N]T` a une taille qui fait partie de son type.

## 15. Slices en détail : la structure de données n°1

Un slice = vue sur un tableau sous-jacent : **pointeur + longueur + capacité**.

```go
// Créations
s1 := []int{1, 2, 3}          // littéral
s2 := make([]int, 3)          // len=3, cap=3, zéros
s3 := make([]int, 0, 10)      // len=0, cap=10 (pré-allocation !)
var s4 []int                 // slice nil : len=0, cap=0, utilisable directement

fmt.Println(s4 == nil)        // true — append fonctionne quand même

// Ajout (append peut réallouer : TOUJOURS réaffecter le résultat)
s3 = append(s3, 42)
s3 = append(s3, 1, 2, 3)      // plusieurs d'un coup
s3 = append(s3, s1...)        // ... déplie un slice

// Découpage (partage le tableau sous-jacent !)
t := []int{0, 1, 2, 3, 4, 5}
sous := t[1:4]                // [1 2 3], len=3, cap=5
sous[0] = 99
fmt.Println(t[1])             // 99 : modification visible via t !

// Copie indépendante
copie := make([]int, len(sous))
copy(copie, sous)

// Longueur / capacité
fmt.Println(len(t), cap(t))   // 6 6

// Parcours
for i, v := range t {
	fmt.Println(i, v)
}
```

Tableau récapitulatif des opérations :

| Opération | Code | Note |
|---|---|---|
| Créer vide pré-alloué | `make([]T, 0, n)` | Évite les réallocations en boucle |
| Ajouter | `s = append(s, v)` | Réaffecter impérativement |
| Copier | `copy(dst, src)` | Retourne le nb d'éléments copiés |
| Trier (`≥ 1.21` : paquet `slices`) | `slices.Sort(s)` | Ancien : `sort.Ints(s)` |
| Inverser/chercher | `slices.Reverse`, `slices.Contains`, `slices.Index` | Paquet `slices` (`≥ 1.21`) |
| Vider en gardant cap | `s = s[:0]` | Réutilise le tableau |
| Supprimer l'index i | `s = append(s[:i], s[i+1:]...)` | Idiome à connaître |
| Comparer | `slices.Equal(a, b)` | `==` interdit sur slices |

Pense-bête perf : si vous connaissez la taille finale (lecture de N lignes),
`make([]T, 0, n)` évite des dizaines de réallocations. Mesurez avec les
benchmarks (section 62) avant d'optimiser à l'aveugle.

## 16. Maps en détail : dictionnaires

```go
// Création
inventaire := map[string]string{
	"srv-dns-01": "10.0.0.53",
	"srv-ntp-01": "10.0.0.123",
}
compteurs := make(map[string]int) // pré-allouer si taille connue : make(map[string]int, 100)
var lazy map[string]int           // map nil : LECTURE ok, ÉCRITURE = panic !

// Lecture avec test d'existence (idiome "comma ok")
ip, ok := inventaire["srv-dns-01"]
if !ok {
	fmt.Println("hôte inconnu")
}

// Écriture / suppression
inventaire["srv-web-01"] = "10.0.0.80"
delete(inventaire, "srv-ntp-01")

fmt.Println(len(inventaire)) // nombre de clés

// Parcours : ordre aléatoire !
for nom, ip := range inventaire {
	fmt.Println(nom, "->", ip)
}
```

