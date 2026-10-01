---
id: collect-261001-rattrapage/rattrapage/sql-guide-13
title: "Guide SQL complet — du SELECT à l'optimisation"
domain: rattrapage
role: reference
task: reference
actors: []
dates: ["2026-09-01", "2026-10-01"]
keywords: []
source: docs/RAG/collect-261001-rattrapage/sql_guide.md
source_anchor: ""
source_lines: [2507, 2667]
sha256: 098b5da8462ee98b2a993fb318989360d561aa6f2fa39998231316cffafa48f4
---

# Guide SQL complet — du SELECT à l'optimisation

Besoin : lister les onduleurs dont le **test batterie** date de plus de 6 mois
ou n'a jamais eu lieu, avec criticité selon la puissance.

```sql
SELECT
    e.reference,
    e.localisation,
    o.puissance_kva,
    o.date_dernier_test,
    -- jours depuis le dernier test (portable via les deux variantes)
    CURRENT_DATE - o.date_dernier_test AS jours_depuis_test,   -- PG : entier
    CASE
        WHEN o.date_dernier_test IS NULL THEN '🔴 Jamais testé'
        WHEN o.date_dernier_test < CURRENT_DATE - INTERVAL '6 months' AND o.puissance_kva >= 30
            THEN '🔴 Critique : gros onduleur non testé depuis 6 mois'
        WHEN o.date_dernier_test < CURRENT_DATE - INTERVAL '6 months'
            THEN '🟠 À planifier'
        ELSE '🟢 À jour'
    END AS priorite_maintenance,
    -- prochaine échéance : test annuel glissant
    o.date_dernier_test + INTERVAL '1 year' AS prochain_test  -- PG ; MySQL : DATE_ADD(..., INTERVAL 1 YEAR)
FROM equipements e
JOIN onduleurs o ON o.equipement_id = e.id
WHERE e.statut <> 'reforme'
  AND (o.date_dernier_test IS NULL
       OR o.date_dernier_test < CURRENT_DATE - INTERVAL '6 months')
ORDER BY
    CASE WHEN o.date_dernier_test IS NULL THEN 0 ELSE 1 END,  -- jamais testés d'abord
    o.puissance_kva DESC;
```

Variante MySQL de `CURRENT_DATE - date` : `DATEDIFF(CURRENT_DATE, o.date_dernier_test)`.

---

## 74. Cas pratique 3 — Suivi du parc copieurs (coûts par page)

Besoin : coût au millier de pages par copieur pour arbitrer les renouvellements.
(On suppose une table `compteurs_copieurs(copieur_id, date_releve, pages_total)`.)

```sql
-- Consommation entre deux relevés consécutifs (LAG, section 38)
WITH releves AS (
    SELECT
        copieur_id,
        date_releve,
        pages_total,
        pages_total - LAG(pages_total) OVER (
            PARTITION BY copieur_id ORDER BY date_releve
        ) AS pages_periode
    FROM compteurs_copieurs
),
couts AS (
    SELECT equipement_id, SUM(cout_eur) AS cout_total
    FROM interventions
    WHERE type_interv = 'curatif'
    GROUP BY equipement_id
)
SELECT
    e.reference,
    e.modele,
    SUM(r.pages_periode) AS pages_imprimees,
    COALESCE(c.cout_total, 0) AS cout_curatif_eur,
    -- coût pour 1000 pages (protégé contre la division par zéro)
    CASE WHEN SUM(r.pages_periode) > 0
         THEN ROUND(COALESCE(c.cout_total, 0) * 1000.0 / SUM(r.pages_periode), 2)
    END AS cout_pour_1000_pages
FROM equipements e
LEFT JOIN releves r ON r.copieur_id = e.id
LEFT JOIN couts c   ON c.equipement_id = e.id
WHERE e.type_equip = 'copieur'
GROUP BY e.reference, e.modele, c.cout_total
HAVING SUM(r.pages_periode) > 0
ORDER BY cout_pour_1000_pages DESC NULLS LAST;
```

Lecture : en tête du classement = les copieurs les plus chers à la page →
candidats au remplacement ou à la renégociation du contrat de maintenance.

---

## 75. Cas pratique 4 — Reporting mensuel automatisable

Besoin : un rapport mensuel unique, rejouable, qui sort : volume, coûts,
délais moyens de clôture, top 3 des équipements les plus coûteux.

```sql
-- Paramètre : mois à traiter (à changer chaque mois, ou via variable psql)
-- psql : psql -v mois="2026-09-01" -f rapport.sql  → :'mois'
WITH perimetre AS (
    SELECT * FROM interventions
    WHERE date_ouverture >= DATE '2026-09-01'
      AND date_ouverture <  DATE '2026-10-01'
),
delais AS (
    SELECT *,
           -- durée en heures (PG) ; MySQL : TIMESTAMPDIFF(MINUTE,...)/60.0
           EXTRACT(EPOCH FROM (date_cloture - date_ouverture)) / 3600.0 AS delai_h
    FROM perimetre
    WHERE date_cloture IS NOT NULL
)
SELECT 'volume' AS indicateur, COUNT(*)::TEXT AS valeur FROM perimetre
UNION ALL
SELECT 'cout_total_eur', COALESCE(SUM(cout_eur), 0)::TEXT FROM perimetre
UNION ALL
SELECT 'delai_moyen_cloture_h', COALESCE(AVG(delai_h), 0)::TEXT FROM delais
UNION ALL
SELECT 'pct_sla_respecte',
       (SUM(CASE WHEN sla_respecte THEN 1 ELSE 0 END) * 100.0
        / NULLIF(COUNT(*) FILTER (WHERE sla_respecte IS NOT NULL), 0))::TEXT
FROM perimetre;
-- Top 3 équipements (requête séparée, même périmètre)
SELECT e.reference, e.modele, COUNT(p.id) AS nb, SUM(p.cout_eur) AS cout
FROM perimetre p
JOIN equipements e ON e.id = p.equipement_id
GROUP BY e.reference, e.modele
ORDER BY cout DESC NULLS LAST
LIMIT 3;
```

`NULLIF(x, 0)` : retourne NULL si x = 0 → évite la division par zéro **et**
signale « pas de données » (NULL) plutôt qu'un 0 trompeur.

---

## 76. Pense-bête de poche — une page

```
SELECT      colonnes | * (éviter en prod)
FROM        table [AS alias]
JOIN        table ON condition      (INNER / LEFT / RIGHT / FULL / CROSS)
WHERE       filtres lignes          (AND > OR : parenthèses !)
GROUP BY    colonnes de regroupement
HAVING      filtres sur groupes      (après agrégation)
ORDER BY    col [ASC|DESC] [, id]    (toujours PK en dernier pour paginer)
LIMIT n [OFFSET m]                   (OFFSET profond → keyset)

NULL :  IS NULL / IS NOT NULL  (jamais = NULL)
        COALESCE(a, 0)  |  NOT IN ⚠️ NULL → NOT EXISTS
Texte : || (PG) / CONCAT (MySQL) | UPPER LOWER TRIM SUBSTRING LIKE '%motif%'
Dates : CURRENT_DATE | + INTERVAL '7 days' (PG) / DATE_ADD (MySQL)
        mois : >= '2026-09-01' AND < '2026-10-01'
CASE  : CASE WHEN cond THEN x ELSE y END
Agrég : COUNT(*) COUNT(col) SUM AVG MIN MAX (ignorent NULL sauf COUNT(*))
Fenet : ROW_NUMBER() RANK() LAG() LEAD() SUM() OVER (PARTITION BY x ORDER BY y)
CTE   : WITH etape AS (SELECT ...) SELECT * FROM etape
DDL   : CREATE TABLE / ALTER TABLE ADD|DROP COLUMN / DROP TABLE / TRUNCATE
DML   : INSERT (colonnes !) / UPDATE ... WHERE (SELECT d'abord !) / DELETE ... WHERE
TXN   : BEGIN; ... COMMIT;  / ROLLBACK;   (MySQL : DDL = commit implicite !)
Sécur : requêtes PARAMÉTRÉES, jamais de concaténation
Sauve : pg_dump -Fc / mysqldump --single-transaction --routines --triggers
Diag  : EXPLAIN (ANALYZE, BUFFERS) → Seq Scan ? index manquant ? ANALYZE ?
```

À garder sous les yeux : **NULL ≠ 0**, **WHERE avant GROUP BY avant HAVING**,
**condition OUTER JOIN dans le ON**, **paramétrer les entrées**, **backup avant ALTER**.

---

## 77. Glossaire — les 60 termes à connaître

