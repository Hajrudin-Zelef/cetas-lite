---
id: collect-261001-rattrapage/rattrapage/redis-guide-2
title: "Guide Redis — De l'installation à la production"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/redis_guide.md
source_anchor: ""
source_lines: [160, 359]
sha256: 613f12ae372d9b9bb04a473ec7396a6848029bd8aff2f4e6b38486b7b5ac1e1d
---

# Guide Redis — De l'installation à la production

Tableau récapitulatif des options de `SET` :

| Option | Effet |
|---|---|
| `EX secondes` | TTL en secondes |
| `PX millisecondes` | TTL en millisecondes |
| `EXAT timestamp` | Expire à un timestamp Unix (⚠️ 7.0+) |
| `PXAT timestamp-ms` | Expire à un timestamp en ms (⚠️ 7.0+) |
| `NX` | Écrit seulement si la clé **n'existe pas** |
| `XX` | Écrit seulement si la clé **existe déjà** |
| `KEEPTTL` | ⚠️ 6.0+ : conserve le TTL existant |
| `GET` | ⚠️ 6.2+ : retourne l'ancienne valeur (comme GETSET) |

---

## 7. Hashes : concepts

Un hash = un **objet** (ensemble de champs/valeurs) sous une seule clé. Idéal pour les fiches : `user:1042` → `{name, email, role, ...}`.

Avantages vs N strings :

- Une seule clé à expirer/gérer (le TTL s'applique **au hash entier**, pas par champ — ⚠️ 7.4+ introduit l'expiration par champ, voir section 8).
- `HGETALL` récupère l'objet d'un coup ; `HMGET` ne prend que les champs utiles (économie réseau).
- Mémoire optimisée : avec peu de champs, Redis utilise un encodage compact (`listpack`).

Convention de nommage : `objet:id` → `user:1042`, `produit:sku-7788`, `vm:proxmox-12`.

---

## 8. Hashes : commandes essentielles

```bash
# --- Écriture ---
redis-cli HSET user:1042 name "Awa Diallo" email "awa@exemple.sn" role "admin"
# → 3 (nombre de champs créés)

redis-cli HSETNX user:1042 role "user"   # n'écrit que si le champ n'existe pas → 0 ici

# --- Lecture ---
redis-cli HGET user:1042 name
# "Awa Diallo"

redis-cli HMGET user:1042 name email role
# 1) "Awa Diallo"
# 2) "awa@exemple.sn"
# 3) "admin"

redis-cli HGETALL user:1042
# 1) "name"  2) "Awa Diallo"  3) "email" ...

redis-cli HKEYS user:1042    # liste des champs
redis-cli HVALS user:1042    # liste des valeurs
redis-cli HLEN user:1042     # nombre de champs → 3
redis-cli HEXISTS user:1042 email  # → 1

# --- Compteurs dans un hash ---
redis-cli HINCRBY stats:page:accueil vues 1
redis-cli HINCRBYFLOAT stats:panier total 49.90

# --- Suppression ---
redis-cli HDEL user:1042 role     # supprime le champ
redis-cli DEL user:1042           # supprime tout le hash

# --- Itération sûre (gros hashes) ---
redis-cli HSCAN user:1042 0
# Comme SCAN : curseur, à boucler jusqu'à curseur 0 (voir section 91)

# --- Expiration par champ (⚠️ Redis 7.4+) ---
redis-cli HEXPIRE user:1042 3600 FIELDS 1 email   # le champ email expire dans 1 h
redis-cli HTTL user:1042 FIELDS 1 email           # TTL restant du champ
```

> **Piège classique :** `HGETALL` sur un hash de 100 000 champs bloque le serveur le temps du transfert. En prod, préférez `HMGET`/`HSCAN`.

---

## 9. Lists : concepts

Liste ordonnée de strings, optimisée pour les **insertions/suppressions aux deux extrémités** (implémentation : quicklist). C'est la structure historique des **files d'attente**.

- `LPUSH` = empile à gauche (tête), `RPUSH` = à droite (queue).
- `LPOP`/`RPOP` = dépile. Avec `BLPOP`/`BRPOP`, le client **attend** (bloquant) qu'un élément arrive — parfait pour un worker.
- `LLEN`, `LRANGE ma_liste 0 -1` (tout), `LINDEX`, `LTRIM` (garder les N derniers : historique borné).

Modèle file FIFO : producteurs `RPUSH`, consommateurs `BLPOP` (ou `BRPOPLPUSH` pour fiabiliser, voir section 48).

---

## 10. Lists : commandes essentielles

```bash
# --- File d'attente FIFO ---
redis-cli RPUSH jobs:emails "mail:1" "mail:2" "mail:3"
redis-cli LLEN jobs:emails          # → 3
redis-cli LPOP jobs:emails          # "mail:1" (le plus ancien)
redis-cli RPOP jobs:emails          # "mail:3" (le plus récent)

# --- Consommation bloquante (worker) ---
redis-cli BLPOP jobs:emails 30
# attend jusqu'à 30 s un élément ; retourne ["jobs:emails","mail:2"] ou nil après timeout
# Variante multi-files : BLPOP jobs:urgent jobs:emails 10  (priorité à jobs:urgent)

# --- Lecture sans dépiler ---
redis-cli LRANGE jobs:emails 0 -1     # tout
redis-cli LRANGE jobs:emails 0 9      # les 10 premiers (pagination)
redis-cli LINDEX jobs:emails 0        # premier élément
redis-cli LPOS jobs:emails "mail:2"   # ⚠️ 6.0.6+ : position d'un élément

# --- Pile (LIFO) ---
redis-cli LPUSH pile "a" "b" "c"
redis-cli LPOP pile   # "c"

# --- Historique borné (les 100 derniers logs) ---
redis-cli RPUSH logs:app "event..." 
redis-cli LTRIM logs:app -100 -1      # ne garde que les 100 derniers
# ⚠️ LTRIM sur une liste en cours d'écriture : à combiner dans un pipeline/Lua

# --- Transfert atomique entre listes (file fiable, voir section 48) ---
redis-cli BRPOPLPUSH jobs:emails jobs:encours 30  # ⚠️ 6.2+ : préférer LMOVE
redis-cli LMOVE jobs:emails jobs:encours LEFT RIGHT  # ⚠️ 6.2+

# --- Insertion ciblée ---
redis-cli LINSERT jobs:emails BEFORE "mail:2" "mail:urgent"
redis-cli LSET jobs:emails 0 "mail:1-modifie"   # écrase par index
redis-cli LREM jobs:emails 1 "mail:2"          # supprime 1 occurrence
```

---

## 11. Sets : concepts

Ensemble **non ordonné** de membres **uniques**. Les doublons sont ignorés silencieusement. Opérations ensemblistes natives : **intersection**, **union**, **différence** — calculées côté serveur, très rapides.

Cas d'usage :

- Tags d'articles : `SADD article:55:tags php redis cache`
- Abonnés, votants (unicité garantie), IPs bannies
- « Utilisateurs en ligne » avec expiration globale
- Recommandations : `SINTER user:1:achats user:2:achats` → goûts communs

Complexité : O(1) pour SADD/SREM/SISMEMBER, O(N) pour les opérations multi-sets.

---

## 12. Sets : commandes essentielles

```bash
# --- Bases ---
redis-cli SADD tags:article:55 php redis cache redis   # → 3 (doublon ignoré)
redis-cli SMEMBERS tags:article:55
redis-cli SISMEMBER tags:article:55 php    # → 1
redis-cli SMISMEMBER tags:article:55 php java  # ⚠️ 6.2+ : 1) 1  2) 0
redis-cli SCARD tags:article:55             # → 3
redis-cli SREM tags:article:55 php         # → 1

# --- Tirage aléatoire ---
redis-cli SRANDMEMBER tags:article:55        # 1 membre au hasard
redis-cli SRANDMEMBER tags:article:55 2      # 2 membres (peut répéter si > card... non : sans répétition si positif)
redis-cli SPOP tags:article:55              # retire ET retourne un membre au hasard

# --- Opérations ensemblistes ---
redis-cli SADD user:1:achats livre souris clavier
redis-cli SADD user:2:achats livre ecran clavier
redis-cli SINTER user:1:achats user:2:achats   # → livre, clavier
redis-cli SUNION  user:1:achats user:2:achats  # → les 5
redis-cli SDIFF   user:1:achats user:2:achats  # → souris (chez 1 mais pas 2)

# --- Stocker le résultat (évite de ramener des Mo au client) ---
redis-cli SINTERSTORE reco:1:2 user:1:achats user:2:achats
redis-cli SUNIONSTORE tous:achats user:1:achats user:2:achats
redis-cli SDIFFSTORE diff:1:2 user:1:achats user:2:achats

# --- Déplacer un membre entre sets ---
redis-cli SMOVE user:1:achats user:2:achats souris

# --- Itération ---
redis-cli SSCAN tags:article:55 0
```

> **Astuce prod :** pour « N membres au hasard sans les retirer » sur un gros set, `SRANDMEMBER key -N` (négatif = avec répétitions possibles) est plus rapide que `SMEMBERS` + tirage côté client.

---

## 13. Sorted Sets : concepts

Comme un Set, mais chaque membre a un **score** (flottant). Redis maintient l'ordre **trié par score** (puis lexicographique à score égal). Implémentation : skiplist + dict → O(log N) pour ZADD/ZRANGE.

Cas d'usage rois :

- **Classements / leaderboards** (score = points)
- **Plannings / files à priorité** (score = timestamp → `ZRANGEBYSCORE` pour « les tâches dues »)
- **Rate limiting** à fenêtre glissante (score = timestamp, voir section 50)
- **Time-series** simples

Le membre doit être unique ; pour le même membre, `ZADD` **met à jour** le score (pratique pour les classements évolutifs).

---

## 14. Sorted Sets : commandes essentielles

