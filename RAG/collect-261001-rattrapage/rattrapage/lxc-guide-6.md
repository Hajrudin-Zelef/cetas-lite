---
id: collect-261001-rattrapage/rattrapage/lxc-guide-6
title: "LXC & LXD — Guide ultra-complet"
domain: rattrapage
role: reference
task: reference
actors: ["Nvidia"]
dates: []
keywords: ["attention", "gpu", "memory", "nvidia"]
source: docs/RAG/collect-261001-rattrapage/lxc_guide.md
source_anchor: ""
source_lines: [1397, 1674]
sha256: 241dbb08936a2de17c88c36853974eead5db2ffdac897379998ddcc85b964520
---

# LXC & LXD — Guide ultra-complet

> 🛡️ **Règle absolue :** restez en **non privilégié**. Si un logiciel exige
> le mode privilégié, cherchez d'abord les capabilities fines
> (`security.syscalls.*`, devices spécifiques) avant de céder.

---

## 47. Sécurité — Bonnes pratiques d'isolation

**Checklist par conteneur :**

- [ ] Non privilégié (`security.privileged=false`)
- [ ] Nesting désactivé sauf besoin (`security.nesting=false` par défaut)
- [ ] Limites CPU/RAM/processus définies (anti-DoS du voisin)
- [ ] Pas de device `disk` vers l'hôte sauf besoin métier documenté
- [ ] Réseau : le conteneur n'a que les interfaces nécessaires
- [ ] `boot.autostart` cohérent (évite les démarrages surprises)
- [ ] Mises à jour automatiques ou planifiées (section 62)

**Isolation réseau :** un bridge par zone de confiance (section 25).
Un conteneur compromis sur `dmz0` ne doit pas voir `prod0`.

---

## 48. Sécurité — AppArmor et seccomp

LXD applique par défaut des profils **AppArmor** et un filtre **seccomp**
qui bloquent les appels système dangereux dans le conteneur.

```bash
# Voir le profil AppArmor appliqué
lxc config get c1 raw.apparmor

# Surcharger (exemple : autoriser un syscall pour une appli exotique)
lxc config set c1 raw.apparmor |
# via edit :
lxc config edit c1
```

```yaml
config:
  raw.apparmor: |
    mount options=(rw, bind),
```

**Seccomp :**

```bash
# Mode par défaut (filtre strict)
lxc config get c1 security.syscalls.deny_default
# Personnaliser la liste
lxc config set c1 security.syscalls.allow "personality sched_setattr"
```

> ⚠️ Ne touchez à AppArmor/seccomp que si une application légitime est
> bloquée (logs `audit`/`dmesg` sur l'hôte). Documentez chaque exception.

---

## 49. Sécurité — idmap et UID/GID

Le **user namespace** mappe les UIDs du conteneur vers une plage haute
de l'hôte :

```
Conteneur : root (0)  →  Hôte : 1000000
Conteneur : www-data (33) → Hôte : 1000033
```

**Partager un dossier hôte avec les bons droits :**

```bash
# Le dossier hôte doit appartenir à la plage mappée
sudo chown -R 1000000:1000000 /srv/partage-c1

lxc config device add c1 partage disk \
  source=/srv/partage-c1 \
  path=/mnt/partage
```

**Trouver la plage d'un conteneur :**

```bash
lxc config get c1 volatile.idmap.next
# ou
grep -A2 "c1" /var/snap/lxd/common/lxd/containers/c1/rootfs 2>/dev/null
cat /proc/$(lxc info c1 | awk '/^PID/{print $2}')/uid_map
```

**idmap custom (faire correspondre un UID précis) :**

```bash
lxc config set c1 raw.idmap "uid 1000 1000
gid 1000 1000"
```

---

## 50. Sécurité — Durcir un conteneur (checklist opérationnelle)

```bash
# 1. Base : non privilégié, pas de nesting inutile
lxc config set c1 security.privileged false
lxc config set c1 security.nesting false

# 2. Interdire les montages sensibles (déjà restreint par défaut en non-privilégié)
# 3. Limiter les syscalls réseau bruts si inutile
lxc config set c1 security.syscalls.deny_compat true

# 4. Lecture seule du rootfs quand possible (applis stateless)
#    → monter les zones d'écriture en volumes séparés
lxc config device add c1 app-data disk pool=default source=appdata path=/data

# 5. Dans le conteneur : SSH durci, pas de root distant
lxc exec c1 -- sed -i 's/^#*PermitRootLogin.*/PermitRootLogin no/' /etc/ssh/sshd_config
```

**Checklist finale :**

- [ ] `security.privileged=false`, `security.nesting=false` (sauf besoin)
- [ ] Limites de ressources posées
- [ ] SSH : clés uniquement, root interdit
- [ ] Pare-feu dans le conteneur (ufw/nftables) ou au bridge
- [ ] Mises à jour auto configurées
- [ ] Logs centralisés (section 67)
- [ ] Snapshot pré-change + sauvegarde externalisée

---

## 51. Nesting — Docker dans LXC : le principe

Faire tourner **Docker à l'intérieur d'un conteneur LXC** : très pratique
(un "hôte Docker" léger et jetable), mais ça demande d'assouplir l'isolation.

```bash
# Configuration minimale requise
lxc config set c-docker security.nesting true

# Recommandé : autoriser les appels système nécessaires
lxc config set c-docker security.syscalls.intercept.mknod true
lxc config set c-docker security.syscalls.intercept.setxattr true

lxc restart c-docker
```

**Puis dans le conteneur :**

```bash
lxc exec c-docker -- bash
[ct] # apt update && apt install -y docker.io
[ct] # systemctl enable --now docker
[ct] # docker run hello-world
```

---

## 52. Nesting — Configuration requise détaillée

**Profil `docker-host` prêt à l'emploi :**

```yaml
config:
  security.nesting: "true"
  security.syscalls.intercept.mknod: "true"
  security.syscalls.intercept.setxattr: "true"
  linux.kernel_modules: ip_tables,ip6_tables,overlay,br_netfilter
  limits.memory: 8GiB
  boot.autostart: "true"
description: Hote Docker dans LXC
devices:
  eth0: {name: eth0, network: lxdbr0, type: nic}
  root: {path: /, pool: default, type: disk, size: 60GiB}
name: docker-host
```

```bash
lxc profile create docker-host
cat profil-docker-host.yaml | lxc profile edit docker-host
lxc launch ubuntu:24.04 docker1 --profile default --profile docker-host
```

**Points d'attention :**

- Préférez le driver de stockage **overlay2** de Docker (défaut) ; évitez
  `vfs` (lent) et le double COW ZFS-sur-ZFS (désactivez la compression
  redondante si besoin).
- **Ne pas** empiler : Docker dans LXC dans Docker = problèmes garantis.
- Sur Proxmox, la même option existe : `features: nesting=1` (section 71).
- Alternative moderne : **Podman rootless** dans le conteneur (moins d'exigences).

---

## 53. Devices — Disques (disk device)

```bash
# Dossier de l'hôte monté dans le conteneur
lxc config device add c1 data disk source=/srv/data-c1 path=/data

# Options de montage
lxc config device add c1 data disk \
  source=/srv/data-c1 \
  path=/data \
  readonly=true \
  shift=true

# Volume custom du pool (recommandé vs dossier hôte)
lxc storage volume create default vol-data size=50GiB
lxc config device add c1 data disk pool=default source=vol-data path=/data

# Limiter la taille du rootfs
lxc config device set c1 root size 30GiB

# Voir les devices
lxc config device list c1
lxc config device show c1 data
```

**`shift=true` :** LXD décale les UIDs à la volée pour que les fichiers
apparaissent avec les bons propriétaires des deux côtés. Pratique pour les
dossiers partagés hôte ↔ conteneur.

---

## 54. Devices — GPU

```bash
# Passer TOUS les GPU au conteneur
lxc config device add c1 gpu gpu

# Passer un GPU précis (NVIDIA, par PCI)
lxc config device add c1 gpu0 gpu pci=0000:01:00.0 id=0

# Vérifier dans le conteneur
lxc exec c1 -- nvidia-smi
lxc exec c1 -- ls /dev/dri
```

**Prérequis :** pilotes installés sur l'**hôte** (le conteneur utilise le
device via le noyau hôte). Pour NVIDIA, le toolkit container fonctionne
aussi bien dans LXC que dans Docker.

**Cas d'usage :** transcodage vidéo (Jellyfin/Plex), inférence IA légère,
build avec accélération.

---

## 55. Devices — USB

```bash
# Lister les USB de l'hôte
lsusb

# Passer une clé USB précise (par IDs vendeur/produit)
lxc config device add c1 usbkey usb vendorid=0951 productid=1666

# Passer un port USB entier
lxc config device add c1 usbport usb bus=1 device=4

# Dongle série (ex. Zigbee / onduleur en USB)
lxc config device add c1 zigbee unix-char path=/dev/ttyUSB0 source=/dev/ttyUSB0
```

**Exemple concret — supervision d'onduleur (lien avec le guide onduleurs) :**

```bash
# L'onduleur est branché en USB sur l'hôte, NUT tourne dans le conteneur
lxc config device add c-nut ups usb vendorid=051d productid=0002
lxc exec c-nut -- apt install -y nut
```

> ⚠️ Un device USB passé au conteneur disparaît de l'hôte. Documentez
> quel conteneur détient quel périphérique physique.

---

## 56. Devices — Proxy (redirection de ports) — approfondi

Déjà vu en section 29, voici les modes avancés :

