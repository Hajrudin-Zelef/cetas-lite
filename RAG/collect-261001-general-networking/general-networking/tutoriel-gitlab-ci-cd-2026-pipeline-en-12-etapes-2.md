---
id: collect-261001-general-networking/general-networking/tutoriel-gitlab-ci-cd-2026-pipeline-en-12-etapes-2
title: "git version 2.47.1"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["agent", "mai"]
source: docs/RAG/collect-261001-general-networking/tutoriel-gitlab-ci-cd-2026-pipeline-en-12-etapes.md
source_anchor: ""
source_lines: [111, 247]
sha256: 5316a1bbde4704eda83f219bb29f008aa56f4d7dd4b23c52e62804aaffc1b420
---

# git version 2.47.1

| Concept | Description | Exemple | 
|---|---|---|
| Pipeline | Ensemble ordonné de stages déclenché par un événement git | Push sur main, merge request | 
| Stage | Phase séquentielle contenant un ou plusieurs jobs | build, test, deploy | 
| Job | Unité d’exécution avec un script et une image Docker | unit-test, lint, sast-scan | 
| Runner | Agent qui exécute les jobs (partagé SaaS ou self-hosted) | Docker executor, Kubernetes executor | 
| Artefact | Fichier produit par un job, transmissible au stage suivant | rapport de couverture, binaire compilé | 
| Cache | Fichiers persistés entre pipelines pour accélérer les builds | node_modules/, .m2/repository | 
| Variable CI/CD | Valeur injectable dans les jobs (secrète ou publique) | $DOCKER_PASSWORD, $APP_ENV | 

Les runners partagés de GitLab SaaS utilisent des machines virtuelles éphémères sous Linux (architecture x86_64) avec Docker executor. Chaque job démarre dans un conteneur propre, ce qui garantit l’isolation. Pour les projets nécessitant plus de contrôle, vous pouvez enregistrer vos propres runners sur des serveurs dédiés ou dans un cluster Kubernetes.

Depuis que les **CI/CD Components** — accompagnés de leurs inputs typés — et leur catalogue sont passés en disponibilité générale avec GitLab 17.0 (sorti en mai 2025) — un statut toujours confirmé par les notes de version officielles GitLab actualisées en septembre 2026 —, la plateforme supporte pleinement ces blocs de configuration réutilisables que vous pouvez importer depuis le catalogue GitLab. Cela remplace les anciens templates `include` et offre un versionnement sémantique des configurations partagées, dont le processus de publication versionnée est détaillé dans la documentation mise à jour en août 2026.

## Étape 3 : Écrire Votre Premier Fichier .gitlab-ci.yml

Le fichier `.gitlab-ci.yml` placé à la racine du dépôt est le cœur de votre pipeline GitLab CI/CD. Voici un premier pipeline fonctionnel qui couvre le lint, les tests unitaires et le build Docker :

```
# .gitlab-ci.yml
stages:
  - validate
  - test
  - build
  - deploy
variables:
  NODE_ENV: "test"
  DOCKER_IMAGE: "$CI_REGISTRY_IMAGE:$CI_COMMIT_SHORT_SHA"
# Cache global pour node_modules
cache:
  key: ${CI_COMMIT_REF_SLUG}
  paths:
    - node_modules/
# Stage: validate
lint:
  stage: validate
  image: node:22-alpine
  before_script:
    - npm ci
  script:
    - npx eslint src/ tests/ --max-warnings 0
  rules:
    - if: $CI_PIPELINE_SOURCE == "merge_request_event"
    - if: $CI_COMMIT_BRANCH == "main"
# Stage: test
unit-tests:
  stage: test
  image: node:22-alpine
  before_script:
    - npm ci
  script:
    - npm test -- --ci --reporters=default --reporters=jest-junit
  coverage: '/All files[^|]*\|[^|]*\s+([\d\.]+)/'
  artifacts:
    when: always
    reports:
      junit: junit.xml
      coverage_report:
        coverage_format: cobertura
        path: coverage/cobertura-coverage.xml
    paths:
      - coverage/
    expire_in: 30 days
# Stage: build
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
    - docker build -t $DOCKER_IMAGE .
    - docker push $DOCKER_IMAGE
  rules:
    - if: $CI_COMMIT_BRANCH == "main"
# Stage: deploy (staging)
deploy-staging:
  stage: deploy
  image: alpine:3.20
  before_script:
    - apk add --no-cache openssh-client
    - eval $(ssh-agent -s)
    - echo "$SSH_PRIVATE_KEY" | ssh-add -
  script:
    - ssh -o StrictHostKeyChecking=no $DEPLOY_USER@$STAGING_HOST
      "docker pull $DOCKER_IMAGE && docker compose up -d"
  environment:
    name: staging
    url: https://staging.example.com
  rules:
    - if: $CI_COMMIT_BRANCH == "main"
```
Ce pipeline définit quatre stages exécutés séquentiellement : `validate` vérifie le style du code, `test` lance les tests unitaires avec couverture, `build` crée et pousse l’image Docker vers le registre GitLab intégré, et `deploy` met à jour l’environnement de staging via SSH.

Les directives `rules` contrôlent quand chaque job s’exécute. Ici, le lint tourne sur les merge requests et les pushs sur main, tandis que le build et le déploiement ne se déclenchent que sur main. Cette approche évite de gaspiller des minutes CI/CD sur des branches de feature pour les étapes coûteuses.

Poussez ce fichier et observez le pipeline se déclencher automatiquement dans l’onglet **CI/CD > Pipelines** de votre projet GitLab.

## Étape 4 : Configurer les Variables CI/CD et les Secrets

Les variables CI/CD permettent d’injecter des valeurs sensibles (clés API, mots de passe, tokens) dans vos pipelines sans les stocker dans le code source. GitLab propose trois niveaux de portée : instance (admin), groupe et projet — et depuis mars 2025, les nouvelles variables d’environnement CI du monolithe sont préfixées `GLCI_` afin de les distinguer clairement des variables historiques. Depuis la version 18.0 (mai 2025), GitLab a également introduit des permissions de job token à granularité fine, renforçant le contrôle des accès inter-projets déclenchés par les pipelines, et la version 17.11 (février 2026) a ajouté les inputs CI/CD structurés pour sécuriser davantage le déclenchement des pipelines — une évolution que Releasebot a détaillée dans sa mise à jour du 5 mai 2026 (version 2026.05), soulignant l’apport de paramètres typés et de règles de validation pour ces inputs.

Pour ajouter une variable, naviguez vers **Settings > CI/CD > Variables** dans votre projet. Cliquez sur **Add variable** et configurez :

**Key** : le nom de la variable (ex. `SSH_PRIVATE_KEY`). **Value** : la valeur secrète. **Type** : Variable (texte) ou File (fichier temporaire). **Flags** : cochez *Protected* pour limiter l’accès aux branches protégées, et *Masked* pour masquer la valeur dans les logs.

Depuis GitLab 18.9, vous pouvez également utiliser les **job inputs** pour les jobs manuels. Cette fonctionnalité permet de passer des paramètres dynamiques à un job sans relancer tout le pipeline. Par exemple, un job de déploiement peut accepter un numéro de version en entrée :

```
# Exemple de job avec inputs (GitLab 18.9+)
deploy-production:
  stage: deploy
  image: alpine:3.20
  script:
    - echo "Déploiement version $DEPLOY_VERSION sur production"
    - ./deploy.sh --version $DEPLOY_VERSION --env production
  environment:
    name: production
    url: https://app.example.com
  rules:
    - if: $CI_COMMIT_TAG
      when: manual
  allow_failure: false
```
GitLab fournit aussi des **variables prédéfinies** accessibles dans tous les jobs. Les plus utiles sont `$CI_COMMIT_SHA` (hash complet du commit), `$CI_COMMIT_SHORT_SHA` (8 premiers caractères), `$CI_PIPELINE_ID`, `$CI_REGISTRY` (URL du registre Docker intégré), `$CI_PROJECT_DIR` (chemin du code cloné) et `$CI_MERGE_REQUEST_IID` (numéro de la merge request).

Pour les déploiements multi-environnements, utilisez les **variables d’environnement scopées**. Créez une variable `DATABASE_URL` avec des valeurs différentes pour les environnements staging et production. GitLab injectera automatiquement la bonne valeur selon l’environnement cible du job.

## Étape 5 : Installer et Configurer un GitLab Runner Self-Hosted

Les runners partagés de GitLab SaaS conviennent pour démarrer, mais les projets professionnels bénéficient d’un runner self-hosted pour des raisons de performance, de sécurité et de coût — une alternative gérée existe désormais pour les clients GitLab Dedicated, dont les runners CI/CD hébergés sont disponibles en accès limité depuis la version 17.8, sortie début 2025. Pour un runner self-hosted classique, l’installation prend moins de 10 minutes sur une machine Linux.

Installez le GitLab Runner sur votre serveur (Ubuntu/Debian) :

