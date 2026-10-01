---
id: collect-250926-servers-hardware/servers-hardware/ryzen-z2-extreme-8-curs-zen-5-contre-intel-arc-g3-2026-1
title: "ryzen-z2-extreme-8-curs-zen-5-contre-intel-arc-g3-2026"
domain: servers-hardware
role: reference
task: reference
actors: ["AMD", "Intel", "Microsoft"]
dates: []
keywords: ["intel", "amd", "benchmark", "benchmarks", "copilot", "gpu", "lpddr5x"]
source: docs/RAG/clean4/ryzen-z2-extreme-8-curs-zen-5-contre-intel-arc-g3-2026.md
source_anchor: ""
source_lines: [1, 39]
sha256: 5e8567d90e0f13170850682baf28630ba777a260eb562c5a4c6bb1c52cf6160c
---

# ryzen-z2-extreme-8-curs-zen-5-contre-intel-arc-g3-2026

Pendant cinq ans, une seule marque a fait tourner presque toutes les consoles portables PC du marché : AMD. Du Steam Deck à la ROG Ally, en passant par la première Legion Go et la MSI Claw, chaque machine emblématique embarquait un APU Ryzen taillé sur mesure. Au cœur de la cuvée 2026, c’est le **Ryzen Z2 Extreme** qui porte ce quasi-monopole : 8 cœurs Zen 5, un GPU Radeon 890M en RDNA 3.5 et un NPU de 50 TOPS, le tout dans un boîtier conçu pour tenir dans deux mains. Mais ce 30 juin 2026, le contexte a radicalement changé. Avec le lancement des premières consoles portables motorisées par les puces **Intel Arc G3**, AMD affronte pour la première fois un rival crédible sur son terrain de prédilection. Décryptage d’un processeur devenu, presque par accident, l’enjeu central de la bataille la plus disputée du jeu portable.

Le **Ryzen Z2 Extreme** n’est pas une nouveauté tombée du ciel : la puce a été dévoilée au CES 2025 et équipe désormais les trois consoles portables les plus convoitées de l’année – la Lenovo Legion Go 2, la ROG Xbox Ally X et la MSI Claw A8. Sa mission est limpide : offrir un bond de performances suffisant pour justifier des tarifs qui flirtent désormais avec ceux des ordinateurs portables, tout en gardant l’autonomie et la chaleur sous contrôle. Dans un marché européen secoué par la flambée du prix de la RAM et l’inflation des composants, chaque watt et chaque image par seconde comptent. Cet article fait le point sur l’architecture du Z2 Extreme, ses performances réelles, les machines qui l’embarquent et son face-à-face inédit avec l’Arc G3 d’Intel.

## Ryzen Z2 Extreme : ce que contient réellement la puce d’AMD

Le **Ryzen Z2 Extreme** appartient à la famille Strix Point d’AMD, la même architecture que l’on retrouve dans les ordinateurs portables Ryzen AI 300. Sa particularité tient à sa configuration hybride : sur ses 8 cœurs (16 threads) officialisés par AMD dès janvier 2025, trois sont des Zen 5 « performance » capables de grimper jusqu’à 5,0 GHz, et cinq sont des Zen 5C « compacts », optimisés pour l’efficacité énergétique et plafonnés à 3,3 GHz. Cette répartition 3+5 n’a rien d’anecdotique : elle permet au processeur de basculer entre puissance brute et économie de batterie selon que l’on joue branché sur secteur ou sur la seule autonomie de l’appareil. Le SoC embarque par ailleurs 8 Mo de cache L2 et 16 Mo de cache L3, soit 24 Mo au total, et un TDP configurable de 15 à 35 watts avec une valeur par défaut de 28 watts – des chiffres qu’AMD reconfirmait encore en juillet 2025 et que Tom’s Hardware avait détaillés dès l’annonce de janvier 2025 en pointant ce même cache de 24 Mo et ce pic à 5,0 GHz. Des essais comparatifs menés par Rutab.net en août 2025 montrent par ailleurs que la puce exploite jusqu’à 24 Go de LPDDR5X-8000, contre de la LPDDR5X-7500 seulement pour la Z1 Extreme – un gain de bande passante mémoire qui profite directement au GPU intégré. Un curseur que les fabricants exploitent pour ajuster le compromis entre performances et silence.

Côté graphismes, c’est le véritable atout du **ryzen z2 extreme** : une partie graphique Radeon 890M dotée de 16 unités de calcul gravées en RDNA 3.5, cadencées jusqu’à 2,9 GHz. C’est la version complète de l’iGPU 890M, là où la précédente génération devait se contenter d’un Radeon 780M en RDNA 3 limité à 12 unités – un écart de 16 contre 12 unités de calcul que PCGamesN a également documenté en comparant les deux puces cœur par cœur. À cela s’ajoute un NPU XDNA 2 affichant jusqu’à 50 TOPS, suffisant pour décrocher la certification Copilot+ de Microsoft et alimenter les fonctions d’intelligence artificielle locales ; Laptop Mag précisait en juin 2025 que ces 50 TOPS sont délivrés dans une enveloppe thermique plafonnée à 35 watts pour la variante Ryzen AI Z2 Extreme, tandis que Windows Central mesurait dès juillet 2025 un fonctionnement effectif entre 15 et 30 watts pour ces mêmes 16 cœurs RDNA 3.5. Selon la fiche officielle d’AMD, le Z2 Extreme est présenté comme le processeur le plus rapide jamais conçu par la marque pour le jeu portable. Le tableau ci-dessous résume ses caractéristiques techniques.

| Caractéristique | Ryzen Z2 Extreme | 
|---|---|
| Architecture CPU | Zen 5 + Zen 5C (3+5 cœurs) | 
| Cœurs / threads | 8 / 16 | 
| Fréquence CPU | 2,0 GHz base – jusqu’à 5,0 GHz | 
| iGPU | Radeon 890M, 16 CU RDNA 3.5 | 
| Fréquence GPU | jusqu’à 2,9 GHz | 
| NPU | XDNA 2 – 50 TOPS | 
| Cache | 8 Mo L2 + 16 Mo L3 (24 Mo) | 
| TDP configurable | 15 à 35 W | 
| Famille | Strix Point (lancée début 2025) | 

## Une famille Z2 à plusieurs visages : Extreme, AI, standard et Go

L’une des sources de confusion les plus fréquentes chez les acheteurs européens tient à la nomenclature d’AMD. Le terme « Z2 » recouvre en réalité plusieurs puces très différentes, et toutes ne se valent pas. Notebookcheck avait déjà repéré ce piège dès le dévoilement au CES 2025, où AMD annonçait d’un coup trois puces distinctes – Z2 Extreme, Z2 et Z2 Go – sous une même bannière. Au sommet trône le **Ryzen Z2 Extreme**, avec ses 8 cœurs Zen 5 et son Radeon 890M complet. AMD le décline en une variante baptisée **Ryzen AI Z2 Extreme**, techniquement identique mais qui conserve l’intégralité du NPU XDNA 2 de 50 TOPS pour les usages Copilot+ ; la propre grille de comparaison des puces pour consoles portables publiée par AMD en juin 2025 confirme noir sur blanc que seule cette variante « AI » embarque le NPU à 50 TOPS, le Z2 Extreme standard en étant dépourvu. C’est cette version « AI » qui équipe la ROG Xbox Ally X, là où certaines déclinaisons du Z2 Extreme standard bridaient la partie neuronale.

En dessous, le **Ryzen Z2** classique repose sur une architecture plus ancienne (un mélange Zen 4) et un iGPU moins fourni, tandis que le **Ryzen Z2 Go**, taillé pour les machines d’entrée de gamme, descend encore d’un cran avec une partie graphique RDNA 2. Pour le joueur, la leçon est simple : seul le Z2 Extreme – et sa variante AI – délivre le niveau de performances que l’on associe spontanément au haut de gamme portable. Cette segmentation explique aussi pourquoi deux consoles affichant fièrement la mention « Z2 » peuvent offrir des expériences radicalement différentes. Le tableau suivant clarifie la hiérarchie.

| Puce | Cœurs CPU | iGPU | NPU | Positionnement | 
|---|---|---|---|---|
| Ryzen AI Z2 Extreme | 8 (Zen 5 + 5C) | Radeon 890M (16 CU, RDNA 3.5) | XDNA 2, 50 TOPS | Haut de gamme + IA | 
| Ryzen Z2 Extreme | 8 (Zen 5 + 5C) | Radeon 890M (16 CU, RDNA 3.5) | réduit / désactivé | Haut de gamme | 
| Ryzen Z2 | 8 (Zen 4) | Radeon (12 CU, RDNA 3) | limité | Milieu de gamme | 
| Ryzen Z2 Go | 4 (Zen 3+) | Radeon (RDNA 2) | absent | Entrée de gamme | 

## Benchmarks du Ryzen Z2 Extreme face à la Z1 Extreme

La grande question que se posent les possesseurs d’une première génération de consoles portables est limpide : le saut vers le **ryzen z2 extreme benchmark** justifie-t-il un nouvel achat ? Les premiers tests indépendants apportent une réponse nuancée. Sur Geekbench 6, à une enveloppe thermique de 25 watts, le Z2 Extreme devance la Z1 Extreme d’environ 26 % en calcul mono-cœur et 27 % en multi-cœur, selon les mesures relayées par NotebookCheck. Un gain réel, mais qui reste dans la fourchette d’une évolution générationnelle classique plutôt que d’une révolution.

