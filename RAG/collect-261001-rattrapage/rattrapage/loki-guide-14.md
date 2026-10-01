---
id: collect-261001-rattrapage/rattrapage/loki-guide-14
title: "Grafana Loki — Le guide complet"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["agent"]
source: docs/RAG/collect-261001-rattrapage/loki_guide.md
source_anchor: ""
source_lines: [3273, 3341]
sha256: b4782dfd948dafdbb30eb64fc85a7e796bbcefd941f992eba16e10df3f9bc0c7
---

# Grafana Loki — Le guide complet

1. Loki n'indexe que les **labels** (léger) et compresse le texte brut
   dans des chunks ; Elasticsearch construit un **index plein texte**
   (index inversé) de chaque mot — plus puissant en recherche, beaucoup
   plus coûteux.
2. L'ensemble des lignes de log partageant **exactement le même jeu de
   labels**. Le nombre de flux actifs détermine la mémoire des ingesters.
3. Parce que chaque valeur unique crée un **flux distinct en mémoire** :
   un million de requêtes = un million de flux → explosion de la
   cardinalité et OOM de l'ingester. On utilise un champ parsé ou la
   structured metadata à la place.
4. Il compacte les fichiers d'index TSDB, **applique la rétention**
   (suppression des chunks/index expirés) et déduplique.
5. Le **nombre de lignes de log par seconde** du job nginx, calculé sur
   une fenêtre glissante de 5 minutes.
6. `|=` filtre sur une **chaîne littérale** (rapide) ; `|~` filtre sur
   une **expression régulière** (plus coûteux). On place `|=` en premier.
7. Il mémorise **l'offset de lecture de chaque fichier** suivi. Le
   supprimer provoque une relecture complète → doublons massifs.
8. Pour ne pas perdre les chunks en mémoire (jusqu'à ~30 min de logs)
   en cas de crash/redémarrage de l'ingester.
9. Risque de `schema mismatch` : l'historique devient illisible ou les
   requêtes renvoient vide. Il faut **ajouter** une nouvelle entrée avec
   un `from:` futur, jamais modifier l'existante.
10. Avec `absent_over_time({job="syslog", host="srvX"}[10m])` dans une
    règle du ruler, ou en comparant
    `count(count_over_time({job="syslog"}[1h]) by (host))` à l'inventaire.

---

## 85. Pour aller plus loin

### Documentation officielle

- Docs Loki : https://grafana.com/docs/loki/
- Référence LogQL : https://grafana.com/docs/loki/latest/query/
- Exemples de configuration : https://github.com/grafana/loki/tree/main/cmd/loki

### À explorer ensuite

| Sujet | Pourquoi |
|---|---|
| **Mode scalable** (`-target=read/write/backend`) | l'étape entre monolithique et microservices |
| **Recording rules** du ruler | pré-calculer les métriques critiques (dashboards instantanés) |
| **Loki + Tempo** (traces) | derived fields `trace_id` : du log à la trace en un clic |
| **LogQL metric queries avancées** | `quantile_over_time`, histogrammes de latence |
| **Alloy** (successeur de Promtail) | l'agent unifié Grafana (logs, métriques, traces) |
| **S3 + lifecycle** | transition des vieux chunks vers stockage froid |
| **Tests de charge** | simuler le pic d'ingestion avant la prod |

### Livres et formations

- *Logging in Action* (Manning) — les fondamentaux de la centralisation
- La documentation Prometheus (les concepts LogQL en sont dérivés)
- Les dashboards communautaires Grafana pour Loki (ID à importer depuis
  grafana.com/dashboards, recherche « Loki »)

### Prochaines étapes pour votre infra

1. Maquette monolithique + MinIO en lab (1 semaine).
2. 3 serveurs pilotes en Promtail (2 semaines d'observation).
3. Seuils d'alerte calibrés, runbooks écrits.
4. Généralisation au parc + dashboards par service.
5. Revue trimestrielle : cardinalité, croissance disque, rétention.

---

*Guide rédigé pour Zelef — centralisation de logs d'infrastructure.*
*Versions : Loki 3.x, Promtail 3.x, Grafana 10/11, Debian 12 / Ubuntu 22.04+.*
*Les valeurs numériques sont des ordres de grandeur : validez sur votre infra.*
