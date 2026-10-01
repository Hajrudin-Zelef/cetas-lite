---
id: collect-261001-general-networking/general-networking/tutoriel-playwright-1-59-tests-e2e-en-13-etapes-2026-6
title: "Choisir : TypeScript, dossier ./tests, ajouter GitHub Actions = yes, installer navigateurs = yes"
domain: general-networking
role: reference
task: reference
actors: ["AWS", "Apple", "Microsoft", "Stripe"]
dates: []
keywords: ["aws", "graviton", "open source"]
source: docs/RAG/collect-261001-general-networking/tutoriel-playwright-1-59-tests-e2e-en-13-etapes-2026.md
source_anchor: ""
source_lines: [531, 640]
sha256: 2eaac7efc4c225493090e25ed2b08fc9ecbe92092086e2590d49d29adaa7c8bc
---

# Choisir : TypeScript, dossier ./tests, ajouter GitHub Actions = yes, installer navigateurs = yes

```
$ npx playwright test
Running 47 tests using 4 workers
  ok  1 [chromium] tests/login.spec.ts:8:3 connexion reussie (1.2s)
  ok  2 [chromium] tests/login.spec.ts:22:3 message d erreur (0.8s)
  ok  3 [chromium] tests/checkout.spec.ts:5:3 paiement Stripe (2.1s)
  ok  4 [firefox]  tests/login.spec.ts:8:3 connexion reussie (1.4s)
  ok  5 [webkit]   tests/login.spec.ts:8:3 connexion reussie (1.9s)
  KO  6 [chromium] tests/visuel.spec.ts:8:3 regression visuelle (3.2s)
       Snapshot comparison failed:
         50 pixels (ratio 0.001 of all image pixels) are different.
       Expected: tests/visuel.spec.ts-snapshots/homepage-chromium-linux.png
       Received: test-results/visuel-spec-ts-regression-visuelle/homepage-actual.png
       Diff:     test-results/visuel-spec-ts-regression-visuelle/homepage-diff.png
  46 passed, 1 failed (1m 18s)
To open last HTML report run:
  npx playwright show-report
```
Le rapport HTML auto-hébergé (`npx playwright show-report`) ouvre une interface web sur le port 9323 avec filtres par projet, statut, durée, tags. Chaque test échoué offre un lien vers sa trace, ses captures, son log console et sa vidéo enregistrée. Le rapport JUnit (`results.xml`) est consommé directement par GitLab, Jenkins ou Azure DevOps pour afficher les résultats dans l’UI native du serveur CI.

## Projet complet : structure de fichiers et arborescence finale

Au terme des 13 étapes, votre projet doit ressembler à l’arborescence suivante. Cette structure est utilisée en production par plusieurs équipes françaises de la French Tech et reste maintenable même au-delà de 500 tests.

```
tutoriel-playwright-2026/
.github/
  workflows/
    playwright.yml          # CI parallele sharding 4
playwright/
  .auth/
    user.json               # Storage state (gitignored)
pages/                      # Page Object Model
  LoginPage.ts
  DashboardPage.ts
  CheckoutPage.ts
fixtures/
  index.ts                  # Fixtures TypeScript typees
tests/
  auth.setup.ts             # Setup global d authentification
  login.spec.ts
  checkout.spec.ts
  visuel.spec.ts            # Regression visuelle
  mobile.spec.ts            # iPhone, Pixel
  api/
    articles.spec.ts        # Tests API REST
    users.spec.ts
  components/
    Button.spec.tsx         # Component testing
test-results/               # Gitignored
playwright-report/          # Gitignored
playwright.config.ts
package.json
tsconfig.json
.gitignore
```
Le `package.json` final exporte cinq scripts clés : `test` (suite complète), `test:smoke` (tag @smoke uniquement), `test:ui` (mode UI interactif), `test:update-snapshots` (régénérer baselines visuelles), `test:report` (ouvrir le dernier rapport HTML). Cette convention de nommage est devenue standard dans la communauté Playwright française et facilite l’onboarding des nouveaux développeurs.

## Questions fréquentes sur Playwright en 2026

### Playwright remplace-t-il complètement Cypress en 2026 ?

Pour les nouveaux projets, oui dans la majorité des cas. Playwright offre des fonctionnalités supérieures (WebKit, sharding gratuit, API testing native) sans paywall. Cypress garde un avantage pour les équipes déjà investies dans son écosystème ou utilisant Cypress Component Testing avec Angular, où Playwright n’a pas encore d’équivalent direct.

### Quelle version de Node.js choisir pour Playwright 1.59 ?

Node.js 20 LTS est l’option recommandée en avril 2026. Node.js 18 reste supporté mais arrivera en fin de vie en avril 2025 (déjà dépassé). Node.js 22 fonctionne aussi mais reste « current » jusqu’à octobre 2026. Pour la stabilité CI, restez sur 20.

### Combien de temps pour exécuter 200 tests Playwright sur GitHub Actions ?

Sans sharding, environ 8 à 12 minutes sur ubuntu-22.04. Avec sharding sur 4 runners, environ 2 minutes 15 secondes plus 30 secondes de merge des rapports. Le coût additionnel reste nul sur les minutes incluses GitHub gratuites pour les repos publics (et 0,008 $/min pour repos privés).

### Peut-on tester une PWA ou une app Electron avec Playwright ?

Oui. Pour les PWA, utilisez les configurations device standards et activez le mode offline via `context.setOffline(true)`. Pour Electron, Playwright propose `_electron` dans son API expérimentale qui pilote directement l’application packagée. C’est ce que VS Code utilise pour ses propres tests.

### Comment intégrer Playwright avec un backend Django ou Symfony ?

Playwright est agnostique du backend. Pour Django, démarrez votre serveur via la configuration `webServer` dans `playwright.config.ts` avec `command: 'python manage.py runserver'`. Pour Symfony, c’est `symfony serve --no-tls`. Notre tutoriel Django 5.2 LTS couvre l’intégration côté Python et notre guide API Platform 4.2 traite Symfony.

### Playwright fonctionne-t-il sur ARM (Apple Silicon, AWS Graviton) ?

Oui depuis la version 1.40. Les binaires Chromium et WebKit sont compilés en ARM64 natif pour macOS Apple Silicon et Linux ARM64. AWS Graviton fonctionne nativement, ce qui réduit le coût CI d’environ 20 % par rapport aux runners x86 chez plusieurs cloud providers.

### Quel est l’impact RGPD des tests Playwright en CI ?

Aucun si vous utilisez des données de test fictives et n’envoyez pas de traces vers un service tiers. Les traces Playwright restent dans GitHub Actions artifacts par défaut. Si vous activez un service comme Currents.dev ou DeploySentinel, vérifiez leur conformité RGPD et hébergement européen. Playwright lui-même ne télémetre rien depuis l’installation, sauf si vous activez explicitement `PLAYWRIGHT_TELEMETRY=1`.

### Comment migrer une suite Selenium existante vers Playwright ?

Procédez progressivement test par test, plutôt qu’en big bang. Playwright et Selenium peuvent coexister dans le même projet. Commencez par les nouveaux tests sur Playwright, puis migrez les tests les plus flaky de Selenium. La fonction `locator.normalize()` de Playwright 1.59 aide à convertir les sélecteurs XPath Selenium vers du `getByRole` recommandé.

## Pour aller plus loin

### Related Coverage

- Tutoriel Vitest 4.1 : Tests JavaScript en 13 Étapes — Complément idéal pour tests unitaires et composants en JSDOM
- Tutoriel Pytest Python 2026 : Tests Automatisés en 13 Étapes — L’équivalent Python pour vos backends FastAPI ou Django
- Tutoriel Selenium Python 2026 : Automatisation Web en 13 Étapes — Comparaison technique et migration progressive
- Tutoriel GitHub Actions 2026 : CI/CD Complet — Le pipeline parallélisé qui complète votre setup Playwright
- Tutoriel Next.js : App Full-Stack en 13 Étapes — Le framework idéal à tester avec Playwright en 2026
- Tutoriel Docker Compose 2026 : Stack Production — Reproduire votre environnement de production pour les tests
- Tutoriel Fastify 5.8 : API REST Node.js — Pour le backend que vous testez via Playwright API

## Conclusion : Playwright, le choix par défaut pour les tests E2E en 2026

En 13 étapes, vous avez construit une suite de tests E2E professionnelle couvrant authentification persistante, API REST, mocking réseau, régression visuelle, émulation mobile, accessibilité ARIA et CI parallélisée sur 4 shards. Avec Playwright 1.59 et son écosystème open source, vous obtenez l’équivalent d’un produit SaaS facturé 9 000 € par an, sans aucune dépendance externe ni paywall. Cette autonomie technique est précieuse en 2026 où les équipes produits cherchent à reprendre le contrôle de leur stack de tests face à la consolidation du marché du SaaS de testing.

Le prochain jalon majeur est Playwright 2.0 attendu fin 2026 avec stabilisation du component testing et nouvelle API d’inspection AI-assisté. En attendant, votre suite actuelle reste 100 % compatible avec les futures versions grâce à la politique de stabilité long-terme de Microsoft. Pour suivre les sorties, abonnez-vous au flux des releases GitHub Playwright et au changelog officiel.

