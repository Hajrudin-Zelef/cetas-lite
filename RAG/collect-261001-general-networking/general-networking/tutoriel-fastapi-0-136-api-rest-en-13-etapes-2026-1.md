---
id: collect-261001-general-networking/general-networking/tutoriel-fastapi-0-136-api-rest-en-13-etapes-2026-1
title: "Installation d'uv (Astral) – outil officiel recommandé en 2026"
domain: general-networking
role: reference
task: reference
actors: ["Hugging Face", "Microsoft", "Mistral", "OpenAI"]
dates: []
keywords: ["benchmarks", "mistral"]
source: docs/RAG/collect-261001-general-networking/tutoriel-fastapi-0-136-api-rest-en-13-etapes-2026.md
source_anchor: ""
source_lines: [1, 52]
sha256: 6dc54bce74f68817850e6af115b0923ef6dcf3b7d84dceb3073d0e87e5a2fce5
---

# Installation d'uv (Astral) – outil officiel recommandé en 2026

**FastAPI 0.141.1** est la dernière version publiée sur le dépôt GitHub officiel du projet en juillet 2026, confirmant la bascule définitive de l’écosystème Python vers le tout-asynchrone et la validation typée stricte. Reflet de ce format d’apprentissage plus structuré qui s’impose en 2026, ce **tutoriel FastAPI** Tech Insider en 13 étapes vous guide de l’installation à la mise en production d’une API REST complète, avec authentification JWT, base PostgreSQL, streaming JSON Lines (la grande nouveauté de la version 0.134.0) et déploiement Docker. Comptez 90 minutes pour boucler le projet final, qui sert aujourd’hui de squelette de référence pour les équipes backend à Paris, Berlin et Amsterdam.

Le framework, désormais utilisé en production par **Microsoft, Uber et Netflix** selon Planeks, s’appuie sur Starlette ≥ 0.46.0 et Pydantic v2 exclusivement depuis la version 0.128.0 de décembre 2025. Le rythme de publication s’est encore accéléré depuis : la 0.138.0 est sortie le 20 juin 2026 d’après le tableau des versions d’EOL Risk, suivie de la 0.138.1 en juin puis de la 0.140.1 en juillet 2026 selon les notes de version officielles. Ce tutoriel couvre toutes les nouveautés 2025-2026 : **support Python 3.13 free-threaded**, streaming binaire avec `yield`, dépendances asynchrones, et tests automatisés à 100 % de couverture. À la fin de ce guide, vous disposerez d’une API REST production-ready, conteneurisée et déployable en un seul `docker compose up`.

## Pourquoi FastAPI s’impose en 2026 face à Flask et Django

FastAPI, créé en 2018 par Sebastián Ramírez (alias *tiangolo*), a connu une accélération spectaculaire entre 2024 et 2026. Le framework asynchrone-first dépasse aujourd’hui Flask en téléchargements PyPI mensuels et talonne Django sur les nouveaux projets backend en France. Selon le Stack Overflow Developer Survey 2024, FastAPI figure parmi les frameworks web Python les plus appréciés, avec un taux de satisfaction supérieur à 75 %, devant Flask et derrière uniquement Django sur la métrique d’adoption brute — un engouement que confirme un tutoriel GeeksforGeeks mis à jour en juin 2026, qui met en avant la validation automatique, le support asynchrone natif et la documentation interactive comme principaux moteurs d’adoption.

Trois facteurs expliquent cette domination en 2026. D’abord, la **validation Pydantic v2**, écrite en Rust, offre des gains de performance de 5 à 50× par rapport à v1 sur le parsing JSON, ce qui rapproche FastAPI des frameworks Go ou Node.js sur les benchmarks TechEmpower. Ensuite, la **génération automatique d’OpenAPI 3.1** et de la documentation Swagger UI/ReDoc supprime des centaines de lignes de boilerplate. Enfin, l’écosystème asynchrone Python (asyncpg, httpx, aioredis) a mûri au point que la plupart des bibliothèques métier disposent d’une variante `async` stable.

L’adoption en France suit la tendance mondiale. Les offres d’emploi backend Python sur Welcome to the Jungle et Indeed mentionnent FastAPI dans plus de 60 % des annonces senior pour 2026, contre 35 % en 2024. Les startups deeptech parisiennes (Mistral AI, Hugging Face, Dust) utilisent FastAPI pour exposer leurs modèles d’inférence, profitant du support natif des streams binaires introduit dans la version 0.134.0. Les ESN traditionnelles migrent progressivement leurs APIs Flask 2.x vers FastAPI pour bénéficier du typage strict et de la documentation auto-générée.

## Prérequis : versions exactes pour suivre ce tutoriel FastAPI

Avant d’attaquer le projet, alignez votre environnement sur les versions ci-dessous. FastAPI a abandonné Pydantic v1 depuis la version 0.128.0 de décembre 2025, et la ligne actuellement supportée est la 0.139.2, active depuis juillet 2026 et toujours listée comme la version en cours de maintenance par EOL Risk en août 2026 — ce qui implique des dépendances plus récentes que celles que vous trouverez dans les anciens tutoriels en ligne. Si vous travaillez sur Windows, privilégiez WSL2 Ubuntu 24.04 pour éviter les soucis de compilation des wheels asyncpg et orjson.

| Composant | Version recommandée | Version minimale | Notes 2026 | 
|---|---|---|---|
| Python | 3.12.7 | 3.10 | 3.13 free-threaded supporté en bêta | 
| FastAPI | 0.136.1 | 0.128.0 | Pydantic v1 retiré depuis 0.128.0 | 
| Pydantic | 2.11.x | 2.5 | v1 non supporté | 
| Uvicorn | 0.34.x | 0.30 | HTTP/2 et WebSockets natifs | 
| Starlette | 0.46.x | 0.46.0 | Requise par 0.134.0+ | 
| SQLModel | 0.0.24 | 0.0.22 | Compatible Pydantic v2 | 
| PostgreSQL | 17.4 | 15 | JSONB, GIN indexes | 
| Docker | 27.4 | 24.0 | BuildKit activé par défaut | 

Côté outillage, installez **uv** 0.5+ d’Astral comme gestionnaire de paquets : il remplace pip et venv avec des installations 10 à 100× plus rapides. Pour l’éditeur, VS Code avec l’extension Pylance et Ruff suffit. Si vous utilisez PyCharm Professional 2025.1, le support FastAPI est natif depuis cette version (autocomplétion sur les `Annotated[Depends()]`, navigation OpenAPI). Côté CLI, prévoyez `httpie` ou `curl` pour tester les endpoints, et `jq` pour formater les réponses JSON.

## Étape 1 : Installer Python 3.12 et créer un environnement virtuel

Sur Ubuntu 24.04 LTS (Noble Numbat), Python 3.12 est désormais disponible dans les dépôts officiels. Sur macOS, utilisez Homebrew. Sur Windows, l’installeur officiel python.org reste la voie la plus simple, ou WSL2 si vous visez la production Linux. Voici les commandes pour chaque OS, en privilégiant uv comme gestionnaire d’environnement.

```
# Installation d'uv (Astral) – outil officiel recommandé en 2026
curl -LsSf https://astral.sh/uv/install.sh | sh
# Création du projet
mkdir tutoriel-fastapi-2026 && cd tutoriel-fastapi-2026
uv init --python 3.12
uv venv
source .venv/bin/activate     # macOS/Linux
# .venv\Scripts\activate      # Windows PowerShell
# Vérification
python --version
# Python 3.12.7
```
L’utilisation d’uv divise par 10 le temps d’installation des dépendances par rapport à pip classique, selon les benchmarks Astral de janvier 2026. Sur un projet typique avec 40 dépendances, on passe d’environ 22 secondes (pip) à 2,1 secondes (uv). Cette différence devient critique en CI/CD lorsque vous exécutez plusieurs centaines de builds par jour. Si vous préférez rester sur pip, remplacez `uv add` par `pip install` dans les étapes suivantes ; le code applicatif reste identique.

## Étape 2 : Installer FastAPI 0.136.1 et ses dépendances clés

FastAPI 0.139.0, publiée le 1er juillet 2026, a ajouté la gestion des dépendances via `app.frontend()` et requiert toujours Pydantic ≥ 2.5 et Starlette ≥ 0.46.0. Comme le recommande le tutoriel 2026 de Sebastián Ramírez (tiangolo) pour tout nouveau projet, l’installation se fait via `uv add "fastapi[standard]"`. Pour ce tutoriel, nous installons également Uvicorn (serveur ASGI), SQLModel (ORM), asyncpg (driver PostgreSQL asynchrone), python-jose (JWT) et passlib (hashing bcrypt). Évitez d’épingler des versions trop strictes en développement, mais figez-les en production via `uv lock`.

