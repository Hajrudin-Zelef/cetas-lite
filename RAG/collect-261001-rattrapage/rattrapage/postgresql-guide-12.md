---
id: collect-261001-rattrapage/rattrapage/postgresql-guide-12
title: "PostgreSQL en production — Guide complet pour sysadmin"
domain: rattrapage
role: reference
task: reference
actors: []
dates: ["2026-09-01", "2026-09-26"]
keywords: ["incident", "pruning"]
source: docs/RAG/collect-261001-rattrapage/postgresql_guide.md
source_anchor: ""
source_lines: [2471, 2683]
sha256: 184d2f8a654f144ab153d3ca051371ca9e3ca1b1d02c55d8029f8b4e2fa62f49
---

# PostgreSQL en production — Guide complet pour sysadmin

**Contexte :** application de supervision énergie : 500 capteurs,
1 mesure/capteur/minute, rétention 2 ans, 50 utilisateurs, pics
de lecture le matin.

**1. Volume :**

```
500 capteurs × 1440 mesures/jour = 720 000 lignes/jour
≈ 263 M lignes/an → ~530 M lignes / 2 ans
Ligne ≈ 40 octets + index ≈ 60 octets → ~50–60 Go pour 2 ans
```

**2. Partitionnement** (section 42) : table `mesures` partitionnée par
mois → purge par `DETACH` + requêtes filtrées par temps (pruning).

**3. Serveur :**

| Ressource | Choix | Justification |
|---|---|---|
| RAM | 32 Go | `shared_buffers` 8 Go, cache OS pour les lectures du matin |
| CPU | 8 vCPU | Pics d'agrégation ; parallélisme (`max_parallel_workers_per_gather = 4`) |
| Disque données | 500 Go SSD | 60 Go × marge de croissance 3 ans + index + marge 20 % |
| Disque WAL | 100 Go dédiés | `max_wal_size` 8 Go + archivage |
| Réseau | 10 Gb/s si réplication + sauvegardes | `pg_basebackup` sans impacter la prod |

**4. Paramètres clés :** `shared_buffers=8GB`, `effective_cache_size=24GB`,
`work_mem=32MB`, `maintenance_work_mem=1GB`, `checkpoint_completion_target=0.9`,
`autovacuum` agressif sur `mesures` (scale_factor 0.02).

**5. Exploitation :** standby streaming (reporting + bascule),
Barman (base hebdo + WAL continu, RPO ~ minutes), dump logique
quotidien, `pg_stat_statements` actif, alertes section 65.

---

## 78. Cas pratique 2 : diagnostiquer une requête lente (pas à pas)

**Ticket :** « Le dashboard énergie met 12 s à charger le matin. »

```sql
-- Étape 1 : identifier la requête (pg_stat_statements)
SELECT query, calls, round(mean_exec_time::numeric) AS moy_ms,
       round(total_exec_time::numeric) AS total_ms
FROM pg_stat_statements
ORDER BY total_exec_time DESC LIMIT 5;
-- Coupable : agrégation journalière sur mesures sans filtre de partition

-- Étape 2 : EXPLAIN ANALYZE avec les vrais paramètres
EXPLAIN (ANALYZE, BUFFERS)
SELECT date_trunc('day', horodatage) AS jour, avg(valeur)
FROM mesures
WHERE capteur_id IN (SELECT id FROM capteurs WHERE site_id = 5)
  AND horodatage >= '2026-09-01' AND horodatage < '2026-09-26'
GROUP BY 1;
-- Constat : Seq Scan sur 40 M de lignes, "Rows Removed by Filter: 39900000"
-- + Sort sur disque (temp files)

-- Étape 3 : corriger
-- a) Index manquant sur (site_id) de capteurs ? déjà là.
-- b) La requête ne profite pas du partitionnement : le filtre est sur
--    horodatage -> OK, mais l'index composite manque :
CREATE INDEX CONCURRENTLY idx_mesures_capteur_date
  ON mesures (capteur_id, horodatage DESC);

-- c) Réécrire : éviter la fonction sur la colonne partitionnée
--    (ici date_trunc est dans le SELECT, pas le WHERE : OK)

-- Étape 4 : re-mesurer
-- EXPLAIN ANALYZE -> Bitmap Heap Scan + Index, 12 s -> 0,4 s

-- Étape 5 : pérenniser
-- - ANALYZE après le CREATE INDEX (fait automatiquement par CONCURRENTLY)
-- - Ajouter la requête au dashboard de supervision (pg_stat_statements)
-- - Documenter dans le journal d'exploitation
```

**Leçon :** 80 % des « requêtes lentes » = index manquant ou
statistiques périmées. Les 20 % restants = réécriture ou
partitionnement.

---

## 79. Cas pratique 3 : restaurer à un instant T (scénario réel)

**Incident :** à 14h32, un script supprime par erreur les interventions
du mois (`DELETE FROM interventions WHERE ...` sans garde-fou).

```bash
# 1. STOP : empêcher d'autres écritures qui écraseraient le WAL utile
#    (si possible : passer l'appli en lecture seule)

# 2. Point de restauration : juste avant 14h32
#    Sur le serveur de restauration isolé :
sudo pg_ctlcluster 17 main stop
rm -rf /var/lib/postgresql/17/main/*
barman recover --target-time "2026-09-26 14:31:00+02" \
  srv-bdd latest /var/lib/postgresql/17/main/
sudo pg_ctlcluster 17 main start

# 3. Vérifier : compter les interventions du mois
sudo -u postgres psql -p 5433 -d inventaire -c \
  "SELECT count(*) FROM interventions WHERE debut >= '2026-09-01';"
# -> le chiffre attendu d'avant l'incident : OK

# 4. Réintégrer SEULEMENT les lignes manquantes (pas tout écraser !)
#    Depuis le cluster restauré (port 5433) vers la prod :
pg_dump -p 5433 -d inventaire -t interventions \
  --data-only --inserts | psql -h srv-bdd -d inventaire
# Mieux : dump dans une table _recup puis INSERT ... SELECT anti-doublons :
psql -h srv-bdd -d inventaire <<'EOF'
CREATE TEMP TABLE _recup (LIKE interventions INCLUDING ALL);
\copy _recup FROM 'interventions_recup.csv' WITH (FORMAT csv, HEADER)
INSERT INTO interventions SELECT * FROM _recup r
  WHERE NOT EXISTS (SELECT 1 FROM interventions i WHERE i.id = r.id);
EOF

# 5. Post-mortem : pourquoi le script a-t-il pu faire ça ?
#    -> droits trop larges ? pas de transaction ? pas de backup avant batch ?
#    Actions : rôle batch restreint, REVOKE DELETE direct, dry-run obligatoire
```

---

## 80. Cas pratique 4 : mise en place d'un standby de reporting

**Besoin :** soulager le primaire des requêtes BI du matin.

```bash
# 1. Primaire : slot + pg_hba (sections 50/53 déjà faits)
# 2. Cloner
sudo -u postgres pg_basebackup -h srv-primaire -U replicateur \
  -D /var/lib/postgresql/17/main -P -R --slot=standby_bi
# 3. postgresql.conf du standby : autoriser les requêtes longues
#    max_standby_streaming_delay = 30min   # requêtes BI longues
#    hot_standby_feedback = on             # évite les annulations (risque bloat primaire)
# 4. Démarrer, vérifier le lag (section 53)
# 5. Chaîne de connexion BI -> srv-standby (lecture seule)
# 6. Alerte lag > 10 min (section 65)
```

**Pièges :** le standby n'accepte **aucune** écriture (même les tables
temporaires → utilisez `temp_tablespaces` ou faites-les sur le primaire) ;
les séquences y sont « en retard » (normal : `nextval` non répliqué finement).

---

## 81. Pense-bête de poche

### Connexion

```bash
sudo -u postgres psql
psql -h srv -U user -d base -c "SELECT 1;"
PGPASSWORD='...' psql -h srv -U user -d base   # scripts (ou .pgpass)
```

### psql express

```
\l \du \dn \dt+ \di \d table \x \timing \copy \q
```

### SQL express

```sql
SELECT version();  SHOW config_file;  SHOW data_directory;
SELECT pg_reload_conf();                       -- reload conf
SELECT pg_terminate_backend(pid);              -- tuer une session
CREATE EXTENSION pg_stat_statements;
EXPLAIN (ANALYZE, BUFFERS) SELECT ...;
```

### Sauvegarde / restauration express

```bash
pg_dump -Fc -f base-$(date +%F).dump base
pg_dumpall --globals-only -f globals.sql
pg_basebackup -D /srv/bkp/base -Ft -z -P --wal-method=stream
pg_restore -j 4 -d base base-2026-09-26.dump
barman backup srv-bdd && barman list-backup srv-bdd
```

### Supervision express

```sql
SELECT * FROM pg_stat_archiver;                -- WAL OK ?
SELECT client_addr, state FROM pg_stat_replication;
SELECT pg_is_in_recovery();                    -- t = standby
SELECT datname, pg_size_pretty(pg_database_size(datname)) FROM pg_database;
```

### Fichiers

```
/etc/postgresql/17/main/postgresql.conf
/etc/postgresql/17/main/pg_hba.conf
/var/lib/postgresql/17/main/          (PGDATA)
/var/log/postgresql/
```

### Numéros à connaître

```
5432            port par défaut
25 % RAM        shared_buffers (point de départ)
50-75 % RAM     effective_cache_size
100             max_connections par défaut
```

---

## 82. Glossaire

