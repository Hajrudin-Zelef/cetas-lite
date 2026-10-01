---
id: collect-261001-rattrapage/rattrapage/tutoriel-github-actions-2026-ci-cd-complet-4
title: ".github/dependabot.yml"
domain: rattrapage
role: reference
task: reference
actors: ["AWS", "Apple", "Microsoft", "Nvidia"]
dates: []
keywords: ["aws", "copilot", "distribution", "gpu", "graviton", "nvidia"]
source: docs/RAG/collect-261001-rattrapage/tutoriel-github-actions-2026-ci-cd-complet.md
source_anchor: ""
source_lines: [457, 586]
sha256: a2d94cd7d760ca0335dec2bd1d757d20372aeb5505137b8deebeb795c5febbd5
---

# .github/dependabot.yml

```
name: Déployer sur AWS via OIDC
on:
  push:
    branches: ['main']
permissions:
  id-token: write
  contents: read
jobs:
  deploy-aws:
    name: Déployer sur AWS S3 + CloudFront
    runs-on: ubuntu-24.04
    environment: production
    steps:
      - uses: actions/checkout@v4
      - name: Configurer les credentials AWS via OIDC
        uses: aws-actions/configure-aws-credentials@v4
        with:
          role-to-assume: arn:aws:iam::123456789012:role/GitHubActionsDeployRole
          aws-region: eu-west-3
      - uses: actions/setup-node@v4
        with:
          node-version: '22'
          cache: 'npm'
      - run: npm ci && npm run build
      - name: Synchroniser avec S3
        run: |
          aws s3 sync dist/ s3://mon-bucket-production \
            --delete \
            --cache-control "public, max-age=31536000" \
            --exclude "index.html" \
            --exclude "*.json"
          
          aws s3 cp dist/index.html s3://mon-bucket-production/index.html \
            --cache-control "public, max-age=0, must-revalidate"
      - name: Invalider le cache CloudFront
        run: |
          aws cloudfront create-invalidation \
            --distribution-id ${{ vars.CLOUDFRONT_DIST_ID }} \
            --paths "/*"
          echo "Cache CloudFront invalidé avec succès"
```
Pour que ce workflow fonctionne, vous devez d'abord configurer le fournisseur d'identité OIDC dans AWS IAM. Créez un Identity Provider de type OpenID Connect avec l'URL `https://token.actions.githubusercontent.com` et l'audience `sts.amazonaws.com`. Puis créez un rôle IAM avec une politique de confiance qui autorise votre dépôt GitHub spécifique.

La région `eu-west-3` (Paris) est utilisée ici pour la conformité avec les réglementations européennes sur les données. Pour des architectures cloud plus avancées avec Terraform, consultez notre tutoriel Terraform AWS.

## Étape 10 : Configurer les Notifications et le Monitoring

Un pipeline CI/CD performant nécessite des notifications fiables pour réagir rapidement aux échecs. GitHub Actions offre plusieurs mécanismes pour alerter votre équipe en cas de problème.

```
name: CI avec Notifications
on:
  push:
    branches: ['main']
jobs:
  build-and-notify:
    runs-on: ubuntu-24.04
    steps:
      - uses: actions/checkout@v4
      - uses: actions/setup-node@v4
        with:
          node-version: '22'
          cache: 'npm'
      - run: npm ci
      - run: npm test
      - run: npm run build
  notify-success:
    needs: build-and-notify
    if: success()
    runs-on: ubuntu-24.04
    steps:
      - name: Notification Slack - Succès
        uses: slackapi/slack-github-action@v2
        with:
          webhook: ${{ secrets.SLACK_WEBHOOK_URL }}
          webhook-type: incoming-webhook
          payload: |
            {
              "text": "✅ Déploiement réussi",
              "blocks": [
                {
                  "type": "section",
                  "text": {
                    "type": "mrkdwn",
                    "text": "*Déploiement réussi* sur `${{ github.ref_name }}`\nCommit: `${{ github.sha }}` par ${{ github.actor }}\n<${{ github.server_url }}/${{ github.repository }}/actions/runs/${{ github.run_id }}|Voir les logs>"
                  }
                }
              ]
            }
  notify-failure:
    needs: build-and-notify
    if: failure()
    runs-on: ubuntu-24.04
    steps:
      - name: Notification Slack - Échec
        uses: slackapi/slack-github-action@v2
        with:
          webhook: ${{ secrets.SLACK_WEBHOOK_URL }}
          webhook-type: incoming-webhook
          payload: |
            {
              "text": "❌ Échec du build sur ${{ github.ref_name }}"
            }
```
**Sortie attendue** : un message Slack formaté apparaît dans votre canal avec le statut du build, le nom de la branche, l'auteur du commit et un lien direct vers les logs. La condition `if: success()` ou `if: failure()` garantit que seule la notification appropriée est envoyée.

Au-delà de Slack, vous pouvez également configurer des notifications par email via les paramètres GitHub, utiliser des webhooks personnalisés ou intégrer des outils comme PagerDuty pour les alertes critiques en production. La clé est de trouver le bon équilibre entre être informé et ne pas être submergé par les notifications.

## Étape 11 : Optimiser les Performances avec les Runners Arm64 et GPU

Depuis le quatrième trimestre 2025, les runners Arm64 sont en disponibilité générale sur GitHub Actions. Ils offrent des performances 20 % supérieures pour certaines charges de travail et un coût identique aux runners Linux standard (0,008 $ par minute). Les runners GPU avec NVIDIA A100 (40 Go) ou H100 (80 Go) sont également disponibles depuis mars 2026 pour les workloads de machine learning. Côté réseau privé, CodeZine rapporte qu'en avril 2026 les hosted runners avec Azure private networking ont gagné le failover VNET en preview publique, une avancée qui renforce la résilience des architectures d'entreprise fortement isolées.

| Type de Runner | vCPU / RAM | Prix/minute (USD) | Cas d'Usage | Disponibilité | 
|---|---|---|---|---|
| Linux standard | 2 / 7 Go | 0,008 $ | Builds Node.js, Python, Go | GA | 
| Linux Arm64 | 4 / 16 Go | 0,008 $ | Builds Docker multi-arch, Graviton | GA (Q4 2025) | 
| Windows standard | 2 / 7 Go | 0,016 $ | Builds .NET, applications Windows | GA | 
| macOS standard | 3 / 14 Go | 0,08 $ | Builds iOS, applications macOS | GA | 
| Linux large | 64 / 256 Go | 0,064 $ | Builds monorepo, compilations lourdes | GA (2025) | 
| GPU (A100/H100) | Variable | 0,08 $+ | ML/IA, entraînement de modèles | GA (Mars 2026) | 

Pour utiliser un runner Arm64, il suffit de changer la valeur de `runs-on` à `ubuntu-24.04-arm64`. Cette simple modification peut réduire significativement le temps de build pour les projets qui ciblent des architectures ARM, comme les déploiements sur AWS Graviton ou les applications Docker multi-architecture.

Les runners GPU sont particulièrement intéressants pour les équipes qui développent des modèles d'IA. Avec CUDA 12.4 préinstallé et les frameworks TensorFlow et PyTorch disponibles, vous pouvez exécuter des tests d'entraînement et de validation directement dans votre pipeline CI/CD. Pour les développeurs qui travaillent avec des outils de code IA, notre comparatif GitHub Copilot vs Cursor peut vous aider à choisir le meilleur assistant.

## Étape 12 : Projet Complet — Application Node.js avec Pipeline CI/CD de Production

Rassemblons toutes les étapes précédentes dans un projet complet et fonctionnel. Voici la structure d'une application Node.js avec un pipeline CI/CD GitHub Actions prêt pour la production :

