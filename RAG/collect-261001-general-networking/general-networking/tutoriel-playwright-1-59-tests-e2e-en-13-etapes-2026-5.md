---
id: collect-261001-general-networking/general-networking/tutoriel-playwright-1-59-tests-e2e-en-13-etapes-2026-5
title: "Choisir : TypeScript, dossier ./tests, ajouter GitHub Actions = yes, installer navigateurs = yes"
domain: general-networking
role: reference
task: reference
actors: ["Microsoft"]
dates: []
keywords: ["apache"]
source: docs/RAG/collect-261001-general-networking/tutoriel-playwright-1-59-tests-e2e-en-13-etapes-2026.md
source_anchor: ""
source_lines: [449, 530]
sha256: 0361c1aba343ca6a3d22a6e46a42bd9f4e9e533521ce847f8a393930240de98b
---

# Choisir : TypeScript, dossier ./tests, ajouter GitHub Actions = yes, installer navigateurs = yes

```
// tests/Button.spec.tsx
import { test, expect } from '@playwright/experimental-ct-react';
import { Button } from '../src/components/Button';
test('rendu du bouton CTA', async ({ mount }) => {
  const component = await mount(
    <Button variant="primary" onClick={() => console.log('clicked')}>
      Commencer maintenant
    </Button>
  );
  await expect(component).toHaveText('Commencer maintenant');
  await expect(component).toHaveCSS('background-color', 'rgb(37, 99, 235)');
});
test('bouton desactive', async ({ mount }) => {
  const component = await mount(<Button disabled>Indisponible</Button>);
  await expect(component).toBeDisabled();
  await component.click({ force: true });
  // Aucune action ne devrait etre declenchee
});
```
Le component testing reste marqué experimental en 1.59, mais il est utilisé en production par GitHub, Microsoft et Shopify. Sa stabilisation est annoncée pour Playwright 2.0 prévu fin 2026. Pour tester des composants plus simples sans navigateur, restez sur Vitest 4.1 qui reste deux à trois fois plus rapide pour les tests unitaires purs.

## Playwright vs Cypress vs Selenium : tableau comparatif 2026

Le marché du test E2E compte trois acteurs majeurs en 2026. Voici une comparaison factuelle basée sur les versions stables actuelles et les métriques publiques GitHub, npm et State of JavaScript Survey 2025.

| Critère | Playwright 1.59 | Cypress 14.x | Selenium 4.x | 
|---|---|---|---|
| Étoiles GitHub (avril 2026) | 89 635 | ~47 000 | ~30 000 | 
| Navigateurs supportés | Chromium, Firefox, WebKit | Chromium, Firefox, Edge, Electron | Tous (via WebDriver) | 
| Test runner natif | Oui (@playwright/test) | Oui (intégré) | Non (Mocha, Pytest…) | 
| Auto-waiting | Oui (web-first) | Oui (limité) | Non (waits manuels) | 
| Parallélisation locale | Native, illimitée | Payante (Cypress Cloud) | Via Grid manuel | 
| API testing | Oui (request fixture) | Oui (cy.request) | Non natif | 
| Mobile emulation | 100+ devices preset | Viewport seulement | Via Appium | 
| Component testing | React, Vue, Svelte, Solid | React, Vue, Angular, Svelte | Non | 
| Trace/replay viewer | Oui (Trace Viewer) | Oui (Test Replay payant) | Non | 
| Licence | Apache 2.0 | MIT (cloud payant) | Apache 2.0 | 
| Coût annuel équipe 10 dev | 0 EUR | ~9 000 EUR (Cloud) | 0 EUR (auto-hébergé) | 
| Vitesse run 100 tests | ~45 sec | ~110 sec | ~180 sec | 
| Satisfaction StateOfJS 2025 | 88 % | 76 % | 54 % | 

Playwright domine sur le rapport vitesse/coût/cross-browser. Cypress conserve un avantage marketing et un écosystème de plugins plus mature dans certaines niches (testing-library, accessibilité), mais l’écart se resserre depuis 2024. Selenium reste pertinent pour les contextes legacy ou les besoins multi-langages (Java, C#, Ruby) où Playwright officiel n’existe qu’en JavaScript, Python, .NET et Java.

## Pièges courants et erreurs à éviter

Voici les 8 erreurs les plus fréquentes rencontrées en mission audit chez des équipes françaises en 2026. Les éviter dès la conception de votre suite épargne des semaines de refactoring.

- **Sélecteurs CSS fragiles** : utiliser`.btn-primary` ou`div > div:nth-child(3)` au lieu de`getByRole('button', { name: '...' })` . Solution : exécuter`locator.normalize()` sur l’ancien code et migrer progressivement.
- **Sleep et waitForTimeout** : ces appels créent des tests lents et flaky. Remplacer par`expect(locator).toBeVisible()` qui retentent jusqu’au timeout.
- **Tests dépendants entre eux** : un test qui dépend de l’état laissé par le précédent casse en parallèle. Chaque test doit créer ses propres données et les nettoyer.
- **Storage state non rafraîchi** : si votre token JWT expire en 1 heure et que vous régénérez`user.json` quotidiennement, vos tests échouent le matin. Régénérez avant chaque run CI.
- **Baselines visuelles cross-OS** : une baseline macOS échoue systématiquement en CI Linux à cause du rendu de polices. Régénérez en Docker Linux.
- **Trop de retries** :`retries: 5` masque les vrais bugs. Garder à 2 maximum et investiguer chaque flake.
- **Timeouts insuffisants en CI** : un test qui passe en 2 sec local peut prendre 8 sec en CI partagée. Augmenter`actionTimeout` à 15 secondes pour CI.
- **Pas de cleanup en afterAll** : créer 1 000 articles de test sans les supprimer pollue la base de données et finit par casser le seed.

## Dépannage : 8 erreurs Playwright fréquentes et solutions

| Erreur | Cause probable | Solution | 
|---|---|---|
| `Host system is missing dependencies` | Bibliothèques Linux manquantes | `npx playwright install --with-deps` | 
| `Test timeout of 30000ms exceeded` | Auto-wait dépasse le délai | Augmenter `timeout` dans la config ou via`test.setTimeout()` | 
| `Browser closed unexpectedly` | OOM kill, RAM insuffisante en CI | Passer à un runner 4 vCPU/16 Go ou réduire workers | 
| `strict mode violation` | Locator matche plusieurs éléments | Préciser avec `.first()` ,`.nth(0)` ou meilleur sélecteur | 
| `page.goto: net::ERR_CONNECTION_REFUSED` | webServer pas encore démarré | Augmenter `webServer.timeout` à 180 sec | 
| `expect(received).toHaveScreenshot` failed | Baseline différente (OS, polices) | Régénérer en Docker Linux : `--update-snapshots` | 
| `Cannot find module @playwright/test` | Installation incomplète | `rm -rf node_modules && npm ci` | 
| `TimeoutError: locator.click` | Élément masqué par overlay/modal | Fermer modal d’abord ou utiliser `force: true` | 

## Astuces avancées pour 2026 : sharding intelligent, traces sélectives

Trois techniques permettent de pousser Playwright au-delà du tutoriel standard. La première est le **sharding intelligent** : au lieu de répartir aléatoirement les tests, utilisez `--shard` en combinaison avec un fichier de durées historiques pour équilibrer les runners. Un test de 90 secondes ne doit pas se retrouver isolé sur un runner pendant que 3 autres terminent en 20 secondes. Le plugin communautaire `playwright-balance` automatise cet équilibrage à partir de `results.json` de la run précédente.

La deuxième technique est l’usage de **tags dans les noms de tests** pour exécution sélective. Marquez vos tests critiques avec `@smoke` et lancez uniquement la suite smoke avant chaque déploiement : `npx playwright test --grep @smoke`. La suite complète tourne la nuit. Cette segmentation divise par 5 le temps de feedback sur les pull requests sans sacrifier la couverture nocturne.

La troisième est l’**intégration avec Docker Compose** pour reproduire l’environnement de production en CI. Démarrez PostgreSQL, Redis, votre backend et un mailcatcher (MailHog) via `docker compose up` avant les tests E2E. Cette approche, détaillée dans notre tutoriel Docker Compose Stack Production, permet de tester les flows complets (inscription, e-mail de confirmation, première connexion) sans dépendre de services externes payants.

## Exemples de sortie : rapport HTML et JUnit

Voici à quoi ressemble une exécution complète de la suite de ce tutoriel sur un projet Next.js 15 hébergé en local. Vous obtenez trois sorties simultanées grâce à la configuration du reporter dans `playwright.config.ts`.

