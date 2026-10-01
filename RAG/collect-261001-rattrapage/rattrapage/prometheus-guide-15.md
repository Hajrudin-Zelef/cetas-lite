---
id: collect-261001-rattrapage/rattrapage/prometheus-guide-15
title: "Guide Prometheus — Supervision métrique complète"
domain: rattrapage
role: reference
task: reference
actors: ["Google"]
dates: []
keywords: ["distribution"]
source: docs/RAG/collect-261001-rattrapage/prometheus_guide.md
source_anchor: ""
source_lines: [3112, 3255]
sha256: 65d523296b30368a5c3ceae0c8e1ffb51fe627a290dd9cc20ae5bf2a2cc6a4ff
---

# ── Cardinalité ────────────────────────────────────────
prometheus_tsdb_head_series    # séries actives
topk(10, count by (__name__) ({__name__=~".+"}))
```

## 101. Glossaire

| Terme | Définition |
|---|---|
| **Alerte** | Règle dont l'expression est vraie → envoyée à Alertmanager |
| **Alertmanager** | Composant qui déduplique, groupe, route et notifie les alertes |
| **Cardinalité** | Nombre de séries uniques ; le facteur limitant n°1 |
| **Chunk** | Bloc d'échantillons compressés dans la TSDB |
| **Compaction** | Fusion des blocs TSDB de 2 h en blocs plus gros |
| **Counter** | Compteur qui ne fait qu'augmenter (requêtes, octets) |
| **Exporter** | Programme qui expose des métriques sur `/metrics` |
| **Fédération** | Un Prometheus scrape les métriques d'un autre |
| **Gauge** | Jauge qui monte et descend (température, mémoire) |
| **Grouping** | Regroupement d'alertes similaires en une notification |
| **Head** | Partie en mémoire de la TSDB (données récentes) |
| **Histogram** | Distribution d'observations en buckets cumulés |
| **Inhibition** | Une alerte grave coupe les notifications des alertes dérivées |
| **Instance** | Label `host:port` d'une cible scrapée |
| **Job** | Ensemble de cibles scrapées avec la même config |
| **Label** | Paire clé=valeur identifiant une série |
| **mTLS** | TLS mutuel : client et serveur s'authentifient |
| **PromQL** | Langage de requêtes de Prometheus |
| **Pushgateway** | Point d'entrée pour les jobs batch (push) |
| **Range vector** | Série de valeurs sur une durée `[5m]` |
| **Recording rule** | Requête pré-calculée et stockée comme métrique |
| **Relabeling** | Réécriture des labels avant/après scrape |
| **Remote read/write** | Lecture/écriture vers un stockage distant |
| **Rétention** | Durée/taille de conservation des données |
| **Scrape** | Action d'aller chercher les métriques (pull) |
| **Série temporelle** | Métrique + jeu unique de labels |
| **Service Discovery (SD)** | Découverte automatique des cibles |
| **Silence** | Coupure volontaire et tracée des notifications |
| **SLI / SLO** | Indicateur / objectif de niveau de service |
| **SNMP** | Protocole de supervision des équipements réseau |
| **Subquery** | Requête PromQL imbriquée `[durée:pas]` |
| **Summary** | Histogram-like avec quantiles calculés côté client |
| **Target** | Cible à scraper (`host:port`) |
| **TSDB** | Base de données temporelle embarquée |
| **up** | Métrique de santé du scrape (1 = OK, 0 = échec) |
| **Vector matching** | Jointure de vecteurs sur les labels (`on`/`ignoring`) |
| **WAL** | Write-Ahead Log : journal des écritures récentes |

## 102. Quiz : validez vos acquis (10 questions + réponses)

**Q1.** Quelle est la différence fondamentale entre un `counter` et une `gauge` ?
Pourquoi ne doit-on jamais faire `rate()` sur une gauge ?
> **R1.** Le counter ne fait qu'augmenter (remise à zéro au reboot) ; la gauge
> monte et descend. `rate()` suppose un compteur monotone : sur une gauge, les
> baisses "normales" seraient interprétées comme des remises à zéro et fausseraient
> le calcul. Sur une gauge on utilise `delta()`, `deriv()` ou la valeur brute.

**Q2.** Votre alerte `up == 0` sans `for:` vous réveille à 3 h du matin alors que
tout va bien. Que s'est-il passé et comment corriger ?
> **R2.** Un scrape isolé a échoué (micro-coupure réseau, GC de l'exporter) et
> l'alerte est partie immédiatement. Correction : `for: 5m` — l'alerte ne tire
> que si la cible reste down 5 minutes.

**Q3.** Que se passe-t-il si vous ajoutez un label `user_id` (10 000 valeurs) à
`http_requests_total` qui a déjà 3 labels à 10 valeurs chacun ?
> **R3.** Explosion de cardinalité : 10 000 × 10³ = 10 millions de séries
> potentielles → RAM saturée, OOM probable. Il faut normaliser (route générique)
> ou agréger côté application, jamais d'identifiant unique en label.

**Q4.** Dans quel ordre s'appliquent `relabel_configs` et `metric_relabel_configs`,
et à quoi sert chacun ?
> **R4.** `relabel_configs` d'abord (avant le scrape, sur les métadonnées de
> découverte : filtrer/réécrire les cibles). `metric_relabel_configs` ensuite
> (après le scrape, sur les séries : jeter des métriques avant stockage).
> Le premier économise aussi la bande passante réseau, pas le second.

**Q5.** Écrivez la requête du p95 de latence HTTP par route sur 5 minutes.
> **R5.** `histogram_quantile(0.95, sum by (route, le) (rate(http_request_duration_seconds_bucket[5m])))`
> Points clés : `rate()` à l'intérieur sur les buckets (compteurs), `sum by (le)`
> avant le quantile, `le` conservé.

**Q6.** À quoi sert `honor_labels: true` et dans quel cas est-il obligatoire ?
> **R6.** Il conserve les labels (`job`, `instance`…) exposés par la cible en cas
> de conflit avec ceux du scrape. Obligatoire en **fédération** : sans lui,
> toutes les séries fédérées se retrouvent avec `job="federate"` et on perd
> l'origine.

**Q7.** Votre disque Prometheus est plein à 98 %. Donnez les 3 premières actions,
dans l'ordre.
> **R7.** 1) Vérifier ce qui consomme (`df`, `du`) et purger les vieux snapshots
> oubliés dans `snapshots/`. 2) Réduire temporairement la rétention
> (`retention.time`) et redémarrer pour purger. 3) En dernier recours, à froid,
> supprimer les plus vieux blocs TSDB (jamais `wal/` à chaud).

**Q8.** Quelle est la différence entre une inhibition et un silence dans
Alertmanager ? Quand utiliser chacun ?
> **R8.** L'**inhibition** est automatique et permanente (règle de config) : une
> alerte grave coupe les alertes dérivées (panne site → pas d'alerte par serveur).
> Le **silence** est manuel et temporaire : on coupe les notifications pendant
> une maintenance planifiée, avec auteur, commentaire et expiration.

**Q9.** Pourquoi `predict_linear(node_filesystem_avail_bytes[1h], 86400)` est-il
une mauvaise idée en alerte ?
> **R9.** Fenêtre trop courte (1 h) : la prédiction est ultra-sensible aux
> à-coups (gros fichier temporaire) → faux positifs. Il faut au minimum `[6h]`,
> et garder en tête que la tendance est supposée linéaire.

**Q10.** Vous avez deux Prometheus en HA qui scrapent les mêmes cibles. Pourquoi
faut-il des `external_labels` différents et un cluster Alertmanager ?
> **R10.** Les `external_labels` (`replica: prom-01/02`) distinguent les deux
> sources dans la fédération et le debug. Le cluster Alertmanager (gossip)
> déduplique : sans lui, chaque Prometheus enverrait ses propres notifications
> → chaque alerte reçue en double.

## 103. Pour aller plus loin

**Documentation officielle**

- https://prometheus.io/docs/introduction/overview/ — concepts
- https://prometheus.io/docs/prometheus/latest/querying/basics/ — PromQL
- https://prometheus.io/docs/alerting/latest/configuration/ — Alertmanager
- https://prometheus.io/docs/operating/integrations/ — liste des exporters

**Dashboards Grafana (IDs à importer)**

- 1860 — Node Exporter Full
- 11074 — Blackbox Exporter
- 3662 — Prometheus 2.0 Stats (santé du serveur lui-même)

**Livres**

- *Prometheus: Up & Running* (O'Reilly) — la référence, à jour sur les concepts
- *Site Reliability Engineering* (Google, gratuit en ligne) — les chapitres
  alerting et SLO, la doctrine derrière les bonnes pratiques

**Prochaines étapes concrètes**

1. Monter une maquette : 1 Prometheus + 3 node_exporters + 1 blackbox en une
   après-midi (sections 7, 8, 45, 51, 52).
2. Brancher Alertmanager avec email + Telegram (sections 60-63) et tester une
   vraie alerte de bout en bout.
3. Ajouter le SNMP pour 2 switchs et 1 onduleur (sections 54-56).
4. Mettre en place la paire HA + snapshots quotidiens (sections 74, 76).
5. Revue des alertes après 1 mois : supprimer le bruit, ajuster les seuils.

