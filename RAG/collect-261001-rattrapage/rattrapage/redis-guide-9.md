---
id: collect-261001-rattrapage/rattrapage/redis-guide-9
title: "Guide Redis — De l'installation à la production"
domain: rattrapage
role: reference
task: reference
actors: []
dates: ["2026-09-26"]
keywords: ["leaderboard"]
source: docs/RAG/collect-261001-rattrapage/redis_guide.md
source_anchor: ""
source_lines: [1546, 1760]
sha256: 9427f60af15bffed95d4960ad64101a5d67a20417005f6c2ea082708e28cfeb2
---

# Guide Redis — De l'installation à la production

```bash
# Autorise 100 requêtes sur les 60 dernières secondes :
redis-cli EVAL "
local now = tonumber(ARGV[1])
local window = tonumber(ARGV[2])
local limit = tonumber(ARGV[3])
redis.call('ZREMRANGEBYSCORE', KEYS[1], 0, now - window)
local c = redis.call('ZCARD', KEYS[1])
if c < limit then
  redis.call('ZADD', KEYS[1], now, now .. ':' .. ARGV[4])
  redis.call('EXPIRE', KEYS[1], window)
  return 1
else
  return 0
end
" 1 ratelimit:glissant:7 1769557200000 60000 100 $(uuidgen)
# → 1 autorisé, 0 refusé. Précis mais 1 clé + N entrées par client.
```

### 50.3 Token bucket (lisse les pics)

```bash
# Seau de 100 tokens, refill 10/s — implémentation Lua classique :
redis-cli EVAL "
local key = KEYS[1]
local capacity = tonumber(ARGV[1])
local refill_per_sec = tonumber(ARGV[2])
local now = tonumber(ARGV[3])
local data = redis.call('HMGET', key, 'tokens', 'ts')
local tokens = tonumber(data[1]) or capacity
local ts = tonumber(data[2]) or now
tokens = math.min(capacity, tokens + (now - ts) / 1000 * refill_per_sec)
if tokens >= 1 then
  tokens = tokens - 1
  redis.call('HSET', key, 'tokens', tokens, 'ts', now)
  redis.call('EXPIRE', key, 3600)
  return 1
else
  redis.call('HSET', key, 'tokens', tokens, 'ts', now)
  return 0
end
" 1 bucket:apikey:7 100 10 1769557200000
```

> Réponse HTTP standard en cas de refus : `429 Too Many Requests` + headers `Retry-After`, `X-RateLimit-Limit`, `X-RateLimit-Remaining`.

---

## 51. Sessions utilisateur

Stocker les sessions en Redis : invalidation instantanée, partage entre N serveurs web, expiration automatique.

```bash
# --- Création (à la connexion) ---
# Token de session opaque, jamais devinable :
TOKEN=$(openssl rand -hex 32)
redis-cli HSET session:$TOKEN user_id 1042 role "admin" ip "10.0.5.44" created 1769557200
redis-cli EXPIRE session:$TOKEN 1800     # 30 min d'inactivité

# --- Lecture + prolongation (sliding expiration) à chaque requête ---
redis-cli HGETALL session:$TOKEN
redis-cli EXPIRE session:$TOKEN 1800

# --- Déconnexion / révocation ---
redis-cli DEL session:$TOKEN             # immédiat sur tous les frontaux

# --- Révocation globale d'un utilisateur (ex. compte compromis) ---
# Maintenir un index : SADD sessions:user:1042 $TOKEN (à la création)
redis-cli SMEMBERS sessions:user:1042
# puis DEL de chaque session + DEL de l'index (script Lua pour l'atomicité)
```

Bonnes pratiques :

- Cookie `HttpOnly; Secure; SameSite=Lax`, valeur = token opaque (jamais l'ID utilisateur).
- **Sliding expiration** : `EXPIRE` à chaque requête authentifiée.
- Données minimales dans la session (ID + rôle) ; le reste en base.
- En cas de compromission : rotation du token + révocation via l'index.

---

## 52. Patterns de cache : cache-aside (lazy loading)

Le pattern le plus répandu : l'application **lit d'abord Redis**, et ne va en base qu'en cas de miss.

```python
# Pseudo-code Python (redis-py)
def get_produit(sku):
    key = f"cache:produit:{sku}"
    data = redis.get(key)
    if data:
        return json.loads(data)          # HIT
    # MISS :
    row = db.query("SELECT * FROM produits WHERE sku = %s", sku)
    if row is None:
        return None
    payload = json.dumps(row)
    ttl = 3600 + random.randint(-300, 300)   # jitter anti-avalanche
    redis.set(key, payload, ex=ttl)
    return row
```

```bash
# Équivalent manuel en redis-cli :
redis-cli GET cache:produit:sku-7788        # MISS → nil
# ... requête SQL ...
redis-cli SET cache:produit:sku-7788 '{"nom":"Onduleur 3kVA","prix":450000}' EX 3600
redis-cli GET cache:produit:sku-7788        # HIT
```

Avantages : simple, ne cache que ce qui est demandé (pas de remplissage inutile).
Inconvénients : premier appel lent (miss), **données potentiellement périmées** entre deux expirations → à combiner avec l'invalidation (section 54).

---

## 53. Patterns de cache : write-through / write-behind

**Write-through :** l'écriture va **d'abord** (ou en même temps) dans le cache **et** en base, de façon synchrone.

```python
def update_prix(sku, nouveau_prix):
    db.execute("UPDATE produits SET prix=%s WHERE sku=%s", nouveau_prix, sku)
    redis.set(f"cache:produit:{sku}", json.dumps({"prix": nouveau_prix}), ex=3600)
    # ou redis.delete(...) pour forcer un rechargement (variante write-invalidate)
```

| Pattern | Cohérence | Latence d'écriture | Quand |
|---|---|---|---|
| Cache-aside | Éventuelle (TTL) | Faible | Lecture >> écriture (catalogue, fiches) |
| Write-through | Forte | Plus élevée (2 écritures) | Données critiques lues souvent |
| Write-behind | Faible (async) | Très faible | Compteurs, stats (risque de perte) |
| Refresh-ahead | Bonne | Faible (refresh en tâche de fond) | Clés chaudes à SLA strict |

> En pratique : **cache-aside + invalidation explicite à l'écriture** couvre 80 % des besoins (section 54).

---

## 54. Invalidation de cache : stratégies

Le cache périmé est la source n°1 de bugs « incompréhensibles ». Stratégies, de la plus simple à la plus robuste :

```bash
# --- 1. TTL (expiration passive) ---
redis-cli SET cache:produit:42 "..." EX 600   # périmé au pire 10 min

# --- 2. Invalidation explicite à l'écriture ---
# Quand le produit 42 change : 
redis-cli DEL cache:produit:42

# --- 3. Invalidation par motif (⚠️ SCAN, jamais KEYS) ---
# Tous les caches liés à la catégorie 7 :
redis-cli --scan --pattern "cache:categorie:7:*" | xargs -L 100 redis-cli DEL
# Mieux : maintenir un index (Set) des clés à invalider par entité :
redis-cli SADD idx:produit:42 "cache:produit:42" "cache:recherche:page:3"
# ... à l'invalidation : SMEMBERS puis DEL (script Lua)

# --- 4. Versioning des clés (zéro DEL) ---
# Clé = cache:produit:42:v3 ; à chaque modification, on incrémente la version :
redis-cli INCR version:produit:42   # → 4, les lecteurs utilisent v4
# Les anciennes versions expirent via leur TTL. Aucune suppression, aucun SCAN.

# --- 5. Pub/Sub d'invalidation entre N frontaux (chaque frontal a son cache local) ---
redis-cli PUBLISH inval:produit "42"   # chaque frontal DEL son cache local
```

Règle d'or : **définissez la tolérance au stale par donnée** (prix d'un produit : 0 ; nombre de vues : 10 min OK) et choisissez la stratégie en conséquence.

---

## 55. Cas pratique : compteur temps réel

Besoin : afficher en temps réel le nombre de dossiers traités par l'équipe, remis à zéro chaque jour.

```bash
# --- Incrément à chaque dossier clôturé (atomique, sans race) ---
redis-cli INCR dossiers:2026-09-26:clotures
# Le premier INCR du jour crée la clé à 1 ; on pose l'expiration une fois :
redis-cli EVAL "
local c = redis.call('INCR', KEYS[1])
if c == 1 then redis.call('EXPIRE', KEYS[1], 86400 * 2) end
return c
" 1 dossiers:2026-09-26:clotures

# --- Lecture pour le dashboard ---
redis-cli GET dossiers:2026-09-26:clotures

# --- Compteurs multiples en un hash (par technicien) ---
redis-cli HINCRBY dossiers:2026-09-26:par-tech awa 1
redis-cli HGETALL dossiers:2026-09-26:par-tech
redis-cli EXPIRE dossiers:2026-09-26:par-tech 172800

# --- Top du jour ---
# (si besoin de classement : Sorted Set)
redis-cli ZINCRBY top:tech:2026-09-26 1 awa
redis-cli ZRANGE top:tech:2026-09-26 0 4 REV WITHSCORES
```

> Pour des compteurs « à la seconde près » multi-sites : un `INCR` par site + agrégation périodique, ou un stream d'événements + agrégat.

---

## 56. Cas pratique : leaderboard / classements

```bash
# --- Enregistrer un score (mise à jour si meilleur... ou pas) ---
# Score du joueur = son max :
redis-cli EVAL "
local cur = tonumber(redis.call('ZSCORE', KEYS[1], ARGV[1]) or '-1')
local new = tonumber(ARGV[2])
if new > cur then return redis.call('ZADD', KEYS[1], new, ARGV[1]) else return 0 end
" 1 leaderboard:quiz awa 850

# --- Top 10 ---
redis-cli ZRANGE leaderboard:quiz 0 9 REV WITHSCORES

