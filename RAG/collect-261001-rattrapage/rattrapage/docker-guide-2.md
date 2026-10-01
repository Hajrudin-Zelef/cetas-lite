---
id: collect-261001-rattrapage/rattrapage/docker-guide-2
title: "Guide Docker complet — Production & Sysadmin"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["arr"]
source: docs/RAG/collect-261001-rattrapage/docker_guide.md
source_anchor: ""
source_lines: [192, 427]
sha256: af6a81c74b96efef3910066f1157c91094a17d4a95cda2697ef4fccc11cc625c
---

# Voir les conteneurs en cours
docker ps
docker ps -a          # même les arrêtés
docker ps --format "table {{.Names}}\t{{.Status}}\t{{.Ports}}"

# Voir les logs
docker logs mon-nginx
docker logs -f --tail 50 mon-nginx   # suivi temps réel, 50 dernières lignes

# Entrer dans le conteneur
docker exec -it mon-nginx bash
docker exec -it mon-nginx sh         # si pas de bash (images alpine)

# Arrêter / démarrer / redémarrer
docker stop mon-nginx      # SIGTERM puis SIGKILL après 10 s
docker start mon-nginx
docker restart mon-nginx
docker stop -t 30 mon-nginx  # délai de grâce de 30 s

# Supprimer (arrêté d'abord, sauf -f)
docker rm mon-nginx
docker rm -f mon-nginx      # force : stop + rm
```

## 8. Images : pull, tag, lister, nettoyer

```bash
docker pull nginx:1.27
docker pull postgres:16-alpine

docker images
docker images --format "table {{.Repository}}\t{{.Tag}}\t{{.Size}}"

# Tagger une image (ex : pour un registre privé)
docker tag mon-app:1.0 registry.interne:5000/mon-app:1.0

# Inspecter une image : couches, variables d'env, point d'entrée
docker image inspect nginx:1.27 | less

# Historique des couches d'une image
docker history nginx:1.27

# Supprimer une image
docker rmi nginx:1.27
docker rmi $(docker images -q)   # TOUTES les images : dangereux !

# Nettoyage : conteneurs arrêtés, réseaux et images non utilisés
docker system prune        # demande confirmation
docker system prune -a --volumes  # radical : + images + volumes (DANGEREUX)
docker system df           # voir l'espace utilisé par Docker
```

> 💡 `docker system df` est votre ami : il montre combien pèsent images,
> conteneurs, volumes et build cache.

## 9. Publier des ports et nommer les conteneurs

```bash
# -p [IP_hôte:]port_hôte:port_conteneur
docker run -d -p 8080:80 nginx              # toutes les interfaces
docker run -d -p 127.0.0.1:8080:80 nginx    # localhost uniquement (plus sûr)
docker run -d -p 8080:80 -p 8443:443 nginx  # plusieurs ports

# --name : nom explicite (sinon nom aléatoire type "stoic_bose")
docker run -d --name api-prod -p 3000:3000 mon-api:2.1

# --hostname : nom d'hôte vu depuis l'intérieur
docker run -d --hostname web1 --name web1 nginx
```

Tableau de décision d'exposition :

| Besoin | Commande |
|---|---|
| Site public | `-p 80:80 -p 443:443` |
| Admin locale uniquement | `-p 127.0.0.1:8080:80` |
| Conteneurs entre eux (compose) | pas de `-p`, le réseau interne suffit |
| Debug temporaire | `-p 8080:80`, à retirer ensuite |

## 10. Exécuter des commandes : exec, run, attacher

```bash
# Une commande ponctuelle dans un conteneur qui tourne
docker exec mon-nginx nginx -t            # tester la conf nginx
docker exec mon-nginx cat /etc/nginx/nginx.conf

# Shell interactif
docker exec -it mon-nginx bash

# Lancer un conteneur jetable pour un test
docker run --rm -it ubuntu:24.04 bash
docker run --rm alpine:3.20 ping -c 3 8.8.8.8

# Voir les processus d'un conteneur
docker top mon-nginx

# Statistiques temps réel (CPU, RAM, réseau, I/O)
docker stats
docker stats --no-stream --format "table {{.Name}}\t{{.CPUPerc}}\t{{.MemUsage}}"
```

## 11. Inspect : tout savoir sur un conteneur

```bash
docker inspect mon-nginx | less
# Extraire des champs précis (format Go template)
docker inspect --format '{{.State.Status}}' mon-nginx
docker inspect --format '{{.NetworkSettings.IPAddress}}' mon-nginx
docker inspect --format '{{range .Mounts}}{{.Source}} -> {{.Destination}}{{end}}' mon-nginx
```

Champs les plus utiles : `.State` (statut, code de sortie, heure de démarrage),
`.Mounts` (volumes), `.NetworkSettings` (IP, ports), `.Config` (image, env,
entrypoint).

## 12. Aide-mémoire des commandes quotidiennes

| Action | Commande |
|---|---|
| Lister conteneurs actifs | `docker ps` |
| Lister tout | `docker ps -a` |
| Logs temps réel | `docker logs -f <nom>` |
| Suivre les ressources | `docker stats` |
| Redémarrer | `docker restart <nom>` |
| Mettre à jour une image | `docker pull <image>:<tag>` + `docker compose up -d` |
| Espace disque Docker | `docker system df` |
| Tout nettoyer (prudent) | `docker system prune` |
| Version / infos | `docker version`, `docker info` |

## 13. Le Dockerfile : principes

- Un `Dockerfile` est une **recette** : chaque instruction crée une **couche**
  (layer) en lecture seule empilée sur la précédente.
- Les couches sont **mises en cache** : reconstruire après un petit changement
  ne refait que les couches modifiées et suivantes → ordonnez les instructions
  de la moins changeante à la plus changeante.
- Règle d'or : **un conteneur = un processus principal** (le `CMD`).

```dockerfile
# Exemple minimal : serveur web statique
FROM nginx:1.27-alpine
COPY ./site/ /usr/share/nginx/html/
EXPOSE 80
```

```bash
# Construire l'image
docker build -t mon-site:1.0 .
docker build -t mon-site:1.0 -f Dockerfile.prod .

# Lancer
docker run -d --name site -p 8080:80 mon-site:1.0
```

## 14. Instruction FROM : choisir sa base

```dockerfile
FROM ubuntu:24.04
FROM python:3.12-slim
FROM nginx:1.27-alpine
FROM node:20-alpine
```

- Préférez les images **officielles** (Docker Hub, badge « Official Image »).
- Préférez les variantes **`-slim`** (Debian allégé) ou **`-alpine`**
  (musl libc, ~5 Mo) pour réduire la surface d'attaque et le temps de pull.
- ⚠️ Alpine + musl : certaines applis (binaires compilés glibc, wheels Python
  exotiques) peuvent poser problème → testez, sinon restez sur `-slim`.
- **Épinglez toujours un tag précis** (`python:3.12-slim`), jamais `latest`
  en production.

## 15. Instructions RUN, COPY, ADD

```dockerfile
# RUN : exécute une commande pendant le build (nouvelle couche à chaque RUN)
RUN apt-get update && apt-get install -y --no-install-recommends \
      curl ca-certificates \
    && rm -rf /var/lib/apt/lists/*
# ^ combinez update+install+nettoyage dans UN SEUL RUN,
#   sinon le cache apt reste dans une couche intermédiaire.

# COPY : copie fichiers/dossiers du contexte vers l'image (préféré)
COPY ./app/ /app/
COPY requirements.txt /app/

# ADD : comme COPY + décompression auto des .tar + téléchargement d'URL
# À ÉVITER sauf pour décompresser une archive locale :
ADD https://example.com/fichier.tar.gz /tmp/   # déconseillé : préférez curl + RUN
```

Bonnes pratiques `RUN` :

- Chaînez avec `&&` et nettoyez dans la même instruction.
- `--no-install-recommends` avec apt pour limiter la taille.
- En fin de RUN apt : `rm -rf /var/lib/apt/lists/*`.

## 16. Instructions CMD vs ENTRYPOINT

```dockerfile
# CMD : commande par défaut, REMPLAÇABLE par des arguments de `docker run`
CMD ["nginx", "-g", "daemon off;"]

# ENTRYPOINT : commande FIXE, les arguments de `docker run` s'y AJOUTENT
ENTRYPOINT ["python", "app.py"]
CMD ["--port", "8000"]
# docker run mon-image --port 9000  =>  python app.py --port 9000
```

- Forme **exec** (`["..."]`, JSON) : recommandée — le processus devient PID 1
  et reçoit correctement les signaux (SIGTERM lors du `docker stop`).
- Forme **shell** (`CMD python app.py`) : lance `/bin/sh -c`, le signal
  n'atteint pas toujours l'app → arrêt brutal après le timeout.

## 17. Instructions ENV, WORKDIR, EXPOSE, USER

```dockerfile
ENV PYTHONUNBUFFERED=1 \
    APP_ENV=production
# ^ les logs Python sortent immédiatement (pas de buffer)

WORKDIR /app
# ^ crée le dossier et s'y place ; tous les COPY/RUN/CMD suivants sont relatifs

EXPOSE 8000
# ^ documentation du port d'écoute (ne PUBLIE pas le port, c'est -p qui le fait)

# Créer un utilisateur non-root et l'utiliser (SÉCURITÉ, voir section 52)
RUN useradd -m -u 1001 appuser && chown -R appuser:appuser /app
USER appuser
```

> 🔐 **Ne jamais laisser tourner une appli en root dans le conteneur** si on
> peut l'éviter. `USER` doit apparaître avant `CMD`.

## 18. Instructions VOLUME, LABEL, HEALTHCHECK, ARG

