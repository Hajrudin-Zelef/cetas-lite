---
id: collect-260926-rattrapage/rattrapage/r-localllama-comments-1oonomc-why-the-strix-halo-is-a-poor-purchase-for-most-5862438c-2
title: "r-localllama-comments-1oonomc-why-the-strix-halo-is-a-poor-purchase-for-most-5862438c"
domain: rattrapage
role: reference
task: reference
actors: ["AMD"]
dates: []
keywords: ["llama", "gpu", "llama.cpp", "moe", "prefill"]
source: docs/RAG/lot-rattrapage/ai-llm/r-localllama-comments-1oonomc-why-the-strix-halo-is-a-poor-purchase-for-most-5862438c.md
source_anchor: ""
source_lines: [9, 128]
sha256: be07cfbe8487b903afb24f39accd6d3a5e05d8912890e0beb482c7dd3f69515a
---

# r-localllama-comments-1oonomc-why-the-strix-halo-is-a-poor-purchase-for-most-5862438c

    J'ai vu pas mal de posts qui font la promo du Strix Halo comme un bon achat, et je me suis souvent demandé si j'aurais dû l'acheter moi-même. J'ai appris beaucoup de choses sur la façon dont ces modèles sont exécutés depuis. Dans ce post, j'aimerais partager des mesures empiriques, d'où je pense que ces chiffres viennent, et expliquer pourquoi peu de gens devraient acheter ce système. J'espère que ça vous sera utile !
Modèle testé
- 
      llama.cpp
- 
      Gpt-oss-120b
- 
      Un des modèles de la plus haute qualité qui peut tourner sur du matériel milieu de gamme.
- 
      La taille totale de ce modèle est d'environ 59 Go, dont environ 57 Go de couches expertes.
Systèmes testés
Premier système :
- 
      128 Go Strix Halo
- 
      Quad channel LPDDR5-8000
Deuxième système (mon système) :
- 
      Dual channel DDR5-6000 + pcie5 x16 + une rtx 5090
- 
      Une rtx 5090 avec la plus grande taille de contexte nécessite environ 2/3 des experts (38 Go de données) pour résider dans la RAM du système.
- 
      cuda backed
- 
      mmap off
- 
      batch 4096
- 
      ubatch 4096
Voici les chiffres soumis par les utilisateurs pour le Strix Halo :
| test | t/s | 
|---|---|
| pp4096 | 1012.63 ± 0.63 | 
| tg128 | 52.31 ± 0.05 | 
| pp4096 @ d20000 | 357.27 ± 0.64 | 
| tg128 @ d20000 | 32.46 ± 0.03 | 
| pp4096 @ d48000 | 230.60 ± 0.26 | 
| tg128 @ d48000 | 32.76 ± 0.05 | 
Qu'est-ce qu'on peut en tirer ?
Les performances ne sont acceptables qu'avec un contexte de 0. Au fur et à mesure que le contexte grandit, les performances pp s'effondrent. Les performances tg subissent également un léger ralentissement.
Et voici les chiffres de mon système :
| test | t/s | 
|---|---|
| pp4096 | 4065.77 ± 25.95 | 
| tg128 | 39.35 ± 0.05 | 
| pp4096 @ d20000 | 3267.95 ± 27.74 | 
| tg128 @ d20000 | 36.96 ± 0.24 | 
| pp4096 @ d48000 | 2497.25 ± 66.31 | 
| tg128 @ d48000 | 35.18 ± 0.62 | 
Attends une seconde, comment les chiffres de décodage sont-ils si proches ? Le Strix Halo a une mémoire 2,5 fois plus rapide que mon système.
Regardons de plus près gpt-oss-120b. Ce modèle fait 59 Go. Il y a environ 0,76 Go de données de couche qui sont lues pour chaque token. Puisque chaque token a besoin de ces données, elles sont conservées dans la VRAM. Chaque token doit également lire 4 experts arbitraires, ce qui représente 1,78 Go supplémentaires. Étant donné que nous pouvons faire tenir 1/3 des experts dans la VRAM, cela porte la répartition totale à 1,35 Go dans la VRAM et 1,18 Go dans la RAM du système avec un contexte de 0.
Maintenant, la VRAM sur une 5090 est beaucoup plus rapide que la mémoire unifiée du Strix Halo et que le dual channel DDR5-6000. Au final, en effectuant ~53 % de vos lectures dans une VRAM ultra rapide et 47 % de vos lectures dans une RAM système un peu lente, le temps de décodage est très similaire avec de petites tailles de contexte par rapport à l'exécution de toutes vos lectures dans la mémoire modérément rapide du Strix Halo.
Pourquoi le Strix Halo a-t-il un ralentissement lorsque le contexte de décodage augmente ?
Probablement parce que lorsque la taille de votre contexte augmente, le décodage doit également lire le plus grand cache KV.
Et pourquoi mon système subit-il moins de ralentissement à mesure que le contexte augmente ?
Vous pouvez voir que, bien qu'avec un contexte de 0, le Strix Halo ait une avance en tg, il chute rapidement une fois que vous avez du contexte à traiter et que mon système gagne. C'est parce que tout le cache KV est stocké dans la VRAM, qui a des lectures de mémoire ultra rapides. Le temps de décodage est dominé par la lecture lente de la mémoire dans la RAM du système, donc cela ne fait pas bouger l'aiguille.
Pourquoi les temps de préremplissage se dégradent-ils si rapidement sur le Strix Halo ?
Bonne question ! J'adorerais savoir !
Puis-je simplement ajouter un GPU à la machine Strix Halo pour améliorer mon préremplissage ?
Malheureusement non. La capacité à utiliser un GPU pour améliorer les temps de préremplissage dépend fortement de la bande passante pcie et le Strix Halo n'offre que pcie x4.
Mesures réelles de l'effet de la bande passante pcie sur le préremplissage
Ces tests ont été effectués en modifiant les paramètres du BIOS sur ma machine.
| config | prefill tps | 
|---|---|
| pcie5 x16 | ~4100 | 
| pcie4 x16 | ~2700 | 
| pcie4 x4 | ~1000 | 
Pourquoi la bande passante pci est-elle si importante ?
Voici ma meilleure compréhension de haut niveau de ce que llama.cpp fait avec un gpu + cpu moe :
- 
      Tout d'abord, il exécute le routeur sur les 4096 tokens pour déterminer quels experts il faut pour chaque token.
- 
      Chaque token utilisera 4 des 128 experts, donc en moyenne chaque expert sera mappé à 128 tokens (4096 * 4 / 128).
- 
      Ensuite, pour chaque expert, téléchargez les poids sur le GPU et exécutez-les sur tous les tokens qui ont besoin de cet expert.
- 
      Cela en vaut la peine car le préremplissage est gourmand en calcul et l'exécuter simplement sur le CPU est beaucoup plus lent.
- 
      Ce processus est pipeliné : vous téléchargez les poids pour le token suivant, tout en exécutant le calcul pour le courant.
- 
      Maintenant, tous les experts pour gpt-oss-120b font ~57 Go. Cela prendra ~0,9 s pour le télécharger en utilisant pcie5 x16 à son maximum de 64 Go/s. Cela place un plafond en pp de ~4600tps.
- 
      Pour pcie4 x16, vous n'obtiendrez que 32 Go/s, donc votre maximum est d'environ ~2300tps. Pour pcie4 x4 comme le Strix Halo via occulink, c'est 1/4 de ce nombre.
- 
      En pratique, aucun des deux n'obtiendra sa pleine bande passante, mais les ratios absolus se maintiennent.
Autres avantages d'un ordinateur normal avec une rtx 5090
- 
      Meilleur refroidissement
- 
      Boîtier de meilleure qualité
- 
      Une 5090 aura presque certainement une valeur de revente plus élevée qu'une machine Strix Halo
- 
      Plus extensible
- 
      CPU plus puissant
- 
      Gaming haut de gamme
- 
      Les modèles qui tiennent entièrement dans la VRAM se décoderont également plusieurs fois plus vite qu'un Strix Halo.
- 
      La génération d'images sera beaucoup plus rapide.
À quoi sert le Strix Halo
- 
      Consommation d'énergie au repos extrêmement faible
- 
      C'est petit
- 
      Peut-être que tout ce qui vous intéresse, ce sont les chatbots avec presque 0 contexte
TLDR
Si vous pouvez vous permettre 1000 à 1500 $ de plus, vous feriez beaucoup mieux de simplement construire un ordinateur avec une rtx 5090. Le rapport qualité-prix est tellement plus fort. Même si vous ne voulez pas dépenser ce genre d'argent, vous devriez vous demander si votre cas d'utilisation est réellement couvert par le Strix Halo. Peut-être ne rien acheter à la place.
Corrections
Veuillez me corriger sur tout ce que j'ai mal compris ! Je ne suis qu'un novice !
EDIT :
WOW ! Le kit ddr5 que j'ai acheté en juin a vu son prix doubler depuis que je l'ai acheté. Peut-être que 50 % de plus est maintenant une sous-estimation.
Section des commentaires
Tes chiffres pour Strix Halo sont pas bons. Voici mes derniers chiffres gpt-oss-120b sur llama.cpp avec ROCm 7.10 :
Info de dernière minute : les ordinateurs plus chers sont plus rapides que les ordinateurs moins chers. Plus d'infos à 23h.
Ça doit être des chiffres qui datent. Parce que le Strix Halo est meilleur que ça maintenant et s'améliore de jour en jour. Voici un nouveau test qui vient de se terminer il y a une minute.
Bien sûr, même si le Strix Halo ne peut pas espérer avoir la puissance de calcul pour rivaliser avec la 5090 pour PP. En TG, j'ose dire qu'il se bat à armes égales avec la 5090. Même avec un contexte important.
