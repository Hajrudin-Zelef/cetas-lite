---
id: collect-261001-rattrapage/rattrapage/sql-guide-6
title: "Guide SQL complet — du SELECT à l'optimisation"
domain: rattrapage
role: reference
task: reference
actors: []
dates: ["2026-09-01", "2026-09-30"]
keywords: ["arr"]
source: docs/RAG/collect-261001-rattrapage/sql_guide.md
source_anchor: ""
source_lines: [1026, 1267]
sha256: 428e77bdb4d6ff21bb6a445baaf8b7683ed7e9276d97a07733a7d47cd0895066
---

# Guide SQL complet — du SELECT à l'optimisation

-- Cas légitime :
SELECT * FROM equipements e JOIN onduleurs o USING (equipement_id);
-- Note : la colonne fusionnée n'apparaît qu'une fois dans SELECT *.
```

`NATURAL JOIN` joint sur **toutes** les colonnes de même nom, implicitement.
**À proscrire** : l'ajout futur d'une colonne homonyme change silencieusement la jointure.
Règle d'équipe : `ON` explicite, toujours.

---

## 29. Ordre et parenthésage des JOIN multiples

```sql
-- Chaîne de 3 tables : interventions → équipements (+ onduleurs) → techniciens
SELECT i.id, e.reference, o.puissance_kva, emp.nom AS tech
FROM interventions i
JOIN equipements e ON e.id = i.equipement_id
LEFT JOIN onduleurs o ON o.equipement_id = e.id   -- NULL si pas un onduleur
JOIN employes emp    ON emp.id = i.technicien_id;
```

L'ordre compte pour les OUTER JOIN enchaînés : `(A LEFT JOIN B) LEFT JOIN C`
≠ `A LEFT JOIN (B LEFT JOIN C)` dans certains cas. En cas de doute, parenthéser
explicitement — le SGBD respecte les parenthèses.

Astuce de lecture : partir de la table « sujet » (souvent la table de faits :
`interventions`), puis rattacher les dimensions (`equipements`, `employes`).

---

## 30. Sous-requêtes avec IN

Une sous-requête est un SELECT dans un SELECT. Avec `IN` :

```sql
-- Équipements ayant eu au moins une intervention curative prioritaire
SELECT reference, modele FROM equipements
WHERE id IN (
    SELECT equipement_id FROM interventions
    WHERE type_interv = 'curatif' AND priorite <= 2
);

-- IN avec liste ET sous-requête combinées : impossible directement → UNION dans la sous-requête
SELECT * FROM equipements
WHERE id IN (
    SELECT equipement_id FROM interventions WHERE priorite = 1
    UNION
    SELECT 105  -- ajout manuel : l'onduleur sous alarme
);
```

Rappel critique (section 13) : si la sous-requête peut renvoyer NULL,
`NOT IN` renvoie zéro ligne. `IN` (positif), lui, n'a pas ce problème :
`x IN (1, 2, NULL)` vaut TRUE si x=1, FALSE si x=3, UNKNOWN si x=NULL→ filtré.

---

## 31. EXISTS / NOT EXISTS — le test d'existence

`EXISTS` renvoie TRUE dès qu'**une** ligne correspond — sans NULL-piège, sans doublons.

```sql
-- Techniciens ayant réalisé au moins une intervention (semi-join)
SELECT nom, prenom FROM employes e
WHERE EXISTS (
    SELECT 1 FROM interventions i WHERE i.technicien_id = e.id
);

-- Équipements SANS intervention (anti-join) — la forme recommandée
SELECT reference FROM equipements e
WHERE NOT EXISTS (
    SELECT 1 FROM interventions i WHERE i.equipement_id = e.id
);
```

Comparatif des anti-joins (résultat identique, sémantiques différentes) :

| Forme | NULL-piège ? | Lisibilité |
|---|---|---|
| `NOT EXISTS (...)` | ✅ non | très bonne |
| `LEFT JOIN ... WHERE droite.id IS NULL` | ✅ non | bonne |
| `NOT IN (sous-requête)` | ❌ **OUI si NULL** | trompeuse |

**Règle d'équipe : pour « les X sans Y », utiliser `NOT EXISTS`.**
C'est aussi celle que l'optimiseur traite le mieux en général (semi/anti-join natifs
dans les plans d'exécution).

---

## 32. Sous-requêtes corrélées — la requête dans la boucle

Une sous-requête **corrélée** référence une colonne de la requête externe :
elle est réévaluée pour chaque ligne candidate (conceptuellement).

```sql
-- La dernière intervention de CHAQUE équipement
SELECT e.reference,
       (SELECT MAX(date_ouverture) FROM interventions i
         WHERE i.equipement_id = e.id) AS derniere_interv
FROM equipements e;

-- Employés payés plus que la moyenne de leur... (ici moyenne globale, non corrélée)
-- Version corrélée : techniciens dont le coût total dépasse la moyenne des techniciens
SELECT e.nom, tot.total
FROM employes e
JOIN (SELECT technicien_id, SUM(cout_eur) AS total
      FROM interventions GROUP BY technicien_id) tot ON tot.technicien_id = e.id
WHERE tot.total > (SELECT AVG(total) FROM
                   (SELECT SUM(cout_eur) AS total FROM interventions
                    GROUP BY technicien_id) t);
```

⚠️ Performance : une sous-requête corrélée dans le SELECT s'exécute N fois
(N = lignes externes). Souvent mieux : la réécrire en JOIN + GROUP BY,
ou utiliser une fonction fenêtre (section 36+).

---

## 33. Sous-requêtes dans FROM (tables dérivées) et LATERAL

```sql
-- Table dérivée : obligatoire d'aliaser (t)
SELECT t.type_interv, t.nb, t.cout_total
FROM (
    SELECT type_interv, COUNT(*) AS nb, SUM(cout_eur) AS cout_total
    FROM interventions
    GROUP BY type_interv
) AS t
WHERE t.nb > 1
ORDER BY t.cout_total DESC;
```

`LATERAL` (PG) / `CROSS APPLY` (SQL Server) : sous-requête du FROM qui référence
les tables précédentes — les « 3 dernières interventions de chaque équipement » :

```sql
-- PostgreSQL
SELECT e.reference, r.id, r.date_ouverture
FROM equipements e
CROSS JOIN LATERAL (
    SELECT i.id, i.date_ouverture FROM interventions i
    WHERE i.equipement_id = e.id
    ORDER BY i.date_ouverture DESC
    LIMIT 3
) AS r;

-- MySQL 8+ : pas de LATERAL → utiliser ROW_NUMBER() (section 37)
```

---

## 34. CTE — WITH, la requête nommée

Une CTE (Common Table Expression) nomme une sous-requête réutilisable.
Lisibilité +++, et souvent le même plan que la sous-requête.

```sql
WITH
interv_par_tech AS (
    SELECT technicien_id, COUNT(*) AS nb, SUM(cout_eur) AS cout_total
    FROM interventions
    GROUP BY technicien_id
),
techs AS (
    SELECT id, nom, prenom FROM employes WHERE actif
)
SELECT t.nom, t.prenom, i.nb, i.cout_total
FROM techs t
LEFT JOIN interv_par_tech i ON i.technicien_id = t.id
ORDER BY i.nb DESC NULLS LAST;
```

Avantages sur les sous-requêtes imbriquées :

- Chaque étape a un **nom** qui documente l'intention.
- On peut chaîner / réutiliser (la même CTE référencée 2 fois).
- Le débogage se fait étape par étape (exécuter chaque CTE seule).

⚠️ MySQL < 8.0 et MariaDB < 10.2 : pas de CTE → tables dérivées (section 33).
Aujourd'hui, MySQL 8+ / MariaDB 10.2+ / PG : OK partout.

Note PG : jusqu'à la v11, les CTE étaient des « barrières d'optimisation »
(toujours matérialisées). Depuis PG 12, l'optimiseur peut les « aplatir ».
`MATERIALIZED` / `NOT MATERIALIZED` permettent de forcer le comportement.

---

## 35. CTE récursives — hiérarchies et séries

```sql
-- Chaîne hiérarchique : qui est sous les ordres de qui (niveau par niveau)
WITH RECURSIVE hierarchie AS (
    -- Ancre : le sommet
    SELECT id, nom, manager_id, 0 AS niveau
    FROM employes WHERE manager_id IS NULL
    UNION ALL
    -- Récursion : les subordonnés du niveau précédent
    SELECT e.id, e.nom, e.manager_id, h.niveau + 1
    FROM employes e
    JOIN hierarchie h ON e.manager_id = h.id
)
SELECT nom, niveau FROM hierarchie ORDER BY niveau, nom;
```

Deuxième usage roi : **générer des séries** (calendrier, sans table calendrier) :

```sql
-- Tous les jours de septembre 2026, avec le nombre d'interventions (0 inclus)
WITH RECURSIVE jours AS (
    SELECT DATE '2026-09-01' AS j
    UNION ALL
    SELECT j + INTERVAL '1 day' FROM jours WHERE j < DATE '2026-09-30'
)
SELECT j.j AS jour, COUNT(i.id) AS nb_interv
FROM jours j
LEFT JOIN interventions i ON i.date_ouverture::date = j.j
GROUP BY j.j
ORDER BY j.j;
-- MySQL : j + INTERVAL 1 DAY ; i : DATE(i.date_ouverture) = j.j
```

⚠️ Toujours une condition d'arrêt (`WHERE j < ...`), sinon récursion infinie
(PG coupe à 100 itérations par défaut ? non — PG ne limite pas, MySQL/MariaDB :
`cte_max_recursion_depth` = 1000 par défaut).

---

## 36. Fonctions fenêtre — OVER, le game changer

Une fonction fenêtre calcule sur un **ensemble de lignes liées** à la ligne courante,
**sans réduire** le nombre de lignes (contrairement à GROUP BY).

```sql
SELECT
    nom, service, salaire,
    AVG(salaire) OVER () AS moyenne_globale,              -- toute la table
    AVG(salaire) OVER (PARTITION BY service) AS moy_service, -- par groupe
    salaire - AVG(salaire) OVER (PARTITION BY service) AS ecart_service
FROM employes;
```

