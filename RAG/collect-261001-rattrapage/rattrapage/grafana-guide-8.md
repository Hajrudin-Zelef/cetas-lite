---
id: collect-261001-rattrapage/rattrapage/grafana-guide-8
title: "Guide Grafana — Dashboards, visualisation et alerting"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["datacenter", "memory", "valuation"]
source: docs/RAG/collect-261001-rattrapage/grafana_guide.md
source_anchor: ""
source_lines: [1505, 1683]
sha256: ac172024d8a4354c31b0362dedb84c7fa24c6b2f03f8c281a9eff420fa1be23a
---

# Guide Grafana — Dashboards, visualisation et alerting

| Transformation | Effet | Cas d'usage |
|---|---|---|
| **Merge** | Fusionne plusieurs requêtes en une table | Combiner CPU + RAM + disque |
| **Join by field** | Jointure sur un champ (ex. instance) | Enrichir avec une table métier |
| **Filter by name** | Ne garde que certains champs | Nettoyer une table SQL |
| **Organize fields** | Renomme, réordonne, masque | Tables propres |
| **Calculate field** | Nouveau champ calculé | `utilisation = 100 * (1 - libre/total)` |
| **Group by** | Agrégation | Moyenne par datacenter |
| **Sort by** | Tri | Top N |
| **Limit** | Tronque | Top 10 |
| **Series to rows** | Séries → lignes | Alimenter une table depuis PromQL |
| **Labels to fields** | Labels → colonnes | Table d'inventaire depuis Prometheus |
| **Config from query** | Pilote seuils/unités depuis une requête | Seuils dynamiques par client |

Exemple : table « serveurs les plus chargés » depuis Prometheus :

1. Requête A : `avg by (instance) (rate(node_cpu_seconds_total{mode!="idle"}[5m]))`
2. Requête B : `node_memory_MemTotal_bytes - node_memory_MemAvailable_bytes`
3. Transform **Merge** → **Labels to fields** → **Organize fields**
   (renommer, masquer `Time`) → **Sort by** CPU décroissant → **Limit** 10.

> **Transformations vs requêtes :** si le backend peut le faire (SQL,
> PromQL), faites-le côté backend (plus rapide, moins de données
> transférées). Les transformations servent à la **présentation** et aux
> combinaisons impossibles côté backend.

---

## 40. Annotations (événements superposés aux courbes)

Les **annotations** affichent des événements (déploiements, incidents,
maintenances) comme des lignes verticales sur tous les panels time series.
Corréler « la latence a explosé » avec « déploiement à 14h32 » devient immédiat.

**Annotation manuelle** : `Ctrl+clic` (ou `Cmd+clic`) sur un graphe →
« Add annotation ». Visible par tous les viewers du dashboard.

**Annotation par requête** (recommandée) — exemples :

```promql
# Prometheus : changements détectés via Alertmanager ou un metric maison
ALERTS{alertstate="firing", alertname="DeploiementEnCours"}
```

```logql
# Loki : lignes de log marquant un déploiement
{job="deploy"} |= "DEPLOY OK"
```

```sql
-- PostgreSQL : table des maintenances planifiées
SELECT
  debut AS "time",
  CONCAT(titre, ' — ', description) AS "text",
  'maintenance' AS "tags"
FROM maintenances
WHERE $__timeFilter(debut);
```

Réglages : Dashboard settings → Annotations → New. Choisissez la datasource,
la requête, et mappez les champs `time` / `text` / `tags`.

> **Bonne pratique d'équipe :** imposez une annotation à chaque changement
> en production (via une table SQL alimentée par votre outil de déploiement,
> ou un webhook). Le MTTR chute quand la cause est visible sur la courbe.

---

## 41. Liens et drill-down (data links, dashboard links)

**Data links** (par point de donnée) : clic droit sur un point → ouvrir un
lien construit avec les valeurs du point.

Exemples (Panel → Data links) :

```
# Voir les logs Loki de cette instance au moment du point
http://grafana.mondomaine.fr/explore?orgId=1&left={"datasource":"Loki",
"queries":[{"expr":"{instance=\"${__data.fields.instance}\"}",
"range":{"from":"${__value.time}","to":"${__value.time}"}}]}

# Simple : ouvrir le dashboard détaillé du serveur
/d/UID_DASHBOARD_DETAIL/serveur-detail?var-serveur=${__data.fields.instance}&from=${__from}&to=${__to}
```

Variables disponibles : `${__data.fields.<nom>}`, `${__value.time}`,
`${__value.raw}`, `${__from}`, `${__to}`.

**Dashboard links** (en haut du dashboard) : liens vers des dashboards liés
(même tag, ou URL externe vers wiki/runbook). Ex. : depuis « Vue d'ensemble »,
liens vers « Détail serveur », « Détail réseau », « Logs ».

> Le drill-down bien conçu remplace 10 dashboards : **vue d'ensemble →
> détail par entité → logs/traces**. Trois niveaux suffisent dans 95 % des cas.

---

## 42. Alerting unifié : concepts (règles, instances, états)

Rappel du vocabulaire (section 4), avec le cycle de vie précis :

```
Requêtes (A, B, C...) → Expressions (reduce, math) → Condition (seuil)
        │ évaluée toutes les N minutes (evaluation interval)
        ▼
   ┌─────────┐   condition vraie    ┌─────────┐   pendant 'for'   ┌─────────┐
   │ Normal  │ ──────────────────► │ Pending │ ─────────────────► │ Firing  │
   └─────────┘                     └─────────┘                   └────┬────┘
        ▲                                                            │ résolu
        └────────────────────────────────────────────────────────────┘
```

- **Evaluation group** : les règles sont regroupées ; le groupe définit
  l'intervalle d'évaluation commun (ex. `1m`). Mettez les règles critiques
  dans un groupe à intervalle court, les règles lentes (taux journaliers)
  dans un groupe à `10m` ou `1h`.
- **For** : durée pendant laquelle la condition doit rester vraie avant de
  passer en Firing. `for = 0` (ou vide) = alerte immédiate (à réserver aux
  cas vraiment urgents) ; `for = 5m` évite les faux positifs sur les pics.
- **Labels** : `severity=critical`, `equipe=systemes`… Ils servent au
  **routage** (section 46) et au **regroupement** des notifications.
- **Annotations** : `summary` (titre court) et `description` (détail, avec
  template `{{ $values.B.Value }}`) — c'est ce que reçoit l'astreinte.
- **NoData / Error handling** : que faire si la requête ne renvoie rien
  (`NoData` → Normal / Alerting / KeepLast / Ok) ou échoue (`Error` →
  Alerting / Ok / KeepLast). **À régler consciemment** : le défaut
  (`NoData = NoData`, `Error = Error`) peut laisser une panne silencieuse.

---

## 43. Créer une règle d'alerte — exemple commenté pas à pas

**Alerting → Alert rules → New alert rule.** Exemple : « Disque > 85 % depuis
10 minutes sur un serveur ».

**Étape 1 — Requêtes :**

```
A (Prometheus) :
100 * (1 - (node_filesystem_avail_bytes{fstype!~"tmpfs|overlay",mountpoint!~"/snap.*"}
            / node_filesystem_size_bytes{fstype!~"tmpfs|overlay",mountpoint!~"/snap.*"}))
```

**Étape 2 — Expressions :**

```
B : Reduce — Input A, Function Last, Mode Strict
    → réduit chaque série à sa dernière valeur
C : Math — $B > 85
    → condition : vrai/faux par série (instance × mountpoint)
```

(Alternative : **Classic condition** `WHEN last() OF A IS ABOVE 85` — plus
simple mais moins flexible que Reduce+Math.)

**Étape 3 — Détails :**

| Champ | Valeur |
|---|---|
| Rule name | `Disque plein imminent` |
| Folder | `Systemes` |
| Evaluation group | `systemes-1m` (interval `1m`) |
| For | `10m` |
| Labels | `severity=warning`, `equipe=systemes` |
| Annotations | `summary = "Disque à {{ $values.C.Value }}% sur {{ $labels.instance }} ({{ $labels.mountpoint }})"`, `description = "…"` |
| NoData | `NoData` (créer une alerte NoData visible) |
| Error | `Error` |

**Étape 4 — Aperçu** : le bouton **Preview** montre les instances qui
seraient créées **maintenant** avec les valeurs actuelles. Utilisez-le
systématiquement avant de sauvegarder.

**Étape 5 — Contact point** : laisser la policy par défaut au début, vérifier
la réception, puis router finement (section 46).

---

## 44. Expressions d'alerte (reduce, math, resample)

