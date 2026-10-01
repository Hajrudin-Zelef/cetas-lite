---
id: collect-261001-rattrapage/rattrapage/rust-guide-10
title: "Guide Rust ultra-complet — Langage système sûr, rapide et moderne"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["agent"]
source: docs/RAG/collect-261001-rattrapage/rust_guide.md
source_anchor: ""
source_lines: [2093, 2336]
sha256: ec76c298f3a5bba5978b5142ee8c06da84a83d8d4f81994a510b94a934721a6c
---

# Guide Rust ultra-complet — Langage système sûr, rapide et moderne

L'async Rust = concurrence **coopérative** sur peu de threads (idéal I/O : milliers de
connexions, sondes réseau, scrapers). `async fn` ne fait rien tant qu'on ne l'**exécute pas**
dans un runtime (`tokio` = standard de facto).

```toml
[dependencies]
tokio = { version = "1", features = ["full"] }
```

```rust
use std::time::Duration;

// async fn = fonction qui retourne une Future (paresseuse !)
async fn sonder_async(hote: &str) -> bool {
    // Ici on simule ; en vrai : tokio::net::TcpStream::connect(...).await
    tokio::time::sleep(Duration::from_millis(50)).await;  // .await = point de suspension
    !hote.is_empty()
}

#[tokio::main]                       // macro : construit le runtime et bloque dessus
async fn main() {
    // Exécution séquentielle
    println!("a : {}", sonder_async("10.0.0.1").await);

    // Exécution CONCURRENTE : join! attend les deux
    let (r1, r2) = tokio::join!(
        sonder_async("10.0.0.1"),
        sonder_async("10.0.0.2"),
    );
    println!("{r1} {r2}");

    // spawn : tâche détachée (comme un thread, mais légère)
    let tache = tokio::spawn(sonder_async("10.0.0.3"));
    println!("spawn : {}", tache.await.unwrap());

    // Timeout : combinateur indispensable en supervision
    let res = tokio::time::timeout(Duration::from_millis(10), sonder_async("x")).await;
    println!("timeout écoulé : {}", res.is_err());
}
```

| Concept | À retenir |
|---|---|
| `Future` | Calcul paresseux, pollé par le runtime |
| `.await` | Suspend la tâche SANS bloquer le thread |
| `tokio::spawn` | Tâche `'static` (posséder ses données, comme les threads) |
| `tokio::join!` | Attendre plusieurs futures en parallèle |
| `tokio::select!` | Premier arrivé gagne (course, timeout, annulation) |
| `Send` across `.await` | Les données conservées entre deux `.await` doivent être `Send` |

> ⚠️ **Ne jamais bloquer** (`std::thread::sleep`, I/O bloquante) dans du code async :
> ça fige tout le runtime. Utiliser `tokio::time::sleep`, `tokio::fs`, `tokio::task::spawn_blocking`
> pour le code bloquant legacy.

---

## 47. FFI et binaires système : parler au C et à l'OS

```rust
use std::ffi::{CStr, CString};
use std::os::raw::c_char;

// Déclarer une fonction C existante
extern "C" {
    fn strlen(s: *const c_char) -> usize;
}

fn longueur_c(texte: &str) -> usize {
    let c_texte = CString::new(texte).expect("pas de \\0 intérieur");
    // SAFETY: c_texte est une C-string valide terminée par \0.
    unsafe { strlen(c_texte.as_ptr()) }
}

// Exposer une fonction Rust au C
#[no_mangle]                          // garde le nom du symbole
pub extern "C" fn addition_c(a: i32, b: i32) -> i32 {
    a + b
}

fn main() {
    println!("len C : {}", longueur_c("hello"));
    // Lire une chaîne C (ex : retour d'API) :
    // let s: &str = unsafe { CStr::from_ptr(ptr).to_str().unwrap() };
}
```

Types de pont :

| Rust | C | Note |
|---|---|---|
| `CString` / `CStr` | `char*` nul-terminé | `CString::new` échoue si `\0` intérieur |
| `*const T` / `*mut T` | pointeurs bruts | Pas de garantie : `unsafe` pour déréférencer |
| `std::os::raw::{c_int, c_char…}` | types C | Tailles selon plateforme |
| `#[repr(C)] struct` | `struct` C | Layout compatible garanti |

Binaire système portable :

```bash
rustup target add x86_64-unknown-linux-musl
cargo build --release --target x86_64-unknown-linux-musl
file target/x86_64-unknown-linux-musl/release/mon_outil   # statically linked
```

> Crate `libc` : bindings bruts des appels système. Pour du code portable, préfère la std
> (`std::process`, `std::fs`) ou des crates (`nix` pour l'Unix idiomatique).

---

## 48. Les erreurs du borrow checker les plus fréquentes (décryptées)

### E0502 — `cannot borrow as mutable because it is also borrowed as immutable`

```rust
let mut v = vec![1, 2, 3];
let r = &v;
// v.push(4);            // ❌ E0502 : r emprunte encore v
println!("{r}");         // dernier usage de r
v.push(4);               // ✅ OK : r est mort (NLL)
```

### E0505 — `cannot move out because it is borrowed`

```rust
let v = vec![String::from("a")];
let r = &v[0];
// let s = v;            // ❌ E0505 : v est emprunté par r
println!("{r}");
let s = v;               // ✅ après dernier usage de r
println!("{} {}", s[0], r.len());  // ❌ r est mort ici en fait... voir correction :
// ✅ version propre :
let v2 = vec![String::from("a")];
let s2 = v2;             // move direct, sans emprunt actif
println!("{}", s2[0]);
```

### E0515 — `cannot return reference to local variable`

```rust
// ❌ fn bad() -> &str { let s = String::from("x"); &s }  // s meurt à la fin !
 // ✅ retourner une valeur possédée :
fn good() -> String {
    String::from("x")
}
// ✅ ou emprunter un paramètre (lifetime) :
fn premier<'a>(s: &'a str) -> &'a str { s }
```

### E0382 — `borrow of moved value` (le classique)

```rust
let s = String::from("hello");
prend(s);                            // move dans la fonction
// println!("{s}");                   // ❌ E0382
fn prend(_s: String) {}
// ✅ passer &s si la fonction n'a pas besoin de posséder :
fn emprunte(s: &str) { println!("{s}"); }
let s2 = String::from("hello");
emprunte(&s2);
println!("{s2}");                    // ✅ toujours valide
```

### E0499 — `cannot borrow as mutable more than once`

```rust
let mut x = 5;
let a = &mut x;
// let b = &mut x;       // ❌ E0499
*a += 1;
println!("{a}");         // dernier usage de a
let b = &mut x;          // ✅ OK maintenant
*b += 1;
```

### E0597 — `does not live long enough`

```rust
// ❌ let r: &str; { let s = String::from("x"); r = &s; }  // s meurt avant r
// ✅ faire vivre le propriétaire assez longtemps :
let s = String::from("x");
let r: &str = &s;
println!("{r}");
```

> Méthode de debug : lis le **dernier usage** de chaque emprunt. 90 % des erreurs E05xx
> se résolvent en **réordonnant** le code (utiliser l'emprunt avant la mutation) ou en
> **clonant** explicitement quand le partage de propriété est voulu (`Rc`/`Arc`).

---

## 49. Éditions 2021 et 2024 : ce qui change

| Sujet | Édition 2021 | Édition 2024 (Rust 1.85+) |
|---|---|---|
| `cargo new` | Édition 2021 par défaut | Préciser `edition = "2024"` dans Cargo.toml |
| Closures | Capturent les variables entières | Capturent **champ par champ** (plus précis, moins d'erreurs d'emprunt) |
| `gen` | Identifiant normal | Mot-clé réservé (générateurs) → renommer ou `r#gen` |
| `unsafe extern` | `extern` simple | `unsafe extern` requis pour déclarer des FFI |
| Résolution `match` ergonomique | Déjà là (2018) | Inchangée, mais match sur `&Option` encore adouci |
| Prelude | `TryInto`, `FromIterator`… | Ajouts : `Future`, `IntoFuture` dans le prelude |

Choisir dans `Cargo.toml` :

```toml
[package]
edition = "2024"
```

> Les éditions **ne cassent jamais** le code existant : un projet 2018 compile avec un
> compilateur 2024. Migrer = `cargo fix --edition` + revue manuelle. Pour un projet neuf
> en 2026 : **édition 2024** directement.

---

## 50. Workspaces : plusieurs crates, un seul build

```toml
# Cargo.toml à la racine
[workspace]
members = ["agent", "exporter", "commun"]
resolver = "2"          # résolveur moderne (recommandé depuis 2021)
```

```
mon_projet/
├── Cargo.toml          # [workspace]
├── Cargo.lock          # unique pour tout le workspace
├── target/             # build partagé
├── agent/
│   ├── Cargo.toml      # name = "agent" ; commun = { path = "../commun" }
│   └── src/main.rs
├── exporter/
│   ├── Cargo.toml
│   └── src/main.rs
└── commun/
    ├── Cargo.toml      # [lib]
    └── src/lib.rs
```

```bash
cargo build -p agent            # compiler un membre
cargo run -p exporter -- --help
cargo test --workspace          # tout tester
```

