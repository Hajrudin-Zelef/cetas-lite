---
id: collect-261001-rattrapage/rattrapage/postgresql-guide-6
title: "PostgreSQL en production — Guide complet pour sysadmin"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["cost"]
source: docs/RAG/collect-261001-rattrapage/postgresql_guide.md
source_anchor: ""
source_lines: [1156, 1379]
sha256: edb12ad3d0b2d76a9d742d4ff45e486d4401e62c69e865d4ccd564ac65cd39cd
---

# Import/export côté client (droits du user local, recommandé)
psql -h srv-bdd -U app_ecrivain -d inventaire \
  -c "\copy capteurs(code, site_id) FROM 'capteurs.csv' WITH (FORMAT csv, HEADER)"

# Gros import : désactiver temporairement les index/triggers est rarement
# utile ; préférez COPY (déjà très rapide) + maintenance_work_mem élevée
```

**Options `WITH` utiles :** `FORMAT csv`, `HEADER`, `DELIMITER ';'`,
`QUOTE '"'`, `ENCODING 'UTF8'`, `NULL '\N'`.

---

# Partie D — Performance : index, EXPLAIN, VACUUM, transactions

## 36. Index B-tree : le défaut qui couvre 90 % des cas

```sql
-- Index simple sur une colonne de filtrage/jointure
CREATE INDEX idx_mesures_capteur ON mesures (capteur_id);

-- Index composite : ordre = ordre des filtres (égalité d'abord)
CREATE INDEX idx_mesures_capteur_date ON mesures (capteur_id, horodatage DESC);

-- Index unique (contrainte + accès)
CREATE UNIQUE INDEX uq_capteurs_code ON capteurs (code);

-- Création sans bloquer les écritures (prod !)
CREATE INDEX CONCURRENTLY idx_mesures_site ON mesures (site_id);
```

**Règles :**

- Indexez les **clés étrangères** (sinon chaque DELETE/UPDATE du parent
  fait un seq scan de l'enfant).
- `CREATE INDEX CONCURRENTLY` en production : pas de verrou exclusif,
  mais 2× plus long et non utilisable dans un bloc transactionnel.
- Un index inutile coûte : écritures ralenties + espace + maintenance.

```sql
-- Trouver les index jamais utilisés (candidats à la suppression)
SELECT schemaname, relname, indexrelname, idx_scan
FROM pg_stat_user_indexes
WHERE idx_scan = 0 AND indexrelname NOT LIKE '%pkey%'
ORDER BY pg_relation_size(indexrelid) DESC;
```

---

## 37. Index partiels, d'expression et couvrants

```sql
-- Index PARTIEL : seulement les lignes qui nous intéressent (petit, rapide)
CREATE INDEX idx_alertes_actives ON supervision.alertes (site_id)
  WHERE resolue = false;

-- Index D'EXPRESSION : recherche insensible à la casse, fonctions
CREATE INDEX idx_capteurs_code_lower ON capteurs (lower(code));

-- Index COUVRANT (INCLUDE) : index-only scan, sans toucher la table
CREATE INDEX idx_mesures_couvrant
  ON mesures (capteur_id, horodatage DESC)
  INCLUDE (valeur);

-- Vérifier qu'une requête utilise l'index partiel : le WHERE doit
-- impliquer la même condition (resolue = false)
EXPLAIN SELECT * FROM supervision.alertes
WHERE resolue = false AND site_id = 3;
```

---

## 38. GIN, GiST, BRIN : les index spécialisés

| Type | Usage | Exemple |
|---|---|---|
| B-tree | Égalités, plages, tri (défaut) | `WHERE id = 42`, `ORDER BY date` |
| GIN | JSONB, tableaux, recherche plein texte | `config @> '{...}'`, `to_tsvector` |
| GiST | Géométrie, plages, `inet` | PostGIS, `ip << '10.0.0.0/8'` |
| BRIN | **Grosses tables** ordonnées physiquement (séries temporelles) | `horodatage` sur des milliards de lignes |
| Hash | Égalités simples (rarement mieux que B-tree) | Cas spécifiques |

```sql
-- GIN trigramme : recherche floue (LIKE '%...%') rapide
CREATE EXTENSION IF NOT EXISTS pg_trgm;
CREATE INDEX idx_sites_nom_trgm ON sites USING gin (nom gin_trgm_ops);
SELECT * FROM sites WHERE nom ILIKE '%atel%';

-- BRIN : minuscule, parfait pour les mesures horodatées insérées en ordre
CREATE INDEX idx_mesures_horodatage_brin ON mesures USING brin (horodatage);

-- Plein texte : tsvector + GIN
ALTER TABLE interventions ADD COLUMN doc_tsv tsvector
  GENERATED ALWAYS AS (to_tsvector('french', description)) STORED;
CREATE INDEX idx_interventions_fts ON interventions USING gin (doc_tsv);
SELECT * FROM interventions
WHERE doc_tsv @@ plainto_tsquery('french', 'onduleur batterie');
```

> ⚠️ `pg_trgm` et le FTS : pensez à la **configuration de langue**
> (`'french'`) pour la pertinence en français.

---

## 39. EXPLAIN : lire un plan d'exécution

```sql
EXPLAIN
SELECT * FROM mesures WHERE capteur_id = 42 ORDER BY horodatage DESC LIMIT 100;
```

```
Limit  (cost=8.45..8.47 rows=100 ...)
  ->  Sort  (cost=8.45..8.70 ...)
        Sort Key: horodatage DESC
        ->  Index Scan using idx_mesures_capteur_date on mesures
              Index Cond: (capteur_id = 42)
```

**Vocabulaire du plan :**

| Nœud | Signification |
|---|---|
| `Seq Scan` | Balayage complet : OK sur petite table, suspect sur grosse |
| `Index Scan` | Parcours d'index + accès table |
| `Index Only Scan` | Tout dans l'index (idéal, cf. `INCLUDE`) |
| `Bitmap Heap Scan` | Combine plusieurs index, puis accès table triés |
| `Nested Loop` | Boucle imbriquée : bien si la table interne est petite/indexée |
| `Hash Join` / `Merge Join` | Jointures sur gros volumes |
| `Sort` | Tri explicite : coûteux sans index |

**`cost=`** : unité arbitraire (coût d'une page disque = 1).
**`rows=`** : estimation — si elle est fausse d'un facteur 10+,
lancez `ANALYZE` (statistiques périmées, section 43).

---

## 40. EXPLAIN ANALYZE : le plan réel, avec temps

```sql
EXPLAIN (ANALYZE, BUFFERS, TIMING OFF)
SELECT s.nom, avg(m.valeur)
FROM mesures m JOIN capteurs c ON c.id = m.capteur_id
JOIN sites s ON s.id = c.site_id
WHERE m.horodatage > now() - INTERVAL '1 day'
GROUP BY s.nom;
```

**Options :**

| Option | Apport |
|---|---|
| `ANALYZE` | Exécute vraiment la requête (⚠️ `INSERT/UPDATE/DELETE` modifient !) |
| `BUFFERS` | Pages lues : `shared hit` (cache) vs `read` (disque) |
| `TIMING OFF` | Moins de surcharge de mesure sur requêtes rapides |
| `VERBOSE` | Détail des colonnes/filtres |

**Lecture :**

- `actual time=... rows=... loops=...` : temps réel par boucle.
- Écart `rows` estimé vs réel → `ANALYZE` ou statistiques ciblées.
- `Buffers: shared hit=... read=...` : beaucoup de `read` = disque
  sollicité → `shared_buffers` / index / `effective_cache_size`.
- Un `Seq Scan` avec `Rows Removed by Filter: 999000` crie "index manquant".

> ⚠️ `EXPLAIN ANALYZE` **exécute** la requête : pour un `DELETE`,
> encapsulez dans une transaction avec `ROLLBACK`.

```sql
BEGIN;
EXPLAIN (ANALYZE) DELETE FROM _tmp_import WHERE importe_le < now() - INTERVAL '1 an';
ROLLBACK;
```

---

## 41. Optimisation : la méthode en 7 étapes

1. **Mesurer** : `log_min_duration_statement = 1000` + `pg_stat_statements`
   (section 61) pour trouver les coupables, pas les suspects.
2. **EXPLAIN ANALYZE** sur la requête réelle, avec les vrais paramètres.
3. **Vérifier les estimations** : `rows` estimé vs réel → `ANALYZE`,
   voire `ALTER TABLE ... ALTER COLUMN ... SET STATISTICS 1000`.
4. **Index manquant ?** `Seq Scan` + gros `Rows Removed by Filter`.
5. **Index inutilement large ?** `Index Scan` qui lit 50 % de la table →
   filtre plus sélectif ou partitionnement.
6. **Réécrire** : `NOT IN` → `NOT EXISTS`, CTE matérialisée ou non,
   `OR` → `UNION`, `DISTINCT ON` vs `GROUP BY`.
7. **Re-mesurer** : le gain doit être visible sur `EXPLAIN ANALYZE`
   **et** sur la charge réelle (`pg_stat_statements`).

**Anti-patterns fréquents :**

```sql
-- ❌ Fonction sur colonne indexée : tue l'index
WHERE date_trunc('day', horodatage) = CURRENT_DATE
-- ✅ Plage : index utilisable
WHERE horodatage >= date_trunc('day', now())
  AND horodatage <  date_trunc('day', now()) + INTERVAL '1 day';

-- ❌ SELECT * puis tri applicatif
-- ✅ Ne sélectionner que les colonnes utiles + LIMIT + ORDER BY indexé

-- ❌ N+1 : une requête par ligne dans l'appli
-- ✅ Une seule requête avec JOIN / IN
```

---

## 42. Partitionnement : quand et comment

**Quand :** tables > ~10–50 Go, données à durée de vie (rétention),
requêtes filtrant sur la clé de partition (souvent le temps).

```sql
-- Table partitionnée par mois sur horodatage (déclaratif, PG 10+)
CREATE TABLE mesures (
  id          bigint GENERATED ALWAYS AS IDENTITY,
  capteur_id  bigint NOT NULL,
  horodatage  timestamptz NOT NULL,
  valeur      double precision NOT NULL,
  PRIMARY KEY (id, horodatage)          -- la clé de partition fait partie de la PK
) PARTITION BY RANGE (horodatage);

