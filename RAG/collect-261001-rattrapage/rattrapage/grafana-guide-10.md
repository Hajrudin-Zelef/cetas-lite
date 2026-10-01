---
id: collect-261001-rattrapage/rattrapage/grafana-guide-10
title: "Guide Grafana — Dashboards, visualisation et alerting"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/grafana_guide.md
source_anchor: ""
source_lines: [1892, 2161]
sha256: c50701898c0952c50ac1f7762e301a000b094e00cd9c0f59a55f1c6ef38d9e7a
---

# Guide Grafana — Dashboards, visualisation et alerting

Comportement : au démarrage, Grafana **crée ou met à jour** les objets
provisionnés. Il ne supprime pas ce qui a été retiré du YAML sauf option
explicite (`deleteDatasources`, dossiers dashboards avec `purge`).

> **Règle d'équipe :** tout dashboard « officiel » (supervision, métier) est
> provisionné. Les dashboards personnels/brouillons restent éditables à la
> main dans l'interface.

---

## 50. Provisioning des datasources (YAML)

`/etc/grafana/provisioning/datasources/datasources.yaml` :

```yaml
apiVersion: 1

# Optionnel : supprime les datasources qui ne sont plus dans ce fichier
deleteDatasources:
  - name: Old-Prometheus
    orgId: 1

datasources:
  - name: Prometheus
    type: prometheus
    access: proxy
    url: http://prometheus.mondomaine.fr:9090
    isDefault: true
    editable: false                 # verrouille l'édition dans l'UI
    jsonData:
      httpMethod: POST
      timeInterval: 15s
      queryTimeout: 60s
      manageAlerts: true            # alerting unifié sur cette datasource
    # Secrets : via variables d'environnement (jamais en clair dans Git !)
    secureJsonData:
      basicAuthPassword: $PROM_BASIC_PASSWORD

  - name: Loki
    type: loki
    access: proxy
    url: http://loki.mondomaine.fr:3100
    jsonData:
      maxLines: 1000

  - name: MySQL-Metier
    type: mysql
    url: bdd.mondomaine.fr:3306
    user: grafana_ro
    jsonData:
      database: metier
      maxOpenConns: 10
      maxIdleConns: 2
      connMaxLifetime: 14400
    secureJsonData:
      password: $MYSQL_GRAFANA_PASSWORD
```

Les variables `$VAR` sont résolues depuis l'environnement du processus
`grafana-server` (fichier `/etc/default/grafana-server` ou `Environment=` dans
l'override systemd). Appliquer :

```bash
sudo systemctl restart grafana-server
sudo tail -n 50 /var/log/grafana/grafana.log | grep -i provisioning
```

---

## 51. Provisioning des dashboards (YAML)

Principe en deux temps : le YAML dit **où** sont les dashboards, les fichiers
JSON sont les dashboards eux-mêmes.

`/etc/grafana/provisioning/dashboards/dashboards.yaml` :

```yaml
apiVersion: 1

providers:
  - name: 'Systemes'
    orgId: 1
    folder: 'Supervision'          # dossier cible dans Grafana
    type: file
    disableDeletion: false
    updateIntervalSeconds: 30      # relecture des JSON toutes les 30 s
    allowUiUpdates: false          # false = la source de vérité est Git
    options:
      path: /etc/grafana/dashboards/systemes
```

Workflow :

```bash
# 1. Construire le dashboard dans l'UI, l'exporter (Share → Export → JSON)
# 2. Le déposer versionné :
sudo mkdir -p /etc/grafana/dashboards/systemes
sudo cp supervision-serveurs.json /etc/grafana/dashboards/systemes/
# 3. Git :
cd /srv/git/grafana && git add . && git commit -m "Dashboard serveurs v3" && git push
# 4. Appliquer (updateIntervalSeconds le fait tout seul sous 30 s,
#    sinon restart pour forcer)
```

> `allowUiUpdates: false` + édition dans l'UI = modifications **écrasées**
> au prochain rechargement. Si vous voulez autoriser des ajustements locaux,
> mettez `true`, mais documentez que Git reste la référence.

---

## 52. Provisioning de l'alerting (contact points, policies, règles)

`/etc/grafana/provisioning/alerting/contact-points.yaml` :

```yaml
apiVersion: 1
contactPoints:
  - orgId: 1
    name: equipe-systemes
    receivers:
      - uid: systemes-mail
        type: email
        settings:
          addresses: systemes@mondomaine.fr
          subject: '[{{ .Status | toUpper }}] {{ .GroupLabels.alertname }}'
      - uid: systemes-slack
        type: slack
        settings:
          url: $SLACK_WEBHOOK_URL
          channel: '#alertes-systemes'
```

`/etc/grafana/provisioning/alerting/policies.yaml` :

```yaml
apiVersion: 1
policies:
  - orgId: 1
    receiver: equipe-systemes
    group_by: ['alertname']
    group_wait: 30s
    group_interval: 5m
    repeat_interval: 4h
    routes:
      - receiver: astreinte-tel
        object_matchers:
          - [severity, =, critical]
        group_wait: 0s
        repeat_interval: 1h
```

`/etc/grafana/provisioning/alerting/rules.yaml` (format unifié) :

```yaml
apiVersion: 1
groups:
  - orgId: 1
    name: systemes-1m
    folder: Systemes
    interval: 1m
    rules:
      - uid: disque-plein-imminent
        title: Disque plein imminent
        condition: C
        data:
          - refId: A
            relativeTimeRange: {from: 600, to: 0}
            datasourceUid: prometheus
            model:
              expr: >-
                100 * (1 - (node_filesystem_avail_bytes{fstype!~"tmpfs|overlay"}
                / node_filesystem_size_bytes{fstype!~"tmpfs|overlay"}))
          - refId: B
            datasourceUid: __expr__
            model: {type: reduce, expression: A, reducer: last}
          - refId: C
            datasourceUid: __expr__
            model: {type: math, expression: '$B > 85'}
        for: 10m
        labels: {severity: warning, equipe: systemes}
        annotations:
          summary: 'Disque à {{ $values.C.Value }}% sur {{ $labels.instance }}'
```

> **Le YAML d'alerting est verbeux** : construisez la règle dans l'UI,
> vérifiez-la avec Preview, puis exportez-la (Alerting → règle → Export) pour
> obtenir le YAML exact à versionner.

---

## 53. Workflow Git pour le provisioning

Structure de dépôt conseillée :

```
grafana-config/
├── provisioning/
│   ├── datasources/datasources.yaml
│   ├── dashboards/dashboards.yaml
│   ├── alerting/contact-points.yaml
│   ├── alerting/policies.yaml
│   └── alerting/rules.yaml
├── dashboards/
│   ├── systemes/supervision-serveurs.json
│   ├── reseau/supervision-reseau.json
│   └── onduleurs/supervision-onduleur.json
├── grafana.ini                # modèle (sans secrets)
└── README.md                  # procédure d'install + conventions
```

Secrets : **jamais en clair dans Git**. Trois options, par ordre de
préférence :

1. **Gestionnaire de secrets** (Vault, Bitwarden, etc.) qui écrit
   `/etc/default/grafana-server` au déploiement.
2. **Variables d'environnement** sur le serveur (`$PROM_BASIC_PASSWORD`),
   renseignées à la main une fois.
3. **SOPS/age** pour chiffrer les YAML dans le dépôt (si vous devez absolument
   versionner les secrets).

Déploiement :

```bash
# Sur le serveur Grafana
cd /etc/grafana
sudo git -C /srv/git/grafana-config pull   # ou rsync/ansible
sudo cp -r /srv/git/grafana-config/provisioning/* /etc/grafana/provisioning/
sudo cp -r /srv/git/grafana-config/dashboards/* /etc/grafana/dashboards/
sudo chown -R root:grafana /etc/grafana/provisioning /etc/grafana/dashboards
sudo systemctl restart grafana-server
```

---

## 54. Organisations, équipes et dossiers

**Organisation** : cloisonnement de haut niveau (dashboards, datasources,
utilisateurs). Cas d'usage : une orga `Production`, une orga `Kiosque`
(écrans muraux, voir section 72), voire une orga par grande entité.

**Équipe (Team)** : groupe d'utilisateurs au sein d'une orga. On donne les
permissions aux **équipes**, pas aux individus.

**Dossier (Folder)** : regroupe des dashboards et porte des permissions.

Schéma conseillé pour un service systèmes :

```
Orga "Main Org."
├── Équipe "Systemes" (rôle Editor)
├── Équipe "Reseau"   (rôle Editor)
├── Équipe "Direction" (rôle Viewer)
├── Dossier "Supervision"      → Systemes: Edit, Reseau: View
├── Dossier "Metier"          → Direction: View, Systemes: Edit
└── Dossier "Brouillons"      → Systemes: Edit (pas de provisioning)
```

Commandes utiles (via l'API) :

```bash
# Créer une équipe
curl -s -X POST -H "Content-Type: application/json" \
  -d '{"name":"Systemes","email":"systemes@mondomaine.fr"}' \
  http://admin:<A_COMPLETER>@localhost:3000/api/teams | python3 -m json.tool
```

---

## 55. Rôles : Viewer, Editor, Admin, Server Admin

