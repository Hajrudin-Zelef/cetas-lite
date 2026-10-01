---
id: collect-261001-rattrapage/rattrapage/docker-guide-10
title: "Guide Docker complet — Production & Sysadmin"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["arr", "memory", "scout"]
source: docs/RAG/collect-261001-rattrapage/docker_guide.md
source_anchor: ""
source_lines: [2310, 2476]
sha256: 23c1846f2567dcbf867e91d89f2b6b2c7a547bd408f9399b929b69cf04276cbb
---

# Guide Docker complet — Production & Sysadmin

### Erreur 4 — Secrets en clair dans le YAML versionné
Symptôme : mot de passe root de la BDD dans git. Remède : secrets compose +
`*_FILE` (sections 53, 54).

### Erreur 5 — `--privileged` « pour que ça marche »
Symptôme : faille de sécurité béante. Remède : `--cap-add` ciblé, durcissement
(section 56).

### Erreur 6 — Tout en root dans les conteneurs
Symptôme : compromission d'un conteneur = root facile. Remède : `USER`
dans le Dockerfile + `user:` en compose (sections 17, 56).

### Erreur 7 — Ports publiés sur toutes les interfaces par défaut
Symptôme : `-p 8080:80` expose le service à tout le réseau alors qu'on voulait
du local. Remède : `-p 127.0.0.1:8080:80` (section 9).

### Erreur 8 — `depends_on` sans healthcheck
Symptôme : l'API démarre avant la BDD et crash en boucle. Remède :
`condition: service_healthy` + retry applicatif (section 28).

### Erreur 9 — Images jamais mises à jour
Symptôme : CVE critiques vieilles de 2 ans dans les images. Remède : rebuild
mensuel, `docker scout`/`trivy` en CI (sections 57, 62).

### Erreur 10 — Aucune limite de ressources
Symptôme : un conteneur qui fuit la RAM tue l'hôte entier (OOM killer).
Remède : `memory`/`cpus`/`pids-limit` partout (section 51).

### Erreur 11 — Le socket Docker monté n'importe où
Symptôme : un conteneur compromis = hôte compromis. Remède : ne monter le
socket que pour Portainer/Traefik/watchtower, en `:ro` quand possible
(section 55).

### Erreur 12 — Pas de sauvegarde testée des volumes
Symptôme : le jour où le disque meurt, on découvre que le « backup » ne
restaurait rien. Remède : dumps + archives + **test de restauration
trimestriel** (section 61).

## 78. Sécurité — récapitulatif du chapitre (checklist d'audit)

À passer en revue pour chaque hôte Docker de production :

- [ ] Installation depuis le dépôt officiel, moteur à jour (section 4, 64)
- [ ] Groupe `docker` réservé aux admins (section 55)
- [ ] Socket non monté dans les conteneurs applicatifs (section 55)
- [ ] Aucun `--privileged` ; capabilities minimales (section 56)
- [ ] Conteneurs non-root (`USER`/`user:`), `read_only` quand possible (section 56)
- [ ] `no-new-privileges:true` (section 56)
- [ ] Images officielles/vérifiées, tags épinglés ou digests (section 57)
- [ ] Scan de vulnérabilités (scout/trivy) avant mise en prod (section 57)
- [ ] Secrets hors git, via fichiers `*_FILE`, `chmod 600` (sections 53, 54)
- [ ] Ports publiés restreints (`127.0.0.1:` si local) (section 9)
- [ ] Réseaux dédiés par stack, `internal: true` pour les données (section 32)
- [ ] Rotation des logs activée (section 50)
- [ ] Limites de ressources sur tous les services (section 51)
- [ ] Healthchecks sur les services critiques (section 30)
- [ ] Sauvegardes des volumes + test de restauration (section 61)
- [ ] Mises à jour planifiées (images + moteur) (sections 62, 64)

## 79. Pense-bête de poche — conteneurs

```bash
docker ps -a                    # tout voir
docker logs -f --tail 50 NOM    # suivre les logs
docker exec -it NOM sh          # entrer dedans
docker stop/start/restart NOM   # cycle de vie
docker inspect NOM | less       # tout savoir
docker stats                    # ressources temps réel
docker top NOM                  # processus
docker cp NOM:/chemin ./local   # copier un fichier hors du conteneur
docker cp ./local NOM:/chemin   # ... ou vers le conteneur
docker rename ANCIEN NOUVEAU    # renommer
docker update --memory 1g NOM   # changer une limite à chaud (selon option)
```

## 80. Pense-bête de poche — images, volumes, réseaux

```bash
# Images
docker images
docker pull IMAGE:TAG
docker rmi IMAGE
docker tag SRC DST
docker history IMAGE
docker system df

# Volumes
docker volume ls
docker volume inspect NOM
docker volume prune              # ⚠️ orphelins

# Réseaux
docker network ls
docker network create --driver bridge --subnet 172.30.0.0/24 NET
docker network connect NET CONTENEUR

# Nettoyage
docker system prune              # doux
docker system prune -a --volumes # radical ⚠️
```

## 81. Pense-bête de poche — compose

```bash
docker compose up -d              # lancer
docker compose ps                 # état
docker compose logs -f SERVICE    # logs d'un service
docker compose exec SERVICE sh    # shell dans un service
docker compose run --rm SERVICE CMD  # one-shot
docker compose pull && docker compose up -d  # mettre à jour
docker compose down               # arrêter (garde les volumes)
docker compose down -v            # arrêter + SUPPRIME les volumes ⚠️
docker compose config             # config finale résolue
docker compose build SERVICE      # rebuild une image
```

## 82. Glossaire

| Terme | Définition |
|---|---|
| Bridge | Driver réseau par défaut : NAT entre conteneurs et hôte |
| BuildKit | Moteur de build moderne (parallèle, cache avancé, secrets de build) |
| cgroup | Mécanisme noyau limitant CPU/RAM par groupe de processus |
| Containerd | Daemon gérant le cycle de vie des conteneurs pour dockerd |
| Couche (layer) | Différence immuable empilée pour former une image |
| Daemon (dockerd) | Service qui construit, lance et supervise les conteneurs |
| Digest | Empreinte SHA256 immuable d'une image (`image@sha256:...`) |
| Dockerfile | Recette de construction d'une image |
| Healthcheck | Test périodique de santé d'un conteneur |
| Host (réseau) | Mode où le conteneur partage la pile réseau de l'hôte |
| Image | Modèle immuable pour créer des conteneurs |
| Macvlan | Driver donnant au conteneur une MAC/IP propre sur le LAN |
| Multi-stage | Build en plusieurs étapes pour des images finales légères |
| Namespace | Isolation noyau (PID, réseau, montage…) entre conteneurs |
| Orchestrateur | Système répartissant les conteneurs (Swarm, Kubernetes) |
| Overlay | Réseau virtuel multi-hôtes (Swarm) |
| Overlay2 | Driver de stockage moderne (union filesystem) |
| runc | Exécutable bas niveau lançant les conteneurs (OCI) |
| Registry | Entrepôt d'images (Docker Hub, Harbor, registre privé) |
| Restart policy | Comportement au crash/arrêt (`unless-stopped`…) |
| Rootless | Mode Docker sans daemon root (sécurité renforcée) |
| Secret | Donnée sensible injectée via fichier, jamais en variable claire |
| Socket Docker | `/var/run/docker.sock` : API du daemon (équivaut à root) |
| Stack | Ensemble de services déployés ensemble (compose / Swarm) |
| Tag | Étiquette de version d'image (`1.4.2`, `latest`…) |
| tmpfs | Montage en RAM, non persisté |
| Volume | Stockage persistant géré par Docker |
| Bind mount | Dossier/fichier de l'hôte monté dans un conteneur |

## 83. Quiz — 10 questions (réponses en fin de section)

1. Quelle différence entre une image et un conteneur ?
2. Pourquoi faut-il éviter le snap et le paquet `docker.io` d'Ubuntu pour
   installer Docker ?
3. Que se passe-t-il si on ne configure pas la rotation des logs ?
4. Pourquoi `depends_on` seul ne suffit-il pas à garantir qu'une API trouve
   sa base de données prête ?
5. Citez 3 raisons de préférer un volume nommé à un bind mount pour une base
   de données.
6. Que risque-t-on à monter `/var/run/docker.sock` dans un conteneur ?
7. Quelle est la différence entre `ARG` et `ENV` dans un Dockerfile ?
8. Pourquoi faut-il éviter le tag `latest` en production ?
9. Que fait `docker compose down -v` et pourquoi est-ce dangereux ?
10. Citez 4 éléments de la checklist de durcissement d'un conteneur.

**Réponses :**

