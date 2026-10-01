---
id: collect-261001-general-networking/general-networking/fr-review-asus-ascent-qn10-review-550f5c4b-4
title: "fr-review-asus-ascent-qn10-review-550f5c4b"
domain: general-networking
role: reference
task: reference
actors: ["Intel", "Qualcomm"]
dates: []
keywords: ["benchmark", "diffusion", "gpu", "intel", "panther lake", "valuation"]
source: docs/RAG/collect-261001-general-networking/fr-review-asus-ascent-qn10-review-550f5c4b.md
source_anchor: ""
source_lines: [95, 130]
sha256: 91c6469a533914bbb0e8bddd6b1aaa7ad70096650ce6cdecea119e1d8a8ff6f7
---

# fr-review-asus-ascent-qn10-review-550f5c4b

| 7-Zip (GIPS, plus la valeur est élevée, mieux c'est) | ASUS Ascent QN10 | ASUS NUC 14 Pro (Core Ultra 7 165H) | Lenovo ThinkCentre Neo 50q QC | 
|---|---|---|---|
| Compression | 96.729 | N/D | 53.096 | 
| Décompression | 120.205 | N/D | 46.998 | 
| Note totale | 108.467 | N/D | 50.047 | 
Le total de 108 467 GIPS du QN10 dépasse largement le double des 50 047 GIPS du Neo 50q, la décompression à 120 205 GIPS constituant la performance la plus remarquable. À titre de comparaison, ce total surpasse tous les ordinateurs portables de notre récente comparaison Panther Lake ; il s'agit d'un véritable débit multicœur, et non d'une version allégée pour appareils mobiles. La colonne NUC 14 est vide car le benchmark de 7-Zip n'a pas produit de résultat sur ce système lors de notre test initial, et il n'a pas été inclus dans le nouveau test.
croque-y
Le test y-cruncher mesure la vitesse à laquelle le processeur calcule un grand nombre de décimales de Pi, ce qui sollicite fortement le CPU et la mémoire. Le test BBP, quant à lui, extrait les décimales hexadécimales de Pi, une tâche presque exclusivement limitée par le CPU. Les résultats sont exprimés en secondes ; plus le temps est court, mieux c'est. Le benchmark s'exécute nativement sur architecture Arm. Le test de 5 milliards de décimales sur le QN10 a été omis en raison de la mémoire disponible, limite prévue pour une machine de 32 Go, et le NUC 14 de 16 Go n'a pas pu dépasser le milliard de décimales lors des tests standard.
| y-cruncher (secondes, plus c'est court, mieux c'est) | ASUS Ascent QN10 | ASUS NUC 14 Pro (Core Ultra 7 165H) | Lenovo ThinkCentre Neo 50q QC | 
|---|---|---|---|
| Pi 1B | 71.531 | 36.561 | 186.905 | 
| Pi 2.5B | 218.332 | N/D | 497.704 | 
| Pi BBP 1B | 5.844 | 2.306 | N/D | 
| Pi BBP 10B | 67.701 | 25.144 | N/D | 
| Pi BBP 100B | 778.722 | 296.970 | N/D | 
Le QN10 traite 1 milliard de chiffres en 71.5 secondes, soit 2.6 fois plus vite que le Neo 50q (186.9 secondes). Cependant, le NUC 14 Pro, avec un temps de 36.6 secondes, démontre que le sous-système mémoire d'Intel reste largement supérieur dans cette tâche gourmande en bande passante, avec un gain de près de 2 fois. Les tests BBP, qui sollicitent uniquement le processeur, révèlent une situation inverse pour les calculs de grande envergure : le NUC traite 100 milliards de chiffres hexadécimaux en 297.0 secondes, contre 778.7 secondes pour le QN10. Le NUC 16 Go n'a pas pu dépasser le milliard de chiffres lors des tests standard ; la comparaison à 2.5 milliards de chiffres est donc à l'avantage du Lenovo, où le QN10 (218.3 secondes contre 497.7 secondes) affiche un écart de 2.3 fois par rapport à la génération précédente.
Mixeur
Le test de performances de Blender mesure le rendu à l'aide de trois scènes 3D différentes : Monstre, Boutique de bric-à-brac et Salle de classe. Les résultats sont exprimés en échantillons par minute ; plus le score est élevé, meilleures sont les performances. Les scores n'étant pas comparables entre les versions de Blender, nous ne présentons ici que les résultats de la version actuelle, Blender 5.2. Blender 5.2 intègre un chemin d'accès natif pour le processeur ARM64, utilisé par le QN10, mais le GPU Adreno n'étant pas pris en charge pour le rendu, contrairement à nos tests d'ordinateurs portables, ce tableau ne concerne que les performances du processeur pour les deux systèmes.
| Processeur Blender 5.2 (échantillons/min, plus c'est élevé, mieux c'est) | ASUS Ascent QN10 | ASUS NUC 14 Pro (Core Ultra 7 165H) | Lenovo ThinkCentre Neo 50q QC | 
|---|---|---|---|
| Monster | 188.04 | 109.16 | N/D | 
| Brocanteur | 133.21 | 81.08 | N/D | 
| Salle de classe | 93.93 | 57.41 | N/D | 
Les 18 cœurs Oryon permettent un rendu de Monster à 188 échantillons par minute contre 109 pour le NUC 14, soit une avance de 72 % qui se maintient sur les trois scènes. Ce chiffre surpasse tous les résultats obtenus avec les processeurs x86 ultra-fins de notre comparatif d'ordinateurs portables. L'absence de prise en charge du GPU est toutefois problématique : les moteurs de rendu qui s'appuient sur le GPU ne pourront pas exploiter pleinement ses capacités, et il s'agit là d'une lacune logicielle, et non matérielle.
UL Procyon IA
La suite d'IA d'UL Procyon couvre trois charges de travail. La génération d'images par IA mesure l'inférence de diffusion stable, des NPU basse consommation aux GPU haut de gamme ; nous présentons ici le test Stable Diffusion 1.5 INT8, conçu pour les accélérateurs basse consommation. La génération de texte par IA exécute des LLM locaux et évalue les performances de génération, en indiquant le temps d'obtention du premier jeton et le nombre de jetons par seconde ; la version Arm couvre Phi et Llama3. La vision par ordinateur par IA mesure l'inférence pour la classification d'images, la détection d'objets, la segmentation et les modèles de super-résolution ; la suite Computer Vision 2, plus récente, s'exécute via le chemin natif de chaque fournisseur. Des scores plus élevés indiquent de meilleures performances. Les étiquettes du moteur sont ici plus importantes que partout ailleurs dans cette évaluation : sur le QN10, la génération d'images et de texte s'est exécutée sur le NPU Hexagon via la pile QNN de Qualcomm, la vision par ordinateur s'est exécutée via SNPE sur le NPU puis via WinML, et le NUC 14 a utilisé ses meilleurs chemins disponibles, OpenVINO sur l'iGPU Arc et WinML. Aucun de ces chiffres n'est comparable aux résultats OpenVINO et ONNX de nos tests d'ordinateurs portables x86, car les scores Procyon ne sont comparables qu'au sein d'un même moteur et d'un même appareil.
| IA Procyon (plus élevé est mieux) | ASUS Ascent QN10 | ASUS NUC 14 Pro (Core Ultra 7 165H) | Lenovo ThinkCentre Neo 50q QC | 
|---|---|---|---|
| Génération d'images, SD 1.5 INT8 | 1 371 (QNN, NPU) | 1,004 (OpenVINO, iGPU) | N/D | 
| Génération de texte Phi | 1 371 (QNN, NPU) | 528 (OpenVINO, iGPU) | N/D | 
| Génération de texte Llama3 | 1 371 (QNN, NPU) | 434 (OpenVINO, iGPU) | N/D | 
| Vision par ordinateur | 4 254 (SNPE, NPU) | 135 (WinML, GPU) | N/D | 
| Vision par ordinateur 2 (WinML) | 2,161 | N/D | N/D | 
| Vision par ordinateur 2 (SNPE) | 1,897 | N/D | N/D | 
Considéré selon ses propres critères, le Hexagon 80 TOPS offre les performances suivantes : 5 457 en génération d'images SD 1.5 INT8, soit environ 1.9 fois plus que ce que nous avons mesuré sur des NPU x86 de classe 50 TOPS avec leurs meilleurs moteurs ; la génération de texte atteint 35.7 jetons par seconde sur Phi et 25.2 sur Llama3 avec un premier jeton en moins de 0.6 seconde, le tout entièrement sur le NPU sans solliciter le CPU. Le test de génération de texte hybride HTP+CPU a obtenu un score légèrement inférieur à celui du chemin NPU seul, ce qui suggère que le NPU n'est pas le facteur limitant. Le test Computer Vision 2, avec un score de 2 161 via WinML contre 1 897 via SNPE, est un rare exemple où le chemin générique surpasse le chemin du fournisseur. La comparaison avec le NUC 14 souligne la mise en garde concernant les moteurs : ses meilleurs chemins, OpenVINO sur l’iGPU Arc, se classent à 1 004 en génération d’images et à 528 sur Phi, et son exécution de génération de texte NPU a échoué complètement ; ainsi, l’avance de cinq fois du QN10 en génération d’images mêle des avantages matériels et logiciels qui ne peuvent être totalement dissociés.
Conclusion
