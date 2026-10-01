---
id: collect-261001-cisco/cisco/tutoriel-n8n-2026-automatiser-vos-workflows-ia-1
title: "Vérifier le statut du conteneur"
domain: cisco
role: reference
task: reference
actors: ["Anthropic", "OpenAI"]
dates: []
keywords: ["agents", "distribution", "open source"]
source: docs/RAG/collect-261001-cisco/tutoriel-n8n-2026-automatiser-vos-workflows-ia.md
source_anchor: ""
source_lines: [1, 153]
sha256: f24fe2459f0f199fed43afa2f5fbdab87fa9282bceb9e963e5908535702c514b
---

# Vérifier le statut du conteneur

L’automatisation des workflows est devenue un enjeu majeur pour les équipes techniques en 2026. Avec plus de 175 000 étoiles sur GitHub et une communauté de plusieurs centaines de milliers d’utilisateurs, **n8n** s’est imposé comme la plateforme d’automatisation open source de référence. Ce tutoriel complet vous guide pas à pas, de l’installation Docker à la création de workflows IA avancés avec agents autonomes, en passant par l’intégration de plus de 200 nœuds natifs.

Que vous soyez développeur, DevOps ou data engineer, ce guide pratique couvre tout ce dont vous avez besoin pour maîtriser n8n en 2026 : déploiement self-hosted, configuration des credentials, création de workflows complexes avec branchements conditionnels, intégration d’API REST, automatisation avec des modèles d’IA, et mise en production sécurisée. Chaque étape inclut du code fonctionnel, des captures de sortie et des solutions aux erreurs courantes.

## Prérequis et Environnement de Développement

Avant de commencer ce tutoriel n8n, assurez-vous de disposer des outils suivants installés et configurés sur votre machine. Ce guide a été testé sur Ubuntu 22.04/24.04, macOS Sonoma/Sequoia et Windows 11 avec WSL2. Les versions indiquées sont les versions minimales recommandées pour mars 2026.

| Outil | Version minimale | Rôle | Commande de vérification | 
|---|---|---|---|
| Docker | 25.0+ | Conteneurisation de n8n | `docker --version` | 
| Docker Compose | 2.24+ | Orchestration multi-conteneurs | `docker compose version` | 
| Node.js | 20.11+ LTS | Optionnel : installation npm | `node --version` | 
| Git | 2.40+ | Versionnement des workflows | `git --version` | 
| curl | 8.0+ | Tests d’API et webhooks | `curl --version` | 
| Un éditeur de code | VS Code 1.96+ | Édition des fichiers de config | `code --version` | 

Vous aurez également besoin d’un compte sur au moins un service tiers pour tester les intégrations : Gmail, Slack, GitHub ou Notion. Pour les workflows IA, une clé API OpenAI ou Anthropic est recommandée. Côté hardware, prévoyez au minimum 2 Go de RAM et 10 Go d’espace disque pour une installation de développement confortable.

Si Docker n’est pas encore installé, suivez notre tutoriel Docker Compose qui détaille l’installation complète sur chaque système d’exploitation. Pour les utilisateurs sous Windows, WSL2 avec la distribution Ubuntu est fortement recommandé pour une expérience optimale avec n8n.

## Étape 1 : Installer n8n avec Docker Compose

La méthode recommandée pour déployer n8n en 2026 est Docker Compose. Cette approche garantit un environnement reproductible, facilite les mises à jour et permet d’ajouter des services complémentaires comme PostgreSQL ou Redis. Avec la version 2.0 de n8n, le conteneur Docker officiel inclut toutes les dépendances nécessaires pour les workflows IA.

Commencez par créer un répertoire dédié à votre projet n8n et le fichier de configuration Docker Compose :

```
mkdir -p ~/n8n-project && cd ~/n8n-project
cat > docker-compose.yml << 'EOF'
version: '3.8'
services:
  n8n:
    image: n8nio/n8n:latest
    container_name: n8n
    restart: unless-stopped
    ports:
      - "5678:5678"
    environment:
      - N8N_BASIC_AUTH_ACTIVE=true
      - N8N_BASIC_AUTH_USER=admin
      - N8N_BASIC_AUTH_PASSWORD=VotreMotDePasseSecurise2026!
      - N8N_HOST=localhost
      - N8N_PORT=5678
      - N8N_PROTOCOL=http
      - WEBHOOK_URL=http://localhost:5678/
      - GENERIC_TIMEZONE=Europe/Paris
      - TZ=Europe/Paris
      - N8N_LOG_LEVEL=info
      - N8N_DIAGNOSTICS_ENABLED=false
    volumes:
      - n8n_data:/home/node/.n8n
    networks:
      - n8n-network
volumes:
  n8n_data:
networks:
  n8n-network:
    driver: bridge
EOF
docker compose up -d
```
Après quelques secondes de téléchargement de l'image, n8n sera accessible à l'adresse `http://localhost:5678`. La première connexion vous demandera de créer un compte propriétaire. Ce compte est distinct de l'authentification basique configurée dans le fichier Docker Compose — cette dernière protège l'accès au niveau du serveur web.

Pour vérifier que l'installation fonctionne correctement, exécutez la commande suivante :

```
# Vérifier le statut du conteneur
docker compose ps
# Sortie attendue :
# NAME   IMAGE              COMMAND                  SERVICE   CREATED          STATUS          PORTS
# n8n    n8nio/n8n:latest   "tini -- /docker-ent…"   n8n       30 seconds ago   Up 28 seconds   0.0.0.0:5678->5678/tcp
# Vérifier les logs
docker compose logs n8n --tail 20
# Tester l'API de santé
curl -s http://localhost:5678/healthz
# Sortie attendue : {"status":"ok"}
```
Si vous obtenez le statut `ok`, votre instance n8n est opérationnelle. En cas d'erreur de port occupé, modifiez le mapping de ports dans le fichier `docker-compose.yml` (par exemple `5679:5678`). Pour les environnements de production, nous verrons plus loin comment configurer PostgreSQL comme base de données au lieu de SQLite par défaut.

## Étape 2 : Configurer la Base de Données PostgreSQL pour la Production

Par défaut, n8n utilise SQLite pour stocker les workflows et les données d'exécution. Cette configuration convient au développement, mais pour un déploiement en production avec plusieurs utilisateurs et des volumes d'exécution élevés, PostgreSQL est indispensable. La migration vers PostgreSQL améliore les performances de 3 à 5 fois selon la documentation officielle de n8n 2.0.

Mettez à jour votre fichier `docker-compose.yml` pour inclure PostgreSQL :

```
version: '3.8'
services:
  postgres:
    image: postgres:16-alpine
    container_name: n8n-postgres
    restart: unless-stopped
    environment:
      POSTGRES_USER: n8n
      POSTGRES_PASSWORD: n8nDbPassword2026Secure
      POSTGRES_DB: n8n
      POSTGRES_NON_ROOT_USER: n8n
    volumes:
      - postgres_data:/var/lib/postgresql/data
    networks:
      - n8n-network
    healthcheck:
      test: ["CMD-SHELL", "pg_isready -U n8n -d n8n"]
      interval: 10s
      timeout: 5s
      retries: 5
  n8n:
    image: n8nio/n8n:latest
    container_name: n8n
    restart: unless-stopped
    ports:
      - "5678:5678"
    environment:
      - DB_TYPE=postgresdb
      - DB_POSTGRESDB_HOST=postgres
      - DB_POSTGRESDB_PORT=5432
      - DB_POSTGRESDB_DATABASE=n8n
      - DB_POSTGRESDB_USER=n8n
      - DB_POSTGRESDB_PASSWORD=n8nDbPassword2026Secure
      - N8N_BASIC_AUTH_ACTIVE=true
      - N8N_BASIC_AUTH_USER=admin
      - N8N_BASIC_AUTH_PASSWORD=VotreMotDePasseSecurise2026!
      - N8N_HOST=localhost
      - N8N_PORT=5678
      - N8N_PROTOCOL=http
      - WEBHOOK_URL=http://localhost:5678/
      - GENERIC_TIMEZONE=Europe/Paris
      - TZ=Europe/Paris
      - N8N_LOG_LEVEL=info
      - N8N_DIAGNOSTICS_ENABLED=false
      - EXECUTIONS_DATA_PRUNE=true
      - EXECUTIONS_DATA_MAX_AGE=168
    volumes:
      - n8n_data:/home/node/.n8n
    networks:
      - n8n-network
    depends_on:
      postgres:
        condition: service_healthy
volumes:
  n8n_data:
  postgres_data:
networks:
  n8n-network:
    driver: bridge
```
Notez les ajouts importants : le `healthcheck` PostgreSQL garantit que n8n ne démarre qu'après la disponibilité de la base de données. La variable `EXECUTIONS_DATA_PRUNE=true` avec `EXECUTIONS_DATA_MAX_AGE=168` (7 jours en heures) empêche l'accumulation de données d'exécution qui pourrait saturer votre espace disque. Relancez l'ensemble avec `docker compose down && docker compose up -d`.

Pour les entreprises françaises soumises aux exigences du RGPD, le self-hosting avec PostgreSQL sur des serveurs européens (OVH, Scaleway, Hetzner) garantit que toutes les données de workflow restent dans l'UE. C'est un avantage décisif de n8n face aux solutions SaaS américaines comme Zapier ou Make.

## Étape 3 : Créer Votre Premier Workflow d'Automatisation

