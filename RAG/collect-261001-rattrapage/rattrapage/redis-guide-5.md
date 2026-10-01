---
id: collect-261001-rattrapage/rattrapage/redis-guide-5
title: "Guide Redis — De l'installation à la production"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["backlog", "memory"]
source: docs/RAG/collect-261001-rattrapage/redis_guide.md
source_anchor: ""
source_lines: [763, 952]
sha256: 9d03587c4db47754797de7b70a223eac765320f5921166c5921cd7582a8ddab9
---

# --- File d'attente TCP ---
tcp-backlog 511                 # à augmenter (ex. 2048) si "too many connections" sous charge
timeout 0                       # 0 = jamais de déconnexion idle (mettre 300 en prod exposée)
tcp-keepalive 300               # sondes keepalive : détecte les clients morts

# --- Threads E/S (⚠️ 6.0+) ---
io-threads 4                    # 4-8 sur machines multi-cœurs sous forte charge réseau
io-threads-do-reads yes         # ⚠️ 6.0+ : aussi les lectures en threads
```

Réglages **système Linux** indispensables (sans eux, Redis loggue des WARNING au démarrage) :

```bash
# 1. overcommit mémoire (sinon fork() pour BGSAVE peut échouer)
echo "vm.overcommit_memory = 1" | sudo tee /etc/sysctl.d/99-redis.conf

# 2. backlog SOMAXCONN cohérent avec tcp-backlog
echo "net.core.somaxconn = 2048" | sudo tee -a /etc/sysctl.d/99-redis.conf
sudo sysctl --system

# 3. Désactiver les Transparent Huge Pages (pics de latence !)
echo never | sudo tee /sys/kernel/mm/transparent_hugepage/enabled
# persistant via systemd :
sudo tee /etc/systemd/system/disable-thp.service <<'EOF'
[Unit]
Description=Disable Transparent Huge Pages
[Service]
Type=oneshot
ExecStart=/bin/sh -c 'echo never > /sys/kernel/mm/transparent_hugepage/enabled'
[Install]
WantedBy=multi-user.target
EOF
sudo systemctl enable --now disable-thp.service

# 4. Vérification au démarrage (aucun WARNING attendu)
sudo grep -i warning /var/log/redis/redis-server.log | head
```

---

## 27. redis.conf : mémoire (maxmemory et politiques d'éviction)

**Sans `maxmemory`, Redis grossit jusqu'à l'OOM killer.** C'est LE réglage prod n°1.

```ini
maxmemory 4gb                    # ex. 70-80 % de la RAM si Redis est seul sur la machine
maxmemory-policy noeviction      # défaut : renvoie des erreurs à l'écriture quand plein
```

Politiques d'éviction (`maxmemory-policy`) :

| Politique | Comportement | Quand l'utiliser |
|---|---|---|
| `noeviction` | Erreur OOM sur écriture | Données critiques, persistance forte |
| `allkeys-lru` | Évince les clés les moins récemment utilisées | **Cache général (choix par défaut conseillé)** |
| `allkeys-lfu` | Évince les moins fréquemment utilisées (⚠️ 4.0+) | Cache avec fréquences contrastées |
| `volatile-lru` | LRU parmi les clés **avec TTL** | Mix cache + données persistantes |
| `volatile-lfu` | LFU parmi les clés avec TTL | Idem, version LFU |
| `volatile-random` | Aléatoire parmi clés avec TTL | Rarement |
| `allkeys-random` | Aléatoire parmi toutes | Tests uniquement |
| `volatile-ttl` | Les TTL les plus courts d'abord | Files à expiration |

```bash
# Réglage à chaud :
redis-cli CONFIG SET maxmemory 4gb
redis-cli CONFIG SET maxmemory-policy allkeys-lru

# Surveiller la pression mémoire :
redis-cli INFO memory | grep -E "used_memory_human|maxmemory_human|evicted_keys|mem_fragmentation_ratio"
```

> **Dimensionnement :** gardez une marge pour le **fork** (BGSAVE/AOF rewrite duplique transitoirement la mémoire — avec `vm.overcommit_memory=1`, le COW limite le besoin réel, mais prévoyez quand même ~1,5× la donnée en RAM libre sur les pics).
> **Réplicas :** `replica-ignore-maxmemory` (défaut `yes`) — un réplica n'évince pas, il suit le maître.

---

## 28. redis.conf : persistance RDB (snapshots)

Le RDB écrit un **snapshot binaire** complet à intervalles (`dump.rdb`). Compact, rapide à charger, idéal pour les **backups**.

```ini
# --- Déclencheurs : save <secondes> <modifications> ---
save 3600 1        # 1 modif  en 1 h
save 300 100       # 100 modifs en 5 min
save 60 10000      # 10 000 modifs en 1 min
# save ""          # désactive totalement le RDB

dbfilename dump.rdb
dir /var/lib/redis
rdbcompression yes       # compresse (LZF) — laissez yes
rdbchecksum yes          # checksum CRC64 — laissez yes
rdb-del-sync-files no    # ⚠️ 7.0+
```

Commandes associées :

```bash
redis-cli BGSAVE    # snapshot en arrière-plan (fork) → "Background saving started"
redis-cli SAVE      # snapshot BLOQUANT — ne jamais utiliser en prod !
redis-cli LASTSAVE  # timestamp Unix du dernier snapshot réussi
ls -lh /var/lib/redis/dump.rdb
```

Avantages : fichier unique, petit, chargement rapide → parfait pour **backup/restore** et **redémarrages rapides**.
Inconvénients : **perte des écritures** depuis le dernier snapshot (jusqu'à plusieurs minutes). Pour une perte ~0, combinez avec l'AOF (section 30).

---

## 29. redis.conf : persistance AOF (journal)

L'AOF enregistre **chaque écriture** dans un journal (`appendonly.aof`, en pratique un manifeste + fichiers depuis Redis 7 : voir `appenddirname`). Rejeu au redémarrage = données à jour.

```ini
appendonly yes
appendfilename "appendonly.aof"   # ⚠️ 7.0+ : devient un préfixe (multi-part AOF)
appenddirname "appendonlydir"     # ⚠️ 7.0+ : répertoire des segments

# --- Durabilité vs performance ---
appendfsync always    # fsync à chaque écriture : sûr, LENT (~divise le débit par 10)
appendfsync everysec  # fsync 1×/s : compromis standard — perte max ~1 s  ✅ RECOMMANDÉ
appendfsync no        # laisse l'OS décider : rapide, perte jusqu'à ~30 s

# --- Réécriture (compaction) ---
auto-aof-rewrite-percentage 100   # réécrit quand l'AOF a doublé (+100 %)
auto-aof-rewrite-min-size 64mb     # ... et dépasse 64 Mo
aof-rewrite-incremental-fsync yes  # fsync par morceaux pendant le rewrite (évite les pics)

# --- Tolérance à la corruption de queue ---
aof-load-truncated yes   # tronque un AOF à queue corrompue au lieu de refuser de démarrer
```

Commandes :

```bash
redis-cli BGREWRITEAOF   # compaction manuelle en arrière-plan
redis-cli INFO persistence | grep -E "aof_|rdb_"
```

> ⚠️ **7.0+ : Multi-Part AOF.** L'AOF n'est plus un fichier unique mais un répertoire (`appendonlydir/`) avec fichiers base + incrémentaux. Adaptez vos scripts de backup (sauvegardez le **répertoire entier**, pas un fichier).

---

## 30. RDB vs AOF : que choisir ?

| Critère | RDB seul | AOF seul (`everysec`) | RDB + AOF (recommandé) |
|---|---|---|---|
| Perte max de données | Minutes (selon `save`) | ~1 seconde | ~1 seconde |
| Vitesse de redémarrage | ⚡ Très rapide | 🐢 Lent (rejeu) | Rapide (charge le RDB, rejoue l'incrémental) |
| Taille sur disque | Petite | Grande (avant rewrite) | Moyenne |
| Charge CPU/E-S | Faible (snapshots espacés) | Continue modérée | Modérée |
| Backup simple | ✅ Un fichier à copier | ⚠️ Répertoire à copier | ✅ RDB pour backup, AOF pour durabilité |

**Recommandation production : activez les deux.**

```ini
# /etc/redis/conf.d/99-production.conf
save 3600 1
save 300 100
save 60 10000
appendonly yes
appendfsync everysec
```

Cas particuliers :

- **Cache pur reconstituable** (sessions recréables, cache devant SQL) : RDB seul, voire **aucune persistance** + `replicaof` pour la HA. Assumez la perte.
- **File de tâches critique** : AOF `everysec` minimum, `always` si chaque message compte (avec le coût perf).
- **Gros dataset, redémarrages fréquents** : RDB indispensable pour des restarts rapides.

---


## 31. Expiration : EXPIRE, TTL, PEXPIRE

Toute clé peut avoir une **durée de vie**. À expiration, Redis la supprime automatiquement.

```bash
redis-cli SET session:xyz "donnees"
redis-cli EXPIRE session:xyz 1800        # expire dans 30 min → 1
redis-cli TTL session:xyz               # → 1799 (secondes restantes)
# TTL retourne : -2 si la clé n'existe pas, -1 si pas d'expiration

redis-cli PEXPIRE session:xyz 90000     # en millisecondes
redis-cli PTTL session:xyz              # → 89999

redis-cli EXPIREAT session:xyz 1769560800    # timestamp Unix absolu
redis-cli PEXPIREAT session:xyz 1769560800000

redis-cli PERSIST session:xyz           # retire l'expiration → 1 (clé devient persistante)

