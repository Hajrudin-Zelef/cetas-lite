---
id: collect-261001-rattrapage/rattrapage/lxc-guide-7
title: "LXC & LXD — Guide ultra-complet"
domain: rattrapage
role: reference
task: reference
actors: []
dates: ["2026-09-25", "2026-09-26"]
keywords: ["agent", "arr", "memory"]
source: docs/RAG/collect-261001-rattrapage/lxc_guide.md
source_anchor: ""
source_lines: [1675, 1956]
sha256: 060d44c41c3c260f0a9be6c709800a982adb2913ba0731d0d1dd2c791d0f743c
---

# LXC & LXD — Guide ultra-complet

```bash
# Mode NAT (défaut) : processus proxy sur l'hôte
lxc config device add c1 web proxy \
  listen=tcp:0.0.0.0:80 connect=tcp:127.0.0.1:80 nat=true

# Proxy UDP (ex. WireGuard, DNS)
lxc config device add c1 wg proxy \
  listen=udp:0.0.0.0:51820 connect=udp:127.0.0.1:51820

# Proxy vers un socket Unix
lxc config device add c1 dockersock proxy \
  listen=unix:/var/host-docker.sock connect=unix:/var/run/docker.sock \
  uid=1000 gid=1000 mode=0660

# Voir les connexions actives
lxc config device show c1 web
```

**Limites :** un proxy par port ; pour des dizaines de services, préférez
un reverse-proxy (section 78) ou macvlan.

---

## 57. Devices — NIC avancés

```bash
# Adresse MAC fixée (utile pour DHCP externe / filtrage)
lxc config device set c1 eth0 hwaddr 00:16:3e:aa:bb:cc

# MTU personnalisé
lxc config device set c1 eth0 mtu 9000

# File d'attente / offloading
lxc config device set c1 eth0 mtu 1500

# Plusieurs interfaces
lxc config device add c1 eth1 nic network=dmz0 name=eth1
lxc config device add c1 eth2 nic network=prod0 name=eth2

# Limiter le débit réseau du conteneur
lxc config device set c1 eth0 limits.ingress 100Mbit
lxc config device set c1 eth0 limits.egress 100Mbit
```

**Ordre des interfaces :** `eth0`, `eth1`… dans l'ordre d'ajout. Pour un
routage prévisible, fixez les noms (`name=wan`, `name=lan`) plutôt que de
compter sur l'ordre.

---

## 58. Backups — Export / Import (méthode native)

```bash
# Exporter un conteneur (arrêté de préférence) vers un fichier
lxc stop c1
lxc export c1 /backups/c1-2026-09-26.tar.gz

# Avec les snapshots inclus
lxc export c1 /backups/c1-full.tar.gz --instance-only=false

# Compressé et optimisé (backend ZFS : rapide)
lxc export c1 /backups/c1.tar.gz --optimized-storage

# Importer sur le même hôte ou un autre
lxc import /backups/c1-2026-09-26.tar.gz
lxc import /backups/c1-2026-09-26.tar.gz c1-restore
```

**Contenu de l'export :** rootfs + config + snapshots (si inclus).
**Ne contient pas :** les volumes custom attachés séparément → exportez-les
aussi (`lxc storage volume export`).

---

## 59. Backups — Stratégie snapshots + rsync / Borg

**Architecture recommandée :**

```
Conteneur (prod)
   │  snapshot ZFS @daily (récupération rapide locale)
   ▼
Export hebdo → /backups/ (fichiers .tar.gz)
   │  Borg vers NAS distant (dédupliqué, chiffré)
   ▼
Règle 3-2-1 : 3 copies, 2 supports, 1 hors site
```

```bash
# Export des volumes custom aussi
lxc storage volume export default vol-data /backups/vol-data.tar.gz
```

**Avec Borg (déduplication + chiffrement) :**

```bash
# Init du dépôt (une fois)
borg init --encryption=repokey /mnt/nas/borg-lxd

# Sauvegarde des exports
borg create /mnt/nas/borg-lxd::lxd-{now:%Y-%m-%d} /backups/*.tar.gz

# Rétention
borg prune /mnt/nas/borg-lxd --keep-daily=7 --keep-weekly=4 --keep-monthly=6
```

---

## 60. Backups — Script de sauvegarde automatisée

`/usr/local/bin/lxd-backup.sh` :

```bash
#!/bin/bash
# Sauvegarde LXD : export de tous les conteneurs + volumes, puis Borg.
set -euo pipefail

BACKUP_DIR="/backups/lxd"
DATE=$(date +%Y-%m-%d)
RETENTION_DAYS=14
BORG_REPO="/mnt/nas/borg-lxd"

mkdir -p "$BACKUP_DIR/$DATE"

echo "[1/4] Export des conteneurs..."
for ct in $(lxc list -f csv -c n status=running); do
    echo "  -> $ct"
    lxc export "$ct" "$BACKUP_DIR/$DATE/${ct}.tar.gz" --quiet
done

echo "[2/4] Export des volumes custom..."
for vol in $(lxc storage volume list default -f csv -c n | grep -v '^$'); do
    # on ignore les volumes root des instances (type container)
    if lxc storage volume show default "$vol" | grep -q "type: custom"; then
        echo "  -> $vol"
        lxc storage volume export default "$vol" \
            "$BACKUP_DIR/$DATE/vol-${vol}.tar.gz" --quiet
    fi
done

echo "[3/4] Envoi vers Borg..."
borg create --stats "$BORG_REPO::lxd-$DATE" "$BACKUP_DIR/$DATE"

echo "[4/4] Nettoyage local (+$RETENTION_DAYS jours)..."
find "$BACKUP_DIR" -maxdepth 1 -type d -mtime +$RETENTION_DAYS -exec rm -rf {} +
borg prune "$BORG_REPO" --keep-daily=7 --keep-weekly=4 --keep-monthly=6

echo "Sauvegarde terminee : $DATE"
```

```bash
chmod +x /usr/local/bin/lxd-backup.sh
# Cron : tous les jours à 2h
(crontab -l 2>/dev/null; echo "0 2 * * * /usr/local/bin/lxd-backup.sh >> /var/log/lxd-backup.log 2>&1") | crontab -
```

> 🔐 Stockez la passphrase Borg hors du script (variable d'environnement
> `BORG_PASSPHRASE` via un fichier protégé `chmod 600` sourcé par le cron).

---

## 61. Backups — Restauration complète (procédure)

**Scénario :** le conteneur `c-db1` est corrompu, on restaure la sauvegarde
de la veille.

```bash
# 1. Extraire l'archive Borg si besoin
borg extract /mnt/nas/borg-lxd::lxd-2026-09-25

# 2. Supprimer (ou renommer) le conteneur cassé
lxc stop c-db1 --force
lxc rename c-db1 c-db1-hs
# ou : lxc delete c-db1 --force

# 3. Réimporter
lxc import /backups/lxd/2026-09-25/c-db1.tar.gz

# 4. Restaurer le volume de données
lxc storage volume delete default pgdata --force  # si corrompu aussi
lxc storage volume import default /backups/lxd/2026-09-25/vol-pgdata.tar.gz pgdata

# 5. Démarrer et vérifier
lxc start c-db1
lxc exec c-db1 -- systemctl status postgresql
lxc exec c-db1 -- psql -U postgres -c "SELECT 1;"
```

**Testez votre restauration** au moins une fois par trimestre sur un
conteneur jetable (`c-db1-test`). Une sauvegarde non testée n'est pas une
sauvegarde.

---

## 62. Mises à jour — Hôte et conteneurs

**Hôte :**

```bash
# LXD via snap : mises à jour automatiques (canal suivi)
sudo snap refresh lxd

# Système hôte
sudo apt update && sudo apt upgrade -y
```

**Conteneurs — manuel :**

```bash
# Tous les conteneurs Ubuntu/Debian d'un coup
for ct in $(lxc list -f csv -c n status=running); do
    echo "=== $ct ==="
    lxc exec "$ct" -- apt update
    lxc exec "$ct" --env DEBIAN_FRONTEND=noninteractive -- \
        apt upgrade -y
done
```

**Conteneurs — automatique (`unattended-upgrades`) :**

```bash
# Dans chaque conteneur (ou via le template / profil cloud-init)
lxc exec c1 -- apt install -y unattended-upgrades
lxc exec c1 -- dpkg-reconfigure -plow unattended-upgrades
```

**Stratégie :** sécurité en auto (`unattended-upgrades`), montées de version
majeures planifiées avec snapshot préalable (section 38).

---

## 63. Mises à jour — Stratégie rolling (zéro coupure)

Pour un service en plusieurs instances (ex. 3 webs derrière un reverse-proxy) :

```
1. Snapshot de web1                    (lxc snapshot web1 avant-maj)
2. MAJ web1, redémarrage, tests        (lxc exec web1 -- apt upgrade -y)
3. Si OK → web2, puis web3
4. Si KO → lxc restore web1 avant-maj  (retour en < 1 min)
```

**Automatisation Ansible (extrait) :**

```yaml
- hosts: lxd_hosts
  serial: 1   # un conteneur à la fois
  tasks:
    - name: Snapshot avant MAJ
      command: lxc snapshot {{ inventory_hostname }} avant-maj-{{ ansible_date_time.date }}
    - name: Mise à jour
      command: lxc exec {{ inventory_hostname }} -- apt upgrade -y
      environment:
        DEBIAN_FRONTEND: noninteractive
```

---

## 64. Supervision — Métriques natives LXD

```bash
# État temps réel d'un conteneur
lxc info c1
# → CPU, mémoire, disque, réseau, processus, load

# Métriques au format OpenMetrics (pour Prometheus !)
curl --unix-socket /var/snap/lxd/common/lxd/unix.socket \
  lxd/1.0/metrics | head -50

# Avec authentification par certificat si HTTPS distant
```

**Métriques exposées :** `lxd_cpu_*`, `lxd_memory_*`, `lxd_disk_*`,
`lxd_network_*`, `lxd_procs_*` par instance. C'est la source la plus
fiable : elle vient des cgroups, pas d'un agent dans le conteneur.

---

## 65. Supervision — Prometheus + Grafana

**`prometheus.yml` (scrape du socket LXD via un exporter ou HTTPS) :**

