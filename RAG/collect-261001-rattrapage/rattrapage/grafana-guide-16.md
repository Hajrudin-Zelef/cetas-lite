---
id: collect-261001-rattrapage/rattrapage/grafana-guide-16
title: "Guide Grafana — Dashboards, visualisation et alerting"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["agent", "distribution", "incident", "valuation"]
source: docs/RAG/collect-261001-rattrapage/grafana_guide.md
source_anchor: ""
source_lines: [3203, 3405]
sha256: b360b49a958fe31cadf9eef84fe6f5f7900ff0ef14eded7e0665349bd51373f6
---

# Guide Grafana — Dashboards, visualisation et alerting

| Terme | Définition |
|---|---|
| Alerting unifié | Moteur d'alertes de Grafana 9+ (règles, contact points, policies) |
| Annotation | Événement affiché sur les graphes (déploiement, incident) |
| Contact point | Destination des notifications (e-mail, Slack, webhook…) |
| Dashboard | Page composée de panels répondant à une question |
| Datasource | Connexion nommée vers un backend (Prometheus, Loki, SQL…) |
| Explore | Mode de requêtage ad hoc sans créer de dashboard |
| For | Durée pendant laquelle une condition doit persister avant Firing |
| Gauge | Panel jauge (valeur vs seuils) |
| Heatmap | Densité d'une distribution dans le temps |
| Instance (alerte) | Occurrence concrète d'une règle (ex. par serveur) |
| Kiosk (mode) | Affichage plein écran sans menus (TV murale) |
| Label | Paire clé/valeur attachée à une série ou une alerte |
| LogQL | Langage de requête de Loki |
| Mute timing | Plage récurrente de suppression des notifications |
| Notification policy | Arbre de routage des alertes vers les contact points |
| Panel | Bloc de visualisation dans un dashboard |
| Playlist | Rotation automatique de dashboards |
| PromQL | Langage de requête de Prometheus |
| Provisioning | Configuration as code via YAML/JSON versionnés |
| RBAC | Contrôle d'accès basé sur les rôles (permissions fines) |
| Recording rule | Règle Prometheus pré-calculant une métrique |
| Reduce | Expression réduisant une série à un nombre |
| Resample | Expression rééchantillonnant une série |
| Row | Ligne regroupant des panels dans un dashboard |
| Scrape | Collecte périodique de métriques par Prometheus |
| Service account | Compte technique pour l'API (remplace les API keys) |
| Silence | Suppression temporaire des notifications (maintenance) |
| Snapshot | Copie figée d'un dashboard avec ses données |
| Stat | Panel affichant un grand chiffre + tendance |
| Threshold | Seuil déclenchant un changement de couleur/état |
| Transformation | Traitement des données après requête (merge, join…) |
| Variable (template) | Paramètre de dashboard (ex. `$serveur`) |

---

## 80. Quiz — 10 questions + réponses

**Q1.** Où sont stockées les métriques affichées par Grafana ?
**R1.** Nulle part dans Grafana : elles restent dans les backends (Prometheus,
Loki, SQL…). Grafana ne stocke que sa configuration (dashboards,
utilisateurs, règles) dans sa base interne.

**Q2.** Pourquoi faut-il renseigner `root_url` dans `grafana.ini` ?
**R2.** Parce que les liens insérés dans les notifications d'alerte et les
redirections OAuth sont construits à partir de cette URL. Sans elle, ils
pointent vers `localhost` et sont inutilisables.

**Q3.** Quelle différence entre le mode d'accès `Server` et `Browser` d'une
datasource ?
**R3.** En mode Server, le backend Grafana interroge la datasource (les
secrets restent côté serveur — recommandé). En mode Browser, c'est le
navigateur de chaque utilisateur qui interroge directement (expose l'URL et
les identifiants).

**Q4.** Pourquoi une règle d'alerte ne peut-elle pas utiliser les variables
de dashboard (`$serveur`) ?
**R4.** Parce qu'une alerte doit être déterministe : elle s'évalue en tâche
de fond, sans dashboard ouvert ni sélection d'utilisateur. Il faut écrire la
requête sans variable (regex explicite si besoin).

**Q5.** À quoi sert le paramètre `for` dans une règle d'alerte ?
**R5.** À exiger que la condition reste vraie pendant une durée donnée avant
de passer en Firing, ce qui filtre les pics transitoires et réduit les faux
positifs.

**Q6.** Que signifie `group_by: [alertname]` dans une notification policy ?
**R6.** Que les instances d'alerte partageant le même `alertname` sont
regroupées en **une seule** notification (ex. « 12 disques pleins » au lieu
de 12 e-mails).

**Q7.** Quelle commande sauvegarde proprement la base SQLite de Grafana ?
**R7.** `sqlite3 /var/lib/grafana/grafana.db ".backup '/chemin/sauvegarde.db'"`.
Un simple `cp` à chaud risque de produire une base corrompue.

**Q8.** Citez deux choses à vérifier en premier quand une datasource ne
répond plus.
**R8.** (1) La joignabilité du backend **depuis le serveur Grafana**
(`curl`/`nc`) ; (2) les identifiants, testés à la main hors Grafana. Dans
80 % des cas, c'est le réseau ou les identifiants, pas Grafana.

**Q9.** Pourquoi ne faut-il jamais appliquer `rate()` à une gauge ?
**R9.** Parce que `rate()` calcule une variation par seconde en supposant un
**compteur** qui ne fait qu'augmenter. Sur une gauge (température, niveau),
le résultat est absurde. On utilise `avg_over_time`, `max_over_time`, etc.

**Q10.** Que doit contenir une sauvegarde complète de Grafana pour permettre
une restauration ?
**R10.** La base interne (SQLite/PostgreSQL), `/etc/grafana` (grafana.ini,
provisioning, ldap.toml), les dashboards JSON versionnés (Git), la
`secret_key`, et la liste des plugins. Sans la `secret_key`, les secrets des
datasources sont illisibles après restauration.

---

## 81. Pour aller plus loin

**Écosystème Grafana :**

- **Grafana Mimir** : backend Prometheus scalable (long terme, multi-tenant).
- **Grafana Tempo** : tracing distribué (corrélation métriques → logs → traces).
- **Grafana Pyroscope** : profiling continu (pourquoi ce CPU ?).
- **Grafana OnCall** : gestion d'astreinte (planning, escalades).
- **Grafana Alloy** : collecteur unifié (successeur de l'agent Grafana).

**Sujets à creuser :**

- Recording rules Prometheus pour les dashboards lents (pré-calcul).
- Exemplars : du graphe au log à la trace en un clic.
- LBC (library panels) : panels réutilisables entre dashboards.
- Tests de charge des dashboards (k6) avant une démo à la direction.
- Haute disponibilité de Grafana (plusieurs instances + base partagée).

**Documentation officielle :** https://grafana.com/docs/grafana/latest/
**Communauté :** https://community.grafana.com/

---

## 82. Annexe A — `grafana.ini` de référence commenté

```ini
#################### Serveur ####################
[server]
protocol = http
http_addr = 127.0.0.1
http_port = 3000
root_url = https://grafana.mondomaine.fr/
;serve_from_sub_path = false

#################### Chemins ####################
[paths]
data = /var/lib/grafana
logs = /var/log/grafana
plugins = /var/lib/grafana/plugins
provisioning = /etc/grafana/provisioning

#################### Base interne ####################
[database]
# SQLite par défaut ; PostgreSQL recommandé en production :
#host = 127.0.0.1:5432
#name = grafana
#user = grafana
#password = '<A_COMPLETER>'
#ssl_mode = require

#################### Logs ####################
[log]
mode = console file
level = info
[log.file]
log_rotate = true
daily_rotate = true
max_days = 7

#################### Sécurité ####################
[security]
admin_user = admin
admin_password = '<A_COMPLETER>'
allow_sign_up = false
secret_key = '<A_COMPLETER_32_CARACTERES_ALEATOIRES>'
login_maximum_inactive_lifetime_duration = 7d
login_maximum_lifetime_duration = 30d
cookie_secure = true
cookie_samesite = lax
;hide_version = true

[auth.anonymous]
enabled = false

#################### Utilisateurs ####################
[users]
default_theme = dark
allow_sign_up = false

#################### SMTP ####################
[smtp]
enabled = true
host = smtp.mondomaine.fr:587
user = grafana@mondomaine.fr
password = '<A_COMPLETER>'
from_address = grafana@mondomaine.fr
from_name = Grafana Supervision
startTLS_policy = MandatoryStartTLS

#################### Métriques internes ####################
[metrics]
enabled = true
basic_auth_username = '<A_COMPLETER>'
basic_auth_password = '<A_COMPLETER>'

#################### Dashboards ####################
[dashboards]
default_timezone = Europe/Paris

#################### Alerting ####################
[unified_alerting]
enabled = true
# Évaluation et rétention :
#evaluation_timeout = 30s
#max_attempts = 3
```

