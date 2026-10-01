---
id: collect-261001-rattrapage/rattrapage/rust-guide-6
title: "Guide Rust ultra-complet — Langage système sûr, rapide et moderne"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/rust_guide.md
source_anchor: ""
source_lines: [1172, 1426]
sha256: f9889090e39ae7e4a0ed573b51ab23cb04994a0eaef73b3b4c9845bb781dfb9d
---

# Guide Rust ultra-complet — Langage système sûr, rapide et moderne

fn mini_serveur() -> std::io::Result<()> {
    let listener = TcpListener::bind("0.0.0.0:7878")?;
    println!("écoute sur 7878");
    for flux in listener.incoming() {          // itérateur de connexions
        let mut stream = flux?;               // TcpStream
        let mut buf = [0u8; 1024];
        let n = stream.read(&mut buf)?;
        stream.write_all(b"HTTP/1.1 200 OK\r\n\r\npong")?;
        println!("requête : {} octets", n);
    }
    Ok(())
}

fn sonde_tcp(hote: &str, port: u16, timeout_ms: u64) -> bool {
    let addr = format!("{hote}:{port}");
    let timeout = Duration::from_millis(timeout_ms);
    match addr.parse::<std::net::SocketAddr>() {
        Ok(sa) => TcpStream::connect_timeout(&sa, timeout).is_ok(),
        Err(_) => false,
    }
}

fn envoi_udp() -> std::io::Result<()> {
    let sock = UdpSocket::bind("0.0.0.0:0")?;     // port éphémère
    sock.send_to(b"metrique=1", "10.0.0.9:8125")?; // ex : DogStatsD
    Ok(())
}

fn main() {
    println!("sonde : {}", sonde_tcp("10.0.0.1", 443, 500));
    let _ = envoi_udp();
    // mini_serveur().unwrap();
}
```

> Pour du HTTP, du TLS, de l'async ou des milliers de connexions : crates `reqwest`,
> `tokio`, `rustls` (§27, §45). La std reste parfaite pour sondes et prototypes.

---

## 27. HTTP client/serveur : l'écosystème (reqwest, axum)

```toml
[dependencies]
reqwest = { version = "0.12", features = ["json", "rustls-tls"] }
tokio = { version = "1", features = ["full"] }
serde = { version = "1", features = ["derive"] }
axum = "0.7"
```

```rust
use serde::{Deserialize, Serialize};

#[derive(Debug, Deserialize)]
struct Sante { statut: String }

#[derive(Serialize)]
struct Rapport { hote: String, ok: bool }

// Client HTTP (bloquant possible avec feature "blocking", async par défaut)
async fn verifier(url: &str) -> Result<bool, reqwest::Error> {
    let client = reqwest::Client::builder()
        .timeout(std::time::Duration::from_secs(5))
        .build()?;
    let rep: Sante = client.get(url).send().await?.json().await?;
    Ok(rep.statut == "ok")
}

// Serveur HTTP minimal avec axum
async fn lancer_serveur() {
    use axum::{routing::get, Json, Router};
    async fn sante() -> Json<Rapport> {
        Json(Rapport { hote: hostname::get().unwrap().to_string_lossy().into(), ok: true })
    }
    let app = Router::new().route("/sante", get(sante));
    let listener = tokio::net::TcpListener::bind("0.0.0.0:3000").await.unwrap();
    axum::serve(listener, app).await.unwrap();
}
```

> `reqwest` avec `rustls-tls` : TLS en pur Rust, pas de dépendance OpenSSL système.
> Binaire plus portable (musl !). `axum` = framework web moderne bâti sur `tokio`/`hyper`.

---

## 28. Tests unitaires : intégrés au langage

```rust
// Fonctions à tester
fn addition(a: i32, b: i32) -> i32 { a + b }
fn diviser(a: f64, b: f64) -> Option<f64> {
    if b == 0.0 { None } else { Some(a / b) }
}

// Les tests vivent dans le même fichier, module cfg(test)
#[cfg(test)]
mod tests {
    use super::*;                          // importe le code testé

    #[test]
    fn addition_basique() {
        assert_eq!(addition(2, 3), 5);
    }

    #[test]
    fn division_par_zero() {
        assert_eq!(diviser(1.0, 0.0), None);
    }

    #[test]
    #[should_panic(expected = "dépassement")]
    fn panique_attendue() {
        panic!("dépassement volontaire");
    }

    #[test]
    fn avec_resultat() -> Result<(), String> {   // test qui retourne Result : élégant
        if addition(2, 2) == 4 { Ok(()) } else { Err("mauvais calcul".into()) }
    }
}
```

```bash
cargo test                  # tous les tests
cargo test addition         # filtre par nom
cargo test -- --nocapture   # affiche les println!
cargo test --release        # en mode optimisé
```

Macros d'assertion :

| Macro | Usage |
|---|---|
| `assert!(cond)` | Condition booléenne |
| `assert_eq!(a, b)` / `assert_ne!` | Égalité (nécessite `PartialEq + Debug`) |
| `panic!("msg")` / `unreachable!()` / `unimplemented!()` | Échecs explicites |

---

## 29. Tests d'intégration et tests de documentation

```
mon_outil/
├── src/
│   └── lib.rs          # la lib (fonctions publiques)
├── src/main.rs         # le binaire (utilise la lib)
└── tests/
    └── integration.rs  # tests d'intégration : crate externe
```

```rust
// tests/integration.rs — compile comme une crate séparée
use mon_outil::addition;   // n'a accès qu'à l'API PUBLIQUE

#[test]
fn via_api_publique() {
    assert_eq!(addition(40, 2), 42);
}
```

```rust
// Test de documentation : le code dans /// est COMPILÉ et EXÉCUTÉ
/// Additionne deux entiers.
///
/// ```
/// assert_eq!(mon_outil::addition(2, 3), 5);
/// ```
pub fn addition(a: i32, b: i32) -> i32 { a + b }
```

> Les doctests garantissent que **tes exemples de doc restent vrais**. `cargo test` les lance.
> Stratégie : unitaires pour la logique interne, intégration pour le comportement public,
> doctests pour les exemples.

---

## 30. Debug : eprintln!, dbg!, gdb/lldb

```rust
fn main() {
    let config = vec!["a", "b", "c"];

    // dbg! : affiche fichier:ligne + valeur, REND la valeur (pratique en pipeline)
    let n = dbg!(config.len());            // [src/main.rs:4] config.len() = 3

    // eprintln! : vers stderr (ne pollue pas stdout = les données)
    eprintln!("DEBUG : {} éléments", config.len());

    // pretty-print
    eprintln!("{config:#?}");

    // traçage conditionnel via variable d'env (sans crate)
    if std::env::var("DEBUG").is_ok() {
        eprintln!("[debug] étape 1 ok");
    }
}
```

Debug natif :

```bash
# Symboles déjà inclus en debug. Lancer sous gdb :
rust-gdb ./target/debug/mon_outil
# ou lldb :
rust-lldb ./target/debug/mon_outil
```

> Pour du vrai logging structuré : crates `log` + `env_logger` (façade + implémentation,
> niveau via `RUST_LOG=debug`), ou `tracing` (async-aware, standard moderne).
> `dbg!` = provisoire ; ne jamais le laisser en prod (il écrit sur stderr).

---

## 31. Clippy et rustfmt : qualité automatique

```bash
cargo fmt --check        # vérifie le formatage (CI)
cargo fmt                # reformate tout
cargo clippy             # ~700 lints : maladresses, perfs, style
cargo clippy -- -D warnings   # les warnings deviennent des erreurs (CI stricte)
```

Exemples de ce que clippy détecte :

```rust
fn avant(hotes: &Vec<String>) -> usize {   // clippy: &Vec -> &[String] (ptr_arg)
    let mut total = 0;
    for h in hotes {                        // clippy ok ici, mais...
        total += h.len();
    }
    return total;                           // clippy: needless_return
}

fn apres(hotes: &[String]) -> usize {
    hotes.iter().map(|h| h.len()).sum()     // clippy: suggère l'itérateur
}
```

Fichier `clippy.toml` / attributs :

```rust
#![warn(clippy::pedantic)]        // en tête de crate : niveau exigeant
#[allow(clippy::too_many_arguments)]  // exception locale justifiée
fn f(a: u8, b: u8, c: u8, d: u8, e: u8, f2: u8, g: u8, h: u8) -> u8 {
    a + b + c + d + e + f2 + g + h
}
```

> Règle d'équipe : `cargo fmt --check` + `cargo clippy -- -D warnings` dans la CI.
> Zéro warning toléré sur le code neuf.

---

## 32. Bonnes pratiques & style : le contrat d'équipe

