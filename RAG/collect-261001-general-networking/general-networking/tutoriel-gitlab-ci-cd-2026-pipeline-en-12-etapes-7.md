---
id: collect-261001-general-networking/general-networking/tutoriel-gitlab-ci-cd-2026-pipeline-en-12-etapes-7
title: "git version 2.47.1"
domain: general-networking
role: reference
task: reference
actors: ["AWS", "Microsoft"]
dates: []
keywords: ["agent", "attention", "aws", "copilot", "open source"]
source: docs/RAG/collect-261001-general-networking/tutoriel-gitlab-ci-cd-2026-pipeline-en-12-etapes.md
source_anchor: ""
source_lines: [826, 1017]
sha256: 121d07ae523036d56a5b062e2445034aefe41a9f3e1abc2388f439d5bd0218ef
---

# git version 2.47.1

```
# .gitlab-ci.yml — Pipeline de production complet
stages:
  - validate
  - test
  - build
  - deploy
  - release
variables:
  NODE_ENV: "test"
  DOCKER_IMAGE: "$CI_REGISTRY_IMAGE:$CI_COMMIT_SHORT_SHA"
include:
  - template: Security/SAST.gitlab-ci.yml
  - template: Security/Dependency-Scanning.gitlab-ci.yml
  - template: Security/Secret-Detection.gitlab-ci.yml
# Template partagé
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
# === VALIDATE ===
lint:
  extends: .node-defaults
  stage: validate
  script:
    - npx eslint src/ tests/ --max-warnings 0
  rules:
    - if: $CI_PIPELINE_SOURCE == "merge_request_event"
    - if: $CI_COMMIT_BRANCH == "main"
# === TEST ===
unit-tests:
  extends: .node-defaults
  stage: test
  needs: ["lint"]
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
integration-tests:
  extends: .node-defaults
  stage: test
  needs: ["lint"]
  services:
    - name: postgres:16-alpine
      alias: db
  variables:
    POSTGRES_DB: test_db
    POSTGRES_USER: test_user
    POSTGRES_PASSWORD: test_pass
    DATABASE_URL: "postgresql://test_user:test_pass@db:5432/test_db"
  script:
    - npm run test:integration
  rules:
    - if: $CI_COMMIT_BRANCH == "main"
    - if: $CI_PIPELINE_SOURCE == "merge_request_event"
# === BUILD ===
docker-build:
  stage: build
  image: docker:27
  services:
    - docker:27-dind
  needs: ["unit-tests"]
  variables:
    DOCKER_TLS_CERTDIR: "/certs"
  before_script:
    - docker login -u $CI_REGISTRY_USER -p $CI_REGISTRY_PASSWORD $CI_REGISTRY
  script:
    - docker build
      --cache-from $CI_REGISTRY_IMAGE:latest
      --tag $DOCKER_IMAGE
      --tag $CI_REGISTRY_IMAGE:latest .
    - docker push $DOCKER_IMAGE
    - docker push $CI_REGISTRY_IMAGE:latest
  rules:
    - if: $CI_COMMIT_BRANCH == "main"
# === DEPLOY ===
deploy-staging:
  stage: deploy
  image: alpine:3.20
  needs: ["docker-build", "integration-tests"]
  before_script:
    - apk add --no-cache openssh-client curl
    - eval $(ssh-agent -s)
    - echo "$SSH_PRIVATE_KEY" | ssh-add -
  script:
    - ssh -o StrictHostKeyChecking=no deploy@$STAGING_HOST
      "docker pull $DOCKER_IMAGE &&
       docker compose -f docker-compose.staging.yml up -d &&
       sleep 5 && curl -f http://localhost:3000/health"
  environment:
    name: staging
    url: https://staging.example.com
  rules:
    - if: $CI_COMMIT_BRANCH == "main"
deploy-production:
  stage: deploy
  image: alpine:3.20
  needs: ["deploy-staging"]
  before_script:
    - apk add --no-cache openssh-client curl
    - eval $(ssh-agent -s)
    - echo "$SSH_PRIVATE_KEY_PROD" | ssh-add -
  script:
    - ssh -o StrictHostKeyChecking=no deploy@$PRODUCTION_HOST
      "docker pull $DOCKER_IMAGE &&
       docker compose -f docker-compose.prod.yml up -d"
  environment:
    name: production
    url: https://app.example.com
  rules:
    - if: $CI_COMMIT_BRANCH == "main"
      when: manual
  allow_failure: false
# === RELEASE ===
create-release:
  stage: release
  image: registry.gitlab.com/gitlab-org/release-cli:latest
  rules:
    - if: $CI_COMMIT_TAG =~ /^v\d+\.\d+\.\d+$/
  script:
    - echo "Release $CI_COMMIT_TAG"
  release:
    tag_name: $CI_COMMIT_TAG
    description: "Release $CI_COMMIT_TAG"
```
Ce pipeline s'exécute en environ 5 à 8 minutes grâce aux DAG (`needs`), au cache optimisé et au build Docker avec cache de couches. Sur un runner self-hosted avec SSD, les temps descendent à 3-4 minutes — des durées que vous pouvez désormais suivre plus finement grâce au backend d'analytics migré vers ClickHouse depuis GitLab 18.0, qui améliore la précision des métriques de pipeline (documentation mise à jour en août 2026).

## GitLab CI/CD vs GitHub Actions : Quand Choisir Quoi

La question du choix entre GitLab CI/CD et GitHub Actions revient fréquemment. Voici un comparatif objectif basé sur les fonctionnalités de 2026.

| Critère | GitLab CI/CD | GitHub Actions | 
|---|---|---|
| Hébergement | SaaS ou Self-Managed | SaaS uniquement (GHES limité) | 
| Sécurité intégrée | SAST, DAST, Secret Detection, Dependency Scanning | Via actions tierces (CodeQL) | 
| Container Registry | Intégré nativement | ghcr.io (séparé) | 
| Minutes gratuites | 400/mois (SaaS Free) | 2 000/mois (Free) | 
| Environnements | Natif avec review apps | Environments (basique) | 
| DAG / Parallélisme | needs + parallel matrix | needs (similaire) | 
| Compliance | Politiques de conformité (Ultimate) | Branch protection rules | 
| IA intégrée | GitLab Duo Chat | Copilot (payant séparé) | 

GitLab CI/CD est le meilleur choix pour les entreprises qui veulent une plateforme DevSecOps tout-en-un, avec le contrôle d'une instance self-managed et des fonctionnalités de conformité avancées. GitHub Actions est plus adapté aux projets open source et aux équipes qui utilisent déjà l'écosystème GitHub pour la gestion du code et des issues.

Pour les organisations européennes soumises au RGPD et au NIS2, GitLab self-managed offre un avantage significatif : l'hébergement des données CI/CD reste entièrement sous votre contrôle, sur vos propres serveurs ou un cloud souverain.

### Articles Connexes

Pour approfondir les sujets abordés dans ce tutoriel GitLab CI/CD, consultez nos guides complémentaires :

- Comment Construire un Pipeline CI/CD avec GitHub Actions — le concurrent direct de GitLab CI/CD
- Tutoriel Docker pour Débutants — maîtrisez les conteneurs utilisés par les runners GitLab
- Tutoriel Docker Compose — orchestrez vos applications multi-conteneurs
- GitHub vs GitLab 2026 — comparatif complet des deux plateformes
- Déployer avec Kubernetes et Helm — l'étape suivante après Docker pour vos déploiements
- Infrastructure AWS avec Terraform — automatisez votre infrastructure cloud

## FAQ : Questions Fréquentes sur GitLab CI/CD

### Combien coûte GitLab CI/CD ?

Le tier Free offre 400 minutes CI/CD par mois sur les runners partagés SaaS. Le tier Premium (29 $/utilisateur/mois) ajoute les merge request approvals et les environnements protégés. Le tier Ultimate (99 $/utilisateur/mois) débloque le DAST, le dependency scanning avancé et les politiques de conformité. Pour les runners self-hosted, seul le coût du serveur s'applique.

### Puis-je utiliser GitLab CI/CD avec un dépôt GitHub ?

Oui, GitLab supporte le mirroring de dépôts GitHub. Vous pouvez configurer un mirror automatique qui synchronise votre dépôt GitHub vers GitLab, déclenchant les pipelines CI/CD à chaque push. Cette approche est utile pour les équipes qui veulent conserver GitHub pour le code mais utiliser GitLab pour la CI/CD.

### Quelle est la différence entre cache et artefacts ?

Le cache persiste entre les pipelines (best-effort) et sert à stocker les dépendances (node_modules, pip packages). Les artefacts sont garantis dans le même pipeline et transmettent les résultats de build entre les stages. Utilisez le cache pour les dépendances, les artefacts pour les binaires compilés et les rapports de test.

### Comment déboguer un pipeline qui échoue silencieusement ?

Activez le mode debug avec la variable `CI_DEBUG_TRACE: "true"` dans les variables du pipeline (pas dans .gitlab-ci.yml pour éviter de commiter). Cela affiche chaque commande exécutée et les variables d'environnement. Attention : les variables masquées sont visibles en mode debug.

### GitLab CI/CD supporte-t-il les runners Windows et macOS ?

