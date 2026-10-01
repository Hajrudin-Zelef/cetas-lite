---
id: collect-261001-general-networking/general-networking/tailwind-css-v4-tuto-en-13-etapes-2026-4
title: "Sortie attendue : v20.x.x ou supérieur (v22.x.x recommandé)"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/tailwind-css-v4-tuto-en-13-etapes-2026.md
source_anchor: ""
source_lines: [345, 407]
sha256: d0fccc3ef8454602c33f8b924609d59cca8791cb1d802e5e739223c5fc374623
---

# Sortie attendue : v20.x.x ou supérieur (v22.x.x recommandé)

- **Conserver les anciennes directives `@tailwind`.** En v4,`@tailwind base; @tailwind components; @tailwind utilities;` ne fonctionne plus. Utilisez l’unique`@import "tailwindcss";` . C’est l’erreur n°1 lors d’une migration.
- **Chercher un `tailwind.config.js`.** Il n’existe plus par défaut. Toute personnalisation passe par`@theme` dans le CSS. Si vous avez vraiment besoin d’un fichier JS (plugin tiers), chargez-le explicitement avec`@config "./tailwind.config.js";` .
- **Mélanger la chaîne PostCSS de la v3.** Ne gardez ni`postcss` , ni`autoprefixer` , ni`postcss-import` si vous utilisez`@tailwindcss/vite` . Ils entrent en conflit avec le plugin et provoquent des doublons de préfixes.
- **Cibler des navigateurs trop anciens.** La v4 exige Safari 16.4+, Chrome 111+ et Firefox 128+. Si votre audience utilise des navigateurs antérieurs, restez sur Tailwind 3.4 – la v4 ne propose pas de fallback automatique.
- **Abuser de `@apply`.** Recréer des dizaines de classes`.card` ,`.title` ,`.wrapper` avec`@apply` annule l’intérêt du modèle utility-first. Préférez des composants de framework et réservez`@apply` aux quelques motifs vraiment répétés.

## Dépannage : 8 problèmes fréquents et leurs solutions

Si quelque chose ne fonctionne pas, ce tableau couvre les huit symptômes les plus signalés par la communauté francophone et internationale.

| Symptôme | Cause probable | Solution | 
|---|---|---|
| Aucun style appliqué | `style.css` non importé dans`main.js` | Ajouter `import './style.css'` | 
| Les classes ne s’actualisent pas | Plugin Vite absent de `vite.config.js` | Vérifier `plugins: [tailwindcss()]` | 
| `Cannot find module @tailwindcss/vite` | Paquet non installé | `npm i @tailwindcss/vite` | 
| `@tailwind` sans effet | Syntaxe v3 obsolète | Remplacer par `@import "tailwindcss"` | 
| Classe personnalisée ignorée | Variable absente de `@theme` | Déclarer le token dans `@theme` | 
| Mode sombre ne bascule pas | `@custom-variant dark` manquant | Ajouter la déclaration + classe `dark` | 
| CSS énorme en production | Build non lancé / cache | `npm run build` et vider`dist/` | 
| Couleurs ternes | Navigateur sans support P3 | Mettre à jour le navigateur | 

En cas de doute, supprimez le dossier `node_modules` et le fichier `package-lock.json`, puis relancez `npm install`. Une grande partie des erreurs de Tailwind v4 vient de versions mixtes (un paquet en 3.x, un autre en 4.x) restées dans le cache. Le guide de migration officiel détaille par ailleurs chaque changement entre la v3 et la v4.

## Astuces avancées pour aller plus loin

Une fois les bases maîtrisées, plusieurs techniques font passer votre productivité au niveau supérieur. Première astuce : exploitez les **valeurs arbitraires** entre crochets, par exemple `top-[117px]` ou `bg-[#1da1f2]`, pour les cas où aucun utilitaire prédéfini ne convient – sans quitter le HTML. Tailwind les génère à la volée.

Deuxième astuce : combinez les variants. `dark:hover:lg:bg-brand` est parfaitement valide et applique le fond seulement en mode sombre, au survol, sur grand écran. Cette composabilité élimine quasiment tout besoin d’écrire des media queries manuelles. Troisième astuce : utilisez le variant `group` et `peer` pour styliser un élément en fonction de l’état d’un parent ou d’un frère – par exemple afficher un libellé quand un champ reçoit le focus, avec `peer-focus:text-brand`.

Enfin, pour les équipes : adoptez l’écosystème officiel. **Tailwind Plus** (anciennement Tailwind UI) fournit des centaines de composants prêts à l’emploi, et **Catalyst** propose un kit de composants React entièrement typé. Ces produits officiels, recensés sur le site de Tailwind, accélèrent considérablement la mise en place d’interfaces cohérentes tout en restant 100 % compatibles avec la configuration v4 décrite ici.

## Intégration avec Next.js, Laravel et les autres

Ce tutoriel utilise Vite, mais la v4 s’intègre à tout l’écosystème. Pour **Next.js**, qui repose sur une chaîne PostCSS, installez `@tailwindcss/postcss` et déclarez-le dans `postcss.config.mjs` au lieu du plugin Vite. Pour **Laravel**, le plugin Vite officiel s’utilise exactement comme dans ce guide puisque Laravel embarque Vite depuis plusieurs versions. Pour un site statique sans bundler, le paquet `@tailwindcss/cli` compile votre CSS en une commande.

```
# Chaine PostCSS (Next.js, Webpack)
npm install tailwindcss @tailwindcss/postcss
# CLI autonome (sites statiques, prototypes)
npm install tailwindcss @tailwindcss/cli
npx @tailwindcss/cli -i ./src/input.css -o ./dist/output.css --watch
```
Dans tous les cas, le fichier CSS reste identique : un `@import "tailwindcss";` suivi de votre bloc `@theme`. C’est l’un des grands atouts de la v4 – la logique de personnalisation est portable d’un environnement à l’autre. Que vous travailliez sur une application Next.js 16 ou une interface React 19.2, votre thème Tailwind se transpose sans modification.

## Tailwind v4 face au CSS classique et à Bootstrap

Beaucoup de développeurs français se demandent encore s’il faut adopter Tailwind ou rester sur une approche traditionnelle. La différence tient au modèle mental. Avec le CSS classique ou un préprocesseur comme Sass, vous écrivez des sélecteurs et des règles dans des fichiers séparés, puis vous jonglez entre HTML et CSS en inventant des noms de classes (le fameux problème du nommage BEM). Avec Tailwind, vous composez l’interface directement dans le balisage à l’aide d’utilitaires atomiques, sans changer de contexte et sans inventer de noms. Le CSS final ne contient que ce que vous utilisez réellement.

Face à Bootstrap, la distinction est encore plus nette. Bootstrap fournit des composants pré-stylés (boutons, cartes, modales) avec une apparence reconnaissable, ce qui accélère le prototypage mais uniformise les interfaces. Tailwind, lui, ne fournit aucun composant tout fait par défaut : il vous donne des briques de bas niveau pour construire *votre* design system. Résultat : des interfaces uniques, un poids CSS minimal et une maintenance simplifiée. Le tableau ci-dessous résume les principaux arbitrages.

| Critère | CSS / Sass classique | Bootstrap 5 | Tailwind CSS v4 | 
|---|---|---|---|
| Approche | Sélecteurs nommés | Composants prêts | Utilitaires atomiques | 
| Nommage de classes | Manuel (BEM) | Imposé | Aucun à inventer | 
| Personnalisation | Totale mais lente | Limitée (overrides) | Totale via `@theme` | 
| Poids CSS final | Croît avec le projet | ~25 ko+ (gzip) | Quelques ko (purge auto) | 
| Uniformité visuelle | Variable | Forte (reconnaissable) | Sur-mesure | 
| Courbe d’apprentissage | Faible | Faible | Moyenne (mémoriser les classes) | 

Le bon choix dépend du contexte. Pour un prototype jetable à livrer en une journée, Bootstrap reste pertinent. Pour une application maintenue sur le long terme, avec une identité visuelle propre et des contraintes de performance, Tailwind v4 est aujourd’hui le standard de l’industrie : selon l’enquête State of CSS 2025 relayée par AIBase News en janvier 2026, 51 % des développeurs déclarent l’utiliser, pour 75 millions de téléchargements mensuels, et W3Techs mesurait au 4 septembre 2026 une part de marché de 1,7 % des sites dont la technologie CSS est identifiée (0,3 % de l’ensemble du web), tandis que TechnologyChecker.io détectait déjà Tailwind CSS sur 2 787 808 domaines au 28 mars 2026 – une adoption confirmée par les enquêtes annuelles auprès des développeurs front-end et par l’intégration dans des frameworks majeurs.

## Optimiser performance et accessibilité en production

