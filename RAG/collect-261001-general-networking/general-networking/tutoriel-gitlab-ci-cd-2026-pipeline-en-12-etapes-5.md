---
id: collect-261001-general-networking/general-networking/tutoriel-gitlab-ci-cd-2026-pipeline-en-12-etapes-5
title: "git version 2.47.1"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["arr", "mai"]
source: docs/RAG/collect-261001-general-networking/tutoriel-gitlab-ci-cd-2026-pipeline-en-12-etapes.md
source_anchor: ""
source_lines: [606, 745]
sha256: 313854cc9c1a1c7fea354bfc3ab8c3d87ee4dd844b7e4d1902291f7c0260c5b9
---

# git version 2.47.1

## Étape 11 : Pipelines Avancés avec Rules, Needs et DAG

Les pipelines GitLab CI/CD avancés utilisent les DAG (Directed Acyclic Graphs) via le mot-clé `needs` pour paralléliser les stages et réduire drastiquement la durée d'exécution.

Sans `needs`, un pipeline à 4 stages s'exécute strictement en séquence : validate (1 min) → test (3 min) → build (2 min) → deploy (1 min) = 7 minutes. Avec `needs`, les jobs peuvent démarrer dès que leurs dépendances directes sont terminées, réduisant le chemin critique.

```
# Pipeline DAG optimisé
stages:
  - validate
  - test
  - build
  - deploy
lint:
  stage: validate
  script: npx eslint src/
security-scan:
  stage: validate
  script: echo "SAST scan"
unit-tests:
  stage: test
  needs: ["lint"]    # Démarre dès que lint est fini
  script: npm test
integration-tests:
  stage: test
  needs: ["lint"]    # Parallèle avec unit-tests
  script: npm run test:integration
docker-build:
  stage: build
  needs: ["unit-tests"]  # N'attend pas integration-tests
  script: docker build -t $DOCKER_IMAGE .
deploy-staging:
  stage: deploy
  needs: ["docker-build", "integration-tests", "security-scan"]
  script: ./deploy.sh staging
# Résultat : le pipeline passe de 7 min à ~4 min
# lint (1 min) → unit-tests + integration-tests (3 min, parallèle)
#                → docker-build (2 min, commence dès unit-tests fini)
#                → deploy (1 min, attend tous les prérequis)
```
Les **rules** avancées permettent de contrôler finement quand les jobs s'exécutent. Voici les patterns les plus utiles :

```
# Exécuter uniquement si des fichiers spécifiques changent
test-frontend:
  rules:
    - if: $CI_PIPELINE_SOURCE == "merge_request_event"
      changes:
        - "src/frontend/**/*"
        - "package.json"
# Skip CI avec le message de commit
all-jobs:
  rules:
    - if: $CI_COMMIT_MESSAGE =~ /\[skip ci\]/
      when: never
    - when: on_success
# Pipelines programmés (cron)
nightly-tests:
  rules:
    - if: $CI_PIPELINE_SOURCE == "schedule"
  script: npm run test:e2e
# Déclenchement conditionnel par variable
deploy:
  rules:
    - if: $CI_COMMIT_TAG =~ /^v\d+\.\d+\.\d+$/
      variables:
        DEPLOY_ENV: "production"
    - if: $CI_COMMIT_BRANCH == "main"
      variables:
        DEPLOY_ENV: "staging"
```
Le déclenchement par `changes` est particulièrement utile pour les monorepos : les tests frontend ne s'exécutent que si le code frontend a changé, évitant des minutes CI/CD inutiles. Les pipelines programmés permettent de lancer des suites de tests e2e complètes la nuit sans bloquer les développeurs pendant la journée.

## Étape 12 : Utiliser les CI/CD Components et les Templates Réutilisables

GitLab 18.x a introduit les **CI/CD Components**, un système de configuration modulaire qui remplace les anciens fichiers `include`. Les components sont versionnés, testés et publiés dans un catalogue centralisé, et depuis GitLab 18.9 (mai 2026), une vue d'analytics du catalogue affiche le nombre d'utilisations par projet sur une fenêtre glissante de 30 jours — le détail par composant individuel restant toutefois réservé au tier Ultimate. Depuis août 2026, cette Component Analytics repère en plus les versions de components devenues obsolètes à travers l'ensemble des projets, facilitant les campagnes de mise à jour groupées.

Pour utiliser un component depuis le catalogue GitLab :

```
# Utiliser des components du catalogue
include:
  # Component officiel de déploiement
  - component: gitlab.com/components/[email protected]
    inputs:
      environment: staging
      strategy: rolling
  # Component de notification Slack
  - component: gitlab.com/components/[email protected]
    inputs:
      channel: "#deployments"
      webhook_url: $SLACK_WEBHOOK
  # Template local réutilisable
  - local: .gitlab/ci/templates/node-defaults.yml
# Créer votre propre template réutilisable
# .gitlab/ci/templates/node-defaults.yml
.node-defaults:
  image: node:22-alpine
  before_script:
    - npm ci
  cache:
    key:
      files:
        - package-lock.json
    paths:
      - node_modules/
# Héritage du template
lint:
  extends: .node-defaults
  stage: validate
  script:
    - npx eslint src/
tests:
  extends: .node-defaults
  stage: test
  script:
    - npm test
```
Le mot-clé `extends` permet l'héritage de configuration, éliminant la duplication. Un template `.node-defaults` centralise l'image, le `before_script` et le cache. Chaque job qui en hérite n'a plus qu'à définir son stage et son script spécifique.

Pour les organisations avec plusieurs projets, créez un dépôt de templates partagés au niveau du groupe GitLab. Les projets enfants incluent ces templates via une URL relative, garantissant la cohérence des pipelines à travers l'organisation.

## 5 Pièges Courants et Comment les Éviter

Après des centaines de pipelines GitLab CI/CD — et alors que l'usage des pipelines CI/CD a bondi de 40 % en glissement annuel à fin septembre 2026 selon les points saillants de la conférence de résultats relayés par MarketBeat, tandis que le nombre de clients GitLab générant plus de 100 000 $ d'ARR est passé de 1 344 à 1 571 (+17 %) d'après le dernier 10-Q déposé auprès de la SEC —, certains pièges reviennent systématiquement. Voici les erreurs les plus fréquentes et leurs solutions.

**Piège 1 : Le cache qui ne fonctionne pas entre les branches.** Par défaut, la clé de cache utilise `$CI_COMMIT_REF_SLUG`, ce qui crée un cache par branche. Les branches de feature partent donc toujours d'un cache vide. Solution : utilisez `key: files: [package-lock.json]` pour partager le cache entre toutes les branches ayant les mêmes dépendances.

**Piège 2 : Docker-in-Docker qui échoue avec "Cannot connect to the Docker daemon".** Ce problème survient quand `DOCKER_TLS_CERTDIR` n'est pas configuré ou quand le service DinD n'a pas eu le temps de démarrer. Solution : ajoutez `DOCKER_TLS_CERTDIR: "/certs"` dans les variables et utilisez la même version majeure pour l'image et le service (ex. `docker:27` et `docker:27-dind`).

**Piège 3 : Les jobs qui tournent alors qu'ils ne devraient pas.** La syntaxe `only/except` est dépréciée et produit des comportements inattendus combinée avec `rules`. Solution : utilisez exclusivement `rules` avec des conditions explicites. Ne mélangez jamais `only/except` et `rules` dans le même pipeline.

**Piège 4 : Les artefacts qui font exploser le stockage GitLab.** Sans `expire_in`, les artefacts sont conservés indéfiniment. Un projet actif avec des rapports de couverture de 50 Mo peut consommer des Go de stockage en quelques semaines. Solution : définissez toujours `expire_in` (7 jours pour les rapports de test, 30 jours pour les builds).

**Piège 5 : Les variables masquées qui apparaissent quand même dans les logs.** Le flag `Masked` ne fonctionne que si la valeur fait au moins 8 caractères et ne contient pas de caractères spéciaux. Solution : encodez les secrets complexes en base64 avant de les stocker, et décodez-les dans le script du job.

## Dépannage : 8 Erreurs Fréquentes et Leurs Solutions

Voici les erreurs les plus courantes rencontrées dans les pipelines GitLab CI/CD, avec les commandes de diagnostic et les solutions testées.

