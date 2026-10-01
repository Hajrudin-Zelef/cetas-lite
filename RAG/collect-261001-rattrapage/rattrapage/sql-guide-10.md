---
id: collect-261001-rattrapage/rattrapage/sql-guide-10
title: "Guide SQL complet — du SELECT à l'optimisation"
domain: rattrapage
role: reference
task: reference
actors: []
dates: ["2026-01-01", "2027-01-01"]
keywords: ["attention"]
source: docs/RAG/collect-261001-rattrapage/sql_guide.md
source_anchor: ""
source_lines: [1914, 2109]
sha256: 0710c8775b4945efccd71e260867538223ef59202b5fd2a8aac93e48b3cb7d4d
---

# Guide SQL complet — du SELECT à l'optimisation

Intérêts : masquer la complexité, **sécuriser** (exposer sans GRANT sur les tables de base),
garantir une définition unique (« le parc actif, c'est ÇA »).

Limites :

- Une vue simple est **réévaluée à chaque usage** (pas de stockage).
- Mise à jour via une vue : possible si elle mappe 1-1 à une table (sinon triggers
  `INSTEAD OF` sous PG).
- `CREATE OR REPLACE` ne peut pas changer les colonnes existantes de type incompatible →
  `DROP VIEW` + `CREATE` (attention aux dépendances).

---

## 57. Vues matérialisées — le cache requêtable

```sql
-- PostgreSQL : stocke le résultat, à rafraîchir explicitement
CREATE MATERIALIZED VIEW mv_couts_mensuels AS
SELECT DATE_TRUNC('month', date_ouverture) AS mois,
       type_interv, COUNT(*) AS nb, SUM(cout_eur) AS cout_total
FROM interventions
GROUP BY 1, 2;

CREATE UNIQUE INDEX ON mv_couts_mensuels (mois, type_interv);

-- Rafraîchir (chaque nuit via cron/pg_cron par ex.)
REFRESH MATERIALIZED VIEW CONCURRENTLY mv_couts_mensuels;
```

| | Vue simple | Vue matérialisée |
|---|---|---|
| Stockage | non | oui (table physique) |
| Fraîcheur | temps réel | au dernier REFRESH |
| Vitesse lecture | = requête sous-jacente | très rapide |
| Écriture | — | coût au REFRESH |

⚠️ **Portabilité** — MySQL/MariaDB : **pas** de vues matérialisées natives.
Émulation : table + `INSERT/REPLACE` périodique via EVENT, ou outil externe.
`CONCURRENTLY` (PG) : rafraîchit sans bloquer les lectures (exige un index unique).

Cas d'usage : tableaux de bord (Grafana/Metabase) sur des agrégats lourds,
requêtes répétées à l'identique.

---

## 58. Index — quand et pourquoi

Un index = une structure (souvent B-tree) qui évite de **scanner toute la table**.

```sql
-- Index sur les colonnes filtrées/jointes
CREATE INDEX idx_interv_equipement ON interventions (equipement_id);
CREATE INDEX idx_interv_date       ON interventions (date_ouverture DESC);
CREATE INDEX idx_equip_statut_type ON interventions (statut, type_equip);  -- composite

-- Index unique = contrainte + performance
CREATE UNIQUE INDEX uq_equip_reference ON equipements (reference);

-- Index partiel (PG) : n'indexer que ce qui sert
CREATE INDEX idx_interv_en_cours ON interventions (date_ouverture)
    WHERE date_cloture IS NULL;

-- Index d'expression (PG) : requêtes avec LOWER()
CREATE INDEX idx_equip_marque_lower ON equipements (LOWER(marque));
```

Quand indexer :

- ✅ Colonnes de `JOIN` (FK), de `WHERE` sélectif, de `ORDER BY` / `GROUP BY`.
- ✅ Contraintes UNIQUE/PK (automatique).
- ❌ Colonnes à faible cardinalité seules (`statut` avec 4 valeurs sur 10 M de lignes :
  l'index ne sert presque jamais — sauf partiel).
- ❌ Tables minuscules (< quelques milliers de lignes) : le scan séquentiel est plus rapide.
- ❌ Colonnes écrites intensivement et jamais lues par filtre : chaque INSERT/UPDATE
  paie le coût de maintenance de l'index.

Sur-indexer ralentit les écritures ; sous-indexer tue les lectures.
**Mesurer** avec EXPLAIN (section 59), pas au doigt mouillé.

⚠️ MySQL : les index sur expressions exigent MySQL 8.0.13+ (colonnes générées avant).
Les index partiels n'existent pas sous MySQL (index complet seulement).

---

## 59. Lire un EXPLAIN — diagnostiquer une requête lente

```sql
-- PostgreSQL
EXPLAIN (ANALYZE, BUFFERS, FORMAT TEXT)
SELECT e.reference, COUNT(i.id)
FROM equipements e
LEFT JOIN interventions i ON i.equipement_id = e.id
GROUP BY e.reference;
```

Ce qu'il faut y voir :

| Nœud | Signification | Bon / mauvais signe |
|---|---|---|
| `Seq Scan` | lit toute la table | OK si petite table ou gros % lu ; sinon suspect |
| `Index Scan` / `Index Only Scan` | utilise un index | bon signe |
| `Bitmap Heap Scan` | index + accès table | normal pour sélectivité moyenne |
| `Nested Loop` | boucle imbriquée | bien pour petits volumes ; catastrophique si l'intérieur est gros |
| `Hash Join` / `Merge Join` | jointures par hachage/tri | normal sur gros volumes |
| `Sort` | tri explicite | vérifier qu'un index ne l'éviterait pas |

```sql
-- MySQL
EXPLAIN FORMAT=TREE
SELECT e.reference, COUNT(i.id)
FROM equipements e
LEFT JOIN interventions i ON i.equipement_id = e.id
GROUP BY e.reference;
-- Colonnes clés : type (ALL = scan complet !), key (index utilisé), rows (estimation), Extra
```

Méthode express (5 minutes) :

1. `EXPLAIN` la requête → repérer `Seq Scan`/`ALL` sur de grosses tables.
2. Vérifier que les colonnes de JOIN/WHERE ont un index (section 58).
3. Regarder les **estimations vs réel** (`ANALYZE`) : un écart ×100 = statistiques
   obsolètes → `ANALYZE table;` (PG) / `ANALYZE TABLE` (MySQL).
4. Chercher les fonctions sur colonnes indexées (`WHERE LOWER(x) = ...` tue l'index
   simple → index d'expression ou normalisation à l'écriture).

---

## 60. Optimisation — les 10 réflexes avant de toucher au hardware

1. **Indexer les jointures** : toute FK filtrée/jointe mérite un index.
2. **Éviter `SELECT *`** : moins de colonnes = moins d'I/O, index couvrants possibles.
3. **Filtrer tôt** : WHERE sélectif avant JOIN quand c'est équivalent.
4. **Pas de fonction sur colonne indexée** dans WHERE : `WHERE YEAR(d) = 2026` →
   `WHERE d >= '2026-01-01' AND d < '2027-01-01'`.
5. **LIMIT dès que possible** sur les requêtes d'exploration.
6. **EXISTS > IN** pour les sous-requêtes corrélées/volumineuses (souvent équivalent
   pour l'optimiseur moderne, mais EXISTS documente l'intention).
7. **Éviter les `OR` sur plusieurs colonnes** : `WHERE a = 1 OR b = 2` casse l'indexage →
   réécrire en `UNION` de deux requêtes indexées.
8. **Pagination keyset** au lieu d'`OFFSET` profond (section 11).
9. **Agréger en base**, pas dans l'application : `GROUP BY` côté SQL, pas 10 000 lignes
   ramenées en Python pour compter.
10. **Statistiques à jour** : `ANALYZE` régulier (autovacuum le fait sous PG ;
    vérifier qu'il n'est pas désactivé).

```sql
-- ❌ Lent : fonction sur colonne indexée + OR
SELECT * FROM interventions
WHERE YEAR(date_ouverture) = 2026 OR priorite = 1;

-- ✅ Rapide : plage sargable + UNION
SELECT * FROM interventions
WHERE date_ouverture >= '2026-01-01' AND date_ouverture < '2027-01-01'
UNION
SELECT * FROM interventions
WHERE priorite = 1 AND NOT (date_ouverture >= '2026-01-01' AND date_ouverture < '2027-01-01');
```

---

## 61. Dépannage — « la requête rame » pas à pas

Checklist d'intervention (dans l'ordre) :

- [ ] **1. Reproduire** : la lenteur est-elle constante ? (`EXPLAIN ANALYZE` 2-3 fois)
- [ ] **2. Isoler** : quelle partie ? Découper la requête (CTE par CTE, enlever les JOIN un par un).
- [ ] **3. Plan** : `Seq Scan` sur grosse table ? Mauvaises estimations (rows estimés vs réels) ?
- [ ] **4. Index** : colonnes de JOIN/WHERE indexées ? Index **utilisé** (pas seulement existant) ?
- [ ] **5. Verrous** : une autre session bloque ? (PG : `pg_locks` + `pg_stat_activity` ; MySQL : `SHOW ENGINE INNODB STATUS`, `performance_schema`)
- [ ] **6. Ressources** : I/O disque saturé ? `work_mem` trop petit → tri sur disque (`Sort Method: external merge` dans EXPLAIN) ?
- [ ] **7. Statistiques** : `ANALYZE` récent ? (table venant de subir un gros import ?)

Requêtes de diagnostic :

```sql
-- PostgreSQL : requêtes en cours et leur durée
SELECT pid, now() - query_start AS duree, state, LEFT(query, 120)
FROM pg_stat_activity
WHERE state <> 'idle'
ORDER BY duree DESC;

-- PostgreSQL : qui bloque qui
SELECT blocked.pid AS pid_bloque, blocking.pid AS pid_bloquant,
       blocked.query AS requete_bloquee
FROM pg_locks blocked
JOIN pg_locks blocking ON blocking.locktype = blocked.locktype
    AND blocking.pid <> blocked.pid;

-- MySQL : processus en cours
SHOW FULL PROCESSLIST;
-- MySQL : transactions InnoDB qui attendent
SELECT * FROM information_schema.INNODB_TRX ORDER BY trx_started;
```

Tuer une requête runaway (avec discernement) :

