---
id: collect-261001-rattrapage/rattrapage/prometheus-guide-7
title: "Guide Prometheus — Supervision métrique complète"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/prometheus_guide.md
source_anchor: ""
source_lines: [1330, 1588]
sha256: c3d1e0fd9d6d5b7cf2bd057ad24b20f8b7527564d6e9f78e99245b16856d3cc2
---

# Guide Prometheus — Supervision métrique complète

  icmp_ping:                     # ping ICMP (nécessite capabilities)
    prober: icmp
    timeout: 10s
    icmp:
      preferred_ip_protocol: "ip4"

  tcp_connect:                   # simple ouverture TCP (ex. : port 443, 25, 3306)
    prober: tcp
    timeout: 10s

  dns_resolve:                   # résolution DNS
    prober: dns
    dns:
      query_name: "intranet.lan"
      query_type: "A"
      valid_rcodes: ["NOERROR"]
```

Service systemd (note : ICMP nécessite `CAP_NET_RAW`) :

```ini
[Service]
User=blackbox_exporter
ExecStart=/usr/local/bin/blackbox_exporter \
  --config.file=/etc/blackbox_exporter/config.yml \
  --web.listen-address=127.0.0.1:9115
AmbientCapabilities=CAP_NET_RAW
```

## 52. blackbox_exporter : configuration Prometheus

Le pattern clé : `relabel_configs` réécrit `__address__` vers le blackbox et
passe la vraie cible en paramètre `__param_target`.

```yaml
  - job_name: 'blackbox-http'
    metrics_path: /probe
    params:
      module: [http_2xx]          # module par défaut
    file_sd_configs:
      - files: ['/etc/prometheus/file_sd/blackbox.json']
    relabel_configs:
      - source_labels: [__address__]
        target_label: __param_target     # la vraie cible -> ?target=
      - source_labels: [__param_target]
        target_label: instance          # garder le nom lisible
      - target_label: __address__       # ...et scraper le blackbox lui-même
        replacement: '127.0.0.1:9115'

  - job_name: 'blackbox-icmp'
    metrics_path: /probe
    params:
      module: [icmp_ping]
    static_configs:
      - targets:
          - '192.168.10.1'      # passerelle
          - '192.168.10.20'     # onduleur
          - '8.8.8.8'           # contrôle sortie internet
    relabel_configs:
      - source_labels: [__address__]
        target_label: __param_target
      - source_labels: [__param_target]
        target_label: instance
      - target_label: __address__
        replacement: '127.0.0.1:9115'
```

`/etc/prometheus/file_sd/blackbox.json` :

```json
[
  {"targets": ["https://intranet.lan", "https://webmail.lan"],
   "labels": {"module": "http_2xx", "equipe": "infra"}},
  {"targets": ["smtp.lan:25", "imap.lan:993"],
   "labels": {"module": "tcp_connect", "equipe": "infra"}}
]
```

> Pour choisir le module par cible via le fichier SD, ajouter un relabeling
> `source_labels: [module] → target_label: __param_module`.

## 53. blackbox_exporter : métriques et alertes

Métriques principales (préfixe `probe_`) :

| Métrique | Signification |
|---|---|
| `probe_success` | 1 = sonde OK, 0 = échec |
| `probe_duration_seconds` | durée totale de la sonde |
| `probe_http_status_code` | code HTTP retourné |
| `probe_http_duration_seconds{phase}` | détail par phase (connect, tls, processing…) |
| `probe_ssl_earliest_cert_expiry` | timestamp d'expiration du certificat |
| `probe_icmp_duration_seconds` | latence ping |
| `probe_dns_lookup_time_seconds` | temps de résolution DNS |

```promql
# Sonde en échec depuis plus de 5 min
probe_success == 0

# Latence HTTP p95 par cible
histogram_quantile(0.95, sum by (instance, le) (rate(probe_http_duration_seconds_bucket[5m])))

# Certificat qui expire dans moins de 21 jours
(probe_ssl_earliest_cert_expiry - time()) / 86400 < 21

# Détail : où part le temps ? (DNS ? TLS ? serveur ?)
avg by (instance, phase) (probe_http_duration_seconds)
```

> `probe_ssl_earliest_cert_expiry` : **l'alerte certificat** la plus fiable qui
> soit — oubliez les scripts maison.

## 54. snmp_exporter : principe et générateur

Le **snmp_exporter** interroge les équipements SNMP (switchs, routeurs,
onduleurs, imprimantes) et expose le résultat à Prometheus.

Particularité : la configuration (`snmp.yml`) est **générée** par le
`generator` à partir d'un `generator.yml` qui liste les MIBs à parcourir.
On ne l'écrit pas à la main.

```bash
# Générateur : cloner le repo et construire le snmp.yml
git clone https://github.com/prometheus/snmp_exporter.git
cd snmp_exporter/generator
# 1. Placer vos MIBs dans mibs/
# 2. Décrire les modules dans generator.yml :
```

`generator.yml` (extrait) :

```yaml
modules:
  switch_standard:          # module "générique" pour switchs
    walk: [sysUpTime, interfaces, ifXTable, ipNetToMediaTable]
    lookups:
      - source_labels: [ifIndex]
        target_label: ifDescr
    overrides:
      ifDescr:
        ignore: true        # on utilise ifAlias à la place si renseigné

  ups_eaton:                # module onduleur (MIB PowerMIB / XUPS-MIB)
    walk:
      - 1.3.6.1.4.1.534.1  # XUPS-MIB : tout l'arbre Eaton
    version: 2
    auth:
      community: public
```

```bash
go run generator.go generate   # produit snmp.yml
sudo cp snmp.yml /etc/snmp_exporter/snmp.yml
```

> Sans les MIBs du constructeur, le générateur ne peut pas nommer les OID :
> téléchargez-les (Eaton XUPS-MIB, APC PowerNet-MIB, MIBs constructeur du switch)
> et placez-les dans `mibs/`.

## 55. snmp_exporter : supervision des switchs

`prometheus.yml` — pattern identique au blackbox (paramètre `target`) :

```yaml
  - job_name: 'snmp-switch'
    metrics_path: /snmp
    params:
      module: [switch_standard]
      # auth: [v3-user]          # si SNMPv3 (voir section 56)
    static_configs:
      - targets:
          - '192.168.10.2'      # SW-CORE-01
          - '192.168.10.3'      # SW-ACC-01
    relabel_configs:
      - source_labels: [__address__]
        target_label: __param_target
      - source_labels: [__param_target]
        target_label: instance
      - target_label: __address__
        replacement: '127.0.0.1:9116'   # snmp_exporter
```

Métriques utiles :

```promql
# Trafic par interface (en Mbit/s), avec le nom d'interface si ifAlias renseigné
sum by (instance, ifDescr) (
  rate(ifHCInOctets[5m]) or rate(ifInOctets[5m])
) * 8 / 1e6

# Interfaces en erreur
increase(ifInErrors[1h]) > 0 or increase(ifOutErrors[1h]) > 0

# Interface down (mais pas administrativement désactivée)
ifOperStatus != 1 and ifAdminStatus == 1

# Uptime de l'équipement (détecte un reboot : chute brutale)
sysUpTime
```

⚠️ Compteurs 32 bits (`ifInOctets`) vs 64 bits (`ifHCInOctets`) : sur les liens
rapides, le 32 bits boucle en quelques minutes → `rate()` fausse les calculs.
Toujours préférer `ifHCInOctets`/`ifHCOutOctets` quand l'équipement les expose.

## 56. snmp_exporter : supervision des onduleurs (UPS)

Les onduleurs exposent l'état batterie, la charge, les tensions via SNMP.
Deux grandes familles de MIB :

| Constructeur | MIB | OID racine typique |
|---|---|---|
| Eaton | XUPS-MIB | `1.3.6.1.4.1.534.1` |
| APC / Schneider | PowerNet-MIB | `1.3.6.1.4.1.318.1.1.1` |
| Standard | UPS-MIB (RFC 1628) | `1.3.6.1.2.1.33` |

Métriques typiques après génération (noms normalisés par le générateur) :

```promql
# Charge de l'onduleur en %
xupsOutputLoad            # Eaton
upsAdvOutputLoad          # APC

# Niveau batterie en %
xupsBatteryCapacity       # Eaton
upsAdvBatteryCapacity     # APC

# Tension batterie (V)
xupsBatteryVoltage

# Temps restant estimé (minutes)
xupsBatteryRunTimeRemaining
upsAdvBatteryRunTimeRemaining

# Température batterie (°C)
xupsBatteryTemperature
upsAdvBatteryTemperature
```

Alertes onduleur critiques :

```promql
# Sur batterie (pas sur secteur) — CRITIQUE
xupsInputSource == 2 or upsAdvInputLineStatus == 2

# Batterie faible (< 30 %)
xupsBatteryCapacity < 30 or upsAdvBatteryCapacity < 30

# Batterie à remplacer (drapeau constructeur)
xupsBatteryNeedsReplacement == 1 or upsAdvBatteryReplaceIndicator == 1

# Surcharge (> 90 %)
xupsOutputLoad > 90

# Température batterie anormale (> 35 °C : vieillissement accéléré)
xupsBatteryTemperature > 35
```

Configuration SNMPv3 (recommandée : SNMPv2c = community en clair) :

