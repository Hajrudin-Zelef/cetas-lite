---
id: collect-261001-rattrapage/rattrapage/tutoriel-github-actions-2026-ci-cd-complet-1
title: ".github/dependabot.yml"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["attention", "gpu"]
source: docs/RAG/collect-261001-rattrapage/tutoriel-github-actions-2026-ci-cd-complet.md
source_anchor: ""
source_lines: [1, 119]
sha256: 80dc829dc8520a6b65eb054760494a53eed0903c0ace51d93f9458faebe2cc13
---

# .github/dependabot.yml

GitHub Actions s’est imposé comme la plateforme de CI/CD la plus adoptée au monde, avec plus de 15 000 actions disponibles sur le marketplace et une intégration native avec l’écosystème GitHub. En mars 2026, les nouvelles fonctionnalités comme les runners Arm64, les runners GPU, les actions immuables et les workflows réutilisables améliorés transforment radicalement les pipelines de déploiement. Ce tutoriel complet vous guide pas à pas pour maîtriser GitHub Actions, de la configuration initiale jusqu’au déploiement en production d’un projet complet.

Que vous soyez développeur backend, frontend ou DevOps, ce guide pratique couvre les 12 étapes essentielles pour automatiser vos builds, tests et déploiements avec GitHub Actions. Chaque section inclut des exemples de code fonctionnels, des sorties attendues et les pièges courants à éviter. À la fin de ce tutoriel, vous aurez un pipeline CI/CD complet et prêt pour la production.

## Prérequis et Environnement de Travail pour GitHub Actions

Avant de commencer ce tutoriel GitHub Actions, assurez-vous de disposer des outils et versions suivants. La compatibilité de votre environnement est essentielle pour reproduire chaque étape sans erreur.

| Outil | Version Minimale | Version Recommandée (Mars 2026) | Rôle | 
|---|---|---|---|
| Git | 2.40+ | 2.46 | Gestion de version et push vers GitHub | 
| Node.js | 18 LTS | 22.x LTS | Runtime JavaScript pour le projet exemple | 
| Docker | 24.0+ | 27.1 | Conteneurisation et tests d’intégration | 
| GitHub CLI (gh) | 2.40+ | 2.65 | Gestion des workflows depuis le terminal | 
| Compte GitHub | Free | Pro ou Team | Accès aux runners et minutes CI/CD | 
| VS Code | 1.85+ | 1.98 | Édition avec extension GitHub Actions | 

Concernant les minutes gratuites, GitHub offre 2 000 minutes par mois sur le plan Free, 3 000 sur Pro, 10 000 sur Team et plus de 50 000 sur Enterprise. Les dépôts publics bénéficient de minutes illimitées. Le coût des runners Linux standard s’élève à 0,008 $ par minute sur les dépôts privés, tandis que les runners macOS coûtent 0,08 $ par minute — soit dix fois plus. Depuis le 1er janvier 2026, GitHub a réduit ces tarifs de 15 à 39 % selon l’analyse tarifaire de WarpBuild, rendant les runners hébergés encore plus compétitifs face aux solutions self-hosted.

Installez l’extension **GitHub Actions** pour VS Code, qui offre la validation YAML en temps réel, l’autocomplétion des noms d’actions et la visualisation des logs directement dans l’éditeur. Utilisez également **actionlint**, le linter officiel, pour détecter les erreurs de syntaxe avant de pousser vos workflows.

## Étape 1 : Créer Votre Premier Workflow GitHub Actions

Un workflow GitHub Actions est un fichier YAML placé dans le répertoire `.github/workflows/` de votre dépôt. Chaque workflow définit un ou plusieurs jobs qui s’exécutent sur des runners hébergés par GitHub ou auto-hébergés. Commençons par créer un workflow simple qui s’exécute à chaque push sur la branche principale.

Créez la structure de répertoires et le fichier de workflow :

```
mkdir -p .github/workflows
cat > .github/workflows/ci.yml << 'EOF'
name: CI Pipeline
on:
  push:
    branches: ['main', 'develop']
  pull_request:
    branches: ['main']
permissions:
  contents: read
defaults:
  run:
    shell: bash
jobs:
  build:
    name: Build & Test
    runs-on: ubuntu-24.04
    timeout-minutes: 15
    steps:
      - name: Checkout du code
        uses: actions/checkout@v4
      - name: Configuration de Node.js
        uses: actions/setup-node@v4
        with:
          node-version: '22'
          cache: 'npm'
      - name: Installation des dépendances
        run: npm ci
      - name: Linting du code
        run: npm run lint
      - name: Exécution des tests
        run: npm test
      - name: Build de production
        run: npm run build
EOF
```
**Sortie attendue** : après un push sur la branche `main`, le workflow apparaît dans l'onglet Actions de votre dépôt. Chaque étape affiche un indicateur vert (succès) ou rouge (échec) avec les logs détaillés. Le runner `ubuntu-24.04` est la version par défaut de `ubuntu-latest` depuis janvier 2025, basée sur Ubuntu Noble Numbat avec Node.js 22.x, Python 3.12 et Go 1.23 préinstallés.

Notez l'utilisation de `permissions: contents: read` en début de fichier. C'est une bonne pratique de sécurité qui limite les droits du token GITHUB_TOKEN au strict minimum nécessaire. Par défaut, ce token dispose de permissions étendues qui ne sont pas toujours requises. Depuis le déploiement progressif lancé le 24 avril 2026, ce token stateless (préfixe `ghs_`) est passé d'environ 40 à près de 520 caractères, un changement de format documenté par Dev.to qu'il faut anticiper si vos scripts parsent ou stockent le token dans des variables de taille fixe.

## Étape 2 : Configurer le Cache et les Artefacts pour des Builds Rapides

Les performances de votre pipeline GitHub Actions dépendent fortement de la gestion du cache. L'action `actions/cache@v4`, mise à jour en novembre 2025, offre la compression automatique (réduction de 50 % de la taille), la recherche parallèle des clés de restauration et une limite de stockage doublée à 10 Go (contre 5 Go précédemment).

```
name: CI avec Cache Optimisé
on:
  push:
    branches: ['main']
jobs:
  build:
    runs-on: ubuntu-24.04
    steps:
      - uses: actions/checkout@v4
      - name: Setup Node.js avec cache npm
        uses: actions/setup-node@v4
        with:
          node-version: '22'
          cache: 'npm'
      - name: Cache des dépendances npm
        uses: actions/cache@v4
        id: npm-cache
        with:
          path: |
            ~/.npm
            node_modules
          key: ${{ runner.os }}-node-22-${{ hashFiles('**/package-lock.json') }}
          restore-keys: |
            ${{ runner.os }}-node-22-
            ${{ runner.os }}-node-
      - name: Installation conditionnelle
        if: steps.npm-cache.outputs.cache-hit != 'true'
        run: npm ci
      - name: Build et upload artefact
        run: npm run build
      - name: Upload des artefacts de build
        uses: actions/upload-artifact@v4
        with:
          name: build-output
          path: dist/
          retention-days: 30
          compression-level: 6
```
**Sortie attendue** : lors du premier run, le cache est créé (« Cache not found for key... »). Lors des runs suivants, vous verrez « Cache restored from key... » et l'étape d'installation est ignorée grâce à la condition `if`. Cela réduit le temps de build de 60 à 90 secondes en moyenne. L'action `upload-artifact@v4`, mise à jour en décembre 2025, supporte désormais la compression Brotli et Zstandard, réduisant la taille des artefacts de 30 % par rapport à la v3.

Attention : les actions `actions/cache@v3` et `upload-artifact@v3` ont atteint leur fin de vie en janvier 2026. Si vous utilisez encore ces versions, migrez immédiatement vers la v4 pour éviter les échecs de workflow.

## Étape 3 : Implémenter les Tests avec une Stratégie Matricielle

La stratégie matricielle de GitHub Actions permet d'exécuter vos tests en parallèle sur plusieurs combinaisons de systèmes d'exploitation, versions de runtime et configurations. En 2026, la limite passe à 1 000 combinaisons par matrice, et la directive `generate` permet de créer des matrices dynamiques à partir d'appels API.

