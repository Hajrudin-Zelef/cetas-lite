---
id: collect-261001-rattrapage/rattrapage/sql-guide-12
title: "Guide SQL complet — du SELECT à l'optimisation"
domain: rattrapage
role: reference
task: reference
actors: []
dates: ["2026-09-26"]
keywords: []
source: docs/RAG/collect-261001-rattrapage/sql_guide.md
source_anchor: ""
source_lines: [2303, 2506]
sha256: e0e46b7f57e9e86affe4c987ea7e4c9d87a461acf47930dd2637d8c80dd80b76
---

# Guide SQL complet — du SELECT à l'optimisation

Les doublons viennent de 3 sources : données dupliquées, JOIN en éventail (fan-out),
UNION ALL vs UNION.

```sql
-- 1. Détecter les doublons métier (même référence 2 fois)
SELECT reference, COUNT(*) AS nb
FROM equipements
GROUP BY reference
HAVING COUNT(*) > 1;

-- 2. Le fan-out : joindre 2 tables "plusieurs" fait exploser les lignes
-- Un équipement avec 3 interventions ET 2 contrats → 3 × 2 = 6 lignes !
-- → agréger AVANT de joindre :
WITH nb_i AS (
    SELECT equipement_id, COUNT(*) AS nb_interv FROM interventions GROUP BY 1
),
nb_c AS (
    SELECT equipement_id, COUNT(*) AS nb_contrats FROM contrats_maintenance GROUP BY 1
)
SELECT e.reference, COALESCE(nb_i.nb_interv, 0), COALESCE(nb_c.nb_contrats, 0)
FROM equipements e
LEFT JOIN nb_i ON nb_i.equipement_id = e.id
LEFT JOIN nb_c ON nb_c.equipement_id = e.id;

-- 3. Dédupliquer en gardant une ligne (la plus récente) — PostgreSQL
DELETE FROM interventions a
USING interventions b
WHERE a.id > b.id
  AND a.equipement_id = b.equipement_id
  AND a.date_ouverture = b.date_ouverture;
-- MySQL : même logique en auto-jointure DELETE a FROM interventions a JOIN interventions b ON ...
```

Réflexe : après tout JOIN multiple, `COUNT(*)` de contrôle vs `COUNT(DISTINCT id)`
de la table de faits. Écart = fan-out.

---

## 68. Piège : GROUP BY abusif et agrégats trompeurs

```sql
-- ❌ Moyenner des moyennes : FAUX (les groupes n'ont pas le même poids)
SELECT AVG(moy_par_type) FROM (
    SELECT type_interv, AVG(cout_eur) AS moy_par_type FROM interventions GROUP BY 1
) t;

-- ✅ La vraie moyenne globale
SELECT SUM(cout_eur) * 1.0 / COUNT(cout_eur) FROM interventions;
-- = AVG(cout_eur) tout simplement. La moyenne des moyennes n'est juste que si
-- les groupes ont tous le même effectif — rarement le cas.

-- ❌ Sommer après un JOIN qui duplique (fan-out, section 67)
SELECT e.reference, SUM(i.cout_eur)
FROM equipements e JOIN interventions i ON i.equipement_id = e.id
GROUP BY e.reference;   -- OK ici (1 seul "plusieurs")...

-- ...mais avec 2 tables plusieurs, les sommes sont multipliées. Toujours agréger avant de joindre.
```

Autre classique : `HAVING` sans `GROUP BY` (groupe unique implicite — valide mais obscur),
et `GROUP BY` sur une colonne calculée différemment entre SELECT et GROUP BY
(PG exige la même expression ; utiliser l'alias via CTE ou la position `GROUP BY 1`).

---

## 69. Piège : le problème N+1

Le N+1 : 1 requête pour la liste + N requêtes pour le détail de chaque ligne.
Côté application (boucle), pas côté SQL pur — mais c'est le sysadmin qui le voit
dans les logs : 1 001 requêtes pour afficher 1 000 équipements.

```python
# ❌ N+1 : 1 + N requêtes
equips = db.execute("SELECT id, reference FROM equipements").fetchall()
for e in equips:
    nb = db.execute("SELECT COUNT(*) FROM interventions WHERE equipement_id = %s", (e.id,))
```

```sql
-- ✅ 1 seule requête : le JOIN + GROUP BY fait le travail
SELECT e.id, e.reference, COUNT(i.id) AS nb_interv
FROM equipements e
LEFT JOIN interventions i ON i.equipement_id = e.id
GROUP BY e.id, e.reference;
```

Signes dans les logs : des centaines de requêtes identiques à un paramètre près,
en rafale. Correction : jointure, `IN (...)` avec la liste des ids, ou chargement
eager de l'ORM (`selectinload`/`prefetch_related`).

---

## 70. Bonnes pratiques — nommage et style

```sql
-- ✅ Style recommandé dans ce guide
SELECT e.reference,
       e.modele,
       COUNT(i.id) AS nb_interventions
FROM   equipements AS e
LEFT   JOIN interventions AS i
       ON i.equipement_id = e.id
WHERE  e.statut = 'en_service'
GROUP  BY e.reference, e.modele
HAVING COUNT(i.id) > 0
ORDER  BY nb_interventions DESC;
```

Conventions :

- **Tables** : pluriel, minuscules, snake_case (`interventions`, `contrats_maintenance`).
- **Colonnes** : snake_case, sans accent, sans espace (`date_ouverture`, pas `Date d'ouverture`).
- **PK** : `id` (simple). **FK** : `<table_singulier>_id` (`equipement_id`).
- **Booléens** : adjectif/participe (`actif`, `sla_respecte`, `realisee`) — pas `flag1`.
- **Mots-clés** en MAJUSCULES, un par ligne pour les clauses principales.
- **Alias** courts mais lisibles (`e`, `i`, `emp` — pas `a`, `b`, `c`).
- **Virgules en début de ligne** ou fin : choisir UNE convention d'équipe et s'y tenir
  (début de ligne = diffs git plus propres).
- Commenter l'**intention** (`-- SLA : interventions critiques non clôturées`),
  pas l'évidence (`-- sélectionne les employés`).

---

## 71. Bonnes pratiques — requêtes et exploitation

- [ ] `SELECT` explicite (pas de `*`) sauf exploration ad hoc.
- [ ] `WHERE` + `ORDER BY` déterministe (+ PK) avant tout `LIMIT`/`OFFSET`.
- [ ] Requêtes paramétrées dès qu'une valeur vient de l'extérieur (section 63).
- [ ] Transactions pour tout lot multi-instructions (sections 51-54).
- [ ] `EXPLAIN` avant de mettre en production une requête sur table volumineuse.
- [ ] Tester sur un **échantillon** (`TABLESAMPLE` PG / `LIMIT`) avant le full scan.
- [ ] Nommer les contraintes pour des erreurs lisibles (section 42).
- [ ] Dater/ signer les scripts : en-tête `-- auteur, date, objet, ticket`.
- [ ] Versionner les scripts SQL (git) comme le reste du code.
- [ ] Ne jamais lancer un script DDL/DML trouvé sur internet sans le lire :
  chercher `DROP`, `DELETE`, `TRUNCATE`, `;--` dedans.

```sql
-- En-tête type d'un script d'exploitation
-- =====================================================================
-- Objet   : Purge des interventions clôturées depuis plus de 3 ans
-- Auteur  : Service Systèmes & Énergies
-- Date    : 2026-09-26
-- Ticket  : GLPI #4521
-- Cible   : sql_formation (PROD après validation sur PREPROD)
-- Rollback: restaurer depuis sauvegarde_2026-09-26.dump
-- =====================================================================
BEGIN;
-- 1. Contrôle du périmètre
SELECT COUNT(*) FROM interventions
WHERE date_cloture < CURRENT_DATE - INTERVAL '3 years';
-- 2. (après validation du count) DELETE ...
-- COMMIT;
```

---

## 72. Cas pratique 1 — Tableau de bord du service (commenté)

Besoin : pour la réunion hebdo, un tableau par technicien : interventions ouvertes,
clôturées, coût total, respect SLA, ancienneté de la plus vieille ouverte.

```sql
WITH stats AS (
    SELECT
        technicien_id,
        -- pivot manuel portable (section 17)
        SUM(CASE WHEN date_cloture IS NULL THEN 1 ELSE 0 END) AS ouvertes,
        SUM(CASE WHEN date_cloture IS NOT NULL THEN 1 ELSE 0 END) AS cloturees,
        COALESCE(SUM(cout_eur), 0) AS cout_total,
        -- % SLA sur les clôturées évaluées (évite la division entière, section 15)
        CASE WHEN COUNT(*) FILTER (WHERE sla_respecte IS NOT NULL) = 0 THEN NULL
             ELSE SUM(CASE WHEN sla_respecte THEN 1 ELSE 0 END) * 100.0
                / COUNT(*) FILTER (WHERE sla_respecte IS NOT NULL)
        END AS pct_sla,
        MIN(CASE WHEN date_cloture IS NULL THEN date_ouverture END) AS plus_vieille_ouverte
    FROM interventions
    WHERE technicien_id IS NOT NULL
    GROUP BY technicien_id
)
SELECT
    e.nom || ' ' || e.prenom AS technicien,   -- PG ; MySQL : CONCAT(e.nom,' ',e.prenom)
    COALESCE(s.ouvertes, 0)  AS ouvertes,
    COALESCE(s.cloturees, 0) AS cloturees,
    COALESCE(s.cout_total, 0) AS cout_total_eur,
    ROUND(s.pct_sla, 1) AS sla_pct,
    s.plus_vieille_ouverte,
    -- alerte : une ouverte depuis plus de 7 jours
    CASE WHEN s.plus_vieille_ouverte < NOW() - INTERVAL '7 days'
         THEN '⚠️ À relancer' ELSE 'OK' END AS vigilance
FROM employes e
LEFT JOIN stats s ON s.technicien_id = e.id
WHERE e.actif
ORDER BY ouvertes DESC, e.nom;
```

Points à noter : `COALESCE` après LEFT JOIN (les techniciens sans stats ont des NULL),
`FILTER` PG remplacé par `CASE` si besoin de portabilité stricte, `ROUND(pct,1)`
pour un affichage propre.

---

## 73. Cas pratique 2 — Plan de maintenance des onduleurs

