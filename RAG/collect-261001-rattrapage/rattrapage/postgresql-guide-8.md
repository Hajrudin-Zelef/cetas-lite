---
id: collect-261001-rattrapage/rattrapage/postgresql-guide-8
title: "PostgreSQL en production — Guide complet pour sysadmin"
domain: rattrapage
role: reference
task: reference
actors: []
dates: ["2026-09-20", "2026-09-26"]
keywords: ["arr"]
source: docs/RAG/collect-261001-rattrapage/postgresql_guide.md
source_anchor: ""
source_lines: [1624, 1852]
sha256: 51b7f578541763036a68a73b74f10c65158c4bfbe754249062f133ea780cd1c1
---

# PostgreSQL en production — Guide complet pour sysadmin

```bash
# Sauvegarde physique complète (prérequis : wal_level=replica, rôle REPLICATION)
pg_basebackup -h srv-bdd -U replicateur -D /srv/sauvegardes/postgresql/physique/base-$(date +%F) \
  -Ft -z -P --wal-method=stream

# Options clés
# -Ft -z   : tar + gzip
# -P       : progression
# --wal-method=stream : inclut le WAL nécessaire à la cohérence
# -R       : écrit la conf de réplication (pour cloner un standby)
```

**Prérequis serveur (primaire) :**

```ini
# postgresql.conf
wal_level = replica
max_wal_senders = 5
max_replication_slots = 5
```

```sql
-- Rôle de réplication (pas de superuser nécessaire)
CREATE ROLE replicateur WITH REPLICATION LOGIN PASSWORD '...';
```

```
# pg_hba.conf : autoriser la réplication depuis le standby/outil
host replication replicateur 10.0.2.10/32 scram-sha-256
```

---

## 51. PITR : archivage WAL et restauration à un instant T

Le **PITR** (Point-In-Time Recovery) rejoue le WAL archivé sur une base
physique : restauration à **n'importe quel instant** couvert.

### 51.1. Configurer l'archivage (primaire)

```ini
# postgresql.conf
archive_mode = on
archive_command = 'test ! -f /srv/sauvegardes/postgresql/wal/%f && cp %p /srv/sauvegardes/postgresql/wal/%f'
archive_timeout = 600   # force un segment toutes les 10 min max
```

```sql
-- Vérifier que l'archivage fonctionne
SELECT * FROM pg_stat_archiver;
-- failed_count doit rester à 0 ; last_failed_time NULL
```

> ⚠️ `archive_command` doit **échouer** si la copie échoue (sinon PG
> croit le segment archivé). Le `test ! -f` évite d'écraser.
> Surveillez l'espace : un archivage en panne fait **grossir pg_wal**
> jusqu'à remplir le disque (alerte section 63).

### 51.2. Restaurer à un instant T (procédure)

```bash
# 1. Arrêter le cluster cible (serveur de restauration isolé !)
sudo pg_ctlcluster 17 main stop

# 2. Vider PGDATA et restaurer la base physique
rm -rf /var/lib/postgresql/17/main/*
tar -xzf /srv/sauvegardes/postgresql/physique/base-2026-09-20.tar.gz \
  -C /var/lib/postgresql/17/main/

# 3. Fichier de recovery : rejouer le WAL jusqu'à l'instant voulu
cat > /var/lib/postgresql/17/main/postgresql.auto.conf <<'EOF'
restore_command = 'cp /srv/sauvegardes/postgresql/wal/%f %p'
recovery_target_time = '2026-09-26 14:30:00+02'
recovery_target_action = 'promote'
EOF
touch /var/lib/postgresql/17/main/recovery.signal
chown -R postgres:postgres /var/lib/postgresql/17/main

# 4. Démarrer : PG rejoue le WAL jusqu'à 14h30 puis promeut
sudo pg_ctlcluster 17 main start
# Suivre : tail -f /var/log/postgresql/postgresql-17-main.log
# Message attendu : "recovery stopping before ... consistent recovery state reached"

# 5. Vérifier les données, puis OUVRIR aux applis (jamais avant)
```

**Cibles de restauration :** `recovery_target_time`, `recovery_target_xid`,
`recovery_target_lsn`, `recovery_target_name` (point nommé via
`pg_create_restore_point('avant_migration')`).

> ⚠️ Restaurez **toujours** sur un serveur isolé d'abord. Une fois promu,
> le standby restauré a une nouvelle timeline : il ne peut pas redevenir
> standby de l'ancien primaire sans reconstruction.

---

## 52. Barman : l'archivage WAL industrialisé

**Barman** (pgBarman) centralise : base backups + WAL streaming +
rétention + restauration en une commande.

```bash
# Installation (serveur de sauvegarde dédié)
sudo apt install -y barman

# /etc/barman.d/srv-bdd.conf
[srv-bdd]
description = "Primaire production"
conninfo = host=srv-bdd user=barman dbname=postgres
backup_method = rsync
streaming_archiver = on
slot_name = barman
retention_policy = RECOVERY WINDOW OF 14 DAYS
wal_retention_policy = MAIN
```

```bash
# Vérifier la conf, créer le slot, lancer la première sauvegarde
barman check srv-bdd
barman receive-wal --create-slot srv-bdd
barman backup srv-bdd

# Restaurer la dernière sauvegarde + PITR à un instant T
barman recover --target-time "2026-09-26 14:30:00+02" \
  srv-bdd latest /var/lib/postgresql/17/main/

# Lister / purger selon la rétention
barman list-backup srv-bdd
barman cron   # à mettre en cron : maintenance + rétention
```

**Pourquoi Barman plutôt que des scripts maison :** slots de réplication
dédiés (pas de perte de WAL), `barman check` diagnostique, restauration
testable en une commande, politiques de rétention déclaratives.

---

## 53. Réplication streaming : primaire + standby

```bash
# 1. Sur le primaire : rôle + pg_hba (voir section 50), puis slot
sudo -u postgres psql -c \
  "SELECT pg_create_physical_replication_slot('standby1');"

# 2. Sur le futur standby : cloner depuis le primaire (cluster arrêté)
sudo pg_ctlcluster 17 main stop
sudo -u postgres rm -rf /var/lib/postgresql/17/main/*
sudo -u postgres pg_basebackup -h srv-primaire -U replicateur \
  -D /var/lib/postgresql/17/main -P -R --slot=standby1

# 3. Démarrer le standby : -R a créé standby.signal + primary_conninfo
sudo pg_ctlcluster 17 main start

# 4. Vérifier des deux côtés
# Primaire :
sudo -u postgres psql -c \
  "SELECT client_addr, state, sync_state, replay_lag FROM pg_stat_replication;"
# Standby (lecture seule) :
sudo -u postgres psql -c "SELECT pg_is_in_recovery();"  -- t -> c'est un standby
```

**Usages du standby :**

- Requêtes **lecture seule** (reporting) : `hot_standby = on` (défaut).
- Sauvegardes déportées (`pg_basebackup` depuis le standby).
- Bascule en cas de panne (section 54).

**Retard de réplication (lag) :**

```sql
-- Sur le primaire : lag par standby
SELECT client_addr, state,
       pg_size_pretty(pg_wal_lsn_diff(pg_current_wal_lsn(), replay_lsn)) AS lag
FROM pg_stat_replication;
```

> ⚠️ Un standby en retard + `max_standby_archive_delay`/`max_standby_streaming_delay`
> trop courts = requêtes annulées sur le standby ("canceling statement
> due to conflict"). Pour du reporting lourd, augmentez ces délais ou
> utilisez `hot_standby_feedback = on` (au prix d'un risque de bloat
> sur le primaire).

---

## 54. Failover manuel : promouvoir le standby

Procédure de bascule planifiée (ou d'urgence si le primaire est mort).

```bash
# --- BASCULE PLANIFIÉE (sans perte) ---
# 1. Couper les écritures applicatives (maintenance applicative)
# 2. Attendre la synchronisation complète
sudo -u postgres psql -h srv-primaire -c \
  "SELECT pg_size_pretty(pg_wal_lsn_diff(pg_current_wal_lsn(), replay_lsn))
   FROM pg_stat_replication WHERE client_addr = '10.0.2.11';"
# -> 0 bytes : synchronisé

# 3. Promouvoir le standby
sudo -u postgres pg_ctlcluster 17 main promote
# ou : SELECT pg_promote();

# 4. Vérifier
sudo -u postgres psql -c "SELECT pg_is_in_recovery();"  -- f : primaire

# 5. Repointer les applis (DNS, VIP, ou chaîne de connexion)
# 6. Reconstruire l'ancien primaire comme standby (pg_basebackup inversé)
```

```bash
# --- BASCULE D'URGENCE (primaire injoignable) ---
# 1. S'assurer que le primaire est VRAIMENT mort (pas de split-brain !)
#    STONITH / arrêt électrique si doute.
# 2. Promouvoir (perte = WAL non répliqué)
sudo -u postgres pg_ctlcluster 17 main promote
# 3. Repointer les applis, vérifier les données critiques
# 4. Quand le primaire revient : le RECONSTRUIRE comme standby
```

**Anti-split-brain :** un seul écrivain à la fois. En manuel : procédure
écrite + checklist. En automatique : Patroni (section 55) avec DCS
(etcd/Consul).

---

## 55. Patroni : tour d'horizon de la haute disponibilité auto

**Patroni** orchestre le failover automatique via un magasin distribué
(DCS : etcd le plus courant).

