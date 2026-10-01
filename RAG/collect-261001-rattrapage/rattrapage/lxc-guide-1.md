---
id: collect-261001-rattrapage/rattrapage/lxc-guide-1
title: "LXC & LXD — Guide ultra-complet"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["arr", "datacenter"]
source: docs/RAG/collect-261001-rattrapage/lxc_guide.md
source_anchor: ""
source_lines: [1, 233]
sha256: 72f66bc33eff6030602e5a69805695b77cbd3539b62d47cd035c4b0b0412c4da
---

# LXC & LXD — Guide ultra-complet

> **Public :** administrateurs systèmes, chefs de service, utilisateurs Proxmox VE.
> **Objectif :** maîtriser les conteneurs Linux de A à Z — installation, réseau, stockage,
> sécurité, production, supervision, sauvegarde, dépannage — avec un lien permanent
> vers Proxmox VE (l'outil préféré de Zelef).
> **Convention :** les commandes préfixées `#` s'exécutent en root sur l'hôte,
> celles préfixées `$` en utilisateur simple. À l'intérieur d'un conteneur, on note `[ct] #`.

---

## 0. Comment utiliser ce guide

- Lisez les sections 1 à 10 pour démarrer en une heure.
- Les sections 11 à 50 couvrent le quotidien et l'exploitation.
- Les sections 51 à 74 traitent production, sécurité et Proxmox.
- Les sections 75 à 86 sont la pratique : cas concrets, erreurs, dépannage, quiz.
- Le **pense-bête de poche** (section 83) est imprimable pour le bureau.

---

## 1. Introduction — Pourquoi LXC ?

LXC (*Linux Containers*) est la technologie de conteneurisation native du noyau Linux :
des processus isolés grâce aux **namespaces** (pid, net, mnt, uts, ipc, user, cgroup)
et limités en ressources par les **cgroups**. Pas d'hyperviseur, pas de noyau invité :
le conteneur partage le noyau de l'hôte.

**Ce que ça change concrètement :**

| Aspect | Machine virtuelle | Conteneur LXC |
|---|---|---|
| Démarrage | 30 s à 2 min | < 2 secondes |
| Surcoût mémoire | Noyau invité complet (~200-500 Mo) | Quelques Mo |
| Densité | Dizaines par hôte | Centaines par hôte |
| Isolation | Totale (noyau séparé) | Forte (noyau partagé) |
| Cas d'usage | OS hétérogènes, noyaux spécifiques | Services Linux, densité, rapidité |

**En une phrase :** là où une VM virtualise du matériel, LXC virtualise un système
d'exploitation — c'est plus léger, plus rapide, et largement suffisant pour la
majorité des services Linux en production.

---

## 2. LXC vs LXD vs Docker vs Proxmox — Le grand comparatif

| Critère | LXC (pur) | LXD | Docker / Podman | Proxmox VE (LXC) |
|---|---|---|---|---|
| Nature | Bibliothèque + outils bas niveau | Gestionnaire de conteneurs système | Conteneurs applicatifs | Hyperviseur + conteneurs |
| Unité gérée | Conteneur brut | Conteneur système complet | Image applicative | CT (via `pct`) + VM (via `qm`) |
| Réseau intégré | Non (manuel) | Oui (bridges, OVN) | Oui (bridge docker0) | Oui (vmbr0, SDN) |
| Stockage intégré | Non | Oui (pools ZFS/btrfs/LVM/dir) | Volumes/drivers | Oui (ZFS, LVM, Ceph…) |
| Snapshots | Manuel (selon FS) | Natifs, planifiables | Layers d'image | Natifs (selon stockage) |
| Clustering | Non | Oui (natif) | Swarm/K8s (externe) | Oui (corosync) |
| Idéal pour | Embarqué, cas spécifiques | Infra système, VPS-like | Microservices, CI/CD | Datacenter complet |

**Retenez :** LXC = le moteur, LXD = la voiture complète construite autour.
Proxmox VE utilise LXC en interne via sa propre couche (`pct`) — voir sections 71-74.

---

## 3. Que choisir : LXC, LXD, Docker ou VM ?

**Arbre de décision :**

```
Besoin d'un autre noyau / OS non-Linux (Windows, BSD) ?
└── OUI → Machine virtuelle (KVM / Proxmox VE)
└── NON (Linux uniquement)
    ├── Un seul processus applicatif, image portable, CI/CD ?
    │   └── OUI → Docker / Podman
    └── NON → Un "mini-serveur" complet (systemd, ssh, plusieurs services) ?
        ├── Gestion centralisée type datacenter ? → Proxmox VE (LXC via pct)
        └── Hôte Debian/Ubuntu simple, automatisation ? → LXD
```

**Règle pratique de Zelef :**

- **Proxmox VE** = la plateforme de production (cluster, HA, backups intégrés).
- **LXD** = parfait sur un hôte isolé, un labo, un serveur dédié sans Proxmox,
  ou quand on veut l'API REST moderne et le clustering léger sans tout Proxmox.
- **Docker** = à l'intérieur d'un conteneur LXC si besoin (nesting, section 51),
  ou directement sur l'hôte pour des applis 12-factor.

---

## 4. Prérequis système et noyau

**Matériel :**

- CPU 64 bits (x86_64 ou ARM64), 2 cœurs minimum, 4 recommandés.
- RAM : 4 Go minimum, 8 Go+ recommandés (les conteneurs sont légers, mais ZFS adore la RAM).
- Disque : 20 Go minimum ; SSD fortement recommandé.
- Virtualisation matérielle **non requise** (pas d'hyperviseur).

**Noyau :** Linux ≥ 5.4 recommandé (cgroups v2 bien supportés à partir de 5.x ;
Ubuntu 22.04/24.04 et Debian 11/12/13 conviennent).

**Vérifications préalables :**

```bash
# Version du noyau
uname -r

# Architecture
uname -m

# cgroups v2 actif ? (doit afficher "cgroup2fs")
stat -fc %T /sys/fs/cgroup

# Vérifier le support LXC du noyau (si lxc installé)
lxc-checkconfig
```

**Sortie attendue de `lxc-checkconfig` :** la plupart des lignes en vert `enabled`.
Les lignes `missing` sur `Cgroup v1 ...` sont normales sur un système cgroup v2.

**Espace disque pour les images :** prévoir ~500 Mo par image de base
(Ubuntu ~ 200 Mo, Debian ~ 150 Mo téléchargés, plus une fois décompressés).

---

## 5. Installation sur Ubuntu — via snap (méthode recommandée)

Sur Ubuntu, LXD est distribué et maintenu via **snap** (version toujours à jour).

```bash
# Mise à jour du système
sudo apt update && sudo apt upgrade -y

# Installation de LXD
sudo snap install lxd

# Vérification
lxd --version
lxc --version
```

**Ajouter votre utilisateur au groupe `lxd`** (évite `sudo` à chaque commande) :

```bash
sudo usermod -aG lxd "$USER"
# Puis DÉCONNEXION / RECONNEXION (ou newgrp lxd) pour prendre en compte le groupe
newgrp lxd
```

> ⚠️ **Sécurité :** un membre du groupe `lxd` équivaut à root sur l'hôte
> (un conteneur privilégié peut monter le disque hôte). Ne mettez dans ce groupe
> que des administrateurs de confiance. Voir section 46.

**Canal snap :** par défaut `latest/stable`. Pour une version LTS figée :

```bash
sudo snap refresh lxd --channel=5.21/stable   # exemple : branche LTS 5.21
```

---

## 6. Installation sur Debian — dépôts et alternatives

Debian propose LXD dans ses dépôts (version parfois en retard d'une release) :

```bash
sudo apt update && sudo apt upgrade -y

# Option A : paquet Debian (simple, version du dépôt)
sudo apt install -y lxd lxd-client

# Vérification
lxd --version
```

**Option B : snap sur Debian** (pour la dernière version, comme Ubuntu) :

```bash
sudo apt install -y snapd
sudo snap install core
sudo snap install lxd
```

**Comparatif :**

| Méthode | Avantages | Inconvénients |
|---|---|---|
| `apt install lxd` | Intégré au système, mises à jour avec APT | Version en retard (ex. 5.x sur Debian 12) |
| `snap install lxd` | Toujours à jour, canaux LTS | Dépendance snapd, chemins `/snap` |

**Recommandation :** en production, préférez le **canal LTS du snap**
(stabilité + correctifs de sécurité), ou le paquet APT si votre politique
interdit snapd.

---

## 7. Installation de LXC pur (sans LXD) — pour comprendre les fondations

LXD repose sur LXC. Installer LXC seul est utile pour comprendre, ou pour des
usages embarqués/minimaux.

```bash
# Debian / Ubuntu
sudo apt install -y lxc lxc-templates uidmap

# Vérifier la configuration noyau
sudo lxc-checkconfig

# Créer un conteneur Debian (template download)
sudo lxc-create -t download -n c1 -- --dist debian --release bookworm --arch amd64

# Démarrer / attacher / arrêter
sudo lxc-start -n c1
sudo lxc-attach -n c1
sudo lxc-stop -n c1

# Lister
sudo lxc-ls --fancy
```

**Quand utiliser LXC pur ?**

- Systèmes très contraints, scripts sur mesure.
- Comprendre ce que LXD automatise (réseau, stockage, images).
- **En pratique :** préférez LXD pour tout usage serveur standard.

---

## 8. `lxd init` — Initialisation interactive

```bash
sudo lxd init
```

**Déroulé typique (réponses recommandées entre crochets) :**

