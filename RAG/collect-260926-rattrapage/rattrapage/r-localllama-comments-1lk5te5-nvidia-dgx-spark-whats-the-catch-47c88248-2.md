---
id: collect-260926-rattrapage/rattrapage/r-localllama-comments-1lk5te5-nvidia-dgx-spark-whats-the-catch-47c88248-2
title: "r-localllama-comments-1lk5te5-nvidia-dgx-spark-whats-the-catch-47c88248"
domain: rattrapage
role: reference
task: reference
actors: ["AMD", "Apple", "Nvidia"]
dates: []
keywords: ["nvidia", "amd", "fp4", "gpu", "int4"]
source: docs/RAG/lot-rattrapage/servers-reviews/r-localllama-comments-1lk5te5-nvidia-dgx-spark-whats-the-catch-47c88248.md
source_anchor: ""
source_lines: [9, 45]
sha256: cb8b1d4380aae30e4ccadd69b42245f0233b6b2b51a515ac2312b05d689636d4
---

# r-localllama-comments-1lk5te5-nvidia-dgx-spark-whats-the-catch-47c88248

    En ce moment, je m'entraîne/affine des modèles de transformateurs pour l'audio (environ 50 millions de paramètres) avec ma super 3090 et pour l'affinage, ça marche nickel, alors que l'entraînement de zéro, c'est quasi impossible, parce que c'est lent et j'ai pas assez de VRAM.
J'ai découvert le DGX Spark et je regardais celui d'Asus à 3000$, mais je vois pas le piège. Sur la plupart des sites, les gens se plaignent et disent que ça vaut pas le coup et tout, mais à part la bande passante mémoire plus lente (2-3 fois moins vite que la 3090 si les specs sont vraies) - je vois pas d'inconvénients ?
Le truc le plus impressionnant pour moi, c'est la mémoire unifiée de 128 Go, qui, je suppose, pourrait être utilisée comme VRAM et qui accélérerait pas mal mon boulot.
Y'a un truc à surveiller quand on prend le DGX Spark ?
Section des commentaires
> bande passante mémoire plus lente
C'est un gros problème, à mon avis.
Les gens comprennent généralement mal le but du DGX Spark parce qu'ils refusent obstinément d'écouter Nvidia quand ils l'expliquent : C'est un kit de développement. Il est destiné à reproduire l'architecture et la pile logicielle d'une station de travail ou d'un serveur DGX prêt pour la production afin que les développeurs puissent tester leurs trucs sur le Spark et si ça marche là-bas (peu importe la lenteur), ça marchera sur le grand frère.
Il n'a jamais été conçu comme un produit autonome pour l'inférence ou la formation au-delà du test pour savoir si ce que vous essayez de faire fonctionnera réellement.
La mémoire plus lente sera "un plus", mais ce sera quand même beaucoup plus rapide pour l'entraînement que d'échanger dans la RAM du système. Notez également que seuls 96 Go de la mémoire unifiée peuvent être alloués au GPU. Mais même avec ces réserves, je parie que beaucoup d'amateurs utilisent le Spark pour entraîner de petits modèles et effectuer des ajustements complets de modèles de taille moyenne.
Ok, mais on peut avoir des MacBooks avec 128 Go de mémoire unifiée avec une bande passante de dingue de 546 Go/s contre le dgx spark à 276 je crois.
Si tu veux CUDA, prends deux 3090 avec de la DDR5 pour soutenir les 60 Go finaux.
Ça fait un peu marrer, parce que 98% du dev, tu peux le faire sur un laptop avec une 5060 Ti. Si t'es prêt à en prendre un autre, tu pourrais probablement gérer 99,99% du dev — et ce serait toujours moins cher qu'un DGX Spark. Bien sûr, le DGX Spark a plus de VRAM, mais pour le dev, t'as généralement juste besoin d'un prototype à petite échelle, puis tu l'adaptes pour le déploiement.
"98% du dev peut se faire sur un laptop avec une 5060 Ti" tu peux élaborer un peu plus, stp ?
Ah, compris, merci ! C'est un peu dommage, je cherchais des solutions similaires de type mini PC pour m'entraîner.
Tu m'as au moins fait gagner 1 heure🙏
Le truc, c'est que c'est super lent et que ça ne peut servir qu'à tester. Ce n'est pas un poste de travail prêt pour la production, c'est juste quelque chose que les devs peuvent utiliser pour s'assurer que tout fonctionne avant de le déployer sur un cluster ou autre chose.
J'ai lu toutes les spécifications et je vois vraiment pas en quoi c'est un mauvais produit, en dehors du domaine des LLM. Comme je l'ai dit, je travaille pas avec des modèles de langage et pour les transformateurs audio comme ceux que j'utilise, ils dépendent pas des tokens ou quoi que ce soit de similaire, et la bande passante mémoire, c'est pas un truc qui doit être super rapide.
Cela dit - j'ai jamais entraîné sur un truc avec 200-300GB/s franchement, tous les GPU que j'ai utilisés ont au moins 800GB/s. Peut-être que je devrais essayer, mais je reste sceptique que ça fasse une grosse différence.
L'entraînement, ce n'est pas juste une question de TFLOPs. Il faut aussi déplacer les données assez vite pour que les unités de calcul soient occupées. Chaque étape d'entraînement doit lire et écrire les poids, les activations, les gradients et les états de l'optimiseur. Ça représente généralement trois à cinq fois la taille du modèle en trafic mémoire pour chaque étape.
Exemple : un modèle à 1 milliard de paramètres en FP16, ça fait environ 2 Go de poids. Un cycle forward + backward + update peut déplacer environ 20 Go de données.
Maintenant, regardez deux GPU avec la même puissance de calcul :
Un GPU avec une bande passante de 300 Go/s a besoin d'environ 67 ms juste pour déplacer ces données.
Un GPU avec une bande passante de 900 Go/s n'a besoin que d'environ 22 ms.
Les unités mathématiques sont les mêmes, mais la carte à faible bande passante passe beaucoup de temps à attendre la mémoire. C'est pourquoi les GPU axés sur l'entraînement comme les A100, H100 et 4090 se situent tous dans la plage des 800–1000 Go/s.
tl;dr : une fois que le modèle est assez gros, l'entraînement est généralement limité par la bande passante mémoire plutôt que par le calcul. Passer de 900 Go/s à 300 Go/s peut vous ralentir de deux à trois fois, même si les FLOPs semblent les mêmes.
Je suis un peu perdu, parce que la note TOPS maximale supposée pour l'AMD Strix Halo 395+ est de 126, alors que la documentation de Nvidia dit que le DGX Spark peut fournir 1 000 TOPS avec une précision de 4 points. Maintenant, peut-être que la clause de non-responsabilité "avec une précision de 4 points" leur permet de gonfler un peu le chiffre, mais il semble que le DGX Spark surclassera facilement le Strix Halo 395+.
Cela dit, Strix Halo gère toujours les charges de travail X86_64 et est souvent livré avec Windows, tandis que le DGX Spark exécutera une version d'Ubuntu optimisée par NVidia pour ses capacités de charge de travail IA.
Pendant ce temps, on est peut-être à 4 mois de la sortie du M5 d'Apple. Donc, au moment où le DGX Spark sera livré (ou peu de temps après), les gens auront le choix entre lui, un M5 et un 395+, tous offrant des bus de mémoire intégrés qui permettent au GPU/NPU d'accéder à une plus grande quantité de RAM.
"À la précision 4 points" doit discrètement faire référence à INT4/FP4, ce qui est complètement nul pour le moment. Un nombre de 4 bits ne peut contenir que 16 valeurs différentes. C'est inutile pour la plupart des modèles.
Quelques modèles peuvent être optimisés pour fonctionner correctement en 4 bits. Mais cela nécessite de nouvelles architectures qui utilisent des réseaux très profonds avec des milliers de couches (comme ResNet) pour obtenir des résultats utiles - car le 4 bits nécessite beaucoup de couches pour créer des résultats qui approchent ce que d'autres réseaux peuvent faire avec peu de couches et de très grands nombres à la place.
Quelqu'un devrait inventer une nouvelle architecture pour faire quelque chose d'impressionnant en 4 bits avant que je ne m'enthousiasme pour "1 trillion d'opérations en 4 bits par seconde".
Il reste à voir si le 4 bits sera un jour plus utile que les modèles 8/16/24/32 bits.
La plupart des nouveaux modèles ollama sont quantifiés à 4 bits par défaut. 4 bits, c'est super.
Évidemment, c'est pas aussi bien que 8 ou 16 en termes de précision.
On est en janvier 2026 maintenant, et on a maintenant des données qui expliquent le DGX Spark et définissent un peu mieux sa niche.
Certaines des réponses ci-dessous indiquent son utilisation comme un environnement de développement de type production, permettant à chaque développeur d'avoir un environnement de développement et de test qui est plus ou moins garanti de fonctionner comme en production. C'est vrai, mais ce n'est pas toute la vérité parce que déployer une boîte sur mon bureau est le même niveau de difficulté que déployer dans le cloud - vous changez juste l'adresse IP vers laquelle vous poussez vos changements. C'est peut-être ce que Nvidia a dit que vous "pourriez" vouloir l'utiliser, mais ce n'est pas le principal cas d'utilisation.
