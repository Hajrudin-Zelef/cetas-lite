---
id: collect-261001-rattrapage/rattrapage/tutoriel-github-actions-2026-ci-cd-complet-2
title: ".github/dependabot.yml"
domain: rattrapage
role: reference
task: reference
actors: ["AWS"]
dates: []
keywords: ["aws"]
source: docs/RAG/collect-261001-rattrapage/tutoriel-github-actions-2026-ci-cd-complet.md
source_anchor: ""
source_lines: [120, 264]
sha256: b93b9e82b7808403806c22552a37d76be698bcfc4946d25f0cdf70dbe33c2b39
---

# .github/dependabot.yml

```
name: Tests Matriciels
on:
  pull_request:
    branches: ['main']
jobs:
  test:
    name: Test Node ${{ matrix.node }} sur ${{ matrix.os }}
    runs-on: ${{ matrix.os }}
    strategy:
      fail-fast: false
      max-parallel: 6
      matrix:
        os: [ubuntu-24.04, windows-2025, macos-15]
        node: [18, 20, 22]
        exclude:
          - os: macos-15
            node: 18
    steps:
      - uses: actions/checkout@v4
      - name: Setup Node.js ${{ matrix.node }}
        uses: actions/setup-node@v4
        with:
          node-version: ${{ matrix.node }}
          cache: 'npm'
      - run: npm ci
      - run: npm test
      - name: Upload rapport de couverture
        if: matrix.os == 'ubuntu-24.04' && matrix.node == 22
        uses: actions/upload-artifact@v4
        with:
          name: coverage-report
          path: coverage/
```
**Sortie attendue** : GitHub Actions crée 8 jobs parallèles (3 OS × 3 versions - 1 exclusion). Chaque job apparaît individuellement dans l'interface avec son statut. L'option `fail-fast: false` garantit que tous les jobs s'exécutent même si l'un échoue, ce qui est indispensable pour identifier les incompatibilités spécifiques à un environnement.

Remarquez que le rapport de couverture n'est uploadé que pour une seule combinaison (Ubuntu + Node 22) grâce à la condition `if`. Cette approche évite de créer des artefacts dupliqués inutiles. Pour les projets qui nécessitent des matrices dynamiques, vous pouvez utiliser `fromJSON()` combiné avec un job préliminaire qui génère la matrice via un script ou un appel API.

## Étape 4 : Sécuriser Vos Workflows avec les Bonnes Pratiques 2026

La sécurité des pipelines CI/CD est un enjeu critique en 2026. Depuis juillet 2025, GitHub recommande l'épinglage (pinning) des actions par SHA comme pratique par défaut. Cette mesure empêche les attaques de type supply chain où un mainteneur malveillant modifie le code d'une action populaire. Selon la feuille de route sécurité 2026 relayée par Dev.to, GitHub prévoit d'aller plus loin avec le verrouillage des dépendances de workflow (workflow dependency locking), attendu en preview publique au deuxième ou troisième trimestre 2026 avant une disponibilité générale visée au troisième ou quatrième trimestre.

Voici les règles de sécurité essentielles pour vos workflows GitHub Actions :

**Épinglage des actions par SHA** : au lieu d'utiliser `actions/checkout@v4`, utilisez le hash complet du commit : `actions/checkout@692973e3d937129bcbf40652eb9f2f61becf3332`. Configurez Dependabot pour mettre à jour automatiquement ces SHA lorsque de nouvelles versions sont publiées. Créez un fichier `.github/dependabot.yml` :

```
# .github/dependabot.yml
version: 2
updates:
  - package-ecosystem: "github-actions"
    directory: "/"
    schedule:
      interval: "weekly"
    commit-message:
      prefix: "ci"
    labels:
      - "dependencies"
      - "ci"
```
**Authentification OIDC sans secrets** : au lieu de stocker des clés AWS ou GCP dans les secrets GitHub, utilisez l'authentification OpenID Connect. Cette approche génère des tokens éphémères valides uniquement pour la durée du workflow, éliminant ainsi le risque de fuite de credentials longue durée. En avril 2026, GitHub a enrichi ce mécanisme avec de nouvelles propriétés personnalisées (custom properties) pour OIDC, permettant d'affiner les politiques de confiance, en parallèle de l'ajout de surcharges entrypoint/command pour les conteneurs de service dans Actions.

**Permissions minimales** : définissez toujours `permissions` au niveau du workflow ou du job. Ne laissez jamais les permissions par défaut, qui sont trop larges. Un workflow de build n'a besoin que de `contents: read`, tandis qu'un workflow de déploiement peut nécessiter `deployments: write` et `id-token: write` pour l'OIDC.

**Secrets chiffrés** : GitHub Actions offre la gestion native des secrets chiffrés à trois niveaux — organisation, dépôt et environnement. Les secrets d'environnement sont les plus restrictifs et doivent être privilégiés pour les credentials de production. Combinés avec les règles de protection d'environnement (approbation manuelle, délai d'attente), ils constituent une défense robuste contre les déploiements non autorisés. Pour aller plus loin sur la sécurité des pipelines, consultez notre guide sur l'architecture Zero Trust.

## Étape 5 : Créer des Workflows Réutilisables pour Votre Organisation

Les workflows réutilisables de GitHub Actions permettent de définir des pipelines standards que plusieurs dépôts peuvent appeler. En 2026, les améliorations incluent le support des matrices dans les inputs, les outputs depuis les jobs appelés et une augmentation significative de la limite d'appels imbriqués.

Créez d'abord le workflow réutilisable dans un dépôt dédié (par exemple `org/.github`) :

```
# .github/workflows/reusable-deploy.yml
name: Déploiement Réutilisable
on:
  workflow_call:
    inputs:
      environment:
        description: 'Environnement cible'
        required: true
        type: string
      node-version:
        description: 'Version de Node.js'
        required: false
        type: string
        default: '22'
    secrets:
      DEPLOY_TOKEN:
        required: true
    outputs:
      deploy-url:
        description: 'URL du déploiement'
        value: ${{ jobs.deploy.outputs.url }}
jobs:
  deploy:
    name: Déployer sur ${{ inputs.environment }}
    runs-on: ubuntu-24.04
    environment: ${{ inputs.environment }}
    outputs:
      url: ${{ steps.deploy.outputs.url }}
    steps:
      - uses: actions/checkout@v4
      - uses: actions/setup-node@v4
        with:
          node-version: ${{ inputs.node-version }}
          cache: 'npm'
      - run: npm ci
      - run: npm run build
      - name: Déployer
        id: deploy
        run: |
          echo "Déploiement sur ${{ inputs.environment }}..."
          # Votre logique de déploiement ici
          echo "url=https://${{ inputs.environment }}.example.com" >> $GITHUB_OUTPUT
        env:
          DEPLOY_TOKEN: ${{ secrets.DEPLOY_TOKEN }}
```
Ensuite, appelez ce workflow depuis n'importe quel dépôt de votre organisation :

```
# .github/workflows/deploy-staging.yml
name: Deploy Staging
on:
  push:
    branches: ['develop']
jobs:
  deploy:
    uses: org/.github/.github/workflows/reusable-deploy.yml@main
    with:
      environment: staging
      node-version: '22'
    secrets:
      DEPLOY_TOKEN: ${{ secrets.STAGING_DEPLOY_TOKEN }}
```
**Sortie attendue** : le workflow appelant déclenche le workflow réutilisable, qui s'exécute dans le contexte du dépôt appelant. Les logs montrent les deux niveaux d'exécution. L'output `deploy-url` est accessible dans les jobs suivants du workflow appelant via `needs.deploy.outputs.deploy-url`.

Cette architecture de workflows réutilisables est devenue la norme dans les organisations qui gèrent plus de 10 dépôts. Elle garantit la cohérence des pipelines CI/CD, réduit la duplication de code YAML et facilite les mises à jour centralisées des pratiques de déploiement. Si vous gérez également votre infrastructure avec du code, notre tutoriel Ansible complète parfaitement cette approche.

## Étape 6 : Configurer les Environnements et le Déploiement Progressif

Les environnements GitHub Actions permettent de définir des règles de protection pour chaque étape de votre pipeline de déploiement. En mars 2026, une nouvelle fonctionnalité permet d'utiliser les environnements sans créer automatiquement un déploiement, grâce à la clé `deployment: false`.

Configurez trois environnements dans les paramètres de votre dépôt : **development** (aucune restriction), **staging** (approbation manuelle) et **production** (approbation manuelle + délai de 10 minutes). Chaque environnement possède ses propres secrets et variables.

