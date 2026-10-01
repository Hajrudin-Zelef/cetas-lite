---
id: collect-261001-general-networking/general-networking/tutoriel-playwright-1-59-tests-e2e-en-13-etapes-2026-1
title: "Choisir : TypeScript, dossier ./tests, ajouter GitHub Actions = yes, installer navigateurs = yes"
domain: general-networking
role: reference
task: reference
actors: ["Hugging Face", "Microsoft", "Stripe"]
dates: []
keywords: ["apache", "mai", "open source"]
source: docs/RAG/collect-261001-general-networking/tutoriel-playwright-1-59-tests-e2e-en-13-etapes-2026.md
source_anchor: ""
source_lines: [1, 37]
sha256: 30496bb32920c38e39e1e122602c02ef7ffd327da71a67cdf89e521159330224
---

# Choisir : TypeScript, dossier ./tests, ajouter GitHub Actions = yes, installer navigateurs = yes

**Publié le 27 avril 2026 · Niveau : Intermédiaire · Durée : 60 minutes**

Playwright s’impose en 2026 comme la référence des tests bout-en-bout (E2E) pour les applications web modernes. Après la sortie de **Playwright 1.59** début avril 2026, Microsoft a enchaîné les mises à jour jusqu’au patch **v1.62.1** publié le 30 juillet 2026 (relevé par DevUpdate.io), consolidant une avance technique déjà confortable face à Cypress et Selenium : pilotage cross-browser (Chromium 141, Firefox 142, WebKit 26), test runner natif `@playwright/test`, mode UI interactif, Trace Viewer, et désormais l’API `page.screencast` pour enregistrer et annoter des sessions complètes. Signe de cette accélération, les téléchargements hebdomadaires npm ont atteint 52 millions en mai 2026 selon TestDino, faisant de Playwright le framework de test le plus utilisé au monde. Ce tutoriel Playwright vous emmène de l’installation jusqu’au pipeline CI parallélisé, en 13 étapes structurées et testables localement.

Le projet GitHub microsoft/playwright, qui affichait déjà plus de 78 600 étoiles fin décembre 2025 selon Testdino, comptabilise désormais plus de **510 000 dépôts dépendants** sur GitHub au 30 juillet 2026 et dépasse aujourd’hui les **89 600 étoiles**, restant l’outil de test E2E le plus téléchargé sur npm. Concrètement, vous allez construire une suite de tests complète pour une application de démonstration : authentification, formulaires, API REST, captures de régression visuelle, mocking réseau, et reporting HTML. Le projet final est exécutable sur GitHub Actions avec sharding sur 4 runners, soit une suite qui passe de 8 minutes à moins de 2 minutes en parallèle.

## Pourquoi Playwright en 2026 : 9 900 recherches mensuelles et adoption industrielle

Le test E2E n’est plus optionnel. D’après les données de recherche DataForSEO pour la France (avril 2026), le mot-clé « playwright » génère **9 900 recherches mensuelles** avec une concurrence faible, contre 27 100 pour Cypress (concurrence élevée) et 1 000 pour Selenium. Cette tendance se confirme au niveau mondial : les téléchargements hebdomadaires npm du package ont oscillé entre 20 et 33 millions fin 2025 et début 2026 selon TestDino, avant de culminer à 52 millions dès mai 2026, soit une multiplication par plus de 100 en cinq ans. Cette croissance s’explique par trois facteurs structurels : la complexité croissante des SPA React et Vue 3, l’obligation réglementaire de tester l’accessibilité (directive EAA effective en juin 2025), et la migration massive vers le tooling JavaScript moderne avec Bun et Vite. Là où Selenium WebDriver exigeait des configurations laborieuses et où Cypress reste bloqué sur Chromium par défaut, Playwright propose une API unifiée pour Chromium, Firefox et WebKit, avec auto-waiting et assertions web-first.

L’adoption en entreprise progresse rapidement : Microsoft, GitHub, Stripe, Vercel, Shopify et Hugging Face utilisent Playwright en production, aux côtés de plus de 4 400 entreprises vérifiées recensées par TestDino en décembre 2025. Selon la State of JavaScript Survey 2025, Playwright a dépassé Cypress en taux de satisfaction (88 % contre 76 %), et son taux d’adoption chez les équipes QA a bondi de 12 % en 2023 à 45,1 % en 2025 selon Agamisoft (citant le State of Testing), un niveau confirmé à 45,1 % dès mai 2026 par TestDino. La courbe d’apprentissage est plus douce qu’avec Selenium : un développeur familier de JavaScript ou TypeScript devient productif en moins de trois jours. Mieux encore, le runner intégré `@playwright/test` élimine la dépendance à Jest ou Mocha pour orchestrer la parallélisation, les retries et le sharding.

Sur le plan économique, Playwright reste open source sous licence Apache 2.0, sans paywall ni dashboard SaaS imposé. Cypress Cloud commence à 75 $ par mois pour 100 000 tests enregistrés, tandis que BrowserStack Automate facture 169 $ par utilisateur. Playwright se contente d’un runner local, d’un rapport HTML auto-hébergé et de l’intégration native à GitHub Actions, Azure Pipelines, GitLab CI ou Jenkins. Pour un cabinet de conseil ou une équipe produit française, l’économie annuelle dépasse souvent les 8 000 € sur une suite de taille moyenne.

## Prérequis et versions exactes : Node.js 20, Playwright 1.59, navigateurs bundled

Avant de plonger dans le code, validez votre environnement. Playwright 1.59 (publiée le 1er avril 2026 selon les release notes officielles) a depuis été suivie par la v1.61.1, puis par le patch v1.62.1 le 30 juillet 2026 selon DevUpdate.io, portant le compteur à 165 releases cumulées répertoriées par Microsoft sur GitHub à cette date. Entre-temps, l’éditeur Checkly a lui-même embarqué Playwright v1.58.2 dans son Runtime 2026.04 publié le 31 mars 2026, signe que l’écosystème CI/CD suit désormais la cadence de release Microsoft de très près ; la version courante requiert Node.js 18 ou supérieur, mais nous recommandons fortement Node.js 20 LTS pour bénéficier du support natif de `--watch` et des performances V8 améliorées. Sur macOS, vérifiez que vous tournez sous macOS 13 (Ventura) minimum ; sur Linux, Ubuntu 22.04 ou Debian 12 sont les configurations testées par Microsoft.

| Composant | Version requise | Version recommandée 2026 | Commande de vérification | 
|---|---|---|---|
| Node.js | ≥ 18.0 | 20.18 LTS | `node --version` | 
| npm | ≥ 9 | 10.9 | `npm --version` | 
| Playwright | 1.59.1 | 1.59.1 (stable avril 2026) | `npx playwright --version` | 
| Chromium bundled | 141.0.7390 | 141.0.7390.37 | Installé via `npx playwright install` | 
| Firefox bundled | 142.0 | 142.0.1 | Idem | 
| WebKit bundled | 26.0 | 26.0 | Idem | 
| OS supportés | — | Ubuntu 22/24, macOS 13+, Windows 10+ | `uname -a` | 
| RAM minimum | 4 Go | 8 Go (16 Go pour CI parallèle) | `free -h` | 

Une remarque importante sur les binaires de navigateurs : Playwright télécharge ses propres builds, séparés des navigateurs installés sur la machine. Cette isolation garantit la reproductibilité entre développeurs et CI, mais consomme environ 580 Mo dans `~/.cache/ms-playwright`. Si vous travaillez en environnement contraint (laptop SSD 256 Go), prévoyez l’espace. Vous pouvez également pointer vers un Chrome existant via `channel: 'chrome'` dans la configuration, ce qui évite le téléchargement de Chromium.

Côté éditeur, l’extension VS Code « Playwright Test for VSCode » publiée par Microsoft est quasi indispensable. Elle ajoute l’exécution test-par-test depuis la gutter, le debugger intégré, le picker de sélecteurs et l’enregistrement Codegen sans quitter l’IDE. Elle est compatible avec WebStorm via le plugin Playwright officiel JetBrains.

## Étape 1 : Installation initiale et scaffolding du projet

Créez un nouveau dossier et initialisez le projet Playwright via la commande officielle d’installation. Cette commande génère la structure de répertoires, un fichier `playwright.config.ts` typé, deux tests d’exemple et le workflow GitHub Actions correspondant. Choisissez TypeScript lorsqu’elle vous le propose : la majorité de la documentation 2026 utilise TypeScript et l’auto-complétion sur les sélecteurs est précieuse.

