---
id: collect-261001-rattrapage/rattrapage/sql-guide-3
title: "Guide SQL complet — du SELECT à l'optimisation"
domain: rattrapage
role: reference
task: reference
actors: ["Oracle"]
dates: []
keywords: ["attention"]
source: docs/RAG/collect-261001-rattrapage/sql_guide.md
source_anchor: ""
source_lines: [361, 571]
sha256: e499a869e8a487df8c91dcbb4026be5cf550905797261e00e0c0f681f2cdd49e
---

# Guide SQL complet — du SELECT à l'optimisation

| Motif | Sens |
|---|---|
| `%` | 0, 1 ou plusieurs caractères |
| `_` | exactement 1 caractère |
| `[abc]` | ⚠️ MySQL uniquement avec REGEXP, pas LIKE standard |

Recherche insensible à la casse :

```sql
-- PostgreSQL
SELECT * FROM equipements WHERE marque ILIKE 'apc%';
-- Portable (les deux SGBD)
SELECT * FROM equipements WHERE LOWER(marque) LIKE 'apc%';
```

⚠️ Performance : un `LIKE '%motif'` (motif en tête) **ne peut pas** utiliser un index B-tree
classique → scan complet. Sur des millions de lignes de logs, envisager un index trigramme
(PostgreSQL `pg_trgm`) ou full-text (voir section 60).

---

## 10. ORDER BY — trier

```sql
-- Tri simple
SELECT nom, salaire FROM employes ORDER BY salaire DESC;

-- Tri multiple : salaire décroissant, puis nom croissant
SELECT nom, salaire FROM employes ORDER BY salaire DESC, nom ASC;

-- Trier par position (1 = 1re colonne) — pratique mais fragile si on modifie le SELECT
SELECT nom, salaire * 12 FROM employes ORDER BY 2 DESC;

-- Trier par alias
SELECT nom, salaire * 12 AS annuel FROM employes ORDER BY annuel DESC;

-- NULLS FIRST / LAST (standard ; MySQL : NULLS considérés comme plus petits par défaut)
SELECT nom, date_dernier_test FROM employes e
LEFT JOIN ... -- exemple section 24
ORDER BY date_dernier_test DESC NULLS LAST;
```

⚠️ **Portabilité** — `NULLS FIRST` / `NULLS LAST` : standard SQL, supporté par PostgreSQL.
MySQL/MariaDB ne les acceptent pas ; émuler avec `ORDER BY date_dernier_test IS NULL, date_dernier_test DESC`
(`FALSE`=0 avant `TRUE`=1, donc les non-NULL d'abord).

Tri stable : `ORDER BY` sur une colonne non unique ne garantit **aucun** ordre pour les ex æquo.
Pour un ordre déterministe (pagination !), toujours ajouter la clé primaire en dernier :

```sql
SELECT * FROM interventions ORDER BY date_ouverture DESC, id DESC;
```

---

## 11. LIMIT / OFFSET — paginer

```sql
-- Les 5 interventions les plus récentes
SELECT * FROM interventions ORDER BY date_ouverture DESC LIMIT 5;

-- Page 2 de 10 lignes (lignes 11 à 20)
SELECT * FROM interventions ORDER BY id LIMIT 10 OFFSET 10;
```

⚠️ **Portabilité** — `LIMIT n OFFSET m` : PostgreSQL, MySQL, MariaDB, SQLite.
Syntaxe standard SQL:2008 (SQL Server, Oracle 12c+) : `OFFSET 10 ROWS FETCH NEXT 10 ROWS ONLY`.

Pièges de la pagination par OFFSET :

1. **Sans ORDER BY déterministe**, les pages peuvent contenir des doublons ou sauter des lignes
   si des insertions ont lieu entre deux pages.
2. **OFFSET grand = lent** : `OFFSET 100000` force le SGBD à lire et jeter 100 000 lignes.
   Pour de gros volumes, préférer la pagination par curseur (keyset) :

```sql
-- Pagination keyset : "la page suivante après l'id 1004"
SELECT * FROM interventions
WHERE id > 1004
ORDER BY id
LIMIT 10;
```

C'est la méthode à utiliser pour exporter des millions de lignes de logs ou d'historique.

---

## 12. SELECT DISTINCT — éliminer les doublons

```sql
-- Quelles marques sont présentes dans le parc ?
SELECT DISTINCT marque FROM equipements;

-- DISTINCT sur plusieurs colonnes = combinaison unique
SELECT DISTINCT type_equip, marque FROM equipements ORDER BY 1, 2;

-- Compter les valeurs distinctes
SELECT COUNT(DISTINCT technicien_id) AS nb_techniciens_actifs FROM interventions;
```

Attention : `DISTINCT` s'applique à **toute** la ligne projetée, pas à une colonne isolée.
`SELECT DISTINCT nom, salaire` déduplique les paires (nom, salaire).

⚠️ `COUNT(DISTINCT col)` ignore les NULL (comme toutes les fonctions d'agrégat, section 19).
MySQL accepte `COUNT(DISTINCT a, b)` (plusieurs colonnes) ; PostgreSQL non —
émuler avec `COUNT(DISTINCT (a, b))` (ligne composite) ou `COUNT(DISTINCT a || '|' || b)`.

---

## 13. NULL et la logique à trois valeurs — LE concept clé

En SQL, `NULL` signifie **« valeur inconnue / absente »**, pas zéro, pas chaîne vide.
Toute comparaison avec NULL donne NULL (ni vrai ni faux) : c'est la **logique à trois valeurs**
(TRUE / FALSE / UNKNOWN), et `WHERE` ne garde que les lignes **TRUE**.

```sql
-- ❌ NE MARCHE PAS : rien ne vaut NULL, pas même NULL
SELECT * FROM interventions WHERE date_cloture = NULL;   -- 0 ligne, toujours !
SELECT * FROM interventions WHERE date_cloture <> NULL;   -- 0 ligne aussi !

-- ✅ CORRECT
SELECT * FROM interventions WHERE date_cloture IS NULL;      -- en cours
SELECT * FROM interventions WHERE date_cloture IS NOT NULL;  -- clôturées
```

Le piège `NOT IN` avec NULL — **à connaître absolument** :

```sql
-- Si la sous-requête renvoie ne serait-ce qu'UN NULL, NOT IN renvoie... RIEN.
SELECT * FROM employes
WHERE id NOT IN (SELECT technicien_id FROM interventions);
-- technicien_id peut être NULL (intervention non assignée) → résultat vide, sans erreur !

-- ✅ Solutions sûres
SELECT * FROM employes e
WHERE NOT EXISTS (SELECT 1 FROM interventions i WHERE i.technicien_id = e.id);

-- ou filtrer les NULL :
SELECT * FROM employes
WHERE id NOT IN (SELECT technicien_id FROM interventions WHERE technicien_id IS NOT NULL);
```

Propagation de NULL dans les calculs :

```sql
SELECT
    cout_eur,                          -- 850.00
    cout_eur + 100,                    -- 950.00
    NULL + 100,                        -- NULL ! tout calcul avec NULL donne NULL
    COALESCE(cout_eur, 0) + 100        -- 100 si cout_eur est NULL (COALESCE = premier non-NULL)
FROM interventions;
```

Fonctions utiles :

| Fonction | Rôle | Exemple |
|---|---|---|
| `COALESCE(a, b, ...)` | premier argument non NULL | `COALESCE(sla_respecte, FALSE)` |
| `NULLIF(a, b)` | NULL si a = b | `NULLIF(division_par_zero, 0)` pattern |
| `GREATEST/LEAST` | max/min en ignorant... ⚠️ NULL les contamine sous PG | — |

⚠️ `GREATEST(10, NULL)` → NULL sous PostgreSQL (contamination). MySQL : idem.
Toujours `COALESCE` avant, ou filtrer.

---

## 14. Fonctions sur les chaînes de caractères

```sql
SELECT
    UPPER(nom)                 AS nom_maj,        -- 'ZELEF'
    LOWER(marque)              AS marque_min,     -- 'apc'
    LENGTH(modele)             AS longueur,       -- nb de caractères
    TRIM('  test  ')           AS nettoye,        -- 'test'
    SUBSTRING(modele FROM 1 FOR 4) AS prefixe,   -- 'Easy'  (standard)
    POSITION('UPS' IN modele)  AS pos,            -- position (1-based), 0 si absent
    REPLACE(modele, 'UPS', 'ASI') AS fr,          -- remplacement
    nom || ' ' || prenom      AS nom_complet      -- concaténation (standard ||)
FROM equipements;
```

⚠️ **Portabilité** — concaténation :

| SGBD | Syntaxe |
|---|---|
| Standard / PostgreSQL / SQLite | `'a' \|\| 'b'` |
| MySQL/MariaDB | `CONCAT('a','b')` (par défaut `\|\|` = OU logique !) |
| MySQL avec `PIPES_AS_CONCAT` | `\|\|` fonctionne aussi |

```sql
-- Extraction : tout ce qui est après le dernier espace (ex. "40 kVA" → "kVA")
-- PostgreSQL :
SELECT SUBSTRING(modele FROM '([0-9]+ kVA)$') FROM equipements;
-- MySQL 8+ :
SELECT REGEXP_SUBSTR(modele, '[0-9]+ kVA$') FROM equipements;

-- Découper une référence 'UPS-001' :
SELECT
    SPLIT_PART(reference, '-', 1) AS prefixe,   -- PostgreSQL : 'UPS'
    SPLIT_PART(reference, '-', 2) AS numero     -- '001'
FROM equipements;
-- MySQL : SUBSTRING_INDEX(reference, '-', 1) / SUBSTRING_INDEX(reference, '-', -1)
```

Règle sysadmin : normaliser les données **à l'écriture** (CHECK, trigger, application),
pas avec des fonctions dans chaque requête — sinon les index ne servent plus (section 60).

---

## 15. Fonctions numériques

