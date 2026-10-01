---
id: collect-261001-rattrapage/rattrapage/loki-guide-7
title: "Grafana Loki — Le guide complet"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/loki_guide.md
source_anchor: ""
source_lines: [1548, 1820]
sha256: 7091fb7e16c4189907e4121e33233ac68904c650e2cf60f005bbe9eb607440a0
---

# Grafana Loki — Le guide complet

```yaml
limits_config:
  max_query_length: 721h          # 30 j max
  max_query_parallelism: 32       # parallélisme par requête
  max_entries_limit_per_query: 5000
  max_series_per_query: 1000      # garde-fou cardinalité
  max_cache_freshness_per_query: 10m
```

Et côté query-frontend : `split_queries_by_interval: 30m` +
`max_outstanding_per_tenant` (section 8).

---

## 40. Alerting : les règles du ruler Loki

Le **ruler** évalue des requêtes LogQL périodiquement et envoie les
alertes à Alertmanager. Les règles vivent dans des fichiers YAML
(`rules_directory`, voir section 13).

`/var/lib/loki/rules/fake/alerts.yaml` (le répertoire porte le nom du
tenant ; `fake` quand `auth_enabled: false`) :

```yaml
groups:
  - name: nginx-errors
    interval: 1m
    rules:
      - alert: TropDErreurs500
        expr: |
          sum(rate({job="nginx"} | regexp `(?P<status>\d{3})` | status=~"5.." [5m])) > 1
        for: 5m
        labels:
          severity: critical
          service: nginx
        annotations:
          summary: "Trop d'erreurs 500 sur nginx ({{ $labels.host }})"
          description: "Plus de 1 erreur 500/s depuis 5 minutes."

      - alert: AucunLogServeur
        expr: |
          absent_over_time({job="syslog", host="srv01"}[10m])
        for: 10m
        labels:
          severity: warning
        annotations:
          summary: "srv01 n'envoie plus de logs"
```

Configuration du ruler dans `loki.yaml` :

```yaml
ruler:
  alertmanager_url: http://localhost:9093
  enable_api: true
  storage:
    type: local
    local:
      directory: /var/lib/loki/rules
  ring:
    kvstore:
      store: inmemory
  # Évaluer les règles sur les 15 dernières minutes max :
  query_offset: 2m
```

Rechargez sans redémarrer : `curl -X POST http://localhost:3100/ruler/reload`
(ou `SIGHUP`).

---

## 41. Alertmanager : alerter sur trop d'erreurs 500

Exemple complet : Loki → Alertmanager → notification.

`alertmanager.yaml` minimal :

```yaml
global:
  smtp_smarthost: 'smtp.entreprise.lan:587'
  smtp_from: 'alertes@entreprise.lan'
  smtp_auth_username: 'alertes@entreprise.lan'
  smtp_auth_password: 'MOT_DE_PASSE'   # idéalement via fichier ou variable

route:
  receiver: 'equipe-systemes'
  group_by: ['alertname', 'service']
  group_wait: 30s
  group_interval: 5m
  repeat_interval: 4h
  routes:
    - matchers:
        - severity = "critical"
      receiver: 'astreinte'
      repeat_interval: 30m

receivers:
  - name: 'equipe-systemes'
    email_configs:
      - to: 'systemes@entreprise.lan'
  - name: 'astreinte'
    email_configs:
      - to: 'astreinte@entreprise.lan'
```

Règle Loki correspondante (fichier de la section 40) :

```logql
sum(rate({job="nginx"} | regexp `(?P<status>\d{3})` | status=~"5.." [5m])) > 1
```

> 💡 **Seuils** : commencez avec des seuils **larges** (`> 1/s` pendant
> 5 min) et resserrez après 2 semaines d'observation. Une alerte qui sonne
> tout le temps finit ignorée — c'est pire que pas d'alerte.

Tester une règle sans attendre : évaluez l'expression dans Grafana
Explore sur les 7 derniers jours et vérifiez visuellement les
dépassements.

---

## 42. Alerting : bonnes pratiques

| Principe | Exemple |
|---|---|
| Alerter sur des **symptômes**, pas des causes | taux d'erreur 5xx, pas « CPU à 80 % » |
| `for:` toujours renseigné | `for: 5m` minimum (évite le flapping) |
| Labels `severity` + `service` | routage et tri dans Alertmanager |
| Annotations actionnables | runbook, dashboard Grafana en lien |
| Pas d'alerte sans destinataire | chaque alerte critique → astreinte |
| Réévaluer les seuils | revue mensuelle des alertes bruyantes |

Modèle d'annotations recommandé :

```yaml
        annotations:
          summary: "Résumé en une ligne"
          description: "Ce qui se passe, depuis quand, ampleur"
          runbook: "https://wiki.entreprise.lan/runbooks/nginx-500"
          dashboard: "https://grafana.entreprise.lan/d/nginx-logs"
```

---

## 43. Rétention : le table manager (méthode historique)

Avec l'index **BoltDB**, la rétention est gérée par le **table manager** :
les tables périodiques (`index_20260101`) sont supprimées quand elles
dépassent `retention_period`.

```yaml
# Valable UNIQUEMENT avec l'index boltdb-shipper (legacy)
storage_config:
  boltdb_shipper:
    active_index_directory: /var/lib/loki/index
    cache_location: /var/lib/loki/index-cache

table_manager:
  retention_deletes_enabled: true
  retention_period: 744h   # 31 jours
```

> ⚠️ Avec l'index **TSDB** (recommandé depuis Loki 2.8), le table manager
> ne gère **plus** la rétention : c'est le **compactor** qui s'en charge
> (section 44). Ne mélangez pas les deux.

---

## 44. Rétention : le compactor (méthode moderne)

Avec l'index TSDB, le **compactor** fait trois choses : il compacte les
fichiers d'index, **supprime les chunks et index expirés** (rétention),
et déduplique.

```yaml
compactor:
  working_directory: /var/lib/loki/compactor
  compaction_interval: 10m
  retention_enabled: true
  retention_delete_delay: 2h      # délai de grâce avant suppression réelle
  retention_delete_worker_count: 150

limits_config:
  retention_period: 744h          # 31 jours, global
```

Vérifier que la rétention tourne :

```bash
# Le compactor a-t-il tourné récemment ?
curl -s http://localhost:3100/metrics | grep -E "loki_compactor.*(last|applied)"
# Marqueurs de suppression :
curl -s http://localhost:3100/metrics | grep loki_compactor_deleted
```

> 💡 La suppression est **asynchrone** : entre l'expiration théorique et
> la libération disque, comptez `retention_delete_delay` + le temps du
> cycle de compaction. Ne paniquez pas si le disque ne se libère pas
> instantanément.

---

## 45. Rétention par tenant

Avec `auth_enabled: true`, chaque tenant peut avoir sa propre rétention
(utile : 90 jours pour la prod, 7 jours pour le lab).

```yaml
auth_enabled: true

limits_config:
  retention_period: 744h   # défaut : 31 jours

  # Surcharges par tenant (fichier séparé recommandé) :
  per_tenant_override_config: /etc/loki/overrides.yaml
```

`/etc/loki/overrides.yaml` :

```yaml
overrides:
  "tenant-prod":
    retention_period: 2160h    # 90 jours
  "tenant-lab":
    retention_period: 168h     # 7 jours
```

> ⚠️ La rétention par tenant exige le compactor avec `retention_enabled:
> true` **et** un stockage objet (ou filesystem en mono-tenant de test).
> Les chunks partagés entre tenants ne sont supprimés que quand tous les
> tenants les ont expirés.

---

## 46. Intégration Grafana : déclarer la datasource

Dans Grafana : **Connections → Data sources → Add → Loki**.

| Champ | Valeur |
|---|---|
| URL | `http://loki.local:3100` |
| Auth | selon votre setup (section 55/57) |
| `X-Scope-OrgID` header | nom du tenant si `auth_enabled: true` |
| Maximum lines | `1000` (défaut raisonnable) |

En provisioning (`/etc/grafana/provisioning/datasources/loki.yaml`) :

```yaml
apiVersion: 1
datasources:
  - name: Loki
    type: loki
    access: proxy
    url: http://loki.local:3100
    jsonData:
      maxLines: 1000
      derivedFields:
        # Cliquer sur un trace_id ouvre Tempo/Jaeger :
        - datasourceUid: tempo
          matcherRegex: "trace_id=(\\w+)"
          name: Trace
          url: "$${__value.raw}"
```

Les **derived fields** transforment un champ du log (ex : `trace_id`)
en lien cliquable vers Tempo : le pont logs ↔ traces.

---

## 47. Grafana Explore : explorer les logs

Explore (**Compass → Explore**) est l'outil d'investigation :

