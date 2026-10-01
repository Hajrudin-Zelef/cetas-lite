---
id: collect-261001-rattrapage/rattrapage/redis-guide-3
title: "Guide Redis — De l'installation à la production"
domain: rattrapage
role: reference
task: reference
actors: []
dates: ["2026-09-25", "2026-09-26"]
keywords: ["consumer", "leaderboard"]
source: docs/RAG/collect-261001-rattrapage/redis_guide.md
source_anchor: ""
source_lines: [360, 534]
sha256: d0c15490a6d213661ee5ab8550eaf1fb39a1a453482f2b862cfa070082579e79
---

# Guide Redis — De l'installation à la production

```bash
# --- Classement ---
redis-cli ZADD leaderboard 1500 "awa" 2300 "moussa" 1800 "fatou"
redis-cli ZRANGE leaderboard 0 -1 WITHSCORES
# 1) "awa" 2) "1500" 3) "fatou" 4) "1800" 5) "moussa" 6) "2300"

# ⚠️ 6.2+ : ZRANGE remplace ZREVRANGE (option REV)
redis-cli ZRANGE leaderboard 0 2 REV WITHSCORES   # top 3 décroissant

redis-cli ZRANK leaderboard "fatou"       # rang croissant → 1
redis-cli ZREVRANK leaderboard "fatou"   # rang décroissant → 1
redis-cli ZSCORE leaderboard "moussa"    # "2300"
redis-cli ZCARD leaderboard              # → 3

# --- Mise à jour / incrément ---
redis-cli ZINCRBY leaderboard 100 "awa"  # awa passe à 1600
redis-cli ZADD leaderboard XX 2000 "awa" # XX : ne met à jour que si existe
redis-cli ZADD leaderboard NX 100 "ibrahim"  # NX : n'ajoute que si absent

# --- Requêtes par score (planning / fenêtre glissante) ---
redis-cli ZADD taches:due 1769557200 "job:1" 1769557300 "job:2"
redis-cli ZRANGEBYSCORE taches:due 0 1769557250        # tâches dues
# ⚠️ 6.2+ : préférer ZRANGE ... BYSCORE :
redis-cli ZRANGE taches:due 0 1769557250 BYSCORE

# --- Requêtes par membre (ordre lexicographique, ⚠️ scores égaux) ---
redis-cli ZADD idx 0 "aaa" 0 "aab" 0 "abc"
redis-cli ZRANGE idx "[aab" "[abc" BYLEX   # → aab, abc

# --- Suppressions ---
redis-cli ZREM leaderboard "awa"
redis-cli ZREMRANGEBYRANK leaderboard 0 9      # supprime les 10 plus bas
redis-cli ZREMRANGEBYSCORE taches:due 0 1769557250  # purge les tâches dues traitées
redis-cli ZPOPMIN taches:due 1   # ⚠️ 5.0+ : dépile le plus petit score (file à priorité !)
redis-cli ZPOPMAX leaderboard 3  # top 3 dépilés

# --- Comptage dans un intervalle ---
redis-cli ZCOUNT leaderboard 1000 2000
```

---

## 15. Streams : concepts

Les Streams (⚠️ Redis 5.0+) sont des **journaux append-only** : chaque entrée a un **ID** (`timestamp-séquence`, ex. `1769557200000-0`, ou `*` pour auto) et des champs clé-valeur. C'est la structure la plus robuste pour les **files de messages** et l'**event sourcing**.

Concepts clés :

- **Consumer groups** : plusieurs consommateurs se partagent un stream ; chaque message est distribué à **un seul** consommateur du groupe.
- **Pending entries** : messages lus mais non acquittés (`XACK`). `XPENDING`/`XAUTOCLAIM` permettent de **récupérer** les messages d'un worker mort — fiabilité impossible avec les simples lists.
- `XREAD BLOCK` : lecture bloquante comme `BLPOP`, mais multi-streams.
- `MAXLEN ~ N` : borne approximative de la taille (le `~` évite un trim exact coûteux).

> Pour une file simple sans besoin de rejeu ni de groupes : les **lists** suffisent (section 48). Pour du robuste : **streams** (section 49).

---


## 16. Streams : commandes essentielles

```bash
# --- Production de messages ---
redis-cli XADD events:commandes * user 1042 action "panier_valide" montant 149.90
# → "1769557200000-0" (ID auto)

redis-cli XADD events:commandes MAXLEN ~ 10000 * user 1043 action "panier_valide"
# borne la taille à ~10 000 entrées (trim approximatif, peu coûteux)

# --- Lecture simple ---
redis-cli XLEN events:commandes
redis-cli XRANGE events:commandes - +            # tout (borne - / +)
redis-cli XRANGE events:commandes - + COUNT 10   # les 10 premiers
redis-cli XREVRANGE events:commandes + - COUNT 5  # les 5 derniers
redis-cli XREAD COUNT 10 STREAMS events:commandes 0   # depuis le début
redis-cli XREAD BLOCK 5000 STREAMS events:commandes $  # bloque 5 s, $ = nouveaux seulement

# --- Consumer groups (le vrai mode "file robuste") ---
redis-cli XGROUP CREATE events:commandes grp:facturation 0 MKSTREAM
# 0 = depuis le début ; $ = seulement les nouveaux ; MKSTREAM crée le stream s'il n'existe pas

redis-cli XREADGROUP GROUP grp:facturation worker-1 COUNT 5 STREAMS events:commandes ">"
# ">" = messages jamais distribués. Retourne les entrées à traiter.

# --- Acquittement ---
redis-cli XACK events:commandes grp:facturation 1769557200000-0
# Sans XACK, le message reste "pending" (non traité) — voir récupération ci-dessous.

# --- Surveillance des pending ---
redis-cli XPENDING events:commandes grp:facturation
# résumé : nb pending, plus petit/plus grand ID, par consommateur

redis-cli XPENDING events:commandes grp:facturation - + 10 worker-1
# détail des 10 pending du worker-1 (idle, retry count...)

# --- Récupérer les messages d'un worker mort (⚠️ 7.0+ : XAUTOCLAIM) ---
redis-cli XAUTOCLAIM events:commandes grp:facturation worker-2 60000 0 COUNT 10
# réclame les messages idle depuis > 60 s, à partir du début

# --- Infos ---
redis-cli XINFO STREAM events:commandes
redis-cli XINFO GROUPS events:commandes
redis-cli XINFO CONSUMERS events:commandes grp:facturation

# --- Suppression / trim ---
redis-cli XDEL events:commandes 1769557200000-0   # supprime une entrée (ne libère pas toujours la mémoire de suite)
redis-cli XTRIM events:commandes MAXLEN ~ 5000    # trim manuel
```

Cycle de vie robuste d'un worker : `XREADGROUP` → traitement → `XACK`. Un superviseur périodique fait `XAUTOCLAIM` sur les idle > N secondes. Voir le cas pratique section 86.

---

## 17. Bitmaps : introduction

Un bitmap manipule des **bits individuels** dans une string. Idéal pour des **statistiques de présence** : « l'utilisateur 1042 s'est-il connecté le 2026-09-26 ? » → 1 bit par utilisateur et par jour.

```bash
# --- Marquer / lire ---
redis-cli SETBIT presence:2026-09-26 1042 1   # user 1042 présent
redis-cli SETBIT presence:2026-09-26 2048 1
redis-cli GETBIT presence:2026-09-26 1042     # → 1
redis-cli GETBIT presence:2026-09-26 9999     # → 0 (jamais défini)

# --- Comptage ---
redis-cli BITCOUNT presence:2026-09-26        # nb de bits à 1 = nb de présents
redis-cli BITCOUNT presence:2026-09-26 0 0    # sur le premier octet seulement

# --- Opérations bit à bit (rétention : présents J ET J-1) ---
redis-cli BITOP AND retention:j_et_j1 presence:2026-09-26 presence:2026-09-25
redis-cli BITCOUNT retention:j_et_j1

# --- Trouver le premier bit à 1 ---
redis-cli BITPOS presence:2026-09-26 1        # position du premier présent
```

Ordre de grandeur mémoire : **1 million d'utilisateurs = ~125 Ko par jour**. Imbattable pour des cohortes/DAU. Limite : l'offset max est 2^32-1 ; au-delà, découpez en plusieurs clés.

---

## 18. HyperLogLog : introduction

`HyperLogLog` estime le **nombre d'éléments distincts** avec une erreur standard de **0,81 %**, pour **12 Ko** de mémoire — quel que soit le volume (milliards d'éléments).

```bash
redis-cli PFADD visiteurs:2026-09-26 "user:1042" "user:2048" "user:1042"
# → 1 (le HLL a été modifié ; le doublon est ignoré)

redis-cli PFCOUNT visiteurs:2026-09-26          # estimation des uniques
redis-cli PFCOUNT visiteurs:2026-09-26 visiteurs:2026-09-25  # union sur plusieurs jours
redis-cli PFMERGE visiteurs:semaine visiteurs:2026-09-2*     # fusion persistante
```

Cas d'usage : **visiteurs uniques**, IPs distinctes, recherches uniques. ⚠️ C'est une **estimation** : n'utilisez pas pour de la facturation ou des quotas stricts. Pour un comptage exact, utilisez un Set (mémoire proportionnelle au volume).

---

## 19. Geo : index géospatiaux (introduction)

Redis indexe des points (longitude, latitude) et répond à « quoi dans un rayon de X km ? ».

```bash
# --- Ajouter des points ---
redis-cli GEOADD agences:sn -17.4441 14.7167 "agence:dakar" -16.9269 14.6937 "agence:pikine"
# ordre : LONGITUDE LATITUDE membre

# --- Distance ---
redis-cli GEODIST agences:sn "agence:dakar" "agence:pikine" km
# → "17.1234" (km ; unités : m, km, ft, mi)

# --- Recherche autour d'un point (⚠️ 6.2+ : GEOSEARCH, GEORADIUS déprécié) ---
redis-cli GEOSEARCH agences:sn FROMLONLAT -17.4441 14.7167 BYRADIUS 20 km WITHDIST
# → agences dans un rayon de 20 km, avec distances

redis-cli GEOSEARCH agences:sn FROMMEMBER "agence:dakar" BYBOX 30 30 km COUNT 5

