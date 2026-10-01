---
id: collect-261001-rattrapage/rattrapage/redis-guide-11
title: "Guide Redis — De l'installation à la production"
domain: rattrapage
role: reference
task: reference
actors: []
dates: ["2026-09-25", "2026-09-26"]
keywords: ["arr"]
source: docs/RAG/collect-261001-rattrapage/redis_guide.md
source_anchor: ""
source_lines: [1941, 2150]
sha256: 89890ad8ede3a44f258012d8431ecc8db6605a3dac8eda029f0a2f7819807b1f
---

# Guide Redis — De l'installation à la production

- **Toujours un nombre impair** de Sentinels, répartis sur **au moins 2–3 zones/racks** (sinon une panne réseau = pas de quorum = pas de failover).
- `sentinel monitor` : le quorum y est déclaré ; il doit être la **majorité**.
- En cas de **partition réseau** (split-brain) : le côté minoritaire ne fait pas de failover (pas de quorum) ; le côté majoritaire élit un nouveau maître. L'ancien maître isolé passe en lecture seule s'il voit qu'il n'a plus de quorum ? Non — Redis seul ne le fait pas ; c'est `min-replicas-to-write` qui protège :

```ini
# Sur chaque nœud Redis : refuse les écritures si moins de N réplicas joignables
min-replicas-to-write 1
min-replicas-max-lag 10   # ... avec un lag <= 10 s
```

Avec `min-replicas-to-write 1`, un maître isolé de ses réplicas **refuse les écritures** → pas de divergence silencieuse. **Recommandé en prod.**

Notifications Sentinel utiles (à brancher sur votre supervision) :

```bash
redis-cli -p 26379 PSUBSCRIBE "+*"
# Événements : +sdown, -sdown, +odown, +try-failover, +failover-end, +switch-master, +reboot...
```

---

## 62. Cluster : concepts (sharding, hash slots)

Le **Cluster** répartit les données sur N maîtres (**sharding**) : chaque clé est assignée à l'un des **16 384 hash slots** via `slot = CRC16(clé) mod 16384`.

```
Clé "user:1042" → CRC16 → slot 7421 → nœud maître B
```

- Chaque maître a 0..N réplicas (failover **interne** au cluster, sans Sentinel).
- Minimum prod : **3 maîtres + 3 réplicas** (6 nœuds). En dessous de 3 maîtres, pas de quorum.
- Les clients « cluster-aware » maintiennent la table des slots et suivent les **redirections** (`-MOVED`, `-ASK`).
- **Hash tags** : `{...}` force le même slot → `user:{1042}:name` et `user:{1042}:email` sont co-localisées. Indispensable pour les opérations multi-clés (MGET, transactions, Lua).

Contraintes à connaître avant d'adopter :

- Une seule base (`SELECT` interdit, seule la base 0).
- Les opérations multi-clés doivent cibler le **même slot** (sinon `CROSSSLOT` error).
- `KEYS`/`SCAN` par nœud (pas de vue globale native).
- Complexité opérationnelle nettement supérieure à Sentinel.

---

## 63. Cluster : mise en place (6 nœuds)

```ini
# --- /etc/redis/conf.d/99-cluster.conf sur CHAQUE nœud ---
port 6379
cluster-enabled yes
cluster-config-file /var/lib/redis/nodes-6379.conf
cluster-node-timeout 5000
appendonly yes
```

```bash
# 1. Démarrer les 6 instances (10.0.5.41 à 10.0.5.46)

# 2. Créer le cluster (3 maîtres, 1 réplica chacun) :
redis-cli --cluster create \
  10.0.5.41:6379 10.0.5.42:6379 10.0.5.43:6379 \
  10.0.5.44:6379 10.0.5.45:6379 10.0.5.46:6379 \
  --cluster-replicas 1 --cluster-yes
# --cluster-yes : ne pas demander confirmation (scriptable)

# 3. Vérifier :
redis-cli --cluster info 10.0.5.41:6379
# cluster_state:ok, cluster_known_nodes:6, cluster_slots_assigned:16384

redis-cli --cluster check 10.0.5.41:6379   # audit : slots, réplicas, clés orphelines

# 4. Tester avec un client cluster-aware :
redis-cli -c -h 10.0.5.41 SET user:{1042}:name "Awa"   # -c = suit les redirections
redis-cli -c -h 10.0.5.41 GET user:{1042}:name

# 5. Resharding (ajouter un nœud / rééquilibrer) :
redis-cli --cluster reshard 10.0.5.41:6379   # interactif
# ou : --cluster rebalance 10.0.5.41:6379 --cluster-use-empty-masters
```

> Le port **16379** (bus cluster, = port+10000) doit être ouvert entre les nœuds, en plus du 6379. Oubli classique des firewalls.

---

## 64. Cluster : clients, redirections et pièges

```bash
# Sans -c, une clé d'un autre slot renvoie une redirection :
redis-cli -h 10.0.5.41 GET user:1042
# (error) MOVED 7421 10.0.5.42:6379
# Avec -c, redis-cli suit automatiquement :
redis-cli -c -h 10.0.5.41 GET user:1042
# -> Redirected to slot [7421] located at 10.0.5.42:6379
# "Awa"
```

Pièges :

| Piège | Symptôme | Remède |
|---|---|---|
| Client non cluster-aware | `MOVED` en erreur | Lib cluster (redis-py `RedisCluster`, Lettuce, ioredis) |
| Multi-clés cross-slot | `CROSSSLOT Keys in request don't hash to the same slot` | Hash tags `{...}` |
| Transactions Lua multi-clés | `CROSSSLOT` | Toutes les clés avec le même tag |
| `SELECT 1` | `ERR SELECT is not allowed in cluster mode` | Une instance par usage |
| Slot non couvert (`cluster_state:fail`) | Écritures refusées | Au moins 1 maître par slot joignable ; `cluster-require-full-coverage no` pour dégrader gracieusement |

```bash
# Voir la répartition des slots :
redis-cli --cluster info 10.0.5.41:6379 | grep slots
redis-cli -c -h 10.0.5.41 CLUSTER SLOTS | head -20
redis-cli -c -h 10.0.5.41 CLUSTER NODES
```

---

## 65. Sauvegarde : stratégie RDB (BGSAVE)

Le RDB est votre **format de backup** : un fichier unique, copiable à chaud.

```bash
# --- Backup manuel ---
redis-cli BGSAVE
# Vérifier la fin :
redis-cli LASTSAVE
tail -5 /var/log/redis/redis-server.log | grep -i "background saving"

# --- Script de backup quotidien (à mettre en cron) ---
sudo tee /usr/local/bin/redis-backup.sh <<'EOF'
#!/bin/bash
set -euo pipefail
SRC="/var/lib/redis/dump.rdb"
DST="/srv/backups/redis/dump-$(date +%F).rdb"
mkdir -p /srv/backups/redis
redis-cli BGSAVE
# attend la fin du BGSAVE (max 120 s) :
for i in $(seq 1 120); do
  [ "$(redis-cli --raw LASTSAVE)" != "$START" ] 2>/dev/null && break
  sleep 1
done
cp -a "$SRC" "$DST"
gzip -9 "$DST"
find /srv/backups/redis -name "dump-*.rdb.gz" -mtime +30 -delete
echo "Backup OK : $DST.gz"
EOF
sudo chmod +x /usr/local/bin/redis-backup.sh
```

> ⚠️ Ne copiez **jamais** `dump.rdb` pendant un `BGSAVE` sans vérifier `LASTSAVE` : vous risqueriez un fichier tronqué. La boucle ci-dessus attend la fin.

---

## 66. Sauvegarde : AOF et réécriture

Avec le **Multi-Part AOF** (7.0+), sauvegardez le **répertoire entier** :

```bash
APPENDDIR="/var/lib/redis/appendonlydir"
ls -lh "$APPENDDIR"
# appendonly.aof.1.base.rdb  appendonly.aof.1.incr.aof  appendonly.aof.manifest

# Backup cohérent : déclencher un BGREWRITEAOF puis copier le répertoire
redis-cli BGREWRITEAOF
# attendre la fin :
redis-cli INFO persistence | grep aof_rewrite_in_progress  # → 0
tar -czf /srv/backups/redis/aof-$(date +%F).tar.gz -C /var/lib/redis appendonlydir
```

Vérifier l'intégrité d'un backup AOF/RDB avant de s'en servir :

```bash
# RDB :
redis-check-rdb /srv/backups/redis/dump-2026-09-26.rdb
# AOF (ancien format mono-fichier) :
redis-check-aof --fix appendonly.aof   # --fix tronque la queue corrompue (destructif !)
```

> **Règle 3-2-1** : 3 copies, 2 supports différents, 1 hors site. Pour Redis : RDB local + copie sur stockage réseau + copie hors site (rsync/rclone vers un bucket).

---

## 67. Restauration : procédure pas à pas

```bash
# === RESTAURATION RDB (instance arrêtée) ===
sudo systemctl stop redis-server
sudo cp /var/lib/redis/dump.rdb /var/lib/redis/dump.rdb.avant-restauration
sudo cp /srv/backups/redis/dump-2026-09-25.rdb /var/lib/redis/dump.rdb
sudo chown redis:redis /var/lib/redis/dump.rdb
sudo systemctl start redis-server
redis-cli DBSIZE        # nombre de clés restaurées
redis-cli INFO keyspace

# === RESTAURATION AOF (7.0+, répertoire) ===
sudo systemctl stop redis-server
sudo mv /var/lib/redis/appendonlydir /var/lib/redis/appendonlydir.avant
sudo tar -xzf /srv/backups/redis/aof-2026-09-25.tar.gz -C /var/lib/redis
sudo chown -R redis:redis /var/lib/redis/appendonlydir
sudo systemctl start redis-server
# Surveillez le log : le rejeu peut prendre du temps sur un gros AOF
tail -f /var/log/redis/redis-server.log | grep -i "aof\|ready"

# === Restauration sur une AUTRE machine (migration) ===
# 1. rsync du dump.rdb, 2. mêmes étapes, 3. vérifiez les ACL/bind avant d'ouvrir
```

**Testez vos restaurations** (au moins trimestriellement) sur une machine isolée : un backup jamais testé n'est pas un backup. Mesurez le **RTO** (temps de restauration) réel.

---

## 68. Monitoring : INFO, les sections à connaître

