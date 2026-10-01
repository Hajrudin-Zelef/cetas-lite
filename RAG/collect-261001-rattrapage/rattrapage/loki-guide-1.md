---
id: collect-261001-rattrapage/rattrapage/loki-guide-1
title: "Grafana Loki — Le guide complet"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/loki_guide.md
source_anchor: ""
source_lines: [1, 203]
sha256: b33a5b8b2e47a9f4989334af733c7d942aac96fc6cda17cca50c222ab1796630
---

# Grafana Loki — Le guide complet

> **Centraliser, indexer et exploiter vos logs comme vos métriques.**
> Guide pratique en français pour mettre Loki en production : de l'installation
> sur Debian/Ubuntu à l'alerting, en passant par Promtail, LogQL, la rétention
> et le dépannage.
>
> Versions couvertes : **Loki 2.9 / 3.x**, **Promtail 2.9 / 3.x**, Grafana 10/11.
> Les commandes ciblent **Debian 12 / Ubuntu 22.04 / 24.04**.
>
> ⚠️ Les valeurs de configuration sont des **ordres de grandeur** : adaptez-les
> à votre volumétrie et validez toujours sur un environnement de test avant la
> production.

---

## Table des matières

| # | Section |
|---|---------|
| 1 | Pourquoi Loki ? |
| 2 | Concepts : le flux de logs (log stream) |
| 3 | Les chunks : l'unité de stockage |
| 4 | Pas d'index plein texte — la différence avec Elasticsearch |
| 5 | Architecture : vue d'ensemble |
| 6 | Le distributeur (distributor) |
| 7 | L'ingester |
| 8 | Le querier et le query-frontend |
| 9 | Le stockage : filesystem vs stockage objet |
| 10 | Installation sur Debian/Ubuntu : le binaire |
| 11 | Installation : service systemd |
| 12 | Installation : Docker |
| 13 | Premier fichier `loki.yaml` (mode monolithique) |
| 14 | Vérifier que Loki démarre |
| 15 | Modes de déploiement : monolithique vs microservices |
| 16 | Promtail : installation du binaire |
| 17 | Promtail : configuration de base |
| 18 | `scrape_configs` et `static_configs` |
| 19 | `__path__` : globs, pièges et bonnes pratiques |
| 20 | Parser logfmt |
| 21 | Parser JSON |
| 22 | Parser regex |
| 23 | Multiline : gérer les stack traces Java et Python |
| 24 | `relabel_configs`, `labeldrop`, `labelkeep` |
| 25 | Labels : bonnes pratiques |
| 26 | La cardinalité : l'ennemi numéro 1 |
| 27 | Exemples concrets : labels autorisés vs interdits |
| 28 | Structured metadata (Loki 2.8+) |
| 29 | LogQL : anatomie d'une requête |
| 30 | Les sélecteurs de flux |
| 31 | Opérateurs de filtrage : `\|=`, `!=`, `\|~`, `!~` |
| 32 | Filtrer sur les labels extraits par les parsers |
| 33 | `rate()` : mesurer un taux d'erreur |
| 34 | `count_over_time()`, `sum`, `avg` et compagnie |
| 35 | Agréger par labels : `sum by (...)` |
| 36 | Requêtes métriques depuis les logs |
| 37 | Détecter l'absence de logs (dead man's switch) |
| 38 | Sous-requêtes et fenêtres temporelles |
| 39 | Optimiser les requêtes LogQL lentes |
| 40 | Alerting : les règles du ruler Loki |
| 41 | Alertmanager : alerter sur trop d'erreurs 500 |
| 42 | Alerting : bonnes pratiques |
| 43 | Rétention : le table manager (méthode historique) |
| 44 | Rétention : le compactor (méthode moderne) |
| 45 | Rétention par tenant |
| 46 | Intégration Grafana : déclarer la datasource |
| 47 | Grafana Explore : explorer les logs |
| 48 | Dashboards logs : exemples de panneaux |
| 49 | Variables de dashboard LogQL |
| 50 | Cas pratique 1 : centraliser le syslog de 20 serveurs |
| 51 | Cas pratique 2 : suivre les erreurs nginx |
| 52 | Cas pratique 3 : détecter les attaques SSH (avec fail2ban) |
| 53 | Cas pratique 4 : logs applicatifs au format JSON |
| 54 | Cas pratique 5 : audit des commandes sudo |
| 55 | Sécurité : authentification et multi-tenant |
| 56 | Sécurité : TLS de bout en bout |
| 57 | Sécurité : reverse proxy nginx devant Loki |
| 58 | Sécurité : durcissement |
| 59 | Supervision de Loki : ses propres métriques |
| 60 | Alertes sur la santé de Loki |
| 61 | Sauvegarde : chunks sur filesystem |
| 62 | Sauvegarde : stockage objet (S3) |
| 63 | Restauration après sinistre |
| 64 | Mise à jour de Loki |
| 65 | Mise à jour de Promtail |
| 66 | Migration : changer de schéma de stockage |
| 67 | Dépannage : les logs n'arrivent pas |
| 68 | Dépannage : requêtes trop lentes |
| 69 | Dépannage : disque plein |
| 70 | 10 cas de dépannage concrets |
| 71 | Erreur 1 : `too many outstanding requests` |
| 72 | Erreur 2 : `entry out of order` |
| 73 | Erreur 3 : `per-stream rate limit exceeded` |
| 74 | Erreur 4 : `label value too long` / `max label value length` |
| 75 | Erreur 5 : `schema mismatch` / erreurs de schéma |
| 76 | Erreur 6 : HTTP 429 `too many requests` |
| 77 | Erreur 7 : explosion de cardinalité |
| 78 | Erreur 8 : `query timeout` |
| 79 | Erreur 9 : le compactor ne tourne pas |
| 80 | Erreur 10 : Promtail ne lit pas les fichiers |
| 81 | Checklist de mise en production |
| 82 | Pense-bête de poche LogQL |
| 83 | Glossaire |
| 84 | Quiz : 10 questions + réponses |
| 85 | Pour aller plus loin |

---

## 1. Pourquoi Loki ?

Grafana Loki est un système d'agrégation de logs **inspiré de Prometheus** :
au lieu d'indexer le contenu complet de chaque ligne de log (comme
Elasticsearch), il indexe uniquement un petit ensemble de **labels**
(métadonnées clé=valeur) et stocke les lignes brutes compressées dans des
**chunks**.

Conséquences pratiques :

| Aspect | Loki | Elasticsearch |
|---|---|---|
| Index | labels uniquement (léger) | plein texte (lourd) |
| Requêtes | LogQL (proche de PromQL) | Query DSL / KQL |
| Coût de stockage | faible (chunks compressés) | élevé (index inversé) |
| Recherche plein texte | filtrée après sélection des flux | native et rapide |
| Écosystème | Grafana, Prometheus, Alertmanager | Kibana, Logstash |

**Quand choisir Loki :**
- vous utilisez déjà Prometheus + Grafana (même philosophie, même langage) ;
- vous voulez un coût d'infrastructure maîtrisé ;
- vos requêtes partent d'un périmètre (serveur, application) avant de chercher
  du texte.

**Quand éviter Loki :**
- besoin de recherche plein texte massive et instantanée sur tout le corpus ;
- analytique complexe sur le contenu des logs (dans ce cas, un SIEM ou
  Elasticsearch reste pertinent).

---

## 2. Concepts : le flux de logs (log stream)

Dans Loki, un **flux (stream)** = un ensemble unique de labels + les lignes
de log qui lui sont associées.

Exemple : toutes les lignes suivantes appartiennent au **même flux** car
elles partagent exactement les mêmes labels :

```
{job="nginx", host="web01", level="error"}
```

```
{job="nginx", host="web01", level="error"} 2026-09-26T10:00:01Z GET /api 500
{job="nginx", host="web01", level="error"} 2026-09-26T10:00:02Z GET /api 500
```

En revanche, ceci constitue un **autre flux** (le label `host` diffère) :

```
{job="nginx", host="web02", level="error"}
```

Règles fondamentales :

1. **Le nombre de flux actifs = le nombre de combinaisons uniques de labels.**
   C'est lui qui détermine la charge mémoire des ingesters.
2. Les labels doivent être **stables et peu nombreux** : `job`, `host`,
   `app`, `environment`, `level`.
3. Tout ce qui varie à chaque ligne (ID de requête, IP client, timestamp,
   message d'erreur complet) **ne doit jamais être un label**.

Retenez cette phrase, elle reviendra partout dans ce guide :

> **Les labels décrivent *d'où vient* le log. Le contenu du log décrit
> *ce qui s'est passé*. Ne mélangez jamais les deux.**

---

## 3. Les chunks : l'unité de stockage

Loki regroupe les lignes d'un flux en **chunks** (morceaux) compressés :

- un chunk contient les lignes d'**un seul flux** ;
- taille cible : ~1,5 Mo compressé (paramètre `chunk_target_size`) ;
- un chunk est d'abord conservé en mémoire dans l'ingester, puis **flushé**
  vers le stockage (filesystem ou objet) quand il est plein ou trop vieux
  (`chunk_idle_period`, `max_chunk_age`) ;
- les chunks sont **immuables** une fois flushés : on ne les modifie plus,
  on en crée de nouveaux.

Cycle de vie d'un chunk :

```
lignes de log → ingester (mémoire) → chunk plein/vieux → flush → stockage
                                                              → index (labels)
```

L'**index** (BoltDB, TSDB) ne contient que : *quel chunk contient quel flux
pour quelle période*. C'est pour cela que Loki est si économe : l'index est
minuscule comparé à un index plein texte.

---

