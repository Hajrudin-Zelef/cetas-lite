---
id: collect-261001-rattrapage/rattrapage/redis-guide-16
title: "Guide Redis — De l'installation à la production"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["consumer", "latency", "memory"]
source: docs/RAG/collect-261001-rattrapage/redis_guide.md
source_anchor: ""
source_lines: [2948, 3140]
sha256: eceb910dea98bc2bb7c33d5befe024dbaa5f3d11010f7ed981fedcb71762e7f5
---

# --- DIVERS PROD ---
hz 10
activerehashing yes
# io-threads 4                   # à activer si saturation réseau mesurée
# io-threads-do-reads yes
notify-keyspace-events ""
# tls-port 6379 / port 0         # si TLS (voir section 38)
```

---

## 93. Checklist de mise en production

À imprimer et cocher pour chaque nouvelle instance :

**Installation**
- [ ] Paquet du dépôt officiel, version épinglée (`apt-mark hold`)
- [ ] `vm.overcommit_memory=1`, `somaxconn`, THP `never` (section 26)
- [ ] Service systemd `enabled`, redémarrage testé

**Configuration**
- [ ] `bind` restreint, `protected-mode yes`
- [ ] `maxmemory` + politique d'éviction adaptée
- [ ] Persistance : RDB + AOF `everysec` (ou choix assumé, section 30)
- [ ] `timeout`, `tcp-keepalive`, `maxclients` réglés
- [ ] `slowlog-log-slower-than` ≤ 5000 µs

**Sécurité**
- [ ] ACL : `default off`, users applicatifs moindre-privilège (section 36)
- [ ] Secrets en coffre, jamais dans git
- [ ] `rename-command` FLUSHALL/FLUSHDB/DEBUG
- [ ] Firewall : 6379 (et 16379 si cluster, 26379 si sentinel) filtrés
- [ ] TLS si réseau non fiable
- [ ] Scan externe : port non visible d'Internet

**Haute disponibilité**
- [ ] Réplication : 1 maître + 2 réplicas, `min-replicas-to-write 1`
- [ ] Sentinel × 3 (quorum 2) **ou** Cluster × 6 — testé par failover réel
- [ ] Clients Sentinel/Cluster-aware

**Sauvegarde**
- [ ] Script de backup RDB quotidien + rétention 30 j
- [ ] Copie hors site (règle 3-2-1)
- [ ] **Restauration testée** et chronométrée (RTO connu)

**Supervision**
- [ ] `redis_exporter` + dashboard Grafana (ou équivalent)
- [ ] Alertes : down, mémoire, évictions, hit rate, bgsave, lien réplication, latence
- [ ] Health-check scripté (section 72)

**Documentation**
- [ ] Topologie et mots de passe (coffre) documentés
- [ ] Procédure de failover et de restauration écrites
- [ ] Astreinte informée des seuils d'alerte

---

## 94. Pense-bête de poche

```
CONNEXION
  redis-cli -h HOST -p 6379 -a PASS        connexion
  redis-cli ping                           PONG ?
  redis-cli -c -h HOST                     mode cluster

STRINGS
  SET k v [EX s|PX ms] [NX|XX]             écriture (+TTL, conditionnelle)
  GET k / MGET k1 k2 / GETDEL k            lectures
  INCR k / INCRBY k n / DECR k             compteurs atomiques
  APPEND k v / STRLEN k                    concat / longueur

HASHES
  HSET k f v [f v...] / HGET k f / HMGET k f1 f2
  HGETALL k / HDEL k f / HEXISTS k f / HLEN k
  HINCRBY k f n

LISTS (files)
  RPUSH k v / LPUSH k v / LLEN k
  LPOP k / RPOP k / BLPOP k timeout        (BLPOP = bloquant)
  LRANGE k 0 -1 / LTRIM k -100 -1          (historique borné)
  LMOVE src dst LEFT RIGHT                (transfert atomique)

SETS
  SADD k m / SREM k m / SMEMBERS k / SISMEMBER k m
  SINTER / SUNION / SDIFF k1 k2            (opérations ensemblistes)
  SINTERSTORE dst k1 k2                    (stocke le résultat)

SORTED SETS
  ZADD k score m / ZSCORE k m / ZRANK k m / ZREVRANK k m
  ZRANGE k 0 -1 [REV] [WITHSCORES]         (⚠️ 6.2+ : REV remplace ZREVRANGE)
  ZRANGE k min max BYSCORE / BYLEX
  ZINCRBY k n m / ZPOPMIN k / ZPOPMAX k
  ZREMRANGEBYSCORE k min max

STREAMS
  XADD k MAXLEN ~ N * f v ...              produire
  XREAD BLOCK ms STREAMS k $               lire (bloquant)
  XGROUP CREATE k grp 0 MKSTREAM           créer un groupe
  XREADGROUP GROUP grp c COUNT n STREAMS k ">"
  XACK k grp id [id...]                    acquitter
  XAUTOCLAIM k grp c ms 0 COUNT n          récupérer les pannes
  XLEN k / XPENDING k grp

CLÉS / TTL
  DEL k / UNLINK k (non bloquant) / EXISTS k
  EXPIRE k s / PEXPIRE k ms / TTL k / PTTL k / PERSIST k
  SCAN 0 MATCH "pat:*" COUNT 1000         (JAMAIS keys *)
  RENAME k nk / TYPE k / DUMP k / RESTORE k ttl v

TRANSACTIONS / LUA
  MULTI ... EXEC / DISCARD / WATCH k
  EVAL "script" nbcles k1 [k2...] arg1 [arg2...]
  EVALSHA sha nbcles k1... / SCRIPT LOAD / SCRIPT FLUSH

PUB/SUB
  SUBSCRIBE c / PSUBSCRIBE "pat:*" / PUBLISH c msg
  SSUBSCRIBE / SPUBLISH                    (sharded, ⚠️ 7.0+)

ADMIN
  INFO [section] / CONFIG GET|SET param / CONFIG REWRITE
  BGSAVE / BGREWRITEAOF / LASTSAVE / DBSIZE / FLUSHDB (⚠️)
  SLOWLOG GET n / LATENCY DOCTOR / MEMORY DOCTOR
  CLIENT LIST / CLIENT KILL / MONITOR (⚠️ perf)
  ACL LIST / ACL SETUSER ... / ACL SAVE
  REPLICAOF host port / REPLICAOF NO ONE

SENTINEL (port 26379)
  SENTINEL masters / SENTINEL sentinels mymaster
  SENTINEL get-master-addr-by-name mymaster
  SENTINEL FAILOVER mymaster

CLUSTER
  redis-cli --cluster create h1:6379 ... --cluster-replicas 1
  redis-cli --cluster check h1:6379 / --cluster info h1:6379
  CLUSTER NODES / CLUSTER SLOTS / CLUSTER INFO
```

---

## 95. Glossaire

| Terme | Définition |
|---|---|
| ACL | Access Control List : utilisateurs, commandes et clés autorisés (Redis 6+) |
| AOF | Append Only File : journal des écritures, rejoué au redémarrage |
| Big key | Clé anormalement grosse, source de latence/mémoire |
| BLPOP/BRPOP | Variantes **bloquantes** de LPOP/RPOP (attendent un élément) |
| Cache-aside | Pattern : l'appli lit Redis, va en base si miss |
| Cluster | Mode sharding : 16 384 slots répartis sur N maîtres |
| Consumer group | Groupe de consommateurs se partageant un stream |
| CROSSSLOT | Erreur : opération multi-clés sur des slots différents |
| Éviction | Suppression d'une clé par `maxmemory-policy` (≠ expiration) |
| Expiration | Suppression automatique à fin de TTL (passive + active) |
| Failover | Bascule automatique maître → réplica (Sentinel/Cluster) |
| Hash slot | 1 des 16 384 compartiments du cluster (`CRC16(clé) mod 16384`) |
| Hash tag | `{...}` dans une clé pour forcer le même slot |
| Hit rate | `hits / (hits + misses)` : efficacité du cache |
| HyperLogLog | Compteur approximatif d'éléments distincts (12 Ko, erreur 0,81 %) |
| IO threads | Threads d'E/S réseau (Redis 6+, l'exécution reste mono-thread) |
| Jitter | Variation aléatoire des TTL pour éviter les avalanches d'expiration |
| Latence p99 | 99 % des requêtes sont plus rapides que cette valeur |
| LFU / LRU | Least Frequently/Recently Used : politiques d'éviction |
| Lua (script) | Script exécuté côté serveur, atomiquement |
| Maître / réplica | Nœud qui accepte les écritures / copie en lecture seule |
| maxmemory | Plafond mémoire ; au-delà : évictions ou erreurs OOM |
| MOVED / ASK | Redirections cluster vers le bon nœud |
| MULTI/EXEC | Transaction : exécution atomique d'une file de commandes |
| OOM | Out Of Memory : mémoire épuisée |
| Pipeline | Envoi groupé de commandes (1 aller-retour réseau) |
| Protected mode | Garde-fou refusant les connexions distantes non authentifiées |
| Pub/Sub | Messagerie publish/subscribe, sans persistance |
| Quorum | Majorité de Sentinels requise pour un failover |
| RDB | Snapshot binaire ponctuel (`dump.rdb`) |
| Redlock | Algorithme de verrou distribué sur N instances |
| Réplication | Copie asynchrone maître → réplica(s) |
| RESP | Protocole texte/binaire client-serveur de Redis |
| Sentinel | Supervision + failover automatique (port 26379) |
| Sharding | Partitionnement des données sur plusieurs nœuds |
| Single-threaded | Exécution des commandes sur un seul thread (atomicité) |
| Slowlog | Journal des commandes lentes (> seuil en µs) |
| Sorted Set | Ensemble trié par score (skiplist) |
| Split-brain | Partition réseau créant deux « maîtres » divergents |
| Stream | Journal append-only avec consumer groups (Redis 5+) |
| THP | Transparent Huge Pages : à désactiver (pics de latence) |
| TTL | Time To Live : durée de vie d'une clé |
| WAIT | Attente de la réplication sur N réplicas (semi-synchrone) |
| Write-through | Écriture synchrone cache + base |
| XACK / XPENDING | Acquittement / messages non acquittés d'un consumer group |

---

## 96. Quiz : 10 questions + réponses

