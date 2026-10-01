---
id: collect-250926-servers-hardware/servers-hardware/parts-de-marche-gpu-2026-nvidia-90-amd-a-8-2
title: "Windows (PowerShell)"
domain: servers-hardware
role: reference
task: reference
actors: ["AMD", "Intel", "Nvidia", "Samsung", "TSMC"]
dates: []
keywords: ["amd", "attention", "dram", "gpu", "hbm", "intel", "nvidia", "research"]
source: docs/RAG/clean4/parts-de-marche-gpu-2026-nvidia-90-amd-a-8.md
source_anchor: ""
source_lines: [29, 84]
sha256: 45ef2c4bc65fe5445f4695b23b0cec8db116e7561a4cee8ace411d11593382ef
---

# Windows (PowerShell)

Au-delà de la répartition des parts de marché GPU, c’est le volume global qui inquiète, même si la tendance s’est retournée depuis. Le marché des cartes graphiques dédiées avait reculé de **0,6 % sur le trimestre** au T1 2026, pour s’établir précisément à 11,82 millions d’unités, avec un taux d’attachement d’un GPU dédié désormais proche de 76 % côté bureau, selon les données de Jon Peddie Research reprises par PCCentral en juin 2026 – avant de rebondir de **12,2 % sur le trimestre suivant**, pour atteindre 12,5 millions de cartes livrées au deuxième trimestre 2026, selon Igor’sLAB. Un chiffre en hausse de 28,3 % sur un an au T1, mais qui traduisait surtout un point de comparaison faible début 2025, avant que le rebond du T2 ne vienne nuancer le récit d’une contraction continue.

Le tableau côté processeurs est encore plus sombre : les livraisons de CPU pour PC ont chuté de 15 % au premier trimestre, les processeurs de bureau plongeant même de 24 % sur le trimestre. Jon Peddie Research anticipe une croissance annuelle moyenne **négative de 3,3 %** pour le marché des cartes graphiques dédiées d’ici 2029. Autrement dit, le segment ne devrait pas retrouver ses volumes d’antan.

### Tableau : chiffres clés du marché des cartes graphiques

| Indicateur (source : Jon Peddie Research) | Valeur | 
|---|---|
| Cartes graphiques dédiées livrées (T1 2026) | 11,8 millions d’unités | 
| Évolution trimestrielle (T4 2025 → T1 2026) | −0,6 % | 
| Évolution annuelle (T1 2025 → T1 2026) | +28,3 % | 
| Livraisons totales de GPU de bureau en 2025 | 44,28 millions (+28 % vs 2024) | 
| Taux d’attachement d’un GPU dédié (bureau) | 76 % (+33 % sur le trimestre) | 
| Livraisons de processeurs PC (T1 2026) | −15 % (CPU de bureau : −24 %) | 
| Croissance annuelle moyenne du marché AIB → 2029 | −3,3 % | 
| Part la plus basse jamais atteinte par AMD/ATI | 5 % (T4 2025) | 

## La flambée de la mémoire, coupable désigné

Pourquoi le marché des cartes graphiques s’essouffle-t-il alors que la demande de puissance graphique n’a jamais été aussi forte ? La réponse tient en un mot : la mémoire. La flambée des prix de la DRAM et de la mémoire vidéo, alimentée par l’appétit insatiable des centres de données pour l’IA, renchérit brutalement le coût de fabrication de chaque carte. L’analyse de Jon Peddie Research est sans détour : le marché des cartes dédiées est « pris en étau » par le haut, à cause de la hausse des coûts et des droits de douane américains, et par le bas, à cause des ordinateurs portables et des processeurs à graphique intégré de plus en plus performants.

Cette crise de la mémoire, que nous avons documentée dans notre enquête sur la flambée du prix de la RAM (+171 %), frappe l’ensemble de la chaîne : barrettes DDR5, mémoire GDDR7 des cartes graphiques, mais aussi la mémoire HBM ultra-rapide réservée aux accélérateurs IA. Or, face à des marges bien plus généreuses sur la HBM, les fabricants de mémoire – Samsung, SK Hynix, Micron – arbitrent naturellement en faveur des puces destinées aux datacenters, au détriment du grand public.

Conséquence directe : une carte graphique milieu de gamme intègre aujourd’hui des composants dont le coût a bondi de plusieurs dizaines de pourcents en un an. Les constructeurs répercutent cette hausse sur les prix conseillés, ou rognent sur les volumes produits. Dans les deux cas, c’est le joueur qui paie l’addition.

## L’IA rebat les cartes : quand les GPU partent dans les datacenters

Le second grand responsable de ce déséquilibre s’appelle intelligence artificielle. Le cœur de métier de Nvidia n’est plus le jeu vidéo depuis longtemps : selon les estimations de Silicon Analysts, les revenus de sa division centres de données ont atteint **193,7 milliards de dollars** sur l’exercice fiscal 2026, clos en janvier 2026, portant sa part du marché des accélérateurs IA à environ **80 %** – une domination que les mêmes analystes projettent encore proche de **75 %** sur un marché mondial de l’IA désormais évalué à plus de 200 milliards de dollars pour 2026. La firme a d’ailleurs profité de sa keynote asiatique début juin pour confirmer la montée en production de sa plateforme IA de nouvelle génération, comme nous le détaillions dans notre article sur les usines à IA Nvidia DSX.

Dans ce contexte, chaque tranche de silicium gravée chez TSMC devient un arbitrage économique. Faut-il produire un accélérateur IA facturé plusieurs dizaines de milliers de dollars, ou une carte GeForce grand public vendue quelques centaines d’euros ? Le calcul est vite fait, et il ne concerne pas que Nvidia : les puces Instinct d’AMD, elles aussi destinées aux centres de données, n’ont généré qu’entre **7 et 8 milliards de dollars** de revenus en 2025, soit à peine 5 à 7 % du marché mondial des accélérateurs IA selon Silicon Analysts – un montant qui n’en incite pas moins Lisa Su à détourner vers l’IA une partie des ressources d’ingénierie qui auraient pu profiter à Radeon. Les capacités de production les plus avancées, la mémoire HBM et l’attention des ingénieurs se concentrent sur l’IA, reléguant le marché des cartes graphiques pour joueurs au rang de priorité secondaire.

Ce phénomène explique en partie pourquoi Nvidia peut se permettre de dominer le marché GPU grand public sans y consacrer l’essentiel de ses ressources : sa position est si solide qu’elle se maintient presque par inertie, pendant que l’entreprise concentre ses efforts sur la manne des datacenters.

## Tableau comparatif : parts de marché GPU 2025-2026

Le tableau ci-dessous retrace l’évolution des parts de marché GPU sur le segment des cartes graphiques dédiées de bureau, d’après les relevés trimestriels de Jon Peddie Research. Les valeurs sont arrondies et concernent les cartes de bureau (AIB).

| Fabricant | T1 2025 | T4 2025 | T1 2026 | 
|---|---|---|---|
| Nvidia (GeForce) | 92 % | 94 % | ~90 % | 
| AMD (Radeon) | 8 % | 5 % *(record bas)* | ~8 % | 
| Intel (Arc) | < 1 % | ~1 % | ~1 % *(+0,4 pt)* | 
| Total unités livrées | ≈ 9,2 M | ≈ 11,9 M | 11,8 M | 

Deux enseignements ressortent. D’abord, la remarquable stabilité de Nvidia autour de 90-94 % sur toute la période. Ensuite, le fait qu’Intel, malgré des volumes encore modestes, soit le seul à grignoter des points – au détriment aussi bien d’AMD que, marginalement, de Nvidia.

## Le paradoxe européen : AMD résiste chez les assembleurs

Ces parts de marché GPU mondiales méritent toutefois une lecture nuancée à l’échelle européenne. Les chiffres de Jon Peddie Research agrègent l’ensemble des livraisons, y compris celles destinées aux fabricants d’ordinateurs de marque (OEM), qui privilégient massivement Nvidia. Or le marché de l’assemblage « maison » (*do it yourself*), très vivace en Europe, raconte une tout autre histoire.

Chez le détaillant allemand Mindfactory, baromètre suivi de près par la communauté des monteurs de PC européens, AMD a historiquement capté une part bien supérieure à sa moyenne mondiale – parfois proche de la moitié des cartes vendues à l’unité. La raison est méthodologique : ces données reflètent les choix d’acheteurs éclairés, sensibles au rapport performance-prix et à la quantité de mémoire vidéo, là où les données globales sont écrasées par les commandes OEM et les ordinateurs portables.

Autrement dit, le joueur européen qui monte lui-même sa machine se montre nettement plus ouvert à Radeon que ne le suggèrent les 8 % mondiaux. Un contre-pied qui rappelle qu’AMD conserve une base de fidèles solide sur le Vieux Continent, même si elle ne suffit pas à peser sur les statistiques planétaires.

## Des prix qui s’envolent en France et en Europe

