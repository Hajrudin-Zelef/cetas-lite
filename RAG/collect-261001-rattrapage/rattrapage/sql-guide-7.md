---
id: collect-261001-rattrapage/rattrapage/sql-guide-7
title: "Guide SQL complet — du SELECT à l'optimisation"
domain: rattrapage
role: reference
task: reference
actors: ["Oracle"]
dates: ["2020-01-01"]
keywords: []
source: docs/RAG/collect-261001-rattrapage/sql_guide.md
source_anchor: ""
source_lines: [1268, 1495]
sha256: 382034e29e9707907e0834a1ee5b07a0dbee41f4196e7dbb1f85849907b46b97
---

# Guide SQL complet — du SELECT à l'optimisation

Anatomie : `FONCTION(...) OVER (PARTITION BY ... ORDER BY ... ROWS/RANGE ...)`.

- `PARTITION BY` = le « GROUP BY » de la fenêtre (optionnel).
- `ORDER BY` = ordre **dans** la fenêtre (obligatoire pour les classements).
- Sans PARTITION : une seule fenêtre = toute la table.

Cas sysadmin typique : comparer chaque intervention au coût moyen **de son type**,
sur la même ligne, sans perdre le détail :

```sql
SELECT id, type_interv, cout_eur,
       AVG(cout_eur) OVER (PARTITION BY type_interv) AS cout_moyen_type,
       cout_eur - AVG(cout_eur) OVER (PARTITION BY type_interv) AS ecart
FROM interventions;
```

---

## 37. ROW_NUMBER, RANK, DENSE_RANK

```sql
SELECT nom, salaire,
       ROW_NUMBER() OVER (ORDER BY salaire DESC) AS rang_strict,
       RANK()       OVER (ORDER BY salaire DESC) AS rang,
       DENSE_RANK() OVER (ORDER BY salaire DESC) AS rang_dense
FROM employes;
```

| nom | salaire | ROW_NUMBER | RANK | DENSE_RANK |
|---|---|---|---|---|
| Zelef | 5200 | 1 | 1 | 1 |
| Diallo | 3100 | 2 | 2 | 2 |
| Ndiaye | 2950 | 3 | 3 | 3 |
| Sow | 2800 | 4 | 4 | 4 |
| Kane | 1400 | 5 | 5 | 5 |

Avec ex æquo (deux salaires à 3100) : ROW_NUMBER donne 2 et 3 (arbitraire mais unique),
RANK donne 2, 2 puis 4 (trou), DENSE_RANK donne 2, 2 puis 3 (sans trou).

Recette « top-N par groupe » (les 2 dernières interventions de chaque équipement) :

```sql
WITH classees AS (
    SELECT i.*,
           ROW_NUMBER() OVER (PARTITION BY equipement_id
                              ORDER BY date_ouverture DESC) AS rn
    FROM interventions i
)
SELECT * FROM classees WHERE rn <= 2;
```

⚠️ On **ne peut pas** mettre une fonction fenêtre dans WHERE (évaluée après) →
toujours via CTE ou sous-requête comme ci-dessus.

---

## 38. LAG / LEAD — regarder la ligne voisine

```sql
-- Évolution du coût : comparer chaque intervention à la précédente (par équipement)
SELECT equipement_id, date_ouverture, cout_eur,
       LAG(cout_eur) OVER (PARTITION BY equipement_id
                           ORDER BY date_ouverture) AS cout_precedent,
       cout_eur - LAG(cout_eur) OVER (PARTITION BY equipement_id
                                     ORDER BY date_ouverture) AS delta,
       LEAD(date_ouverture) OVER (PARTITION BY equipement_id
                                  ORDER BY date_ouverture) AS prochaine_date
FROM interventions;
```

Arguments : `LAG(colonne [, décalage [, défaut]])`.
Exemple métier : détecter les équipements dont le coût d'intervention **augmente**
trois fois de suite (signe de fin de vie → arbitrage remplacement) :

```sql
WITH c AS (
    SELECT equipement_id, cout_eur,
           LAG(cout_eur, 1) OVER w AS c1,
           LAG(cout_eur, 2) OVER w AS c2
    FROM interventions
    WINDOW w AS (PARTITION BY equipement_id ORDER BY date_ouverture)
)
SELECT DISTINCT equipement_id FROM c
WHERE cout_eur > c1 AND c1 > c2;
```

(`WINDOW w AS (...)` : clause standard pour factoriser la définition de fenêtre —
supportée par PG ; MySQL 8+ aussi depuis 8.0.)

---

## 39. Agrégats et cadres de fenêtre (ROWS / RANGE)

Par défaut, avec ORDER BY dans OVER, le cadre est
`RANGE BETWEEN UNBOUNDED PRECEDING AND CURRENT ROW` : cumul progressif.

```sql
-- Coût cumulé des interventions dans le temps
SELECT date_ouverture::date AS jour, cout_eur,
       SUM(cout_eur) OVER (ORDER BY date_ouverture
                           ROWS BETWEEN UNBOUNDED PRECEDING AND CURRENT ROW) AS cumul
FROM interventions
ORDER BY date_ouverture;

-- Moyenne mobile sur 3 interventions (lignes)
SELECT id, cout_eur,
       AVG(cout_eur) OVER (ORDER BY date_ouverture
                           ROWS BETWEEN 1 PRECEDING AND 1 FOLLOWING) AS moy_mobile_3
FROM interventions;
```

Cadres utiles :

| Cadre | Sens |
|---|---|
| `ROWS BETWEEN UNBOUNDED PRECEDING AND CURRENT ROW` | cumul depuis le début |
| `ROWS BETWEEN 6 PRECEDING AND CURRENT ROW` | 7 dernières lignes (moyenne mobile) |
| `RANGE BETWEEN INTERVAL '30 days' PRECEDING AND CURRENT ROW` | 30 derniers jours (valeurs, pas lignes) |

⚠️ `RANGE` avec offset n'accepte qu'**une** expression d'ordre, et les intervalles
de dates en RANGE sont PG-spécifiques. Pour du portable : `ROWS`.

---

## 40. UNION, INTERSECT, EXCEPT — combiner des requêtes

```sql
-- Tous les identifiants "personnes ou équipements" suivis (ex. inventaire global)
SELECT id, nom AS libelle, 'employe' AS nature FROM employes
UNION
SELECT id, reference, 'equipement' FROM equipements
ORDER BY 1;
```

| Opérateur | Résultat | Doublons |
|---|---|---|
| `UNION` | lignes des deux | dédupliqués |
| `UNION ALL` | lignes des deux | **conservés** (plus rapide) |
| `INTERSECT` | lignes communes | dédupliqués |
| `EXCEPT` | lignes de la 1re absentes de la 2e | dédupliqués |

Règles :

- Même **nombre** de colonnes, types **compatibles** (les noms viennent de la 1re requête).
- `ORDER BY` / `LIMIT` : une seule fois, à la fin, sur l'ensemble.
- `UNION ALL` est plus rapide (pas de tri de déduplication) : l'utiliser dès qu'on sait
  qu'il n'y a pas de doublons ou qu'on les veut.

```sql
-- Techniciens n'ayant fait QUE du curatif (jamais de préventif)
SELECT DISTINCT technicien_id FROM interventions WHERE type_interv = 'curatif'
EXCEPT
SELECT DISTINCT technicien_id FROM interventions WHERE type_interv = 'preventif';
-- MySQL/MariaDB : pas de EXCEPT/INTERSECT → réécrire avec NOT EXISTS (section 31)
```

⚠️ **Portabilité** — `INTERSECT` / `EXCEPT` : PostgreSQL, SQL Server, Oracle, SQLite.
MySQL/MariaDB : **absents** → passer par `IN`/`EXISTS`/`NOT EXISTS`.

---

## 41. Ordre d'exécution logique — à graver

```
FROM / JOIN        → quelles tables, comment jointes
WHERE              → filtre les lignes
GROUP BY           → regroupe
HAVING             → filtre les groupes
SELECT             → projette / calcule (DISTINCT ici)
   └── fenêtres    → OVER() évalué après SELECT de base
ORDER BY           → trie
LIMIT / OFFSET     → pagine
```

Conséquences directes :

1. On ne peut pas utiliser un alias du SELECT dans WHERE (évalué avant).
2. On ne peut pas mettre une fonction fenêtre dans WHERE ou GROUP BY.
3. `ORDER BY` peut utiliser les alias (évalué après SELECT).
4. `WHERE` ne voit pas les agrégats → `HAVING`.

```sql
-- ❌ WHERE sur un alias
SELECT salaire * 12 AS annuel FROM employes WHERE annuel > 40000;
-- ✅
SELECT salaire * 12 AS annuel FROM employes WHERE salaire * 12 > 40000;
-- ✅ (plus propre) : CTE
WITH s AS (SELECT *, salaire * 12 AS annuel FROM employes)
SELECT * FROM s WHERE annuel > 40000;
```

---

## 42. DDL — CREATE TABLE proprement

```sql
CREATE TABLE IF NOT EXISTS maintenances_planifiees (
    id              SERIAL PRIMARY KEY,              -- PG : auto-incrément
    equipement_id   INTEGER NOT NULL REFERENCES equipements(id),
    date_prevue     DATE NOT NULL,
    type_maintenance VARCHAR(20) NOT NULL DEFAULT 'preventif',
    commentaire     TEXT,
    realisee        BOOLEAN NOT NULL DEFAULT FALSE,
    CONSTRAINT ck_date_future CHECK (date_prevue >= DATE '2020-01-01'),
    CONSTRAINT uq_equip_date UNIQUE (equipement_id, date_prevue)
);
```

⚠️ **Portabilité auto-incrément** :

| SGBD | Syntaxe |
|---|---|
| PostgreSQL | `SERIAL` / `BIGSERIAL` (ou `GENERATED ALWAYS AS IDENTITY`, standard) |
| MySQL/MariaDB | `INT AUTO_INCREMENT PRIMARY KEY` |
| Standard SQL:2008 | `GENERATED ALWAYS AS IDENTITY` (PG ≥ 10, Oracle, DB2) |

Conseils de création :

- `IF NOT EXISTS` : rend les scripts rejouables (déploiements, Ansible).
- Nommer les contraintes (`CONSTRAINT ck_...`) : les messages d'erreur deviennent lisibles.
- Choisir le bon type **dès le départ** (section 43) : changer après coûte cher.
- Toujours une PK (même technique) : sans PK, pas de réplication logique fiable,
  pas d'UPDATE ciblé propre, ORM en difficulté.

---

## 43. Types de données — choisir juste

