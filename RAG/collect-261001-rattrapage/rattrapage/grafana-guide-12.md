---
id: collect-261001-rattrapage/rattrapage/grafana-guide-12
title: "Guide Grafana — Dashboards, visualisation et alerting"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["arr", "datacenter"]
source: docs/RAG/collect-261001-rattrapage/grafana_guide.md
source_anchor: ""
source_lines: [2401, 2648]
sha256: 1b120ffff1a9fddc8010aeef144651a2ef5f4c71c1a72d407431829417fcbabb
---

# Guide Grafana — Dashboards, visualisation et alerting

```bash
# Vérification rapide d'exposition
ss -ltnp | grep -E "3000|443"
# Le port 3000 ne doit être visible que sur 127.0.0.1
curl -s -o /dev/null -w "%{http_code}\n" http://localhost:3000/api/health
```

---

## 61. Superviser Grafana lui-même (métriques internes)

Grafana expose ses propres métriques Prometheus sur `/metrics` (à activer) :

```ini
[metrics]
enabled = true
# Restreindre l'accès : ne pas exposer publiquement
basic_auth_username = '<A_COMPLETER>'
basic_auth_password = '<A_COMPLETER>'
```

Job Prometheus :

```yaml
scrape_configs:
  - job_name: grafana
    static_configs:
      - targets: ['grafana.mondomaine.fr:3000']
    metrics_path: /metrics
    basic_auth:
      username: '<A_COMPLETER>'
      password: '<A_COMPLETER>'
```

Métriques clés à surveiller :

```promql
# Disponibilité de Grafana (blackbox ou scrape up)
up{job="grafana"}

# Requêtes API en erreur (5xx)
sum(rate(grafana_api_response_status_total{code=~"5.."}[5m]))

# Latence p95 des réponses API
histogram_quantile(0.95, sum by (le) (rate(grafana_api_response_time_seconds_bucket[5m])))

# Évaluations d'alertes en échec
sum(rate(grafana_alerting_rule_evaluation_failures_total[5m])) > 0

# Notifications en échec
sum(rate(grafana_alerting_notification_failed_total[5m])) > 0

# Nombre de sessions actives (ordre de grandeur d'usage)
grafana_stat_total_users
grafana_stat_total_dashboards
grafana_stat_total_alerts
```

**Règles d'alerte minimales sur Grafana lui-même :**

1. `up{job="grafana"} == 0` depuis 2 min → critical (l'outil de supervision
   est en panne !).
2. `sum(rate(grafana_alerting_notification_failed_total[5m])) > 0` → warning
   (les alertes ne partent plus — voir section 74).
3. Latence API p95 > 2 s depuis 10 min → warning.

> **Paradoxe du superviseur :** si Grafana tombe, qui vous prévient ?
> Prévoyez un **canal externe** : une sonde blackbox indépendante (un second
> Prometheus, un service externe type health check) qui surveille
> `https://grafana.mondomaine.fr/api/health`.

---

## 62. Sauvegarde : base SQLite

Avec SQLite (défaut), la base est un simple fichier : `/var/lib/grafana/grafana.db`.
**Ne copiez jamais le fichier à chaud avec `cp`** : risque de base corrompue.

```bash
#!/bin/bash
# /usr/local/bin/backup-grafana.sh
set -euo pipefail
DATE=$(date +%F)
DEST="/srv/backups/grafana/$DATE"
mkdir -p "$DEST"

# 1. Dump cohérent de la base SQLite (verrouille proprement)
sqlite3 /var/lib/grafana/grafana.db ".backup '$DEST/grafana.db'"

# 2. Configuration + provisioning + dashboards versionnés
tar -czf "$DEST/grafana-etc.tar.gz" -C /etc grafana

# 3. Plugins installés (réinstallables, mais pratique)
tar -czf "$DEST/grafana-plugins.tar.gz" -C /var/lib/grafana plugins 2>/dev/null || true

# 4. Rétention : 30 jours
find /srv/backups/grafana -maxdepth 1 -type d -mtime +30 -exec rm -rf {} +

echo "Sauvegarde Grafana OK : $DEST"
```

```bash
sudo chmod +x /usr/local/bin/backup-grafana.sh
# Cron quotidien à 2h
echo "0 2 * * * root /usr/local/bin/backup-grafana.sh >> /var/log/grafana-backup.log 2>&1" \
  | sudo tee /etc/cron.d/grafana-backup
```

Vérifiez que `sqlite3` est installé (`apt install sqlite3`). Testez le dump
restauré **une fois par trimestre** (section 65) : une sauvegarde non testée
n'est pas une sauvegarde.

---

## 63. Sauvegarde : base PostgreSQL

Avec PostgreSQL, utilisez les outils natifs (dump cohérent à chaud) :

```bash
#!/bin/bash
# /usr/local/bin/backup-grafana-pg.sh
set -euo pipefail
DATE=$(date +%F_%H%M)
DEST="/srv/backups/grafana/$DATE"
mkdir -p "$DEST"

# 1. Dump logique (format custom, parallélisable à la restauration)
pg_dump -h 127.0.0.1 -U grafana -Fc -f "$DEST/grafana.dump" grafana

# 2. Configuration + provisioning
tar -czf "$DEST/grafana-etc.tar.gz" -C /etc grafana

# 3. Rétention : 30 jours
find /srv/backups/grafana -maxdepth 1 -type d -mtime +30 -exec rm -rf {} +

echo "Sauvegarde Grafana (PG) OK : $DEST"
```

Authentification sans mot de passe interactif : fichier `/root/.pgpass`
(`chmod 600`) :

```
127.0.0.1:5432:grafana:grafana:<A_COMPLETER>
```

Alternative robuste : `pg_basebackup` pour une sauvegarde physique complète
du cluster, en complément du dump logique.

> **Point de vigilance :** la `secret_key` (section 13) **doit** être
> sauvegardée avec la base. Sans elle, les secrets des datasources
> provisionnées (mots de passe, tokens) sont illisibles après restauration.
> Stockez-la dans le coffre **et** dans la sauvegarde chiffrée.

---

## 64. Sauvegarde : provisioning, plugins, stratégie 3-2-1

Ce que la base ne contient pas (ou pas bien) :

| Élément | Où | Sauvegarde |
|---|---|---|
| `grafana.ini` | `/etc/grafana/grafana.ini` | Git + tar quotidien |
| Provisioning YAML | `/etc/grafana/provisioning/` | **Git** (source de vérité) |
| Dashboards JSON | `/etc/grafana/dashboards/` | **Git** |
| Plugins | `/var/lib/grafana/plugins/` | Liste (`grafana-cli plugins ls`) + tar |
| `ldap.toml`, certificats | `/etc/grafana/` | Git privé ou tar chiffré |
| `secret_key` | `grafana.ini` | Coffre + sauvegarde chiffrée |

**Stratégie 3-2-1** : 3 copies, 2 supports différents, 1 hors site.

```bash
# Exemple : copie chiffrée vers un stockage distant (après le backup local)
tar -czf - -C /srv/backups/grafana "$(date +%F)" | \
  gpg --encrypt --recipient supervision@mondomaine.fr | \
  ssh backup@nas.mondomaine.fr "cat > /backups/grafana-$(date +%F).tar.gz.gpg"
```

Checklist sauvegarde :

- [ ] Backup quotidien automatisé (base + /etc/grafana).
- [ ] Git poussé à chaque changement de provisioning/dashboard.
- [ ] `secret_key` et mots de passe dans le coffre.
- [ ] Copie hors site (NAS, S3, autre datacenter).
- [ ] **Test de restauration trimestriel** (section 65) avec compte-rendu.

---

## 65. Restauration — procédure complète

Scénario : serveur Grafana perdu, on reconstruit depuis les sauvegardes.

```bash
# 1. Réinstaller (même version majeure !)
# Voir section 6, puis NE PAS démarrer le service tout de suite.

# 2. Restaurer la configuration
sudo tar -xzf /srv/backups/grafana/<DATE>/grafana-etc.tar.gz -C /

# 3a. SQLite : restaurer le fichier (service arrêté)
sudo systemctl stop grafana-server
sudo cp /srv/backups/grafana/<DATE>/grafana.db /var/lib/grafana/grafana.db
sudo chown grafana:grafana /var/lib/grafana/grafana.db

# 3b. PostgreSQL : recréer la base et restaurer le dump
sudo -u postgres psql -c "DROP DATABASE IF EXISTS grafana;"
sudo -u postgres psql -c "CREATE DATABASE grafana OWNER grafana;"
pg_restore -h 127.0.0.1 -U grafana -d grafana /srv/backups/grafana/<DATE>/grafana.dump

# 4. Plugins
sudo tar -xzf /srv/backups/grafana/<DATE>/grafana-plugins.tar.gz -C /var/lib/grafana
# ou réinstallation propre :
sudo grafana-cli plugins install alexanderzobnin-zabbix-app

# 5. Démarrer et vérifier
sudo systemctl start grafana-server
sleep 5
curl -s http://localhost:3000/api/health | python3 -m json.tool
sudo tail -n 30 /var/log/grafana/grafana.log
```

Points de contrôle post-restauration :

- [ ] `/api/health` → `database: ok`.
- [ ] Login admin OK.
- [ ] Datasources : **Save & test** au vert sur chacune.
- [ ] Dashboards présents (compter : `grafana_stat_total_dashboards`).
- [ ] Règles d'alerte présentes et évaluées (Alerting → état).
- [ ] Test d'envoi d'une notification (contact point → Test).

> **Restaurer sur une version différente** : Grafana migre le schéma à la
> montée de version, jamais à la descente. Restaurez toujours sur une version
> **égale ou supérieure** à celle de la sauvegarde.

---

## 66. Mise à jour de Grafana — procédure sans casse

```bash
# 0. Sauvegarde complète AVANT (sections 62-64) — non négociable
sudo /usr/local/bin/backup-grafana.sh

# 1. Lire la release note (breaking changes, migrations)
# https://grafana.com/docs/grafana/latest/breaking-changes/

# 2. Mise à jour via APT
sudo apt-get update
sudo apt-get install -y --only-upgrade grafana

