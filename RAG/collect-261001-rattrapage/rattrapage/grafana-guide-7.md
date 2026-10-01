---
id: collect-261001-rattrapage/rattrapage/grafana-guide-7
title: "Guide Grafana — Dashboards, visualisation et alerting"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["datacenter", "valuation"]
source: docs/RAG/collect-261001-rattrapage/grafana_guide.md
source_anchor: ""
source_lines: [1299, 1504]
sha256: e4815a488f6f2fb3d1135b40a07ad42e800dc6af485932b510cbc7d42ec7ec59
---

# Guide Grafana — Dashboards, visualisation et alerting

```markdown
## Dashboard Supervision Serveurs
**Question :** mes serveurs sont-ils sains ?
**Mainteneur :** Équipe Systèmes — <A_COMPLETER>
**Runbook :** https://wiki.mondomaine.fr/runbooks/serveurs
```

**Dashboard list** : affiche les dashboards d'un dossier ou taggés — utile
pour une page d'accueil / un sommaire par équipe.

**Alert list** : liste les alertes en cours (firing/pending), filtrable par
dossier ou dashboard. Parfait en haut d'un dashboard « tour de contrôle ».

---

## 34. Options de panel : unités, min/max, seuils, couleurs

Ces réglages (onglet **Field** / **Standard options**) s'appliquent à presque
tous les panels : c'est 50 % de la lisibilité.

- **Unit** : `percent (0-100)` pour un %, `bytes (IEC)` pour du disque/RAM
  (affiche Go automatiquement), `s` / `ms` pour les durées, `ops` / `reqps`
  pour les débits. **Un graphe sans unité est un graphe ambigu.**
- **Min / Max** : fixez-les quand la borne a un sens (0-100 % pour un taux).
  Ne fixez pas de max arbitraire sur une métrique sans borne (trafic réseau).
- **Decimals** : 1 ou 2 suffisent. `99.87 %`, pas `99.8712345678 %`.
- **Thresholds** : mode `Absolute` (valeurs) ou `Percentage`. La **couleur de
  base** = l'état normal (souvent vert). Ajoutez les seuils dans l'ordre
  croissant.
- **Color scheme** : `Single color` (sobre, recommandé), `From thresholds`
  (la couleur suit les seuils), `Classic palette` (une couleur par série).

Convention d'équipe conseillée (à écrire dans votre wiki) :

| Couleur | Sens |
|---|---|
| Vert | OK / normal |
| Orange | Dégradé / warning (agir sous 4 h) |
| Rouge | Critique (agir maintenant) |
| Bleu | Information (pas d'action) |
| Gris | Pas de données / désactivé |

---

## 35. Overrides de champs (field overrides)

Les **overrides** permettent d'appliquer des réglages différents à **certaines
séries** d'un même panel, sans le dupliquer.

Exemple : sur un graphe « trafic réseau par interface », afficher `eth0` en
ligne épaisse et masquer les interfaces `veth*` :

1. Panel → onglet **Overrides** → **Add field override**.
2. **Fields with name matching regex** : `/^veth/` → **Hide in → Tooltip,
   Viz, Legend** (ou filtrez déjà dans la requête, c'est mieux).
3. **Fields with name** : `eth0` → **Graph styles → Line width = 3**.

Autres usages :

- Une série en **barres** au milieu de courbes (ex. déploiements) :
  override → **Graph styles → Draw mode = Bars**.
- Unités différentes par série : override → **Standard options → Unit**.
  (Mieux : deux axes Y via **Axis → Placement = Right**.)
- Renommer proprement : override → **Standard options → Display name**.

> **Priorité :** les réglages panel s'appliquent d'abord, puis les overrides
> dans l'ordre de la liste. En cas de conflit, le **dernier** override gagne.

---

## 36. Variables et templates — le principe

Une **variable** rend un dashboard réutilisable : au lieu de coder en dur
`instance="srv-web-03"`, on écrit `instance="$serveur"` et l'utilisateur
choisit le serveur dans un menu déroulant en haut du dashboard.

Cycle de vie d'une variable :

1. **Dashboard settings → Variables → New** : nom (`serveur`), type, requête.
2. Utilisation dans les requêtes : `$serveur` ou `${serveur}` (forme longue
   obligatoire si le nom est suivi de texte : `${serveur}_suffixe`).
3. Dans les titres de panels : `$serveur` fonctionne aussi.

Exemple minimal (Prometheus) :

- Nom : `serveur`
- Type : **Query**
- Data source : `Prometheus`
- Query : `label_values(node_uname_info, instance)`
- Refresh : **On Dashboard Load**
- Multi-value : oui, Include All option : oui

Requête du panel :

```promql
100 - (avg by (instance) (rate(node_cpu_seconds_total{mode="idle",instance=~"$serveur"}[5m])) * 100)
```

Notez `=~` (regex match) au lieu de `=` : indispensable quand la variable
peut valoir plusieurs instances ou `.*` (All).

> **Règle :** dès qu'un dashboard est décliné plus de 2 fois pour des cibles
> différentes, il doit utiliser des variables. Le copier-coller de dashboards
> est une dette qui se paie en maintenance.

---

## 37. Variables : les 7 types en détail

| Type | Usage | Exemple |
|---|---|---|
| **Query** | Liste dynamique depuis une datasource | `label_values(up, job)` |
| **Custom** | Liste fixe saisie à la main | `prod,staging,lab` |
| **Interval** | Pas de temps (`$pas`) | `1m,5m,15m,1h` |
| **Data source** | Choisir la datasource | Type `prometheus` → `Prom-Prod`, `Prom-Lab` |
| **Constant** | Valeur invisible figée | `seuil_critique = 90` |
| **Textbox** | Saisie libre | Rechercher un hostname |
| **Ad hoc filters** | Filtres ajoutés à la volée | (Prometheus/Loki/ES uniquement) |

Requêtes utiles par datasource (type Query) :

```promql
# Prometheus : valeurs d'un label, éventuellement filtrées par une autre variable
label_values(node_uname_info{datacenter="$dc"}, instance)
# Tous les jobs
label_values(up, job)
```

```logql
# Loki : valeurs d'un label
label_values({job="syslog"}, instance)
```

```sql
-- MySQL/PostgreSQL : liste depuis une table métier
SELECT DISTINCT datacenter FROM sites ORDER BY 1;
```

```yaml
# Zabbix (plugin) : groupes d'hôtes, hôtes, applications…
# Dans le champ Query du plugin : choisir "Groupes", "Hôtes", etc.
```

Options importantes :

- **Refresh** : `On Dashboard Load` (défaut raisonnable) ; `On Time Range
  Change` si la liste dépend de la plage.
- **Sort** : alphabétique ou numérique selon le cas.
- **Include All option** : ajoute `All` (valeur `.*` en regex). Custom All
  value possible (ex. `.+` pour exclure le vide).
- **Hide** : `Variable` pour masquer une variable technique (ex. constante
  de seuil) tout en l'utilisant.

---

## 38. Chaînage de variables, regex et multi-sélection

**Chaînage** : la requête d'une variable peut référencer une autre variable.
Ordre de définition = ordre d'évaluation (Grafana trie automatiquement les
dépendances, mais gardez un ordre logique : `dc` → `serveur`).

```
Variable dc      : label_values(node_uname_info, datacenter)
Variable serveur : label_values(node_uname_info{datacenter="$dc"}, instance)
```

Quand `$dc` change, la liste des serveurs se rafraîchit. C'est le pattern
« drill-down » : datacenter → cluster → serveur → interface.

**Regex** (champ Regex de la variable) : filtrer/transformer les valeurs.

```
# Ne garder que les vrais serveurs (pas les conteneurs éphémères)
/^(srv|db|web)-.*/

# Extraire le hostname sans le port
/([^:]+):.*/
```

**Multi-value + All** : quand l'utilisateur sélectionne plusieurs valeurs,
Grafana les injecte comme une regex `(a|b|c)`. D'où l'obligation d'utiliser
`=~` dans les requêtes PromQL/LogQL :

```promql
# CORRECT avec variable multi
node_cpu_seconds_total{instance=~"$serveur"}
# FAUX (ne marchera qu'avec une seule valeur)
node_cpu_seconds_total{instance="$serveur"}
```

**Répétition de panels/rows** (Repeat options) : répéter un panel ou une row
**pour chaque valeur** de la variable. Ex. : une row « Serveur » répétée pour
chaque `$serveur` sélectionné → un dashboard qui affiche N serveurs en
parallèle sans duplication manuelle.

> **Limite :** ne répétez pas plus de ~12 panels par dashboard ; au-delà, la
> page devient lourde et les requêtes explosent. Pour de gros parcs, préférez
> le drill-down (un serveur à la fois) + un dashboard « vue d'ensemble ».

---

## 39. Transformations (merge, join, filter, calculate field…)

Les **transformations** (onglet Transform) modifient les données **après**
requête, sans toucher au backend. Les plus utiles :

