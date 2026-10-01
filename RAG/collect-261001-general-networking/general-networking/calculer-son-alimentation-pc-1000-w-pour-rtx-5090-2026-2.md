---
id: collect-261001-general-networking/general-networking/calculer-son-alimentation-pc-1000-w-pour-rtx-5090-2026-2
title: "Exemple : RTX 5090 + Ryzen 9 9950X3D"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["gpu"]
source: docs/RAG/collect-261001-general-networking/calculer-son-alimentation-pc-1000-w-pour-rtx-5090-2026.md
source_anchor: ""
source_lines: [55, 107]
sha256: 4cf0ddbb121da487f6c1c7108ea2b06e0b771319c07b87f07b01b0c454b9c44c
---

# Exemple : RTX 5090 + Ryzen 9 9950X3D

Le premier est la **répartition de puissance par rail**, indiquée dans un petit tableau appelé “étiquette de puissance” (power label), généralement imprimé sur le côté du boîtier du PSU ou disponible dans le manuel PDF téléchargeable. Sur un modèle single rail, vous verrez une seule ligne “+12V” avec un ampérage maximal élevé (par exemple 83A sur un modèle 1000 W). Sur un modèle multi-rail, vous verrez plusieurs lignes +12V1, +12V2, etc., chacune avec son propre ampérage limité : c’est cette information qui vous permettra, le cas échéant, de savoir sur quel rail brancher votre GPU sans risquer un déclenchement prématuré de la protection.

Le deuxième élément est le **nombre et le type de connecteurs PCIe disponibles**. Un PSU ATX 3.1 destiné aux RTX 50 doit proposer au moins un connecteur 12V-2×6 natif (parfois listé “PCIe 5.0” ou “600W 12V-2×6” dans la fiche produit). Méfiez-vous des modèles qui annoncent la compatibilité RTX 50 uniquement via un adaptateur fourni dans la boîte : ce n’est pas équivalent à un câble natif en matière de fiabilité thermique, comme expliqué plus haut.

Le troisième élément est la **courbe de rendement en fonction de la charge**, parfois publiée par les fabricants sérieux (Seasonic, Corsair, be quiet!) sous forme de graphique. Un bon PSU atteint son rendement maximal entre 40 % et 60 % de sa charge nominale. C’est une raison supplémentaire de ne pas systématiquement acheter le modèle le plus puissant disponible : une alimentation 1600 W utilisée en permanence à 300 W (moins de 20 % de charge) fonctionne dans une zone de rendement dégradée, alors qu’une alimentation 850 W bien dimensionnée pour le même usage tournerait autour de 35-40 % de charge, une zone beaucoup plus efficace.

## Alimentation SFX vs ATX : quel format pour votre boîtier

La méthode de calcul du wattage reste identique quel que soit le format physique de l’alimentation, mais le choix du format a une influence directe sur les modèles disponibles et sur le budget. Le format **ATX** (150 x 86 x 140 mm environ) reste la référence pour les boîtiers standards et moyens-tours : c’est le format qui offre le plus grand choix de puissances, de la 450 W à la 1600 W, avec le meilleur rapport prix/watt.

Le format **SFX** et sa variante allongée **SFX-L** (125 x 63,5 x 100 mm environ) ciblent les boîtiers compacts type mini-ITX. Leur taille réduite impose une densité de composants plus élevée, ce qui se traduit généralement par un ventilateur plus petit tournant plus vite, donc un niveau sonore plus élevé à puissance égale, et un prix supérieur de 20 à 40 % par rapport à un modèle ATX équivalent. Pour une configuration compacte avec une RTX 5070 ou 5070 Ti, un modèle SFX-L 750-850 W Gold reste tout à fait viable ; au-delà, pour une RTX 5090 en boîtier mini-ITX, vérifiez soigneusement l’espace disponible car peu de modèles SFX dépassent 1000 W sur le marché européen en 2026.

Un dernier point pratique souvent négligé : dans un boîtier compact, la longueur des câbles fournis avec un PSU SFX peut être un facteur limitant. Vérifiez la longueur du câble 12V-2×6 avant achat si votre GPU se trouve loin de l’emplacement de l’alimentation, en particulier sur les boîtiers avec compartiment PSU séparé.

## Étape 1 à 3 : lister et additionner la consommation de vos composants

Voici la première partie concrète de la méthode de calcul, en trois étapes.

**Étape 1 : notez le TGP exact de votre GPU.** Reportez-vous au tableau ci-dessus ou à la fiche technique officielle du fabricant. N’utilisez jamais un chiffre approximatif trouvé sur un forum : le TGP officiel est la seule donnée fiable pour un calcul de sécurité.

**Étape 2 : additionnez la consommation du CPU.** Un processeur haut de gamme comme un Ryzen 9 9950X3D ou un Core Ultra 9 consomme généralement entre 150 W et 250 W en pleine charge avec le turbo actif, selon le réglage PBO ou le profil de refroidissement. Un CPU milieu de gamme (Ryzen 7, Core i5/i7) se situe plutôt entre 65 W et 125 W.

**Étape 3 : ajoutez la plateforme.** Comptez environ 80 à 150 W pour l’ensemble carte mère, RAM, SSD NVMe, ventilateurs de boîtier et éclairage RGB. Une configuration avec plusieurs SSD, un système de watercooling AIO et un éclairage RGB complet se situe en haut de cette fourchette.

Exemple concret pour une configuration RTX 5080 : GPU 360 W + CPU 200 W + plateforme 120 W = **680 W** de consommation théorique cumulée sous charge maximale simultanée (un scénario rare mais qu’il faut couvrir).

## Étape 4 à 6 : appliquer la marge de sécurité et choisir la certification

**Étape 4 : ajoutez une marge de 30 à 40 % pour les pics transitoires.** Les GPU modernes, en particulier les RTX 50, produisent des pics de consommation instantanés (transients) qui dépassent largement le TGP moyen pendant quelques millisecondes. Des tests indépendants ont mesuré des pics système autour de 1044 W sur une configuration RTX 5090 dont le TGP carte seule est de 575 W. Sans cette marge, l’alimentation peut déclencher sa protection de surintensité (OCP) et couper brutalement le PC.

**Étape 5 : reprenez notre exemple RTX 5080.** 680 W de base x 1,35 (marge de 35 %) = 918 W. On arrondit à la puissance commerciale supérieure disponible, soit **850 W ou 1000 W** selon la marge de confort souhaitée. C’est cohérent avec la recommandation constructeur de 850 W pour une RTX 5080.

**Étape 6 : choisissez la certification 80 PLUS.** Pour un PC gamer 2026, visez au minimum le niveau Gold. Une alimentation mieux certifiée chauffe moins, dure plus longtemps et consomme moins d’électricité sur la facture annuelle, un argument qui compte davantage avec la hausse du prix de l’énergie en Europe.

## Comprendre les certifications 80 PLUS : tableau des rendements

La certification 80 PLUS garantit un rendement énergétique minimum mesuré à 20 %, 50 % et 100 % de charge. Plus le rendement est élevé, moins d’énergie est perdue sous forme de chaleur entre la prise secteur et vos composants. Voici les seuils officiels par palier.

| Certification | Rendement à 20 % | Rendement à 50 % | Rendement à 100 % | 
|---|---|---|---|
| 80 PLUS (standard) | 80 % | 80 % | 80 % | 
| 80 PLUS Bronze | 82 % | 85 % | 82 % | 
| 80 PLUS Silver | 85 % | 88 % | 85 % | 
| 80 PLUS Gold | 88 % | 90 % | 87 % | 
| 80 PLUS Platinum | 90 % | 92 % | 89 % | 
| 80 PLUS Titanium | 92 % (+90 % à 10 %) | 94 % | 90 % | 

Pour une configuration au-delà de 1000 W (typiquement une RTX 5090), privilégiez Platinum ou Titanium : les pertes de rendement à cette puissance représentent plusieurs dizaines de watts dissipés en chaleur, ce qui sollicite davantage la ventilation du boîtier.

L’écart de rendement entre Bronze et Platinum se traduit aussi directement sur la facture d’électricité. Pour un usage intensif (jeu 4 heures par jour à pleine charge GPU), la différence de pertes entre un modèle Bronze et un modèle Platinum sur une configuration RTX 5090 représente environ 15 à 25 kWh consommés en plus par an avec le modèle Bronze, soit quelques euros par an au tarif réglementé en France. L’écart reste modeste sur une seule année, mais un PSU Platinum ou Titanium dure généralement plus longtemps et chauffe moins l’intérieur du boîtier, ce qui réduit indirectement la charge de travail des autres ventilateurs.

## Étape 7 à 9 : single rail vs multi rail et choix du modèle

