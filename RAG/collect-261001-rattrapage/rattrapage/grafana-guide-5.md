---
id: collect-261001-rattrapage/rattrapage/grafana-guide-5
title: "Guide Grafana — Dashboards, visualisation et alerting"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/grafana_guide.md
source_anchor: ""
source_lines: [853, 1087]
sha256: 18240bdf67973f07680f186f97ab6a07a2285c57627b94fbd859a844a8335cc5
---

# Guide Grafana — Dashboards, visualisation et alerting

## 21. LogQL avancé (parsing, filtres, métriques issues des logs)

```logql
# --- Parser des logs JSON ---
{job="app"} | json | level="error" | line_format "{{.msg}} (user={{.user}})"
# Compter les erreurs par utilisateur
sum by (user) (count_over_time({job="app"} | json | level="error" [5m]))

# --- Parser des logs syslog/nginx (regexp nommée) ---
{job="nginx"}
  | regexp `(?P<ip>\\S+) - - \\[.*?\\] "(?P<method>\\S+) (?P<path>\\S+).*?" (?P<code>\\d+)`
  | code="500"
  | line_format "{{.ip}} -> {{.method}} {{.path}}"

# --- Extraire un nombre et en faire une métrique ---
{job="app"} | json | unwrap duration_ms | __error__=""
# p95 de la durée par endpoint :
quantile_over_time(0.95, {job="app"} | json | unwrap duration_ms [5m]) by (endpoint)

# --- Détecter un pattern d'attaque (brute force SSH) ---
sum by (instance) (
  count_over_time({job="sshd"} |= "Failed password" [5m])
) > 10

# --- Logs d'un conteneur précis sur les dernières 24h, triés ---
{container="api"} | logfmt | level="error"
```

> **Performance :** les requêtes LogQL scannent les chunks. Limitez la plage
> de temps, soyez précis sur les labels (`{job=...}`), évitez `|~` en début de
> pipeline sur de gros volumes. Pour des alertes fréquentes, préférez des
> **règles d'enregistrement** côté Loki (ruler) plutôt qu'une requête ad hoc
> évaluée chaque minute.

---

## 22. Datasource Zabbix (plugin)

Le plugin **Zabbix** (`alexanderzobnin-zabbix-app`) n'est pas inclus par défaut :
il s'installe depuis le catalogue.

```bash
# Installation du plugin (redémarrage nécessaire)
sudo grafana-cli plugins install alexanderzobnin-zabbix-app
sudo systemctl restart grafana-server

# Vérification
sudo grafana-cli plugins ls | grep zabbix
```

Configuration de la datasource :

| Champ | Valeur |
|---|---|
| Name | `Zabbix` |
| URL | `http://zabbix.mondomaine.fr/api_jsonrpc.php` |
| Access | `Server (default)` |
| Username / Password | compte API Zabbix dédié `grafana` (rôle lecture seule) |
| Trends | coché (utilise les trends pour les longues plages — **beaucoup** plus rapide) |
| Cache | `1h` pour les longues plages |

Bonnes pratiques Zabbix + Grafana :

- Créez un **utilisateur Zabbix dédié** avec uniquement les droits de lecture
  sur les groupes d'hôtes concernés. Jamais le compte Admin.
- Activez **Trends** : sans ça, un dashboard sur 30 jours scanne l'historique
  brut et met Zabbix à genoux.
- Les **variables** fonctionnent avec les groupes d'hôtes / hôtes / items
  Zabbix (`$groupe`, `$hote`) : requêtes de type « Zabbix hosts in group ».
- **Problèmes Zabbix → Grafana** : le panel natif du plugin affiche les
  triggers en cours ; pour l'alerting unifié, préférez recréer des règles
  Grafana sur les métriques plutôt que d'importer les triggers.

---

## 23. Datasource MySQL — métriques métier

C'est ici que Grafana devient un outil **métier** : chiffre d'affaires,
tickets ouverts, SLA, taux de disponibilité applicative — en SQL direct.

**Connections → Data sources → Add data source → MySQL.**

| Champ | Valeur |
|---|---|
| Host | `bdd.mondomaine.fr:3306` |
| Database | `metier` |
| User / Password | compte lecture seule `grafana_ro` |
| Max open connections | `10` |
| Max idle | `2` |

Créez le compte en lecture seule côté MySQL :

```sql
CREATE USER 'grafana_ro'@'%' IDENTIFIED BY '<A_COMPLETER>';
GRANT SELECT, SHOW VIEW ON metier.* TO 'grafana_ro'@'%';
FLUSH PRIVILEGES;
```

Exemples de requêtes (macros Grafana `$__timeFilter`, `$__timeGroup`) :

```sql
-- Commandes par jour sur la plage du dashboard
SELECT
  $__timeGroup(created_at, '1d') AS "time",
  COUNT(*) AS "Commandes",
  SUM(total_ht) AS "CA HT"
FROM commandes
WHERE $__timeFilter(created_at)
GROUP BY 1
ORDER BY 1;

-- Tickets ouverts par priorité (panel Table ou Bar chart)
SELECT
  priorite AS "Priorite",
  COUNT(*) AS "Tickets ouverts"
FROM tickets
WHERE statut NOT IN ('clos','resolu')
GROUP BY priorite;

-- Taux de disponibilité mensuel (panel Stat)
SELECT
  ROUND(100 * SUM(CASE WHEN statut='OK' THEN 1 ELSE 0 END) / COUNT(*), 2)
  AS "Dispo %"
FROM sondes_dispo
WHERE $__timeFilter(horodatage);
```

> **Sécurité :** le compte SQL de Grafana doit être **lecture seule**, sur les
> seules bases nécessaires. Grafana n'a aucune raison d'écrire dans votre base
> métier.

---

## 24. Datasource PostgreSQL — métriques métier

Même principe que MySQL, avec le dialecte PostgreSQL.

| Champ | Valeur |
|---|---|
| Host | `bdd.mondomaine.fr:5432` |
| Database | `metier` |
| User / Password | `grafana_ro` (lecture seule) |
| SSL Mode | `require` (ou `verify-full` avec CA) |
| Version | `14+` (active les bonnes macros) |

```sql
-- Rôle lecture seule côté PostgreSQL
CREATE USER grafana_ro WITH PASSWORD '<A_COMPLETER>';
GRANT CONNECT ON DATABASE metier TO grafana_ro;
GRANT USAGE ON SCHEMA public TO grafana_ro;
GRANT SELECT ON ALL TABLES IN SCHEMA public TO grafana_ro;
ALTER DEFAULT PRIVILEGES IN SCHEMA public GRANT SELECT ON TABLES TO grafana_ro;
```

Exemples :

```sql
-- Latence médiane d'API par endpoint (Time series)
SELECT
  $__timeGroup(horodatage, '5m') AS "time",
  endpoint AS metric,
  percentile_cont(0.5) WITHIN GROUP (ORDER BY duree_ms) AS "p50 ms"
FROM api_logs
WHERE $__timeFilter(horodatage)
GROUP BY 1, 2
ORDER BY 1;

-- File d'attente : jobs en échec dernières 24h (Stat)
SELECT COUNT(*) AS "Jobs en echec"
FROM jobs
WHERE statut = 'failed'
  AND horodatage > NOW() - INTERVAL '24 hours';

-- Heatmap des connexions par heure et par jour (Heatmap)
SELECT
  $__timeGroup(horodatage, '1h') AS "time",
  EXTRACT(DOW FROM horodatage)::int AS "jour_semaine",
  COUNT(*) AS "connexions"
FROM sessions
WHERE $__timeFilter(horodatage)
GROUP BY 1, 2;
```

> **Astuce :** la colonne temporelle doit s'appeler `time` (ou être la
> première colonne de type timestamp). Les autres colonnes deviennent des
> séries (mode « Time series ») ou des champs (mode « Table ») selon le
> **Format** choisi dans le panel : `Time series` ou `Table`.

---

## 25. Datasource InfluxDB (introduction)

InfluxDB reste présent dans beaucoup d'entreprises (historique, IoT, sondes
maison). Deux dialectes selon la version :

| Version | Langage | Réglage Grafana |
|---|---|---|
| InfluxDB 1.x | InfluxQL (proche SQL) | Query language = `InfluxQL` |
| InfluxDB 2.x / 3.x | Flux | Query language = `Flux` |

Configuration (InfluxDB 2.x) :

| Champ | Valeur |
|---|---|
| URL | `http://influxdb.mondomaine.fr:8086` |
| Organization | `mon-org` |
| Token | `<A_COMPLETER>` (token lecture seule sur le bucket) |
| Default Bucket | `supervision` |

Exemples :

```flux
// Flux : température moyenne par sonde sur 1h
from(bucket: "supervision")
  |> range(start: v.timeRangeStart, stop: v.timeRangeStop)
  |> filter(fn: (r) => r._measurement == "temperature" and r._field == "value")
  |> aggregateWindow(every: 1h, fn: mean)
```

```sql
-- InfluxQL : même idée
SELECT mean("value") FROM "temperature"
WHERE $timeFilter GROUP BY time(1h), "sonde"
```

> **Conseil de migration :** si vous démarrez un nouveau projet métriques,
> préférez **Prometheus** (écosystème, PromQL, alerting unifié). Gardez
> InfluxDB pour l'existant et l'IoT léger.

---

## 26. Datasources : tests et bonnes pratiques

Checklist à appliquer à **chaque** datasource :

