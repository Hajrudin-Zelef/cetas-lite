---
id: collect-261001-general-networking/general-networking/tutoriel-playwright-1-59-tests-e2e-en-13-etapes-2026-4
title: "Choisir : TypeScript, dossier ./tests, ajouter GitHub Actions = yes, installer navigateurs = yes"
domain: general-networking
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: ["agent"]
source: docs/RAG/collect-261001-general-networking/tutoriel-playwright-1-59-tests-e2e-en-13-etapes-2026.md
source_anchor: ""
source_lines: [309, 448]
sha256: 16b918b25e6d070cdfaadf5c1796280823d20aa17c649266d330eb19dd873dd8
---

# Choisir : TypeScript, dossier ./tests, ajouter GitHub Actions = yes, installer navigateurs = yes

Un bouton dont la couleur passe accidentellement de bleu à gris, une typographie qui change de Inter à Times après un déploiement de polices, un layout qui casse sur Firefox uniquement : ces régressions visuelles échappent à toutes les assertions classiques. La méthode `toHaveScreenshot()` de Playwright capture une image, la compare à une baseline et signale toute différence supérieure à un seuil configurable.

```
// tests/visuel.spec.ts
import { test, expect } from '@playwright/test';
test('regression visuelle page d accueil', async ({ page }) => {
  await page.goto('/');
  // Masquer les elements dynamiques (compteurs, dates)
  await expect(page).toHaveScreenshot('homepage.png', {
    fullPage: true,
    mask: [page.getByTestId('compteur-temps-reel'), page.locator('time')],
    maxDiffPixels: 50,
    threshold: 0.2,
  });
});
test('composant bouton CTA', async ({ page }) => {
  await page.goto('/');
  const cta = page.getByRole('button', { name: 'Commencer gratuitement' });
  await expect(cta).toHaveScreenshot('cta-bouton.png', {
    animations: 'disabled',
  });
});
```
Lors de la première exécution (`--update-snapshots`), Playwright génère les images de référence dans `tests/visuel.spec.ts-snapshots/`. Versionnez-les avec Git et activez Git LFS si vous dépassez les 100 Mo cumulés. Pour éviter les faux positifs entre macOS et Linux (rendu de polices différent), vos baselines doivent provenir du même OS que la CI : exécutez `--update-snapshots` dans un container Docker Linux pour garantir la cohérence.

## Étape 10 : Trace Viewer et UI Mode pour le débogage

Quand un test échoue en CI sans raison apparente, le Trace Viewer est votre meilleur allié. Il enregistre chaque action, capture le DOM à chaque étape, journalise les requêtes réseau, les logs console et les sources des assertions. Avec le nouveau mode `retain-on-failure-and-retries` de Playwright 1.59, vous récupérez les traces de chaque tentative ratée, ce qui révèle les patterns de tests flaky.

```
# Executer avec traces forcees en local
npx playwright test --trace on
# Visualiser une trace generee
npx playwright show-trace test-results/login-spec-ts-chromium/trace.zip
# Mode UI interactif (recommande pour le developpement)
npx playwright test --ui
# Debug pas-a-pas avec inspecteur
PWDEBUG=1 npx playwright test tests/login.spec.ts
# Lancer un test specifique en headed mode avec slow motion
npx playwright test login.spec.ts --headed --slow-mo=500
```
L’UI Mode (`--ui`) est la fonctionnalité préférée des équipes Playwright depuis sa stabilisation en 1.42. Elle ouvre une fenêtre dédiée affichant tous vos tests, leur arbre d’exécution, leur historique de runs, leurs traces, et permet de relancer un test isolément en un clic. Combinée avec l’extension VS Code, l’UI Mode élimine quasi totalement le besoin de `console.log` dans les tests.

## Étape 11 : Émulation mobile, accessibilité et tests cross-browser

Playwright émule plus de 100 devices via `devices['iPhone 15']`, `devices['Pixel 7']`, ou `devices['Galaxy S9+']`. Chaque device configure user-agent, viewport, device-scale-factor, support tactile, géolocalisation et permissions. La nouveauté de Playwright 1.59 est `page.ariaSnapshot()` qui capture l’arbre d’accessibilité pour valider la conformité WCAG 2.2 et la directive EAA européenne entrée en vigueur en juin 2025.

```
// tests/mobile.spec.ts
import { test, expect, devices } from '@playwright/test';
test.use({ ...devices['iPhone 15'], locale: 'fr-FR' });
test('navigation mobile iPhone 15', async ({ page }) => {
  await page.goto('/');
  await expect(page.getByRole('button', { name: 'Menu' })).toBeVisible();
  await page.getByRole('button', { name: 'Menu' }).tap();
  await expect(page.getByRole('navigation')).toBeVisible();
});
// Accessibilite avec ariaSnapshot (Playwright 1.59+)
test('arbre d accessibilite conforme', async ({ page }) => {
  await page.goto('/');
  const snapshot = await page.locator('main').ariaSnapshot();
  expect(snapshot).toContain('heading "Accueil" [level=1]');
  expect(snapshot).toContain('button "Commencer gratuitement"');
});
// Geolocalisation simulee Paris
test('geolocalisation Paris', async ({ browser }) => {
  const context = await browser.newContext({
    geolocation: { latitude: 48.8566, longitude: 2.3522 },
    permissions: ['geolocation'],
    locale: 'fr-FR',
  });
  const page = await context.newPage();
  await page.goto('/agences-proches');
  await expect(page.getByText('Paris')).toBeVisible();
});
```
Pour l’audit d’accessibilité complet, combinez Playwright avec `@axe-core/playwright` qui exécute le moteur Axe sur chaque page testée et signale les violations WCAG. Une intégration standard ajoute 80 ms par test mais permet de bloquer un déploiement si le score d’accessibilité chute. Plusieurs sites publics français (impots.gouv.fr, ameli.fr) suivent désormais cette approche en production.

## Étape 12 : CI parallèle avec GitHub Actions et sharding sur 4 runners

Une suite de 200 tests E2E qui prend 8 minutes en local devient inutilisable sur chaque pull request si la CI ne parallélise pas. Le sharding officiel de Playwright divise la suite en N parties exécutées sur N runners distincts, puis fusionne les rapports HTML. Sur 4 runners GitHub-hosted, vous passez de 8 minutes à environ 2 minutes 15 secondes.

```
# .github/workflows/playwright.yml
name: Playwright Tests
on:
  push:
    branches: [main]
  pull_request:
    branches: [main]
jobs:
  test:
    timeout-minutes: 30
    runs-on: ubuntu-22.04
    strategy:
      fail-fast: false
      matrix:
        shard: [1, 2, 3, 4]
    steps:
      - uses: actions/checkout@v4
      - uses: actions/setup-node@v4
        with:
          node-version: 20
          cache: 'npm'
      - run: npm ci
      - run: npx playwright install --with-deps chromium
      - run: npx playwright test --shard=${{ matrix.shard }}/4
      - uses: actions/upload-artifact@v4
        if: always()
        with:
          name: blob-report-${{ matrix.shard }}
          path: blob-report
          retention-days: 7
  merge-reports:
    if: always()
    needs: test
    runs-on: ubuntu-22.04
    steps:
      - uses: actions/checkout@v4
      - uses: actions/setup-node@v4
        with: { node-version: 20, cache: 'npm' }
      - run: npm ci
      - uses: actions/download-artifact@v4
        with:
          path: all-blob-reports
          pattern: blob-report-*
          merge-multiple: true
      - run: npx playwright merge-reports --reporter html ./all-blob-reports
      - uses: actions/upload-artifact@v4
        with:
          name: playwright-html-report
          path: playwright-report
          retention-days: 14
```
L’astuce du `blob` reporter, introduit en Playwright 1.37, sérialise les résultats de chaque shard dans un format intermédiaire, puis le job `merge-reports` les agrège en un rapport HTML unique téléchargeable comme artifact. Pour un pipeline complet incluant lint, tests unitaires Vitest et tests E2E, consultez notre guide GitHub Actions CI/CD 2026 qui détaille les caches optimisés et les stratégies de réutilisation des workflows.

## Étape 13 : Component testing pour React et Vue 3

Au-delà du E2E classique, Playwright propose depuis la version 1.22 le component testing pour React, Vue, Svelte et Solid. Chaque composant est monté dans un vrai navigateur (pas dans JSDOM comme Vitest), ce qui élimine la classe entière de bugs liée aux divergences DOM. Pour une bibliothèque de design system, c’est souvent la couche de tests qui détecte les régressions visuelles avant la production.

