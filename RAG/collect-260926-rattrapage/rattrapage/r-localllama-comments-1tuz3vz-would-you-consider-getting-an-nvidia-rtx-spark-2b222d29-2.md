---
id: collect-260926-rattrapage/rattrapage/r-localllama-comments-1tuz3vz-would-you-consider-getting-an-nvidia-rtx-spark-2b222d29-2
title: "r-localllama-comments-1tuz3vz-would-you-consider-getting-an-nvidia-rtx-spark-2b222d29"
domain: rattrapage
role: reference
task: reference
actors: ["AMD", "Alibaba", "Apple", "Qualcomm"]
dates: []
keywords: ["llama", "nvidia", "amd", "gpu", "llama.cpp", "open source"]
source: docs/RAG/lot-rattrapage/servers-reviews/r-localllama-comments-1tuz3vz-would-you-consider-getting-an-nvidia-rtx-spark-2b222d29.md
source_anchor: ""
source_lines: [9, 51]
sha256: 3b7d41a93fb23addbaaf6afc9b13e85127de082ad557e1f448fe05e19b133935
---

# r-localllama-comments-1tuz3vz-would-you-consider-getting-an-nvidia-rtx-spark-2b222d29

    Si oui, pourquoi ? Si non, dis aussi pourquoi.
J'envisagerais un si c'est plus rapide en inférence AI locale que mon matériel actuel et qu'il peut encore gérer le jeu de manière décente.
L'idée de la mémoire unifiée de 128 Go est plutôt intéressante, mais je ne suis pas sûr pour Windows sur Arm et la compatibilité des jeux.
Achèterais-tu un, ou resterais-tu avec un ordinateur portable RTX x86 normal ?
Section des commentaires
Seulement si je peux mettre Linux dessus au lieu de Windows.
Faire de l'inférence AI sur un ordinateur portable semble une mauvaise idée à cause de la chaleur, mais qu'est-ce que j'en sais
J'attendrai les avis... puis j'achèterai un MacBook M5 de toute façon
M6 bientôt
Je pensais à ça aussi.
Non. Avec une bande passante mémoire de 300 Go/s, c'est dans une position étrange où il ne peut ni bien jouer à des jeux ni exécuter des inférences avec des vitesses de token utilisables. Vous n'obtiendrez vraiment des tk/s utilisables qu'à 600 Go/s ou plus.
Pour jouer à des jeux dessus, vous aurez également besoin de jeux qui fonctionnent bien sur ARM. Je ne suis pas sûr de la façon dont ça se passe. Obtenir un ordinateur portable normal, avec un bon GPU, sera le meilleur choix pour tous les cas d'utilisation.
Non, parce que c'est complètement stupide de payer plus de 4 000 $ pour 128 Go de RAM à 300 Go/s.
Pour référence, des barrettes ECC DDR5-4800 de 16 Go coûtent moins de 200 $ chacune. Pour simplifier, disons 200 $. 12 barrettes coûtent 2400 $ et vous donnent 192 Go de RAM. Un bundle avec 96 cœurs Epyc Genoa + carte mère est à environ 2000 $ sur eBay. Cela fait 4400 $ pour 192 Go de RAM avec... 460 Go/s de bande passante mémoire. Ajoutez une alimentation, un dissipateur thermique et un boîtier pour 600 $, pour un total d'environ 5000 $. Ça pourrait être légèrement plus cher qu'un ordinateur portable Spark, mais vous obtenez 50 % de mémoire en plus et 50 % de bande passante mémoire en plus.
Pas mal, on ne peut pas contester les chiffres, mais est-ce que ça peut faire tourner Crysis????
La comparaison la plus probable sera presque certainement contre un M5 MBP. Ça dépendra du prix du RTX Spark (il pourrait être en dessous de 5000 $ pour un M5 Max MBP entièrement équipé), mais le M5 Max va également jusqu'à 128 Go de RAM, a plus de bande passante mémoire, et probablement plus de performance GPU, et suffisamment d'efficacité énergétique pour réellement gérer n'importe quel type de charge de travail sur batterie.
Aussi, des rumeurs sur le M6 et le Macbook Ultra sont à l'horizon.
Si je veux un vrai bon laptop et pas une boîte d'inférence portable bâclée, le MBP va gagner facilement. La première génération de RTX Spark va être vraiment sous-dimensionnée en termes de... trucs normaux de laptop. Peut-être qu'nVidia pourra régler les problèmes avec de l'argent illimité dans les générations futures, mais je ne suis pas vraiment en train de lire les feuilles de thé à ce sujet en ce moment. Et de toute façon, tu ne peux pas acheter du matériel futur.
Si tu veux vraiment du matériel nvidia Spark, achète un DGX Spark, mets-le quelque part avec du réseau, et prends un macbook neo avec quel que soit le budget restant pour faire des choses de laptop.
Aussi, je veux dire que j'ai des doutes sur les CPU MediaTek, car MediaTek est plutôt connu pour les téléphones Android
déchets électroniquesbudget, mais apparemment, leur haut de gamme est maintenant compétitif avec Qualcomm ? Je n'ai pas vraiment suivi de près puisque les téléphones phares du marché américain n'utilisent pas MediaTek lol.
TLDR : Non.
Non, seulement l'ordinateur portable Strix Halo. Open source et x86. Et de très bonnes critiques.
C'est à peu près la même chose qu'un DGX Spark sans l'incroyable interconnexion. La bande passante de la mémoire sera la même, la RAM atteindra des niveaux similaires. Attendez-vous à ce qu'il soit thermiquement limité dans la plupart des formats de portables.
Je ne veux pas d'un chauffage d'appoint à 5000 $ sur mes genoux.
J'y ai réfléchi pour être honnête, mais ce n'est pas le meilleur rapport qualité-prix. Les AMD Ryzen™ AI Max+ 395 actuellement, j'ai le beelink, mais framework ou corsair sont super aussi.
Je suis un utilisateur de MacBook quand je ne suis pas sur mon PC, je ne suis pas vraiment un gamer donc... ouais, pour l'inférence de l'IA absolument. C'est bien sûr si l'argent n'était pas un problème lol
Même sans avoir discuté du coût, faire une sorte de calcul scientifique sur Windows est une non-starter en ce qui me concerne.
S'ils avaient sorti avec Debian sur ARM, j'aurais envisagé.
Je ne comprends vraiment pas. Ça devrait être un système pour des inférences locales d'IA mais la bande passante est très faible. Beaucoup de mémoire ne vous permettra pas d'utiliser des modèles plus grands à cause de la faible bande passante, ce qui les rendra très lents. Si ça coûte plus qu'un AMD Strix Halo, alors ça n'aura aucun sens, sinon juste du marketing. J'espérais une bande passante plus grande, au moins comparable à celle du nouveau Mac ou même meilleure, mais ce n'est pas le cas. Je ne comprends vraiment pas.
J'attends aussi avec impatience d'en acheter un pour ma première plateforme llm autre qu'un PC normal.
Si c'est ta première, je ne le vois pas comme une mauvaise option. En plus, tu as accès aux cœurs CUDA. Pas une mauvaise idée. ET tu peux jouer dessus !!!
*Vérifie le portefeuille*
Ouais....non.
trop cher, avec une mémoire unifiée pas si rapide et déjà infecté par le meilleur logiciel espion et la bloatware de ms ? C'est un non pour moi :D
Mouais, je préfère juste me connecter en SSH à une machine serveur pour des charges de travail comme ça. Le codage agentique / l'inférence lourde est trop lourd pour du portable.
Ça a l'air assez intéressant et si l'IA fonctionne vraiment et est aussi locale qu'ils le prétendent, c'est vraiment un énorme pas en avant pour les LLMs locaux. Cependant, je ne joue pas sur mon ordinateur portable et les LLMs locaux vont réduire l'autonomie de ta batterie en déplacement. Je vais simplement utiliser mon propre serveur LLM pour faire l'inférence pendant que l'ordinateur portable restera frais et presque silencieux pendant de nombreuses heures.
J'aimerais en avoir un avec Linux
Non. La raison : le coût et ce que vous obtenez.
J'ai un spark basé sur GB10 et ses 20 cœurs ARM sont lents. Beaucoup plus lents qu'un PC normal. Vous pouvez le voir compiler llama.cpp environ 3 fois plus longtemps que mon PC Ryzen. (Je vois aussi qu'un cœur CPU est saturé à 100 % quand il infère, et je pense que cela indique qu'il y a une limitation du débit d'inférence due au CPU aussi lors de l'utilisation de llama.cpp.)
Deuxièmement, la bande passante. Ce n'est pas compétitif, et trop lent pour un appareil de 2026. Pour le même prix, vous pouvez déjà acheter une machine avec plus de deux fois la bande passante RAM dans le M5 Max, un appareil de 128 Go. Donc, je pense que vous pouvez obtenir le double de la vitesse d'inférence d'une machine disponible aujourd'hui, et c'est vraiment le problème. Si vous arrivez avec un GPU à 300 Go/s à des prix niveau Apple, vous pouvez simplement remettre votre truc directement sur l'étagère. Je n'en veux pas.
Donc non, la proposition de valeur est à mon avis terrible. Aussi much que j'aime les ordinateurs portables ARM en principe et que je n'ai rien contre le CPU, j'ai besoin de l'ARM fait par Apple plutôt que de l'ARM fait par nvidia. Ils semblent être à la traîne.
Et si j'obtiens un appareil Apple, je ferai probablement tourner llama.cpp dessus. Il devrait faire fonctionner ces mêmes UD-Q8_K_XL Qwen3.6-27b en 8 bits que j'utilise déjà sur GB10. L'expérience est assez agréable, sauf pour le taux de génération de jetons qui ne l'est pas. Et ce problème est dû à la bande passante qui n'est pas corrigée.
