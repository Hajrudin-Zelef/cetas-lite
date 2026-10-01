---
id: collect-261001-rattrapage/rattrapage/loki-guide-5
title: "Grafana Loki — Le guide complet"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["agent", "memory"]
source: docs/RAG/collect-261001-rattrapage/loki_guide.md
source_anchor: ""
source_lines: [1009, 1277]
sha256: 15f1e3c1ab90a78b4c6e514cc20c4c0323aa7a44e85767218b91bea476ed53fb
---

# Grafana Loki — Le guide complet

Configuration :

```yaml
    pipeline_stages:
      - multiline:
          firstline: '^\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}'
          max_wait_time: 3s
          max_lines: 128
```

- `firstline` : regex qui détecte le **début** d'un nouveau log.
  Tout ce qui ne matche pas est rattaché au log précédent.
- `max_wait_time` : au bout de 3 s sans nouvelle ligne, le bloc est flushé.
- `max_lines` : sécurité contre les blocs infinis.

Variante Java (stack traces indentées) :

```yaml
      - multiline:
          firstline: '^\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2},\d{3}'
          max_wait_time: 3s
```

> ⚠️ `firstline` doit être **très fiable** : si elle matche au milieu d'une
> trace, le bloc est coupé en deux. Testez sur des exemples réels.

---

## 24. `relabel_configs`, `labeldrop`, `labelkeep`

Comme Prometheus, Promtail permet de **réécrire les labels** avant envoi.

```yaml
scrape_configs:
  - job_name: kubernetes-pods
    kubernetes_sd_configs:
      - role: pod
    relabel_configs:
      # Ne garder que le namespace 'prod'
      - source_labels: [__meta_kubernetes_namespace]
        regex: prod
        action: keep
      # Créer un label 'app' depuis une annotation
      - source_labels: [__meta_kubernetes_pod_annotation_app]
        target_label: app
      # Supprimer les labels temporaires __meta_*
      - regex: __meta_.*
        action: labeldrop
    pipeline_stages:
      # Supprimer un label devenu inutile
      - labeldrop:
          - filename
      # Ou ne garder qu'une liste blanche (radical)
      # - labelkeep:
      #     - job
      #     - host
      #     - level
```

Actions `relabel_configs` utiles :

| Action | Effet |
|---|---|
| `keep` / `drop` | filtrer les cibles selon une regex |
| `replace` | (défaut) créer/modifier un label |
| `labeldrop` / `labelkeep` | supprimer/garder par regex sur le **nom** |
| `hashmod` | sharding (rarement utile pour Loki) |

---

## 25. Labels : bonnes pratiques

Le choix des labels est **la décision la plus importante** d'un
déploiement Loki. Règles d'or :

1. **Peu de labels** : 3 à 6 par flux suffisent largement.
   Exemples : `job`, `host`, `env`, `app`, `level`, `unit`.
2. **Valeurs stables et bornées** : la liste des valeurs possibles doit
   être connue et petite (dizaines, centaines — pas millions).
3. **Jamais de données à forte cardinalité** : pas d'IP client, pas d'ID
   de requête, pas de timestamp, pas de message d'erreur complet, pas
   d'URL complète.
4. **Cohérence des noms** : `host` partout (pas `hostname` ici et
   `server` là). Documentez votre convention.
5. **Préférez les labels aux filtres texte** pour ce que vous interrogez
   souvent : `{level="error"}` est infiniment plus rapide que `|= "error"`.

Convention proposée pour une PME / un service systèmes :

| Label | Valeurs exemples |
|---|---|
| `job` | `syslog`, `nginx`, `app-paiement`, `systemd-journal` |
| `host` | `srv01`, `web02` (nom court, stable) |
| `env` | `prod`, `preprod`, `lab` |
| `level` | `debug`, `info`, `warn`, `error` |
| `unit` | `nginx.service`, `sshd.service` (journald) |

---

## 26. La cardinalité : l'ennemi numéro 1

La **cardinalité** d'un label = le nombre de valeurs distinctes qu'il peut
prendre. Le nombre de flux actifs ≈ produit des cardinalités.

Exemple catastrophe :

| Label | Cardinalité |
|---|---|
| `job` | 20 |
| `host` | 20 |
| `request_id` | 1 000 000 / jour |
| **Flux potentiels** | **400 000 000** |

Chaque flux actif consomme de la mémoire dans l'ingester (quelques Ko).
400 millions de flux = **OOM kill garanti** de l'ingester, ingestion
bloquée, trous dans les logs.

**Comment détecter** : métrique `loki_ingester_memory_streams` en
croissance continue, ou requêtes `logql` qui timeout sur
`sum by (...)` avec trop de séries.

**Comment corriger** :
1. Identifier le label coupable : `topk(10, count by (label) (...))` ne
   marche pas directement — utilisez plutôt l'API
   `/loki/api/v1/label/{nom}/values` et comptez.
2. Le **rétrograder** : transformer le label en champ parsé à la requête
   (`| json | request_id="abc"` au lieu d'un label).
3. Ou utiliser la **structured metadata** (section 28) pour les champs à
   cardinalité moyenne.

> 🚨 **Règle simple** : si un label peut dépasser ~quelques milliers de
> valeurs distinctes, ce n'est pas un label. Point.

Limites de garde-fou dans `limits_config` :

```yaml
limits_config:
  max_streams_per_user: 10000      # rejette les nouveaux flux au-delà
  max_global_streams_per_user: 0   # 0 = désactivé ; à régler en cluster
  max_label_names_per_series: 15   # pas plus de 15 labels par flux
  max_label_name_length: 1024
  max_label_value_length: 2048
```

---

## 27. Exemples concrets : labels autorisés vs interdits

### ✅ Labels sains

```yaml
labels:
  job: nginx
  host: web01
  env: prod
  level: error
```
→ 4 labels, cardinalités faibles, flux stables et peu nombreux.

### ❌ Labels dangereux

```yaml
labels:
  job: nginx
  host: web01
  client_ip: 203.0.113.45      # ❌ cardinalité énorme (tous les visiteurs)
  request_id: a3f9c1d2         # ❌ unique par requête = un flux par requête !
  url: /api/users/12345        # ❌ quasi-unique par appel
  user_agent: "Mozilla/5.0..." # ❌ des milliers de variantes
```

### 🔧 Correction : parser au lieu de labelliser

```yaml
    pipeline_stages:
      - regex:
          expression: '^(?P<ip>\S+) .* "(?P<method>\S+) (?P<path>\S+).*?" (?P<status>\d{3})'
      # On garde UNIQUEMENT method et status en labels (faible cardinalité)
      - labels:
          method:
          status:
      # ip et path restent des champs : filtrables en LogQL avec | ip="..."
```

Requête équivalente, sans explosion de cardinalité :

```logql
{job="nginx"} | regexp `^(?P<ip>\S+)` | ip="203.0.113.45"
```

---

## 28. Structured metadata (Loki 2.8+)

Entre les labels (indexés, faible cardinalité) et le texte brut, il existe
un juste milieu : la **structured metadata**. Ce sont des paires clé=valeur
attachées au log, **non indexées**, mais interrogeables en LogQL.

Cas d'usage typique : `request_id`, `trace_id`, `user_id` — cardinalité
trop forte pour des labels, mais besoin de filtrer dessus ponctuellement.

```yaml
    pipeline_stages:
      - json:
          expressions:
            request_id: request_id
            trace_id: trace_id
      - structured_metadata:
          request_id:
          trace_id:
```

Requête :

```logql
{job="api"} | request_id="a3f9c1d2"
```

> ⚠️ La structured metadata n'est **pas indexée** : le filtre s'applique
> après sélection des flux, comme `|=`. À utiliser avec un sélecteur de
> flux déjà sélectif. Elle augmente la taille des chunks (compter ~+10-20 %).

---

## 29. LogQL : anatomie d'une requête

LogQL a deux types de requêtes, comme PromQL :

**1. Requête de logs** (retourne des lignes) :

```logql
{job="nginx"} |= "error" | json | status >= 500
```

Pipeline de gauche à droite :
1. `{job="nginx"}` — sélecteur de flux (index, rapide) ;
2. `|= "error"` — filtre texte : garde les lignes contenant "error" ;
3. `| json` — parse le JSON, extrait les champs ;
4. `| status >= 500` — filtre sur le champ parsé.

**2. Requête métrique** (retourne un nombre par instant) :

```logql
sum(rate({job="nginx"} |= "error" [5m]))
```

Elle enveloppe une requête de logs dans une **fonction d'agrégation**
sur une fenêtre `[5m]`. C'est ce type qui sert pour les dashboards et
les alertes.

---

## 30. Les sélecteurs de flux

Le sélecteur `{...}` choisit les flux via les labels indexés. Opérateurs :

| Opérateur | Sens | Exemple |
|---|---|---|
| `=` | égalité exacte | `{job="nginx"}` |
| `!=` | différent | `{job!="debug"}` |
| `=~` | regex (RE2) | `{host=~"web0[1-3]"}` |
| `!~` | regex négative | `{env!~"dev\|lab"}` |

Exemples :

```logql
# Tous les logs nginx de web01
{job="nginx", host="web01"}

