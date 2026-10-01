---
id: collect-261001-rattrapage/rattrapage/sandbox-guide-9
title: "Guide complet du sandboxing sous Linux"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["sandbox", "exploit", "memory"]
source: docs/RAG/collect-261001-rattrapage/sandbox_guide.md
source_anchor: ""
source_lines: [1796, 1965]
sha256: 5563b12daa7424c7ca11ef70657f2278913f4bdc851f2e1c5125cdb71090c150
---

# Guide complet du sandboxing sous Linux

```bash
sudo apt install -y podman
podman run --rm -it --userns=auto debian:bookworm-slim bash
# Chaque conteneur mappe ses UID vers une plage non privilégiée de l'hôte
# (/etc/subuid, /etc/subgid)
```

Avec les user namespaces, "root" dans le conteneur = UID banal sur l'hôte. Même en cas d'évasion du conteneur, l'attaquant n'a aucun privilège.

## 71. Analyser une image avant de l'exécuter

```bash
# Lister les couches et leur taille
docker history monimage:1.2.3

# Scanner les vulnérabilités (Trivy)
trivy image monimage:1.2.3

# Inspecter sans exécuter : extraire le filesystem
docker export $(docker create monimage:1.2.3) | tar -t | less
mkdir /tmp/img && tar -xf <(docker export $(docker create monimage:1.2.3)) -C /tmp/img
grep -r "password\|secret\|BEGIN PRIVATE KEY" /tmp/img/etc /tmp/img/app 2>/dev/null | head

# Vérifier l'utilisateur par défaut
docker inspect monimage:1.2.3 --format '{{.Config.User}}'
# vide = root ! -> prévoyez --user au run
```

## 72. Cas pratique : exécuter un outil non fiable dans Docker

Scénario : un utilitaire de conversion téléchargé hors dépôt, image inconnue.

```bash
#!/bin/bash
# run-douteux.sh — exécute un conteneur "jetable" très confiné
set -euo pipefail
IMAGE="$1"; shift
WORK=$(mktemp -d)
cp -a ./entree/. "$WORK"/

docker run --rm -i \
  --read-only --tmpfs /tmp:noexec,nosuid \
  --cap-drop=ALL --security-opt no-new-privileges \
  --network none \                 # pas de réseau du tout
  --user 10000:10000 \
  --pids-limit 50 --memory 256m --memory-swap 256m \
  --volume "$WORK":/work:rw \
  --workdir /work \
  "$IMAGE" "$@"

echo "--- Résultat dans $WORK/sortie (à inspecter avant usage) ---"
ls -l "$WORK"
```

`--network none` + `--read-only` + `--user` + `--cap-drop=ALL` : l'outil peut au pire corrompre `/work`, qui est un répertoire temporaire.

## 73. Supervision des conteneurs : logs d'audit

```bash
# Logs d'un conteneur
docker logs --since 1h monapp-durci
docker logs -f --tail 100 monapp-durci

# Événements Docker (création, mort, OOM...)
docker events --since 1h --filter event=die --filter event=oom

# Un conteneur tué par seccomp : status 31/SYS dans les events
docker events --filter event=die | grep -i "31"

# Refus AppArmor liés à Docker
sudo journalctl -k | grep "apparmor.*DENIED.*docker"

# Centraliser : driver json-file + logrotate, ou syslog/journald
# /etc/docker/daemon.json
{
  "log-driver": "journald",
  "log-opts": { "tag": "docker/{{.Name}}" }
}
```

Alertes à mettre en place (voir §80) : conteneur en boucle de restart, OOM kills répétés, refus seccomp/AppArmor, image `:latest` déployée.

## 74. Erreurs classiques avec Docker

1. **`--privileged` "pour que ça marche"** : donne accès aux devices de l'hôte, désactive seccomp/AppArmor. Cherchez l'option précise (`--device`, `--cap-add`) au lieu du marteau.
2. **Monter `docker.sock`** dans un conteneur (Portainer mal configuré, CI) : le conteneur peut lancer des conteneurs privilégiés → contrôle de l'hôte.
3. **Tourner en root dans le conteneur** par défaut : combinez `USER` dans le Dockerfile et `--user` au run.
4. **`:latest` en production** : déploiements non reproductibles, mises à jour surprises. Épinglez `image:tag@sha256:digest`.
5. **Secrets dans les variables d'environnement** : visibles via `docker inspect`. Utilisez des fichiers (`_FILE`) ou un gestionnaire de secrets.
6. **Aucune limite de ressources** : un conteneur qui fuit = OOM de l'hôte. Toujours `mem_limit` + `pids_limit`.

---

# PARTIE VIII — BONNES PRATIQUES ET DÉFENSE EN PROFONDEUR

## 75. Principe du moindre privilège : méthode d'application

Pour chaque programme/service, posez-vous ces 6 questions et traduisez en directives :

| Question | Si non | Traduction technique |
|---|---|---|
| A-t-il besoin de tout le filesystem ? | Non | `ProtectSystem=strict` / `--read-only` |
| A-t-il besoin de `/home` ? | Non | `ProtectHome=tmpfs` |
| A-t-il besoin du réseau ? | Non | `PrivateNetwork=true` / `--network none` |
| A-t-il besoin de root/capabilities ? | Non | `NoNewPrivileges=true`, `CapabilityBoundingSet=` vide |
| A-t-il besoin de tous les syscalls ? | Non | `SystemCallFilter=@system-service` |
| A-t-il besoin de ressources illimitées ? | Non | `MemoryMax=`, `TasksMax=`, `CPUQuota=` |

**Méthode itérative :** appliquez tout, testez, et ne relâchez que ce qui casse un usage légitime — en documentant pourquoi dans un commentaire de l'unité.

## 76. Défense en profondeur : empiler les couches

Aucune couche n'est infaillible ; leur combinaison rend l'exploitation exponentiellement plus difficile :

```
┌─────────────────────────────────────────────┐
│  Couche 7 : Supervision (logs, alertes)     │  <- détecte l'échec des autres
├─────────────────────────────────────────────┤
│  Couche 6 : MAC (AppArmor/SELinux)          │  <- confine par profil
├─────────────────────────────────────────────┤
│  Couche 5 : Seccomp (filtre syscalls)       │  <- bloque mount, ptrace...
├─────────────────────────────────────────────┤
│  Couche 4 : Capabilities (retrait)          │  <- pas de CAP_SYS_ADMIN
├─────────────────────────────────────────────┤
│  Couche 3 : Namespaces (pid, net, mnt...)   │  <- vue isolée
├─────────────────────────────────────────────┤
│  Couche 2 : cgroups (mémoire, CPU, pids)    │  <- pas d'épuisement
├─────────────────────────────────────────────┤
│  Couche 1 : Utilisateur dédié, FS read-only │  <- base saine
└─────────────────────────────────────────────┘
        Noyau à jour + mises à jour auto
```

Exemple concret (service web) : `User=monapp` + `ProtectSystem=strict` + `NoNewPrivileges` + `SystemCallFilter` + profil AppArmor + `MemoryMax` + alertes sur les DENIED. Un attaquant doit enchaîner : exploit appli → évasion namespace → contournement seccomp → faille noyau → évasion AppArmor, le tout en silence.

## 77. Durcissement de l'hôte : prérequis au sandboxing

Un sandbox sur un hôte pourri ne sert à rien :

```bash
# 1. Mises à jour automatiques de sécurité
sudo apt install -y unattended-upgrades
sudo dpkg-reconfigure -plow unattended-upgrades

# 2. Noyau : user namespaces OK, mais restreindre si inutile
sysctl kernel.unprivileged_userns_clone   # 1 nécessaire pour Firejail/podman

# 3. ptrace : restreindre aux processus du même utilisateur
# /etc/sysctl.d/10-ptrace.conf
kernel.yama.ptrace_scope = 1

# 4. kexec / hibernation : désactiver si inutile (anti-persistance firmware-like)
# /etc/sysctl.d/10-kexec.conf
kernel.kexec_load_disabled = 1

# 5. dmesg restreint aux root (cache les adresses noyau)
kernel.dmesg_restrict = 1

# 6. Appliquer
sudo sysctl --system
```

```bash
# 7. Audit du système : lynis (rapport)
sudo apt install -y lynis
sudo lynis audit system --quick | grep -E "Warning|Suggestion" | head -20
```

## 78. Mises à jour des sandboxes eux-mêmes

