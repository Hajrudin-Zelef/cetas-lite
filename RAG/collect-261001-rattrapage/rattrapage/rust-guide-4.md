---
id: collect-261001-rattrapage/rattrapage/rust-guide-4
title: "Guide Rust ultra-complet — Langage système sûr, rapide et moderne"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["cost"]
source: docs/RAG/collect-261001-rattrapage/rust_guide.md
source_anchor: ""
source_lines: [681, 923]
sha256: 865255ac1e512434f04d2da0951103013551ca99f74a56ecda7aabd503b7024b
---

# Guide Rust ultra-complet — Langage système sûr, rapide et moderne

    // ok_or : convertir Option -> Result
    let r: Result<u16, &str> = trouver_port(&services, "dns").ok_or("service inconnu");
    println!("{q} {r:?}");
}
```

Méthodes essentielles de `Option<T>` :

| Méthode | Effet |
|---|---|
| `.unwrap()` | Panic si `None` — **à éviter** hors tests/prototypes |
| `.expect("msg")` | Panic avec message — acceptable si l'invariant est prouvé |
| `.unwrap_or(défaut)` / `.unwrap_or_default()` / `.unwrap_or_else(f)` | Repli sans panic |
| `.map(f)` / `.and_then(f)` | Transformation chaînée |
| `.filter(prédicat)` | `Some` → `None` si faux |
| `.ok_or(err)` / `.ok_or_else(f)` | Vers `Result` |
| `.is_some()` / `.is_none()` | Test |

> Règle d'équipe : `unwrap()` dans du code applicatif = code smell. `expect()` avec un message
> qui explique l'invariant = tolérable. Valeur par défaut ou propagation `?` = idéal.

---

## 17. `Result<T, E>` : les erreurs sont des valeurs (pas d'exceptions)

Pas d'exceptions en Rust. Une fonction qui peut échouer retourne `Result<T, E>` :
`Ok(valeur)` ou `Err(erreur)`. **L'appelant DOIT traiter les deux cas** (sinon warning
`unused_must_use`, et clippy hurle).

```rust
use std::fs;
use std::io;

fn lire_config(chemin: &str) -> Result<String, io::Error> {
    fs::read_to_string(chemin)   // retourne déjà un Result : on le propage tel quel
}

fn main() {
    // Style 1 : match exhaustif
    match lire_config("/etc/mon_outil.conf") {
        Ok(contenu) => println!("config : {} octets", contenu.len()),
        Err(e) => eprintln!("impossible de lire la config : {e}"),
    }

    // Style 2 : propagation avec `?` (dans une fonction qui retourne Result)
    if let Err(e) = demarrer() {
        eprintln!("échec démarrage : {e}");
        std::process::exit(1);
    }
}

// L'opérateur `?` : si Err, retourne immédiatement l'erreur à l'appelant
fn demarrer() -> Result<(), io::Error> {
    let cfg = lire_config("/etc/mon_outil.conf")?;   // <- le `?` magique
    println!("démarrage avec {} octets de config", cfg.len());
    Ok(())
}
```

**Pourquoi pas d'exceptions ?** Parce qu'une exception rend invisible le chemin d'erreur :
tu ne sais pas, en lisant une signature, ce qui peut casser. Avec `Result`, la signature
`fn f() -> Result<T, E>` **documente** l'échec. Le `?` garde le code lisible sans cacher l'erreur.

---

## 18. Traits : le polymorphisme à la Rust

Un **trait** définit un comportement partagé (comme une interface, mais avec super-pouvoirs :
méthodes par défaut, implémentation pour types existants, génériques).

```rust
trait Resume {
    fn resume(&self) -> String;
    // Méthode par défaut : implémentation fournie, surchargeable
    fn resume_court(&self) -> String {
        let r = self.resume();
        r.chars().take(40).collect()
    }
}

struct Onduleur { modele: String, charge_pct: u8 }
struct GroupeElectrogene { modele: String, carburant_l: f64 }

impl Resume for Onduleur {
    fn resume(&self) -> String {
        format!("UPS {} : {}%", self.modele, self.charge_pct)
    }
}
impl Resume for GroupeElectrogene {
    fn resume(&self) -> String {
        format!("GE {} : {} L", self.modele, self.carburant_l)
    }
}

// Fonction générique bornée par le trait
fn afficher<T: Resume>(equipement: &T) {
    println!("{}", equipement.resume());
}

// Syntaxe moderne équivalente (impl Trait)
fn afficher2(equipement: &impl Resume) {
    println!("{}", equipement.resume_court());
}

// Plusieurs bornes : T: Resume + Clone
// where clause pour lisibilité :
fn comparer<T>(a: &T, b: &T) -> bool
where T: Resume + PartialEq {
    a == b && a.resume() == b.resume()
}
```

### Traits dérivables (derive)

```rust
#[derive(Debug, Clone, PartialEq)]
struct Config {
    hote: String,
    port: u16,
}

fn main() {
    let c1 = Config { hote: "srv".into(), port: 443 };
    let c2 = c1.clone();            // Clone explicite (jamais implicite !)
    println!("{c1:?}");             // Debug : {:#?} pour pretty-print
    println!("égaux : {}", c1 == c2);
}
```

| Derive | Rôle |
|---|---|
| `Debug` | Affichage développeur `{:?}` |
| `Clone` | Copie explicite `.clone()` |
| `Copy` | Copie implicite (types simples uniquement) |
| `PartialEq` / `Eq` | `==` / `!=` |
| `PartialOrd` / `Ord` | `<`, tri |
| `Hash` | Clé de `HashMap` |
| `Default` | `Config::default()` |

---

## 19. Génériques : écrire une fois, typer fort

```rust
// Fonction générique : le type est décidé à l'appel, vérifié à la compilation
fn premier<T>(liste: &[T]) -> Option<&T> {
    liste.first()
}

// Struct générique avec contrainte
struct Pile<T> {
    elements: Vec<T>,
}

impl<T> Pile<T> {
    fn new() -> Self { Self { elements: Vec::new() } }
    fn empiler(&mut self, v: T) { self.elements.push(v); }
    fn depiler(&mut self) -> Option<T> { self.elements.pop() }
}

fn main() {
    let nombres = [10, 20, 30];
    println!("{:?}", premier(&nombres));       // Option<&i32>
    let mots = ["a", "b"];
    println!("{:?}", premier(&mots));          // Option<&&str>

    let mut p: Pile<String> = Pile::new();     // monomorphisation : un code par type utilisé
    p.empiler("hello".to_string());
    println!("{:?}", p.depiler());
}
```

> **Monomorphisation** : le compilateur génère une version spécialisée par type concret.
> Zéro coût à l'exécution (pas de boxing, pas de dispatch dynamique), binaire un peu plus gros.
> Pour du dispatch dynamique, voir les *trait objects* (`dyn Trait`, §66).

---

## 20. Collections : Vec, HashMap, String et consorts

```rust
use std::collections::HashMap;

fn main() {
    // ---- Vec<T> : tableau dynamique ----
    let mut v = Vec::new();
    v.push("srv-01");
    v.push("srv-02");
    let mut v2 = vec![80u16, 443, 8080];       // macro vec!
    v2.push(8443);
    println!("2e : {}", v2[1]);                 // indexation (panic si hors bornes)
    println!("sûr : {:?}", v2.get(99));         // get -> Option, jamais de panic

    // Itération sans consommer / en consommant / mutable
    for x in &v2 { println!("lecture {x}"); }
    for x in &mut v2 { *x += 1; }
    let somme: u16 = v2.iter().sum();

    // ---- HashMap<K, V> ----
    let mut latences = HashMap::new();
    latences.insert("srv-01", 12u64);
    latences.insert("srv-02", 45u64);
    // entry : insérer seulement si absent (évite double lookup)
    latences.entry("srv-03").or_insert(99);
    if let Some(l) = latences.get("srv-01") {
        println!("latence srv-01 : {l} ms");
    }
    for (hote, lat) in &latences {
        println!("{hote} -> {lat} ms");
    }

    // ---- String vs &str ----
    let mut s = String::new();
    s.push('a');                 // un char
    s.push_str("bc");            // une tranche
    let s2 = format!("{s}-suffixe");  // format! retourne String
    // let bad = s + "x";        // &str + &str n'existe PAS ; String + &str existe
    let ok = s + "x";            // String + &str -> String (move de s)
    println!("{ok} somme={somme}");
}
```

| Collection | Quand l'utiliser |
|---|---|
| `Vec<T>` | Liste ordonnée, taille dynamique (le défaut) |
| `VecDeque<T>` | File double-ended (push/pop aux deux bouts en O(1)) |
| `HashMap<K,V>` | Association clé→valeur (clés `Hash + Eq`) |
| `BTreeMap<K,V>` | Comme HashMap mais trié par clé (`Ord`) |
| `HashSet<T>` / `BTreeSet<T>` | Ensemble, unicité |
| `String` | Texte possédé, modifiable (UTF-8 garanti) |
| `&str` | Vue texte empruntée |

> ⚠️ `String` est **toujours** de l'UTF-8 valide. L'indexation `s[0]` est interdite
> (un « caractère » = nombre variable d'octets). Utilise `.chars()`, `.bytes()`, ou les slices
> sur frontières UTF-8 (`.get(0..2)` retourne `Option`).

---

## 21. Itérateurs : la puissance fonctionnelle sans surcoût

Les itérateurs Rust sont **paresseux** (lazy) et **zero-cost** : `map`/`filter` ne font rien
tant qu'on ne consomme pas, et le code généré égale une boucle manuelle.

