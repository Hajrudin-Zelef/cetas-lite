---
id: collect-261001-rattrapage/rattrapage/loki-guide-9
title: "Grafana Loki — Le guide complet"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/loki_guide.md
source_anchor: ""
source_lines: [2111, 2384]
sha256: 7202135fa9df1f8400c41587152988dca2c15bc63d5f4adb77d86b84a315c74e
---

# Grafana Loki — Le guide complet

```yaml
      - alert: AttaqueSSHBruteForce
        expr: sum(rate({job="sshd"} |= "Failed password" [5m])) by (host) > 0.5
        for: 10m
        labels:
          severity: warning
          service: ssh
        annotations:
          summary: "Possible brute-force SSH sur {{ $labels.host }}"
          description: "Vérifier fail2ban : fail2ban-client status sshd"
```

### Corrélation avec fail2ban

```bash
# IP bannies actuellement :
sudo fail2ban-client status sshd
# Croiser avec Loki : l'IP top-attaquante est-elle bannie ?
```

> 💡 Piste d'amélioration : un script qui lit le top des IP dans Loki
> (API `/loki/api/v1/query`) et les injecte dans un ipset de blocage
> préventif. À évaluer avec prudence (risque de faux positifs).

---

## 53. Cas pratique 4 : logs applicatifs au format JSON

**Objectif** : une application Java qui logue en JSON, avec niveau de log
en label et `trace_id` en structured metadata.

### Log d'exemple

```json
{"ts":"2026-09-26T10:00:01.123Z","level":"ERROR","logger":"com.app.Paiement","msg":"Paiement refusé","trace_id":"a3f9c1d2e4b5","amount":129.99,"user_id":"u-4821"}
```

### Promtail

```yaml
scrape_configs:
  - job_name: app-paiement
    static_configs:
      - targets: [localhost]
        labels:
          job: app-paiement
          host: app01
          env: prod
          __path__: /opt/paiement/logs/app.json.log
    pipeline_stages:
      - multiline:
          firstline: '^\{"ts"'
          max_wait_time: 3s
      - json:
          expressions:
            level: level
            message: msg
            trace_id: trace_id
            ts: ts
      - labels:
          level:
      - structured_metadata:
          trace_id:
      - timestamp:
          source: ts
          format: RFC3339Nano
      - output:
          source: message
```

### Requêtes

```logql
# Toutes les erreurs de l'app
{job="app-paiement", level="ERROR"}

# Suivre une transaction de bout en bout via son trace_id
{job="app-paiement"} | trace_id="a3f9c1d2e4b5"

# Taux d'erreur par minute
sum(count_over_time({job="app-paiement", level="ERROR"}[1m]))

# Erreurs contenant "Paiement refusé" sur 7 jours
{job="app-paiement", level="ERROR"} |= "Paiement refusé"
```

---

## 54. Cas pratique 5 : audit des commandes sudo

**Objectif** : tracer qui a fait `sudo` quoi, où et quand — besoin
classique d'un chef de service systèmes.

### Source : journald (où sudo logue nativement)

```yaml
scrape_configs:
  - job_name: journal
    journal:
      path: /var/log/journal
      max_age: 12h
      labels:
        job: systemd-journal
        host: srv01
    relabel_configs:
      - source_labels: ['__journal__systemd_unit']
        target_label: unit
    pipeline_stages:
      # Ne garder que sudo (filtre à l'ingestion = moins de volume)
      - match:
          selector: '{unit="sudo.service"}'
          stages:
            - regex:
                expression: '(?P<sudo_user>\S+) : TTY=\S+ ; PWD=(?P<pwd>\S+) ; USER=(?P<target_user>\S+) ; COMMAND=(?P<command>.+)'
            - labels:
                sudo_user:
                target_user:
            - structured_metadata:
                command:
```

### Requêtes d'audit

```logql
# Toutes les commandes sudo d'un utilisateur
{job="systemd-journal", unit="sudo.service", sudo_user="martin"}

# Qui a touché au fichier shadow ?
{job="systemd-journal", unit="sudo.service"} |= "/etc/shadow"

# Commandes sudo par utilisateur sur 30 jours (rapport mensuel)
sum by (sudo_user) (count_over_time({job="systemd-journal", unit="sudo.service"}[30d]))

# sudo vers root par un compte de service (suspect)
{job="systemd-journal", unit="sudo.service", target_user="root", sudo_user=~"svc-.*"}
```

> ⚠️ `command` en structured metadata : une commande complète est quasi
> unique → jamais en label. Et pensez au RGPD : les logs sudo contiennent
> des données personnelles, cadrez la rétention (section 45).

---

## 55. Sécurité : authentification et multi-tenant

Par défaut (`auth_enabled: false`), **quiconque atteint le port 3100 peut
lire tous les logs et en pousser**. En production : activez l'authentification.

```yaml
auth_enabled: true
```

Avec `auth_enabled: true`, chaque requête doit porter le header
`X-Scope-OrgID: <tenant>` :

```bash
# Push vers le tenant 'prod'
curl -H 'X-Scope-OrgID: prod' -X POST \
  http://localhost:3100/loki/api/v1/push ...

# Requête sur le tenant 'prod'
curl -H 'X-Scope-OrgID: prod' \
  'http://localhost:3100/loki/api/v1/query_range?query={job="nginx"}'
```

Promtail côté client :

```yaml
clients:
  - url: http://loki.local:3100/loki/api/v1/push
    tenant_id: prod
```

Stratégie de tenants recommandée :

| Tenant | Contenu | Rétention |
|---|---|---|
| `prod` | serveurs de production | 90 jours |
| `infra` | équipements réseau, hyperviseurs | 90 jours |
| `lab` | tests, préproduction | 7 jours |

> 💡 Loki ne gère pas lui-même les mots de passe : placez un reverse
> proxy (nginx, section 57) avec authentification basique ou OIDC devant,
> qui injecte le `X-Scope-OrgID` selon l'utilisateur.

---

## 56. Sécurité : TLS de bout en bout

Chiffrez au minimum l'ingestion et l'interface Grafana ↔ Loki.

### TLS côté serveur Loki

```yaml
server:
  http_listen_port: 3100
  http_tls_config:
    cert_file: /etc/loki/tls/loki.crt
    key_file: /etc/loki/tls/loki.key
  grpc_listen_port: 9096
  grpc_tls_config:
    cert_file: /etc/loki/tls/loki.crt
    key_file: /etc/loki/tls/loki.key
```

### Promtail vers Loki en HTTPS

```yaml
clients:
  - url: https://loki.local:3100/loki/api/v1/push
    tls_config:
      ca_file: /etc/promtail/tls/ca.crt
      cert_file: /etc/promtail/tls/client.crt   # si mTLS
      key_file: /etc/promtail/tls/client.key
      # insecure_skip_verify: false  ← ne JAMAIS mettre à true en prod
```

Générez les certificats avec votre CA interne (voir le guide Debian :
section CA interne) ou Let's Encrypt pour les noms publics.

---

## 57. Sécurité : reverse proxy nginx devant Loki

Le reverse proxy centralise TLS + authentification basique + injection
du tenant.

```nginx
# /etc/nginx/sites-available/loki
upstream loki {
    server 127.0.0.1:3100;
}

server {
    listen 443 ssl;
    server_name loki.entreprise.lan;

    ssl_certificate     /etc/ssl/certs/loki.entreprise.lan.crt;
    ssl_certificate_key /etc/ssl/private/loki.entreprise.lan.key;

    # Authentification basique (fichier htpasswd)
    auth_basic "Loki - accès restreint";
    auth_basic_user_file /etc/nginx/htpasswd_loki;

    location / {
        proxy_pass http://loki;
        proxy_set_header Host $host;
        # Tenant selon l'utilisateur authentifié :
        # (ici : tout le monde → tenant 'prod' ; affinez avec des maps)
        proxy_set_header X-Scope-OrgID "prod";
        proxy_read_timeout 300s;
        proxy_send_timeout 300s;
        client_max_body_size 10M;   # les pushs peuvent être volumineux
    }
}
```

Pour un mapping utilisateur → tenant fin, utilisez `map $remote_user` :

```nginx
map $remote_user $loki_tenant {
    default      "lab";
    "admin"      "prod";
    "reseau"     "infra";
}
# puis : proxy_set_header X-Scope-OrgID $loki_tenant;
```

---

## 58. Sécurité : durcissement

Checklist de durcissement d'un Loki exposé au réseau interne :

