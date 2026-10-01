---
id: collect-261001-rattrapage/rattrapage/lxc-guide-3
title: "LXC & LXD — Guide ultra-complet"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["attention", "memory"]
source: docs/RAG/collect-261001-rattrapage/lxc_guide.md
source_anchor: ""
source_lines: [498, 793]
sha256: 62698b223e41128a3433679ce89d2c0902695496d04f0b5cf2b4d1247b123366
---

# LXC & LXD — Guide ultra-complet

```bash
# Commande simple (non interactive)
lxc exec c1 -- ls -la /root

# Avec variables d'environnement
lxc exec c1 --env MODE=prod -- ./deploy.sh

# Mode interactif : shell
lxc exec c1 -- bash

# En tant qu'utilisateur non-root
lxc exec c1 --user 1000 --group 1000 -- whoami
lxc exec c1 -- su - deploy -c "whoami"
```

**Subtilités :**

- Le `--` sépare les options `lxc` de la commande à exécuter (obligatoire
  si la commande commence par `-`).
- Sans TTY alloué, utilisez `--mode non-interactive` pour forcer (utile
  dans les scripts/Ansible).
- `lxc exec` passe par l'API LXD : fonctionne même si le réseau du
  conteneur est cassé (gros avantage pour le dépannage !).

**Équivalent Proxmox :** `pct exec <vmid> -- <commande>` (section 71).

---

## 16. `lxc shell` — Accès shell simplifié

```bash
# Ouvre un shell (bash si dispo, sinon sh)
lxc shell c1

# Équivaut à :
lxc exec c1 -- bash
```

**Choisir le shell :**

```bash
lxc exec c1 -- sh      # minimaliste, toujours présent
lxc exec c1 -- bash    # confort
lxc exec c1 -- zsh     # si installé dans le conteneur
```

**Astuce :** définissez un alias shell pour vos conteneurs fréquents :

```bash
# ~/.bashrc
alias c1sh='lxc exec c1 -- bash'
alias c1root='lxc exec c1 -- su -'
```

---

## 17. `lxc info` / `lxc list` — Inspection

```bash
# Vue d'ensemble tabulaire
lxc list

# Colonnes personnalisées
lxc list -c n,s,4,P,profiles
# n=name s=state 4=ipv4 P=pid profiles=profils

# Format CSV (pour scripts)
lxc list -f csv -c n,s

# Format JSON (pour automatisation)
lxc list -f json | jq '.[].name'

# Détails d'un conteneur
lxc info c1

# Détails + état étendu (processus, mémoire...)
lxc info c1 --show-log
```

**Filtrer :**

```bash
lxc list status=running          # seulement les démarrés
lxc list "web.*"                 # par motif de nom
lxc list -c n --fast             # rapide (sans résolution IP)
```

---

## 18. Fichiers — `lxc file push/pull/edit`

```bash
# Envoyer un fichier vers le conteneur
lxc file push ./nginx.conf c1/etc/nginx/nginx.conf

# Avec propriétaire et permissions
lxc file push --uid 0 --gid 0 --mode 0644 ./app.conf c1/etc/app/app.conf

# Récupérer un fichier depuis le conteneur
lxc file pull c1/etc/nginx/nginx.conf ./nginx-c1.conf

# Récursif (dossier)
lxc file push -r ./site/ c1/var/www/html/ --create-dirs
lxc file pull -r c1/var/log/nginx/ ./logs-c1/

# Éditer directement (ouvre $EDITOR)
lxc file edit c1/etc/hosts

# Créer un dossier
lxc file push --create-dirs /dev/null c1/opt/app/placeholder
```

**Cas d'usage typiques :**

- Déployer une config sans SSH ni Ansible pour un petit parc.
- Récupérer des logs d'un conteneur au réseau cassé.
- **Équivalent Proxmox :** `pct push` / `pct pull` (section 71).

---

## 19. Configuration — `lxc config`

Toute la configuration d'un conteneur est une paire clé/valeur :

```bash
# Voir toute la config (fusionnée : profils + locale)
lxc config show c1 --expanded

# Voir seulement la config locale (sans l'héritage des profils)
lxc config show c1

# Définir une valeur
lxc config set c1 limits.memory 2GiB

# Supprimer une valeur (retour à l'héritage)
lxc config unset c1 limits.memory

# Éditer en YAML dans $EDITOR
lxc config edit c1

# Métadonnées
lxc config set c1 user.commentaire "Serveur web production"
lxc config get c1 user.commentaire
```

**Clés les plus utilisées :**

| Clé | Exemple | Effet |
|---|---|---|
| `limits.cpu` | `2` ou `0-3` | CPUs alloués |
| `limits.memory` | `4GiB` | RAM max |
| `limits.memory.swap` | `false` | Interdire le swap |
| `boot.autostart` | `true` | Démarrage auto avec l'hôte |
| `boot.autostart.delay` | `30` | Délai avant démarrage (s) |
| `boot.autostart.priority` | `10` | Ordre de démarrage |
| `security.nesting` | `true` | Autoriser Docker dans LXC |
| `security.privileged` | `true` | Conteneur privilégié ⚠️ |
| `snapshots.schedule` | `@daily` | Snapshots automatiques |
| `snapshots.expiry` | `7d` | Expiration des snapshots |

---

## 20. Profils — Le concept

Un **profil** est un ensemble réutilisable de configuration (devices + config)
appliqué à plusieurs conteneurs. C'est l'équivalent des "rôles" Ansible ou des
templates Proxmox.

```
┌──────────┐   ┌──────────� |   ┌──────────┐
│ profil   │   │ profil   │   │ profil   │
│ default  │ + │ web      │ + │ limits-m │ = config finale de c1
└──────────┘   └──────────┘   └──────────┘
     (réseau + disque)   (nginx, ports)   (2 CPU / 4 Go)
```

**Ordre d'application :** les profils s'appliquent dans l'ordre de la liste ;
la **config locale** du conteneur a toujours le dernier mot.

```bash
# Voir les profils d'un conteneur
lxc config show c1 | grep -A5 profiles

# Changer l'ordre / la liste des profils
lxc profile assign c1 default,web,limits-m
```

---

## 21. Profils — Créer et éditer

```bash
# Créer un profil vide
lxc profile create web

# Éditer (ouvre $EDITOR avec un squelette YAML)
lxc profile edit web
```

**Exemple : profil `web` complet :**

```yaml
config:
  limits.cpu: "2"
  limits.memory: 2GiB
  boot.autostart: "true"
  user.role: webserver
description: Profil serveurs web mutualises
devices:
  eth0:
    name: eth0
    network: lxdbr0
    type: nic
  root:
    path: /
    pool: default
    type: disk
    size: 20GiB
  http:
    connect: tcp:127.0.0.1:80
    listen: tcp:0.0.0.0:8080
    type: proxy
name: web
```

**Autres commandes :**

```bash
lxc profile list
lxc profile show web
lxc profile rename web web-prod
lxc profile copy web web-staging
lxc profile delete web-staging
```

---

## 22. Profils — Appliquer, exemples concrets

```bash
# Appliquer au lancement
lxc launch ubuntu:24.04 web1 --profile default --profile web

# Appliquer à un conteneur existant (à chaud pour la plupart des clés)
lxc profile add c1 web
lxc profile remove c1 web

# Remplacer toute la liste
lxc profile assign c1 default,db
```

**Bibliothèque de profils suggérée pour une équipe :**

| Profil | Contenu | Usage |
|---|---|---|
| `default` | réseau + disque de base | tout conteneur |
| `limits-s/m/l/xl` | CPU/RAM calibrés | gabarits de taille |
| `web` | limites + proxy 80/443 | serveurs web |
| `db` | RAM élevée, I/O prioritaires, shutdown long | bases de données |
| `docker-host` | `security.nesting=true`, disque 50 Go | hôtes Docker |
| `no-ipv6` | `ipv6.address=none` | réseau IPv4 only |

> ✅ **Bonne pratique :** tout ce qui est répété 3 fois devient un profil.
> Versionnez vos profils (`lxc profile show web > profiles/web.yaml` dans Git).

---

## 23. Réseau — Concepts (bridge, macvlan, routed, physical)

| Type | Description | IP du conteneur | Cas d'usage |
|---|---|---|---|
| `bridged` | Attaché à un bridge Linux | DHCP du bridge ou statique | Standard, recommandé |
| `macvlan` | Interface virtuelle sur l'interface physique | DHCP du LAN | IP LAN directe |
| `routed` | Routage via l'hôte (pas de bridge) | Statique, routée | Environnements sans bridge |
| `physical` | Interface physique dédiée au conteneur | Selon réseau | SR-IOV, cas avancés |
| `ovn` | Réseau virtuel OVN | Géré par OVN | Overlay multi-hôtes |
| `proxy` | Redirection de ports (pas un vrai NIC) | — | Exposer un port |

**Recommandation :** `bridged` sur `lxdbr0` (ou votre bridge) dans 90 % des cas.
`macvlan` quand le conteneur doit apparaître comme une machine du LAN
(attention : l'hôte ne peut pas joindre le conteneur en macvlan — limitation
du noyau, voir section 26).

---

## 24. Réseau — Le bridge `lxdbr0` par défaut

Créé par `lxd init`, c'est un bridge Linux privé avec NAT :

```bash
# Détails
lxc network show lxdbr0

# Voir les baux DHCP attribués
lxc network show lxdbr0 | grep -A20 "Used by"

