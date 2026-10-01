---
id: collect-261001-rattrapage/rattrapage/rust-guide-5
title: "Guide Rust ultra-complet — Langage système sûr, rapide et moderne"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["consumer"]
source: docs/RAG/collect-261001-rattrapage/rust_guide.md
source_anchor: ""
source_lines: [924, 1171]
sha256: e9e7abbbc7c2a673431640ba88679e5949c8b6c2f8049b8b9636fd0cf00a405b
---

# Guide Rust ultra-complet — Langage système sûr, rapide et moderne

```rust
fn main() {
    let logs = vec![
        "INFO démarrage ok",
        "WARN disque 82%",
        "ERROR connexion refusée",
        "INFO cycle terminé",
        "ERROR timeout 10.0.0.5",
    ];

    // Pipeline : iterator adaptors (paresseux) + consumer (déclenche)
    let erreurs: Vec<&str> = logs
        .iter()                       // -> impl Iterator<Item = &&str>
        .map(|l| l.trim())            // adaptateur : transforme
        .filter(|l| l.starts_with("ERROR"))  // adaptateur : filtre
        .collect();                   // consommateur : exécute et collecte
    println!("{erreurs:?}");

    // Autres consommateurs courants
    let nb_erreurs = logs.iter().filter(|l| l.starts_with("ERROR")).count();
    let total: usize = logs.iter().map(|l| l.len()).sum();
    let max = logs.iter().max_by_key(|l| l.len());

    // Chaînage, zip, enumerate, take/skip
    let pairs: Vec<(usize, &&str)> = logs.iter().enumerate().collect();
    let premiers: Vec<&&str> = logs.iter().take(2).collect();

    // into_iter() consomme le Vec (prend possession des éléments)
    let possedes: Vec<String> = vec!["a".to_string(), "b".to_string()]
        .into_iter()
        .map(|s| s.to_uppercase())
        .collect();
    println!("{nb_erreurs} {total} {max:?} {pairs:?} {premiers:?} {possedes:?}");
}
```

| Adaptateur (paresseux) | Consommateur (exécute) |
|---|---|
| `map`, `filter`, `filter_map` | `collect`, `sum`, `count` |
| `take`, `skip`, `enumerate` | `for_each`, `fold`, `reduce` |
| `zip`, `chain`, `flat_map` | `find`, `any`, `all`, `position` |
| `map_while`, `take_while` | `min/max`, `min_by_key`, `partition` |

> `collect()` a besoin du type cible : `let v: Vec<_> = ...` ou `let s: String = ...`.
> Le `_` laisse l'inférence faire le reste.

---

## 22. Gestion d'erreurs avancée : thiserror et anyhow

Deux crates complémentaires (écosystème standard de facto) :

| Crate | Rôle | Quand |
|---|---|---|
| `thiserror` | **Définir** des types d'erreur propres (bibliothèques) | Tu exposes une API |
| `anyhow` | **Propager** des erreurs avec contexte (applications) | Ton binaire, ton CLI |

```toml
[dependencies]
thiserror = "2"
anyhow = "1"
```

```rust
use thiserror::Error;

// -- Côté bibliothèque : erreur typée, exhaustive --
#[derive(Error, Debug)]
pub enum InventaireErreur {
    #[error("hôte injoignable : {0}")]
    Injoignable(String),
    #[error("fichier {chemin} illisible")]
    FichierIllisible { chemin: String },
    #[error(transparent)]                 // délègue le Display à l'erreur source
    Io(#[from] std::io::Error),           // #[from] génère le From<> pour `?`
}

// -- Côté application : anyhow + contexte --
use anyhow::{Context, Result};

fn charger_inventaire(chemin: &str) -> Result<Vec<String>> {
    let contenu = std::fs::read_to_string(chemin)
        .with_context(|| format!("lecture de l'inventaire {chemin}"))?;
    Ok(contenu.lines().map(|l| l.to_string()).collect())
}

fn main() -> Result<()> {                  // main peut retourner Result !
    let hotes = charger_inventaire("/etc/inventaire.txt")?;
    println!("{} hôtes chargés", hotes.len());
    Ok(())
}
```

Sortie d'erreur avec contexte :

```
Error: lecture de l'inventaire /etc/inventaire.txt

Caused by:
    No such file or directory (os error 2)
```

> `fn main() -> Result<()>` : si `Err`, Rust affiche l'erreur via `Debug` et sort avec code 1.
> Avec `anyhow`, tu obtiens la chaîne de contextes. Pour un CLI soigné, ajoute une crate
> comme `miette` (rapports d'erreur enrichis).

---

## 23. Modules et crates : organiser un projet

```rust
// src/main.rs
mod inventaire;          // charge src/inventaire.rs (ou src/inventaire/mod.rs)
mod reseau;

use inventaire::Hote;    // import dans la portée
use reseau::ping;

fn main() {
    let h = Hote::new("srv-01");
    println!("{} joignable : {}", h.nom, ping("10.0.0.1"));
}
```

```rust
// src/inventaire.rs
pub struct Hote {         // `pub` : visible hors du module
    pub nom: String,      // champ public
    ip: String,           // champ privé au module
}

impl Hote {
    pub fn new(nom: &str) -> Self {
        Self { nom: nom.to_string(), ip: String::new() }
    }
    pub(crate) fn definir_ip(&mut self, ip: &str) {  // visible dans toute la crate
        self.ip = ip.to_string();
    }
}
```

Visibilité, du plus restreint au plus large :

| Mot-clé | Portée |
|---|---|
| (privé, défaut) | Module courant + descendants |
| `pub(crate)` | Toute la crate |
| `pub(super)` | Module parent |
| `pub(in crate::chemin)` | Chemin précis |
| `pub` | Partout (API publique) |

> Une **crate** = une unité de compilation (un projet). Un **package** Cargo peut contenir
> plusieurs crates (binaire + lib). `crates.io` = le registre public (~150 000 crates).

---

## 24. Entrées-sorties et fichiers : le terrain du sysadmin

```rust
use std::fs;
use std::io::{self, BufRead, BufReader, Write};
use std::path::{Path, PathBuf};

fn main() -> io::Result<()> {
    // --- Lecture/écriture "tout d'un coup" (fichiers petits/moyens) ---
    let contenu = fs::read_to_string("/etc/hostname")?;
    fs::write("/tmp/test.txt", "hello")?;

    // --- Lecture ligne à ligne (logs !) : BufReader indispensable ---
    let fichier = fs::File::open("/var/log/syslog")?;
    let lecteur = BufReader::new(fichier);
    let mut nb_warn = 0;
    for ligne in lecteur.lines() {
        let ligne = ligne?;                       // chaque ligne est un io::Result
        if ligne.contains("WARN") { nb_warn += 1; }
    }

    // --- Écriture bufferisée ---
    let mut sortie = fs::File::create("/tmp/rapport.txt")?;
    writeln!(sortie, "avertissements : {nb_warn}")?;

    // --- Parcours de répertoire ---
    for entree in fs::read_dir("/etc")? {
        let entree = entree?;
        let chemin = entree.path();
        if chemin.extension().and_then(|e| e.to_str()) == Some("conf") {
            println!("{}", chemin.display());
        }
    }

    // --- Path / PathBuf : manipuler sans String ---
    let mut p = PathBuf::from("/etc");
    p.push("mon_outil");                          // gère les séparateurs selon l'OS
    p.set_extension("conf");
    println!("{} existe={}", p.display(), Path::new(&p).exists());
    Ok(())
}
```

> `writeln!` vers un fichier retourne un `Result` : ne l'ignore pas (`?` ou `let _ =`).
> Pour des logs volumineux, `BufReader` + `.lines()` évite de charger des Go en RAM.

---

## 25. Arguments CLI, variables d'environnement, codes de sortie

```rust
use std::env;
use std::process::ExitCode;

fn main() -> ExitCode {
    // --- Arguments : env::args() (le 1er = nom du programme) ---
    let args: Vec<String> = env::args().collect();
    // let args: Vec<String> = env::args().skip(1).collect();  // sans le nom

    if args.len() < 2 {
        eprintln!("usage : {} <fichier> [--verbose]", args[0]);
        return ExitCode::from(2);                  // 2 = erreur d'usage (convention)
    }

    // --- Variables d'environnement ---
    let verbose = env::var("VERBOSE").is_ok();     // présence
    let timeout: u64 = env::var("TIMEOUT_MS")      // valeur parsée
        .ok()
        .and_then(|v| v.parse().ok())
        .unwrap_or(1000);
    // env::var("OBLIGATOIRE").expect("...")       // panic si absente

    println!("fichier={} verbose={verbose} timeout={timeout}", args[1]);
    ExitCode::SUCCESS                              // 0
}
```

> Pour un vrai CLI (sous-commandes, `--help` auto, validation), utilise **clap**
> (dérive : `#[derive(Parser)]`). `env::args` suffit pour les outils internes simples.
> Convention Unix : stdout = données, stderr = logs/erreurs, code 0 = succès.

---

## 26. Réseau : TCP/UDP avec la std

La std couvre TCP/UDP bloquant — suffisant pour sondes, petits serveurs, outils.

```rust
use std::io::{Read, Write};
use std::net::{TcpListener, TcpStream, UdpSocket};
use std::time::Duration;

