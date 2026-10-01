---
id: collect-261001-rattrapage/rattrapage/java-guide-4
title: "Guide Java — du zéro au production"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/java_guide.md
source_anchor: ""
source_lines: [723, 929]
sha256: 6a0403755925cad4eb47f079d451f307b24e9b864be7201dd05e93e6ba0c2663
---

# Guide Java — du zéro au production

// Usage
EtatEquipement etat = EtatEquipement.DEGRADE;
System.out.println(etat.getLibelle());       // Fonctionnement dégradé
System.out.println(etat.ordinal());          // 1 (position, évitez de vous en servir en persistance)
System.out.println(etat.name());             // "DEGRADE"

for (EtatEquipement e : EtatEquipement.values()) System.out.println(e);
EtatEquipement x = EtatEquipement.valueOf("EN_SERVICE"); // depuis une chaîne (lève IllegalArgumentException si inconnu)

// Les enum se comparent avec == sans risque (singletons garantis)
if (etat == EtatEquipement.HORS_SERVICE) { /* ... */ }
```

**À retenir** : remplacez vos constantes `int`/`String` magiques par des enums typées — le compilateur devient votre garde-fou.

---

## 20. Records : des classes de données compactes (Java 16+)

```java
// Équivaut à une classe immutable avec constructeur, getters, equals, hashCode, toString
public record Mesure(String capteurId, double valeur, java.time.Instant horodatage) {}

Mesure m = new Mesure("temp-01", 21.5, java.time.Instant.now());
System.out.println(m.capteurId());  // accesseurs = nom du composant, pas getXxx()
System.out.println(m);              // Mesure[capteurId=temp-01, valeur=21.5, horodatage=...]

// Record avec validation dans le constructeur compact
public record PlageIp(String cidr) {
    public PlageIp {
        if (!cidr.contains("/")) throw new IllegalArgumentException("CIDR invalide : " + cidr);
    }
}
```

Records = parfaits pour DTO, résultats de requêtes, messages, clés de Map. Ils sont **implicitement final** et leurs composants sont **final**.

---

## 21. Packages : organiser le code

```text
src/main/java/
└── com/entreprise/supervision/
    ├── Main.java
    ├── model/
    │   ├── Equipement.java
    │   └── Onduleur.java
    ├── service/
    │   └── SupervisionService.java
    └── util/
        └── ReseauUtil.java
```

```java
package com.entreprise.supervision.model;   // 1re ligne du fichier

import com.entreprise.supervision.util.ReseauUtil;
import java.util.List;                       // les imports suivent le package

public class Equipement { /* ... */ }
```

Conventions : nom de package en minuscules, en **reverse-DNS** (`com.entreprise.projet`) ; un package ≈ un module fonctionnel (`model`, `service`, `dao`, `web`).

---

## 22. Modules Java (JPMS, Java 9+)

Pour les grosses applications : encapsulation au niveau package + dépendances explicites.

```java
// module-info.java à la racine des sources
module com.entreprise.supervision {
    requires java.sql;              // dépendance explicite
    requires spring.boot;           // (nom automatique si pas modulaire)
    exports com.entreprise.supervision.api;   // packages visibles de l'extérieur
    // com.entreprise.supervision.impl reste interne : inaccessible !
}
```

En pratique, beaucoup de projets (dont Spring Boot) s'en passent encore. **Conseil** : commencez sans modules ; adoptez JPMS quand le projet dépasse ~50 classes ou doit exposer une API publique propre.

---

## 23. Modificateurs de visibilité

| Modificateur | Classe | Package | Sous-classe | Monde |
|---|---|---|---|---|
| `private` | ✅ | ❌ | ❌ | ❌ |
| (défaut) | ✅ | ✅ | ❌ | ❌ |
| `protected` | ✅ | ✅ | ✅ | ❌ |
| `public` | ✅ | ✅ | ✅ | ✅ |

Autres modificateurs clés :
- `static` : appartient à la classe, pas à l'instance (méthodes utilitaires, constantes, compteurs partagés).
- `final` : variable = constante ; méthode = non redéfinissable ; classe = non héritable.
- `abstract` : classe non instanciable / méthode sans corps.
- `synchronized` : verrou d'exclusion mutuelle (voir section 40).
- `volatile` : garantit la visibilité inter-threads d'une variable (pas d'atomicité !).

**Règle d'or** : visibilité la plus restrictive possible. `private` par défaut, on élargit si besoin prouvé.

---

## 24. Génériques (Generics)

```java
// Sans génériques (avant Java 5) : casts partout, erreurs à l'exécution
List brute = new ArrayList();
brute.add("texte");
Integer i = (Integer) brute.get(0);  // ClassCastException à l'exécution !

// Avec génériques : vérifié À LA COMPILATION
List<String> noms = new ArrayList<>();
noms.add("UPS-01");
// noms.add(42);  // ERREUR de compilation : c'est le but !
String premier = noms.get(0);       // pas de cast

// Méthode générique
public static <T> T premierOuNull(List<T> liste) {
    return liste.isEmpty() ? null : liste.get(0);
}

// Classe générique
public class Paire<K, V> {
    private final K cle;
    private final V valeur;
    public Paire(K cle, V valeur) { this.cle = cle; this.valeur = valeur; }
    public K getCle() { return cle; }
    public V getValeur() { return valeur; }
}
Paire<String, Double> mesure = new Paire<>("temp-01", 21.5);

// Wildcards : <? extends T> lecture, <? super T> écriture (règle PECS)
public static double somme(List<? extends Number> nombres) {
    double total = 0;
    for (Number n : nombres) total += n.doubleValue();
    return total;
}
somme(List.of(1, 2.5, 3L));  // OK : Integer, Double, Long sont des Number
```

**Type erasure** : les génériques n'existent qu'à la compilation ; à l'exécution, `List<String>` = `List`. Conséquence : impossible de faire `new T()` ou `instanceof List<String>`.

**À retenir** : génériques partout où un conteneur/algorithme est indépendant du type contenu — zéro cast, erreurs détectées tôt.

---

## 25. Collections : vue d'ensemble

| Interface | Implémentations | Ordonné | Doublons | `null` | Accès |
|---|---|---|---|---|---|
| `List` | `ArrayList`, `LinkedList` | Oui (insertion) | Oui | Oui | Index |
| `Set` | `HashSet`, `LinkedHashSet`, `TreeSet` | Selon impl. | Non | 1 (sauf TreeSet) | — |
| `Map` | `HashMap`, `LinkedHashMap`, `TreeMap` | Selon impl. | Clés uniques | Oui (sauf TreeMap) | Clé |
| `Queue`/`Deque` | `ArrayDeque`, `PriorityQueue` | FIFO/priorité | Oui | Non (ArrayDeque) | Tête/queue |

```java
// Fabriques immutables (Java 9+) : parfaites pour constantes et tests
List<String> baies = List.of("A", "B", "C");       // immutable ! add() -> UnsupportedOperationException
Set<String> roles = Set.of("admin", "tech");
Map<String, Integer> ports = Map.of("http", 80, "https", 443);

// Copie défensive
List<String> modifiable = new ArrayList<>(baies);
```

**Choisir la bonne implémentation** : `ArrayList` par défaut ; `HashMap`/`HashSet` par défaut ; `TreeMap`/`TreeSet` si tri naturel nécessaire ; `LinkedHashMap` si ordre d'insertion à préserver (caches LRU).

---

## 26. List : ArrayList en détail

```java
List<String> onduleurs = new ArrayList<>();
onduleurs.add("UPS-01");
onduleurs.add("UPS-02");
onduleurs.add(0, "UPS-00");            // insertion à l'index 0
System.out.println(onduleurs.get(1));  // UPS-01
System.out.println(onduleurs.size()); // 3
onduleurs.set(1, "UPS-01B");           // remplacement
onduleurs.remove("UPS-02");           // par valeur
onduleurs.remove(0);                  // par index
System.out.println(onduleurs.contains("UPS-01B")); // true

// Tri
onduleurs.sort(String::compareTo);              // naturel
onduleurs.sort(Comparator.comparingInt(String::length)); // par longueur

// Conversion tableau <-> liste
String[] tableau = onduleurs.toArray(new String[0]);
List<String> depuis = new ArrayList<>(Arrays.asList(tableau));
```

`ArrayList` = tableau redimensionnable : accès index en O(1), insertion au milieu en O(n). `LinkedList` = liste chaînée : insertion/suppression en tête en O(1), mais accès index en O(n) — **rarement le bon choix** en pratique.

---

## 27. Set : unicité garantie

```java
Set<String> ipsVues = new HashSet<>();
ipsVues.add("10.0.0.1");
ipsVues.add("10.0.0.1");   // ignoré silencieusement
System.out.println(ipsVues.size()); // 1

