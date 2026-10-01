---
id: collect-261001-rattrapage/rattrapage/rust-guide-2
title: "Guide Rust ultra-complet — Langage système sûr, rapide et moderne"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["arr"]
source: docs/RAG/collect-261001-rattrapage/rust_guide.md
source_anchor: ""
source_lines: [201, 434]
sha256: ff4c8fc72babbf9f203cec04aba78aaa077c737834d35652dfba23e89a48f87d
---

# Guide Rust ultra-complet — Langage système sûr, rapide et moderne

```rust
fn main() {
    let a: i32 = -42;
    let b = 42u8;                       // suffixe de type sur littéral
    let c = 0xFF_FF;                    // hexadécimal, _ comme séparateur
    let d = 0b1010_0110;                // binaire
    let e = 1_000_000;                  // lisibilité
    let f = 3.14f32;
    let g: bool = true;
    let h: char = '🦀';                 // la mascotte de Rust : Ferris le crabe

    // Conversions : explicites, jamais implicites
    let petit: u8 = 200;
    let grand: u32 = petit as u32;      // `as` : transtypage (tronque si dépassement)
    let sur: i32 = petit as i32;

    // Dépassement d'entier : panic en debug, wrap en release !
    // let boom: u8 = 255 + 1;          // panic en debug
    let wrap = 255u8.wrapping_add(1);    // 0, explicite
    let sat = 255u8.saturating_add(1);  // 255, explicite
    let chk = 255u8.checked_add(1);     // None, explicite
    println!("{wrap} {sat} {chk:?}");
}
```

> ⚠️ En `debug`, un dépassement arithmétique fait paniquer le programme.
> En `release`, il s'enroule silencieusement (wrapping). Utilise les méthodes
> `wrapping_*` / `saturating_*` / `checked_*` / `overflowing_*` quand le dépassement est possible.

---

## 7. Tuples, tableaux, slices

```rust
fn main() {
    // Tuple : taille fixe, types hétérogènes
    let srv: (&str, u16, bool) = ("srv-web-01", 443, true);
    let nom = srv.0;                    // accès par index
    let (hote, port, _actif) = srv;     // déstructuration

    // Tableau : taille fixe, type homogène, sur la pile
    let ports: [u16; 4] = [80, 443, 8080, 8443];
    let zeros = [0u8; 16];              // [valeur; taille]
    println!("{}", ports[1]);           // indexation (panic si hors bornes)

    // Slice : vue sur une séquence contiguë (taille inconnue à la compilation)
    let tranche: &[u16] = &ports[1..3]; // [443, 8080]
    println!("{:?}", tranche);

    // Chaînes : &str (vue) vs String (possédée) — voir § ownership
    let s1: &str = "immuable, dans le binaire";
    let mut s2 = String::from("modifiable, sur le tas");
    s2.push_str(" !");
}
```

| Structure | Taille | Où | Copiable ? |
|---|---|---|---|
| Tuple | Fixe (connue à la compilation) | Pile (si ses éléments y sont) | Si tous les éléments sont `Copy` |
| `[T; N]` | Fixe | Pile | Si `T: Copy` |
| `&[T]` (slice) | Dynamique (fat pointer : ptr + len) | Vue sur pile/tas/statique | Oui (la vue, pas les données) |
| `Vec<T>` | Dynamique | Tas | Non (ownership) |

---

## 8. Contrôle de flux : if, boucles, labels

```rust
fn main() {
    // if est une EXPRESSION : retourne une valeur
    let charge = 78;
    let statut = if charge > 90 { "critique" } else if charge > 70 { "alerte" } else { "ok" };

    // loop : boucle infinie, `break` peut retourner une valeur
    let mut n = 0;
    let resultat = loop {
        n += 1;
        if n == 10 { break n * 2; }      // break avec valeur
    };

    // while classique
    while n > 0 { n -= 1; }

    // for sur itérateur (pas d'index manuel = pas d'erreur off-by-one)
    let ports = [80, 443, 8080];
    for p in ports {                     // `ports` est Copy ([u16; 3]) : pas de move ici
        println!("port {p}");
    }
    for (i, p) in ports.iter().enumerate() {
        println!("{i}: {p}");
    }
    for i in 1..=5 {                     // 1..=5 inclusif ; 1..5 exclusif
        println!("{i}");
    }

    // Labels pour sortir de boucles imbriquées
    'exterieur: loop {
        loop {
            break 'exterieur;            // sort des DEUX boucles
        }
    }
}
```

> En Rust, `if`/`match`/`loop` sont des expressions : on écrit `let x = if ... { ... } else { ... };`.
> Les blocs retournent la dernière expression **sans point-virgule**.

---

## 9. Fonctions : signatures, retours, expressions

```rust
// snake_case obligatoire (clippy râle sinon)
fn ping(hote: &str, timeout_ms: u64) -> bool {
    // pas de `return` nécessaire : dernière expression = valeur de retour
    !hote.is_empty() && timeout_ms > 0
}

fn afficher(b: bool) {                  // -> () implicite
    println!("joignable : {b}");
}

// Fonctions divergentes : ne retournent jamais
fn mourir(msg: &str) -> ! {
    eprintln!("FATAL : {msg}");
    std::process::exit(1);
}

fn main() {
    let ok = ping("10.0.0.1", 500);
    afficher(ok);
    // mourir("test");                   // le programme s'arrête ici
}
```

Règles d'or :

- Les paramètres et le retour sont **toujours typés** (pas d'inférence sur les signatures).
- Le corps est un bloc d'expressions ; `return` existe mais on le réserve aux sorties anticipées.
- Point-virgule final = la fonction retourne `()` au lieu de la valeur. C'est l'erreur n°1 des débutants.

---

## 10. Ownership : le concept central (pas à pas)

L'ownership répond à UNE question : **qui libère la mémoire, et quand ?**
En C tu appelles `free`, en Go/Java un GC s'en charge, en Rust **le compilateur** insère la
libération automatiquement grâce à 3 règles.

### Règle 1 — Chaque valeur a exactement un propriétaire

```rust
fn main() {
    let s = String::from("données");  // s est propriétaire de la String
}                                    // fin de portée : s est "droppée", mémoire libérée
```

### Règle 2 — Un seul propriétaire à la fois (move)

```rust
fn main() {
    let s1 = String::from("hello");
    let s2 = s1;              // MOVE : s1 n'est plus valide
    // println!("{s1}");      // ERREUR : borrow of moved value
    println!("{s2}");         // OK
}
```

Pourquoi ? `String` = (pointeur, longueur, capacité) sur la pile + octets sur le tas.
Copier bêtement le pointeur donnerait deux propriétaires → double `free` → faille.
Rust **invalide l'ancien propriétaire** au lieu de copier : c'est le *move*.

### Les types `Copy` : l'exception

```rust
fn main() {
    let a: i32 = 5;
    let b = a;                // COPIE bit à bit, a reste valide
    println!("{a} {b}");      // OK : i32 est Copy (entiers, flottants, bool, char, tuples de Copy...)
}
```

`Copy` = données entièrement sur la pile, copie triviale. `String`, `Vec<T>`, `Box<T>` ne sont PAS `Copy`.

### Règle 3 — Quand le propriétaire sort de portée, la valeur est droppée

```rust
fn prendre(s: String) {       // s devient propriétaire (move depuis l'appelant)
    println!("{s}");
}                             // s droppée ici

fn main() {
    let s = String::from("x");
    prendre(s);               // move
    // println!("{s}");       // ERREUR : valeur déplacée
}
```

Pour **prêter** sans transférer : les références, §11.

> 🧠 Modèle mental : une valeur = un jeton physique. Tu peux le donner (move), le prêter
> (borrow), ou le cloner explicitement (`s.clone()` = copie profonde, coûteuse mais explicite).

---

## 11. Borrowing : références immuables et mutables

```rust
fn longueur(s: &String) -> usize {   // &String : emprunt immuable
    s.len()
}                                   // la référence meurt, RIEN n'est libéré

fn ajouter(s: &mut String) {         // &mut String : emprunt mutable
    s.push_str(" + suffixe");
}

fn main() {
    let s = String::from("hello");
    println!("len = {}", longueur(&s));   // on prête, on garde la propriété
    println!("{s}");                      // s toujours valide

    let mut m = String::from("hello");
    ajouter(&mut m);
    println!("{m}");
}
```

### Les deux règles d'emprunt (le cœur du borrow checker)

1. **OU** un nombre quelconque de références immuables `&T`,
2. **OU** exactement **une** référence mutable `&mut T`.

Jamais les deux en même temps. C'est ce qui élimine les data races **à la compilation**.

