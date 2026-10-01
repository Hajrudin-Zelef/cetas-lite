---
id: collect-261001-rattrapage/rattrapage/prometheus-guide-2
title: "Guide Prometheus — Supervision métrique complète"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["arr", "datacenter", "valuation"]
source: docs/RAG/collect-261001-rattrapage/prometheus_guide.md
source_anchor: ""
source_lines: [184, 431]
sha256: 24e65298d304a2014e76064326c11295b26f8506f52ff434ec51d6e851f0e005
---

# Guide Prometheus — Supervision métrique complète

| | Histogram | Summary |
|---|---|---|
| Quantiles calculés | côté serveur (PromQL) | côté client |
| Agrégeable (`sum` sur plusieurs instances) | ✅ oui | ❌ non |
| Coût | buckets à dimensionner | calculé en continu |

**Recommandation** : préférer **Histogram** pour tout ce qui doit s'agréger
(par service, par datacenter). Summary uniquement pour du debug local.

## 7. Installation : méthode binaire (Debian/Ubuntu)

La méthode la plus simple et la plus robuste en production : binaire statique
officiel, aucune dépendance.

```bash
# 1. Variables
PROM_VERSION="3.0.1"
ARCH="linux-amd64"   # ou linux-arm64 pour Raspberry Pi / ARM

# 2. Téléchargement et vérification (depuis releases GitHub officielles)
cd /tmp
wget "https://github.com/prometheus/prometheus/releases/download/v${PROM_VERSION}/prometheus-${PROM_VERSION}.${ARCH}.tar.gz"
wget "https://github.com/prometheus/prometheus/releases/download/v${PROM_VERSION}/sha256sums.txt"
grep "prometheus-${PROM_VERSION}.${ARCH}.tar.gz" sha256sums.txt | sha256sum -c -

# 3. Installation
tar xzf "prometheus-${PROM_VERSION}.${ARCH}.tar.gz"
cd "prometheus-${PROM_VERSION}.${ARCH}"
sudo mkdir -p /etc/prometheus /var/lib/prometheus
sudo cp prometheus promtool /usr/local/bin/
sudo cp -r consoles console_libraries /etc/prometheus/
sudo cp prometheus.yml /etc/prometheus/prometheus.yml

# 4. Utilisateur dédié (ne JAMAIS lancer en root en production)
sudo useradd --no-create-home --shell /usr/sbin/nologin prometheus
sudo chown -R prometheus:prometheus /etc/prometheus /var/lib/prometheus
sudo chmod 755 /usr/local/bin/prometheus /usr/local/bin/promtool

# 5. Vérification de la config avant tout démarrage
promtool check config /etc/prometheus/prometheus.yml
```

⚠️ `promtool check config` est votre meilleur ami : lancez-le **avant chaque
redémarrage** après modification de la config.

## 8. Installation : service systemd

Fichier `/etc/systemd/system/prometheus.service` :

```ini
[Unit]
Description=Prometheus Server
Documentation=https://prometheus.io/docs/
Wants=network-online.target
After=network-online.target

[Service]
User=prometheus
Group=prometheus
Type=simple
Restart=always
RestartSec=5
# Limites anti-OOM : ajuster selon la RAM disponible
MemoryMax=4G
# Démarrage
ExecStart=/usr/local/bin/prometheus \
  --config.file=/etc/prometheus/prometheus.yml \
  --storage.tsdb.path=/var/lib/prometheus \
  --storage.tsdb.retention.time=30d \
  --storage.tsdb.retention.size=50GB \
  --web.console.templates=/etc/prometheus/consoles \
  --web.console.libraries=/etc/prometheus/console_libraries \
  --web.listen-address=0.0.0.0:9090 \
  --web.enable-lifecycle \
  --query.max-concurrency=20 \
  --query.timeout=2m

[Install]
WantedBy=multi-user.target
```

```bash
sudo systemctl daemon-reload
sudo systemctl enable --now prometheus
sudo systemctl status prometheus
# Vérifier que l'UI répond :
curl -s http://localhost:9090/-/healthy && echo " OK"
```

`--web.enable-lifecycle` active `/-/reload` (rechargement de config sans restart)
et `/-/quit` (arrêt propre). **Indispensable en production.**

## 9. Flags de démarrage importants

| Flag | Défaut | Usage |
|---|---|---|
| `--config.file` | `prometheus.yml` | Chemin du fichier de configuration |
| `--storage.tsdb.path` | `data/` | Répertoire des données TSDB |
| `--storage.tsdb.retention.time` | `15d` | Durée de rétention |
| `--storage.tsdb.retention.size` | `512MB` | Taille max (purge au premier atteint) |
| `--web.listen-address` | `0.0.0.0:9090` | Adresse d'écoute HTTP |
| `--web.external-url` | — | URL publique (derrière reverse proxy) |
| `--web.route-prefix` | — | Préfixe de chemin (ex. `/prometheus`) |
| `--web.enable-lifecycle` | off | `/-/reload` et `/-/quit` |
| `--web.enable-admin-api` | off | API admin (snapshots, delete series) |
| `--query.timeout` | `2m` | Timeout max d'une requête |
| `--query.max-concurrency` | `20` | Requêtes concurrentes max |
| `--query.max-samples` | `50M` | Échantillons max par requête (anti-OOM) |
| `--storage.tsdb.no-lockfile` | off | À activer si le FS ne supporte pas flock (NFS…) |

Recharger la configuration **sans redémarrer** :

```bash
curl -X POST http://localhost:9090/-/reload
# ou, si protégé :
kill -HUP $(pidof prometheus)
```

## 10. Structure de prometheus.yml

```yaml
# /etc/prometheus/prometheus.yml
global:
  scrape_interval: 15s          # fréquence de scrape par défaut
  scrape_timeout: 10s           # doit être < scrape_interval !
  evaluation_interval: 15s      # fréquence d'évaluation des règles
  external_labels:              # labels ajoutés à TOUTES les séries (fédération, HA)
    datacenter: 'dc-paris'
    replica: 'prom-01'

# Fichiers de règles (recording + alerting)
rule_files:
  - '/etc/prometheus/rules/*.yml'

# Destinations d'alertes
alerting:
  alertmanagers:
    - static_configs:
        - targets: ['127.0.0.1:9093']

# Les jobs de scrape
scrape_configs:
  - job_name: 'prometheus'
    static_configs:
      - targets: ['127.0.0.1:9090']
```

⚠️ Règles d'or :

- `scrape_timeout` < `scrape_interval`, toujours.
- `external_labels` : **obligatoire** en HA (distinguer les réplicas) et fédération.
- Après chaque modification : `promtool check config` puis `/-/reload`.

## 11. scrape_configs : le job de base

Chaque `scrape_config` définit **un job** = un ensemble de cibles scrapées
de la même façon.

```yaml
scrape_configs:
  - job_name: 'node-linux'
    scrape_interval: 15s
    scrape_timeout: 10s
    metrics_path: /metrics        # chemin HTTP (défaut)
    scheme: http                  # ou https
    static_configs:
      - targets:
          - 'srv-web01:9100'
          - 'srv-web02:9100'
          - 'srv-db01:9100'
        labels:
          datacenter: 'dc-paris'
          equipe: 'infra'
```

Labels automatiques ajoutés par Prometheus à chaque série scrapée :

- `job` = nom du job (`node-linux`),
- `instance` = `host:port` de la cible (`srv-web01:9100`).

On peut les réécrire via `relabel_configs` (section 19).

## 12. static_configs : quand les cibles sont fixes

`static_configs` = liste de cibles en dur. Parfait pour un parc stable
et de taille modérée (< 50 serveurs).

```yaml
  - job_name: 'onduleurs-snmp'
    static_configs:
      - targets: ['192.168.10.20', '192.168.10.21']
        labels:
          site: 'siege'
          equipement: 'ups-eaton-93ps'
```

Limites : chaque ajout/retrait = modification du fichier + reload.
Pour un parc dynamique, préférer `file_sd_configs` (section suivante).

## 13. file_sd_configs : découverte par fichier

Le fichier JSON/YAML listant les cibles est **relu automatiquement**
(`refresh_interval`), sans reload de Prometheus. Idéal quand un inventaire
(Ansible, CMDB, script) génère la liste.

`/etc/prometheus/file_sd/serveurs.json` :

```json
[
  {
    "targets": ["srv-web01:9100", "srv-web02:9100"],
    "labels": {"datacenter": "dc-paris", "role": "web"}
  },
  {
    "targets": ["srv-db01:9100"],
    "labels": {"datacenter": "dc-paris", "role": "db", "critique": "oui"}
  }
]
```

```yaml
scrape_configs:
  - job_name: 'node-dynamique'
    file_sd_configs:
      - files:
          - '/etc/prometheus/file_sd/*.json'
        refresh_interval: 30s
```

> **Astuce d'exploitation** : générez ce fichier depuis votre inventaire Ansible
> (`ansible-inventory --list`) via un template : une seule source de vérité.

## 14. Découverte DNS (SRV, A, AAAA)

Pour les environnements avec DNS à jour (Consul DNS, DNS d'entreprise) :

```yaml
  - job_name: 'dns-sd'
    dns_sd_configs:
      - names: ['_prometheus._tcp.dc-paris.lan']  # enregistrements SRV
        type: 'SRV'
        port: 9100                                # port par défaut si absent du SRV
        refresh_interval: 30s
      - names: ['exporters.dc-paris.lan']         # enregistrements A/AAAA
        type: 'A'
        port: 9100
```

