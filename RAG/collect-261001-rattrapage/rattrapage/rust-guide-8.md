---
id: collect-261001-rattrapage/rattrapage/rust-guide-8
title: "Guide Rust ultra-complet — Langage système sûr, rapide et moderne"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["consumer"]
source: docs/RAG/collect-261001-rattrapage/rust_guide.md
source_anchor: ""
source_lines: [1660, 1928]
sha256: 6eb31fc2f7e17cd575c085cfeb474de21c48fc5a346487743470920a6ca7a940
---

# Guide Rust ultra-complet — Langage système sûr, rapide et moderne

```rust
// ❌ std::fs::read_to_string("/etc/fichier").unwrap();  // panic si absent
// ✅
fn main() -> anyhow::Result<()> {
    let c = std::fs::read_to_string("/etc/fichier")?;     // erreur propre + code 1
    println!("{c}");
    Ok(())
}
```

### Erreur 8 — Muter via une référence immuable

```rust
// ❌ let v = vec![1]; v.push(2);   // cannot borrow as mutable
// ✅
let mut v = vec![1];
v.push(2);
```

### Erreur 9 — `&String` vs `&str` : accepter trop étroit

```rust
// ❌ fn f(s: &String)  — refuse les littéraux &str
// ✅
fn f(s: &str) { println!("{s}"); }   // accepte &String (deref coercion) ET &str
```

### Erreur 10 — Ignorer le `Result` retourné (`unused_must_use`)

```rust
use std::io::Write;
// ❌ std::io::stdout().write_all(b"x");   // warning : unused Result
// ✅
let _ = std::io::stdout().write_all(b"x"); // assumé explicitement
// ✅ ou : std::io::stdout().write_all(b"x")?;
```

### Erreur 11 — Pattern `match` non exhaustif

```rust
enum S { A, B, C }
// ❌ match s { S::A => {}, S::B => {} }   // non-exhaustive patterns: `C` not covered
// ✅ ajouter S::C => {} ou _ => {}
```

### Erreur 12 — Shadowing involontaire qui masque un bug

```rust
let mut total = 0;
let total = total + 5;   // ⚠️ shadowing : le `mut` initial ne sert plus ; souvent un oubli
println!("{total}");
```

---

## 37. Cas pratique 1 — Analyseur de logs (commenté)

Objectif : compter les niveaux de log d'un fichier, avec gestion d'erreurs propre.

```rust
use std::collections::HashMap;
use std::env;
use std::fs::File;
use std::io::{BufRead, BufReader};
use std::process::ExitCode;

/// Extrait le niveau (INFO/WARN/ERROR/...) en tête de ligne.
/// Retourne None si la ligne ne commence pas par un niveau connu.
fn niveau(ligne: &str) -> Option<&str> {
    let premier = ligne.split_whitespace().next()?;
    match premier {
        "INFO" | "WARN" | "ERROR" | "DEBUG" => Some(premier),
        _ => None,
    }
}

fn analyser(chemin: &str) -> Result<HashMap<String, u64>, Box<dyn std::error::Error>> {
    let fichier = File::open(chemin)?;                    // ? propage io::Error
    let mut compteurs: HashMap<String, u64> = HashMap::new();
    for ligne in BufReader::new(fichier).lines() {
        let ligne = ligne?;
        if let Some(n) = niveau(&ligne) {
            *compteurs.entry(n.to_string()).or_insert(0) += 1;  // entry : idiome canonique
        }
    }
    Ok(compteurs)
}

fn main() -> ExitCode {
    let chemin = env::args().nth(1).unwrap_or_else(|| "/var/log/syslog".to_string());
    match analyser(&chemin) {
        Ok(c) => {
            let mut cles: Vec<_> = c.keys().collect();
            cles.sort();                                  // affichage déterministe
            for k in cles {
                println!("{k}: {}", c[k]);
            }
            ExitCode::SUCCESS
        }
        Err(e) => {
            eprintln!("erreur : {e}");
            ExitCode::FAILURE
        }
    }
}
```

Points à noter :

- `Box<dyn Error>` : type d'erreur générique pour un petit outil (anyhow ferait pareil, en mieux).
- `entry().or_insert()` : l'idiome compteur, à connaître par cœur.
- Tri des clés : sortie déterministe = diffable, scriptable.

---

## 38. Cas pratique 2 — Sonde réseau parallèle (threads + channels)

Objectif : pinger N hôtes en parallèle, agréger les résultats via un channel.

```rust
use std::net::TcpStream;
use std::sync::mpsc;
use std::thread;
use std::time::Duration;

#[derive(Debug)]
struct ResultatSonde {
    hote: String,
    joignable: bool,
    latence_ms: u128,
}

fn sonder(hote: &str, port: u16) -> ResultatSonde {
    let debut = std::time::Instant::now();
    let joignable = TcpStream::connect_timeout(
        &format!("{hote}:{port}").parse().unwrap(),
        Duration::from_millis(800),
    )
    .is_ok();
    ResultatSonde { hote: hote.to_string(), joignable, latence_ms: debut.elapsed().as_millis() }
}

fn main() {
    let hotes = vec!["10.0.0.1", "10.0.0.2", "10.0.0.3", "10.0.0.4"];
    let (tx, rx) = mpsc::channel();       // multiple producers, single consumer

    for hote in hotes {
        let tx = tx.clone();              // un émetteur par thread
        let hote = hote.to_string();      // posséder les données envoyées au thread
        thread::spawn(move || {           // `move` : le thread possède son environnement
            let r = sonder(&hote, 443);
            let _ = tx.send(r);           // envoi (Result ignoré si rx fermé)
        });
    }
    drop(tx);                             // ferme l'émetteur d'origine : rx se termine

    for r in rx {                         // itère jusqu'à fermeture du channel
        let statut = if r.joignable { "OK  " } else { "KO  " };
        println!("{statut} {:12} {} ms", r.hote, r.latence_ms);
    }
}
```

> `move` sur la closure : obligatoire, car le thread peut survivre à la fonction.
> Le borrow checker garantit qu'aucune donnée ne meurt pendant que le thread l'utilise.

---

## 39. Cas pratique 3 — Mini inventaire JSON (serde)

```toml
[dependencies]
serde = { version = "1", features = ["derive"] }
serde_json = "1"
anyhow = "1"
```

```rust
use anyhow::{Context, Result};
use serde::{Deserialize, Serialize};
use std::fs;

#[derive(Debug, Serialize, Deserialize)]
struct Equipement {
    hostname: String,
    ip: String,
    #[serde(default = "port_defaut")]      // valeur par défaut si champ absent
    port: u16,
    #[serde(default)]
    tags: Vec<String>,                     // défaut = Vec::new()
}

fn port_defaut() -> u16 { 22 }

#[derive(Debug, Serialize, Deserialize)]
struct Inventaire {
    version: u8,
    equipements: Vec<Equipement>,
}

fn charger(chemin: &str) -> Result<Inventaire> {
    let brut = fs::read_to_string(chemin)
        .with_context(|| format!("lecture {chemin}"))?;
    let inv: Inventaire = serde_json::from_str(&brut)
        .with_context(|| format!("parse JSON {chemin}"))?;
    Ok(inv)
}

fn main() -> Result<()> {
    // Exemple de JSON accepté :
    // {"version":1,"equipements":[{"hostname":"sw-01","ip":"10.0.1.2","tags":["core"]}]}
    let inv = charger("inventaire.json")?;
    println!("{} équipements (v{})", inv.equipements.len(), inv.version);
    let json = serde_json::to_string_pretty(&inv)?;   // sérialisation inverse
    fs::write("inventaire.pretty.json", json)?;
    Ok(())
}
```

> `serde` = LE pilier de l'écosystème (JSON, YAML, TOML, CSV…). `#[serde(default)]`,
> `rename`, `skip_serializing_if` : à connaître pour des formats de config tolérants.

---

## 40. Pense-bête de poche (une page)

```
--- Ownership ---
let s = String::from("x");   // s possède
let t = s;                   // MOVE : s invalide
let u = t.clone();           // copie profonde explicite
fn f(s: &str) {}             // emprunt : le prêteur garde la propriété

--- Emprunts ---
&r      immuable, N autorisés
&mut r  mutable, UN seul, pas de & en même temps

--- Erreurs ---
fn f() -> Result<T, E>       // pas d'exceptions
val?                         // propage Err, ou unwrap le Ok
opt.ok_or("msg")?            // Option -> Result

--- Match ---
match v { P1 => .., P2(x) => .., _ => .. }   // exhaustif !
if let Some(x) = opt { .. }
let Some(x) = opt else { return; };

--- Collections ---
vec![1,2,3]   v.push(x)   v.get(i) -> Option
map.entry(k).or_insert(v)    // idiome compteur/insertion
s.chars() / s.bytes()        // jamais s[i]

--- Itérateurs ---
.iter() .iter_mut() .into_iter()
.map .filter .collect::<Vec<_>>() .sum() .count()

--- Cargo ---
cargo check   (rapide)      cargo clippy -- -D warnings
cargo fmt     (toujours)    cargo test -- --nocapture

--- Conventions ---
&str en param, String en retour | &[T] pas &Vec<T>
unwrap() interdit en prod     | eprintln! pour les logs
```

---

## 41. Glossaire

