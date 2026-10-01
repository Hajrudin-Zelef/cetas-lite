---
id: collect-261001-rattrapage/rattrapage/sandbox-guide-5
title: "Guide complet du sandboxing sous Linux"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["sandbox", "agent", "attention"]
source: docs/RAG/collect-261001-rattrapage/sandbox_guide.md
source_anchor: ""
source_lines: [845, 1046]
sha256: 6463f684f0442765f140c45e19ab88d9323bbdb7c3607a9eee662b2bba6463d7
---

# Guide complet du sandboxing sous Linux

Sous Wayland, l'isolation est native : chaque application ne voit que ses propres fenêtres. **Si vous utilisez Firejail pour des applis graphiques sensibles, préférez une session Wayland.**

## 32. Supervision : voir ce que fait un sandbox

```bash
# Processus et namespaces
firejail --list
firejail --tree

# Événements d'audit du sandbox (tentatives bloquées)
firejail --audit firefox
# Exemple de sortie : syscall bloqué, accès fichier refusé

# Journal en temps réel d'un sandbox nommé
firejail --name=test --noprofile bash
# dans un autre terminal :
firejail --join=test
```

Activer la journalisation des violations seccomp :

```bash
# Les processus tués par seccomp apparaissent dans le journal noyau
sudo dmesg | grep -i seccomp
sudo journalctl -k | grep -i "audit.*seccomp"
```

## 33. Dépannage Firejail : cas concrets

**Cas 1 — L'application ne démarre pas : écran noir / segfault immédiat**
Cause fréquente : `private-bin` oublie un binaire nécessaire, ou `seccomp` bloque un appel légitime.
Diagnostic :

```bash
# Lancer sans seccomp pour isoler la cause
firejail --profile=votre.profile --ignore=seccomp programme
# Si ça marche : le filtre seccomp est en cause -> ajoutez
# seccomp.keep=mmap,mmap2,openat,... (approche par allowlist ciblée)
# Lancer avec strace dans le sandbox pour voir l'appel bloqué
firejail --noprofile --seccomp strace -f -e trace=%network programme
```

**Cas 2 — "cannot create /run/firejail/..." ou erreur SUID**
Le binaire firejail doit être SUID ou les user namespaces activés :

```bash
ls -l /usr/bin/firejail
sysctl kernel.unprivileged_userns_clone   # doit être 1
```

**Cas 3 — Pas de son dans le navigateur sandboxé**
`nosound` ou `private-dev` bloque l'accès à PulseAudio/PipeWire. Retirez `nosound` et ajoutez le socket :

```ini
# pour PipeWire/PulseAudio
whitelist ${RUNUSER}/pulse/native
whitelist ${RUNUSER}/pipewire-0
```

**Cas 4 — Téléchargements invisibles**
Avec `private=` ou `whitelist`, le répertoire `~/Téléchargements` de l'hôte n'est pas le même que dans le sandbox. Utilisez `--whitelist=~/Téléchargements` explicite ou copiez les fichiers après fermeture.

## 34. Erreurs classiques avec Firejail

1. **Lancer en root** : `sudo firejail firefox` casse l'isolation (le sandbox hérite des privilèges). Toujours lancer en utilisateur normal.
2. **Empiler les profils sans tester** : un profil trop strict casse l'appli, l'utilisateur le désactive entièrement. Durcissez progressivement.
3. **Oublier `firecfg` après mise à jour** : relancez `sudo firecfg` après chaque mise à jour de firejail pour régénérer les liens.
4. **Croire que `--private` protège les données** : `--private` donne un home *vide*, il ne chiffre rien. Les fichiers restent accessibles hors sandbox.
5. **Négliger X11** : sans `x11 xephyr` (ou Wayland), un keylogger sandboxé lit quand même le clavier sous X11.

## 35. Checklist Firejail pour un poste

- [ ] `firejail` et `firejail-profiles` installés, version ≥ 0.9.70
- [ ] `kernel.unprivileged_userns_clone = 1`
- [ ] `sudo firecfg` exécuté (lanceurs intégrés)
- [ ] Profils personnalisés pour : navigateur, lecteur PDF, client mail
- [ ] Procédure "PDF suspect" connue et testée
- [ ] Session Wayland préférée (sinon `x11 xephyr` pour les applis à risque)
- [ ] Test mensuel : `firejail --list` et revue des profils après mises à jour

---

# PARTIE IV — SANDBOXING SYSTEMD : CONFINER LES SERVICES

## 36. Pourquoi sandboxer avec systemd ?

Sur un serveur, la majorité des processus sont des services systemd. Plutôt qu'ajouter un outil tiers, systemd intègre nativement : namespaces, seccomp, capabilities, cgroups, montages read-only. Avantage décisif : **zéro dépendance supplémentaire**, configuration déclarative, audit via `systemd-analyze security`.

```bash
# Vérifier la version (>= 247 pour un bon support, >= 252 idéal)
systemctl --version

# Voir le niveau de confinement actuel d'un service
systemd-analyze security nginx.service
```

## 37. ProtectSystem : verrouiller le système de fichiers

| Valeur | Effet |
|---|---|
| `no` (défaut) | Aucune protection |
| `yes` / `true` | `/usr` et `/boot` en lecture seule, `/etc` en lecture seule |
| `full` | Comme `yes` + `/var`, `/srv`, `/home` en lecture seule |
| `strict` | Tout le filesystem en lecture seule + `/proc`, `/sys`, `/dev` masqués |

```ini
[Service]
# Service web typique : binaires et config en lecture seule
ProtectSystem=strict
ReadWritePaths=/var/lib/monapp /var/log/monapp
# /tmp et /var/tmp privés (namespaces mnt)
PrivateTmp=true
```

`ReadWritePaths=` perce des exceptions dans le read-only. `ReadOnlyPaths=` fait l'inverse (rend un chemin read-only avec `ProtectSystem=no`).

## 38. ProtectHome : protéger les données utilisateurs

| Valeur | Effet |
|---|---|
| `no` (défaut) | Accès normal |
| `read-only` | `/home`, `/root`, `/run/user` en lecture seule |
| `tmpfs` | Masqués par des tmpfs vides (invisibles) |
| `yes` | Comme `tmpfs` (alias historique) |

```ini
[Service]
# Un service n'a JAMAIS besoin de lire /home
ProtectHome=tmpfs
# Exception ciblée si nécessaire :
# ReadWritePaths=/home/monapp/data
```

**Réflexe :** sauf besoin explicite, mettez `ProtectHome=tmpfs` sur tous vos services. C'est la directive au meilleur ratio protection/effort.

## 39. PrivateTmp, PrivateDevices, PrivateNetwork

```ini
[Service]
# /tmp et /var/tmp privés au service (invisibles des autres processus)
PrivateTmp=true

# /dev minimal : null, zero, random, urandom, full, tty, ptmx...
# Bloque l'accès aux disques (/dev/sda), au hardware brut
PrivateDevices=true

# Namespace réseau vide (que lo). Pour un agent qui n'a pas besoin du réseau.
PrivateNetwork=true

# Combinaison typique pour un worker de traitement local
PrivateTmp=true
PrivateDevices=true
PrivateNetwork=true
```

Attention : `PrivateNetwork=true` casse la résolution DNS et toute connexion. Pour un service qui doit sortir mais pas écouter en local, préférez `RestrictAddressFamilies` (§43).

## 40. NoNewPrivileges : interdire l'élévation

```ini
[Service]
# Le processus et ses enfants ne pourront JAMAIS gagner de privilèges
# (setuid, setgid, capabilities fichier, sudo... deviennent inopérants)
NoNewPrivileges=true
```

C'est la directive la plus importante après `ProtectSystem`/`ProtectHome`. Elle neutralise toute une classe d'exploitations (binaire setuid vulnérable, `sudo` mal configuré invoqué depuis le service). **Effet de bord :** les programmes qui ont *légitimement* besoin de setuid (ex. `ping` avec capabilities fichier) cesseront de fonctionner — prévoyez des capabilities explicites (§46).

## 41. ProtectKernelTunables, ProtectKernelModules, ProtectControlGroups

```ini
[Service]
# /proc/sys, /sys en lecture seule : impossible de modifier les paramètres noyau
ProtectKernelTunables=true

# Interdit le chargement/déchargement de modules noyau
ProtectKernelModules=true

# /sys/fs/cgroup en lecture seule : le service ne peut pas modifier ses limites
# ni s'échapper de son cgroup
ProtectControlGroups=true

# Masque /proc/kallsyms, /proc/kcore... (adresses noyau utiles aux exploits)
ProtectKernelLogs=true
```

Ces quatre directives sont sans risque pour 95 % des services applicatifs (web, bases de données, workers). Mettez-les par défaut.

## 42. SystemCallFilter : seccomp pour les services

```ini
[Service]
# Allowlist : le service ne peut appeler que les syscalls du groupe @system-service
SystemCallFilter=@system-service
# Bloque aussi les architectures 32 bits (x32, i386) : réduit la surface
SystemCallArchitectures=native
# Action en cas de violation : tuer le processus (défaut) ou errno
SystemCallErrorNumber=EPERM
```

Groupes prédéfinis utiles (`man systemd.exec` pour la liste complète) :

