---
id: collect-261001-general-networking/general-networking/fr-review-coolit-5d9994f0-2
title: "fr-review-coolit-5d9994f0"
domain: general-networking
role: reference
task: reference
actors: ["Intel"]
dates: []
keywords: ["dram", "gpu", "intel"]
source: docs/RAG/collect-261001-general-networking/fr-review-coolit-5d9994f0.md
source_anchor: ""
source_lines: [16, 22]
sha256: 5bf68a19ea2cbead30e7229548ba66b1ea2196805babd2cf1bd92cbdcba0be09
---

# fr-review-coolit-5d9994f0

Pour analyser plus précisément la consommation électrique du R760, nous avons configuré notre module d'analyse de consommation électrique Quarch QTL2843 . Nous avons testé le serveur avec ses dissipateurs thermiques à air d'origine, puis avec les plaques froides CoolIT. Afin de solliciter fortement les processeurs, nous avons effectué un calcul de Pi jusqu'à 50 milliards de décimales, ce qui impose une charge très importante au processeur et à la mémoire DRAM. Notre objectif était de pousser les processeurs au maximum afin de forcer le fonctionnement des ventilateurs à leur vitesse maximale.
L’impact de la mise en œuvre du DLC était immédiatement évident. Lors de l'exécution du R760 dans la configuration refroidie par air, les ventilateurs tournent à 100 % pendant la charge de travail, comme prévu. Avec la configuration DLC, le R760 a choisi de faire tourner les ventilateurs à 32%, une baisse spectaculaire. Cela équivaut à une économie de 200 watts sur un seul serveur. Ce n'est pas seulement la vitesse du ventilateur qui ressort, les processeurs eux-mêmes ont rapporté environ la moitié de la température avec le DLC, 41/42°C contre 88/89°C lorsqu'ils sont refroidis par air.
Mais ce ne sont pas seulement les économies d’énergie qui ressortent du refroidissement liquide. Nous avons constaté une petite amélioration des performances, à laquelle nous ne nous attendions pas. Grâce aux plaques froides offrant un meilleur refroidissement, le processeur peut fonctionner au maximum. Dans la configuration refroidie par air, le R760 a effectué le calcul de 50 milliards de Pi en 369 secondes. Dans la configuration DLC, le R760 est allé un peu plus vite, livrant le calcul en 347 secondes. Cela représente un gain de performances d'environ 6%, ce qui nous permet de tirer un peu plus parti des processeurs Intel.
Réflexions finales
Nous commençons tout juste le refroidissement liquide en laboratoire et sommes ravis d'avoir travaillé avec CoolIT sur ce premier effort. Les plaques froides fonctionnent parfaitement sur le PowerEdge R760 et le collecteur et le CDU s'assemblent et « fonctionnent » sans aucun souci ni bricolage continu. Pour ceux qui craignent d’introduire du liquide dans le centre de données, une simplicité continue est essentielle. Nous n'avons également eu aucune fuite ni autre événement plus catastrophique, ce qui était attendu : il s'agit d'un équipement d'entreprise avec un taux de défaillance extrêmement faible.
Pour les entreprises qui cherchent à intégrer des systèmes d’IA haute puissance dans leur centre de données, le refroidissement liquide est une évidence. Les serveurs GPU à 8 voies vont abandonner le refroidissement par air, optant pour des boucles DLC comme celle-ci ou au minimum, une boucle fermée et un radiateur. Quoi qu’il en soit, une certaine quantité de liquide s’infiltrera dans le centre de données. Avec les économies d'électricité substantielles et l'amélioration modeste des performances, il existe de nombreuses raisons pour lesquelles les entreprises devraient adopter des serveurs DLC.
CoolIT est un leader incontesté dans ce domaine et sa relation avec Dell met sur le marché une grande variété de solutions de refroidissement liquide d'une manière facilement consommable, avec très peu de soucis. Nous sommes impatients d'explorer davantage notre petite boucle et avons hâte de voir davantage de serveurs refroidis par liquide dans le laboratoire.
