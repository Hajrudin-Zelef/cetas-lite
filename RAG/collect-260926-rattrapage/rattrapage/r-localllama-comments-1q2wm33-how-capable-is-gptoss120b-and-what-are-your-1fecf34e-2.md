---
id: collect-260926-rattrapage/rattrapage/r-localllama-comments-1q2wm33-how-capable-is-gptoss120b-and-what-are-your-1fecf34e-2
title: "r-localllama-comments-1q2wm33-how-capable-is-gptoss120b-and-what-are-your-1fecf34e"
domain: rattrapage
role: reference
task: reference
actors: ["Anthropic", "Apple", "DeepSeek", "MiniMax", "Mistral", "Moonshot", "OpenAI", "OpenRouter", "Z.ai"]
dates: []
keywords: ["llama", "claude", "deepseek", "glm", "gpu", "kimi", "llama.cpp", "moe", "mxfp4", "open source", "reasoning"]
source: docs/RAG/lot-rattrapage/ai-llm/r-localllama-comments-1q2wm33-how-capable-is-gptoss120b-and-what-are-your-1fecf34e.md
source_anchor: ""
source_lines: [9, 49]
sha256: 96a47bb0ed8e9913764952f9aba95e890e4a08d62674c58c4c1dda121f12d179
---

# r-localllama-comments-1q2wm33-how-capable-is-gptoss120b-and-what-are-your-1fecf34e

    J'ai une RTX 3090 et j'envisage d'en prendre une autre pour pouvoir faire tourner OSS-120b. Ça m'intéresse surtout de discuter avec sur des documents privés, des analyses statistiques, des connaissances/analyses STEM et un peu de code.
Est-ce que ça vaut le coup d'investir ? Je suis pas contre la spéculation dans ce post - vous pensez quoi de possible pour des modèles plus petits dans ce délai que je pourrais faire tourner avec deux RTX 3090 cette année ?
Section des commentaires
Comment tu voudrais faire tourner oss-120b d'une manière que tu peux pas faire en ce moment ? 2x3090 ça te donne que 48GB de vram et oss-120b il faut au moins 60, je crois.
Tu peux le faire tourner avec llama.cpp maintenant (si t'as la RAM) et ça pourrait aller plus vite avec llama.cpp avec deux 3090, mais tu peux tester ses capacités tout de suite.
Ok, merci. J'ai 32 Go de DDR4, est-ce que je devrais upgrader ?
Aussi, comment est la sécurité avec OpenRouter ? Ils disent qu'ils ne collectent pas nos chats.
If you only have 32GB system RAM you will not be able to run gpt-oss 120B even with 2x 3090 because the model itself is already about 65GB. The minimum you will need is 64GB system RAM. With a 24GB 3090/4090 and 64GB DDR5-4800 you can load it with full 128K context if you use -ncmoe 28 with llamacpp. You have DDR4 so it will be slower, but probably still around 15 tok/s (someone surely has some numbers, I only have systems with 32GB of DDR4-2666 so can't test for exact values
Openrouter, c'est un intermédiaire. Ils acheminent tes requêtes vers différents (des dizaines) fournisseurs d'hébergement de modèles. Tu dois tenir compte à la fois des politiques d'Openrouter et de celles de ces autres entreprises, puisque tu interagis avec les deux quand tu utilises openrouter.
Ces modèles sont super performants, ces derniers jours j'ai joué avec mon nouveau Mac Studio avec 128 Go, je dois avouer qu'au début j'étais super déçu des performances, j'essayais de faire tourner des modèles avec LMStudio + Claude Code Router.
LMStudio a un gros défaut, c'est que le chargement du contexte prend une éternité, Claude Code Router était aussi super inefficace pour une raison ou une autre et les modèles produisaient n'importe quoi.
Ensuite, j'ai décidé d'essayer Llama.cpp et Jan.ai comme outil pratique pour récupérer le modèle et les charger. Beaucoup mieux que LMStudio (l'interface est moins sympa mais pas mal). Et en combinant Jan.ai avec Opencode : un client open source assez impressionnant inspiré de Claude Code, beaucoup mieux en fait, il est plus joli, comprend les Skills, les slash commands, les Subagents, etc.
J'ai essayé avec gpt-oss-20b et j'ai pu Vibe Code avec un niveau de résultats très comparable à ceux que j'obtiens avec Claude Code + Sonnet (probablement quelque chose de similaire à Sonnet 3.5).
Ensuite, j'ai pu faire tenir Minimax avec un niveau de quantification q3. Et les résultats étaient encore meilleurs et la vitesse toujours assez acceptable.
Ma conclusion est qu'il est tout à fait possible d'utiliser ça pour des projets personnels/familiaux et d'éviter de devoir payer un abonnement.
Alors, tu aimes bien Minimax ? On dirait que c'est un super modèle.
Je suis un grand fan de gpt-oss-120b. C'est un modèle très sympa et polyvalent qui tourne à plus de 100t/s sur trois 3090 ou 45t/s sur trois Mi50 de 32 Go. C'est mon modèle de prédilection pour la plupart des choses, y compris la plupart des tâches d'administration système.
Pour 2026, je pense que ça continuera la tendance des années précédentes : des modèles plus petits sortiront, aussi performants que des modèles 5 à 10 fois plus grands. Le matériel continuera de coûter plus cher.
Ça marche super bien sur les systèmes limités par la bande passante mémoire comme Strix Halo.
Il a quelques-uns des comportements clairsemés habituels (il devient paresseux, ok pour la connaissance du monde, médiocre en codage). Ce qui est génial avec le modèle, c'est qu'il hallucine rarement par rapport aux autres modèles 4 bits (surtout dans cette gamme de tailles).
Avant, je l'utilisais pour l'architecture logicielle (maintenant je reste avec Minimax M2.1 IQ3_M... il hallucine parfois, mais c'est étonnamment rare, au point que j'oublie qu'il est quantifié)
Commentaire supprimé par un membre de l’équipe de modération
Si t'as plein de RAM normale, genre 96 Go, tu peux faire tourner des modèles moe de 100B en utilisant le déchargement tensoriel sur une seule 3090.
Mais, genre, combien de t/s ? Si c'est 20, ce qui est en gros la vitesse de lecture, c'est top.
9-10 à Q5KM (glm 4.5 air). Ça va jamais être proche d'un modèle GPU complètement chargé.
Je recommande vivement 3x RTX3090 pour faire tourner GPT-OSS-120b. Avec trois, ça tournera entièrement depuis le GPU avec un contexte complet de 128k et environ 110t/s sur llama.cpp. Pour moi, c'est le truc que j'utilise tous les jours. Connaissances, cuisine, contrôle Home Assistant, un peu de programmation, de l'aide avec linux / docker… J'espère vraiment que OpenAI travaille sur une mise à jour qui inclura aussi la vision.
Je suis tout à fait d'accord. Difficile de faire mieux que la combinaison de vitesse et de performance qu'on obtient avec oss-120b. Je suis en train de réfléchir à l'idée d'acheter une 4ème 3090 pour pouvoir faire tourner oss-120b et z-image en même temps…
Je sais que c'est populaire de cracher sur gpt-oss ici, mais ça tape vraiment dans le mille pour une utilisation générale.
C'est super rapide sur Apple Silicon et Strix Halo. (60-70 t/s pour gpt-oss-120b-mxfp4 sur M1 Ultra, comparé à ~20 t/s pour MiniMax M2.1 UD_Q2_K_XL)
Le cache KV est très efficace :
Metal KV buffer size = 4608.00 MiBpour l'intégralité du contexte de 128K. Compare ça à MiniMax M2 qui a besoin d'environ 30 Go pour un contexte de 128K.
L'ensemble du modèle + cache KV n'utilise que ~65 GiB de mémoire, donc il vous reste encore beaucoup de place pour d'autres tâches sur des machines de 128 Go.
Effort de raisonnement réglable pour que vous puissiez passer par défaut à élevé, mais juste passer
lowàreasoning_effortdanschat_template_kwargssi vous voulez juste une réponse rapide.
C'est assez intelligent pour sa catégorie de taille. Bien sûr, ça ne va pas rivaliser avec GLM, Kimi K2, DeepSeek, etc. de taille normale, mais c'est quelque chose qui peut être exécuté sur la machine de la plupart des gens.
Si vous avez des problèmes avec le garde-fou par défaut, vous pouvez simplement exécuter la version hérétique. Pour la plupart des tâches de codage/agénétiques, la version de base devrait bien fonctionner.
C'est dingue comme les gens ont la mémoire courte ; gpt-oss-120b est à peu près aussi performant que openai chat-gpt 4, et c'est le 3.5 qui a rendu le monde fou. C'était genre, il y a 18 mois. C'est un super modèle. Pouvoir faire tourner ça sur du matos local, c'est révolutionnaire.
Tl;dr : Ajouter de la RAM est indispensable, et si le budget n'est pas un problème, ajoutez à la fois de la RAM et une autre 3090.
Ajouter une 3090 aidera à faire tourner d'autres modèles plus petits. Mais vous ne remarquerez aucune différence si ce que vous voulez faire tourner est gpt-oss-120b.
Vous ne pourrez pas faire tenir gpt-oss-120b sur la VRAM avec 2x3090. Vous devrez toujours décharger une partie du modèle sur la RAM, et quand la RAM est impliquée, vous êtes limité par la vitesse de la RAM et du CPU, peu importe le nombre de 3090 que vous avez.
Si votre besoin est de faire tourner gpt-oss-120b, ajouter 64 Go de DDR4 supplémentaires est votre meilleure option. Mélanger de la RAM DDR4, c'est ok - je l'ai fait.
