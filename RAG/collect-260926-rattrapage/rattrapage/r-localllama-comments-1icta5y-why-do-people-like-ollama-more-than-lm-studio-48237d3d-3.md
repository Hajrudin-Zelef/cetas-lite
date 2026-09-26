---
id: collect-260926-rattrapage/rattrapage/r-localllama-comments-1icta5y-why-do-people-like-ollama-more-than-lm-studio-48237d3d-3
title: "r-localllama-comments-1icta5y-why-do-people-like-ollama-more-than-lm-studio-48237d3d"
domain: rattrapage
role: reference
task: reference
actors: ["Hugging Face", "OpenAI"]
dates: []
keywords: ["llama", "arr", "gguf", "llama.cpp", "open source"]
source: docs/RAG/lot-rattrapage/ai-llm/r-localllama-comments-1icta5y-why-do-people-like-ollama-more-than-lm-studio-48237d3d.md
source_anchor: ""
source_lines: [39, 55]
sha256: d380bbb1096bc4268d5543d4a32dc40b9ec8c9d74257ac4a73058b6fc5210b1b
---

# r-localllama-comments-1icta5y-why-do-people-like-ollama-more-than-lm-studio-48237d3d

LM Studio apporte le package complet avec sa propre interface. Sur macOS, c'est aussi l'une des options les plus faciles pour exécuter des modèles MLX. Contrairement à ce que les gens semblent penser, il peut également fonctionner en arrière-plan et servir de backend pour d'autres applis (API OpenAI + leur propre API désormais). En ce qui concerne ses défauts, c'est "une autre app Electron", c'est plus lourd (tu n'as peut-être pas besoin de tout), et c'est évidemment fermé (en partie, puisque ça repose de toute façon sur des technologies open-source).
Il y a bien sûr d'autres options, mais c'est mon expérience avec ces deux-là. Je les utilise toutes les deux, ainsi que PrivateLLM (pour ses Omniquants), ce qui me permet de faire facilement fonctionner et comparer des modèles sur différents moteurs et techniques de quantification (llama.cpp, mlx, mlc) afin d'utiliser ce qui est le mieux pour moi.
Je pense que beaucoup d'entre nous sont des programmeurs, personnellement je préfère les applications open-source. Bien sûr, en tant que programmeur, je n'ai pas de problèmes avec l'utilisation de la ligne de commande.
Ollama n'utilise-t-il pas aussi des GGUF ?
Il en va de même pour Ollama. Vous pouvez chercher des modèles en utilisant le site web d'Ollama. Télécharger des modèles, c'est juste une commande
ollama pull <model_name>. Il en va de même pour enlever ou servir un modèle.
Et tu peux utiliser Ollama avec HuggingFace GGUF maintenant aussi.
Je ne veux pas d'une interface graphique. Je veux un moteur qui peut tourner en arrière-plan sans avoir besoin d'être surveillé par moi et que je peux pointer vers Open WebUI. C'est plus facile avec Ollama.
Si je devais deviner, c'est parce qu'Ollama est open source alors que LM Studio ne l'est pas. Personnellement, j'ai toujours utilisé LM Studio car il a une interface très propre et les choses sont juste faciles à faire/à gérer. Ollama propose également une CLI, ce qui pourrait être une autre raison
Préoccupations concernant la vie privée et les pièces jointes.
Le fait que LM Studio soit logiciel propriétaire ouvre un vecteur de fuite de code et ils peuvent l'arrêter à tout moment, vous devrez de toute façon migrer. Ils peuvent aussi commencer à facturer l'application à tout moment et/ou ajouter des publicités. Pour moi, la confidentialité du code est la plus importante. Je ne peux pas, en bonne conscience, utiliser le code d'une entreprise dans quelque chose d'aussi fermé que LM Studio.
Je préfère Ollama pour sa simplicité, mais le moteur MLX de LM Studio est incroyable. Les modèles fonctionnent plus vite et avec moins de mémoire. J'aimerais qu'Ollama l'adopte.
J'utilise LM Studio même en production. Tu peux le démarrer en tant que service. Quand le produit est nouveau, tu dois le surveiller, tester différents paramètres et modèles, c'est assez bien dans LM Studio. Je l'ai utilisé depuis 0.1 jusqu'à la nouvelle 0.3, la nouvelle interface et le support multi-modèles sont plutôt bons.
Ollama est mignon
Comment ça ?
Je suis encore plus confus sur la raison pour laquelle quelqu'un utiliserait ollama ou lmstudio au lieu du serveur llama.cpp original. Un énorme avantage de llama.cpp, en plus d'être le plus à jour en ce qui concerne le travail avec les fichiers gguf, est que vous n'avez pas besoin de les fusionner d'abord, et vous êtes libre de choisir n'importe quel modèle, pas seulement ceux fournis par les interfaces ollama ou lmstudio.
Et le serveur llama a sa propre interface graphique pour ceux qui aiment discuter avec leurs modèles.
