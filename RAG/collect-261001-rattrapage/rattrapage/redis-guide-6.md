---
id: collect-261001-rattrapage/rattrapage/redis-guide-6
title: "Guide Redis — De l'installation à la production"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["agent"]
source: docs/RAG/collect-261001-rattrapage/redis_guide.md
source_anchor: ""
source_lines: [953, 1149]
sha256: 50e276c4671738658b207788b7b157ef5c418aa40695c4d6219757cea5228b9d
---

# --- Variantes conditionnelles (⚠️ 7.0+) ---
redis-cli EXPIRE session:xyz 1800 NX    # ne pose le TTL que s'il n'y en a pas
redis-cli EXPIRE session:xyz 1800 XX    # ... que s'il y en a déjà un
redis-cli EXPIRE session:xyz 1800 GT    # ... que si le nouveau TTL est plus grand
redis-cli EXPIRE session:xyz 60 LT      # ... que s'il est plus petit
```

> Depuis Redis 7, `SET` accepte aussi `EXAT`/`PXAT` (timestamps absolus). Et `GETEX`/`GETDEL` combinent lecture + gestion du TTL (section 6).

---

## 32. Expiration : stratégies interne (passive + active)

Redis combine deux mécanismes :

1. **Expiration passive :** à l'accès d'une clé expirée, elle est supprimée et traitée comme absente. Coût nul si la clé n'est jamais relue.
2. **Expiration active :** 10×/seconde (par défaut), Redis échantillonne 20 clés avec TTL par base ; si > 25 % sont expirées, il recommence immédiatement (jusqu'à 25 % du temps CPU). Avec `hz` augmenté, le cycle est plus fréquent.

```ini
hz 10   # défaut : fréquence des tâches de fond (expiration active, rehash...)
# hz 50 à 100 sur serveurs chargés en expirations — au prix d'un peu de CPU
```

```bash
redis-cli INFO stats | grep expired_keys     # clés expirées (passif+actif)
redis-cli INFO stats | grep evicted_keys     # clés évincées par maxmemory (différent !)
```

> **Distinction importante :** `expired_keys` = TTL atteint ; `evicted_keys` = `maxmemory` atteinte. Les deux font baisser le hit rate, mais les causes et remèdes diffèrent.

---

## 33. Expiration : pièges et bonnes pratiques

| Piège | Explication | Remède |
|---|---|---|
| TTL sur clé écrasée | `SET` **sans** option efface le TTL existant | Utilisez `KEEPTTL` (⚠️ 6.0+) ou reposez le TTL |
| `GETSET`/`SETRANGE` | conservent le TTL (ce sont des écritures partielles) | OK, mais à savoir |
| Horloges désynchronisées | `EXPIREAT` avec timestamp absolu + NTP en vrac | chrony partout, préférez les durées relatives |
| Avalanche d'expirations | 1 M de clés avec le même TTL → pic de CPU/suppressions | **Jitter** : TTL aléatoire ±10 % |
| Clés « immortelles » accidentelles | Oubli du TTL sur cache → mémoire qui gonfle | `volatile-lru` + TTL systématique sur le cache |
| TTL trop court | Hit rate effondré, base SQL saturée | Mesurez le hit rate avant de baisser (section 69) |

Pattern du jitter en application (pseudo-code) :

```python
import random
ttl_base = 3600
ttl = ttl_base + random.randint(-300, 300)   # ±5 min
redis.set("cache:produit:42", valeur, ex=ttl)
```

---

## 34. Persistance et expiration : cas limite

- Un **réplica** n'expire pas lui-même : c'est le **maître** qui propage les `DEL` d'expiration. Un réplica isolé (ex. `REPLICAOF NO ONE` après failover) peut donc resservir des clés expirées jusqu'à sa propre expiration active.
- Au **chargement RDB**, les clés déjà expirées sont ignorées (maître) ou conservées-mais-inaccessibles (réplica, expirées au premier accès).
- Dans l'**AOF**, l'expiration est journalisée comme un `DEL` explicite au moment où elle se produit.

---

## 35. Sécurité : AUTH avec requirepass

Le minimum vital dès que Redis écoute autre chose que localhost.

```ini
# redis.conf
requirepass "UnMotDePasseLongEtAleatoireIci"
```

```bash
# Côté client :
redis-cli -a "UnMotDePasseLongEtAleatoireIci" ping
# ou (mieux : n'apparaît pas dans l'historique avec -a... en fait si ; préférez) :
redis-cli
127.0.0.1:6379> AUTH "UnMotDePasseLongEtAleatoireIci"
# OK

# Vérifier qui est authentifié / combien de commandes :
redis-cli INFO clients
```

Limites de `requirepass` :

- Mot de passe **unique** pour tout le monde, transmis en clair (sans TLS).
- Aucune granularité : un client authentifié peut tout faire, y compris `FLUSHALL`.

> **En production moderne : préférez les ACLs (section 36).** `requirepass` reste utile en dev ou comme filet minimal.

Générez des secrets robustes, jamais dans ce guide ni en clair dans un ticket :

```bash
openssl rand -hex 32   # 64 caractères hexadécimaux
```

---

## 36. Sécurité : ACLs (Redis 6+, la bonne pratique)

Les **ACL** (Access Control Lists) donnent des **utilisateurs** avec : mot de passe, commandes autorisées, clés accessibles, canaux pub/sub.

```bash
# --- Création à chaud ---
redis-cli ACL SETUSER appli_cache on ">MotDePasseAppli" \
  +get +set +del +expire +ttl \
  -@all \
  ~cache:* ~session:* \
  &events:*

# Lecture de la règle : utilisateur actif, commandes +get+set+del+expire+ttl,
# clés limitées aux motifs cache:* et session:*, canaux events:*

# --- Vérifications ---
redis-cli ACL LIST
redis-cli ACL WHOAMI
redis-cli ACL GETUSER appli_cache

# --- Test d'interdiction ---
redis-cli --user appli_cache --pass "MotDePasseAppli" FLUSHALL
# (error) NOPERM this user has no permissions to run the 'flushall' command

# --- Utilisateur admin restreint au réseau ---
redis-cli ACL SETUSER admin on ">MotDePasseAdmin" allcommands allkeys

# --- Persister les ACL (sinon perdues au restart !) ---
redis-cli ACL SAVE        # nécessite : aclfile /etc/redis/users.acl dans redis.conf
```

Fichier `/etc/redis/users.acl` (généré par `ACL SAVE`, éditable à la main) :

```ini
user default off                          # désactive l'utilisateur anonyme par défaut !
user admin on >MotDePasseAdmin allcommands allkeys
user appli_cache on >MotDePasseAppli +get +set +del +expire +ttl -@all ~cache:* ~session:*
```

```ini
# redis.conf
aclfile /etc/redis/users.acl
# (ne pas combiner requirepass + aclfile : requirepass définit juste le mdp de "default")
```

Catégories de commandes utiles : `@connection`, `@string`, `@hash`, `@list`, `@set`, `@sortedset`, `@stream`, `@keyspace`, `@admin`, `@dangerous`.

**Recette prod :** `default off` + un user par application avec le minimum de commandes et des motifs de clés stricts. Chaque fuite de credential n'expose alors qu'un périmètre.

---

## 37. Sécurité : bind et protected-mode

```ini
bind 127.0.0.1 ::1 10.0.5.21   # UNIQUEMENT les interfaces nécessaires
protected-mode yes               # garde-fou : refuse le distant si pas d'AUTH/bind restreint
port 6379
```

Matrice de décision :

| Exposition | Configuration |
|---|---|
| Localhost uniquement (dev, agent local) | `bind 127.0.0.1 ::1`, protected-mode yes |
| Réseau privé (VLAN applicatif) | `bind 10.x.x.x`, AUTH/ACL obligatoire, firewall |
| Internet | **INTERDIT** sans TLS + ACL + firewall. Préférez un VPN/bastion. |

Vérifiez l'exposition réelle (pas seulement la config) :

```bash
ss -ltnp | grep 6379
# 127.0.0.1:6379 ✅  vs  0.0.0.0:6379 ⚠️ (toutes interfaces !)

# Depuis un poste externe, testez que c'est fermé :
nmap -p 6379 redis-prod.exemple.sn
sudo iptables -L -n | grep 6379   # ou règles nftables / security groups
```

> Des milliers d'instances Redis exposées sans mot de passe sont régulièrement compromises (ransomware `FLUSHALL` + demande de rançon dans une clé). **Ne soyez pas l'un d'eux.**

---

## 38. Sécurité : TLS (chiffrement en transit)

Redis 6+ supporte le TLS natif (compiler avec `BUILD_TLS=yes`, paquet officiel Debian : inclus).

```bash
# --- 1. Générer une CA et des certificats (exemple autosigné, prod = PKI d'entreprise) ---
mkdir -p /etc/redis/tls && cd /etc/redis/tls
openssl genrsa -out ca.key 4096
openssl req -x509 -new -nodes -key ca.key -sha256 -days 3650 -out ca.crt -subj "/CN=Redis CA"
openssl genrsa -out redis.key 2048
openssl req -new -key redis.key -out redis.csr -subj "/CN=redis-prod.exemple.sn"
openssl x509 -req -in redis.csr -CA ca.crt -CAkey ca.key -CAcreateserial \
  -out redis.crt -days 825 -sha256

# --- 2. redis.conf ---
```

