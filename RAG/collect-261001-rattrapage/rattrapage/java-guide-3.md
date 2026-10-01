---
id: collect-261001-rattrapage/rattrapage/java-guide-3
title: "Guide Java — du zéro au production"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["arr"]
source: docs/RAG/collect-261001-rattrapage/java_guide.md
source_anchor: ""
source_lines: [493, 722]
sha256: b1858ab4b7fe19b9b20b28fb2bf9a4dcd0a754064ee36994d3e85ce71844627e
---

# Guide Java — du zéro au production

    // Passage par valeur : Java passe TOUJOURS par valeur
    public static void essayer(int x) { x = 99; }  // sans effet dehors
    public static void remplir(int[] t) { t[0] = 99; } // ...mais la valeur copiée peut être une référence !

    public static void main(String[] args) {
        int x = 5;
        essayer(x);
        System.out.println(x); // 5 : inchangé
        int[] t = {1, 2};
        remplir(t);
        System.out.println(t[0]); // 99 : le contenu du tableau a changé
    }
}
```

**À retenir** : Java est *pass-by-value* ; pour un objet, c'est la **référence** qui est copiée — d'où la confusion classique.

### Checklist « méthode propre »

- [ ] Nom = verbe d'action (`calculerTotal`, pas `truc`)
- [ ] Une méthode = une responsabilité (viser < 30 lignes)
- [ ] Paramètres ≤ 4 (au-delà : objet dédié ou builder)
- [ ] Retour explicite plutôt que `null` quand c'est possible (`Optional`, liste vide)
- [ ] Javadoc sur les méthodes publiques d'une bibliothèque

---

## 13. Classes et objets : les fondamentaux

La programmation orientée objet repose sur 4 piliers : **encapsulation, héritage, polymorphisme, abstraction**. Tout le reste du guide en découle.

```java
public class Onduleur {
    // Attributs (état)
    private String nom;
    private double puissanceKva;
    private boolean enService;

    // Constructeur (section 15)
    public Onduleur(String nom, double puissanceKva) {
        this.nom = nom;
        this.puissanceKva = puissanceKva;
        this.enService = true;
    }

    // Méthodes (comportement)
    public void mettreHorsService() {
        this.enService = false;
    }

    public String decrire() {
        return nom + " (" + puissanceKva + " kVA) - " +
               (enService ? "en service" : "hors service");
    }

    public static void main(String[] args) {
        Onduleur o1 = new Onduleur("UPS-Baie-A", 40.0);  // instanciation
        System.out.println(o1.decrire());
        o1.mettreHorsService();
        System.out.println(o1.decrire());
    }
}
```

Vocabulaire : **classe** = moule ; **objet** (instance) = exemplaire ; **attribut** = donnée ; **méthode** = action ; `this` = l'instance courante.

---

## 14. Constructeurs

```java
public class Batterie {
    private final String technologie;  // final : assigné une seule fois
    private final int capaciteAh;
    private int cycles;

    // Constructeur principal
    public Batterie(String technologie, int capaciteAh) {
        if (capaciteAh <= 0) throw new IllegalArgumentException("capacité > 0 requise");
        this.technologie = technologie;
        this.capaciteAh = capaciteAh;
    }

    // Constructeur secondaire qui délègue avec this(...)
    public Batterie(int capaciteAh) {
        this("VRLA", capaciteAh);  // doit être la PREMIÈRE instruction
    }

    // Pas de constructeur écrit => Java génère un constructeur par défaut sans argument.
    // Dès qu'on en écrit un, le défaut disparaît !
}
```

**Bonnes pratiques** : validez les arguments (fail-fast avec `IllegalArgumentException`) ; préférez des objets **immutables** (`final` partout) quand c'est possible — ils sont thread-safe par construction.

---

## 15. Encapsulation : private + getters/setters

```java
public class CompteurEnergie {
    private double kwh;   // inaccessible directement de l'extérieur

    public double getKwh() { return kwh; }

    public void ajouter(double delta) {
        if (delta < 0) throw new IllegalArgumentException("delta négatif interdit");
        this.kwh += delta;
    }
    // Pas de setKwh() : l'état ne change que via une règle métier.
}
```

L'encapsulation n'est pas de la paranoïa : c'est ce qui permet de **changer l'implémentation interne sans casser le code client** (ex. passer de `double` à `BigDecimal` sans toucher les appelants). Dans IntelliJ : `Alt+Insert` → Getters/Setters.

---

## 16. Héritage : extends

```java
public class Equipement {
    protected String nom;              // visible par les sous-classes
    public Equipement(String nom) { this.nom = nom; }
    public String decrire() { return "Équipement : " + nom; }
}

public class Onduleur extends Equipement {
    private double puissanceKva;

    public Onduleur(String nom, double puissanceKva) {
        super(nom);                    // appel OBLIGATOIRE au constructeur parent (1re ligne)
        this.puissanceKva = puissanceKva;
    }

    @Override                          // annotation : le compilateur vérifie la redéfinition
    public String decrire() {
        return super.decrire() + " - " + puissanceKva + " kVA";
    }
}
```

Règles :
- Une classe n'hérite que d'**une** seule classe (pas d'héritage multiple de classes).
- `final class` = non héritable ; `final` sur une méthode = non redéfinissable.
- **Préférez la composition à l'héritage** : héritez pour un vrai « est-un » (`Onduleur` est un `Equipement`), composez pour un « a-un » (`Onduleur` *a* une `Batterie`).

---

## 17. Polymorphisme

```java
public class Supervision {
    public static void afficher(Equipement e) {
        // Le bon decrire() est choisi À L'EXÉCUTION selon le type réel : liaison dynamique
        System.out.println(e.decrire());
    }

    public static void main(String[] args) {
        Equipement e1 = new Equipement("Switch");
        Equipement e2 = new Onduleur("UPS-01", 40.0);  // référence parent, objet enfant
        afficher(e1);  // Équipement : Switch
        afficher(e2);  // Équipement : UPS-01 - 40.0 kVA  (version Onduleur !)
    }
}
```

C'est le cœur des architectures extensibles : votre code de supervision manipule `Equipement`, et chaque nouvel équipement n'exige **aucune** modification de ce code (principe ouvert/fermé).

---

## 18. Classes abstraites et interfaces

```java
// Classe abstraite : base incomplète, ne peut pas être instanciée
public abstract class Capteur {
    protected String id;
    public Capteur(String id) { this.id = id; }
    public abstract double lire();          // les enfants DOIVENT l'implémenter
    public String getId() { return id; }    // méthode concrète partagée
}

// Interface : contrat pur (une classe peut en implémenter PLUSIEURS)
public interface Alertable {
    void alerter(String message);
    default void alerterCritique(String message) {   // méthode par défaut (Java 8+)
        alerter("[CRITIQUE] " + message);
    }
}

public class CapteurTemperature extends Capteur implements Alertable {
    public CapteurTemperature(String id) { super(id); }
    @Override public double lire() { return 21.5; /* lecture réelle ici */ }
    @Override public void alerter(String message) { System.out.println(id + " : " + message); }
}
```

| | Classe abstraite | Interface |
|---|---|---|
| Instanciable | Non | Non |
| Héritage multiple | Non (1 seule) | Oui (plusieurs) |
| Attributs d'instance | Oui | Non (constantes uniquement) |
| Constructeur | Oui | Non |
| Cas d'usage | Base commune avec état partagé | Contrat / capacité (`Alertable`, `Comparable`) |

**Depuis Java 8**, les interfaces peuvent avoir des méthodes `default` et `static` ; utilisez-les pour faire évoluer un contrat sans casser les implémentations existantes.

---

## 19. Énumérations (enum)

Bien plus qu'une liste de constantes : une `enum` Java est une vraie classe.

```java
public enum EtatEquipement {
    EN_SERVICE("Opérationnel", 1),
    DEGRADE("Fonctionnement dégradé", 2),
    HORS_SERVICE("Arrêté", 3);

    private final String libelle;
    private final int criticite;

    EtatEquipement(String libelle, int criticite) {
        this.libelle = libelle;
        this.criticite = criticite;
    }

    public String getLibelle() { return libelle; }
    public boolean estCritique() { return criticite >= 3; }
}

