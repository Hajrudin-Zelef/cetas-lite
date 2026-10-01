---
id: collect-261001-cisco/cisco/tutoriel-n8n-2026-automatiser-vos-workflows-ia-4
title: "Vérifier le statut du conteneur"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["arr"]
source: docs/RAG/collect-261001-cisco/tutoriel-n8n-2026-automatiser-vos-workflows-ia.md
source_anchor: ""
source_lines: [334, 488]
sha256: 74a6ee77ff3132f8e1971bf7cb05ed9cc7925d3b7f8e8920155b9ab361b27dea
---

# Vérifier le statut du conteneur

**Retry on Fail** : Chaque nœud peut être configuré pour retenter automatiquement en cas d'échec. Dans les paramètres du nœud, activez « Retry On Fail » et définissez le nombre maximum de tentatives (3 est une valeur sûre) et le délai entre les tentatives. Utilisez un délai exponentiel pour les API externes qui pourraient être temporairement surchargées.

**Error Output** : Depuis n8n 2.0, chaque nœud dispose d'une sortie d'erreur séparée. Connectez cette sortie à un flux de traitement alternatif plutôt que de laisser le workflow s'arrêter complètement. Par exemple, si l'envoi d'un email échoue, enregistrez le message dans une file d'attente pour un envoi ultérieur.

Pour le monitoring avancé, n8n s'intègre nativement avec les outils d'observabilité. Ajoutez ces variables d'environnement à votre `docker-compose.yml` pour activer le streaming de logs vers un SIEM :

```
# Variables d'environnement pour le monitoring n8n
# À ajouter dans la section environment du service n8n
# Activer les métriques Prometheus
- N8N_METRICS=true
- N8N_METRICS_PREFIX=n8n_
# Log structuré en JSON pour ingestion par ELK/Loki
- N8N_LOG_OUTPUT=console
- N8N_LOG_LEVEL=warn
# Webhook de notification d'erreurs (optionnel)
- N8N_ERROR_TRIGGER_TYPE=all
# Pour exposer le endpoint /metrics pour Prometheus
# Accessible sur http://localhost:5678/metrics
# Métriques disponibles :
#   n8n_workflow_success_total
#   n8n_workflow_error_total  
#   n8n_workflow_production_active
#   n8n_workflow_execution_duration_seconds
```
Avec les métriques Prometheus activées, vous pouvez créer des dashboards Grafana pour surveiller le taux d'erreur, la durée d'exécution et le nombre de workflows actifs. Configurez des alertes quand le taux d'erreur dépasse 5 % ou quand un workflow critique n'a pas été exécuté depuis plus de 2 heures. Pour aller plus loin avec la supervision d'infrastructure, notre tutoriel GitHub Actions CI/CD montre comment intégrer le monitoring dans votre pipeline de déploiement continu.

## Étape 9 : Déployer n8n en Production avec HTTPS et Reverse Proxy

Le passage en production nécessite une configuration sécurisée avec HTTPS, un nom de domaine et un reverse proxy. Caddy est le choix recommandé en 2026 pour sa simplicité : il gère automatiquement les certificats SSL via Let's Encrypt sans configuration supplémentaire. Cette étape est indispensable pour que les webhooks fonctionnent avec des services tiers qui exigent HTTPS.

Voici le fichier `docker-compose.prod.yml` complet pour un déploiement production :

```
version: '3.8'
services:
  caddy:
    image: caddy:2-alpine
    container_name: n8n-caddy
    restart: unless-stopped
    ports:
      - "80:80"
      - "443:443"
    volumes:
      - ./Caddyfile:/etc/caddy/Caddyfile:ro
      - caddy_data:/data
      - caddy_config:/config
    networks:
      - n8n-network
  postgres:
    image: postgres:16-alpine
    container_name: n8n-postgres
    restart: unless-stopped
    environment:
      POSTGRES_USER: ${DB_USER}
      POSTGRES_PASSWORD: ${DB_PASSWORD}
      POSTGRES_DB: n8n
    volumes:
      - postgres_data:/var/lib/postgresql/data
    networks:
      - n8n-network
    healthcheck:
      test: ["CMD-SHELL", "pg_isready -U ${DB_USER} -d n8n"]
      interval: 10s
      timeout: 5s
      retries: 5
  n8n:
    image: n8nio/n8n:latest
    container_name: n8n
    restart: unless-stopped
    environment:
      - DB_TYPE=postgresdb
      - DB_POSTGRESDB_HOST=postgres
      - DB_POSTGRESDB_PORT=5432
      - DB_POSTGRESDB_DATABASE=n8n
      - DB_POSTGRESDB_USER=${DB_USER}
      - DB_POSTGRESDB_PASSWORD=${DB_PASSWORD}
      - N8N_HOST=${N8N_DOMAIN}
      - N8N_PORT=5678
      - N8N_PROTOCOL=https
      - WEBHOOK_URL=https://${N8N_DOMAIN}/
      - GENERIC_TIMEZONE=Europe/Paris
      - TZ=Europe/Paris
      - N8N_LOG_LEVEL=warn
      - N8N_METRICS=true
      - EXECUTIONS_DATA_PRUNE=true
      - EXECUTIONS_DATA_MAX_AGE=336
      - N8N_DIAGNOSTICS_ENABLED=false
      - N8N_ENCRYPTION_KEY=${ENCRYPTION_KEY}
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
  caddy_data:
  caddy_config:
networks:
  n8n-network:
    driver: bridge
```
Créez le fichier `Caddyfile` à la racine de votre projet :

```
# Caddyfile pour n8n en production
# Remplacez n8n.votredomaine.fr par votre domaine
n8n.votredomaine.fr {
    reverse_proxy n8n:5678 {
        flush_interval -1
        transport http {
            keepalive 30s
        }
    }
    
    header {
        X-Content-Type-Options nosniff
        X-Frame-Options SAMEORIGIN
        Referrer-Policy strict-origin-when-cross-origin
        Strict-Transport-Security "max-age=31536000; includeSubDomains"
    }
    
    log {
        output file /data/access.log {
            roll_size 10mb
            roll_keep 5
        }
    }
}
```
Créez un fichier `.env` pour les variables sensibles (ne commitez jamais ce fichier dans Git) :

```
# .env - Variables d'environnement production
# IMPORTANT : Ne jamais committer ce fichier dans Git !
N8N_DOMAIN=n8n.votredomaine.fr
DB_USER=n8n_prod
DB_PASSWORD=MotDePasseUltraSecurise2026!GenerezAvecOpenSSL
ENCRYPTION_KEY=GenerezAvec_openssl_rand_hex_32
# Pour générer la clé de chiffrement :
# openssl rand -hex 32
```
Avant de lancer, assurez-vous que votre domaine pointe vers l'IP de votre serveur (enregistrement DNS A). Lancez ensuite avec `docker compose -f docker-compose.prod.yml up -d`. Caddy obtiendra automatiquement un certificat SSL Let's Encrypt en quelques secondes. Vérifiez avec `curl -I https://n8n.votredomaine.fr/healthz` que vous obtenez un code 200 avec un certificat valide.

Pour les hébergeurs français et européens, OVH Cloud, Scaleway et Hetzner offrent des VPS à partir de 4-6 €/mois amplement suffisants pour une instance n8n de petite à moyenne taille (jusqu'à 50 000 exécutions/mois). Cette approche self-hosted revient nettement moins cher que le plan cloud n8n à 20 $/mois tout en gardant le contrôle total sur les données — un point crucial pour la conformité RGPD.

## Étape 10 : Sauvegarder et Versionner les Workflows avec Git

En environnement professionnel, vos workflows n8n représentent de la logique métier critique. Les perdre équivaut à perdre du code de production. n8n 2.0 intègre nativement le versionnement Git, permettant de synchroniser automatiquement vos workflows avec un dépôt GitHub ou GitLab. Cette fonctionnalité, disponible dans les plans Team et Enterprise, transforme n8n en un outil véritablement adapté aux pratiques DevOps modernes.

Pour les installations self-hosted avec le plan Community, vous pouvez exporter et versionner manuellement vos workflows via l'API REST de n8n :

