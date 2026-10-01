---
id: collect-250926-servers-hardware/servers-hardware/nvidia-dsx-40-de-gpu-par-usine-a-ia-2026-1
title: "nvidia-dsx-40-de-gpu-par-usine-a-ia-2026"
domain: servers-hardware
role: reference
task: reference
actors: ["Nvidia"]
dates: []
keywords: ["gpu", "nvidia", "attention", "blackwell", "energy", "mai", "open source", "rubin", "vera rubin"]
source: docs/RAG/clean4/nvidia-dsx-40-de-gpu-par-usine-a-ia-2026.md
source_anchor: ""
source_lines: [1, 47]
sha256: a205cc72a08db164e739ce181dd114c48ffe00808d40541b4adcdc28a4159e83
---

# nvidia-dsx-40-de-gpu-par-usine-a-ia-2026

Le 31 mai 2026, lors de la conférence GTC Taipei, NVIDIA a dévoilé **DSX**, une plateforme présentée comme le « mode d’emploi » complet pour construire des **usines à IA** (AI factories). L’annonce intervient quelques jours seulement après des résultats trimestriels records – 81,62 milliards de dollars de chiffre d’affaires au premier trimestre de l’exercice 2027 – et confirme un basculement industriel : le centre de données n’est plus un simple hébergeur de serveurs, mais une chaîne de production de « tokens » d’intelligence artificielle dont l’unité de mesure est désormais le mégawatt, voire le gigawatt.

Pour la France et l’Europe, le sujet est loin d’être anecdotique. Deux des partenaires industriels phares de DSX – le français **Schneider Electric** et l’éditeur **Dassault Systèmes** – figurent au cœur de l’écosystème. Au même moment, des projets d’infrastructure géants se dessinent sur le continent, du hub de Dunkerque aux ambitions de calcul souverain. Cette analyse décortique ce qu’est réellement NVIDIA DSX, ce que la plateforme change pour le marché, et pourquoi la bataille des usines à IA va structurer la décennie technologique européenne.

## NVIDIA DSX : qu’est-ce qu’une « usine à IA » exactement ?

NVIDIA DSX est une plateforme de conception, de simulation, d’exploitation et d’optimisation d’usines à IA couvrant l’intégralité de la pile technologique : des puces et des systèmes jusqu’aux logiciels d’infrastructure, en passant par les installations physiques et les technologies des partenaires. Selon le constructeur, l’objectif est simple à énoncer mais redoutablement difficile à atteindre : produire de l’intelligence artificielle au **coût du token le plus bas possible**.

La notion d’« usine à IA » mérite une définition précise. Contrairement à un centre de données classique, qui loue de la capacité de stockage et de calcul à des usages variés, une usine à IA est une installation conçue d’un seul tenant pour entraîner et faire tourner des modèles d’intelligence artificielle. Sa production se mesure en tokens générés par seconde, sa contrainte première est l’énergie disponible, et son rendement dépend autant de la climatisation et de la topologie réseau que de la puissance brute des processeurs graphiques. C’est précisément cette complexité que DSX entend industrialiser.

« NVIDIA est la seule entreprise qui construit l’usine à IA complète », a déclaré Jensen Huang, fondateur et directeur général de NVIDIA, en présentant la plateforme. La formule résume l’ambition stratégique : ne plus seulement vendre des cartes graphiques, mais fournir une référence d’ingénierie de bout en bout que les opérateurs – hyperscalers, États, industriels – peuvent dupliquer pour bâtir des installations de plusieurs centaines de mégawatts en limitant les erreurs de conception coûteuses.

## Les six briques de la plateforme DSX

DSX n’est pas un produit unique mais un ensemble modulaire de composants logiciels et de références matérielles. L’annonce du 31 mai 2026 détaille six briques principales, chacune répondant à une étape du cycle de vie d’une usine à IA, de la planche à dessin à l’exploitation quotidienne.

| Composant DSX | Rôle | Bénéfice clé revendiqué | 
|---|---|---|
| DSX Reference Design | Plan de référence de l’usine à IA complète | Réduit les erreurs de conception et le temps de déploiement | 
| DSX Sim | Jumeau numérique pour simuler énergie, refroidissement et réseau | Test virtuel avant construction physique | 
| DSX MaxLPS | Optimisation du débit de tokens par mégawatt | Jusqu’à 40 % de GPU en plus à budget énergétique constant | 
| DSX Flex | Gestion flexible de la charge et de l’alimentation | Adaptation aux contraintes du réseau électrique local | 
| DSX Exchange | Place de marché de composants et modules certifiés | Approvisionnement standardisé auprès des partenaires | 
| DSX OS | Système d’exploitation open source de l’usine à IA | Ordonnancement, résilience et multi-tenant unifiés | 

La brique la plus commentée est **DSX OS**, un système d’exploitation modulaire et open source qui prend en charge la gestion du cycle de vie, l’ordonnancement, la cohérence d’exécution, l’automatisation de la supervision, la résilience et l’exploitation multi-locataires. En ouvrant ce socle logiciel, NVIDIA cherche à reproduire la stratégie qui a fait le succès de CUDA : verrouiller l’écosystème non pas par la fermeture, mais par l’ubiquité d’une couche logicielle que tout le monde finit par adopter.

## De 100 mégawatts à plusieurs gigawatts : la nouvelle échelle

Le changement d’échelle est vertigineux. Selon NVIDIA, DSX est conçu pour soutenir des usines à IA allant de **100 mégawatts à plusieurs gigawatts**. À titre de comparaison, un centre de données traditionnel de grande taille consomme rarement plus de quelques dizaines de mégawatts. Une usine à IA d’un gigawatt équivaut donc à la consommation électrique d’une ville moyenne, et nécessite une planification énergétique digne d’un projet industriel lourd.

Cette inflation des besoins explique pourquoi NVIDIA a intégré dès la conception des partenaires de l’énergie comme GE Vernova, Hitachi, Siemens Energy et la jeune pousse Emerald AI. L’enjeu n’est plus seulement d’acheter des GPU, mais de sécuriser des dizaines de mégawatts d’électricité, de concevoir des systèmes de refroidissement liquide à très haute densité, et de raccorder l’ensemble au réseau sans le déstabiliser. Le jumeau numérique DSX Sim permet précisément de modéliser ces contraintes avant le premier coup de pioche.

Pour l’Europe, où le foncier industriel est rare et où l’acceptabilité sociale des centres de données énergivores fait débat, cette approche par la simulation pourrait s’avérer décisive. Un opérateur capable de prouver, chiffres à l’appui, la performance énergétique de son installation avant construction dispose d’un argument de poids face aux autorités locales et aux régulateurs.

## DSX MaxLPS : 40 % de GPU en plus à budget énergétique constant

Le chiffre qui a retenu l’attention des analystes est celui de **DSX MaxLPS**. NVIDIA affirme que cette technologie permet aux opérateurs de faire fonctionner jusqu’à **40 % de GPU supplémentaires** à leur point de fonctionnement le plus efficace sur le plan énergétique, avec un impact minimal sur la performance des charges de travail, à budget de puissance fixe.

Cette promesse touche au nerf de la guerre. Dans un monde où l’électricité est devenue la ressource limitante de l’industrie de l’IA, gagner 40 % de capacité de calcul sans consommer un watt de plus revient à reculer d’autant la contrainte physique qui plafonne la croissance. En d’autres termes, MaxLPS transforme une limite énergétique en levier économique : pour un même raccordement, l’opérateur produit davantage de tokens, donc davantage de revenus.

### Pourquoi le « token par mégawatt » devient l’indicateur roi

L’industrie a longtemps raisonné en FLOPS (opérations en virgule flottante par seconde) ou en téraoctets de mémoire. DSX impose un nouveau référentiel : le débit de tokens par mégawatt. Cet indicateur réconcilie la performance de calcul et la réalité physique de l’alimentation électrique. Il explique pourquoi les contrats d’approvisionnement en énergie, l’efficacité du refroidissement et la densité par baie comptent désormais autant que la génération de la puce. C’est un changement de paradigme comptable autant que technique.

## Vera Rubin : le successeur de Blackwell attendu au second semestre 2026

