---
id: collect-261001-rattrapage/rattrapage/java-guide-5
title: "Guide Java — du zéro au production"
domain: rattrapage
role: reference
task: reference
actors: ["Lambda"]
dates: []
keywords: ["consumer"]
source: docs/RAG/collect-261001-rattrapage/java_guide.md
source_anchor: ""
source_lines: [930, 1148]
sha256: 617874c38403b266bd7fd3eedbd2eec47a0842d7b9fff4f44483522718bec1d7
---

# Guide Java — du zéro au production

Set<String> ordonne = new TreeSet<>();        // trié (éléments Comparable)
Set<String> insertion = new LinkedHashSet<>(); // ordre d'insertion préservé

// Opérations ensemblistes
Set<String> a = new HashSet<>(Set.of("1", "2", "3"));
Set<String> b = new HashSet<>(Set.of("3", "4"));
a.retainAll(b);  // intersection : a = {"3"}
```

**Contrat crucial** : si vous mettez vos objets dans un `HashSet`/`HashMap`, redéfinissez **toujours** `equals()` **et** `hashCode()` ensemble (les records le font gratuitement — encore une raison de les utiliser).

---

## 28. Map : HashMap en détail

```java
Map<String, String> dns = new HashMap<>();
dns.put("ups-01", "10.0.0.11");
dns.put("switch-01", "10.0.0.12");
System.out.println(dns.get("ups-01"));          // 10.0.0.11
System.out.println(dns.get("inconnu"));         // null
System.out.println(dns.getOrDefault("inconnu", "0.0.0.0")); // 0.0.0.0

// Parcours
for (Map.Entry<String, String> e : dns.entrySet()) {
    System.out.println(e.getKey() + " -> " + e.getValue());
}
dns.forEach((nom, ip) -> System.out.println(nom + " -> " + ip));

// Méthodes modernes très utiles
dns.putIfAbsent("ups-01", "x");                 // n'écrase pas l'existant
dns.computeIfAbsent("routeur-01", k -> resoudre(k)); // calcule seulement si absent (cache lazy)
dns.merge("compteur", 1, Integer::sum);          // accumulateur : compteur += 1

// Compter des occurrences : l'idiome merge
Map<String, Integer> occurrences = new HashMap<>();
for (String code : codesAlarme) {
    occurrences.merge(code, 1, Integer::sum);
}
```

**Performance** : `HashMap` = O(1) moyen. En concurrence : `ConcurrentHashMap` (jamais `Collections.synchronizedMap` sauf besoin spécifique).

---

## 29. Iterator et ConcurrentModificationException

```java
List<String> liste = new ArrayList<>(List.of("a", "b", "c"));

// ❌ INTERDIT : modifier pendant un for-each
for (String s : liste) {
    if (s.equals("b")) liste.remove(s);  // ConcurrentModificationException !
}

// ✅ Via l'itérateur explicite
Iterator<String> it = liste.iterator();
while (it.hasNext()) {
    if (it.next().equals("b")) it.remove();  // OK
}

// ✅ Ou en Java 8+ : removeIf
liste.removeIf(s -> s.equals("b"));
```

**À retenir** : un for-each = un itérateur caché ; toute modification structurelle pendant l'itération (hors `Iterator.remove()`) explose.

---

## 30. Exceptions : checked vs unchecked

```
Throwable
├── Error               → ne PAS catcher (OutOfMemoryError, StackOverflowError)
└── Exception
    ├── RuntimeException (UNCHECKED) → NullPointerException, IllegalArgumentException...
    └── autres (CHECKED)             → IOException, SQLException : déclaration obligatoire
```

```java
import java.io.IOException;
import java.nio.file.*;

// CHECKED : le compilateur EXIGE try/catch ou throws
public static String lire(Path p) throws IOException {
    return Files.readString(p);
}

// UNCHECKED : pas d'obligation, mais à prévenir par validation
public static double diviser(double a, double b) {
    if (b == 0) throw new IllegalArgumentException("diviseur nul");
    return a / b;
}

public static void main(String[] args) {
    try {
        System.out.println(lire(Path.of("config.txt")));
    } catch (NoSuchFileException e) {
        System.err.println("Fichier absent : " + e.getFile());
    } catch (IOException e) {
        System.err.println("Erreur I/O : " + e.getMessage());
    }
    // Ordre des catch : du PLUS spécifique au PLUS général !
}
```

**Philosophie** : checked = conditions extérieures récupérables (fichier, réseau, SQL) ; unchecked = bugs de programmation (null, argument invalide, état illégal). Ne catchez **jamais** `Exception` ou `Throwable` en bloc « pour être tranquille » — vous masquerez les bugs.

---

## 31. try-with-resources : fermeture automatique (Java 7+)

Toute ressource `AutoCloseable` (fichiers, sockets, connexions JDBC, scanners) **doit** être ouverte ainsi :

```java
// Fermeture garantie, même en cas d'exception, dans l'ordre inverse d'ouverture
try (var lecteur = Files.newBufferedReader(Path.of("journal.log"));
     var ecrivain = Files.newBufferedWriter(Path.of("resume.txt"))) {
    String ligne;
    while ((ligne = lecteur.readLine()) != null) {
        ecrivain.write(ligne.toUpperCase());
        ecrivain.newLine();
    }
} catch (IOException e) {
    System.err.println("Échec : " + e.getMessage());
}
// lecteur et ecrivain sont fermés ici, automatiquement

// Multi-catch
try {
    risquer();
} catch (IOException | IllegalStateException e) {
    System.err.println("Problème : " + e);
}
```

**À retenir** : tout `open()`/`new ...Reader/Writer/Stream/Connection` sans try-with-resources est un bug en puissance (fuite de descripteurs).

---

## 32. Créer ses propres exceptions

```java
// Exception métier checked (récupérable : ex. équipement injoignable)
public class EquipementInjoignableException extends Exception {
    public EquipementInjoignableException(String ip) {
        super("Équipement injoignable : " + ip);
    }
    public EquipementInjoignableException(String ip, Throwable cause) {
        super("Équipement injoignable : " + ip, cause);
    }
}

// Exception métier unchecked (bug/validation : ex. paramètre invalide)
public class ConfigurationInvalideException extends RuntimeException {
    public ConfigurationInvalideException(String message) { super(message); }
}

// Usage : ne jamais "avaler" une exception
try {
    interroger(ip);
} catch (EquipementInjoignableException e) {
    logger.warn("Ping échoué sur {}", ip, e);  // on LOGUE avec la stack trace...
    planifierNouvelEssai(ip);                   // ...et on AGIT
}
```

**Anti-pattern** : `catch (Exception e) {}` vide. Si vous ne savez pas quoi faire, laissez remonter (`throws`) — c'est toujours mieux que le silence.

---

## 33. Lambdas : fonctions anonymes concises (Java 8+)

```java
// Avant : classe anonyme verbeuse
new Thread(new Runnable() {
    @Override public void run() { System.out.println("démarré"); }
}).start();

// Après : lambda
new Thread(() -> System.out.println("démarré")).start();

// Formes
Runnable r = () -> System.out.println("sans paramètre");
Consumer<String> affiche = s -> System.out.println(s);            // 1 param, pas de retour
BiFunction<Integer, Integer, Integer> add = (a, b) -> a + b;     // 2 params, retour
Supplier<Double> aleatoire = () -> Math.random();                // pas de param, retour
Predicate<String> nonVide = s -> !s.isBlank();                   // test -> boolean
Function<String, Integer> longueur = String::length;             // référence de méthode

// Références de méthode : les 4 formes
Function<String, Integer> f1 = String::length;        // méthode d'instance (param = receveur)
BiFunction<String, String, Boolean> f2 = String::equals;
Supplier<List<String>> f3 = ArrayList::new;          // constructeur
Function<Integer, int[]> f4 = int[]::new;            // constructeur de tableau

// Les lambdas capturent les variables locales si elles sont "effectively final"
int seuil = 20;
Predicate<Integer> alerte = v -> v > seuil;  // OK tant que seuil n'est plus modifié après
```

Interfaces fonctionnelles clés (`java.util.function`) : `Predicate<T>`, `Function<T,R>`, `Consumer<T>`, `Supplier<T>`, `BiFunction`, `UnaryOperator`, `BinaryOperator`.

---

## 34. Streams : traitement déclaratif des données (Java 8+)

Un stream = **pipeline** : source → opérations intermédiaires (lazy) → opération terminale.

```java
record Alarme(String equipement, String code, int criticite) {}

List<Alarme> alarmes = List.of(
    new Alarme("UPS-01", "BATTERIE_FAIBLE", 3),
    new Alarme("SW-01", "PORT_DOWN", 2),
    new Alarme("UPS-01", "SURCHARGE", 4),
    new Alarme("RT-01", "PORT_DOWN", 2)
);

