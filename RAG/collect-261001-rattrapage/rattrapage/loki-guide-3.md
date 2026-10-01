---
id: collect-261001-rattrapage/rattrapage/loki-guide-3
title: "Grafana Loki — Le guide complet"
domain: rattrapage
role: reference
task: reference
actors: []
dates: ["2024-01-01"]
keywords: ["agent"]
source: docs/RAG/collect-261001-rattrapage/loki_guide.md
source_anchor: ""
source_lines: [406, 705]
sha256: 2a39bb0b8c330b0becf153463c477510723ff57094c729174b9f06034a7f2028
---

# 3. Vérification du checksum (fortement recommandé)
curl -sSL "https://github.com/grafana/loki/releases/download/v${LOKI_VERSION}/SHA256SUMS" -o SHA256SUMS
sha256sum -c SHA256SUMS --ignore-missing

# 4. Installation
sudo apt-get install -y unzip
unzip -o loki-linux-${ARCH}.zip
unzip -o promtail-linux-${ARCH}.zip
sudo install -m 0755 loki-linux-${ARCH} /usr/local/bin/loki
sudo install -m 0755 promtail-linux-${ARCH} /usr/local/bin/promtail

# 5. Utilisateur dédié + répertoires
sudo useradd --system --no-create-home --shell /usr/sbin/nologin loki
sudo mkdir -p /etc/loki /var/lib/loki /var/log/loki
sudo chown -R loki:loki /var/lib/loki /var/log/loki

# 6. Vérification
loki --version
promtail --version
```

Alternative : paquets `.deb` non officiels ou dépôt APT tiers — le binaire
reste la méthode la plus maîtrisée et la plus documentée.

---

## 11. Installation : service systemd

Fichier `/etc/systemd/system/loki.service` :

```ini
[Unit]
Description=Grafana Loki - log aggregation system
Documentation=https://grafana.com/docs/loki/
After=network-online.target
Wants=network-online.target

[Service]
Type=simple
User=loki
Group=loki
ExecStart=/usr/local/bin/loki -config.file=/etc/loki/loki.yaml
Restart=on-failure
RestartSec=5
# Durcissement
NoNewPrivileges=true
ProtectSystem=strict
ProtectHome=true
ReadWritePaths=/var/lib/loki /var/log/loki /tmp
PrivateTmp=true
LimitNOFILE=65536

[Install]
WantedBy=multi-user.target
```

Activation :

```bash
sudo systemctl daemon-reload
sudo systemctl enable --now loki
sudo systemctl status loki --no-pager
# Logs du service :
sudo journalctl -u loki -f
```

> 💡 `LimitNOFILE=65536` est important : Loki ouvre beaucoup de fichiers
> (chunks, index, WAL). La limite par défaut (1024) provoque des erreurs
> `too many open files` en charge.

---

## 12. Installation : Docker

Pour un lab ou un déploiement conteneurisé, l'image officielle suffit :

```bash
mkdir -p ~/loki/{config,data}
cat > ~/loki/config/loki.yaml <<'EOF'
auth_enabled: false
server:
  http_listen_port: 3100
common:
  path_prefix: /loki
  storage:
    filesystem:
      chunks_directory: /loki/chunks
      rules_directory: /loki/rules
  replication_factor: 1
  ring:
    instance_addr: 127.0.0.1
    kvstore:
      store: inmemory
schema_config:
  configs:
    - from: 2024-01-01
      store: tsdb
      object_store: filesystem
      schema: v13
      index:
        prefix: index_
        period: 24h
EOF

docker run -d --name loki --restart unless-stopped \
  -p 3100:3100 \
  -v ~/loki/config/loki.yaml:/etc/loki/loki.yaml:ro \
  -v ~/loki/data:/loki \
  grafana/loki:3.3.2 -config.file=/etc/loki/loki.yaml
```

Avec Docker Compose (Loki + Promtail + Grafana), voir la section 50
pour un exemple complet prêt à l'emploi.

---

## 13. Premier fichier `loki.yaml` (mode monolithique)

Voici une configuration **monolithique complète et fonctionnelle** pour
démarrer sur un serveur Debian/Ubuntu avec stockage filesystem.
C'est la base que nous enrichirons dans tout le guide.

```yaml
# /etc/loki/loki.yaml
auth_enabled: false          # true en production multi-tenant (section 55)

server:
  http_listen_port: 3100
  grpc_listen_port: 9096
  log_level: info

common:
  path_prefix: /var/lib/loki
  storage:
    filesystem:
      chunks_directory: /var/lib/loki/chunks
      rules_directory: /var/lib/loki/rules
  replication_factor: 1
  ring:
    instance_addr: 127.0.0.1
    kvstore:
      store: inmemory        # en cluster : consul ou etcd

ingester:
  wal:
    enabled: true
    dir: /var/lib/loki/wal
  lifecycler:
    address: 127.0.0.1
    ring:
      kvstore:
        store: inmemory
      replication_factor: 1

schema_config:
  configs:
    - from: 2024-01-01
      store: tsdb
      object_store: filesystem
      schema: v13
      index:
        prefix: index_
        period: 24h

storage_config:
  tsdb_shipper:
    active_index_directory: /var/lib/loki/tsdb-index
    cache_location: /var/lib/loki/tsdb-cache

compactor:
  working_directory: /var/lib/loki/compactor
  compaction_interval: 10m
  retention_enabled: true
  retention_delete_delay: 2h
  retention_delete_worker_count: 150

limits_config:
  retention_period: 744h     # 31 jours
  ingestion_rate_mb: 16
  ingestion_burst_size_mb: 32
  max_query_parallelism: 16
  max_query_length: 721h
  split_queries_by_interval: 30m

query_range:
  align_queries_with_step: true
  results_cache:
    cache:
      embedded_cache:
        enabled: true
        max_size_mb: 100

ruler:
  alertmanager_url: http://localhost:9093   # si Alertmanager installé
  storage:
    type: local
    local:
      directory: /var/lib/loki/rules
```

> ⚠️ `auth_enabled: false` = **aucune authentification**. OK en lab sur
> réseau de confiance, **interdit** exposé sur Internet. Voir section 55.

---

## 14. Vérifier que Loki démarre

```bash
# Le port d'écoute répond ?
curl -s http://localhost:3100/ready
# → "ready" (peut prendre quelques secondes au premier démarrage)

# Les métriques Prometheus internes :
curl -s http://localhost:3100/metrics | head -20

# Page d'état (ingesters, schémas, build info) :
curl -s http://localhost:3100/loki/api/v1/status/buildinfo
```

Checklist de premier démarrage :

- [ ] `curl localhost:3100/ready` renvoie `ready`
- [ ] aucune erreur `level=error` dans `journalctl -u loki`
- [ ] les répertoires `/var/lib/loki/{chunks,tsdb-index,wal}` se créent
- [ ] le port 3100 n'est exposé que sur le réseau prévu (pare-feu)

Test d'ingestion manuel (sans Promtail) :

```bash
curl -s -X POST http://localhost:3100/loki/api/v1/push \
  -H 'Content-Type: application/json' \
  -d '{
    "streams": [{
      "stream": {"job": "test", "host": "srv01"},
      "values": [["'$(date +%s%N)'", "premier log de test"]]
    }]
  }'
# → HTTP 204 = accepté
```

Puis vérifiez dans Grafana Explore ou via l'API :
`curl -s 'http://localhost:3100/loki/api/v1/query_range?query={job="test"}'`.

---

## 15. Modes de déploiement : monolithique vs microservices

| Critère | Monolithique (`-target=all`) | Microservices |
|---|---|---|
| Processus | 1 seul | 6+ (distributor, ingester, querier…) |
| Complexité ops | faible | élevée (ring, gossip, scaling) |
| Scalabilité | verticale (plus gros serveur) | horizontale (plus de réplicas) |
| Stockage | filesystem possible | objet quasi obligatoire |
| Quand | < ~2 To/jour, équipe réduite | gros volumes, haute dispo |

**Recommandation pour Zelef** : commencez **monolithique** avec stockage
objet (MinIO). C'est simple à opérer, à sauvegarder et à superviser.
Passez en microservices (ou en mode « scalable monolithique » : 3 cibles
`read`, `write`, `backend`) uniquement si un composant devient le goulot.

Le mode intermédiaire « scalable » (Loki 2.9+) :

```bash
# 3 déploiements du même binaire, cibles différentes :
loki -target=read      # query-frontend + querier
loki -target=write     # distributor + ingester
loki -target=backend   # compactor + ruler + index-gateway
```

---

## 16. Promtail : installation du binaire

Promtail est l'**agent** qui lit les fichiers de log sur chaque machine et
les pousse vers Loki. Installez-le **sur chaque serveur supervisé**.

```bash
# Binaire déjà récupéré à la section 10, sinon :
PROMTAIL_VERSION="3.3.2"
cd /tmp
curl -sSLO "https://github.com/grafana/loki/releases/download/v${PROMTAIL_VERSION}/promtail-linux-amd64.zip"
unzip -o promtail-linux-amd64.zip
sudo install -m 0755 promtail-linux-amd64 /usr/local/bin/promtail

sudo useradd --system --no-create-home --shell /usr/sbin/nologin promtail
# Promtail doit LIRE les logs : ajoutez-le aux groupes concernés
sudo usermod -aG adm,systemd-journal promtail

sudo mkdir -p /etc/promtail /var/lib/promtail
sudo chown -R promtail:promtail /var/lib/promtail
```

Service systemd `/etc/systemd/system/promtail.service` :

```ini
[Unit]
Description=Promtail - log shipper for Loki
After=network-online.target
Wants=network-online.target

