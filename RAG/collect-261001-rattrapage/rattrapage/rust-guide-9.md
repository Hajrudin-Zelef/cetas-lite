---
id: collect-261001-rattrapage/rattrapage/rust-guide-9
title: "Guide Rust ultra-complet — Langage système sûr, rapide et moderne"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["arr", "cost"]
source: docs/RAG/collect-261001-rattrapage/rust_guide.md
source_anchor: ""
source_lines: [1929, 2092]
sha256: 6b87347f10f7d41467220a8b1fa3617f4a9a7e6d9d46d5ce2c78b5e3fe2d6bea
---

# Guide Rust ultra-complet — Langage système sûr, rapide et moderne

| Terme | Définition |
|---|---|
| Ownership | Système de possession unique : chaque valeur a un propriétaire qui la libère |
| Borrow checker | Analyse statique du compilateur qui valide emprunts et durées de vie |
| Move | Transfert de propriété ; l'ancien propriétaire devient invalide |
| Borrow (emprunt) | Prêt temporaire via `&T` / `&mut T`, sans transfert de propriété |
| Lifetime (`'a`) | Durée de vie d'une référence, souvent élidée |
| NLL | Non-Lexical Lifetimes : l'emprunt meurt à son dernier usage |
| Trait | Contrat de comportement partagé (interface + méthodes par défaut) |
| Générique | Code paramétré par type, monomorphisé à la compilation |
| Crate | Unité de compilation ; `crates.io` = registre public |
| Cargo | Gestionnaire de build, dépendances, tests, doc |
| rustup | Gestionnaire de toolchains Rust |
| `Result<T,E>` | `Ok(T)` ou `Err(E)` : erreur = valeur, pas d'exception |
| `Option<T>` | `Some(T)` ou `None` : remplace `null` |
| Panic | Arrêt brutal sur bug irrécupérable (débordement, unwrap sur None) |
| `unsafe` | Bloc autorisant 5 opérations sensibles, sous contrat du programmeur |
| RAII | La libération suit la portée (drop automatique) |
| Zero-cost abstraction | Abstraction sans surcoût à l'exécution (itérateurs, génériques) |
| Monomorphisation | Génération d'une version par type concret pour les génériques |
| FFI | Foreign Function Interface : appeler du C (ou être appelé) |
| Édition | Version du langage (2015/2018/2021/2024) : évolutions sans casser l'existant |
| Shadowing | Redéclaration `let` qui masque la précédente (type modifiable) |
| Deref coercion | Conversion auto `&String`→`&str`, `&Vec<T>`→`&[T]` |
| Send / Sync | Traits marqueurs : transférable entre threads / partageable entre threads |
| Tokio | Runtime async de référence |
| Miri | Interpréteur (nightly) détectant les comportements indéfinis |

---

## 42. Quiz : 10 questions (réponses en §43)

1. Que se passe-t-il quand le propriétaire d'une `String` sort de portée ?
2. Peut-on avoir deux `&mut` sur la même valeur en même temps ? Pourquoi ?
3. `let x = vec![1,2]; let y = x;` — que vaut `x` après ? Et si c'était un `i32` ?
4. À quoi sert l'opérateur `?` et dans quel type de fonction peut-on l'utiliser ?
5. Pourquoi Rust n'a-t-il pas d'exceptions ? Quel type les remplace ?
6. `match` doit être exhaustif : qu'est-ce que ça garantit concrètement ?
7. Différence entre `&String` et `&str` en paramètre de fonction : lequel préférer ?
8. Que signifie `#[derive(Debug, Clone)]` au-dessus d'un struct ?
9. `String` est-elle `Copy` ? Pourquoi (où sont ses données) ?
10. À quoi sert `cargo check` par rapport à `cargo build` ?

---

## 43. Quiz : réponses

1. La valeur est **droppée** : `Drop` libère la mémoire du tas automatiquement (RAII).
   Aucun `free` manuel, aucun GC.
2. **Non.** Règle d'emprunt : soit N `&T`, soit un seul `&mut T`. Ça élimine les data races
   à la compilation.
3. `x` est **déplacé** (move) : invalide. Avec `i32` (type `Copy`), `x` reste valide
   (copie bit à bit).
4. `?` **propage l'erreur** : sur `Err(e)`/`None`, retourne immédiatement `e` à l'appelant
   (avec conversion `From`). Utilisable dans toute fonction retournant `Result` ou `Option`.
5. Les exceptions rendent les chemins d'erreur **invisibles** dans les signatures.
   `Result<T, E>` rend l'échec **explicite et obligatoire à traiter**.
6. Le compilateur **refuse d'oublier un cas** (ajout d'une variante d'enum = erreur de
   compilation aux endroits à mettre à jour, pas bug silencieux).
7. Préférer **`&str`** : accepte `&String` (deref coercion) ET les littéraux, donc plus général.
8. Génère automatiquement les implémentations `Debug` (affichage `{:?}`) et `Clone` (`.clone()`).
9. **Non** : ses octets sont sur le tas (pointeur + len + cap sur la pile). La copie serait
   un double pointeur → double free. D'où le move.
10. `cargo check` **vérifie** (types, emprunts) sans générer de code machine : bien plus
    rapide en itération. `cargo build` produit le binaire.

---

## 44. Pour aller plus loin

| Ressource | Type | Note |
|---|---|---|
| *The Rust Programming Language* (le « Book ») | Livre officiel gratuit | La référence, en ligne |
| *Rust by Example* | Tutoriels | Apprendre par l'exemple |
| docs.rs | Documentation | Doc de toutes les crates |
| Rustlings | Exercices CLI | `cargo install rustlings` : petits exercices corrigés |
| *The Rustonomicon* | Livre | L'unsafe en profondeur |
| *Asynchronous Programming in Rust* | Livre officiel | L'async/await |
| This Week in Rust | Newsletter | Suivi de l'écosystème |
| crates.io | Registre | Chercher des crates (trier par téléchargements) |
| users.rust-lang.org | Forum | Questions/réponses |
| `cargo audit`, `cargo deny` | Outils | Sécurité des dépendances en CI |

Prochaines étapes conseillées pour un sysadmin :

1. Réécrire un script Python d'inventaire en Rust (serde + clap).
2. Écrire un exporter Prometheus minimal (axum).
3. Packager en binaire statique musl + image `scratch` Docker.
4. Ajouter `cargo audit` à ta CI.

---

## 45. Concurrence : threads, Arc, Mutex, channels

```rust
use std::sync::{mpsc, Arc, Mutex};
use std::thread;
use std::time::Duration;

fn main() {
    // --- thread::spawn : threads OS, `move` obligatoire ---
    let poignee = thread::spawn(|| {
        thread::sleep(Duration::from_millis(100));
        42
    });
    println!("résultat thread : {}", poignee.join().unwrap());  // join -> Result

    // --- Arc<Mutex<T>> : état partagé entre threads ---
    // Arc = Atomic Reference Counted (compteur thread-safe) ; Mutex = exclusion mutuelle
    let compteur = Arc::new(Mutex::new(0u64));
    let mut poignees = vec![];
    for _ in 0..10 {
        let c = Arc::clone(&compteur);       // clone le POINTEUR, pas la donnée
        poignees.push(thread::spawn(move || {
            let mut n = c.lock().unwrap();   // MutexGuard : verrou jusqu'à fin de portée
            *n += 1;
        }));
    }
    for p in poignees { p.join().unwrap(); }
    println!("compteur = {}", *compteur.lock().unwrap());   // 10

    // --- Channels : "ne communique pas en partageant, partage en communiquant" ---
    let (tx, rx) = mpsc::channel();
    thread::spawn(move || {
        for i in 0..3 { tx.send(format!("msg-{i}")).unwrap(); }
    });
    for recu in rx {                       // se termine quand tous les tx sont droppés
        println!("reçu : {recu}");
    }
}
```

| Primitive | Rôle |
|---|---|
| `thread::spawn` | Thread OS, closure `move`, `JoinHandle::join()` |
| `mpsc::channel()` | Multi-producteur, mono-consommateur |
| `Arc<T>` | Propriété partagée thread-safe (compteur atomique) |
| `Mutex<T>` / `RwLock<T>` | Exclusion mutuelle / lecteurs-multiples |
| `AtomicU64`… | Compteurs sans verrou (`Ordering::SeqCst` par défaut sûr) |
| `thread::scope` (Rust 1.63+) | Threads scopés : emprunter la pile parente sans `Arc` ! |

```rust
// Scoped threads : plus besoin d'Arc pour emprunter des données locales
fn main2() {
    let donnees = vec![1, 2, 3, 4];
    let mut somme = 0;
    thread::scope(|s| {
        let (g, d) = donnees.split_at(2);
        let h1 = s.spawn(|| g.iter().sum::<i32>());
        let h2 = s.spawn(|| d.iter().sum::<i32>());
        somme = h1.join().unwrap() + h2.join().unwrap();
    });   // tous les threads sont joints ici, garanti
    println!("somme = {somme}");
}
```

> **Send / Sync** : `Send` = peut traverser les threads (presque tout) ; `Sync` = `&T`
> partageable (`Mutex<T>` est `Sync` si `T: Send`). `Rc<T>` et `Cell<T>` ne sont PAS `Send` :
> le compilateur refuse de les envoyer dans un thread. C'est la garantie anti-data-race.

---

## 46. Async / Tokio : introduction

