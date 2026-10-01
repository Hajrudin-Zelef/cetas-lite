---
id: collect-261001-rattrapage/rattrapage/lxc-guide-4
title: "LXC & LXD — Guide ultra-complet"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/lxc_guide.md
source_anchor: ""
source_lines: [794, 1103]
sha256: 73bfc0781492123e168f53e2d48d294bb0090df816ba15a63e967c1f929d783f
---

# Leases DHCP (via dnsmasq géré par LXD)
cat /var/snap/lxd/common/lxd/networks/lxdbr0/dnsmasq.leases 2>/dev/null
# ou sur install APT :
cat /var/lib/lxd/networks/lxdbr0/dnsmasq.leases 2>/dev/null
```

**Ce que fait LXD pour vous :**

- Crée le bridge et lui assigne l'IP passerelle (ex. `10.10.10.1/24`).
- Lance `dnsmasq` : DHCP + DNS local (résolution des noms de conteneurs !).
- Configure le NAT (iptables/nftables) vers l'extérieur.

**Test DNS intégré :**

```bash
lxc exec c1 -- ping -c2 c2   # résout le nom d'un autre conteneur du bridge !
```

---

## 25. Réseau — Créer son propre bridge

Pour séparer les environnements (prod / dev / dmz) :

```bash
# Bridge isolé sans NAT (réseau privé pur)
lxc network create dmz0 \
  ipv4.address=192.168.50.1/24 \
  ipv4.nat=false \
  ipv6.address=none

# Bridge avec NAT et DHCP custom
lxc network create prod0 \
  ipv4.address=10.20.0.1/24 \
  ipv4.nat=true \
  ipv4.dhcp.ranges=10.20.0.100-10.20.0.200 \
  ipv6.address=none

# Attacher un conteneur
lxc network attach prod0 c1 eth1
# ou via device :
lxc config device add c1 eth1 nic network=prod0 name=eth1
```

**Options utiles :**

```bash
lxc network set prod0 dns.domain prod.lan
lxc network set prod0 dns.mode managed
lxc network set prod0 ipv4.firewall true
```

---

## 26. Réseau — macvlan (IP directe sur le LAN)

Le conteneur obtient une IP du **LAN physique** (comme une VM bridgée Proxmox
sur vmbr0) :

```bash
# Créer un profil macvlan
lxc profile create lan
lxc profile device add lan eth0 nic \
  nictype=macvlan \
  parent=enp0s3 \
  name=eth0

# Lancer avec ce profil
lxc launch ubuntu:24.04 srv1 --profile lan
```

> ⚠️ **Limitation noyau :** en macvlan, l'**hôte ne peut pas communiquer**
> avec ses propres conteneurs macvlan (le parent filtre). Les conteneurs se
> voient entre eux, le LAN les voit, mais pas l'hôte. Si vous avez besoin
> d'hôte ↔ conteneur, préférez un bridge avec une interface physique attachée.

**Vérifier le parent :** `parent=` doit être l'interface physique réelle
(`ip link show` pour lister).

---

## 27. Réseau — VLAN (tagged 802.1Q)

```bash
# 1. Créer l'interface VLAN sur l'hôte (exemple VLAN 20 sur enp0s3)
sudo ip link add link enp0s3 name enp0s3.20 type vlan id 20
sudo ip link set enp0s3.20 up

# 2. Rendre persistant via netplan (/etc/netplan/60-vlan.yaml)
```

```yaml
network:
  version: 2
  ethernets:
    enp0s3: {}
  vlans:
    enp0s3.20:
      id: 20
      link: enp0s3
```

```bash
# 3. Bridge LXD sur le VLAN
lxc network create vlan20 \
  --type=bridge \
  bridge.external_interfaces=enp0s3.20 \
  ipv4.address=none \
  ipv6.address=none

# 4. Attacher
lxc config device add c1 eth1 nic network=vlan20 name=eth1
```

Le conteneur reçoit alors une IP du DHCP du VLAN 20 (celui de votre infra).

---

## 28. Réseau — IP statiques et DHCP

**IP statique gérée par LXD (recommandé) :** LXD réserve l'IP dans dnsmasq,
le conteneur reste en DHCP mais reçoit toujours la même IP :

```bash
# À chaud (redémarrage du conteneur requis pour prise en compte DHCP)
lxc config device set c1 eth0 ipv4.address 10.10.10.50
lxc restart c1
```

**IP statique dans le conteneur (netplan) :** possible mais LXD ne la connaît
pas (le `lxc list` n'affichera pas l'IP). Préférez la méthode LXD.

**Réservations DHCP visibles :**

```bash
lxc network show lxdbr0 | grep -B2 -A2 "10.10.10.50"
```

**Désactiver le DHCP sur un réseau :**

```bash
lxc network set lxdbr0 ipv4.dhcp false
```

---

## 29. Réseau — NAT et port forwarding (device proxy)

Pour exposer un service d'un conteneur vers l'hôte ou le LAN **sans macvlan** :

```bash
# Rediriger le port 8080 de l'hôte vers le port 80 du conteneur
lxc config device add c1 http proxy \
  listen=tcp:0.0.0.0:8080 \
  connect=tcp:127.0.0.1:80

# HTTPS
lxc config device add c1 https proxy \
  listen=tcp:0.0.0.0:8443 \
  connect=tcp:127.0.0.1:443

# Restreindre l'écoute à localhost (admin seulement)
lxc config device add c1 ssh-admin proxy \
  listen=tcp:127.0.0.1:2222 \
  connect=tcp:127.0.0.1:22

# Lister / supprimer
lxc config device list c1
lxc config device remove c1 http
```

**Fonctionnement :** un processus proxy tourne sur l'hôte et relaie le trafic.
Simple et efficace pour quelques ports ; pour du vrai reverse-proxy multi-sites,
voir le cas pratique section 78.

---

## 30. Réseau — IPv6

Par défaut `lxd init` propose `auto` (ULA `fd42:...`) ou `none`.

```bash
# Activer l'IPv6 ULA sur lxdbr0
lxc network set lxdbr0 ipv6.address auto
lxc network set lxdbr0 ipv6.nat true

# IPv6 routé public (si votre hébergeur route un /64 vers l'hôte)
lxc network set prod0 ipv6.address 2001:db8:abcd:10::1/64
lxc network set prod0 ipv6.nat false

# Désactiver IPv6 sur un conteneur
lxc config device set c1 eth0 ipv6.address none
```

**Vérification dans le conteneur :**

```bash
lxc exec c1 -- ip -6 addr show eth0
lxc exec c1 -- ping -c2 ipv6.google.com
```

> 📝 En entreprise, documentez votre plan d'adressage IPv6 comme l'IPv4.
> L'ULA (`fd00::/8`) convient pour l'interne sans exposition publique.

---

## 31. Réseau — Fan networking (overlay simple)

Le **fan** est un overlay VXLAN simplifié : chaque hôte obtient un sous-réseau
d'un grand /8, les conteneurs communiquent entre hôtes sans configuration réseau
complexe.

```bash
# Sur chaque hôte du "fan", avec un sous-réseau unique par hôte
lxc network create fan0 --type=bridge \
  bridge.mode=fan \
  fan.overlay_subnet=240.0.0.0/8 \
  fan.underlay_subnet=10.0.0.0/16
```

**Cas d'usage :** maquette multi-hôtes rapide, labo. Pour de la production
multi-hôtes sérieuse, préférez **OVN** (routage distribué, voir section 69).

---

## 32. Stockage — Concepts (pools, volumes)

```
Storage pool "default" (driver: zfs)
├── volume c1          (disque root du conteneur c1)
├── volume c2
├── volume custom/data (volume custom attaché à c1 en /data)
└── snapshot c1/snap0  (snapshot ZFS natif)
```

```bash
# Lister les pools
lxc storage list

# Infos détaillées (espace utilisé/disponible)
lxc storage info default

# Lister les volumes d'un pool
lxc storage volume list default
```

**Règle d'or :** un conteneur = un volume `root`. Les données applicatives
volumineuses (bases, uploads) méritent des **volumes custom** séparés
(sauvegarde et migration indépendantes).

---

## 33. Stockage — Pool `dir` (simple, sans dépendance)

Le driver `dir` stocke les conteneurs comme de simples dossiers. Aucune
fonctionnalité avancée (pas de snapshots instantanés, copies lentes).

```bash
lxc storage create dirpool dir source=/var/lib/lxd-disk
```

| ✅ Avantages | ❌ Inconvénients |
|---|---|
| Fonctionne partout, zéro prérequis | Pas de snapshots natifs |
| Facile à inspecter / rsync | Copies complètes (lentes) |
| | Pas de quotas efficaces |

**Usage :** dépannage, tout petit labo, ou quand ZFS/btrfs sont impossibles.
**En production : préférez ZFS.**

---

## 34. Stockage — Pool ZFS (recommandé)

```bash
# Sur un disque dédié (PRODUCTION) — efface le disque !
lxc storage create zfspool zfs source=/dev/sdb

# Sur une partition existante
lxc storage create zfspool zfs source=/dev/sdb1

# En fichier loop (TEST uniquement — performances limitées)
lxc storage create zfspool zfs size=50GiB

# Migrer le pool par défaut vers un disque réel après coup
lxc storage create prod-zfs zfs source=/dev/sdb
lxc storage set default source /dev/sdb   # selon contexte
```

**Pourquoi ZFS :**

- Snapshots et clones **instantanés** (COW).
- Copies de conteneurs quasi gratuites (`lxc copy`).
- Compression native (`compression=on` — souvent `lz4` par défaut).
- `zfs list`, `zfs snapshot` accessibles directement sur l'hôte.

**Réglages utiles :**

```bash
# Compression (si non active)
sudo zfs set compression=lz4 default/containers

# Voir l'espace réel
sudo zfs list -o name,used,avail,compressratio
```

---

## 35. Stockage — Pool btrfs

