---
id: collect-261001-rattrapage/rattrapage/tutoriel-github-actions-2026-ci-cd-complet-3
title: ".github/dependabot.yml"
domain: rattrapage
role: reference
task: reference
actors: ["AWS"]
dates: []
keywords: ["aws"]
source: docs/RAG/collect-261001-rattrapage/tutoriel-github-actions-2026-ci-cd-complet.md
source_anchor: ""
source_lines: [265, 456]
sha256: 176b28eeee5d9e2586b4ed9982a12f913ec90100235894554f04a97f2939eff6
---

# .github/dependabot.yml

```
name: Pipeline de Déploiement Complet
on:
  push:
    branches: ['main']
jobs:
  test:
    name: Tests unitaires
    runs-on: ubuntu-24.04
    steps:
      - uses: actions/checkout@v4
      - uses: actions/setup-node@v4
        with:
          node-version: '22'
          cache: 'npm'
      - run: npm ci
      - run: npm test
  deploy-staging:
    name: Déployer en staging
    needs: test
    runs-on: ubuntu-24.04
    environment:
      name: staging
      url: https://staging.example.com
    steps:
      - uses: actions/checkout@v4
      - uses: actions/setup-node@v4
        with:
          node-version: '22'
          cache: 'npm'
      - run: npm ci && npm run build
      - name: Déployer sur staging
        run: |
          echo "Déploiement staging en cours..."
          # rsync, scp, aws s3 sync, etc.
        env:
          DEPLOY_KEY: ${{ secrets.STAGING_KEY }}
  integration-tests:
    name: Tests d'intégration
    needs: deploy-staging
    runs-on: ubuntu-24.04
    steps:
      - uses: actions/checkout@v4
      - name: Tests E2E avec Playwright
        run: |
          npx playwright install --with-deps
          npx playwright test --config=e2e/playwright.config.ts
        env:
          BASE_URL: https://staging.example.com
  deploy-production:
    name: Déployer en production
    needs: integration-tests
    runs-on: ubuntu-24.04
    environment:
      name: production
      url: https://example.com
    steps:
      - uses: actions/checkout@v4
      - uses: actions/setup-node@v4
        with:
          node-version: '22'
          cache: 'npm'
      - run: npm ci && npm run build
      - name: Déployer en production
        run: |
          echo "Déploiement production en cours..."
        env:
          DEPLOY_KEY: ${{ secrets.PRODUCTION_KEY }}
```
**Sortie attendue** : le pipeline s'exécute séquentiellement — tests → staging → tests d'intégration → production. À l'étape de staging et production, si des règles de protection sont configurées, le workflow se met en pause et affiche « Waiting for review ». Un reviewer désigné doit approuver manuellement pour continuer. Cette approche de déploiement progressif est fondamentale pour garantir la qualité en production.

Pour les applications conteneurisées, combinez cette stratégie avec notre tutoriel Docker Compose pour orchestrer vos services lors du déploiement.

## Étape 7 : Intégrer Docker et les Registres de Conteneurs

L'intégration de Docker avec GitHub Actions est devenue un standard pour les équipes qui déploient des applications conteneurisées. Le runner `ubuntu-24.04` inclut Docker 27.1 préinstallé, et GitHub Container Registry (ghcr.io) offre un stockage gratuit et illimité pour les dépôts publics.

```
name: Build et Push Docker
on:
  push:
    tags: ['v*']
permissions:
  contents: read
  packages: write
jobs:
  docker:
    name: Build image Docker
    runs-on: ubuntu-24.04
    steps:
      - uses: actions/checkout@v4
      - name: Login GitHub Container Registry
        uses: docker/login-action@v3
        with:
          registry: ghcr.io
          username: ${{ github.actor }}
          password: ${{ secrets.GITHUB_TOKEN }}
      - name: Metadata Docker
        id: meta
        uses: docker/metadata-action@v5
        with:
          images: ghcr.io/${{ github.repository }}
          tags: |
            type=semver,pattern={{version}}
            type=semver,pattern={{major}}.{{minor}}
            type=sha,prefix=
      - name: Build et Push
        uses: docker/build-push-action@v6
        with:
          context: .
          push: true
          tags: ${{ steps.meta.outputs.tags }}
          labels: ${{ steps.meta.outputs.labels }}
          cache-from: type=gha
          cache-to: type=gha,mode=max
      - name: Vérifier l'image
        run: |
          docker pull ghcr.io/${{ github.repository }}:sha-${GITHUB_SHA::7}
          docker inspect ghcr.io/${{ github.repository }}:sha-${GITHUB_SHA::7}
```
**Sortie attendue** : lorsqu'un tag `v1.2.3` est poussé, le workflow construit l'image Docker, la tague avec `1.2.3`, `1.2` et le SHA du commit, puis la pousse vers ghcr.io. Le cache GitHub Actions (`type=gha`) accélère les builds suivants de 40 à 70 % en réutilisant les couches Docker inchangées.

L'action `docker/metadata-action` est particulièrement puissante car elle gère automatiquement les tags selon votre stratégie de versioning sémantique. Pour un tag `v2.1.0`, elle produit les tags `2.1.0`, `2.1` et le SHA — ce qui permet aux utilisateurs de votre image de choisir leur niveau de stabilité. Pour approfondir les concepts Docker, consultez notre comparatif Docker vs Kubernetes.

## Étape 8 : Automatiser les Releases et le Changelog

L'automatisation des releases avec GitHub Actions permet de générer automatiquement des changelogs, créer des releases GitHub et publier des packages. Ce workflow utilise les **Conventional Commits** pour catégoriser les changements et déterminer la version suivante automatiquement.

```
name: Release Automatique
on:
  push:
    branches: ['main']
permissions:
  contents: write
  pull-requests: write
jobs:
  release:
    name: Créer une release
    runs-on: ubuntu-24.04
    steps:
      - uses: actions/checkout@v4
        with:
          fetch-depth: 0
      - name: Déterminer la prochaine version
        id: version
        run: |
          LATEST=$(git describe --tags --abbrev=0 2>/dev/null || echo "v0.0.0")
          echo "Version actuelle: $LATEST"
          
          # Analyser les commits depuis le dernier tag
          COMMITS=$(git log ${LATEST}..HEAD --oneline)
          
          if echo "$COMMITS" | grep -q "BREAKING CHANGE\|feat!:"; then
            BUMP="major"
          elif echo "$COMMITS" | grep -q "^feat"; then
            BUMP="minor"
          else
            BUMP="patch"
          fi
          
          # Calculer la nouvelle version
          IFS='.' read -r major minor patch <<< "${LATEST#v}"
          case $BUMP in
            major) major=$((major + 1)); minor=0; patch=0 ;;
            minor) minor=$((minor + 1)); patch=0 ;;
            patch) patch=$((patch + 1)) ;;
          esac
          
          NEW_VERSION="v${major}.${minor}.${patch}"
          echo "version=$NEW_VERSION" >> $GITHUB_OUTPUT
          echo "Nouvelle version: $NEW_VERSION"
      - name: Générer le changelog
        id: changelog
        run: |
          LATEST=$(git describe --tags --abbrev=0 2>/dev/null || echo "")
          {
            echo "changelog<
```
> $GITHUB_OUTPUT
      - name: Créer le tag et la release
        uses: softprops/action-gh-release@v2
        with:
          tag_name: ${{ steps.version.outputs.version }}
          name: Release ${{ steps.version.outputs.version }}
          body: ${{ steps.changelog.outputs.changelog }}
          generate_release_notes: true **Sortie attendue** : chaque push sur `main` analyse les messages de commit, détermine le type de bump (major, minor, patch), génère un changelog structuré et crée une release GitHub avec les notes automatiques. Cette automatisation élimine les erreurs humaines et garantit un historique de releases cohérent et traçable.

## Étape 9 : Déployer sur AWS avec OIDC (Sans Secrets Statiques)

L'authentification OIDC (OpenID Connect) entre GitHub Actions et AWS est la méthode recommandée en 2026 pour les déploiements cloud. Elle élimine le besoin de stocker des clés d'accès AWS dans les secrets GitHub en générant des tokens éphémères valides uniquement pour la durée du workflow.

