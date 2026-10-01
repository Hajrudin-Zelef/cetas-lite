---
id: collect-261001-rattrapage/rattrapage/grafana-guide-2
title: "Guide Grafana — Dashboards, visualisation et alerting"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["attention", "datacenter", "incident"]
source: docs/RAG/collect-261001-rattrapage/grafana_guide.md
source_anchor: ""
source_lines: [164, 346]
sha256: 592743fb58ca8d88447452825980402a343e7a4c39ebf2d7eaa88eacc43b8193
---

# Guide Grafana — Dashboards, visualisation et alerting

- **Dashboard** : une page composée de panels. Il répond idéalement à **une
  seule question** (« l'état de santé de mes serveurs ce matin ? »).
- **Panel** : un bloc de visualisation (courbe, jauge, tableau…). Chaque panel
  exécute une ou plusieurs requêtes et affiche le résultat.
- **Row** (ligne) : regroupement visuel de panels, repliable.
- **Variable (template)** : un paramètre du dashboard (`$serveur`, `$datacenter`).
  L'utilisateur choisit des valeurs en haut du dashboard, les requêtes des
  panels s'adaptent. C'est ce qui rend un dashboard **réutilisable** au lieu
  d'en cloner un par serveur.
- **Annotation** : un événement ponctuel (déploiement, incident) affiché comme
  une ligne verticale sur les graphes.
- **Playlist** : une rotation automatique de dashboards (TV murale).
- **Snapshot** : une copie figée d'un dashboard (données incluses), partageable
  sans accès aux datasources.

Cycle de vie conseillé : dashboard construit à la main → exporté en JSON →
versionné en Git → **provisionné** automatiquement au démarrage de Grafana.

---

## 4. Concepts : alerting unifié

Depuis Grafana 9, l'**alerting unifié** remplace l'ancien système d'alertes par
panel. Vocabulaire :

- **Règle d'alerte (alert rule)** : une ou plusieurs requêtes + des expressions
  (reduce, math) + une condition. Évaluée périodiquement (ex. toutes les minutes).
- **Instance d'alerte** : une alerte concrète produite par la règle (ex. « CPU >
  90 % sur srv-web-03 »). Une règle peut produire N instances (une par série).
- **États** : `Normal` → `Pending` (la condition est vraie depuis moins que la
  durée `for`) → `Firing` (alerte déclenchée) → `Resolved`.
- **Contact point** : la destination d'une notification (e-mail, Slack, webhook…).
- **Notification policy** : l'arbre de routage qui décide, selon les labels,
  quel contact point reçoit quoi, avec quel regroupement (`group_by`) et quels
  délais (`group_wait`, `group_interval`, `repeat_interval`).
- **Silence** : met en sourdine des notifications correspondant à des labels,
  pour une durée donnée (maintenance planifiée).
- **Mute timing** : plage récurrente de silence (ex. « pas de notifications
  non critiques le week-end »).

L'alerting unifié gère aussi les alertes **externes** (ex. Alertmanager
Prometheus) et les règles **enregistrées dans les datasources compatibles**
(Prometheus, Loki, Mimir) : Grafana peut les afficher et les router sans les
réévaluer lui-même.

---

## 5. Architecture de Grafana

```
┌─────────────┐      HTTP(S)      ┌──────────────┐      ┌────────────────┐
│  Navigateur │ ◄──────────────► │    Grafana   │ ───► │  Datasources   │
│ (frontend)  │                  │   (backend   │      │ Prometheus,    │
└─────────────┘                  │    Go + API) │ ◄─── │ Loki, Zabbix,  │
                                 └──────┬───────┘      │ SQL, InfluxDB… │
                                        │ SQLite /     └────────────────┘
                                        │ Postgres
                                 ┌──────▼───────┐
                                 │ Base interne │
                                 │ (dashboards, │
                                 │ users, alert │
                                 │  rules…)     │
                                 └──────────────┘
```

- **Backend Go** : sert l'API HTTP, exécute les requêtes datasource en mode
  `Server`, évalue les règles d'alerte, envoie les notifications.
- **Frontend** : application web (React) affichée dans le navigateur.
- **Base interne** (SQLite par défaut, PostgreSQL/MySQL recommandé en
  production) : stocke dashboards, utilisateurs, permissions, règles d'alerte,
  historique. **Ce n'est pas là que vivent vos métriques.**
- **Provisioning** : au démarrage, Grafana lit `/etc/grafana/provisioning/` et
  crée/met à jour datasources, dashboards, contact points, etc.

Conséquence pratique : **sauvegarder Grafana = sauvegarder sa base interne +
le dossier provisioning** (sections 62-64). Les métriques elles, se
sauvegardent côté backends (Prometheus, Loki…).

---

## 6. Installation sur Debian/Ubuntu (dépôt officiel)

On installe depuis le **dépôt APT officiel** de Grafana : mises à jour via
`apt upgrade`, signatures vérifiées, pas de binaire téléchargé à la main.

```bash
# 1. Pré-requis
sudo apt-get install -y apt-transport-https software-properties-common wget gnupg

# 2. Clé GPG officielle Grafana
sudo mkdir -p /etc/apt/keyrings/
wget -q -O - https://apt.grafana.com/gpg.key | gpg --dearmor | \
  sudo tee /etc/apt/keyrings/grafana.gpg > /dev/null

# 3. Dépôt stable
echo "deb [signed-by=/etc/apt/keyrings/grafana.gpg] https://apt.grafana.com stable main" | \
  sudo tee /etc/apt/sources.list.d/grafana.list

# 4. Installation
sudo apt-get update
sudo apt-get install -y grafana

# 5. Service systemd : démarrage auto + lancement
sudo systemctl daemon-reload
sudo systemctl enable --now grafana-server

# 6. Vérification
systemctl status grafana-server --no-pager
ss -ltnp | grep 3000
curl -s http://localhost:3000/login | head -c 200; echo
```

Fichiers installés par le paquet :

| Chemin | Rôle |
|---|---|
| `/usr/sbin/grafana-server` | Binaire |
| `/etc/grafana/grafana.ini` | Configuration (à personnaliser) |
| `/etc/grafana/provisioning/` | Provisioning as code |
| `/var/lib/grafana/` | Base SQLite + plugins + données |
| `/var/log/grafana/grafana.log` | Journal applicatif |
| `/usr/share/grafana/` | Frontend statique |

> **Note :** le paquet crée l'utilisateur système `grafana` (sans shell).
> Ne lancez jamais Grafana en root.

Vérifier la version installée :

```bash
grafana-server -v
# ou via l'API (après premier login) :
curl -s http://admin:<A_COMPLETER>@localhost:3000/api/health | python3 -m json.tool
```

---

## 7. Installation via Docker (aperçu)

Utile pour un test ou un labo. En production d'équipe, préférez le paquet APT
(intégration systemd, chemins standards, sauvegardes simples).

```bash
docker volume create grafana-data

docker run -d --name grafana \
  --restart unless-stopped \
  -p 3000:3000 \
  -e GF_SECURITY_ADMIN_USER=admin \
  -e GF_SECURITY_ADMIN_PASSWORD='<A_COMPLETER>' \
  -e GF_INSTALL_PLUGINS=grafana-clock-panel \
  -v grafana-data:/var/lib/grafana \
  grafana/grafana:11.5.2
```

Points d'attention Docker :

- Les variables d'environnement `GF_<SECTION>_<CLE>` surchargent `grafana.ini`
  (ex. `GF_SERVER_HTTP_PORT`, `GF_DATABASE_PATH`). Pratique, mais documentez-les.
- Sans volume, **toute la configuration est perdue** à la suppression du
  conteneur.
- Épinglez un tag de version précis (`:11.5.2`), jamais `:latest` en production.

---

## 8. Premier lancement et première connexion

1. Ouvrez `http://<serveur>:3000` dans un navigateur.
2. Identifiants initiaux : `admin` / `admin`.
3. Grafana **impose** le changement du mot de passe admin à la première
   connexion. Choisissez un mot de passe fort, stocké dans votre coffre.
4. L'assistant de bienvenue propose d'ajouter une datasource : cliquez
   **« Add your first data source »** et choisissez (ex. Prometheus).

Premiers réglages immédiats (roue dentée → Administration) :

- **General → Default preferences** : thème (Dark/Light), fuseau horaire
  (`Europe/Paris` si vos équipes sont en France — voir section 10).
- **Users** : créez les comptes nominatifs, ne partagez pas le compte `admin`.
- **Authentication** : désactivez l'inscription publique si elle est active
  (`allow_sign_up = false`, voir section 13).

Vérification santé via l'API :

