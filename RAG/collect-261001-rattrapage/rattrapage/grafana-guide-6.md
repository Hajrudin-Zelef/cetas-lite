---
id: collect-261001-rattrapage/rattrapage/grafana-guide-6
title: "Guide Grafana — Dashboards, visualisation et alerting"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["distribution"]
source: docs/RAG/collect-261001-rattrapage/grafana_guide.md
source_anchor: ""
source_lines: [1088, 1298]
sha256: d6b158ac307f071dbf183310f1f3de697362c8aa19a4f66ce71204c8b4a7546c
---

# Guide Grafana — Dashboards, visualisation et alerting

- [ ] Mode d'accès `Server` (sauf besoin explicite du mode Browser, documenté).
- [ ] **Save & test** au vert, depuis le serveur Grafana (pas depuis votre PC).
- [ ] `Scrape interval` / `timeInterval` cohérent avec la réalité du backend.
- [ ] Compte dédié, droits minimaux (lecture seule partout où c'est possible).
- [ ] Secrets chiffrés (provisioning : via variables d'environnement ou
      gestionnaire de secrets, jamais en clair dans Git — section 53).
- [ ] Timeout adapté (60 s par défaut ; 120 s pour les requêtes SQL lourdes).
- [ ] Une seule datasource **par défaut** ; nommez les autres explicitement
      (`Prom-Prod`, `Prom-Lab`, `Loki-Prod`).
- [ ] **Health check** : le panel « datasource qui ne répond pas » (section 73)
      décrit quoi faire quand le test échoue.

Diagnostic réseau depuis le serveur Grafana (à faire **avant** d'accuser Grafana) :

```bash
# Le backend est-il joignable depuis le serveur Grafana ?
curl -s -o /dev/null -w "%{http_code}\n" http://prometheus.mondomaine.fr:9090/-/healthy
curl -s http://loki.mondomaine.fr:3100/ready
nc -vz bdd.mondomaine.fr 5432

# Résolution DNS vue par Grafana
getent hosts prometheus.mondomaine.fr
```

---

## 27. Panels : time series

Le panel **Time series** est le cheval de labour : 80 % de vos graphes.

Réglages qui font la différence :

- **Style → Line width / Fill opacity** : `1` et `10` pour des courbes
  lisibles ; `Fill opacity 0` + `Line width 2` pour comparer peu de séries.
- **Connect null values → Never** : un trou dans les données doit se **voir**
  (un collecteur en panne qui affiche une belle courbe continue, c'est un
  mensonge).
- **Show points → Never** (sauf séries clairsemées).
- **Gradient mode → Scheme** : joli mais trompeur sur les comparaisons ;
  réservez-le aux dashboards de présentation.
- **Stacking → Normal** : pour des composants qui s'additionnent
  (CPU user + system + iowait). **Jamais** de stacking sur des pourcentages
  indépendants.
- **Tooltip mode → All** : voir toutes les séries au survol.

Exemple : CPU par mode, stacké, avec légende propre :

```promql
sum by (mode) (rate(node_cpu_seconds_total{instance="$serveur"}[$__rate_interval]))
```

- Legend : `{{mode}}`
- Unit : `percent (0-100)` → non, ici ce sont des secondes/s : laissez
  `short`, ou convertissez en % avec la requête de la section 17.
- Seuils : `80` orange, `95` rouge (section 34).

---

## 28. Panels : gauge (jauge)

La **jauge** répond à « où en est-on par rapport à une limite ? ».
Un seul nombre, un arc de cercle, des seuils colorés.

Cas d'usage : % disque utilisé, température salle, charge onduleur,
niveau d'un bac, progression d'un quota.

Réglages :

- **Show threshold labels / markers** : coché (on voit où sont les seuils).
- **Orientation → Auto**, **Show unfilled area** : coché.
- **Value → Calculation → Last** (la valeur actuelle).
- **Min/Max** : renseignez-les **toujours** (une jauge sans max ment :
  75 % de quoi ?).
- **Thresholds** : ex. disque : `70` orange, `90` rouge. **Base color** verte.

Exemple (disque le plus plein du serveur sélectionné) :

```promql
100 * max(1 - (node_filesystem_avail_bytes{fstype!~"tmpfs|overlay",instance="$serveur"}
              / node_filesystem_size_bytes{fstype!~"tmpfs|overlay",instance="$serveur"}))
```

> **Anti-pattern :** 12 jauges identiques pour 12 serveurs. Utilisez une
> **variable** `$serveur` (section 36) ou un panel **Stat** répété
> (option Repeat, section 38).

---

## 29. Panels : stat (grandes valeurs)

Le panel **Stat** affiche un grand chiffre + un libellé + une mini-tendance
(sparkline). Idéal pour les indicateurs « d'un coup d'œil ».

Bandeau type en haut d'un dashboard serveurs :

| Panel | Requête | Calculation | Couleur |
|---|---|---|---|
| Serveurs UP | `count(up == 1)` | Last | verte si > 0 |
| Alertes firing | `sum(ALERTS{alertstate="firing"})` | Last | rouge si > 0 |
| CPU moyen | `100 - (avg(rate(node_cpu_seconds_total{mode="idle"}[5m]))*100)` | Last | seuils 70/90 |
| Dispo 30 j | `100 * avg_over_time(up[30d])` | Last | seuils 99/99.9 |

Réglages :

- **Graph mode → Area** (sparkline derrière le chiffre).
- **Color mode → Background** pour les indicateurs critiques (toute la tuile
  devient rouge : visible à 5 mètres sur un écran mural).
- **Text mode → Value and name**.
- **Unit** : toujours renseignée (`percent`, `bytes`, `s`, `none`…).

---

## 30. Panels : table

Le panel **Table** affiche des données tabulaires, avec tri, filtre et
mise en forme par colonne.

Cas d'usage : inventaire (serveur, OS, version), top des filesystems,
tickets ouverts, dernières alertes.

Exemple : filesystems > 70 % d'utilisation :

```promql
100 * (1 - (node_filesystem_avail_bytes{fstype!~"tmpfs|overlay"}
            / node_filesystem_size_bytes{fstype!~"tmpfs|overlay"})) > 70
```

Mise en forme :

- **Column Settings** : renommez `Value` → `Utilisation %`, alignez à droite.
- **Cell options → Cell type → Colored background** + seuils 80/90 : la
  colonne devient un indicateur visuel.
- **Cell type → Data links** : lien vers le dashboard détaillé du serveur
  (`/d/<uid>/detail-serveur?var-serveur=${__data.fields.instance}`).
- Désactivez les colonnes parasites (`Time`) via **Organize fields**
  (transformation, section 39).

> **Table ou Time series ?** Si la question est « combien / qui / quoi
> maintenant » → Table. Si c'est « comment ça évolue » → Time series.

---

## 31. Panels : heatmap

La **heatmap** montre la densité d'une distribution dans le temps : chaque
colonne = un intervalle, chaque ligne = un « bucket », la couleur = le nombre
d'observations. Parfait pour les **latences** et les répartitions.

Source idéale : un histogramme Prometheus.

```promql
sum by (le) (rate(http_request_duration_seconds_bucket[5m]))
```

Réglages :

- **Calculate from histogram → Yes** (Grafana reconstruit les buckets depuis
  les séries `_bucket`).
- **Color scheme → Oranges/Reds** (ou `Spectrum`).
- **Y axis → Unit** : `s` (secondes).
- **Tooltip → Show histogram** : coché.

Lecture : une tache chaude qui monte = la latence se dégrade ; deux taches
distinctes = **distribution bimodale** (ex. cache hit vs cache miss) —
information qu'une simple moyenne ne montre jamais.

Autre usage : heatmap des connexions par heure (requête SQL, section 24) ou
des erreurs par code et par heure.

---

## 32. Panels : bar chart, pie chart, state timeline, status history

**Bar chart** (barres) : comparer des catégories à un instant T.

```promql
# Top 5 filesystems par utilisation
topk(5, 100 * (1 - (node_filesystem_avail_bytes / node_filesystem_size_bytes)))
```
- Orientation → Horizontal, tri décroissant : lecture immédiate.

**Pie chart** (camembert) : répartition d'un tout. À utiliser avec parcimonie
(5 parts max, sinon illisible). Ex. : répartition des OS du parc, des
tickets par statut.

**State timeline** : états discrets dans le temps (up/down, open/closed).
Idéal pour « l'historique de disponibilité des services » :

```promql
up{job="blackbox"}
```
- **Merge equal consecutive values** : coché. Couleurs : vert = 1, rouge = 0.

**Status history** : comme State timeline mais en blocs annotés (ex. statut
des jobs de sauvegarde par nuit : OK / WARNING / FAILED).

> **Règle :** un panel = un type de question. Ne mélangez pas des gauges, des
> courbes et des camemberts pour dire la même chose : choisissez la forme la
> plus directe.

---

## 33. Panels : text, dashboard list, alert list

**Text** (Markdown/HTML) : documentation **dans** le dashboard. Utilisez-le
pour :

- expliquer ce que montre le dashboard et qui le maintient ;
- documenter la procédure (« si cette alerte est rouge, faire… ») ;
- afficher les liens utiles (wiki, runbook).

