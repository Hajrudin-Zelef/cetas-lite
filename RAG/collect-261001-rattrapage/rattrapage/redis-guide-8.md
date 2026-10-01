---
id: collect-261001-rattrapage/rattrapage/redis-guide-8
title: "Guide Redis — De l'installation à la production"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["acquisition", "consumer", "distribution"]
source: docs/RAG/collect-261001-rattrapage/redis_guide.md
source_anchor: ""
source_lines: [1365, 1545]
sha256: 09f4ca69a060ee2c6faa4f2947ee58b46b67386e7861437b69b3bc2809154de1
---

# Guide Redis — De l'installation à la production

- Pas de rejeu, pas d'ACK, pas de file d'attente → **ne pas utiliser pour des événements critiques** (préférez les Streams, section 16).
- Un subscriber lent fait gonfler les **output buffers** du serveur → risque OOM (réglez `client-output-buffer-limit pubsub 32mb 8mb 60`).
- `UNSUBSCRIBE` / `PUNSUBSCRIBE` / `PUBSUB CHANNELS` / `PUBSUB NUMSUB canal` pour l'introspection.

Cas d'usage légitimes : invalidation de cache entre instances applicatives, notifications temps réel non critiques, bus d'événements éphémères.

---

## 45. Keyspace notifications : réagir aux événements

Redis peut **publier** des événements internes (expiration, éviction, écritures...) sur des canaux `__keyspace@<db>__:<clé>` et `__keyevent@<db>__:<op>`.

```ini
# redis.conf — quoi notifier (chaîne de flags) :
notify-keyspace-events Ex
# K = keyspace, E = keyevent, g = génériques, $ = strings, l = lists, s = sets,
# h = hashes, z = sorted sets, x = expirations, e = évictions, A = alias de "g$lshzx"
```

```bash
# Exemple : être notifié des expirations (utile : "session expirée → nettoyage")
redis-cli CONFIG SET notify-keyspace-events Ex

# Terminal 1 :
redis-cli PSUBSCRIBE "__keyevent@0__:expired"

# Terminal 2 :
redis-cli SET session:temp "x" EX 5
# ... 5 s plus tard, terminal 1 reçoit : 1) "pmessage" 2) "__keyevent@0__:expired" 3) "__keyevent@0__:expired" 4) "session:temp"

# Savoir QUI a expiré quoi par clé :
# Terminal 1 :
redis-cli SUBSCRIBE "__keyspace@0__:session:temp"
```

> ⚠️ Les notifications d'expiration ne sont **pas garanties** en cas de failover : un événement peut être manqué. Ne basez pas une logique métier critique dessus sans filet (ex. un job de réconciliation périodique).

---


## 46. Verrous distribués : le pattern SET NX PX

Besoin : « un seul worker à la fois exécute le job 42 ». La recette correcte tient en 3 commandes :

```bash
# --- 1. Acquisition (atomique grâce à NX + PX combinés) ---
# token unique par détenteur (uuid) — indispensable pour ne libérer que SON verrou
TOKEN="worker-1:$(uuidgen)"
redis-cli SET verrou:job42 "$TOKEN" NX PX 30000
# → OK  : verrou acquis pour 30 s
# → nil : déjà pris, réessayer plus tard (avec backoff + jitter)

# --- 2. Renouvellement (si le traitement dépasse le TTL) ---
redis-cli EVAL "
if redis.call('GET', KEYS[1]) == ARGV[1] then
  return redis.call('PEXPIRE', KEYS[1], ARGV[2])
else return 0 end
" 1 verrou:job42 "$TOKEN" 30000

# --- 3. Libération (atomique : seulement si on est propriétaire) ---
redis-cli EVAL "
if redis.call('GET', KEYS[1]) == ARGV[1] then
  return redis.call('DEL', KEYS[1])
else return 0 end
" 1 verrou:job42 "$TOKEN"
```

Règles impératives :

- **Token unique** par détenteur (jamais une valeur fixe comme `"1"`).
- **TTL toujours** (`PX`) : sans lui, un worker qui crash laisse un verrou éternel.
- Libération/renouvellement **en Lua** avec vérification du token (le `GET`+`DEL` en deux temps est une race condition classique).
- Le TTL doit couvrir la durée du traitement **ou** être renouvelé périodiquement (< TTL/3).
- Ce pattern est valable sur **une instance unique** (ou un maître avec réplication synchrone). Pour du multi-maître, voir Redlock (section 47).

> ⚠️ Avec réplication **asynchrone** : si le maître tombe avant d'avoir répliqué le `SET NX`, un nouveau maître peut accorder le verrou à un autre worker → **double détention possible**. Si c'est inacceptable, utilisez `WAIT` (réplication semi-synchrone) ou Redlock.

---

## 47. Redlock : concepts et avertissements

**Redlock** (algorithme proposé par l'auteur de Redis) étend le verrou à **N instances indépendantes** (ex. 5) : le verrou est acquis s'il est obtenu sur une **majorité** (quorum) en moins que le TTL.

Principe (simplifié) :

1. Générer un token unique, noter l'heure de début.
2. `SET NX PX` séquentiellement sur les N instances.
3. Si quorum atteint **et** temps écoulé < TTL → verrou acquis (TTL effectif = TTL initial − temps écoulé).
4. Sinon, libérer partout (script Lua) et réessayer.

**Avertissements honnêtes (débat Martin Kleppmann vs Salvatore) :**

- Redlock suppose des **horloges à peu près synchronisées** et des **pauses GC/processus bornées** — hypothèses fragiles.
- Il ne protège pas contre un processus en pause longue qui « se réveille » après expiration du verrou (fencing tokens requis pour une vraie sûreté).
- **Recommandation pratique :** pour 95 % des besoins (leader election, jobs exclusifs), le **verrou mono-instance** (section 46) + `WAIT 1 1000` suffit. N'utilisez Redlock que si vous avez déjà N instances indépendantes **et** que vous comprenez ses limites. Les bibliothèques (`redlock-py`, etc.) implémentent l'algorithme — ne le codez pas à la main.

---

## 48. Files d'attente avec les Lists

Le pattern historique, simple et rapide. Producteurs → `RPUSH`, workers → `BLPOP`.

```bash
# --- Producteur ---
redis-cli RPUSH jobs:impression "job: impression doc 7788, imprimante kyocera-3"

# --- Worker fiable : BRPOPLPUSH / LMOVE (⚠️ 6.2+) ---
# Le job est déplacé atomiquement vers une liste "en cours" :
redis-cli LMOVE jobs:impression jobs:impression:encours RIGHT LEFT
# → "job: ..." (le worker le traite...)
# Succès :
redis-cli LREM jobs:impression:encours 1 "job: ..."
# Échec / worker mort : un superviseur remet les jobs de :encours vers :impression
# après un délai (scan périodique).

# --- Priorités : plusieurs files ---
redis-cli BLPOP jobs:urgent jobs:impression 30   # jobs:urgent d'abord
```

Limites des lists pour les files :

- Pas d'ACK natif (le pattern `:encours` est artisanal).
- Pas de rejeu, pas de multi-consommateurs avec distribution garantie.
- Un message = une string (sérialisez votre payload en JSON).

**Quand passer aux Streams (section 49) :** besoin d'ACK, de rejeu, de groupes de consommateurs, ou de visibilité sur les pending.

---

## 49. Files d'attente avec les Streams (consumer groups)

Le pattern robuste : distribution à un seul worker par groupe + ACK + récupération des pannes.

```bash
# --- Setup (une fois) ---
redis-cli XGROUP CREATE jobs:emails grp:workers 0 MKSTREAM

# --- Worker (boucle) ---
# 1. Récupère jusqu'à 10 jobs jamais distribués, bloque 5 s :
redis-cli XREADGROUP GROUP grp:workers worker-1 COUNT 10 BLOCK 5000 STREAMS jobs:emails ">"
# 2. Traite chaque entrée...
# 3. Acquitte :
redis-cli XACK jobs:emails grp:workers 1769557200000-0 1769557200001-0

# --- Superviseur : récupère les jobs bloqués (worker mort) ---
# Toutes les 60 s : réclame les pending idle depuis > 120 s
redis-cli XAUTOCLAIM jobs:emails grp:workers superviseur 120000 0 COUNT 50
# puis XACK après retraitement

# --- Producteur ---
redis-cli XADD jobs:emails MAXLEN ~ 50000 * type "email" to "awa@exemple.sn" template "facture"
```

Tableau de décision file :

| Besoin | Lists | Streams |
|---|---|---|
| Simple, fire-and-forget | ✅ | — |
| ACK / au-moins-une-fois | artisanal (`:encours`) | ✅ natif |
| Rejeu / audit | ❌ | ✅ (`XREAD` depuis un ID) |
| Multi-workers sans doublon | ⚠️ (BLPOP distribue, sans garantie) | ✅ consumer groups |
| Latence minimale | ✅ légèrement plus rapide | très bonne aussi |

---

## 50. Rate limiting : 3 algorithmes

### 50.1 Fenêtre fixe (simple, ~suffisant)

```bash
# Autorise 100 requêtes / 60 s par clé API :
redis-cli EVAL "
local c = redis.call('INCR', KEYS[1])
if c == 1 then redis.call('EXPIRE', KEYS[1], ARGV[1]) end
return c
" 1 ratelimit:apikey:7 60
# L'appli autorise si c <= 100. Inconvénient : pic de 200 req à cheval sur 2 fenêtres.
```

### 50.2 Fenêtre glissante avec Sorted Set (précis)

