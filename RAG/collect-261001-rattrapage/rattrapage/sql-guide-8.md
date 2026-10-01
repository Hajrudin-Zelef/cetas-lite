---
id: collect-261001-rattrapage/rattrapage/sql-guide-8
title: "Guide SQL complet — du SELECT à l'optimisation"
domain: rattrapage
role: reference
task: reference
actors: []
dates: ["2024-06-01"]
keywords: []
source: docs/RAG/collect-261001-rattrapage/sql_guide.md
source_anchor: ""
source_lines: [1496, 1697]
sha256: 3530a38c51a2b448f5023608eb87e82403a0fd4e099871b061732d0e52522453
---

# Guide SQL complet — du SELECT à l'optimisation

| Besoin | PostgreSQL | MySQL/MariaDB | Notes |
|---|---|---|---|
| Entier 32 bits | `INTEGER` | `INT` | PK courantes |
| Entier 64 bits | `BIGINT` | `BIGINT` | gros volumes |
| Auto-incrément | `SERIAL`/`IDENTITY` | `INT AUTO_INCREMENT` | §42 |
| Décimal exact (monnaie) | `NUMERIC(p,s)` | `DECIMAL(p,s)` | **toujours** pour l'argent |
| Flottant approximatif | `DOUBLE PRECISION` | `DOUBLE` | mesures, jamais monnaie |
| Chaîne courte | `VARCHAR(n)` | `VARCHAR(n)` | n = garde-fou, pas perf |
| Texte long | `TEXT` | `TEXT`/`MEDIUMTEXT` | descriptions, logs |
| Booléen | `BOOLEAN` | `BOOLEAN` (= TINYINT(1)) | ⚠️ MySQL : 0/1 |
| Date / heure | `DATE`, `TIMESTAMP`, `TIMESTAMPTZ` | `DATE`, `DATETIME`, `TIMESTAMP` | §16 |
| Binaire | `BYTEA` | `BLOB` | fichiers, empreintes |
| JSON | `JSONB` | `JSON` | semi-structuré (§79) |
| Réseau (PG) | `INET`, `CIDR`, `MACADDR` | — (VARCHAR) | adresses IP : typé > texte |

Règles :

- **Monnaie → NUMERIC/DECIMAL**, jamais FLOAT (0.1 + 0.2 ≠ 0.3 en binaire).
- IP, emails : VARCHAR + CHECK de format, ou types natifs (PG `INET` permet
  les opérateurs réseau `>>`, `<<`).
- `VARCHAR(255)` par habitude : OK, mais mettre une vraie limite métier quand elle existe
  (`statut VARCHAR(20)` + CHECK).
- ⚠️ MySQL `VARCHAR` sans longueur de charset utf8mb4 : 4 octets/caractère pour l'indexation
  (limite 191 caractères pour index unique en InnoDB ancien).

---

## 44. ALTER TABLE — faire évoluer sans casser

```sql
-- Ajouter une colonne (avec défaut pour les lignes existantes)
ALTER TABLE equipements ADD COLUMN IF NOT EXISTS
    puissance_kw NUMERIC(8,2);

-- Ajouter une contrainte
ALTER TABLE equipements ADD CONSTRAINT ck_puissance_positive
    CHECK (puissance_kw IS NULL OR puissance_kw > 0);

-- Renommer (sans toucher aux données)
ALTER TABLE equipements RENAME COLUMN localisation TO site;

-- Changer le type (⚠️ peut réécrire toute la table)
ALTER TABLE equipements ALTER COLUMN site TYPE VARCHAR(80);
-- MySQL : ALTER TABLE equipements MODIFY site VARCHAR(80);

-- Poser / retirer NOT NULL
ALTER TABLE equipements ALTER COLUMN marque SET NOT NULL;    -- PG
ALTER TABLE equipements ALTER COLUMN marque DROP NOT NULL;   -- PG

-- Supprimer une colonne
ALTER TABLE equipements DROP COLUMN IF EXISTS puissance_kw;
```

En production, sur de grosses tables :

1. `ADD COLUMN` avec DEFAULT **non volatile** : instantané sous PG 11+ (métadonnée),
   mais réécrit la table sous MySQL < 8 (ALGORITHM=INPLACE l'évite souvent).
2. Tester sur une copie (`CREATE TABLE test AS SELECT * FROM prod LIMIT 0` + jeu d'essai).
3. Faire les ALTER en heure creuse, avec un backup avant (section 68).
4. ⚠️ `ALTER ... TYPE` avec conversion : `USING` (PG) pour exprimer la conversion.

---

## 45. DROP TABLE et TRUNCATE

```sql
DROP TABLE IF EXISTS maintenances_planifiees;            -- supprime structure + données
DROP TABLE IF EXISTS a, b, c CASCADE;                    -- PG : CASCADE supprime les dépendances (vues, FK)

TRUNCATE TABLE interventions;                            -- vide la table, rapide
TRUNCATE TABLE interventions RESTART IDENTITY;           -- + remet les compteurs à zéro (PG)
-- MySQL : TRUNCATE TABLE interventions;  (remet AUTO_INCREMENT à 1, pas de RESTART IDENTITY)
```

| | `DELETE` sans WHERE | `TRUNCATE` | `DROP` |
|---|---|---|---|
| Vitesse | lent (ligne à ligne) | très rapide | instantané |
| WHERE possible | oui | non | — |
| Déclenche triggers DELETE | oui | non (PG) | — |
| Rollback possible | oui (transaction) | oui sous PG, **non sous MySQL** (DDL implicite) | non |
| Remet les séquences à zéro | non | option | — |

⚠️ Sous MySQL, `TRUNCATE` = DDL → **commit implicite**, impossible à annuler.
En cas de doute sur une table métier : `DELETE` en transaction, vérification, puis `COMMIT`.

---

## 46. Clés primaires (PRIMARY KEY)

```sql
-- En ligne
CREATE TABLE sites (
    id   SERIAL PRIMARY KEY,
    code VARCHAR(10) NOT NULL UNIQUE,
    nom  VARCHAR(80) NOT NULL
);

-- Clé composite
CREATE TABLE interventions_techniciens (
    intervention_id INTEGER NOT NULL REFERENCES interventions(id),
    technicien_id   INTEGER NOT NULL REFERENCES employes(id),
    role            VARCHAR(20) NOT NULL DEFAULT 'principal',
    PRIMARY KEY (intervention_id, technicien_id)
);
```

Propriétés : UNIQUE + NOT NULL, **une seule** par table, crée un index automatiquement.
Choix PK :

- **Surrogate** (id technique auto) : stable, étroite, jamais métier → choix par défaut.
- **Naturelle** (ex. `reference` unique) : parlante mais risque de changement
  (renumérotation d'inventaire → cascade de mises à jour).
- Éviter les PK composites larges : chaque FK et chaque index secondaire les embarque.

---

## 47. Clés étrangères (FOREIGN KEY) — l'intégrité référentielle

```sql
CREATE TABLE interventions (
    ...
    equipement_id INTEGER NOT NULL
        REFERENCES equipements(id)
        ON DELETE RESTRICT      -- refuse la suppression d'un équipement suivi
        ON UPDATE CASCADE,      -- propage un changement d'id (rare avec surrogate)
    ...
);
```

Comportements `ON DELETE` :

| Option | Effet |
|---|---|
| `RESTRICT` / `NO ACTION` | refuse si des enfants existent (défaut, le plus sûr) |
| `CASCADE` | supprime les enfants en cascade ⚠️ |
| `SET NULL` | met la FK à NULL (colonne doit accepter NULL) |
| `SET DEFAULT` | met la valeur par défaut |

```sql
-- Ajouter après coup + nommer la contrainte
ALTER TABLE interventions
    ADD CONSTRAINT fk_interv_equip
    FOREIGN KEY (equipement_id) REFERENCES equipements(id)
    ON DELETE RESTRICT;

-- Trouver les orphelins AVANT de poser la FK (sinon l'ALTER échoue)
SELECT i.* FROM interventions i
LEFT JOIN equipements e ON e.id = i.equipement_id
WHERE e.id IS NULL;
```

⚠️ MySQL : les FK n'existent qu'en **InnoDB** (pas MyISAM !). Vérifier le moteur :
`SHOW TABLE STATUS`. Une FK silencieusement ignorée = corruption logique future.

---

## 48. UNIQUE, CHECK, NOT NULL, DEFAULT

```sql
CREATE TABLE contrats_maintenance (
    id          SERIAL PRIMARY KEY,
    equipement_id INTEGER NOT NULL UNIQUE REFERENCES equipements(id),
                       -- UNIQUE + FK : relation 1-1 (un contrat par équipement)
    prestataire VARCHAR(80) NOT NULL,
    date_debut  DATE NOT NULL,
    date_fin    DATE NOT NULL,
    montant_annuel NUMERIC(10,2) NOT NULL DEFAULT 0 CHECK (montant_annuel >= 0),
    CONSTRAINT ck_dates CHECK (date_fin > date_debut),
    CONSTRAINT uq_presta_periode UNIQUE (prestataire, date_debut)
);
```

| Contrainte | Garantit |
|---|---|
| `NOT NULL` | pas de valeur absente |
| `UNIQUE` | pas de doublon (plusieurs NULL autorisés en standard !) |
| `CHECK` | règle métier au plus près des données |
| `DEFAULT` | valeur si non fournie |

⚠️ `UNIQUE` et NULL : en standard, deux NULL ne sont **pas** considérés égaux →
plusieurs NULL autorisés dans une colonne UNIQUE. (SQL Server fait exception.)
PG 15+ : `UNIQUE NULLS NOT DISTINCT` pour interdire les doublons de NULL.

Philosophie : **les contraintes sont la dernière ligne de défense**.
L'application peut avoir un bug, le script d'import aussi ; la base, elle, refuse
la donnée incohérente. Coût : quasi nul en écriture unitaire, énorme en fiabilité.

---

## 49. INSERT — insérer des données

```sql
-- Toujours lister les colonnes (résiste aux ALTER TABLE)
INSERT INTO equipements (id, reference, type_equip, marque, modele, date_achat, localisation)
VALUES (106, 'SW-001', 'switch', 'Cisco', 'Catalyst 9300', '2024-06-01', 'Salle serveurs');

-- Plusieurs lignes en une fois (1 seul aller-retour réseau)
INSERT INTO equipements (id, reference, type_equip, marque, modele)
VALUES
 (107, 'UPS-004', 'onduleur', 'Vertiv', 'Liebert EXS 30 kVA'),
 (108, 'CPY-002', 'copieur',  'Kyocera','TASKalfa 3554ci');

