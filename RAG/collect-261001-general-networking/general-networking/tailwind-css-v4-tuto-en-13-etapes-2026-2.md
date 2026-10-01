---
id: collect-261001-general-networking/general-networking/tailwind-css-v4-tuto-en-13-etapes-2026-2
title: "Sortie attendue : v20.x.x ou supérieur (v22.x.x recommandé)"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/tailwind-css-v4-tuto-en-13-etapes-2026.md
source_anchor: ""
source_lines: [71, 212]
sha256: 1442d4f41f3e09c86d27f39e4753997b8410cdfcf0b2b66140d5171d575c43bb
---

# Sortie attendue : v20.x.x ou supérieur (v22.x.x recommandé)

`npm install tailwindcss @tailwindcss/vite`
Notez la différence fondamentale avec la v3 : il n’y a **plus de commande** `npx tailwindcss init` pour générer un `tailwind.config.js`, et l’on n’installe ni `postcss` ni `autoprefixer` séparément. Le plugin Vite gère tout. Selon le contexte, Tailwind v4 propose désormais quatre points d’entrée officiels : `@tailwindcss/vite` pour Vite, `@tailwindcss/postcss` pour une chaîne PostCSS (utile avec Next.js), un plugin dédié pour **Webpack** introduit avec la version 4.2.0, publiée le 18 février 2026 selon InfoQ (aux côtés de quatre nouvelles palettes de couleurs), et `@tailwindcss/cli` pour une compilation en ligne de commande sans bundler.

| Intégration | Paquet | Cas d’usage | 
|---|---|---|
| Vite | `@tailwindcss/vite` | Vite, projets vanilla, React, Vue | 
| PostCSS | `@tailwindcss/postcss` | Next.js, Webpack, Laravel Mix | 
| CLI | `@tailwindcss/cli` | Build autonome, scripts, prototypes | 

## Étape 3 : Activer le plugin dans la configuration Vite

Créez (ou ouvrez) le fichier `vite.config.js` à la racine du projet et enregistrez le plugin Tailwind. C’est la seule configuration JavaScript que ce tutoriel requiert.

```
// vite.config.js
import { defineConfig } from 'vite'
import tailwindcss from '@tailwindcss/vite'
export default defineConfig({
  plugins: [
    tailwindcss(),
  ],
})
```
C’est tout. Le plugin scanne automatiquement vos fichiers `.html`, `.js`, `.jsx`, `.vue`, etc., détecte les classes Tailwind utilisées et génère uniquement le CSS nécessaire. Plus besoin de configurer manuellement le tableau `content` comme en v3 : la détection est désormais automatique et plus intelligente, ce qui élimine l’une des principales sources d’erreur de l’ancienne version.

## Étape 4 : Importer Tailwind dans votre feuille CSS

Ouvrez `src/style.css`, supprimez tout son contenu généré par Vite, et remplacez-le par une seule ligne. C’est ici que la philosophie CSS-first de la v4 prend tout son sens.

```
/* src/style.css */
@import "tailwindcss";
```
Cette unique directive `@import "tailwindcss"` remplace les trois anciennes lignes de la v3 (`@tailwind base; @tailwind components; @tailwind utilities;`). Elle charge le reset (Preflight), les composants et l’ensemble des utilitaires. Assurez-vous ensuite que ce fichier est bien importé dans votre point d’entrée JavaScript.

```
// src/main.js
import './style.css'
```
Lancez maintenant le serveur de développement pour vérifier que tout fonctionne.

```
npm run dev
# Sortie attendue :
  VITE v7.x.x  ready in 312 ms
  ->  Local:   http://localhost:5173/
  ->  Network: use --host to expose
  ->  press h + enter to show help
```
Ouvrez `http://localhost:5173/`. Le style par défaut du navigateur a changé (marges réinitialisées, police sans-serif) : c’est la preuve que le Preflight de Tailwind est actif. Si rien ne change, consultez la section dépannage en fin d’article.

## Étape 5 : Écrire votre premier composant avec les utilitaires

Remplacez le contenu du `<body>` de `index.html` par une structure de page d’accueil. Le principe de Tailwind est de composer le style directement dans le HTML via des classes utilitaires : `flex` pour la flexbox, `p-6` pour le padding, `text-xl` pour la taille de police, etc.

```
<!-- index.html (extrait du body) -->
<body class="bg-slate-50 text-slate-900 antialiased">
  <header class="border-b border-slate-200">
    <nav class="mx-auto flex max-w-5xl items-center justify-between p-6">
      <span class="text-lg font-bold tracking-tight">Studio Lumiere</span>
      <ul class="flex gap-6 text-sm font-medium">
        <li><a href="#" class="hover:text-indigo-600">Accueil</a></li>
        <li><a href="#" class="hover:text-indigo-600">Projets</a></li>
        <li><a href="#" class="hover:text-indigo-600">Contact</a></li>
      </ul>
    </nav>
  </header>
  <main class="mx-auto max-w-5xl px-6 py-20">
    <h1 class="text-4xl font-bold tracking-tight sm:text-6xl">
      Concevez plus vite avec Tailwind</h1>
    <p class="mt-6 max-w-2xl text-lg text-slate-600">
      Un studio de design qui transforme vos idees en interfaces
      performantes et accessibles.</p>
    <button class="mt-8 rounded-lg bg-indigo-600 px-5 py-3
      font-semibold text-white transition hover:bg-indigo-500">
      Demarrer un projet</button>
  </main>
</body>
```
Enregistrez le fichier : grâce au moteur Oxide et au HMR de Vite, la page se met à jour instantanément. Vous disposez déjà d’un en-tête, d’un titre responsive (le préfixe `sm:` agrandit la police sur écran moyen) et d’un bouton avec effet de survol. Chaque classe correspond à une propriété CSS – c’est la base du modèle utility-first.

## Étape 6 : Personnaliser le thème avec la directive @theme

Voici le cœur de la nouveauté v4. Pour personnaliser couleurs, polices, espacements ou points de rupture, on n’édite plus un fichier JavaScript : on déclare des variables CSS dans un bloc `@theme`. Ajoutez ce bloc dans `src/style.css`, juste après l’import.

```
/* src/style.css */
@import "tailwindcss";
@theme {
  --color-brand: #4f46e5;
  --color-brand-dark: #4338ca;
  --color-surface: #0f172a;
  --font-display: "Poppins", ui-sans-serif, system-ui, sans-serif;
  --spacing-section: 6rem;
  --breakpoint-3xl: 120rem;
}
```
Chaque variable génère automatiquement les utilitaires correspondants. `--color-brand` crée les classes `bg-brand`, `text-brand`, `border-brand`, etc. `--font-display` crée `font-display`. `--breakpoint-3xl` crée le préfixe responsive `3xl:`. Vous pouvez désormais remplacer `bg-indigo-600` par `bg-brand` dans votre bouton. Et puisque ces tokens sont de véritables variables CSS, vous pouvez aussi les consommer en CSS classique : `color: var(--color-brand);`.

### Une palette P3 et des variables accessibles partout

L’annonce de la v4 met en avant une palette de couleurs redessinée en espace **P3**, offrant des teintes plus vives sur les écrans modernes. Parce que tous les tokens sont exposés comme variables CSS natives, vous pouvez les manipuler dynamiquement en JavaScript ou les interpoler avec `color-mix()`, ce qui était impossible proprement quand la configuration vivait dans un fichier JS isolé du runtime.

## Étape 7 : Implémenter le mode sombre

Le mode sombre s’active avec le variant `dark:`. Par défaut, il suit la préférence système (`prefers-color-scheme`). Ajoutez les variantes sombres à vos éléments existants.

```
<body class="bg-slate-50 text-slate-900 antialiased
  dark:bg-slate-950 dark:text-slate-100">
<button class="mt-8 rounded-lg bg-brand px-5 py-3 font-semibold
  text-white transition hover:bg-brand-dark
  dark:bg-indigo-500 dark:hover:bg-indigo-400">
  Demarrer un projet</button>
```
Pour basculer le thème manuellement (avec un bouton plutôt que la préférence système), redéfinissez le variant `dark` en CSS pour qu’il réagisse à une classe sur l’élément `<html>`.

```
/* src/style.css */
@import "tailwindcss";
@custom-variant dark (&:where(.dark, .dark *));
```
Ajoutez ensuite un peu de JavaScript pour mémoriser le choix de l’utilisateur dans le `localStorage` – un comportement attendu sur tout site professionnel.

```
// src/main.js
import './style.css'
const root = document.documentElement
const stored = localStorage.getItem('theme')
if (stored === 'dark' ||
   (!stored && window.matchMedia('(prefers-color-scheme: dark)').matches)) {
  root.classList.add('dark')
}
document.querySelector('#theme-toggle')?.addEventListener('click', () => {
  root.classList.toggle('dark')
  localStorage.setItem('theme',
    root.classList.contains('dark') ? 'dark' : 'light')
})
```
## Étape 8 : Rendre la page responsive

