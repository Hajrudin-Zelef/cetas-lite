---
id: collect-261001-rattrapage/rattrapage/java-guide-2
title: "Guide Java — du zéro au production"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/java_guide.md
source_anchor: ""
source_lines: [231, 492]
sha256: 0a9bbaf8ce45b969f57b878242b63c4d92062e63f294eb5408a4b62bf8c9f115
---

# Guide Java — du zéro au production

    // Méthode
    public static double prixTTC(double prixHT) {
        return prixHT * (1 + TVA);
    }
}
```

Règles de base :
- Instructions terminées par `;`.
- Blocs délimités par `{ }`.
- Sensible à la casse : `prix` ≠ `Prix`.
- Conventions de nommage : `MaClasse` (PascalCase), `maMethode`/`maVariable` (camelCase), `MA_CONSTANTE` (UPPER_SNAKE).

---

## 6. Types primitifs et variables

| Type | Taille | Plage / valeurs | Défaut (attribut) |
|---|---|---|---|
| `byte` | 8 bits | -128 à 127 | 0 |
| `short` | 16 bits | -32 768 à 32 767 | 0 |
| `int` | 32 bits | ±2,1 milliards | 0 |
| `long` | 64 bits | ±9×10¹⁸ | 0L |
| `float` | 32 bits | ~±3,4×10³⁸ | 0.0f |
| `double` | 64 bits | ~±1,8×10³⁰⁸ | 0.0 |
| `char` | 16 bits | caractère Unicode | '\u0000' |
| `boolean` | — | `true` / `false` | false |

```java
int compteur = 0;
long population = 8_000_000_000L;   // _ = séparateur de lisibilité, L obligatoire
double prix = 19.99;
float ratio = 0.5f;                  // f obligatoire, sinon double
char lettre = 'A';
boolean actif = true;

var nom = "Zelef";        // inférence de type (Java 10+) : nom est un String
// var x;                 // INTERDIT : var exige une initialisation immédiate

final int PORT = 8080;    // constante : ne peut plus être réassignée
```

**Pièges** :
- `int / int` = division **entière** : `5 / 2` vaut `2`, pas `2.5`. Écrivez `5 / 2.0`.
- Dépassement silencieux : `int` qui déborde ne lève **aucune** erreur (il « boucle »). Pour des calculs critiques, utilisez `long` ou `Math.addExact` (lève `ArithmeticException` en cas de débordement).
- N'utilisez **jamais** `float`/`double` pour de la monnaie : utilisez `BigDecimal` (voir section 35).

**À retenir** : `int`/`long`/`double`/`boolean` couvrent 95 % des besoins ; `var` allège le code mais n'est pas du typage dynamique.

---

## 7. Opérateurs

```java
int a = 10, b = 3;
System.out.println(a + b);  // 13
System.out.println(a - b);  // 7
System.out.println(a * b);  // 30
System.out.println(a / b);  // 3  (division entière !)
System.out.println(a % b);  // 1  (modulo = reste)

a++; b--;                   // incrémentation / décrémentation
a += 5;                     // a = a + 5 ; idem -= *= /= %=

boolean ok = (a > 5) && (b < 10);  // && et || sont "court-circuit" : la suite n'est pas évaluée si inutile
boolean ko = (a > 5) & (b < 10);   // & évalue toujours les deux côtés (rarement voulu sur des booléens)

int max = (a > b) ? a : b;         // ternaire

// Opérateurs bit à bit (masques, flags, protocoles réseau)
int flags = 0b1010;
int masque = flags & 0b1100;       // ET bit à bit
int ajout  = flags | 0b0001;        // OU bit à bit
int bascule= flags ^ 0b1111;        // OU exclusif
int decale = flags << 2;           // décalage à gauche (= ×4)
```

Priorités à connaître : `*` `/` `%` avant `+` `-` ; comparaisons avant `&&` avant `||`. En cas de doute : **parenthèses**.

---

## 8. Chaînes de caractères : String

`String` est **immutable** : chaque « modification » crée un nouvel objet.

```java
String s = "Bonjour";
String t = "Bonjour";
System.out.println(s == t);        // true ici (pool de String), mais...
System.out.println(s.equals(t));   // true — TOUJOURS comparer avec equals() !

String nom = new String("test");
System.out.println(s == nom);      // false ! (objets différents)
System.out.println(s.equals(nom)); // true

// Concaténation
String msg = "Valeur : " + 42 + " €";   // "Valeur : 42 €"

// Méthodes essentielles
"  hello  ".trim();                // "hello"  (Java 11+ : strip() gère mieux Unicode)
"hello".toUpperCase();             // "HELLO"
"hello".length();                  // 5
"hello".charAt(1);                 // 'e'
"hello".substring(1, 4);           // "ell"
"hello world".contains("world");   // true
"hello".startsWith("he");          // true
"a,b,c".split(",");                // ["a","b","c"]
String.join(", ", "a", "b", "c");  // "a, b, c"
"hello".replace('l', 'L');         // "heLLo"

// Text blocks (Java 15+) : chaînes multilignes propres
String json = """
    {
      "nom": "onduleur-01",
      "puissance_kva": 40
    }
    """;
System.out.println(json);
```

**Concaténer en boucle = piège de performance** : chaque `+` crée un objet. Utilisez `StringBuilder` :

```java
StringBuilder sb = new StringBuilder();
for (int i = 0; i < 10000; i++) {
    sb.append("ligne ").append(i).append('\n');
}
String resultat = sb.toString();
```

**À retenir** : `equals()` pour comparer, jamais `==` ; `StringBuilder` pour construire en boucle ; text blocks pour le multiligne.

---

## 9. Structures de contrôle : if / switch

```java
int note = 14;
if (note >= 16) {
    System.out.println("Très bien");
} else if (note >= 10) {
    System.out.println("Passable");
} else {
    System.out.println("Insuffisant");
}

// Switch expression (Java 14+) : retourne une valeur, sans break oublié
String mention = switch (note / 4) {
    case 5, 4 -> "Excellent";
    case 3 -> "Bien";
    case 2 -> "Passable";
    default -> "Insuffisant";
};

// Pattern matching dans switch (Java 21+)
static String decrire(Object o) {
    return switch (o) {
        case Integer i -> "entier : " + i;
        case String s  -> "chaîne de " + s.length() + " caractères";
        case null      -> "rien (null)";
        default        -> "autre type : " + o.getClass().getSimpleName();
    };
}
```

Notez `case null` : depuis Java 21, un `switch` sur `null` ne lève plus forcément `NullPointerException` si vous gérez le cas.

---

## 10. Boucles

```java
// for classique
for (int i = 0; i < 5; i++) {
    System.out.println("i = " + i);
}

// for-each (à privilégier sur les collections/tableaux)
String[] baies = {"baie-A", "baie-B", "baie-C"};
for (String baie : baies) {
    System.out.println(baie);
}

// while
int n = 3;
while (n > 0) {
    System.out.println(n);
    n--;
}

// do-while : exécuté au moins une fois
int code;
do {
    code = lireCodePin();
} while (code != 1234);

// break / continue avec labels (rare mais utile pour boucles imbriquées)
externe:
for (int i = 0; i < 10; i++) {
    for (int j = 0; j < 10; j++) {
        if (i * j > 50) break externe;
    }
}
```

**Boucle infinie volontaire** (serveur, worker) : `while (true) { ... }` avec une condition de sortie propre.

---

## 11. Tableaux

```java
int[] puissances = new int[4];          // [0,0,0,0]
int[] notes = {10, 20, 30};             // initialisation directe
String[] noms = new String[3];

puissances[0] = 10;
System.out.println(notes.length);       // 3 (attribut, pas méthode !)
System.out.println(notes[notes.length - 1]); // dernier élément : 30

// Parcours
for (int i = 0; i < notes.length; i++) System.out.println(notes[i]);
for (int note : notes) System.out.println(note);

// Tableaux multidimensionnels
int[][] matrice = new int[3][3];
matrice[1][2] = 7;
int[][] triangle = { {1}, {2, 3}, {4, 5, 6} };

// Utilitaires
java.util.Arrays.sort(notes);
System.out.println(java.util.Arrays.toString(notes)); // [10, 20, 30]
int[] copie = java.util.Arrays.copyOf(notes, notes.length);
```

**Limites** : taille fixe à la création ; `notes[5]` sur un tableau de 3 → `ArrayIndexOutOfBoundsException`. Pour une taille dynamique : `ArrayList` (section 25).

---

## 12. Méthodes (fonctions)

```java
public class Calculatrice {

    // Signature : visibilité + (static?) + type de retour + nom + paramètres
    public static int additionner(int a, int b) {
        return a + b;
    }

    // Surcharge (overloading) : même nom, paramètres différents
    public static double additionner(double a, double b) {
        return a + b;
    }

    // Nombre variable d'arguments (varargs)
    public static int somme(int... valeurs) {
        int total = 0;
        for (int v : valeurs) total += v;
        return total;
    }
    // somme(1, 2, 3) -> 6 ; somme() -> 0

