---
id: collect-261001-general-networking/general-networking/dlss-5-vs-fsr-4-1-vs-xess-2-comparatif-2026-1
title: "dlss-5-vs-fsr-4-1-vs-xess-2-comparatif-2026"
domain: general-networking
role: reference
task: reference
actors: ["AMD", "Intel", "Nvidia"]
dates: []
keywords: ["amd", "benchmarks", "blackwell", "diffusion", "foundry", "fp8", "gpu", "intel", "latency", "nvidia"]
source: docs/RAG/collect-261001-general-networking/dlss-5-vs-fsr-4-1-vs-xess-2-comparatif-2026.md
source_anchor: ""
source_lines: [1, 28]
sha256: 2ea28565fcf34ee970710c8ae2b505c11efbbed8e5fbe3ec1c414b420e3912a9
---

# dlss-5-vs-fsr-4-1-vs-xess-2-comparatif-2026

Le 3 septembre 2026, Nvidia a mis en ligne DLSS 5 sur *NBA 2K27* avec un chiffre choc : jusqu’à 590 images par seconde en 1440p sur une RTX 5090. Trois jours plus tard, Tom’s Hardware et Notebookcheck démontraient que le rendu natif de cette même carte tournait en réalité autour de 62 fps, le reste provenant d’une génération d’images par IA. La polémique a immédiatement relancé un débat plus large : DLSS 5, FSR 4.1 “Redstone” d’AMD et XeSS 2 d’Intel jouent-ils encore au même jeu ? Ce comparatif détaille les trois technologies d’upscaling, leurs GPU compatibles, leurs vrais gains de performance mesurés par la presse spécialisée, et leur impact direct sur le prix des cartes graphiques vendues en France.

## DLSS 5, FSR 4.1 et XeSS 2 : pourquoi ce trio s’impose en septembre 2026

Les trois grands fabricants de GPU ont livré leur nouvelle génération d’upscaling à quelques mois d’écart. Nvidia a dégainé **DLSS 5** avec sa fonction “3D-Guided Neural Rendering”, intégrée pour la première fois dans *NBA 2K27* via le pilote Game Ready 616.64 WHQL. AMD a de son côté déployé **FSR Upscaling 4.1**, brique centrale de la suite “Redstone” annoncée au CES 2026, qui ajoute la génération de frames, la régénération de rayons et un cache de radiance. Intel, enfin, a étendu **XeSS 2** avec la génération de frames et la technologie Xe Low Latency sur ses cartes Arc B-series.

Ce qui rend la comparaison utile maintenant, c’est que les trois technologies ne visent plus seulement à améliorer la définition d’image : elles cherchent à multiplier le nombre d’images affichées sans que le GPU les calcule réellement toutes. Cette bascule vers la “frame generation” massive change la façon de lire un test de performance, et elle a des conséquences concrètes sur le choix d’achat d’une carte graphique en France, où les prix de la gamme RTX 50 et RX 9000 restent élevés. Le mot-clé “dlss 5” affiche à lui seul plus de 6 600 recherches mensuelles en France selon les données de suivi de mots-clés de septembre 2026, preuve que le sujet dépasse largement le cercle des passionnés. Pour les lecteurs qui découvrent seulement l’ancienne génération, notre guide pour activer DLSS 4 sur RTX 50 reste utile en attendant l’extension officielle de DLSS 5 à davantage de jeux.

## DLSS 5 : la 3D-Guided Neural Rendering de Nvidia

DLSS 5 introduit un concept que Nvidia baptise “3D-Guided Neural Rendering”. Contrairement aux versions précédentes qui se contentaient de reconstruire une image à partir d’une résolution plus basse, ce nouveau modèle combine les données de la scène 3D fournies par le moteur de jeu (éclairage, matériaux, ombres de contact, diffusion sous-cutanée sur la peau) avec un réseau de neurones qui ajuste le rendu final en temps réel. Selon Nvidia, cette approche rapproche le rendu en temps réel de la qualité d’un rendu cinéma, un argument que le communiqué officiel de Nvidia met en avant pour justifier le lancement exclusif sur *NBA 2K27*.

Techniquement, DLSS 5 est réservé au lancement aux **GPU GeForce RTX 50** (bureau et portables) ainsi qu’à GeForce NOW Ultimate, qui s’appuie sur des serveurs cloud équipés de puces de classe RTX 5080. Les cartes RTX 40 et RTX 30 en sont exclues à ce stade, même si Nvidia a indiqué vouloir étendre le support officiel aux RTX 40 une fois les modèles optimisés sur la génération Blackwell, selon des propos rapportés par la presse spécialisée. Des moddeurs sont parvenus à faire tourner DLSS 5 de force sur des RX 9070 XT d’AMD, mais Tom’s Hardware précise que la carte ne dépasse pas 30 fps en 1080p dans cette configuration non officielle, loin de l’expérience visée.

La sortie a immédiatement suscité une controverse. Les chiffres marketing de Nvidia annoncent 370 fps en 4K et 590 fps en 1440p sur RTX 5090 avec les réglages Ultra et le ray tracing activé. Mais Tom’s Hardware a démontré que ce chiffre de 370 fps correspond à une image réellement calculée à environ 61 fps native, le reste étant généré par la génération de frames multiple (Multi Frame Generation). Notebookcheck, qui a repris une méthodologie proche de celle de Digital Foundry, a mesuré une chute de 128,4 fps à 62,7 fps en 4K Ultra sur RTX 5090 une fois DLSS 5 activé avec le mode neural rendering complet, soit une perte de rendu natif de 51,2 %. Sur RTX 5070, la perte grimpe à 55 %, et à 41 % sur RTX 5060. Cette dissociation entre fps affichés et fps réellement calculés a valu à DLSS 5 le surnom de “filtre à bouillie IA” (“AI slop filter”) dans une partie de la presse et des communautés de joueurs, un qualificatif né précisément de cet écart entre communication et mesures indépendantes.

## FSR 4.1 Redstone : la réponse d’AMD, plus ouverte mais moins spectaculaire

AMD n’a pas attendu DLSS 5 pour répondre. Sa suite **FSR Redstone**, présentée au CES 2026, regroupe l’upscaling par apprentissage automatique (FSR Upscaling 4.1), la génération de frames, la régénération de rayons et un cache de radiance. Le chief software officer d’AMD, Andrej Zdravković, a confirmé lors d’une table ronde presse que l’ensemble complet de la suite Redstone restait pour l’instant réservé aux GPU RDNA 4, c’est-à-dire les Radeon RX 9070, RX 9070 XT et RX 9060 XT, un choix similaire à la stratégie de Nvidia avec ses RTX 50.

La différence de philosophie se joue ailleurs. Selon des informations reprises par Club386 et IGN, AMD envisage d’ouvrir une partie de la technologie FSR 4 à des GPU concurrents, y compris certaines cartes Nvidia dotées d’unités IA compatibles FP8 (architecture Ada et plus récente), même si le cœur de la suite Redstone resterait propre à RDNA 4. C’est un positionnement inverse de celui de Nvidia, qui verrouille DLSS 5 à son propre matériel. La documentation technique d’AMD publiée sur GPUOpen confirme cette volonté de compatibilité élargie, tout en reconnaissant que le modèle transformeur de DLSS 4.5 garde une avance en qualité d’image pure.

Côté performance, les chiffres publics restent plus rares que pour DLSS 5. AMD revendique un gain moyen de 3,5 fois les performances sur RX 9070 XT en combinant upscaling et génération de frames sur l’ensemble de sa suite Redstone, une moyenne multi-jeux et non un chiffre par titre. Un test comparatif plus précis, réalisé sur une Radeon RX 9060 XT, mesure un gain concret de 44 % en mode FSR 4.1 Qualité par rapport au rendu natif, avec une moyenne de 61 fps et un minimum de 55 fps dans le jeu testé. Sur le plan de l’adoption, FSR 4 en mode upscaling par IA équipe désormais plus de 200 jeux, et la génération de frames dédiée à la RDNA 4 est présente dans une trentaine de titres selon les données publiées par la presse spécialisée en hardware, un rythme d’adoption plus rapide que celui de DLSS 5 qui ne compte qu’un seul jeu au lancement. Nous avions déjà détaillé l’arrivée de FSR 4.1 sur les Radeon RX 7000, une extension aux GPU de génération précédente qui distingue la stratégie logicielle d’AMD de celle de Nvidia.

## XeSS 2 : Intel mise sur la génération de frames pour exister face aux deux géants

Intel reste le troisième acteur du marché des GPU dédiés, mais XeSS 2 affiche des gains impressionnants sur le papier, précisément parce que le point de départ (les cartes Arc) est plus modeste en puissance brute. La documentation officielle de performance d’Intel, disponible sur son centre de documentation des benchmarks Arc, indique un gain moyen de 85 % avec XeSS classique (upscaling seul) et de 187 % avec XeSS 2 une fois la génération de frames et la technologie anti-latence Xe Low Latency activées, sur une sélection de jeux testés sur Arc B580.

