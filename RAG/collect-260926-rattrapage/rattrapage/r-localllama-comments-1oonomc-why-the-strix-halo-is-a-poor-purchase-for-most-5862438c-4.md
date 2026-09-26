---
id: collect-260926-rattrapage/rattrapage/r-localllama-comments-1oonomc-why-the-strix-halo-is-a-poor-purchase-for-most-5862438c-4
title: "r-localllama-comments-1oonomc-why-the-strix-halo-is-a-poor-purchase-for-most-5862438c"
domain: rattrapage
role: reference
task: reference
actors: ["AMD", "Apple"]
dates: []
keywords: ["amd", "gguf", "gpu", "moe"]
source: docs/RAG/lot-rattrapage/ai-llm/r-localllama-comments-1oonomc-why-the-strix-halo-is-a-poor-purchase-for-most-5862438c.md
source_anchor: ""
source_lines: [158, 174]
sha256: b9730ba2f8424a516558ee1535f6332bed3f0efb3a715a68257f4c99eea1e3f8
---

# r-localllama-comments-1oonomc-why-the-strix-halo-is-a-poor-purchase-for-most-5862438c

Je m'en fous de la vitesse de ta bande passante mémoire, tu vas te faire défoncer. J'ai vu la même chose se produire sur les anciens Apple Studio M1 par rapport au 5090. Le 5090 déchire tout jusqu'à ce qu'il atteigne un mur. Et plus l'allocation de mémoire dépasse 32 Go, plus le 5090 souffre. La mémoire devient plus précieuse que la bande passante, parce que tu ne brasses pas constamment la mémoire, limité par le PCIe, ou que tu dois diviser le surplus vers les CPU qui ne peuvent tout simplement pas rivaliser avec les GPU.
Tu as trouvé un seul point de données et tu as décidé d'en faire une généralisation complète.
a) 5090 seul, ça coûte près de 3000 $ ces temps-ci, et avec le reste du système, les prix avoisinent les 4000 $.
b) gpt-oss-120b, c'est MOE. Forcément, ce sera plus rapide sur la 5090, car seule une petite partie est réellement chargée. Essaie maintenant un modèle dense de taille moyenne sur la 5090 et compare-le à l'AMD 395.
c) C'est quoi la machine Strix Halo ? Un laptop ou un miniPC ? Vu qu'il n'y a aucune info. Je demande parce qu'il y a un écart de perf dû à la consommation entre un laptop (85C) et un MiniPC (140W).
d) C'est quoi les chiffres quand Lemonade est utilisé pour l'exécution hybride (iGPU + NPU) ?
gpt-oss-120b-mxfp-GGUF via Lemonade est supporté.
Y'a tellement de commentaires haineux. Changer sa config RTX 5090 pour une RTX 3090 et t'auras 70% de ses perfs pour -2000 dollars.
Alors, est-ce que les performances s'améliorent vraiment quand tu quantifies le cache kv en q8 ou moins ?
Le cache KV quant flingue les perfs de pré-remplissage pour moi. Je sais pas pourquoi !
J'ai un rig 5090, un de 7 litres en fait, donc c'est même pas si peu portable que ça qu'une boîte Strix Halo, mais ça n'a vraiment aucun sens de diviser le travail sur la mémoire principale du système, c'est juste un goulot d'étranglement énorme.
Les perfs du Strix Halo, comme beaucoup l'ont montré, s'améliorent et 30+tok/s sont atteignables avec un contexte large. Ça veut dire que c'est utilisable.
Je pense que si vous en avez besoin, ce serait vraiment bien, mais c'est la prochaine itération de ces puces Halo qui commencera vraiment à devenir intéressante. S'ils peuvent continuer à ajouter encore plus de canaux de mémoire, et bien sûr il y aura plus de calcul disponible, alors on commencera à voir 100tok/s sortir de ce modèle 120b et à ce stade, on parle de suffisamment rapide pour une utilisation générale.
Ce sera aussi tellement bien pour les algorithmes généraux du CPU de pouvoir exploiter toute cette bande passante mémoire. Une fois que vous commencez à dépasser la moitié d'un To/s, c'est un autre jeu.
Je pense aussi qu'une fois que les logiciels rattraperont leur retard, il y aura un avantage de réactivité pour les systèmes à mémoire unifiée qui pourront sauter le transfert de bus.
Ça veut dire que les jours où ça avait encore un sens de construire un PC de bureau dans un petit format sont comptés. Comme il se doit. L'unifié est juste logique.
On dirait que t'as juste aucune envie d'apprendre quoi que ce soit sur l'utilisation du strix halo. Peut-être qu'on devrait demander l'avis de gens qui veulent vraiment apprendre les choses correctement.
