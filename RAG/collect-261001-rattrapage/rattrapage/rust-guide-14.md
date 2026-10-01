---
id: collect-261001-rattrapage/rattrapage/rust-guide-14
title: "Guide Rust ultra-complet — Langage système sûr, rapide et moderne"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["agent", "datacenter", "incident", "memory"]
source: docs/RAG/collect-261001-rattrapage/rust_guide.md
source_anchor: ""
source_lines: [3116, 3346]
sha256: a6461d95e20badf2a0efd7871280c6ee3a284dffa6afb0b0805fc173d8820911
---

# Guide Rust ultra-complet — Langage système sûr, rapide et moderne

```rust
fn main() {
    // 1. Pré-allouer quand la taille est connue
    let mut v = Vec::with_capacity(10_000);   // évite les réallocations successives
    let mut s = String::with_capacity(1024);

    // 2. Réutiliser au lieu de recréer (boucles chaudes)
    let mut tampon = String::new();
    for i in 0..1000 {
        tampon.clear();                        // garde la capacité allouée
        tampon.push_str(&format!("ligne {i}"));
    }

    // 3. Emprunter plutôt que cloner dans les signatures
    // fn f(s: &str) plutôt que fn f(s: String) quand on ne stocke pas

    // 4. Cow : clone uniquement si mutation nécessaire
    use std::borrow::Cow;
    fn normaliser(entree: &str) -> Cow<str> {
        if entree.chars().all(|c| !c.is_whitespace()) {
            Cow::Borrowed(entree)              // zéro allocation
        } else {
            Cow::Owned(entree.split_whitespace().collect::<Vec<_>>().join(" "))
        }
    }
    println!("{}", normaliser("hello"));
    println!("{}", normaliser("hel lo"));

    // 5. Éviter format! en boucle chaude : write! dans un buffer réutilisé
    use std::fmt::Write as _;
    let mut buf = String::with_capacity(64);
    write!(buf, "val={}", 42).unwrap();
    println!("{buf} cap={}", v.capacity());
    let _ = s;
}
```

---

## 72. Panics : quand c'est légitime (et comment les contenir)

Légitime :

- Invariant interne violé = bug (`unreachable!`, `assert!` sur préconditions internes).
- Échec d'initialisation irrécupérable au démarrage (config absente → exit propre plutôt).
- Tests et prototypes.

Illégitime : fichier manquant, réseau coupé, entrée utilisateur invalide → `Result`.

```rust
fn main() {
    // catch_unwind : dernier recours (frontière FFI, plugin, worker isolé)
    let resultat = std::panic::catch_unwind(|| {
        // code qui pourrait paniquer (dépendance fragile)
        let v = vec![1, 2, 3];
        v[10]   // panic : index out of bounds
    });
    match resultat {
        Ok(val) => println!("ok : {val}"),
        Err(_) => eprintln!("le worker a paniqué, on continue"),
    }

    // assert! avec message : documente l'invariant
    let seuil = 90u8;
    assert!(seuil <= 100, "seuil {seuil} incohérent : doit être <= 100");
}
```

> En lib : ne panique jamais sur des entrées externes. En binaire : `panic = "abort"`
> en release si tu veux des binaires plus petits ET aucun unwinding à travers du C.

---

## 73. Données et fichiers : CSV, TOML, YAML

```toml
[dependencies]
csv = "1"
serde = { version = "1", features = ["derive"] }
toml = "0.8"
```

```rust
use serde::{Deserialize, Serialize};

#[derive(Debug, Deserialize, Serialize)]
struct Mesure {
    horodatage: String,
    equipement: String,
    watts: f64,
}

fn lire_csv(chemin: &str) -> Result<Vec<Mesure>, csv::Error> {
    let mut lecteur = csv::Reader::from_path(chemin)?;
    lecteur.deserialize().collect()   // Vec<Mesure> via FromIterator<Result>
}

fn ecrire_csv(mesures: &[Mesure]) -> Result<(), csv::Error> {
    let mut ecrivain = csv::Writer::from_writer(std::io::stdout());
    for m in mesures {
        ecrivain.serialize(m)?;
    }
    ecrivain.flush()?;
    Ok(())
}

#[derive(Debug, Deserialize)]
struct ConfToml {
    general: General,
}
#[derive(Debug, Deserialize)]
struct General {
    site: String,
    #[serde(default)]
    verbose: bool,
}

fn lire_toml(texte: &str) -> Result<ConfToml, toml::de::Error> {
    toml::from_str(texte)
}

fn main() {
    // CSV : mesures.csv -> horodatage,equipement,watts
    // TOML :
    //   [general]
    //   site = "datacenter-nord"
    let cfg = lire_toml("[general]\nsite = \"datacenter-nord\"\n").unwrap();
    println!("site : {}", cfg.general.site);
    let _ = (lire_csv, ecrire_csv);
}
```

---

## 74. Base de données : SQL avec sqlx (intro)

```toml
[dependencies]
sqlx = { version = "0.8", features = ["runtime-tokio", "sqlite"] }
tokio = { version = "1", features = ["full"] }
```

```rust
#[derive(Debug)]
struct EquipementDb { id: i64, hostname: String }

async fn demo() -> Result<(), sqlx::Error> {
    // SQLite en mémoire : parfait pour tests/outils embarqués
    let pool = sqlx::SqlitePool::connect("sqlite::memory:").await?;

    sqlx::query("CREATE TABLE equipements (id INTEGER PRIMARY KEY, hostname TEXT NOT NULL)")
        .execute(&pool).await?;

    // Requête VÉRIFIÉE À LA COMPILATION avec query! (nécessite DATABASE_URL au build)
    // let eq = sqlx::query_as!(EquipementDb,
    //     "SELECT id, hostname FROM equipements WHERE id = ?", 1u32)
    //     .fetch_one(&pool).await?;

    // Version dynamique (sans vérification compile-time) :
    sqlx::query("INSERT INTO equipements (hostname) VALUES (?)")
        .bind("srv-01")
        .execute(&pool).await?;
    let rows: Vec<EquipementDb> = sqlx::query_as("SELECT id, hostname FROM equipements")
        .fetch_all(&pool).await
        .map(|v| v.into_iter().map(|(id, hostname): (i64, String)|
            EquipementDb { id, hostname }).collect())?;

    println!("{rows:?}");
    Ok(())
}

#[tokio::main]
async fn main() {
    if let Err(e) = demo().await { eprintln!("db : {e}"); }
}
```

> `query!` vérifie le SQL **à la compilation** contre une vraie base (via `DATABASE_URL`) :
> colonne renommée = erreur de build, pas incident de prod. Postgres/MySQL/SQLite supportés.

---

## 75. Checklist de revue de code Rust (avant chaque merge)

- [ ] `cargo fmt --check` passe
- [ ] `cargo clippy -- -D warnings` passe
- [ ] Aucun `unwrap()`/`expect()` injustifié dans le code applicatif
- [ ] Aucun `panic!` sur erreur récupérable (I/O, réseau, parsing)
- [ ] Erreurs avec contexte (`anyhow`) côté binaire, typées (`thiserror`) côté lib
- [ ] `&str`/`&[T]` en paramètres plutôt que `&String`/`&Vec<T>`
- [ ] Pas de clone défensif sans raison (emprunts privilégiés)
- [ ] Lifetimes explicites uniquement où nécessaire (retour de référence)
- [ ] `unsafe` encapsulé + commentaire `// SAFETY:`
- [ ] Tests : nouveau comportement = nouveau test ; bug = test de non-régression
- [ ] Doc `///` sur les items publics (+ doctests qui passent)
- [ ] `Cargo.lock` committé (binaire) ; features additives
- [ ] `cargo audit` sans vulnérabilité bloquante
- [ ] Logs sur stderr / données sur stdout ; codes de sortie conventionnels
- [ ] Mesuré en `--release` si perf revendiquée

---

## 76. Anti-patterns : ce qu'il ne faut plus faire (venu d'autres langages)

| Habitude d'ailleurs | Équivalent Rust idiomatique |
|---|---|
| `null` / `None` checks partout (Python) | `Option<T>` + combinateurs (`map`, `?`) |
| Exceptions (Java/Python) | `Result<T, E>` + `?` + thiserror/anyhow |
| Héritage de classes (Java/C++) | Traits + composition (pas d'héritage en Rust) |
| Getters/setters systématiques | Champs `pub` si trivial, méthodes si logique |
| Variables globales mutables (C) | `OnceLock` / `Mutex` / paramètres explicites |
| `char*` et `strcpy` (C) | `String`/`&str` (UTF-8 garanti) |
| Threads + mémoire partagée sans garde (C) | `Arc<Mutex<T>>` ou channels (vérifié à la compilation) |
| `void*` générique (C) | Génériques `T` ou `dyn Trait` |
| Macros préprocesseur `#define` (C) | `const`, `macro_rules!`, generics |
| GC implicite (Go/Java) | Ownership : libération déterministe, zéro GC |

---

## 77. Ressources système : aller plus loin côté sysadmin

Pistes de projets pour ancrer les réflexes :

1. **Exporter de métriques** : lire `/proc`, exposer en Prometheus (axum + endpoint `/metrics`).
2. **Agent d'inventaire** : SSH/parallèle (threads), sortie JSON (serde), config TOML.
3. **Rotate de logs** : surveiller un fichier, compresser, retenir N archives (fs + chrono).
4. **Watchdog de service** : boucle de sondes, alertes, backoff exponentiel.
5. **Outil de sauvegarde** : parcours récursif (`walkdir`), hash (sha2), manifest JSON.

Crates utiles au sysadmin :

