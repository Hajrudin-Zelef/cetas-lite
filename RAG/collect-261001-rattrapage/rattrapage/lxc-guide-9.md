---
id: collect-261001-rattrapage/rattrapage/lxc-guide-9
title: "LXC & LXD — Guide ultra-complet"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["datacenter", "memory"]
source: docs/RAG/collect-261001-rattrapage/lxc_guide.md
source_anchor: ""
source_lines: [2211, 2443]
sha256: b75093b3dd6ce5bac17d7d01039a85f62dd0b078d4946a9f3607cf569ffabb3d
---

# 4. Créer un CT avec ce rootfs (méthode manuelle)
pct create 100 local:vztmpl/ubuntu-24.04-standard_24.04-2_amd64.tar.zst \
  --rootfs local-lvm:20 --hostname c1-migre ...
# Puis remplacer le contenu du volume par le rootfs exporté
```

> 📝 Il n'existe pas de bouton "import LXD" dans Proxmox : la migration
> passe par l'export du rootfs + recréation. Pour des applis conteneurisées
> proprement (Docker, Ansible), **re-déployer** est souvent plus propre que
> migrer.

**De Proxmox vers LXD :** `vzdump` → extraire le `.tar.zst` → `lxc import`
via un rootfs re-packagé, ou `lxc publish` après recréation manuelle.

---

## 74. Proxmox VE — Bonnes pratiques templates

- **Templates standardisés :** un template par OS majeur, mis à jour
  mensuellement (`pveam download` + patch + `vzdump` en template).
- **Cloud-init :** utilisez-le pour hostname, clés SSH, IP (`--ciuser`,
  `--sshkey`, `--ipconfig0`) → conteneurs reproductibles.
- **Unprivileged par défaut** (`--unprivileged 1`), nesting seulement si besoin.
- **Backups `vzdump` planifiés** (datacenter → backup → ajouter) : c'est
  l'avantage n°1 de Proxmox sur LXD standalone.
- **Séparez OS et données** : disque système sur `local-lvm`, données sur
  volume dédié (monté via `pct set 100 -mp0 ...`).

```bash
# Exemple : création reproductible avec cloud-init
pct create 101 local:vztmpl/debian-12-standard_12.7-1_amd64.tar.zst \
  --hostname db1 --cores 4 --memory 8192 \
  --rootfs ceph-vm:30 \
  --net0 name=eth0,bridge=vmbr0,ip=10.0.0.21/24,gw=10.0.0.1 \
  --ciuser admin --sshkey ~/.ssh/id_ed25519.pub \
  --unprivileged 1 --onboot 1 --startup order=2,up=30
```

---

## 75. Cas pratique 1 — Conteneur web (nginx + PHP)

**Objectif :** un serveur web isolé, exposé sur le port 8080 de l'hôte.

```bash
# 1. Création avec le profil web (section 21)
lxc launch ubuntu:24.04 web1 --profile default --profile web

# 2. Installation
lxc exec web1 -- apt update
lxc exec web1 -- apt install -y nginx php-fpm php-mysql

# 3. Déployer le site
lxc file push -r ./monsite/ web1/var/www/html/ --create-dirs
lxc exec web1 -- chown -R www-data:www-data /var/www/html

# 4. Config nginx (via file edit ou push)
lxc file push ./web1-nginx.conf web1/etc/nginx/sites-available/monsite
lxc exec web1 -- ln -sf /etc/nginx/sites-available/monsite /etc/nginx/sites-enabled/
lxc exec web1 -- nginx -t && systemctl reload nginx

# 5. Exposer le port (si pas déjà dans le profil)
lxc config device add web1 http proxy listen=tcp:0.0.0.0:8080 connect=tcp:127.0.0.1:80

# 6. Snapshot de mise en service
lxc snapshot web1 mise-en-service
```

**Vérification :** `curl http://hote:8080` → page du site.
**Supervision :** alerte si le port 8080 ne répond plus (section 65).

---

## 76. Cas pratique 2 — Conteneur de build / CI

**Objectif :** un runner de build jetable, reconstruit à chaque pipeline.

```bash
# 1. Profil build (fort CPU, disque large, éphémère)
lxc profile create build
lxc profile set build limits.cpu 8
lxc profile set build limits.memory 16GiB
lxc profile device add build root disk path=/ pool=default size=100GiB

# 2. Lancement depuis la golden image (section 40)
lxc launch build-golden-v3 build-$CI_PIPELINE_ID --profile default --profile build

# 3. Build
lxc exec build-$CI_PIPELINE_ID -- su - builder -c "cd /src && make -j8"

# 4. Récupérer les artefacts
lxc file pull -r build-$CI_PIPELINE_ID/home/builder/dist/ ./artefacts/

# 5. Destruction (zéro trace)
lxc delete build-$CI_PIPELINE_ID --force
```

**Avantages vs Docker :** environnement système complet (systemd, multi-services),
isolation forte, démarrage en 2 s, coût nul de reconstruction.

---

## 77. Cas pratique 3 — Isolation d'un service (PostgreSQL)

**Objectif :** PostgreSQL dédié, données sur volume séparé, sauvegardé.

```bash
# 1. Volume de données
lxc storage volume create default pgdata-prod size=200GiB

# 2. Conteneur avec profil db
lxc launch ubuntu:24.04 pg1 --profile default --profile db

# 3. Attacher le volume
lxc config device add pg1 pgdata disk pool=default source=pgdata-prod \
  path=/var/lib/postgresql

# 4. Installer et configurer
lxc exec pg1 -- apt update && lxc exec pg1 -- apt install -y postgresql-16
lxc exec pg1 -- systemctl enable --now postgresql

# 5. Durcissement : n'écouter que sur le bridge interne
lxc exec pg1 -- bash -c "echo \"listen_addresses = '10.10.10.60'\" >> /etc/postgresql/16/main/postgresql.conf"
lxc config device set pg1 eth0 ipv4.address 10.10.10.60
lxc restart pg1

# 6. Snapshots + backup du volume
lxc config set pg1 snapshots.schedule "@daily"
lxc config set pg1 snapshots.expiry 7d
```

**Restauration :** volume `pgdata-prod` réattachable à un nouveau conteneur
en cas de problème OS (données intactes).

---

## 78. Cas pratique 4 — Reverse proxy multi-sites

**Objectif :** un conteneur nginx qui route vers N conteneurs web internes.

```
Internet → hote:80/443 → [ct reverse-proxy] → web1 (10.10.10.11)
                                          → web2 (10.10.10.12)
                                          → api  (10.10.10.13)
```

```bash
# 1. Conteneur proxy
lxc launch ubuntu:24.04 rproxy --profile default
lxc exec rproxy -- apt install -y nginx certbot python3-certbot-nginx

# 2. Exposer 80/443 de l'hôte vers le conteneur
lxc config device add rproxy http proxy listen=tcp:0.0.0.0:80 connect=tcp:127.0.0.1:80
lxc config device add rproxy https proxy listen=tcp:0.0.0.0:443 connect=tcp:127.0.0.1:443

# 3. Config nginx (exemple : deux vhosts)
lxc file push ./rproxy-sites.conf rproxy/etc/nginx/conf.d/sites.conf
```

```nginx
# /etc/nginx/conf.d/sites.conf dans rproxy
upstream web1 { server 10.10.10.11:80; }
upstream web2 { server 10.10.10.12:80; }

server {
    listen 80;
    server_name site1.exemple.fr;
    location / { proxy_pass http://web1; proxy_set_header Host $host; }
}
server {
    listen 80;
    server_name site2.exemple.fr;
    location / { proxy_pass http://web2; proxy_set_header Host $host; }
}
```

```bash
lxc exec rproxy -- nginx -t && lxc exec rproxy -- systemctl reload nginx

# 4. TLS avec Let's Encrypt (le challenge HTTP passe par le proxy)
lxc exec rproxy -- certbot --nginx -d site1.exemple.fr -d site2.exemple.fr
```

---

## 79. Cas pratique 5 — Environnement de dev par développeur

**Objectif :** chaque dev a son conteneur identique, recréable en 1 commande.

```bash
# 1. Golden image "dev" (outils communs préinstallés)
lxc launch ubuntu:24.04 dev-template
lxc exec dev-template -- apt install -y git build-essential python3-venv \
  nodejs npm docker.io zsh tmux htop
# ... dotfiles, config git globale ...
lxc stop dev-template
lxc publish dev-template --alias dev-golden-v1

# 2. Création pour un nouveau dev (1 commande)
lxc launch dev-golden-v1 dev-alice --profile default
lxc config device add dev-alice home disk \
  source=/home/alice source=/dev/null 2>/dev/null || true
# Mieux : volume persistant par dev
lxc storage volume create default dev-alice-home size=50GiB
lxc config device add dev-alice home disk pool=default \
  source=dev-alice-home path=/home/dev

# 3. Le dev s'y connecte en SSH ou via lxc exec
lxc exec dev-alice -- su - dev
```

**Reset du poste dev :** `lxc delete dev-alice --force` + relance →
environnement neuf en 10 secondes, home préservé sur le volume.

---

## 80. Erreurs classiques — 12 pièges et leurs solutions

### Erreur 1 : `Error: Get "http://unix/1.0": dial unix ... permission denied`

**Cause :** l'utilisateur n'est pas dans le groupe `lxd`.
**Solution :**

```bash
sudo usermod -aG lxd "$USER"
newgrp lxd   # ou reconnexion complète
```

### Erreur 2 : Le conteneur démarre mais n'a pas d'IP

**Causes possibles :** dnsmasq du bridge en conflit, DHCP désactivé, image sans
client DHCP.

