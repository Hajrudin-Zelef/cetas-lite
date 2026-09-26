---
id: collect-260926-rattrapage/rattrapage/r-localllama-comments-1r569eb-psa-nvidia-dgx-spark-has-terrible-cuda-software-111adb08-1
title: "r-localllama-comments-1r569eb-psa-nvidia-dgx-spark-has-terrible-cuda-software-111adb08"
domain: rattrapage
role: reference
task: reference
actors: ["AMD", "Apple", "Google", "Nvidia", "vLLM"]
dates: []
keywords: ["llama", "nvidia", "amd", "blackwell", "diffusion", "gpu", "llama.cpp", "nvfp4", "vllm"]
source: docs/RAG/lot-rattrapage/servers-reviews/r-localllama-comments-1r569eb-psa-nvidia-dgx-spark-has-terrible-cuda-software-111adb08.md
source_anchor: ""
source_lines: [1, 33]
sha256: bf8da3474a4e1fd77339f312ac1c535da3f1437d3c76e06589400da1f21516b0
---

# r-localllama-comments-1r569eb-psa-nvidia-dgx-spark-has-terrible-cuda-software-111adb08

Merci pour ton avis !
Explique-nous pourquoi ce contenu n’est pas utile.
Avis important : La puce NVIDIA DGX Spark présente une compatibilité CUDA et logicielle déplorable ; elle ressemble davantage à une puce pour consoles de jeux portables.
J'ai passé la semaine dernière à expérimenter avec le DGX Spark et je suis sur le point de le renvoyer. Même si j'avais compris les limitations de la bande passante mémoire et des performances, j'aime l'écosystème CUDA et j'étais prêt à payer le prix fort. Malheureusement, mes expériences ont été assez mauvaises, et je soupçonne qu'il s'agit en fait de restes de jeux sur console portable que NVIDIA s'est empressé de transformer en produit pour concurrencer Apple et Strix Halo.
Le plus gros problème : DGX Spark n'est pas Blackwell pour centres de données, ce n'est même pas Blackwell pour le jeu, il a sa propre architecture spéciale sm121. Beaucoup de logiciels ne fonctionnent pas avec, ou ont été patchés pour exécuter des codepaths sm80 (Ampere, vieux de 6 ans !) ce qui signifie qu'il ne profite pas des optimisations blackwell.
Interrogé à ce sujet sur le forum d'assistance NVIDIA, un représentant officiel de NVIDIA a dit:
      Les kernels de classe sm80 peuvent s'exécuter sur DGX Spark car le comportement des Tensor Cores est très similaire, en particulier pour les GEMM/MMA (plus proche du modèle MMA de style GeForce Ampere). DGX Spark n'a pas tcgen05 comme jetson Thor ou GB200, en raison de l'espace de la puce avec les RT Cores et l'algorithme DLSS
Pardon ?? La raison pour laquelle nous obtenons des tensor cores réduits (pas de vrai blackwell) est à cause des RT Cores et de "l'algorithme DLSS" ? C'est un kit de développement IA ; pourquoi aurais-je besoin de RT Cores, et en plus, comment le DLSS entre-t-il en jeu ? Cela me fait penser qu'ils ont essayé de transformer un GPU de jeu portable (qui a besoin/prend en charge la mémoire unifiée) en un piètre concurrent pour un marché auquel ils n'étaient pas préparés.
De plus, dans le même message, le représentant a posté ce qui semble être des hallucinations de LLM, mentionnant que des problèmes ont été corrigés dans des numéros de version et des versions de bibliothèques logicielles qui n'existent pas.
Soyez juste prudent lorsque vous achetez un DGX Spark. Vous n'obtenez pas vraiment une expérience CUDA moderne. Oui, tout fonctionne bien si vous faites semblant de n'avoir qu'un Ampere, mais tenter d'utiliser les fonctionnalités de Blackwell est un exercice d'inutilité.
De plus, pour quelque chose qui est censé être prêt 'out of the box', beaucoup de gens (y compris moi-même et servethehome) signalent des problèmes de base comme la sortie d'affichage HDMI. Je pensais à l'origine que mon Spark était DOA ; non ; il refuse juste de fonctionner avec mon viewsonic 1080p144 (qui fonctionne avec tous les autres GPU ; y compris mes NVIDIA) ; et j'ai dû passer à mon moniteur 4K60. Cher NVIDIA, vous ne devriez pas avoir de problèmes de sortie d'affichage de base...
Section des commentaires
C'était un peu galère quand j'ai déballé le premier. J'ai cherché sur Google, mis à jour un package, compilé un binaire, et quelques heures plus tard, c'était parti.
Clé en main ? Non. Mais c'est pas non plus Slackware Linux.
Quelques semaines plus tard, quand le deuxième est arrivé, j'ai pu faire un ‘apt-get’ de ce dont j'avais besoin (cuda13 pour vLLM pour avoir le bon compilateur).
Je n'y connais pas assez en puces pour savoir si les accusations sont vraies, mais je dirais que j'obtiens une plus grande gamme de capacités plus facilement en ayant CUDA disponible que sur mon Mac (meilleure génération d'images / vidéos, entre autres) mais je peux aussi clairement voir des contre-arguments.
Merci pour ton retour d'expérience, je suis surpris que ce soit pire avant.
Je suppose que j'avais juste de plus grandes attentes en matière de compatibilité CUDA complète et de support logiciel NVIDIA, étant donné le prix élevé et les affirmations marketing / fiches techniques de NVIDIA. Je voulais expérimenter le pré-entraînement NVFP4 d'un petit LLM, et je n'ai toujours pas de solution viable.
La transformation de Hadamard, qui est essentielle à la stabilité de l'entraînement NVFP4 (selon NVIDIA), n'étant pas disponible sur sm120/sm121, c'était comme un autre coup de poing ; surtout après avoir appris que NVIDIA propose bien DC blackwell (avec des cœurs tensoriels complets) dans le Jetson Thor pour à peu près le même prix et le même format.
Pourquoi ont-ils utilisé blackwell grand public (avec sm121) pour DGX spark au lieu de ce qu'ils ont utilisé dans le Jetson Thor ? On dirait qu'ils avaient construit des puces GB10 pour un but différent (les cœurs RT et DLSS sont sans intérêt pour la recherche en IA), et qu'ils avaient besoin de se débarrasser des GB10.
Les fiches techniques de DGX Spark et de Jetson Thor indiquent toutes les deux "NVIDIA Blackwell GPU with fifth-gen Tensor Core technology". Sauf qu'elles sont loin d'être les mêmes. Je m'y serais attendu si c'était AMD/ROCm, mais j'essaie de payer la taxe CUDA pour ne pas avoir à gérer ça :)
Pour moi, le plus gros problème, c'est que si tu t'intéresses à la génération d'images/vidéos, à cause d'un mauvais support soit dans ComfyUI, soit dans la librairie transformers utilisée par Comfy, charger un safetensor prend 2 fois plus de mémoire ! Le modèle est chargé une fois dans la "RAM" puis aussi dans la "VRAM", ce qui bouffe 2 fois plus de mémoire. Donc t'as pas un "GPU" de 120 Go comme tu le penses, t'en as un de 60 Go. C'est vraiment la m.
Y'a pas ces problèmes côté LLM avec llama.cpp et vllm. Pour les LLM, tu peux même connecter 2 Sparks ensemble avec un câble QSFP, charger un modèle de 240 Go (ou plus petit, je suppose), et la vitesse augmente au lieu de "se dégrader moins".
S'ils réparaient le problème avec les modèles de diffusion, ce serait une machine très respectable.
Oui ! C'est aussi un des problèmes que j'ai rencontrés en jouant un peu ; j'ai remarqué une double utilisation de la RAM pour ComfyUI. Je l'ai installé en utilisant leur playbook officiel.
Les produits NVIDIA sont généralement beaucoup mieux supportés que le DGX Spark. Ce n'est pas mon premier devkit NVIDIA, j'ai acheté des tonnes de Jetsons différents avant, au fil des décennies, pour bricoler. Je ne dis pas que c'est complètement inutile, mais ça ne ressemble pas à un produit NVIDIA.
Ni Strix Halo ni Apple n'ont ces problèmes.
C'est pas une puce pour console portable. On sait déjà qu'elle est quasiment identique à celle qu'ils vont mettre dans leurs laptops. Et on sait aussi que les laptops ont été retardés à cause de Windows. C'est pour ça qu'ils ont sorti ça entre-temps (pour récupérer un peu les frais de développement).
Cela dit, y'a des rumeurs qui ont surgi plus tard, qui disent qu'il y a plein de bugs dans la conception hardware de la version laptop. Et encore une fois, si c'est aussi bien identique, alors ils vont probablement exister aussi dans Spark.
Aha oui ! Ça aurait du sens comme puce pour un laptop, j'avais oublié que GB10 arrive sur les laptops.
Mais je crois que mon point tient toujours. DGX et NVIDIA ont une réputation et des attentes élevées en matière de qualité et de support fournisseur ; et le prix correspond à ces attentes. Je ne pense pas que ce produit devrait porter le moniteur DGX dans son état actuel.
Je peux gérer les problèmes sm121 et nvfp4 si NVIDIA les a résolus en logiciel. NVIDIA ne l'a pas fait, et beaucoup de leurs playbooks officiels sont cassés et ne peuvent pas être exécutés tels quels sur les dernières versions.
Le sm120 arch ne supporte pas non plus tcgen05. Différence majeure entre les cartes serveur Blackwell et leurs éditions workstation.
