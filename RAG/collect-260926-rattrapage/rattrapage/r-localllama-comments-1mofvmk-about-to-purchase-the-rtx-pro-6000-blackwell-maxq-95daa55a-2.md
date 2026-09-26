---
id: collect-260926-rattrapage/rattrapage/r-localllama-comments-1mofvmk-about-to-purchase-the-rtx-pro-6000-blackwell-maxq-95daa55a-2
title: "r-localllama-comments-1mofvmk-about-to-purchase-the-rtx-pro-6000-blackwell-maxq-95daa55a"
domain: rattrapage
role: reference
task: reference
actors: ["Alibaba", "Meta", "Nvidia"]
dates: []
keywords: ["blackwell", "llama", "benchmarks", "gpu", "llama.cpp", "nvidia"]
source: docs/RAG/lot-rattrapage/ai-llm/r-localllama-comments-1mofvmk-about-to-purchase-the-rtx-pro-6000-blackwell-maxq-95daa55a.md
source_anchor: ""
source_lines: [16, 53]
sha256: 858d6345a84095718c6f3fb85d1fd9cd8f13e87ec91f4b6d3829292e8d62b84e
---

# r-localllama-comments-1mofvmk-about-to-purchase-the-rtx-pro-6000-blackwell-maxq-95daa55a

Section des commentaires
J'ai une RTX 8000 et je viens de commander une Pro 6000 Max Q. Elle arrive demain d'après le suivi FedEx :)
Je vais d'abord l'installer sur Linux, puis sur Windows plus tard. Je vais surtout l'utiliser pour construire des modèles d'IA à partir de zéro. Donc, je n'aurai peut-être pas de vitesses d'inférence tout de suite.
https://m.youtube.com/watch?v=-RaKDJuEzUE#bottom-sheet Cette vidéo vient de la page YouTube de Micro Center et c'est la seule que j'ai vue utilisant un GPU MAX q. Toutes les autres vidéos YouTube étaient la version studio avec une limite de puissance fixée à 300.
Ils ont seulement testé gpt oss 120b mais avec une vitesse de token très rapide.
Si vous êtes nerveux, mais que vous avez l'argent, achetez-le et si vous ne pensez pas que ça vaut le coup, renvoyez-le, assurez-vous que l'endroit où vous achetez a une politique de retour de 30 jours.
S'il te plaît, dis-moi comment ça marche et si tu as des problèmes de drivers. J'ai besoin que les deux cartes fonctionnent.
Ok, je m'en occupe !
T'as vraiment besoin d'une meilleure latence ? 70t/s, c'est déjà pas mal. Tu vas probablement taper dans les 100-110t/s avec Max Q. Y'a moyen que MIG te file du parallélisme tensor/expert si tu le sépares en 2x48GB. En théorie, ça pourrait te faire grimper dans les 200t/s (?). Pareil, le pipeline devrait booster le débit. Mais bon, je suis pas sûr que TP/EP soient supportés par llama.cpp.
Je peux tester ce setup MIG demain. Mon chiffre de 110t/s, c'est de mémoire (je crois que j'avais fait tourner Q6 de la variante instruct à cette vitesse).
Ben, c'est pas juste une meilleure latence, mais aussi une architecture à jour, deux fois plus de VRAM et, espérons-le, une meilleure longévité. C'est le moment pour moi de faire une mise à niveau, de toute façon.
J'ai le 600w pro. Je pense que je pourrais le brider, non ? J'ai LM Studio ouvert en ce moment.
Ouais, tu devrais pouvoir le sous-volter à 300W.
Sans rapport avec tout ça, je suis en train de traiter un payload de 260 515 tokens sur qwen3-30b-a3b-2507 Q8_0, donc ça va prendre encore quelques minutes avant que je puisse lancer le test qui t'intéresse. Avec des prompts courts, c'est genre ~150t/s, si je me souviens bien.
Attends-toi à devoir pas mal bricoler pour que ça marche, plutôt qu'une expérience clé en main.
Peut-être que le package que tu utilises ne tourne pas sur la version de cuda dont tu as besoin pour Blackwell et tu dois configurer des trucs à la main. Peut-être que tu dois baisser la vitesse du pcie à des générations plus anciennes pour que ton PC reste stable. Peut-être que ça marche avec plus de 4g sur ta carte mère, peut-être pas. Ce genre de trucs.
Alors, la performance, c'est un peu la loterie. Mais d'après la config requise, j'ai tout le matos qu'il faut, donc je suis pas trop inquiet de ce côté-là. J'ai CUDA 12.4, donc peut-être un souci de compatibilité, quoi.
Vu que j'ai déjà une RTX 8000 Quadro, est-ce que le driver des maxq serait compatible avec l'autre carte ? C'est une architecture bien plus ancienne, mais ça reste dans la famille RTX.
Ce mec a fait des tests comparatifs sur plusieurs modèles avec une version non-max q (mais il a limité la consommation électrique pour certains, ce qui est peut-être similaire ?). Ça peut varier selon d'autres variables du système, mais ça donne une idée : https://www.reddit.com/r/LocalLLaMA/comments/1kvf8d2/nvidia_rtx_pro_6000_workstation_96gb_benchmarks/
Y a-t-il une grosse différence entre faire tourner Qwen3-30b-a3b-q8 et Qwen3-30b-a3b-q4 ? J'ai entendu dire qu'on peut faire tourner le q4 sur une RTX 3090, c'est pour un pote.
Je peux pas me prononcer sur la qualité vu que j'ai jamais utilisé le Q4, mais le Q8 est rapide sur ma carte graphique RTX 8000 Quadro avec 48 Go de VRAM, genre 70t/s, et cette carte a seulement 600 Go/s de bande passante mémoire.
Si tu n'en achètes qu'un, prends la version pro. Y'a pas de raison de jeter 10-15% de perf vu le prix de ces cartes par rapport à une mise à niveau de l'alim.
T'as pas dit quelle est l'ancienne carte, donc je vois pas comment on peut te dire à quel point ça va être plus rapide, mais un truc comme ça, ça tourne super vite sur à peu près n'importe quoi, donc ça va aussi tourner vite ici.
Windows, c'est toujours un peu la loterie, perso je le ferais pas, mais avec une seule carte graphique, t'as plus de chances que ça se passe bien. Utilise WSL2 pour l'expérience la plus proche de Linux et croise les doigts.
Tu vas pas atteindre le TDP max juste en faisant tourner un modèle pour de l'inférence en flux unique. Faut soit faire du batching (20+ de large) soit de l'entraînement pour ça. Donc je m'inquiéterais pas trop pour ça.
Le problème avec le Blackwell original, c'est la génération de chaleur. Je veux le MaxQ à cause de la consommation électrique plus faible, mais aussi une carte plus froide avec un ventilateur soufflant. Mon GPU actuel, un RTX 8000 Quadro 48 Go, a également un ventilateur soufflant, ce qui serait donc un bon ajustement pour le MaxQ afin que je puisse faire tourner les deux sur ma carte mère x670 Taichi.
Prends la version pro, et fixe une limite de puissance si ça t'inquiète vraiment. Tu peux peut-être ajouter un ou deux ventilateurs de boîtier. Je doute que ce soit un problème. Tu parles d'un GPU qui coûte plusieurs fois le prix du système actuel dans lequel tu vas le mettre. Ne laisse pas la queue remuer le chien.
16 fois le détail
Genre 10-15% contre la version normale, du moins d'après ce que j'ai entendu, ce qui le mettrait au même niveau qu'une 5090 normale.
À Q8, Qwen3-30b-a3b ne rentre pas tout à fait dans une seule 5090, donc je dois le diviser entre mes deux cartes et ça me donne, à peu près, environ 110 t/s. Pourrait probablement faire de l'optimisation et l'obtenir plus haut, mais il n'y a que quelques heures dans une journée et c'est bien comme ça.
Pour en revenir au Max-Q, si vous n'allez faire tourner qu'une seule carte et que vous avez un boîtier avec une bonne circulation d'air, prenez la PRO 6000 normale. Si votre boîtier est à l'étroit ou si vous avez l'intention de faire tourner plusieurs cartes, optez pour le Max-Q.
une partie de la règle : please linux
Commentaire supprimé par un membre de l’équipe de modération
Ouais, mais on parle d'une RTX 8000 Quadro contre la MaxQ. C'est une différence de 3 générations. L'écart devrait être assez important dans ce cas, vu que ma carte a une bande passante mémoire de 600 Go/s.
Commentaire supprimé par un membre de l’équipe de modération
Ils sont proches, mais le 6000 a un débit de jetons plus élevé. Quelqu'un les benchait l'autre jour sur localllama.
Question au hasard : pour les amateurs qui se soucient moins de la consommation d'énergie (en supposant que ce type en fait partie), pourquoi n'achète-t-il pas genre 4 5090 et ne profite-t-il pas de 3,5 fois plus de cœurs CUDA avec la même VRAM ?
Commentaire supprimé par un membre de l’équipe de modération
