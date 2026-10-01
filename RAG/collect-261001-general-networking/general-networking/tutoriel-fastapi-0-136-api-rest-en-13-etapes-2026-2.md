---
id: collect-261001-general-networking/general-networking/tutoriel-fastapi-0-136-api-rest-en-13-etapes-2026-2
title: "Installation d'uv (Astral) – outil officiel recommandé en 2026"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/tutoriel-fastapi-0-136-api-rest-en-13-etapes-2026.md
source_anchor: ""
source_lines: [53, 167]
sha256: 0d719aadea0a8958b1c18046d809babc7e6df3551fc966c732f0adc9dce9a46a
---

# Installation d'uv (Astral) – outil officiel recommandé en 2026

```
# Dépendances principales – versions vérifiées avril 2026
uv add 'fastapi[standard]==0.136.1'
uv add 'uvicorn[standard]==0.34.0'
uv add 'pydantic==2.11.0'
uv add 'pydantic-settings==2.7.0'
uv add 'sqlmodel==0.0.24'
uv add 'asyncpg==0.30.0'
uv add 'python-jose[cryptography]==3.3.0'
uv add 'passlib[bcrypt]==1.7.4'
uv add 'python-multipart==0.0.18'
# Dépendances de développement
uv add --dev 'pytest==8.3.4' 'pytest-asyncio==0.25.0' 'httpx==0.28.1'
uv add --dev 'ruff==0.8.4' 'mypy==1.14.0'
# Génération du fichier de lock pour la prod
uv lock
```
Le suffixe `[standard]` sur FastAPI installe automatiquement les extras les plus utilisés : `email-validator` pour les EmailStr Pydantic, `itsdangerous` pour les sessions, `jinja2` pour le rendu HTML et `fastapi-cli` pour la commande `fastapi dev` introduite en 2024. En août 2026, le même tutoriel officiel de FastAPI recommande désormais de lancer le serveur via `uv run fastapi dev main.py`, qui remplace avantageusement `uvicorn main:app --reload` en gérant automatiquement le rechargement, les logs et le watchdog. Vérifiez votre installation avec `fastapi --version` qui doit retourner `FastAPI CLI version: 0.0.7` ou supérieure.

## Étape 3 : Créer le squelette de l’application FastAPI

Adoptez dès le départ une structure de projet inspirée de Domain-Driven Design : un répertoire `app/` avec sous-modules `api/`, `models/`, `db/`, `core/` et `tests/`. Cette organisation, popularisée par le template *full-stack-fastapi-template* officiel de tiangolo, scale jusqu’à plusieurs dizaines de routes sans devenir illisible.

```
tutoriel-fastapi-2026/
├── app/
│   ├── __init__.py
│   ├── main.py              # Instance FastAPI + montage des routers
│   ├── core/
│   │   ├── config.py        # Settings via pydantic-settings
│   │   └── security.py      # JWT + hashing
│   ├── db/
│   │   ├── session.py       # AsyncSession SQLModel
│   │   └── init_db.py
│   ├── models/
│   │   └── article.py       # Modèles SQLModel
│   ├── schemas/
│   │   └── article.py       # Schémas Pydantic d'entrée/sortie
│   └── api/
│       ├── deps.py          # Dépendances communes
│       └── routes/
│           ├── articles.py
│           └── auth.py
├── tests/
├── docker-compose.yml
├── Dockerfile
├── pyproject.toml
└── .env
```
Créez maintenant le fichier `app/main.py` avec une route racine minimale. FastAPI 0.136.1 introduit un nouveau pattern de configuration via `FastAPI(lifespan=...)` qui remplace les anciens événements `@app.on_event("startup")` dépréciés depuis 0.110.0. Cette approche utilise un context manager asynchrone pour gérer le cycle de vie des connexions DB, des pools Redis ou des modèles ML.

```
# app/main.py
from contextlib import asynccontextmanager
from fastapi import FastAPI
from app.api.routes import articles, auth
from app.db.session import engine
from sqlmodel import SQLModel
@asynccontextmanager
async def lifespan(app: FastAPI):
    # Startup : créer les tables, ouvrir les pools
    async with engine.begin() as conn:
        await conn.run_sync(SQLModel.metadata.create_all)
    yield
    # Shutdown : fermer proprement
    await engine.dispose()
app = FastAPI(
    title="API Articles 2026",
    version="1.0.0",
    description="Tutoriel FastAPI en 13 étapes",
    lifespan=lifespan,
)
app.include_router(auth.router, prefix="/auth", tags=["auth"])
app.include_router(articles.router, prefix="/articles", tags=["articles"])
@app.get("/")
async def root() -> dict[str, str]:
    return {"message": "FastAPI 0.136 – Bonjour Paris"}
```
Lancez le serveur avec `fastapi dev app/main.py`. Ouvrez `http://127.0.0.1:8000/docs` dans votre navigateur : Swagger UI s’affiche automatiquement avec votre route racine. La documentation alternative ReDoc est disponible sur `/redoc`. Les deux interfaces consomment l’OpenAPI 3.1 généré dynamiquement à `/openapi.json`, que vous pouvez importer dans Postman, Insomnia ou Bruno pour générer un client.

## Étape 4 : Définir les modèles Pydantic v2 pour la validation

Pydantic v2, sortie en juin 2023 et obligatoire avec FastAPI 0.128+, repose sur un cœur Rust nommé *pydantic-core*. Les performances sont multipliées par 5 à 50 selon les opérations, et le typage devient strict par défaut. Pour ce tutoriel, nous séparons les schémas Pydantic (entrée/sortie API) des modèles SQLModel (persistance DB) afin d’éviter les fuites de champs sensibles comme les hashes de mots de passe.

```
# app/schemas/article.py
from datetime import datetime
from pydantic import BaseModel, ConfigDict, Field, field_validator
class ArticleBase(BaseModel):
    title: str = Field(min_length=3, max_length=200)
    content: str = Field(min_length=10)
    published: bool = False
    @field_validator("title")
    @classmethod
    def title_capitalize(cls, v: str) -> str:
        return v.strip().capitalize()
class ArticleCreate(ArticleBase):
    pass
class ArticleUpdate(BaseModel):
    title: str | None = Field(default=None, min_length=3, max_length=200)
    content: str | None = None
    published: bool | None = None
class ArticleRead(ArticleBase):
    model_config = ConfigDict(from_attributes=True)
    id: int
    author_id: int
    created_at: datetime
    updated_at: datetime
```
Notez l’usage de `Field()` avec `min_length` et `max_length` : ces contraintes apparaissent automatiquement dans Swagger UI et déclenchent une réponse HTTP 422 avec un message JSON détaillé en cas de violation. Le décorateur `@field_validator` remplace `@validator` de Pydantic v1. La syntaxe `str | None` exige Python 3.10+ et est strictement équivalente à `Optional[str]` de Python 3.9. Le `ConfigDict(from_attributes=True)` active le mode ORM (anciennement `orm_mode = True`) pour convertir un objet SQLModel en réponse Pydantic.

## Étape 5 : Configurer la base PostgreSQL avec SQLModel et asyncpg

SQLModel, créé par Sebastián Ramírez en 2021, fusionne SQLAlchemy 2.0 et Pydantic en une seule classe. Vous définissez le schéma une fois et obtenez à la fois un modèle ORM et un schéma de validation. Pour la version 0.0.24 sortie en mars 2026, la compatibilité avec Pydantic v2 est complète et les sessions asynchrones sont supportées de bout en bout via asyncpg.

