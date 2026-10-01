---
id: collect-261001-rattrapage/rattrapage/docker-guide-1
title: "Guide Docker complet — Production & Sysadmin"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["attention", "pruning"]
source: docs/RAG/collect-261001-rattrapage/docker_guide.md
source_anchor: ""
source_lines: [1, 191]
sha256: 8b871e3d2fc9851246f440fe03a9df0f7133ebe60db29bf325da784b139fb5bf
---

# Guide Docker complet — Production & Sysadmin

> **Public :** administrateurs systèmes, chefs de service systèmes/énergies, équipes
> exploitation. Angle : pratique, production, robustesse.
> **Version :** Docker Engine 25.x/26.x, Compose v2 (plugin `docker compose`).
> Tous les exemples ont été conçus pour Debian 12/13 et Ubuntu 22.04/24.04 LTS.
> *Avertissement : testez toujours les commandes sensibles (iptables, macvlan,
> pruning) sur un environnement de pré-production avant la production.*

---

## 1. Pourquoi Docker en production ?

- **Isolation** : chaque application tourne dans son propre espace (namespaces,
  cgroups), sans interférer avec ses voisines.
- **Reproductibilité** : une image fige l'application + ses dépendances. « Ça
  marchait sur ma machine » disparaît.
- **Densité** : bien moins lourd qu'une VM (pas de noyau invité).
- **Déploiement** : `docker compose up -d` et la stack est en ligne.
- **Rollback** : une image est taguée, on peut revenir à la version précédente en
  une commande.

## 2. Le vocabulaire indispensable

| Terme | Définition | Analogie |
|---|---|---|
| Image | Modèle immuable (lecture seule) qui contient l'app et ses dépendances | Le moule |
| Conteneur | Instance en cours d'exécution d'une image | Le gâteau sorti du moule |
| Volume | Stockage persistant géré par Docker | Le disque dur externe |
| Bind mount | Dossier de l'hôte monté dans le conteneur | Un dossier partagé |
| Réseau | Réseau virtuel reliant les conteneurs | Le switch virtuel |
| Registre (registry) | Entrepôt d'images (Docker Hub, registre privé) | Le magasin de moules |
| Dockerfile | Recette texte pour construire une image | La recette de cuisine |
| Compose | Fichier YAML décrivant une stack multi-conteneurs | Le plan du repas complet |

- **Image ≠ conteneur** : on peut lancer 10 conteneurs depuis la même image.
- **Conteneur = éphémère par défaut** : tout ce qui est écrit dans sa couche
  inscriptible disparaît à sa suppression, sauf volumes et bind mounts.
- **Tag** : étiquette de version d'image (`nginx:1.27`, `postgres:16`).
  `latest` = tag mouvant, à éviter en production (voir section 63).

## 3. Architecture : client, daemon, registres

```
┌─────────────┐      API REST      ┌──────────────────┐
│ docker CLI  │ ◄────────────────► │ dockerd (daemon) │
└─────────────┘   /var/run/        └────────┬─────────┘
                  docker.sock               │
                                            ▼
                                   ┌──────────────────┐
                                   │ containerd + runc│  (exécution réelle)
                                   └────────┬─────────┘
                                            │
                        ┌───────────────────┼───────────────────┐
                        ▼                   ▼                   ▼
                   conteneurs          images/volumes       réseaux
```

- Le **daemon** (`dockerd`) fait tout le travail lourd ; la CLI n'est qu'un
  client qui lui parle via le socket Unix `/var/run/docker.sock`.
- **containerd** gère le cycle de vie des conteneurs, **runc** les lance.
- Le **socket Docker = root sur l'hôte** : quiconque peut y écrire peut lancer
  un conteneur `--privileged` et prendre le contrôle total de la machine
  (voir section 55, Sécurité).

## 4. Installation sur Debian/Ubuntu — dépôt officiel (pas le snap !)

> ⚠️ **Ne pas utiliser** le paquet `docker.io` des dépôts Ubuntu (souvent en
> retard de plusieurs versions) ni le snap (problèmes de permissions, de
> volumes et de réseau). Toujours le dépôt officiel `download.docker.com`.

```bash
# 1. Nettoyer d'anciennes installations conflictuelles
sudo apt-get remove -y docker.io docker-doc docker-compose \
  docker-compose-v2 podman-docker containerd runc || true

# 2. Prérequis
sudo apt-get update
sudo apt-get install -y ca-certificates curl gnupg

# 3. Clé GPG officielle Docker
sudo install -m 0755 -d /etc/apt/keyrings
curl -fsSL https://download.docker.com/linux/debian/gpg \
  | sudo gpg --dearmor -o /etc/apt/keyrings/docker.gpg
sudo chmod a+r /etc/apt/keyrings/docker.gpg
# Sur Ubuntu : remplacer "debian" par "ubuntu" dans les deux URL/commandes

# 4. Ajouter le dépôt officiel
echo "deb [arch=$(dpkg --print-architecture) \
  signed-by=/etc/apt/keyrings/docker.gpg] \
  https://download.docker.com/linux/debian \
  $(. /etc/os-release && echo "$VERSION_CODENAME") stable" \
  | sudo tee /etc/apt/sources.list.d/docker.list > /dev/null

# 5. Installer le moteur + le plugin Compose
sudo apt-get update
sudo apt-get install -y docker-ce docker-ce-cli containerd.io \
  docker-buildx-plugin docker-compose-plugin

# 6. Vérifier
docker --version
docker compose version
sudo docker run --rm hello-world
```

Checklist post-installation :

- [ ] `docker run --rm hello-world` fonctionne
- [ ] Le service démarre au boot : `sudo systemctl is-enabled docker`
- [ ] Version du moteur ≥ 24 : `docker version`
- [ ] Le plugin compose répond : `docker compose version`

## 5. Gestion du groupe docker et démarrage automatique

```bash
# Permettre à votre utilisateur d'utiliser docker SANS sudo
# (attention : membre du groupe docker = équivalent root, voir section 55)
sudo usermod -aG docker "$USER"
newgrp docker   # ou reconnectez-vous

# Activer le démarrage automatique
sudo systemctl enable --now docker
sudo systemctl enable --now containerd

# Vérifier l'état
sudo systemctl status docker --no-pager
docker info
docker info --format '{{.ServerVersion}}'
```

> 🔐 En production sur serveur partagé, **préférez `sudo docker ...`** plutôt
> que d'ajouter tout le monde au groupe `docker`.

## 6. Configuration du daemon : /etc/docker/daemon.json

Fichier de configuration central du daemon. Toute modification nécessite un
rechargement : `sudo systemctl reload docker` (ou restart si le reload ne
suffit pas).

```json
{
  "log-driver": "json-file",
  "log-opts": {
    "max-size": "10m",
    "max-file": "3"
  },
  "storage-driver": "overlay2",
  "exec-opts": ["native.cgroupdriver=systemd"],
  "live-restore": true,
  "default-address-pools": [
    { "base": "172.26.0.0/16", "size": 24 }
  ],
  "dns": ["9.9.9.9", "1.1.1.1"],
  "icc": true,
  "iptables": true,
  "userland-proxy": false,
  "default-ulimits": {
    "nofile": { "Name": "nofile", "Hard": 65536, "Soft": 65536 }
  },
  "metrics-addr": "127.0.0.1:9323",
  "experimental": false
}
```

Explication des clés :

| Clé | Rôle |
|---|---|
| `log-driver` / `log-opts` | Driver de logs + **rotation** (indispensable : sans rotation, les logs remplissent le disque) |
| `storage-driver` | `overlay2` = le driver moderne et recommandé |
| `live-restore` | Les conteneurs **survivent** au redémarrage du daemon (mises à jour sans coupure) |
| `default-address-pools` | Plages IP des réseaux bridge (à adapter pour éviter les collisions avec votre LAN/VPN) |
| `dns` | DNS utilisés par les conteneurs (si vide : ceux de l'hôte) |
| `userland-proxy: false` | Désactive le proxy userspace (performances), nécessite iptables actif |
| `metrics-addr` | Expose les métriques Prometheus du daemon (voir section 58) |

```bash
# Appliquer la configuration
sudo systemctl reload docker
# Si le reload échoue (changement profond), faire :
sudo systemctl restart docker
# Vérifier que la config est prise en compte
docker info | grep -i -E "logging|storage|cgroup"
```

## 7. Premières commandes : cycle de vie d'un conteneur

```bash
# Lancer un conteneur nginx en tâche de fond
docker run -d --name mon-nginx -p 8080:80 nginx:1.27

