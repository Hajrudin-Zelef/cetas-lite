---
id: collect-261001-rattrapage/rattrapage/tutoriel-docker-compose-2026-guide-complet-multi-conteneurs-3
title: "Vérifier les versions installées"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["attention"]
source: docs/RAG/collect-261001-rattrapage/tutoriel-docker-compose-2026-guide-complet-multi-conteneurs.md
source_anchor: ""
source_lines: [260, 406]
sha256: eabbef0c0abd68665b16151771e00c28654011cfc28fafc8fa2af2867b44023b
---

        # Timeouts
        proxy_connect_timeout 30s;
        proxy_read_timeout 60s;
    }
    # Health check endpoint
    location /health {
        proxy_pass http://api_backend/health;
    }
    # Page d'accueil
    location / {
        return 200 '{"message": "Docker Compose Tutorial API - tech-insider.org"}';
        add_header Content-Type application/json;
    }
}
```
L’élément clé ici est la directive `upstream` qui référence `api:8000`. Docker Compose crée automatiquement un réseau interne où chaque service est accessible par son nom. Ainsi, `api` résout vers l’adresse IP du conteneur FastAPI – aucune configuration réseau manuelle n’est nécessaire. C’est l’un des avantages majeurs de Docker Compose par rapport à l’utilisation manuelle de `docker run` avec des liens réseau explicites.

**Piège courant n°4 :** N’utilisez pas `localhost` ou `127.0.0.1` dans la configuration Nginx pour référencer d’autres conteneurs. Dans le réseau Docker, `localhost` fait référence au conteneur Nginx lui-même, pas aux autres services. Utilisez toujours le nom du service défini dans `compose.yaml`.

## Étape 5 : Écrire le Fichier compose.yaml Complet

Voici le cœur de notre projet : le fichier `compose.yaml` qui orchestre les quatre services. Ce fichier utilise les fonctionnalités modernes de Docker Compose disponibles en 2026, incluant les health checks, les dépendances conditionnelles, les profils et la gestion avancée des volumes. Chaque directive est documentée pour vous permettre de comprendre et d’adapter cette configuration à vos propres projets.

```
# compose.yaml - Docker Compose Tutorial 2026
# Documentation : https://docs.docker.com/compose/
services:
  # --- Base de données PostgreSQL ---
  db:
    image: postgres:16-alpine
    container_name: tutorial-db
    restart: unless-stopped
    environment:
      POSTGRES_DB: ${POSTGRES_DB:-appdb}
      POSTGRES_USER: ${POSTGRES_USER:-appuser}
      POSTGRES_PASSWORD: ${POSTGRES_PASSWORD:?Mot de passe requis}
    ports:
      - "5432:5432"
    volumes:
      - postgres_data:/var/lib/postgresql/data
      - ./db/init.sql:/docker-entrypoint-initdb.d/init.sql:ro
    healthcheck:
      test: ["CMD-SHELL", "pg_isready -U ${POSTGRES_USER:-appuser} -d ${POSTGRES_DB:-appdb}"]
      interval: 10s
      timeout: 5s
      retries: 5
      start_period: 30s
    networks:
      - backend
  # --- Cache Redis ---
  redis:
    image: redis:7-alpine
    container_name: tutorial-redis
    restart: unless-stopped
    ports:
      - "6379:6379"
    volumes:
      - redis_data:/data
    healthcheck:
      test: ["CMD", "redis-cli", "ping"]
      interval: 10s
      timeout: 3s
      retries: 3
    command: redis-server --maxmemory 128mb --maxmemory-policy allkeys-lru
    networks:
      - backend
  # --- API FastAPI ---
  api:
    build:
      context: ./api
      dockerfile: Dockerfile
    container_name: tutorial-api
    restart: unless-stopped
    environment:
      DB_HOST: db
      REDIS_HOST: redis
      POSTGRES_DB: ${POSTGRES_DB:-appdb}
      POSTGRES_USER: ${POSTGRES_USER:-appuser}
      POSTGRES_PASSWORD: ${POSTGRES_PASSWORD:?Mot de passe requis}
    depends_on:
      db:
        condition: service_healthy
      redis:
        condition: service_healthy
    ports:
      - "8000:8000"
    networks:
      - backend
    develop:
      watch:
        - action: sync+restart
          path: ./api/main.py
          target: /app/main.py
  # --- Reverse Proxy Nginx ---
  nginx:
    image: nginx:1.27-alpine
    container_name: tutorial-nginx
    restart: unless-stopped
    ports:
      - "80:80"
    volumes:
      - ./nginx/nginx.conf:/etc/nginx/conf.d/default.conf:ro
    depends_on:
      - api
    networks:
      - backend
volumes:
  postgres_data:
    driver: local
  redis_data:
    driver: local
networks:
  backend:
    driver: bridge
```
Plusieurs points importants méritent votre attention dans ce fichier. La syntaxe `${POSTGRES_PASSWORD:?Mot de passe requis}` force Docker Compose à échouer si la variable n’est pas définie – un garde-fou essentiel en production. La directive `depends_on` avec `condition: service_healthy` garantit que l’API ne démarre qu’une fois PostgreSQL et Redis prêts à accepter des connexions. Sans cette condition, l’API pourrait démarrer avant la base de données et crasher immédiatement. Enfin, la section `develop.watch` active la fonctionnalité Compose Watch qui synchronise automatiquement vos fichiers locaux vers le conteneur pendant le développement.

## Étape 6 : Configurer les Variables d’Environnement

La gestion des variables d’environnement est un aspect critique de tout projet Docker Compose. En 2026, les bonnes pratiques imposent de ne jamais coder en dur les mots de passe, clés API ou paramètres sensibles dans le fichier `compose.yaml`. Docker Compose prend en charge plusieurs mécanismes pour injecter ces valeurs de manière sécurisée.

Créez le fichier `.env` à la racine du projet. Ce fichier est automatiquement chargé par Docker Compose :

```
# .env - Variables d'environnement (NE PAS VERSIONNER)
POSTGRES_DB=appdb
POSTGRES_USER=appuser
POSTGRES_PASSWORD=MonMotDePasse2026!Securise
REDIS_HOST=redis
API_PORT=8000
```
Docker Compose applique un ordre de priorité précis pour résoudre les variables d’environnement. Comprendre cet ordre est essentiel pour déboguer les problèmes de configuration :

| Priorité | Source | Exemple | Utilisation | 
|---|---|---|---|
| 1 (la plus haute) | Variables shell de l’hôte | `export POSTGRES_PASSWORD=xxx` | Override temporaire en CI/CD | 
| 2 | Fichier .env spécifié | `docker compose --env-file prod.env up` | Environnements multiples | 
| 3 | Fichier .env par défaut | Fichier `.env` à la racine | Développement local | 
| 4 | Valeurs par défaut dans compose.yaml | `${VAR:-default}` | Fallback sécurisé | 
| 5 (la plus basse) | Dockerfile ENV | `ENV APP_PORT=8000` | Valeurs de base de l’image | 

**Piège courant n°5 :** Les espaces autour du signe `=` dans le fichier `.env` sont inclus dans la valeur. La ligne `POSTGRES_PASSWORD = secret` définira le mot de passe comme `" secret"` (avec un espace devant), ce qui causera des erreurs d’authentification difficiles à diagnostiquer. Utilisez toujours `POSTGRES_PASSWORD=secret` sans espaces.

## Étape 7 : Lancer et Vérifier les Services

Tout est en place. Il est temps de construire les images et de démarrer l’ensemble de la stack. Docker Compose offre plusieurs commandes essentielles que tout développeur doit maîtriser. Voici le processus complet de lancement et de vérification de notre application multi-conteneurs.

