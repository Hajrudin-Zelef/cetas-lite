---
id: collect-261001-general-networking/general-networking/fr-review-asus-expertcenter-pro-et900n-g3-review-7a4e59ce-4
title: "fr-review-asus-expertcenter-pro-et900n-g3-review-7a4e59ce"
domain: general-networking
role: reference
task: reference
actors: ["Alibaba", "DeepSeek", "MiniMax", "Mistral", "Z.ai", "vLLM"]
dates: []
keywords: ["deepseek", "fp8", "glm", "leaderboard", "llama", "mistral", "mxfp4", "nvfp4", "vllm"]
source: docs/RAG/collect-261001-general-networking/fr-review-asus-expertcenter-pro-et900n-g3-review-7a4e59ce.md
source_anchor: ""
source_lines: [74, 97]
sha256: b87e421fada7d2e0f56b5775d1ae2d754b6bbce1a9678378d4ad9649b39e620e
---

# fr-review-asus-expertcenter-pro-et900n-g3-review-7a4e59ce

Sur les 28 paires modèle-charge de travail, 24 présentent un écart inférieur à 2 %, la plupart de ces écarts étant de quelques fractions de pour cent en faveur du WS300. Quatre se situent en dehors de cette fourchette, dont trois à charge de travail égale : Llama 3.1 8B FP8 avec un écart de 3.5 % par rapport au WS300 (25 602 contre 26 536 jetons de sortie par seconde), Qwen3 Coder 30B BF16 avec un écart de 4.6 % (10 000 contre 10 486) et Nemotron 3 Ultra 550B avec un écart de 5.2 % (159 contre 168), ainsi que Qwen3 Coder 30B BF16 à nouveau avec un écart de 2.1 % sur les invites longues. Chaque chiffre présenté ici correspond à un seul test sur chaque tour ; par conséquent, un écart de 2 à 5 % entre trois modèles sur 14 ne peut être imputé au châssis. Nous nous contentons de constater les données et l'écart entre les deux. Les modèles qui utilisent le pool de mémoire cohérent, MiniMax M3 et GLM-5.2, affichent des performances équivalentes, voire supérieures, pour les deux charges de travail ; c'est le résultat qui nous semble le plus important pour la plateforme : le chemin de déchargement Grace offre les mêmes performances dans le châssis ASUS que dans celui de MSI.
| Modèle | WS300 512/512 | ET900N G3 8 192/1 024 | Delta | WS300 8,192/1,024 | ET900N G3 8 192/1 024 | Delta | 
|---|---|---|---|---|---|---|
| GPT-OSS-20B MXFP4 | 22,161 | 22,041 | -0.5% | 9,572 | 9,555 | -0.2% | 
| GPT-OSS-120B MXFP4 | 9,294 | 9,280 | -0.2% | 5,038 | 5,023 | -0.3% | 
| Lama 3.1 8B BF16 | 17,278 | 17,074 | -1.2% | 3,520 | 3,498 | -0.6% | 
| Llama 3.1 8B FP8 | 26,536 | 25,602 | -3.5% | 6,189 | 6,164 | -0.4% | 
| Llama 3.1 8B NVFP4 | 28,698 | 28,400 | -1.0% | 6,744 | 6,737 | -0.1% | 
| Mistral Petit 24B BF16 | 7,922 | 7,861 | -0.8% | 1,643 | 1,633 | -0.6% | 
| Mistral Small 24B FP8 | 11,157 | 11,075 | -0.7% | 2,316 | 2,302 | -0.6% | 
| Codeur Qwen3 30B BF16 | 10,486 | 10,000 | -4.6% | 3,830 | 3,749 | -2.1% | 
| Codeur Qwen3 30B FP8 | 12,425 | 12,439 | + 0.1% | 3,851 | 3,837 | -0.4% | 
| Flash DeepSeek V4 FP8 | 1,766 | 1,766 | 0.0 % | 948 | 954 | + 0.6% | 
| MiniMax M2.7 NVFP4 | 4,801 | 4,793 | -0.2% | 1,743 | 1,744 | + 0.1% | 
| MiniMax M3 NVFP4 | 1,041 | 1,041 | 0.0 % | 278 | 282 | + 1.3% | 
| GLM-5.2 NVFP4 | 138 | 139 | + 0.4% | 118 | 120 | + 1.7% | 
| Nemotron 3 Ultra 550B NVFP4 | 168 | 159 | -5.2% | 142 | 140 | -1.1% | 
GPT-OSS-20B, le cas dense à débit le plus élevé de l'ensemble, et GLM-5.2, le cas de déchargement, montrent le chevauchement sans ratio : sur les deux, l'exécution du WS300 suit point par point celle de l'ET900N G3.
Nous étofferons cette comparaison au fur et à mesure que de nouvelles tours GB300 seront testées au laboratoire. La ZGX de HP sera la prochaine sur la liste, suivie d'autres modèles. Chaque tour sera soumise au même balayage, ce qui permettra de mettre en évidence les différences entre les implémentations des constructeurs, sur un graphique par tour, les graphiques par modèle présentant les courbes brutes.
Conclusion
L'ASUS ExpertCenter Pro ET900N G3 remplit parfaitement son rôle de seconde version d'une plateforme de référence : elle confirme les performances de la première. Sur 14 modèles et deux charges de travail, son débit vLLM maximal s'est avéré inférieur de seulement 2 % à celui du MSI XpertStation WS300 dans 24 des 28 comparaisons, passant de 1 766 jetons de sortie par seconde sur DeepSeek V4 Flash à 22 041 sur GPT-OSS-20B, avec quatre valeurs aberrantes ponctuelles inférieures de 2 à 5 %. De plus, il a géré GLM-5.2 avec 215 Go d'experts dans la mémoire Grace à des débits que quatre cartes RTX PRO 6000 n'ont pas pu atteindre. Le processeur GB300 Superchip établit la norme, et ASUS a réalisé un excellent travail.
Les améliorations apportées par ASUS sont pratiques : les poignées permettent à deux personnes de déplacer une tour de 27 kg sans avoir à saisir les panneaux ; le ventilateur placé au-dessus des cages QSFP112 résout un problème potentiel pour ceux qui utilisent des périphériques optiques 400G pendant des heures ; et l’alimentation Titanium réduit la consommation électrique de quelques watts. Le processeur Arm, le circuit de 20 A, la sortie vidéo uniquement gérée par le BMC et la puissance de 1 600 W partagée lorsqu’une carte graphique de grande taille est installée sont des contraintes techniques communes aux deux tours.
La station GB300 DGX reste le dispositif le plus performant que nous ayons jamais installé sur un bureau, et ASUS propose une version qui constitue un excellent moyen de l'acquérir.
Sur notre Meilleurs ordinateurs de bureau pour l'IA locale leaderboard, the ET900N G3 now stands alongside the WS300 as a tested alternative in the Best GB300 System slot.
