---
id: collect-261001-rattrapage/rattrapage/lxc-guide-8
title: "LXC & LXD — Guide ultra-complet"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["agent", "arr", "cost", "datacenter", "memory"]
source: docs/RAG/collect-261001-rattrapage/lxc_guide.md
source_anchor: ""
source_lines: [1957, 2210]
sha256: 93501c7e5c0abbafbbf1722357dec7c5db7878eed2ebbc60b84fa9c3737504e8
---

# LXC & LXD — Guide ultra-complet

```yaml
scrape_configs:
  - job_name: 'lxd'
    static_configs:
      - targets: ['hote-lxd1:8443']
    scheme: https
    tls_config:
      cert_file: /etc/prometheus/lxd-client.crt
      key_file: /etc/prometheus/lxd-client.key
      insecure_skip_verify: true
    metrics_path: /1.0/metrics
```

**Alertes utiles (`alerts.yml`) :**

```yaml
groups:
  - name: lxd
    rules:
      - alert: ConteneurArrete
        expr: lxd_instance_state == 0
        for: 5m
        labels: {severity: critical}
      - alert: MemoireConteneurHaute
        expr: lxd_memory_MemAvailable_bytes / lxd_memory_MemTotal_bytes < 0.1
        for: 10m
        labels: {severity: warning}
```

**Dashboards Grafana :** un panel par conteneur (CPU, RAM, I/O, réseau),
un panel d'ensemble (top 10 consommateurs).

---

## 66. Supervision — Zabbix / Netdata

**Zabbix :** agent dans chaque conteneur (ou supervision SNMP de l'hôte +
`lxc list` via UserParameter) :

```bash
# Sur l'hôte : UserParameter pour Zabbix
# /etc/zabbix/zabbix_agentd.d/lxd.conf
UserParameter=lxd.containers.running,lxc list -f csv -c s | grep -c RUNNING
UserParameter=lxd.container.discovery,lxc list -f json | jq -c '{data:[.[]|{ "{#CT}": .name }]}'
```

**Netdata :** le plugin cgroups détecte automatiquement les conteneurs
LXC/LXD et affiche leurs métriques sans configuration. Idéal pour un
diagnostic rapide :

```bash
# Installer Netdata sur l'hôte (conteneurs détectés auto)
bash <(curl -Ss https://get.netdata.cloud/kickstart.sh)
```

---

## 67. Supervision — Logs et alertes

```bash
# Logs du daemon LXD
lxc monitor --type=logging --pretty

# Logs d'un conteneur (console)
lxc info c1 --show-log

# Logs dans le conteneur (systemd)
lxc exec c1 -- journalctl -p err -b

# Événements cycle de vie (démarrages, arrêts, snapshots)
lxc monitor --type=lifecycle --pretty
```

**Centralisation :** envoyez les logs vers rsyslog/Loki/ELK :

```bash
# Exemple : forwarder journald du conteneur vers l'hôte
lxc exec c1 -- journalctl -f -o json | /usr/local/bin/ship-to-loki.sh
```

**Alertes minimales à mettre en place :**

- [ ] Un conteneur attendu est arrêté > 5 min
- [ ] Espace disque pool > 85 %
- [ ] Échec de sauvegarde (cron)
- [ ] Snapshot automatique en échec

---

## 68. Production — Checklist de mise en production

Avant de déclarer un conteneur "prod", validez :

**Configuration :**

- [ ] Profil(s) appliqué(s), config versionnée dans Git
- [ ] Limites CPU/RAM/I/O/processus définies
- [ ] `boot.autostart=true` + priorité/délai cohérents
- [ ] Réseau : bridge dédié, IP statique réservée, DNS OK
- [ ] Stockage : volume custom pour les données, taille rootfs limitée

**Sécurité :**

- [ ] Non privilégié, nesting désactivé (sauf besoin documenté)
- [ ] SSH clés uniquement, root interdit, fail2ban si exposé
- [ ] Mises à jour de sécurité automatiques
- [ ] Pare-feu (hôte et/ou conteneur)

**Exploitation :**

- [ ] Snapshots planifiés + expiration
- [ ] Sauvegarde externalisée testée (restauration OK)
- [ ] Supervision : métriques + alertes (arrêt, disque, mémoire)
- [ ] Logs centralisés
- [ ] Documentation : rôle, contacts, procédure de redémarrage
- [ ] Snapshot "mise-en-prod" pris le jour J

---

## 69. Production — Clustering LXD (haute disponibilité)

LXD peut former un **cluster** : plusieurs hôtes gérés comme un seul,
avec répartition des conteneurs.

```bash
# Sur le 1er nœud : activer le clustering à l'init
sudo lxd init
# Would you like to use LXD clustering? yes
# → définir un trust password

# Sur les autres nœuds : joindre
sudo lxd init
# Would you like to use LXD clustering? yes
# → join existing cluster, adresse du 1er nœud + trust password

# Voir les membres
lxc cluster list

# Lancer un conteneur sur un nœud précis
lxc launch ubuntu:24.04 c1 --target hote2

# Migrer à chaud (avec CRIU si configuré, sinon à froid)
lxc move c1 --target hote3
lxc stop c1 && lxc move c1 --target hote3 && lxc start c1
```

**Prérequis réseau :** les nœuds doivent partager le stockage (ou utiliser
des volumes répliqués) et un réseau de cluster dédié. Pour du stockage
partagé : **Ceph** (via `lxc storage create ... ceph`).

**Limites honnêtes :** le clustering LXD répartit la gestion, mais la
**haute disponibilité automatique** (redémarrage sur panne) reste moins
intégrée que le HA de Proxmox VE. Évaluez selon votre SLA.

---

## 70. Production — Quotas et facturation interne

```bash
# Quota disque par projet (LXD "projects")
lxc project create equipe-reseau
lxc project set equipe-reseau limits.containers 10
lxc project set equipe-reseau limits.memory 32GiB
lxc project set equipe-reseau limits.cpu 16
lxc project set equipe-reseau limits.disk 200GiB

# Basculer sur le projet et créer dedans
lxc project switch equipe-reseau
lxc launch ubuntu:24.04 c1   # créé dans le projet, soumis aux quotas
lxc project switch default
```

**Facturation interne (chargeback) :** exportez `lxc list -f json` +
métriques Prometheus, calculez CPU/RAM/disque par projet/équipe, et
produisez un rapport mensuel. Les `user.*` keys permettent de taguer
(`user.cost-center=SRV-042`).

---

## 71. Proxmox VE — LXC sous Proxmox : `pct`

Proxmox VE intègre LXC nativement. Chaque conteneur a un **VMID** numérique.

| Action | LXD | Proxmox VE |
|---|---|---|
| Créer | `lxc launch ubuntu:24.04 c1` | `pct create 100 local:vztmpl/ubuntu-24.04.tar.zst ...` |
| Démarrer/arrêter | `lxc start/stop c1` | `pct start/stop 100` |
| Shell | `lxc exec c1 -- bash` | `pct enter 100` / `pct exec 100 -- bash` |
| Fichiers | `lxc file push/pull` | `pct push/pull 100 ...` |
| Config | `lxc config set` | `pct set 100 --memory 4096` |
| Snapshot | `lxc snapshot` | `pct snapshot 100 snap0` |
| Backup | `lxc export` | `vzdump 100` (intégré !) |
| Liste | `lxc list` | `pct list` |
| Supprimer | `lxc delete` | `pct destroy 100` |

**Créer un conteneur sous Proxmox :**

```bash
# 1. Télécharger un template (interface ou CLI)
pveam update
pveam download local ubuntu-24.04-standard_24.04-2_amd64.tar.zst

# 2. Créer le CT
pct create 100 local:vztmpl/ubuntu-24.04-standard_24.04-2_amd64.tar.zst \
  --hostname web1 \
  --cores 2 --memory 4096 --swap 0 \
  --rootfs local-lvm:20 \
  --net0 name=eth0,bridge=vmbr0,ip=dhcp \
  --unprivileged 1 \
  --features nesting=1 \
  --onboot 1 \
  --password 'mot-de-passe-temporaire'
```

---

## 72. Proxmox VE — Différences LXD vs `pct`

| Aspect | LXD | Proxmox `pct` |
|---|---|---|
| Nommage | Nom libre (`web1`) | VMID numérique (100) |
| Images | Remotes dynamiques (`ubuntu:`) | Templates `pveam` (téléchargés) |
| Profils | Natif, puissant | Pas d'équivalent (templates + cloud-init partiel) |
| Réseau | `lxdbr0`, OVN, fan | `vmbr0`, SDN (EVPN/QinQ), firewall intégré |
| Stockage | Pools LXD | Stockages Proxmox (ZFS, LVM, Ceph, NFS…) |
| Snapshots | Planifiables nativement | Natifs (selon stockage), planifiés via backup |
| Backup | Manuel (`export`) / script | **`vzdump` intégré + planifié** |
| Clustering/HA | Cluster LXD | Cluster corosync + **HA automatique** |
| API | REST moderne | API REST Proxmox (bien outillée : Terraform…) |
| Nesting | `security.nesting=true` | `features: nesting=1` |
| Privilégié | `security.privileged` | `--unprivileged 0` (déconseillé idem) |

**En résumé :** LXD = excellent sur hôte standalone ; Proxmox = plateforme
datacenter complète (HA, backups planifiés, firewall, SDN).

---

## 73. Proxmox VE — Migrer entre LXD et Proxmox

**De LXD vers Proxmox :**

```bash
# 1. Sur l'hôte LXD : exporter
lxc stop c1
lxc export c1 /tmp/c1.tar.gz

# 2. Transférer vers Proxmox
scp /tmp/c1.tar.gz root@proxmox:/tmp/

# 3. Sur Proxmox : extraire le rootfs
mkdir -p /tmp/c1-rootfs
tar -xzf /tmp/c1.tar.gz -C /tmp/c1-rootfs
# Le rootfs se trouve dans ./rootfs/ de l'archive

