---
id: collect-260926-rattrapage/rattrapage/r-localllama-comments-1r569eb-psa-nvidia-dgx-spark-has-terrible-cuda-software-111adb08-2
title: "r-localllama-comments-1r569eb-psa-nvidia-dgx-spark-has-terrible-cuda-software-111adb08"
domain: rattrapage
role: reference
task: reference
actors: ["AMD", "Apple", "Nvidia", "vLLM"]
dates: []
keywords: ["llama", "nvidia", "amd", "blackwell", "claude", "fp4", "fp8", "gpu", "llama.cpp", "nvfp4", "vllm"]
source: docs/RAG/lot-rattrapage/servers-reviews/r-localllama-comments-1r569eb-psa-nvidia-dgx-spark-has-terrible-cuda-software-111adb08.md
source_anchor: ""
source_lines: [34, 63]
sha256: d0a6c813cf15d3bf6db4afb187ed770de8ab92a92e198c566cff6c7be58b60ba
---

# r-localllama-comments-1r569eb-psa-nvidia-dgx-spark-has-terrible-cuda-software-111adb08

C'est fondamentalement une architecture différente qui porte le même nom. sm100 ne supporte même pas (mx/nv)fp4 dans les instructions mma traditionnelles et ne les supporte que dans tcgen05, alors que sm120 n'a pas du tout tcgen05.
C'est en partie pour ça que cette génération d'architecture graphique pour le jeu est si mal supportée dans beaucoup de nouvelles librairies par rapport au passé. Hopper a implémenté wgmma de manière exclusive, mais c'était largement basé sur la même idée que le mma traditionnel, avec un support asynchrone et d'exécution coopérative supplémentaire, alors que tcgen05 ressemble plus à du matériel dédié externe au SM.
Est-ce que le jetson Thor serait un meilleur choix ?
La fragmentation sm121 est le problème le plus intéressant ici. La gamme d'accélérateurs de NVIDIA compte désormais suffisamment de variantes d'architecture pour que votre code CUDA ne soit pas portable sur l'ensemble de leur gamme de produits. Pour tous ceux qui exécutent de l'inférence en production, cela signifie que votre cible de déploiement compte autant que le choix de votre modèle. La mémoire unifiée est vraiment utile pour les grandes fenêtres contextuelles, mais vous êtes enfermé dans un écosystème logiciel qui est encore à la traîne par rapport au matériel.
Quand j'ai fait mes recherches, j'ai choisi Strix Halo en partie parce que les points forts de Spark ne semblaient pas si forts, surtout pour le prix - c'était il y a 3 ou 4 mois, donc les choses ont peut-être bougé, à l'époque Spark coûtait presque le double et était en fait moins bon pour la plupart de mes utilisations, alors que maintenant c'est plutôt 60-70% plus cher (et quand même je préférerais avoir le Strix Halo même au même prix)
Je réévaluerai pour la prochaine génération quand Medusa Halo sortira, pour l'instant pour moi Spark est un peu bizarre dans l'espace mémoire unifié, avec un prix aussi élevé que les Mac Studios avec plus de mémoire et un environnement très mature, et avec les GPU nVidia haut de gamme qui sont la solution pour CUDA et généralement de hautes performances - donc si vous êtes un chercheur avec des besoins spécifiques, vous voudrez probablement toujours aller vers les GPU complets, bien que pour le moment vous soyez bien sûr complètement hors du marché pour 100 Go+ de GPU - par rapport à un Spark relativement accessible, mais encore une fois, vous saurez mieux si le produit est fait pour vous ou non
Mdr, c'est un titre de post scandaleux. Genre, le Spark, c'est en fait une Nintendo Switch reconditionnée ou un truc du genre.
nvfp4 est-il déjà pris en charge ?
Quand j'ai acheté mon Spark en octobre, j'étais aussi frustré que toi, mais maintenant, c'est beaucoup mieux. Côté architecture, c'est le même Blackwell que les RTX6000 Pro et RTX5090, mais un code arch sm121 séparé n'aide vraiment pas avec la compatibilité logicielle.
Cela dit, c'est quand même un super petit appareil, surtout si tu as un cluster.
Le problème principal, c'est qu'Nvidia a été super louche avec la fiche technique quand/avant la sortie du DGX Spark. Ils ont carrément caché plein de détails importants et se sont appuyés sur le battage médiatique et les influenceurs tech pour refourguer le dgx à des consommateurs qui ne se doutaient de rien. J'ai fait un post là-dessus puis, et il y a une longue discussion sur llama.cpp sur ses performances par rapport à Thor, étant donné que ce dernier est beaucoup moins cher.
Je ne suis toujours pas convaincu par ce qu'ils proposent, donc je suis toujours curieux de voir à quoi vous l'utilisez, et quelle est sa valeur ajoutée par rapport aux appareils Ryzen AI Max+ 395 qui coûtent moins de la moitié de son prix. Je comprends que cuda est un écosystème mature, mais pour la plupart des utilisations de llm, l'inférence AMD est presque au même niveau. En plus, le Thor a l'air beaucoup plus intéressant (avec un CPU plus lent, mais je suppose que le cas d'utilisation principal n'est pas d'héberger des services docker).
Mets à jour vers le dernier kernel (6.17 enfin sorti il y a quelques jours) pour les problèmes de moniteur.
Les forums gb10 sont le meilleur endroit pour les infos.
Le support sm121 arrive, ainsi que des corrections pour nvfp4.
Retourne le tien et récupère ton argent si tu peux pas attendre/gérer les solutions de contournement.
Ça répond bien à tous mes besoins.
J'en ai un et j'en suis assez content, mais soyons honnêtes. Le support sm121 et nvfp4 est "En cours" depuis 6 mois...
NVidia s'attend à ce qu'on fasse le gros du boulot.
Le plus drôle, c'est de payer la taxe NVIDIA spécifiquement pour la compatibilité CUDA et de ne pas l'avoir. Tout l'argument de vente pour choisir NVIDIA plutôt qu'Apple Silicon ou Strix Halo, c'est l'écosystème logiciel, et s'il est cassé, vous payez en gros plus pour moins. Pour l'inférence LLM pure, le Mac Studio est imbattable à des prix similaires depuis un moment maintenant.
Nan, même le M3 Ultra à 10 000 $ a juste de la mémoire rapide, mais un calcul fp32 de 26 TFlops au rythme d'un escargot et rien d'autre. Pas de fp16 ou fp8. Strix halo a 46 TFlops de f16, DGX a plus de 100TF en f16 et les maudits 400-500TF en nvfp4. Même la RTX 3060 a 40TF en f16, ce qui est juste embarrassant.
Le problème, c'est que tout le monde et sa mère benchmaxxing ces appareils en leur demandant d'écrire une histoire ou du code sans aucun contexte, sans outils ou en continuant réellement la conversation. Personne ne fait ça en pratique. Tu vas dépasser les 10 000 tokens en un rien de temps avec la recherche web ou l'entrée d'images/vision, et la vitesse de génération s'effondre rapidement. Mais personne ne fait ça parce que c'est plus dur que de coller une seule invite encore et encore, et ils aiment voir les personnages bouger vite à l'écran, au lieu de réellement construire quelque chose.
J'attends la deuxième génération de ces systèmes. Mais je me demande combien de temps. Probablement beaucoup d'années :)
C'est sûr que quelqu'un avec du claude code va vite régler ça /s
T'as tout à fait raison de souligner ça !
CUDA marche plutôt bien maintenant, la plupart du temps. Le code de l'architecture sm121 est vraiment une galère, parce qu'en termes de fonctionnalités, c'est pareil que sm120, à part la mémoire unifiée.
vLLM marche bien, le seul problème qui reste, c'est le support correct de NVFP4, mais ils sont dessus.
Pourquoi diable brancherais-tu ta bougie sur un moniteur ? Il n'existe pas de GPU pour le gaming. On dirait que tu n'as aucune idée de ce dont tu parles.
Carrément d'accord. J'ai eu plein de problèmes pour faire tourner du code ML de base, même en utilisant les conteneurs fournis, en compilant à partir des sources, les nightly builds, etc. C'est super frustrant comparé à la plupart des autres sorties de GPU qui ont un support dès le premier jour de cuda/nvcc/pytorch/tensorflow.
Cela dit, je pense que c'est un mauvais produit pour le ML pur. J'utilise les cœurs de ray tracing et en fait, c'est le truc qui me fait acheter. Mais sans support logiciel, ça ne sert à rien pour l'instant.
