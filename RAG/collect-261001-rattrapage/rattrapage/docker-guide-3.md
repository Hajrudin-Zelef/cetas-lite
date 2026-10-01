---
id: collect-261001-rattrapage/rattrapage/docker-guide-3
title: "Guide Docker complet — Production & Sysadmin"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["arr"]
source: docs/RAG/collect-261001-rattrapage/docker_guide.md
source_anchor: ""
source_lines: [428, 721]
sha256: 8409faae968eb2866d5ac5093e2feb98d1b0f58f8e6d20957e4833248b300148
---

# Guide Docker complet — Production & Sysadmin

```dockerfile
ARG APP_VERSION=1.0.0
# ^ variable de BUILD (docker build --build-arg APP_VERSION=2.0 .)
#   Non persistée dans l'image finale (contrairement à ENV).

LABEL maintainer="infra@entreprise.fr" \
      version="${APP_VERSION}" \
      description="API de production"
# ^ métadonnées : traçabilité, inventaire.

VOLUME ["/data"]
# ^ déclare un point de montage persistant (préférez le déclarer dans compose).

HEALTHCHECK --interval=30s --timeout=5s --start-period=10s --retries=3 \
  CMD curl -f http://localhost:8000/health || exit 1
# ^ Docker surveille la santé de l'app ; visible dans `docker ps` et exploitable
#   par l'orchestrateur (restart, load balancer).
```

Différence **ARG vs ENV** :

| | ARG | ENV |
|---|---|---|
| Disponible au build | ✅ | ✅ |
| Persisté dans l'image/conteneur | ❌ | ✅ |
| Usage typique | version, options de compilation | config runtime |

⚠️ Ne mettez **jamais de secret** dans un `ARG` ou un `ENV` du Dockerfile :
ils restent visibles dans l'historique de l'image (`docker history`).

## 19. Multi-stage builds : des images petites et propres

Principe : une étape **build** (avec compilateurs, SDK) puis une étape
**runtime** qui ne récupère que les artefacts. Résultat : image finale légère,
sans outils de compilation (= moins de vulnérabilités).

```dockerfile
# ---------- Étape 1 : build ----------
FROM golang:1.22-alpine AS builder
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o /app/mon-api ./cmd/api

# ---------- Étape 2 : runtime minimal ----------
FROM alpine:3.20
RUN adduser -D -u 1001 appuser
COPY --from=builder /app/mon-api /usr/local/bin/mon-api
USER appuser
EXPOSE 8000
ENTRYPOINT ["mon-api"]
```

Exemple Node.js :

```dockerfile
FROM node:20-alpine AS builder
WORKDIR /app
COPY package*.json ./
RUN npm ci --only=production=false
COPY . .
RUN npm run build

FROM node:20-alpine
WORKDIR /app
ENV NODE_ENV=production
COPY package*.json ./
RUN npm ci --only=production && npm cache clean --force
COPY --from=builder /app/dist ./dist
USER node
CMD ["node", "dist/index.js"]
```

Exemple Python :

```dockerfile
FROM python:3.12-slim AS builder
WORKDIR /app
COPY requirements.txt .
RUN pip install --no-cache-dir --prefix=/install -r requirements.txt

FROM python:3.12-slim
WORKDIR /app
COPY --from=builder /install /usr/local
COPY ./src ./src
RUN useradd -m -u 1001 appuser && chown -R appuser /app
USER appuser
CMD ["python", "-m", "src.main"]
```

## 20. Bonnes pratiques de taille d'image

Checklist anti-obésité :

- [ ] Base `alpine` ou `slim` quand c'est compatible
- [ ] Multi-stage pour les langages compilés / builds frontend
- [ ] Un seul `RUN` pour `apt-get update + install + rm -rf /var/lib/apt/lists/*`
- [ ] `--no-install-recommends` (apt), `--no-cache-dir` (pip), `npm cache clean`
- [ ] `.dockerignore` complet (voir section 21)
- [ ] Pas de secrets, pas de `.git`, pas de fichiers de test inutiles
- [ ] `docker images` : visez < 100 Mo pour une API simple, < 300 Mo sinon

```bash
# Analyser ce qui pèse dans une image
docker history --human --format "{{.CreatedBy}} : {{.Size}}" mon-image:1.0
# Outil visuel (à installer) : dive
# https://github.com/wagoodman/dive
```

## 21. Le .dockerignore : ne pas envoyer n'importe quoi au daemon

Le **contexte de build** (le dossier envoyé au daemon) doit être minimal :
chaque fichier inutile ralentit le build et peut fuiter dans l'image.

```gitignore
# .dockerignore
.git
.gitignore
.github/
.vscode/
.idea/
node_modules/
__pycache__/
*.pyc
.pytest_cache/
venv/
.env
.env.*
*.log
Dockerfile*
docker-compose*.yml
README.md
tests/
docs/
```

```bash
# Vérifier ce qui part dans le contexte (buildx)
docker build --progress=plain -t test . 2>&1 | grep -i "transferring context"
```

## 22. Build : cache, tags, bonnes habitudes

```bash
# Build simple avec tag versionné
docker build -t registry.interne:5000/mon-app:1.4.2 .

# Plusieurs tags d'un coup
docker build -t mon-app:1.4.2 -t mon-app:latest .

# Sans utiliser le cache (build 100 % propre, ex : release)
docker build --no-cache -t mon-app:1.4.2 .

# Passer un ARG
docker build --build-arg APP_VERSION=1.4.2 -t mon-app:1.4.2 .

# BuildKit (par défaut sur les versions récentes) : parallélise et sécurise
DOCKER_BUILDKIT=1 docker build -t mon-app:1.4.2 .

# Voir les logs détaillés d'un build qui échoue
docker build --progress=plain -t mon-app:1.4.2 . 2>&1 | tail -60
```

Convention de tags recommandée :

- `1.4.2` : version exacte (déploiement reproductible)
- `1.4` : dernière 1.4.x (mises à jour mineures auto)
- `latest` : uniquement pour le dev local, **jamais** en prod

## 23. Cas pratique : Dockerfile d'une API Python (FastAPI) production-ready

```dockerfile
# Dockerfile
FROM python:3.12-slim AS builder
WORKDIR /app
COPY requirements.txt .
RUN pip install --no-cache-dir --prefix=/install -r requirements.txt

FROM python:3.12-slim
ENV PYTHONUNBUFFERED=1 \
    PYTHONDONTWRITEBYTECODE=1
WORKDIR /app
COPY --from=builder /install /usr/local
COPY ./app ./app
RUN useradd --create-home --uid 1001 appuser \
    && chown -R appuser:appuser /app
USER appuser
EXPOSE 8000
HEALTHCHECK --interval=30s --timeout=5s --retries=3 \
  CMD python -c "import urllib.request; urllib.request.urlopen('http://localhost:8000/health')" || exit 1
CMD ["uvicorn", "app.main:app", "--host", "0.0.0.0", "--port", "8000"]
```

```txt
# requirements.txt
fastapi==0.115.0
uvicorn[standard]==0.30.0
```

```bash
docker build -t mon-api:1.0.0 .
docker run -d --name api --user 1001 -p 8000:8000 mon-api:1.0.0
curl http://localhost:8000/health
docker inspect --format '{{.State.Health.Status}}' api
```

## 24. Cas pratique : Dockerfile frontend (build statique + nginx)

```dockerfile
# Build du frontend (ex : Vite/React)
FROM node:20-alpine AS builder
WORKDIR /app
COPY package*.json ./
RUN npm ci
COPY . .
RUN npm run build   # produit /app/dist

# Serveur statique léger
FROM nginx:1.27-alpine
COPY --from=builder /app/dist /usr/share/nginx/html
COPY nginx.conf /etc/nginx/conf.d/default.conf
EXPOSE 80
```

```nginx
# nginx.conf : SPA avec fallback + cache des assets
server {
    listen 80;
    root /usr/share/nginx/html;
    location / {
        try_files $uri $uri/ /index.html;
    }
    location /assets/ {
        expires 1y;
        add_header Cache-Control "public, immutable";
    }
}
```

## 25. Docker Compose : le chef d'orchestre local

Compose décrit **toute une stack** (services, réseaux, volumes) dans un fichier
`compose.yaml` (ou `docker-compose.yml`), puis la pilote d'un bloc.

```bash
docker compose up -d        # lance la stack en tâche de fond
docker compose ps           # état des services
docker compose logs -f      # logs de tous les services
docker compose logs -f web  # logs d'un seul service
docker compose down         # arrête + supprime conteneurs et réseaux
docker compose down -v       # + supprime les VOLUMES (données ! prudent)
docker compose pull         # met à jour les images de la stack
docker compose up -d --build  # rebuild les images locales puis relance
```

> Le nom du projet (préfixe des conteneurs/réseaux/volumes) vient du dossier
> ou de `-p monprojet` / `name: monprojet` en tête du YAML.

## 26. Compose : anatomie d'un fichier

```yaml
name: demo

services:
  web:
    image: nginx:1.27-alpine
    container_name: demo-web
    ports:
      - "8080:80"
    volumes:
      - ./site:/usr/share/nginx/html:ro
    networks:
      - front
    restart: unless-stopped
    depends_on:
      api:
        condition: service_healthy

  api:
    build: ./api
    environment:
      DATABASE_URL: "postgres://app:secret@db:5432/appdb"
    networks:
      - front
      - back
    healthcheck:
      test: ["CMD", "curl", "-f", "http://localhost:8000/health"]
      interval: 30s
      timeout: 5s
      retries: 3
      start_period: 10s
    restart: unless-stopped

