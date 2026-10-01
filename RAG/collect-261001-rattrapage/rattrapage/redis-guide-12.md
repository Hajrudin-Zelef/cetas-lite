---
id: collect-261001-rattrapage/rattrapage/redis-guide-12
title: "Guide Redis — De l'installation à la production"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["latency", "memory"]
source: docs/RAG/collect-261001-rattrapage/redis_guide.md
source_anchor: ""
source_lines: [2151, 2340]
sha256: 2a9655a8161de152c0624dbea83bcf8daab044883e6bbe9ac5db9dc8fb1c6e7a
---

# Guide Redis — De l'installation à la production

`INFO` est la commande de supervision n°1. Sections utiles :

```bash
redis-cli INFO server       # version, uptime, process_id
redis-cli INFO clients      # connected_clients, blocked_clients
redis-cli INFO memory        # used_memory_human, mem_fragmentation_ratio
redis-cli INFO persistence   # rdb_last_bgsave_status, aof_last_rewrite_time_sec
redis-cli INFO stats         # instantaneous_ops_per_sec, hit/miss, expired/evicted
redis-cli INFO replication   # role, connected_slaves, master_link_status
redis-cli INFO cpu           # used_cpu_sys/user
redis-cli INFO keyspace      # clés par base : db0:keys=...,expires=...
redis-cli INFO commandstats  # appels par commande (⚠️ 7.0+ : INFO commandstats)
redis-cli INFO errorstats    # erreurs par type (⚠️ 7.0+)
redis-cli INFO latencystats  # percentiles de latence (⚠️ 7.0+, voir section 71)
```

Exemple d'indicateurs clés en une commande :

```bash
redis-cli INFO stats | grep -E "instantaneous_ops|keyspace_hits|keyspace_misses|expired_keys|evicted_keys|total_commands_processed"
# instantaneous_ops_per_sec:1240
# keyspace_hits:98213
# keyspace_misses:1204
# → hit rate = 98213 / (98213+1204) ≈ 98,8 % ✅
```

---

## 69. Monitoring : métriques et seuils d'alerte

Tableau des alertes à configurer (Prometheus/Grafana, Zabbix, ou autre) :

| Métrique | Source | Seuil warning | Seuil critique | Signification |
|---|---|---|---|---|
| `up` / ping | `redis-cli ping` | — | down | Instance injoignable |
| `connected_clients` | INFO clients | > 80 % maxclients | = maxclients | Saturation connexions |
| `blocked_clients` | INFO clients | > 10 | > 50 | Workers bloqués (BLPOP/XREAD) anormalement nombreux |
| `used_memory_rss` / RAM | INFO memory + OS | > 80 % RAM | > 90 % | Risque OOM |
| `mem_fragmentation_ratio` | INFO memory | > 1,5 ou < 1 | > 2 | Fragmentation (restart si > 2 durable) |
| `evicted_keys` (dérivée) | INFO stats | > 0/min | > 1000/min | maxmemory trop bas ou fuite |
| `keyspace hit rate` | hits/(hits+misses) | < 90 % | < 70 % | Cache inefficace |
| `rdb_last_bgsave_status` | INFO persistence | — | err | Échec snapshot |
| `aof_last_bgrewrite_status` | INFO persistence | — | err | Échec rewrite AOF |
| `master_link_status` | INFO replication | — | down | Réplica décroché |
| `instantaneous_ops_per_sec` | INFO stats | chute brutale | 0 avec trafic | Blocage / saturation |
| `slowlog` récent | SLOWLOG GET | latence p99 | commandes > 100 ms | Commandes lentes (KEYS, HGETALL...) |
| Certificat TLS | fichier | expire < 30 j | expire < 7 j | Renouvellement |

Exemple d'exporter : **redis_exporter** (Oliver006) :

```bash
# Binaire ou Docker, avec AUTH :
docker run -d --name redis_exporter -p 9121:9121 \
  -e REDIS_ADDR="redis://10.0.5.21:6379" \
  -e REDIS_PASSWORD="SecretRedisProd" \
  oliver006/redis_exporter:latest
# → métriques Prometheus sur :9121/metrics
curl -s localhost:9121/metrics | grep -E "^redis_up|^redis_memory_used_bytes"
```

---

## 70. Monitoring : SLOWLOG (chasse aux commandes lentes)

Le slowlog enregistre les commandes dépassant un seuil (microsecondes). En prod, c'est votre **détecteur de `KEYS *`** et de `HGETALL` géants.

```bash
# --- Configuration (10 ms par défaut ; en prod : 1-5 ms) ---
redis-cli CONFIG SET slowlog-log-slower-than 2000   # 2000 µs = 2 ms
redis-cli CONFIG SET slowlog-max-len 256            # nb d'entrées conservées

# --- Consultation ---
redis-cli SLOWLOG GET 10
# 1) 1) (integer) 12        → id
#    2) (integer) 1769557200 → timestamp
#    3) (integer) 45213      → durée en µs (45 ms !)
#    4) 1) "KEYS" 2) "*"     → la commande coupable

redis-cli SLOWLOG LEN
redis-cli SLOWLOG RESET   # vider (après analyse)

# --- Rendre persistant ---
# /etc/redis/conf.d/99-production.conf :
# slowlog-log-slower-than 2000
# slowlog-max-len 256
```

> Un slowlog plein de `KEYS`, `HGETALL`, `SMEMBERS` sur de grosses clés = **code applicatif à corriger** (SCAN, HMGET, SSCAN...). Le slowlog ne ment pas.

---

## 71. Monitoring : latence (LATENCY DOCTOR)

Redis 7 mesure la latence par commande (percentiles) sans surcoût notable :

```bash
redis-cli CONFIG SET latency-tracking yes
redis-cli CONFIG SET latency-tracking-info-percentiles "50 99 99.9"

# Attendre un peu de trafic, puis :
redis-cli INFO latencystats
# latency_percentiles_usec_get:p50=8.031,p99=15.103,p99.9=22.015
# latency_percentiles_usec_hgetall:p50=120.5,p99=4500.0,...

# Diagnostic automatique :
redis-cli LATENCY DOCTOR
# Dave, I have found 2 latency spikes... (analyse les 160 derniers événements)

redis-cli LATENCY HISTORY command    # 160 derniers échantillons de la classe "command"
redis-cli LATENCY LATEST
redis-cli LATENCY RESET
```

Causes classiques de pics de latence :

| Cause | Indice | Remède |
|---|---|---|
| Commande lourde (KEYS, SORT...) | slowlog corrélé | Réécrire en SCAN / Lua borné |
| Fork BGSAVE/AOF | pics réguliers | `vm.overcommit_memory=1`, THP off |
| Swap | latence aléatoire | `vm.swappiness=1`, assez de RAM |
| AOF `always` | latence d'écriture | passer en `everysec` |
| Réseau | latence client >> serveur | `LATENCY HISTORY` vs mesure client |

---

## 72. Supervision : exemple d'alerting complet (script)

Script de health-check à brancher sur votre supervision (Nagios/Zabbix : code retour + message) :

```bash
#!/bin/bash
# /usr/local/bin/redis-health.sh — code 0 OK, 1 WARNING, 2 CRITICAL
HOST="${1:-127.0.0.1}"; PORT="${2:-6379}"
AUTH_ARGS=(-a "$REDIS_PASS")   # export REDIS_PASS depuis le coffre

ping_out=$(redis-cli -h "$HOST" -p "$PORT" "${AUTH_ARGS[@]}" ping 2>&1)
[ "$ping_out" != "PONG" ] && { echo "CRITICAL: ping=$ping_out"; exit 2; }

mem=$(redis-cli -h "$HOST" -p "$PORT" "${AUTH_ARGS[@]}" --raw INFO memory | grep used_memory_rss | cut -d: -f2)
maxmem=$(redis-cli -h "$HOST" -p "$PORT" "${AUTH_ARGS[@]}" --raw CONFIG GET maxmemory | tail -1)
if [ "$maxmem" != "0" ] && [ "$mem" -gt $(( maxmem * 90 / 100 )) ]; then
  echo "WARNING: mémoire à plus de 90% de maxmemory"; exit 1
fi

bgsave=$(redis-cli -h "$HOST" -p "$PORT" "${AUTH_ARGS[@]}" --raw INFO persistence | grep rdb_last_bgsave_status | cut -d: -f2 | tr -d '\r')
[ "$bgsave" != "ok" ] && { echo "CRITICAL: dernier BGSAVE en échec"; exit 2; }

role=$(redis-cli -h "$HOST" -p "$PORT" "${AUTH_ARGS[@]}" --raw INFO replication | grep '^role' | cut -d: -f2 | tr -d '\r')
if [ "$role" = "slave" ]; then
  link=$(redis-cli -h "$HOST" -p "$PORT" "${AUTH_ARGS[@]}" --raw INFO replication | grep master_link_status | cut -d: -f2 | tr -d '\r')
  [ "$link" != "up" ] && { echo "CRITICAL: lien réplication down"; exit 2; }
fi

echo "OK: redis $HOST:$PORT (role=$role)"
exit 0
```

---

## 73. Gestion mémoire : fragmentation et jemalloc

```bash
redis-cli INFO memory | grep -E "used_memory_human|used_memory_rss_human|mem_fragmentation_ratio|mem_allocator"
# used_memory_human:3.20G        → ce que Redis alloue (logique)
# used_memory_rss_human:4.10G    → ce que l'OS voit (résident)
# mem_fragmentation_ratio:1.28   → rss / used
# mem_allocator:jemalloc-5.3.0
```

Lecture du ratio :

| Ratio | Interprétation | Action |
|---|---|---|
| ~1,0–1,3 | Sain | Rien |
| > 1,5 | Fragmentation notable | Surveiller ; `MEMORY PURGE` (défrag jemalloc partielle) |
| > 2,0 durable | Fragmentation sévère | Planifier un **restart** (failover propre) |
| < 1,0 | Swap probable ! | Urgent : RAM/swap (le RSS < alloué = pages swappées) |

```bash
redis-cli MEMORY PURGE    # demande à l'allocateur de rendre la mémoire (effet variable)
redis-cli MEMORY STATS    # détail par classe d'allocation
redis-cli MEMORY DOCTOR   # diagnostic texte (⚠️ 4.0+)
```

> La défragmentation active (`activedefrag yes`, ⚠️ 4.0+) peut aider sur des workloads avec beaucoup d'allocations/libérations, au prix d'un peu de CPU. À tester avant d'activer en prod.

---

## 74. Gestion mémoire : traquer les big keys

