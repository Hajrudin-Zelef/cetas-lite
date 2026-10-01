---
id: collect-261001-rattrapage/rattrapage/rust-guide-11
title: "Guide Rust ultra-complet — Langage système sûr, rapide et moderne"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["agent", "incident"]
source: docs/RAG/collect-261001-rattrapage/rust_guide.md
source_anchor: ""
source_lines: [2337, 2616]
sha256: 2944516572539c9e7050f4a4dcd76464b3e00a38206825f9008ae3196f82b7a5
---

# Guide Rust ultra-complet — Langage système sûr, rapide et moderne

> Dépendance `path` entre membres = pas de version à gérer, recompilation incrémentale.
> Idéal : binaire agent + lib commune + exporter, tous versionnés ensemble.

---

## 51. Features Cargo : compilation conditionnelle propre

```toml
# Cargo.toml de la crate
[features]
default = ["json"]
json = ["dep:serde_json"]        # syntaxe moderne : dépendance optionnelle
yaml = ["dep:serde_yaml"]
tls = []

[dependencies]
serde_json = { version = "1", optional = true }
serde_yaml = { version = "0.9", optional = true }
```

```rust
// src/lib.rs
#[cfg(feature = "json")]
pub fn vers_json<T: serde::Serialize>(v: &T) -> String {
    serde_json::to_string(v).unwrap()
}

#[cfg(feature = "tls")]
pub fn client_tls() -> &'static str { "client TLS activé" }
```

```bash
cargo build                        # features default
cargo build --no-default-features   # sans rien
cargo build --features yaml,tls     # sélection
cargo build --all-features          # tout (CI)
```

> Les features doivent être **additives** : activer une feature ne doit jamais casser
> un code qui compile sans elle. Convention respectée par tout l'écosystème.

---

## 52. Macros : `macro_rules!` et derive procédurales

```rust
// Macro déclarative : pattern matching sur du code
macro_rules! hashmap {
    ( $( $cle:expr => $val:expr ),* $(,)? ) => {{
        let mut m = std::collections::HashMap::new();
        $( m.insert($cle, $val); )*
        m
    }};
}

fn main() {
    let seuils = hashmap! { "cpu" => 90, "disque" => 85, "ram" => 95 };
    println!("seuil cpu : {}", seuils["cpu"]);
}
```

| Type de macro | Usage | Exemples connus |
|---|---|---|
| Déclarative (`macro_rules!`) | Petits DSL, répétitions | `vec!`, `hashmap!` maison |
| Derive procédurale | Génère du code depuis un struct | `#[derive(Serialize)]`, `Parser` de clap |
| Attribut procédural | Transforme une fonction | `#[tokio::main]`, `#[test]` |
| Fonction procédurale | Appel `ma_macro!(...)` | `sqlx::query!` (vérifié à la compilation !) |

> Règle : **n'écris pas de macro** tant qu'une fonction générique + trait suffit.
> Les macros procédurales = crates séparées (`proc-macro = true`), à réserver aux cas
> où la génération de code apporte une vraie valeur (serde, clap, sqlx).

---

## 53. Sérialisation : serde en profondeur

```rust
use serde::{Deserialize, Serialize};

#[derive(Debug, Serialize, Deserialize)]
struct Interface {
    nom: String,
    #[serde(rename = "adresseIP")]          // nom différent dans le JSON
    ip: String,
    #[serde(default)]                       // absent -> Default::default()
    mtu: u16,
    #[serde(skip_serializing_if = "Option::is_none")]
    description: Option<String>,            // omis si None
    #[serde(skip)]                          // jamais sérialisé
    #[serde(default)]
    cache_interne: String,
}

#[derive(Debug, Serialize, Deserialize)]
#[serde(tag = "type")]                      // enum à tag interne
enum EquipementReseau {
    Switch { ports: u16 },
    Routeur { bgp: bool },
}
// JSON : {"type":"Switch","ports":48}

#[derive(Debug, Deserialize)]
struct ConfigSouple {
    #[serde(default)]
    hotes: Vec<String>,
    #[serde(flatten)]                       // champs inconnus capturés
    extra: std::collections::HashMap<String, serde_json::Value>,
}

fn main() -> Result<(), serde_json::Error> {
    let json = r#"{"nom":"eth0","adresseIP":"10.0.0.1","description":"uplink"}"#;
    let itf: Interface = serde_json::from_str(json)?;
    println!("{itf:?}");                    // mtu=0 (défaut), cache_interne=""
    let eq = EquipementReseau::Switch { ports: 48 };
    println!("{}", serde_json::to_string(&eq)?);
    Ok(())
}
```

> `#[serde(deny_unknown_fields)]` : strict sur les configs (détecte les fautes de frappe).
> `flatten` + `HashMap<String, Value>` : tolérant sur les formats qui évoluent.

---

## 54. CLI sérieuses avec clap (derive)

```toml
[dependencies]
clap = { version = "4", features = ["derive"] }
anyhow = "1"
```

```rust
use clap::{Parser, Subcommand};

/// Sonde réseau parallèle — exemple d'outil sysadmin
#[derive(Parser, Debug)]
#[command(name = "sonde", version, about = "Sonde TCP multi-hôtes")]
struct Cli {
    /// Hôtes à sonder (IP ou nom)
    #[arg(required = true)]
    hotes: Vec<String>,

    /// Port TCP à tester
    #[arg(short, long, default_value_t = 443)]
    port: u16,

    /// Timeout en millisecondes
    #[arg(long, default_value_t = 1000)]
    timeout_ms: u64,

    /// Sortie JSON
    #[arg(long)]
    json: bool,

    #[command(subcommand)]
    commande: Option<Commandes>,
}

#[derive(Subcommand, Debug)]
enum Commandes {
    /// Liste les hôtes sans sonder
    Lister,
}

fn main() -> anyhow::Result<()> {
    let cli = Cli::parse();   // --help et --version générés automatiquement
    println!("{cli:?}");
    Ok(())
}
```

```bash
cargo run -- --help
cargo run -- 10.0.0.1 10.0.0.2 -p 22 --timeout-ms 500 --json
cargo run -- lister --help
```

> clap v4 derive : un struct annoté = un CLI complet (help, erreurs, complétion shell
> via `clap_complete`). Pour les scripts internes simples, `std::env::args` suffit (§25).

---

## 55. Logging et tracing en production

```toml
[dependencies]
tracing = "0.1"
tracing-subscriber = { version = "0.3", features = ["env-filter", "json"] }
```

```rust
use tracing::{debug, error, info, instrument, warn};

#[instrument]                                  // span auto : nom + args de la fonction
fn sonder(hote: &str, port: u16) -> bool {
    debug!(hote, port, "début de sonde");
    let ok = !hote.is_empty() && port != 0;
    if ok { info!(hote, "joignable"); } else { warn!(hote, "injoignable"); }
    ok
}

fn main() {
    // Format JSON + niveau via RUST_LOG (ex : RUST_LOG=debug)
    tracing_subscriber::fmt()
        .json()
        .with_env_filter("sonde=info")
        .init();

    if !sonder("10.0.0.1", 443) {
        error!("sonde critique échouée");
        std::process::exit(1);
    }
}
```

| Niveau | Usage |
|---|---|
| `error!` | Panne nécessitant une action |
| `warn!` | Dégradation, à surveiller |
| `info!` | Événements métier normaux |
| `debug!` | Détail pour diagnostiquer |
| `trace!` | Très verbeux (boucles, paquets) |

> `tracing` > `log` pour l'async (spans suivent les tâches). En JSON + journald/ELK :
> tes logs deviennent requêtables. `RUST_LOG` en prod, jamais de niveau codé en dur.

---

## 56. Configuration : layered config (fichier + env + CLI)

Ordre de priorité idiomatique : **CLI > variables d'env > fichier > défauts**.

```toml
[dependencies]
config = "0.14"
serde = { version = "1", features = ["derive"] }
```

```rust
use serde::Deserialize;

#[derive(Debug, Deserialize)]
struct Config {
    #[serde(default = "defaut_hote")]
    hote: String,
    #[serde(default = "defaut_port")]
    port: u16,
    #[serde(default)]
    verbose: bool,
}
fn defaut_hote() -> String { "127.0.0.1".to_string() }
fn defaut_port() -> u16 { 8080 }

fn charger() -> Result<Config, config::ConfigError> {
    config::Config::builder()
        .add_source(config::File::with_name("/etc/sonde/config").required(false))
        .add_source(config::Environment::with_prefix("SONDE"))  // SONDE_PORT=9090
        .build()?
        .try_deserialize()
}

fn main() {
    match charger() {
        Ok(c) => println!("config : {c:?}"),
        Err(e) => { eprintln!("config invalide : {e}"); std::process::exit(2); }
    }
}
```

> Valide la config **au démarrage** et échoue vite (fail fast) avec un message qui dit
> QUELLE clé pose problème. Un daemon qui démarre à moitié configuré = incident garanti.

---

## 57. Gestion des processus : spawn, signaux, exit codes

```rust
use std::process::{Command, Stdio};

