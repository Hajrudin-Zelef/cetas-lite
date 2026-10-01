---
id: collect-261001-rattrapage/rattrapage/docker-guide-9
title: "Guide Docker complet — Production & Sysadmin"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["apache", "arr", "attention", "memory"]
source: docs/RAG/collect-261001-rattrapage/docker_guide.md
source_anchor: ""
source_lines: [2095, 2309]
sha256: 53ea8c1403a9dfec3ea28360e3e50c5adc2a2556da399a4416429a1eeba6896d
---

# Guide Docker complet — Production & Sysadmin

  db:
    image: mariadb:11.4
    environment:
      MYSQL_DATABASE: nextcloud
      MYSQL_USER: nc
      MYSQL_PASSWORD_FILE: /run/secrets/db_password
      MYSQL_ROOT_PASSWORD_FILE: /run/secrets/db_root_password
    secrets: [db_password, db_root_password]
    volumes: [nc-db:/var/lib/mysql]
    command: --transaction-isolation=READ-COMMITTED --binlog-format=ROW
    healthcheck:
      test: ["CMD", "healthcheck.sh", "--connect"]
      interval: 10s
      timeout: 5s
      retries: 5
    restart: unless-stopped

  cache:
    image: redis:7-alpine
    command: ["redis-server", "--requirepass", "changeme-redis"]
    volumes: [nc-redis:/data]
    restart: unless-stopped

  cron:
    image: nextcloud:29-apache
    entrypoint: ["/cron.sh"]           # tâches de fond Nextcloud
    volumes: [nc-html:/var/www/html, ./data:/var/www/html/data]
    depends_on: [app]
    restart: unless-stopped

volumes:
  nc-html:
  nc-db:
  nc-redis:

secrets:
  db_password:      { file: ./secrets/db_password.txt }
  db_root_password: { file: ./secrets/db_root_password.txt }
  admin_user:       { file: ./secrets/admin_user.txt }
  admin_password:   { file: ./secrets/admin_password.txt }
```

Points d'attention Nextcloud :

- Mettez un **reverse proxy TLS** devant (section 67/68), Nextcloud exige
  HTTPS pour un usage sérieux.
- `trusted_domains` et `overwrite.cli.url` dans `config.php` après le premier
  démarrage.
- Sauvegarde = **les 3 volumes** (`nc-html`, `nc-db`, `./data`) + dump MariaDB.
- Montées de version : une **majeure à la fois** (29 → 30, jamais 29 → 31).

## 71. Dépannage — le conteneur crash en boucle

```bash
# 1. Voir pourquoi il meurt
docker ps -a                       # colonne STATUS : Exited (1) ...
docker logs --tail 100 mon-app     # l'erreur est souvent dans les 20 dernières lignes
docker inspect --format '{{.State.Error}}' mon-app

# 2. Cas fréquents et remèdes
```

| Symptôme dans les logs | Cause probable | Remède |
|---|---|---|
| `connection refused` vers `db` | base pas prête | `depends_on` + healthcheck + retry dans l'app |
| `permission denied` sur `/data` | UID du conteneur ≠ propriétaire du bind mount | `chown -R 1001:1001 /srv/data` ou `user:` adapté |
| `address already in use` | port hôte déjà pris | changer le port publié, `ss -tlnp` pour trouver le coupable |
| `no such file` au démarrage | mauvais `WORKDIR`/`COPY` dans le Dockerfile | vérifier les chemins, rebuild |
| `exec format error` | image pour une autre archi (ARM vs x86) | `--platform linux/amd64` ou image multi-arch |
| OOMKilled (`docker inspect` → `OOMKilled: true`) | limite mémoire trop basse | augmenter `memory`, chercher la fuite |

```bash
# Démarrer un shell malgré un CMD qui crash (pour inspecter)
docker run --rm -it --entrypoint sh mon-app:1.0
```

## 72. Dépannage — réseau : le conteneur ne joint rien

```bash
# 1. Le DNS intégré répond-il ?
docker exec app getent hosts db
docker exec app cat /etc/resolv.conf   # doit contenir 127.0.0.11

# 2. Le réseau existe-t-il, qui est dedans ?
docker network inspect mon-reseau --format '{{range .Containers}}{{.Name}} {{end}}'

# 3. Connecter/déconnecter à chaud
docker network connect mon-reseau app
docker network disconnect mon-reseau app

# 4. Tester depuis un conteneur jetable sur le même réseau
docker run --rm --network mon-reseau alpine:3.20 \
  sh -c "apk add -q curl && curl -sv http://api:8000/health" 2>&1 | tail -20

# 5. Côté hôte : iptables et IP forwarding
sudo iptables -L DOCKER-USER -n -v
sysctl net.ipv4.ip_forward   # doit valoir 1
```

Causes classiques :

- Conteneurs sur **deux réseaux différents** sans réseau commun.
- `network_mode: host` + tentative de joindre par nom de service (pas de DNS
  dans ce mode).
- Firewall de l'hôte (ufw) qui bloque le forwarding : avec ufw, réglez
  `DEFAULT_FORWARD_POLICY="ACCEPT"` dans `/etc/default/ufw`.
- Collision de sous-réseaux avec le VPN/LAN (voir `default-address-pools`,
  section 6).

## 73. Dépannage — disque plein : /var/lib/docker déborde

```bash
# Diagnostic
df -h /var/lib/docker
docker system df
sudo du -sh /var/lib/docker/* | sort -h | tail

# Les gros consommateurs de logs
sudo du -sh /var/lib/docker/containers/*/*-json.log 2>/dev/null | sort -h | tail -5

# Nettoyage progressif (du moins au plus agressif)
docker container prune -f                    # conteneurs arrêtés
docker image prune -f                        # images non utilisées (dangling)
docker image prune -a -f --filter "until=720h"  # images non utilisées depuis 30 j
docker builder prune -f                      # cache de build
docker volume prune -f                       # volumes orphelins (⚠️ vérifiez avant !)
docker system prune -a -f --volumes          # TOUT (dernier recours)
```

Prévention :

- Rotation des logs configurée (section 50) — **non négociable**.
- Cron hebdo : `docker image prune -a -f --filter "until=720h"`.
- Alerte disque à 80 % (supervision, section 58).
- Ne jamais laisser `docker system prune -a` dans un cron aveugle sur un
  hôte de prod (il peut supprimer une image encore référencée par un
  conteneur arrêté que vous comptiez relancer).

## 74. Dépannage — permissions sur les volumes (bind mounts)

Le piège : le conteneur tourne avec l'UID 1001, mais `/srv/data` appartient à
root sur l'hôte → `permission denied`.

```bash
# 1. Identifier l'UID utilisé dans le conteneur
docker exec app id
# uid=1001(appuser) gid=1001(appuser)

# 2. Aligner le propriétaire côté hôte
sudo chown -R 1001:1001 /srv/data

# 3. Ou forcer l'UID au lancement (si l'image le supporte)
docker run -d --user 1001:1001 -v /srv/data:/data mon-app:1.0
```

```yaml
# En compose
services:
  app:
    user: "1001:1001"
    volumes: ["/srv/data:/data"]
```

> Les **volumes nommés** n'ont pas ce problème : Docker les initialise avec
> les permissions du contenu de l'image. C'est une raison de plus de les
> préférer pour les données.

## 75. Dépannage — le daemon ne démarre plus

```bash
# 1. Logs du daemon
sudo journalctl -u docker --no-pager | tail -50

# 2. Cause n°1 : daemon.json invalide (JSON mal formé)
sudo python3 -m json.tool /etc/docker/daemon.json
# ou : sudo dockerd --validate  (selon version)

# 3. Cause n°2 : disque plein (voir section 73)

# 4. Cause n°3 : conflit iptables / réseau après changement de config
# Revenir à une config minimale et restart :
sudo systemctl restart docker

# 5. Vérifier le socket
ls -l /var/run/docker.sock
sudo systemctl status containerd --no-pager
```

## 76. Dépannage — méthode générale (checklist)

Face à n'importe quel problème Docker, dans l'ordre :

1. [ ] `docker ps -a` : le conteneur existe ? quel statut / code de sortie ?
2. [ ] `docker logs --tail 100 <nom>` : l'erreur est dans les logs 9 fois sur 10
3. [ ] `docker inspect <nom>` : mounts, networks, env, health, OOMKilled
4. [ ] `docker stats` : saturation CPU/RAM ?
5. [ ] `docker system df` + `df -h` : disque plein ?
6. [ ] Reproduire en minimal : `docker run --rm -it image sh`
7. [ ] `docker compose config` : le YAML interpolé est-il celui attendu ?
8. [ ] `journalctl -u docker` : le daemon lui-même se plaint-il ?

## 77. Les 12 erreurs classiques (et comment les éviter)

### Erreur 1 — `latest` en production
Symptôme : un `pull` + `up -d` change de version majeure sans prévenir, tout
casse. Remède : tags versionnés épinglés (sections 22, 57).

### Erreur 2 — Pas de rotation de logs
Symptôme : `/var/lib/docker` plein à 100 %, daemon qui ne répond plus.
Remède : `max-size`/`max-file` dans `daemon.json` (sections 6, 50).

### Erreur 3 — Données dans le conteneur au lieu d'un volume
Symptôme : `docker compose down -v` ou recréation → données perdues.
Remède : volume nommé pour **toutes** les données (sections 45, 61).

