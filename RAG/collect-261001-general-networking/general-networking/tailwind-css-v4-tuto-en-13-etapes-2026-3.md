---
id: collect-261001-general-networking/general-networking/tailwind-css-v4-tuto-en-13-etapes-2026-3
title: "Sortie attendue : v20.x.x ou supérieur (v22.x.x recommandé)"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["attention"]
source: docs/RAG/collect-261001-general-networking/tailwind-css-v4-tuto-en-13-etapes-2026.md
source_anchor: ""
source_lines: [213, 344]
sha256: 4aff106c9bdcf6186dab12ac3da9e252510ba60716611e7d2bdce661a2128466
---

# Sortie attendue : v20.x.x ou supérieur (v22.x.x recommandé)

Tailwind est mobile-first : une classe sans préfixe s’applique à toutes les tailles, et les préfixes `sm:`, `md:`, `lg:`, `xl:`, `2xl:` ne s’activent qu’à partir du point de rupture indiqué. Construisons une grille de cartes de projets qui passe d’une colonne sur mobile à trois colonnes sur grand écran.

```
<section class="mx-auto max-w-5xl px-6 py-16">
  <div class="grid grid-cols-1 gap-6 md:grid-cols-2 lg:grid-cols-3">
    <article class="rounded-xl border border-slate-200 p-6
      dark:border-slate-800">
      <h3 class="font-display text-xl font-semibold">Identite visuelle</h3>
      <p class="mt-2 text-slate-600 dark:text-slate-400">
        Logos, chartes et systemes de design complets.</p>
    </article>
    <!-- Dupliquez l'article pour les autres projets -->
  </div>
</section>
```
Redimensionnez la fenêtre du navigateur : la grille passe de 1 à 2 puis 3 colonnes. Les points de rupture par défaut sont `sm` (40rem / 640px), `md` (48rem / 768px), `lg` (64rem / 1024px), `xl` (80rem / 1280px) et `2xl` (96rem / 1536px). Vous pouvez en ajouter via `@theme`, comme le `3xl` défini à l’étape 6.

## Étape 9 : Utiliser les container queries natives

Les **container queries**, l’une des grandes nouveautés de la v4, sont désormais intégrées sans aucun plugin. Contrairement aux media queries qui réagissent à la taille de l’écran, elles réagissent à la taille du *conteneur parent* – idéal pour des composants vraiment réutilisables. Marquez un conteneur avec `@container`, puis stylez ses enfants avec le préfixe `@`.

```
<div class="@container">
  <div class="flex flex-col gap-4 @md:flex-row @md:items-center">
    <img src="/avatar.jpg" alt="" class="size-16 rounded-full" />
    <div>
      <p class="font-semibold">Camille Roux</p>
      <p class="text-sm text-slate-500">Directrice artistique</p>
    </div>
  </div>
</div>
```
Ici, `@md:flex-row` ne s’applique pas selon la largeur de l’écran mais selon celle du conteneur `@container`. Placez ce composant dans une barre latérale étroite ou dans une zone large : il s’adapte tout seul à l’espace disponible. C’est un changement profond pour la conception de bibliothèques de composants. Pour approfondir, la documentation MDN sur les container queries CSS détaille la spécification sous-jacente.

## Étape 10 : Factoriser des composants réutilisables

Répéter de longues listes de classes nuit à la lisibilité. Deux approches existent en v4. La première, recommandée par l’équipe Tailwind, consiste à utiliser des composants de votre framework (un composant `Button` en React ou Vue). La seconde, en CSS pur, s’appuie sur `@apply` à l’intérieur d’une couche de composants.

```
/* src/style.css */
@import "tailwindcss";
@layer components {
  .btn-primary {
    @apply rounded-lg bg-brand px-5 py-3 font-semibold text-white
           transition hover:bg-brand-dark
           focus:outline-none focus:ring-2 focus:ring-brand focus:ring-offset-2;
  }
}
```
Vous pouvez maintenant écrire `<button class="btn-primary">Envoyer</button>` partout. Attention toutefois : abuser de `@apply` recrée le problème que Tailwind cherche à résoudre (du CSS dispersé et difficile à maintenir). Réservez-le aux motifs vraiment récurrents comme les boutons, les champs de formulaire ou les badges.

### Exemple : un champ de formulaire accessible

```
@layer components {
  .input-field {
    @apply w-full rounded-md border border-slate-300 px-4 py-2 text-slate-900
           placeholder:text-slate-400
           focus:border-brand focus:ring-2 focus:ring-brand/30
           dark:border-slate-700 dark:bg-slate-900 dark:text-slate-100;
  }
}
```
L’utilitaire `ring-brand/30` illustre la syntaxe d’opacité moderne : le `/30` applique 30 % d’opacité à la couleur, en interne via `color-mix()`. C’est plus lisible et plus performant que les anciens utilitaires d’opacité séparés.

## Étape 11 : Ajouter transitions et transformations 3D

La v4 introduit des utilitaires de **transformation 3D** natifs, permettant de manipuler les éléments dans l’espace directement depuis le HTML. Combinez `perspective`, `rotate-x`, `rotate-y` et `transform-3d` pour un effet de carte qui pivote au survol.

```
<div class="perspective-distant">
  <article class="transform-3d transition-transform duration-500
    hover:rotate-y-12 hover:rotate-x-6
    rounded-xl bg-white p-8 shadow-lg dark:bg-slate-900">
    <h3 class="text-xl font-semibold">Carte interactive</h3>
    <p class="mt-2 text-slate-600 dark:text-slate-400">
      Survolez cette carte pour voir l'effet 3D.</p>
  </article>
</div>
```
Pour les animations d’apparition, combinez les utilitaires `transition`, `duration-*` et `ease-*`. Tailwind v4 a aussi élargi les utilitaires de dégradés (dégradés coniques et radiaux) et ajouté `color-mix()` au cœur du moteur, ce qui rend les transitions de couleur bien plus douces qu’auparavant.

## Étape 12 : Compiler pour la production

Tailwind ne génère que les classes réellement utilisées : votre CSS de production reste minuscule, souvent quelques kilo-octets après compression. Lancez le build Vite.

```
npm run build
# Sortie attendue :
vite v7.x.x building for production...
14 modules transformed.
dist/index.html                  1.85 kB | gzip:  0.74 kB
dist/assets/index-a1b2c3d4.css   9.42 kB | gzip:  2.61 kB
dist/assets/index-e5f6g7h8.js    1.12 kB | gzip:  0.63 kB
built in 487ms
```
Le fichier CSS final ne pèse ici que 9,42 ko (2,61 ko gzippé) car le moteur Oxide a éliminé toutes les classes inutilisées. Testez le résultat localement avec `npm run preview` avant de déployer le dossier `dist/` sur votre hébergeur (Vercel, Netlify, Cloudflare Pages, ou un serveur classique). Aucune étape de purge manuelle n’est nécessaire : c’est automatique.

## Étape 13 : Le projet complet assemblé

Voici la feuille de style finale qui regroupe thème, mode sombre et composants – le fichier `src/style.css` complet de notre projet de production.

```
/* src/style.css -- version finale */
@import "tailwindcss";
@custom-variant dark (&:where(.dark, .dark *));
@theme {
  --color-brand: #4f46e5;
  --color-brand-dark: #4338ca;
  --font-display: "Poppins", ui-sans-serif, system-ui, sans-serif;
  --spacing-section: 6rem;
  --breakpoint-3xl: 120rem;
}
@layer components {
  .btn-primary {
    @apply rounded-lg bg-brand px-5 py-3 font-semibold text-white
           transition hover:bg-brand-dark
           focus:outline-none focus:ring-2 focus:ring-brand focus:ring-offset-2;
  }
  .input-field {
    @apply w-full rounded-md border border-slate-300 px-4 py-2
           focus:border-brand focus:ring-2 focus:ring-brand/30
           dark:border-slate-700 dark:bg-slate-900 dark:text-slate-100;
  }
}
```
Avec ces trois fichiers – `vite.config.js`, `src/style.css` et `index.html` – vous disposez d’un projet Tailwind CSS v4 complet, responsive, accessible, avec mode sombre persistant, container queries et animations 3D. La structure est volontairement framework-agnostique : copiez la même configuration dans un projet React, Vue ou Svelte et tout fonctionne à l’identique.

## 5 pièges courants à éviter avec Tailwind v4

La migration vers la v4 et le modèle CSS-first introduisent des erreurs spécifiques que rencontrent la plupart des développeurs. Voici les cinq plus fréquentes et comment les contourner.

