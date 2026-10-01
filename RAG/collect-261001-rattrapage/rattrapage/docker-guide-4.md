---
id: collect-261001-rattrapage/rattrapage/docker-guide-4
title: "Guide Docker complet — Production & Sysadmin"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["arr", "memory"]
source: docs/RAG/collect-261001-rattrapage/docker_guide.md
source_anchor: ""
source_lines: [722, 1031]
sha256: 593a13ced8b25db8725525320968ee01406709db2ecf3dc4192106a86a73928a
---

# Guide Docker complet — Production & Sysadmin

  db:
    image: postgres:16-alpine
    environment:
      POSTGRES_USER: app
      POSTGRES_PASSWORD_FILE: /run/secrets/db_password
      POSTGRES_DB: appdb
    secrets:
      - db_password
    volumes:
      - db-data:/var/lib/postgresql/data
    networks:
      - back
    restart: unless-stopped

networks:
  front:
  back:

volumes:
  db-data:

secrets:
  db_password:
    file: ./secrets/db_password.txt
```

## 27. Compose : services, images et build

```yaml
services:
  # 1. Depuis une image du registre
  web:
    image: nginx:1.27-alpine

  # 2. Depuis un Dockerfile local
  api:
    build:
      context: ./api
      dockerfile: Dockerfile.prod
      args:
        APP_VERSION: "1.4.2"
      target: runtime        # étape du multi-stage à utiliser

  # 3. Commande et point d'entrée surchargés
  worker:
    image: mon-app:1.4.2
    command: ["python", "-m", "worker"]
    entrypoint: ["/bin/sh", "-c"]
```

```bash
# Rebuild uniquement le service api puis relance
docker compose build api
docker compose up -d api

# Forcer le rebuild sans cache
docker compose build --no-cache api
```

## 28. Compose : depends_on et ordre de démarrage

```yaml
services:
  api:
    depends_on:
      db:
        condition: service_healthy   # attend que db soit SAIN
      cache:
        condition: service_started   # attend le simple démarrage
  db:
    image: postgres:16-alpine
    healthcheck:
      test: ["CMD-SHELL", "pg_isready -U app"]
      interval: 10s
      timeout: 5s
      retries: 5
```

- `depends_on` ne garantit **que l'ordre**, pas que l'app soit prête :
  combinez avec `healthcheck` + `condition: service_healthy`.
- L'application doit aussi **réessayer** sa connexion à la base au démarrage
  (backoff), car en production l'ordre ne suffit jamais.

## 29. Compose : restart policies

```yaml
services:
  web:
    restart: unless-stopped   # redémarre toujours, sauf arrêt manuel explicite
```

| Politique | Comportement |
|---|---|
| `no` | jamais redémarré (défaut) |
| `always` | toujours redémarré, même après un `docker stop` + reboot hôte |
| `unless-stopped` | comme `always`, mais respecte un arrêt manuel (**recommandé en prod**) |
| `on-failure[:N]` | redémarre seulement si code de sortie ≠ 0 (N = max de tentatives) |

> En production : `unless-stopped` partout, sauf jobs ponctuels (`no`).

## 30. Compose : healthcheck

```yaml
services:
  api:
    healthcheck:
      test: ["CMD", "curl", "-f", "http://localhost:8000/health"]
      # variantes :
      # test: ["CMD-SHELL", "pg_isready -U app -d appdb"]
      # test: ["CMD", "wget", "--no-verbose", "--tries=1", "--spider", "http://localhost/"]
      interval: 30s        # entre deux contrôles
      timeout: 5s          # délai max d'un contrôle
      retries: 3           # échecs avant "unhealthy"
      start_period: 15s    # grâce au démarrage (pas compté comme échec)
      start_interval: 5s   # (optionnel) intervalle pendant start_period
```

```bash
docker inspect --format '{{.State.Health.Status}}' projet-api-1
docker compose ps   # la colonne STATUS affiche (healthy) / (unhealthy)
```

Un service `unhealthy` : ne le laisse jamais tourner en prod sans alerte —
c'est le signal d'un déploiement raté ou d'une dépendance tombée.

## 31. Compose : variables d'environnement

```yaml
services:
  api:
    environment:
      APP_ENV: production          # valeur en clair dans le YAML
      LOG_LEVEL: info
      DB_HOST: db
      # Ne JAMAIS mettre de mot de passe en clair ici en prod :
      DB_PASSWORD: ${DB_PASSWORD}  # lu depuis l'environnement du shell / .env
    env_file:
      - ./config/app.env           # fichier clé=valeur
      - ./config/override.env
```

```bash
# .env à côté du compose.yaml (chargé automatiquement par compose)
DB_PASSWORD=S3cret_tres_long
APP_ENV=production

# Priorité (du plus fort au plus faible) :
# 1. `docker compose run -e VAR=...`
# 2. environment: du YAML (${VAR} interpolée depuis le shell/.env)
# 3. env_file:
# 4. valeurs par défaut ${VAR:-defaut} dans le YAML
```

```yaml
# Valeur par défaut si la variable est absente
image: "mon-app:${APP_VERSION:-latest}"
```

> ⚠️ Le fichier `.env` contient des secrets : `chmod 600 .env` et **jamais**
> dans git (ajoutez-le au `.gitignore`).

## 32. Compose : networks

```yaml
services:
  web:
    networks:
      - front
  api:
    networks:
      front:
        aliases: ["api-interne"]   # nom DNS supplémentaire
      back:
  db:
    networks:
      - back

networks:
  front:                            # réseau applicatif
  back:                             # réseau données, isolé du web
    internal: true                  # PAS d'accès Internet pour ce réseau
```

- Chaque réseau compose = un **bridge** avec **DNS intégré** : les conteneurs
  se joignent par **nom de service** (`db`, `api`).
- `internal: true` : parfait pour le réseau base de données.
- Réseau externe pré-existant (ex : celui de Traefik, voir section 73) :

```yaml
networks:
  proxy:
    name: proxy
    external: true
```

## 33. Compose : volumes et bind mounts

```yaml
services:
  db:
    image: postgres:16-alpine
    volumes:
      # Volume nommé (géré par Docker, recommandé pour les données)
      - db-data:/var/lib/postgresql/data
      # Bind mount (dossier de l'hôte)
      - ./backups:/backups
      # Fichier unique, en lecture seule
      - ./config/postgresql.conf:/etc/postgresql/postgresql.conf:ro
      # tmpfs (RAM, non persisté, pour caches/secrets temporaires)
      - type: tmpfs
        target: /tmp
        tmpfs_size: 100m

volumes:
  db-data:
    driver: local
    # driver_opts pour monter un NFS par ex. (voir section 45)
```

Syntaxe courte vs longue :

| Besoin | Syntaxe courte |
|---|---|
| Volume nommé | `db-data:/var/lib/postgresql/data` |
| Bind mount | `./config:/etc/app:ro` |
| Lecture seule | ajouter `:ro` |

> `:ro` partout où l'appli n'a pas besoin d'écrire : principe du moindre
> privilège, aussi pour les fichiers.

## 34. Compose : ressources, logging, divers

```yaml
services:
  api:
    deploy:
      resources:
        limits:
          cpus: "1.5"
          memory: 512M
        reservations:
          memory: 256M
    logging:
      driver: json-file
      options:
        max-size: "10m"
        max-file: "3"
    ulimits:
      nofile:
        soft: 65536
        hard: 65536
    sysctls:
      net.core.somaxconn: 1024
    stop_grace_period: 30s     # délai avant SIGKILL (défaut 10 s)
    init: true                 # mini-init (tini) comme PID 1 : récole les zombies
```

- `deploy.resources` fonctionne avec `docker compose` (ignoré en partie par
  l'ancien `docker-compose` v1, mais le plugin v2 l'applique).
- `init: true` : recommandé si l'app ne gère pas les processus fils.

## 35. Compose : profils et surcharges par environnement

```yaml
services:
  web:
    image: nginx:1.27-alpine
    profiles: ["prod"]        # ne démarre que si le profil est activé
  mailhog:
    image: mailhog/mailhog
    profiles: ["dev"]         # outil de dev uniquement
```

```bash
docker compose --profile prod up -d
docker compose --profile dev up -d
```

Surcharges par environnement (pattern classique) :

```bash
# compose.yaml        : base commune
# compose.prod.yaml   : surcharge production
# compose.dev.yaml    : surcharge développement
docker compose -f compose.yaml -f compose.prod.yaml up -d
docker compose -f compose.yaml -f compose.dev.yaml up -d
```

```yaml
# compose.prod.yaml : exemple de surcharge
services:
  web:
    restart: unless-stopped
    logging:
      options: { max-size: "10m", max-file: "5" }
```

## 36. Compose en pratique : commandes de tous les jours

```bash
# État, logs, top
docker compose ps
docker compose top
docker compose logs -f --tail=100 api

# (Re)lancer un seul service
docker compose up -d --build api
docker compose restart api
docker compose stop api && docker compose rm -f api

