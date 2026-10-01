---
id: collect-261001-rattrapage/rattrapage/tutoriel-docker-compose-2026-guide-complet-multi-conteneurs-5
title: "Vérifier les versions installées"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["arr", "memory"]
source: docs/RAG/collect-261001-rattrapage/tutoriel-docker-compose-2026-guide-complet-multi-conteneurs.md
source_anchor: ""
source_lines: [484, 638]
sha256: 673180928f0126e5f0a7dd74866280737c3f4aaa71e12bbba48b48a36d1d6fbf
---

# Vérifier les versions installées

```
# Lister les volumes créés par Docker Compose
docker volume ls --filter name=compose-tutorial
# DRIVER    VOLUME NAME
# local     compose-tutorial_postgres_data
# local     compose-tutorial_redis_data
# Inspecter un volume pour voir où il est stocké
docker volume inspect compose-tutorial_postgres_data
# [{"CreatedAt":"2026-03-23T10:00:00Z",
#   "Driver":"local",
#   "Mountpoint":"/var/lib/docker/volumes/compose-tutorial_postgres_data/_data",
#   "Name":"compose-tutorial_postgres_data"}]
# Sauvegarder les données PostgreSQL
docker compose exec db pg_dump -U appuser appdb > backup_$(date +%Y%m%d).sql
# Restaurer une sauvegarde
cat backup_20260323.sql | docker compose exec -T db psql -U appuser appdb
# Supprimer les volumes orphelins (nettoyage)
docker volume prune --filter "label!=keep"
```
L’option `:ro` (read-only) que nous utilisons pour le script d’initialisation et la configuration Nginx est une bonne pratique de sécurité. Elle empêche le conteneur de modifier ces fichiers, réduisant la surface d’attaque en cas de compromission du conteneur. En production, appliquez `:ro` à tous les bind mounts de configuration.

**Piège courant n°6 :** La commande `docker compose down` ne supprime PAS les volumes par défaut. Vos données PostgreSQL survivent aux arrêts/redémarrages. Si vous voulez un reset complet, utilisez explicitement `docker compose down -v`. C’est un comportement qui surprend beaucoup de développeurs qui pensent repartir de zéro après un `down`.

## Étape 11 : Optimiser pour la Production avec les Profils et Overrides

Docker Compose offre des mécanismes puissants pour gérer différents environnements (développement, staging, production) à partir d’une même base de configuration. Les **profils** et les **fichiers override** sont les deux approches principales en 2026, chacune adaptée à des cas d’usage spécifiques.

Les profils permettent de regrouper des services par contexte. Par exemple, vous pouvez créer un profil “debug” qui inclut des outils de monitoring, ou un profil “test” qui ajoute un conteneur de tests automatisés. Voici comment étendre notre `compose.yaml` avec des profils :

```
# Ajout au compose.yaml - Services avec profils
services:
  # ... services existants ...
  # Outil de monitoring (profil debug uniquement)
  adminer:
    image: adminer:4
    container_name: tutorial-adminer
    ports:
      - "8080:8080"
    networks:
      - backend
    profiles:
      - debug
  # Service de tests (profil test uniquement)
  tests:
    build:
      context: ./api
    container_name: tutorial-tests
    command: python -m pytest tests/ -v
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
    networks:
      - backend
    profiles:
      - test
# Utilisation :
# docker compose up -d                        → Services de base uniquement
# docker compose --profile debug up -d        → + Adminer
# docker compose --profile test run tests     → Exécuter les tests
```
L’approche par fichiers override est complémentaire. Créez un fichier `compose.override.yaml` pour le développement (chargé automatiquement) et un `compose.prod.yaml` pour la production :

```
# compose.override.yaml - Surcharges pour le développement (chargé automatiquement)
services:
  api:
    build:
      target: base
    volumes:
      - ./api:/app  # Hot-reload en développement
    environment:
      - DEBUG=true
# compose.prod.yaml - Configuration de production
services:
  api:
    restart: always
    deploy:
      resources:
        limits:
          cpus: "1.0"
          memory: 512M
    environment:
      - DEBUG=false
  
  db:
    restart: always
    deploy:
      resources:
        limits:
          cpus: "2.0"
          memory: 1G
  nginx:
    restart: always
    ports:
      - "443:443"
# Lancer en production :
# docker compose -f compose.yaml -f compose.prod.yaml up -d
```
En production, les limites de ressources (`deploy.resources.limits`) empêchent un conteneur de consommer toute la mémoire ou le CPU de l’hôte. C’est une configuration essentielle qui prévient les situations où un pic de trafic sur l’API affame la base de données en ressources, provoquant un effondrement en cascade de toute l’application.

## Étape 12 : Mettre en Place la Supervision et les Logs

La supervision est l’ultime étape pour transformer votre projet Docker Compose en une application prête pour la production. En 2026, les bonnes pratiques imposent une centralisation des logs et une surveillance active de la santé des services. Docker Compose facilite grandement cette mise en place grâce à ses pilotes de logging intégrés et sa compatibilité avec les outils de monitoring modernes. D’ailleurs, Docker Desktop 4.70.0, publié le 20 avril 2026 avec Docker Compose v5.1.2, a introduit de nouveaux indices en ligne de commande (CLI hints) pour la commande `compose logs`, facilitant le diagnostic rapide sans quitter le terminal.

Par défaut, Docker utilise le pilote de logging `json-file`, qui stocke les logs localement. Pour éviter que les logs ne remplissent le disque, configurez une rotation automatique dans votre `compose.yaml` :

```
# Ajout à chaque service dans compose.yaml
services:
  api:
    # ... configuration existante ...
    logging:
      driver: json-file
      options:
        max-size: "10m"    # Taille maximale par fichier de log
        max-file: "3"      # Nombre de fichiers de rotation
        tag: "{{.Name}}"   # Tag avec le nom du conteneur
# Commandes de monitoring essentielles :
# Suivre les logs de tous les services en temps réel
docker compose logs -f --tail=50
# Logs d'un seul service avec timestamps
docker compose logs -f --timestamps api
# Statistiques de ressources en temps réel
docker stats --format "table {{.Name}}\t{{.CPUPerc}}\t{{.MemUsage}}\t{{.NetIO}}"
# NAME              CPU %     MEM USAGE / LIMIT     NET I/O
# tutorial-api      0.15%     85.2MiB / 512MiB      1.2kB / 850B
# tutorial-db       0.08%     45.7MiB / 1GiB        2.1kB / 1.5kB
# tutorial-nginx    0.02%     12.3MiB / 256MiB      980B / 720B
# tutorial-redis    0.05%     8.1MiB / 128MiB       450B / 320B
# Vérifier les événements Docker
docker compose events --json
```
Pour une supervision plus avancée, vous pouvez intégrer Prometheus et Grafana à votre stack Docker Compose. Ces outils permettent de collecter des métriques, de créer des tableaux de bord et de configurer des alertes. Cependant, pour la plupart des projets de taille moyenne, les commandes `docker stats` et `docker compose logs` combinées avec une rotation des logs constituent une solution suffisante et bien plus simple à maintenir.

## Dépannage : 8 Problèmes Courants et Leurs Solutions

Même avec une configuration soignée, vous rencontrerez inévitablement des problèmes lors de l’utilisation de Docker Compose. Voici les huit situations les plus fréquentes, avec leurs causes et solutions détaillées, basées sur les retours de milliers de développeurs en 2026.

### Problème 1 : “port is already allocated”

**Cause :** Un autre processus (ou un ancien conteneur) utilise déjà le port demandé sur l’hôte. C’est le problème le plus fréquent lors du premier lancement.

**Solution :** Identifiez le processus avec `sudo lsof -i :5432` (Linux/macOS) ou `netstat -ano | findstr :5432` (Windows). Arrêtez le processus conflictuel ou modifiez le port dans `compose.yaml` : `"5433:5432"` pour mapper sur un port différent de l’hôte.

### Problème 2 : “no matching manifest for linux/arm64”

