---
id: collect-261001-rattrapage/rattrapage/prometheus-guide-6
title: "Guide Prometheus — Supervision métrique complète"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["memory"]
source: docs/RAG/collect-261001-rattrapage/prometheus_guide.md
source_anchor: ""
source_lines: [1093, 1329]
sha256: 57690ed80c732df8d76fc9fe74cbcbae8ba264ad4293248f2356c73d899f251d
---

# Guide Prometheus — Supervision métrique complète

1. **Vérifier que la métrique existe** : dans Graph, taper le début du nom
   (autocomplétion) ou `count({__name__=~"node_.*"})`.
2. **Simplifier** : partir de la métrique brute, ajouter les filtres un par un.
3. **Vérifier les labels** : cliquer une série → voir ses labels réels
   (souvent la jointure échoue à cause d'un label inattendu).
4. **Expliquer la lenteur** : préfixer par rien de spécial — utiliser l'API
   avec `stats=all` ou le flag `--query.stats-enabled` :

```bash
curl -s 'http://localhost:9090/api/v1/query?query=up&stats=all' | head -c 2000
```

Réflexes :

- Résultat vide ? → vérifier les noms de labels (`instance` vs `hostname`),
  les regex, le `on()`/`ignoring()`.
- Trop lent ? → réduire la fenêtre, ajouter un filtre `job`, créer une
  recording rule.
- Valeurs absurdes sur `rate()` ? → fenêtre trop courte (< 4× scrape_interval).

---

---

# PARTIE III — Les exporters

---

## 45. node_exporter : l'exporter système de référence

`node_exporter` expose ~1500 métriques matérielles/OS sur `:9100/metrics`.
Installation (même méthode que Prometheus : binaire + systemd) :

```bash
NODE_VERSION="1.9.0"
cd /tmp
wget "https://github.com/prometheus/node_exporter/releases/download/v${NODE_VERSION}/node_exporter-${NODE_VERSION}.linux-amd64.tar.gz"
tar xzf "node_exporter-${NODE_VERSION}.linux-amd64.tar.gz"
sudo cp "node_exporter-${NODE_VERSION}.linux-amd64/node_exporter" /usr/local/bin/
sudo useradd --no-create-home --shell /usr/sbin/nologin node_exporter
```

`/etc/systemd/system/node_exporter.service` :

```ini
[Unit]
Description=Node Exporter
Wants=network-online.target
After=network-online.target

[Service]
User=node_exporter
Group=node_exporter
Type=simple
Restart=always
ExecStart=/usr/local/bin/node_exporter \
  --collector.systemd \
  --collector.processes \
  --collector.tcpstat \
  --no-collector.arp \
  --no-collector.bcache \
  --no-collector.infiniband \
  --web.listen-address=127.0.0.1:9100

[Install]
WantedBy=multi-user.target
```

> Activez uniquement les collecteurs utiles : chaque collecteur = des séries
> en plus. `--collector.disable-defaults` + activation sélective = le plus propre
> sur un grand parc.

## 46. node_exporter : CPU et charge

Métriques clés :

| Métrique | Type | Usage |
|---|---|---|
| `node_cpu_seconds_total{cpu,mode}` | counter | % CPU par mode |
| `node_load1`, `node_load5`, `node_load15` | gauge | charge système |
| `node_procs_running`, `node_procs_blocked` | gauge | processus |
| `node_context_switches_total` | counter | changements de contexte |
| `node_forks_total` | counter | forks |

```promql
# Utilisation CPU % (tous modes sauf idle), par instance
100 * (1 - avg by (instance) (rate(node_cpu_seconds_total{mode="idle"}[5m])))

# Détail par mode (user/system/iowait...) — dashboard uniquement
100 * avg by (instance, mode) (rate(node_cpu_seconds_total[5m]))

# Charge vs nombre de CPU : alerte si load5 > 2× CPU pendant 15 min
node_load5 > 2 * count by (instance) (node_cpu_seconds_total{mode="idle"})

# iowait élevé = disques ou stockage en souffrance
100 * avg by (instance) (rate(node_cpu_seconds_total{mode="iowait"}[5m])) > 20
```

⚠️ `node_load1` seul ne veut rien dire sans le comparer au nombre de CPU.

## 47. node_exporter : mémoire et swap

| Métrique | Usage |
|---|---|
| `node_memory_MemTotal_bytes` | RAM totale |
| `node_memory_MemAvailable_bytes` | RAM réellement disponible (le bon indicateur) |
| `node_memory_SwapTotal_bytes`, `node_memory_SwapFree_bytes` | swap |
| `node_memory_Cached_bytes`, `node_memory_Buffers_bytes` | cache |

```promql
# % mémoire utilisée (méthode fiable)
100 * (1 - node_memory_MemAvailable_bytes / node_memory_MemTotal_bytes)

# Swap utilisé en Mio
(node_memory_SwapTotal_bytes - node_memory_SwapFree_bytes) / 1024 / 1024

# Pression mémoire : disponible < 10 % pendant 10 min → alerte
node_memory_MemAvailable_bytes / node_memory_MemTotal_bytes < 0.1
```

⚠️ Ne jamais utiliser `MemFree` seul : Linux utilise la RAM libre comme cache,
c'est normal. `MemAvailable` est l'indicateur pertinent (noyau ≥ 3.14).

## 48. node_exporter : disques et filesystems

| Métrique | Usage |
|---|---|
| `node_filesystem_size_bytes{device,mountpoint,fstype}` | taille |
| `node_filesystem_avail_bytes` | espace dispo (utilisateur) |
| `node_filesystem_files`, `node_filesystem_files_free` | inodes |
| `node_filesystem_readonly` | 1 si monté en lecture seule |
| `node_disk_io_time_seconds_total{device}` | temps d'occupation disque |
| `node_disk_read_bytes_total`, `node_disk_written_bytes_total` | débits |

```promql
# % utilisé par filesystem (exclure les pseudo-fs)
100 * (1 - node_filesystem_avail_bytes{fstype!~"tmpfs|overlay|squashfs|vfat"}
          / node_filesystem_size_bytes)

# Inodes : le piège classique (% espace OK mais inodes pleins !)
100 * (1 - node_filesystem_files_free / node_filesystem_files) > 80

# Saturation disque : % d'utilisation du disque (0-100)
100 * rate(node_disk_io_time_seconds_total[5m])

# Disque qui va saturer dans moins de 3 jours
predict_linear(node_filesystem_avail_bytes{fstype="ext4"}[6h], 3*24*3600) < 0
```

> **Toujours** superviser les inodes en plus de l'espace : des millions de
> petits fichiers (sessions PHP, mails) remplissent les inodes bien avant l'espace.

## 49. node_exporter : réseau

| Métrique | Usage |
|---|---|
| `node_network_receive_bytes_total{device}` | octets reçus |
| `node_network_transmit_bytes_total{device}` | octets émis |
| `node_network_receive_errors_total` | erreurs RX |
| `node_network_receive_drop_total` | paquets jetés |
| `node_network_carrier` | 1 si lien UP |

```promql
# Débit en Mbit/s par interface (hors loopback/virtuelles)
sum by (instance, device) (rate(node_network_receive_bytes_total{device!~"lo|veth.*|docker.*"}[5m])) * 8 / 1e6

# Erreurs ou drops : doit rester à 0
increase(node_network_receive_errors_total[1h]) > 0
  or increase(node_network_transmit_errors_total[1h]) > 0

# Lien tombé
node_network_carrier{device!~"lo.*"} == 0
```

## 50. node_exporter : systemd et processus

Avec `--collector.systemd` :

```promql
# Service en échec
node_systemd_unit_state{name="nginx.service", state="failed"} == 1

# Tous les services failed, par machine
count by (instance) (node_systemd_unit_state{state="failed"} == 1)

# Service désactivé mais attendu actif (exemple)
node_systemd_unit_state{name="postgresql.service", state="active"} == 0
```

Avec `--collector.processes` : `namedprocess_namegroup_*` (nécessite config).
Alternative légère : `node_procs_running` pour détecter un fork-bomb.

## 51. blackbox_exporter : principe et installation

Le **blackbox_exporter** sonde des cibles **depuis** Prometheus : HTTP(S),
ICMP (ping), TCP, DNS. Pattern : Prometheus scrape le blackbox en lui passant
la cible en paramètre.

Installation :

```bash
BB_VERSION="0.26.0"
cd /tmp
wget "https://github.com/prometheus/blackbox_exporter/releases/download/v${BB_VERSION}/blackbox_exporter-${BB_VERSION}.linux-amd64.tar.gz"
tar xzf "blackbox_exporter-${BB_VERSION}.linux-amd64.tar.gz"
sudo cp "blackbox_exporter-${BB_VERSION}.linux-amd64/blackbox_exporter" /usr/local/bin/
sudo mkdir -p /etc/blackbox_exporter
sudo useradd --no-create-home --shell /usr/sbin/nologin blackbox_exporter
```

`/etc/blackbox_exporter/config.yml` — modules de sondes :

```yaml
modules:
  http_2xx:                      # sonde HTTP(S) : code 2xx attendu
    prober: http
    timeout: 10s
    http:
      valid_status_codes: [200, 301, 302]
      method: GET
      follow_redirects: true
      preferred_ip_protocol: "ip4"

  http_post_api:                 # sonde d'API avec POST + Basic Auth
    prober: http
    timeout: 10s
    http:
      method: POST
      headers:
        Content-Type: application/json
      body: '{"ping": true}'
      basic_auth:
        username: "monitor"
        password: "changeme"
      fail_if_body_not_matches_regexp:
        - '"status":"ok"'

