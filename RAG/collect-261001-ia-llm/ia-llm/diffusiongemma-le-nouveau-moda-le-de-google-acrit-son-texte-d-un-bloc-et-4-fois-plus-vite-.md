---
id: collect-261001-ia-llm/ia-llm/diffusiongemma-le-nouveau-moda-le-de-google-acrit-son-texte-d-un-bloc-et-4-fois-plus-vite-
title: "DiffusionGemma : le nouveau modÃ¨le de Google Ã©crit son texte d'un bloc, et 4 fois plus vite"
domain: ia-llm
role: reference
task: reference
actors: ["Anthropic", "Apple", "Google", "Hugging Face", "Nvidia"]
dates: []
keywords: ["diffusion", "apache", "claude", "gemini", "gpu", "mixture of experts", "nvidia"]
source: docs/RAG/collect-261001-ia-llm/diffusiongemma-le-nouveau-moda-le-de-google-acrit-son-texte-d-un-bloc-et-4-fois-plus-vite-korben.md
source_anchor: ""
source_lines: [1, 31]
sha256: 8e7f82bc7e615a17424f9356958c2817c84567b26e967c6aea6d67317938fa61
---

# DiffusionGemma : le nouveau modÃ¨le de Google Ã©crit son texte d'un bloc, et 4 fois plus vite

## Ce quâil faut retenir RÃ©sumÃ© gÃ©nÃ©rÃ© par IA

1. DiffusionGemma gÃ©nÃ¨re le texte en bloc plutÃ´t que token par token, atteignant plus de 1 000 tokens/seconde sur H100 et environ 700 sur RTX 5090, soit quatre fois plus vite que les Gemma classiques de taille comparable.
2. Le modÃ¨le fonctionne comme un gÃ©nÃ©rateur d'images : il pose un canevas de 256 tokens fictifs, le raffine en plusieurs passes, puis finalise le bloc entier d'un coup, ce qui permet de rÃ©soudre des tÃ¢ches non linÃ©aires comme le Sudoku Ã 80% de rÃ©ussite.
3. DiffusionGemma (26 milliards de paramÃ¨tres, 3,8 milliards actifs) tient en 18 Go de mÃ©moire vidÃ©o compressÃ©e, sort sous licence Apache 2.0 avec poids tÃ©lÃ©chargeables sur Hugging Face, et fonctionne sur Mac via MLX.

Plus de 1 000 tokens par seconde sur une seule carte H100, l'accÃ©lÃ©rateur que Nvidia vend aux centres de donnÃ©es, et environ 700 sur une RTX 5090, sa carte gaming haut de gamme. C'est le dÃ©bit que Google DeepMind annonce pour DiffusionGemma, son nouveau modÃ¨le d'IA ouvert, Ã peu prÃ¨s quatre fois ce que produisent les modÃ¨les Gemma classiques de taille comparable.

Toute la diffÃ©rence se joue dans la faÃ§on de gÃ©nÃ©rer le texte. Les modÃ¨les de langage habituels sont autorÃ©gressifs : ils Ã©crivent de gauche Ã droite, un token Ã la fois, le token Ã©tant le petit morceau de mot que manipule une IA. DiffusionGemma fait tout autrement.

Il travaille comme les gÃ©nÃ©rateurs d'images, qui partent d'un nuage de bruit et le dÃ©bruitent petit Ã petit jusqu'Ã la photo demandÃ©e. Le modÃ¨le pose un canevas de 256 tokens fictifs, repasse dessus plusieurs fois pour affiner ses estimations, puis finalise le bloc entier d'un coup.

Sous le capot, on a un Mixture of Experts de 26 milliards de paramÃ¨tres, une architecture oÃ¹ seule une petite partie du modÃ¨le se rÃ©veille Ã chaque calcul, 3,8 milliards ici. Du coup le tout tient dans 18 Go de mÃ©moire vidÃ©o en version compressÃ©e, soit une grosse carte graphique grand public.

L'intÃ©rÃªt en local, c'est que cette approche dÃ©place le goulot d'Ã©tranglement de la bande passante mÃ©moire, la vitesse Ã laquelle la carte lit ses propres donnÃ©es, vers le calcul pur. Dans le cloud, les serveurs mutualisent les requÃªtes de milliers d'utilisateurs et leurs puces tournent en permanence, alors que votre GPU Ã la maison passe le plus clair de son temps Ã attendre les donnÃ©es. La diffusion occupe ces cycles perdus.

Et puis il y a les tÃ¢ches non linÃ©aires, oÃ¹ l'ordre d'Ã©criture ne suit pas l'ordre de lecture. Google a mÃªme affinÃ© une version sur le Sudoku, un casse-tÃªte rÃ©putÃ© impossible pour les modÃ¨les classiques puisque chaque case dÃ©pend de cases pas encore Ã©crites. DiffusionGemma, qui corrige son canevas en continu, atteint 80% de rÃ©ussite en faisant tomber les Ã©tapes de calcul de 48 Ã 12.

Tout n'est pas rose pour autant. Dans une image, un pixel ratÃ© passe inaperÃ§u. Un token mal prÃ©dit, lui, peut rendre un paragraphe entier incohÃ©rent et forcer Ã tout recommencer. Et pour une rÃ©ponse de cinq mots, dÃ©grossir un canevas complet gaspille du calcul. C'est d'ailleurs pour Ã§a que les gros Gemini du cloud n'y passent pas.

Le modÃ¨le est expÃ©rimental, mais il sort sous licence Apache 2.0, la mÃªme que le reste de la famille Gemma 4, donc utilisable commercialement sans restriction. Les poids se tÃ©lÃ©chargent dÃ¨s maintenant sur Hugging Face, la plateforme de rÃ©fÃ©rence des modÃ¨les ouverts, avec une optimisation menÃ©e main dans la main avec Nvidia. MLX, l'outil d'Apple pour faire tourner l'IA en local, est aussi de la partie, les Mac sont donc servis.

Si vous voulez mon avis, c'est sur ces modÃ¨les locaux que Google est le plus intÃ©ressant en ce moment, bien plus que sur Gemini.

Source : ARS Technica

## Commentaires

starfix!dans Surfshark ne vous rend pas invMorganedans Discord devine votre Ã¢ge sansts3rv1dans Les Ray-Ban Display arrivent eponpondans Openpilot - La NHTSA passe lesfabiendans Claude Code vous fait choisir
