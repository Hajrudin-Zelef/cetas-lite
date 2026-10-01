---
id: collect-261001-rattrapage/rattrapage/docker-guide-5
title: "Guide Docker complet — Production & Sysadmin"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["arr", "attention"]
source: docs/RAG/collect-261001-rattrapage/docker_guide.md
source_anchor: ""
source_lines: [1032, 1295]
sha256: c6276ccf510bb4b2a2ede9fdb9fa1866d3bc4f20bf8ab4ed903f0897ff164b7c
---

# Exécuter une commande / shell dans un service
docker compose exec api bash
docker compose exec db psql -U app -d appdb

# Lancer un conteneur one-shot (ne démarre pas les dépendances liées)
docker compose run --rm api python -m pytest

# Voir la config finale après interpolation et surcharges
docker compose config

# Mettre à jour toute la stack
docker compose pull && docker compose up -d

# Tout arrêter proprement
docker compose down
```

## 37. Registres : Docker Hub

```bash
# Connexion (stockée chiffrée dans ~/.docker/config.json)
docker login
docker login registry.interne:5000

# Pousser une image (le nom doit commencer par votre namespace)
docker tag mon-app:1.4.2 zelef/mon-app:1.4.2
docker push zelef/mon-app:1.4.2

# Se déconnecter
docker logout
```

- Docker Hub : gratuit pour les images publiques, quotas de pull pour les
  anonymes (limites périodiques) → en prod, **authentifiez** le daemon
  (`docker login`) ou utilisez un miroir/registre interne.
- Préférez les images **officielles** (`nginx`, `postgres`…) et les éditeurs
  **vérifiés** (badge). Méfiez-vous des images au nom approximatif
  (typosquatting : `nglnx` au lieu de `nginx`).

## 38. Registre privé : le lancer soi-même

Le registre officiel est lui-même un conteneur :

```yaml
# compose.yaml du registre privé
services:
  registry:
    image: registry:3
    container_name: registry
    ports:
      - "127.0.0.1:5000:5000"   # ou 5000:5000 si accès réseau voulu
    volumes:
      - registry-data:/var/lib/registry
    environment:
      REGISTRY_STORAGE_DELETE_ENABLED: "true"
    restart: unless-stopped

volumes:
  registry-data:
```

```bash
docker compose up -d
docker tag mon-app:1.4.2 localhost:5000/mon-app:1.4.2
docker push localhost:5000/mon-app:1.4.2
docker pull localhost:5000/mon-app:1.4.2
```

- En HTTPS avec certificat : montez les certs et mettez
  `REGISTRY_HTTP_TLS_CERTIFICATE` / `REGISTRY_HTTP_TLS_KEY`.
- Registre **insecure** (HTTP, labo uniquement) : à déclarer côté daemon :

```json
{ "insecure-registries": ["192.168.1.50:5000"] }
```

> En production, préférez Harbor ou Nexus (RBAC, scan de vulnérabilités,
> réplication) plutôt que le registre nu.

## 39. Réseaux Docker : vue d'ensemble

```bash
docker network ls
docker network inspect mon-reseau
```

| Driver | Usage | Isolement |
|---|---|---|
| `bridge` | défaut, conteneurs d'un même hôte | NAT vers l'hôte |
| `host` | conteneur partage la pile réseau de l'hôte | aucun (perf maximale) |
| `macvlan` | conteneur avec sa propre MAC/IP sur le LAN | apparaît comme une machine physique |
| `overlay` | multi-hôtes (Swarm) | chiffré en option |
| `none` | pas de réseau | total |
| `internal` (option) | bridge sans accès Internet | sortant bloqué |

## 40. Bridge : le réseau par défaut (et ses limites)

```bash
# Créer un réseau bridge dédié (TOUJOURS préféré au bridge par défaut "bridge")
docker network create --driver bridge --subnet 172.30.0.0/24 mon-bridge

docker run -d --name web --network mon-bridge nginx:1.27-alpine
docker run -d --name api --network mon-bridge mon-api:1.0

# Depuis api, "web" se résout en DNS :
docker exec api ping -c 2 web
docker exec api getent hosts web
```

Pourquoi éviter le bridge par défaut `bridge` :

- pas de **résolution DNS par nom** entre conteneurs (il faut des `--link`,
  obsolètes) ;
- tous les conteneurs « par défaut » partagent le même segment.

> Règle : **un réseau bridge dédié par application/stack**.

## 41. Host : quand la performance réseau prime

```bash
docker run -d --name dns --network host coredns/coredns:1.11.1
```

- Le conteneur utilise directement les interfaces de l'hôte : **zéro NAT**,
  latence minimale.
- Inconvénients : pas d'isolation des ports (conflits possibles), pas de
  mapping `-p`, surface d'attaque élargie.
- Cas d'usage : serveurs DNS, exporters réseau, applis très sensibles à la
  latence. À éviter pour une appli web standard.

## 42. Macvlan : le conteneur sur le LAN physique

Le conteneur obtient sa **propre adresse MAC et IP** sur votre réseau local :
il apparaît comme une machine physique (utile pour DHCP, supervision SNMP,
vieux protocoles qui n'aiment pas le NAT).

```bash
# Créer le réseau macvlan (adaptez l'interface et le sous-réseau !)
docker network create -d macvlan \
  --subnet=192.168.1.0/24 \
  --gateway=192.168.1.1 \
  -o parent=eth0 \
  lan-net

docker run -d --name srv --network lan-net --ip 192.168.1.60 nginx:1.27-alpine
```

Points d'attention :

- L'hôte **ne peut pas joindre** directement ses conteneurs macvlan (limite du
  noyau) : créez une interface macvlan sur l'hôte si besoin.
- Réservez les IP dans le DHCP du LAN pour éviter les collisions.
- Certains hyperviseurs / cartes réseau filtrent les MAC multiples (mode
  promiscuous requis sur l'interface).

## 43. Overlay : le multi-hôtes (Swarm)

- Disponible en mode **Swarm** (voir section 66) : un réseau logique réparti
  sur plusieurs machines.
- Chiffrement du trafic inter-nœuds : `--opt encrypted`.
- En mono-hôte avec compose, inutile : restez sur `bridge`.

```bash
# Exemple (manager Swarm)
docker network create -d overlay --opt encrypted mon-overlay
```

## 44. DNS intégré et communication inter-conteneurs

- Sur tout réseau **dédié** (bridge custom, overlay), Docker fournit un
  serveur DNS embarqué (`127.0.0.11` dans le conteneur).
- Chaque conteneur est joignable par son **nom** ou **nom de service**
  (compose), plus les `aliases` éventuels.
- `docker exec app getent hosts db` : vérification rapide.

```bash
# Tester la connectivité entre deux conteneurs
docker run --rm --network mon-bridge alpine:3.20 \
  sh -c "apk add -q curl && curl -s -o /dev/null -w '%{http_code}' http://web:80"
```

Personnaliser le DNS d'un conteneur :

```yaml
services:
  app:
    dns:
      - 9.9.9.9
      - 1.1.1.1
    dns_search:
      - entreprise.lan
```

## 45. Volumes : le stockage persistant géré par Docker

```bash
docker volume create db-data
docker volume ls
docker volume inspect db-data
docker run -d --name db -v db-data:/var/lib/postgresql/data postgres:16-alpine

# Où sont les données sur l'hôte ?
sudo ls /var/lib/docker/volumes/db-data/_data
```

- Les volumes survivent à la suppression du conteneur (`docker rm`) et même
  à `docker compose down` (seul `down -v` les supprime).
- Sauvegarde : voir section 61 — **c'est LE point critique** en production.

Volume sur stockage réseau (NFS) :

```yaml
volumes:
  nfs-data:
    driver: local
    driver_opts:
      type: nfs
      o: "addr=192.168.1.10,rw,nfsvers=4"
      device: ":/exports/docker"
```

## 46. Bind mounts : quand monter un dossier de l'hôte

```bash
# Dossier de config en lecture seule
docker run -d --name web \
  -v /srv/web/nginx.conf:/etc/nginx/nginx.conf:ro \
  -v /srv/web/html:/usr/share/nginx/html:ro \
  nginx:1.27-alpine
```

| | Volume nommé | Bind mount |
|---|---|---|
| Géré par | Docker | vous |
| Chemin hôte | `/var/lib/docker/volumes/...` | libre (`/srv/...`) |
| Sauvegarde | via conteneur helper (section 61) | rsync/cp direct |
| Permissions | Docker initialise avec le contenu de l'image | **UID/GID de l'hôte** (piège classique, section 70) |
| Cas d'usage | bases de données, données applicatives | configs, code en dev |

## 47. tmpfs : le stockage en RAM

```yaml
services:
  app:
    tmpfs:
      - /run
      - /tmp:size=100m,mode=1777
```

- Contenu **jamais écrit sur disque**, perdu à l'arrêt du conteneur.
- Idéal pour caches, sockets temporaires, et pour rendre un conteneur en
  lecture seule (`read_only: true`, voir section 53) tout en gardant des
  zones inscriptibles.

## 48. Partager des données entre conteneurs

```bash
# volumes_from : héritage (ancien, encore vu dans la nature)
docker run -d --name data-holder -v /data busybox true
docker run --volumes-from data-holder alpine ls /data

# Moderne : même volume nommé monté dans deux services
```

