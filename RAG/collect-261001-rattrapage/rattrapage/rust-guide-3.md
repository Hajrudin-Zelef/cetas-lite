---
id: collect-261001-rattrapage/rattrapage/rust-guide-3
title: "Guide Rust ultra-complet — Langage système sûr, rapide et moderne"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["arr"]
source: docs/RAG/collect-261001-rattrapage/rust_guide.md
source_anchor: ""
source_lines: [435, 680]
sha256: e1dc0eb308dd2e5cb19a10f08d6dc7d32e0a83cf93af98b6834070cdb269dcb4
---

# Guide Rust ultra-complet — Langage système sûr, rapide et moderne

```rust
fn main() {
    let mut s = String::from("hello");
    let r1 = &s;              // OK : emprunt immuable
    let r2 = &s;              // OK : plusieurs immuables
    // let r3 = &mut s;       // ERREUR : cannot borrow as mutable (r1, r2 actifs)
    println!("{r1} {r2}");     // r1, r2 utilisés ici...
    let r3 = &mut s;          // OK maintenant (NLL : les emprunts meurent à leur dernier usage)
    r3.push('!');
}
```

> **NLL (Non-Lexical Lifetimes**, édition 2018+) : un emprunt se termine à son **dernier usage**,
> pas à la fin du bloc. Le code ci-dessus compile grâce à ça.

---

## 12. Lifetimes : quand le compilateur a besoin d'aide

La plupart du temps, les durées de vie sont **élidées** (déduites). Elles deviennent explicites
quand une fonction **retourne une référence** liée à ses paramètres.

```rust
// Le compilateur refuse : de quel paramètre vient la référence retournée ?
// fn plus_long(x: &str, y: &str) -> &str {
//     if x.len() > y.len() { x } else { y }
// }

// On annote : 'a = "vit au moins aussi longtemps que x ET y"
fn plus_long<'a>(x: &'a str, y: &'a str) -> &'a str {
    if x.len() > y.len() { x } else { y }
}

fn main() {
    let a = String::from("court");
    let b = String::from("beaucoup plus long");
    let r = plus_long(&a, &b);
    println!("{r}");
}
```

### Les 3 règles d'élision (quand on peut omettre `'a`)

1. Chaque paramètre référence reçoit sa propre lifetime.
2. S'il n'y a **qu'un** paramètre référence en entrée, sa lifetime est donnée à la sortie.
3. Si un paramètre est `&self`/`&mut self` (méthode), sa lifetime est donnée à la sortie.

```rust
fn premier_mot(s: &str) -> &str {   // OK par la règle 2 : fn premier_mot<'a>(s: &'a str) -> &'a str
    s.split_whitespace().next().unwrap_or("")
}
```

### `'static` : vit jusqu'à la fin du programme

```rust
let s: &'static str = "compilé dans le binaire";   // littéraux de chaîne : toujours 'static
```

> En pratique : tu écris rarement des lifetimes à la main sauf pour des fonctions qui
> retournent des références, des structs qui contiennent des références (§13), et du code async.

---

## 13. Structs : données structurées

```rust
// Struct classique (champs nommés)
struct Serveur {
    hostname: String,
    ip: String,
    port: u16,
    actif: bool,
}

// Tuple struct (champs anonymes)
struct Coordonnees(f64, f64);

// Unit struct (marqueur, zéro octet)
struct Supervision;

// Struct avec lifetime : contient une référence empruntée
struct Vue<'a> {
    extrait: &'a str,
}

impl Serveur {
    // Méthode associée (constructeur idiomatique) — pas de `self`
    fn new(hostname: &str, ip: &str, port: u16) -> Self {
        Self { hostname: hostname.to_string(), ip: ip.to_string(), port, actif: true }
    }

    // Méthode : &self = emprunt immuable
    fn adresse(&self) -> String {
        format!("{}:{}", self.ip, self.port)
    }

    // Méthode : &mut self = emprunt mutable
    fn desactiver(&mut self) {
        self.actif = false;
    }
}

fn main() {
    let mut srv = Serveur::new("web-01", "10.0.0.11", 443);
    println!("{}", srv.adresse());
    srv.desactiver();

    // Syntaxe de mise à jour struct (struct update syntax)
    let srv2 = Serveur { hostname: "web-02".into(), ..srv };  // srv partiellement déplacé !
    // println!("{}", srv.hostname);  // ERREUR : hostname a été déplacé dans srv2
    println!("{}", srv.ip);           // ERREUR aussi : `..srv` déplace le reste
}
```

| Forme | Usage typique |
|---|---|
| `struct S { champ: T }` | Données métier (99 % des cas) |
| `struct S(T, U)` | Newtype pattern : `struct Port(u16)` pour typer fort |
| `struct S;` | Marqueur de type, zéro coût |

> **Newtype pattern** : `struct Metres(f64);` vs `struct Secondes(f64);` — le compilateur
> t'empêche de mélanger des unités. Précieux pour configs réseau/énergie.

---

## 14. Enums : le couteau suisse de la modélisation

```rust
// Enum simple (comme en C)
enum Statut { Ok, Alerte, Critique }

// Enum avec données : chaque variante est un "constructeur" différent
enum Evenement {
    Ping { hote: String, latence_ms: u64 },  // variante struct
    Log(String),                              // variante tuple
    Arret,                                    // variante unit
    Metrique(String, f64),
}

// Enum générique : le célèbre Option<T>
enum MonOption<T> { Some(T), None }

// Méthodes sur enum
impl Evenement {
    fn resume(&self) -> String {
        match self {
            Evenement::Ping { hote, latence_ms } => format!("ping {hote}: {latence_ms}ms"),
            Evenement::Log(msg) => format!("log: {msg}"),
            Evenement::Arret => "arrêt".to_string(),
            Evenement::Metrique(nom, val) => format!("{nom}={val}"),
        }
    }
}
```

> Contrairement au C, une enum Rust peut porter des données **et** le compilateur connaît
> toujours la variante active (tag implicite). C'est la base de `Option` et `Result`.

---

## 15. Pattern matching : `match` et `if let`

`match` est **exhaustif** : le compilateur exige de traiter TOUS les cas. Oublier un cas = erreur
de compilation, pas bug en prod.

```rust
enum Alarme { Temperature(f64), Disque(u8), Reseau(String) }

fn traiter(a: Alarme) {
    match a {
        Alarme::Temperature(t) if t > 80.0 => println!("CRITIQUE : {t}°C"),
        Alarme::Temperature(t) => println!("temp : {t}°C"),
        Alarme::Disque(pct) => println!("disque à {pct}%"),
        Alarme::Reseau(msg) => println!("réseau : {msg}"),
        // aucun `_` nécessaire ici : tous les cas sont couverts
    }
}

fn main() {
    // match comme expression
    let code = 2;
    let txt = match code {
        0 => "ok",
        1 | 2 => "avertissement",      // plusieurs motifs avec |
        3..=5 => "erreur",             // plage inclusive
        _ => "inconnu",                // joker (catch-all)
    };

    // if let : sucre syntaxique pour un seul cas intéressant
    let opt: Option<String> = Some("srv-01".to_string());
    if let Some(nom) = opt {
        println!("hôte : {nom}");
    }

    // while let : boucle tant que le motif correspond
    let mut pile = vec![1, 2, 3];
    while let Some(x) = pile.pop() {
        println!("dépilé : {x}");
    }

    // let-else (édition 2021+, Rust 1.65) : extraction ou sortie anticipée
    let entree: Option<u16> = Some(8080);
    let port = match entree { Some(p) => p, None => return };
    // équivalent :
    // let Some(port) = entree else { return };
    println!("port {port} {txt}");
    traiter(Alarme::Temperature(85.5));
}
```

| Construction | Quand l'utiliser |
|---|---|
| `match` | Tous les cas comptent, ou valeur de retour voulue |
| `if let` | Un seul cas intéressant, les autres ignorés |
| `while let` | Consommer jusqu'à épuisement (`pop`, itérateurs, channels) |
| `let-else` | Extraire ou sortir (`return`, `break`, `continue`) |

---

## 16. `Option<T>` : l'absence de valeur, sans `null`

Rust **n'a pas de `null`**. L'absence est explicite via `Option<T>` : `Some(valeur)` ou `None`.
Impossible d'oublier le cas « pas de valeur » : le compilateur l'exige.

```rust
fn trouver_port(services: &[(&str, u16)], nom: &str) -> Option<u16> {
    services.iter().find(|(n, _)| *n == nom).map(|(_, p)| *p)
}

fn main() {
    let services = [("http", 80), ("https", 443)];

    // Méthodes chaînables (le cœur du style Rust)
    let p = trouver_port(&services, "https")
        .map(|p| p + 1000)          // transforme si Some
        .filter(|p| *p < 2000)      // garde si prédicat vrai
        .unwrap_or(8080);           // valeur par défaut si None
    println!("port : {p}");

    // unwrap_or_else : défaut calculé paresseusement
    let q = trouver_port(&services, "smtp").unwrap_or_else(|| {
        eprintln!("smtp introuvable, défaut 25");
        25
    });

