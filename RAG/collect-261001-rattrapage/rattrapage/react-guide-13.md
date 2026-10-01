---
id: collect-261001-rattrapage/rattrapage/react-guide-13
title: "Guide React — Le manuel complet"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/react_guide.md
source_anchor: ""
source_lines: [2866, 2899]
sha256: 493c7fae7a0c6db11782686f95706ae50dae05fe75e3c9f84e633df18d22e196
---

# Guide React — Le manuel complet

### Librairies incontournables pour un SI interne
| Besoin | Librairie | Pourquoi |
|---|---|---|
| Data fetching + cache | **TanStack Query** | Cache, retry, polling, déduplication — remplace 80 % des `useEffect` de fetch |
| Formulaires complexes | **React Hook Form** | Performant (non contrôlé), validation intégrée (Zod) |
| Graphiques | **Recharts** / **ECharts** | Courbes de charge, températures, historiques |
| Tableaux avancés | **TanStack Table** | Tri, filtres, pagination, virtualisation |
| État global | **Zustand** | Simple, sans boilerplate (alternative à Redux) |
| Dates | **date-fns** | Léger, immutable, tree-shakable |

### Sujets à creuser ensuite
- **TanStack Query** en détail : `useQuery`, `useMutation`, invalidation du cache.
- **Zustand** : store global en ~20 lignes pour les préférences / l'utilisateur connecté.
- **React Hook Form + Zod** : validation typée des formulaires.
- **Tests E2E** : Playwright pour les parcours critiques (créer un ticket de bout en bout).
- **PWA** : transformer le dashboard en app installable (mode hors-ligne du technicien sur site).
- **Accessibilité (a11y)** : ARIA, navigation clavier, contrastes — obligatoire pour un outil utilisé quotidiennement.
- **i18n** : `react-i18next` si l'outil doit exister en FR/EN.

### Ressources officielles
- https://react.dev — la documentation officielle (tutoriels, référence des hooks) — **le point de départ n°1**
- https://vite.dev — documentation Vite
- https://reactrouter.com — documentation React Router
- https://tanstack.com — TanStack Query / Table / Virtual

### Feuille de route conseillée (4 semaines, à temps partiel)
- **Semaine 1** : sections 1→18 de ce guide + mini-app "liste d'équipements" (useState, props, map/keys).
- **Semaine 2** : sections 19→30 + dashboard avec `usePolling` et Error Boundaries.
- **Semaine 3** : React Router + formulaire de ticket + tests Vitest sur les utilitaires.
- **Semaine 4** : TypeScript (convertir un composant), TanStack Query, build + déploiement Nginx.

---

*Fin du guide — bon code, et que vos dashboards restent verts. 🟢*
