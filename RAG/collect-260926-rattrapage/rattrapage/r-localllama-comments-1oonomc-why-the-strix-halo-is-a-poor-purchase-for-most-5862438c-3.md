---
id: collect-260926-rattrapage/rattrapage/r-localllama-comments-1oonomc-why-the-strix-halo-is-a-poor-purchase-for-most-5862438c-3
title: "r-localllama-comments-1oonomc-why-the-strix-halo-is-a-poor-purchase-for-most-5862438c"
domain: rattrapage
role: reference
task: reference
actors: ["AMD", "Nvidia", "Z.ai"]
dates: []
keywords: ["llama", "amd", "glm", "gpu", "llama.cpp", "moe", "nvidia"]
source: docs/RAG/lot-rattrapage/ai-llm/r-localllama-comments-1oonomc-why-the-strix-halo-is-a-poor-purchase-for-most-5862438c.md
source_anchor: ""
source_lines: [129, 157]
sha256: c4404befa32292b968467a24b34bbf8f6cb2ab1e0e20d7d7b20ddbd7b6d273c3
---

# r-localllama-comments-1oonomc-why-the-strix-halo-is-a-poor-purchase-for-most-5862438c

Ouais, sans vouloir t'offenser, mais tu conseilles aux gens de dépenser potentiellement des milliers de plus et de consommer beaucoup plus d'énergie pour faire de l'inférence... Je vois pas l'intérêt, dans mon cas j'ai un desktop super rapide et solide comme un roc, minuscule, qui est excellent pour le prix... ça me donne pas les vitesses d'inférence d'un modèle de pointe, mais c'est carrément dingue pour jouer avec des modèles locaux sans se ruiner. Je suis beaucoup plus content avec ça plutôt que de construire un desktop pour égaler les vitesses de ma Strix avec une 5090 en plus.
Je veux dire, faut un peu comparer des setups qui coûtent le même prix, non ?
La vraie réponse, c'est de brancher 4 5090 sur ta strix halo.
Je rigole, je rigole, 4 3090, ça le fait.
Ma config actuelle, c'est 128 Go de strix halo avec 2 3090 et une 4090 de 48 Go, ce qui me permet de charger des modèles plus gros comme GLM-4.6
Je suis passé de mon PC de bureau 2x3090 x 128 Go DDR5 à un Halo Strix et je ne pourrais pas être plus content. GLM 4.5 Air qui fait de l'inférence à 120w est plus rapide que le même modèle qui tourne sur mon PC de bureau de 800w. Et maintenant, mon PC est de nouveau libre pour jouer !
C'est du pipeau, ça ?
Plusieurs testeurs ont montré que le modèle OSS fonctionne nettement mieux sur le matériel Nvidia, mais les différences de performance sont moindres lors de l'utilisation de modèles sans les couches expertes.
Cela dit, ce week-end dernier, NewEgg n'avait pas de 5090 à moins de 3200 $, avant les autres composants, contre 2000 $ pour la Strix…
Je suis pas fan de Strix Halo et je trouve ça un peu cher et sur-vendu, mais la plupart des gens n'ont pas de RTX5090 et un système capable de faire tourner de la DDR5 6000.
Au fait, il y a eu un post il y a quelques jours avec un fork de llama.cpp qui améliore les performances avec l'augmentation du contexte.
Tu as considéré la RAM (100 Go/s) comme le goulot d'étranglement pour la vitesse, mais en réalité, c'est le pcie 5.0 à 64 Go/s. Ça va diminuer ta vitesse théorique nette, où 47% est servi depuis la RAM.
En plus, t'as pas calculé pour plusieurs KVCache, t'as juste pris 20k de contexte. Dans les tâches réelles, le contexte grandit beaucoup plus vite, et c'est ça le problème. Si tu pouvais calculer pour plusieurs tailles de contexte, ce serait plus juste. Genre 20k, 40k, 60k, 80k, 100k, 150k, 200k.
Ce que je trouve intéressant dans ce post, c'est qu'il s'intitule "Pourquoi le Strix Halo est un mauvais achat pour la plupart des gens", mais nulle part il n'explique pourquoi les gens envisageraient d'acheter un Halo et quels sont les cas d'utilisation les plus courants, voire même des cas d'utilisation, et en quoi c'est un mauvais choix pour la plupart de ces cas d'utilisation. Comment pouvez-vous dire que c'est un mauvais achat pour la plupart des gens sans établir ce que la plupart des gens qui pourraient l'acheter veulent pouvoir faire et quels problèmes et limitations ils pourraient rencontrer avec lui par rapport à une 5090 ou d'autres options ?
J'ai un Halo. J'ai aussi une 5090 (et une RTX Pro 6000 aussi). Pour l'usage pour lequel je l'ai acheté, le Halo est BEAUCOUP plus utile et considérablement plus rapide que la 5090. La 6000 pourrait bien sûr le détruire pour les mêmes utilisations, mais alors je gaspillerais la 6000 pour quelque chose que le Halo peut faire assez bien, et à quel point ce serait stupide ? La 5090 est également BEAUCOUP mieux adaptée aux autres tâches qu'elle effectue qu'à ce pour quoi j'utilise le Halo.
Votre argument ne soutient pas votre thèse, et vous ne comprenez clairement pas autant de choses sur ce sujet que vous voulez le croire... Vous connaissez peut-être quelques détails techniques, mais vous ne comprenez presque rien des MULTIPLES façons dont les gens peuvent utiliser ces outils, ce qui est d'une importance capitale pour ce sujet.
J'ai un minisforum s1 max qui arrive et j'ai hâte de le tester à fond.
Tu vas adorer. N'écoute pas ce type. Tous ceux que je connais avec des Strix Halo les adorent. AMD améliore ROCm de plus en plus. Ils viennent d'envoyer des Strix Halo aux mainteneurs de llamacpp pour qu'ils voient quelles optimisations de performances ils peuvent faire.
Dépenser encore 2 000 balles pour une 5090, c'est dingue. Tu peux littéralement pas battre le rapport qualité-prix d'un système Strix Halo. J'ai eu le mien pour 1 650 il y a un moment et c'est mon truc quotidien. En dehors de l'IA, j'ai 128 Go de RAM super rapide associés à un CPU presque aussi performant qu'un 9950. Même comme labo domestique, c'est une affaire de fou.
Ça donne quoi les chiffres avec qwen3 coder et un grand contexte ?
Intéressant. Dis-nous combien coûterait ton ordi + GPU aujourd'hui chez un grand détaillant.
Ce post spécifique a été supprimé via Redact. La motivation est inconnue, mais cela pourrait inclure la vie privée, la sécurité, l'opsec, ou un désir général de réduire son empreinte numérique.
dam paltry fact hungry jar aspiring head brave cooperative fear
D'accord avec la plupart.
Mais pour la dernière partie, je crois que Strix Halo + GPU a encore du potentiel. À mon avis, le comportement actuel lié à la bande passante PCIe est en fait dû à une implémentation inférieure de llama.cpp.
L'heuristique de base pertinente est la suivante : pour les poids MoE déchargés sur la RAM, si la taille du lot est petite (pour le décodage, c'est 1), alors c'est définitivement limité par la bande passante mémoire, donc on le calcule simplement sur le CPU. Si la taille du lot est plus grande (surtout le préremplissage), alors la complexité computationnelle surmontera la bande passante mémoire/PCIe, donc on transfère les poids vers le GPU.
Le plus gros problème ici est de savoir quelle est la taille de "grand" ? Actuellement, llama.cpp utilise un nombre très grossier : 32. Oui, c'est un nombre fixe, quelle que soit votre configuration CPU/GPU. Faisons quelques calculs rapides. Supposons que les 120b soient tous des paramètres MoE. Un lot de 32 éléments nécessitera exactement 4*32=128 multiplications d'experts, c'est-à-dire 120G OPs. Maintenant, la performance dépend du taux de réutilisation des experts, c'est-à-dire combien d'experts doivent être lus. Si l'utilisation des experts est répartie uniformément, alors nous devons lire 60 Go de données pour 120G OPs. Les CPU grand public modernes pourraient facilement faire des centaines de GFLOPS, donc cela ne vaut évidemment pas la peine d'envoyer les données via PCIe. En réalité, il y aura une certaine réutilisation des experts, donc la meilleure stratégie varie en fonction du modèle/de l'entrée/de la taille du lot. Il y a une PR dans lk_llama il y a des mois qui s'attaque à cela (https://github.com/ikawrakow/ik_llama.cpp/pull/520). Avec quelques ajustements de paramètres, ils peuvent obtenir ~2x de performances PP avec une petite taille de lot.
Maintenant, passons au cas de Strix Halo. En suivant les calculs ci-dessus, vous verrez qu'envoyer les poids via PCIe ne vaudra jamais la peine pour Strix Halo - même avec une taille de lot de 4096. Le GPU+NPU de Strix Halo a une capacité théorique de 126 TOPs, c'est-à-dire facilement ~100x plus rapide qu'un CPU grand public classique. Et sa bande passante RAM est ~4x la bande passante PCIe 5 x16. Ce serait fou d'envoyer les poids via PCIe au lieu de calculer in-situ dans la RAM.
Si tout ce que tu veux faire, c'est faire tourner des MoEs sur ton système qui sont à peine plus gros que tes 32 Go de RAM, alors oui, tu as raison. Mais disons que tu veux faire plus grand, quelque chose qui remplirait presque 96 Go de RAM.
