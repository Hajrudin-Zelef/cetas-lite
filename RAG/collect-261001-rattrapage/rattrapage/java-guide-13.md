---
id: collect-261001-rattrapage/rattrapage/java-guide-13
title: "Guide Java — du zéro au production"
domain: rattrapage
role: reference
task: reference
actors: ["Lambda"]
dates: []
keywords: ["packaging"]
source: docs/RAG/collect-261001-rattrapage/java_guide.md
source_anchor: ""
source_lines: [2540, 2651]
sha256: 09ad2b29c5e031778c62170e65838667858b981a388a51cbad2bb3bc50725e09
---

# Guide Java — du zéro au production

| Terme | Définition |
|---|---|
| JDK | Kit de développement (compilateur + outils + runtime) |
| JRE | Runtime seul (plus distribué séparément depuis Java 11) |
| JVM | Machine virtuelle qui exécute le bytecode |
| Bytecode | Code intermédiaire des `.class`, indépendant de l'OS |
| JIT | Compilation « juste-à-temps » du bytecode en natif à chaud |
| Heap | Zone mémoire des objets, gérée par le GC |
| GC | Ramasse-miettes (G1, ZGC, Shenandoah, Parallel) |
| LTS | Version à support long (8, 11, 17, 21, 25) |
| Maven/Gradle | Outils de build et gestion de dépendances |
| POJO | Simple objet Java, sans contrainte de framework |
| DTO | Objet de transfert de données (souvent un record) |
| Boxing | Conversion auto `int` ↔ `Integer` (a un coût) |
| Erasure | Effacement des génériques à la compilation |
| Immutable | Objet non modifiable après création (String, records, List.of) |
| Checked exception | Exception que le compilateur force à gérer (`IOException`) |
| Unchecked exception | `RuntimeException` : bug probable, non forcée |
| Stream | Pipeline lazy de traitement de données (≠ `InputStream` !) |
| Lambda | Fonction anonyme `(a, b) -> a + b` |
| Thread virtuel | Thread léger JVM (Java 21+), idéal I/O massif |
| Deadlock | Deux threads qui s'attendent mutuellement (verrous croisés) |
| Race condition | Résultat dépendant de l'ordre d'exécution des threads |
| NPE | `NullPointerException` : déréférencement de `null` |
| JAR | Archive Java exécutable/portable |
| JPMS | Système de modules Java (`module-info.java`) |
| Javadoc | Documentation générée depuis les commentaires `/** */` |
| Réflexion | Inspection/manipulation des classes à l'exécution |
| Backoff | Attente croissante entre retries (exponentiel) |
| Idempotent | Opération répétable sans effet supplémentaire (GET, PUT) |

---

## 68. Quiz : 10 questions + réponses

**Q1.** Que vaut `System.out.println(5 / 2);` et pourquoi ?
<details><summary>Réponse</summary>Affiche <code>2</code> : division entière entre deux <code>int</code>, la partie décimale est tronquée. Pour 2.5 : <code>5 / 2.0</code>.</details>

**Q2.** Pourquoi ce test échoue-t-il parfois : `if (scanner.nextLine() == "oui")` ?
<details><summary>Réponse</summary><code>==</code> compare les références, pas le contenu. Deux String au contenu identique peuvent être des objets différents. Utilisez <code>"oui".equals(scanner.nextLine())</code>.</details>

**Q3.** Quelle exception survient ici, et comment la corriger ?
```java
List<String> l = new ArrayList<>(List.of("a", "b"));
for (String s : l) { if (s.equals("a")) l.remove(s); }
```
<details><summary>Réponse</summary><code>ConcurrentModificationException</code> : modification pendant itération. Corriger avec <code>l.removeIf(s -> s.equals("a"))</code> ou un <code>Iterator</code> explicite et <code>it.remove()</code>.</details>

**Q4.** Checked ou unchecked : `IOException` ? `NullPointerException` ? `SQLException` ?
<details><summary>Réponse</summary><code>IOException</code> et <code>SQLException</code> : checked (gestion ou <code>throws</code> obligatoire). <code>NullPointerException</code> : unchecked (<code>RuntimeException</code>).</details>

**Q5.** Que garantit `try-with-resources` ?
<details><summary>Réponse</summary>La fermeture automatique des ressources <code>AutoCloseable</code> (fichiers, sockets, connexions), même si une exception survient, dans l'ordre inverse d'ouverture.</details>

**Q6.** Quelle est la différence entre `orElse` et `orElseGet` sur un `Optional` ?
<details><summary>Réponse</summary><code>orElse(valeur)</code> évalue son argument <strong>toujours</strong> (même si l'Optional est présent). <code>orElseGet(() -> ...)</code> n'appelle le supplier que si l'Optional est vide — à préférer pour les valeurs coûteuses.</details>

**Q7.** Pourquoi `map.get(cle)` peut-il être dangereux sans précaution ?
<details><summary>Réponse</summary>Il retourne <code>null</code> si la clé est absente → risque de NPE en chaîne. Utilisez <code>getOrDefault</code>, <code>computeIfAbsent</code>, ou testez <code>containsKey</code>.</details>

**Q8.** Dans quel cas préférer un virtual thread (Java 21+) à un thread plateforme ?
<details><summary>Réponse</summary>Tâches I/O-bound nombreuses (milliers de connexions/pings/appels HTTP) avec code bloquant simple. Pas de gain pour du calcul CPU pur.</details>

**Q9.** Pourquoi ce code est-il une faille de sécurité ?
```java
String sql = "SELECT * FROM users WHERE login = '" + login + "'";
Statement st = cx.createStatement();
st.executeQuery(sql);
```
<details><summary>Réponse</summary>Injection SQL : un login comme <code>' OR '1'='1</code> détourne la requête. Corriger avec <code>PreparedStatement</code> et paramètres <code>?</code>.</details>

**Q10.** Que fait ce code et quel est le piège ?
```java
String total = "";
for (String s : lignes) total += s + "\n";
```
<details><summary>Réponse</summary>Concatène les lignes, mais chaque <code>+=</code> crée un nouveau String (O(n²)). Sur gros volume : <code>StringBuilder</code> ou <code>String.join("\n", lignes)</code>.</details>

---

## 69. Pour aller plus loin

### Par priorité pour un responsable systèmes

1. **Spring Boot** — docs.spring.io : votre porte d'entrée vers les API REST, Spring Data JPA (BDD sans SQL manuel), Spring Security, Actuator (health/metrics pour votre supervision).
2. **Concurrence avancée** — *Java Concurrency in Practice* (Goetz) : la référence, toujours d'actualité sur les fondamentaux.
3. **JVM en production** — apprenez `jstat`/`jfr` (Java Flight Recorder : profiling sans surcoût, intégré au JDK).
4. **Protocoles métier** — bibliothèques : SNMP4J (SNMP), j2mod (Modbus), Paho (MQTT) : parfaites pour parler à onduleurs et équipements.
5. **Tests** — Testcontainers (vraie BDD Docker dans les tests), AssertJ (assertions lisibles).

### Feuille de route 6 mois (2 h/semaine)

| Mois | Objectif | Livrable |
|---|---|---|
| 1 | Bases + collections + exceptions | Outil CLI inventaire (cas pratique 1) |
| 2 | Streams, Optional, NIO, JDBC | Import CSV→BDD (cas pratique 3) |
| 3 | JUnit/Mockito, logging | Couverture > 70 % sur vos outils |
| 4 | Spring Boot REST | API d'inventaire exposée (cas pratique 2 + section 48) |
| 5 | Concurrence, virtual threads | Collecteur parallèle d'alarmes |
| 6 | Packaging, Docker, runbook | Outil packagé + déployé avec checklist section 65 |

### Ressources

- Documentation officielle : **docs.oracle.com/javase**
- Tutoriels : **dev.java** (site officiel, moderne)
- Pratique : **exercism.org** (track Java, mentorat gratuit)
- Questions : **stackoverflow.com** (tag java)
- Versions : adoptez une LTS et suivez les notes de version d'**OpenJDK**

---

*Fin du guide — bon code, et que vos NPE soient rares et vos builds verts.* 🍵
