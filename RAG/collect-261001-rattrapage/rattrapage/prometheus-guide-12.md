---
id: collect-261001-rattrapage/rattrapage/prometheus-guide-12
title: "Guide Prometheus — Supervision métrique complète"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["arr", "valuation"]
source: docs/RAG/collect-261001-rattrapage/prometheus_guide.md
source_anchor: ""
source_lines: [2516, 2727]
sha256: e87e1de258e7f607383816ccece1b81149e12cd22128a1a6dde29e08fef3959b
---

# Guide Prometheus — Supervision métrique complète

Nginx devant Prometheus (TLS + basic auth + restriction IP) :

```nginx
server {
    listen 443 ssl;
    server_name prometheus.lan;

    ssl_certificate     /etc/ssl/prometheus.lan.crt;
    ssl_certificate_key /etc/ssl/prometheus.lan.key;
    ssl_protocols TLSv1.2 TLSv1.3;

    # API reload protégée : seul le réseau admin
    location = /-/reload {
        allow 10.0.0.0/8;
        deny all;
        auth_basic "admin";
        auth_basic_user_file /etc/nginx/.htpasswd;
        proxy_pass http://127.0.0.1:9090;
    }

    location / {
        allow 10.0.0.0/8;       # réseau interne uniquement
        deny all;
        auth_basic "Prometheus";
        auth_basic_user_file /etc/nginx/.htpasswd;
        proxy_pass http://127.0.0.1:9090;
        proxy_set_header Host $host;
    }
}
```

Côté scrape en HTTPS vers les exporters :

```yaml
  - job_name: 'node-tls'
    scheme: https
    tls_config:
      ca_file: '/etc/prometheus/certs/ca.crt'
      # cert_file: '/etc/prometheus/certs/client.crt'   # si mTLS
      # key_file: '/etc/prometheus/certs/client.key'
      insecure_skip_verify: false   # jamais true en production !
    basic_auth:
      username: 'prometheus'
      password_file: '/etc/prometheus/node_pass'  # fichier, pas en clair
    static_configs:
      - targets: ['srv-dmz01:9100']
```

## 84. Sécurité : durcissement en production

Checklist :

- [ ] Utilisateur dédié non-root (`prometheus`, `node_exporter`…).
- [ ] `web.listen-address=127.0.0.1:9090` si accès via reverse proxy/SSH.
- [ ] Mots de passe **jamais en clair** : `password_file`, `*_token_file`.
- [ ] `--web.enable-admin-api` : à n'activer que si besoin (snapshots/delete),
  et protéger via ACL IP.
- [ ] `--web.enable-lifecycle` : protéger `/-/quit` et `/-/reload` (ACL IP).
- [ ] Fichiers de config en `640` propriétaire `prometheus:prometheus`
  (ils contiennent des secrets).
- [ ] SNMP : **v3 authPriv**, jamais v2c en clair sur un réseau non maîtrisé.
- [ ] Alertmanager : `bot_token_file`, `auth_password_file` (pas en clair).
- [ ] Sauvegardes chiffrées si elles quittent le site.
- [ ] Mettre à jour régulièrement (faille Go ou UI possible) — voir section 86.
- [ ] Journaliser les accès admin (nginx access_log sur `/-/reload`, `/api/v1/admin`).

---

---

# PARTIE VII — Supervision de Prometheus, mise à jour, dépannage

---

## 85. Supervision de Prometheus lui-même (méta-monitoring)

"Qui surveille le surveillant ?" — **un second Prometheus** (celui de la paire
HA, section 76) ou, à minima, des checks externes. Métriques vitales que
Prometheus expose sur lui-même :

```promql
# --- Santé du scrape ---
up{job="node-linux"} == 0                          # cibles down
increase(prometheus_target_scrapes_exceeded_sample_limit_total[1h]) > 0   # sample_limit atteinte
rate(prometheus_target_scrape_pool_exceeded_target_limit_total[5m]) > 0

# --- Performance des requêtes ---
histogram_quantile(0.95, rate(prometheus_engine_query_duration_seconds_bucket[5m])) > 1
# p95 des requêtes > 1 s = requêtes trop lourdes

# --- Mémoire / TSDB ---
prometheus_tsdb_head_series                          # séries actives (tendance !)
prometheus_tsdb_head_chunks
rate(prometheus_tsdb_head_gc_duration_seconds_sum[5m])  # GC trop long = mémoire sous pression
prometheus_tsdb_compactions_failed_total

# --- Stockage ---
prometheus_tsdb_retention_limit_bytes - prometheus_tsdb_size_retention_total_bytes < 5e9
# moins de 5 Go de marge avant purge forcée

# --- Règles ---
prometheus_rule_evaluation_failures_total > 0        # règles en erreur
prometheus_rule_group_interval_seconds - prometheus_rule_group_last_duration_seconds < 0
# évaluation plus longue que l'intervalle = règles trop lourdes

# --- Alertmanager ---
prometheus_notifications_dropped_total > 0           # alertes perdues !
rate(prometheus_notifications_failed_total[5m]) > 0
```

Alertes méta-minimales à mettre en place **en priorité** :

```yaml
- alert: PrometheusDown
  expr: up{job="prometheus"} == 0
  for: 5m
  labels: {severity: critical}
- alert: PrometheusSeriesExplosion
  expr: increase(prometheus_tsdb_head_series[1h]) > 100000
  for: 30m
  labels: {severity: warning}
  annotations: {summary: "Explosion du nombre de séries : vérifier la cardinalité"}
- alert: PrometheusNotificationsDropped
  expr: increase(prometheus_notifications_dropped_total[10m]) > 0
  for: 5m
  labels: {severity: critical}
```

> **Règle d'or** : les alertes `PrometheusDown` et `AlertmanagerDown` doivent
> être supervisées par un système **externe** (deuxième site, check HTTP externe)
> : si tout votre monitoring est down, personne ne vous préviendra.

## 86. Mise à jour de Prometheus

Procédure sans perte de données (TSDB compatible ascendante entre mineures) :

```bash
# 1. Sauvegarde : snapshot + config
curl -X POST http://127.0.0.1:9090/api/v1/admin/tsdb/snapshot
cp -a /etc/prometheus /root/prometheus-config-backup-$(date +%F)

# 2. Arrêt propre (flush du WAL)
sudo systemctl stop prometheus

# 3. Remplacement du binaire
cd /tmp
wget "https://github.com/prometheus/prometheus/releases/download/vX.Y.Z/prometheus-X.Y.Z.linux-amd64.tar.gz"
tar xzf prometheus-X.Y.Z.linux-amd64.tar.gz
sudo cp prometheus-X.Y.Z.linux-amd64/prometheus /usr/local/bin/
sudo cp prometheus-X.Y.Z.linux-amd64/promtool /usr/local/bin/

# 4. Vérification config (les flags dépréciés peuvent casser le démarrage !)
promtool check config /etc/prometheus/prometheus.yml

# 5. Redémarrage + vérification
sudo systemctl start prometheus
sleep 5
curl -s http://127.0.0.1:9090/-/healthy && echo OK
curl -s http://127.0.0.1:9090/api/v1/status/buildinfo | grep version
```

Points de vigilance :

- Lire les **release notes** : certains flags sont renommés/supprimés entre
  majeures (ex. 2.x → 3.x : vérifier `--storage.tsdb.*` et la syntaxe des
  `matchers` Alertmanager).
- Ne jamais sauter plus d'une majeure sans étape intermédiaire.
- Tester d'abord sur un Prometheus **de preprod** ou sur le second membre
  de la paire HA (rolling upgrade).
- Après upgrade : vérifier `Targets` (tout UP), `Rules` (0 erreur), et
  `prometheus_tsdb_head_series` (pas d'explosion).

## 87. Dépannage : une cible est DOWN

Symptôme : `up{instance="srv-web01:9100"} == 0`, pastille rouge dans Targets
avec message d'erreur.

| Message d'erreur | Cause probable | Action |
|---|---|---|
| `connection refused` | Exporter arrêté | `systemctl status node_exporter` sur la cible |
| `connection timed out` | Firewall / réseau | `telnet srv-web01 9100`, vérifier iptables/nftables |
| `context deadline exceeded` | Scrape trop lent | Augmenter `scrape_timeout`, alléger l'exporter |
| `server returned HTTP 404` | Mauvais `metrics_path` | Vérifier le chemin (`/metrics` ?) |
| `401 Unauthorized` | Basic auth | Vérifier `basic_auth` / `password_file` |
| `certificate verify failed` | TLS | `ca_file`, nom DNS vs CN du certificat |
| `exceeded sample limit` | Trop de séries | Voir cardinalité (section 81) |

Méthode express depuis le serveur Prometheus :

```bash
# Reproduire le scrape à la main (mêmes paramètres que Prometheus)
curl -s -m 10 http://srv-web01:9100/metrics | head -20
curl -s -m 10 http://srv-web01:9100/metrics | grep -c '^node_'
```

## 88. Dépannage : requêtes lentes ou timeout

Symptômes : Grafana "timeout", `query.timeout` atteint, UI lente.

Diagnostic :

```promql
# Requêtes les plus lentes (p95 par chemin)
histogram_quantile(0.95, sum by (path, le) (rate(prometheus_engine_query_duration_seconds_bucket[5m])))

# File d'attente des requêtes (si > 0 durablement : sous-dimensionné)
prometheus_engine_queries
prometheus_engine_query_log_enabled
```

Remèdes (dans l'ordre) :

