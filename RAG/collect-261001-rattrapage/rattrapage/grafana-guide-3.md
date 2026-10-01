---
id: collect-261001-rattrapage/rattrapage/grafana-guide-3
title: "Guide Grafana — Dashboards, visualisation et alerting"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/grafana_guide.md
source_anchor: ""
source_lines: [347, 613]
sha256: 1136b57189ce270b1a1eabdf5c34f1ff79e4ef293f92b37ad92762c8d58be867
---

# Guide Grafana — Dashboards, visualisation et alerting

```bash
curl -s http://localhost:3000/api/health | python3 -m json.tool
# {"commit":"...","database":"ok","version":"11.5.2"}
```

Si `database` vaut autre chose que `ok`, regardez `/var/log/grafana/grafana.log`
avant toute chose.

---

## 9. Tour d'horizon de l'interface

Barre latérale gauche (icônes) :

| Icône / Menu | Contenu |
|---|---|
| **Dashboards** | Parcourir, créer, importer, playlists |
| **Explore** | Requêtage ad hoc (PromQL/LogQL/SQL) sans créer de dashboard |
| **Alerting** | Règles, contact points, policies, silences |
| **Connections → Data sources** | Ajouter/configurer/tester les datasources |
| **Administration → Users / Teams** | Comptes, équipes, organisations |
| **Administration → Settings** | Lecture seule de la config effective |

Le menu **Explore** mérite d'être connu : c'est l'établi du sysadmin.
On y teste une requête en 10 secondes, avec complétion, avant de l'intégrer
dans un panel. Réflexe : *Explore d'abord, dashboard ensuite.*

Raccourcis utiles : `/` ou `Ctrl+K` pour la recherche globale de dashboards.
Le sélecteur de plage de temps (en haut à droite) s'applique à tout le
dashboard ; apprenez `Ctrl+Z`… non, il n'y a pas d'annulation globale :
**pensez à sauvegarder** (icône disquette ou `Ctrl+S` dans l'éditeur).

---

## 10. Configuration `grafana.ini` — serveur et chemins

Le fichier `/etc/grafana/grafana.ini` est en format INI (`[section]` puis
`clé = valeur`). Toute ligne commençant par `;` est commentée. **Après chaque
modification : `systemctl restart grafana-server`.**

```ini
[server]
# Protocole et port d'écoute
protocol = http
http_port = 3000

# Adresse d'écoute : localhost si un reverse proxy est devant,
# 0.0.0.0 si Grafana est exposé directement (alors mettez HTTPS, section 14)
http_addr = 127.0.0.1

# URL publique vue par les navigateurs (utilisée dans les liens des alertes !)
# INDISPENSABLE : sans elle, les liens "voir l'alerte" pointent vers localhost
root_url = https://grafana.mondomaine.fr/

# Domaine utilisé pour les cookies (à ajuster si sous-domaine)
;domain = grafana.mondomaine.fr

[paths]
# Rarement à toucher avec le paquet APT
data = /var/lib/grafana
logs = /var/log/grafana
plugins = /var/lib/grafana/plugins
provisioning = /etc/grafana/provisioning

[users]
# Page d'accueil par défaut après login
default_theme = dark
```

> **Piège classique :** `root_url` mal renseignée = liens cassés dans les
> e-mails d'alerte et redirections OAuth en échec. C'est le réglage n°1 à
> vérifier quand « les liens des alertes ne marchent pas ».

Fuseau horaire par défaut des dashboards (pratique pour une équipe en France) :

```ini
[date_formats]
# Laisser défaut sauf besoin spécifique

[dashboards]
# Fuseau par défaut proposé à la création (l'utilisateur peut le changer)
default_timezone = Europe/Paris
```

---

## 11. Configuration — base de données interne

Par défaut Grafana utilise **SQLite** (`/var/lib/grafana/grafana.db`).
C'est suffisant pour : une petite équipe, un seul nœud, < 50 dashboards.

Passez à **PostgreSQL** (ou MySQL) quand : plusieurs dizaines d'utilisateurs,
alerting intensif, besoin de sauvegarde à chaud fiable, ou haute
disponibilité en perspective.

```ini
[database]
# SQLite (défaut) :
#path = /var/lib/grafana/grafana.db

# PostgreSQL (recommandé en production) :
path =
host = 127.0.0.1:5432
name = grafana
user = grafana
password = '<A_COMPLETER>'
ssl_mode = require
# Options utiles :
#max_open_conn = 10
#max_idle_conn = 2
#conn_max_lifetime = 14400

# MySQL/MariaDB (alternative) :
#host = 127.0.0.1:3306
#name = grafana
#user = grafana
#password = '<A_COMPLETER>'
```

Création de la base côté PostgreSQL :

```sql
CREATE USER grafana WITH PASSWORD '<A_COMPLETER>';
CREATE DATABASE grafana OWNER grafana;
```

Grafana crée et migre le schéma **tout seul** au démarrage (`migrations`).
Ne touchez jamais aux tables à la main.

**Migration SQLite → PostgreSQL :** il n'existe pas d'outil officiel
automatique. Procédure fiable : exporter les dashboards en JSON (ou via le
provisioning), basculer la config, réimporter. D'où l'intérêt de gérer les
dashboards **as code dès le départ** (section 49).

---

## 12. Configuration — journalisation

```ini
[log]
# Niveaux : debug, info, warn, error, critical
mode = console file
level = info

[log.file]
log_rotate = true
max_lines = 1000000
max_size_shift = 28        # 256 Mo par fichier (2^28)
daily_rotate = true
max_days = 7               # rétention 7 jours
```

En dépannage, passez temporairement en `debug` :

```bash
# Sans redémarrer : via l'API admin (pratique !)
curl -s -X POST -H "Content-Type: application/json" \
  -d '{"level":"debug"}' \
  http://admin:<A_COMPLETER>@localhost:3000/api/admin/log/level
# Revenir en info ensuite :
curl -s -X POST -H "Content-Type: application/json" \
  -d '{"level":"info"}' \
  http://admin:<A_COMPLETER>@localhost:3000/api/admin/log/level
```

Les logs contiennent aussi les **requêtes datasource en échec** et les
**évaluations d'alertes** : c'est le premier endroit à regarder quand
« ça ne marche plus » (voir section 73-75).

> Si vous centralisez vos logs (Loki, rsyslog), ajoutez `/var/log/grafana/*.log`
> à votre collecte : superviser l'outil de supervision, c'est la section 61.

---

## 13. Configuration — sécurité applicative (secret, cookies, sessions)

```ini
[security]
# Identifiants admin créés au premier démarrage si la base est vide.
# Mettez un mot de passe fort ; Grafana force de toute façon son changement.
admin_user = admin
admin_password = '<A_COMPLETER>'

# INTERDIT en production : l'auto-inscription crée des comptes Viewer publics
allow_sign_up = false

# Clé de chiffrement des secrets (datasources, tokens). CRITIQUE :
# - générez-la une fois, sauvegardez-la avec la base (section 62-64)
# - si vous la perdez/changez, les secrets provisionnés deviennent illisibles
secret_key = '<A_COMPLETER_32_CARACTERES_ALEATOIRES>'

# Durée de vie des sessions de login
login_maximum_inactive_lifetime_duration = 7d
login_maximum_lifetime_duration = 30d
token_rotation_interval_minutes = 10

# Désactiver la page de login basique si vous utilisez uniquement OAuth/LDAP
;disable_login_form = true

# Masquer la version dans les en-têtes et l'API (hygiène de base)
;hide_version = true

[auth]
# Interdire le login anonyme sauf besoin kiosque (voir section 72)
disable_login_form = false

[auth.anonymous]
enabled = false
# Si activé pour un écran mural : limitez STRICTEMENT l'organisation et le rôle
#enabled = true
#org_name = Kiosque
#org_role = Viewer
```

Générer une `secret_key` robuste :

```bash
python3 -c "import secrets; print(secrets.token_hex(32))"
```

Cookies (à durcir si HTTPS, voir sections 14 et 59) :

```ini
[security]
cookie_secure = true      # cookie transmis uniquement en HTTPS
cookie_samesite = lax     # 'lax' par défaut ; 'strict' si pas d'iframe
```

---

## 14. Configuration — HTTPS/TLS natif

Deux options : TLS natif dans Grafana, ou TLS terminé par un reverse proxy
(Nginx). **En production, préférez le reverse proxy** (section 59) : gestion
centralisée des certificats, Let's Encrypt, en-têtes de sécurité.

TLS natif (utile en labo ou sans proxy) :

```ini
[server]
protocol = https
http_port = 3000
cert_file = /etc/grafana/cert.pem
cert_key = /etc/grafana/key.pem
# Pour forcer TLS 1.2 minimum :
;min_tls_version = "1.2"
```

```bash
# Permissions : le binaire tourne en tant que 'grafana'
sudo chown root:grafana /etc/grafana/key.pem
sudo chmod 640 /etc/grafana/key.pem
sudo systemctl restart grafana-server
```

Vérification :

```bash
curl -vk https://localhost:3000/login 2>&1 | grep -E "SSL|subject|issuer"
```

---

## 15. Configuration — SMTP (notifications e-mail)

Indispensable pour les alertes e-mail et les invitations d'utilisateurs.

