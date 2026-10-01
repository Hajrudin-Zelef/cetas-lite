---
id: collect-261001-rattrapage/rattrapage/rust-guide-12
title: "Guide Rust ultra-complet — Langage système sûr, rapide et moderne"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["apache", "licenses"]
source: docs/RAG/collect-261001-rattrapage/rust_guide.md
source_anchor: ""
source_lines: [2617, 2878]
sha256: ee06acc1478167c2c126a6a246574c33cf4eee26a5f91569e01e0ff05e9b6dd0
---

# Guide Rust ultra-complet — Langage système sûr, rapide et moderne

fn main() -> std::io::Result<()> {
    // Lancer un processus et capturer sa sortie
    let sortie = Command::new("ip")
        .args(["-brief", "addr"])
        .output()?;                                   // attend + capture stdout/stderr
    println!("code : {}", sortie.status.code().unwrap_or(-1));
    println!("{}", String::from_utf8_lossy(&sortie.stdout));

    // Succès/échec explicite
    let statut = Command::new("ping")
        .args(["-c", "1", "-W", "1", "10.0.0.1"])
        .stdout(Stdio::null())
        .status()?;
    println!("ping ok : {}", statut.success());

    // Pipeline : echo hello | wc -c (via deux Command chaînés, ou shell -c si besoin)
    // Préférer l'API directe au shell pour éviter l'injection :
    let _ = Command::new("systemctl")
        .args(["is-active", "sshd"])
        .status()?;
    Ok(())
}
```

Signaux (Unix) : la std ne les gère pas → crate `tokio::signal` ou `signal-hook`.

```rust
// Avec tokio :
// tokio::signal::ctrl_c().await?;   // attend Ctrl-C proprement pour shutdown gracieux
```

> **Jamais** de `format!("ping {user_input}")` passé à `sh -c` : injection de commande.
> `Command::new` + `.args()` = pas d'interprétation shell, arguments sûrs par construction.

---

## 58. Temps, durées, sleep : std::time et chrono

```rust
use std::time::{Duration, Instant, SystemTime, UNIX_EPOCH};

fn main() {
    // Mesure d'intervalle (monotonique) : Instant
    let debut = Instant::now();
    std::thread::sleep(Duration::from_millis(120));
    println!("écoulé : {:?}", debut.elapsed());          // jamais affecté par NTP

    // Horloge murale : SystemTime (peut reculer !)
    let maintenant = SystemTime::now();
    let epoch = maintenant.duration_since(UNIX_EPOCH).unwrap().as_secs();
    println!("epoch : {epoch}");

    // Durées : arithmétique sûre
    let timeout = Duration::from_secs(5);
    let total = timeout.checked_add(Duration::from_secs(10)).unwrap();
    println!("{}s", total.as_secs());
}
```

```toml
[dependencies]
chrono = "0.4"
```

```rust
// Dates humaines : chrono (l'écosystème ; `time` est l'alternative)
fn exemple_chrono() {
    use chrono::{Local, Utc};
    let utc = Utc::now();
    println!("{}", utc.to_rfc3339());              // 2026-09-26T23:34:28+00:00
    println!("{}", Local::now().format("%d/%m/%Y %H:%M"));
}
```

> `Instant` pour mesurer, `SystemTime`/`chrono` pour horodater. Ne jamais calculer
> une durée avec l'horloge murale (NTP, DST, ajustements manuels).

---

## 59. Expressions régulières et parsing de texte

```toml
[dependencies]
regex = "1"
```

```rust
use regex::Regex;

fn main() {
    // Parser des lignes de log : extraire IP et code
    let re = Regex::new(r#"^(?P<ip>\d+\.\d+\.\d+\.\d+).*?"code":(?P<code>\d+)"#).unwrap();
    let ligne = r#"10.0.0.5 - - [26/Sep/2026] "GET /" "code":200"#;
    if let Some(caps) = re.captures(ligne) {
        println!("ip={} code={}", &caps["ip"], &caps["code"]);
    }

    // Valider un hostname (simple)
    let re_host = Regex::new(r"^[a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?$").unwrap();
    println!("valide : {}", re_host.is_match("srv-web-01"));

    // Splitter sur plusieurs séparateurs
    let re_sep = Regex::new(r"[;,\s]+").unwrap();
    let champs: Vec<&str> = re_sep.split("a;b, c  d").collect();
    println!("{champs:?}");
}
```

> Compile tes `Regex` **une fois** (coûteux) : `static` + `OnceLock` (std depuis 1.80),
> ou `lazy_static`/`once_cell` sur anciennes versions. Pour du parsing structuré
> (protocoles, formats), préfère `nom` ou `winnow` aux regex géantes.

---

## 60. Lecture/écriture binaire et endianness

```rust
fn main() {
    // Entiers <-> octets : endianness explicite (indispensable en réseau/protocoles)
    let n: u32 = 0x01020304;
    let be = n.to_be_bytes();      // big-endian (ordre réseau)
    let le = n.to_le_bytes();      // little-endian
    println!("be={be:?} le={le:?}");
    assert_eq!(u32::from_be_bytes(be), n);

    // Construire un paquet : en-tête fixe + payload
    let mut paquet = Vec::with_capacity(8);
    paquet.extend_from_slice(&0xCAFEu16.to_be_bytes());  // magic
    paquet.extend_from_slice(&(42u16).to_be_bytes());    // longueur
    paquet.extend_from_slice(b"data");                  // payload (4 octets)
    println!("paquet : {paquet:02X?}");

    // Parser : slices + from_be_bytes (jamais de cast sauvage)
    let magic = u16::from_be_bytes(paquet[0..2].try_into().unwrap());
    let len = u16::from_be_bytes(paquet[2..4].try_into().unwrap());
    println!("magic={magic:#X} len={len}");
}
```

> `try_into().unwrap()` sur slice de taille fixe = conversion vérifiée.
> Pour des protocoles complexes : crates `binrw` / `deku` (dérive déclarative).

---

## 61. Documentation : rustdoc comme un pro

```rust
/// Sonde TCP avec timeout.
///
/// # Arguments
///
/// * `hote` - IP ou nom DNS (ex : `"10.0.0.1"`).
/// * `port` - Port TCP (ex : `443`).
///
/// # Exemple
///
/// ```
/// assert!(mon_outil::sonder("127.0.0.1", 22));
/// ```
///
/// # Erreurs
///
/// Ne retourne pas d'erreur : `false` = injoignable ou timeout.
///
/// # Panics
///
/// Ne panique jamais.
pub fn sonder(hote: &str, port: u16) -> bool {
    !hote.is_empty() && port != 0
}
```

Sections conventionnelles : `# Arguments`, `# Returns`, `# Errors`, `# Panics`,
`# Safety` (si unsafe), `# Examples`. Le code des exemples = **doctests** (§29).

```bash
cargo doc --open --no-deps     # doc de TA crate uniquement
```

> Une fonction publique sans `///` = dette. `cargo doc` + `#![warn(missing_docs)]`
> en tête de lib pour l'exiger.

---

## 62. Édition 2024 en pratique : ce qui change vraiment ton code

Au-delà du tableau §49, les impacts concrets :

```rust
// 1. Captures de closures précises (édition 2024)
// Avant (2021) : la closure capturait `config` ENTIER par référence...
// Après (2024) : elle ne capture que `config.port` -> moins de conflits d'emprunt
struct Config { port: u16, hote: String }
fn exemple_ed2024() {
    let mut config = Config { port: 8080, hote: "x".into() };
    let lire_port = || config.port;      // 2024 : capture config.port uniquement
    println!("{}", lire_port());
    config.hote = "y".into();            // ✅ OK en 2024 (conflit en 2021 !)
}

// 2. `gen` est réservé : utiliser r#gen si besoin
// let r#gen = 5;

// 3. unsafe extern pour les FFI (édition 2024)
unsafe extern "C" {
    fn getpid() -> i32;
}
```

> Migration 2021 → 2024 : `cargo fix --edition`, puis `cargo test`.
> Les gains (captures précises) valent le coup sur du code concurrent.

---

## 63. Dépendances : choisir, auditer, verrouiller

```bash
cargo tree                    # qui dépend de quoi
cargo tree -i serde            # qui amène serde ?
cargo tree --depth 1
cargo outdated                 # versions plus récentes (cargo install cargo-outdated)
```

`Cargo.lock` :

- **Binaires** : committer TOUJOURS (reproductibilité du déploiement).
- **Bibliothèques** : ne PAS committer (laisser le résolveur choisir ; convention crates.io).

Sécurité supply chain :

```bash
cargo install cargo-audit cargo-deny
cargo audit                   # CVE connues dans tes dépendances
cargo deny check               # licences + advisories + sources (CI)
```

```toml
# deny.toml minimal
[advisories]
vulnerability = "deny"
[licenses]
allow = ["MIT", "Apache-2.0", "BSD-3-Clause", "ISC"]
```

> Checklist avant d'ajouter une crate : maintenance active ? téléchargements ?
> `unsafe` dedans (`cargo geiger`) ? licence compatible ? alternative dans la std ?

---

## 64. Cross-compilation et release : livrer un binaire

```bash
# Cibles utiles au sysadmin
rustup target add x86_64-unknown-linux-musl      # Linux statique (containers scratch)
rustup target add aarch64-unknown-linux-gnu      # ARM64 (Raspberry Pi, serveurs ARM)
rustup target add x86_64-pc-windows-gnu          # Windows depuis Linux

cargo build --release --target x86_64-unknown-linux-musl
```

`Cargo.toml` : profil release soigné :

