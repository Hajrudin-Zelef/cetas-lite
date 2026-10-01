---
id: collect-261001-rattrapage/rattrapage/grafana-guide-4
title: "Guide Grafana — Dashboards, visualisation et alerting"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["datacenter", "memory"]
source: docs/RAG/collect-261001-rattrapage/grafana_guide.md
source_anchor: ""
source_lines: [614, 852]
sha256: f5af54f20eb6e7991b2f9ddf6b64f61f758e0fe9cf1bd68b1d835ab06272f9e9
---

# Guide Grafana — Dashboards, visualisation et alerting

```ini
[smtp]
enabled = true
host = smtp.mondomaine.fr:587
user = grafana@mondomaine.fr
password = '<A_COMPLETER>'
from_address = grafana@mondomaine.fr
from_name = Grafana Supervision
# Si votre relais exige STARTTLS (cas général sur le port 587) :
startTLS_policy = MandatoryStartTLS
#skip_verify = false   # ne jamais passer à true en production
```

Test depuis l'interface : **Alerting → Contact points → votre contact point
e-mail → Test** (envoie un e-mail de test immédiat).

Test en ligne de commande (diagnostic hors Grafana) :

```bash
# Vérifier que le relais répond
nc -vz smtp.mondomaine.fr 587
# Dialogue SMTP manuel minimal
openssl s_client -starttls smtp -connect smtp.mondomaine.fr:587
```

> **Erreur fréquente :** `Failed to send notification: dial tcp: i/o timeout`.
> Dans 90 % des cas c'est un pare-feu sortant (port 25/587 bloqué) ou un
> `host` mal saisi, pas un problème Grafana. Testez avec `nc`/`openssl`
> **depuis le serveur Grafana** avant d'accuser la configuration.

---

## 16. Datasource Prometheus — ajout et réglages

**Connections → Data sources → Add data source → Prometheus.**

Réglages essentiels :

| Champ | Valeur conseillée | Commentaire |
|---|---|---|
| Name | `Prometheus` (ou `Prom-Prod`) | Nom affiché dans les panels |
| URL | `http://prometheus.mondomaine.fr:9090` | URL vue **depuis le serveur Grafana** (mode Server) |
| Access | `Server (default)` | Les secrets restent côté serveur |
| Default | coché | Datasource pré-sélectionnée |
| Scrape interval | `15s` | Doit correspondre au `scrape_interval` de Prometheus (sert aux calculs `$__rate_interval`) |
| Query timeout | `60s` | À augmenter si requêtes lourdes |
| HTTP method | `POST` | Recommandé (URL moins longues, cache) |

Authentification si Prometheus est protégé :

- **Basic auth** : cocher, saisir user/password (stockés chiffrés avec `secret_key`).
- **TLS / CA perso** : coller le certificat CA dans « TLS/SSL Auth Details ».
- **Custom HTTP headers** : ex. `Authorization: Bearer <A_COMPLETER>` pour un
  reverse proxy avec token.

Bouton **Save & test** : Grafana exécute `GET /api/v1/status/buildinfo`.
« Successfully queried the Prometheus API » = la connexion fonctionne.

En provisioning YAML (voir section 50) :

```yaml
apiVersion: 1
datasources:
  - name: Prometheus
    type: prometheus
    access: proxy
    url: http://prometheus.mondomaine.fr:9090
    isDefault: true
    jsonData:
      httpMethod: POST
      timeInterval: 15s
      queryTimeout: 60s
```

---

## 17. PromQL de base pour vos dashboards

Dans un panel Time series, datasource Prometheus, onglet **Code** (pas Builder)
dès que vous êtes à l'aise : le code est versionnable et copiable.

```promql
# CPU : pourcentage d'utilisation moyen par instance (5 min)
100 - (avg by (instance) (rate(node_cpu_seconds_total{mode="idle"}[5m])) * 100)

# Mémoire : % utilisé par instance
100 * (1 - (node_memory_MemAvailable_bytes / node_memory_MemTotal_bytes))

# Charge système (load1)
node_load1

# Espace disque : % utilisé par point de montage (hors tmpfs)
100 * (1 - (node_filesystem_avail_bytes{fstype!~"tmpfs|overlay"}
            / node_filesystem_size_bytes{fstype!~"tmpfs|overlay"}))

# Réseau : débit entrant par interface (bits/s)
rate(node_network_receive_bytes_total{device!~"lo|veth.*|docker.*"}[5m]) * 8

# Uptime d'un nœud (secondes)
time() - node_boot_time_seconds

# Nombre d'instances UP vues par Prometheus
count(up == 1)

# Disponibilité d'une cible en % sur 1h
100 * avg_over_time(up[1h])
```

Astuces d'affichage :

- **Legend** : `{{instance}}` ou `{{device}}` pour des légendes lisibles
  (sinon vous aurez la série brute illisible).
- **Min step** : laissez `$__interval` par défaut ; augmentez sur les longues
  plages pour alléger Prometheus.
- Utilisez **`$__rate_interval`** dans `rate(...[$__rate_interval])` : Grafana
  calcule l'intervalle optimal selon la plage affichée et le scrape interval.

---

## 18. PromQL avancé (rate, histogrammes, label_replace)

```promql
# --- Quantile de latence HTTP (histogramme) ---
# p95 de la durée des requêtes par handler
histogram_quantile(0.95,
  sum by (handler, le) (rate(http_request_duration_seconds_bucket[5m]))
)

# --- Taux d'erreur HTTP 5xx en % ---
100 * sum(rate(http_requests_total{code=~"5.."}[5m]))
    / sum(rate(http_requests_total[5m]))

# --- Top 5 des consommateurs CPU (par job) ---
topk(5, sum by (job) (rate(process_cpu_seconds_total[5m])))

# --- Prédiction de saturation disque à 7 jours (régression linéaire) ---
# Nombre de jours avant saturation du /var sur chaque instance
(node_filesystem_avail_bytes{mountpoint="/var",fstype!="tmpfs"}
  / deriv(node_filesystem_size_bytes{mountpoint="/var"}[1h]) ) / 86400
# Variante simple : alerte si < 10 % libre ET décroît
(node_filesystem_avail_bytes / node_filesystem_size_bytes) < 0.1
  and deriv(node_filesystem_avail_bytes[1h]) < 0

# --- Renommer / nettoyer un label ---
label_replace(up, "serveur", "$1", "instance", "([^:]+):.*")

# --- Joindre une info (ex. ajouter le datacenter depuis un autre metric) ---
node_cpu_seconds_total * on(instance) group_left(dc) node_dc_info

# --- Compter les redémarrages de conteneurs (increase sur 1h) ---
increase(kube_pod_container_status_restarts_total[1h]) > 3

# --- Exclure proprement du bruit ---
sum by (instance) (rate(node_cpu_seconds_total{mode!="idle"}[5m]))
```

> **Règle d'or PromQL :** `rate()` s'applique **toujours** à un compteur
> (`_total`). Sur une gauge (température, niveau), utilisez `avg_over_time`,
> `max_over_time`, `deriv`. Appliquer `rate()` à une gauge = chiffres absurdes.

---

## 19. Datasource Loki — ajout et réglages

**Connections → Data sources → Add data source → Loki.**

| Champ | Valeur conseillée |
|---|---|
| Name | `Loki` |
| URL | `http://loki.mondomaine.fr:3100` |
| Access | `Server (default)` |
| Default | selon votre usage |

Options spécifiques Loki :

- **Maximum lines** : `1000` (limite de lignes retournées par requête).
- **Derived fields** : extrayez un `traceID` des logs pour sauter vers Tempo
  (si vous avez du tracing) — optionnel.
- **Alerting** : si vous voulez des règles d'alerte Loki gérées par Grafana,
  activez « Manage alerts via Alerting UI » (nécessite le ruler Loki configuré).

**Save & test** : Grafana interroge `/loki/api/v1/labels`.

Provisioning YAML :

```yaml
apiVersion: 1
datasources:
  - name: Loki
    type: loki
    access: proxy
    url: http://loki.mondomaine.fr:3100
    jsonData:
      maxLines: 1000
```

> **Corrélation Loki ↔ Prometheus :** dans la datasource Loki, renseignez le
> champ **Derived fields** ou utilisez les **exemplars** côté Prometheus pour
> passer d'une courbe à « voir les logs de cette instance à ce moment » en un
> clic (data link, section 41).

---

## 20. LogQL de base

LogQL ressemble à PromQL mais s'applique aux **flux de logs** (sélectionnés par
labels, ex. `{job="syslog"}`).

```logql
# Tous les logs d'un job sur la plage affichée
{job="syslog"}

# Filtrer : lignes contenant "error" (insensible à la casse avec (?i))
{job="syslog"} |= "error"
{job="syslog"} |~ "(?i)error|fail|timeout"

# Exclure : tout sauf le bruit
{job="syslog"} != "DEBUG" |!= "healthcheck"

# Plusieurs labels
{job="nginx", instance="srv-web-03"} |= "500"

# Compter les lignes d'erreur par minute (métrique issue des logs !)
count_over_time({job="syslog"} |= "error" [1m])

# Taux d'erreurs nginx par seconde
sum by (instance) (rate({job="nginx"} |= " 500 " [5m]))
```

Dans Grafana, le panel **Logs** affiche les lignes avec coloration par niveau
(`level=error` détecté automatiquement si présent). Les labels extraits sont
cliquables pour filtrer.

**Bon réflexe :** commencez large (`{job="syslog"}`), ajoutez les filtres
`|=` un par un, puis transformez en métrique (`count_over_time`, `rate`) pour
les dashboards et les alertes.

---

