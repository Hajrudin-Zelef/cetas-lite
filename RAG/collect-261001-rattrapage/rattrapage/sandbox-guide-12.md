---
id: collect-261001-rattrapage/rattrapage/sandbox-guide-12
title: "Guide complet du sandboxing sous Linux"
domain: rattrapage
role: reference
task: reference
actors: ["Google"]
dates: ["2026-09-26"]
keywords: ["sandbox", "benchmarks", "incident"]
source: docs/RAG/collect-261001-rattrapage/sandbox_guide.md
source_anchor: ""
source_lines: [2299, 2425]
sha256: 8f68f0d2016d702f302a4aa1437fff208e26eb6bc2eae7d66d164c2360f05072
---

# Guide complet du sandboxing sous Linux

| Terme | Définition |
|---|---|
| **Namespace** | Vue isolée d'une ressource noyau (processus, réseau, montages, UID...) |
| **cgroup** | Groupe de contrôle : limite et mesure les ressources (CPU, RAM, pids, I/O) |
| **chroot** | Redéfinit la racine `/` apparente d'un processus (pas une sécurité) |
| **Seccomp-bpf** | Filtre les appels système autorisés pour un processus |
| **Capability** | Morceau du pouvoir root (ex. `CAP_NET_BIND_SERVICE`) attribuable séparément |
| **MAC** | Mandatory Access Control : politique obligatoire (AppArmor, SELinux), au-dessus des permissions Unix (DAC) |
| **DAC** | Discretionary Access Control : les classiques `rwx` Unix |
| **Profil AppArmor** | Liste des accès autorisés pour un programme donné |
| **Contexte SELinux** | Étiquette `user:role:type:level` collée sur chaque objet |
| **User namespace** | Namespace qui mappe root interne vers un UID non privilégié externe |
| **pivot_root** | Change la racine d'un namespace de montage (base des conteneurs) |
| **Bubblewrap (bwrap)** | Outil minimal de sandbox par namespaces, brique de Flatpak |
| **nspawn** | `systemd-nspawn` : chroot amélioré avec namespaces + cgroups |
| **SIGSYS (31)** | Signal envoyé quand seccomp tue un processus |
| **W^X** | Write XOR Execute : une page mémoire ne peut être à la fois inscriptible et exécutable |
| **Fork bomb** | Programme qui se réplique jusqu'à épuisement des processus (`:(){ :|:& };:`) |
| **Évasion (escape)** | Sortie du sandbox par l'attaquant (faille noyau, mauvaise config) |
| **Défense en profondeur** | Empilement de couches de sécurité indépendantes |
| **Moindre privilège** | Ne donner que les droits strictement nécessaires |
| **Drop-in systemd** | Fichier `/etc/systemd/system/X.d/*.conf` qui surcharge une unité sans la modifier |
| **tmpfs** | Système de fichiers en RAM, vide au montage, jeté au démontage |

## 93. Quiz : 10 questions pour valider

**Q1.** Un processus root s'échappe-t-il facilement d'un chroot classique ?
**Q2.** Quelle directive systemd rend `/usr` et `/boot` (entre autres) en lecture seule pour un service ?
**Q3.** Que signifie un service tué avec `status=31/SYS` dans `journalctl` ?
**Q4.** Citez 3 namespaces Linux et ce qu'ils isolent.
**Q5.** Pourquoi faut-il éviter `sudo firejail firefox` ?
**Q6.** Quelle option Docker donne au conteneur un accès quasi total à l'hôte et doit être bannie ?
**Q7.** AppArmor : quelle est la différence entre les modes `complain` et `enforce` ?
**Q8.** Quel fichier SELinux restaure-t-on avec `restorecon` après avoir déplacé un fichier web ?
**Q9.** Que protège `NoNewPrivileges=true` ?
**Q10.** Sous X11, pourquoi un programme sandboxé avec Firejail peut-il quand même espionner le clavier, et quelle est la parade ?

<details>
<summary><b>Réponses</b></summary>

**R1.** Oui, trivialement (ex. descripteur ouvert sur `/` avant `chroot` + `fchdir`, voir §15). Le chroot n'est pas une barrière de sécurité contre root.
**R2.** `ProtectSystem=strict` (ou `yes`/`full` selon le niveau). Avec `ReadWritePaths=` pour les exceptions.
**R3.** Le processus a été tué par seccomp (SIGSYS = signal 31) : `SystemCallFilter` bloque un appel système que le programme tente d'utiliser.
**R4.** Exemples : `pid` (vue des processus), `net` (interfaces, routes, ports), `mnt` (arborescence des montages), `user` (mapping UID/GID), `uts` (hostname), `ipc` (mémoire partagée).
**R5.** Parce que le sandbox hérite alors des privilèges root : l'isolation (user namespace, `noroot`) devient illusoire et une évasion donne root sur l'hôte.
**R6.** `--privileged` : désactive seccomp/AppArmor, donne accès aux devices de l'hôte. À remplacer par `--device` / `--cap-add` ciblés.
**R7.** `complain` : journalise les violations sans bloquer (mode apprentissage). `enforce` : bloque et journalise.
**R8.** Le **contexte SELinux** (type) du fichier : `mv` conserve l'ancien type, d'où des refus même avec les bonnes permissions Unix. `restorecon` réapplique le type prévu par la politique.
**R9.** Contre toute élévation de privilèges du processus et de ses enfants : setuid/setgid, capabilities fichier et `sudo` deviennent inopérants.
**R10.** Parce que le protocole X11 permet à tout client de lire les événements clavier/souris et de capturer l'écran. Parades : `firejail --x11=xephyr` (serveur X imbriqué) ou passer à Wayland (isolation native par fenêtre).
</details>

## 94. Pour aller plus loin

**Documentation officielle :**
- `man systemd.exec` — LA référence des directives de sandboxing (très complète)
- `man firejail-profile` — toutes les directives de profils Firejail
- `man apparmor.d` — syntaxe des profils AppArmor
- Namespaces : `man 7 namespaces`, `man 1 unshare`, `man 8 ip-netns`
- cgroups v2 : `Documentation/admin-guide/cgroup-v2.rst` du noyau

**Outils à explorer ensuite :**
- **Bubblewrap (`bwrap`)** : la brique minimale (`flatpak` repose dessus) — idéal pour scripter ses propres sandboxes
- **gVisor (Google)** : sandbox userspace qui intercepte les syscalls (isolation proche d'une VM pour des conteneurs)
- **Kata Containers** : conteneurs dans des micro-VM (compromis Docker/KVM)
- **Qubes OS** : poste de travail où chaque appli tourne dans sa VM (paranoïa assumée)
- **Trivy / Grype** : scan de vulnérabilités des images

**Lectures :**
- "Linux Containers in 500 lines of Go" — comprendre les namespaces en les codant
- La documentation seccomp-bpf (`man 2 seccomp`)
- CIS Benchmarks (Debian, Docker) : checklists de durcissement auditables

**Exercices pratiques :**
1. Prenez un de vos services et faites passer son score `systemd-analyze security` sous 3.0.
2. Écrivez un profil AppArmor en `complain` pour un outil interne, passez en `enforce` après une semaine.
3. Montez un pipeline "PDF suspect" (§79) et testez-le avec un PDF EICAR-like inoffensif.
4. Auditez vos `docker run` / compose avec la checklist §66.

## 95. Modèle de fiche service confiné (à remplir par l'équipe)

```markdown
# Fiche confinement — <nom du service>
- Responsable :
- Date de mise en confinement :
- Score systemd-analyze security : ___ (objectif ≤ 3.0)
- Utilisateur dédié : ___
- Directives appliquées : ProtectSystem=___ ProtectHome=___ ...
- Exceptions justifiées :
  - ReadWritePaths=___ parce que ___
  - cap_add ___ parce que ___
- Profil AppArmor : ___ (enforce/complain)
- Limites ressources : MemoryMax=___ TasksMax=___ CPUQuota=___
- Alertes configurées : ___
- Test d'intrusion léger (date + résultat) : ___
- Prochaine revue : ___
```

## 96. Anti-sèche sysctl durcissement (rappel §77)

```bash
# /etc/sysctl.d/10-durcissement.conf — appliquer avec sysctl --system
kernel.yama.ptrace_scope = 1          # ptrace restreint
kernel.kexec_load_disabled = 1        # pas de kexec
kernel.dmesg_restrict = 1             # dmesg pour root uniquement
kernel.unprivileged_bpf_disabled = 2  # eBPF non privilégié désactivé
net.core.bpf_jit_harden = 2          # durcit le JIT eBPF
kernel.perf_event_paranoid = 3        # perf restreint
# kernel.unprivileged_userns_clone = 1  # GARDER à 1 si Firejail/podman utilisés !
```

## 97. Ce que ce guide ne couvre pas (périmètre honnête)

- L'analyse de malware avancée (sandbox dynamique type Cuckoo, INetSim) : autre métier.
- Le durcissement du noyau par compilation (options `CONFIG_*`, grsecurity).
- La sécurité des orchestrateurs (Kubernetes : PodSecurity, NetworkPolicies, OPA/Gatekeeper).
- Le chiffrement (LUKS, TLS) : complémentaire, pas remplacé par le sandboxing.
- La réponse à incident (forensique) : le sandbox limite l'impact, il ne remplace pas un plan de réponse.

## 98. Historique des révisions du guide

| Date | Version | Changement |
|---|---|---|
| 2026-09-26 | 1.0 | Création : 8 parties, 98 sections, quiz, pense-bête |

## 99. Licence et réutilisation interne

