---
id: collect-261001-rattrapage/rattrapage/sql-guide-2
title: "Guide SQL complet — du SELECT à l'optimisation"
domain: rattrapage
role: reference
task: reference
actors: []
dates: ["2026-09-01", "2026-09-10", "2026-09-20", "2026-09-21", "2026-09-22", "2026-09-25", "2026-09-26", "2026-10-01"]
keywords: ["arr", "attention", "valuation"]
source: docs/RAG/collect-261001-rattrapage/sql_guide.md
source_anchor: ""
source_lines: [155, 360]
sha256: a283138a6e94685137c5c02c920c1387c8a8490255b324b3b76d09fba4ca7a45
---

# Guide SQL complet — du SELECT à l'optimisation

INSERT INTO interventions (id, equipement_id, technicien_id, date_ouverture, date_cloture,
                           type_interv, priorite, description, cout_eur, sla_respecte) VALUES
 (1001, 102, 2, '2026-09-20 08:15:00', '2026-09-20 14:30:00', 'curatif',   1, 'Onduleur en défaut batterie, remplacement bloc 7', 850.00, TRUE),
 (1002, 103, 4, '2026-09-21 09:00:00', '2026-09-22 11:00:00', 'curatif',   3, 'Bourrage papier récurrent tiroir 2',                 120.00, TRUE),
 (1003, 105, 2, '2026-09-25 07:45:00', NULL,                 'preventif', 2, 'Test batteries + thermographie armoire',             0.00,   NULL),
 (1004, 101, 5, '2026-09-26 10:00:00', NULL,                 'curatif',   2, 'Alarme température ambiante local TGBT',             0.00,   NULL),
 (1005, 104, 3, '2026-09-10 14:00:00', '2026-09-10 16:00:00', 'installation','4','Mise en baie + câblage',                         300.00, TRUE);
```

> Astuce sysadmin : gardez ce script dans `~/workspace/sql/00_init.sql`. Le rejouer
> (DROP + CREATE) remet la base de démo dans un état propre avant chaque session d'entraînement.

---

## 4. SELECT — lire des données

Le `SELECT` est l'ordre le plus utilisé. Forme canonique :

```sql
SELECT colonne1, colonne2
FROM   table
WHERE  condition
ORDER  BY colonne1;
```

Exemples progressifs :

```sql
-- 1. Tout lire (à éviter en production sur de grosses tables !)
SELECT * FROM employes;

-- 2. Choisir les colonnes (toujours préférer la liste explicite)
SELECT nom, prenom, poste FROM employes;

-- 3. Renommer à l'affichage avec AS
SELECT nom AS "Nom", prenom AS "Prénom", salaire * 12 AS salaire_annuel
FROM employes;

-- 4. Constantes et calculs dans le SELECT
SELECT reference, modele, 'PARC-SETIS' AS inventaire, puissance_kva * 0.8 AS puissance_kw_utile
FROM equipements e
JOIN onduleurs o ON o.equipement_id = e.id;
```

Règles d'or :

- En exploitation, **ne jamais** faire `SELECT *` sur une table métier : on ramène des colonnes
  inutiles (parfois lourdes : TEXT, BYTEA), on casse les applis si une colonne est ajoutée,
  et on empêche l'optimiseur d'utiliser un index couvrant (section 58).
- `AS` est optionnel (`SELECT nom n FROM employes` fonctionne) mais l'écrire rend la requête lisible.
- Les alias avec espaces ou accents exigent des guillemets doubles `"Prénom"` (standard).
  ⚠️ MySQL accepte aussi les backticks `` `Prénom` `` — non portable, à éviter.

---

## 5. WHERE — filtrer les lignes

`WHERE` ne garde que les lignes vérifiant la condition. Les opérateurs de base :

```sql
-- Égalité / inégalité
SELECT * FROM employes WHERE service = 'Maintenance';
SELECT * FROM equipements WHERE statut <> 'reforme';   -- <> est le standard ; != accepté partout

-- Comparaisons numériques
SELECT nom, salaire FROM employes WHERE salaire >= 3000;

-- Dates : toujours au format ISO entre apostrophes
SELECT * FROM interventions WHERE date_ouverture >= '2026-09-01';

-- Plusieurs conditions : AND a priorité sur OR → parenthèses obligatoires
SELECT * FROM interventions
WHERE type_interv = 'curatif'
  AND (priorite <= 2 OR sla_respecte = FALSE);
```

Ordre d'évaluation logique : `NOT` > `AND` > `OR`. En cas de doute, **parenthèses explicites**.
C'est l'erreur n°1 dans les filtres d'astreinte (« pannes critiques OU préventif du mois »
mal parenthésé = tout le préventif de l'année qui remonte).

```sql
-- ❌ FAUX (AND évalué d'abord) : ramène TOUT le préventif, même priorité 5
SELECT * FROM interventions
WHERE type_interv = 'curatif' AND priorite <= 2 OR type_interv = 'preventif';

-- ✅ CORRECT
SELECT * FROM interventions
WHERE (type_interv = 'curatif' AND priorite <= 2)
   OR type_interv = 'preventif';
```

---

## 6. Opérateurs de comparaison — le détail

| Opérateur | Sens | Exemple |
|---|---|---|
| `=` | égal | `statut = 'en_panne'` |
| `<>` | différent (standard) | `statut <> 'reforme'` |
| `!=` | différent (accepté partout, non standard) | — |
| `<`, `<=`, `>`, `>=` | comparaisons | `priorite <= 2` |
| `BETWEEN a AND b` | entre a et b **inclus** | `cout_eur BETWEEN 100 AND 1000` |
| `IN (...)` | dans une liste | `type_equip IN ('onduleur','copieur')` |
| `LIKE` | motif (section 10) | `modele LIKE 'Easy UPS%'` |
| `IS NULL` / `IS NOT NULL` | test de NULL | `date_cloture IS NULL` |
| `EXISTS` | existence (section 31) | — |

Points d'attention :

```sql
-- BETWEEN est INCLUSIF des deux bornes
SELECT * FROM interventions WHERE cout_eur BETWEEN 100 AND 1000;
-- équivaut à : cout_eur >= 100 AND cout_eur <= 1000

-- Comparer des dates : BETWEEN sur TIMESTAMP inclut minuit pile de la borne haute
-- Pour "tout le mois de septembre", préférer :
SELECT * FROM interventions
WHERE date_ouverture >= '2026-09-01' AND date_ouverture < '2026-10-01';

-- Comparaison de chaînes : sensible à la casse selon le collationnement
-- 'APC' <> 'apc' en général. Pour insensible à la casse :
-- PostgreSQL : WHERE marque ILIKE 'apc'
-- MySQL (collation *_ci par défaut) : WHERE marque = 'apc' suffit déjà
-- Portable : WHERE LOWER(marque) = 'apc'
```

⚠️ **Portabilité** — `ILIKE` n'existe que sous PostgreSQL. `LOWER(col) = 'apc'` marche partout.

---

## 7. Opérateurs logiques AND / OR / NOT

```sql
-- Interventions critiques non clôturées
SELECT id, description FROM interventions
WHERE priorite <= 2 AND date_cloture IS NULL;

-- Équipements à surveiller : en panne OU en maintenance
SELECT reference, statut FROM equipements
WHERE statut = 'en_panne' OR statut = 'en_maintenance';

-- NOT : tout sauf...
SELECT * FROM employes WHERE NOT (actif = FALSE);
-- équivaut à : WHERE actif <> FALSE  (mais attention à NULL, section 13 !)
```

Table de vérité avec NULL (à connaître par cœur, voir section 13) :

| A | B | A AND B | A OR B | NOT A |
|---|---|---|---|---|
| TRUE | TRUE | TRUE | TRUE | FALSE |
| TRUE | FALSE | FALSE | TRUE | FALSE |
| TRUE | NULL | **NULL** | TRUE | **NULL** |
| FALSE | NULL | FALSE | **NULL** | **NULL** |

Conséquence pratique : `WHERE NOT (sla_respecte = TRUE)` **exclut** les lignes où
`sla_respecte` est NULL (interventions en cours). Pour les inclure :

```sql
SELECT * FROM interventions
WHERE sla_respecte IS DISTINCT FROM TRUE;   -- standard SQL, supporté par PostgreSQL
-- MySQL/MariaDB : WHERE NOT (sla_respecte <=> TRUE)  -- opérateur <=> "NULL-safe equal"
-- Portable : WHERE sla_respecte = FALSE OR sla_respecte IS NULL
```

---

## 8. BETWEEN, IN, NOT IN

```sql
-- BETWEEN : intervalles inclusifs
SELECT nom, salaire FROM employes WHERE salaire BETWEEN 2500 AND 3500;

-- IN : liste de valeurs (plus lisible qu'une cascade de OR)
SELECT reference, type_equip FROM equipements
WHERE type_equip IN ('onduleur', 'copieur', 'serveur');

-- NOT IN : ⚠️ DANGER avec NULL (voir section 13)
SELECT * FROM employes WHERE id NOT IN (2, 3);

-- IN avec sous-requête (section 30)
SELECT * FROM equipements
WHERE id IN (SELECT equipement_id FROM interventions WHERE priorite = 1);
```

Performance : `IN` avec une petite liste de constantes est parfait. Avec une sous-requête
retournant des milliers de lignes, `EXISTS` (section 31) est souvent meilleur — l'optimiseur
moderne gère bien les deux, mais `EXISTS` s'arrête dès la première correspondance trouvée.

---

## 9. LIKE — recherche par motif

```sql
-- % : n'importe quelle séquence (y compris vide)
SELECT * FROM equipements WHERE modele LIKE 'Easy UPS%';   -- commence par
SELECT * FROM equipements WHERE modele LIKE '%kVA';        -- finit par
SELECT * FROM equipements WHERE modele LIKE '%UPS%';       -- contient

-- _ : exactement un caractère
SELECT * FROM equipements WHERE reference LIKE 'UPS-___'; -- UPS- + 3 caractères

-- Échapper un % ou _ littéral avec ESCAPE (ou le backslash sous MySQL)
SELECT * FROM equipements WHERE modele LIKE '100\%%' ESCAPE '\';  -- contient "100%"
```

