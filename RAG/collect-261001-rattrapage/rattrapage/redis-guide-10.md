---
id: collect-261001-rattrapage/rattrapage/redis-guide-10
title: "Guide Redis — De l'installation à la production"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["leaderboard"]
source: docs/RAG/collect-261001-rattrapage/redis_guide.md
source_anchor: ""
source_lines: [1761, 1940]
sha256: 0b2f576ac31b3eb9e7a6b3c694df8f7288bec2e5ffca7e02916c42d98e2c2b38
---

# --- Rang d'un joueur + voisins (pagination autour de lui) ---
redis-cli ZREVRANK leaderboard:quiz awa     # → 3 (0-based)
# voisins : ZRANGE leaderboard:quiz <rank-2> <rank+2> REV WITHSCORES

# --- Classement hebdo avec expiration ---
redis-cli ZADD leaderboard:2026-W39 1200 awa 950 moussa
redis-cli EXPIRE leaderboard:2026-W39 604800   # 7 jours

# --- Pourcentage / percentile (côté appli) ---
# rank = ZREVRANK, total = ZCARD → percentile = (total - rank) / total * 100
```

---

## 57. Réplication : maître / réplica

Un **réplica** copie en continu un **maître** (asynchrone par défaut). Usages : lectures réparties, backup à chaud, bascule manuelle.

```bash
# --- Sur le futur réplica (ou dans son redis.conf : replicaof 10.0.5.21 6379) ---
redis-cli REPLICAOF 10.0.5.21 6379
redis-cli INFO replication
# role:slave
# master_host:10.0.5.21
# master_link_status:up

# --- Sur le maître : vérifier ---
redis-cli INFO replication | grep -E "role|connected_slaves"

# --- Promouvoir un réplica en maître (bascule manuelle) ---
redis-cli REPLICAOF NO ONE
# (puis reconfigurer l'appli / le DNS / le VIP vers lui)

# --- Réplication en chaîne (réplica de réplica, pour répartir la charge réseau) ---
# Sur replica-2 : REPLICAOF <ip-replica-1> 6379
```

Points critiques :

- **Asynchrone par défaut** : un `SET` acquitté peut ne pas encore être sur le réplica si le maître meurt dans la foulée. `WAIT 1 1000` (attend l'ACK d'1 réplica, timeout 1 s) donne une **durabilité semi-synchrone** pour les écritures critiques.
- Un réplica est **en lecture seule** par défaut (`replica-read-only yes`) — laissez ainsi.
- `repl-diskless-sync yes` (défaut) : le maître envoie le RDB via socket sans toucher le disque — préférable sauf réseau très lent.
- Protège le maître : `maxclients`, et `client-output-buffer-limit replica 256mb 64mb 60` (un réplica lent ne doit pas OOM le maître).

---

## 58. Réplication : configuration pas à pas (2 nœuds)

```ini
# --- /etc/redis/conf.d/99-replication.conf sur le RÉPLICA ---
replicaof 10.0.5.21 6379
replica-read-only yes
replica-serve-stale-data yes   # sert les données (périmées) si le lien est coupé
```

```bash
# 1. Sur le maître : créer l'utilisateur de réplication (ACL)
redis-cli ACL SETUSER replicateur on ">SecretReplication" +psync +replconf +ping ~*

# 2. Sur le réplica : configurer l'authentification de réplication
#    /etc/redis/conf.d/99-replication.conf :
#    masterauth "SecretReplication"      # ⚠️ déprécié au profit de :
#    masteruser "replicateur"
#    (les deux lignes : user + password)

# 3. Redémarrer le réplica, vérifier la synchro :
redis-cli -h 10.0.5.22 INFO replication | grep master_link_status  # → up
redis-cli -h 10.0.5.21 INFO replication | grep connected_slaves   # → 1

# 4. Test : écrire sur le maître, lire sur le réplica
redis-cli -h 10.0.5.21 SET test:repli "ok"
redis-cli -h 10.0.5.22 GET test:repli   # → "ok"
redis-cli -h 10.0.5.22 DEL test:repli:rien
redis-cli -h 10.0.5.22 SET ecriture "x"
# (error) READONLY You can't write against a read only replica.

# 5. Nettoyer le test
redis-cli -h 10.0.5.21 DEL test:repli
```

> **Topologie prod minimale avec HA automatique :** 1 maître + 2 réplicas + 3 Sentinels (section 59). Sans Sentinel, la bascule est **manuelle**.

---

## 59. Sentinel : architecture et failover automatique

**Sentinel** surveille des maîtres Redis, détecte les pannes et **élit automatiquement** un réplica comme nouveau maître. C'est un processus séparé (`redis-sentinel`, même binaire avec un autre mode).

```
        ┌──────────┐  ┌──────────┐  ┌──────────┐
        │ Sentinel │  │ Sentinel │  │ Sentinel │
        │   (1)    │  │   (2)    │  │   (3)    │
        └────┬─────┴──┴────┬─────┴──┴────┬─────┘
             │             │             │
        ┌────▼─────────────▼─────────────▼────┐
        │        Maître Redis (6379)          │
        └────┬───────────────────────┬───────┘
             │                       │
      ┌──────▼──────┐         ┌──────▼──────┐
      │ Réplica (1) │         │ Réplica (2) │
      └─────────────┘         └─────────────┘
```

Concepts :

- **Quorum** : nombre de Sentinels devant être d'accord pour déclarer un maître `ODOWN` (objectively down) et lancer le failover. Quorum = **majorité** (2 sur 3, 3 sur 5).
- Les Sentinels **se découvrent** entre eux via le maître (pas besoin de tous les lister partout).
- Les clients **doivent** utiliser une lib « Sentinel-aware » (elle demande aux Sentinels « qui est le maître ? ») — un client pointé sur une IP fixe ne profitera pas du failover.
- `down-after-milliseconds` : délai avant de déclarer un nœud `SDOWN` (subjectively down). Prod : 5 000–10 000 ms (trop bas = failovers intempestifs).
- `failover-timeout` : temps max d'un failover (défaut 3 min).

---

## 60. Sentinel : configuration pas à pas

```bash
# --- 1. Fichier /etc/redis/sentinel.conf (identique sur les 3 sentinels, sauf port/myid auto) ---
sudo tee /etc/redis/sentinel.conf <<'EOF'
port 26379
dir /var/lib/redis
daemonize no
supervised systemd
logfile /var/log/redis/redis-sentinel.log

# Surveille le maître "mymaster" à 10.0.5.21:6379, quorum = 2
sentinel monitor mymaster 10.0.5.21 6379 2

# Authentification vers les nœuds Redis (si ACL/requirepass)
sentinel auth-pass mymaster SecretRedisProd
# sentinel auth-user mymaster sentinel_user   # ⚠️ 6.0+ avec ACL

# Délais
sentinel down-after-milliseconds mymaster 8000
sentinel failover-timeout mymaster 180000
sentinel parallel-syncs mymaster 1
EOF

# --- 2. Service systemd (paquet Debian : redis-sentinel.service existe si installé) ---
# Sinon, copiez /lib/systemd/system/redis-sentinel.service depuis le paquet.

# --- 3. Démarrer les 3 sentinels, vérifier ---
sudo systemctl enable --now redis-sentinel
redis-cli -p 26379 SENTINEL masters
redis-cli -p 26379 SENTINEL sentinels mymaster   # → les 3 sentinels se voient
redis-cli -p 26379 SENTINEL replicas mymaster

# --- 4. Tester le failover (EN MAINTENANCE, jamais un vendredi 18h) ---
# Stopper le maître :
sudo systemctl stop redis-server   # sur 10.0.5.21
# Observer (quelques secondes) :
redis-cli -p 26379 SENTINEL get-master-addr-by-name mymaster
# → nouvelle IP du maître élu, ex. 10.0.5.22
# Sur l'ancien maître, au redémarrage, il rejoint comme réplica automatiquement.
```

**Côté application** (exemple Python) :

```python
from redis.sentinel import Sentinel
sentinel = Sentinel([("10.0.5.31", 26379), ("10.0.5.32", 26379), ("10.0.5.33", 26379)],
                    socket_timeout=0.5)
master = sentinel.master_for("mymaster", socket_timeout=0.5)  # écritures
replica = sentinel.slave_for("mymaster", socket_timeout=0.5)  # lectures
master.set("k", "v")
```

> **Piège :** Sentinel gère le failover, **pas** le sharding ni la répartition de charge. Pour > RAM d'une machine : Cluster (sections 62–64).

---


## 61. Sentinel : quorum, split-brain et garde-fous

| Sentinels | Quorum conseillé | Pannes tolérées |
|---|---|---|
| 3 | 2 | 1 |
| 5 | 3 | 2 |

Règles :

