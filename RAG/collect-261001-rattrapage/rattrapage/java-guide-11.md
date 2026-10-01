---
id: collect-261001-rattrapage/rattrapage/java-guide-11
title: "Guide Java — du zéro au production"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/java_guide.md
source_anchor: ""
source_lines: [2159, 2363]
sha256: 300289c8f4022f20b8a2596e96fb8cec5c51b02efd602029cf2d249650179973
---

# Guide Java — du zéro au production

```java
// Dépendance Maven : com.fasterxml.jackson.core:jackson-databind:2.17.x
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.datatype.jsr310.JavaTimeModule;

record EquipementJson(String nom, double puissanceKva, EtatEquipement etat, Instant maj) {}

ObjectMapper mapper = new ObjectMapper();
mapper.registerModule(new JavaTimeModule());  // support java.time
// mapper.configure(DeserializationFeature.FAIL_ON_UNKNOWN_PROPERTIES, false); // tolérant

// Sérialiser
String json = mapper.writeValueAsString(new EquipementJson("UPS-01", 40.0, EtatEquipement.EN_SERVICE, Instant.now()));
String joli = mapper.writerWithDefaultPrettyPrinter().writeValueAsString(obj);

// Désérialiser
EquipementJson e = mapper.readValue(json, EquipementJson.class);

// Structures libres
JsonNode node = mapper.readTree(json);
System.out.println(node.get("nom").asText());
```

**À retenir** : `ObjectMapper` est thread-safe après configuration → **une instance partagée** (static final). Jackson + records = combo idéal pour les DTO d'API.

---

## 60. Logging : SLF4J + Logback

Ne faites **jamais** `System.out.println` en production : pas de niveaux, pas de fichier, pas de rotation.

```java
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

public class Collecteur {
    private static final Logger log = LoggerFactory.getLogger(Collecteur.class);

    public void interroger(String ip) {
        log.info("Interrogation de {}", ip);                    // {} = placeholder, pas de concaténation
        try {
            // ...
        } catch (Exception e) {
            log.warn("Échec sur {}, nouvel essai planifié", ip, e); // stack trace en dernier arg
        }
        log.debug("Détail trame : {}", () -> trameHexaCouteuse()); // supplier lazy (Logback 1.3+)
    }
}
```

`logback.xml` minimal (console + fichier rotatif) :
```xml
<configuration>
  <appender name="FICHIER" class="ch.qos.logback.core.rolling.RollingFileAppender">
    <file>/var/log/collecteur/app.log</file>
    <rollingPolicy class="ch.qos.logback.core.rolling.TimeBasedRollingPolicy">
      <fileNamePattern>/var/log/collecteur/app.%d{yyyy-MM-dd}.log</fileNamePattern>
      <maxHistory>30</maxHistory>
    </rollingPolicy>
    <encoder><pattern>%d{HH:mm:ss.SSS} %-5level %logger{36} - %msg%n</pattern></encoder>
  </appender>
  <root level="INFO">
    <appender-ref ref="FICHIER"/>
  </root>
  <logger name="com.entreprise" level="DEBUG"/>
</configuration>
```

Niveaux : `ERROR` (panne) > `WARN` (anormal mais géré) > `INFO` (jalons métier) > `DEBUG` (diagnostic) > `TRACE` (très verbeux).

---

## 61. Cas pratique 1 — Outil CLI : inventaire réseau depuis un CSV

```java
import java.nio.file.*;
import java.util.*;
import java.util.stream.*;

/** Lit un CSV "nom;ip;puissance_kva", ping chaque hôte en parallèle (virtual threads),
 *  et écrit un rapport. Usage : java Inventaire.java inventaire.csv rapport.txt */
public class Inventaire {

    record Equipement(String nom, String ip, double puissanceKva) {}

    public static void main(String[] args) throws Exception {
        if (args.length < 2) {
            System.err.println("Usage : java Inventaire.java <entree.csv> <rapport.txt>");
            System.exit(2);
        }
        List<Equipement> equipements;
        try (Stream<String> lignes = Files.lines(Path.of(args[0]))) {
            equipements = lignes
                .filter(l -> !l.isBlank() && !l.startsWith("#"))
                .map(l -> l.split(";"))
                .filter(c -> c.length == 3)
                .map(c -> new Equipement(c[0].trim(), c[1].trim(), Double.parseDouble(c[2].trim())))
                .toList();
        }

        // Ping parallèle : un virtual thread par hôte (Java 21+)
        Map<Equipement, Boolean> resultats = new ConcurrentHashMap<>();
        try (var executor = Executors.newVirtualThreadPerTaskExecutor()) {
            for (Equipement e : equipements) {
                executor.submit(() -> resultats.put(e, ping(e.ip())));
            }
        } // attend la fin de tous les pings

        long joignables = resultats.values().stream().filter(b -> b).count();
        double puissanceTotale = equipements.stream().mapToDouble(Equipement::puissanceKva).sum();

        String rapport = """
            Rapport d'inventaire — %s
            =====================================
            Équipements : %d (joignables : %d, injoignables : %d)
            Puissance totale : %.1f kVA
            Injoignables :
            %s
            """.formatted(
                java.time.LocalDate.now(),
                equipements.size(), joignables, equipements.size() - joignables,
                puissanceTotale,
                resultats.entrySet().stream()
                    .filter(en -> !en.getValue())
                    .map(en -> "  - " + en.getKey().nom() + " (" + en.getKey().ip() + ")")
                    .collect(Collectors.joining("\n")));

        Files.writeString(Path.of(args[1]), rapport);
        System.out.println(rapport);
    }

    private static boolean ping(String ip) {
        try {
            return java.net.InetAddress.getByName(ip).isReachable(2000);
        } catch (Exception e) {
            return false;
        }
    }
}
```

**Ce que ce cas montre** : NIO (`Files.lines`), records, streams, text block `.formatted()`, virtual threads, `ConcurrentHashMap`. Un outil interne réel tient en ~70 lignes.

---

## 62. Cas pratique 2 — Mini client d'API REST avec retry

```java
import java.net.URI;
import java.net.http.*;
import java.time.Duration;

/** Interroge une API de supervision avec retry exponentiel. */
public class ApiClient {

    private static final HttpClient CLIENT = HttpClient.newBuilder()
        .connectTimeout(Duration.ofSeconds(5))
        .build();
    private static final int MAX_ESSAIS = 4;

    public static String getAvecRetry(String url) throws Exception {
        HttpRequest requete = HttpRequest.newBuilder()
            .uri(URI.create(url))
            .timeout(Duration.ofSeconds(10))
            .header("Accept", "application/json")
            .GET().build();

        Exception dernierEchec = null;
        for (int essai = 1; essai <= MAX_ESSAIS; essai++) {
            try {
                HttpResponse<String> rep = CLIENT.send(requete, HttpResponse.BodyHandlers.ofString());
                if (rep.statusCode() == 200) return rep.body();
                if (rep.statusCode() >= 400 && rep.statusCode() < 500)
                    throw new IllegalStateException("Erreur client HTTP " + rep.statusCode());
                dernierEchec = new IllegalStateException("HTTP " + rep.statusCode());
            } catch (java.io.IOException | InterruptedException e) {
                dernierEchec = e;   // réseau : on réessaie
                if (e instanceof InterruptedException) Thread.currentThread().interrupt();
            }
            long attenteMs = (long) (500 * Math.pow(2, essai - 1)); // backoff exponentiel
            Thread.sleep(attenteMs);
        }
        throw new IllegalStateException("Échec après " + MAX_ESSAIS + " essais", dernierEchec);
    }

    public static void main(String[] args) throws Exception {
        System.out.println(getAvecRetry("http://localhost:8080/api/equipements"));
    }
}
```

**Points à noter** : on ne retry **jamais** les 4xx (erreur du client, inutile) ; backoff exponentiel pour ne pas marteler ; `InterruptedException` → restaurer le flag d'interruption.

---

## 63. Cas pratique 3 — Import CSV vers base avec JDBC batch

```java
import java.nio.file.*;
import java.sql.*;
import java.util.stream.Stream;

/** Importe equipements.csv (nom;puissance;etat) en base via batch JDBC. */
public class ImportJdbc {

