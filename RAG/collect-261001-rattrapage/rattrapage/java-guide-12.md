---
id: collect-261001-rattrapage/rattrapage/java-guide-12
title: "Guide Java — du zéro au production"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["arr"]
source: docs/RAG/collect-261001-rattrapage/java_guide.md
source_anchor: ""
source_lines: [2364, 2539]
sha256: 9a83ccf41334c002500ea49c0c646a01dfce71a5d41f3dda3cfafa52d35006b9
---

# Guide Java — du zéro au production

    public static void main(String[] args) throws Exception {
        String url = System.getenv("JDBC_URL");       // jamais en dur
        String user = System.getenv("JDBC_USER");
        String pass = System.getenv("JDBC_PASSWORD");

        try (Connection cx = DriverManager.getConnection(url, user, pass);
             Stream<String> lignes = Files.lines(Path.of(args[0]))) {

            cx.setAutoCommit(false);
            String sql = "INSERT INTO equipement(nom, puissance_kva, etat) VALUES (?, ?, ?) " +
                         "ON CONFLICT (nom) DO UPDATE SET puissance_kva = EXCLUDED.puissance_kva";
            try (PreparedStatement ps = cx.prepareStatement(sql)) {
                final int[] compteur = {0};
                lignes.filter(l -> !l.isBlank()).map(l -> l.split(";")).forEach(c -> {
                    try {
                        ps.setString(1, c[0].trim());
                        ps.setDouble(2, Double.parseDouble(c[1].trim()));
                        ps.setString(3, c[2].trim());
                        ps.addBatch();
                        if (++compteur[0] % 500 == 0) ps.executeBatch(); // flush tous les 500
                    } catch (SQLException e) { throw new RuntimeException(e); }
                });
                ps.executeBatch();  // reste
                cx.commit();
                System.out.println("Importé : " + compteur[0] + " lignes");
            } catch (Exception e) {
                cx.rollback();
                throw e;
            } finally {
                cx.setAutoCommit(true);
            }
        }
    }
}
```

**Points à noter** : batch (`addBatch`/`executeBatch`) = 10 à 100× plus rapide que des inserts un par un ; transaction unique + rollback ; upsert PostgreSQL (`ON CONFLICT`).

---

## 64. Cas pratique 4 — Planificateur de relevés avec logs et arrêt propre

```java
import java.time.*;
import java.util.concurrent.*;
import org.slf4j.*;

/** Toutes les 5 min : relève les compteurs, logue, s'arrête proprement sur SIGTERM. */
public class Planificateur {

    private static final Logger log = LoggerFactory.getLogger(Planificateur.class);

    public static void main(String[] args) {
        ScheduledExecutorService planif = Executors.newSingleThreadScheduledExecutor(
            r -> { Thread t = new Thread(r, "releve"); t.setDaemon(false); return t; });

        // Arrêt propre : flush + fin du tour en cours (max 30 s)
        Runtime.getRuntime().addShutdownHook(new Thread(() -> {
            log.info("Arrêt demandé, fin du relevé en cours...");
            planif.shutdown();
            try {
                if (!planif.awaitTermination(30, TimeUnit.SECONDS)) planif.shutdownNow();
            } catch (InterruptedException e) {
                planif.shutdownNow();
                Thread.currentThread().interrupt();
            }
            log.info("Arrêt terminé.");
        }, "shutdown-hook"));

        planif.scheduleAtFixedRate(() -> {
            try {
                Instant debut = Instant.now();
                int n = releverCompteurs();   // votre logique métier
                log.info("Relevé OK : {} équipements en {} ms",
                    n, Duration.between(debut, Instant.now()).toMillis());
            } catch (Exception e) {
                log.error("Relevé en échec", e);  // on logue, le planificateur continue
            }
        }, 0, 5, TimeUnit.MINUTES);

        log.info("Planificateur démarré (Ctrl+C pour arrêter).");
    }

    private static int releverCompteurs() {
        // ... SNMP / Modbus / HTTP vers les équipements ...
        return 42;
    }
}
```

**Points à noter** : `addShutdownHook` = indispensable dans les conteneurs (SIGTERM) ; le `catch` dans la tâche empêche le planificateur de mourir silencieusement après une exception ; nommer les threads aide énormément au diagnostic (`jstack`).

---

## 65. Checklists : nouveau projet & mise en production

### Nouveau projet

- [ ] JDK LTS (17/21) + `maven.compiler.release` aligné
- [ ] `.gitignore` (target/, .idea/, *.iml, .DS_Store)
- [ ] `README.md` : prérequis, build (`mvn package`), lancement, config (variables d'env.)
- [ ] Encodage UTF-8 forcé (`project.build.sourceEncoding`)
- [ ] Dépendances : JUnit 5 + Mockito + SLF4J/Logback dès le jour 1
- [ ] Formateur de code configuré (Spotless)
- [ ] `main` minimal qui démarre et logue sa version

### Mise en production

- [ ] `-Xmx` / `MaxRAMPercentage` réglé pour le conteneur
- [ ] Secrets via variables d'environnement ou coffre (jamais dans l'image)
- [ ] Logs vers stdout **ou** fichier rotatif (pas les deux en double), niveau INFO par défaut
- [ ] Healthcheck (`/actuator/health` en Spring Boot, ou endpoint `/health` maison)
- [ ] Shutdown hook (section 64) pour SIGTERM propre
- [ ] Heap dump on OOM (`-XX:+HeapDumpOnOutOfMemoryError`)
- [ ] Sauvegarde/restauration testée si état local (BDD embarquée, fichiers)
- [ ] Versions épinglées (Docker tag précis, pas `latest`)
- [ ] Runbook : où sont les logs, comment redémarrer, qui appeler

---

## 66. Pense-bête de poche

```text
COMPILATION / EXÉCUTION
  javac Main.java && java Main        java Main.java (script, Java 11+)
  mvn compile | mvn package | mvn test
  jshell                              REPL interactif

BASE
  var x = ...;        inférence (init obligatoire)
  "a".equals(b)       comparaison chaînes (jamais ==)
  s = s.trim();       String immutable : réassigner !
  (double) a / b      éviter la division entière
  new BigDecimal("1.10")  monnaie (jamais double, jamais new BigDecimal(1.1))

COLLECTIONS
  List.of / Set.of / Map.of   immutables (Java 9+)
  new ArrayList<>(...)         modifiable
  map.getOrDefault(k, def)     map.merge(k, 1, Integer::sum)
  list.removeIf(p)             suppression sûre pendant parcours

EXCEPTIONS
  try (var r = ...) { }       try-with-resources : TOUJOURS pour I/O, JDBC, sockets
  catch du spécifique au général
  throw new IllegalArgumentException("...")   validation d'arguments

STREAMS
  .stream().filter().map().collect(...)   .toList() = immutable
  Collectors.groupingBy / counting / summarizingDouble
  Optional : map / filter / orElse / orElseGet / orElseThrow / ifPresent

CONCURRENCE
  Executors.newVirtualThreadPerTaskExecutor()  I/O massif (Java 21+)
  Executors.newFixedThreadPool(n)              CPU-bound : n = nb cœurs
  shutdown() + awaitTermination                TOUJOURS arrêter les pools
  volatile = flag visible   synchronized/AtomicX = lire-modifier-écrire

DATES
  Instant.now()  horodatage UTC   LocalDate.now()  ZonedDateTime (affichage)
  DateTimeFormatter.ofPattern("dd/MM/yyyy HH:mm")  (thread-safe)

RÉSEAU / BDD
  HttpClient.newBuilder()...build()   UNE instance partagée
  PreparedStatement + ?               jamais de concat SQL
  HikariCP                            pool de connexions en prod

DEBUG
  jstack <pid>   threads    jmap -histo <pid>   heap
  -XX:+HeapDumpOnOutOfMemoryError
  Lire la stack trace du HAUT + première ligne "Caused by"
```

---

## 67. Glossaire

