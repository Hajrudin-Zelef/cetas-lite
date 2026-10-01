---
id: collect-261001-general-networking/general-networking/tutoriel-playwright-1-59-tests-e2e-en-13-etapes-2026-3
title: "Choisir : TypeScript, dossier ./tests, ajouter GitHub Actions = yes, installer navigateurs = yes"
domain: general-networking
role: reference
task: reference
actors: ["Google", "Microsoft", "Stripe"]
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/tutoriel-playwright-1-59-tests-e2e-en-13-etapes-2026.md
source_anchor: ""
source_lines: [150, 308]
sha256: cb66cd050c827a5a5868af566daca12dfc139dbe67571b15c65540c775d74d17
---

# Choisir : TypeScript, dossier ./tests, ajouter GitHub Actions = yes, installer navigateurs = yes

```
// pages/LoginPage.ts
import { Page, Locator, expect } from '@playwright/test';
export class LoginPage {
  readonly page: Page;
  readonly emailInput: Locator;
  readonly passwordInput: Locator;
  readonly submitButton: Locator;
  readonly errorAlert: Locator;
  constructor(page: Page) {
    this.page = page;
    this.emailInput = page.getByLabel('Adresse e-mail');
    this.passwordInput = page.getByLabel('Mot de passe');
    this.submitButton = page.getByRole('button', { name: 'Se connecter' });
    this.errorAlert = page.getByRole('alert');
  }
  async goto() {
    await this.page.goto('/connexion');
  }
  async login(email: string, password: string) {
    await this.emailInput.fill(email);
    await this.passwordInput.fill(password);
    await this.submitButton.click();
  }
  async expectErrorMessage(message: string) {
    await expect(this.errorAlert).toContainText(message);
  }
}
// fixtures.ts
import { test as base } from '@playwright/test';
import { LoginPage } from './pages/LoginPage';
type MesFixtures = { loginPage: LoginPage };
export const test = base.extend<MesFixtures>({
  loginPage: async ({ page }, use) => {
    const lp = new LoginPage(page);
    await lp.goto();
    await use(lp);
  },
});
export { expect } from '@playwright/test';
```
Avec cette structure, un test devient extrêmement lisible : `await loginPage.login('[email protected]', 'pwd')`. Si demain la page de connexion change de markup, vous modifiez un seul fichier au lieu de 40 tests. Les fixtures supportent également le scope `worker` pour partager une ressource coûteuse (instance Postgres, container Redis) entre tous les tests d’un worker, exactement comme avec Pytest.

## Étape 6 : Authentification persistante avec storageState

Se reconnecter avant chaque test est lent et fragile. Le pattern recommandé par la documentation officielle Playwright Auth consiste à se connecter une fois dans un fichier de setup global, puis à sauvegarder l’état du navigateur (cookies, localStorage, sessionStorage) dans un fichier JSON. Tous les tests suivants chargent cet état et démarrent déjà authentifiés.

```
// auth.setup.ts
import { test as setup, expect } from '@playwright/test';
const authFile = 'playwright/.auth/user.json';
setup('authentification', async ({ page }) => {
  await page.goto('/connexion');
  await page.getByLabel('Adresse e-mail').fill(process.env.TEST_EMAIL!);
  await page.getByLabel('Mot de passe').fill(process.env.TEST_PASSWORD!);
  await page.getByRole('button', { name: 'Se connecter' }).click();
  await page.waitForURL('/tableau-de-bord');
  await expect(page.getByTestId('user-menu')).toBeVisible();
  await page.context().storageState({ path: authFile });
});
// playwright.config.ts (extrait)
export default defineConfig({
  projects: [
    { name: 'setup', testMatch: /.*\.setup\.ts/ },
    {
      name: 'chromium',
      use: { ...devices['Desktop Chrome'], storageState: 'playwright/.auth/user.json' },
      dependencies: ['setup'],
    },
  ],
});
```
Cette stratégie réduit le temps total d’une suite de 50 tests d’environ 6 minutes à 90 secondes. Pour les applications avec OAuth (Google, GitHub, Microsoft), Playwright supporte également l’injection directe de tokens via `page.context().addCookies()`, ce qui évite le passage par le flow OAuth complet. Pour les SSO d’entreprise complexes (SAML, OIDC avec Keycloak), reportez-vous à notre tutoriel Keycloak 26.6 SSO/OIDC qui couvre l’intégration côté serveur d’identité.

## Étape 7 : Tester l’API REST avec request fixture

Beaucoup d’équipes ne réalisent pas que Playwright n’est pas qu’un outil UI : la fixture `request` permet d’écrire des tests d’API REST sans démarrer aucun navigateur. C’est extrêmement rapide (10 à 30 ms par test) et particulièrement utile pour tester un backend Fastify, Express ou NestJS avant le déploiement. Les mêmes assertions et le même test runner sont utilisés, ce qui unifie la stack de tests.

```
// tests/api/articles.spec.ts
import { test, expect } from '@playwright/test';
const API_BASE = process.env.API_URL || 'http://localhost:4000';
test.describe('API REST Articles', () => {
  let articleId: string;
  let authToken: string;
  test.beforeAll(async ({ request }) => {
    const res = await request.post(`${API_BASE}/auth/login`, {
      data: { email: '[email protected]', password: 'admin' },
    });
    expect(res.status()).toBe(200);
    authToken = (await res.json()).token;
  });
  test('POST /articles cree un article', async ({ request }) => {
    const res = await request.post(`${API_BASE}/articles`, {
      headers: { Authorization: `Bearer ${authToken}` },
      data: { titre: 'Mon article 2026', contenu: 'Lorem ipsum...' },
    });
    expect(res.status()).toBe(201);
    const body = await res.json();
    expect(body).toHaveProperty('id');
    expect(body.titre).toBe('Mon article 2026');
    articleId = body.id;
  });
  test('GET /articles/:id renvoie l article', async ({ request }) => {
    const res = await request.get(`${API_BASE}/articles/${articleId}`);
    expect(res.status()).toBe(200);
    const body = await res.json();
    expect(body.titre).toBe('Mon article 2026');
  });
  test.afterAll(async ({ request }) => {
    await request.delete(`${API_BASE}/articles/${articleId}`, {
      headers: { Authorization: `Bearer ${authToken}` },
    });
  });
});
```
Pour un backend construit avec notre tutoriel Fastify 5.8 API REST TypeScript ou un NestJS 11 avec Prisma, vous pouvez exécuter la suite API sur chaque pull request en 12 secondes là où une suite UI complète prendrait 4 minutes. C’est l’approche recommandée pour les boucles de feedback rapides en développement.

## Étape 8 : Mocking réseau et interception de requêtes

Tester un parcours utilisateur qui dépend d’une API tierce (Stripe, SendGrid, Mapbox) sans appeler ces services en production est un défi classique. Playwright propose `page.route()` pour intercepter toute requête sortante et la remplacer par une réponse contrôlée. Cette technique transforme des tests E2E imprévisibles en tests parfaitement déterministes.

```
// tests/checkout.spec.ts
import { test, expect } from '@playwright/test';
test('paiement Stripe avec mock', async ({ page }) => {
  // Intercepter l'appel Stripe et renvoyer succes
  await page.route('**/v1/payment_intents**', async (route) => {
    await route.fulfill({
      status: 200,
      contentType: 'application/json',
      body: JSON.stringify({
        id: 'pi_3OxXxXxXxXxXxXxX',
        status: 'succeeded',
        amount: 4999,
        currency: 'eur',
      }),
    });
  });
  // Bloquer Google Analytics pour des tests propres
  await page.route('**/google-analytics.com/**', (route) => route.abort());
  await page.goto('/panier');
  await page.getByRole('button', { name: 'Payer 49,99 EUR' }).click();
  await expect(page.getByText('Paiement confirme')).toBeVisible();
});
// Simuler une erreur reseau
test('gestion erreur API', async ({ page }) => {
  await page.route('**/api/checkout**', (route) =>
    route.fulfill({ status: 503, body: 'Service Unavailable' })
  );
  await page.goto('/panier');
  await page.getByRole('button', { name: 'Payer' }).click();
  await expect(page.getByRole('alert')).toContainText('Service temporairement indisponible');
});
```
Playwright 1.59 conserve les requêtes interceptées dans le Trace Viewer avec un nouveau toggle pretty-print pour les bodies JSON et form-data. Vous pouvez désormais débugger un mock défaillant en visualisant la requête réelle et la réponse renvoyée côte à côte, sans recourir à `console.log`. Pour les payloads volumineux (GraphQL, JSON-RPC), cette amélioration divise le temps d’investigation par cinq.

## Étape 9 : Tests de régression visuelle avec toHaveScreenshot

