---
id: collect-261001-rattrapage/rattrapage/loki-guide-13
title: "Grafana Loki — Le guide complet"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["agent", "apache", "arr", "latency"]
source: docs/RAG/collect-261001-rattrapage/loki_guide.md
source_anchor: ""
source_lines: [3072, 3272]
sha256: 8a498d3aac3bfc39aadb8ecab2d5ca6b330580dc36f6ea8d1aaf1e243e6d7d23
---

# Grafana Loki — Le guide complet

> 💡 Une requête qui timeout **toujours** même optimisée signale un
> problème de dimensionnement ou de cardinalité, pas un problème de timeout.

---

## 79. Erreur 9 : le compactor ne tourne pas

**Symptômes** : la rétention ne supprime rien (section 70, cas 8),
l'index TSDB grossit indéfiniment, ou les requêtes sur de vieilles
données restent lentes (pas de compaction).

**Vérifications** :
```bash
# Le compactor est-il actif ?
curl -s http://localhost:3100/metrics | grep -E "loki_boltdb_shipper_compactor_running|loki_compactor"
# Dernière compaction réussie ?
curl -s http://localhost:3100/metrics | grep loki_compactor_last_successful_run
```

**Causes fréquentes** :
1. `retention_enabled: true` oublié dans le bloc `compactor:` → la
   rétention ne s'applique jamais (compaction seule OK).
2. En mode microservices : le déploiement `backend` (ou le `-target`
   incluant le compactor) n'est pas lancé.
3. `working_directory` non accessible en écriture.
4. Conflit de leadership en cluster (ring KV) : un seul compactor actif
   à la fois, c'est normal — mais il en faut **un**.

---

## 80. Erreur 10 : Promtail ne lit pas les fichiers

**Symptômes** : `promtail_targets_active == 0`, aucun log du job,
aucune erreur réseau.

**Checklist** (dans l'ordre) :
1. **Le glob matche-t-il ?**
   ```bash
   sudo -u promtail ls -l /var/log/nginx/*.log
   # → si "permission denied" ou "no such file", c'est ici que ça coince
   ```
2. **Permissions** : l'utilisateur `promtail` est-il dans les bons
   groupes ? (`adm`, `systemd-journal` — section 16)
3. **systemd durci** : `ProtectSystem=strict` sans `ReadOnlyPaths`
   → Promtail ne voit rien (section 16).
4. **positions.yaml** : corrompu ? (JSON invalide après crash disque)
   → renommer le fichier et redémarrer (relecture depuis le début :
   doublons temporaires, à n'utiliser qu'en dernier recours).
5. **Debug** : passer le log en debug temporairement :
   ```yaml
   server:
     log_level: debug
   ```
   puis `journalctl -u promtail -f | grep -i "tail\|target"`.

---

## 81. Checklist de mise en production

### Avant le go-live

- [ ] `loki.yaml` versionné (Git) et sauvegardé
- [ ] Stockage objet (S3/MinIO) configuré, ou disque dimensionné (section 69)
- [ ] WAL de l'ingester activé (section 7)
- [ ] Rétention configurée **et** compactor vérifié (sections 44, 79)
- [ ] `auth_enabled: true` + tenants + reverse proxy TLS (sections 55-57)
- [ ] Labels relus par un pair : pas de cardinalité dangereuse (sections 25-27)
- [ ] `max_streams_per_user` et rate limits fixés (sections 6, 26)
- [ ] Query-frontend + cache activés (section 8)
- [ ] Règles d'alerte Loki déployées et **testées** (sections 40-41)
- [ ] Supervision de Loki dans Prometheus + alertes santé (sections 59-60)
- [ ] Sauvegarde testée **avec restauration** (sections 61-63)
- [ ] Runbook : « Loki down » / « disque plein » / « ingestion arrêtée »
- [ ] Promtail déployé sur tous les serveurs (Ansible), `positions.yaml` persisté

### Les 30 premiers jours

- [ ] Revoir les seuils d'alerte (bruit vs signal)
- [ ] Vérifier la croissance disque réelle vs dimensionnement
- [ ] Auditer les labels : nouveaux jobs = nouvelle revue cardinalité
- [ ] Tester une restauration complète sur un environnement isolé
- [ ] Documenter les dashboards et requêtes utiles dans le wiki d'équipe

---

## 82. Pense-bête de poche LogQL

### Sélecteurs

```logql
{job="nginx"}                    # égalité
{job="nginx", host="web01"}      # ET implicite
{job=~"nginx|apache"}            # regex
{env!="prod"}                    # négation
{host=~"web0[12]"}               # regex partielle
```

### Filtres de ligne

```logql
|= "error"        # contient
!= "healthcheck"  # ne contient pas
|~ "5\d\d"        # regex
!~ "timeout"      # regex négative
```

### Parsers

```logql
| json                       # parse JSON
| logfmt                     # parse logfmt
| regexp `(?P<s>\d{3})`       # regex nommée
| pattern `<_> Failed <w>`   # pattern simplifié
| line_format `{{.msg}}`     # reformate la ligne affichée
| label_format host=`{{.h}}` # (re)nomme un label
```

### Filtres sur champs parsés

```logql
| json | level="error" | latency_ms > 1000
| logfmt | status=~"5.." | path!="/health"
```

### Fonctions métriques

```logql
rate({job="nginx"}[5m])                          # lignes/s
count_over_time({job="nginx"}[1h])               # lignes sur 1h
bytes_rate({job="nginx"}[5m])                    # octets/s
sum by (host) (rate({job="nginx"}[5m]))           # par host
topk(5, sum by (host) (count_over_time({job="nginx"}[1h])))
avg_over_time({j} | json | unwrap latency_ms [5m])
quantile_over_time(0.99, {j} | json | unwrap latency_ms [5m])
absent_over_time({job="syslog", host="srv01"}[10m])
```

### Recettes express

```logql
# Taux d'erreur 5xx nginx
sum(rate({job="nginx"} |~ " 5\d\d " [5m]))

# Brute-force SSH par minute
sum(count_over_time({job="sshd"} |= "Failed password" [1m]))

# Top IP attaquantes (structured metadata 'ip')
topk(20, sum by (ip) (count_over_time({job="sshd"} |= "Failed password" [24h])))

# Serveurs silencieux (pas de logs depuis 10 min)
# → absent_over_time + ruler (section 37)
```

---

## 83. Glossaire

| Terme | Définition |
|---|---|
| **Chunk** | bloc de logs compressé d'un seul flux ; unité de stockage |
| **Distributeur** | composant qui reçoit les pushs et route vers les ingesters |
| **Flux (stream)** | ensemble de lignes partageant exactement les mêmes labels |
| **Ingester** | composant qui écrit les chunks en mémoire puis les flushe |
| **Label** | paire clé=valeur indexée décrivant l'origine d'un log |
| **LogQL** | langage de requête de Loki (inspiré de PromQL) |
| **Cardinalité** | nombre de valeurs distinctes d'un label ; à garder faible |
| **Compactor** | compacte l'index TSDB et applique la rétention |
| **Ruler** | évalue les règles d'alerte/enregistrement sur les logs |
| **Querier** | exécute les requêtes sur ingesters + stockage |
| **Query-frontend** | cache, découpe et parallélise les requêtes |
| **Promtail** | agent qui lit les fichiers de log et les pousse vers Loki |
| **Pipeline stages** | étapes de transformation des logs dans Promtail |
| **Structured metadata** | paires clé=valeur non indexées, filtrables en LogQL |
| **TSDB** | format d'index moderne de Loki (remplace BoltDB) |
| **WAL** | journal write-ahead de l'ingester (survit aux redémarrages) |
| **Tenant** | espace de noms isolé quand `auth_enabled: true` |
| **Rétention** | durée de conservation des logs avant suppression |
| **rate()** | fonction LogQL : lignes par seconde sur une fenêtre |
| **unwrap** | extrait une valeur numérique pour les agrégations |
| **Derived fields** | champs cliquables dans Grafana (ex : trace_id → Tempo) |

---

## 84. Quiz : 10 questions + réponses

### Questions

1. Quelle est la différence fondamentale d'indexation entre Loki et
   Elasticsearch ?
2. Qu'est-ce qu'un flux (stream) dans Loki ?
3. Pourquoi ne faut-il jamais mettre un `request_id` en label ?
4. Que fait le compactor ?
5. Que signifie `rate({job="nginx"}[5m])` ?
6. Quelle est la différence entre `|=` et `|~` ?
7. À quoi sert le fichier `positions.yaml` de Promtail ?
8. Pourquoi faut-il activer le WAL de l'ingester en production ?
9. Que se passe-t-il si on modifie une entrée `schema_config` existante ?
10. Comment détecter qu'un serveur n'envoie plus de logs ?

### Réponses

