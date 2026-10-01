---
id: collect-261001-rattrapage/rattrapage/java-guide-8
title: "Guide Java — du zéro au production"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["packaging"]
source: docs/RAG/collect-261001-rattrapage/java_guide.md
source_anchor: ""
source_lines: [1554, 1789]
sha256: f8f694404a284c118fe2116ed426f705728606b615f66d739eaf1ad3d10887a6
---

# Guide Java — du zéro au production

**Indispensable en production** : un pool de connexions (HikariCP — inclus par défaut dans Spring Boot). Ouvrir une connexion TCP par requête = lent et fragile.

**Transaction** :
```java
cx.setAutoCommit(false);
try {
    // ... plusieurs updates ...
    cx.commit();
} catch (SQLException e) {
    cx.rollback();   // tout ou rien
    throw e;
} finally {
    cx.setAutoCommit(true);
}
```

---

## 45. Tests unitaires avec JUnit 5

```java
import org.junit.jupiter.api.*;
import static org.junit.jupiter.api.Assertions.*;

class CalculatriceTest {

    @Test
    @DisplayName("L'addition de deux positifs fonctionne")
    void addition() {
        assertEquals(5, Calculatrice.additionner(2, 3));
    }

    @Test
    void divisionParZeroLeveUneException() {
        assertThrows(IllegalArgumentException.class, () -> Calculatrice.diviser(1, 0));
    }

    @Test
    void sommeVarargs() {
        assertAll(
            () -> assertEquals(6, Calculatrice.somme(1, 2, 3)),
            () -> assertEquals(0, Calculatrice.somme())
        );
    }

    // Test paramétré : le même test avec plusieurs jeux de données
    @ParameterizedTest
    @CsvSource({"2, 3, 5", "0, 0, 0", "-1, 1, 0"})
    void additionParametree(int a, int b, int attendu) {
        assertEquals(attendu, Calculatrice.additionner(a, b));
    }

    @Test
    @Timeout(2)   // échoue si > 2 secondes
    void neDoitPasBoucler() { /* ... */ }
}
```

Cycle de vie : `@BeforeEach` (setup avant chaque test), `@AfterEach`, `@BeforeAll`/`@AfterAll` (static, une fois). **Règle** : un test = un comportement ; nom explicite ; pas de dépendance entre tests ; pas d'accès réseau/disque dans les tests unitaires (mockez !).

---

## 46. Mockito : simuler les dépendances

```java
import static org.mockito.Mockito.*;
import static org.junit.jupiter.api.Assertions.*;

// On teste SupervisionService sans vraie base ni vrai réseau
class SupervisionServiceTest {

    @Test
    void alerteSiBatterieFaible() {
        // 1. Créer le mock
        CapteurBatterie capteur = mock(CapteurBatterie.class);
        // 2. Programmer son comportement
        when(capteur.niveauCharge("UPS-01")).thenReturn(12);  // 12 %
        // 3. Injecter et tester
        SupervisionService service = new SupervisionService(capteur);
        assertTrue(service.batterieCritique("UPS-01"));
        // 4. Vérifier les interactions
        verify(capteur).niveauCharge("UPS-01");
        verify(capteur, never()).niveauCharge("UPS-02");
    }

    @Test
    void relanceEnCasDErreur() {
        CapteurBatterie capteur = mock(CapteurBatterie.class);
        when(capteur.niveauCharge(anyString()))
            .thenThrow(new RuntimeException("timeout"))  // 1er appel
            .thenReturn(80);                             // 2e appel
        SupervisionService service = new SupervisionService(capteur);
        assertFalse(service.batterieCritique("UPS-01")); // le retry a fonctionné
        verify(capteur, times(2)).niveauCharge("UPS-01");
    }
}
```

**Mockito** = `mock()` + `when(...).thenReturn(...)` + `verify()`. Ne mockez que les **frontières** (réseau, BDD, horloge) — jamais la classe testée elle-même.

---

## 47. Debug : méthode et outils

**Méthode en 5 étapes** :
1. **Reproduire** : cas minimal, déterministe si possible.
2. **Lire** : la stack trace du **haut** (votre code) vers le bas ; la 1re ligne « Caused by » est souvent la vraie cause.
3. **Hypothèse** : une seule à la fois.
4. **Vérifier** : breakpoint ou log ciblé, pas 50 `println`.
5. **Corriger + test de non-régression** : ajoutez le test qui aurait attrapé le bug.

**Dans IntelliJ** : clic dans la gouttière = breakpoint ; `F8` pas-à-pas, `F7` entrer dans, `Shift+F8` sortir ; « Evaluate Expression » (`Alt+F8`) pour inspecter ; breakpoints conditionnels (clic droit) pour les boucles.

**Flags JVM utiles au diagnostic** :
```bash
java -Xlog:gc*:file=gc.log          # logs GC
jstack <pid>                         # dump des threads (deadlocks !)
jmap -histo <pid>                    # histogramme du heap (fuites mémoire)
java -XX:+HeapDumpOnOutOfMemoryError -XX:HeapDumpPath=/tmp/dump.hprof -jar app.jar
```

**Lire une stack trace** :
```text
Exception in thread "main" java.lang.NullPointerException: Cannot invoke "String.length()" because "nom" is null
        at com.entreprise.App.valider(App.java:42)      <- VOTRE code, ligne 42
        at com.entreprise.App.main(App.java:15)
```
Depuis Java 14, les NPE sont « helpful » : le message dit **quelle** variable est null. (Voir aussi section 53.)

---

## 48. Intro Spring Boot : exposer une API REST

Spring Boot = le standard pour les applications/services Java en entreprise. Il embarque un serveur (Tomcat) : **un jar exécutable suffit**.

`pom.xml` (parent + starter web) :
```xml
<parent>
  <groupId>org.springframework.boot</groupId>
  <artifactId>spring-boot-starter-parent</artifactId>
  <version>3.3.5</version>
</parent>
<dependencies>
  <dependency>
    <groupId>org.springframework.boot</groupId>
    <artifactId>spring-boot-starter-web</artifactId>
  </dependency>
</dependencies>
```

```java
package com.entreprise.api;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.web.bind.annotation.*;
import java.util.*;
import java.util.concurrent.ConcurrentHashMap;

@SpringBootApplication
public class ApiApplication {
    public static void main(String[] args) {
        SpringApplication.run(ApiApplication.class, args);  // démarre sur http://localhost:8080
    }
}

// DTO immutable : un record suffit
record EquipementDto(String nom, double puissanceKva, String etat) {}

@RestController
@RequestMapping("/api/equipements")
class EquipementController {

    private final Map<String, EquipementDto> base = new ConcurrentHashMap<>();

    @GetMapping
    public Collection<EquipementDto> lister() {
        return base.values();
    }

    @GetMapping("/{nom}")
    public EquipementDto detail(@PathVariable String nom) {
        EquipementDto e = base.get(nom);
        if (e == null) throw new EquipementIntrouvableException(nom);
        return e;
    }

    @PostMapping
    @ResponseStatus(HttpStatus.CREATED)
    public EquipementDto creer(@RequestBody EquipementDto dto) {
        base.put(dto.nom(), dto);
        return dto;
    }

    @DeleteMapping("/{nom}")
    @ResponseStatus(HttpStatus.NO_CONTENT)
    public void supprimer(@PathVariable String nom) {
        base.remove(nom);
    }
}

// Gestion d'erreur centralisée
@ResponseStatus(HttpStatus.NOT_FOUND)
class EquipementIntrouvableException extends RuntimeException {
    EquipementIntrouvableException(String nom) { super("Équipement introuvable : " + nom); }
}
```

Test :
```bash
curl -X POST localhost:8080/api/equipements -H 'Content-Type: application/json' \
  -d '{"nom":"UPS-04","puissance_kva":60,"etat":"EN_SERVICE"}'
curl localhost:8080/api/equipements/UPS-04
```

**À retenir** : `@RestController` + `@GetMapping`/`@PostMapping` + records DTO = API REST fonctionnelle en 50 lignes. La suite logique : Spring Data JPA (BDD), Spring Security, Actuator (supervision `/actuator/health`).

---

## 49. Packaging : jar, jlink, jpackage

```bash
# 1. Jar exécutable avec Maven (plugin shade : inclut les dépendances)
# pom.xml -> maven-shade-plugin, puis :
mvn package
java -jar target/outil-1.0.0.jar

# MANIFESTE : pour un jar lançable sans plugin, indiquez la classe principale
# src/main/resources/META-INF/MANIFEST.MF :
#   Main-Class: com.entreprise.outil.Main

# 2. Runtime sur mesure (Java 11+) : embarque uniquement les modules JDK nécessaires
jlink --module-path "$JAVA_HOME/jmods" --add-modules java.base,java.sql,java.net.http \
      --output runtime-perso --compress=2 --strip-debug
# -> dossier runtime-perso/bin/java (~40 Mo au lieu de ~300 Mo)

