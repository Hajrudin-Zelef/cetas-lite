---
id: collect-261001-rattrapage/rattrapage/java-guide-9
title: "Guide Java — du zéro au production"
domain: rattrapage
role: reference
task: reference
actors: ["Google", "Oracle"]
dates: []
keywords: ["benchmark"]
source: docs/RAG/collect-261001-rattrapage/java_guide.md
source_anchor: ""
source_lines: [1790, 1946]
sha256: ca131888f02631f63aa04685f903853bdeda3954a7811ecbb34107d467d2f09a
---

# 3. Installateur natif (Java 14+) : .deb, .msi, .dmg avec runtime embarqué
jpackage --name outil-inventaire --input target/ --main-jar outil-1.0.0.jar \
         --main-class com.entreprise.outil.Main --type deb \
         --dest installateurs/
```

**En pratique entreprise** : jar « fat » (shade) pour les outils internes simples ; image Docker `eclipse-temurin:21-jre` + jar pour les services ; `jpackage` quand il faut un installateur livré à des non-techniciens.

---

## 50. La JVM : comprendre la machine

Le bytecode `.class` s'exécute sur la JVM, qui le compile **à chaud** (JIT) en code natif après quelques milliers d'appels. Conséquences pratiques :

- **Warmup** : les premières secondes/minutes sont plus lentes (montée en charge progressive en prod, pas de benchmark à froid).
- **Heap** : zone mémoire des objets, gérée par le GC. Dimensionnée via `-Xms` (initial) et `-Xmx` (max).
- **Stack** : une pile par thread (appels de méthodes, variables locales). `-Xss` règle sa taille.
- **Metaspace** : définitions des classes.

```bash
# Réglages mémoire typiques pour un service (ex. : 2 Go max, GC G1 par défaut)
java -Xms512m -Xmx2g -jar app.jar

# Java 21+ : dimensionnement automatique raisonnable dans les conteneurs (détecte les cgroups)
java -XX:MaxRAMPercentage=75.0 -jar app.jar   # utilise 75 % de la RAM du conteneur
```

**À retenir** : en conteneur, fixez `-Xmx` ou `MaxRAMPercentage` — sinon la JVM peut se croire sur une machine à 128 Go et se faire tuer par l'OOM-killer.

---

## 51. Garbage Collector : l'essentiel

| GC | Depuis | Profil |
|---|---|---|
| G1 (défaut) | Java 9 | Équilibré, pauses < ~200 ms, bon choix par défaut |
| ZGC | Java 15 (prod) | Pauses < 1 ms, heaps énormes (To) |
| Shenandoah | Java 15 | Pauses courtes, bon sur heaps moyens |
| Parallel | Toujours | Débit max, batchs (pauses longues acceptables) |

```bash
java -XX:+UseZGC -Xmx8g -jar app.jar        # choisir ZGC
java -Xlog:gc*:file=/var/log/app/gc.log:time,uptime,level,tags -jar app.jar
```

**Fuites mémoire en Java ?** Oui : références conservées par erreur (cache sans éviction, listeners jamais désinscrits, `ThreadLocal` oublié, collections static qui grossissent). Symptôme : heap qui monte en dents de scie jusqu'à `OutOfMemoryError`. Diagnostic : `jmap -histo`, heap dump + analyse (Eclipse MAT).

**Bonnes pratiques** :
- Ne faites **jamais** `System.gc()` en production (force un GC complet, sauf cas très particulier).
- Réutilisez les objets coûteux via des pools uniquement si mesuré nécessaire (le GC moderne alloue très vite).
- Préférez les objets **courte durée de vie** : le GC adore ça (collectés jeunes, presque gratuits).

---

## 52. Performance : mesurer avant d'optimiser

**Règle n°1** : pas d'optimisation sans mesure. 97 % du temps, le goulot n'est pas là où vous croyez.

```bash
# Micro-benchmark sérieux : JMH (jamais de boucle System.nanoTime() artisanale !)
# Dépendance : org.openjdk.jmh:jmh-core:1.37
```

Checklist des gains faciles (par ordre d'impact typique) :

1. **I/O** : requêtes SQL N+1, appels HTTP en série → batcher / paralléliser (virtual threads).
2. **Algorithmique** : `List.contains` en O(n) dans une boucle → `HashSet` en O(1).
3. **Concaténation String en boucle** → `StringBuilder` (section 8).
4. **Boxing** : `Stream<Integer>` + `mapToInt` plutôt que `map` (évite `Integer` ↔ `int`).
5. **Logs** : `logger.debug("x=" + couteux())` évalue même si le niveau est INFO → utilisez `logger.debug("x={}", () -> couteux())` ou gardez la concaténation hors du log.
6. **Collections dimensionnées** : `new ArrayList<>(tailleConnue)`, `new HashMap<>(n)` évitent les reallocations.
7. **Regex précompilées** : `Pattern.compile` une fois en `static final`, pas dans la boucle.

```java
// ❌ Lent : boxing + contains O(n) + regex recompilée à chaque tour
for (String ip : ips) {
    if (listeNoire.contains(ip) && ip.matches("\\d+\\.\\d+\\.\\d+\\.\\d+")) alerter(ip);
}

// ✅ Rapide : HashSet O(1) + Pattern précompilé
private static final Pattern IPV4 = Pattern.compile("\\d+\\.\\d+\\.\\d+\\.\\d+");
private final Set<String> listeNoire = new HashSet<>(...);
for (String ip : ips) {
    if (IPV4.matcher(ip).matches() && listeNoire.contains(ip)) alerter(ip);
}
```

---

## 53. Sécurité : les réflexes

| Risque | Contre-mesure |
|---|---|
| Injection SQL | `PreparedStatement` systématique (section 44), jamais de concaténation |
| Injection de commande | `ProcessBuilder` avec liste d'arguments, jamais `Runtime.exec(String)` avec entrée utilisateur |
| Path traversal | Validez que `chemin.normalize().startsWith(racineAutorisee)` |
| XXE (XML) | Désactivez DTD/entités externes sur le parser |
| Désérialisation Java native | **Bannissez** `ObjectInputStream` sur données non fiables (gadget chains) ; préférez JSON |
| Secrets en dur | Variables d'environnement / coffre (Vault), jamais dans le code ni le jar |
| Dépendances vulnérables | `mvn dependency:check` / OWASP Dependency-Check / Dependabot en CI |
| TLS | `HttpsURLConnection`/`HttpClient` vérifient par défaut ; ne désactivez **jamais** la vérif. du certificat « pour tester » en prod |
| Mots de passe | BCrypt/Argon2 (jamais MD5/SHA-1, jamais en clair) |

```java
// ✅ Lecture d'un secret : variable d'environnement
String dbPassword = System.getenv("DB_PASSWORD");
if (dbPassword == null) throw new IllegalStateException("DB_PASSWORD non défini");

// ✅ Path traversal : confinement dans un répertoire
Path racine = Path.of("/srv/rapports").toAbsolutePath();
Path demande = racine.resolve(nomFichier).normalize();
if (!demande.startsWith(racine)) throw new SecurityException("Chemin interdit : " + nomFichier);

// ❌ INTERDIT : désactiver la vérification TLS
// TrustManager qui accepte tout -> faille béante. Ne jamais commiter ça.
```

**Hygiène** : `mvn versions:display-dependency-updates` régulièrement ; image Docker minimale ; principe du moindre privilège (utilisateur non-root dans le conteneur).

---

## 54. Bonnes pratiques & style de code

### Style (Google Java Style / conventions Oracle)

- Indentation 4 espaces (ou 2 — mais **constant** dans le projet) ; accolades même ligne pour les blocs.
- 120 colonnes max ; un import par ligne, pas de `.*` (sauf tests).
- Noms explicites : `equipementsHorsService`, pas `list2`.
- Formateur automatique : **Spotless** ou le formateur IntelliJ appliqué à chaque commit (via pre-commit hook).

### Principes

- **DRY** : factorisez la 3e duplication, pas la 1re.
- **KISS** : la solution simple et lisible bat la solution « intelligente ».
- **Fail fast** : `Objects.requireNonNull(param, "param")` en tête de méthode publique.
- **Immutabilité** : `final` champs, records, collections non modifiables renvoyées (`List.copyOf`).
- **Pas de `null` silencieux** : `Optional`, listes vides, exceptions explicites.
- **Javadoc** sur toute API publique : `@param`, `@return`, `@throws`.

```java
/**
 * Calcule l'autonomie estimée en minutes.
 *
 * @param chargeKva charge actuelle en kVA, doit être &gt; 0
 * @param capaciteAh capacité batterie en Ah, doit être &gt; 0
 * @return autonomie en minutes, arrondie à l'entier inférieur
 * @throws IllegalArgumentException si un paramètre est &lt;= 0
 */
public static long autonomieMinutes(double chargeKva, double capaciteAh) {
    Objects.requireNonNull(chargeKva, "chargeKva"); // pour objets ; ici double -> check manuel
    if (chargeKva <= 0 || capaciteAh <= 0) throw new IllegalArgumentException("paramètres > 0 requis");
    return (long) (capaciteAh * 12 * 60 / (chargeKva * 1000));
}
```

### Checklist revue de code

