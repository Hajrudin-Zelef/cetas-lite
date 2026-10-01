---
id: collect-261001-general-networking/general-networking/tutoriel-fastapi-0-136-api-rest-en-13-etapes-2026-3
title: "Installation d'uv (Astral) – outil officiel recommandé en 2026"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["cost"]
source: docs/RAG/collect-261001-general-networking/tutoriel-fastapi-0-136-api-rest-en-13-etapes-2026.md
source_anchor: ""
source_lines: [168, 282]
sha256: f415bcff5473525d5897bcd700a59451a10c9d7ac8d38cda78a3646de8160598
---

# Installation d'uv (Astral) – outil officiel recommandé en 2026

```
# app/db/session.py
from collections.abc import AsyncGenerator
from sqlalchemy.ext.asyncio import AsyncSession, create_async_engine
from sqlmodel.ext.asyncio.session import AsyncSession as SQLModelSession
from app.core.config import settings
engine = create_async_engine(
    settings.DATABASE_URL,
    echo=False,
    pool_size=20,
    max_overflow=10,
    pool_pre_ping=True,
)
async def get_session() -> AsyncGenerator[SQLModelSession, None]:
    async with SQLModelSession(engine) as session:
        yield session
```
```
# app/models/article.py
from datetime import datetime, UTC
from sqlmodel import Field, SQLModel, Relationship
class Article(SQLModel, table=True):
    id: int | None = Field(default=None, primary_key=True)
    title: str = Field(index=True, max_length=200)
    content: str
    published: bool = Field(default=False)
    author_id: int = Field(foreign_key="user.id", index=True)
    created_at: datetime = Field(default_factory=lambda: datetime.now(UTC))
    updated_at: datetime = Field(default_factory=lambda: datetime.now(UTC))
class User(SQLModel, table=True):
    id: int | None = Field(default=None, primary_key=True)
    email: str = Field(unique=True, index=True, max_length=255)
    hashed_password: str
    is_active: bool = Field(default=True)
    created_at: datetime = Field(default_factory=lambda: datetime.now(UTC))
```
Le paramètre `pool_size=20` dimensionne le pool de connexions PostgreSQL. Sur une instance Postgres 17.4 standard avec `max_connections=100`, ce réglage permet de servir 5 workers Uvicorn sans saturer la base. Pour la production à fort trafic, ajoutez PgBouncer en transaction pooling devant Postgres : vous gagnez un facteur 5 à 10 sur le nombre de connexions logiques sans dégradation. Le flag `pool_pre_ping=True` est crucial : il évite les `OperationalError: connection has been closed` lorsqu’un load balancer ferme silencieusement les connexions inactives.

## Étape 6 : Implémenter les routes CRUD avec dépendances asynchrones

FastAPI 0.95+ a introduit la syntaxe `Annotated[Type, Depends(...)]` qui remplace l’ancienne signature `param: Type = Depends(...)`. Cette nouvelle forme est compatible avec mypy strict et permet aux IDE de remonter correctement les types. Tous les exemples ci-dessous utilisent la syntaxe Annotated, qui est désormais la convention recommandée par tiangolo dans la documentation officielle 2026.

```
# app/api/routes/articles.py
from typing import Annotated
from fastapi import APIRouter, Depends, HTTPException, status
from sqlmodel import select
from sqlmodel.ext.asyncio.session import AsyncSession
from app.api.deps import CurrentUser, SessionDep
from app.models.article import Article
from app.schemas.article import ArticleCreate, ArticleRead, ArticleUpdate
router = APIRouter()
@router.get("/", response_model=list[ArticleRead])
async def list_articles(
    session: SessionDep,
    skip: int = 0,
    limit: int = 20,
) -> list[Article]:
    statement = select(Article).where(Article.published == True).offset(skip).limit(limit)
    result = await session.exec(statement)
    return result.all()
@router.post("/", response_model=ArticleRead, status_code=status.HTTP_201_CREATED)
async def create_article(
    payload: ArticleCreate,
    session: SessionDep,
    user: CurrentUser,
) -> Article:
    article = Article.model_validate(payload, update={"author_id": user.id})
    session.add(article)
    await session.commit()
    await session.refresh(article)
    return article
@router.get("/{article_id}", response_model=ArticleRead)
async def get_article(article_id: int, session: SessionDep) -> Article:
    article = await session.get(Article, article_id)
    if not article:
        raise HTTPException(status.HTTP_404_NOT_FOUND, "Article introuvable")
    return article
@router.patch("/{article_id}", response_model=ArticleRead)
async def update_article(
    article_id: int,
    payload: ArticleUpdate,
    session: SessionDep,
    user: CurrentUser,
) -> Article:
    article = await session.get(Article, article_id)
    if not article:
        raise HTTPException(status.HTTP_404_NOT_FOUND, "Article introuvable")
    if article.author_id != user.id:
        raise HTTPException(status.HTTP_403_FORBIDDEN, "Accès refusé")
    update_data = payload.model_dump(exclude_unset=True)
    for key, value in update_data.items():
        setattr(article, key, value)
    session.add(article)
    await session.commit()
    await session.refresh(article)
    return article
@router.delete("/{article_id}", status_code=status.HTTP_204_NO_CONTENT)
async def delete_article(
    article_id: int,
    session: SessionDep,
    user: CurrentUser,
) -> None:
    article = await session.get(Article, article_id)
    if not article or article.author_id != user.id:
        raise HTTPException(status.HTTP_404_NOT_FOUND)
    await session.delete(article)
    await session.commit()
```
Le pattern `SessionDep = Annotated[AsyncSession, Depends(get_session)]` dans `deps.py` permet de réutiliser la dépendance dans toutes les routes sans dupliquer la signature complète. Cette factorisation est devenue le standard depuis le template officiel *full-stack-fastapi-template* de tiangolo. La méthode `model_dump(exclude_unset=True)` de Pydantic v2 retourne uniquement les champs réellement envoyés par le client, ce qui permet des updates partielles propres sans écraser les valeurs existantes par `None`.

## Étape 7 : Sécuriser l’API avec JWT et OAuth2 password flow

FastAPI fournit nativement le squelette OAuth2 via `OAuth2PasswordBearer` et `OAuth2PasswordRequestForm`. Pour l’authentification réelle, nous utilisons des JWT signés HS256 avec python-jose et le hashing bcrypt via passlib. En production, augmentez le coût bcrypt à 12 (par défaut) ou 14 selon votre tolérance latence : un hash bcrypt cost 12 prend environ 250 ms sur un Xeon Gold 6326, ce qui reste acceptable pour un endpoint `/login` appelé rarement.

