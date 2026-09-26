---
id: collect-260926-rattrapage/rattrapage/r-localllama-comments-1gfvo69-mac-mini-m4-pro-64-gb-best-compromise-2c52f0dc
title: "r-localllama-comments-1gfvo69-mac-mini-m4-pro-64-gb-best-compromise-2c52f0dc"
domain: rattrapage
role: reference
task: reference
actors: ["Apple", "Cerebras", "Groq", "Nvidia"]
dates: []
keywords: ["llama", "benchmarks", "diffusion", "gpu", "llama.cpp", "nvidia"]
source: docs/RAG/lot-rattrapage/ai-llm/r-localllama-comments-1gfvo69-mac-mini-m4-pro-64-gb-best-compromise-2c52f0dc.md
source_anchor: ""
source_lines: [1, 51]
sha256: 2209c156aaa179566b51b978609f0c028e5524b54f6e1870a9fee4f3c4148eb9
---

# r-localllama-comments-1gfvo69-mac-mini-m4-pro-64-gb-best-compromise-2c52f0dc

Merci pour ton avis !
Explique-nous pourquoi ce contenu n’est pas utile.
      Mac Mini M4 Pro 64 Go : Meilleur compromis ? 
        
        
        
    
    
    Maintenant que nous connaissons toute la gamme de produits M4, je pense que la meilleure option pour mon utilisation serait d'opter pour un Mac Mini M4 Pro avec 64 Go de RAM. Cela me coûterait 2100 euros (j'achèterai via mon entreprise, sans payer de TVA en Europe).
Je l'utiliserais comme serveur d'inférence à distance (LLM pour le codage, Stable Diffusion) à partir de mon Macbook Pro 2019, et comme serveur NAS pour stocker mes photos.
Plus tard, j'achèterai un Macbook Air ou un Macbook Pro d'occasion, selon les offres à venir.
En comparaison, un Macbook Pro avec M4 Pro et 48 Go de RAM me coûterait 2813 euros. Un Macbook Pro avec M4 MAX et 64 Go de RAM me coûterait 3900 euros.
Qu'en penses-tu ? Ou devrais-je plutôt attendre des offres sur les Macbook Pro avec M3 Pro / Max ?
P.S : Je sais qu'un rig NVIDIA donnerait de meilleures performances / me coûterait moins cher, mais je ne veux pas mettre les mains dans le cambouis ces jours-ci, je recherche une solution robuste dans l'écosystème Apple avec lequel je suis habitué à travailler.
Section des commentaires
TLDR : Un Mac Studio M2 max avec GPU 38 cœurs de 64 Go peut être obtenu à un prix similaire et serait plus rapide
42 t/s contre 25 t/s pour le Mac mini M4
Le Mac Studio M1 Ultra avec GPU 64 cœurs de 64 Go obtient 60 t/s, donc ça vaut aussi le coup d'y réfléchir.
Jetez un œil ici
https://github.com/ggerganov/llama.cpp/discussions/4167
pour avoir une idée des tokens par seconde que vous pouvez obtenir en utilisant un modèle 7B sur différents Apple Silicon.
Les colonnes Q8_0 TG [t/s] et Q4_0 TG [t/s] représentent les tokens par seconde pour les quantifications Q8 et Q4.
Le Mac Mini M4 Pro n'a que 20 cœurs GPU. La puce la plus proche est le M2 Pro à 19 cœurs à 23,01 Q8.
Le cœur supplémentaire et la largeur de bande mémoire feront légèrement augmenter cela, mais pas de beaucoup. 25 t/s est probablement réaliste.
Alors, il semble que ça livre du lourd après tout ? https://9to5mac.com/2024/11/01/new-mac-mini-m4-pro-geekbench/
Ce sont des benchmarks CPU. En gros, c'est zéro pertinence par rapport à la façon dont ça va fonctionner en exécutant des modèles de langage, car tu utiliseras le GPU.
Pour la génération de tokens LLM, la bande passante mémoire prime sur la puissance CPU. Si c'est votre utilisation principale, alors obtenir un Ultra M1 ou M2 d'occasion sera plus rapide. Même pour des travaux intensifs en calcul, les cœurs GPU seraient inférieurs mais il y en a tellement plus.
Le M4 Pro est puissant pour de nombreuses tâches et vous pourriez inclure l'inférence dans le mix, mais si l'inférence est le but central, l'un ou l'autre Ultra gagnera facilement.
Quelqu'un sait si Stable Diffusion fonctionnerait bien sur un max d'occasion / un mini neuf / tout ce qui pourrait être dans la gamme de prix ? Et disons que je veux guider la génération avec une carte de profondeur. Je peux juste télécharger et utiliser une interface ouverte (je ne connais que SD Forge rn) et utiliser le modèle civitai que je veux sur l'écosystème Apple ?
Comment les M1 ultra et M2 ultra se comparent-ils pour les tps ?
Ma préoccupation concerne la puissance de traitement, pas la mémoire.
Ohh, élabore ? Je suis curieux.
Les GPU Nvidia ont des milliers de cœurs. Nous essayons de rivaliser avec 10 ou 20 cœurs. Bien qu'Apple offre plus de mémoire, le souci, c'est que les modèles plus grands nécessitent plus de puissance de calcul et nous pourrions être limités là-dessus. Je n'ai pas encore de chiffres solides.
Ça ne va pas être un très bon serveur d'inférence. Je te conseillerais une version Linux avec une nvidia 30909(s) si tu es sérieux.
Nous ne connaissons pas toute la gamme de produits m4. Nous ne le saurons pas avant quelques mois, lorsque le studio sera mis à jour vers m4 ultra.
Demandez à quelqu'un comme SomeOddCoderGuy comment faire fonctionner de grands modèles avec un long contexte sur Mac. Les modèles de 70B tiendraient dans 64 Go de RAM à Q4 ou Q5, mais vous pourriez attendre longtemps si vous avez beaucoup d'évaluations de prompts à faire.
Tous les milliers de cœurs CUDA sur RTX s'additionnent.
Seul le CPU "Max" a la bande passante mémoire pour des vitesses d'interférence à peu près correctes.
C'est ce que je pense. Tu peux faire tourner un modèle de 70 à 90 milliards dessus. Bien que je ne sois pas sûr du taux de tokens.
Ce ne sera pas bon.
Mon avis 12 t/s
Si un m1 fait 70b à 8-9 t/s, alors j'imagine que le m4 va être pas mal meilleur. Mais 70b pour la plupart des gens, ce n'est pas une amélioration significative par rapport à un modèle à 8b.
Peux-tu expliquer la qualité entre toutes les tailles de modèle ?
Compare avec 3x3090 (3*24Go=72Go).
Watts ?
On ne parle pas de watts en dehors d'Apple 😂. Ça ne va pas du tout donner une bonne image, tu sais.
Commentaire supprimé par le membre
Sd et llms fonctionnent incroyablement bien sur Apple Silicon.
Commentaire supprimé par le membre
En fait, l'architecture elle-même ne joue pas un rôle si important si la charge de travail bénéficie d'un parallélisme massif au lieu de la performance d'un CPU monocœur, ce qui est exactement le cas pour les calculs matriciels et la raison pour laquelle des dizaines de milliers de cœurs dans les GPU sont bénéfiques.
Édition : En faisant référence aux architectures CPU populaires comme x86 et RISC, pas aux nouvelles architectures propriétaires natives de matrice de Cerebras, SambaNova, Groq, etc.
