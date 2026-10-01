---
id: collect-261001-rattrapage/rattrapage/sandbox-guide-6
title: "Guide complet du sandboxing sous Linux"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["sandbox", "agent"]
source: docs/RAG/collect-261001-rattrapage/sandbox_guide.md
source_anchor: ""
source_lines: [1047, 1323]
sha256: 2e37f75207ebe3186a851dab7e95124c0c9d2cc44c1a2479e3cca02ceaee8cfc
---

# Guide complet du sandboxing sous Linux

| Groupe | Contenu indicatif |
|---|---|
| `@system-service` | Base pour un service système standard |
| `@basic-io` | Lecture/écriture fichiers, ioctl de base |
| `@network-io` | Sockets, connect, send/recv |
| `@file-system` | open, stat, mkdir... |
| `@privileged` | mount, sethostname... (à éviter) |
| `@resources` | nice, setrlimit... |
| `@debug` | ptrace, process_vm_readv (à éviter) |

Approche pragmatique : commencez par `@system-service`, testez, et resserrez avec `~@mount @debug` (syntaxe d'exclusion) si besoin :

```ini
SystemCallFilter=@system-service @network-io
SystemCallFilter=~@privileged ~@debug ~@resources
```

**Piège :** un filtre trop strict tue le service au démarrage (souvent pendant l'init, ex. un appel `name_to_handle_at` inhabituel). Diagnostic : `journalctl -u monservice` montre `code=killed, status=31/SYS` (31 = SIGSYS, le signal seccomp).

## 43. RestrictAddressFamilies : contrôler le réseau au niveau socket

```ini
[Service]
# Le service ne peut créer que des sockets TCP/UDP/Unix
RestrictAddressFamilies=AF_UNIX AF_INET AF_INET6
# Version radicale : aucun socket du tout
# RestrictAddressFamilies=no
```

Exemples par profil :

```ini
# Serveur web : TCP + sockets unix (php-fpm, etc.)
RestrictAddressFamilies=AF_UNIX AF_INET AF_INET6

# Worker local sans réseau
RestrictAddressFamilies=AF_UNIX

# Agent SNMP qui a besoin de raw sockets : ajoutez AF_PACKET avec prudence
# RestrictAddressFamilies=AF_UNIX AF_INET AF_PACKET
```

Combiné à `PrivateNetwork=true`, c'est redondant ; combiné à un réseau normal, c'est un filet anti-exfiltration par canaux exotiques (ex. `AF_PACKET` pour sniffer).

## 44. Limiter les ressources : cgroups via systemd

```ini
[Service]
# Mémoire : le service est tué (OOM) au-delà de 1 Go
MemoryMax=1G
# Variante souple : pression avant kill
# MemoryHigh=800M

# CPU : 50% d'un cœur max
CPUQuota=50%

# Nombre de tâches (processus+threads) : anti fork-bomb
TasksMax=100

# I/O disque : 10 Mo/s en lecture, 5 Mo/s en écriture
IOReadBandwidthMax=/dev/sda 10M
IOWriteBandwidthMax=/dev/sda 5M
```

Vérification :

```bash
systemctl show monservice.service -p MemoryMax,CPUQuota,TasksMax
systemd-cgtop
```

## 45. RestrictNamespaces, RestrictRealtime, MemoryDenyWriteExecute

```ini
[Service]
# Interdit au service de créer des namespaces (anti-évasion vers user ns)
RestrictNamespaces=yes
# Variante fine : autorise seulement user et mnt
# RestrictNamespaces=user mnt

# Interdit le temps réel (anti DoS CPU)
RestrictRealtime=true

# Interdit les pages mémoire à la fois inscriptibles et exécutables (W^X)
# Bloque une partie des exploits (shellcode), peut casser les JIT (Java, Node)
MemoryDenyWriteExecute=true

# Interdit les namespaces user spécifiquement (durcit encore)
RestrictNamespaces=~user
```

## 46. Capabilities : le minimum vital

```ini
[Service]
# Le service ne garde AUCUNE capability...
CapabilityBoundingSet=
# ... sauf celles listées ici (exemple : binder un port < 1024)
AmbientCapabilities=CAP_NET_BIND_SERVICE
# Alternative moderne au setuid root pour les ports privilégiés
```

Tableau de décision :

| Besoin | Directive |
|---|---|
| Écouter sur le port 80/443 | `AmbientCapabilities=CAP_NET_BIND_SERVICE` (+ `authbind` en alternative) |
| Envoyer des pings | `CAP_NET_RAW` |
| Changer l'heure | `CAP_SYS_TIME` (rare, préférez chrony) |
| Rien de spécial | `CapabilityBoundingSet=` vide |

## 47. Exemple complet : unité durcie pour une application web

```ini
# /etc/systemd/system/monapp.service
[Unit]
Description=Mon application web interne
After=network-online.target
Wants=network-online.target

[Service]
Type=simple
User=monapp
Group=monapp
WorkingDirectory=/opt/monapp
ExecStart=/opt/monapp/bin/monapp --config /etc/monapp/config.yaml
Restart=on-failure
RestartSec=5

# --- Filesystem ---
ProtectSystem=strict
ProtectHome=tmpfs
PrivateTmp=true
PrivateDevices=true
ReadWritePaths=/var/lib/monapp /var/log/monapp
ProtectKernelTunables=true
ProtectKernelModules=true
ProtectControlGroups=true

# --- Privilèges ---
NoNewPrivileges=true
CapabilityBoundingSet=
RestrictNamespaces=yes

# --- Syscalls ---
SystemCallFilter=@system-service
SystemCallFilter=~@privileged ~@debug
SystemCallArchitectures=native

# --- Réseau ---
RestrictAddressFamilies=AF_UNIX AF_INET AF_INET6

# --- Ressources ---
MemoryMax=1G
CPUQuota=100%
TasksMax=200

# --- Divers ---
RestrictRealtime=true
LockPersonality=true
UMask=0027

[Install]
WantedBy=multi-user.target
```

```bash
sudo systemctl daemon-reload
sudo systemctl enable --now monapp.service
systemd-analyze security monapp.service   # viser <= 3.0 (EXPOSED)
```

## 48. systemd-analyze security : lire le score

```bash
systemd-analyze security monapp.service
```

Sortie typique :

```
  NAME                                        DESCRIPTION
✗ PrivateNetwork=                             Service has access to the host's network
✗ User=/DynamicUser=                          Service runs as root user
✗ CapabilityBoundingSet=~CAP_SYS_ADMIN        Service may have elevated capabilities
✓ NoNewPrivileges=                           Service cannot gain new privileges
...
→ Overall exposure level for monapp.service: 7.5 UNSAFE 🛑
```

Échelle :

| Score | Niveau | Interprétation |
|---|---|---|
| 0 – 1.5 | `SAFE` | Très bien confiné |
| 1.6 – 3.0 | `EXPOSED` | Correct, revoir les ✗ restants |
| 3.1 – 8.0 | `UNSAFE` | Confinement insuffisant |
| 8.1 – 10 | `DANGEROUS` | Aucune protection (défaut) |

**Méthode :** partez du score, corrigez les ✗ un par un en testant le service à chaque étape. Un service qui ne démarre plus après un changement = directive trop stricte → assouplissez *cette* directive uniquement.

```bash
# Comparer avant/après
systemd-analyze security nginx.service > /tmp/avant.txt
# ... modifiez l'unité ...
systemd-analyze security nginx.service > /tmp/apres.txt
diff /tmp/avant.txt /tmp/apres.txt
```

## 49. Cas concret : durcir nginx pas à pas

État initial : `nginx.service` fourni par Debian, score ~9.6 DANGEROUS (tourne en root pour le master).

```bash
# 1. Créer un drop-in (ne jamais modifier le fichier du paquet)
sudo systemctl edit nginx.service
```

```ini
# /etc/systemd/system/nginx.service.d/durcissement.conf
[Service]
# nginx master a besoin de root pour binder le 443 et lire les certs :
# on garde User=root mais on verrouille le reste
ProtectSystem=strict
ReadWritePaths=/var/lib/nginx /var/log/nginx /run
ProtectHome=tmpfs
PrivateTmp=true
PrivateDevices=true
NoNewPrivileges=true
ProtectKernelTunables=true
ProtectKernelModules=true
ProtectControlGroups=true
RestrictAddressFamilies=AF_UNIX AF_INET AF_INET6 AF_NETLINK
SystemCallFilter=@system-service
SystemCallFilter=~@privileged
SystemCallArchitectures=native
RestrictNamespaces=yes
RestrictRealtime=true
MemoryDenyWriteExecute=true
```

```bash
sudo systemctl daemon-reload
sudo systemctl restart nginx.service
systemd-analyze security nginx.service
# Objectif : passer de 9.6 DANGEROUS à ~2.5 EXPOSED
curl -sI https://localhost | head -3   # vérifier que ça sert toujours
```

Note : `AF_NETLINK` est nécessaire à nginx pour la résolution et les logs. Si `MemoryDenyWriteExecute` casse un module tiers (ex. ModSecurity avec JIT PCRE), retirez-le.

## 50. Cas concret : service utilisateur (user service) sandboxé

Les services `--user` bénéficient des mêmes directives :

```ini
# ~/.config/systemd/user/syncthing-perso.service
[Unit]
Description=Syncthing personnel confiné

[Service]
ExecStart=/usr/bin/syncthing -no-browser
Restart=on-failure
ProtectSystem=strict
ProtectHome=read-only
ReadWritePaths=%h/Sync %h/.config/syncthing
PrivateTmp=true
NoNewPrivileges=true
SystemCallFilter=@system-service
RestrictAddressFamilies=AF_UNIX AF_INET AF_INET6
MemoryMax=512M
TasksMax=100

[Install]
WantedBy=default.target
```

