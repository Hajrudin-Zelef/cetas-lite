---
id: collect-261001-general-networking/general-networking/dlss-5-vs-fsr-4-1-vs-xess-2-comparatif-2026-3
title: "dlss-5-vs-fsr-4-1-vs-xess-2-comparatif-2026"
domain: general-networking
role: reference
task: reference
actors: ["AMD", "Intel", "Nvidia"]
dates: []
keywords: ["amd", "benchmarks", "blackwell", "fp8", "gpu", "intel", "latency", "nvidia"]
source: docs/RAG/collect-261001-general-networking/dlss-5-vs-fsr-4-1-vs-xess-2-comparatif-2026.md
source_anchor: ""
source_lines: [57, 103]
sha256: 65996fdcbe38d8544e319ffe06908eb0de7b60f0f0a072474dab06008d2e93a3
---

# dlss-5-vs-fsr-4-1-vs-xess-2-comparatif-2026

| Critère | DLSS 5 (Nvidia) | FSR 4.1 Redstone (AMD) | XeSS 2 (Intel) | 
|---|---|---|---|
| Date de lancement public | 3 septembre 2026 | CES 2026 (suite Redstone) | 2026, extension avec génération de frames | 
| Fonction phare | 3D-Guided Neural Rendering | Upscaling ML + génération de frames + régénération de rayons + cache de radiance | Upscaling ML + génération de frames + Xe Low Latency | 
| GPU natifs supportés | GeForce RTX 50 (bureau et portable), GeForce NOW Ultimate | Radeon RX 9070, RX 9070 XT, RX 9060 XT (suite complète) | Arc B580 (données vérifiées), architecture Xe2 | 
| Support GPU plus anciens | RTX 40 annoncé pour plus tard, RTX 30 non supporté | Upscaling FSR 4 classique déjà étendu aux RX 7000 ; ouverture explorée vers RDNA 3 et GPU tiers | Chemin de repli DP4a possible sur GPU non-Intel (architecture antérieure) | 
| Premier jeu compatible | NBA 2K27 (exclusif au lancement) | Plus de 200 jeux en upscaling FSR 4, une trentaine en génération de frames RDNA 4 | F1 24 parmi les premiers titres XeSS 2 | 
| Gain moyen mesuré (upscaling seul) | Non isolé par la presse, mesuré combiné | +44 % (RX 9060 XT, mode Qualité, test indépendant) | +85 % en moyenne (données Intel), +66,7 % à +215 % selon le jeu testé | 
| Gain moyen mesuré (upscaling + génération de frames) | Chute du rendu natif de 41 à 55 % selon la carte une fois le mode complet activé | Jusqu’à 3,5x en moyenne multi-jeux selon AMD | +187 % en moyenne (données Intel), jusqu’à +376 % sur F1 24 en RT | 
| Ouverture multi-vendeur | Fermé à l’écosystème RTX | Upscaling de base ouvert, exploration d’un support Nvidia via FP8 | Conçu pour être API-agnostique, chemin de repli sur GPU non-Intel | 
| Disponible en cloud gaming | Oui, via GeForce NOW Ultimate | Non confirmé pour un service cloud grand public | Non confirmé pour un service cloud grand public | 
| Principale critique | Écart marketing/rendu natif, surnommé “filtre à bouillie IA” | Suite complète encore réservée à une seule génération de GPU (RDNA 4) | Couverture matérielle limitée en haut de gamme (Arc B770 peu documenté) | 
| Modèle de calcul | Réseau neuronal guidé par les données 3D du moteur de jeu | Upscaling par apprentissage automatique, plus léger que DLSS selon AMD | Réseau temporel ML, accéléré par les unités XMX sur Arc | 

## Tableau des prix : l’impact sur les cartes graphiques vendues en France

Le choix d’une technologie d’upscaling est indissociable du prix de la carte graphique qui la supporte. Voici les tarifs observés en France mi-septembre 2026, entre le MSRP officiel et les prix constatés chez les revendeurs.

| Carte graphique | Technologie supportée | MSRP officiel (Europe) | Prix constaté en France (sept. 2026) | 
|---|---|---|---|
| RTX 5070 | DLSS 5 | Non communiqué séparément par Nvidia | Environ 780 € à 1 000 €, médiane proche de 900 € selon les comparateurs de prix | 
| RTX 5070 Ti | DLSS 5 | 884 € TTC | Prix constatés au-dessus du MSRP selon la disponibilité | 
| RTX 5080 | DLSS 5 | 1 179 € TTC (Founders Edition) | Variable selon les modèles des partenaires | 
| RTX 5090 | DLSS 5 | 2 349 € TTC (Founders Edition) | Rapporté jusqu’à 3 700-5 500 € lors des pics de pénurie en 2026 | 
| RX 9070 | FSR 4.1 Redstone | Non communiqué séparément par AMD | À partir d’environ 930 € (modèle Sapphire) | 
| RX 9070 XT | FSR 4.1 Redstone | Non communiqué séparément par AMD | Environ 780 € à 980 € selon les partenaires (ASRock, Sapphire, XFX, ASUS) | 
| RX 9060 XT | FSR 4.1 Redstone | Non confirmé publiquement en euros | Non confirmé par une source vérifiée | 
| Arc B580 | XeSS 2 | Environ 390 € en France selon les comparateurs | Environ 390 € | 

Ce tableau illustre un déséquilibre net : DLSS 5 impose l’achat d’une carte RTX 50, dont le ticket d’entrée réel dépasse 780 € en France, pour profiter de la fonctionnalité la plus avancée du marché. FSR 4.1 se situe dans une fourchette de prix comparable côté AMD, sans avantage tarifaire décisif. Seul Intel casse les prix avec l’Arc B580 à moins de 400 €, mais au prix d’une puissance brute et d’une couverture logicielle encore loin de celles de ses deux concurrents. Pour approfondir l’écart de prix précis entre les modèles Nvidia, consultez notre comparatif RTX 5070 vs RTX 5070 Ti ainsi que notre analyse RTX 5080 vs RTX 5090, qui détaillent les écarts de VRAM et de fps entre chaque palier de la gamme Blackwell.

## Benchmarks réels : ce que mesurent Tom’s Hardware, Notebookcheck et Club386

Au-delà des chiffres marketing, trois sources indépendantes permettent de comparer les technologies sur des bases concrètes. Tom’s Hardware a testé DLSS 5 sur l’ensemble de la gamme RTX 50 dans NBA 2K27 et confirme que même la RTX 5060, la carte d’entrée de gamme Blackwell, parvient à maintenir près de 60 fps en moyenne en 1080p avec DLSS 5 activé, un résultat honorable mais très éloigné des centaines d’images par seconde mises en avant par Nvidia lors de l’annonce.

Notebookcheck a documenté la perte de performance native engendrée par le mode neural rendering complet de DLSS 5 : 51,2 % sur RTX 5090, 55 % sur RTX 5070 et 41 % sur RTX 5060 en 4K Ultra. Cette perte s’explique par le coût de calcul du réseau de neurones qui ajuste l’éclairage et les matériaux en temps réel, une charge additionnelle par rapport à un simple upscaling de résolution. Côté Intel, le test de l’Arc B580 par Club386 montre au contraire un bénéfice net à chaque étage : le rendu natif à 51 fps grimpe à 85 fps avec l’upscaling seul, puis à 145 fps une fois la génération de frames ajoutée, sans le même effet de compensation négative observé côté DLSS 5.

Sur le terrain d’AMD, les données publiques restent plus parcellaires. Le constructeur communique une moyenne agrégée de 3,5x sur l’ensemble de la suite Redstone, sans détail par jeu comparable aux tests indépendants menés sur DLSS 5 ou XeSS 2. Le seul chiffre isolé disponible, un gain de 44 % en mode Qualité sur RX 9060 XT, confirme que FSR 4.1 apporte un bénéfice réel, mais sur un échantillon plus restreint que ses deux concurrents.

## Ce que dit Nvidia sur l’avenir du rendu neuronal

Jensen Huang, fondateur et PDG de Nvidia, a personnellement défendu l’ambition de DLSS 5 lors de l’annonce officielle. Il a déclaré : “Vingt-cinq ans après l’invention par Nvidia du shader programmable, nous réinventons à nouveau l’infographie” (source : communiqué officiel Nvidia GeForce). Il a également résumé la portée qu’il attribue à cette technologie en la comparant à une rupture générationnelle : “DLSS 5, c’est le moment GPT du graphisme, il mélange le rendu artisanal avec l’IA générative pour offrir un bond spectaculaire de réalisme visuel tout en préservant le contrôle dont les artistes ont besoin pour leur expression créative” (source : communiqué officiel Nvidia GeForce).

Ces déclarations tranchent avec la réception critique de la presse et d’une partie des joueurs, qui reprochent précisément à Nvidia de présenter des gains de performance qui reposent majoritairement sur de l’image générée plutôt que calculée. L’écart entre le discours d’entreprise et les mesures indépendantes de Tom’s Hardware et Notebookcheck illustre la tension actuelle du marché : plus les technologies d’upscaling avancent, plus la définition même de “performance” devient complexe à communiquer au grand public. Ce débat n’est pas neuf : notre comparatif précédent sur DLSS 4.5 vs FSR 4.1 sur GeForce NOW avait déjà documenté des écarts de fps de plus de 128 % entre technologies, avant même l’arrivée du rendu neuronal complet de DLSS 5.

## 5 cas d’usage concrets : quelle techno choisir selon votre profil

