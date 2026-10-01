---
id: collect-261001-general-networking/general-networking/tutoriel-gitlab-ci-cd-2026-pipeline-en-12-etapes-4
title: "git version 2.47.1"
domain: general-networking
role: reference
task: reference
actors: ["AWS"]
dates: []
keywords: ["agent", "arr", "aws"]
source: docs/RAG/collect-261001-general-networking/tutoriel-gitlab-ci-cd-2026-pipeline-en-12-etapes.md
source_anchor: ""
source_lines: [421, 605]
sha256: 2ee49b11677e63e0b934da85719e19970af0925971503ae4133cc0936541add5
---

# git version 2.47.1

```
# Tests unitaires avec couverture
unit-tests:
  stage: test
  image: node:22-alpine
  before_script:
    - npm ci
  script:
    - npm test -- --ci --coverage --reporters=default --reporters=jest-junit
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
    expire_in: 7 days
# Tests d'intégration
integration-tests:
  stage: test
  image: node:22-alpine
  services:
    - name: postgres:16-alpine
      alias: db
  variables:
    POSTGRES_DB: test_db
    POSTGRES_USER: test_user
    POSTGRES_PASSWORD: test_pass
    DATABASE_URL: "postgresql://test_user:test_pass@db:5432/test_db"
  before_script:
    - npm ci
  script:
    - npm run test:integration
  rules:
    - if: $CI_COMMIT_BRANCH == "main"
    - if: $CI_PIPELINE_SOURCE == "merge_request_event"
# Sortie attendue des tests :
# PASS tests/app.test.js
#   API Health Check
#     ✓ GET /health retourne status ok (23 ms)
#     ✓ GET /api/users retourne la liste (8 ms)
#
# ----------|---------|----------|---------|---------|
# File      | % Stmts | % Branch | % Funcs | % Lines |
# ----------|---------|----------|---------|---------|
# All files |   94.12 |      100 |     100 |   93.33 |
#  app.js   |   94.12 |      100 |     100 |   93.33 |
# ----------|---------|----------|---------|---------|
```
L’expression régulière dans `coverage` permet à GitLab d’extraire le pourcentage de couverture depuis la sortie Jest. Ce chiffre apparaît ensuite dans les badges du projet et dans les merge requests. Le rapport Cobertura active la visualisation de la couverture ligne par ligne directement dans le diff de la merge request.

Les **services** dans le job d’intégration lancent un conteneur PostgreSQL éphémère accessible via l’alias `db`. Ce mécanisme est identique pour Redis, MongoDB ou tout autre service nécessaire aux tests. GitLab gère le cycle de vie des conteneurs de service automatiquement.

Pour les projets exigeant un seuil de couverture minimal, ajoutez une vérification dans le script :

```
unit-tests:
  script:
    - npm test -- --ci --coverage
    - |
      COVERAGE=$(cat coverage/coverage-summary.json | 
        python3 -c "import sys,json; print(json.load(sys.stdin)['total']['lines']['pct'])")
      echo "Couverture: ${COVERAGE}%"
      if [ $(echo "$COVERAGE < 80" | bc -l) -eq 1 ]; then
        echo "ERREUR: Couverture insuffisante (${COVERAGE}% < 80%)"
        exit 1
      fi
```
## Étape 9 : Ajouter le Scanning de Sécurité (SAST et Dependency Scanning)

L'un des avantages majeurs de GitLab CI/CD par rapport à GitHub Actions est l'intégration native des outils de sécurité. Le SAST (Static Application Security Testing) analyse le code source, tandis que le Dependency Scanning vérifie les vulnérabilités connues dans les bibliothèques tierces.

Pour activer le scanning de sécurité, ajoutez les templates officiels GitLab :

```
# Ajouter en haut du .gitlab-ci.yml
include:
  - template: Security/SAST.gitlab-ci.yml
  - template: Security/Dependency-Scanning.gitlab-ci.yml
  - template: Security/Secret-Detection.gitlab-ci.yml
# Les jobs suivants sont automatiquement ajoutés :
# - semgrep-sast (analyse statique du code)
# - nodejs-scan-sast (scan spécifique Node.js)
# - gemnasium-dependency_scanning (vulnérabilités NPM)
# - secret_detection (clés API, tokens dans le code)
# Personnalisation optionnelle
variables:
  SAST_EXCLUDED_PATHS: "tests/, node_modules/, coverage/"
  DS_EXCLUDED_ANALYZERS: "bundler-audit"
  SECRET_DETECTION_EXCLUDED_PATHS: "tests/fixtures/"
```
Ces templates ajoutent automatiquement des jobs au stage `test` de votre pipeline. Les résultats apparaissent dans l'onglet **Security** du projet et directement dans les merge requests sous forme de widget. Depuis GitLab 18.9, l'auto-remédiation propose automatiquement des merge requests pour corriger les vulnérabilités détectées, avec des niveaux de sévérité configurables de LOW à CRITICAL.

Le Secret Detection est particulièrement critique : il empêche les commits contenant des clés API AWS, des tokens GitLab, des mots de passe en clair et d'autres secrets. Ce scan s'exécute avant le merge, bloquant les merge requests si un secret est détecté.

Pour les équipes sur le tier Ultimate, le DAST (Dynamic Application Security Testing) teste l'application déployée en environnement de staging, simulant des attaques XSS, injection SQL et CSRF sur les endpoints exposés. C'est un complément essentiel au SAST pour une couverture de sécurité complète.

## Étape 10 : Configurer le Déploiement Multi-Environnements

Un pipeline GitLab CI/CD mature déploie sur plusieurs environnements : staging pour la validation, production pour les utilisateurs finaux, et potionnellement des environnements éphémères par merge request (review apps).

Voici la configuration complète pour un déploiement multi-environnements :

```
# Déploiement staging automatique
deploy-staging:
  stage: deploy
  image: alpine:3.20
  before_script:
    - apk add --no-cache openssh-client curl
    - eval $(ssh-agent -s)
    - echo "$SSH_PRIVATE_KEY" | ssh-add -
  script:
    - ssh -o StrictHostKeyChecking=no deploy@$STAGING_HOST
      "docker pull $CI_REGISTRY_IMAGE:$CI_COMMIT_SHORT_SHA &&
       docker compose -f docker-compose.staging.yml up -d &&
       sleep 5 &&
       curl -f http://localhost:3000/health"
  environment:
    name: staging
    url: https://staging.example.com
    on_stop: stop-staging
  rules:
    - if: $CI_COMMIT_BRANCH == "main"
# Arrêt de l'environnement staging
stop-staging:
  stage: deploy
  image: alpine:3.20
  script:
    - ssh deploy@$STAGING_HOST "docker compose -f docker-compose.staging.yml down"
  environment:
    name: staging
    action: stop
  rules:
    - if: $CI_COMMIT_BRANCH == "main"
      when: manual
# Déploiement production manuel avec approbation
deploy-production:
  stage: deploy
  image: alpine:3.20
  before_script:
    - apk add --no-cache openssh-client curl
    - eval $(ssh-agent -s)
    - echo "$SSH_PRIVATE_KEY_PROD" | ssh-add -
  script:
    - echo "Déploiement de $CI_REGISTRY_IMAGE:$CI_COMMIT_SHORT_SHA en production"
    - ssh -o StrictHostKeyChecking=no deploy@$PRODUCTION_HOST
      "docker pull $CI_REGISTRY_IMAGE:$CI_COMMIT_SHORT_SHA &&
       docker compose -f docker-compose.prod.yml up -d --no-deps app &&
       sleep 10 &&
       curl -f http://localhost:3000/health || 
       (docker compose -f docker-compose.prod.yml rollback && exit 1)"
  environment:
    name: production
    url: https://app.example.com
  rules:
    - if: $CI_COMMIT_BRANCH == "main"
      when: manual
  allow_failure: false
# Review Apps (environnement éphémère par MR)
deploy-review:
  stage: deploy
  image: docker:27
  services:
    - docker:27-dind
  script:
    - docker pull $CI_REGISTRY_IMAGE:$CI_COMMIT_SHORT_SHA
    - docker run -d --name review-$CI_MERGE_REQUEST_IID
      -p $((3000 + $CI_MERGE_REQUEST_IID)):3000
      $CI_REGISTRY_IMAGE:$CI_COMMIT_SHORT_SHA
  environment:
    name: review/$CI_MERGE_REQUEST_SOURCE_BRANCH_NAME
    url: http://$CI_SERVER_HOST:$((3000 + $CI_MERGE_REQUEST_IID))
    on_stop: stop-review
    auto_stop_in: 1 week
  rules:
    - if: $CI_PIPELINE_SOURCE == "merge_request_event"
```
Les review apps créent un environnement éphémère pour chaque merge request. L'équipe peut tester les changements sur une URL dédiée avant le merge. La directive `auto_stop_in: 1 week` garantit que les environnements orphelins sont automatiquement nettoyés.

Le déploiement production utilise `when: manual` pour exiger une validation humaine. Combiné avec les merge request approvals du tier Premium, cela crée un processus de release contrôlé et auditable. Le script inclut un rollback automatique si le health check échoue après déploiement.

