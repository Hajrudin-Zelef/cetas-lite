---
id: collect-261001-rattrapage/rattrapage/docker-guide-6
title: "Guide Docker complet — Production & Sysadmin"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["memory"]
source: docs/RAG/collect-261001-rattrapage/docker_guide.md
source_anchor: ""
source_lines: [1296, 1574]
sha256: b17fb6244df824b3072435d1e034b9824dce8f6801be7dd794f128a31d813e8e
---

# Guide Docker complet — Production & Sysadmin

```yaml
services:
  app:
    image: mon-app:1.0
    volumes: ["shared:/data"]
  backup:
    image: alpine:3.20
    volumes: ["shared:/data:ro"]
    command: ["sh", "-c", "tar czf /backup/data.tgz -C /data ."]
    volumes:
      - shared:/data:ro
      - ./backups:/backup

volumes:
  shared:
```

## 49. Logs : drivers et lecture

```bash
docker logs --tail 100 -f mon-conteneur
docker logs --since 1h mon-conteneur
docker logs --since 2026-09-26T08:00 --until 2026-09-26T09:00 mon-conteneur
docker compose logs -f --tail=200 api db
```

Drivers de logs (`--log-driver` ou `daemon.json`) :

| Driver | Usage |
|---|---|
| `json-file` | défaut, fichiers locaux avec rotation |
| `syslog` | envoi vers syslog/rsyslog de l'hôte |
| `journald` | journal systemd (`journalctl -u docker`) |
| `gelf` | Graylog / Logstash |
| `none` | aucun log (debug extrême uniquement) |

```bash
# Logs via journald
docker run -d --log-driver journald --name app mon-app:1.0
journalctl CONTAINER_NAME=app -f
```

> ⚠️ Avec un driver distant (syslog, gelf), `docker logs` ne fonctionne plus
> en local : prévoyez votre lecteur de logs centralisé.

## 50. Rotation des logs : éviter le disque plein

**Le piège n°1 des débutants** : sans rotation, un conteneur verbeux remplit
`/var/lib/docker` jusqu'au plantage de l'hôte. La rotation se configure :

1. **Globalement** dans `daemon.json` (recommandé, voir section 6) :

```json
{
  "log-driver": "json-file",
  "log-opts": { "max-size": "10m", "max-file": "3" }
}
```

2. **Par conteneur/service** (surcharge) :

```yaml
services:
  api:
    logging:
      driver: json-file
      options:
        max-size: "10m"
        max-file: "3"
```

```bash
# Vérifier la taille des logs d'un conteneur
docker inspect --format '{{.LogPath}}' mon-conteneur
sudo du -sh /var/lib/docker/containers/*/*-json.log | sort -h | tail
```

## 51. Limites de ressources : CPU et mémoire

Sans limites, un conteneur peut affamer l'hôte. En production, **toujours**
poser des limites.

```bash
docker run -d --name api \
  --memory=512m --memory-swap=1g \
  --cpus=1.5 \
  --pids-limit=200 \
  mon-api:1.0
```

| Option | Effet |
|---|---|
| `--memory=512m` | RAM max ; dépassement → le conteneur est **tué (OOM)** |
| `--memory-swap=1g` | RAM + swap max (mettre = à `--memory` pour **désactiver le swap**) |
| `--cpus=1.5` | 1,5 cœur max |
| `--pids-limit=200` | anti fork-bomb |
| `--restart=unless-stopped` | redémarre après un OOM kill |

En compose (voir section 34) :

```yaml
deploy:
  resources:
    limits: { cpus: "1.5", memory: 512M }
    reservations: { memory: 256M }
```

```bash
# Surveiller
docker stats --no-stream
# Vérifier qu'une limite est appliquée
docker inspect --format '{{.HostConfig.Memory}}' api   # en octets
```

> 💡 Dimensionnez `memory` d'après `docker stats` en charge réelle + 20 % de
> marge, pas au doigt mouillé.

## 52. Variables d'environnement : les 4 voies (et leurs pièges)

```bash
# 1. -e en ligne de commande
docker run -d -e APP_ENV=production -e LOG_LEVEL=info mon-app:1.0

# 2. --env-file
docker run -d --env-file ./config/app.env mon-app:1.0

# 3. compose: environment: / env_file: (voir section 31)

# 4. Valeurs par défaut dans l'image (ENV du Dockerfile)
```

Règles :

- Les variables **non secrètes** (env, log level, ports) : `environment:` OK.
- Les **secrets** (mots de passe, clés API, tokens) : **jamais** en clair
  dans le YAML versionné → utilisez les **secrets Docker** ou un fichier
  `.env` en `chmod 600` hors git (section 54).
- `docker inspect` affiche les variables d'un conteneur en clair : toute
  personne avec accès Docker les voit.

## 53. Secrets Docker (Swarm) et fichiers de secrets en compose

En **Swarm** : secrets chiffrés, montés en `tmpfs` sous `/run/secrets/`.

```bash
# Créer un secret (Swarm)
echo "motdepasse_tres_long" | docker secret create db_password -
docker secret ls
```

En **compose simple** (mono-hôte), l'équivalent pragmatique :

```yaml
services:
  db:
    image: postgres:16-alpine
    environment:
      POSTGRES_PASSWORD_FILE: /run/secrets/db_password  # _FILE : lu depuis le fichier
    secrets:
      - db_password
  api:
    image: mon-api:1.0
    secrets:
      - db_password

secrets:
  db_password:
    file: ./secrets/db_password.txt   # chmod 600, hors git !
```

- Les images officielles (`postgres`, `mysql`, `redis`…) supportent les
  variantes `*_FILE` : le secret est lu depuis le fichier, **jamais** exposé
  en variable d'environnement visible par `docker inspect`.
- Permissions : `chmod 600 secrets/`, ajoutez `secrets/` au `.gitignore`.

## 54. Gestion des secrets : stratégie complète

| Niveau | Solution | Quand |
|---|---|---|
| Labo / dev | `.env` en `chmod 600`, hors git | poste isolé |
| Mono-hôte prod | secrets compose (`file:`) + `*_FILE` | PME, un serveur |
| Multi-hôtes | Vault / Doppler / Infisical, injection au déploiement | équipe, plusieurs serveurs |
| Swarm | `docker secret` natif | cluster Swarm |

Anti-patterns à bannir :

```dockerfile
# ❌ INTERDIT : secret dans le Dockerfile (visible dans docker history)
ENV DB_PASSWORD=S3cret
ARG API_KEY=xxx
```

```yaml
# ❌ INTERDIT : secret en clair dans un YAML versionné
environment:
  DB_PASSWORD: S3cret
```

Checklist secrets :

- [ ] Aucun secret dans les Dockerfiles ni dans git (`git log -S password`)
- [ ] `secrets/` et `.env` dans le `.gitignore`, `chmod 600`
- [ ] Rotation des mots de passe documentée (qui, quand, comment)
- [ ] Les images officielles utilisent les variantes `*_FILE`

## 55. SÉCURITÉ — le socket Docker et le groupe docker

> 🔴 **Point le plus important du guide** : l'accès au daemon Docker
> (`/var/run/docker.sock` ou le groupe `docker`) équivaut à **root sur l'hôte**,
> sans mot de passe :

```bash
# N'importe quel membre du groupe docker peut faire ceci :
docker run --rm -it --privileged -v /:/host alpine:3.20 chroot /host bash
# => shell root sur l'HÔTE. Game over.
```

Règles :

- N'ajoutez au groupe `docker` que des administrateurs de confiance.
- **Ne montez jamais** `/var/run/docker.sock` dans un conteneur, sauf outil
  d'administration légitime (Portainer, watchtower) — et comprenez le risque.
- Sur un serveur multi-utilisateurs, imposez `sudo docker` + audit.
- Protégez le socket : `srw-rw---- root:docker /var/run/docker.sock`.

## 56. SÉCURITÉ — bannir --privileged et durcir le runtime

```bash
# ❌ JAMAIS en production (donne accès à tous les devices de l'hôte)
docker run --privileged ...

# ✅ À la place : accordez uniquement la capability nécessaire
docker run --cap-add=NET_ADMIN --cap-drop=ALL mon-app:1.0
docker run --cap-drop=ALL --cap-add=CHOWN --cap-add=SETUID ...
```

Checklist de durcissement d'un conteneur :

```yaml
services:
  api:
    image: mon-api:1.0.0
    user: "1001:1001"          # ne pas tourner en root (même avec USER dans le Dockerfile, ceinture + bretelles)
    read_only: true            # système de fichiers racine en lecture seule
    tmpfs: ["/tmp", "/run"]    # zones inscriptibles en RAM
    cap_drop: ["ALL"]          # retire toutes les capabilities...
    cap_add: ["NET_BIND_SERVICE"]  # ...puis n'ajoute que le nécessaire
    security_opt:
      - no-new-privileges:true # empêche l'élévation via setuid
    pids_limit: 200
```

Vérification :

```bash
docker exec api id -u          # doit afficher 1001, pas 0
docker inspect --format '{{.HostConfig.ReadonlyRootfs}}' projet-api-1
```

## 57. SÉCURITÉ — images : choisir, épingler, scanner

- **Images officielles** ou éditeurs vérifiés uniquement.
- **Épinglez le digest** pour les déploiements critiques (le tag peut être
  réécrit, le digest non) :

```bash
docker pull nginx:1.27-alpine
docker inspect --format '{{index .RepoDigests 0}}' nginx:1.27-alpine
# => nginx@sha256:abc123...
```

```yaml
services:
  web:
    image: nginx:1.27-alpine@sha256:abc123...  # immuable, garanti
```

- **Scannez** les images avant production :

