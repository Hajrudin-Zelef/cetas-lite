---
id: collect-261001-rattrapage/rattrapage/java-guide-1
title: "Guide Java — du zéro au production"
domain: rattrapage
role: reference
task: reference
actors: ["AWS", "Oracle"]
dates: []
keywords: ["aws", "distribution"]
source: docs/RAG/collect-261001-rattrapage/java_guide.md
source_anchor: ""
source_lines: [1, 230]
sha256: ba3145a7303388970323b41ffe4c4a0f259917648d6e81531aafbb5c8a69ef64
---

# Guide Java — du zéro au production

> **Public visé** : administrateurs systèmes, responsables d'équipe technique et développeurs d'outils internes d'entreprise.
> **Niveau** : débutant à intermédiaire confirmé. **Version cible** : Java 17+ (LTS). Les nouveautés Java 21 sont signalées explicitement.
> **Ton** : direct, tutoriel + référence. Chaque section se termine par un point « À retenir » quand c'est utile.

---

## 1. Pourquoi Java en 2026 ?

Java reste le pilier des applications d'entreprise : banques, opérateurs télécoms, éditeurs de supervision, outils internes critiques. Pourquoi il vous concerne, vous, chef de service systèmes & énergies :

- **Portabilité réelle** : « write once, run anywhere » — un jar tourne sur Linux, Windows, conteneurs Docker sans recompilation.
- **Écosystème mature** : Maven Central contient des millions de bibliothèques (JDBC, SNMP, Modbus, LDAP, etc.).
- **Longévité** : les LTS (8, 11, 17, 21, 25) garantissent des années de support. Un outil interne écrit aujourd'hui tournera dans 10 ans.
- **Performance** : la JVM (JIT, GC modernes) rivalise avec le natif pour la plupart des charges métier.
- **Spring Boot** : l'écosystème dominant pour exposer vos outils en API REST / dashboards web.

**À retenir** : pour des outils internes (inventaires, collecteurs SNMP, batchs d'import, petites API), Java + Spring Boot est un choix sûr et durable.

---

## 2. Installation du JDK

### 2.1 Choisir une distribution

| Distribution | Fournisseur | Licence | Remarque |
|---|---|---|---|
| Temurin | Eclipse Adoptium | Gratuite (GPL+CE) | Choix recommandé, builds testés |
| Oracle JDK | Oracle | Gratuite sous NFTC | OK pour dev, vérifier les conditions en prod |
| Corretto | Amazon | Gratuite | Bien pour AWS |
| Zulu | Azul | Gratuite (communauté) | Large choix de versions |
| GraalVM | Oracle | Gratuite | Compilation native, polyglotte |

**Règle** : en entreprise, privilégiez une LTS : **Java 17** (supportée jusqu'en 2029+) ou **Java 21** (jusqu'en 2031+).

### 2.2 Installation sur Debian/Ubuntu

```bash
# Temurin via le dépôt Adoptium
sudo apt install -y wget apt-transport-https gpg
wget -qO - https://packages.adoptium.net/artifactory/api/gpg/key/public | gpg --dearmor | sudo tee /etc/apt/trusted.gpg.d/adoptium.gpg > /dev/null
echo "deb https://packages.adoptium.net/artifactory/deb $(awk -F= '/^VERSION_CODENAME/{print$2}' /etc/os-release) main" | sudo tee /etc/apt/sources.list.d/adoptium.list
sudo apt update && sudo apt install -y temurin-21-jdk

# Vérification
java -version
javac -version
```

### 2.3 Installation sur Windows

1. Téléchargez l'installeur `.msi` Temurin 21 depuis adoptium.net.
2. Cochez « Set JAVA_HOME » et « Add to PATH » pendant l'installation.
3. Vérifiez dans PowerShell : `java -version`.

### 2.4 Variables d'environnement

```bash
# ~/.bashrc ou ~/.profile
export JAVA_HOME=/usr/lib/jvm/temurin-21-jdk-amd64
export PATH="$JAVA_HOME/bin:$PATH"
```

Vérifiez que `javac` pointe bien vers la bonne version : `which javac` puis `javac -version`.

### 2.5 JDK vs JRE

- **JDK** (Java Development Kit) : compilateur `javac` + outils + runtime. **C'est ce qu'il faut installer pour développer.**
- **JRE** : runtime seul (n'existe plus en distribution séparée depuis Java 11 ; on génère un runtime sur mesure avec `jlink`/`jpackage`, voir section 48).

**À retenir** : installez toujours un JDK LTS (17 ou 21), jamais « la dernière version » non-LTS pour de la production.

---

## 3. Tooling : Maven, Gradle, IDE

### 3.1 Maven (recommandé pour débuter)

Maven gère les dépendances et le cycle de build via un fichier `pom.xml`.

```bash
# Installation Debian/Ubuntu
sudo apt install -y maven
mvn -version
```

Créer un projet :

```bash
mvn archetype:generate -DgroupId=com.entreprise -DartifactId=outil-inventaire \
  -DarchetypeArtifactId=maven-archetype-quickstart -DinteractiveMode=false
cd outil-inventaire
mvn compile     # compile
mvn package     # produit target/outil-inventaire-1.0-SNAPSHOT.jar
mvn test        # lance les tests
```

`pom.xml` minimal avec Java 21 :

```xml
<project>
  <modelVersion>4.0.0</modelVersion>
  <groupId>com.entreprise</groupId>
  <artifactId>outil-inventaire</artifactId>
  <version>1.0.0</version>
  <properties>
    <maven.compiler.release>21</maven.compiler.release>
    <project.build.sourceEncoding>UTF-8</project.build.sourceEncoding>
  </properties>
  <dependencies>
    <dependency>
      <groupId>org.junit.jupiter</groupId>
      <artifactId>junit-jupiter</artifactId>
      <version>5.10.3</version>
      <scope>test</scope>
    </dependency>
  </dependencies>
</project>
```

### 3.2 Gradle (alternative moderne)

```bash
# Via SDKMAN (recommandé)
curl -s "https://get.sdkman.io" | bash
sdk install gradle
gradle init   # assistant interactif
./gradlew build
```

`build.gradle` minimal (syntaxe Kotlin DSL) :

```kotlin
plugins { java }
java { toolchain { languageVersion.set(JavaLanguageVersion.of(21)) } }
repositories { mavenCentral() }
dependencies { testImplementation("org.junit.jupiter:junit-jupiter:5.10.3") }
tasks.test { useJUnitPlatform() }
```

### 3.3 IDE : IntelliJ IDEA (recommandé)

- **IntelliJ IDEA Community** (gratuit) : le meilleur support Java/Maven/Gradle. Refactoring, debug, inspections.
- **Eclipse** : gratuit, historique, encore très présent en entreprise.
- **VS Code** + extension « Extension Pack for Java » : léger, correct pour petits projets.

**Conseil** : prenez IntelliJ IDEA Community. Le gain de productivité (complétion, corrections automatiques) vaut largement l'installation.

### 3.4 Outils en ligne de commande du JDK à connaître

| Outil | Rôle |
|---|---|
| `javac` | Compilateur |
| `java` | Lanceur d'applications |
| `jar` | Manipulation d'archives jar |
| `javadoc` | Génération de documentation |
| `jshell` | REPL interactif (idéal pour tester un bout de code) |
| `jconsole` / `jvisualvm` | Supervision JVM (threads, heap) |
| `jstat`, `jmap`, `jstack` | Diagnostic production |
| `jlink`, `jpackage` | Runtimes sur mesure et installateurs |

Essayez `jshell` dès maintenant : c'est le bac à sable parfait pour ce guide.

```text
$ jshell
|  Welcome to JShell -- Version 21
jshell> int x = 40 + 2;
x ==> 42
jshell> System.out.println("Hello " + x);
Hello 42
```

**À retenir** : JDK Temurin 21 + Maven + IntelliJ IDEA Community = stack de départ saine et gratuite.

---

## 4. Premier programme : compilation et exécution

Fichier `Hello.java` (le nom du fichier **doit** correspondre à la classe publique) :

```java
public class Hello {
    public static void main(String[] args) {
        System.out.println("Bonjour, Java !");
        if (args.length > 0) {
            System.out.println("Argument reçu : " + args[0]);
        }
    }
}
```

```bash
javac Hello.java        # produit Hello.class (bytecode)
java Hello              # exécute
java Hello monde        # avec un argument
```

Depuis Java 11, on peut lancer un fichier source directement (pratique pour les scripts) :

```bash
java Hello.java monde
```

**Anatomie** :
- `public class Hello` : tout le code Java vit dans des classes.
- `public static void main(String[] args)` : point d'entrée du programme.
- `System.out.println(...)` : affiche sur la sortie standard.

**À retenir** : un fichier = une classe publique du même nom ; `main` est la porte d'entrée.

---

## 5. Syntaxe de base : le squelette du langage

```java
// Commentaire sur une ligne
/* Commentaire
   sur plusieurs lignes */
/** Commentaire Javadoc (génère de la doc) */

package com.entreprise.outil;   // doit être la 1re instruction (hors commentaires)

import java.util.ArrayList;     // import d'une classe
import java.util.*;             // import d'un package (évitez * en équipe)

public class Syntaxe {
    // Attribut (champ) de classe
    private static final double TVA = 0.20;

