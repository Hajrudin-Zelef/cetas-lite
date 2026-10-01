---
id: collect-261001-rattrapage/rattrapage/redis-guide-1
title: "Guide Redis — De l'installation à la production"
domain: rattrapage
role: reference
task: reference
actors: []
dates: ["2026-09-26"]
keywords: ["open source"]
source: docs/RAG/collect-261001-rattrapage/redis_guide.md
source_anchor: ""
source_lines: [1, 159]
sha256: 419a032a324797ba48796f575463ddef9c0df3d2adc35febf13a65e05e3bb897
---

# Guide Redis — De l'installation à la production

> **Public :** sysadmins, chefs de service systèmes, DevOps.
> **Version couverte :** Redis 7.x (7.0 → 7.4). Les commandes marquées ⚠️ existent depuis une version précise, indiquée à chaque fois.
> **Angle :** pratique et production. Chaque section donne des commandes `redis-cli` testables, des tableaux de référence et des checklists opérationnelles.
> **Avertissement :** les mots de passe, adresses IP et noms d'hôtes de ce guide sont fictifs. Adaptez-les à votre infrastructure.

---

## 1. Qu'est-ce que Redis ?

Redis (**RE**mote **DI**ctionary **S**erver) est un **magasin de données en mémoire**, structuré en **clé-valeur**, qui persiste optionnellement sur disque. Il est écrit en C, open source (licence RSALv2/SSPLv1 depuis la 7.4 — point juridique à vérifier avec votre DSI avant déploiement en 7.4+ ; la 7.2 reste la dernière sous BSD).

Pourquoi il est partout en production :

| Caractéristique | Détail |
|---|---|
| Vitesse | ~100 000+ op/s sur une machine modeste, latence sub-milliseconde |
| Structures riches | Strings, Hashes, Lists, Sets, Sorted Sets, Streams, Bitmaps, HyperLogLog, Geo, JSON (module) |
| Persistance | RDB (snapshot) et/ou AOF (journal), combinables |
| Haute dispo | Réplication, Sentinel (failover auto), Cluster (sharding) |
| Atomicité | Commandes atomiques, transactions MULTI/EXEC, scripts Lua |
| Écosystème | Client dans tous les langages, exporters Prometheus, intégration K8s |

Cas d'usage typiques en entreprise : **cache applicatif**, **sessions**, **files d'attente**, **compteurs temps réel**, **rate limiting**, **leaderboards**, **pub/sub** pour événements, **verrous distribués**.

---

## 2. Redis dans un système d'information : où le placer ?

```
                    ┌──────────────┐
                    │  Application │
                    │  (PHP/Python │
                    │   /Java/Go)  │
                    └──────┬───────┘
                           │
              ┌────────────┼────────────┐
              ▼            ▼            ▼
       ┌──────────┐ ┌───────────┐ ┌──────────┐
       │  Redis   │ │ PostgreSQL│ │  Stockage│
       │ (cache,  │ │ / MySQL   │ │  fichiers│
       │ sessions,│ │ (données  │ │          │
       │ files…)  │ │  durables)│ │          │
       └──────────┘ └───────────┘ └──────────┘
```

Règle d'or : **Redis n'est pas votre seule copie des données critiques** (sauf usage assumé type cache volatile). La base relationnelle reste la source de vérité ; Redis accélère et découple.

Topologies courantes :

| Topologie | Usage | Complexité |
|---|---|---|
| Instance unique | Dev, petits caches non critiques | Faible |
| Maître + réplica(s) | Lecture répartie, backup à chaud | Moyenne |
| Sentinel (3+ nœuds) | Failover automatique | Moyenne |
| Cluster (6+ nœuds) | Sharding, > RAM d'une machine | Élevée |

---

## 3. Concepts fondamentaux

### 3.1 Clé-valeur, mais pas que

Chaque donnée est une **clé** (chaîne binaire, conventionnellement `objet:id:champ`, ex. `user:1042:name`) associée à une **valeur typée**. Contrairement à Memcached (que des strings), Redis comprend la structure de la valeur et propose des opérations dessus (incrémenter un compteur, pousser dans une liste, etc.).

### 3.2 En mémoire, avec persistance optionnelle

- Les lectures/écritures se font en **RAM** → vitesse maximale.
- La persistance (RDB/AOF, sections 28–30) permet de survivre à un redémarrage. **Sans persistance, un reboot = perte totale.** C'est le piège n°1 des débutants.

### 3.3 Single-threaded (presque)

Le traitement des commandes est **mono-thread** : chaque commande s'exécute atomiquement, sans verrou. Conséquences pratiques :

- ✅ Pas de race condition sur une commande simple (`INCR` est atomique).
- ⚠️ Une commande lente (ex. `KEYS *` sur des millions de clés) **bloque tout le serveur**. D'où l'interdiction de `KEYS` en prod (utiliser `SCAN`, section 91).
- Depuis Redis 6, les **E/S réseau sont multi-threadées** (`io-threads`), mais l'exécution reste mono-thread.

### 3.4 Bases de données numérotées

Par défaut 16 bases (`0` à `15`), sélection via `SELECT 2`. **En pratique :** préférez **une instance par usage** plutôt que plusieurs bases — le Cluster ne supporte que la base 0, et `FLUSHDB`/`FLUSHALL` sont sources d'incidents.

```bash
redis-cli -n 2 SET ma_cle "valeur"   # -n = numéro de base
```

---

## 4. Les structures de données en un coup d'œil

| Type | Contenu | Cas d'usage typique | Commandes clés |
|---|---|---|---|
| String | Texte / nombre / binaire (≤ 512 Mo) | Cache, compteurs, sessions, verrous | SET, GET, INCR, SETEX |
| Hash | Champ → valeur (objet) | Fiche utilisateur, produit | HSET, HGET, HGETALL |
| List | Liste ordonnée d'insertion | File d'attente, historique | LPUSH, RPOP, LRANGE |
| Set | Ensemble unique non ordonné | Tags, abonnés, votes | SADD, SMEMBERS, SINTER |
| Sorted Set | Ensemble + score (trié) | Classements, planning | ZADD, ZRANGE, ZRANK |
| Stream | Journal append-only | Event sourcing, files robustes | XADD, XREAD, XACK |
| Bitmap | Bits (⚠️ 7.x via SETBIT) | Présence/absence, stats | SETBIT, GETBIT, BITCOUNT |
| HyperLogLog | Compteur approximatif unique | Visiteurs uniques | PFADD, PFCOUNT |
| Geo | Coordonnées géographiques | Points d'intérêt proches | GEOADD, GEORADIUS (⚠️ 6.2+: GEOSEARCH) |

> Les sections 5 à 20 détaillent chaque type avec des exemples `redis-cli` complets.

---

## 5. Strings : concepts

Le type le plus simple et le plus utilisé : une clé → une valeur (chaîne, entier, ou binaire). Limite : **512 Mo** par valeur.

Idées reçues à corriger :

- Un entier stocké en string (`"1042"`) peut être **incrémenté atomiquement** avec `INCR` — Redis convertit à la volée.
- Les strings servent aussi à stocker du **binaire** (images miniatures, blobs sérialisés), mais préférez un stockage objet (S3/NFS) au-delà de quelques Mo.
- `SET` avec options `NX`/`XX`/`EX`/`PX` remplace avantageusement `SETNX`+`EXPIRE` (atomique, voir section 46).

---

## 6. Strings : commandes essentielles

```bash
# --- Bases ---
redis-cli SET user:1042:name "Awa Diallo"
redis-cli GET user:1042:name
# "Awa Diallo"

redis-cli MSET user:1042:ville "Dakar" user:1042:role "admin"
redis-cli MGET user:1042:ville user:1042:role
# 1) "Dakar"
# 2) "admin"

# --- Expiration à la création (atomique) ---
redis-cli SET session:abc123 "donnees..." EX 3600   # expire dans 1 h
redis-cli SET session:abc123 "donnees..." PX 60000  # expire dans 60 s (ms)

# --- Ne créer que si absent (verrous, voir section 46) ---
redis-cli SET verrou:job42 "worker-1" NX PX 30000
# (nil) si la clé existe déjà, "OK" sinon

# --- Compteurs atomiques ---
redis-cli SET compteur:visites 0
redis-cli INCR compteur:visites        # → 1
redis-cli INCRBY compteur:visites 10   # → 11
redis-cli DECR compteur:visites        # → 10
redis-cli INCRBYFLOAT prix:total 19.99  # flottants

# --- Manipulation de chaînes ---
redis-cli APPEND log:2026-09-26 "ligne1\n"   # concatène, retourne la taille
redis-cli STRLEN user:1042:name              # longueur
redis-cli GETRANGE user:1042:name 0 2         # "Awa" (substring)
redis-cli SETRANGE user:1042:name 4 "DIALLO"  # écrase à partir de l'offset

# --- Get-and-set atomique ---
redis-cli GETSET compteur:visites 0   # retourne l'ancienne valeur, met 0
redis-cli GETDEL session:abc123       # ⚠️ 6.2+ : lit puis supprime
redis-cli GETEX session:abc123 EX 3600  # ⚠️ 6.2+ : lit et prolonge le TTL
```

