---
id: collect-261001-general-networking/general-networking/stockage-ps5-vs-xbox-carte-a-280-ssd-40-2026-2
title: "stockage-ps5-vs-xbox-carte-a-280-ssd-40-2026"
domain: general-networking
role: reference
task: reference
actors: ["Microsoft", "Samsung"]
dates: []
keywords: ["foundry"]
source: docs/RAG/collect-261001-general-networking/stockage-ps5-vs-xbox-carte-a-280-ssd-40-2026.md
source_anchor: ""
source_lines: [39, 93]
sha256: 0521dd47ffa2e29aac60fd0e3a2c6a02f88761a4826d31672f4c0d594f3c8bbc
---

# stockage-ps5-vs-xbox-carte-a-280-ssd-40-2026

Cette architecture verrouillée a un avantage réel : aucune configuration à faire, aucun risque de choisir un SSD trop lent ou mal dimensionné. Elle a aussi un coût, littéralement, puisque le joueur Xbox ne peut pas profiter de la guerre des prix entre fabricants de SSD M.2 grand public, contrairement au joueur PS5.

## Tableau comparatif : PS5 vs Xbox, spécification par spécification

Le tableau ci-dessous synthétise les différences techniques majeures entre les deux écosystèmes de stockage en septembre 2026.

| Critère | PlayStation 5 / PS5 Pro | Xbox Series X\|S | 
|---|---|---|
| Stockage interne de base | 825 Go (environ 667 Go utilisables) | 1 To Series X (802 Go utilisables) / 512 Go Series S | 
| Type d’extension officielle | SSD M.2 NVMe standard, emplacement interne | Carte d’extension propriétaire (port externe dédié) | 
| Interface | PCIe Gen4 x4, NVMe | PCIe via architecture Xbox Velocity (propriétaire) | 
| Fabricants disponibles | Samsung, Western Digital, Crucial, et autres marques NVMe | Seagate, WD_BLACK (licenciés Microsoft uniquement) | 
| Capacités proposées | 250 Go à 4 To | 512 Go, 1 To, 2 To | 
| Vitesse de lecture recommandée | ≥ 5 500 Mo/s | Environ 2 à 2,4 Go/s (équivalent au SSD interne) | 
| Dissipateur thermique requis | Oui, obligatoire (intégré ou ajouté) | Intégré à la carte, non configurable | 
| Jeux nouvelle génération sur USB externe | Non (stockage seulement, doit être rapatrié) | Non (stockage seulement, doit être rapatrié) | 
| Jeux ancienne génération sur USB externe | Oui (jeux PS4 jouables directement) | Oui (Xbox 360, One, rétrocompatibles) | 
| Facilité d’installation | Moyenne (démontage d’une trappe, vissage) | Très simple (branchement direct comme une clé USB) | 
| Concurrence sur les prix | Forte (plusieurs marques, promotions fréquentes) | Faible à modérée (marché encore restreint) | 
| Compatible si on change de console | Oui, tout SSD M.2 conforme fonctionne sur n’importe quelle PS5 | Oui, la carte se transfère d’une Xbox Series à une autre | 

Ce tableau met en évidence la différence philosophique fondamentale entre les deux constructeurs. Sony ouvre son écosystème à la concurrence du marché des SSD grand public, avec les avantages (choix, baisse des prix, capacités jusqu’à 4 To) et les inconvénients (nécessité de vérifier soi-même la compatibilité, montage physique) que cela implique. Microsoft verrouille son écosystème pour garantir une expérience sans erreur possible, au prix d’une offre plus restreinte et généralement plus chère au gigaoctet.

## Vitesses réelles : ce que mesurent les bancs d’essai

Les chiffres constructeurs donnent une idée du potentiel théorique, mais ce sont les tests indépendants qui comptent pour un acheteur. Plusieurs médias spécialisés, dont Tom’s Guide France et le guide comparatif de L’Éclaireur Fnac, ont comparé disque dur externe et SSD sur PS5, et la conclusion est sans appel : pour jouer dans de bonnes conditions à un titre PS5 natif, le SSD reste la seule option viable, le disque dur classique n’étant pertinent que pour l’archivage de la ludothèque PS4.

Sur le segment des SSD M.2 pour PS5, les modèles les plus populaires en France en 2026, à savoir le Samsung 990 Pro, le WD Black SN850X, le Crucial T500 et, plus récemment, le WD_BLACK SN850P repéré par Rosenberry Rooms en septembre 2026, affichent des vitesses annoncées par le fabricant entre 7 300 et 7 450 Mo/s en lecture séquentielle. Dans des conditions réelles, une fois installés dans la console, ces mêmes SSD atteignent généralement entre 6 500 et 6 800 Mo/s, un écart normal lié à l’overhead du système PS5. La différence de temps de chargement entre ces disques et le SSD interne d’origine reste, dans l’immense majorité des cas testés sur des jeux comme Horizon ou Ratchet & Clank, inférieure à une seconde. Autrement dit, au-delà d’un certain seuil de performance, l’écart devient invisible pour le joueur.

Côté Xbox, les cartes d’extension affichent des débits mesurés plus modestes en valeur brute, autour de 2 à 2,4 Go/s en séquentiel, avec un débit compressé de l’ordre de 2,4 à 2,8 Go/s grâce à l’architecture Velocity. C’est nettement inférieur aux chiffres des meilleurs SSD PS5, mais le point important, confirmé par les analyses techniques reprises notamment par TechRadar, est que ce débit est calibré pour être strictement identique à celui du SSD interne de la console. Il n’existe donc pas, côté Xbox, de gain de performance possible en changeant de support de stockage : la carte d’extension ne fait que reproduire fidèlement les performances internes, ni plus, ni moins.

Les analyses techniques associées à la sphère Digital Foundry, généralement relayées via Eurogamer, confirment ce constat sur la parité de performance entre stockage interne et carte d’extension Xbox : les temps de chargement observés entre les deux supports diffèrent de quelques dixièmes de seconde à peine, une variation qui relève davantage du bruit de mesure que d’un écart significatif. Pour la carte WD_BLACK C50, les premiers retours évoquent également des performances jugées irréprochables, sur un pied d’égalité avec la carte Seagate et le stockage interne, sans pénalité perceptible.

En résumé sur ce point : côté PS5, mieux vaut viser un SSD annoncé à 7 000 Mo/s ou plus pour avoir une vraie marge de confort ; côté Xbox, la question de la vitesse ne se pose quasiment pas puisque toutes les cartes licenciées offrent la même performance calibrée par Microsoft.

## Combien coûte l’extension de stockage en septembre 2026

C’est le nerf de la guerre. Voici les fourchettes de prix relevées début septembre 2026 chez les revendeurs français, pour les principales références disponibles sur chaque plateforme. Ces montants intègrent déjà l’impact de la hausse des prix de la mémoire flash évoquée plus haut, une tendance confirmée par le lancement en avril 2026 du SanDisk Extreme Portable SSD compatible PS5, facturé 25 080 ¥ en 1 To et 38 280 ¥ en 2 To selon Kakaku.com, pour un débit séquentiel annoncé de 1 000 Mo/s en USB 3.2 Gen2, une garantie de 5 ans et un poids plume de 69,4 g.

| Produit | Capacité | Prix indicatif (Sept. 2026) | Plateforme | 
|---|---|---|---|
| Carte d’extension Seagate | 1 To | 150 – 170 € | Xbox Series X\|S | 
| Carte d’extension Seagate | 2 To | 260 – 300 € | Xbox Series X\|S | 
| Carte WD_BLACK C50 | 512 Go | 90 – 110 € | Xbox Series X\|S | 
| Carte WD_BLACK C50 | 1 To | 150 – 180 € | Xbox Series X\|S | 
| Carte WD_BLACK C50 | 2 To | 250 – 300 € | Xbox Series X\|S | 
| Samsung 990 Pro (dissipateur) | 2 To | 160 – 200 € | PS5 / PS5 Pro | 
| Samsung 990 Pro (dissipateur) | 4 To | 300 – 380 € | PS5 / PS5 Pro | 
| WD Black SN850X (dissipateur) | 2 To | 150 – 190 € | PS5 / PS5 Pro | 
| Crucial T500 (dissipateur) | 2 To | 140 – 170 € | PS5 / PS5 Pro | 
| Crucial T700 PCIe 5.0 (dissipateur) | 2 To | 200 – 260 € | PS5 / PS5 Pro | 
| Disque dur externe USB générique | 2 To | 65 – 90 € | PS5 et Xbox (rétrocompatibilité seulement) | 

Le constat saute aux yeux : à capacité égale de 2 To, le SSD M.2 PS5 le moins cher (Crucial T500, autour de 140 à 170 €) coûte structurellement moins cher que la carte Xbox la moins chère équivalente (WD_BLACK C50 ou Seagate, entre 250 et 300 €). L’écart dépasse régulièrement 100 € en faveur de la PS5 sur ce segment de capacité, un delta qui s’explique par l’absence totale de concurrence directe sur les cartes propriétaires Xbox, contrairement au marché ouvert des SSD M.2 où Samsung, Western Digital et Crucial se livrent une bataille permanente de prix, en particulier lors des soldes et du Black Friday.

