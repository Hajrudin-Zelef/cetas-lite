---
id: collect-261001-rattrapage/rattrapage/tutoriel-docker-compose-2026-stack-prod-en-13-etapes-6
title: "Mise à jour des dépôts et installation des dépendances"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["arr", "attention"]
source: docs/RAG/collect-261001-rattrapage/tutoriel-docker-compose-2026-stack-prod-en-13-etapes.md
source_anchor: ""
source_lines: [632, 654]
sha256: 46a57de65df0263f9679e0ef32eb2d38281e77e032bf71e3870cd68d2b25ded2
---

# Mise à jour des dépôts et installation des dépendances

`docker compose exec NOM_SERVICE COMMANDE`. Exemple : `docker compose exec db psql -U postgres`. Pour un service à l’arrêt, utilisez `docker compose run --rm NOM_SERVICE COMMANDE` qui lance un conteneur jetable, exécute la commande et le supprime.

### Comment scaler horizontalement un service ?

`docker compose up -d --scale api=4` lance quatre instances du service `api`. Attention : le service ne doit pas avoir de `ports:` en mode HOST:CONTAINER (collision), ni de volume bind unique non partageable. Combiné à un load balancer interne (Nginx, Traefik), c’est suffisant pour gérer plusieurs centaines de req/s sur un hôte.

### Comment migrer un projet de Compose v1 à v2 ?

Trois étapes : 1) supprimer la ligne `version:` en tête de fichier ; 2) remplacer toutes les occurrences de `docker-compose` par `docker compose` dans vos scripts CI ; 3) tester avec `docker compose config` pour repérer les warnings de syntaxe désormais stricts (clés inconnues, types incorrects). 95 % des fichiers v1 sont compatibles tels quels.

### Les volumes nommés survivent-ils à `docker compose down` ?

Oui. `docker compose down` seul supprime conteneurs et réseaux mais **conserve** les volumes nommés. Pour tout effacer (utile en début de tests d’intégration), ajoutez `-v` : `docker compose down -v`. Cette opération est irréversible : sauvegardez avant !

### Comment obtenir un dépannage HTTPS rapide en local ?

Remplacez Nginx par **Caddy** dans le service `proxy` et activez `tls internal` dans le Caddyfile : Caddy génère automatiquement un certificat ACME local. Aucune ligne OpenSSL à manipuler. Pour les déploiements publics, Caddy obtient gratuitement un certificat Let’s Encrypt en moins de 10 secondes.

## Conclusion : Une Stack Docker Compose Production-Ready en 60 Minutes

Vous disposez désormais d’une stack **Docker Compose** complète : API FastAPI, PostgreSQL persistant, cache Redis, reverse proxy Nginx avec rate-limiting, monitoring Prometheus/Grafana activable à la demande, gestion sécurisée des secrets, override de développement avec hot-reload via Compose Watch, et procédure de déploiement production avec tags immuables. Toutes les bonnes pratiques 2026 sont intégrées : suppression du champ `version:`, healthchecks systématiques, profiles, limites de ressources, multi-stage Dockerfile.

Prochaines étapes naturelles : configurer un pipeline CI/CD avec GitHub Actions pour tester chaque commit dans une stack éphémère (`docker compose -f compose.test.yml up --abort-on-container-exit`), sauvegarder vos volumes nocturnement vers S3 ou Backblaze B2, et – quand votre charge approche d’un million de requêtes par jour – migrer vers Kubernetes via Kompose. Mais en attendant, Docker Compose vous offre 90 % de la valeur pour 10 % de la complexité.
