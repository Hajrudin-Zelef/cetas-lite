---
id: collect-261001-rattrapage/rattrapage/lxc-guide-11
title: "LXC & LXD — Guide ultra-complet"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["gpu"]
source: docs/RAG/collect-261001-rattrapage/lxc_guide.md
source_anchor: ""
source_lines: [2686, 2804]
sha256: caa9aef1bc5633b4178e6d49019e15ee645a6dd65ba4c76a9cea27a2a92a6f99
---

# LXC & LXD — Guide ultra-complet

=== INFOS ===
lxc info c1                           détails + métriques
lxc info c1 --show-log                log console
lxc monitor --pretty                  événements temps réel

=== PROXMOX (équivalents) ===
pct create|start|stop|destroy          cycle de vie
pct exec|enter <vmid>                 shell
pct push|pull <vmid>                  fichiers
pct snapshot <vmid> <nom>             snapshot
vzdump <vmid>                         backup intégré
```

---

## 84. Glossaire

| Terme | Définition |
|---|---|
| **LXC** | Technologie noyau de conteneurs (namespaces + cgroups) + outils bas niveau |
| **LXD** | Gestionnaire de conteneurs système construit sur LXC (daemon + CLI + API) |
| **Namespace** | Mécanisme noyau d'isolation (pid, net, mnt, uts, ipc, user, cgroup) |
| **Cgroup** | Mécanisme noyau de limitation/comptage des ressources (CPU, RAM, I/O) |
| **Image** | Modèle de système de fichiers pour créer des conteneurs |
| **Remote** | Serveur d'images (`ubuntu:`, `images:`, `local:`) |
| **Profil** | Ensemble réutilisable de config/devices appliqué aux conteneurs |
| **Pool** | Réserve de stockage gérée par LXD (dir, zfs, btrfs, lvm, ceph) |
| **Volume** | Unité de stockage dans un pool (rootfs ou volume custom) |
| **Snapshot** | Photographie instantanée d'un conteneur (restauration rapide) |
| **Device** | Ressource attachée : `disk`, `nic`, `proxy`, `gpu`, `usb`, `unix-char`… |
| **Bridge** | Commutateur virtuel Linux reliant conteneurs et/ou physique |
| **macvlan** | NIC virtuelle exposant le conteneur directement sur le LAN |
| **Nesting** | Exécuter des conteneurs dans un conteneur (ex. Docker dans LXC) |
| **Privilégié** | Conteneur dont root = root hôte (à éviter absolument) |
| **idmap** | Correspondance UID/GID conteneur ↔ hôte (user namespace) |
| **AppArmor** | MAC : confine les programmes par profils (complète LXD) |
| **seccomp** | Filtre les appels système autorisés dans le conteneur |
| **COW** | Copy-On-Write : les snapshots/clones ne dupliquent que les différences |
| **OVN** | Réseau virtuel distribué (overlay multi-hôtes pour LXD) |
| **CRIU** | Migration à chaud (checkpoint/restore) des processus |
| **pct** | CLI Proxmox pour les conteneurs LXC |
| **vzdump** | Outil de backup intégré de Proxmox VE |
| **cloud-init** | Standard de configuration initiale (hostname, SSH, réseau…) |

---

## 85. Quiz — 10 questions + réponses

**Q1. Quelle est la différence fondamentale entre LXC et LXD ?**
> R : LXC est la technologie bas niveau (bibliothèque + outils) ; LXD est le
> gestionnaire complet (daemon, API REST, réseau, stockage, images, clustering)
> construit au-dessus de LXC.

**Q2. Pourquoi un conteneur LXC démarre-t-il en moins de 2 secondes ?**
> R : Il partage le noyau de l'hôte : pas de BIOS, pas de bootloader, pas de
> noyau à charger — seul l'init du conteneur (systemd) démarre.

**Q3. Que signifie `lxc launch` par rapport à `lxc init` ?**
> R : `launch` = `init` (créer) + `start` (démarrer). `init` seul permet de
> configurer le conteneur avant son premier démarrage.

**Q4. Un membre du groupe `lxd` peut-il devenir root sur l'hôte ?**
> R : Oui. Via un conteneur privilégié ou un montage du disque hôte, c'est
> une équivalence root. Ne jamais y mettre d'utilisateur non-admin.

**Q5. Pourquoi préférer ZFS comme backend de stockage ?**
> R : Snapshots/clones instantanés (COW), copies quasi gratuites, compression
> native, quotas — idéal pour l'exploitation intensive de conteneurs.

**Q6. Quelle commande expose le port 80 d'un conteneur sur le port 8080 de l'hôte ?**
> R : `lxc config device add c1 http proxy listen=tcp:0.0.0.0:8080 connect=tcp:127.0.0.1:80`

**Q7. Que faut-il activer pour faire tourner Docker dans un conteneur LXC ?**
> R : `security.nesting=true` (et idéalement les interceptions de syscalls
> `mknod`/`setxattr`), puis redémarrer le conteneur.

**Q8. Un snapshot LXD remplace-t-il une sauvegarde externalisée ?**
> R : Non. Il protège des erreurs humaines et permet un rollback rapide, mais
> il vit sur le même disque/hôte : il ne protège ni du crash disque ni du site.

**Q9. Quelle est la limitation connue du macvlan concernant l'hôte ?**
> R : L'hôte ne peut pas communiquer avec ses propres conteneurs macvlan
> (filtrage au niveau du parent par le noyau).

**Q10. Citez trois équivalences entre LXD et Proxmox `pct`.**
> R : `lxc exec` ↔ `pct exec`/`pct enter` ; `lxc file push/pull` ↔
> `pct push/pull` ; `lxc snapshot` ↔ `pct snapshot` ; `lxc export` ↔ `vzdump`.

---

## 86. Pour aller plus loin

**Documentation officielle :**

- LXD — https://documentation.ubuntu.com/lxd/ (référence complète, API REST)
- LinuxContainers.org — https://linuxcontainers.org/ (LXC, images `images:`)
- Proxmox VE Wiki — https://pve.proxmox.com/wiki/ (LXC, `pct`, `vzdump`)

**Sujets d'approfondissement :**

1. **API REST LXD** : tout ce que fait la CLI passe par l'API — scriptez en
   Python/Go (`lxc query` pour explorer).
2. **OVN** : réseaux overlay multi-hôtes avec routage distribué et ACLs.
3. **Projets LXD** : multi-tenant, quotas, RBAC fin.
4. **Terraform provider LXD/Proxmox** : infra as code pour vos conteneurs.
5. **Ansible** : module `community.general.lxd_container`, inventaire dynamique.
6. **Ceph + LXD/Proxmox** : stockage distribué pour clusters.
7. **Confidential computing / Kata** : quand le noyau partagé ne suffit plus.

**Routines d'équipe suggérées :**

- Revue mensuelle : images à jour ? snapshots qui s'accumulent ? quotas OK ?
- Test trimestriel de restauration de sauvegarde (section 61).
- Mise à jour des golden images après chaque patch Tuesday / release Ubuntu.
- Ce guide est vivant : ajoutez-y vos propres erreurs (section 80) au fil de l'eau.

---

*Fin du guide — bon courage avec vos conteneurs !* 🚀
