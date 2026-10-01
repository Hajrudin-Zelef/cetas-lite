---
id: collect-261001-rattrapage/rattrapage/rust-guide-7
title: "Guide Rust ultra-complet — Langage système sûr, rapide et moderne"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["agent", "benchmarks"]
source: docs/RAG/collect-261001-rattrapage/rust_guide.md
source_anchor: ""
source_lines: [1427, 1659]
sha256: 39edaa2cbb644076cd871216565f5cb24073f04aa5670759f087b6b4287b5797
---

# Guide Rust ultra-complet — Langage système sûr, rapide et moderne

1. **`cargo fmt` systématique** avant chaque commit (hook pre-commit ou CI).
2. **Zéro `unwrap()`** dans les chemins applicatifs ; `expect("invariant: ...")` si prouvé.
3. **Nomme explicitement** : `hote`, `latence_ms`, pas `x`, `tmp` (hors boucles triviales).
4. **Petites fonctions** : une fonction = une responsabilité ; < 50 lignes en général.
5. **Erreurs typées** (`thiserror`) dans les libs, `anyhow` + contexte dans les binaires.
6. **Pas de `panic!` pour une erreur récupérable** (fichier manquant, réseau coupé).
7. **Documentation** : `///` sur toute fonction publique, avec exemple (doctest).
8. **`&str` en paramètre, `String` en retour** quand c'est possible (flexibilité d'appel).
9. **`&[T]` plutôt que `&Vec<T>`** en paramètre (accepte slices et Vec).
10. **Évite les clones défensifs** : préfère les emprunts ; clone quand la sémantique l'exige.
11. **Tests** : chaque correction de bug = un test de non-régression.
12. **Édition** : choisis 2021 ou 2024 et reste cohérent dans le workspace.

---

## 33. Performance : mesurer avant d'optimiser

```bash
cargo build --release              # toujours mesurer en release !
```

```rust
use std::time::Instant;

fn main() {
    // Mesure artisanale
    let debut = Instant::now();
    let v: Vec<u64> = (0..10_000_000).map(|x| x * 2).collect();
    println!("collect : {:?} ({} éléments)", debut.elapsed(), v.len());

    // Comparer deux approches : écrire un bench...
}
```

Benchmarks sérieux :

```toml
[dev-dependencies]
criterion = "0.5"

[[bench]]
name = "mon_bench"
harness = false
```

```rust
// benches/mon_bench.rs
use criterion::{black_box, criterion_group, criterion_main, Criterion};

fn bench_filtrage(c: &mut Criterion) {
    let donnees: Vec<u64> = (0..100_000).collect();
    c.bench_function("filtre pair", |b| {
        b.iter(|| donnees.iter().filter(|x| *x % 2 == 0).count())
    });
    // black_box empêche le compilateur d'optimiser le calcul hors de la mesure
    c.bench_function("avec black_box", |b| {
        b.iter(|| black_box(&donnees).iter().filter(|x| *x % 2 == 0).count())
    });
}
criterion_group!(benches, bench_filtrage);
criterion_main!(benches);
```

Leviers classiques :

| Levier | Effet typique |
|---|---|
| `cargo build --release` (opt-level=3) | Base obligatoire |
| Éviter les allocations en boucle (`String::with_capacity`) | Moins de pression sur l'allocateur |
| Itérateurs plutôt que boucles indexées | Même perf, moins d'erreurs |
| `lto = true`, `codegen-units = 1` | Binaire plus rapide, compilation plus lente |
| `jemalloc` / `mimalloc` (crates) | Allocateur plus rapide selon workload |
| Éviter `Mutex` en section chaude | Contention = poison de la perf |

> Règle : profile d'abord (`perf`, `flamegraph` via `cargo install flamegraph`), optimise ensuite.
> 90 % du temps passé dans 10 % du code.

---

## 34. Sécurité mémoire : ce que le compilateur garantit (et pas)

**Garanti à la compilation** (code safe) :

- Pas de déréférencement de pointeur nul (pas de `null`).
- Pas de *use-after-free* (ownership).
- Pas de double libération (move).
- Pas de data race (règles d'emprunt + `Send`/`Sync`).
- Pas de dépassement de buffer non vérifié (indexation contrôlée, slices avec longueur).

**NON garanti** (à toi de gérer) :

| Risque | Mitigation |
|---|---|
| Panic par `unwrap` / index hors bornes | `get()`, `?`, valeurs par défaut |
| Dépassement arithmétique silencieux en release | `checked_*` / `saturating_*` |
| Fuite mémoire (`mem::forget`, cycles `Rc`) | Rare ; `Weak` pour casser les cycles |
| Logique métier fausse | Tests, types forts (newtype) |
| Bloc `unsafe` | Audit obligatoire (voir §35) |
| Dépendances vulnérables | `cargo audit`, `cargo deny` en CI |

> « Fearless concurrency » n'est pas un slogan : si ça compile, les data races sont exclus
> (hors `unsafe`). C'est LA raison d'écrire un agent multi-thread en Rust plutôt qu'en C.

---

## 35. `unsafe` : quand et comment (avec parcimonie)

`unsafe` ne désactive pas le borrow checker : il autorise 5 opérations normalement interdites :

1. Déréférencer un pointeur brut (`*const T`, `*mut T`).
2. Appeler une fonction `unsafe` (dont les FFI C).
3. Accéder/muter une variable `static mut` (déconseillé ; préférer `Atomic*`/`Mutex`).
4. Implémenter un trait `unsafe` (`Send`, `Sync` manuels).
5. Accéder aux champs d'une `union`.

```rust
// Exemple canonique : FFI vers libc (voir §46)
extern "C" {
    fn getpid() -> i32;
}

fn main() {
    let pid = unsafe { getpid() };   // l'appel FFI est unsafe : l'appelant assume le contrat C
    println!("pid = {pid}");
}
```

Contrat d'équipe pour `unsafe` :

- [ ] Encapsuler dans une **abstraction sûre** : fonction safe qui contient un petit bloc unsafe.
- [ ] Documenter l'invariant avec `// SAFETY: ...` (convention officielle).
- [ ] Justifier pourquoi le safe ne suffit pas.
- [ ] `cargo geiger` / audit en revue de code.
- [ ] Miri (`cargo +nightly miri test`) pour traquer les UB dans les tests.

```rust
/// Retourne le pid du processus.
///
/// # Safety
/// Aucune : fonction 100 % sûre, l'unsafe est encapsulé.
fn mon_pid() -> i32 {
    // SAFETY: getpid ne peut pas échouer et ne touche pas de mémoire partagée.
    unsafe { getpid() }
}
```

---

## 36. Erreurs classiques des débutants (10+, avec corrections)

### Erreur 1 — `cannot borrow as mutable` : deux emprunts incompatibles

```rust
// ❌ NE COMPILE PAS
// let mut v = vec![1, 2, 3];
// let premier = &v[0];
// v.push(4);                    // push peut réallouer -> invaliderait `premier`
// println!("{premier}");

// ✅ Copier la valeur (types Copy) ou terminer l'emprunt d'abord
let mut v = vec![1, 2, 3];
let premier = v[0];              // i32 est Copy : pas d'emprunt
v.push(4);
println!("{premier}");
```

### Erreur 2 — `borrow of moved value`

```rust
// ❌ let s = String::from("x"); let t = s; println!("{s}");
// ✅
let s = String::from("x");
let t = s.clone();               // copie explicite si on veut garder s
println!("{s} {t}");
// ✅ ou : passer par référence
fn affiche(s: &str) { println!("{s}"); }
let s2 = String::from("y");
affiche(&s2);
println!("{s2}");                // toujours valide
```

### Erreur 3 — Oublier que `if`/`match` sans `;` retourne une valeur

```rust
fn statut(charge: u8) -> &'static str {
    if charge > 90 {
        "critique";   // ❌ le `;` transforme en statement -> la fonction retourne ()
    } else {
        "ok";         // ❌ idem
    }
}
// ✅ retirer les points-virgules finaux
fn statut_ok(charge: u8) -> &'static str {
    if charge > 90 { "critique" } else { "ok" }
}
```

### Erreur 4 — Comparer `String` et `&str` avec `==`… ça marche, mais l'addition non

```rust
let a = "hello".to_string();
// let b = a + " world";         // ❌ : String + &str existe, mais &str + &str NON
let b = a + " world";            // ✅ en fait ça compile : String + &str -> String
let c = "hello";
// let d = c + " world";         // ❌ vraiment interdit : &str + &str n'existe pas
let d = format!("{c} world");    // ✅ format! pour concaténer des &str
println!("{b} {d}");
```

### Erreur 5 — `Vec` déplacé par `for x in v`

```rust
let v = vec![1, 2, 3];
// for x in v { }                 // ❌ into_iter : v est déplacé (sauf si Copy... non, Vec ne l'est pas)
// println!("{:?}", v);           // ❌ v n'existe plus
// ✅
let v = vec![1, 2, 3];
for x in &v { println!("{x}"); } // emprunt : v survit
println!("{:?}", v);
```

### Erreur 6 — Lifetime manquante sur struct avec référence

```rust
// ❌ struct Config { chemin: &str }   // missing lifetime specifier
// ✅
struct Config<'a> { chemin: &'a str }
// Mieux : posséder plutôt qu'emprunter quand c'est possible
struct ConfigPossedee { chemin: String }   // zéro annotation, zéro souci
```

### Erreur 7 — `unwrap()` sur `Result` dans `main` sans réfléchir

