---
id: collect-261001-rattrapage/rattrapage/rust-guide-13
title: "Guide Rust ultra-complet — Langage système sûr, rapide et moderne"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["attention"]
source: docs/RAG/collect-261001-rattrapage/rust_guide.md
source_anchor: ""
source_lines: [2879, 3115]
sha256: 3beabf752b32d9a20fb20a28da4357e2223fc8d030827fea98b5f1d0135fc66d
---

# Guide Rust ultra-complet — Langage système sûr, rapide et moderne

```toml
[profile.release]
opt-level = 3
lto = true
codegen-units = 1
strip = true          # ou "debuginfo" pour garder les symboles séparés
panic = "abort"       # binaire plus petit (pas de unwinding) — attention : pas de catch
```

Dockerfile multi-stage typique :

```dockerfile
FROM rust:1.82 AS build
WORKDIR /app
COPY . .
RUN cargo build --release --target x86_64-unknown-linux-musl

FROM scratch
COPY --from=build /app/target/x86_64-unknown-linux-musl/release/mon_outil /mon_outil
ENTRYPOINT ["/mon_outil"]
```

> Binaire musl statique + image `scratch` = ~5 Mo, zéro CVE d'OS, déploiement trivial.
> Vérifie `ldd` / `file` après build pour confirmer le statique.

---

## 65. Déboguer le borrow checker : méthode pas à pas

Quand le compilateur dit non, applique cette procédure :

1. **Lis l'erreur EN ENTIER** : Rust donne la ligne, la cause, et souvent la solution
   (`help: ...`). Les E05xx sont les plus pédagogiques du marché.
2. **Identifie les 3 acteurs** : le propriétaire, l'emprunt actif, l'opération refusée.
3. **Trouve le dernier usage** de l'emprunt : peux-tu réordonner (utiliser avant de muter) ?
4. **Questionne le besoin** : faut-il vraiment partager ? Souvent on peut :
   - cloner explicitement (petites données),
   - passer par valeur puis retourner (transfert aller-retour),
   - restructurer (calculer d'abord, muter ensuite).
5. **Si partage réel** : `Rc<T>` (mono-thread) / `Arc<T>` (multi-thread) + `RefCell`/`Mutex`
   pour la mutabilité intérieure.
6. **Ne contourne JAMAIS** avec `unsafe` pour faire taire le borrow checker.

Exemple de restructuration (le pattern le plus rentable) :

```rust
// ❌ calcul + mutation entremêlés
// let mut cache = std::collections::HashMap::new();
// let v = cache.get("cle");          // emprunt immuable
// cache.insert("cle", 1);            // ❌ E0502
// println!("{v:?}");

// ✅ séparer : phase de lecture, puis phase d'écriture
let mut cache = std::collections::HashMap::new();
let existe = cache.contains_key("cle");   // lecture terminée (bool : Copy)
if !existe {
    cache.insert("cle", 1);               // écriture, plus d'emprunt actif
}
println!("{existe}");
```

---

## 66. Trait objects (`dyn Trait`) : dispatch dynamique

Quand le type concret n'est pas connu à la compilation (plugins, handlers hétérogènes) :

```rust
trait Sondable {
    fn sonder(&self) -> bool;
    fn nom(&self) -> &str;
}

struct Ping { hote: String }
struct Http { url: String }

impl Sondable for Ping {
    fn sonder(&self) -> bool { !self.hote.is_empty() }
    fn nom(&self) -> &str { &self.hote }
}
impl Sondable for Http {
    fn sonder(&self) -> bool { self.url.starts_with("https") }
    fn nom(&self) -> &str { &self.url }
}

fn main() {
    // Vec de types DIFFÉRENTS via trait object (taille connue via Box)
    let sondes: Vec<Box<dyn Sondable>> = vec![
        Box::new(Ping { hote: "10.0.0.1".into() }),
        Box::new(Http { url: "https://srv".into() }),
    ];
    for s in &sondes {
        println!("{} : {}", s.nom(), s.sonder());
    }
}
```

| | Génériques (`T: Trait`) | Trait objects (`dyn Trait`) |
|---|---|---|
| Dispatch | Statique (monomorphisation) | Dynamique (vtable) |
| Coût | Zéro à l'exécution | Indirection légère |
| Types hétérogènes en collection | Non | Oui |
| Taille connue à la compilation | Oui | Non → `Box`/`&` |

> Règle : génériques par défaut (perf) ; `dyn Trait` quand l'hétérogénéité l'exige.
> Un trait objet-safe = pas de méthodes génériques, pas de `Self` en retour.

---

## 67. Intériorité mutabilité : Cell, RefCell, OnceLock

Parfois la propriété est partagée (`&T`) mais on doit muter : **mutabilité intérieure**,
vérifiée à l'exécution plutôt qu'à la compilation.

```rust
use std::cell::{Cell, RefCell};
use std::sync::OnceLock;

// Cell<T: Copy> : get/set sans emprunt
let c = Cell::new(0);
c.set(c.get() + 1);

// RefCell<T> : emprunt vérifié à l'EXÉCUTION (panique si règles violées)
let rc = RefCell::new(vec![1, 2]);
rc.borrow_mut().push(3);          // panic si déjà emprunté (mut ou immut)
// RefCell n'est PAS Send/Sync : mono-thread uniquement !

// OnceLock<T> : initialisation paresseuse thread-safe (std depuis 1.80)
static CONFIG: OnceLock<String> = OnceLock::new();
fn config() -> &'static str {
    CONFIG.get_or_init(|| {
        std::env::var("APP_CONFIG").unwrap_or_else(|_| "/etc/app.conf".into())
    })
}

fn main() {
    println!("{}", c.get());
    println!("{:?}", rc.borrow());
    println!("{}", config());
}
```

| Type | Thread-safe | Vérification | Usage |
|---|---|---|---|
| `Cell<T: Copy>` | Non | — | Compteurs simples mono-thread |
| `RefCell<T>` | Non | Exécution (panic si abus) | Partage `&` + mutation mono-thread |
| `Mutex<T>` | Oui | Exécution (blocage) | Partage multi-thread |
| `OnceLock<T>` | Oui | — | Global initialisé une fois |

---

## 68. Smart pointers : Box, Rc, Arc

```rust
use std::rc::Rc;
use std::sync::Arc;

fn main() {
    // Box<T> : allocation unique sur le tas, propriétaire unique
    let b = Box::new([0u8; 1024]);       // gros tableau : évite de saturer la pile
    println!("len {}", b.len());

    // Rc<T> : propriété PARTAGÉE mono-thread (reference counted)
    let partage = Rc::new(String::from("config"));
    let a = Rc::clone(&partage);         // compteur += 1 (pas cher)
    let b2 = Rc::clone(&partage);
    println!("{} refs = {}", a, Rc::strong_count(&partage));  // 3

    // Arc<T> : comme Rc mais atomique -> threads (voir §45)
    let partage_t = Arc::new(42);
    let _c = Arc::clone(&partage_t);

    // Rc/Arc + RefCell/Mutex pour partager ET muter :
    let mutable_partage = Rc::new(RefCell::new(0));
    *mutable_partage.borrow_mut() += 1;
    println!("{}", mutable_partage.borrow());
}
```

> `Rc::clone(&x)` plutôt que `x.clone()` : signale visuellement qu'on ne clone pas la donnée.
> Cycles `Rc` → fuite : casser avec `Weak::new()`.

---

## 69. Erreurs FFI/libc courantes et garde-fous

- Oublier `CString` : passer `&str` brut à du C = octets non terminés → UB.
- `CStr::from_ptr` sur pointeur nul ou non terminé → UB. Vérifier le contrat de l'API C.
- `String::from_raw_parts` : à réserver aux cas où TU as alloué via l'API Rust.
- Thread C qui rappelle du Rust : déclarer `extern "C" fn` côté Rust, attention aux panics
  (un panic qui traverse la FFI = UB → `catch_unwind` aux frontières).

---

## 70. Patterns utiles : builder, newtype, type state

```rust
// --- Builder : construction lisible d'objets complexes ---
#[derive(Debug)]
struct Sonde { hote: String, port: u16, timeout_ms: u64, retries: u8 }

impl Sonde {
    fn builder(hote: &str) -> SondeBuilder {
        SondeBuilder { hote: hote.to_string(), port: 443, timeout_ms: 1000, retries: 3 }
    }
}

struct SondeBuilder { hote: String, port: u16, timeout_ms: u64, retries: u8 }
impl SondeBuilder {
    fn port(mut self, p: u16) -> Self { self.port = p; self }
    fn timeout_ms(mut self, t: u64) -> Self { self.timeout_ms = t; self }
    fn retries(mut self, r: u8) -> Self { self.retries = r; self }
    fn build(self) -> Sonde {
        Sonde { hote: self.hote, port: self.port, timeout_ms: self.timeout_ms, retries: self.retries }
    }
}

// --- Newtype : unités impossibles à mélanger ---
struct Millisecondes(u64);
struct Pourcent(u8);
fn attendre(d: Millisecondes) { std::thread::sleep(std::time::Duration::from_millis(d.0)); }

fn main() {
    let s = Sonde::builder("10.0.0.1").port(22).retries(5).build();
    println!("{s:?}");
    attendre(Millisecondes(200));
    // attendre(Pourcent(50));   // ❌ ne compile pas : type différent !
}
```

> Builder = quand > 3 paramètres ou beaucoup d'optionnels. Newtype = quand une confusion
> d'unités peut casser du matériel (ms vs s, W vs kW, % vs ratio).

---

## 71. Optimiser les allocations : le guide express

