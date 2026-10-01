---
id: collect-261001-general-networking/general-networking/tutoriel-playwright-1-59-tests-e2e-en-13-etapes-2026-2
title: "Choisir : TypeScript, dossier ./tests, ajouter GitHub Actions = yes, installer navigateurs = yes"
domain: general-networking
role: reference
task: reference
actors: ["Apple", "Stripe"]
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/tutoriel-playwright-1-59-tests-e2e-en-13-etapes-2026.md
source_anchor: ""
source_lines: [38, 149]
sha256: f2a741e136e117c12e13d6c79c8fb50e586338b54658affb0325fe3e0c906fa8
---

# Choisir : TypeScript, dossier ./tests, ajouter GitHub Actions = yes, installer navigateurs = yes

```
mkdir tutoriel-playwright-2026 && cd tutoriel-playwright-2026
npm init -y
npm init playwright@latest
# Choisir : TypeScript, dossier ./tests, ajouter GitHub Actions = yes, installer navigateurs = yes
npx playwright --version
# Output: Version 1.59.1
ls -la
# .github/workflows/playwright.yml
# tests/example.spec.ts
# tests-examples/demo-todo-app.spec.ts
# playwright.config.ts
# package.json
```
L’installation télécharge environ 580 Mo de binaires. Sur une connexion fibre, comptez 90 secondes. Sur une connexion entreprise filtrée, configurez le proxy via les variables `HTTPS_PROXY` et `PLAYWRIGHT_DOWNLOAD_HOST`. La commande `npx playwright install --with-deps` installe également les dépendances système sur Linux (libxkbcommon, libnss3, libgbm1), étape souvent oubliée qui produit l’erreur classique « Host system is missing dependencies to run browsers ».

## Étape 2 : Configurer playwright.config.ts pour la production

Le fichier `playwright.config.ts` orchestre l’ensemble du comportement de la suite. Une configuration de débutant tient en 10 lignes, mais une configuration prête pour la production exige des choix explicites sur les retries, les timeouts, le reporting, la base URL et les projets multi-navigateurs. Voici un fichier de référence pour une équipe française en 2026, optimisé pour CI parallèle et debugging local efficace.

```
// playwright.config.ts
import { defineConfig, devices } from '@playwright/test';
export default defineConfig({
  testDir: './tests',
  fullyParallel: true,
  forbidOnly: !!process.env.CI,
  retries: process.env.CI ? 2 : 0,
  workers: process.env.CI ? 4 : undefined,
  reporter: [
    ['html', { open: 'never' }],
    ['list'],
    ['junit', { outputFile: 'results.xml' }],
  ],
  use: {
    baseURL: process.env.BASE_URL || 'http://localhost:3000',
    trace: 'retain-on-failure-and-retries',
    screenshot: 'only-on-failure',
    video: 'retain-on-failure',
    actionTimeout: 10_000,
    navigationTimeout: 30_000,
    locale: 'fr-FR',
    timezoneId: 'Europe/Paris',
  },
  projects: [
    { name: 'chromium', use: { ...devices['Desktop Chrome'] } },
    { name: 'firefox',  use: { ...devices['Desktop Firefox'] } },
    { name: 'webkit',   use: { ...devices['Desktop Safari'] } },
    { name: 'iphone-15', use: { ...devices['iPhone 15'] } },
    { name: 'pixel-8',   use: { ...devices['Pixel 7'] } },
  ],
  webServer: {
    command: 'npm run dev',
    url: 'http://localhost:3000',
    reuseExistingServer: !process.env.CI,
    timeout: 120_000,
  },
});
```
Trois choix méritent une explication. D’abord, `trace: 'retain-on-failure-and-retries'` est la nouvelle option introduite dans Playwright 1.59 : elle conserve toutes les traces des tentatives ratées, ce qui simplifie l’analyse d’un test flaky qui finit par passer au 3e essai. Ensuite, `locale: 'fr-FR'` et `timezoneId: 'Europe/Paris'` sont essentiels pour les tests de paiement Stripe ou Lemon Squeezy qui adaptent leur interface selon la locale. Enfin, `webServer` démarre automatiquement votre application avant les tests et la coupe à la fin, supprimant le besoin de jongler avec Docker Compose en local.

## Étape 3 : Écrire votre premier test E2E avec les locators web-first

Playwright recommande depuis la version 1.27 l’usage exclusif de locators basés sur les rôles ARIA et les attributs accessibles. Cette approche présente deux avantages : les sélecteurs survivent aux refontes CSS et ils valident simultanément l’accessibilité de l’application. La nouvelle méthode `locator.normalize()` de Playwright 1.59 réécrit automatiquement les locators vers ces patterns recommandés, ce qui facilite la migration d’anciens tests Selenium ou Cypress.

```
// tests/login.spec.ts
import { test, expect } from '@playwright/test';
test.describe('Authentification utilisateur', () => {
  test.beforeEach(async ({ page }) => {
    await page.goto('/connexion');
  });
  test('connexion reussie avec identifiants valides', async ({ page }) => {
    await page.getByLabel('Adresse e-mail').fill('[email protected]');
    await page.getByLabel('Mot de passe').fill('MotDePasseSecur1!');
    await page.getByRole('button', { name: 'Se connecter' }).click();
    await expect(page).toHaveURL(/\/tableau-de-bord/);
    await expect(page.getByRole('heading', { level: 1 })).toHaveText('Bienvenue');
    await expect(page.getByTestId('user-menu')).toBeVisible();
  });
  test('message d erreur sur identifiants invalides', async ({ page }) => {
    await page.getByLabel('Adresse e-mail').fill('[email protected]');
    await page.getByLabel('Mot de passe').fill('mauvais');
    await page.getByRole('button', { name: 'Se connecter' }).click();
    const alert = page.getByRole('alert');
    await expect(alert).toContainText('Identifiants invalides');
    await expect(page).toHaveURL(/\/connexion/);
  });
});
```
Exécutez ce test avec `npx playwright test tests/login.spec.ts --headed` pour voir le navigateur en action. Notez l’absence totale de `waitFor` ou `sleep` : Playwright attend automatiquement que chaque élément soit actionable avant de cliquer ou de remplir, et chaque assertion `expect()` est web-first, c’est-à-dire qu’elle retentent jusqu’au timeout (5 secondes par défaut). Cette différence avec Selenium élimine 80 % des causes classiques de tests flaky.

## Étape 4 : Générer des tests avec Codegen et l’enregistreur

L’outil Codegen est probablement la fonctionnalité la plus sous-estimée de Playwright. Lancez `npx playwright codegen http://localhost:3000` et une fenêtre Chromium s’ouvre, accompagnée d’un Inspector. Toutes vos interactions sont enregistrées sous forme de code TypeScript prêt à coller dans un fichier de test. Playwright 1.59 a amélioré le générateur : il insère désormais automatiquement les assertions `toBeVisible()` appropriées et propose les locators `getByRole` en priorité.

```
# Lancer Codegen sur une URL
npx playwright codegen https://votreapp.fr
# Lancer Codegen en mode mobile (iPhone 15)
npx playwright codegen --device="iPhone 15" https://votreapp.fr
# Lancer Codegen avec une locale et un timezone specifiques
npx playwright codegen --lang=fr-FR --timezone="Europe/Paris" https://votreapp.fr
# Enregistrer directement dans un fichier
npx playwright codegen -o tests/onboarding.spec.ts https://votreapp.fr
```
Le pattern recommandé en équipe est le suivant : un développeur enregistre un parcours utilisateur complet via Codegen, exporte vers un fichier `.spec.ts`, puis refactorise les assertions et extrait les Page Objects. Cette approche divise par trois le temps d’écriture initial d’une suite de tests. Pour un onboarding e-commerce avec 12 étapes (panier, paiement, confirmation), comptez 8 minutes d’enregistrement contre 45 minutes d’écriture manuelle.

## Étape 5 : Page Object Model et fixtures pour code réutilisable

Dès que votre suite dépasse 20 tests, l’inflation des sélecteurs devient un problème. Le Page Object Model encapsule les locators et les actions dans des classes dédiées, suivant le principe de responsabilité unique. Playwright propose en complément un système de **fixtures** typé qui dépasse de loin les fixtures Pytest ou Mocha : chaque test reçoit ses dépendances par injection, avec lifecycle automatique.

