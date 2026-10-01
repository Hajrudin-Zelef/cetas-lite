---
id: collect-261001-rattrapage/rattrapage/sandbox-guide-3
title: "Guide complet du sandboxing sous Linux"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["sandbox"]
source: docs/RAG/collect-261001-rattrapage/sandbox_guide.md
source_anchor: ""
source_lines: [390, 608]
sha256: 28b537e16fcacd60177d0bba57b9f4d3147c5f22fe6ec8079bbd4473a4b3f036
---

# Guide complet du sandboxing sous Linux

1. Ne lancez **jamais** de programme non fiable en root dans un chroot en croyant être protégé.
2. Si vous devez utiliser chroot, combinez-le avec : utilisateur dédié (` nobody`), `NoNewPrivileges`, montages read-only, et idéalement des namespaces par-dessus (`unshare --map-root-user`).
3. Pour une vraie isolation, préférez `systemd-nspawn` (voir §16) ou `bwrap`.

## 16. systemd-nspawn : le chroot qui devient un conteneur

`systemd-nspawn` est le chaînon manquant entre chroot et Docker : il ajoute namespaces (pid, net, mnt, uts, ipc), cgroups et un vrai PID 1.

```bash
# Lancer le chroot comme un conteneur léger
sudo systemd-nspawn -D /srv/chroot/build

# Avec réseau privé (veth vers l'hôte)
sudo systemd-nspawn -D /srv/chroot/build --network-veth

# Sans privilèges : user namespace
sudo systemd-nspawn -D /srv/chroot/build --private-users=pick

# Lancer une commande unique et quitter
sudo systemd-nspawn -D /srv/chroot/build --as-pid2 \
  /usr/bin/make -C /src

# Lier un répertoire de l'hôte en lecture seule
sudo systemd-nspawn -D /srv/chroot/build \
  --bind-ro=/home/zelef/sources:/src \
  /usr/bin/make -C /src
```

Fichier machine persistant `/etc/systemd/nspawn/build.nspawn` :

```ini
[Exec]
Boot=no
PrivateUsers=yes

[Files]
BindReadOnly=/home/zelef/sources:/src
TemporaryFileSystem=/:ro
Bind=/srv/chroot/build/var/tmp

[Network]
Private=yes
VirtualEthernet=yes
```

Puis `sudo systemd-nspawn -M build /usr/bin/make -C /src`. Les logs partent dans le journal : `journalctl -M build`.

## 17. Cas pratique : compiler du code tiers dans un chroot jetable

Scénario : vous devez compiler un projet téléchargé depuis un dépôt obscur. Le `Makefile` pourrait contenir des commandes malveillantes.

```bash
#!/bin/bash
# build-isole.sh — compile dans un nspawn jetable
set -euo pipefail
SRC="$1"                       # répertoire des sources (hôte)
CHROOT=/srv/chroot/build
MACHINE="build-$$"

# 1. Copie des sources dans un répertoire dédié (pas de bind direct)
WORK=$(mktemp -d)
cp -a "$SRC"/. "$WORK"/
chmod -R a-w "$WORK"           # sources en lecture seule

# 2. Compilation confinée : pas de réseau, user namespace, / en ro
sudo systemd-nspawn --machine="$MACHINE" -D "$CHROOT" \
  --private-users=pick \
  --network-veth --private-network \
  --bind-ro="$WORK:/src" \
  --tmpfs=/tmp --tmpfs=/build \
  --as-pid2 \
  bash -c 'cp -a /src/. /build/ && cd /build && ./configure && make -j"$(nproc)"'

# 3. Récupération des artefacts (on ne fait confiance qu'aux binaires, pas au build)
mkdir -p ./artefacts
sudo cp -a "$WORK"/artefacts/. ./artefacts/ 2>/dev/null || true
rm -rf "$WORK"
echo "OK — vérifiez les artefacts avant usage."
```

Points de vigilance : `--private-network` coupe le réseau (un build qui télécharge des dépendances échouera — c'est voulu) ; les artefacts restent suspects jusqu'à inspection (`strings`, antivirus, test en sandbox).

## 18. Mettre à jour et maintenir un chroot

```bash
# Mise à jour depuis l'hôte sans entrer dedans
sudo chroot /srv/chroot/build apt update
sudo chroot /srv/chroot/build apt -y upgrade
sudo chroot /srv/chroot/build apt -y autoremove --purge
sudo chroot /srv/chroot/build apt clean

# Taille et contenu
sudo du -sh /srv/chroot/build
dpkg-query --admindir=/srv/chroot/build/var/lib/dpkg -W | wc -l

# Supprimer proprement un chroot
sudo umount -R /srv/chroot/build 2>/dev/null
sudo rm -rf /srv/chroot/build
```

Automatisez la mise à jour hebdomadaire avec un timer systemd (voir §44 pour le sandboxing du service lui-même) :

```ini
# /etc/systemd/system/chroot-update.service
[Unit]
Description=Mise à jour du chroot de build

[Service]
Type=oneshot
ExecStart=/usr/bin/chroot /srv/chroot/build /usr/bin/apt-get update
ExecStart=/usr/bin/chroot /srv/chroot/build /usr/bin/apt-get -y upgrade
```

## 19. schroot : gérer plusieurs chroots proprement

`schroot` gère sessions, utilisateurs autorisés et snapshots LVM.

```ini
# /etc/schroot/chroot.d/build.conf
[build]
description=Chroot de compilation Debian
type=directory
directory=/srv/chroot/build
users=zelef
groups=sbuild
root-groups=root
personality=linux
preserve-environment=true
```

```bash
# Entrer comme utilisateur normal (pas root !)
schroot -c build -u zelef

# Session nommée pour un build long
schroot -b -c build -n session-build-1
schroot -r -c session-build-1 -u zelef -- make -C /src
schroot -e -c session-build-1
```

## 20. Limites du chroot : tableau récapitulatif opérationnel

| Risque | chroot seul | chroot + user non-root | nspawn |
|---|---|---|---|
| Lecture `/etc/shadow` hôte | Possible (si monté) | Possible si permissions | Non (fs isolé) |
| Évasion par root | Triviale | N/A (pas root) | Très difficile |
| Accès réseau hôte | Total | Total | Filtrable |
| `ps` voit l'hôte | Oui | Oui | Non |
| Fork bomb | Affecte l'hôte | Affecte l'hôte | Contenue (cgroups) |
| Persistance via cron hôte | Possible | Possible | Non |

**Décision :** chroot = environnement de build reproductible ; nspawn = exécution isolée ; ni l'un ni l'autre ne remplace un audit du code exécuté.

## 21. Checklist chroot avant usage en production

- [ ] Le chroot est-il à jour (`apt upgrade` récent) ?
- [ ] Les montages `/dev`, `/proc`, `/sys` sont-ils démontés après usage ?
- [ ] Aucun secret de l'hôte n'est bind-monté (clés SSH, tokens) ?
- [ ] Le processus s'exécute-t-il en non-root à l'intérieur ?
- [ ] Le réseau est-il nécessaire ? Si non, coupez-le (nspawn `--private-network`).
- [ ] Les artefacts produits sont-ils vérifiés avant réutilisation ?
- [ ] Le chroot est-il reconstruit régulièrement depuis une base saine ?

## 22. Erreurs classiques avec chroot

**Erreur 1 — `chroot: failed to run command '/bin/bash': No such file or directory`**
Cause : le chroot est vide ou l'architecture ne correspond pas (binaire ARM sur hôte x86 sans qemu). Vérifiez `ls /srv/chroot/build/bin/bash` et `file`.

**Erreur 2 — `apt update` échoue : "Temporary failure resolving ..."**
Cause : `/etc/resolv.conf` absent dans le chroot. Copiez celui de l'hôte (voir §14).

**Erreur 3 — Oublier de démonter avant `rm -rf`**
`rm -rf` sur un chroot dont `/dev` est bind-monté peut endommager l'hôte. Toujours `umount -R` d'abord, ou utilisez `schroot -e`.

**Erreur 4 — Croire que `chroot` + mot de passe = sécurité**
Le chroot n'authentifie rien et ne confine rien. Voir §15.

---

# PARTIE III — FIREJAIL : LE SANDBOX DU POSTE DE TRAVAIL

## 23. Firejail : principe et installation

Firejail est un sandbox SUID qui utilise les namespaces Linux, seccomp-bpf et les capabilities pour confiner n'importe quel programme, sans configuration préalable. C'est l'outil le plus rentable pour un poste de travail.

```bash
# Debian / Ubuntu
sudo apt update && sudo apt install -y firejail firejail-profiles

# Vérifier la version (>= 0.9.70 recommandé)
firejail --version

# Vérifier que les user namespaces sont autorisés
sysctl kernel.unprivileged_userns_clone
```

Test immédiat, sans config :

```bash
# Firefox confiné avec le profil par défaut
firejail firefox

# Voir ce que le sandbox bloque
firejail --noprofile --seccomp --private-tmp bash
```

Firejail s'appuie sur `kernel.unprivileged_userns_clone=1` (défaut sur Debian/Ubuntu). Si votre noyau le désactive, Firejail bascule sur son binaire SUID — moins propre, à éviter si possible.

## 24. Profils existants : l'inventaire

Firejail livre plus de 1000 profils dans `/etc/firejail/` :

```bash
ls /etc/firejail/*.profile | wc -l
ls /etc/firejail/ | grep -E "firefox|chrom|thunderbird|evince|okular|vlc|libreoffice"
```

Profils les plus utiles au quotidien :

