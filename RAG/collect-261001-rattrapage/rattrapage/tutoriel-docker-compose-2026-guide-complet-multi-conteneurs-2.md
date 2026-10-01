---
id: collect-261001-rattrapage/rattrapage/tutoriel-docker-compose-2026-guide-complet-multi-conteneurs-2
title: "Vérifier les versions installées"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["attention", "decode"]
source: docs/RAG/collect-261001-rattrapage/tutoriel-docker-compose-2026-guide-complet-multi-conteneurs.md
source_anchor: ""
source_lines: [76, 259]
sha256: 61ee2014ffc5ed5bfc25bc0f1a39dfe690fb4ddd8f2bf65a2b38de3baa9da696
---

# Vérifier les versions installées

Commençons par le fichier `api/requirements.txt` qui liste les dépendances Python :

```
# api/requirements.txt
fastapi==0.115.6
uvicorn[standard]==0.34.0
psycopg2-binary==2.9.10
redis==5.2.1
pydantic==2.10.4
```
Ensuite, créez l’API dans `api/main.py`. Cette application se connecte à PostgreSQL et Redis, expose des endpoints CRUD et inclut un health check – un pattern essentiel pour Docker Compose :

```
# api/main.py
import os
import redis
import psycopg2
from fastapi import FastAPI, HTTPException
from pydantic import BaseModel
from contextlib import asynccontextmanager
# Configuration via variables d'environnement
DB_HOST = os.getenv("DB_HOST", "db")
DB_NAME = os.getenv("POSTGRES_DB", "appdb")
DB_USER = os.getenv("POSTGRES_USER", "appuser")
DB_PASS = os.getenv("POSTGRES_PASSWORD", "changeme")
REDIS_HOST = os.getenv("REDIS_HOST", "redis")
def get_db_connection():
    return psycopg2.connect(
        host=DB_HOST, database=DB_NAME,
        user=DB_USER, password=DB_PASS
    )
redis_client = redis.Redis(host=REDIS_HOST, port=6379, decode_responses=True)
@asynccontextmanager
async def lifespan(app: FastAPI):
    # Initialisation au démarrage
    conn = get_db_connection()
    cur = conn.cursor()
    cur.execute("""
        CREATE TABLE IF NOT EXISTS items (
            id SERIAL PRIMARY KEY,
            name VARCHAR(100) NOT NULL,
            description TEXT,
            created_at TIMESTAMP DEFAULT NOW()
        )
    """)
    conn.commit()
    cur.close()
    conn.close()
    yield
app = FastAPI(title="Compose Tutorial API", version="1.0.0", lifespan=lifespan)
class ItemCreate(BaseModel):
    name: str
    description: str = ""
@app.get("/health")
def health_check():
    """Endpoint de santé pour Docker Compose"""
    try:
        conn = get_db_connection()
        conn.close()
        redis_client.ping()
        return {"status": "healthy", "db": "ok", "cache": "ok"}
    except Exception as e:
        raise HTTPException(status_code=503, detail=str(e))
@app.get("/items")
def list_items():
    # Vérifier le cache Redis
    cached = redis_client.get("items_list")
    if cached:
        import json
        return {"source": "cache", "items": json.loads(cached)}
    
    conn = get_db_connection()
    cur = conn.cursor()
    cur.execute("SELECT id, name, description FROM items ORDER BY id DESC LIMIT 50")
    rows = cur.fetchall()
    cur.close()
    conn.close()
    
    items = [{"id": r[0], "name": r[1], "description": r[2]} for r in rows]
    import json
    redis_client.setex("items_list", 30, json.dumps(items))
    return {"source": "database", "items": items}
@app.post("/items", status_code=201)
def create_item(item: ItemCreate):
    conn = get_db_connection()
    cur = conn.cursor()
    cur.execute(
        "INSERT INTO items (name, description) VALUES (%s, %s) RETURNING id",
        (item.name, item.description)
    )
    item_id = cur.fetchone()[0]
    conn.commit()
    cur.close()
    conn.close()
    redis_client.delete("items_list")  # Invalider le cache
    return {"id": item_id, "name": item.name}
```
Créez maintenant le Dockerfile multi-étapes pour l’API, une approche qui optimise la taille de l’image finale :

```
# api/Dockerfile
FROM python:3.12-slim AS base
WORKDIR /app
# Installer les dépendances système pour psycopg2
RUN apt-get update && apt-get install -y --no-install-recommends \
    libpq-dev gcc \
    && rm -rf /var/lib/apt/lists/*
COPY requirements.txt .
RUN pip install --no-cache-dir -r requirements.txt
COPY main.py .
# Créer un utilisateur non-root pour la sécurité
RUN useradd -m appuser
USER appuser
EXPOSE 8000
CMD ["uvicorn", "main:app", "--host", "0.0.0.0", "--port", "8000"]
```
**Piège courant n°2 :** N’utilisez jamais l’image `python:3.12` complète en production. L’image `slim` pèse environ 150 Mo contre 1 Go pour l’image complète. Pour des images encore plus légères, considérez `python:3.12-alpine`, mais attention : Alpine utilise `musl` au lieu de `glibc`, ce qui peut causer des incompatibilités avec certaines bibliothèques comme `psycopg2`. C’est pourquoi nous utilisons `slim` dans ce tutoriel.

## Étape 3 : Configurer PostgreSQL et le Script d’Initialisation

PostgreSQL est la base de données relationnelle la plus utilisée avec Docker Compose en 2026, grâce à sa robustesse et ses fonctionnalités avancées. Notre configuration utilise PostgreSQL 16, la version LTS la plus récente, avec un script d’initialisation automatique qui crée le schéma de base de données au premier lancement.

Créez le fichier `db/init.sql` qui sera automatiquement exécuté par PostgreSQL lors de la première initialisation du conteneur :

```
-- db/init.sql
-- Script d'initialisation PostgreSQL pour Docker Compose
-- Créer une extension utile
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
-- Table principale
CREATE TABLE IF NOT EXISTS items (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    description TEXT DEFAULT '',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);
-- Index pour les performances
CREATE INDEX IF NOT EXISTS idx_items_name ON items(name);
CREATE INDEX IF NOT EXISTS idx_items_created ON items(created_at DESC);
-- Données de test
INSERT INTO items (name, description) VALUES
    ('Premier élément', 'Créé automatiquement par le script init'),
    ('Deuxième élément', 'Données de démonstration Docker Compose'),
    ('Troisième élément', 'Tutoriel tech-insider.org 2026');
-- Fonction de mise à jour automatique du timestamp
CREATE OR REPLACE FUNCTION update_modified_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE TRIGGER update_items_modtime
    BEFORE UPDATE ON items
    FOR EACH ROW EXECUTE FUNCTION update_modified_column();
```
Le répertoire `/docker-entrypoint-initdb.d/` de l’image officielle PostgreSQL exécute automatiquement tous les fichiers `.sql` et `.sh` qu’il contient lors de la première initialisation. Cela signifie que ce script ne s’exécutera que si le volume de données est vide – un comportement qui surprend souvent les développeurs qui modifient leur script et ne voient pas les changements appliqués.

**Piège courant n°3 :** Si vous modifiez le script `init.sql` après avoir déjà lancé les conteneurs, les changements ne seront pas appliqués car le volume PostgreSQL contient déjà des données. Vous devez supprimer le volume avec `docker compose down -v` puis relancer `docker compose up`. Attention : cette opération détruit toutes les données existantes.

## Étape 4 : Configurer Nginx comme Reverse Proxy

Nginx agit comme point d’entrée unique pour notre application, distribuant les requêtes vers l’API FastAPI. Dans un environnement de production, Nginx gère également la terminaison SSL/TLS, la mise en cache statique et la limitation de débit. Notre configuration Docker Compose simplifie cette mise en place en permettant à Nginx de résoudre automatiquement le nom du service API via le DNS interne de Docker.

```
# nginx/nginx.conf
upstream api_backend {
    server api:8000;
}
server {
    listen 80;
    server_name localhost;
    # Logs
    access_log /var/log/nginx/access.log;
    error_log /var/log/nginx/error.log;
    # Proxy vers l'API FastAPI
    location /api/ {
        proxy_pass http://api_backend/;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
        
