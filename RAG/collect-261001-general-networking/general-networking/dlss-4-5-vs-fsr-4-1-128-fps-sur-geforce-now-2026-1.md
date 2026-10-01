---
id: collect-261001-general-networking/general-networking/dlss-4-5-vs-fsr-4-1-128-fps-sur-geforce-now-2026-1
title: "dlss-4-5-vs-fsr-4-1-128-fps-sur-geforce-now-2026"
domain: general-networking
role: reference
task: reference
actors: ["AMD", "Intel", "Nvidia"]
dates: []
keywords: ["amd", "gpu", "intel", "nvidia"]
source: docs/RAG/collect-261001-general-networking/dlss-4-5-vs-fsr-4-1-128-fps-sur-geforce-now-2026.md
source_anchor: ""
source_lines: [1, 39]
sha256: 3b733a47c101f9b8009aba8f3b8ad663d258d538ff2f0ef19fa38c1bde5326bc
---

# dlss-4-5-vs-fsr-4-1-128-fps-sur-geforce-now-2026

Depuis la fin du déploiement des serveurs RTX 5080 sur GeForce NOW en Europe (Stockholm a été le dernier data center migré, selon les trackers de NVIDIA), la question ne porte plus sur le matériel mais sur le logiciel qui l’exploite. DLSS 4.5, dévoilé au CES 2026, embarque un second modèle de transformeur et un mode de génération d’images dynamique jusqu’à 6x. En face, AMD pousse FSR 4.1 sous la bannière “Redstone” et Intel affine XeSS 2 sur ses GPU Arc. Trois approches d’upscaling par IA, trois philosophies, et un seul terrain de jeu commun : le cloud gaming en streaming, où chaque milliseconde de latence et chaque artefact visuel se voient davantage qu’en local.

Pour les joueurs français et européens abonnés à GeForce NOW Ultimate, la question a un impact concret : quelle technologie d’upscaling profite réellement de l’infrastructure RTX 5080, quelle est la latence ajoutée par la génération d’images en streaming, et vaut-il la peine de payer l’abonnement le plus cher pour en profiter. Ce comparatif détaille les chiffres publiés par NVIDIA, les bancs d’essai indépendants de 2026 et les retours d’expérience concrets sur GeForce NOW.

## DLSS 4.5, FSR 4.1 et XeSS 2 : ce qui a changé en 2026

Les trois technologies d’upscaling par intelligence artificielle ont toutes reçu une mise à jour majeure entre janvier et juillet 2026. NVIDIA a présenté DLSS 4.5 au CES 2026 avec un nouveau modèle baptisé “Transformer 2”, capable selon NVIDIA de traiter jusqu’à cinq fois plus de calculs que le modèle Transformer de première génération utilisé dans DLSS 4. Cette puissance de calcul supplémentaire sert un objectif précis : améliorer la reconstruction des détails fins et la stabilité temporelle de l’image, deux points faibles historiques de l’upscaling par IA sur les scènes en mouvement rapide.

La nouveauté la plus visible de DLSS 4.5 reste le mode “Dynamic Multi Frame Generation”, qui ajuste en temps réel le multiplicateur de génération d’images selon la charge de la scène, avec un plafond porté à 6x. Concrètement, le GPU rend une image native et le modèle IA en génère cinq supplémentaires pour la combler, contre trois avec DLSS 4. NVIDIA a confirmé le déploiement complet de cette fonctionnalité au printemps 2026, exclusivement sur les GPU de la série RTX 50, ce qui inclut le tier RTX 5080 de GeForce NOW.

Côté AMD, FSR 4.1 s’inscrit dans la suite “Redstone” et reste la génération actuelle en 2026. Contrairement à DLSS, qui reste verrouillé aux GPU RTX, FSR 4.1 vise une compatibilité plus large, y compris sur certaines cartes NVIDIA et Intel plus anciennes, même si les meilleures performances restent réservées aux GPU AMD RDNA récents. Intel, de son côté, a fait évoluer XeSS vers sa version 2, avec un accent mis sur la génération d’images agressive : les bancs d’essai 2026 montrent des gains de FPS spectaculaires sur GPU Arc, au prix d’une qualité d’image jugée plus irrégulière que celle de DLSS 4.5.

## De DLSS 1 à DLSS 4.5 : dix ans de course à l’upscaling par IA

Pour comprendre pourquoi DLSS 4.5 change la donne sur GeForce NOW, il faut revenir sur la trajectoire de la technologie. La première version de DLSS, lancée avec les GPU RTX 20, souffrait d’un flou perceptible et d’un support limité à une poignée de jeux. DLSS 2 a marqué le vrai tournant en 2020 avec un réseau de neurones convolutif entraîné sur les supercalculateurs de NVIDIA, capable de reconstruire des détails à partir d’images basse résolution avec une fidélité inédite. DLSS 3, arrivé avec les RTX 40, a ajouté la génération d’images (Frame Generation), doublant le framerate perçu en insérant une image entièrement générée par IA entre deux images rendues.

DLSS 4, lancé avec les RTX 50 début 2025, a remplacé le réseau convolutif par un modèle de transformeur, la même famille d’architecture qui alimente les grands modèles de langage. Ce changement d’architecture a permis une meilleure compréhension du contexte spatial et temporel de chaque scène, réduisant les artefacts de “ghosting” (traînées visuelles) qui handicapaient les versions précédentes. DLSS 4.5, annoncé au CES 2026, pousse ce modèle de transformeur vers une deuxième génération, avec cinq fois plus de capacité de calcul selon NVIDIA, et introduit la génération d’images dynamique à 6x. En dix ans, la technologie est donc passée d’un simple filtre de netteté à un modèle d’IA générative capable de produire l’essentiel des pixels affichés à l’écran, un changement de paradigme qui explique pourquoi les GPU RTX 5080 des serveurs GeForce NOW s’appuient aussi lourdement sur DLSS plutôt que sur la puissance de calcul brute.

FSR a suivi une trajectoire parallèle mais distincte : les versions 1 et 2 s’appuyaient sur des algorithmes spatiaux et temporels classiques, sans réseau de neurones dédié, ce qui expliquait leur compatibilité universelle mais aussi leur qualité d’image inférieure. FSR 4, lancé en 2025, a marqué le premier virage d’AMD vers l’IA avec un modèle d’upscaling dédié fonctionnant sur les unités de calcul IA des GPU RDNA récents, abandonnant partiellement l’approche “compatible avec tout” au profit de la qualité. FSR 4.1, sous la bannière Redstone, affine cette approche en 2026 sans rupture architecturale majeure. Intel, arrivé plus tard sur le marché des GPU discrets avec Arc, a construit XeSS directement autour d’un modèle IA dès la version 1, ce qui explique pourquoi XeSS 2 progresse vite malgré un écosystème de cartes encore restreint.

## Tableau comparatif : DLSS 4.5 vs FSR 4.1 vs XeSS 2

Voici une synthèse des caractéristiques techniques des trois technologies, telles qu’elles se présentent en août 2026.

| Critère | DLSS 4.5 | FSR 4.1 (Redstone) | XeSS 2 | 
|---|---|---|---|
| Éditeur | NVIDIA | AMD | Intel | 
| Modèle IA | Transformer de 2e génération | Réseau de neurones optimisé RDNA | Modèle IA optimisé pour Arc (XMX) | 
| Compatibilité matérielle | GPU RTX uniquement (RTX 20 à 50) | GPU AMD RDNA 3/4 recommandés, support élargi RTX/Arc | Optimal sur Arc, mode DP4a sur autres GPU | 
| Génération d’images max | 6x (Dynamic Multi Frame Generation) | 2x (Frame Generation classique) | 2x (XeSS Frame Generation) | 
| Gain FPS observé (Cyberpunk 2077, 4K Qualité) | +128 % vs natif TAA | Non testé sur cette scène | Non testé sur cette scène | 
| Gain FPS observé (Spider-Man 2, 1440p Performance + FG) | Inférieur à FSR sur ce test précis | +236 % (105,2 im/s) | Non testé sur cette scène | 
| Gain FPS observé (F1 24, 1440p, FG activé) | Non testé sur cette scène | Non testé sur cette scène | +287 % (48 à 186 im/s sur Arc) | 
| Latence ajoutée, 1080p (Battlefield 6) | +6,4 ms (Frame Generation) | Jusqu’à +17,5 ms (Frame Generation) | Non communiqué séparément | 
| Latence ajoutée, 4K (Battlefield 6) | +17 ms | +13,4 ms | Non communiqué séparément | 
| Qualité de reconstruction (avis testeurs 2026) | Meilleure stabilité temporelle et anti-ghosting | Bonne mais en retrait sur mouvements fins | Correcte, plus irrégulière selon les scènes | 
| Disponible sur GeForce NOW RTX 5080 | Oui, natif | Non (service propriétaire NVIDIA) | Non (service propriétaire NVIDIA) | 
| Dernière mise à jour | Juillet 2026 (SL 2.11.1 / NGX 310.6.0) | 2026, suite Redstone | 2026, XeSS 2 | 

