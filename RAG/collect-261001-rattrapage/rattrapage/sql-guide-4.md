---
id: collect-261001-rattrapage/rattrapage/sql-guide-4
title: "Guide SQL complet — du SELECT à l'optimisation"
domain: rattrapage
role: reference
task: reference
actors: []
dates: ["2026-09-01", "2026-09-26", "2026-13-45"]
keywords: []
source: docs/RAG/collect-261001-rattrapage/sql_guide.md
source_anchor: ""
source_lines: [572, 796]
sha256: 4f896414f7efb98a1ec017d81338f7e08de011a428f686ae1fea5aa3ff24bebd
---

# Guide SQL complet — du SELECT à l'optimisation

```sql
SELECT
    ROUND(123.456, 2)      AS arrondi,      -- 123.46
    TRUNC(123.456, 1)      AS tronque_pg,   -- 123.4 (PostgreSQL ; MySQL: TRUNCATE(123.456,1))
    CEIL(123.1)            AS plafond,      -- 124
    FLOOR(123.9)           AS plancher,     -- 123
    ABS(-42)               AS absolu,       -- 42
    MOD(17, 5)             AS reste,        -- 2
    POWER(2, 10)           AS puissance,    -- 1024
    SQRT(144)              AS racine,       -- 12
    17 % 5                 AS reste_op,     -- 2 (opérateur %, portable)
    7 / 2                  AS division;     -- ⚠️ 3 (entier) sous PG si 2 entiers ; 3.5 sous MySQL !
```

⚠️ **Division entière** — le piège classique des pourcentages :

```sql
-- ❌ Sous PostgreSQL : 1/2 = 0 → pourcentage toujours 0 !
SELECT sla_respecte, COUNT(*) * 100 / COUNT(*) FROM interventions GROUP BY 1;

-- ✅ Forcer le décimal : multiplier par 100.0 ou caster
SELECT
    COUNT(*) FILTER (WHERE sla_respecte) * 100.0 / COUNT(*) AS pct_sla
FROM interventions WHERE sla_respecte IS NOT NULL;
-- MySQL : la division / donne déjà un décimal, mais 100.0 ne coûte rien et reste portable.
```

Autres utiles : `RANDOM()` (PG) / `RAND()` (MySQL) pour échantillonner,
`GREATEST/LEAST` avec COALESCE (section 13).

```sql
-- Échantillon aléatoire de 10 interventions pour audit qualité
-- PostgreSQL :
SELECT * FROM interventions ORDER BY RANDOM() LIMIT 10;
-- MySQL :
SELECT * FROM interventions ORDER BY RAND() LIMIT 10;
-- ⚠️ Sur de grosses tables c'est un tri complet : OK pour audit, pas en production continue.
```

---

## 16. Fonctions de date et heure — l'essentiel survie

```sql
-- "Maintenant"
SELECT CURRENT_DATE,              -- 2026-09-26
       CURRENT_TIMESTAMP,         -- 2026-09-26 23:41:00+...
       NOW();                     -- idem (PG et MySQL)

-- Extraire des parties
SELECT
    EXTRACT(YEAR  FROM date_ouverture) AS annee,
    EXTRACT(MONTH FROM date_ouverture) AS mois,
    EXTRACT(DOW   FROM date_ouverture) AS jour_semaine_pg,  -- 0=dimanche (PG)
    DATE_TRUNC('month', date_ouverture) AS mois_tronque     -- 2026-09-01 00:00
FROM interventions;
-- MySQL : YEAR(d), MONTH(d), DAYOFWEEK(d) (1=dimanche), DATE_FORMAT(d,'%Y-%m-01')
```

Arithmétique de dates (⚠️ grosse différence de syntaxe) :

```sql
-- PostgreSQL : intervalles
SELECT date_ouverture + INTERVAL '2 hours' FROM interventions;
SELECT * FROM interventions WHERE date_ouverture >= NOW() - INTERVAL '7 days';
SELECT AGE(date_cloture, date_ouverture) AS duree FROM interventions;  -- type interval

-- MySQL/MariaDB : fonctions
SELECT DATE_ADD(date_ouverture, INTERVAL 2 HOUR) FROM interventions;
SELECT * FROM interventions WHERE date_ouverture >= NOW() - INTERVAL 7 DAY;
SELECT TIMESTAMPDIFF(HOUR, date_ouverture, date_cloture) AS duree_h FROM interventions;
```

Durées d'intervention (portable, en heures décimales) :

```sql
-- PostgreSQL
SELECT id, EXTRACT(EPOCH FROM (date_cloture - date_ouverture))/3600.0 AS duree_h
FROM interventions WHERE date_cloture IS NOT NULL;

-- MySQL
SELECT id, TIMESTAMPDIFF(MINUTE, date_ouverture, date_cloture)/60.0 AS duree_h
FROM interventions WHERE date_cloture IS NOT NULL;
```

Recettes anti-erreur :

- Stocker en `TIMESTAMP`/`DATETIME`, **jamais** en VARCHAR.
- Écrire les littéraux en ISO `'2026-09-26'` / `'2026-09-26 14:30:00'`.
- Comparer un mois avec `>= premier_jour AND < premier_jour_suivant` (section 6), pas avec `BETWEEN`.
- ⚠️ Timezones : `TIMESTAMP WITH TIME ZONE` (PG) vs `TIMESTAMP` (MySQL, converti en UTC au stockage).
  Pour un parc multi-sites, stocker en UTC et convertir à l'affichage.

---

## 17. CASE — le couteau suisse conditionnel

`CASE` évalue des conditions ligne par ligne. Deux formes :

```sql
-- Forme "simple" : égalité sur une expression
SELECT nom,
       CASE poste
           WHEN 'Chef de service' THEN 'Encadrement'
           WHEN 'Apprenti'        THEN 'Formation'
           ELSE 'Technique'
       END AS categorie
FROM employes;

-- Forme "recherchée" : conditions libres (la plus utile)
SELECT reference, statut,
       CASE
           WHEN statut = 'en_panne' THEN '🔴 Intervention immédiate'
           WHEN statut = 'en_maintenance' THEN '🟡 Suivi en cours'
           WHEN garantie_jusqu < CURRENT_DATE THEN '🟠 Garantie expirée'
           ELSE '🟢 OK'
       END AS etat_affiche
FROM equipements;
```

`CASE` dans les agrégats = **pivot manuel** (hyper utile en reporting) :

```sql
-- Nombre d'interventions par type, en colonnes
SELECT
    COUNT(*) AS total,
    COUNT(*) FILTER (WHERE type_interv = 'curatif')   AS curatif,    -- PG
    COUNT(*) FILTER (WHERE type_interv = 'preventif') AS preventif,
    SUM(CASE WHEN sla_respecte THEN 1 ELSE 0 END)     AS sla_ok      -- portable
FROM interventions;
```

⚠️ **Portabilité** — `FILTER (WHERE ...)` : PostgreSQL et SQLite uniquement.
La version `SUM(CASE WHEN ... THEN 1 ELSE 0 END)` marche **partout** : à privilégier
dans le code partagé.

`CASE` retourne NULL si aucun WHEN ne matche et pas de ELSE → toujours mettre un `ELSE`
explicite sauf besoin précis.

---

## 18. CAST et conversions de types

```sql
-- Syntaxe standard
SELECT CAST(salaire AS INTEGER) FROM employes;          -- tronque : 5200
SELECT salaire::INTEGER FROM employes;                  -- opérateur PG uniquement

-- Conversions courantes
SELECT
    CAST('2026-09-26' AS DATE),
    CAST('42' AS INTEGER),
    CAST(123.456 AS NUMERIC(10,2)),
    CAST(id AS VARCHAR) || '-SETIS' AS code_inv
FROM employes;

-- Conversions sûres : TRY_CAST (SQL Server), PG 16+ :
SELECT PG_TRY_CAST('abc' AS INTEGER);   -- NULL au lieu d'erreur (PostgreSQL 16+)
-- MySQL : CAST('abc' AS UNSIGNED) → 0 avec warning (pas d'erreur !)
```

Règle : convertir **explicitement** plutôt que de compter sur la conversion implicite —
les comportements diffèrent (`'2026-13-45'` : erreur PG, warning+NULL MySQL selon sql_mode).

---

## 19. Agrégations — COUNT, SUM, AVG, MIN, MAX

Les fonctions d'agrégat réduisent un ensemble de lignes en **une** valeur.
Point fondamental : **elles ignorent les NULL** (sauf `COUNT(*)`).

```sql
SELECT
    COUNT(*)              AS nb_lignes,        -- compte les lignes, NULL inclus
    COUNT(date_cloture)   AS nb_cloturees,     -- ignore les NULL !
    COUNT(DISTINCT technicien_id) AS nb_tech,  -- valeurs distinctes non NULL
    SUM(cout_eur)         AS cout_total,
    AVG(cout_eur)         AS cout_moyen,       -- moyenne des non-NULL
    MIN(date_ouverture)   AS premiere,
    MAX(date_ouverture)   AS derniere
FROM interventions;
```

| Fonction | Ignore NULL ? | Ensemble vide → |
|---|---|---|
| `COUNT(*)` | non (compte lignes) | 0 |
| `COUNT(col)` | oui | 0 |
| `SUM(col)` | oui | NULL |
| `AVG(col)` | oui | NULL |
| `MIN/MAX(col)` | oui | NULL |

```sql
-- Coût moyen PAR TECHNICIEN, en traitant les interventions sans coût
SELECT technicien_id,
       COUNT(*) AS nb,
       COALESCE(SUM(cout_eur), 0) AS total,          -- SUM vide = NULL → 0
       COALESCE(AVG(cout_eur), 0) AS moyenne
FROM interventions
GROUP BY technicien_id;
```

⚠️ `AVG` sur des entiers : `AVG(priorite)` → PG renvoie `numeric` (décimal ✓),
MySQL renvoie un décimal aussi (4 décimales). Pas de piège ici, contrairement à `/`.

---

## 20. GROUP BY — regrouper avant d'agréger

```sql
-- Coût total par type d'intervention
SELECT type_interv, COUNT(*) AS nb, SUM(cout_eur) AS cout_total
FROM interventions
GROUP BY type_interv;

-- Par mois (tronquer la date d'abord)
SELECT DATE_TRUNC('month', date_ouverture) AS mois, COUNT(*) AS nb
FROM interventions
GROUP BY 1
ORDER BY 1;
-- MySQL : SELECT DATE_FORMAT(date_ouverture,'%Y-%m') AS mois, ... GROUP BY 1
```

Règle d'or : **toute colonne du SELECT non agrégée DOIT être dans le GROUP BY**
(ou être fonctionnellement dépendante de la clé groupée).

