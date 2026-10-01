---
id: collect-261001-rattrapage/rattrapage/prometheus-guide-13
title: "Guide Prometheus — Supervision métrique complète"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["arr", "datacenter", "memory"]
source: docs/RAG/collect-261001-rattrapage/prometheus_guide.md
source_anchor: ""
source_lines: [2728, 2925]
sha256: f0f967fb9c4921463fd7f7eea5d0d1f857ed943b3603ee5e33660806cb6115b8
---

# Guide Prometheus — Supervision métrique complète

1. **Réduire la fenêtre** : `[30d]` → `[7d]`, ou pré-agréger.
2. **Filtrer tôt** : ajouter `{job="..."}` / `{instance="..."}`.
3. **Recording rule** : pré-calculer la requête lourde (section 42).
4. **Augmenter les ressources** : `--query.max-concurrency`, RAM.
5. **Limiter** : `--query.max-samples=50M` (défaut) protège déjà de l'OOM ;
   ne l'augmentez que si vous savez pourquoi.

```promql
# ❌ Lent : 30 jours × toutes les instances × tous les CPU
sum(rate(node_cpu_seconds_total[30d]))

# ✅ Rapide : recording rule pré-calculée
sum(instance:cpu_usage_percent:30s)
```

## 89. Dépannage : disque plein

Symptômes : `No space left on device` dans les logs, WAL qui grossit,
rétention qui ne purge plus (la purge a besoin d'espace pour compacter !).

Plan d'urgence :

```bash
# 1. Vérifier l'espace et ce qui le consomme
df -h /var/lib/prometheus
du -sh /var/lib/prometheus/* | sort -h | tail

# 2. Réduire la rétention TEMPORAIREMENT et recharger (pas de restart)
#    -> passer retention.time de 30d à 7d dans le service systemd, puis :
sudo systemctl daemon-reload && sudo systemctl restart prometheus
#    (le restart est nécessaire ici : les flags ne se rechargent pas via /-/reload)

# 3. Si vraiment bloqué : supprimer manuellement les plus vieux blocs
#    (Prometheus arrêté !) — dernier recours
sudo systemctl stop prometheus
ls -lt /var/lib/prometheus/ | tail -5      # les plus vieux en bas
# rm -rf /var/lib/prometheus/01XXXX...     # bloc précis, JAMAIS wal/ !
sudo systemctl start prometheus
```

⚠️ Ne **jamais** supprimer `wal/` à chaud : perte des ~2 dernières heures
et risque de corruption. Les blocs (`01.../`) sont supprimables à froid.

Prévention : alerte `predict_linear(..._size...)` sur le disque de données
+ dimensionnement section 75 + `retention.size` en garde-fou.

## 90. Dépannage : OOM (mémoire épuisée)

Symptômes : processus tué (`dmesg | grep -i "killed process"`), restart en
boucle, `MemoryMax` atteint sous systemd.

Causes classées par fréquence :

1. **Explosion de cardinalité** : `prometheus_tsdb_head_series` en forte
   hausse → section 81. **Cause n°1, de loin.**
2. **Requêtes trop lourdes** : subqueries sur 30 j, `quantile_over_time`
   sur de longues fenêtres → `--query.max-samples`, recording rules.
3. **Trop de règles** évaluées trop souvent → `interval: 1m` au lieu de 15 s
   sur les groupes lourds.
4. **Rétention énorme + petit disque/RAM** : le head grossit avec la rétention.

Actions immédiates :

```bash
# Limiter la casse : réduire la concurrence des requêtes
--query.max-concurrency=10 --query.max-samples=20000000
# Puis identifier la métrique fautive :
# Status -> TSDB -> top 10, ou :
curl -s http://127.0.0.1:9090/api/v1/status/tsdb | python3 -m json.tool | head -40
```

## 91. Dépannage : mes alertes ne partent pas

Checklist dans l'ordre :

1. **La règle s'évalue-t-elle ?** UI → Rules : état `ok`, pas d'erreur.
   Tester l'expr dans Graph : retourne-t-elle des séries ?
2. **Le `for:` est-il atteint ?** UI → Alerts : `pending` (en attente) vs
   `firing` (envoyée). Une alerte `pending` qui retombe = `for` trop long
   ou expr instable.
3. **Prometheus envoie-t-il à Alertmanager ?**
   ```promql
   rate(prometheus_notifications_sent_total[5m])
   rate(prometheus_notifications_failed_total[5m])   # > 0 = problème d'envoi
   ```
4. **Alertmanager reçoit-il ?** UI `:9093` → Alerts : l'alerte y figure ?
   Sinon : vérifier `alerting.alertmanagers` dans prometheus.yml + firewall :9093.
5. **Le routage matche-t-il ?** `:9093` → Status : voir la config chargée.
   Tester avec `amtool config routes test` :
   ```bash
   amtool config routes test --config.file=/etc/alertmanager/alertmanager.yml \
     severity=critical equipe=infra
   ```
6. **Silence ou inhibition actifs ?** `:9093` → Silences / Inhibitions.
7. **Le receiver fonctionne-t-il ?** Tester le SMTP à la main :
   ```bash
   swaks --to infra@entreprise.lan --server smtp.lan:587 -tls --auth-user alertes@entreprise.lan
   ```

## 92. Dépannage : scrape timeout / cibles lentes

Symptôme : `up=1` en pointillés, `scrape_duration_seconds` proche du timeout,
trous dans les graphiques.

```promql
# Durée de scrape par job (alerte si > 80 % du timeout)
max by (job) (scrape_duration_seconds) > 8   # pour timeout=10s

# Échantillons par scrape (détecte l'exporter qui grossit)
avg by (job) (scrape_samples_scraped)
```

Remèdes :

- Exporter lent (SNMP sur équipement chargé) : `scrape_interval: 60s`,
  `scrape_timeout: 50s` sur ce job uniquement.
- Trop de séries : `sample_limit` + `metric_relabel_configs` (drop).
- Réseau distant : intervalle plus large, ou Prometheus local + fédération.
- Blackbox avec `follow_redirects` sur un site lent : `timeout` du module
  < `scrape_timeout` du job.

## 93. Les 10 erreurs classiques (et comment les éviter)

| # | Erreur | Symptôme | Solution |
|---|---|---|---|
| 1 | Label à haute cardinalité (`user_id`, `path` brut) | RAM qui explose, OOM | Normaliser côté app / `labeldrop` |
| 2 | `rate()` sur une gauge | Valeurs absurdes | `rate`/`increase` = compteurs uniquement |
| 3 | Fenêtre `[1m]` dans les alertes | Alertes yo-yo | `[5m]` minimum, `for: 5m` |
| 4 | Oublier `by (le)` avant `histogram_quantile` | Quantiles faux | Toujours agréger les buckets d'abord |
| 5 | `sum()` sans `by` sur des compteurs multi-labels | Double comptage après agrégation | `sum by (labels utiles)` explicite |
| 6 | `predict_linear` sur fenêtre trop courte | Faux positifs "disque plein" | `[6h]` minimum |
| 7 | Pas de `for:` sur `up == 0` | Page à 3 h pour un micro-blip réseau | `for: 5m` |
| 8 | SNMP en v2c sur réseau non maîtrisé | Community interceptée | SNMPv3 authPriv |
| 9 | Un seul Prometheus sans méta-monitoring | Panne de supervision invisible | Paire HA + check externe |
| 10 | Rétention non dimensionnée | Disque plein un dimanche | Calcul section 75 + alerte disque |

---

---

# PARTIE VIII — Cas pratiques, checklist, pense-bête, glossaire, quiz

---

## 94. Cas pratique 1 : supervision complète d'un serveur Linux

Objectif : en 15 minutes, un serveur supervisé avec alertes pertinentes.

```bash
# Sur le serveur cible : installer node_exporter (section 45)
# Sur Prometheus : ajouter au file_sd (section 13)
```

`/etc/prometheus/file_sd/serveurs.json` :

```json
[
  {"targets": ["srv-web01:9100"],
   "labels": {"datacenter": "dc-paris", "role": "web", "os": "debian12"}}
]
```

Vérifications :

```promql
up{instance="srv-web01:9100"}                                              # == 1 ?
100 * (1 - node_memory_MemAvailable_bytes / node_memory_MemTotal_bytes)    # RAM %
100 * (1 - node_filesystem_avail_bytes{mountpoint="/"} / node_filesystem_size_bytes{mountpoint="/"})  # disque /
```

Alertes appliquées : `InstanceDown`, `DisqueBientotPlein`, `MemoireCritique`,
`ChargeElevee` (section 69). Dashboard Grafana : importer le dashboard
officiel **Node Exporter Full (ID 1860)** puis l'adapter.

## 95. Cas pratique 2 : supervision d'un parc de switchs (SNMP)

Objectif : superviser 12 switchs d'accès + 2 cœurs : trafic, erreurs, ports down.

1. Générer le module SNMP avec les MIBs du constructeur (section 54).
2. Configurer le job `snmp-switch` (section 55).
3. Renseigner les `ifAlias` sur les switchs (description des ports) : c'est ce
   qui rend les graphiques lisibles ("Uplink-Fibre-01" plutôt que "ifIndex 27").

Requêtes du dashboard "Réseau" :

```promql
# Top 5 des ports les plus chargés (tous switchs)
topk(5, sum by (instance, ifDescr) (rate(ifHCInOctets[5m]) + rate(ifHCOutOctets[5m])) * 8 / 1e6)

# Ports en erreur sur 1 h
sum by (instance, ifDescr) (increase(ifInErrors[1h]) + increase(ifOutErrors[1h])) > 0

# Ports down non désactivés (le classique "câble débranché ou équipement éteint")
count by (instance) (ifOperStatus == 2 and ifAdminStatus == 1)
```

Alerte :

