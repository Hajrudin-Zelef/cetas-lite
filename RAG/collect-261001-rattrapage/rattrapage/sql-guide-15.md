---
id: collect-261001-rattrapage/rattrapage/sql-guide-15
title: "Guide SQL complet — du SELECT à l'optimisation"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/sql_guide.md
source_anchor: ""
source_lines: [2815, 2986]
sha256: 11ca1d05990c908790e2baca21079a0e619425c21f2ac7fd3180f6afab107812
---

# Guide SQL complet — du SELECT à l'optimisation

**R8.** 1) **Requêtes paramétrées** (non négociable) ; 2) moindre privilège du compte
applicatif ; 3) ne pas exposer les erreurs SQL brutes (+ validation des entrées
en complément) (sections 62-64).

**R9.** Pagination **keyset** : mémoriser le dernier id/date de la page et faire
`WHERE id > :dernier_id ORDER BY id LIMIT 10` — pas de scan des 100 000 lignes
sautées (section 11).

**R10.** `mysqldump --single-transaction` (dump cohérent InnoDB sans verrou global)
avec `--routines --triggers` pour ne pas perdre procédures stockées et déclencheurs
(section 65).

**Score indicatif :** 8+/10 = autonome au quotidien ; 6-7 = relire les sections
13, 19, 21, 24, 31 ; moins = reprendre depuis la section 4 avec la base de test.

---

## 80. Pour aller plus loin — feuille de route

**Niveau 2 — administration légère (sysadmin) :**

- `pg_stat_statements` (PG) : quelles requêtes consomment le plus (temps total, appels).
  `CREATE EXTENSION pg_stat_statements;` puis top 10 par `total_exec_time`.
- `VACUUM` / `ANALYZE` : comprendre le bloat et l'autovacuum (ne jamais le désactiver
  sans raison mesurée).
- Réplication : primaire/standby (streaming replication PG, replica MySQL) pour la HA
  et pour déporter les gros reportings.
- Partitionnement : `PARTITION BY RANGE (date)` pour les tables d'historique/logs
  (purge par `DROP PARTITION` instantanée).
- `pg_cron` / MySQL EVENT scheduler : automatiser REFRESH des vues matérialisées,
  purges, rapports.

**Niveau 3 — SQL avancé :**

- JSON : `JSONB` (PG) — opérateurs `->`, `->>`, `@>`, index GIN. Stocker du
  semi-structuré (config d'équipements, réponses d'API) tout en restant requêtable.
- Full-text search : `tsvector`/`tsquery` (PG), `MATCH ... AGAINST` (MySQL) —
  mieux que `LIKE '%...%'` sur de gros textes.
- Triggers et procédures stockées (`PL/pgSQL`) : audit automatique
  (`qui a modifié quoi quand`), règles métier complexes.
- Fenêtres avancées : `ntile()`, `percent_rank()`, `cume_dist()`, cadres `RANGE`
  temporels, `IGNORE NULLS` (selon SGBD).
- Modélisation : normalisation jusqu'en 3NF, puis **dénormalisation raisonnée**
  pour la lecture (vues matérialisées plutôt que redondance manuelle).

**Ressources :**

- Documentation officielle PostgreSQL (la référence la plus pédagogique) :
  chapitres « Tutorial » puis « Performance Tips ».
- Use The Index, Luke (use-the-index-luke.com) — le site de référence sur l'indexation.
- « SQL Antipatterns » (Bill Karwin) — les pièges à éviter, avec solutions.
- Exercices : SQLZoo, PGExercises (PostgreSQL), LeetCode Database — 15 min/jour.

**Projet fil rouge conseillé :** modéliser VOTRE parc réel (onduleurs, copieurs,
interventions) dans la base de test, y importer un export CSV, et réécrire les
tableaux de bord du service en SQL. C'est là que le SQL devient un réflexe.

---

## 81. Checklist quotidienne du sysadmin SQL

**Le matin (5 minutes) :**

- [ ] Espace disque des volumes de données et de dumps (`df -h`, taille des derniers dumps).
- [ ] Dernier dump : présent ? taille cohérente ? heure correcte ?
- [ ] Requêtes longues en cours (`pg_stat_activity` / `SHOW FULL PROCESSLIST`).
- [ ] Alertes applicatives liées à la base (connexions refusées, erreurs).

**En intervention :**

- [ ] Toujours sur la **bonne base** (`SELECT current_database();` / `SELECT DATABASE();`)
      — le nombre d'accidents « je croyais être sur la préprod » est légendaire.
- [ ] `BEGIN;` avant tout DML manuel ; `SELECT` de contrôle avant `COMMIT`.
- [ ] Dump ciblé (`pg_dump -t table`) avant ALTER sur table métier.
- [ ] Noter la requête exécutée dans le ticket (traçabilité).

**Le vendredi (15 minutes) :**

- [ ] Test de restauration du dump hebdo sur base jetable (chronométré).
- [ ] `ANALYZE` / vérification autovacuum sur les tables à fort churn.
- [ ] Revue des droits : comptes inactifs à révoquer, mots de passe à rotation.
- [ ] Taille des tables qui grossissent vite (logs, historique) — anticiper le partitionnement.

---

## 82. Modèles de requêtes prêts à copier — exploitation

```sql
-- Top 10 des tables les plus volumineuses (PostgreSQL)
SELECT schemaname || '.' || tablename AS table,
       pg_size_pretty(pg_total_relation_size(schemaname || '.' || tablename)) AS taille
FROM pg_tables
WHERE schemaname NOT IN ('pg_catalog', 'information_schema')
ORDER BY pg_total_relation_size(schemaname || '.' || tablename) DESC
LIMIT 10;

-- MySQL : tailles des tables
SELECT table_name AS tbl,
       ROUND((data_length + index_length) / 1024 / 1024, 1) AS mo
FROM information_schema.tables
WHERE table_schema = DATABASE()
ORDER BY (data_length + index_length) DESC
LIMIT 10;

-- Dernière analyse / vacuum par table (PostgreSQL) — repérer les tables oubliées
SELECT relname AS table, last_vacuum, last_autovacuum, last_analyze, last_autoanalyze,
       n_dead_tup AS lignes_mortes
FROM pg_stat_user_tables
ORDER BY n_dead_tup DESC NULLS LAST
LIMIT 15;

-- Interventions ouvertes triées par ancienneté, avec âge en jours (portable)
SELECT id, equipement_id, description, date_ouverture,
       EXTRACT(DAY FROM (NOW() - date_ouverture)) AS age_jours_pg
       -- MySQL : DATEDIFF(NOW(), date_ouverture) AS age_jours
FROM interventions
WHERE date_cloture IS NULL
ORDER BY date_ouverture;

-- Export CSV propre depuis psql (évite les copier-coller hasardeux)
-- \copy (SELECT * FROM v_parc_actif) TO '/tmp/parc.csv' WITH (FORMAT csv, HEADER)
-- MySQL : SELECT ... INTO OUTFILE '/tmp/parc.csv' FIELDS TERMINATED BY ';' ENCLOSED BY '"';
```

---

## 83. Anti-sèches des différences MySQL ↔ PostgreSQL

| Sujet | PostgreSQL | MySQL / MariaDB |
|---|---|---|
| Pagination | `LIMIT n OFFSET m` | `LIMIT m, n` ou `LIMIT n OFFSET m` |
| Concaténation | `'a' \|\| 'b'` | `CONCAT('a','b')` |
| Insensible à la casse | `ILIKE` | collation `*_ci` par défaut |
| Auto-incrément | `SERIAL` / `IDENTITY` | `AUTO_INCREMENT` |
| UPSERT | `ON CONFLICT ... DO UPDATE` | `ON DUPLICATE KEY UPDATE` |
| Retour d'INSERT | `RETURNING` | `LAST_INSERT_ID()` |
| CTE récursive | `WITH RECURSIVE` | `WITH RECURSIVE` (8.0+/10.2+) |
| Fenêtres | complet | complet (8.0+/10.2+) |
| `FILTER (WHERE)` | ✅ | ❌ → `SUM(CASE...)` |
| FULL OUTER JOIN | ✅ | ❌ → UNION de LEFT+RIGHT |
| INTERSECT/EXCEPT | ✅ | ❌ → EXISTS/NOT EXISTS |
| Vues matérialisées | ✅ | ❌ → table + EVENT |
| Index partiels | ✅ (`WHERE ...`) | ❌ |
| `NULLS FIRST/LAST` | ✅ | ❌ → `ORDER BY col IS NULL, col` |
| Intervalle | `INTERVAL '7 days'` | `INTERVAL 7 DAY` |
| Différence de dates | `d2 - d1` (interval) | `DATEDIFF`/`TIMESTAMPDIFF` |
| Aléatoire | `RANDOM()` | `RAND()` |
| Tronquer nombre | `TRUNC(x, n)` | `TRUNCATE(x, n)` |
| Division `/` | entière si 2 entiers | toujours décimale |
| `BOOLEAN` | vrai type | alias de `TINYINT(1)` |
| TRUNCATE | transactionnel | commit implicite |
| DDL | transactionnel | commit implicite |
| FK | toujours vérifiées | seulement en InnoDB |

En cas de doute sur une fonction : la tester sur les deux avec la base de démo
vaut mieux qu'un pari en production.

---

## 84. Conclusion — les 10 commandements du SQL d'exploitation

1. **Tu testeras sur une copie** avant de toucher à la production.
2. **Tu listeras tes colonnes** : `SELECT *` est réservé à l'exploration.
3. **Tu respecteras NULL** : ni égal, ni différent — il s'apprivoise avec IS, COALESCE, EXISTS.
4. **Tu parenthèseras tes OR** : AND d'abord, le reste ensuite.
5. **Tu mettras tes conditions d'OUTER JOIN dans le ON**, pas dans le WHERE.
6. **Tu paramètreras tes requêtes** : aucune concaténation d'entrée utilisateur.
7. **Tu encadreras tes DML** : BEGIN, contrôle, puis COMMIT ou ROLLBACK.
8. **Tu sauvegarderas avant de migrer**, et tu testeras tes restaurations.
9. **Tu EXPLAINeras avant d'indexer**, et tu ANALYzeras après un gros import.
10. **Tu documenteras tes scripts** : objet, auteur, date, ticket, rollback.

