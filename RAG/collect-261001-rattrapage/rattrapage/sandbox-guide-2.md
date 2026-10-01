---
id: collect-261001-rattrapage/rattrapage/sandbox-guide-2
title: "Guide complet du sandboxing sous Linux"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["sandbox", "decode", "memory"]
source: docs/RAG/collect-261001-rattrapage/sandbox_guide.md
source_anchor: ""
source_lines: [185, 389]
sha256: 4e7923cf665816d3a022a9ed6f842b06182b89e1359d18a4f66a43ea8121b3d2
---

# Guide complet du sandboxing sous Linux

```bash
# Vérifier que cgroups v2 est actif
mount | grep cgroup2
# cgroup2 on /sys/fs/cgroup type cgroup2 (...)

# Créer un cgroup pour un programme suspect
sudo mkdir /sys/fs/cgroup/sandbox.d/

# Limites : 512 Mo de RAM, 50% d'un CPU, 100 processus max
echo 536870912      | sudo tee /sys/fs/cgroup/sandbox.d/memory.max
echo 50000           | sudo tee /sys/fs/cgroup/sandbox.d/cpu.max
# (format cpu.max : quota_us period_us ; 50000 100000 = 50%)
echo 100             | sudo tee /sys/fs/cgroup/sandbox.d/pids.max

# Lancer un programme dans ce cgroup
sudo systemd-run --unit=sandbox-test --scope \
  -p MemoryMax=512M -p CPUQuota=50% -p TasksMax=100 \
  /usr/bin/programme-suspect --option

# Surveiller
systemd-cgtop
cat /sys/fs/cgroup/sandbox.d/memory.current
cat /sys/fs/cgroup/sandbox.d/pids.current
```

Avec systemd, c'est encore plus simple : les directives `MemoryMax=`, `CPUQuota=`, `TasksMax=` dans une unité font exactement cela (voir section 44).

## 9. chroot vs vrai sandbox : la différence fondamentale

`chroot` ne change que la racine apparente du système de fichiers. **Ce n'est pas un mécanisme de sécurité** :

| Propriété | chroot | Vrai sandbox (namespaces) |
|---|---|---|
| Isole le filesystem | Partiellement (racine redéfinie) | Oui (mnt ns + pivot_root) |
| Isole les processus | Non (`ps` voit tout) | Oui (pid ns) |
| Isole le réseau | Non | Oui (net ns) |
| Isole les UID | Non | Oui (user ns) |
| Évasion possible | Oui, triviale si root (voir §15) | Difficile (nécessite faille noyau) |
| Limite les ressources | Non | Oui (cgroups) |
| Filtre les appels système | Non | Oui (seccomp) |

Un processus root dans un chroot peut s'en échapper en quelques lignes de C (double chroot + `chdir`). Retenez : **chroot = commodité d'empaquetage et de build, jamais une barrière de sécurité à lui seul**.

## 10. Seccomp : filtrer les appels système

Seccomp-bpf permet de filtrer les appels système qu'un processus peut faire. C'est le dernier rempart : même si un attaquant contrôle le processus, il ne peut pas appeler `mount()`, `ptrace()`, `reboot()`, etc.

```bash
# Voir le filtre seccomp d'un processus (1 = actif)
grep Seccomp /proc/1/status
# Seccomp: 2  (2 = SECCOMP_MODE_FILTER)

# systemd peut appliquer un filtre par défaut à un service :
# SystemCallFilter=@system-service
# SystemCallArchitectures=native
```

Docker applique un profil seccomp par défaut qui bloque ~44 appels système dangereux (`mount`, `ptrace`, `reboot`, `swapon`, `kexec_load`, ...). Firejail fait de même avec `seccomp`. C'est invisible à l'usage, mais redoutablement efficace.

## 11. Capabilities : découper le pouvoir de root

Le "root" monolithique est découpé en ~40 capabilities. Un sandbox retire celles qui sont inutiles :

```bash
# Voir les capabilities d'un processus
grep Cap /proc/1/status
capsh --decode=00000000a80425fb

# Lister les capabilities d'un binaire fichier
getcap /usr/bin/ping
# /usr/bin/ping cap_net_raw=ep
```

Les plus dangereuses à retirer en priorité dans un sandbox : `CAP_SYS_ADMIN` (équivaut presque à root), `CAP_SYS_PTRACE`, `CAP_NET_ADMIN`, `CAP_SYS_MODULE`, `CAP_DAC_OVERRIDE`.

```bash
# Lancer un programme sans aucune capability (nobody-like)
sudo -u nobody /usr/bin/programme

# Avec capsh : garder uniquement cap_net_bind_service
sudo capsh --drop=all --addamb=cap_net_bind_service \
  --user=nobody -- -c "/usr/bin/mon-serveur"
```

Docker : `--cap-drop=ALL --cap-add=NET_BIND_SERVICE`. Systemd : `CapabilityBoundingSet=` et `AmbientCapabilities=` (voir §46).

## 12. Cartographie des outils : qui fait quoi

| Outil | Namespaces | Seccomp | AppArmor/SELinux | cgroups | Cas d'usage typique |
|---|---|---|---|---|---|
| chroot | Non | Non | Non | Non | Build, rescue |
| Firejail | Oui | Oui | Oui (optionnel) | Non | Applis desktop non fiables |
| systemd (directives) | Oui | Oui | Oui (optionnel) | Oui | Services système |
| bubblewrap (`bwrap`) | Oui | Oui | Non | Non | Brique de base (Flatpak) |
| Docker | Oui | Oui | Oui | Oui | Applis serveur conteneurisées |
| AppArmor seul | Non | Non | Oui | Non | Confinement par profil |
| QEMU/KVM | N/A (VM) | N/A | Oui (sVirt) | Oui | Isolation forte (noyau séparé) |

**Lecture :** pour un poste de travail (PDF suspect, navigateur), Firejail est le meilleur rapport simplicité/efficacité. Pour un service, les directives systemd suffisent souvent. Pour une isolation forte (code totalement hostile), une VM KVM reste la référence.

---

# PARTIE II — CHROOT : COMPRENDRE ET UTILISER SANS ILLUSION

## 13. Créer un chroot Debian minimal avec debootstrap

`debootstrap` construit une arborescence Debian dans un répertoire. C'est l'outil standard.

```bash
# Installation
sudo apt update && sudo apt install -y debootstrap schroot

# Créer un chroot Debian stable (bookworm) dans /srv/chroot/build
sudo mkdir -p /srv/chroot/build
sudo debootstrap --variant=minbase bookworm /srv/chroot/build \
  http://deb.debian.org/debian/

# Variante Ubuntu
sudo debootstrap --variant=minbase noble /srv/chroot/ubuntu-noble \
  http://archive.ubuntu.com/ubuntu/
```

Options utiles :

```bash
# Inclure des paquets dès la création
sudo debootstrap --variant=minbase \
  --include=build-essential,git,ca-certificates \
  bookworm /srv/chroot/build http://deb.debian.org/debian/

# Architecture différente (nécessite qemu-user-static)
sudo apt install -y qemu-user-static
sudo debootstrap --foreign --arch=arm64 bookworm /srv/chroot/arm64 \
  http://deb.debian.org/debian/
sudo chroot /srv/chroot/arm64 /debootstrap/debootstrap --second-stage
```

Taille indicative : un `minbase` bookworm pèse ~250 Mo.

## 14. Entrer dans le chroot et le configurer

```bash
# Monter les pseudo-filesystems indispensables
sudo mount --bind /dev  /srv/chroot/build/dev
sudo mount --bind /dev/pts /srv/chroot/build/dev/pts
sudo mount -t proc proc /srv/chroot/build/proc
sudo mount -t sysfs sys /srv/chroot/build/sys

# Copier la résolution DNS (sinon pas de réseau sortant en noms)
sudo cp /etc/resolv.conf /srv/chroot/build/etc/resolv.conf

# Entrer
sudo chroot /srv/chroot/build /bin/bash

# Dans le chroot : configuration minimale
apt update
apt install -y locales
echo "en_US.UTF-8 UTF-8" >> /etc/locale.gen && locale-gen
apt install -y build-essential
exit
```

Script d'aide `/usr/local/bin/chroot-enter` :

```bash
#!/bin/bash
# Usage : chroot-enter /srv/chroot/build [commande...]
CHROOT="$1"; shift
for m in dev dev/pts proc sys; do
  mountpoint -q "$CHROOT/$m" || {
    case $m in
      dev)     mount --bind /dev "$CHROOT/dev" ;;
      dev/pts) mount --bind /dev/pts "$CHROOT/dev/pts" ;;
      proc)    mount -t proc proc "$CHROOT/proc" ;;
      sys)     mount -t sysfs sys "$CHROOT/sys" ;;
    esac
  }
done
[ -f "$CHROOT/etc/resolv.conf" ] || cp /etc/resolv.conf "$CHROOT/etc/resolv.conf"
exec chroot "$CHROOT" "$@"
```

## 15. Pourquoi chroot n'est PAS une sécurité : la démonstration

Un processus **root** dans un chroot peut s'échapper. Voici le principe (ne l'exécutez que dans un labo) :

```c
/* Évasion chroot classique : le "double chroot" ne suffit pas seul ;
   la méthode fiable passe par open() d'un descripteur hors chroot
   avant le chroot, puis fchdir(). */
#include <unistd.h>
#include <fcntl.h>
#include <stdio.h>
int main(void) {
    int fd = open("/", O_RDONLY);   /* descripteur vers la VRAIE racine */
    chroot("/srv/chroot/build");    /* on "entre" dans le chroot */
    fchdir(fd);                     /* ... et on en ressort aussitôt */
    chroot(".");                    /* on redéfinit la racine ici */
    execl("/bin/bash", "bash", NULL);
    return 0;
}
```

Conclusion opérationnelle :

