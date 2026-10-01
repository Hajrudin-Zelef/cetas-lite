---
id: collect-261001-rattrapage/rattrapage/sql-guide-1
title: "Guide SQL complet — du SELECT à l'optimisation"
domain: rattrapage
role: reference
task: reference
actors: []
dates: ["2015-03-01", "2018-06-15", "2019-01-20", "2019-11-03", "2020-04-22", "2020-09-01", "2021-05-10", "2022-02-14", "2023-08-01", "2024-02-01", "2025-12-01", "2026-08-15", "2026-09-01"]
keywords: ["agents"]
source: docs/RAG/collect-261001-rattrapage/sql_guide.md
source_anchor: ""
source_lines: [1, 154]
sha256: 475a8b662b560c5257955a9710b0e787da958aecafba1a311406b17bd20d138b
---

# Guide SQL complet — du SELECT à l'optimisation

**Public :** Zelef, chef de service Systèmes & Énergies, sysadmin.
**Objectif :** maîtriser SQL au quotidien (requêtes d'exploitation, reporting, dépannage),
sans installation lourde : on travaille sur les bases **déjà en place** (PostgreSQL, MySQL/MariaDB).
**Portabilité :** tout le SQL présenté est portable ; les différences MySQL/MariaDB ↔ PostgreSQL
sont signalées dans des encadrés `⚠️ Portabilité`.

**Tables de référence utilisées dans tout le guide** (entreprise fictive "SETIS", service maintenance) :

- `employes` — les techniciens et agents du service
- `equipements` — parc : onduleurs, copieurs, serveurs, commutateurs
- `interventions` — bons d'intervention (pannes, préventif)
- `onduleurs` — détail technique des UPS suivis

> Convention du guide : les mots-clés SQL en MAJUSCULES, les identifiants en minuscules_snake_case,
> les chaînes entre apostrophes simples `'texte'`, les dates au format ISO `YYYY-MM-DD`.

---

## 1. Pourquoi ce guide et comment l'utiliser

SQL est le langage commun de toutes les bases relationnelles : inventorié en 1974, normalisé
(ISO/IEC 9075), il n'a quasiment pas changé dans ses fondamentaux. Pour un sysadmin/chef de service,
SQL sert tous les jours :

- **Exploitation** : interroger GLPI, Zabbix, l'inventaire, les logs applicatifs stockés en base.
- **Reporting** : nombre d'interventions par mois, MTBF des onduleurs, SLA du service.
- **Dépannage** : trouver la requête qui rame, vérifier une incohérence de données.
- **Sauvegarde/audit** : contrôler ce qui est sauvegardé, extraire des jeux de données.

Méthode de lecture conseillée :

1. Lire les sections 1 à 13 dans l'ordre (socle).
2. Pratiquer chaque exemple sur une base de test (voir section 3).
3. Revenir aux sections 60+ (optimisation, pièges) dès qu'une requête réelle coince.
4. Garder la section 76 (pense-bête) imprimée ou en favori.

⚠️ **Portabilité** — 95 % du guide est du SQL standard. Quand une fonction ou syntaxe diffère,
le guide donne les deux variantes avec une mention explicite du SGBD.

Checklist de démarrage :

- [ ] Identifier les bases accessibles (PostgreSQL ? MariaDB ? SQLite locale ?)
- [ ] Obtenir un accès **lecture seule** pour s'entraîner sans risque
- [ ] Créer la base de test `sql_formation` et y charger les tables de la section 3
- [ ] Installer un client : `psql` (PostgreSQL), `mysql` (MariaDB), ou DBeaver (graphique, multi-SGBD)

---

## 2. Les 5 familles d'ordres SQL

| Famille | Nom | Rôle | Exemples |
|---|---|---|---|
| DQL | Data Query Language | Interroger | `SELECT` |
| DML | Data Manipulation Language | Modifier les données | `INSERT`, `UPDATE`, `DELETE` |
| DDL | Data Definition Language | Définir les structures | `CREATE TABLE`, `ALTER TABLE`, `DROP TABLE` |
| DCL | Data Control Language | Gérer les droits | `GRANT`, `REVOKE` |
| TCL | Transaction Control Language | Gérer les transactions | `COMMIT`, `ROLLBACK`, `SAVEPOINT` |

Ce guide couvre surtout DQL (sections 4–41), puis DDL (42–48), DML (49–53),
TCL (54–55), et les sujets transverses (index, sécurité, sauvegarde, optimisation).

Point crucial pour un sysadmin : **savoir lire** du SQL est aussi important que savoir l'écrire.
Les sections 59 (EXPLAIN) et 69 (dépannage) sont là pour ça.

---

## 3. Jeu de données de référence (à créer une fois)

Tout le guide utilise les mêmes tables. Les créer **une fois** dans une base de test,
et tous les exemples deviennent exécutables.

```sql
-- Base de test (PostgreSQL)
CREATE DATABASE sql_formation;

-- MySQL/MariaDB : CREATE DATABASE sql_formation CHARACTER SET utf8mb4;
-- puis : USE sql_formation;
```

```sql
CREATE TABLE employes (
    id          INTEGER PRIMARY KEY,
    nom         VARCHAR(50)  NOT NULL,
    prenom      VARCHAR(50)  NOT NULL,
    poste       VARCHAR(60)  NOT NULL,
    service     VARCHAR(40)  NOT NULL DEFAULT 'Maintenance',
    salaire     NUMERIC(10,2),
    date_embauche DATE NOT NULL,
    manager_id  INTEGER REFERENCES employes(id),
    actif       BOOLEAN NOT NULL DEFAULT TRUE
);

CREATE TABLE equipements (
    id            INTEGER PRIMARY KEY,
    reference     VARCHAR(30)  NOT NULL UNIQUE,
    type_equip    VARCHAR(30)  NOT NULL,   -- 'onduleur','copieur','serveur','switch'
    marque        VARCHAR(30)  NOT NULL,
    modele        VARCHAR(60)  NOT NULL,
    date_achat    DATE,
    garantie_jusqu DATE,
    localisation  VARCHAR(60),
    statut        VARCHAR(20)  NOT NULL DEFAULT 'en_service'
                  CHECK (statut IN ('en_service','en_panne','en_maintenance','reforme'))
);

CREATE TABLE interventions (
    id              INTEGER PRIMARY KEY,
    equipement_id   INTEGER NOT NULL REFERENCES equipements(id),
    technicien_id   INTEGER REFERENCES employes(id),
    date_ouverture  TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    date_cloture    TIMESTAMP,
    type_interv     VARCHAR(20) NOT NULL DEFAULT 'curatif'
                    CHECK (type_interv IN ('curatif','preventif','installation')),
    priorite        INTEGER NOT NULL DEFAULT 3 CHECK (priorite BETWEEN 1 AND 5),
    description     TEXT NOT NULL,
    cout_eur        NUMERIC(10,2) DEFAULT 0,
    sla_respecte    BOOLEAN
);

CREATE TABLE onduleurs (
    equipement_id   INTEGER PRIMARY KEY REFERENCES equipements(id),
    puissance_kva   NUMERIC(6,2)  NOT NULL,
    topologie       VARCHAR(10)   NOT NULL DEFAULT 'VFI'
                    CHECK (topologie IN ('VFI','VI','VFD')),
    nb_batteries    INTEGER NOT NULL DEFAULT 0,
    date_dernier_test DATE,
    autonomie_min    INTEGER
);
```

Jeu d'exemple (extrait — à insérer pour suivre les exemples) :

```sql
INSERT INTO employes (id, nom, prenom, poste, salaire, date_embauche, manager_id) VALUES
 (1, 'Zelef',   'Karim',  'Chef de service',      5200.00, '2015-03-01', NULL),
 (2, 'Diallo',  'Awa',    'Technicien energie',   3100.00, '2018-06-15', 1),
 (3, 'Ndiaye',  'Moussa', 'Technicien reseau',    2950.00, '2019-01-20', 1),
 (4, 'Sow',     'Fatou',  'Technicienne copieurs',2800.00, '2020-09-01', 1),
 (5, 'Kane',    'Ibrahima','Apprenti',            1400.00, '2024-02-01', 2);

INSERT INTO equipements (id, reference, type_equip, marque, modele, date_achat, localisation, statut) VALUES
 (101, 'UPS-001', 'onduleur', 'APC',     'Easy UPS 3S 40 kVA', '2021-05-10', 'Local TGBT',   'en_service'),
 (102, 'UPS-002', 'onduleur', 'Eaton',   '9PX 11 kVA',         '2019-11-03', 'Salle serveurs','en_panne'),
 (103, 'CPY-001', 'copieur',  'Kyocera', 'TASKalfa 4054ci',    '2022-02-14', 'Bureau compta', 'en_service'),
 (104, 'SRV-001', 'serveur',  'Dell',    'PowerEdge R450',     '2023-08-01', 'Salle serveurs','en_service'),
 (105, 'UPS-003', 'onduleur', 'APC',     'Easy UPS 3M 60 kVA', '2020-04-22', 'Local TGBT',   'en_maintenance');

INSERT INTO onduleurs (equipement_id, puissance_kva, topologie, nb_batteries, date_dernier_test, autonomie_min) VALUES
 (101, 40.00, 'VFI', 32, '2026-08-15', 18),
 (102, 11.00, 'VI',  12, '2025-12-01', 25),
 (105, 60.00, 'VFI', 40, '2026-09-01', 15);

