---
id: collect-261001-rattrapage/rattrapage/tutoriel-docker-compose-2026-stack-prod-en-13-etapes-4
title: "Mise à jour des dépôts et installation des dépendances"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["arr", "memory"]
source: docs/RAG/collect-261001-rattrapage/tutoriel-docker-compose-2026-stack-prod-en-13-etapes.md
source_anchor: ""
source_lines: [401, 537]
sha256: 84423b70078ab88b4450bfe60a3eb3a1e645b0169e0cfb5159ea96ac90a22941
---

# Mise à jour des dépôts et installation des dépendances

```
# compose.override.yml
services:
  api:
    command: ["uvicorn", "main:app", "--host", "0.0.0.0", "--port", "8000", "--reload"]
    environment:
      - LOG_LEVEL=debug
    ports:
      - "8000:8000"  # Exposition directe en dev
    develop:
      watch:
        - action: sync
          path: ./app
          target: /home/appuser/app
          ignore:
            - __pycache__/
            - "*.pyc"
        - action: rebuild
          path: ./app/requirements.txt
  db:
    ports:
      - "5432:5432"  # Accès direct depuis psql sur l'hôte
  cache:
    ports:
      - "6379:6379"
```
```
# Lancer Compose Watch (mode dev)
docker compose up --watch
# Désormais, toute modification d'un .py dans ./app est synchronisée
# en moins de 200 ms vers le conteneur, et uvicorn --reload recharge.
# Si requirements.txt change, le conteneur est reconstruit automatiquement.
```
Compose Watch a deux *actions* principales : `sync` (copie les fichiers) et `rebuild` (reconstruit l’image). Une troisième, `sync+restart`, redémarre le conteneur après synchronisation. Pour Go, Rust ou Java compilés, `rebuild` est généralement préférable car le binaire doit être régénéré.

## Étape 11 : Construire et Démarrer la Stack Complète

Première mise en route. Le build de l’image FastAPI prend environ 30 à 60 secondes selon votre bande passante (téléchargement de Python 3.13 et des dépendances). Les images officielles PostgreSQL, Redis, Nginx, Prometheus et Grafana sont déjà très optimisées en taille. Notez que le Quickstart officiel de Docker, rafraîchi en juillet 2026, part désormais du principe que vous lancez directement `docker compose up` avec le plugin Compose v2.x le plus récent installé – ce tutoriel suit la même hypothèse.

```
# Build des images locales
docker compose build
# Démarrage en arrière-plan, avec monitoring
docker compose --profile monitoring up -d
# Surveiller les logs en direct (Ctrl+C pour quitter, conteneurs continuent)
docker compose logs -f --tail=50
# Vérifier l'état (RUNNING + healthy)
docker compose ps --format "table {{.Service}}\t{{.Status}}\t{{.Ports}}"
```
Sortie attendue après ~30 secondes (le temps que les healthchecks passent au vert) :

```
SERVICE     STATUS                    PORTS
api         Up 32 seconds (healthy)   8000/tcp
cache       Up 35 seconds             6379/tcp
db          Up 35 seconds (healthy)   5432/tcp
grafana     Up 35 seconds             0.0.0.0:3000->3000/tcp
prometheus  Up 35 seconds             9090/tcp
proxy       Up 30 seconds             0.0.0.0:80->80/tcp
```
```
# Tester l'API via Nginx
curl -s http://localhost/health
# {"status":"ok"}
curl -s -X POST http://localhost/items \
  -H "Content-Type: application/json" \
  -d '{"name":"Café Arabica","price":12.50}'
# {"id":1,"name":"Café Arabica","price":12.5}
curl -s http://localhost/items
# {"source":"db","items":[{"id":1,"name":"Café Arabica","price":"12.50"}]}
# Deuxième appel : devrait venir du cache Redis
curl -s http://localhost/items
# {"source":"cache","items":[...]}
```
## Étape 12 : Maîtriser les Commandes Compose Essentielles

Voici un aide-mémoire des commandes que vous utiliserez quotidiennement. La commande `docker compose` accepte les mêmes flags que `docker` sur la plupart des sous-commandes (`--quiet`, `--format json`, etc.). Le « Docker Compose Tutorial 2026 » de Vucense, mis à jour en août 2026, insiste d’ailleurs sur les deux mêmes commandes piliers pour les applications multi-conteneurs : `docker compose up -d` pour démarrer et `docker compose down` pour arrêter proprement. Pensez aussi à aligner vos pipelines CI : en février 2026, les runners GitHub Actions sont passés à Compose **v2.40.3**, un changement qui a modifié le comportement de certains scripts et exemples YAML dans les tutoriels s’appuyant sur ces images.

| Commande | Effet | Cas d’usage | 
|---|---|---|
| `docker compose up -d` | Démarre tous les services en arrière-plan | Démarrage normal | 
| `docker compose up --build` | Force la reconstruction des images | Après modification du Dockerfile | 
| `docker compose down` | Arrête et supprime conteneurs + réseaux | Fin de session dev | 
| `docker compose down -v` | Idem + supprime les volumes nommés | **Reset complet** (perte de données !) | 
| `docker compose restart api` | Redémarre un service unique | Après changement de config | 
| `docker compose logs -f api` | Logs temps réel d’un service | Debug | 
| `docker compose exec api bash` | Ouvre un shell dans le conteneur | Inspection en live | 
| `docker compose run --rm api pytest` | Lance une commande one-shot | Tests, migrations | 
| `docker compose ps` | Liste les services et leur état | Monitoring rapide | 
| `docker compose top` | Affiche les processus dans chaque conteneur | Investigation perf | 
| `docker compose pull` | Met à jour les images depuis le registre | Maintenance | 
| `docker compose config` | Valide et affiche le fichier final résolu | Debug YAML / variables | 

**Astuce avancée :** `docker compose config --services` liste uniquement les noms de services, parfait pour itérer en script Bash. `docker compose config --hash '*'` calcule un hash unique de la configuration : utile en CI pour décider si l’on doit re-déployer.

## Étape 13 : Déployer en Production avec Compose

Pour un déploiement réel, suivez cette checklist. Compose en production reste viable jusqu’à environ 10-20 services sur un hôte unique (ou 2-3 hôtes via Docker Swarm). Au-delà, basculez sur Kubernetes – voir notre comparatif Docker vs Kubernetes.

- **Tag d’image immuable :** ne déployez jamais`:latest` . Utilisez le SHA Git ou un tag versionné (`v1.4.2` ) pour garantir la reproductibilité.
- **Fichiers séparés :**`docker-compose.yml` (base) +`compose.prod.yml` (overrides prod). Lancez avec`docker compose -f docker-compose.yml -f compose.prod.yml up -d` .
- **Logs centralisés :** configurez`logging.driver: journald` ou`fluentd` au lieu du driver`json-file` par défaut, qui peut saturer le disque.
- **Restart policy stricte :**`restart: unless-stopped` partout, sauf pour les jobs one-shot.
- **Limites de ressources :** définissez`cpus` et`memory` sur chaque service pour éviter qu’un OOM dans un conteneur ne fasse tomber le voisin.
- **Sauvegardes :** backup nocturne de chaque named volume via`docker run --rm --volumes-from db -v $(pwd):/backup alpine tar czf /backup/db-$(date +%F).tgz /var/lib/postgresql/data` .
- **Mises à jour rolling :**`docker compose pull && docker compose up -d --no-deps --build api` ne touche que le service ciblé.
- **Reverse proxy externe :** exposez uniquement Nginx/Caddy/Traefik. Tous les autres services restent sur le réseau`backend` sans`ports:` .

```
# compose.prod.yml (override production)
services:
  api:
    image: registry.example.com/myapp/api:${GIT_SHA:-latest}
    build: !reset null  # Désactive le build, force le pull
    logging:
      driver: "json-file"
      options:
        max-size: "10m"
        max-file: "3"
  proxy:
    ports:
      - "80:80"
      - "443:443"
    volumes:
      - ./certs:/etc/nginx/certs:ro
# Déploiement
GIT_SHA=$(git rev-parse --short HEAD) docker compose \
  -f docker-compose.yml -f compose.prod.yml pull
GIT_SHA=$(git rev-parse --short HEAD) docker compose \
  -f docker-compose.yml -f compose.prod.yml up -d --remove-orphans
```
Le drapeau `--remove-orphans` est *essentiel* : il supprime les conteneurs d’anciens services qui ne figurent plus dans le fichier YAML. Sans lui, vous accumulez des zombies à chaque refactor.

## Pièges Courants et Erreurs Classiques à Éviter

Au-delà des deux pièges déjà mentionnés, voici les erreurs récurrentes observées dans les revues de code Docker Compose, classées par fréquence.

