---
id: collect-261001-rattrapage/rattrapage/tutoriel-docker-compose-2026-stack-prod-en-13-etapes-3
title: "Mise à jour des dépôts et installation des dépendances"
domain: rattrapage
role: reference
task: reference
actors: ["AWS"]
dates: []
keywords: ["aws", "memory"]
source: docs/RAG/collect-261001-rattrapage/tutoriel-docker-compose-2026-stack-prod-en-13-etapes.md
source_anchor: ""
source_lines: [204, 400]
sha256: 2c7fa92cdf67497b3710a3e95d3506e9450a039cd3b8cc16383854d712496ab6
---

# Mise à jour des dépôts et installation des dépendances

```
# docker-compose.yml
name: compose-prod-stack
services:
  api:
    build:
      context: ./app
      target: runtime
    image: compose-demo/api:latest
    restart: unless-stopped
    environment:
      - REDIS_URL=redis://cache:6379/0
    secrets:
      - db_password
    networks:
      - backend
    depends_on:
      db:
        condition: service_healthy
      cache:
        condition: service_started
    deploy:
      resources:
        limits: {cpus: '1.0', memory: 512M}
        reservations: {cpus: '0.25', memory: 128M}
    healthcheck:
      test: ["CMD", "curl", "-f", "http://localhost:8000/health"]
      interval: 30s
      timeout: 5s
      retries: 3
  db:
    image: postgres:17-alpine
    restart: unless-stopped
    environment:
      POSTGRES_USER: postgres
      POSTGRES_PASSWORD_FILE: /run/secrets/db_password
      POSTGRES_DB: postgres
    secrets:
      - db_password
    volumes:
      - db_data:/var/lib/postgresql/data
      - ./postgres/init.sql:/docker-entrypoint-initdb.d/init.sql:ro
    networks:
      - backend
    healthcheck:
      test: ["CMD-SHELL", "pg_isready -U postgres"]
      interval: 10s
      timeout: 3s
      retries: 5
  cache:
    image: redis:7.4-alpine
    restart: unless-stopped
    command: ["redis-server", "--maxmemory", "256mb", "--maxmemory-policy", "allkeys-lru"]
    volumes:
      - cache_data:/data
    networks:
      - backend
  proxy:
    image: nginx:1.27-alpine
    restart: unless-stopped
    ports:
      - "80:80"
    volumes:
      - ./nginx/default.conf:/etc/nginx/conf.d/default.conf:ro
    depends_on:
      api:
        condition: service_healthy
    networks:
      - backend
      - frontend
  prometheus:
    image: prom/prometheus:v3.0.1
    restart: unless-stopped
    profiles: ["monitoring"]
    volumes:
      - ./prometheus/prometheus.yml:/etc/prometheus/prometheus.yml:ro
      - prom_data:/prometheus
    networks:
      - backend
  grafana:
    image: grafana/grafana:11.4.0
    restart: unless-stopped
    profiles: ["monitoring"]
    ports:
      - "3000:3000"
    environment:
      GF_SECURITY_ADMIN_PASSWORD__FILE: /run/secrets/grafana_admin
    secrets:
      - grafana_admin
    volumes:
      - grafana_data:/var/lib/grafana
      - ./grafana/provisioning:/etc/grafana/provisioning:ro
    networks:
      - backend
volumes:
  db_data:
  cache_data:
  prom_data:
  grafana_data:
networks:
  backend:
    driver: bridge
  frontend:
    driver: bridge
secrets:
  db_password:
    file: ./secrets/db_password.txt
  grafana_admin:
    file: ./secrets/grafana_admin.txt
```
Notez la directive `name:` en haut du fichier. Elle remplace le nom dérivé du dossier (`compose-prod-stack`) et garantit que vos volumes et conteneurs auront un préfixe stable même si vous renommez le répertoire. C’est crucial pour des migrations ou des sauvegardes prévisibles.

## Étape 7 : Gérer les Secrets de Manière Sécurisée

Compose supporte les **secrets** via la section `secrets:`. Chaque secret est monté en lecture seule sous `/run/secrets/<nom>` dans les conteneurs qui le déclarent. Contrairement aux variables d’environnement, les secrets *n’apparaissent pas* dans `docker inspect`, `ps -auxe` ou les logs. C’est la méthode officielle recommandée depuis Compose v2.

```
# Génération de mots de passe forts
openssl rand -base64 32 | tr -d '\n' > secrets/db_password.txt
openssl rand -base64 24 | tr -d '\n' > secrets/grafana_admin.txt
# Restriction des permissions
chmod 600 secrets/*.txt
# Vérifier que les fichiers ne contiennent pas de saut de ligne final
xxd secrets/db_password.txt | tail -1
```
**Piège n° 2 :** Si `db_password.txt` contient un `\n` final (très fréquent avec `echo`), PostgreSQL refusera la connexion avec une erreur cryptique. Utilisez `echo -n` ou la pipe `tr -d '\n'` comme ci-dessus. Vérifiez avec `xxd` ou `wc -c` que la taille correspond exactement à ce que vous attendez.

Pour la production réelle, ne stockez *jamais* ces fichiers dans le dépôt Git. Utilisez plutôt un coffre-fort comme HashiCorp Vault, AWS Secrets Manager, ou le binding natif Compose à des *external secrets* (`secrets: db_password: external: true`) couplé à `docker secret create`.

## Étape 8 : Configurer Nginx en Reverse Proxy

Nginx fait l’interface entre le monde extérieur et l’API FastAPI. Il gère les en-têtes *X-Forwarded-For*, applique un *rate limiting* simple et compresse les réponses. Pour HTTPS en production, ajoutez Certbot ou utilisez **Caddy** à la place (voir notre comparatif Caddy vs Nginx 2026).

```
# nginx/default.conf
upstream api_backend {
    server api:8000 max_fails=3 fail_timeout=30s;
    keepalive 32;
}
limit_req_zone $binary_remote_addr zone=api_limit:10m rate=10r/s;
server {
    listen 80 default_server;
    server_name _;
    client_max_body_size 10m;
    gzip on;
    gzip_types application/json text/plain text/css application/javascript;
    location /health {
        access_log off;
        return 200 "ok\n";
    }
    location / {
        limit_req zone=api_limit burst=20 nodelay;
        proxy_pass http://api_backend;
        proxy_http_version 1.1;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
        proxy_connect_timeout 5s;
        proxy_read_timeout 30s;
    }
}
```
Notez que `api` dans `upstream` fait référence au nom du service Compose : grâce au DNS interne du réseau `backend`, Nginx résout automatiquement `api` vers l’IP du conteneur. C’est l’un des plus grands avantages de Compose : pas besoin d’IP statique ni de fichier `hosts`.

## Étape 9 : Activer le Monitoring avec Prometheus et Grafana (Profiles)

Les **profiles** permettent d’activer ou non certains services selon le contexte. Dans notre fichier, `prometheus` et `grafana` sont taggés `profiles: ["monitoring"]`. Ils ne démarrent que si l’on lance Compose avec `--profile monitoring`. Idéal pour ne pas charger inutilement la RAM en développement.

```
# prometheus/prometheus.yml
global:
  scrape_interval: 15s
  evaluation_interval: 15s
scrape_configs:
  - job_name: 'fastapi'
    metrics_path: /metrics
    static_configs:
      - targets: ['api:8000']
  - job_name: 'prometheus'
    static_configs:
      - targets: ['localhost:9090']
```
```
# Démarrer la stack sans monitoring (par défaut)
docker compose up -d
# Démarrer la stack AVEC monitoring
docker compose --profile monitoring up -d
# Vérifier les services actifs
docker compose ps
```
Ouvrez Grafana à http://localhost:3000 avec l’utilisateur `admin` et le mot de passe contenu dans `secrets/grafana_admin.txt`. Ajoutez Prometheus comme source de données (`http://prometheus:9090`) et importez un dashboard FastAPI public depuis grafana.com (par exemple le dashboard ID 14282).

## Étape 10 : Override pour le Développement avec Compose Watch

Le fichier `compose.override.yml` est lu *automatiquement* par Compose en plus du fichier principal, sans option supplémentaire. C’est la convention idéale pour ajuster le comportement en local sans toucher au fichier de production. Couplé à **Compose Watch**, vous obtenez du rechargement à chaud comparable à `nodemon` ou `uvicorn --reload`, mais multi-langage.

