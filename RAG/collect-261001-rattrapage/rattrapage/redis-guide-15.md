---
id: collect-261001-rattrapage/rattrapage/redis-guide-15
title: "Guide Redis — De l'installation à la production"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["backlog", "latency"]
source: docs/RAG/collect-261001-rattrapage/redis_guide.md
source_anchor: ""
source_lines: [2734, 2947]
sha256: d10b4720050f6ab6e7d2dbd133c7f5a83eac90c5de2bb7cf18911a49d3767697
---

# Guide Redis — De l'installation à la production

```bash
# ===== SESSION (cf. section 51) =====
TOKEN=$(openssl rand -hex 32)
redis-cli HSET session:$TOKEN user_id 1042 role client
redis-cli EXPIRE session:$TOKEN 1800

# ===== PANIER (hash, 1 champ par ligne) =====
# panier:<user_id> → { sku → quantité }
redis-cli HSET panier:1042 sku-7788 2 sku-9911 1
redis-cli HINCRBY panier:1042 sku-7788 1      # +1 article
redis-cli HINCRBY panier:1042 sku-9911 -1     # -1 article
redis-cli HDEL panier:1042 sku-9911          # retire la ligne (quantité 0 → HDEL)
redis-cli HGETALL panier:1042
redis-cli EXPIRE panier:1042 2592000         # panier conservé 30 jours

# ===== Validation commande : transaction "panier → commande" =====
# 1. Lire le panier (HGETALL)
# 2. INSERT en base SQL (source de vérité)
# 3. Si OK : DEL panier:1042  (en Lua avec vérification, ou MULTI)
redis-cli MULTI
redis-cli DEL panier:1042
redis-cli SADD commandes:user:1042 cmd-2026-0007788
redis-cli EXEC

# ===== Abandon de panier (marketing) : keyspace notifications (section 45) =====
# notify-keyspace-events Ex ; PSUBSCRIBE __keyevent@0__:expired
# → à l'expiration d'une clé panier:*, déclencher l'email de relance
# (avec le filet anti-manqué : job quotidien qui scanne aussi)
```

---

## 89. Mise à jour de Redis : procédure sans coupure

Principe : **réplica d'abord**, bascule, puis l'ancien maître. Jamais de `apt upgrade` aveugle sur le maître en pleine journée.

```bash
# === 1. Préparation ===
# - Lire les release notes (breaking changes, ex. 6→7 : fonctions, ACL par défaut...)
# - Snapshot de sauvegarde vérifié (section 65)
# - Fenêtre de maintenance annoncée

# === 2. Mettre à jour les RÉPLICAS un par un ===
sudo apt-mark unhold redis-server
sudo apt update && sudo apt install -y redis-server=7:7.4.0-1rl1~bookworm1  # version précise
redis-server --version
sudo systemctl restart redis-server
redis-cli INFO replication | grep master_link_status   # → up
redis-cli INFO server | grep redis_version

# === 3. Bascule (avec Sentinel : failover contrôlé ; sans : manuel) ===
# Avec Sentinel :
redis-cli -p 26379 SENTINEL FAILOVER mymaster
# Sans Sentinel : REPLICAOF NO ONE sur le réplica à jour, repointer l'appli

# === 4. Mettre à jour l'ancien maître (devenu réplica) ===
# même procédure apt qu'à l'étape 2

# === 5. Vérifications post-upgrade ===
redis-cli INFO server | grep redis_version     # partout pareil
redis-cli INFO replication                     # topologie saine
# rejouer les tests applicatifs critiques (connexion, lecture/écriture, TTL)
```

> **Compatibilité ascendante des fichiers** : un RDB/AOF écrit par une 7.x ne se relit pas toujours sur une 6.x → **pas de downgrade** sans backup préalable de l'ancienne version. Testez le rollback en labo avant.

---

## 90. Migration : changer de machine / de topologie

```bash
# === MÉTHODE 1 : réplication (recommandée, quasi sans coupure) ===
# 1. Nouvelle machine : installer Redis, même version, même config (sauf replicaof)
# 2. En faire un réplica de l'ancien maître :
redis-cli -h <nouvelle> REPLICAOF <ancien-maitre> 6379
# 3. Attendre master_link_status:up + offset à jour (INFO replication)
# 4. Bascule : REPLICAOF NO ONE sur la nouvelle, repointer appli/DNS
# 5. Éteindre l'ancienne après une période d'observation

# === MÉTHODE 2 : RDB (coupure courte) ===
# 1. BGSAVE sur l'ancien, rsync du dump.rdb
# 2. Stopper les écritures applicatives, dernier BGSAVE, rsync final
# 3. Démarrer le nouveau avec le dump.rdb, repointer l'appli

# === MÉTHODE 3 : MIGRATE (clé par clé, en ligne) ---
redis-cli -h <ancien> MIGRATE <nouvelle> 6379 "" 0 5000 KEYS cache:produit:42
# "" = pas d'AUTH sur la cible (ou mettre le mot de passe) ; 0 = base ; 5000 = timeout ms
# ⚠️ MIGRATE est bloquant par clé : à éviter sur des millions de clés chaudes.

# === Vérification post-migration ===
# Comparer les nombres de clés :
redis-cli -h <ancien> DBSIZE
redis-cli -h <nouveau> DBSIZE
# Échantillonner : pour 100 clés au hasard, comparer les valeurs (script).
```

---

## 91. redis-cli : les options à connaître

```bash
# --- Connexion ---
redis-cli -h 10.0.5.21 -p 6379 -a "$REDIS_PASS"       # hôte/port/auth
redis-cli -n 2                                         # base n°2
redis-cli -u redis://appli_cache:Secret@10.0.5.21:6379/0  # URI
redis-cli -s /var/run/redis/redis-server.sock          # socket Unix (plus rapide en local)
redis-cli --tls --cacert /etc/redis/tls/ca.crt         # TLS
redis-cli --user appli_cache --pass "$PASS"            # ACL (⚠️ 6.0+)
redis-cli --intrinsic-latency 100                       # mesure la latence intrinsèque (100 s)

# --- Sortie exploitable en script ---
redis-cli --raw GET ma_cle            # sans guillemets ni numéros
redis-cli --csv HGETALL user:1042     # format CSV

# --- Scan sans bloquer (JAMAIS keys *) ---
redis-cli --scan --pattern "cache:*"                 # itère toutes les clés
redis-cli --scan --pattern "session:*" | wc -l       # compte
redis-cli --scan --pattern "tmp:*" | xargs -L 1000 redis-cli DEL   # purge par lot

# --- Piping / mass insert ---
cat data.txt | redis-cli --pipe          # protocole brut, ultra-rapide
# data.txt au format : *3\r\n$3\r\nSET\r\n$3\r\nkey\r\n$5\r\nvalue\r\n

# --- Monitoring temps réel ---
redis-cli MONITOR          # ⚠️ TOUTES les commandes en temps réel (divise le débit par ~2 !)
# À n'utiliser que quelques secondes en debug, jamais en continu.

redis-cli --stat           # stats temps réel (ops/s, mémoire, clients)
redis-cli --latency        # latence min/max/avg en continu
redis-cli --bigkeys        # échantillonne les plus grosses clés
redis-cli --hotkeys        # ⚠️ outils récents : clés les plus accédées

# --- Cluster ---
redis-cli -c -h 10.0.5.41              # mode cluster (suit les redirections)
redis-cli --cluster check 10.0.5.41:6379
redis-cli --cluster info 10.0.5.41:6379

# --- Divers ---
redis-cli --rdb /tmp/dump-copy.rdb     # télécharge un RDB depuis le serveur distant
redis-cli -x SET ma_cle < fichier.bin  # -x : lit la valeur depuis stdin
```

---

## 92. redis.conf : modèle production complet (annoté)

```ini
# ============================================================
# redis.conf — modèle production Redis 7.x
# Machine : 8 Go RAM dédiée — adaptez maxmemory et les chemins
# ============================================================

# --- RÉSEAU ---
bind 10.0.5.21 127.0.0.1
protected-mode yes
port 6379
tcp-backlog 2048
timeout 300
tcp-keepalive 300

# --- GÉNÉRAL ---
supervised systemd
daemonize no
loglevel notice
logfile /var/log/redis/redis-server.log
databases 16
# include /etc/redis/conf.d/*.conf   # (à ajouter si vous utilisez des overrides)

# --- SNAPSHOTS RDB ---
save 3600 1
save 300 100
save 60 10000
dbfilename dump.rdb
dir /var/lib/redis
rdbcompression yes
rdbchecksum yes

# --- RÉPLICATION ---
# replicaof 10.0.5.20 6379        # décommenter sur les réplicas
replica-read-only yes
repl-diskless-sync yes
repl-backlog-size 64mb
min-replicas-to-write 1
min-replicas-max-lag 10

# --- SÉCURITÉ ---
# requirepass "..."               # ou aclfile ci-dessous (recommandé)
aclfile /etc/redis/users.acl
rename-command FLUSHALL ""
rename-command FLUSHDB ""
rename-command DEBUG ""

# --- LIMITES ---
maxclients 10000

# --- MÉMOIRE ---
maxmemory 6gb
maxmemory-policy allkeys-lru
maxmemory-eviction-tenacity 10    # ⚠️ 7.0+ : agressivité d'éviction (10 = défaut)
maxmemory-samples 5

# --- AOF ---
appendonly yes
appendfsync everysec
auto-aof-rewrite-percentage 100
auto-aof-rewrite-min-size 64mb
aof-rewrite-incremental-fsync yes
aof-load-truncated yes

# --- SLOWLOG / LATENCE ---
slowlog-log-slower-than 2000
slowlog-max-len 256
latency-tracking yes

