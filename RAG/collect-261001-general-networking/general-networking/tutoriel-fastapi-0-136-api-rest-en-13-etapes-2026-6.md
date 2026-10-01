---
id: collect-261001-general-networking/general-networking/tutoriel-fastapi-0-136-api-rest-en-13-etapes-2026-6
title: "Installation d'uv (Astral) – outil officiel recommandé en 2026"
domain: general-networking
role: reference
task: reference
actors: ["Hugging Face", "Microsoft", "vLLM"]
dates: []
keywords: ["benchmarks", "decode", "diffusion", "inference", "memory", "vllm"]
source: docs/RAG/collect-261001-general-networking/tutoriel-fastapi-0-136-api-rest-en-13-etapes-2026.md
source_anchor: ""
source_lines: [580, 648]
sha256: 5022aa764bd47263a701850dc2eb1579a66f9487e8be37f6768d8f909cb63eb5
---

# Installation d'uv (Astral) – outil officiel recommandé en 2026

```
# Déploiement Scaleway Serverless Containers via CLI
scw container namespace create name=fastapi-2026 region=fr-par
scw registry namespace create name=fastapi-2026
docker tag fastapi-app:latest rg.fr-par.scw.cloud/fastapi-2026/api:latest
docker push rg.fr-par.scw.cloud/fastapi-2026/api:latest
scw container container create \
    namespace-id=$NAMESPACE_ID \
    name=api \
    registry-image=rg.fr-par.scw.cloud/fastapi-2026/api:latest \
    port=8000 \
    cpu-limit=1000 \
    memory-limit=1024 \
    min-scale=1 \
    max-scale=10
scw container container deploy container-id=$CONTAINER_ID
```
Pour une charge prévisible et continue, OVH Public Cloud avec Managed Kubernetes Service à 0,01 € par node-heure (hors VM) reste l’option la plus économique en France. Comptez environ 25 €/mois pour un cluster minimal 3 nœuds B2-7 (2 vCPU, 7 Go) capable d’héberger une API FastAPI servant 500 requêtes/seconde. Pour le scaling vertical, profitez du autoscaling pod-level (HPA) sur la métrique `http_requests_per_second` exposée par Prometheus.

## Pièges courants à éviter en 2026

FastAPI est généralement intuitif, mais certaines erreurs récurrentes coûtent des heures de débogage. Voici les six pièges les plus fréquents que nous voyons en revue de code dans les ESN parisiennes et les startups deeptech.

- **Mélanger sync et async dans la même route** : appeler une fonction synchrone bloquante (requests, time.sleep, psycopg2 sync) dans une route`async def` bloque l’event loop entier. Utilisez`fastapi.concurrency.run_in_threadpool` ou bascule l’appel vers une variante asynchrone (httpx, asyncpg).
- **Oublier le `response_model`** : sans cette annotation, FastAPI sérialise l’objet brut, ce qui peut leaker des champs sensibles comme`hashed_password` . Définissez systématiquement un schéma de sortie séparé du modèle DB.
- **Confondre `Body()`, `Query()` et `Path()`** : Body extrait du JSON, Query du querystring, Path du chemin. Un mauvais decorator donne des 422 incompréhensibles côté frontend.
- **Pydantic v1 dans les imports** : depuis FastAPI 0.128,`from pydantic.v1 import BaseModel` ne fonctionne plus. Migrez avec`bump-pydantic` .
- **Sessions DB partagées entre requêtes** : créer un`AsyncSession` au niveau module et le réutiliser provoque des fuites de transactions. Toujours injecter via`Depends(get_session)` .
- **Lancer Uvicorn sans `--proxy-headers` derrière un reverse proxy** : sans ce flag, request.client.host retourne 127.0.0.1 au lieu de l’IP réelle, cassant le rate limiting et les logs.

## Tableau comparatif : FastAPI vs Flask vs Django REST en 2026

| Critère | FastAPI 0.136.1 | Flask 3.1 | Django REST 3.15 | 
|---|---|---|---|
| Année de création | 2018 | 2010 | 2014 | 
| Async natif | Oui (ASGI) | Partiel (Flask 2+) | Limité (ASGI) | 
| Validation | Pydantic v2 (Rust) | Manuelle / marshmallow | Serializers DRF | 
| Documentation auto | OpenAPI 3.1 + Swagger | flasgger (extension) | drf-spectacular | 
| Type hints stricts | Obligatoires | Optionnels | Optionnels | 
| Performance (req/s, JSON simple) | ≈ 30 000+ | ≈ 5 000-10 000 | ≈ 4 000-8 000 | 
| Courbe apprentissage | Faible | Très faible | Élevée | 
| Écosystème | Asynchrone (httpx, asyncpg) | Synchrone (requests, psycopg2) | Tout-en-un (ORM, admin, auth) | 
| Cas d’usage idéal | API REST moderne, ML serving | Microservices simples | App full-stack monolithique | 

Les chiffres de performance dépendent fortement du scénario testé (JSON simple, ORM, IO bound). Sur les benchmarks TechEmpower Round 22, FastAPI atteint environ 30 000 requêtes par seconde sur le test JSON simple, soit 4 à 6 fois plus que Flask synchrone. Sur des charges IO-bound (requêtes DB, appels HTTP externes), l’écart se creuse encore plus en faveur de FastAPI grâce à l’asynchronisme. Pour des applications full-stack monolithiques avec admin intégré et ORM Django, Django REST reste un choix légitime – FastAPI ne cherche pas à le remplacer sur ce périmètre.

## Dépannage : 8 erreurs FastAPI fréquentes et leurs solutions

- **“AttributeError: module ‘pydantic’ has no attribute ‘BaseSettings'”** – Pydantic v2 a déplacé BaseSettings vers le paquet`pydantic-settings` . Installez`uv add pydantic-settings` et importez via`from pydantic_settings import BaseSettings` .
- **“RuntimeError: Event loop is closed”** en tests – Activez`asyncio_mode = "auto"` dans`pyproject.toml` sous`[tool.pytest.ini_options]` et utilisez la fixture`event_loop` de pytest-asyncio 0.25+.
- **HTTP 422 sur un payload qui semble valide** – Vérifiez que vous envoyez bien`Content-Type: application/json` . Sans ce header, FastAPI cherche les paramètres dans le querystring et échoue.
- **“sqlalchemy.exc.MissingGreenlet”** – Vous appelez une méthode synchrone sur un objet ORM dans un contexte async. Préfixez par`await` ou utilisez`session.run_sync()` .
- **OpenAPI lent à générer (10+ secondes)** – Vous avez probablement plusieurs centaines de routes ou des schémas Pydantic récursifs. Activez`app = FastAPI(openapi_url=None)` en production et générez le JSON OpenAPI au build.
- **“jose.exceptions.ExpiredSignatureError”** imprévu – Vérifiez le décalage horaire entre vos serveurs (NTP) et augmentez`leeway=10` dans`jwt.decode()` pour tolérer 10 secondes de drift.
- **Connexions PostgreSQL qui s’accumulent** – Surveillez`pg_stat_activity` . Probable absence de`await session.close()` ou pool trop large. Ajoutez PgBouncer en transaction pooling.
- **“422 Unprocessable Entity” sans détails** – Activez le middleware de logging des erreurs Pydantic et inspectez`exc.errors()` dans un exception handler personnalisé. Le message inclut le chemin exact du champ invalide.

## Astuces avancées pour optimiser une API FastAPI

Une fois le projet en production, plusieurs optimisations marginales mais cumulatives font gagner facilement 30 à 50 % de débit. Premier réflexe : remplacer `json` standard par `orjson` ou `msgspec` pour la sérialisation. Le gain mesuré sur le projet de référence est de 2 à 3× sur des payloads de 50 ko. FastAPI accepte un encodeur custom via `FastAPI(default_response_class=ORJSONResponse)`.

Deuxième optimisation : précompiler les schémas Pydantic au démarrage avec `BaseModel.model_rebuild()`. Sur une API avec 80 schémas, cette étape supprime le *JIT compilation* au premier appel et réduit la latence p99 du premier appel par client par 200 ms environ. Troisième astuce : remplacer le pool asyncpg standard par `asyncpg.create_pool()` avec `statement_cache_size=2048` pour cacher les plans préparés et économiser un round-trip Postgres par requête.

Pour les charges très élevées, étudiez Granian, un serveur ASGI alternatif écrit en Rust qui prétend battre Uvicorn de 20 à 40 % sur des benchmarks réalistes selon les tests publiés en 2025. La migration depuis Uvicorn est triviale : `granian --interface asgi app.main:app`. Enfin, pour servir des modèles ML lourds (PyTorch, ONNX), externalisez l’inférence vers un sidecar Triton Inference Server ou vers vLLM, et faites de FastAPI le simple *API gateway*. Cette séparation libère l’event loop des opérations CPU-bound qui bloqueraient sinon le serveur entier.

## Cas d’usage et adoption en production

FastAPI s’est imposé en 2025-2026 comme le framework de référence pour exposer des modèles d’IA. Hugging Face utilise FastAPI dans la majorité de ses Spaces et endpoints d’inférence. Microsoft, Uber et Netflix figurent parmi les utilisateurs documentés selon Planeks (planeks.net), pour des cas d’usage allant de l’orchestration de microservices à la diffusion de contenu personnalisé. Le framework supporte 100 % de couverture de tests sur son propre code, ce qui rassure les équipes plateforme et SRE qui auditent les dépendances.

