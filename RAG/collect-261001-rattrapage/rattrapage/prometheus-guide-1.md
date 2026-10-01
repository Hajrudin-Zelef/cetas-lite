---
id: collect-261001-rattrapage/rattrapage/prometheus-guide-1
title: "Guide Prometheus — Supervision métrique complète"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["agent", "datacenter", "memory", "valuation"]
source: docs/RAG/collect-261001-rattrapage/prometheus_guide.md
source_anchor: ""
source_lines: [1, 183]
sha256: 1d2b99c3cbdf0a0ec1fc00cd78dd7e488d5ec0e70e348d0ee9982fab24f5aee6
---

# Guide Prometheus — Supervision métrique complète

> **Public cible** : chefs de service systèmes, administrateurs systèmes & réseaux,
> responsables d'exploitation. **Version** : Prometheus 2.5x / 3.x (concepts stables
> sur toutes les 2.x). **Date** : septembre 2026.

---

## 1. Pourquoi ce guide

Prometheus est devenu le standard de facto de la supervision métrique dans les
infrastructures modernes : serveurs Linux, équipements réseau, bases de données,
applications. Pour un chef de service systèmes, c'est l'outil qui permet de
passer d'une supervision réactive ("le serveur est tombé, on nous appelle") à
une supervision prédictive ("le disque sera plein dans 12 jours, on planifie").

Ce guide couvre tout le cycle de vie : installation, configuration, PromQL
(le langage de requêtes), les exporters (node, blackbox, SNMP — y compris les
onduleurs), Alertmanager, la haute disponibilité, la sauvegarde, le dépannage,
et les pièges à éviter (la cardinalité !).

> **Convention de lecture** : les blocs `yaml` sont des extraits de configuration,
> les blocs `promql` (ou sans langage) sont des requêtes à tester dans l'interface
> web Prometheus (menu **Graph**). Les ⚠️ signalent les pièges de production.

---

## 2. Architecture générale de Prometheus

Prometheus est un **système de monitoring** qui fonctionne sur un principe
simple : il **va chercher** (pull) les métriques sur les cibles, les **stocke**
dans sa base de données temporelle (TSDB), permet de les **interroger** avec
PromQL, et **déclenche des alertes** via Alertmanager.

```
 ┌──────────────┐   scrape (pull)   ┌──────────────────┐
 │ node_exporter│◄──────────────────┤                  │
 ├──────────────┤                   │                  │
 │blackbox_exp. │◄──────────────────┤    PROMETHEUS    │──► Alertmanager ──► email/Telegram/webhook
 ├──────────────┤                   │  (TSDB + PromQL) │
 │ snmp_exporter│◄──────────────────┤                  │
 ├──────────────┤                   │                  │
 │  application │◄──────────────────┤                  │
 │  (instrument)│                   └──────────────────┘
 └──────────────┘                              │
                                               ▼
                                        Grafana (dashboards)
```

Composants :

| Composant | Rôle |
|---|---|
| **Prometheus server** | Scrape, stockage TSDB, évaluation des règles, API HTTP |
| **Exporters** | Exposent les métriques d'un système au format texte sur `/metrics` |
| **Alertmanager** | Reçoit les alertes, les déduplique, groupe, route vers les bons canaux |
| **Pushgateway** | Point d'entrée pour les jobs batch (push) qui ne vivent pas assez longtemps pour être scrapés |
| **Grafana** | Visualisation (dashboards), s'appuie sur l'API PromQL |

## 3. Le modèle pull : comprendre la philosophie

Contrairement à Nagios/Zabbix (modèle push ou checks actifs centralisés),
Prometheus **initie la connexion** vers chaque cible à intervalle régulier
(par défaut toutes les 15 s) sur un endpoint HTTP `/metrics`.

Avantages du pull :

- **Découverte de service simple** : ajouter une cible = ajouter une ligne de config.
- **Santé intrinsèque** : si le scrape échoue, la métrique `up{job="..."}` passe
  à 0 → on sait immédiatement que la cible ne répond plus, sans agent tiers.
- **Pas d'agent à configurer côté serveur** : l'exporter expose, Prometheus vient.
- **Débogage facile** : `curl http://cible:9100/metrics` montre exactement ce que
  Prometheus voit.

Inconvénients :

- Derrière un NAT ou un firewall strict, il faut ouvrir le flux **vers** la cible
  (ou utiliser le Pushgateway / un remote write).
- Pas adapté aux jobs éphémères (< intervalle de scrape) → Pushgateway.

## 4. La TSDB : comment Prometheus stocke les données

La **TSDB** (Time Series Database) embarquée stocke chaque série temporelle
sous forme d'échantillons `(timestamp, valeur)` compressés.

Concepts clés :

- **Série temporelle** = une métrique + un jeu unique de labels.
  Exemple : `node_cpu_seconds_total{cpu="0",mode="idle"}` est une série.
- **Chunks** : les échantillons récents vivent en mémoire (head chunk), puis
  sont flushés sur disque en blocs de 2 h, compactés ensuite.
- **Rétention** : par défaut **15 jours** (`--storage.tsdb.retention.time=15d`)
  ou 512 Mo (`--storage.tsdb.retention.size`). Le premier atteint déclenche la purge.
- **WAL** (Write-Ahead Log) : journal d'écriture pour rejouer les données
  récentes après un crash.

Répertoire de données (par défaut `./data`, configurable) :

```
data/
├── wal/            # journal d'écriture (données < ~2h)
├── chunks_head/    # chunks en mémoire persistés
└── 01J.../         # blocs TSDB (2h chacun, avec index + chunks)
```

⚠️ **Ne jamais modifier ces fichiers à la main.** Pour sauvegarder, utiliser
l'API de snapshots (voir section 74).

## 5. Modèle de données : métriques et labels

Toute donnée Prometheus est une **série temporelle** identifiée par :

1. Un **nom de métrique** : `http_requests_total`, `node_memory_MemTotal_bytes`.
2. Zéro ou plusieurs **labels** (paires clé=valeur) : `{method="GET",status="200"}`.

Le nom + l'ensemble des labels = **l'identité unique** de la série.

```promql
# Une série précise :
http_requests_total{method="POST", handler="/api/login", status="500"}

# Toutes les séries de cette métrique (quel que soit le label) :
http_requests_total
```

Règles de nommage :

- Noms en `snake_case`, en anglais, avec **unité en suffixe** si applicable :
  `_seconds`, `_bytes`, `_total` (pour les compteurs).
- Labels en `snake_case` : `instance`, `job`, `datacenter`, `host`.
- Les labels `job` et `instance` sont ajoutés automatiquement par Prometheus
  lors du scrape (configurables via relabeling).

## 6. Les 4 types de métriques

### 6.1 Counter (compteur)

Valeur qui **ne fait qu'augmenter** (ou reste stable), remise à zéro au redémarrage
du processus. Exemples : nombre de requêtes HTTP, octets réseau reçus.

```promql
# ❌ MAUVAIS : la valeur brute d'un compteur ne veut rien dire
http_requests_total

# ✅ BON : le débit par seconde sur 5 minutes
rate(http_requests_total[5m])
```

⚠️ Toujours utiliser `rate()`, `irate()` ou `increase()` sur un compteur,
jamais la valeur brute ni `delta()`.

### 6.2 Gauge (jauge)

Valeur qui **peut monter et descendre** : température, mémoire utilisée,
nombre de connexions actives, niveau de batterie d'onduleur.

```promql
# ✅ Directement exploitable
node_memory_MemAvailable_bytes
ups_battery_charge_percent
```

### 6.3 Histogram

Échantillonne des observations (tailles de requêtes, durées) dans des
**buckets** configurables et expose :

- `*_bucket{le="..."}` : compteurs cumulés par borne supérieure (`le` = less or equal),
- `*_sum` : somme des observations,
- `*_count` : nombre d'observations.

```promql
# Quantile p95 de la latence HTTP sur 5 min
histogram_quantile(0.95, rate(http_request_duration_seconds_bucket[5m]))
```

⚠️ Les buckets sont **cumulatifs** : `le="0.5"` inclut tout ce qui est ≤ 0,5 s.

### 6.4 Summary

Comme l'histogramme, mais calcule les **quantiles côté client** (dans l'application)
et expose `*_quantile{quantile="0.95"}`, `*_sum`, `*_count`.

