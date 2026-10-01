---
id: collect-261001-rattrapage/rattrapage/rust-guide-1
title: "Guide Rust ultra-complet — Langage système sûr, rapide et moderne"
domain: rattrapage
role: reference
task: reference
actors: []
dates: ["2024-10-17"]
keywords: ["agents"]
source: docs/RAG/collect-261001-rattrapage/rust_guide.md
source_anchor: ""
source_lines: [1, 200]
sha256: 7d9ed8f87060759b156f1460a626241decae5f6de58a75c787f91a76817168f1
---

# Guide Rust ultra-complet — Langage système sûr, rapide et moderne

> **Pour qui ?** Sysadmins et chefs de service systèmes & énergies qui veulent écrire des outils
> système fiables, des CLI performantes et des binaires portables sans sacrifier la sécurité mémoire.
> **Niveau :** de zéro à autonome sur projets réels. **Langue :** français.
> **Édition Rust visée :** 2021 et 2024 (les différences sont signalées explicitement).
> **Ton :** direct, tutoriel + référence, dense. Zéro remplissage.

---

## 1. Pourquoi Rust quand on est sysadmin

- **Binaires uniques, statiques par défaut** : un outil compilé en Rust s'exécute sur une machine
  cible sans runtime, sans interpréteur, sans dépendances dynamiques (ou presque). Idéal pour
  déployer sur des serveurs, des appliances, des conteneurs minimaux.
- **Vitesse du C, sécurité en plus** : le *borrow checker* garantit à la compilation l'absence
  de *use-after-free*, de double `free`, de data races. Fini les segfaults en prod à 3h du matin.
- **Gestion d'erreurs explicite** : pas d'exceptions qui surgissent de nulle part.
  `Result<T, E>` force à traiter les pannes (disque plein, réseau coupé) là où elles arrivent.
- **Excellent outillage** : `cargo` (build, dépendances, tests, doc, publish) fait en une commande
  ce qui demande ailleurs un Makefile + pip + venv + pytest + sphinx.
- **Cas d'usage typiques** : CLI d'inventaire, agents de supervision, exporters Prometheus,
  outils de sauvegarde, parsers de logs, daemons réseau, utilitaires de migration.

> Pense-bête : si ton script Python devient critique, lent ou fragile → candidat à la réécriture en Rust.

---

## 2. Installation avec rustup (la seule méthode recommandée)

Ne passe JAMAIS par le gestionnaire de paquets de la distro en premier choix (`apt install rustc`
fournit souvent une version ancienne). Utilise **rustup**, le gestionnaire de chaînes d'outils officiel.

```bash
# Linux / macOS
curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh
# Choisir l'option 1 (installation par défaut), puis :
source "$HOME/.cargo/env"

# Vérification
rustc --version   # ex : rustc 1.82.0 (f20d54c8a 2024-10-17)
cargo --version
rustup --version
```

```powershell
# Windows : télécharger rustup-init.exe depuis https://rustup.rs
# puis dans PowerShell :
rustc --version
cargo --version
```

### Toolchains : stable, beta, nightly

| Canal | Usage |
|---|---|
| `stable` | Production, par défaut. Mises à jour toutes les 6 semaines. |
| `beta` | Prévisualisation de la prochaine stable. |
| `nightly` | Fonctionnalités expérimentales (`#![feature(...)]`). Instable par définition. |

```bash
rustup toolchain install nightly          # ajouter nightly
rustup default stable                      # toolchain par défaut
rustup override set nightly                # nightly pour le dossier courant
rustup update                              # tout mettre à jour
rustup component add rustfmt clippy        # formateur + linter (indispensables)
rustup component add rust-analyzer         # serveur de langage (IDE)
rustup target add x86_64-unknown-linux-musl  # cible pour binaires 100 % statiques
```

> **Édition** : depuis Rust 1.56 (édition 2021), `cargo new` crée des projets en édition 2021.
> L'édition 2024 est stable depuis Rust 1.85. On la choisit explicitement (voir §62).

---

## 3. Premier programme et anatomie d'un projet Cargo

```bash
cargo new mon_outil        # crée mon_outil/ (binaire)
cargo new --lib ma_lib     # crée ma_lib/ (bibliothèque)
cd mon_outil
cargo run                  # compile + exécute
cargo build                # compile (debug)
cargo build --release      # compile optimisé
```

Arborescence générée :

```
mon_outil/
├── Cargo.toml      # manifeste : nom, version, édition, dépendances
├── Cargo.lock      # versions exactes résolues (à committer pour les binaires)
└── src/
    └── main.rs     # point d'entrée
```

`src/main.rs` minimal :

```rust
fn main() {
    println!("Hello, infra !");
}
```

`Cargo.toml` typique :

```toml
[package]
name = "mon_outil"
version = "0.1.0"
edition = "2021"          # ou "2024"

[dependencies]
# serde = "1.0"           # exemple de dépendance (crates.io)

[profile.release]
opt-level = 3
lto = true                # Link Time Optimization : binaire plus petit/rapide
strip = true              # supprime les symboles de debug
```

Checklist premier projet :

- [ ] `rustup component add rustfmt clippy` exécuté
- [ ] Édition choisie consciemment dans `Cargo.toml`
- [ ] `cargo run` fonctionne
- [ ] Éditeur configuré avec rust-analyzer

---

## 4. `cargo` en détail : le couteau suisse

| Commande | Rôle |
|---|---|
| `cargo new` / `cargo init` | Créer un projet |
| `cargo build` / `cargo build --release` | Compiler |
| `cargo run [--release] [-- <args>]` | Compiler + exécuter (args après `--`) |
| `cargo check` | Vérification rapide sans produire de binaire (itération rapide) |
| `cargo test` | Lancer les tests (unitaires, intégration, doc) |
| `cargo doc --open` | Générer/ouvrir la documentation |
| `cargo fmt` | Formater le code (rustfmt) |
| `cargo clippy` | Linter : détecte les maladresses |
| `cargo add <crate>` | Ajouter une dépendance (nécessite `cargo-edit`, ou édition manuelle) |
| `cargo update` | Mettre à jour les dépendances (respecte `Cargo.lock`) |
| `cargo tree` | Arbre des dépendances |
| `cargo audit` | Vulnérabilités connues (via `cargo install cargo-audit`) |
| `cargo install <crate>` | Installer un binaire Rust (ex : `cargo install ripgrep`) |
| `cargo clean` | Supprimer les artefacts de build |

```bash
cargo run -- --help              # passer des arguments au programme
cargo test -- --nocapture        # voir les println! dans les tests
cargo build --target x86_64-unknown-linux-musl   # binaire statique
```

> Astuce sysadmin : `cargo check` est 3 à 5× plus rapide que `cargo build` pendant le
> développement. Réserve `build --release` à la livraison.

---

## 5. Syntaxe de base : variables, mutabilité, constantes

```rust
fn main() {
    let x = 5;            // immutable par défaut (choix de sécurité)
    // x = 6;             // ERREUR : cannot assign twice to immutable variable
    let mut y = 5;        // mutable explicite
    y = 6;

    let z = 5;
    let z = z + 1;        // shadowing : redéclare, peut changer de type
    let z = "texte";      // légal grâce au shadowing

    const MAX_RETRY: u32 = 3;                 // constante : type obligatoire, SCREAMING_CASE
    static COMPTEUR: u32 = 0;                 // variable statique (rare ; voir Mutex pour mutabilité)
}
```

| Mot-clé | Caractéristique |
|---|---|
| `let` | Liaison immutable par défaut |
| `let mut` | Liaison mutable |
| Shadowing (`let x = ...` répété) | Nouvelle variable, type modifiable, portée limitée |
| `const` | Constante évaluée à la compilation, toujours typée, inlinée |
| `static` | Emplacement mémoire fixe, durée de vie `'static` |

---

## 6. Types primitifs : le tableau de référence

| Catégorie | Types | Détails |
|---|---|---|
| Entiers signés | `i8` `i16` `i32` `i64` `i128` `isize` | `i32` par défaut ; `isize` = taille d'un pointeur |
| Entiers non signés | `u8` `u16` `u32` `u64` `u128` `usize` | `usize` pour indexer les collections |
| Flottants | `f32` `f64` | `f64` par défaut (IEEE 754) |
| Booléen | `bool` | `true` / `false`, 1 octet |
| Caractère | `char` | 4 octets, Unicode scalaire (`'é'`, `'🦀'`) |
| Unité | `()` | Type « rien », retour des fonctions sans valeur |
| Jamais | `!` | Type « ne retourne jamais » (`panic!`, boucles infinies) |

