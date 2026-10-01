---
id: collect-261001-rattrapage/rattrapage/tutoriel-docker-compose-2026-stack-prod-en-13-etapes-2
title: "Mise à jour des dépôts et installation des dépendances"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["benchmark", "decode"]
source: docs/RAG/collect-261001-rattrapage/tutoriel-docker-compose-2026-stack-prod-en-13-etapes.md
source_anchor: ""
source_lines: [60, 203]
sha256: cfa3d2030558ad8af9405548c550e06340bdff613bbde6c68eb065a180c3aa70
---

# Mise à jour des dépôts et installation des dépendances

Le **Compose Specification** définit un format YAML standardisé pour décrire une stack multi-conteneurs. Il est maintenu publiquement sur GitHub (compose-spec/compose-spec) et adopté par Docker, Podman, Nerdctl, et même certains outils Kubernetes ; au 11 septembre 2026, ce dépôt totalise 2 727 lignes de Dockerfile et 786 étoiles, un signe de l’évolution active de la spécification. La bibliothèque associée **compose-go**, qui alimente les outils tiers écrits en Go, a franchi la version 2.10.0 le 25 novembre 2025, avec des builds ultérieurs enrichissant encore son API. Un fichier Compose typique décrit trois sections principales : `services` (les conteneurs), `volumes` (le stockage persistant), et `networks` (les réseaux virtuels). Optionnellement, on retrouve `secrets`, `configs` et `profiles`.

Le moteur reconnaît plusieurs noms de fichier par défaut, dans cet ordre : `compose.yaml`, `compose.yml`, `docker-compose.yaml`, `docker-compose.yml`. La convention moderne, recommandée par la spec, est d’utiliser **`compose.yaml`**. Toutefois, `docker-compose.yml` reste majoritaire dans la pratique, notamment pour la rétrocompatibilité avec d’anciens outils CI/CD – c’est d’ailleurs le choix fait par le guide « Docker Compose Tutorial (Beginner to Production 2026) » d’ITSourceCode, publié en juillet 2026, qui standardise lui aussi un unique `docker-compose.yml` pour orchestrer des stacks multi-services. Nous utiliserons ce dernier dans le tutoriel.

**Piège n° 1 :** Le champ `version:` en tête de fichier est *déprécié* depuis Compose v2.27 et purement et simplement *ignoré* depuis v2.30. Ne l’écrivez plus. Si vous reprenez un ancien fichier, supprimez la ligne ; cela évite des warnings inutiles dans les pipelines.

## Étape 3 : Créer la Structure du Projet

Nous allons construire une application web exemple : une API **FastAPI** en Python qui stocke des données dans **PostgreSQL**, met en cache des réponses dans **Redis**, est exposée via **Nginx** en reverse proxy, et monitorée par **Prometheus** et **Grafana**. Créons d’abord l’arborescence.

```
mkdir -p ~/compose-prod-stack && cd ~/compose-prod-stack
mkdir -p app nginx postgres prometheus grafana/provisioning secrets
# Structure cible :
tree -L 2
# .
# ├── app/
# │   ├── Dockerfile
# │   ├── main.py
# │   └── requirements.txt
# ├── nginx/
# │   └── default.conf
# ├── postgres/
# │   └── init.sql
# ├── prometheus/
# │   └── prometheus.yml
# ├── grafana/
# │   └── provisioning/
# ├── secrets/
# │   ├── db_password.txt
# │   └── grafana_admin.txt
# ├── .env
# ├── .gitignore
# ├── compose.override.yml
# └── docker-compose.yml
```
Le répertoire `secrets/` doit être ajouté à `.gitignore` immédiatement pour éviter tout commit accidentel de mots de passe. Créez ce fichier avec : `echo -e "secrets/\n.env\n*.log\n__pycache__/\n.venv/" > .gitignore`.

## Étape 4 : Écrire l’Application FastAPI et son Dockerfile

L’application expose deux endpoints : `GET /items` qui liste les produits depuis PostgreSQL avec mise en cache Redis, et `POST /items` qui insère un produit. Le code est volontairement minimal pour rester pédagogique.

```
# app/requirements.txt
fastapi==0.115.6
uvicorn[standard]==0.32.1
asyncpg==0.30.0
redis==5.2.1
pydantic==2.10.3
prometheus-fastapi-instrumentator==7.0.2
```
```
# app/main.py
import os, json
from contextlib import asynccontextmanager
import asyncpg, redis.asyncio as redis
from fastapi import FastAPI, HTTPException
from pydantic import BaseModel
from prometheus_fastapi_instrumentator import Instrumentator
DB_PASSWORD = open("/run/secrets/db_password").read().strip()
DSN = f"postgres://app:{DB_PASSWORD}@db:5432/appdb"
REDIS_URL = os.getenv("REDIS_URL", "redis://cache:6379/0")
class Item(BaseModel):
    name: str
    price: float
@asynccontextmanager
async def lifespan(app: FastAPI):
    app.state.pool = await asyncpg.create_pool(DSN, min_size=2, max_size=10)
    app.state.cache = redis.from_url(REDIS_URL, decode_responses=True)
    yield
    await app.state.pool.close()
    await app.state.cache.aclose()
app = FastAPI(lifespan=lifespan, title="Compose Demo API")
Instrumentator().instrument(app).expose(app)
@app.get("/health")
async def health():
    return {"status": "ok"}
@app.get("/items")
async def list_items():
    cached = await app.state.cache.get("items:all")
    if cached:
        return {"source": "cache", "items": json.loads(cached)}
    rows = await app.state.pool.fetch("SELECT id, name, price FROM items")
    items = [dict(r) for r in rows]
    await app.state.cache.setex("items:all", 30, json.dumps(items))
    return {"source": "db", "items": items}
@app.post("/items", status_code=201)
async def create_item(item: Item):
    row = await app.state.pool.fetchrow(
        "INSERT INTO items(name, price) VALUES($1, $2) RETURNING id",
        item.name, item.price
    )
    await app.state.cache.delete("items:all")
    return {"id": row["id"], **item.model_dump()}
```
Le **Dockerfile multi-stage** sépare la compilation des dépendances de l’image finale, ce qui réduit la taille et la surface d’attaque – une pratique désormais généralisée au-delà de Python : le guide « Docker Compose Tutorial: Complete Guide » de Luca Berton, publié en avril 2026, illustre le même schéma multi-stage avec un exemple Node.js 22 pour des stacks orchestrées par Compose. L’image finale tourne sous un utilisateur non-root, conformément aux bonnes pratiques CIS Docker Benchmark.

```
# app/Dockerfile
FROM python:3.13-slim AS builder
WORKDIR /build
COPY requirements.txt .
RUN pip install --no-cache-dir --user -r requirements.txt
FROM python:3.13-slim AS runtime
RUN useradd --create-home --shell /bin/bash appuser \
    && apt-get update && apt-get install -y --no-install-recommends curl \
    && rm -rf /var/lib/apt/lists/*
USER appuser
WORKDIR /home/appuser/app
COPY --from=builder /root/.local /home/appuser/.local
COPY --chown=appuser:appuser . .
ENV PATH="/home/appuser/.local/bin:$PATH" \
    PYTHONUNBUFFERED=1 \
    PYTHONDONTWRITEBYTECODE=1
EXPOSE 8000
HEALTHCHECK --interval=30s --timeout=5s --start-period=10s --retries=3 \
  CMD curl -f http://localhost:8000/health || exit 1
CMD ["uvicorn", "main:app", "--host", "0.0.0.0", "--port", "8000", "--workers", "2"]
```
## Étape 5 : Configurer PostgreSQL avec Initialisation Automatique

L’image officielle `postgres:17-alpine` exécute automatiquement tout fichier `.sql` ou `.sh` placé dans `/docker-entrypoint-initdb.d/` au premier démarrage. Nous y déposerons un script qui crée la table `items` et l’utilisateur applicatif.

```
-- postgres/init.sql
CREATE USER app WITH PASSWORD 'CHANGE_ME_AT_RUNTIME';
CREATE DATABASE appdb OWNER app;
\connect appdb;
CREATE TABLE items (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    price NUMERIC(10, 2) NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW()
);
CREATE INDEX idx_items_name ON items(name);
GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO app;
GRANT USAGE, SELECT ON ALL SEQUENCES IN SCHEMA public TO app;
```
Le mot de passe en dur n’est qu’un placeholder : Compose le réécrira via *secrets* à l’étape 7. Pour la sauvegarde, nous monterons le répertoire `/var/lib/postgresql/data` sur un **named volume** (et non un bind mount), ce qui simplifie la gestion et améliore les performances I/O sous Docker Desktop.

## Étape 6 : Rédiger le docker-compose.yml de Production

Voici le fichier `docker-compose.yml` central. Lisez-le attentivement : chaque section illustre une bonne pratique de production. Notez l’absence du champ `version:`, l’utilisation de `profiles:` pour activer ou non le monitoring, les `healthcheck:` et `depends_on: condition: service_healthy` qui garantissent l’ordre de démarrage correct.

