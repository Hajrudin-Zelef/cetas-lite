---
id: collect-261001-rattrapage/rattrapage/sql-guide-5
title: "Guide SQL complet — du SELECT à l'optimisation"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/sql_guide.md
source_anchor: ""
source_lines: [797, 1025]
sha256: 80a2d10371577d7a5d3d0b5ccd09d7acbe0ca5ad8eff2f596a6faf64fbb1c109
---

# Guide SQL complet — du SELECT à l'optimisation

```sql
-- ❌ Invalide en standard (et refusé par PG) : poste n'est ni agrégé ni groupé
SELECT service, poste, COUNT(*) FROM employes GROUP BY service;

-- ✅ Correct
SELECT service, poste, COUNT(*) FROM employes GROUP BY service, poste;

-- ✅ Autorisé par PG : id est la PK, nom/prenom en dépendent fonctionnellement
SELECT e.id, e.nom, e.prenom, COUNT(i.id) AS nb_interv
FROM employes e LEFT JOIN interventions i ON i.technicien_id = e.id
GROUP BY e.id;
```

⚠️ MySQL avec `ONLY_FULL_GROUP_BY` désactivé (ancien défaut !) **accepte** la forme invalide
et renvoie une valeur **arbitraire** pour la colonne manquante → résultats faux silencieux.
Vérifier : `SELECT @@sql_mode;` doit contenir `ONLY_FULL_GROUP_BY`.

---

## 21. HAVING — filtrer APRÈS agrégation

`WHERE` filtre les lignes **avant** regroupement ; `HAVING` filtre les **groupes après** calcul.

```sql
-- Techniciens ayant coûté plus de 500 € au total
SELECT technicien_id, SUM(cout_eur) AS total
FROM interventions
GROUP BY technicien_id
HAVING SUM(cout_eur) > 500;

-- Ordre logique complet (mémoriser cet ordre !) :
-- FROM → WHERE → GROUP BY → HAVING → SELECT → ORDER BY → LIMIT
SELECT type_equip, COUNT(*) AS nb
FROM equipements
WHERE statut <> 'reforme'          -- 1. filtre lignes
GROUP BY type_equip                -- 2. regroupe
HAVING COUNT(*) >= 2               -- 3. filtre groupes
ORDER BY nb DESC                   -- 4. trie
LIMIT 5;                           -- 5. pagine
```

On peut réutiliser l'alias dans HAVING sous MySQL (`HAVING total > 500`) ;
PostgreSQL l'interdit (utiliser l'expression ou un sous-SELECT/CTE). Pour du portable :
répéter l'expression ou passer par une CTE (section 34).

---

## 22. Regroupements avancés : ROLLUP, CUBE, GROUPING SETS

Pour des sous-totaux en une seule requête (reporting) :

```sql
-- Sous-totaux par type + total général
SELECT type_interv, priorite, COUNT(*) AS nb, SUM(cout_eur) AS cout
FROM interventions
GROUP BY ROLLUP (type_interv, priorite)
ORDER BY type_interv NULLS LAST, priorite NULLS LAST;
-- Lignes où type_interv IS NULL = sous-total / total (utiliser GROUPING() pour distinguer)
```

| Construction | Produit |
|---|---|
| `ROLLUP(a, b)` | (a,b), (a), () — hiérarchique |
| `CUBE(a, b)` | (a,b), (a), (b), () — toutes combinaisons |
| `GROUPING SETS ((a,b),(a),())` | sur mesure |

```sql
-- Distinguer un vrai NULL d'un NULL de sous-total
SELECT type_interv,
       GROUPING(type_interv) AS est_sous_total,
       COUNT(*)
FROM interventions
GROUP BY ROLLUP (type_interv);
```

⚠️ **Portabilité** — supporté par PostgreSQL et MySQL 8+. MariaDB : `WITH ROLLUP`
(ancienne syntaxe MySQL, toujours valide : `GROUP BY type_interv WITH ROLLUP`).

---

## 23. INNER JOIN — l'intersection

`INNER JOIN` ne garde que les lignes qui **matchent des deux côtés**.

```sql
-- Interventions avec le nom du technicien et la référence de l'équipement
SELECT i.id, i.date_ouverture, e.reference, e.modele,
       emp.nom AS technicien, i.description
FROM interventions i
INNER JOIN equipements e ON e.id = i.equipement_id
INNER JOIN employes emp  ON emp.id = i.technicien_id
ORDER BY i.date_ouverture DESC;
```

Bonnes pratiques :

- **Toujours** aliaser (`i`, `e`, `emp`) et **préfixer** chaque colonne (`i.id`, pas `id`)
  dès qu'il y a 2+ tables — sinon `Column 'id' in field list is ambiguous`.
- `INNER` est optionnel : `JOIN` seul = `INNER JOIN`. L'écrire explicite aide la relecture.
- La condition dans `ON`, pas dans `WHERE` (pour INNER ça ne change rien au résultat,
  mais pour OUTER ça change tout — section 24).

---

## 24. LEFT JOIN — garder la gauche même sans match

`LEFT JOIN` garde **toutes** les lignes de la table de gauche ; les colonnes de droite
valent NULL quand rien ne matche. LE join le plus utile au quotidien.

```sql
-- Tous les employés + leur nombre d'interventions (même 0)
SELECT e.nom, e.prenom, COUNT(i.id) AS nb_interv
FROM employes e
LEFT JOIN interventions i ON i.technicien_id = e.id
GROUP BY e.id, e.nom, e.prenom
ORDER BY nb_interv DESC;
-- ⚠️ COUNT(i.id) et pas COUNT(*) : COUNT(*) compterait 1 pour les sans-intervention !

-- Équipements SANS aucune intervention (anti-join)
SELECT e.reference, e.modele
FROM equipements e
LEFT JOIN interventions i ON i.equipement_id = e.id
WHERE i.id IS NULL;
```

Le piège du WHERE sur la table de droite :

```sql
-- ❌ Le WHERE transforme le LEFT JOIN en INNER JOIN !
SELECT e.nom, i.description
FROM employes e
LEFT JOIN interventions i ON i.technicien_id = e.id
WHERE i.priorite = 1;   -- élimine les lignes où i.* est NULL → INNER déguisé

-- ✅ Condition dans le ON : on garde les employés sans intervention prio 1
SELECT e.nom, i.description
FROM employes e
LEFT JOIN interventions i ON i.technicien_id = e.id AND i.priorite = 1;
```

**Règle : condition sur la table optionnelle → dans le `ON`, jamais dans le `WHERE`.**

---

## 25. RIGHT JOIN et FULL OUTER JOIN

```sql
-- RIGHT JOIN : garde toute la table de DROITE
-- (équivalent à un LEFT JOIN en inversant les tables — rarement utilisé)
SELECT e.nom, i.description
FROM interventions i
RIGHT JOIN employes e ON e.id = i.technicien_id;
-- = tous les employés, même sans intervention

-- FULL OUTER JOIN : garde TOUT des deux côtés
SELECT e.reference AS equip, i.id AS interv
FROM equipements e
FULL OUTER JOIN interventions i ON i.equipement_id = e.id;
-- Équipements sans intervention (i.id NULL) + interventions orphelines (e.reference NULL)
```

⚠️ **Portabilité** — MySQL/MariaDB **ne supportent pas** `FULL OUTER JOIN`.
Émulation portable :

```sql
SELECT e.reference, i.id
FROM equipements e LEFT JOIN interventions i ON i.equipement_id = e.id
UNION
SELECT e.reference, i.id
FROM equipements e RIGHT JOIN interventions i ON i.equipement_id = e.id;
-- (UNION déduplique ; les lignes matchées apparaissent une fois)
```

En pratique : on écrit presque toujours `LEFT JOIN` (en choisissant l'ordre des tables),
`RIGHT JOIN` ne sert qu'en lecture de code existant.

---

## 26. CROSS JOIN — produit cartésien

```sql
-- Toutes les combinaisons technicien × type d'équipement (ex. matrice d'habilitation)
SELECT e.nom, t.type_equip
FROM (SELECT DISTINCT nom FROM employes) e
CROSS JOIN (SELECT DISTINCT type_equip FROM equipements) t;
```

Cas d'usage légitimes : matrices, calendriers (jours × équipes), jeux de test.
Danger : `n × m` lignes — 10 000 × 10 000 = 100 M de lignes. Un `JOIN` sans `ON`
(par oubli !) est un CROSS JOIN implicite : **toujours vérifier la condition de jointure**.

```sql
-- ❌ Oubli du ON : produit cartésien silencieux
SELECT * FROM interventions, equipements;
-- ✅ Explicite
SELECT * FROM interventions i JOIN equipements e ON e.id = i.equipement_id;
```

---

## 27. Self-join — une table jointe à elle-même

Indispensable pour les hiérarchies (manager/employé) et les comparaisons intra-table.

```sql
-- Employé + nom de son manager
SELECT e.nom AS employe, m.nom AS manager
FROM employes e
LEFT JOIN employes m ON m.id = e.manager_id;   -- LEFT : le chef n'a pas de manager

-- Équipements achetés la même année qu'un autre (paires)
SELECT a.reference AS eq1, b.reference AS eq2,
       EXTRACT(YEAR FROM a.date_achat) AS annee
FROM equipements a
JOIN equipements b ON EXTRACT(YEAR FROM a.date_achat) = EXTRACT(YEAR FROM b.date_achat)
                  AND a.id < b.id;   -- < évite les doublons (a,b)/(b,a) et l'auto-paire
```

Autre usage : détecter des anomalies (« deux interventions ouvertes le même jour
sur le même équipement ») — voir section 63 sur les doublons.

---

## 28. USING et NATURAL JOIN — à connaître, à éviter

```sql
-- USING : quand la colonne de jointure a le même nom des deux côtés
SELECT * FROM interventions JOIN equipements USING (id);  -- ⚠️ ici id ≠ id ! DANGER

