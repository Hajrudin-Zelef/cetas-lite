---
id: collect-261001-rattrapage/rattrapage/sandbox-guide-1
title: "Guide complet du sandboxing sous Linux"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["sandbox", "attention", "memory"]
source: docs/RAG/collect-261001-rattrapage/sandbox_guide.md
source_anchor: ""
source_lines: [1, 184]
sha256: e9ee050ca0f308f3f3d5cba275a6113e36c323a93cde742eefc571f54fd0985e
---

# Guide complet du sandboxing sous Linux

**Isolation, confinement et sécurité pratique pour sysadmin**
*Public : Zelef, chef de service systèmes & énergies — approche opérationnelle, exemples testés sur Debian/Ubuntu.*

> Ce guide couvre les mécanismes d'isolation Linux (namespaces, cgroups v2, chroot), les outils de confinement (Firejail, systemd, AppArmor, SELinux, Docker) et les cas d'usage concrets : ouvrir un PDF suspect, durcir un navigateur, compiler du code tiers, isoler une application non fiable. Chaque section contient des exemples de code prêts à l'emploi.

---

## 1. Pourquoi sandboxer ? Le problème de départ

Un système Linux classique repose sur la confiance : tout processus lancé par un utilisateur hérite de ses droits et peut, en théorie, lire ses fichiers, ouvrir des sockets réseau, ou exploiter une faille du noyau. Le sandboxing inverse cette logique : **on part du principe que le programme est hostile ou compromis**, et on lui retire tout ce dont il n'a pas strictement besoin.

Cas concrets rencontrés en production :

- Un PDF reçu par mail exploite une faille du lecteur PDF → sans sandbox, l'attaquant lit `~/.ssh/id_rsa`.
- Un paquet `.deb` téléchargé hors dépôt exécute un script `postinst` malveillant → sans confinement, il installe un service persistant.
- Un outil de compilation tiers (`npm install`, `pip install`, `make`) exécute du code arbitraire pendant le build.
- Un navigateur web exécute du JavaScript hostile en permanence.

Le sandbox ne supprime pas les vulnérabilités : il **réduit l'impact** de leur exploitation. C'est la différence entre "la machine est compromise" et "un processus isolé a planté".

## 2. Modèle de menace : à quoi se préparer

Avant de choisir un outil, définissez ce que vous craignez :

| Menace | Exemple | Contre-mesure principale |
|---|---|---|
| Lecture de données sensibles | Malware lisant `~/.ssh/`, `~/Documents` | Namespaces `mnt`, `ProtectHome` |
| Exfiltration réseau | Reverse shell vers l'attaquant | Namespace `net`, `RestrictAddressFamilies` |
| Persistance | Installation d'un cron, d'un service systemd | Filesystem read-only, `ProtectSystem=strict` |
| Élévation de privilèges | Exploitation setuid, `sudo` | `NoNewPrivileges`, user namespace |
| Épuisement de ressources | Fork bomb, remplissage disque | cgroups (pids, memory, cpu) |
| Évasion du sandbox | Exploitation d'une faille noyau | Seccomp, AppArmor, noyau à jour |

**Règle d'or :** un sandbox n'est jamais parfait. On empile les couches (défense en profondeur) : namespace + seccomp + AppArmor + cgroups + filesystem read-only.

## 3. Les namespaces Linux : la brique de base

Un namespace est une vue isolée d'une ressource système. Deux processus dans des namespaces différents voient des ressources différentes. Le noyau Linux en propose 8 :

| Namespace | Flag `clone()` | Isole | Depuis le noyau |
|---|---|---|---|
| PID | `CLONE_NEWPID` | Les identifiants de processus | 2.6.24 |
| NET | `CLONE_NEWNET` | Interfaces, routes, iptables, ports | 2.6.24 |
| MNT | `CLONE_NEWNS` | Les points de montage (arborescence `/`) | 2.4.19 |
| UTS | `CLONE_NEWUTS` | Hostname et nom de domaine NIS | 2.6.19 |
| IPC | `CLONE_NEWIPC` | Mémoire partagée, sémaphores, files de messages | 2.6.19 |
| USER | `CLONE_NEWUSER` | Les UID/GID (root mappé vers un UID non privilégié) | 3.8 |
| CGROUP | `CLONE_NEWCGROUP` | La vue des cgroups | 4.6 |
| TIME | `CLONE_NEWTIME` | L'horloge (monotonic, boottime) | 5.6 |

Vérifier les namespaces d'un processus :

```bash
# Lister les namespaces du processus 1 (init)
ls -l /proc/1/ns/

# Lister les namespaces du shell courant
lsns

# Version lisible : quel processus est dans quel namespace
lsns -t net,mnt,pid,user
```

Exemple de sortie `lsns` :

```
        NS TYPE   NPROCS   PID USER COMMAND
4026531836 pid       187     1 root /sbin/init
4026531837 user      187     1 root /sbin/init
4026531840 mnt       186     1 root /sbin/init
4026531992 net         2  3456 toto /usr/bin/firefox
```

Ici, le processus 3456 (Firefox) est dans un namespace réseau dédié (`4026531992`) : il ne voit pas les interfaces de l'hôte.

## 4. Namespace PID : isoler la vue des processus

Dans un namespace PID, le premier processus devient le PID 1 de ce namespace, et `ps` ne montre que les processus du namespace.

```bash
# Créer un shell dans de nouveaux namespaces pid + mnt + uts
sudo unshare --pid --mount --uts --fork --mount-proc /bin/bash

# Dans ce shell :
ps aux        # ne montre que les processus du namespace
hostname test # le hostname de l'hôte n'est pas modifié
```

Points importants :

- `--mount-proc` est nécessaire pour remonter un `/proc` cohérent avec le nouveau namespace PID.
- Le PID 1 du namespace a un rôle spécial : il adopte les processus orphelins. Si votre PID 1 meurt, tout le namespace meurt.
- `unshare` sans `--fork` ne fonctionne pas pour `pid` : le namespace PID ne s'applique qu'aux processus enfants.

## 5. Namespace NET : isoler le réseau

Un namespace réseau vierge ne contient que l'interface `lo` (éteinte). On peut y créer des paires `veth` pour le relier à l'hôte.

```bash
# Créer un namespace réseau nommé "sandbox"
sudo ip netns add sandbox

# Le lister
ip netns list

# Exécuter une commande dedans
sudo ip netns exec sandbox ip addr
# -> seulement lo, DOWN

# Activer lo et tester
sudo ip netns exec sandbox ip link set lo up
sudo ip netns exec sandbox ping -c 2 127.0.0.1

# Relier à l'hôte avec une paire veth
sudo ip link add veth-host type veth peer name veth-sb
sudo ip link set veth-sb netns sandbox
sudo ip addr add 10.200.0.1/24 dev veth-host
sudo ip link set veth-host up
sudo ip netns exec sandbox ip addr add 10.200.0.2/24 dev veth-sb
sudo ip netns exec sandbox ip link set veth-sb up
sudo ip netns exec sandbox ip route add default via 10.200.0.1

# Tester la connectivité
sudo ip netns exec sandbox ping -c 2 10.200.0.1

# Nettoyer
sudo ip netns del sandbox
sudo ip link del veth-host
```

C'est exactement ce que font Docker et Firejail sous le capot pour isoler le réseau d'un conteneur.

## 6. Namespace MNT : isoler le système de fichiers

Le namespace `mnt` permet de monter/démonter des systèmes de fichiers sans affecter l'hôte. Combiné à `pivot_root` ou `chroot`, c'est la base des conteneurs.

```bash
sudo unshare --mount --fork /bin/bash

# Dans le shell isolé :
mount -t tmpfs tmpfs /tmp     # /tmp de l'hôte masqué, sans l'affecter
mount -o remount,ro /          # racine en lecture seule (dans ce namespace)
mount | grep " / "             # vérifier
exit                           # l'hôte n'a rien vu
```

Attention : sans précaution, les montages se propagent (montages "shared" par défaut avec systemd). Pour un vrai sandbox, passez les montages en `private` :

```bash
mount --make-rprivate /
```

## 7. Namespace USER : le faux root

Le namespace `user` est le plus puissant : il permet à un utilisateur non privilégié d'avoir l'UID 0 **à l'intérieur** du namespace, sans aucun privilège sur l'hôte.

```bash
# Créer un user namespace : uid 0 dedans = votre uid dehors
unshare --user --map-root-user /bin/bash

id -u      # affiche 0 !
touch /test_root  # échoue : pas de droit réel sur l'hôte
```

```bash
# Mapping manuel : uid 0-999 dedans -> 100000-100999 dehors
unshare --user --map-users=0:100000:1000 --map-groups=0:100000:1000 /bin/bash
```

C'est ce mécanisme qui permet à Docker "rootless" et à `podman` de fonctionner sans daemon root. Limite : certaines distributions désactivent les user namespaces non privilégiés (`kernel.unprivileged_userns_clone=0` sur Debian par défaut à 1, heureusement).

Vérifier :

```bash
sysctl kernel.unprivileged_userns_clone
# 1 = autorisé (nécessaire pour Firejail, podman, bubblewrap)
```

## 8. cgroups v2 : limiter les ressources

Les namespaces isolent la **vue**, les cgroups limitent les **ressources** : CPU, mémoire, nombre de processus, I/O disque.

