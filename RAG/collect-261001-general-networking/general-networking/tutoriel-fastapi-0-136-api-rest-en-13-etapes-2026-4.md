---
id: collect-261001-general-networking/general-networking/tutoriel-fastapi-0-136-api-rest-en-13-etapes-2026-4
title: "Installation d'uv (Astral) – outil officiel recommandé en 2026"
domain: general-networking
role: reference
task: reference
actors: ["AWS", "OpenAI"]
dates: []
keywords: ["aws", "decode"]
source: docs/RAG/collect-261001-general-networking/tutoriel-fastapi-0-136-api-rest-en-13-etapes-2026.md
source_anchor: ""
source_lines: [283, 413]
sha256: fa51103bf9985c9b59d0548f1102d70570954aae34604bbcef7af7581d604d2a
---

# Installation d'uv (Astral) – outil officiel recommandé en 2026

```
# app/core/security.py
from datetime import datetime, timedelta, UTC
from jose import jwt, JWTError
from passlib.context import CryptContext
from app.core.config import settings
pwd_ctx = CryptContext(schemes=["bcrypt"], deprecated="auto")
def hash_password(password: str) -> str:
    return pwd_ctx.hash(password)
def verify_password(plain: str, hashed: str) -> bool:
    return pwd_ctx.verify(plain, hashed)
def create_access_token(subject: str | int, expires_minutes: int = 60) -> str:
    expire = datetime.now(UTC) + timedelta(minutes=expires_minutes)
    payload = {"sub": str(subject), "exp": expire, "type": "access"}
    return jwt.encode(payload, settings.SECRET_KEY, algorithm="HS256")
def decode_token(token: str) -> dict:
    try:
        return jwt.decode(token, settings.SECRET_KEY, algorithms=["HS256"])
    except JWTError as e:
        raise ValueError("Token invalide") from e
```
```
# app/api/routes/auth.py
from typing import Annotated
from fastapi import APIRouter, Depends, HTTPException, status
from fastapi.security import OAuth2PasswordRequestForm
from sqlmodel import select
from app.api.deps import SessionDep
from app.core.security import create_access_token, hash_password, verify_password
from app.models.article import User
from pydantic import BaseModel, EmailStr
router = APIRouter()
class TokenResponse(BaseModel):
    access_token: str
    token_type: str = "bearer"
class RegisterRequest(BaseModel):
    email: EmailStr
    password: str
@router.post("/register", response_model=TokenResponse, status_code=201)
async def register(payload: RegisterRequest, session: SessionDep) -> TokenResponse:
    existing = await session.exec(select(User).where(User.email == payload.email))
    if existing.first():
        raise HTTPException(409, "Email déjà enregistré")
    user = User(email=payload.email, hashed_password=hash_password(payload.password))
    session.add(user)
    await session.commit()
    await session.refresh(user)
    return TokenResponse(access_token=create_access_token(user.id))
@router.post("/login", response_model=TokenResponse)
async def login(
    form: Annotated[OAuth2PasswordRequestForm, Depends()],
    session: SessionDep,
) -> TokenResponse:
    result = await session.exec(select(User).where(User.email == form.username))
    user = result.first()
    if not user or not verify_password(form.password, user.hashed_password):
        raise HTTPException(status.HTTP_401_UNAUTHORIZED, "Identifiants invalides")
    return TokenResponse(access_token=create_access_token(user.id))
```
La `SECRET_KEY` doit faire au moins 32 caractères aléatoires en production. Générez-la avec `python -c "import secrets; print(secrets.token_urlsafe(64))"` et stockez-la dans un secret manager (AWS Secrets Manager, HashiCorp Vault, OVH Confidential Computing). Pour un déploiement multi-instances, signez plutôt avec RS256 (clé asymétrique) afin de ne pas dupliquer la clé privée sur chaque worker. La durée d’expiration de 60 minutes est un compromis raisonnable : abaissez-la à 15 minutes et complétez avec un refresh token long-lived stocké dans un cookie HttpOnly Secure pour les applications grand public.

## Étape 8 : Streaming JSON Lines avec yield (nouveauté FastAPI 0.134.0)

La version 0.134.0 sortie le 27 février 2026 a introduit le support natif des streams JSON Lines et binaires via `yield` dans une route. Cette fonctionnalité simplifie radicalement l’export de gros volumes (rapports CSV, exports d’événements, streaming de tokens LLM) sans charger toute la donnée en mémoire. Un endpoint qui retournait avant 200 Mo en une seule réponse JSON peut désormais émettre 1 ligne JSON par enregistrement, consommée incrémentalement par le client.

```
# app/api/routes/articles.py – ajout du streaming
import json
from collections.abc import AsyncIterator
from fastapi.responses import StreamingResponse
@router.get("/export.jsonl")
async def export_articles(session: SessionDep) -> StreamingResponse:
    async def generate() -> AsyncIterator[bytes]:
        statement = select(Article).execution_options(yield_per=100)
        result = await session.stream(statement)
        async for article in result.scalars():
            line = json.dumps({
                "id": article.id,
                "title": article.title,
                "created_at": article.created_at.isoformat(),
            }, ensure_ascii=False)
            yield (line + "\n").encode("utf-8")
    return StreamingResponse(
        generate(),
        media_type="application/x-ndjson",
        headers={"Content-Disposition": "attachment; filename=articles.jsonl"},
    )
```
L’option `execution_options(yield_per=100)` de SQLAlchemy 2.0 force le driver asyncpg à streamer les lignes par paquets de 100 plutôt que de tout charger en RAM. Combinée avec `StreamingResponse` et un générateur asynchrone, cette approche permet d’exporter des tables de plusieurs millions d’enregistrements avec une empreinte mémoire constante d’environ 50 Mo. Les tests internes d’OpenAI sur leurs APIs de streaming utilisent un pattern strictement identique pour émettre les tokens LLM, à ceci près que le format est *Server-Sent Events* (`text/event-stream`) au lieu de *NDJSON*.

## Étape 9 : Middlewares de sécurité, CORS et rate limiting

Une API exposée en production doit traiter au minimum trois préoccupations transverses : CORS (Cross-Origin Resource Sharing), HTTPS strict via HSTS, et rate limiting. FastAPI, héritant de Starlette, expose un système de middlewares ASGI compatible avec l’écosystème entier (slowapi, fastapi-limiter, asgi-correlation-id). Voici une configuration de production solide.

```
# app/main.py – extension
from fastapi.middleware.cors import CORSMiddleware
from fastapi.responses import JSONResponse
from starlette.middleware.trustedhost import TrustedHostMiddleware
from slowapi import Limiter
from slowapi.util import get_remote_address
from slowapi.errors import RateLimitExceeded
from slowapi.middleware import SlowAPIMiddleware
limiter = Limiter(key_func=get_remote_address, default_limits=["100/minute"])
app.state.limiter = limiter
app.add_middleware(SlowAPIMiddleware)
app.add_middleware(
    CORSMiddleware,
    allow_origins=["https://example.fr", "https://app.example.fr"],
    allow_credentials=True,
    allow_methods=["GET", "POST", "PATCH", "DELETE"],
    allow_headers=["Authorization", "Content-Type"],
    max_age=600,
)
app.add_middleware(
    TrustedHostMiddleware,
    allowed_hosts=["api.example.fr", "*.example.fr"],
)
@app.exception_handler(RateLimitExceeded)
async def rate_limit_handler(request, exc):
    return JSONResponse(
        status_code=429,
        content={"detail": "Trop de requêtes, réessayez dans quelques secondes"},
    )
```
Pour le rate limiting distribué entre plusieurs workers ou plusieurs nodes Kubernetes, remplacez le backend mémoire de slowapi par Redis : `Limiter(storage_uri="redis://redis:6379")`. Le coût opérationnel est marginal (un GET INCR par requête) mais l’efficacité est globale. Évitez d’appliquer un rate limit unique par IP : derrière un proxy ou un VPN d’entreprise, des dizaines d’utilisateurs partagent la même IP source. Préférez un rate limit composite `get_remote_address + user_id` pour les endpoints authentifiés.

## Étape 10 : Tester l’API avec pytest et httpx AsyncClient

FastAPI fournit `TestClient` (synchrone, basé sur httpx) pour les tests unitaires rapides. Pour des tests d’intégration end-to-end qui imitent fidèlement le comportement asynchrone réel, utilisez `httpx.AsyncClient` avec `pytest-asyncio`. Cette approche détecte les blocages d’event loop, les fuites de connexions DB et les race conditions impossibles à reproduire avec TestClient.

