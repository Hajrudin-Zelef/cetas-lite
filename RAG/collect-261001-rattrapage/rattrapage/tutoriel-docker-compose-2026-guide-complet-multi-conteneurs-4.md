---
id: collect-261001-rattrapage/rattrapage/tutoriel-docker-compose-2026-guide-complet-multi-conteneurs-4
title: "Vérifier les versions installées"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["arr"]
source: docs/RAG/collect-261001-rattrapage/tutoriel-docker-compose-2026-guide-complet-multi-conteneurs.md
source_anchor: ""
source_lines: [407, 483]
sha256: 0dd3e4d50806e51855ba904aae101fc9860db246793a4d30620a7030d5b29339
---

# Vérifier les versions installées

```
# Construire et lancer tous les services en arrière-plan
docker compose up -d --build
# Sortie attendue :
# [+] Building 15.2s (12/12) FINISHED
#  => [api] docker image pull python:3.12-slim
#  => [api] COPY requirements.txt .
#  => [api] RUN pip install --no-cache-dir -r requirements.txt
#  => [api] COPY main.py .
# [+] Running 5/5
#  ✔ Network compose-tutorial_backend  Created
#  ✔ Volume "compose-tutorial_postgres_data" Created
#  ✔ Container tutorial-db     Healthy
#  ✔ Container tutorial-redis  Healthy
#  ✔ Container tutorial-api    Started
#  ✔ Container tutorial-nginx  Started
# Vérifier l'état de tous les services
docker compose ps
# Sortie attendue :
# NAME              IMAGE                 STATUS                   PORTS
# tutorial-api      compose-tutorial-api  Up 2 minutes             0.0.0.0:8000->8000/tcp
# tutorial-db       postgres:16-alpine    Up 2 minutes (healthy)   0.0.0.0:5432->5432/tcp
# tutorial-nginx    nginx:1.27-alpine     Up 2 minutes             0.0.0.0:80->80/tcp
# tutorial-redis    redis:7-alpine        Up 2 minutes (healthy)   0.0.0.0:6379->6379/tcp
# Consulter les logs de tous les services
docker compose logs --tail=20
# Tester l'API directement
curl http://localhost/health
# {"status":"healthy","db":"ok","cache":"ok"}
# Tester via Nginx
curl http://localhost/api/items
# {"source":"database","items":[{"id":3,"name":"Troisième élément",...}]}
# Créer un nouvel élément
curl -X POST http://localhost/api/items \
  -H "Content-Type: application/json" \
  -d '{"name": "Test Docker Compose", "description": "Créé via le tutoriel"}'
# {"id":4,"name":"Test Docker Compose"}
```
Si tous les services affichent le statut “Up” et que le health check retourne “healthy”, votre stack Docker Compose fonctionne correctement. Le flag `--build` force la reconstruction des images à chaque lancement, ce qui est utile pendant le développement. En production, vous omettrez ce flag et utiliserez des images pré-construites depuis un registre Docker.

## Étape 8 : Utiliser les Commandes Docker Compose Essentielles

Maîtriser les commandes Docker Compose est indispensable pour gérer efficacement vos applications au quotidien. En 2026, avec la transition complète vers Compose V2 (plugin intégré à Docker CLI), toutes les commandes utilisent la syntaxe `docker compose` avec un espace. Voici les commandes les plus importantes, classées par fréquence d’utilisation.

| Commande | Description | Usage courant | 
|---|---|---|
| `docker compose up -d` | Démarrer tous les services en arrière-plan | Lancement quotidien | 
| `docker compose down` | Arrêter et supprimer les conteneurs | Fin de journée | 
| `docker compose down -v` | Arrêter + supprimer les volumes | Reset complet des données | 
| `docker compose logs -f api` | Suivre les logs d’un service en temps réel | Débogage | 
| `docker compose exec db psql -U appuser appdb` | Ouvrir un shell PostgreSQL | Requêtes manuelles | 
| `docker compose restart api` | Redémarrer un service spécifique | Après modification de config | 
| `docker compose build --no-cache` | Reconstruire sans cache | Problèmes de dépendances | 
| `docker compose watch` | Activer le mode développement (Watch) | Développement avec hot-reload | 
| `docker compose top` | Afficher les processus en cours | Diagnostic de performance | 
| `docker compose config` | Valider et afficher la config résolue | Vérifier les variables | 

La commande `docker compose config` est particulièrement précieuse car elle affiche le fichier Compose avec toutes les variables d’environnement résolues. Si vous suspectez un problème de configuration, c’est la première commande à exécuter. Elle vous montrera exactement ce que Docker Compose interprète, incluant les valeurs par défaut et les substitutions de variables.

La fonctionnalité **Compose Watch**, introduite dans les versions récentes et améliorée en 2026, révolutionne l’expérience de développement. En exécutant `docker compose watch`, Docker surveille les fichiers locaux et synchronise automatiquement les modifications vers les conteneurs. L’action `sync+restart` que nous avons configurée dans notre `compose.yaml` synchronise le fichier modifié puis redémarre le service – idéal pour les applications Python qui nécessitent un rechargement.

## Étape 9 : Implémenter les Health Checks et la Gestion des Dépendances

Les health checks sont l’une des fonctionnalités les plus importantes de Docker Compose en production. Sans eux, vos conteneurs sont considérés comme “prêts” dès qu’ils démarrent, même si l’application à l’intérieur n’est pas encore opérationnelle. Cette situation provoque des erreurs en cascade : l’API démarre avant que PostgreSQL n’accepte les connexions, les requêtes échouent, et le conteneur redémarre en boucle.

Notre configuration utilise trois types de health checks adaptés à chaque service. Pour PostgreSQL, `pg_isready` vérifie que le serveur accepte les connexions. Pour Redis, `redis-cli ping` confirme que le cache répond. Pour l’API, un endpoint HTTP `/health` teste la connectivité complète vers la base de données et le cache. Les paramètres `interval`, `timeout`, `retries` et `start_period` permettent de contrôler finement le comportement de vérification.

Le paramètre `start_period` est particulièrement important pour PostgreSQL : il accorde un délai de grâce au conteneur pour s’initialiser avant de commencer à compter les échecs. Sans ce paramètre, un conteneur PostgreSQL chargeant un gros script d’initialisation pourrait être considéré comme défaillant et redémarré prématurément. En 2026, la valeur recommandée pour PostgreSQL est de 30 secondes, ce qui laisse le temps à la base de données de s’initialiser même sur des machines modestes.

La directive `depends_on` avec `condition: service_healthy` crée un chaînage intelligent : Docker Compose attend que chaque dépendance soit non seulement démarrée mais aussi en bonne santé avant de lancer le service suivant. C’est une amélioration majeure par rapport à la simple directive `depends_on` sans condition, qui ne garantissait que l’ordre de démarrage, pas la disponibilité effective du service.

## Étape 10 : Gérer les Volumes et la Persistance des Données

La persistance des données est un sujet fondamental dans Docker Compose. Par défaut, les données stockées dans un conteneur sont éphémères : elles disparaissent lorsque le conteneur est supprimé. Les volumes Docker résolvent ce problème en stockant les données sur le système hôte, indépendamment du cycle de vie des conteneurs.

Notre projet utilise deux types de volumes. Les **volumes nommés** (`postgres_data` et `redis_data`) sont gérés par Docker et stockent les données de manière persistante. Les **bind mounts** (`./db/init.sql` et `./nginx/nginx.conf`) lient directement un fichier ou répertoire de l’hôte dans le conteneur, ce qui est idéal pour les fichiers de configuration que vous modifiez fréquemment.

