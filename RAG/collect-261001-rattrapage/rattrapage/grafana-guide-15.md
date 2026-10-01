---
id: collect-261001-rattrapage/rattrapage/grafana-guide-15
title: "Guide Grafana — Dashboards, visualisation et alerting"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["arr", "distribution", "incident", "memory"]
source: docs/RAG/collect-261001-rattrapage/grafana_guide.md
source_anchor: ""
source_lines: [3053, 3202]
sha256: 555aef85fa6ecd82f5a34601ca3e13b3565f7e13258c67c0072e5ac712dd4fc6
---

# Guide Grafana — Dashboards, visualisation et alerting

**Cas 4 — Légendes illisibles (`{instance="...", job="..."}`)**
Cause : champ Legend vide. Fix : `{{instance}}` ou `{{ifDescr}}` dans Legend.

**Cas 5 — L'alerte sonne pour « All » mais pas pour un serveur précis**
Cause : la variable `$serveur` vaut `All` dans la règle d'alerte — les règles
d'alerte **n'ont pas accès** aux variables de dashboard ! Fix : réécrire la
requête sans variable (ou avec une regex explicite).

**Cas 6 — « Template variables are not supported in alert queries »**
Cause : voir cas 5. C'est une limitation **voulue** : une alerte doit être
déterministe. Fix : dupliquer la logique sans `$var`.

**Cas 7 — Dashboard très lent (> 10 s)**
Causes fréquentes : trop de panels (> 20), requêtes sur 30 j sans downsampling,
`maxLines` Loki trop élevé. Fix : réduire la plage par défaut, activer les
trends Zabbix, paginer les tables, utiliser des enregistrements (recording
rules) côté Prometheus.

**Cas 8 — « User not found » après passage à LDAP**
Cause : `allow_sign_up = false` avec LDAP, ou `search_base_dns` incorrect.
Fix : `allow_sign_up = true` pour l'auto-provisionnement, tester avec
`ldapsearch` depuis le serveur Grafana avant d'accuser Grafana.

**Cas 9 — Webhook d'alerte reçu mais JSON incompréhensible**
Cause : format non documenté côté récepteur. Fix : logger le payload brut
une fois, écrire le parseur, documenter le format dans le wiki (section 45).

**Cas 10 — Après mise à jour, des panels affichent « Panel plugin not found »**
Cause : plugin tiers incompatible avec la nouvelle version. Fix :
`grafana-cli plugins update-all`, vérifier la compatibilité sur la page du
plugin, ou remplacer par un panel natif.

---

## 76. Les 10 erreurs classiques (et comment les éviter)

| # | Erreur | Conséquence | Prévention |
|---|---|---|---|
| 1 | `root_url` non renseignée | Liens d'alerte cassés, OAuth en échec | Section 10 : la renseigner dès l'install |
| 2 | Dashboards non versionnés | Perdus à la première fausse manip | Provisioning + Git (sections 49-53) |
| 3 | `secret_key` perdue/changée | Secrets illisibles après restauration | Coffre + sauvegarde (sections 13, 64) |
| 4 | Alertes avec variables `$var` | Règles refusées ou silencieuses | Pas de variables dans les alertes (cas 5-6) |
| 5 | `for: 0` partout | Fatigue d'alerte, bruit | `for` adapté (section 48) |
| 6 | Datasource en mode Browser | Secrets exposés, CORS | Mode Server par défaut (section 26) |
| 7 | Compte SQL avec droits d'écriture | Risque d'altération des données métier | Lecture seule (sections 23-24) |
| 8 | Pas de sauvegarde testée | Découverte le jour de la panne | Test trimestriel (section 65) |
| 9 | Grafana non supervisé | Panne de l'outil invisible | Métriques internes + sonde externe (section 61) |
| 10 | 50 dashboards copiés-collés | Maintenance infernale | Variables + templates (sections 36-38) |

---

## 77. Cas pratiques commentés (4 scénarios d'exploitation)

### Cas A — Mise en service d'un nouveau site distant

1. Ajouter le site dans Prometheus (scrape) et Loki (labels `site=lyon`).
2. Variables existantes `$site` : le site apparaît **automatiquement** dans
   les dashboards (aucune modification de dashboard !).
3. Dupliquer la règle d'alerte « Disque plein » ? **Non** : la règle existante
   sans filtre de site couvre déjà le nouveau site (vérifier avec Preview).
4. Ajouter une annotation « Mise en service site Lyon » sur le dashboard réseau.
5. Vérifier la réception d'une alerte de test depuis le nouveau site.

### Cas B — Coupure électrique un dimanche à 3h

1. L'alerte `Onduleur sur batterie` (critical) part vers l'astreinte
   (téléphone, pas e-mail).
2. L'astreinte ouvre le dashboard Onduleurs : autonomie restante 22 min,
   charge 45 %.
3. Procédure (panel Text du dashboard) : si autonomie < 15 min → arrêt
   ordonné des VMs non critiques (runbook lié).
4. Le retour secteur est visible sur la State timeline ; l'alerte passe en
   Resolved → notification de résolution.
5. Lundi : revue — pourquoi la coupure ? Le groupe électrogène a-t-il pris
   le relais ? Ajuster les seuils si besoin.

### Cas C — Lenteurs applicatives signalées par les utilisateurs

1. **Explore** : `histogram_quantile(0.95, sum by (le) (rate(http_request_duration_seconds_bucket[5m])))`
   → le p95 a triplé à 10h12.
2. Heatmap des latences : distribution bimodale → piste cache.
3. Annotations : un déploiement à 10h10 → corrélation immédiate.
4. Logs Loki du service à 10h10-10h15 : erreurs de connexion au cache.
5. Rollback du déploiement, annotation « rollback 10h40 », la courbe
   redescend → incident documenté en 20 min sans dashboard jetable.

### Cas D — Audit : « prouvez que la supervision est sauvegardée »

1. Montrer le cron de backup + les 30 derniers dumps (`ls -lh /srv/backups/grafana`).
2. Montrer le dépôt Git du provisioning (historique des changements).
3. **Test de restauration** sur une VM de labo (section 65), chronométré :
   objectif < 1 h pour un Grafana fonctionnel.
4. Compte-rendu signé : ce qui a été restauré, ce qui a été vérifié
   (datasources au vert, alertes évaluées, test de notification).

---

## 78. Pense-bête de poche (commandes, requêtes, raccourcis)

**Services :**

```bash
systemctl status grafana-server
systemctl restart grafana-server
tail -f /var/log/grafana/grafana.log
grafana-server -v
grafana-cli plugins ls
```

**Santé :**

```bash
curl -s http://localhost:3000/api/health | python3 -m json.tool
curl -s -o /dev/null -w "%{http_code}\n" http://prometheus.mondomaine.fr:9090/-/healthy
```

**PromQL express :**

```promql
100 - (avg by (instance) (rate(node_cpu_seconds_total{mode="idle"}[5m])) * 100)
100 * (1 - (node_memory_MemAvailable_bytes / node_memory_MemTotal_bytes))
topk(5, sum by (instance) (rate(node_cpu_seconds_total{mode!="idle"}[5m])))
histogram_quantile(0.95, sum by (le) (rate(http_request_duration_seconds_bucket[5m])))
```

**LogQL express :**

```logql
{job="syslog"} |= "error"
count_over_time({job="syslog"} |= "error" [1m])
{job="app"} | json | level="error"
```

**Raccourcis interface :**

| Touche | Action |
|---|---|
| `Ctrl+K` ou `/` | Recherche globale |
| `Ctrl+S` | Sauvegarder le dashboard (dans l'éditeur) |
| `Ctrl+clic` sur un graphe | Ajouter une annotation |
| `E` | Basculer en édition (dashboard) |
| `Esc` | Quitter l'édition / fermer |

**Checklist « ça ne marche pas » :** backend joignable ? → identifiants ? →
URL exacte ? → logs Grafana → test dans Explore → (alerting) routage/silences.

---

## 79. Glossaire

