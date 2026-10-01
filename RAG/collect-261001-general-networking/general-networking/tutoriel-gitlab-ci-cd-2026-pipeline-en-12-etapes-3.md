---
id: collect-261001-general-networking/general-networking/tutoriel-gitlab-ci-cd-2026-pipeline-en-12-etapes-3
title: "git version 2.47.1"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/tutoriel-gitlab-ci-cd-2026-pipeline-en-12-etapes.md
source_anchor: ""
source_lines: [248, 420]
sha256: bb1e177c28a5d96f40aa1375b84f18aa79a43a1fcec8836882ef4e34d10143b8
---

# git version 2.47.1

```
# Ajouter le dépôt officiel GitLab
curl -L "https://packages.gitlab.com/install/repositories/runner/gitlab-runner/script.deb.sh" | sudo bash
# Installer le runner
sudo apt-get install gitlab-runner
# Vérifier l'installation
gitlab-runner --version
# Version: 18.10.0
# Git revision: abc1234d
# Git branch: 18-10-stable
# Enregistrer le runner auprès de votre instance GitLab
sudo gitlab-runner register \
  --url https://gitlab.com/ \
  --token glrt-VOTRE_TOKEN_ICI \
  --executor docker \
  --docker-image alpine:3.20 \
  --description "runner-docker-prod" \
  --tag-list "docker,linux,production"
# Démarrer le service
sudo gitlab-runner start
# Vérifier le statut
sudo gitlab-runner status
# gitlab-runner: Service is running
```
Le token d’enregistrement se trouve dans **Settings > CI/CD > Runners** de votre projet ou groupe. Depuis GitLab 17+, les tokens de type `glrt-` remplacent les anciens registration tokens pour une meilleure sécurité.

Pour optimiser les performances, configurez le fichier `/etc/gitlab-runner/config.toml` :

```
# /etc/gitlab-runner/config.toml
concurrent = 4
check_interval = 3
[[runners]]
  name = "runner-docker-prod"
  url = "https://gitlab.com/"
  token = "glrt-VOTRE_TOKEN"
  executor = "docker"
  [runners.docker]
    image = "alpine:3.20"
    privileged = false
    disable_entrypoint_overwrite = false
    oom_kill_disable = false
    volumes = ["/cache", "/var/run/docker.sock:/var/run/docker.sock"]
    shm_size = 0
    pull_policy = ["if-not-present"]
  [runners.cache]
    Type = "s3"
    Shared = true
    [runners.cache.s3]
      BucketName = "gitlab-runner-cache"
      BucketLocation = "eu-west-3"
```
Le paramètre `concurrent = 4` permet d’exécuter jusqu’à 4 jobs simultanément. La directive `pull_policy = ["if-not-present"]` évite de re-télécharger les images Docker à chaque job, ce qui réduit la durée d’exécution de 30 à 50 % selon la taille des images.

## Étape 6 : Optimiser le Cache et les Artefacts

La gestion du cache est l’un des leviers les plus efficaces pour accélérer vos pipelines GitLab CI/CD. Sans cache, chaque job télécharge et installe les dépendances depuis zéro, ce qui peut ajouter 2 à 5 minutes par pipeline pour un projet Node.js typique.

GitLab distingue deux mécanismes : le **cache** (persisté entre pipelines, pour les dépendances) et les **artefacts** (transmis entre stages du même pipeline, pour les résultats de build). Confondre les deux est une erreur fréquente.

| Caractéristique | Cache | Artefacts | 
|---|---|---|
| Persistance | Entre pipelines (best-effort) | Dans le même pipeline (garanti) | 
| Cas d’usage | node_modules/, pip cache, .m2/ | Binaires compilés, rapports de test | 
| Disponibilité | Même branche par défaut | Tous les jobs des stages suivants | 
| Politique | pull, push, pull-push | always, on_success, on_failure | 
| Stockage | Runner local ou S3 | Serveur GitLab | 
| Expiration | Configurable | expire_in (défaut : 30 jours) | 

Voici une configuration optimisée du cache pour un projet Node.js :

```
# Configuration cache optimisée
cache:
  key:
    files:
      - package-lock.json
  paths:
    - node_modules/
  policy: pull-push
# Job qui ne fait que lire le cache (plus rapide)
lint:
  stage: validate
  cache:
    key:
      files:
        - package-lock.json
    paths:
      - node_modules/
    policy: pull  # Ne met pas à jour le cache
  script:
    - npx eslint src/
# Artefacts pour transmettre le build au stage deploy
build:
  stage: build
  script:
    - npm run build
  artifacts:
    paths:
      - dist/
    expire_in: 1 week
```
L’utilisation de `package-lock.json` comme clé de cache garantit que le cache est invalidé uniquement lorsque les dépendances changent. Les jobs en lecture seule utilisent `policy: pull` pour éviter les écritures inutiles. Pour les projets Python, remplacez par `requirements.txt` et le chemin `.pip-cache/`.

Un piège courant est de mettre `node_modules/` dans les artefacts au lieu du cache. Cela transfère des centaines de Mo via le serveur GitLab à chaque pipeline, ralentissant tout. Les dépendances vont dans le cache, les résultats de build dans les artefacts.

## Étape 7 : Intégrer Docker et le Container Registry GitLab

GitLab inclut un registre Docker intégré pour chaque projet, accessible via `$CI_REGISTRY_IMAGE`. Cette intégration élimine le besoin d’un Docker Hub ou d’un registre externe pour la plupart des cas d’usage.

Commencez par créer un `Dockerfile` optimisé pour la production :

```
# Dockerfile multi-stage optimisé
FROM node:22-alpine AS builder
WORKDIR /app
COPY package*.json ./
RUN npm ci --only=production
COPY src/ ./src/
FROM node:22-alpine AS production
WORKDIR /app
RUN addgroup -g 1001 appgroup && \
    adduser -u 1001 -G appgroup -s /bin/sh -D appuser
COPY --from=builder /app/node_modules ./node_modules
COPY --from=builder /app/src ./src
COPY package.json ./
USER appuser
EXPOSE 3000
HEALTHCHECK --interval=30s --timeout=3s \
  CMD wget -qO- http://localhost:3000/health || exit 1
CMD ["node", "src/app.js"]
```
Le build multi-stage réduit la taille de l’image finale en excluant les dépendances de développement. L’utilisateur non-root `appuser` renforce la sécurité du conteneur. Le `HEALTHCHECK` permet à Docker Compose et aux orchestrateurs de vérifier l’état de l’application.

Pour builder et pusher l’image dans le pipeline GitLab CI/CD, deux approches existent : Docker-in-Docker (DinD) et le montage du socket Docker. DinD est plus sécurisé mais légèrement plus lent :

```
# Approche Docker-in-Docker (recommandée pour SaaS)
docker-build:
  stage: build
  image: docker:27
  services:
    - docker:27-dind
  variables:
    DOCKER_TLS_CERTDIR: "/certs"
  before_script:
    - docker login -u $CI_REGISTRY_USER -p $CI_REGISTRY_PASSWORD $CI_REGISTRY
  script:
    - docker build
      --cache-from $CI_REGISTRY_IMAGE:latest
      --tag $CI_REGISTRY_IMAGE:$CI_COMMIT_SHORT_SHA
      --tag $CI_REGISTRY_IMAGE:latest
      .
    - docker push $CI_REGISTRY_IMAGE:$CI_COMMIT_SHORT_SHA
    - docker push $CI_REGISTRY_IMAGE:latest
  rules:
    - if: $CI_COMMIT_BRANCH == "main"
# Sortie attendue :
# Step 1/12 : FROM node:22-alpine AS builder
# ---> Using cache
# Successfully built a1b2c3d4e5f6
# Successfully tagged registry.gitlab.com/groupe/demo-cicd-2026:abc12345
# The push refers to repository [registry.gitlab.com/groupe/demo-cicd-2026]
# abc12345: digest: sha256:... size: 1234
```
La directive `--cache-from` réutilise les couches de l’image précédente, ce qui peut réduire le temps de build de 60 à 80 % lorsque seul le code applicatif change. Le double tag (SHA + latest) permet de déployer une version précise tout en gardant une référence à la dernière version stable.

## Étape 8 : Mettre en Place les Tests Automatisés et la Couverture

Un pipeline GitLab CI/CD robuste inclut plusieurs niveaux de tests. GitLab offre une intégration native avec les rapports JUnit, la couverture de code et les merge request widgets qui affichent les résultats directement dans l’interface.

Configurez des jobs séparés pour chaque type de test :

