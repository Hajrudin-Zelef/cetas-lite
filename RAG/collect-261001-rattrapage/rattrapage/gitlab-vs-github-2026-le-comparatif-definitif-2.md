---
id: collect-261001-rattrapage/rattrapage/gitlab-vs-github-2026-le-comparatif-definitif-2
title: ".gitlab-ci.yml"
domain: rattrapage
role: reference
task: reference
actors: ["Anthropic", "Apple", "Google", "Microsoft", "OpenAI"]
dates: []
keywords: ["benchmarks", "copilot", "license", "open source"]
source: docs/RAG/collect-261001-rattrapage/gitlab-vs-github-2026-le-comparatif-definitif.md
source_anchor: ""
source_lines: [65, 172]
sha256: acc9c485e7e94fda166672de36d530dd7fd0bd173acfc794eda13edb4869df97
---

# .gitlab-ci.yml

GitLab CI/CD se configure via un fichier `.gitlab-ci.yml` à la racine du projet. La syntaxe est déclarative et puissante, avec un support natif des environnements, des revues d’applications, des déploiements canary et du feature flagging. Voici un exemple typique :

```
# .gitlab-ci.yml
stages:
  - build
  - test
  - security
  - deploy
build:
  stage: build
  image: node:20-alpine
  script:
    - npm ci
    - npm run build
  artifacts:
    paths:
      - dist/
test:
  stage: test
  script:
    - npm run test:coverage
  coverage: '/All files\s+\|\s+(\d+\.?\d*)\%/'
sast:
  stage: security
  include:
    - template: Security/SAST.gitlab-ci.yml
deploy_production:
  stage: deploy
  environment:
    name: production
    url: https://app.example.com
  script:
    - kubectl apply -f k8s/
  only:
    - main
```
GitHub Actions utilise des fichiers YAML dans `.github/workflows/`. L’approche est plus modulaire, basée sur des « actions » réutilisables que l’on compose ensemble :

```
# .github/workflows/ci.yml
name: CI/CD Pipeline
on:
  push:
    branches: [main]
  pull_request:
    branches: [main]
jobs:
  build-and-test:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - uses: actions/setup-node@v4
        with:
          node-version: 20
          cache: 'npm'
      - run: npm ci
      - run: npm run build
      - run: npm run test:coverage
  deploy:
    needs: build-and-test
    if: github.ref == 'refs/heads/main'
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - uses: azure/k8s-deploy@v5
        with:
          manifests: k8s/
          namespace: production
```
Les deux approches sont efficaces, mais GitLab offre des fonctionnalités avancées comme les **pipelines parents-enfants**, les **pipelines multi-projets** et les **environnements de review** directement dans le fichier de configuration. Sur GitHub, ces patterns nécessitent des workarounds plus complexes ou des actions tierces.

### Minutes CI/CD et Coûts d’Exécution

Un point critique en 2026 concerne les minutes d’exécution incluses. GitHub a fait polémique début 2026 en annonçant puis en repoussant indéfiniment l’introduction de frais de plateforme de **0,002 $/minute** pour les runners auto-hébergés, suite à la réaction négative de la communauté. En revanche, les tarifs des runners hébergés ont été réduits de **25 %** en janvier 2026 (Linux 2-core à 0,006 $/min).

GitHub Team inclut **3 000 minutes/mois** d’Actions et 2 Go de stockage Packages. Le plan gratuit offre 2 000 minutes. GitLab Free propose **400 minutes/mois** de CI/CD et 10 Go de stockage, tandis que les plans payants offrent des quotas significativement plus élevés. Pour les projets avec de gros volumes de builds, le coût des minutes supplémentaires peut rapidement faire pencher la balance.

Comme le souligne **Fireship** dans son analyse comparative : « GitHub Actions est devenu incroyablement puissant avec son marketplace, mais GitLab CI/CD reste le roi pour les pipelines complexes d’entreprise. Si votre pipeline fait plus de 200 lignes, GitLab est probablement votre meilleur choix. »

## Intelligence Artificielle : GitHub Copilot vs GitLab Duo

L’intégration de l’IA dans le workflow de développement est devenue un facteur de différenciation majeur en 2026. Les deux plateformes ont investi massivement dans ce domaine, mais avec des approches distinctes qui reflètent leurs philosophies respectives.

**GitHub Copilot**, lancé en 2021 et constamment amélioré depuis, est l’assistant IA le plus utilisé au monde pour le code. En 2026, il propose trois paliers : Individual à **10 $/mois**, Business à **19 $/mois** et Enterprise à **39 $/mois**. Copilot s’intègre dans VS Code, JetBrains, Neovim et directement dans l’interface GitHub. Ses fonctionnalités incluent la complétion de code en temps réel, le chat contextuel, la génération de pull requests, le résumé d’issues et depuis fin 2025, les **Copilot Workspaces** qui permettent de planifier et implémenter des changements complexes de manière autonome.

**GitLab Duo**, la réponse de GitLab, s’intègre directement dans la plateforme DevOps. Il propose la complétion de code dans l’IDE, le chat contextuel, mais aussi des fonctionnalités spécifiques au workflow DevOps : résumé automatique de merge requests, analyse de vulnérabilités avec suggestions de correction, explication de code dans le contexte des pipelines CI/CD, et génération de tests. GitLab Duo est **inclus dans le plan Ultimate** (99 $/mois), ce qui en fait un argument de vente fort pour les entreprises qui auraient autrement payé Copilot en supplément.

En termes de qualité brute de suggestion de code, GitHub Copilot conserve un avantage grâce à son partenariat avec OpenAI et l’accès aux modèles GPT-4o et Codex. Les benchmarks internes de GitHub montrent un taux d’acceptation des suggestions de **30 à 35 %** et une réduction du temps de codage de **55 %** en moyenne. GitLab Duo, qui utilise des modèles propriétaires et des partenariats avec Anthropic et Google, affiche des performances comparables sur les langages les plus courants mais reste en retrait sur les langages moins populaires.

**MKBHD**, dans sa couverture des outils de productivité tech en février 2026, a noté : « GitHub Copilot est à la génération de code ce que l’iPhone a été au smartphone – pas le premier, mais celui qui a défini la catégorie. GitLab Duo rattrape son retard rapidement, surtout pour les équipes qui veulent tout intégrer dans une seule plateforme. » Cette observation résume bien l’état du marché : Copilot domine en qualité pure, Duo l’emporte en intégration.

## Sécurité et DevSecOps : L’Avantage GitLab

Si le CI/CD est le champ de bataille, la **sécurité applicative** est là où GitLab prend clairement l’avantage en 2026. La plateforme propose un arsenal complet de tests de sécurité intégrés nativement dans les pipelines, sans outil tiers nécessaire :

- **SAST (Static Application Security Testing)** : analyse statique du code source pour détecter les vulnérabilités avant le déploiement
- **DAST (Dynamic Application Security Testing)** : tests de sécurité dynamiques sur les applications déployées
- **Dependency Scanning** : détection de vulnérabilités dans les dépendances tierces
- **Container Scanning** : analyse des images Docker pour identifier les CVE connues
- **Secret Detection** : identification des secrets et credentials exposés dans le code
- **Fuzz Testing** : tests de fuzzing automatisés pour découvrir des bugs de sécurité
- **License Compliance** : vérification automatique des licences open source

Toutes ces fonctionnalités sont disponibles dans **GitLab Ultimate** et s’activent avec quelques lignes dans le fichier `.gitlab-ci.yml` grâce aux templates prédéfinis. Les résultats apparaissent directement dans les merge requests, permettant aux développeurs de corriger les problèmes avant la fusion du code.

GitHub propose un ensemble de fonctionnalités de sécurité via **GitHub Advanced Security (GHAS)**, vendu séparément à **49 $/utilisateur/mois** en plus du plan Enterprise. GHAS inclut le code scanning (via CodeQL), le secret scanning et le dependency review. GitHub propose également **Dependabot** gratuitement pour les mises à jour automatiques des dépendances vulnérables. Bien que puissantes, ces fonctionnalités restent moins complètes que l’offre GitLab : pas de DAST natif, pas de fuzz testing intégré, et la couverture de Container Scanning nécessite des intégrations tierces.

