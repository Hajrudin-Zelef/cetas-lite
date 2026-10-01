---
id: collect-261001-rattrapage/rattrapage/redis-guide-14
title: "Guide Redis — De l'installation à la production"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/redis_guide.md
source_anchor: ""
source_lines: [2536, 2733]
sha256: 12d0ced9feee3d2aeabb83fd9b1b87db450c5e446fc3a94e38212a83b84fb4ae
---

# Guide Redis — De l'installation à la production

| Cause | Remède |
|---|---|
| Un maître down sans réplica | Redémarrer le nœud ; sinon `CLUSTER FORGET` + reshard |
| Slots non assignés après ajout de nœud | `redis-cli --cluster reshard` / `rebalance` |
| `cluster-require-full-coverage yes` (défaut) bloque tout | Temporairement `no` pour dégrader gracieusement (à vos risques) |
| Bus cluster (16379) filtré | Ouvrir le port entre nœuds |

```bash
# Forcer l'oubli d'un nœud mort (sur chaque nœud restant) :
redis-cli -c -h 10.0.5.41 CLUSTER FORGET <node-id-du-mort>
# Réassigner ses slots au préalable via reshard si des données doivent survivre.
```

---

## 82. Dépannage : saturation des connexions clients

Symptômes : `ERR max number of clients reached`, `connected_clients` = `maxclients`.

```bash
redis-cli INFO clients | grep -E "connected_clients|maxclients|blocked_clients"
redis-cli CLIENT LIST | head -5        # qui est connecté ? (addr, cmd, idle...)
redis-cli CLIENT LIST | awk '{print $2}' | sort | uniq -c | sort -rn | head
# → top des IPs par nombre de connexions (fuite de connexions applicative ?)

# Tuer les connexions idle abusives d'une IP :
redis-cli CLIENT KILL ADDR 10.0.5.99:51234
# ou par pattern :
redis-cli CLIENT KILL TYPE normal SKIPME yes
```

Prévention :

```ini
maxclients 10000
timeout 300              # déconnecte les idle > 5 min
```

Et côté applicatif : **pool de connexions** (pas une connexion par requête), avec taille bornée et timeout.

---

## 83. Les 10 erreurs classiques (et comment les éviter)

| # | Erreur | Conséquence | Prévention |
|---|---|---|---|
| 1 | `KEYS *` en production | **Blocage total** du serveur (mono-thread) | `SCAN` + `--scan` ; `rename-command KEYS ""` |
| 2 | Aucune persistance configurée | Perte totale au reboot | RDB + AOF (section 30) |
| 3 | Pas de `maxmemory` | OOM killer tue Redis (ou la machine) | `maxmemory` + politique d'éviction (section 27) |
| 4 | Exposé sur Internet sans AUTH | Compromission, `FLUSHALL` + rançon | bind + ACL + firewall (sections 35–40) |
| 5 | `FLUSHALL` sur la mauvaise instance | Suppression de **toutes** les bases | ACL `-@dangerous`, `rename-command`, confirmation en prod |
| 6 | TTL identiques en masse | Avalanche d'expirations, pic CPU | Jitter ±10 % (section 33) |
| 7 | `SAVE` au lieu de `BGSAVE` | Blocage de plusieurs secondes/minutes | Toujours `BGSAVE` |
| 8 | Verrou sans TTL ni token | Verrou éternel / libération du verrou d'autrui | Pattern section 46 (NX+PX+token+Lua) |
| 9 | Client non Sentinel/Cluster-aware | Pas de failover, `MOVED` en erreur | Libs adaptées (section 60/64) |
| 10 | Backup jamais testé | Restauration impossible le jour J | Test trimestriel chronométré (section 67) |

Bonus — erreurs de config fréquentes :

- `bind 0.0.0.0` « pour que ça marche » → exposition totale.
- `appendfsync always` « pour la sécurité » → débit divisé par ~10 sans raison.
- 16 bases `SELECT` utilisées comme namespaces → cauchemar en Cluster ; préférez des préfixes de clés ou des instances.

---

## 84. Cas pratique 1 : cache devant PostgreSQL (catalogue produits)

Contexte : 50 000 produits, 2 000 req/s en lecture, 10 écritures/s. Objectif : < 5 ms p99, tolérance stale 5 min.

```bash
# --- Convention de clés ---
# cache:produit:<sku>          → JSON du produit, TTL 300 s + jitter
# cache:categorie:<id>:liste   → liste des SKU, TTL 600 s
# version:produit:<sku>        → version pour invalidation (section 54)

# --- Lecture (cache-aside, pseudo-code déjà vu section 52) ---
# 1. GET cache:produit:sku-7788
# 2. MISS → SELECT ... FROM produits WHERE sku='sku-7788'
# 3. SET ... EX <300±30>

# --- Écriture (invalidation explicite) ---
# UPDATE produits SET prix=... WHERE sku='sku-7788'
redis-cli DEL cache:produit:sku-7788
redis-cli DEL cache:categorie:12:liste
# Variante versioning (zéro SCAN) :
redis-cli INCR version:produit:sku-7788

# --- Warm-up après déploiement (pré-remplir le top 1000) ---
# Script appli : SELECT sku FROM produits ORDER BY vues DESC LIMIT 1000
# puis pipeline de SET (1 pipeline de 1000 SET ≈ quelques ms)

# --- Monitoring du cache ---
# hit rate :
#   redis-cli INFO stats | grep keyspace_hits / keyspace_misses
# Si hit rate < 95 % : TTL trop court ? clés trop granulaires ? avalanche ?
```

Dimensionnement : 50 000 produits × ~2 Ko = ~100 Mo → `maxmemory 512mb`, `allkeys-lru`, RDB seul (reconstituable).

---

## 85. Cas pratique 2 : compteur temps réel (votes / supervision)

Contexte : sondage en direct, 10 000 votes/min, affichage temps réel + top 5.

```bash
# --- Vote (1 par utilisateur : Set pour l'unicité + compteur) ---
redis-cli EVAL "
if redis.call('SADD', KEYS[2], ARGV[1]) == 1 then
  redis.call('INCR', KEYS[1])
  return 1
else
  return 0
end
" 2 sondage:42:votes sondage:42:votants user:1042
# → 1 vote compté, 0 déjà voté

# --- Lecture temps réel ---
redis-cli GET sondage:42:votes

# --- Sondages multiples : un hash global ---
redis-cli HINCRBY sondages:votes 42 1
redis-cli HGETALL sondages:votes

# --- Fenêtre glissante : votes des 60 dernières secondes (ZSET, section 50.2) ---
# ZADD sondage:42:fenetre <timestamp_ms> <uuid>  puis ZCOUNT sur [now-60000, now]

# --- Remise à zéro / archivage ---
redis-cli RENAME sondage:42:votes archive:sondage:42:votes   # atomique, sans copie
redis-cli EXPIRE archive:sondage:42:votes 2592000           # archive 30 jours
```

---

## 86. Cas pratique 3 : file de tâches avec workers (streams)

Contexte : génération de rapports PDF, 4 workers, exigence « aucun rapport perdu ».

```bash
# ===== PRODUCTEUR (appli web) =====
redis-cli XADD jobs:rapports MAXLEN ~ 20000 * type "pdf" dossier "D-7788" demandeur "awa"

# ===== WORKER (boucle, pseudo-shell) =====
# 1. Attend un job :
#    XREADGROUP GROUP grp:pdf worker-3 COUNT 1 BLOCK 10000 STREAMS jobs:rapports ">"
# 2. Génère le PDF...
# 3. Succès → XACK jobs:rapports grp:pdf <id>
#    Échec transient → ne pas ACK (restera pending, récupéré par le superviseur)

# ===== SUPERVISEUR (cron toutes les 2 min) =====
# Réclame les jobs idle depuis > 5 min (worker crashé en plein traitement) :
redis-cli XAUTOCLAIM jobs:rapports grp:pdf superviseur 300000 0 COUNT 20
# → les réinjecter dans le traitement ou alerter après 3 tentatives
# (compter les tentatives : XPENDING donne le retry count par entrée)

# ===== OBSERVABILITÉ =====
redis-cli XLEN jobs:rapports                          # taille de la file
redis-cli XPENDING jobs:rapports grp:pdf              # jobs en cours non acquittés
redis-cli XINFO CONSUMERS jobs:rapports grp:pdf       # activité par worker
```

Alertes : `XLEN` > 5 000 (workers sous-dimensionnés), pending > 100 depuis > 10 min (workers en panne).

---

## 87. Cas pratique 4 : rate limiter d'API (100 req/min)

Contexte : API publique, quota 100 req/min par clé API, réponse 429 au-delà.

```bash
# ===== Lua déployé une fois (fenêtre fixe, cf. section 50.1) =====
# Clé : ratelimit:<apikey> — valeur : compteur — TTL : 60 s

# ===== Test manuel =====
for i in $(seq 1 105); do
  c=$(redis-cli --raw EVAL "
local c = redis.call('INCR', KEYS[1])
if c == 1 then redis.call('EXPIRE', KEYS[1], ARGV[1]) end
return c" 1 ratelimit:cle-demo 60)
  if [ "$c" -gt 100 ]; then echo "requête $i → 429 (compteur=$c)"; break; fi
done

# ===== Côté appli (pseudo-code) =====
# c = lua_ratelimit("ratelimit:" + api_key, 60)
# headers: X-RateLimit-Limit: 100, X-RateLimit-Remaining: max(0, 100 - c)
# if c > 100: return 429 + Retry-After: <TTL restant>

# ===== Variante stricte (fenêtre glissante, section 50.2) =====
# À utiliser si les pics "à cheval sur deux fenêtres" sont inacceptables
# (ex. facturation au nombre d'appels).
```

> Pour un rate limiting **global** (tous clients confondus, ex. 10 000 req/s vers un backend) : même pattern avec une clé unique. Pour du **per-IP** : `ratelimit:ip:<ip>`.

---

## 88. Cas pratique 5 : sessions + panier e-commerce

