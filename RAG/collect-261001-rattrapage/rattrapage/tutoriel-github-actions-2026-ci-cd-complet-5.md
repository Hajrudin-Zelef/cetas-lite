---
id: collect-261001-rattrapage/rattrapage/tutoriel-github-actions-2026-ci-cd-complet-5
title: ".github/dependabot.yml"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/tutoriel-github-actions-2026-ci-cd-complet.md
source_anchor: ""
source_lines: [587, 730]
sha256: bfb1576dbef80ca33e4e92bd8e4b94493729bb460f2febcf067a47c0896c858f
---

# .github/dependabot.yml

```
# Structure du projet
mon-app/
├── .github/
│   ├── dependabot.yml
│   └── workflows/
│       ├── ci.yml          # Tests et linting
│       ├── deploy.yml       # Déploiement staging/prod
│       └── release.yml      # Releases automatiques
├── src/
│   ├── index.ts
│   ├── routes/
│   └── services/
├── tests/
│   ├── unit/
│   └── e2e/
├── Dockerfile
├── package.json
└── tsconfig.json
# --- .github/workflows/ci.yml ---
name: CI Complet
on:
  push:
    branches: ['main', 'develop']
  pull_request:
    branches: ['main']
concurrency:
  group: ci-${{ github.ref }}
  cancel-in-progress: true
permissions:
  contents: read
  checks: write
jobs:
  lint:
    name: Linting
    runs-on: ubuntu-24.04
    steps:
      - uses: actions/checkout@v4
      - uses: actions/setup-node@v4
        with:
          node-version: '22'
          cache: 'npm'
      - run: npm ci
      - run: npm run lint
      - run: npx tsc --noEmit
  test:
    name: Tests (${{ matrix.shard }}/4)
    runs-on: ubuntu-24.04
    strategy:
      matrix:
        shard: [1, 2, 3, 4]
    steps:
      - uses: actions/checkout@v4
      - uses: actions/setup-node@v4
        with:
          node-version: '22'
          cache: 'npm'
      - run: npm ci
      - run: npm test -- --shard=${{ matrix.shard }}/4
      - uses: actions/upload-artifact@v4
        if: always()
        with:
          name: test-results-${{ matrix.shard }}
          path: test-results/
  security:
    name: Audit de sécurité
    runs-on: ubuntu-24.04
    steps:
      - uses: actions/checkout@v4
      - run: npm audit --audit-level=high
      - name: Scan CodeQL
        uses: github/codeql-action/analyze@v3
  build:
    name: Build Docker
    needs: [lint, test, security]
    runs-on: ubuntu-24.04
    permissions:
      contents: read
      packages: write
    steps:
      - uses: actions/checkout@v4
      - uses: docker/login-action@v3
        with:
          registry: ghcr.io
          username: ${{ github.actor }}
          password: ${{ secrets.GITHUB_TOKEN }}
      - uses: docker/build-push-action@v6
        with:
          context: .
          push: ${{ github.event_name == 'push' }}
          tags: ghcr.io/${{ github.repository }}:${{ github.sha }}
          cache-from: type=gha
          cache-to: type=gha,mode=max
```
Ce pipeline complet illustre les meilleures pratiques de 2026 : exécution parallèle du linting, des tests fragmentés (sharded) et de l'audit de sécurité, suivie du build Docker conditionnel. La directive `concurrency` annule automatiquement les workflows en cours lorsqu'un nouveau commit est poussé sur la même branche, économisant ainsi des minutes précieuses.

## Les 7 Pièges Courants de GitHub Actions et Comment les Éviter

Après avoir mis en place votre pipeline, voici les erreurs les plus fréquentes que commettent les développeurs avec GitHub Actions, et comment les éviter.

**Piège 1 : Ne pas épingler les actions par SHA.** Utiliser `@v4` au lieu du hash complet expose votre pipeline à des attaques supply chain. Un mainteneur compromis peut modifier le code d'une action populaire et exécuter du code malveillant dans votre CI. Solution : épinglez systématiquement par SHA et utilisez Dependabot pour les mises à jour.

**Piège 2 : Permissions GITHUB_TOKEN trop larges.** Par défaut, le token dispose de permissions étendues. Sans restriction explicite via la directive `permissions`, un workflow compromis peut modifier votre code, créer des releases ou accéder à des packages privés. Définissez toujours les permissions minimales nécessaires.

**Piège 3 : Ignorer le cache.** Sans cache, chaque run télécharge toutes les dépendances depuis zéro. Pour un projet Node.js typique, cela représente 60 à 120 secondes perdues par run. Multipliez par le nombre de PR et de push quotidiens, et le coût s'accumule rapidement.

**Piège 4 : Utiliser `actions/cache@v3` ou `upload-artifact@v3`.** Ces versions sont en fin de vie depuis janvier 2026. Les workflows qui les utilisent échoueront progressivement. Migrez vers la v4 immédiatement.

**Piège 5 : Ne pas utiliser `concurrency`.** Sans cette directive, chaque push déclenche un nouveau workflow, même si le précédent n'est pas terminé. Résultat : des workflows redondants qui consomment vos minutes gratuites. Ajoutez `concurrency: { group: ci-${{ github.ref }}, cancel-in-progress: true }`.

**Piège 6 : Secrets dans les logs.** Les commandes `echo` et `printenv` peuvent accidentellement afficher des secrets dans les logs. GitHub masque automatiquement les secrets connus, mais les valeurs dérivées (comme une URL contenant un token) ne sont pas masquées. Utilisez `::add-mask::VALUE` pour les valeurs sensibles calculées.

**Piège 7 : Timeout par défaut trop long.** Sans `timeout-minutes`, un job peut tourner jusqu'à 6 heures (360 minutes), consommant votre quota silencieusement. Définissez toujours un timeout raisonnable — 15 minutes pour un build standard, 30 pour des tests E2E.

## Guide de Dépannage Complet pour GitHub Actions

Voici les problèmes les plus fréquemment rencontrés avec GitHub Actions et leurs solutions détaillées.

**Problème 1 : « Resource not accessible by integration »**

Ce message apparaît lorsque le GITHUB_TOKEN n'a pas les permissions requises. Vérifiez la section `permissions` de votre workflow. Pour les actions qui créent des PR ou modifient des fichiers, ajoutez `pull-requests: write` ou `contents: write`. Dans les dépôts d'organisation, vérifiez aussi les paramètres au niveau de l'organisation sous Settings → Actions → General → Workflow permissions.

**Problème 2 : « No space left on device »**

Les runners standard disposent d'environ 14 Go d'espace libre. Les projets volumineux ou les builds Docker multicouches peuvent épuiser cet espace. Solutions : utilisez `docker system prune -af` en début de workflow, supprimez les outils préinstallés inutiles avec `sudo rm -rf /usr/local/lib/android /usr/share/dotnet`, ou passez à un runner larger avec 2 To de stockage.

**Problème 3 : Cache introuvable malgré la clé correcte**

Le cache GitHub Actions est lié à la branche. Un cache créé sur `main` est accessible depuis les branches enfants, mais pas l'inverse. Si votre PR ne trouve pas le cache, assurez-vous qu'un workflow a déjà créé le cache sur la branche de base. Utilisez `restore-keys` pour un fallback progressif.

**Problème 4 : « Error: Process completed with exit code 1 » sans détails**

Ce message générique indique qu'une commande a échoué. Ajoutez `set -euxo pipefail` en début de vos scripts shell pour obtenir des logs détaillés. Le flag `-x` affiche chaque commande avant son exécution, et `-u` échoue sur les variables non définies.

**Problème 5 : Workflows qui ne se déclenchent pas**

Vérifiez ces points : le fichier YAML doit être sur la branche par défaut pour les triggers `push` et `pull_request`. Les paths et branches doivent correspondre exactement (sensible à la casse). Les workflows désactivés doivent être réactivés manuellement. Les fork PRs nécessitent une approbation pour les premiers contributeurs.

**Problème 6 : « Timeout exceeded » sur les tests**

Les tests qui passent en local mais échouent en CI sont souvent liés à la puissance limitée des runners (2 vCPU / 7 Go RAM). Solutions : augmentez `timeout-minutes`, utilisez le test sharding pour répartir la charge, ou passez à un runner plus puissant. Pour les tests Playwright ou Cypress, ajoutez `--retries=2` pour gérer les flaky tests.

**Problème 7 : Secrets non disponibles dans les fork PRs**

