---
id: collect-261001-rattrapage/rattrapage/loki-guide-6
title: "Grafana Loki — Le guide complet"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["agents", "apache", "latency", "valuation"]
source: docs/RAG/collect-261001-rattrapage/loki_guide.md
source_anchor: ""
source_lines: [1278, 1547]
sha256: 9165086a278936fd4d784f29489a352c99b08b9cdf0a0ae99ad34dc0ae10441e
---

# Tous les serveurs web sauf web03
{job="nginx", host=~"web0[12]"}

# Tout sauf la préproduction
{env!="preprod"}

# Plusieurs jobs : regex ou alternance
{job=~"nginx|apache"}
```

> 💡 **Performance** : un sélecteur avec au moins un `=` exact est
> idéal. Les regex `=~` sur les labels restent indexées mais coûtent
> plus cher. Évitez `{job=~".+"}` (tous les flux !).

---

## 31. Opérateurs de filtrage : `|=`, `!=`, `|~`, `!~`

Après le sélecteur, les **filtres de ligne** travaillent sur le texte brut :

| Opérateur | Sens | Exemple |
|---|---|---|
| `\|=` | contient | `{job="nginx"} \|= "error"` |
| `!=` | ne contient pas | `{job="nginx"} != "healthcheck"` |
| `\|~` | matche la regex | `{job="nginx"} \|~ "5\d\d"` |
| `!~` | ne matche pas | `{job="nginx"} !~ "timeout"` |

Chaînage (évaluation de gauche à droite) :

```logql
# Erreurs 500 sur /api, hors healthchecks
{job="nginx"} |= "/api" |~ " 5[0-9][0-9] " != "healthcheck"
```

**Ordre = performance** : placez d'abord le filtre le plus sélectif.
`|= "error"` (chaîne littérale, rapide) avant `|~ "5\d\d"` (regex,
plus coûteuse).

> ⚠️ Ces filtres **ne sont pas indexés** : Loki décompresse les chunks
> des flux sélectionnés puis filtre. D'où l'importance d'un sélecteur
> de flux resserré (section 39).

---

## 32. Filtrer sur les labels extraits par les parsers

Les parsers (`| json`, `| logfmt`, `| regexp`, `| pattern`) extraient des
champs utilisables comme des labels temporaires :

```logql
# Logs JSON : erreurs avec latence > 1s
{job="api"} | json | level="error" | latency_ms > 1000

# logfmt
{job="monapp"} | logfmt | level="error" | retries > 2

# Regex à la volée (si rien n'est parsé à l'ingestion)
{job="nginx"} | regexp `(?P<status>\d{3})` | status="500"

# Pattern : syntaxe simplifiée (<champ>)
{job="sshd"} | pattern `<_> Failed password for <user> from <ip> port <port>`
```

Opérateurs de comparaison sur les champs :

| Opérateur | Sens |
|---|---|
| `=`, `!=` | égalité (chaîne) |
| `>`, `>=`, `<`, `<=` | comparaison numérique |
| `=~`, `!~` | regex sur le champ |

Exemple combiné (nginx) :

```logql
{job="nginx"}
  | regexp `^(?P<ip>\S+) \S+ \S+ \[[^\]]+\] "(?P<method>\S+) (?P<path>[^ ]*)[^"]*" (?P<status>\d{3})`
  | status="500"
  | path=~"/api/.*"
```

---

## 33. `rate()` : mesurer un taux d'erreur

`rate()` calcule le **nombre de lignes par seconde** sur une fenêtre :

```logql
# Logs nginx par seconde (tous)
rate({job="nginx"}[5m])

# Erreurs 500 par seconde
rate({job="nginx"} | regexp `(?P<status>\d{3})` | status="500" [5m])

# Pourcentage d'erreurs 5xx : le classique SLI
sum(rate({job="nginx"} | regexp `(?P<status>\d{3})` | status=~"5.." [5m]))
/
sum(rate({job="nginx"}[5m]))
* 100
```

Variantes :

| Fonction | Sens |
|---|---|
| `rate(v[5m])` | lignes/seconde (lisse les bords) |
| `count_over_time(v[5m])` | nombre brut de lignes sur 5 min |
| `bytes_rate(v[5m])` | octets/seconde (utile pour le volume) |
| `bytes_over_time(v[5m])` | octets totaux sur la fenêtre |

> 💡 Pour une alerte « trop d'erreurs », préférez `rate()` (lissé) à
> `count_over_time()` (bruité). Voir section 41.

---

## 34. `count_over_time()`, `sum`, `avg` et compagnie

```logql
# Nombre de lignes par serveur sur 1h
sum by (host) (count_over_time({job="syslog"}[1h]))

# Top 5 des hosts les plus bavards
topk(5, sum by (host) (count_over_time({job="syslog"}[1h])))

# Latence moyenne par endpoint (logs JSON avec latency_ms)
avg by (path) (
  avg_over_time({job="api"} | json | unwrap latency_ms [5m])
)

# p99 de latence
quantile_over_time(0.99, {job="api"} | json | unwrap latency_ms [5m])
```

`unwrap` : extrait une valeur numérique d'un champ pour les fonctions
`_over_time` (`avg_over_time`, `max_over_time`, `sum_over_time`,
`quantile_over_time`, `stdvar_over_time`, `stddev_over_time`).

```logql
# Trafic total en octets par vhost nginx sur 24h
sum by (vhost) (
  sum_over_time({job="nginx"} | regexp `(?P<bytes>\d+)$` | unwrap bytes [24h])
)
```

---

## 35. Agréger par labels : `sum by (...)`

Les agrégations PromQL s'appliquent aux vecteurs issus des fonctions :

```logql
# Taux d'erreur par host
sum by (host) (rate({job="nginx"} |~ " 5\d\d " [5m]))

# Sans le 'by' : taux global
sum(rate({job="nginx"} |~ " 5\d\d " [5m]))

# Erreurs par minute et par service applicatif
sum by (service) (count_over_time({job="apps", level="error"}[1m]))

# Ratio d'erreurs par service (pour un tableau de bord)
sum by (service) (rate({job="apps", level="error"}[5m]))
/
sum by (service) (rate({job="apps"}[5m]))
```

> ⚠️ `sum by (label_trop_cardinal)` (ex : `by (request_id)`) génère
> autant de séries que de valeurs → risque de surcharge du frontend.
> Agrégez toujours sur des labels à faible cardinalité.

---

## 36. Requêtes métriques depuis les logs

Tout dashboard ou alerte part d'une **requête métrique** : des logs
transforment en série temporelle. Recettes courantes :

```logql
# 1. Débit de logs par job (volume d'ingestion par source)
sum by (job) (rate({job=~".+"}[5m]))

# 2. Taux d'erreur applicative global
sum(rate({job="apps", level="error"}[5m]))

# 3. Erreurs SSH par minute (brute force ?)
sum(count_over_time({job="syslog"} |= "Failed password" [1m]))

# 4. Requêtes HTTP par code de statut (nginx parsé à l'ingestion)
sum by (status) (rate({job="nginx"}[5m]))

# 5. Volume de logs en Mo/heure par host (capacité)
sum by (host) (bytes_over_time({job=~".+"}[1h])) / 1024 / 1024
```

Astuce : pré-calculez les métriques critiques avec des **règles
d'enregistrement** du ruler (section 40) pour des dashboards instantanés.

---

## 37. Détecter l'absence de logs (dead man's switch)

Un serveur qui n'envoie plus de logs est souvent un serveur en panne.
LogQL détecte l'absence avec `absent_over_time` ou un `count == 0` :

```logql
# Alerte si aucun log reçu d'un host depuis 10 min
# (à évaluer avec le ruler, section 40)
absent_over_time({job="syslog", host="srv01"}[10m])
```

> ⚠️ `absent_over_time` renvoie 1 quand **aucune ligne** n'existe sur la
> fenêtre. En règle d'alerte, combinez avec un `for: 10m` pour éviter les
> faux positifs lors des redémarrages Promtail.

Alternative sans `absent_over_time` (compatible vieux Loki) :

```logql
# Nombre de hosts actifs vus dans l'heure — à comparer à l'inventaire
count(count_over_time({job="syslog"}[1h]) by (host))
```

Si ce nombre chute, un ou plusieurs agents Promtail sont tombés.

---

## 38. Sous-requêtes et fenêtres temporelles

La fenêtre `[5m]` dans `rate(...[5m])` s'appelle un **range vector**.
Règles :

- Fenêtre trop courte (< 2× l'intervalle de scrape) → résultats bruités
  ou vides.
- Fenêtre trop longue → lissage excessif, alertes en retard.
- Pour les alertes : `[5m]` à `[15m]` est le standard.

```logql
# Comparer le taux d'erreur actuel à celui d'il y a 1h (dérive)
sum(rate({job="api", level="error"}[5m]))
/
sum(rate({job="api", level="error"}[5m] offset 1h))
```

Décalages temporels (`offset`) et comparaisons semaine/semaine :

```logql
# Le trafic d'aujourd'hui vs il y a 7 jours
sum(rate({job="nginx"}[5m]))
/
sum(rate({job="nginx"}[5m] offset 7d))
```

---

## 39. Optimiser les requêtes LogQL lentes

Une requête lente vient presque toujours d'un **sélecteur trop large**.
Méthode en 4 étapes :

1. **Resserrer les labels** : ajoutez `host`, `env`, `level` au sélecteur.
   ```logql
   -- {job=~".+"} |= "error"           -- lent : tous les flux
   ++ {job="nginx", level="error"}     -- rapide : index utilisé
   ```
2. **Réduire la plage** : 1 h plutôt que 7 j pour explorer, puis élargir.
3. **Filtres littéraux d'abord** : `|= "timeout"` avant `| json`
   (moins de lignes à parser).
4. **Limiter** : ajoutez `| line_format` seulement à la fin, utilisez le
   sélecteur de limite de Grafana (ex : 1000 lignes).

Réglages serveur contre les requêtes tueuses (`limits_config`) :

