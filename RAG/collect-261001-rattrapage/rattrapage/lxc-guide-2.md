---
id: collect-261001-rattrapage/rattrapage/lxc-guide-2
title: "LXC & LXD — Guide ultra-complet"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["arr", "attention", "memory"]
source: docs/RAG/collect-261001-rattrapage/lxc_guide.md
source_anchor: ""
source_lines: [234, 497]
sha256: 07c7f064242f635910cefa7b060fd2735de97c66842edd38dae450b023269311
---

# LXC & LXD — Guide ultra-complet

```
Would you like to use LXD clustering? (yes/no) [default=no]: no
Do you want to configure a new storage pool? (yes/no) [default=yes]: yes
Name of the new storage pool [default=default]: default
Name of the storage backend to use (dir, lvm, zfs, btrfs) [default=zfs]: zfs
Create a new ZFS pool? (yes/no) [default=yes]: yes
Would you like to use an existing empty block device? (yes/no) [default=no]: no
Size in GiB of the new loop device (1GiB minimum) [default=30GiB]: 50GiB
Would you like to connect to a MAAS server? (yes/no) [default=no]: no
Would you like to create a new local network bridge? (yes/no) [default=yes]: yes
What should the new bridge be called? [default=lxdbr0]: lxdbr0
What IPv4 address should be used? (CIDR subnet notation, "auto" or "none") [default=auto]: 10.10.10.1/24
Would you like LXD to NAT IPv4 traffic on this bridge? [default=yes]: yes
What IPv6 address should be used? (CIDR subnet notation, "auto" or "none") [default=auto]: none
Would you like the LXD server to be available over the network? (yes/no) [default=no]: no
Would you like stale cached images to be updated automatically? (yes/no) [default=yes]: yes
Would you like a YAML "lxd init" preseed to be printed? (yes/no) [default=no]: no
```

**Points d'attention :**

- **Backend `zfs`** : le meilleur choix (snapshots instantanés, copies COW).
  En loop device c'est parfait pour tester ; en production, utilisez un vrai
  disque ou une partition (section 34).
- **`lxdbr0`** : bridge NATé privé. Les conteneurs obtiennent des IP via DHCP
  interne (dnsmasq géré par LXD).
- **Accès réseau au daemon** : laissez `no` sauf besoin distant explicite
  (dans ce cas, TLS + mot de passe, section 69).

---

## 9. `lxd init` — Mode preseed (reproductible, scriptable)

Pour automatiser (Ansible, cloud-init, documentation d'exploitation) :

```bash
cat <<'EOF' > /tmp/lxd-preseed.yaml
config:
  core.https_address: '[::]:8443'
  core.trust_password: "ChangeMe_TresLong_MotDePasse_2026"
networks:
- name: lxdbr0
  type: bridge
  config:
    ipv4.address: 10.10.10.1/24
    ipv4.nat: "true"
    ipv6.address: none
storage_pools:
- name: default
  driver: zfs
  config:
    size: 50GiB
profiles:
- name: default
  devices:
    eth0:
      name: eth0
      network: lxdbr0
      type: nic
    root:
      path: /
      pool: default
      type: disk
cluster: null
EOF

cat /tmp/lxd-preseed.yaml | sudo lxd init --preseed
```

> 🔐 Ne versionnez **jamais** un preseed contenant `trust_password` dans Git.
> Utilisez Ansible Vault ou un coffre de secrets, puis supprimez le fichier.

**Vérification post-init :**

```bash
lxc network list
lxc storage list
lxc profile show default
```

---

## 10. Première vérification de santé

```bash
# Le daemon répond ?
lxc list

# Réseau : le bridge existe avec son IP
ip addr show lxdbr0

# Stockage : le pool est monté
lxc storage info default
zfs list 2>/dev/null | head

# Version et extensions
lxc info | head -30
```

**Checklist de santé initiale :**

- [ ] `lxc list` ne renvoie pas d'erreur de connexion au socket
- [ ] `lxdbr0` existe avec l'IP configurée
- [ ] Le pool `default` apparaît dans `lxc storage list`
- [ ] Le profil `default` attache `eth0` à `lxdbr0` et `root` au pool
- [ ] `lxc image list` fonctionne (même vide au début)

Si `lxc list` échoue avec une erreur de permission : vérifiez l'appartenance
au groupe `lxd` (section 5) ou utilisez `sudo`.

---

## 11. Images — Remotes et liste

LXD télécharge des images depuis des **remotes** (serveurs d'images).

```bash
# Lister les remotes configurées
lxc remote list

# Remotes par défaut :
#   ubuntu:         images Ubuntu officielles (Canonical)
#   ubuntu-daily:   builds quotidiens Ubuntu
#   images:         images communautaires linuxcontainers.org
#   local:          (défini par défaut) votre serveur local
```

```bash
# Lister les images Ubuntu 24.04 disponibles
lxc image list ubuntu: 24.04

# Chercher Debian
lxc image list images: debian

# Détails d'une image
lxc image info ubuntu:24.04
```

**Types d'images :**

| Type | Description | Usage |
|---|---|---|
| `container` | Image de conteneur système | `lxc launch` |
| `virtual-machine` | Image de VM (LXD gère aussi des VM !) | `lxc launch --vm` |

> 💡 Oui, **LXD gère aussi des machines virtuelles** (`--vm`) via QEMU.
> Ce guide se concentre sur les conteneurs, mais sachez que l'option existe
> pour les cas nécessitant un noyau séparé.

---

## 12. Images — Alias, architectures et cycle de vie

```bash
# Alias pratiques (noms courts)
lxc launch ubuntu:24.04 c-web1        # alias "24.04"
lxc launch ubuntu:22.04 c-web2
lxc launch images:debian/12 c-db1
lxc launch images:alpine/3.20 c-tiny1

# Architecture explicite (utile sur ARM)
lxc launch ubuntu:24.04 c-arm1 --target-arch aarch64
```

**Gestion locale des images :**

```bash
# Lister les images en cache local
lxc image list

# Copier une image d'un remote vers le local avec un alias
lxc image copy ubuntu:24.04 local: --alias ubuntu-2404-base

# Renommer / aliaser
lxc image alias create ubuntu-2404-base ubuntu-lts

# Supprimer une image locale
lxc image delete ubuntu-2404-base

# Forcer le rafraîchissement
lxc image refresh ubuntu:24.04
```

**Expiration du cache :** `images.remote_cache_expiry` (défaut 10 jours) et
`images.auto_update_interval`. Les images inutilisées sont nettoyées
automatiquement — ne vous inquiétez pas de voir le cache grossir.

---

## 13. `lxc launch` — Votre premier conteneur

```bash
# Lancement simple (profil default : réseau lxdbr0 + disque root)
lxc launch ubuntu:24.04 c1

# Suivre le démarrage
lxc list
```

**Sortie de `lxc list` :**

```
+------+---------+---------------------+-----------------------------------------------+-----------+-----------+
| NAME |  STATE  |        IPV4         |                     IPV6                      |   TYPE    | SNAPSHOTS |
+------+---------+---------------------+-----------------------------------------------+-----------+-----------+
| c1   | RUNNING | 10.10.10.52 (eth0)  | fd42::... (eth0)                              | CONTAINER | 0         |
+------+---------+---------------------+-----------------------------------------------+-----------+-----------+
```

**Options utiles au lancement :**

```bash
# Avec un nom d'hôte et des profils spécifiques
lxc launch ubuntu:24.04 web1 --profile default --profile web

# Avec des limites directes (config)
lxc launch ubuntu:24.04 db1 -c limits.cpu=2 -c limits.memory=4GiB

# Avec une config réseau spécifique
lxc launch ubuntu:24.04 c2 -d eth0,ipv4.address=10.10.10.50

# En mode "frozen" (créé mais pas démarré)
lxc init ubuntu:24.04 c3
lxc start c3
```

**Différence `launch` vs `init` :** `launch` = `init` + `start`.
Utilisez `init` quand vous voulez configurer avant le premier démarrage.

---

## 14. Cycle de vie — start, stop, restart, delete, freeze

```bash
lxc start c1                 # démarrer
lxc stop c1                  # arrêt propre (SIGPWR → systemd → timeout puis SIGKILL)
lxc stop c1 --force          # arrêt brutal immédiat
lxc restart c1               # redémarrage
lxc restart c1 --force       # redémarrage brutal
lxc freeze c1                # suspendre (gel des processus, mémoire conservée)
lxc unfreeze c1              # dégeler
lxc delete c1                # supprimer (le conteneur doit être arrêté)
lxc delete c1 --force        # supprimer même si en cours d'exécution
```

**Comportement de `lxc stop` :** LXD envoie un signal d'arrêt propre et attend
`boot.host_shutdown_timeout` (défaut 30s) avant de forcer. Pour des bases de
données, augmentez ce timeout :

```bash
lxc config set c-db1 boot.host_shutdown_timeout 120
```

**Suppression en masse :**

```bash
# Supprimer tous les conteneurs arrêtés (attention !)
lxc list -c n -f csv | xargs -r -n1 lxc delete --force
```

---

## 15. `lxc exec` — Exécuter des commandes dans le conteneur

