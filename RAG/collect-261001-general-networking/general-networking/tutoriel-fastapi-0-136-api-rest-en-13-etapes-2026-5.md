---
id: collect-261001-general-networking/general-networking/tutoriel-fastapi-0-136-api-rest-en-13-etapes-2026-5
title: "Installation d'uv (Astral) – outil officiel recommandé en 2026"
domain: general-networking
role: reference
task: reference
actors: ["AWS"]
dates: []
keywords: ["aws"]
source: docs/RAG/collect-261001-general-networking/tutoriel-fastapi-0-136-api-rest-en-13-etapes-2026.md
source_anchor: ""
source_lines: [414, 579]
sha256: 2d84d8002ecbec562e9cf607a22655dd87719065245bf2ac0aa93ab7a78aac36
---

# Installation d'uv (Astral) – outil officiel recommandé en 2026

```
# tests/conftest.py
import pytest
import pytest_asyncio
from httpx import AsyncClient, ASGITransport
from sqlmodel import SQLModel
from sqlalchemy.ext.asyncio import create_async_engine
from sqlmodel.ext.asyncio.session import AsyncSession
from app.main import app
from app.db.session import get_session
TEST_DB_URL = "postgresql+asyncpg://postgres:postgres@localhost:5432/test_db"
@pytest_asyncio.fixture
async def session():
    engine = create_async_engine(TEST_DB_URL)
    async with engine.begin() as conn:
        await conn.run_sync(SQLModel.metadata.create_all)
    async with AsyncSession(engine) as s:
        yield s
    async with engine.begin() as conn:
        await conn.run_sync(SQLModel.metadata.drop_all)
    await engine.dispose()
@pytest_asyncio.fixture
async def client(session):
    async def override():
        yield session
    app.dependency_overrides[get_session] = override
    transport = ASGITransport(app=app)
    async with AsyncClient(transport=transport, base_url="http://test") as c:
        yield c
    app.dependency_overrides.clear()
```
```
# tests/test_articles.py
import pytest
@pytest.mark.asyncio
async def test_create_and_read_article(client):
    # Inscription
    r = await client.post("/auth/register", json={
        "email": "[email protected]",
        "password": "MotDePasse2026!"
    })
    assert r.status_code == 201
    token = r.json()["access_token"]
    headers = {"Authorization": f"Bearer {token}"}
    # Création
    r = await client.post(
        "/articles/",
        json={"title": "test", "content": "contenu suffisant pour passer", "published": True},
        headers=headers,
    )
    assert r.status_code == 201
    article_id = r.json()["id"]
    # Lecture
    r = await client.get(f"/articles/{article_id}")
    assert r.status_code == 200
    assert r.json()["title"] == "Test"  # validator capitalize
@pytest.mark.asyncio
async def test_unauthorized_create(client):
    r = await client.post("/articles/", json={"title": "x", "content": "y"})
    assert r.status_code == 401
```
Lancez la suite avec `pytest -v --asyncio-mode=auto`. Pour mesurer la couverture, ajoutez `pytest-cov` et exécutez `pytest --cov=app --cov-report=term-missing`. L’objectif raisonnable pour un projet FastAPI est **85 %+ de couverture de lignes** et 100 % sur les modules `core/security` et `api/deps`. La documentation officielle FastAPI annonce 100 % de coverage sur le framework lui-même, ce qui est un excellent indicateur de stabilité pour la production.

## Étape 11 : Conteneuriser l’application avec Docker multi-stage

Pour la production, utilisez un Dockerfile multi-stage qui sépare la phase de build (avec uv et les outils de compilation) de l’image runtime minimale. Cette approche divise la taille finale par 3 à 5 selon les dépendances. L’image officielle `python:3.12-slim` pèse 130 Mo, contre 950 Mo pour `python:3.12` standard.

```
# Dockerfile
# Étape 1 : builder
FROM python:3.12-slim AS builder
ENV UV_COMPILE_BYTECODE=1 UV_LINK_MODE=copy
COPY --from=ghcr.io/astral-sh/uv:0.5.11 /uv /usr/local/bin/uv
WORKDIR /app
COPY pyproject.toml uv.lock ./
RUN uv sync --frozen --no-dev --no-install-project
COPY . .
RUN uv sync --frozen --no-dev
# Étape 2 : runtime minimal
FROM python:3.12-slim AS runtime
RUN apt-get update && apt-get install -y --no-install-recommends \
    libpq5 ca-certificates curl \
    && rm -rf /var/lib/apt/lists/*
WORKDIR /app
COPY --from=builder /app /app
ENV PATH="/app/.venv/bin:$PATH" \
    PYTHONUNBUFFERED=1 \
    PYTHONDONTWRITEBYTECODE=1
RUN useradd --create-home --uid 1001 fastapi
USER fastapi
EXPOSE 8000
HEALTHCHECK --interval=30s --timeout=5s --start-period=10s \
    CMD curl -f http://localhost:8000/health || exit 1
CMD ["uvicorn", "app.main:app", "--host", "0.0.0.0", "--port", "8000", "--workers", "4"]
```
```
# docker-compose.yml
services:
  api:
    build: .
    ports:
      - "8000:8000"
    environment:
      DATABASE_URL: postgresql+asyncpg://postgres:postgres@db:5432/articles
      SECRET_KEY: change-me-in-production
    depends_on:
      db:
        condition: service_healthy
    restart: unless-stopped
  db:
    image: postgres:17.4-alpine
    environment:
      POSTGRES_USER: postgres
      POSTGRES_PASSWORD: postgres
      POSTGRES_DB: articles
    volumes:
      - pgdata:/var/lib/postgresql/data
    healthcheck:
      test: ["CMD", "pg_isready", "-U", "postgres"]
      interval: 5s
      timeout: 3s
      retries: 5
  redis:
    image: redis:7.4-alpine
    restart: unless-stopped
volumes:
  pgdata:
```
Le nombre de workers Uvicorn (`--workers 4`) doit correspondre à `2 × nombre de CPU + 1` selon la règle de Gunicorn, dans la limite de la RAM disponible. Sur une instance Scaleway DEV1-M (3 vCPU, 4 Go), 4 workers consommant 250 Mo chacun saturent la mémoire ; restez à 3 workers pour laisser une marge. En Kubernetes, fixez `--workers 1` et laissez l’autoscaler horizontal HPA répliquer les pods : c’est plus propre et plus observable.

## Étape 12 : Observabilité avec Prometheus, OpenTelemetry et logs structurés

Une API en production sans télémétrie est aveugle. Trois piliers minimaux : métriques Prometheus pour le RED (Rate, Errors, Duration), traces OpenTelemetry pour les chemins critiques inter-services, et logs JSON structurés pour le SIEM. La librairie `prometheus-fastapi-instrumentator` expose automatiquement `/metrics` avec les histogrammes par route.

```
# app/main.py – observabilité
from prometheus_fastapi_instrumentator import Instrumentator
from opentelemetry import trace
from opentelemetry.instrumentation.fastapi import FastAPIInstrumentor
from opentelemetry.exporter.otlp.proto.grpc.trace_exporter import OTLPSpanExporter
from opentelemetry.sdk.trace import TracerProvider
from opentelemetry.sdk.trace.export import BatchSpanProcessor
# Métriques Prometheus
Instrumentator().instrument(app).expose(app, endpoint="/metrics")
# Traces OpenTelemetry
trace.set_tracer_provider(TracerProvider())
trace.get_tracer_provider().add_span_processor(
    BatchSpanProcessor(OTLPSpanExporter(endpoint="http://otel-collector:4317"))
)
FastAPIInstrumentor.instrument_app(app)
import structlog
structlog.configure(
    processors=[
        structlog.processors.TimeStamper(fmt="iso"),
        structlog.processors.add_log_level,
        structlog.processors.JSONRenderer(),
    ],
)
logger = structlog.get_logger()
```
Sur Grafana, importez le dashboard *FastAPI Observability* (ID 18739) qui affiche RPS par endpoint, latence p50/p95/p99 et taux d’erreur 5xx. Pour les traces distribuées, l’OTLP exporter envoie vers un OpenTelemetry Collector qui dispatche ensuite vers Jaeger, Tempo ou Datadog. La rétention par défaut de 7 jours sur Tempo coûte environ 0,15 €/Go ingéré sur Grafana Cloud Pro selon les tarifs publics 2026, ce qui reste compétitif face à Datadog APM (estimé entre 31 et 36 € par host/mois selon l’engagement annuel).

## Étape 13 : Déployer en production sur Scaleway, OVH ou AWS

Pour un déploiement européen RGPD-friendly, trois options dominent en 2026 : Scaleway Serverless Containers (Paris, Amsterdam), OVH Public Cloud Managed Kubernetes (Gravelines, Strasbourg) et Clever Cloud (Paris, Montréal, Sydney). Pour une mise en ligne rapide d’un MVP, Scaleway Serverless Containers facture à l’usage (CPU.ms et RAM.s) et démarre vos containers en moins de 2 secondes au cold start.

