---
id: collect-261001-rattrapage/rattrapage/prometheus-guide-3
title: "Guide Prometheus — Supervision métrique complète"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["datacenter", "memory"]
source: docs/RAG/collect-261001-rattrapage/prometheus_guide.md
source_anchor: ""
source_lines: [432, 656]
sha256: 906d2fc9adf33561a27ae1cce9db6ecd9eb7e205783612bdbac22c6befab79a3
---

# Guide Prometheus — Supervision métrique complète

Le SRV permet de publier `host:port` par service : chaque équipe expose
`_node-exporter._tcp` dans sa zone, Prometheus suit automatiquement.

## 15. Découverte Consul

Si vous utilisez Consul comme registre de services :

```yaml
  - job_name: 'consul-services'
    consul_sd_configs:
      - server: 'consul.dc-paris.lan:8500'
        token: 'votre-token-acl'      # via variable d'env de préférence
        datacenter: 'dc-paris'
        refresh_interval: 30s
    relabel_configs:
      # Ne garder que les services taggés "metrics"
      - source_labels: [__meta_consul_tags]
        regex: '.*,metrics,.*'
        action: keep
      # Utiliser le nom du service comme label "service"
      - source_labels: [__meta_consul_service]
        target_label: service
```

Les méta-labels `__meta_consul_*` (service, node, tags, address…) permettent
un relabeling très fin (voir section 19).

## 16. Découverte Kubernetes

Dans un cluster Kubernetes, Prometheus découvre pods, services, endpoints :

```yaml
  - job_name: 'k8s-pods'
    kubernetes_sd_configs:
      - role: pod
        kubeconfig_file: ''   # in-cluster si vide et RBAC configuré
    relabel_configs:
      # Scraper uniquement les pods annotés prometheus.io/scrape=true
      - source_labels: [__meta_kubernetes_pod_annotation_prometheus_io_scrape]
        regex: 'true'
        action: keep
      - source_labels: [__meta_kubernetes_pod_annotation_prometheus_io_path]
        regex: '(.+)'
        target_label: __metrics_path__
      - source_labels: [__meta_kubernetes_pod_ip, __meta_kubernetes_pod_annotation_prometheus_io_port]
        regex: '([^;]+);([^;]+)'
        replacement: '${1}:${2}'
        target_label: __address__
```

⚠️ En Kubernetes, **toujours** filtrer par annotation/label au relabeling :
scraper tous les pods = explosion de cardinalité.

## 17. Paramètres de scrape avancés

```yaml
  - job_name: 'lent-ou-distant'
    scrape_interval: 60s        # site distant : scraper moins souvent
    scrape_timeout: 50s
    honor_labels: false         # les labels du job écrasent ceux de la cible
    honor_timestamps: true      # faire confiance aux timestamps exposés
    sample_limit: 5000          # coupe le scrape si > 5000 échantillons (anti-cardinalité)
    target_limit: 100           # nombre max de cibles pour ce job
    enable_compression: true    # gzip (défaut true)
```

`sample_limit` est un **garde-fou cardinalité** : si un exporter déraille
(label à haute cardinalité), le scrape est rejeté au lieu de remplir la TSDB.

## 18. Relabeling : principes

Le **relabeling** transforme les labels **avant** le scrape (sur les métadonnées
de découverte) ou **après** (sur les séries ingérées, via `metric_relabel_configs`).

Labels spéciaux (préfixe `__`, supprimés après le scrape sauf réécriture) :

| Label | Signification |
|---|---|
| `__address__` | `host:port` de la cible (modifiable → change la cible scrapée !) |
| `__scheme__` | `http` / `https` |
| `__metrics_path__` | chemin HTTP (défaut `/metrics`) |
| `__meta_*` | métadonnées de la découverte (consul, k8s, ec2…) |

Actions possibles : `replace` (défaut), `keep`, `drop`, `labelmap`,
`labeldrop`, `labelkeep`, `hashmod`, `uppercase`, `lowercase`.

Ordre d'application : les `relabel_configs` sont évaluées **dans l'ordre**,
de haut en bas. L'ordre compte !

## 19. Relabeling : filtrer avec keep et drop

```yaml
    relabel_configs:
      # Ne garder que les cibles dont le label "env" vaut prod ou preprod
      - source_labels: [env]
        regex: 'prod|preprod'
        action: keep

      # Exclure les instances en maintenance (label maintenance=true)
      - source_labels: [maintenance]
        regex: 'true'
        action: drop

      # Exclure un port précis
      - source_labels: [__address__]
        regex: '.*:9115'     # ne pas scraper le blackbox avec ce job
        action: drop
```

> `keep` = liste blanche (tout le reste est jeté), `drop` = liste noire.
> En cas de doute, préférez `keep` : le défaut est de tout scraper.

## 20. Relabeling : réécrire des labels (replace)

```yaml
    relabel_configs:
      # Extraire le hostname depuis __address__ (srv-web01:9100 -> srv-web01)
      - source_labels: [__address__]
        regex: '([^:]+):.*'
        replacement: '${1}'
        target_label: hostname

      # Construire un label "site" depuis le nom DNS
      - source_labels: [__address__]
        regex: '.*\\.([a-z]+)\\.lan:.*'
        replacement: '${1}'
        target_label: site

      # Forcer le port de scrape à 9100 quel que soit le port découvert
      - source_labels: [__address__]
        regex: '([^:]+)(:[0-9]+)?'
        replacement: '${1}:9100'
        target_label: __address__

      # Ajouter un label statique
      - replacement: 'equipe-infra'
        target_label: equipe
```

⚠️ Écrire dans `__address__` **change la cible réellement scrapée** :
c'est la base du pattern blackbox_exporter (section 52).

## 21. metric_relabel_configs : filtrer après scrape

Appliqué **après** le scrape, sur les séries ingérées. Sert à **jeter des
métriques inutiles** avant stockage (économie disque + cardinalité).

```yaml
    metric_relabel_configs:
      # Jeter toutes les métriques Go runtime des exporters (inutiles en supervision)
      - source_labels: [__name__]
        regex: 'go_.*'
        action: drop

      # Jeter les métriques de debug d'un exporter
      - source_labels: [__name__]
        regex: 'process_.*'
        action: drop

      # Ne garder du node_exporter que ce qui sert (exemple restrictif)
      - source_labels: [__name__]
        regex: 'node_cpu_seconds_total|node_memory_.*|node_filesystem_.*|node_network_.*|node_load.*|node_time_seconds|up'
        action: keep
```

⚠️ `metric_relabel_configs` s'applique **après** le scrape réseau : la bande
passante est consommée quand même. Pour réduire aussi le trafic, filtrer côté
exporter (`--collector.disable-defaults`, `--no-collector.*`).

## 22. honor_labels et conflits de labels

Quand une cible expose déjà un label `job` ou `instance` :

```yaml
honor_labels: true   # garde les labels de la CIBLE en cas de conflit
honor_labels: false  # (défaut) les labels du SCRAPE (job/instance) gagnent
```

Cas typique : la **fédération** (section 71) ou le scrape d'un autre Prometheus :
`honor_labels: true` préserve les `job`/`instance` d'origine. Sinon, toutes les
séries fédérées se retrouveraient avec `job="federate"`.

## 23. Vérifier sa configuration : promtool

```bash
# Valider la syntaxe
promtool check config /etc/prometheus/prometheus.yml

# Tester les règles (unit tests !)
promtool check rules /etc/prometheus/rules/*.yml

# Tester une requête PromQL contre une série de test
promtool query instant http://localhost:9090 'up{job="node-linux"}'

# Tester les règles avec des données simulées (fichier test.yml)
promtool test rules test_alertes.yml
```

Exemple de fichier de test `test_alertes.yml` :

```yaml
rule_files:
  - alertes.yml
evaluation_interval: 1m
tests:
  - interval: 1m
    input_series:
      - series: 'up{job="node-linux", instance="srv-web01:9100"}'
        values: 1x10 0x5        # 10x up, puis 5x down
    alert_rule_test:
      - eval_time: 12m
        alertname: InstanceDown
        exp_alerts:
          - exp_labels:
              instance: 'srv-web01:9100'
              severity: 'critical'
```

> Les tests unitaires de règles d'alerte : **peu connu, très précieux**.
> Ils évitent de découvrir une alerte cassée le jour de la panne.

## 24. L'interface web au quotidien

`http://prometheus:9090` — onglets principaux :

