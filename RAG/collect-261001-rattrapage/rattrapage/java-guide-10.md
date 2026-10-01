---
id: collect-261001-rattrapage/rattrapage/java-guide-10
title: "Guide Java — du zéro au production"
domain: rattrapage
role: reference
task: reference
actors: ["Microsoft", "OpenAI"]
dates: []
keywords: ["attention", "mai"]
source: docs/RAG/collect-261001-rattrapage/java_guide.md
source_anchor: ""
source_lines: [1947, 2158]
sha256: 889fd918230f27de78d61cda112bc64db6a5a927742bdf49238ba81d7d6d7de0
---

# Guide Java — du zéro au production

- [ ] Le code compile sans warning nouveau (`-Xlint:all`)
- [ ] Tests unitaires : cas nominaux + limites + erreurs
- [ ] Pas de secret / chemin en dur
- [ ] Ressources fermées (try-with-resources)
- [ ] `equals`/`hashCode` cohérents si usage en Set/Map
- [ ] Logs utiles (pas de `printStackTrace()`, niveaux cohérents)
- [ ] Complexité raisonnable (méthode < 30 lignes, imbrication < 3)

---

## 55. Erreurs classiques des débutants (et comment les éviter)

### 1. `NullPointerException` (NPE)
```java
String nom = null;
nom.length();  // 💥 NPE
// Remèdes : valider en entrée (Objects.requireNonNull), Optional en retour,
// vérifier les retours de Map.get(), initialiser les champs.
```

### 2. `==` sur les String (et les objets en général)
```java
if (etat == "EN_SERVICE") { }      // ❌ compare les références
if ("EN_SERVICE".equals(etat)) { } // ✅ compare le contenu (et null-safe !)
```

### 3. Division entière involontaire
```java
double moyenne = total / nombre;      // ❌ si int/int -> tronqué AVANT conversion
double moyenne = (double) total / nombre; // ✅
```

### 4. `ConcurrentModificationException`
Modifier une collection pendant son for-each → voir section 29 (`removeIf`, `Iterator.remove()`).

### 5. Oublier `break` dans un switch classique / confondre `=` et `==`
```java
if (etat = "HS") { }  // ❌ compile si etat est boolean... et assigne ! (avec String : erreur de compilation, ouf)
```
Préférez les **switch expressions** (section 9) : pas de `break` à oublier.

### 6. Fuite de ressources (fichier/socket non fermé)
```java
var r = Files.newBufferedReader(p); r.readLine(); // ❌ jamais fermé si exception
// ✅ try-with-resources systématique (section 31)
```

### 7. `ArrayIndexOutOfBoundsException` : boucle `<=` au lieu de `<`
```java
for (int i = 0; i <= t.length; i++) { } // ❌ le dernier index valide est length-1
for (int i = 0; i < t.length; i++) { }  // ✅ (ou for-each)
```

### 8. Comparer des `double` avec `==`
```java
if (0.1 + 0.2 == 0.3) { } // ❌ false ! (0.30000000000000004)
if (Math.abs(a - b) < 1e-9) { } // ✅ epsilon ; ou BigDecimal pour l'exact
```

### 9. Ignorer la valeur de retour des méthodes immutables
```java
s.trim();                    // ❌ ne modifie PAS s (String immutable)
s = s.trim();                // ✅
liste = liste.stream().filter(...).toList(); // idem : le stream ne modifie pas la liste
```

### 10. `static` abusif / tout mettre dans `main`
Le `static` = état global partagé. Un programme « tout static » n'est ni testable ni extensible : créez des objets, injectez les dépendances (constructeur).

### 11. Catcher `Exception` et ne rien faire
```java
try { ... } catch (Exception e) { } // ❌ bug invisible garanti
// ✅ logger + action, ou laisser remonter (section 32)
```

### 12. Confondre `List.of(...)` immutable et `ArrayList`
```java
var l = List.of("a"); l.add("b"); // 💥 UnsupportedOperationException
var l2 = new ArrayList<>(List.of("a")); l2.add("b"); // ✅
```

**À retenir** : 80 % des bugs débutants = null, `==` vs `equals`, ressources non fermées, division entière. La checklist de la section 54 les attrape en revue.

---

## 56. Nouveautés modernes : var, switch, text blocks, pattern matching

Récapitulatif versionné (tout ce qui suit exige la version indiquée minimum) :

```java
// var — inférence locale (Java 10+)
var inventaire = new ArrayList<Equipement>();  // ArrayList<Equipement>
var ligne = Files.readString(path);            // String
// Interdit : var x; / var n = null; / var comme type de champ ou de paramètre

// Switch expression — valeur de retour + flèches (Java 14+)
int joursMaintenance = switch (mois) {
    case JANVIER, MARS, MAI -> 2;
    case FEVRIER -> 1;
    default -> 0;
};

// yield — retour depuis un bloc case (Java 14+)
String niveau = switch (criticite) {
    case 1, 2 -> "info";
    case 3 -> {
        notifierAstrainte();
        yield "warning";
    }
    default -> "critique";
};

// Text blocks — multiligne (Java 15+)
String script = """
    #!/bin/bash
    echo "Sauvegarde du %s"
    """.formatted(nom);

// Pattern matching instanceof (Java 16+) : plus de cast explicite
if (obj instanceof Onduleur o) {           // o est typé et non-null ici
    System.out.println(o.getPuissanceKva());
}
// Avant : if (obj instanceof Onduleur) { Onduleur o = (Onduleur) obj; ... }

// Records (Java 16+), classes scellées (Java 17+) : hiérarchies fermées et exhaustives
sealed interface Alerte permits AlerteInfo, AlerteCritique {}
record AlerteInfo(String msg) implements Alerte {}
record AlerteCritique(String msg, int niveau) implements Alerte {}

// Pattern matching switch exhaustif sur sealed (Java 21+)
String traiter(Alerte a) {
    return switch (a) {   // le compilateur VÉRIFIE l'exhaustivité : pas de default nécessaire
        case AlerteInfo i -> "INFO: " + i.msg();
        case AlerteCritique c -> "CRITIQUE(" + c.niveau() + "): " + c.msg();
    };
}

// Virtual threads (Java 21+), String templates (preview Java 21+, finalisé plus tard)
// SequencedCollections : List.getFirst()/getLast(), Map.firstEntry() (Java 21+)
var premier = onduleurs.getFirst();
```

**Conseil version** : écrivez en « Java 17 » par défaut (le plus répandu en entreprise), adoptez les features 21 (virtual threads, pattern matching switch) quand le runtime cible est 21+.

---

## 57. Expressions régulières

```java
import java.util.regex.*;

// Précompiler en constante static final (voir section 52)
private static final Pattern IPV4 =
    Pattern.compile("^(?:\\d{1,3}\\.){3}\\d{1,3}$");
private static final Pattern MAC =
    Pattern.compile("^([0-9A-Fa-f]{2}:){5}[0-9A-Fa-f]{2}$");

System.out.println(IPV4.matcher("10.0.0.300").matches()); // true (format ; pas la validité 0-255 !)
System.out.println(MAC.matcher("AA:BB:CC:DD:EE:FF").matches()); // true

// Groupes de capture : extraire des champs d'une ligne de log
Pattern LOG = Pattern.compile("(\\S+) \\S+ \\S+ \\[([^]]+)] \"(\\S+) (\\S+)");
Matcher m = LOG.matcher("10.0.0.5 - - [26/Sep/2026:10:00:01] \"GET /api/ups HTTP/1.1\"");
if (m.find()) {
    System.out.println("IP=" + m.group(1) + " méthode=" + m.group(3) + " chemin=" + m.group(4));
}

// Remplacements
String propre = "  UPS-01 ; ; UPS-02  ".trim().replaceAll("\\s*;\\s*", ";");

// Mémento
// .         n'importe quel caractère        \d chiffre  \w [a-zA-Z0-9_]  \s espace
// * 0+  + 1+  ? 0/1  {n,m}                 ^ début  $ fin  | ou  () groupe
// (?i) insensible à la casse               \\. point littéral
```

**Attention** : valider une IPv4 « 0-255 » en regex pure est verbeux ; en pratique, combinez regex de format + `InetAddress.getByName()` ou parsing manuel des octets.

---

## 58. Annotations et réflexion (intro)

```java
import java.lang.annotation.*;

// Définir une annotation
@Retention(RetentionPolicy.RUNTIME)          // visible à l'exécution
@Target(ElementType.METHOD)                  // applicable aux méthodes
public @interface Planifiable {
    String cron() default "0 0 * * *";
    String description() default "";
}

// L'utiliser
public class Taches {
    @Planifiable(cron = "0 */5 * * *", description = "Relève compteurs")
    public void releverCompteurs() { /* ... */ }
}

// La lire par réflexion (c'est comme ça que Spring trouve vos @GetMapping)
for (var methode : Taches.class.getDeclaredMethods()) {
    Planifiable p = methode.getAnnotation(Planifiable.class);
    if (p != null) System.out.println(methode.getName() + " -> " + p.cron());
}
```

Annotations du quotidien : `@Override`, `@Deprecated`, `@SuppressWarnings("unchecked")`, `@FunctionalInterface`. La réflexion est puissante mais lente et casse l'encapsulation : réservez-la aux frameworks/outils, pas au code métier.

---

## 59. JSON : sérialisation avec Jackson

