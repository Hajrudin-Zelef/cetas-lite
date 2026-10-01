---
id: collect-261001-rattrapage/rattrapage/docker-guide-7
title: "Guide Docker complet — Production & Sysadmin"
domain: rattrapage
role: reference
task: reference
actors: ["Microsoft"]
dates: ["2026-09-26"]
keywords: ["agent", "arr", "memory", "open source", "scout"]
source: docs/RAG/collect-261001-rattrapage/docker_guide.md
source_anchor: ""
source_lines: [1575, 1823]
sha256: 0d1276ab4dc612b2bec1c299dbca50f4c2c07a03b2cd9c0ab0a51ad88b7d9f61
---

# Guide Docker complet — Production & Sysadmin

```bash
# Docker Scout (intégré aux versions récentes)
docker scout quickview mon-app:1.0
docker scout cves mon-app:1.0

# Trivy (open source, recommandé en CI)
trivy image mon-app:1.0
trivy image --severity HIGH,CRITICAL mon-app:1.0
```

- Rebuildez régulièrement les images (les CVE se corrigent dans les bases
  Debian/Alpine) : au moins mensuellement, ou à chaque CVE critique.
- Signature : vérifiez la provenance (`DOCKER_CONTENT_TRUST=1` / cosign)
  pour les images sensibles.

## 58. Supervision — métriques du daemon et des conteneurs

Activer les métriques Prometheus du daemon (`daemon.json`, voir section 6) :

```json
{ "metrics-addr": "127.0.0.1:9323" }
```

```bash
curl -s 127.0.0.1:9323/metrics | grep -E "^engine_daemon_container"
```

Stack de supervision recommandée :

```yaml
services:
  prometheus:
    image: prom/prometheus:v2.53.0
    volumes: ["./prometheus.yml:/etc/prometheus/prometheus.yml:ro", "prom-data:/prometheus"]
    ports: ["127.0.0.1:9090:9090"]
  cadvisor:
    image: gcr.io/cadvisor/cadvisor:v0.49.1
    volumes:
      - /:/rootfs:ro
      - /var/run:/var/run:ro
      - /sys:/sys:ro
      - /var/lib/docker/:/var/lib/docker:ro
    ports: ["127.0.0.1:8080:8080"]
  grafana:
    image: grafana/grafana:11.1.0
    volumes: ["grafana-data:/var/lib/grafana"]
    ports: ["127.0.0.1:3000:3000"]

volumes: { prom-data:, grafana-data: }
```

Métriques clés à alerter : `container_memory_usage_bytes` proche de la limite,
`container_cpu_usage_seconds_total` en saturation, `engine_daemon_container_states`
(conteneurs `unhealthy`), espace `/var/lib/docker` > 80 %.

## 59. Supervision — healthchecks et alertes

- Chaque service critique a un `healthcheck` (sections 18, 30).
- Remontée d'alertes : un exporteur (Prometheus Alertmanager), ou plus simple :

```bash
# Cron : alerte si un conteneur est unhealthy ou restart en boucle
#!/bin/bash
UNHEALTHY=$(docker ps --filter "health=unhealthy" --format '{{.Names}}')
RESTARTING=$(docker ps --filter "status=restarting" --format '{{.Names}}')
[ -n "$UNHEALTHY$RESTARTING" ] && \
  echo "Docker alerte: unhealthy=[$UNHEALTHY] restarting=[$RESTARTING]" | \
  mail -s "[PROD] Docker alerte" astreinte@entreprise.fr
```

```bash
# Voir l'historique des healthchecks d'un conteneur
docker inspect --format '{{range .State.Health.Log}}{{.Output}}{{end}}' api | tail -5
```

## 60. Journalisation centralisée (option production)

```yaml
services:
  api:
    logging:
      driver: gelf
      options:
        gelf-address: "udp://192.168.1.20:12201"
        tag: "prod-api"
```

Alternatives : `syslog` vers rsyslog central, `journald` + forwarder,
ou un agent (Fluent Bit / Promtail) qui lit les `json-file` locaux et les
expédie vers Loki/Elasticsearch. Gardez **toujours** une copie locale avec
rotation en secours (si le collecteur distant tombe, vous gardez les logs).

## 61. SAUVEGARDE — les volumes (le point critique)

> Un conteneur se reconstruit en une commande. **Les volumes, eux, contiennent
> vos données.** Pas de sauvegarde des volumes = pas de PRA.

Méthode 1 — conteneur helper (recommandée, à chaud, sans arrêter le service) :

```bash
# Sauvegarde d'un volume nommé vers un .tar.gz daté
docker run --rm \
  -v db-data:/data:ro \
  -v /srv/backups:/backup \
  alpine:3.20 \
  tar czf /backup/db-data-$(date +%F).tar.gz -C /data .

# Sauvegarde de TOUS les volumes d'une stack compose (depuis son dossier)
docker compose run --rm -v /srv/backups:/backup db \
  sh -c 'tar czf /backup/db-$(date +%F_%H%M).tar.gz -C /var/lib/postgresql/data .'
```

Méthode 2 — dump applicatif (complément indispensable pour les SGBD) :

```bash
# Dump PostgreSQL AVANT la sauvegarde du volume (cohérence garantie)
docker compose exec db pg_dump -U app appdb | gzip > /srv/backups/appdb-$(date +%F).sql.gz

# Dump MySQL/MariaDB
docker compose exec db sh -c 'mysqldump -u root -p"$MARIADB_ROOT_PASSWORD" --all-databases' \
  | gzip > /srv/backups/mariadb-$(date +%F).sql.gz
```

Restauration :

```bash
# Restaurer un volume depuis l'archive
docker run --rm \
  -v db-data:/data \
  -v /srv/backups:/backup:ro \
  alpine:3.20 \
  sh -c 'rm -rf /data/* && tar xzf /backup/db-data-2026-09-26.tar.gz -C /data'

# Restaurer un dump SQL
gunzip -c /srv/backups/appdb-2026-09-26.sql.gz | docker compose exec -T db psql -U app appdb
```

Stratégie 3-2-1 adaptée à Docker :

- [ ] Dump applicatif + archive du volume, **quotidiens**, horodatés
- [ ] 3 copies : hôte + NAS/synology + hors site (rclone vers S3/Backblaze)
- [ ] Rétention : 7 jours glissants + 4 semaines + 3 mois (script de purge)
- [ ] **Test de restauration trimestriel** (une sauvegarde non testée n'existe pas)
- [ ] `docker compose down -v` **interdit** en prod sans double confirmation

Script de purge type :

```bash
#!/bin/bash
# purge_backups.sh : garde 7 jours
find /srv/backups -name "*.tar.gz" -mtime +7 -delete
find /srv/backups -name "*.sql.gz" -mtime +7 -delete
```

## 62. Mises à jour — les images

Procédure standard (zéro surprise) :

```bash
cd /srv/stacks/mon-app
# 1. Sauvegarde (volumes + dumps) AVANT
./backup.sh
# 2. Récupérer les nouvelles images
docker compose pull
# 3. Relancer (recrée uniquement les conteneurs dont l'image a changé)
docker compose up -d
# 4. Vérifier la santé
docker compose ps
docker compose logs --tail=50 -f
# 5. Nettoyer les vieilles images
docker image prune -a -f --filter "until=72h"
```

Rollback express :

```bash
# Revenir au tag précédent (d'où l'intérêt des tags versionnés !)
# éditer compose.yaml : image: mon-app:1.4.1 (au lieu de 1.4.2)
docker compose up -d
```

> Avec `latest`, impossible de savoir vers quoi on revient : **tags versionnés
> obligatoires** (voir section 22).

## 63. Mises à jour — automatiques ? Watchtower avec prudence

[Watchtower](https://containrrr.dev/watchtower/) met à jour les conteneurs
automatiquement quand une nouvelle image sort.

```yaml
services:
  watchtower:
    image: containrrr/watchtower:1.7.1
    volumes:
      - /var/run/docker.sock:/var/run/docker.sock  # ⚠️ risque assumé (section 55)
    environment:
      WATCHTOWER_SCHEDULE: "0 0 4 * * *"   # tous les jours à 4h
      WATCHTOWER_ROLLING_RESTART: "true"    # un par un, pas tout d'un coup
      WATCHTOWER_INCLUDE_STOPPED: "false"
    restart: unless-stopped
```

⚠️ **Prudence en production** :

- Réservé aux services **non critiques** ou aux environnements de test.
- Une mise à jour auto peut casser un service à 4h du matin sans prévenir.
- En prod critique : mises à jour **manuelles**, en heures ouvrées, avec
  sauvegarde + procédure de rollback (section 62).
- Alternative : Watchtower en mode **notification seule**
  (`WATCHTOWER_MONITOR_ONLY=true`) → vous décidez.

## 64. Mises à jour — le moteur Docker lui-même

```bash
# 1. Sauvegarder les stacks (compose.yaml + volumes, section 61)
# 2. Mettre à jour via le dépôt officiel
sudo apt-get update
sudo apt-get install -y docker-ce docker-ce-cli containerd.io \
  docker-buildx-plugin docker-compose-plugin

# 3. Vérifier
docker version
sudo systemctl status docker --no-pager

# 4. Les conteneurs redémarrent (restart policy) ; avec live-restore (section 6),
#    ils survivent même au restart du daemon
```

Bonnes pratiques :

- Lisez le changelog avant une montée de version majeure.
- Testez d'abord sur un hôte de pré-production.
- Ne laissez pas le moteur prendre 3 ans de retard (CVE sur runc/containerd).

## 65. Swarm : introduction (l'orchestrateur intégré)

Swarm transforme plusieurs hôtes Docker en **cluster** : les services sont
répartis, répliqués, et redémarrés automatiquement en cas de panne d'un nœud.

```bash
# Initialiser le cluster (sur le futur manager)
docker swarm init --advertise-addr 192.168.1.10

# Joindre un worker (commande affichée par l'init)
docker swarm join --token SWMTKN-... 192.168.1.10:2377

# Voir les nœuds
docker node ls

