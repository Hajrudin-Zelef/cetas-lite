---
id: collect-260926-rattrapage/rattrapage/r-localllama-comments-1n99vhp-how-do-i-run-ai-locally-and-what-is-the-most-a1a767b1-3
title: "r-localllama-comments-1n99vhp-how-do-i-run-ai-locally-and-what-is-the-most-a1a767b1"
domain: rattrapage
role: reference
task: reference
actors: ["OpenAI"]
dates: []
keywords: ["llama", "gpu", "llama.cpp", "moe", "vllm"]
source: docs/RAG/lot-rattrapage/ai-llm/r-localllama-comments-1n99vhp-how-do-i-run-ai-locally-and-what-is-the-most-a1a767b1.md
source_anchor: ""
source_lines: [47, 67]
sha256: ce55d61dfd83c30658a057885dbf15a7a29fd7e1d53c87eb61371bf1a3151752
---

# r-localllama-comments-1n99vhp-how-do-i-run-ai-locally-and-what-is-the-most-a1a767b1

Ollama - est PRESQUE identique à llama.cpp, mais a été dérivé dans son propre petit jardin clos, avec ses propres modèles "requantisés". Ollama était le premier outil à télécharger des modèles pour vous, mais llama.cpp peut déjà le faire tout seul.
LM studio - Je n'en ai aucune idée, car je ne l'ai jamais utilisé. J'ai entendu dire qu'il ne peut pas utiliser votre RAM ou utiliser votre CPU ? Je pourrais le confondre avec VLLM, cependant.
VLLM - Pas d'idée non plus.
L’interface web de génération de texte d’Oobabooga - Était le premier outil "tout inclus". Comparé aux outils ci-dessus, celui-ci peut exécuter une multitude de types de quants différents, comme ExLlamav2, ExLlamav3, Llama.cpp, Transformers, etc. Mais il reste toujours un peu à la traîne niveau mise à jour.
Ensuite, il y a plein d'autres applications, ou façons d'inférer vos modèles, mais je considérerais que ces applications sont divisées en 2 équipes différentes :
Équipe 1, ce sont des applications de fichiers exe C++. Utilisées principalement pour exécuter vos trucs sur votre PC avec peu ou pas de prérequis (excepté peut-être des pilotes GPU et des trucs accélérés par GPU comme CUDA).
Équipe 2, ce sont des applications de "l'enfer Python". Vous avez besoin d'au moins les bases de Python, et des outils Python essentiels pour les faire fonctionner. En raison de la nature de Python en tant que langage de programmation, vous ne passerez peut-être pas un très bon moment à essayer différentes applications Python, car elles peuvent planter à cause des environnements ou des conflits de versions de dépendances.
Quoi que vous choisissiez, gardez à l'esprit que si vous rencontrez un problème étrange, n'hésitez pas à consulter l'onglet "Issues" sur la page Github du projet correspondant. Parfois, des choses "anciennes, éprouvées et vraies" peuvent se retrouver cassées, et vous devrez peut-être essayer autre chose, ou attendre un certain temps avant que les gens compétents aient une chance de le réparer. C'est la nature de cette "toute nouvelle chose".
Oh, et il y a une autre façon de les diviser :
CPU+GPU.
GPU uniquement.
Plusieurs GPU uniquement (outils professionnels pour servir un grand nombre de clients).
La plupart d'entre nous n'utilisent que 1. Mais si vous êtes un passionné avancé avec un extra à dépenser, vous pourriez trouver que 2 et 3 sont plus à votre goût.
Super aperçu. Je peux partager quelques backends supplémentaires à considérer selon le cas d'utilisation :
Pour les grands modèles MoE, ik_llama.cpp est l'un des meilleurs backends. J'ai partagé des détailsicisur la façon de le construire et de l'utiliser au cas où quelqu'un voudrait essayer s'il ne l'a pas déjà fait. Ça peut être beaucoup plus rapide comparé à llama.cpp, surtout pour des tailles de contexte plus grandes. Il est notable qu'ik_llama.cpp vient avec un frontend Mikupad intégré. Mais bien sûr, tout autre frontend qui prend en charge l'API compatible avec OpenAIpeut y être connecté, comme SillyTavern ou Open WebUI.
TabbyAPI (exllamav2/v3) est une excellente option lorsque le modèle tient entièrement dans la VRAM, car cela a tendance à fournir de meilleures performances comparé à d'autres backends.
VLLM, déjà mentionné dans votre aperçu, est l'une des meilleures options pour l'inférence par lots et les configurations multi-utilisateurs sur des machines avec beaucoup de VRAM, car il est moins efficace en mémoire d'après mon expérience, et aussi plus compliqué à configurer.
Détermine quel niveau de modèle tu veux exécuter. MoE fonctionne très bien sur la plupart des configurations RAM+CPU. Dis-nous ta configuration pour plus d'infos.
Je te recommanderais KoboldCPP. Tu n'as qu'à exécuter le fichier exe et charger le modèle, et elle a plein de fonctionnalités qui manquent à d'autres plateformes comme LMStudio, comme par exemple le "mode histoire/autocomplétion pure" et les cartes de personnages.
Je trouve aussi que l'API est beaucoup plus compatible.
Que veux-tu en faire ? Quel matériel as-tu ?
