---
id: collect-261001-rattrapage/rattrapage/loki-guide-11
title: "Grafana Loki — Le guide complet"
domain: rattrapage
role: reference
task: reference
actors: []
dates: ["2024-01-01", "2026-10-01"]
keywords: ["agent", "aws"]
source: docs/RAG/collect-261001-rattrapage/loki_guide.md
source_anchor: ""
source_lines: [2613, 2834]
sha256: 7f888c13e8203ea1cdb56911b97c94745251aab3a4de20d1a4716333e72c63ba
---

# Grafana Loki — Le guide complet

## 65. Mise à jour de Promtail

Promtail est **stateless** (tout est dans `positions.yaml`) : la mise à
jour est triviale, mais à grande échelle il faut de la méthode.

```bash
# Sur un serveur : même procédure que Loki
sudo systemctl stop promtail
sudo install -m 0755 promtail-linux-amd64 /usr/local/bin/promtail
sudo systemctl start promtail
sudo journalctl -u promtail --since "5 min ago" --no-pager
```

À 20+ serveurs, automatisez (Ansible, section 50) par **vagues** :

1. vague 1 : 2 serveurs pilotes → 24 h d'observation ;
2. vague 2 : 50 % du parc ;
3. vague 3 : le reste.

> ⚠️ Ne supprimez jamais `positions.yaml` pendant la mise à jour :
> Promtail relirait tous les fichiers depuis le début → **doublons
> massifs** dans Loki (et pic d'ingestion).

---

## 66. Migration : changer de schéma de stockage

Exemple : passer de `boltdb-shipper` + filesystem à `tsdb` + S3.
**On ne modifie jamais une entrée de schéma existante** : on ajoute une
nouvelle entrée avec une date `from:` future.

```yaml
schema_config:
  configs:
    # Ancien schéma : inchangé, garde l'historique lisible
    - from: 2024-01-01
      store: boltdb-shipper
      object_store: filesystem
      schema: v12
      index:
        prefix: index_
        period: 24h
    # Nouveau schéma : actif à partir du 1er octobre 2026
    - from: 2026-10-01
      store: tsdb
      object_store: s3
      schema: v13
      index:
        prefix: index_
        period: 24h

storage_config:
  boltdb_shipper:
    active_index_directory: /var/lib/loki/bolt-index
    cache_location: /var/lib/loki/bolt-cache
  tsdb_shipper:
    active_index_directory: /var/lib/loki/tsdb-index
    cache_location: /var/lib/loki/tsdb-cache
  aws:
    s3: s3://KEY:SECRET@minio.local:9000/loki-chunks
    s3forcepathstyle: true
```

Déroulement :

1. Ajouter l'entrée, redémarrer Loki **avant** la date `from`.
2. Au jour J, les nouveaux chunks partent sur S3 en TSDB.
3. L'ancien index reste requêtable jusqu'à expiration de la rétention.
4. Quand la rétention a purgé l'ancien schéma, retirer son entrée
   (et son `storage_config` associé).

---

## 67. Dépannage : les logs n'arrivent pas

Arbre de décision, dans l'ordre :

```
1. Promtail tourne-t-il ?
   └─ systemctl status promtail / journalctl -u promtail
2. Promtail lit-il les fichiers ?
   └─ curl localhost:9080/metrics | grep promtail_targets_active
   └─ permissions ? (groupes adm/systemd-journal, section 16)
3. Promtail arrive-t-il à pousser ?
   └─ journalctl -u promtail | grep -i "error sending"
   └─ métrique promtail_dropped_bytes_total > 0 ?
4. Loki reçoit-il ?
   └─ rate(loki_distributor_lines_received_total[5m])
   └─ journalctl -u loki | grep -i error
5. La requête est-elle bonne ?
   └─ tester {job="..."} seul, sans filtre, sur 1h
   └─ vérifier les labels réellement présents : /loki/api/v1/labels
```

Commandes clés :

```bash
# Labels connus de Loki (le job existe-t-il vraiment ?)
curl -s http://localhost:3100/loki/api/v1/labels

# Valeurs d'un label
curl -s http://localhost:3100/loki/api/v1/label/job/values

# Dernières erreurs Promtail
sudo journalctl -u promtail -p err --since "30 min ago" --no-pager
```

---

## 68. Dépannage : requêtes trop lentes

| Symptôme | Cause probable | Remède |
|---|---|---|
| Timeout après 60 s | sélecteur trop large | resserrer labels + plage (section 39) |
| Lent seulement sur vieilles données | chunks froids sur objet | normal ; prévoir le cache |
| Lent même sur 1 h | trop de flux (cardinalité) | revoir les labels (section 26) |
| Grafana timeout, API OK | `proxy_read_timeout` trop court | 300 s côté nginx (section 57) |
| Tout est lent d'un coup | querier saturé / disque lent | `max_query_parallelism`, IOPS |

Diagnostic express :

```bash
# Temps de réponse d'une requête simple (doit être < 5 s)
time curl -s 'http://localhost:3100/loki/api/v1/query_range?query={job="syslog"}&limit=10' -o /dev/null

# Requêtes en cours et leur durée (métriques)
curl -s http://localhost:3100/metrics | grep -E "loki_request_duration_seconds_(sum|count)"
```

> 💡 Activez le **query-frontend** avec cache (section 8/13) : la 2ᵉ
> exécution d'un dashboard est souvent 10× plus rapide que la 1ʳᵉ.

---

## 69. Dépannage : disque plein

Le disque plein **bloque l'ingestion** (l'ingester ne peut plus flusher
ni écrire le WAL). Réaction en 4 temps :

1. **Urgence** : libérer de l'espace
   ```bash
   df -h /var/lib/loki
   du -sh /var/lib/loki/*          # qui grossit ? chunks ? wal ? index ?
   # Purger les vieux chunks MANUELLEMENT est risqué : préférez baisser
   # temporairement la rétention (le compactor nettoiera)
   ```
2. **Baisser la rétention** temporairement (`retention_period: 168h`),
   redémarrer, laisser le compactor travailler.
3. **Chercher la cause** : pic d'ingestion ? (`bytes_rate` par job),
   rétention trop longue vs disque, compactor en panne (section 79) ?
4. **Prévenir** : alerte Prometheus sur `node_filesystem_avail_bytes`,
   dimensionnement (compter ~1,5× la rétention en marge).

Dimensionnement indicatif :

```
espace_disque ≈ volume_journalier × jours_rétention × 1,5 (marge + WAL + index)
```

Exemple : 50 Go/jour × 31 jours × 1,5 ≈ **2,4 To**.

---

## 70. 10 cas de dépannage concrets

### Cas 1 — Promtail envoie, Loki rejette : `entry out of order`

**Symptôme** : `journalctl -u promtail` montre des 400 `entry out of order`.
**Cause** : deux sources poussent le même flux avec des timestamps qui se
croisent (ex : deux Promtail lisent le même fichier NFS).
**Remède** : un seul agent par fichier ; ou `max_out_of_order_time`
côté Loki (défaut 0 — à n'augmenter qu'en connaissance de cause).

### Cas 2 — Doublons après redémarrage Promtail

**Symptôme** : lignes en double sur quelques minutes après chaque restart.
**Cause** : `positions.yaml` supprimé ou illisible → relecture depuis le début.
**Remède** : ne jamais effacer `positions.yaml` ; vérifier ses permissions
(`promtail:promtail`, `0644`).

### Cas 3 — `too many outstanding requests` côté client

**Symptôme** : Promtail logue `too many outstanding requests`, des logs
sont droppés.
**Cause** : Loki (ou le réseau) n'absorbe pas le débit ; file d'envoi pleine.
**Remède** : augmenter `batchsize`, vérifier le réseau, scaler Loki ;
temporairement `limits_config.ingestion_rate_mb` plus haut. Détail section 71.

### Cas 4 — Requêtes vides alors que les logs existent

**Symptôme** : `{job="nginx"}` ne retourne rien sur les dernières 24 h.
**Cause fréquente** : le `timestamp` parsé par Promtail est faux (mauvais
`format:`) → les logs sont ingérés… en 1970 ou en 2030.
**Remède** : vérifier avec une requête sans filtre sur une large plage,
corriger le `format:` du stage `timestamp`.

### Cas 5 — Grafana : « Data source error »

**Symptôme** : Explore affiche une erreur de datasource.
**Causes** : URL fausse, Loki down, header `X-Scope-OrgID` manquant
(si `auth_enabled: true`), timeout du proxy.
**Remède** : « Save & Test » de la datasource, `curl` direct sur Loki,
vérifier le header dans la config de la datasource.

### Cas 6 — L'ingester est OOM-killed

**Symptôme** : `dmesg | grep -i oom` montre loki tué, trous dans les logs.
**Cause** : explosion du nombre de flux (cardinalité) ou `chunk_target_size`
trop gros × trop de flux.
**Remède** : trouver le label coupable (section 26), fixer
`max_streams_per_user`, augmenter la RAM ou réduire les labels.

### Cas 7 — Le ruler n'envoie aucune alerte

**Symptôme** : la règle existe mais rien n'arrive dans Alertmanager.
**Checklist** : `alertmanager_url` correct ? `enable_api: true` ?
règles dans `/var/lib/loki/rules/<tenant>/` ? `curl
http://localhost:3100/ruler/rule_groups` les liste-t-il ? L'expression
retourne-t-elle quelque chose dans Explore ?

### Cas 8 — Rétention ne supprime rien

