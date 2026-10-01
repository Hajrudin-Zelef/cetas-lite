---
id: collect-261001-rattrapage/rattrapage/redis-guide-17
title: "Guide Redis — De l'installation à la production"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["benchmark"]
source: docs/RAG/collect-261001-rattrapage/redis_guide.md
source_anchor: ""
source_lines: [3141, 3203]
sha256: ed0be87ee7604d771c2e79ad5c3afa73e901356d43f0a85ab0a219defd4e66b0
---

# Guide Redis — De l'installation à la production

**Q1. Quelle commande liste les clés sans bloquer le serveur en production ?**
> `SCAN` (avec `MATCH` et `COUNT`). `KEYS *` bloque le thread unique d'exécution.

**Q2. Que se passe-t-il si Redis redémarre sans persistance configurée ?**
> Perte totale des données (tout était en RAM). D'où RDB et/ou AOF en production.

**Q3. `INCR` sur une clé inexistante : que retourne-t-il ?**
> `1` : la clé est créée à 0 puis incrémentée, atomiquement.

**Q4. Quelle est la différence entre `volatile-lru` et `allkeys-lru` ?**
> `volatile-lru` n'évince que les clés **avec TTL** ; `allkeys-lru` peut évincer n'importe quelle clé.

**Q5. Pourquoi `SET verrou NX PX 30000` est-il préférable à `SETNX` + `EXPIRE` ?**
> Atomicité : avec `SETNX`+`EXPIRE` en deux temps, un crash entre les deux laisse un verrou sans TTL (éternel).

**Q6. Un subscriber pub/sub reçoit-il les messages publiés avant sa souscription ?**
> Non. Le pub/sub est fire-and-forget, sans persistance ni rejeu. Pour du durable : les Streams.

**Q7. Combien de hash slots compte un cluster Redis, et comment une clé y est-elle assignée ?**
> 16 384 slots ; `slot = CRC16(clé) mod 16384`. Les hash tags `{...}` forcent la co-localisation.

**Q8. Quel est le quorum recommandé pour 3 Sentinels, et pourquoi un nombre impair ?**
> Quorum = 2 (majorité). Un nombre impair évite les égalités lors d'une partition réseau (split-brain).

**Q9. `appendfsync everysec` : quelle perte de données maximale en cas de crash ?**
> Environ 1 seconde d'écritures (le fsync a lieu une fois par seconde).

**Q10. Votre `INFO stats` montre `evicted_keys` qui augmente vite et un hit rate en chute. Diagnostic et action ?**
> `maxmemory` atteint : Redis évince. Actions : vérifier les big keys (`--bigkeys`), augmenter `maxmemory` si la RAM le permet, ajuster la politique d'éviction / les TTL, ou ajouter de la capacité (réplica/cluster).

---

## 97. Pour aller plus loin

**Documentation officielle**
- https://redis.io/docs/ — documentation complète (commandes, concepts, tuning)
- https://redis.io/docs/latest/operate/ — guide d'exploitation (persistance, réplication, cluster)

**Livres**
- *Redis in Action* (Josiah Carlson) — patterns applicatifs
- *Redis Essentials* — prise en main rapide

**Outils**
- `redis_exporter` (Prometheus) + dashboard Grafana « Redis » (ID 763)
- `redis-stat` — dashboard temps réel léger
- `redli` / `iredis` — clients CLI alternatifs (autocomplétion, coloration)

**Sujets d'approfondissement**
- Redis Functions (⚠️ 7.0+) : alternative persistante aux scripts Lua éphémères
- Redis Query Engine (ex-RediSearch) : recherche plein texte et secondaire
- Active-Active (CRDT, Redis Enterprise) : multi-maîtres multi-sites
- Chiffrement au repos : AOF/RDB sur volume chiffré (LUKS)
- Durcissement : CIS Benchmark for Redis (si applicable à votre audit)

**Prochaines étapes conseillées pour votre SI**
1. Monter un labo : 1 maître + 2 réplicas + 3 Sentinels (VM ou conteneurs).
2. Y faire passer un cas réel (cache catalogue ou sessions).
3. Simuler pannes (kill -9 du maître, partition réseau) et chronométrer le failover.
4. Écrire la procédure d'exploitation (backup, restauration, upgrade) **avant** la prod.

---

*Fin du guide — bon courage, et que vos hit rates soient élevés.* 🚀
