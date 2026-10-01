---
id: collect-261001-rattrapage/rattrapage/redis-guide-13
title: "Guide Redis — De l'installation à la production"
domain: rattrapage
role: reference
task: reference
actors: []
dates: ["2026-09-26"]
keywords: ["arr", "backlog", "benchmark", "benchmarks", "latency", "memory"]
source: docs/RAG/collect-261001-rattrapage/redis_guide.md
source_anchor: ""
source_lines: [2341, 2535]
sha256: 58b2d39ec1e5ea86cf9e449a40410a3e9f1fa099304520f9dcf633d892df97d9
---

# Guide Redis — De l'installation à la production

Quelques grosses clés expliquent souvent 80 % de la mémoire. Les trouver **sans bloquer** :

```bash
# --- redis-cli --bigkeys : échantillonne (non bloquant, mais lent sur gros datasets) ---
redis-cli --bigkeys
# Sampled 1000000 keys in the keyspace!
# biggest string found 'cache:rapport:mensuel' has 42,123,456 bytes

# --- MEMORY USAGE : taille précise d'une clé ---
redis-cli MEMORY USAGE cache:produit:42
# (integer) 1240

# --- SCAN + MEMORY USAGE en script (top 20) ---
redis-cli --scan --pattern "*" | head -100000 | \
  while read -r k; do printf "%s %s\n" "$(redis-cli --raw MEMORY USAGE "$k")" "$k"; done | \
  sort -rn | head -20

# --- redis-cli --memkeys (⚠️ 7.x, outils récents) : top mémoire via échantillonnage ---
```

Actions sur big keys trouvées :

- String > 10 Mo : externaliser (S3/NFS) ou compresser côté appli.
- Hash/List/Set énorme : **sharder** la clé (`logs:2026-09-26`, `file:1`, `file:2`...).
- Vérifier l'encodage : `redis-cli DEBUG OBJECT ma_cle` (ou `OBJECT ENCODING`) — `raw` vs `listpack`/`intset` change tout sur les petites collections.

---

## 75. Benchmarks : mesurer avant d'optimiser

```bash
# --- redis-benchmark (fourni avec Redis) ---
# 100 000 requêtes, 50 clients parallèles, que des SET :
redis-benchmark -h 10.0.5.21 -a "$REDIS_PASS" -n 100000 -c 50 -t set,get

# Pipeline (regroupe les requêtes : le vrai gain réseau) :
redis-benchmark -h 10.0.5.21 -a "$REDIS_PASS" -n 100000 -c 50 -t set,get -P 16

# Benchmark réaliste : mélange + tailles de payload :
redis-benchmark -h 10.0.5.21 -a "$REDIS_PASS" -n 200000 -c 100 -d 256 \
  --threads 4 -t set,get,hset,hget,lpush,lpop

# --- Mesurer la latence intrinsèque ---
redis-cli --latency -h 10.0.5.21
# min: 0, max: 12, avg: 0.18 (ms) — Ctrl+C pour arrêter
redis-cli --latency-history -i 1 -h 10.0.5.21   # historique par seconde
```

Règles d'interprétation :

- Benchmarkez **depuis un client sur le même réseau** que la prod (pas depuis votre laptop via VPN).
- Le goulot est presque toujours le **réseau** (allers-retours), pas Redis → utilisez le **pipelining** côté appli.
- Comparez avant/après chaque changement (io-threads, maxmemory-policy, AOF...) avec la **même** commande.

---


## 76. Dépannage : Redis ne démarre pas

```bash
# 1. Lire le log EN PREMIER (90 % des diagnostics sont là) :
sudo tail -50 /var/log/redis/redis-server.log
sudo journalctl -u redis-server --no-pager | tail -30

# 2. Tester la config :
redis-server /etc/redis/redis.conf --test-memory 2
# ou : redis-check-rdb /var/lib/redis/dump.rdb  (si soupçon de RDB corrompu)
```

| Message / symptôme | Cause probable | Remède |
|---|---|---|
| `Can't open the append-only file: Permission denied` | Droits sur `/var/lib/redis` | `chown -R redis:redis /var/lib/redis` |
| `Fatal error loading the DB` | RDB/AOF corrompu | Restaurer le backup (section 67) ; AOF : `redis-check-aof --fix` |
| `Address already in use` | Port 6379 occupé (2e instance ?) | `ss -ltnp \| grep 6379`, tuer le doublon |
| `Can't open the log file` | `/var/log/redis` inexistant ou droits | Créer + chown redis |
| `Unknown directive` | Faute de frappe dans redis.conf | Corriger la ligne indiquée |
| `OOM command not allowed` | Démarre mais refuse tout (vieux RDB + maxmemory) | Augmenter maxmemory temporairement via CONFIG |
| Démarre puis s'arrête (systemd) | `supervised` mal réglé / `daemonize yes` + systemd | Paquet Debian : `supervised systemd`, `daemonize no` |

---

## 77. Dépannage : OOM / évictions massives

Symptômes : erreurs `OOM command not allowed when used memory > 'maxmemory'`, `evicted_keys` qui explose, hit rate en chute.

```bash
# --- Diagnostic ---
redis-cli INFO memory | grep -E "used_memory_human|maxmemory_human|mem_fragmentation"
redis-cli INFO stats | grep -E "evicted_keys|expired_keys"
redis-cli CONFIG GET maxmemory-policy
redis-cli --bigkeys | tail -20        # qui consomme ?

# --- Actions immédiates ---
# 1. Augmenter temporairement (si la RAM le permet) :
redis-cli CONFIG SET maxmemory 6gb
# 2. Purger un cache reconstituable :
redis-cli --scan --pattern "cache:tmp:*" | xargs -L 500 redis-cli DEL
# 3. Vérifier qu'un process ne remplit pas (fuite applicative) :
redis-cli INFO keyspace   # db0:keys=... en croissance continue ?
# 4. Surveiller en continu :
watch -n 2 'redis-cli INFO memory | grep used_memory_human'
```

> Si `used_memory` < `maxmemory` mais erreurs OOM quand même : c'est la **fragmentation** (section 73) ou un **fork** qui a échoué (overcommit, section 26). Lisez le log.

---

## 78. Dépannage : réplication cassée

```bash
# --- Sur le réplica ---
redis-cli INFO replication | grep -E "master_link_status|master_last_io_seconds_ago|slave_repl_offset"
# master_link_status:down  → le lien est coupé

# --- Sur le maître ---
redis-cli INFO replication | grep -E "connected_slaves|slave0"

# --- Log du réplica : indices ---
sudo grep -iE "replication|psync|sync" /var/log/redis/redis-server.log | tail -20
```

| Symptôme | Cause | Remède |
|---|---|---|
| `master_link_status:down` | Réseau / firewall / maître down | ping, `ss`, vérifier le maître |
| Resync complète en boucle | `repl-backlog-size` trop petit vs débit d'écriture | Augmenter `repl-backlog-size` (ex. 64mb) |
| `NOAUTH Authentication required` | Mot de passe réplication faux/absent | `masteruser`/`masterauth` sur le réplica |
| `READONLY` côté appli après failover | L'appli écrit encore sur l'ancien maître | Basculer l'appli (DNS/VIP/Sentinel-aware) |
| Lag qui grandit | Réplica saturé (CPU/réseau) | `INFO stats` sur le réplica, dimensionner |

```ini
# Prévention :
repl-backlog-size 64mb
replica-serve-stale-data yes
```

---

## 79. Dépannage : latence élevée

```bash
# 1. La latence vient-elle du serveur ou du réseau ?
redis-cli --latency-history -i 1        # latence vue du client
redis-cli LATENCY DOCTOR                # diagnostic serveur
redis-cli INFO latencystats             # percentiles par commande

# 2. Chercher le coupable :
redis-cli SLOWLOG GET 20

# 3. Vérifier le système :
uptime                                  # load average
free -h                                 # swap utilisé ?
dmesg | tail -20                        # OOM killer ?
cat /sys/kernel/mm/transparent_hugepage/enabled   # doit être [never]
```

Checklist express : slowlog → THP → swap → fork (BGSAVE en cours ? `INFO persistence`) → `appendfsync always` → saturation réseau → `blocked_clients` anormal.

---

## 80. Dépannage : AOF corrompu au démarrage

```
# Log typique :
# Bad file format reading the append only file: make a backup of your AOF file, then use ./redis-check-aof --fix <filename>
```

```bash
# 1. SAUVEGARDER avant toute chose :
sudo cp -a /var/lib/redis/appendonlydir /srv/backups/redis/appendonlydir.corrompu-$(date +%F)

# 2. Tenter la réparation (tronque la queue corrompue — perte des dernières écritures) :
# Ancien format mono-fichier :
redis-check-aof --fix /var/lib/redis/appendonly.aof
# Format 7.0+ multi-part : identifier le segment incrémental fautif dans appendonlydir/,
# le réparer ou le retirer du manifeste (documentez ce que vous faites !)

# 3. Alternative sûre : repartir du dernier RDB sain + accepter la perte AOF
# (voir procédure de restauration section 67)

# 4. Prévention : onduleur (coupures = 1re cause), appendfsync everysec minimum,
#    aof-load-truncated yes (déjà conseillé section 29)
```

---

## 81. Dépannage : Cluster en état `fail`

```bash
redis-cli --cluster check 10.0.5.41:6379
# [ERR] Not all 16384 slots are covered by nodes.

redis-cli -c -h 10.0.5.41 CLUSTER INFO | grep cluster_state
# cluster_state:fail
```

