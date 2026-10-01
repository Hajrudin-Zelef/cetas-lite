---
id: collect-261001-general-networking/general-networking/dlss-5-vs-fsr-4-1-vs-xess-2-comparatif-2026-2
title: "dlss-5-vs-fsr-4-1-vs-xess-2-comparatif-2026"
domain: general-networking
role: reference
task: reference
actors: ["AMD", "Intel", "Nvidia"]
dates: []
keywords: ["amd", "diffusion", "gpu", "intel", "latency", "nvidia"]
source: docs/RAG/collect-261001-general-networking/dlss-5-vs-fsr-4-1-vs-xess-2-comparatif-2026.md
source_anchor: ""
source_lines: [29, 56]
sha256: ea24db3c97be902d59ff894a74ab49cf500acf1a06574ecc5bafb381d671005f
---

# dlss-5-vs-fsr-4-1-vs-xess-2-comparatif-2026

Ces moyennes se retrouvent dans des tests indépendants. Le site Club386 a mesuré, sur une Arc B580 Limited Edition en 1440p, un passage de 51 fps en rendu natif à 85 fps avec XeSS Performance seul (+66,7 %), puis à 145 fps une fois la génération de frames ajoutée (+184 % par rapport au natif), un résultat cohérent avec la moyenne officielle d’Intel. Sur *F1 24*, l’un des premiers titres compatibles XeSS 2, un test technique chinois relayé par plusieurs médias spécialisés a mesuré un bond de 25,8 fps en ray tracing natif à 81,3 fps avec XeSS Performance (+215 %), puis à 123 fps avec la génération de frames XeSS 2 activée, soit une hausse de 51,3 % supplémentaires par rapport au mode upscaling seul, et une latence divisée par deux grâce à Xe Low Latency (6,61 ms contre 13,4 ms).

La limite actuelle de XeSS 2 est la couverture matérielle : les données solides et vérifiées concernent presque exclusivement l’Arc B580, la carte d’entrée de gamme d’Intel vendue autour de 390 € en France selon les comparateurs de prix. Aucune donnée fiable de performance XeSS 2 sur l’Arc B770 n’a pu être confirmée par une source de presse spécialisée au moment de la rédaction de cet article, ce qui traduit une diffusion encore limitée du haut de gamme Arc en Europe. Pour situer l’Arc B580 face à la concurrence directe côté milieu de gamme, notre comparatif RTX 5060 Ti vs RX 9060 XT détaille les écarts de prix et de performance sur ce segment.

## De DLSS 1 à DLSS 5 : un rappel avant de comparer

Pour comprendre l’ampleur du changement apporté par DLSS 5, un rapide retour en arrière aide à situer les enjeux. Nvidia a lancé la première version de DLSS en 2019, avec un simple réseau de neurones chargé de reconstruire une image en haute résolution à partir d’un rendu plus léger. Les versions 2 et 3 ont ensuite ajouté la reconstruction temporelle puis la première génération de frames, avant que DLSS 4 n’introduise un modèle transformeur remplaçant l’ancien réseau convolutif, avec à la clé une génération de frames multiple capable d’insérer plusieurs images calculées par IA entre deux images réellement rendues. DLSS 5 pousse cette logique encore plus loin en injectant des données de scène 3D directement dans le processus de rendu, une étape que Nvidia présente comme le franchissement d’un nouveau palier plutôt qu’une simple mise à jour incrémentale.

Côté AMD, l’histoire suit une trajectoire un peu différente. FSR 1, sorti en 2021, se contentait d’un filtre spatial sans intelligence artificielle, ce qui expliquait sa compatibilité universelle mais aussi sa qualité d’image en retrait face à DLSS. FSR 2 a introduit la reconstruction temporelle, FSR 3 la génération de frames, et FSR 4 a marqué la bascule vers un véritable modèle d’apprentissage automatique, un changement de paradigme qui explique pourquoi AMD a dû, comme Nvidia, restreindre les fonctions les plus avancées à sa dernière architecture (RDNA 4). FSR 4.1 et la suite Redstone représentent donc la consolidation de cette bascule vers le tout-IA, plus que la création d’une technologie entièrement nouvelle.

Intel, arrivé plus tardivement sur le marché des GPU dédiés avec les Arc A-series en 2022, a construit XeSS directement autour d’un modèle par apprentissage automatique dès sa première version, avec un chemin d’exécution optimisé pour ses propres unités XMX et un chemin de repli DP4a pour les GPU tiers. Cette approche explique pourquoi Intel a pu ajouter la génération de frames et Xe Low Latency dans XeSS 2 sans rupture technologique aussi marquée que celle vécue par AMD entre FSR 3 et FSR 4. Le retard d’Intel se situe moins dans la technologie elle-même que dans le nombre de GPU Arc réellement vendus, ce qui limite le nombre de tests indépendants disponibles comparé aux écosystèmes bien plus larges de Nvidia et AMD.

## 5 exemples concrets déjà visibles dans les jeux et le matériel

**1. NBA 2K27 et le lancement de DLSS 5.** Le jeu de basketball de 2K Games reste, au moment de la rédaction de cet article, le seul titre à intégrer nativement la 3D-Guided Neural Rendering de DLSS 5, déployée via le pilote Game Ready 616.64 WHQL le 3 septembre 2026. C’est ce jeu qui a servi de vitrine, pour le meilleur (les gains de fps bruts) et pour le pire (la controverse sur l’écart entre fps affichés et rendu natif).

**2. F1 24 et les premiers pas de XeSS 2.** Le jeu de course de Codemasters a servi de banc d’essai pour la génération de frames d’Intel, avec des gains mesurés dépassant 200 % par rapport au rendu natif en ray tracing sur une Arc B580, un résultat qui a permis à la presse spécialisée de documenter précisément le fonctionnement de Xe Low Latency en conditions réelles.

**3. L’extension de FSR Upscaling 4.1 dans Unreal Engine 5.8.** AMD a mis à jour son plugin officiel pour le moteur Unreal Engine afin d’y intégrer FSR Upscaling 4.1, permettant aux studios utilisant Unreal Engine 5.8 d’ajouter facilement l’upscaling par IA de la RX 7000 à leurs futurs jeux sans développement spécifique, un exemple concret de la stratégie d’AMD pour accélérer l’adoption logicielle de sa technologie.

**4. Le hack communautaire de DLSS 5 sur RX 9070 XT.** Des moddeurs ont réussi, quelques jours après le lancement, à forcer l’exécution de DLSS 5 sur une carte AMD non prévue pour cette technologie. Le résultat, environ 30 fps en 1080p selon Tom’s Hardware, montre à la fois la porosité technique entre écosystèmes concurrents et la marge de progression nécessaire avant qu’un portage non officiel devienne réellement jouable.

**5. GeForce NOW Ultimate comme vitrine cloud de DLSS 5.** Plutôt que d’attendre une mise à niveau matérielle locale, les abonnés du service de cloud gaming de Nvidia peuvent accéder à DLSS 5 via des serveurs équipés de GPU de classe RTX 5080, un exemple concret de la manière dont le cloud gaming permet de contourner le coût d’achat d’une carte graphique neuve pour profiter d’une technologie encore réservée au haut de gamme.

## Tableau comparatif technique : DLSS 5 vs FSR 4.1 vs XeSS 2

Voici la synthèse technique des trois technologies d’upscaling, avec les caractéristiques vérifiées à partir des annonces officielles et des tests indépendants disponibles en septembre 2026.

