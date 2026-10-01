---
id: collect-261001-rattrapage/rattrapage/sandbox-guide-11
title: "Guide complet du sandboxing sous Linux"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["sandbox", "arr", "memory"]
source: docs/RAG/collect-261001-rattrapage/sandbox_guide.md
source_anchor: ""
source_lines: [2162, 2298]
sha256: 11a914b07846ca109cb04bf4d65fa48a47297887b2d2ae2c1ef55da87b43f5a4
---

# Guide complet du sandboxing sous Linux

```bash
# État des sandboxes en une page (à mettre en cron hebdo / rapport)
{
  echo "=== $(date) — $(hostname) ==="
  echo "--- AppArmor ---"; sudo aa-status | grep -E "profiles are in|profiles are loaded"
  echo "--- Firejail ---"; firejail --list | wc -l | xargs echo "sandboxes actifs:"
  echo "--- Services systemd les plus exposés ---"
  for s in $(systemctl list-units --type=service --state=running --no-legend | awk '{print $1}' | head -20); do
    score=$(systemd-analyze security "$s" 2>/dev/null | grep "Overall exposure" | awk '{print $(NF-1)}')
    echo "$score $s"
  done | sort -n | tail -5
  echo "--- Docker : conteneurs privileged/root ---"
  docker ps -q | xargs -r docker inspect --format '{{.Name}} user={{.Config.User}} privileged={{.HostConfig.Privileged}}'
} > /var/log/sandbox-report.txt
```

---

# PARTIE XI — DÉPANNAGE GÉNÉRAL

## 87. Matrice de diagnostic : "ça ne marche plus depuis le confinement"

| Symptôme | Cause probable | Vérification | Correctif |
|---|---|---|---|
| Service tué au démarrage, `status=31/SYS` | `SystemCallFilter` trop strict | `journalctl -u X` | Élargir le filtre ou `SystemCallErrorNumber=EPERM` pour diagnostiquer |
| `Permission denied` sur un fichier légitime | `ProtectSystem=strict` / AppArmor | `journalctl -k \| grep DENIED` | `ReadWritePaths=` / règle AppArmor |
| Plus de DNS / réseau | `PrivateNetwork` / `net none` | `ip a` dans le contexte | Retirer ou ajuster `RestrictAddressFamilies` |
| L'appli graphique ne se lance pas | `private-bin` incomplet, D-Bus coupé | Lancer sans profil, comparer | Ajouter binaires/abstractions |
| Écriture impossible dans `/tmp` | `PrivateTmp` + `read-only` | `touch /tmp/test` | Ajouter `--tmpfs /tmp` ciblé |
| Conteneur sans son/vidéo | devices masqués | `ls /dev` dans le conteneur | `--device /dev/snd` ciblé |
| Build qui échoue (dépendances) | Pas de réseau dans le sandbox | logs du build | Autoriser le réseau pendant le build, couper à l'exécution |
| `sudo` ne marche plus dans le service | `NoNewPrivileges=true` | `sudo -n true` → échec | Ne pas utiliser sudo dans le service (revoir l'archi) |

## 88. Procédure de diagnostic en 5 minutes

```bash
# 1. Le service est-il tué par le confinement ?
systemctl status monservice.service
journalctl -u monservice.service --since "10 min ago" | tail -30

# 2. SIGSYS (31) = seccomp. DENIED = AppArmor. OOM = cgroup mémoire.
journalctl -k --since "10 min ago" | grep -iE "seccomp|apparmor.*denied|oom-killer" | tail -20

# 3. Bissection : désactiver UNE directive à la fois via drop-in
sudo systemctl edit monservice.service   # commenter la directive suspecte
sudo systemctl daemon-reload && sudo systemctl restart monservice.service

# 4. Tracer les appels bloqués (dernier recours, en recette)
sudo strace -f -e trace=%network,%process -p $(pidof monservice) 2>&1 | tail -40

# 5. Valider le correctif avec le score
systemd-analyze security monservice.service | grep "Overall exposure"
```

## 89. 10+ erreurs classiques (récapitulatif)

1. **Confondre chroot et sandbox** : chroot seul n'arrête ni root, ni le réseau, ni `ps`. Voir §15.
2. **Lancer Firejail en root** (`sudo firejail ...`) : hérite des privilèges, isolation illusoire.
3. **`--privileged` Docker "temporaire"** qui devient permanent : auditez les `docker run` avec un script.
4. **Désactiver AppArmor/SELinux au lieu de diagnostiquer** : on supprime l'alarme incendie au lieu d'éteindre le feu.
5. **Profils trop stricts, jamais testés** : l'utilisateur contourne tout. Durcir progressivement.
6. **Oublier les user namespaces** (`kernel.unprivileged_userns_clone=0`) : Firejail/podman cassés après un durcissement noyau trop agressif.
7. **`ProtectSystem=strict` sans `ReadWritePaths`** : le service ne peut plus écrire ses logs/pidfiles → crash silencieux.
8. **Aucune limite de ressources** : fork bomb ou fuite mémoire = DoS de l'hôte. Toujours `TasksMax` + `MemoryMax`.
9. **Secrets dans l'image Docker / le profil** : `docker inspect` et `/etc/firejail` sont lisibles. Secrets = fichiers montés ou gestionnaire dédié.
10. **Ne pas mettre à jour les sandboxes** : vieux Firejail, image de 2 ans, chroot non patché = sandbox vulnérable lui-même.
11. **X11 sans `xephyr`** : keylogger sandboxé sous X11 = clavier espionné quand même. Wayland ou `x11 xephyr`.
12. **Faire confiance aux artefacts d'un build sandboxé** : le sandbox protège l'hôte, pas la qualité du binaire. Scanner avant déploiement.

## 90. Checklist de mise en production d'un service confiné

- [ ] Unité systemd avec au minimum : `ProtectSystem=strict`, `ProtectHome=tmpfs`, `PrivateTmp=true`, `NoNewPrivileges=true`
- [ ] `systemd-analyze security` ≤ 3.0 (EXPOSED) ou justification écrite des ✗ restants
- [ ] Utilisateur dédié (`User=`), jamais root sauf besoin justifié (ports < 1024 → `AmbientCapabilities`)
- [ ] Limites : `MemoryMax`, `TasksMax`, `CPUQuota` dimensionnées
- [ ] Profil AppArmor en `enforce` si le service est exposé au réseau
- [ ] Logs d'audit collectés (journald/auditd) et alertes configurées (§85)
- [ ] Procédure de rollback testée (ancienne image / ancien binaire)
- [ ] Test d'intrusion léger : que se passe-t-il si le service est compromis ? (lecture `/etc/shadow` ? réseau ? persistance ?)
- [ ] Documentation : pourquoi chaque exception (`ReadWritePaths`, `cap_add`) existe

---

# PARTIE XII — PENSE-BÊTE DE POCHE

## 91. Pense-bête : une page à imprimer

```
=== SANDBOX LINUX — PENSE-BÊTE ===
NAMESPACES
  lsns -t net,mnt,pid,user          # qui est isolé ?
  sudo unshare --pid --mount --uts --fork --mount-proc bash
  sudo ip netns add sb && sudo ip netns exec sb ip a
  unshare --user --map-root-user bash   # faux root

CGROUPS v2
  systemd-run --scope -p MemoryMax=512M -p CPUQuota=50% -p TasksMax=100 CMD

FIREJAIL
  firejail firefox                   # profil auto
  firejail --noprofile --net=none --seccomp --caps.drop=all --noroot CMD
  firejail --list | --tree | --top   # supervision
  firejail --shutdown=<pid|name>
  sudo firecfg                       # intégration bureau (après chaque MAJ)

SYSTEMD (dans [Service])
  ProtectSystem=strict | ProtectHome=tmpfs | PrivateTmp=true
  PrivateDevices=true | NoNewPrivileges=true
  ProtectKernelTunables=true | ProtectKernelModules=true | ProtectControlGroups=true
  SystemCallFilter=@system-service | SystemCallArchitectures=native
  RestrictAddressFamilies=AF_UNIX AF_INET AF_INET6
  RestrictNamespaces=yes | RestrictRealtime=true | MemoryDenyWriteExecute=true
  CapabilityBoundingSet=  (+ AmbientCapabilities=CAP_NET_BIND_SERVICE si besoin)
  MemoryMax=1G | CPUQuota=50% | TasksMax=100
  systemd-analyze security SERVICE   # viser ≤ 3.0

APPARMOR
  sudo aa-status | aa-complain BIN | aa-enforce BIN
  sudo aa-genprof BIN   # apprentissage (utiliser l'appli à fond)
  sudo aa-logprof       # intégrer les refus légitimes
  sudo journalctl -k | grep "apparmor.*DENIED"

DOCKER DURCI
  docker run --read-only --tmpfs /tmp --cap-drop=ALL --cap-add=NET_BIND_SERVICE \
    --security-opt no-new-privileges --network none --user 10000:10000 \
    --pids-limit 100 --memory 512m IMAGE
  INTERDIT: --privileged, docker.sock monté, :latest en prod, root par défaut

DIAGNOSTIC EXPRESS
  status=31/SYS  -> seccomp trop strict (SystemCallFilter)
  DENIED         -> AppArmor (aa-logprof)
  oom-killer     -> MemoryMax trop bas ou fuite
  Permission denied fichier -> ProtectSystem / ReadWritePaths
```

## 92. Glossaire

