---
id: collect-261001-rattrapage/rattrapage/prometheus-guide-4
title: "Guide Prometheus — Supervision métrique complète"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["attention", "datacenter", "incident", "memory"]
source: docs/RAG/collect-261001-rattrapage/prometheus_guide.md
source_anchor: ""
source_lines: [657, 875]
sha256: 286856e1016adcc2001dd6a86665a0fb7e6a8dac78715454f614efd8b39b8cfd
---

# Guide Prometheus — Supervision métrique complète

| Onglet | Usage |
|---|---|
| **Graph** | Tester des requêtes PromQL (table + graphique) |
| **Alerts** | État des règles d'alerte (pending / firing) |
| **Targets** | État de santé de chaque cible (up/down, dernier scrape, erreur) |
| **Rules** | Recording rules évaluées |
| **Service Discovery** | Cibles découvertes par mécanisme, avant relabeling |
| **TSDB Status** | Cardinalité : top 10 des métriques les plus lourdes |
| **Flags / Config** | Vérifier la config chargée et les flags |

Réflexe quotidien : **Status → Targets** : tout doit être vert (`UP`).
Une cible `DOWN` = alerte `InstanceDown` si vos règles sont bien faites
(section 70).

---

---

# PARTIE II — PromQL : le langage de requêtes

---

## 25. PromQL en 5 minutes

PromQL manipule des **vecteurs** :

- **Vecteur instantané** : un ensemble de séries, une valeur chacune, à un instant T.
- **Vecteur de plage (range vector)** : un ensemble de séries, plusieurs valeurs
  sur une durée `[5m]`. Ne peut pas être affiché directement : il faut lui
  appliquer une fonction (`rate()`, `avg_over_time()`…) qui le transforme en
  vecteur instantané.

```promql
node_cpu_seconds_total                    # vecteur instantané (toutes les séries)
node_cpu_seconds_total[5m]                 # range vector (5 min d'historique)
rate(node_cpu_seconds_total[5m])          # vecteur instantané : débit/sec par série
```

## 26. Sélecteurs et filtres de labels

```promql
# Égalité exacte
node_filesystem_avail_bytes{mountpoint="/", fstype="ext4"}

# Négation
node_filesystem_avail_bytes{mountpoint!="/boot"}

# Regex (opérateur =~)
node_network_receive_bytes_total{device=~"eth[0-9]+"}

# Regex négative
node_network_receive_bytes_total{device!~"lo|docker.*|veth.*"}

# Plusieurs filtres (ET logique implicite)
up{job="node-linux", datacenter="dc-paris"}
```

⚠️ `{job="..."}` sans nom de métrique = toutes les séries ayant ce label.
Utile en exploration, **interdit dans les dashboards** (trop large).

## 27. Offset : regarder dans le passé

```promql
# Valeur il y a 1 heure (comparaison avant/après incident)
node_memory_MemAvailable_bytes offset 1h

# Débit d'il y a 7 jours, même heure (comparaison semaine/semaine)
rate(http_requests_total[5m] offset 7d)

# Offset aussi sur les range vectors
avg_over_time(node_load1[1h] offset 24h)
```

Cas d'usage roi : *"est-ce que c'est pire que la semaine dernière à la même heure ?"*

## 28. Opérateurs arithmétiques

`+ - * / % ^` s'appliquent élément par élément entre vecteurs (appariés par labels)
ou entre vecteur et scalaire.

```promql
# Mémoire utilisée en Gio
(node_memory_MemTotal_bytes - node_memory_MemAvailable_bytes) / 1024^3

# Pourcentage d'utilisation disque
100 * (1 - node_filesystem_avail_bytes / node_filesystem_size_bytes)

# Conversion octets -> Gbit/s sur le réseau
rate(node_network_receive_bytes_total[5m]) * 8 / 1e9
```

⚠️ Division par zéro : PromQL retourne **NaN**, pas d'erreur. Filtrer si besoin :
`... / clamp_min(node_filesystem_size_bytes, 1)`.

## 29. Opérateurs de comparaison et `bool`

Par défaut, une comparaison **filtre** (garde les séries où c'est vrai, jette les autres).
Avec `bool`, elle retourne **0 ou 1**.

```promql
# Séries dont l'utilisation disque dépasse 85 % (filtre)
(100 * (1 - node_filesystem_avail_bytes / node_filesystem_size_bytes)) > 85

# Retourne 1 si critique, 0 sinon (utile pour des métriques d'état)
(node_filesystem_avail_bytes / node_filesystem_size_bytes) < bool 0.1
```

Opérateurs : `== != > < >= <=`. Attention, `=` seul n'existe pas en PromQL.

## 30. Opérateurs logiques : and, or, unless

- `and` : intersection (séries présentes des deux côtés, valeur du côté gauche).
- `or` : union (toutes les séries, valeur du côté où elle existe).
- `unless` : séries de gauche absentes à droite.

```promql
# Disques pleins à plus de 85 % ET montés (exclure les pseudo-fs)
(100 * (1 - node_filesystem_avail_bytes / node_filesystem_size_bytes)) > 85
  and on (instance, device, mountpoint) node_filesystem_readonly == 0

# Alertes "disque plein" OU "inode épuisés"
(disque_plein > 85) or (inodes_pleins > 90)

# Cibles down SAUF celles en maintenance planifiée
up == 0 unless on (instance) maintenance_mode == 1
```

## 31. Vector matching : on() et ignoring()

Quand les deux côtés n'ont pas exactement les mêmes labels, précisez la clé
de jointure :

```promql
# Pourcentage de CPU idle par instance : les deux métriques partagent "instance"
avg by (instance) (rate(node_cpu_seconds_total{mode="idle"}[5m]))
  / on (instance) group_left
avg by (instance) (rate(node_cpu_seconds_total[5m]))
```

- `on (labels)` : n'apparier **que** sur ces labels.
- `ignoring (labels)` : apparier sur tous les labels **sauf** ceux-là.
- `group_left` / `group_right` : jointure **plusieurs-vers-un**
  (cardinalité many-to-one), obligatoire sinon erreur.

⚠️ Erreur classique : `vector1 / vector2` sans matching alors que les labels
diffèrent → **résultat vide**, sans message d'erreur. Toujours vérifier qu'on
obtient des données.

## 32. rate(), irate(), increase() : bien les utiliser

```promql
rate(http_requests_total[5m])      # débit moyen/sec sur 5 min (lissé)
irate(http_requests_total[5m])     # débit instantané (2 derniers points)
increase(http_requests_total[1h])  # augmentation totale sur 1 h
```

| Fonction | Usage recommandé |
|---|---|
| `rate()` | Dashboards, alertes (lisse les à-coups) |
| `irate()` | Graphiques volatiles, debug en direct |
| `increase()` | "Combien d'erreurs sur la dernière heure ?" |

```promql
# Erreurs HTTP 5xx par seconde, par service
sum by (service) (rate(http_requests_total{status=~"5.."}[5m]))

# Nombre total de redémarrages de conteneurs sur 24 h
increase(kube_pod_container_status_restarts_total[24h])
```

⚠️ **Fenêtre minimale** : `[5m]` avec un scrape à 15 s = ~20 points, bon lissage.
En dessous de 4× l'intervalle de scrape, `rate()` devient bruité. Pour les
alertes, préférez `[5m]` ou `[10m]`, jamais `[1m]`.

⚠️ `rate()` sur un **compteur qui se réinitialise** (reboot) : Prometheus détecte
la remise à zéro et lisse. `irate()` peut produire un pic artificiel.

## 33. Fonctions sur range vectors (`*_over_time`)

Transforment un range vector en vecteur instantané :

```promql
avg_over_time(node_load1[1h])        # charge moyenne sur 1 h
max_over_time(cpu_temp[30m])         # pic de température sur 30 min
min_over_time(node_memory_MemAvailable_bytes[1h])
sum_over_time(increase_erreurs[24h]) # (rare : préférer increase)
count_over_time(up[5m])              # nombre de scrapes réussis sur 5 min
last_over_time(metrique[5m])         # dernière valeur (utile après absent)
stddev_over_time(latence[1h])        # écart-type : détecter l'instabilité
stdvar_over_time(latence[1h])
quantile_over_time(0.95, latence[1h])# p95 (coûteux : préférer histogram_quantile)
```

```promql
# Disponibilité d'un service sur 30 jours (en %)
100 * avg_over_time(up{job="api"}[30d])
```

## 34. delta(), idelta(), deriv(), predict_linear()

```promql
delta(temperature_celsius[1h])              # variation absolue sur 1 h (gauges !)
idelta(temperature_celsius[5m])             # variation entre 2 derniers points
deriv(temperature_celsius[1h])              # dérivée/sec (tendance)
predict_linear(disk_free_bytes[6h], 86400)  # prévision dans 24 h (86400 s)
```

**Le cas d'usage star** — prédire un disque plein :

```promql
# Heures restantes avant saturation du disque (si tendance linéaire)
(node_filesystem_avail_bytes / deriv(node_filesystem_avail_bytes[6h])) / 3600
# Alerte si < 72 h restantes :
predict_linear(node_filesystem_avail_bytes{fstype!~"tmpfs|overlay"}[6h], 3*24*3600) < 0
```

⚠️ `predict_linear` suppose une tendance **linéaire** : parfait pour un disque
qui se remplit régulièrement, faux pour des à-coups. Fenêtre `[6h]` minimum.

