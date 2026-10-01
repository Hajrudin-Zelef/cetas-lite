---
id: collect-261001-rattrapage/rattrapage/java-guide-6
title: "Guide Java — du zéro au production"
domain: rattrapage
role: reference
task: reference
actors: []
dates: ["2026-09-26"]
keywords: []
source: docs/RAG/collect-261001-rattrapage/java_guide.md
source_anchor: ""
source_lines: [1149, 1342]
sha256: 62b0a42357b7e5811b7443e9b5297d03c7b127740e807059f642f48bd7f40bc0
---

# Guide Java — du zéro au production

// Filtrer + transformer + collecter
List<String> critiques = alarmes.stream()
    .filter(a -> a.criticite() >= 3)          // intermédiaire
    .map(Alarme::equipement)                  // intermédiaire
    .distinct()                               // intermédiaire
    .sorted()                                 // intermédiaire
    .toList();                                // terminale -> [UPS-01] (liste immutable)

// Regrouper
Map<String, List<Alarme>> parEquipement = alarmes.stream()
    .collect(Collectors.groupingBy(Alarme::equipement));

// Compter par code
Map<String, Long> parCode = alarmes.stream()
    .collect(Collectors.groupingBy(Alarme::code, Collectors.counting()));

// Statistiques
DoubleSummaryStatistics stats = alarmes.stream()
    .collect(Collectors.summarizingDouble(Alarme::criticite));
System.out.println("max=" + stats.getMax() + " moyenne=" + stats.getAverage());

// Réduction
int criticiteTotale = alarmes.stream().mapToInt(Alarme::criticite).sum();
Optional<Alarme> pire = alarmes.stream().max(Comparator.comparingInt(Alarme::criticite));

// anyMatch / allMatch / noneMatch / findFirst
boolean urgent = alarmes.stream().anyMatch(a -> a.criticite() >= 4);

// Chaînage plat : flatMap
List<String> tousLesCodes = sites.stream()
    .flatMap(site -> site.getAlarmes().stream())
    .map(Alarme::code)
    .toList();
```

**Règles d'or des streams** :
- Une opération intermédiaire ne fait **rien** sans opération terminale (lazy).
- Un stream est **à usage unique** : le réutiliser lève `IllegalStateException`.
- Pas d'effet de bord dans `map`/`filter` (pas de `list.add` dedans !).
- `parallelStream()` : seulement pour gros volumes + opérations coûteuses et sans état — mesurez avant/après, le parallélisme a un coût.

---

## 35. Optional : en finir avec le null (Java 8+)

```java
// ❌ Avant : null qui traîne, NPE à 3h du matin
Equipement e = trouverParNom(nom);
System.out.println(e.getIp());  // NullPointerException si absent !

// ✅ Après : l'absence est explicite dans le type
public Optional<Equipement> trouverParNom(String nom) {
    return equipements.stream().filter(e -> e.getNom().equals(nom)).findFirst();
}

trouverParNom("UPS-01")
    .ifPresent(e -> System.out.println(e.getIp()));            // action si présent

String ip = trouverParNom("UPS-99")
    .map(Equipement::getIp)                                    // transforme si présent
    .orElse("0.0.0.0");                                        // valeur par défaut

String ip2 = trouverParNom("UPS-99")
    .map(Equipement::getIp)
    .orElseThrow(() -> new EquipementInjoignableException("UPS-99")); // ou exception

// Filtrage
trouverParNom("UPS-01")
    .filter(e -> e.getEtat() == EtatEquipement.EN_SERVICE)
    .ifPresent(e -> System.out.println("OK : " + e.getNom()));

// orElse vs orElseGet : orElse ÉVALUE toujours son argument !
String v1 = opt.orElse(calculCouteux());      // calculCouteux() appelé même si opt présent !
String v2 = opt.orElseGet(() -> calculCouteux()); // lazy : appelé seulement si absent
```

**Règles** : `Optional` = type de **retour** (et chaînage), jamais attribut de classe, jamais paramètre de méthode, jamais `Optional.of(null)` (utilisez `ofNullable`). Ne remplace pas la validation d'arguments (`Objects.requireNonNull`).

---

## 36. Dates et heures : java.time (Java 8+)

Oubliez `Date`/`Calendar` (mutables, piégeux). `java.time` = immutable et thread-safe.

```java
import java.time.*;
import java.time.format.DateTimeFormatter;
import java.time.temporal.ChronoUnit;

LocalDate aujourd = LocalDate.now();                    // 2026-09-26
LocalTime midi = LocalTime.of(12, 30);
LocalDateTime releve = LocalDateTime.of(2026, 9, 26, 14, 0);
ZonedDateTime paris = ZonedDateTime.now(ZoneId.of("Europe/Paris"));
Instant instant = Instant.now();                        // timestamp UTC, parfait pour logs/BDD

// Calculs
LocalDate dans30j = aujourd.plusDays(30);
long jours = ChronoUnit.DAYS.between(derniereMaintenance, aujourd);
Duration panne = Duration.between(debut, fin);           // heures/min/sec
Period garantie = Period.ofYears(2);                    // années/mois/jours

// Formatage / parsing (DateTimeFormatter est thread-safe, contrairement à SimpleDateFormat)
DateTimeFormatter fmt = DateTimeFormatter.ofPattern("dd/MM/yyyy HH:mm");
System.out.println(releve.format(fmt));                 // 26/09/2026 14:00
LocalDateTime parse = LocalDateTime.parse("26/09/2026 14:00", fmt);

// Pour la monnaie : BigDecimal, JAMAIS double
import java.math.BigDecimal;
BigDecimal prix = new BigDecimal("19.99");              // constructeur String, pas double !
BigDecimal ttc = prix.multiply(new BigDecimal("1.20"));
System.out.println(ttc.setScale(2, RoundingMode.HALF_UP)); // 23.99
```

**À retenir** : `Instant` pour horodater/stocker, `ZonedDateTime` pour afficher à l'utilisateur, `BigDecimal(String)` pour l'argent.

---

## 37. Entrées-sorties fichiers : Path et Files (NIO.2)

```java
import java.nio.file.*;
import java.io.IOException;
import java.util.List;

Path p = Path.of("data", "inventaire.csv");   // construction portable (séparateurs gérés)
System.out.println(p.toAbsolutePath());
System.out.println(p.getFileName());          // inventaire.csv
System.out.println(p.getParent());            // data

// Lire
String contenu = Files.readString(p);                       // petit fichier texte
byte[] octets = Files.readAllBytes(p);                      // petit fichier binaire
List<String> lignes = Files.readAllLines(p);                // petites listes de lignes
try (var stream = Files.lines(p)) {                         // GROS fichier : stream lazy ligne par ligne
    stream.filter(l -> l.contains("UPS"))
          .forEach(System.out::println);
}

// Écrire
Files.writeString(p, "UPS-01,40\n", StandardOpenOption.CREATE, StandardOpenOption.APPEND);

// Tester / créer
if (Files.notExists(p)) Files.createDirectories(p.getParent());
System.out.println(Files.isRegularFile(p) + " " + Files.size(p));
System.out.println(Files.getLastModifiedTime(p));

// Copier / déplacer / supprimer
Files.copy(p, Path.of("data", "backup.csv"), StandardCopyOption.REPLACE_EXISTING);
Files.move(p, Path.of("archive", "inventaire.csv"));
Files.deleteIfExists(Path.of("tmp.txt"));

// Parcourir une arborescence
try (var walk = Files.walk(Path.of("data"))) {
    walk.filter(Files::isRegularFile)
        .filter(f -> f.toString().endsWith(".csv"))
        .forEach(System.out::println);
}
```

**À retenir** : `Files` + `Path` remplacent `File` dans tout code moderne ; `Files.lines`/`Files.walk` retournent des streams à fermer (try-with-resources).

---

## 38. NIO avancé : buffers, canaux, WatchService

```java
// Copie rapide via canaux (utile pour gros fichiers)
try (var in = FileChannel.open(Path.of("gros.iso"), StandardOpenOption.READ);
     var out = FileChannel.open(Path.of("copie.iso"),
             StandardOpenOption.CREATE, StandardOpenOption.WRITE)) {
    in.transferTo(0, in.size(), out);   // zéro-copy quand l'OS le permet
}

// Surveiller un répertoire (ex. : dossier de dépôt de fichiers d'inventaire)
try (WatchService watcher = FileSystems.getDefault().newWatchService()) {
    Path dir = Path.of("data/inbox");
    dir.register(watcher, StandardWatchEventKinds.ENTRY_CREATE);
    while (true) {
        WatchKey key = watcher.take();  // bloque jusqu'à un événement
        for (WatchEvent<?> ev : key.pollEvents()) {
            System.out.println("Nouveau fichier : " + ev.context());
            // -> déclencher le traitement d'import ici
        }
        if (!key.reset()) break;
    }
}
```

Pour 99 % des outils internes, `Files` suffit ; gardez `WatchService` pour les scénarios « dossier d'échange » (dépôt CSV/exports).

---

## 39. Threads : les bases

