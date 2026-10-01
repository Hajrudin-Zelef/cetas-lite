---
id: collect-261001-rattrapage/rattrapage/docker-guide-8
title: "Guide Docker complet — Production & Sysadmin"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["apache"]
source: docs/RAG/collect-261001-rattrapage/docker_guide.md
source_anchor: ""
source_lines: [1824, 2094]
sha256: 65c8472a4a072739a6b6cfce42dbd2b414496acae5ba53a51b794a7da2e659b9
---

# Déployer une stack répliquée
docker stack deploy -c compose.yaml mon-app

# Mettre à l'échelle
docker service scale mon-app_api=3

# Voir les tâches (quel conteneur tourne sur quel nœud)
docker service ps mon-app_api
docker stack services mon-app
```

compose.yaml adapté à Swarm :

```yaml
services:
  api:
    image: registry.interne:5000/mon-app:1.4.2
    deploy:
      replicas: 3
      update_config:
        parallelism: 1
        delay: 10s
        failure_action: rollback
      restart_policy:
        condition: on-failure
```

Limites à connaître :

- Swarm est **en maintenance** chez Docker (pas de grosses nouveautés) ;
  pour du gros multi-hôtes moderne, Kubernetes domine.
- Mais pour 2–5 serveurs, Swarm reste **simple, robuste et suffisant**.
- Les `build:` locaux ne marchent pas en Swarm : poussez d'abord l'image
  dans un registre accessible par tous les nœuds.

## 66. Portainer : administrer via une interface web

[Portainer](https://www.portainer.io/) : UI web pour voir/gérer conteneurs,
images, volumes, réseaux, stacks compose.

```yaml
services:
  portainer:
    image: portainer/portainer-ce:2.21.0
    container_name: portainer
    ports:
      - "127.0.0.1:9443:9443"   # HTTPS ; exposez via reverse proxy + auth
    volumes:
      - /var/run/docker.sock:/var/run/docker.sock  # ⚠️ voir section 55
      - portainer-data:/data
    restart: unless-stopped

volumes:
  portainer-data:
```

Règles d'usage :

- **Jamais exposé tel quel sur Internet** : derrière un reverse proxy avec
  TLS + authentification forte (et idéalement restriction IP/VPN).
- Pratique pour : visualiser l'état, déployer une stack compose via l'UI,
  donner un accès **limité** (rôles) à une équipe sans accès SSH.
- Ne remplace pas la maîtrise de la CLI pour le dépannage.

## 67. Reverse proxy : Traefik (recommandé)

Traefik découvre **automatiquement** les conteneurs via des labels et génère
les certificats Let's Encrypt tout seul.

```yaml
services:
  traefik:
    image: traefik:v3.1
    container_name: traefik
    command:
      - --providers.docker=true
      - --providers.docker.exposedbydefault=false
      - --entrypoints.web.address=:80
      - --entrypoints.websecure.address=:443
      - --certificatesresolvers.letsencrypt.acme.httpchallenge.entrypoint=web
      - --certificatesresolvers.letsencrypt.acme.email=admin@entreprise.fr
      - --certificatesresolvers.letsencrypt.acme.storage=/letsencrypt/acme.json
    ports:
      - "80:80"
      - "443:443"
    volumes:
      - /var/run/docker.sock:/var/run/docker.sock:ro  # ⚠️ lecture seule !
      - traefik-data:/letsencrypt
    restart: unless-stopped

  web:
    image: nginx:1.27-alpine
    labels:
      - traefik.enable=true
      - traefik.http.routers.web.rule=Host(`www.entreprise.fr`)
      - traefik.http.routers.web.entrypoints=websecure
      - traefik.http.routers.web.tls.certresolver=letsencrypt
      - traefik.http.services.web.loadbalancer.server.port=80
    networks:
      - proxy

networks:
  proxy:
    name: proxy
    external: true

volumes:
  traefik-data:
```

```bash
docker network create proxy   # réseau partagé (voir section 32)
chmod 600 letsencrypt/acme.json  # si créé à la main
```

## 68. Reverse proxy : alternative nginx manuel

Si vous préférez un nginx classique devant vos conteneurs :

```nginx
# /srv/proxy/nginx.conf
upstream api { server api:8000; }

server {
    listen 80;
    server_name api.entreprise.fr;
    location / {
        proxy_pass http://api;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
    }
}
```

```yaml
services:
  proxy:
    image: nginx:1.27-alpine
    ports: ["80:80", "443:443"]
    volumes:
      - ./nginx.conf:/etc/nginx/conf.d/default.conf:ro
      - ./certs:/etc/nginx/certs:ro
    depends_on: [api]
```

Certificats : utilisez certbot sur l'hôte ou le compagnon
`nginxproxy/acme-companion` pour automatiser Let's Encrypt.

## 69. Cas pratique : stack LAMP complète (commentée)

```yaml
name: lamp

services:
  web:
    image: php:8.3-apache
    container_name: lamp-web
    ports:
      - "127.0.0.1:8080:80"      # exposé en local ; le reverse proxy prendra le relais
    volumes:
      - ./www:/var/www/html:ro              # code source (lecture seule)
      - ./config/php.ini:/usr/local/etc/php/php.ini:ro
      - ./config/vhost.conf:/etc/apache2/sites-enabled/000-default.conf:ro
    environment:
      DB_HOST: db
      DB_NAME: appdb
      DB_USER: app
      # mot de passe via secret, jamais en clair :
    secrets: [db_password]
    depends_on:
      db: { condition: service_healthy }
    healthcheck:
      test: ["CMD", "curl", "-f", "http://localhost/"]
      interval: 30s
      timeout: 5s
      retries: 3
    restart: unless-stopped
    networks: [front, back]

  db:
    image: mariadb:11.4
    container_name: lamp-db
    environment:
      MARIADB_DATABASE: appdb
      MARIADB_USER: app
      MARIADB_PASSWORD_FILE: /run/secrets/db_password
      MARIADB_ROOT_PASSWORD_FILE: /run/secrets/db_root_password
    secrets: [db_password, db_root_password]
    volumes:
      - db-data:/var/lib/mysql          # DONNÉES persistantes (à sauvegarder !)
      - ./backups:/backups
    healthcheck:
      test: ["CMD", "healthcheck.sh", "--connect", "--innodb_initialized"]
      interval: 10s
      timeout: 5s
      retries: 5
    restart: unless-stopped
    networks: [back]                    # pas d'accès direct depuis l'extérieur

  phpmyadmin:
    image: phpmyadmin:5.2
    container_name: lamp-pma
    environment:
      PMA_HOST: db
      PMA_USER: app
      PMA_PASSWORD_FILE: /run/secrets/db_password
    secrets: [db_password]
    ports:
      - "127.0.0.1:8081:80"            # admin locale uniquement
    depends_on:
      db: { condition: service_healthy }
    restart: unless-stopped
    networks: [back]
    profiles: ["admin"]                 # ne démarre qu'avec --profile admin

networks:
  front:
  back:
    internal: true

volumes:
  db-data:

secrets:
  db_password:      { file: ./secrets/db_password.txt }
  db_root_password: { file: ./secrets/db_root_password.txt }
```

```bash
docker compose up -d
docker compose --profile admin up -d   # avec phpMyAdmin
curl http://localhost:8080/
# Sauvegarde :
docker compose exec db sh -c 'mariadb-dump -u root -p"$MARIADB_ROOT_PASSWORD" appdb' \
  | gzip > backups/appdb-$(date +%F).sql.gz
```

## 70. Cas pratique : Nextcloud (fichier + base + cache)

```yaml
name: nextcloud

services:
  app:
    image: nextcloud:29-apache
    container_name: nc-app
    ports: ["127.0.0.1:8080:80"]
    volumes:
      - nc-html:/var/www/html          # code + config + data par défaut
      - ./data:/var/www/html/data      # données utilisateurs (bind, simple à sauvegarder)
    environment:
      MYSQL_HOST: db
      MYSQL_DATABASE: nextcloud
      MYSQL_USER: nc
      MYSQL_PASSWORD_FILE: /run/secrets/db_password
      NEXTCLOUD_ADMIN_USER_FILE: /run/secrets/admin_user
      NEXTCLOUD_ADMIN_PASSWORD_FILE: /run/secrets/admin_password
      REDIS_HOST: cache
    secrets: [db_password, admin_user, admin_password]
    depends_on:
      db: { condition: service_healthy }
      cache: { condition: service_started }
    healthcheck:
      test: ["CMD", "curl", "-f", "http://localhost/status.php"]
      interval: 60s
      timeout: 10s
      retries: 3
    restart: unless-stopped

