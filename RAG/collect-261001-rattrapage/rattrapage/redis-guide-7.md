---
id: collect-261001-rattrapage/rattrapage/redis-guide-7
title: "Guide Redis — De l'installation à la production"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/redis_guide.md
source_anchor: ""
source_lines: [1150, 1364]
sha256: 0933b57044db891f135bf48736aaa029ef170419fa2341a5dc3d1dc52c79d3ce
---

# Guide Redis — De l'installation à la production

```ini
# /etc/redis/conf.d/99-tls.conf
port 0                      # désactive le port non chiffré
tls-port 6379
tls-cert-file /etc/redis/tls/redis.crt
tls-key-file /etc/redis/tls/redis.key
tls-ca-cert-file /etc/redis/tls/ca.crt
tls-replication yes         # chiffre aussi la réplication
tls-cluster yes             # ... et le bus cluster
tls-auth-clients no         # yes = exige un certificat client (mTLS)
```

```bash
# --- 3. Client ---
redis-cli --tls --cacert /etc/redis/tls/ca.crt -p 6379 ping

# --- 4. Permissions ---
sudo chown redis:redis /etc/redis/tls/redis.key
sudo chmod 600 /etc/redis/tls/redis.key
```

> Coût : +quelques % de CPU, latence légèrement accrue. Sur un VLAN privé de confiance, beaucoup d'équipes s'en passent ; dès que le trafic traverse un réseau mutualisé ou à cheval sur sites, **TLS obligatoire**.

---

## 39. Sécurité : rename-command (désarmer les commandes dangereuses)

Même avec AUTH, une erreur applicative (`FLUSHDB` dans un mauvais environnement) peut être catastrophique. Renommez ou désactivez les commandes dangereuses :

```ini
# Désactiver (nom vide = commande inexistante) :
rename-command FLUSHALL ""
rename-command FLUSHDB ""
rename-command DEBUG ""
rename-command CONFIG "CONFIG_9f8a7b6c5d"   # renommer plutôt que couper : garde l'admin possible

# Autres candidates : KEYS, SHUTDOWN, MIGRATE, MODULE, ACL
```

```bash
redis-cli FLUSHALL
# (error) unknown command `flushall`, with args beginning with:
```

> ⚠️ Avec Sentinel/Cluster, `CONFIG` renommée peut casser certains outils d'orchestration : testez. Alternative moderne : les **ACL** (`-@dangerous`) par utilisateur, plus fines que le rename global.

---

## 40. Sécurité : checklist de durcissement

Checklist à cocher avant toute mise en production :

- [ ] `bind` restreint aux interfaces strictement nécessaires
- [ ] `protected-mode yes`
- [ ] AUTH via **ACL** (`default off`, un user par appli, moindre privilège)
- [ ] Secrets forts (64+ caractères), stockés en coffre (Vault/ansible-vault), **jamais** en clair dans git
- [ ] `rename-command` ou ACL `-@dangerous` pour FLUSHALL/FLUSHDB/DEBUG/CONFIG/KEYS
- [ ] Firewall (iptables/nftables/security group) : 6379/6380/16379 ouverts **uniquement** vers les clients légitimes
- [ ] TLS si le réseau n'est pas de confiance
- [ ] Fichiers : `redis.conf` (640, root:redis), `dump.rdb`/`appendonlydir` (600, redis:redis), clés TLS (600)
- [ ] Pas d'exposition Internet (vérifié par scan externe)
- [ ] `supervised systemd`, `daemonize no` (paquet Debian : déjà OK)
- [ ] Logs : `logfile /var/log/redis/redis-server.log`, `loglevel notice`
- [ ] Test d'intrusion interne : tentative de connexion sans AUTH depuis un poste non autorisé → doit échouer

```bash
# Audit express d'une instance :
redis-cli INFO server | grep -E "redis_version|tcp_port"
redis-cli CONFIG GET "requirepass"   # ne doit pas être vide (ou aclfile en place)
redis-cli CONFIG GET "bind"
redis-cli ACL LIST 2>/dev/null | head -5
ss -ltnp | grep -E "6379|6380"
```

---

## 41. Transactions : MULTI / EXEC / WATCH

`MULTI`/`EXEC` exécute une **file de commandes de façon atomique et isolée** (pas d'entrelacement avec d'autres clients). Ce n'est **pas** un rollback : si une commande échoue à l'exécution, les autres sont quand même appliquées.

```bash
redis-cli MULTI
redis-cli INCR compteur:visites
redis-cli SET dernier:visiteur "awa"
redis-cli EXPIRE dernier:visiteur 3600
redis-cli EXEC
# 1) (integer) 1241
# 2) OK
# 3) (integer) 1

# --- Annulation ---
redis-cli MULTI
redis-cli SET a 1
redis-cli DISCARD     # annule la file

# --- Optimistic locking avec WATCH ---
redis-cli WATCH solde:compte:42
# ... lecture du solde côté appli ...
redis-cli MULTI
redis-cli DECRBY solde:compte:42 100
redis-cli EXEC
# → nil si solde:compte:42 a été modifié entre-temps (la transaction est annulée)
# → l'appli doit réessayer (boucle retry)
redis-cli UNWATCH
```

Quand préférer Lua (section 42) : dès que la logique contient des **conditions** (« décrémente seulement si > 0 ») — MULTI/EXEC ne fait pas de branchements.

---

## 42. Lua scripting : EVAL / EVALSHA

Les scripts Lua s'exécutent **côté serveur, atomiquement** : idéal pour les opérations « lire-décider-écrire » sans aller-retour réseau ni race.

```bash
# --- Exemple : décrémente un stock seulement si > 0 ---
redis-cli EVAL "
local stock = tonumber(redis.call('GET', KEYS[1]) or '0')
if stock > 0 then
  return redis.call('DECR', KEYS[1])
else
  return -1
end
" 1 stock:produit:42
# → nouveau stock, ou -1 si rupture

# --- Bonnes pratiques : charger une fois, appeler par SHA ---
redis-cli SCRIPT LOAD "return redis.call('GET', KEYS[1])"
# → "a42059b...sha1..."  (à stocker côté appli)
redis-cli EVALSHA a42059b... 1 ma_cle
# NOSCRIPT si le script a été flushé → fallback sur EVAL, ou SCRIPT EXISTS

# --- Gestion ---
redis-cli SCRIPT FLUSH          # vide le cache de scripts (⚠️ EVALSHA suivants → NOSCRIPT)
redis-cli SCRIPT EXISTS sha1 sha2
```

Règles d'or du Lua Redis :

- **Déterministe** : pas d'accès réseau, pas de temps aléatoire non seedé, pas d'itération d'ordre aléatoire sur les tables.
- **Rapide** : le script bloque le serveur pendant son exécution (mono-thread). Pas de boucle infinie !
- `KEYS[]` = clés touchées (obligatoire pour le Cluster : toutes les clés d'un script doivent être dans le **même hash slot**, utilisez les hash tags `{...}`).
- `ARGV[]` = arguments non-clés.
- En réplication, c'est le **script** qui est répliqué (effets déterministes garantis), pas ses écritures une par une.

---

## 43. Lua : cas pratiques commentés

```bash
# --- 1. Rate limiter à fenêtre fixe atomique (voir aussi section 50) ---
redis-cli EVAL "
local c = redis.call('INCR', KEYS[1])
if c == 1 then redis.call('EXPIRE', KEYS[1], ARGV[1]) end
return c
" 1 ratelimit:api:client:7 60
# → nombre de requêtes dans la fenêtre ; l'appli compare à la limite

# --- 2. Verrou avec renouvellement (prolonge seulement si on est le propriétaire) ---
redis-cli EVAL "
if redis.call('GET', KEYS[1]) == ARGV[1] then
  return redis.call('PEXPIRE', KEYS[1], ARGV[2])
else
  return 0
end
" 1 verrou:job42 worker-1 30000

# --- 3. Libération sûre d'un verrou (ne supprime que si propriétaire) ---
redis-cli EVAL "
if redis.call('GET', KEYS[1]) == ARGV[1] then
  return redis.call('DEL', KEYS[1])
else
  return 0
end
" 1 verrou:job42 worker-1

# --- 4. File à priorité avec scores + limite ---
redis-cli EVAL "
local items = redis.call('ZRANGE', KEYS[1], 0, tonumber(ARGV[1]) - 1)
if #items > 0 then redis.call('ZREM', KEYS[1], unpack(items)) end
return items
" 1 file:prioritaire 5
```

> Les scripts 2 et 3 sont le cœur d'un **verrou distribué correct** (section 46). Copiez-les tels quels dans vos libs plutôt que de réinventer le GET+DEL non atomique.

---

## 44. Pub/Sub : concepts et commandes

Messagerie **fire-and-forget** : les publishers émettent sur des canaux, les subscribers reçoivent en temps réel. **Aucune persistance** : un message publié sans abonné est perdu.

```bash
# Terminal 1 — abonné :
redis-cli SUBSCRIBE alertes:infra
# 1) "subscribe"  2) "alertes:infra"  3) (integer) 1
# ... attend les messages ...

# Terminal 2 — publisher :
redis-cli PUBLISH alertes:infra "onduleur: bypass actif"
# → 1 (nombre d'abonnés ayant reçu)

# --- Motifs (pattern subscribe) ---
# Terminal 1 :
redis-cli PSUBSCRIBE "alertes:*"
# Terminal 2 :
redis-cli PUBLISH alertes:reseau "lien down"   # reçu par le PSUBSCRIBE

# --- Sharded Pub/Sub (⚠️ 7.0+, pour le Cluster) ---
redis-cli SSUBSCRIBE commandes:shard1   # le message reste sur le shard de la clé
redis-cli SPUBLISH commandes:shard1 "hello"
```

Limites à connaître :

