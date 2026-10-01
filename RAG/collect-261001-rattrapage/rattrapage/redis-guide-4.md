---
id: collect-261001-rattrapage/rattrapage/redis-guide-4
title: "Guide Redis — De l'installation à la production"
domain: rattrapage
role: reference
task: reference
actors: []
dates: ["2026-09-26"]
keywords: ["memory"]
source: docs/RAG/collect-261001-rattrapage/redis_guide.md
source_anchor: ""
source_lines: [535, 762]
sha256: 75bc058c4531b93e1b3b192489517d6b757e7329fb0d1e7e04d2d5c54207edca
---

# --- Infos ---
redis-cli GEOPOS agences:sn "agence:dakar"    # coordonnées stockées
redis-cli GEOHASH agences:sn "agence:dakar"   # geohash (interopérabilité)
```

Cas d'usage : agences les plus proches, livreurs dans un rayon, géofencing simple. Pour des polygones complexes ou des millions de points avec requêtes spatiales avancées, préférez PostGIS.

---

## 20. Modules : RedisJSON, RediSearch, RedisTimeSeries (introduction)

Depuis Redis 7, les modules les plus utiles sont regroupés dans **Redis Stack** (même binaire). Sur Debian/Ubuntu, le paquet `redis-stack-server` les inclut.

| Module | Apporte | Exemple |
|---|---|---|
| RedisJSON | Stocker/manipuler du JSON natif (type `JSON`) | `JSON.SET doc:1 $ '{"nom":"Awa"}'` |
| RediSearch | Indexation plein texte + requêtes | `FT.CREATE idx ON JSON ...`, `FT.SEARCH` |
| RedisTimeSeries | Séries temporelles (downsampling) | `TS.ADD sonde:1 * 23.5` |
| RedisBloom | Filtres probabilistes | `BF.ADD` |

```bash
# Exemple RedisJSON (si module chargé)
redis-cli JSON.SET produit:42 $ '{"nom":"Onduleur 3kVA","prix":450000,"stock":12}'
redis-cli JSON.GET produit:42 $.prix
redis-cli JSON.NUMINCRBY produit:42 $.stock -1
```

> **Note production :** chaque module augmente la surface mémoire/CPU et la complexité. N'activez que ce dont vous avez besoin (`loadmodule` sélectif dans redis.conf). Le Cluster impose des contraintes sur les clés multi-champs (hash tags, section 62).

---

## 21. Installation sur Debian/Ubuntu (méthode recommandée)

N'utilisez **pas** le paquet `redis-server` des dépôts Debian/Ubuntu standards : il est souvent **plusieurs versions en retard** (ex. Redis 6 sur d'anciennes LTS). Utilisez le **dépôt officiel Redis**.

```bash
# --- 1. Prérequis ---
sudo apt update
sudo apt install -y curl gnupg lsb-release

# --- 2. Clé GPG et dépôt officiel ---
curl -fsSL https://packages.redis.io/gpg | sudo gpg --dearmor -o /usr/share/keyrings/redis-archive-keyring.gpg
echo "deb [signed-by=/usr/share/keyrings/redis-archive-keyring.gpg] https://packages.redis.io/deb $(lsb_release -cs) main" \
  | sudo tee /etc/apt/sources.list.d/redis.list

# --- 3. Installation ---
sudo apt update
sudo apt install -y redis-server

# --- 4. Vérification ---
redis-server --version
# Redis server v=7.4.0 sha=... malloc=jemalloc-5.3.0 bits=64
redis-cli --version

# --- 5. Service systemd ---
sudo systemctl enable --now redis-server
sudo systemctl status redis-server --no-pager
redis-cli ping
# PONG
```

Le paquet officiel fournit :

- `/etc/redis/redis.conf` — configuration
- `/var/lib/redis/dump.rdb` — snapshot par défaut
- `/var/log/redis/redis-server.log` — logs
- Service `redis-server` supervisé par systemd

> **Épingler la version** en production (évite une montée majeure surprise) :
> ```bash
> sudo apt-mark hold redis-server
> # Pour mettre à jour volontairement : sudo apt-mark unhold redis-server && sudo apt install redis-server
> ```

---

## 22. Installation depuis les sources (compilation)

Utile pour une version précise, un patch, ou une architecture exotique.

```bash
# --- Dépendances de build ---
sudo apt update
sudo apt install -y build-essential tcl pkg-config

# --- Téléchargement et compilation ---
cd /tmp
curl -fsSLO https://download.redis.io/releases/redis-7.4.0.tar.gz
tar xzf redis-7.4.0.tar.gz
cd redis-7.4.0

# Build avec jemalloc (allocateur recommandé, défaut) et TLS :
make -j"$(nproc)" BUILD_TLS=yes

# --- Tests (optionnel mais recommandé, ~quelques minutes) ---
make test
# \o/ All tests passed without errors!

# --- Installation ---
sudo make install PREFIX=/usr/local
# binaires : /usr/local/bin/redis-server, redis-cli, redis-sentinel...

# --- Vérification TLS compilé ---
redis-server --help | grep -i tls || redis-cli --tls --help | head -1
```

> Créez ensuite l'utilisateur système, les répertoires et le service systemd manuellement (reprenez les fichiers du paquet `.deb` comme modèle : `/lib/systemd/system/redis-server.service`).

---

## 23. Installation via Docker (dev / tests)

```bash
# --- Instance éphémère pour tester ---
docker run -d --name redis-dev -p 6379:6379 redis:7.4-alpine redis-server --appendonly yes

# --- Avec persistance et config perso ---
mkdir -p /srv/redis/data
cat > /srv/redis/redis.conf <<'EOF'
bind 0.0.0.0
protected-mode yes
port 6379
requirepass CHANGEZ_MOI
appendonly yes
maxmemory 512mb
maxmemory-policy allkeys-lru
EOF

docker run -d --name redis-prod \
  -p 6379:6379 \
  -v /srv/redis/data:/data \
  -v /srv/redis/redis.conf:/usr/local/etc/redis/redis.conf \
  --restart unless-stopped \
  redis:7.4-alpine /usr/local/etc/redis/redis.conf

redis-cli -a CHANGEZ_MOI ping
```

> ⚠️ En production « vraie », préférez le paquet système ou une installation dédiée : Docker ajoute une couche réseau et complique le tuning (THP, overcommit, voir section 26). Pour du dev et des tests d'intégration, c'est parfait.

---

## 24. Premier démarrage et vérification

Checklist post-installation (à faire **avant** d'ouvrir au réseau) :

```bash
# 1. Le service tourne ?
sudo systemctl is-active redis-server   # → active
ss -ltnp | grep 6379                     # écoute ?

# 2. Répond-il ?
redis-cli ping   # → PONG

# 3. Quelle version / config ?
redis-cli INFO server | grep -E "redis_version|process_id|uptime_in_days"
redis-cli CONFIG GET maxmemory
redis-cli CONFIG GET appendonly

# 4. Persistance active ?
ls -lh /var/lib/redis/   # dump.rdb présent après un BGSAVE

# 5. Test d'écriture/lecture + expiration ---
redis-cli SET test:demarrage "ok" EX 60
redis-cli GET test:demarrage
redis-cli TTL test:demarrage

# 6. Nettoyage du test ---
redis-cli DEL test:demarrage
```

Si `redis-cli ping` répond `PONG` mais que vous prévoyez d'exposer Redis au réseau : **ne passez pas à la suite sans les sections 35–40 (sécurité)**.

---

## 25. redis.conf : organisation du fichier

Le fichier `/etc/redis/redis.conf` fait ~1 400 lignes dont 90 % de commentaires. Structure :

```
INCLUDES        → inclure d'autres fichiers (ex. /etc/redis/conf.d/*.conf)
MODULES         → loadmodule ...
NETWORK         → bind, port, protected-mode, timeout...
GENERAL         → daemonize, supervised, logfile, databases...
SNAPSHOTTING    → save, rdbcompression...
REPLICATION     → replicaof, ...
SECURITY        → requirepass, aclfile...
CLIENTS         → maxclients
MEMORY MGMT     → maxmemory, maxmemory-policy, ...
APPEND ONLY     → appendonly, appendfsync...
...
```

**Bonne pratique :** ne modifiez pas le fichier d'origine à la main en vrac. Créez un fichier d'override :

```bash
sudo mkdir -p /etc/redis/conf.d
sudo tee /etc/redis/conf.d/99-production.conf <<'EOF'
# Overrides production — Zelef, 2026-09-26
maxmemory 4gb
maxmemory-policy allkeys-lru
appendonly yes
EOF
echo "include /etc/redis/conf.d/99-production.conf" | sudo tee -a /etc/redis/redis.conf
sudo systemctl restart redis-server
```

Et pour un réglage à chaud (sans restart), `CONFIG SET` (persisté via `CONFIG REWRITE`) :

```bash
redis-cli CONFIG SET maxmemory 4gb
redis-cli CONFIG SET maxmemory-policy allkeys-lru
redis-cli CONFIG REWRITE   # réécrit redis.conf avec les valeurs courantes
```

> ⚠️ `CONFIG REWRITE` réécrit **tout** le fichier : il perd vos commentaires d'origine. Sauvegardez `redis.conf` avant (ex. `cp redis.conf redis.conf.avant-rewrite`).

---

## 26. redis.conf : réseau et système (bind, port, protected-mode)

```ini
# --- Écoute ---
bind 127.0.0.1 ::1              # défaut : localhost uniquement. En prod : IP(s) privées
# bind 10.0.5.21 127.0.0.1     # exemple : interface privée + loopback
port 6379
protected-mode yes              # refuse les connexions distantes sans AUTH ni bind restreint

