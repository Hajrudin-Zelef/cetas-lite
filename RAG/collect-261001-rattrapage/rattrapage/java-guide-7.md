---
id: collect-261001-rattrapage/rattrapage/java-guide-7
title: "Guide Java — du zéro au production"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["arr"]
source: docs/RAG/collect-261001-rattrapage/java_guide.md
source_anchor: ""
source_lines: [1343, 1553]
sha256: c43fa6a8c6ae1f26360f85dca29bb83fd624cf6d984d38b98b4d0f22dfff8fe3
---

# Guide Java — du zéro au production

```java
// 1. Via Thread + lambda
Thread t = new Thread(() -> {
    System.out.println("Tâche dans " + Thread.currentThread().getName());
});
t.start();          // démarre (NE JAMAIS appeler run() directement !)
t.join();           // attend la fin

// 2. Via Runnable nommé (préférable : sépare la tâche du mécanisme)
Runnable inventaire = () -> scannerSousReseau("10.0.0.0/24");
new Thread(inventaire, "scan-reseau").start();

// 3. Via Callable + Future : tâche qui RETOURNE un résultat
ExecutorService pool = Executors.newFixedThreadPool(4);
Future<Integer> futur = pool.submit(() -> compterEquipements());
System.out.println("Équipements : " + futur.get());  // bloque jusqu'au résultat
pool.shutdown();

// États d'un thread : NEW -> RUNNABLE -> (BLOCKED/WAITING/TIMED_WAITING) -> TERMINATED
```

**Erreurs fréquentes** : appeler `run()` au lieu de `start()` (exécution synchrone dans le thread courant !) ; oublier `shutdown()` sur un pool (la JVM ne se termine pas).

---

## 40. synchronized, volatile, et les pièges classiques

```java
public class Compteur {
    private int valeur = 0;

    // ❌ Sans synchronized : deux threads peuvent lire/écrire en même temps -> comptes perdus
    // ✅ Avec : un seul thread à la fois dans la méthode (verrou = l'instance)
    public synchronized void incrementer() { valeur++; }

    public synchronized int getValeur() { return valeur; }
}

// Bloc synchronisé ciblé (verrouille moins longtemps = mieux)
private final Object verrou = new Object();
public void ajouterMesure(double m) {
    // ... calculs non critiques hors verrou ...
    synchronized (verrou) {
        mesures.add(m);   // section critique courte
    }
}

// volatile : visibilité immédiate inter-threads (mais PAS d'atomicité)
private volatile boolean arretDemande = false;
// Thread worker : while (!arretDemande) { ... }  -> voit le changement sans synchronized
```

**Règle** : `volatile` pour un flag ; `synchronized` (ou `java.util.concurrent`) pour toute opération lire-modifier-écrire (`i++`, `check-then-act`). Pour les compteurs : `AtomicInteger`/`AtomicLong` (sans verrou, plus rapides).

```java
import java.util.concurrent.atomic.AtomicLong;
private final AtomicLong requetes = new AtomicLong();
requetes.incrementAndGet();   // atomique, sans synchronized
```

---

## 41. ExecutorService : la bonne façon de gérer des threads

Ne créez **jamais** des `Thread` à la main en production : utilisez des pools.

```java
import java.util.concurrent.*;
import java.util.List;

ExecutorService pool = Executors.newFixedThreadPool(8);  // 8 threads réutilisés
try {
    List<Callable<String>> taches = List.of(
        () -> ping("10.0.0.11"),
        () -> ping("10.0.0.12"),
        () -> interrogerSnmp("10.0.0.13")
    );
    // invokeAll : lance tout, attend tout (avec timeout global possible)
    List<Future<String>> resultats = pool.invokeAll(taches, 30, TimeUnit.SECONDS);
    for (Future<String> f : resultats) {
        System.out.println(f.isCancelled() ? "TIMEOUT" : f.get());
    }
} finally {
    pool.shutdown();                                  // 1. refuse nouvelles tâches
    if (!pool.awaitTermination(10, TimeUnit.SECONDS)) {
        pool.shutdownNow();                           // 2. force l'arrêt si ça traîne
    }
}

// Planification (cron-like en mémoire)
ScheduledExecutorService planif = Executors.newSingleThreadScheduledExecutor();
planif.scheduleAtFixedRate(() -> releverCompteurs(), 0, 5, TimeUnit.MINUTES);
```

**Choisir la taille du pool** : tâches I/O-bound (réseau, disque) → beaucoup de threads (50-200) ; tâches CPU-bound → `Runtime.getRuntime().availableProcessors()` threads.

---

## 42. Virtual threads : la révolution Java 21+

Problème historique : 1 thread plateforme = ~1 Mo de pile → 10 000 connexions simultanées = 10 Go. Les **virtual threads** sont légers (quelques Ko) et gérés par la JVM.

```java
// AVANT (Java 17) : pool limité, code async complexe pour scaler
// APRÈS (Java 21+) : un thread virtuel par tâche, code BLOQUANT simple qui scale

// Lancement simple
Thread.startVirtualThread(() -> {
    System.out.println("Je suis virtuel : " + Thread.currentThread());
});

// Executor de virtual threads : remplace le pool pour les tâches I/O-bound
try (var executor = Executors.newVirtualThreadPerTaskExecutor()) {
    for (String ip : toutesLesIp) {
        executor.submit(() -> ping(ip));   // 10 000 tâches = OK, code bloquant simple
    }
} // fermeture auto : attend la fin des tâches

// Dans du code existant : remplacez juste la fabrique
ExecutorService pool = Executors.newFixedThreadPool(200);          // avant
ExecutorService vt = Executors.newVirtualThreadPerTaskExecutor(); // après (Java 21+)
```

**Quand les utiliser** : I/O-bound massif (scans réseau, collecteurs SNMP, serveurs HTTP) → idéal. **Quand les éviter** : calcul CPU intensif (aucun gain), code avec `synchronized` bloquant long (peut « pinner » le thread porteur — préférez `ReentrantLock`).

---

## 43. Réseau et HTTP : HttpClient (Java 11+)

```java
import java.net.URI;
import java.net.http.*;
import java.time.Duration;

HttpClient client = HttpClient.newBuilder()
    .connectTimeout(Duration.ofSeconds(5))
    .followRedirects(HttpClient.Redirect.NORMAL)
    .build();

// GET simple
HttpRequest requete = HttpRequest.newBuilder()
    .uri(URI.create("https://api.exemple.com/equipements"))
    .header("Accept", "application/json")
    .timeout(Duration.ofSeconds(10))
    .GET()
    .build();

HttpResponse<String> reponse = client.send(requete, HttpResponse.BodyHandlers.ofString());
System.out.println(reponse.statusCode());
System.out.println(reponse.body());

// POST JSON
String json = """
    {"nom": "UPS-04", "puissance_kva": 60}
    """;
HttpRequest post = HttpRequest.newBuilder()
    .uri(URI.create("https://api.exemple.com/equipements"))
    .header("Content-Type", "application/json")
    .POST(HttpRequest.BodyPublishers.ofString(json))
    .build();
HttpResponse<String> rep = client.send(post, HttpResponse.BodyHandlers.ofString());

// Asynchrone (CompletableFuture)
client.sendAsync(requete, HttpResponse.BodyHandlers.ofString())
      .thenApply(HttpResponse::body)
      .thenAccept(System.out::println);
```

**À retenir** : `HttpClient` est immutable et thread-safe → **une seule instance partagée** dans toute l'application.

---

## 44. Introduction à JDBC : parler aux bases SQL

```java
import java.sql.*;

// PostgreSQL : dépendance Maven -> org.postgresql:postgresql:42.7.x
String url = "jdbc:postgresql://localhost:5432/gmao";
String user = "gmao";            // en vrai : variables d'environnement, jamais en dur !
String pass = System.getenv("DB_PASSWORD");

try (Connection cx = DriverManager.getConnection(url, user, pass)) {

    // INSERT avec paramètres (ANTI-INJECTION SQL : toujours PreparedStatement, jamais de concaténation)
    String sql = "INSERT INTO equipement(nom, puissance_kva, etat) VALUES (?, ?, ?)";
    try (PreparedStatement ps = cx.prepareStatement(sql, Statement.RETURN_GENERATED_KEYS)) {
        ps.setString(1, "UPS-04");
        ps.setDouble(2, 60.0);
        ps.setString(3, "EN_SERVICE");
        ps.executeUpdate();
        try (ResultSet cles = ps.getGeneratedKeys()) {
            if (cles.next()) System.out.println("id=" + cles.getLong(1));
        }
    }

    // SELECT
    try (PreparedStatement ps = cx.prepareStatement("SELECT nom, puissance_kva FROM equipement WHERE etat = ?");
         ) {
        ps.setString(1, "EN_SERVICE");
        try (ResultSet rs = ps.executeQuery()) {
            while (rs.next()) {
                System.out.println(rs.getString("nom") + " : " + rs.getDouble("puissance_kva") + " kVA");
            }
        }
    }
} catch (SQLException e) {
    System.err.println("SQL [" + e.getSQLState() + "] : " + e.getMessage());
}
```

