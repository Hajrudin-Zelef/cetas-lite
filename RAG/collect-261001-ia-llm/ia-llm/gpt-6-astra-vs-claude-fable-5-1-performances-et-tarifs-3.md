---
id: collect-261001-ia-llm/ia-llm/gpt-6-astra-vs-claude-fable-5-1-performances-et-tarifs-3
title: "gpt-6-astra-vs-claude-fable-5-1-performances-et-tarifs"
domain: ia-llm
role: reference
task: reference
actors: ["Anthropic", "OpenAI"]
dates: []
keywords: ["astra", "claude", "gpt-6", "agent", "benchmark", "energy", "fable 5"]
source: docs/RAG/collect-261001-ia-llm/gpt-6-astra-vs-claude-fable-5-1-performances-et-tarifs.md
source_anchor: ""
source_lines: [160, 227]
sha256: 605df47e9186569712671603f7b7d73adfc7d2f369e2ac04ab91f8e29c0378ad
---

# gpt-6-astra-vs-claude-fable-5-1-performances-et-tarifs

Fable 5.1 n’est pas pénalisé par sa grille ici, mais par son volume de sortie. Artificial Analysis l’a mesuré avec environ 1,7x les jetons de sortie de Fable 5 à effort max, d’où un coût par tâche 20 % plus élevé que son prédécesseur malgré la baisse du cache. Sans cette baisse, on serait à environ 5,16 $ : la remise agit vraiment, mais pas assez.

Vous payez ce surcoût pour quelque chose. Fable 5.1 marque 5 points de plus sur le même indice, et 4 points de plus en xhigh. La question est de savoir si 5 points valent 2,25x : à vous de juger, mais l’idée que des prix catalogue identiques donnent des factures identiques ne résiste pas aux mesures.

Les légendes d’OpenAI vont dans le même sens, pour ce que vaut l’auto-déclaration d’un éditeur : coût API estimé par tâche environ 31 % sous Fable 5.1 sur Terminal-Bench Science 0.1, 63 % sous sur Terminal-Bench 4.0 et 86 % sous sur BenchCAD, mesurés aux réglages d’effort choisis.

## Comment GPT-6 Astra et Claude Fable 5.1 ont performé

Astra a légèrement remporté notre test globalement, bien que les deux modèles aient fourni une simulation fonctionnelle du premier coup.

### Le test

J’ai exécuté une tâche difficile contre les deux modèles dans des conditions identiques, plutôt que plusieurs tâches superficielles. Le test retenu ici : une simulation de physique from scratch dans un seul fichier HTML, choisie car elle éprouve exactement ce que les deux éditeurs disent avoir amélioré : la justesse soutenue sur une construction longue, sans bibliothèque d’appui.

Cette occurrence est une version rafraîchie du test. La variante publiée utilisait un conteneur carré en rotation ; ici, on remplace par un hexagone en rotation avec un obstacle central en contre-rotation et une masse proportionnelle à la taille, ce qui augmente la difficulté tout en gardant un confinement clair à juger. Voici l’invite, collée mot pour mot dans un nouveau chat pour chaque modèle :

```
Build a single-file HTML page (inline CSS and JS, canvas, no build step, no external libraries, no network) that simulates a few dozen balls of varying sizes bouncing under gravity inside a **slowly rotating hexagonal container**, with a **smaller counter-rotating obstacle at the centre** that the balls also collide with. 
Ball mass should scale with size, so larger balls shove smaller ones around. 
The balls should collide with each other, with the hexagon's walls, and with the central obstacle, and lose a little energy on each collision so the system settles rather than gaining energy over time. 
Both the hexagon and the inner obstacle keep rotating throughout, so the balls should slosh and re-pile as they turn.
Ship it as one working file named `index.html` that starts animating on load. 
Do not install packages. 
Do not open, screenshot, or headless-render the page (no Playwright, Puppeteer, or Chrome). 
Do not ask me clarifying questions — make reasonable assumptions and note them briefly in a comment at the top of the file.
```
Les deux modèles tournent à haut niveau d’effort de raisonnement, une tentative chacun sans relance, dans le même agent de code et avec les mêmes outils. Barème : passage/échec sur exécution, puis note de 1 à 5 sur la justesse physique, la stabilité sur 30 secondes et la qualité visuelle.

### Ce que GPT-6 Astra a produit

Astra a passé le test d’exécution et a obtenu 5/5 sur les trois axes. Il y est parvenu en 6 tours et 9 appels d’outils, avec un mix intéressant : écriture du fichier, 5 relectures, 2 grep, puis 2 patches avant de conclure.

La simulation est correcte. 44 balles, masse proportionnelle à l’aire, confinement maintenu sur 30 secondes pendant que l’hexagone tourne. Astra a dimensionné l’obstacle central assez grand par rapport à la chambre, de sorte que les balles le percutent régulièrement au lieu de se figer hors de sa portée.

Il a aussi construit une page éditoriale autour de la simulation, avec un titre en sérif, des libellés monospace élargis et un panneau de télémétrie affichant les vitesses de rotation. Personne ne l’avait demandé, mais c’est plutôt agréable. Il a livré des contrôles pause et redémarrage, non demandés, mais utiles pour l’inspection.

Astra fait tourner la chambre à 0,09 rad/s et l’obstacle à -0,16 rad/s, environ 3 fois plus lent que Fable pour la paroi externe et 5 fois plus lent pour l’obstacle. Plus lisible, et une mise à l’épreuve plus douce des frontières mobiles. L’invite demande une rotation lente, donc c’est conforme.

### Ce que Claude Fable 5.1 a produit

Fable a également passé l’exécution, avec 4 en justesse physique, 5 en stabilité et 4 en qualité visuelle. Il n’a eu besoin que de 2 tours et d’un seul appel d’outil : une écriture, sans relecture ni correctif.

La physique est saine sans être parfaite. Les balles adhèrent légèrement aux parois, d’où le point perdu en justesse, et en quelques secondes elles se regroupent dans les angles inférieurs, hors de portée de l’obstacle contre-rotatif qui reste inactif la plupart du temps.

Fable a mis l’accent sur l’instrumentation plutôt que sur la présentation. Un HUD affiche en temps réel le nombre de balles, le framerate, l’énergie cinétique et les deux vitesses de rotation, et un clic dans l’hexagone ajoute une balle au point visé. Ce dernier détail s’avère très pratique pour tester les collisions à la main.

### Résultats

| Mesure | GPT-6 Astra | Claude Fable 5.1 | 
|---|---|---|
| Tours | 6 | 2 | 
| Appels d’outils | 9 | 1 | 
| Exécution | OK | OK | 
| Justesse physique | 5 | 4 | 
| Stabilité dans le temps | 5 | 5 | 
| Qualité visuelle | 5 | 4 | 
| Score barème | 5,0 | 4,3 | 

Astra remporte ce test sur la qualité, avec un écart plus marqué sur la présentation que sur la physique. Leur divergence tient à la définition du travail : Astra a lu l’invite comme un brief à interpréter, ajoutant des contrôles, une mise en page éditoriale et une rotation plus lente, agréable à observer. Fable l’a lue comme un cahier des charges, l’a satisfait en un jet et a consacré le reste à un HUD de diagnostic.

Un seul run par modèle : prenez cela comme un point de données, pas comme un benchmark. Il rapporte les tours, pas les jetons ni le coût, et n’évalue que du code de simulation from scratch, pas l’usage PC ni le raisonnement vus plus haut.

## Quand choisir GPT-6 Astra vs Claude Fable 5.1

Avec des tarifs identiques, la décision dépend de la forme de votre charge, de la consommation réelle de jetons et de l’agent de code que vous utilisez déjà. Les trois chiffres à retenir : le facteur 4x sur les lectures de cache, le seuil de surcoût à 272 K, et le rapport 2,25x sur le coût mesuré par tâche.

### Choisissez GPT-6 Astra si…

