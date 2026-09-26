---
id: collect-260926-rattrapage/rattrapage/r-localllama-comments-1tmhypf-best-ai-agent-for-coding-locally-b64a6dad-2
title: "r-localllama-comments-1tmhypf-best-ai-agent-for-coding-locally-b64a6dad"
domain: rattrapage
role: reference
task: reference
actors: ["AMD", "Alibaba", "Anthropic", "Google", "OpenAI"]
dates: []
keywords: ["agent", "llama", "amd", "chatgpt", "claude", "datacenter", "gemini", "gpu", "gpus", "llama.cpp", "mcp", "mixture of experts"]
source: docs/RAG/lot-rattrapage/ai-llm/r-localllama-comments-1tmhypf-best-ai-agent-for-coding-locally-b64a6dad.md
source_anchor: ""
source_lines: [14, 48]
sha256: a889f871fef2be49c70d0e31ca962c8448028dc6911ca547e1123e956926a965
---

# r-localllama-comments-1tmhypf-best-ai-agent-for-coding-locally-b64a6dad

Je veux coder un site web et une application pour un projet, et je me demandais, quelle est la meilleure IA que je peux utiliser avec mon matériel, et devrais-je utiliser un outil comme Claude Code ou Pi agent pour les faire fonctionner ?
J'ai essayé Gemma4 sur Pi Agent et c'était vraiment bizarre pour une raison quelconque, cependant je pense que Pi Agent était en partie responsable. Devrais-je réessayer en local ? Ça a aussi pris environ 6-7 minutes pour obtenir un résultat.. avec ChatGPT cela prend souvent environ 20 secondes et la qualité est souvent bien meilleure. Le temps n'est pas mon problème, mais je pensais que les IA locales sont presque aussi bonnes que celles de OpenAI et Claude de nos jours ? Bref, pour l'instant, je veux juste coder une page d'atterrissage. Devrais-je juste le faire avec Chat ou y a-t-il de bonnes alternatives pour mon matériel en ce moment ?
Merci d'avance !
Section des commentaires
Vous pouvez exécuter qwen 3,6 35b a3b. Vous pouvez mettre tous les experts sur la carte graphique. Pour un câblage local gratuit, vous pouvez utiliser opencode ou l'agent Hermes avec des compétences en codage
Qu'est-ce que tu veux dire par "l'expert de la carte graphique" ? Tu veux dire faire tourner l'IA sur la carte graphique
c'est un modèle MoE ou "mixture of experts", différent d'un modèle dense comme 27b où toute la bibliothèque et les connaissances sont chargées à l'intérieur du GPU. Avec MoE, c'est comme diviser la bibliothèque et les connaissances, et cela peut vivre à des endroits différents comme le CPU d'où le nom A3B ou seulement 3B (connaissances) actif à un moment donné. En général, c'est une chance pour les GPU plus petits de pouvoir faire fonctionner un grand modèle sans sacrifier trop de performance, mais avec ta bonne carte graphique, tu peux facilement les charger là-bas. CMIIW
Vous pourriez exécuter un modèle MOE si vous déchargez sur le CPU / RAM, le truc est d'équilibrer autant que possible sur le GPU. (voir le paramètre --n-cpu-moe)
Je suis super nouveau je ne sais pas comment faire ça que me recommanderais-tu j'utilise Linux donc je ne sais pas si ça fonctionne avec les pilotes et tout
sur une 9070 XT (16 Go) + 32 Go de RAM, tu as de vraies options, mais d'abord quelques points :
6-7 minutes par réponse, c'est très loin de la cible, quelque chose tournait sur le CPU ou le modèle ne tenait pas vraiment sur le GPU. ROCm + llama.cpp Vulkan devrait te donner 20-40 t/s sur quelque chose comme qwen2.5-coder-14b à Q4. confirme que le GPU fait bien le travail via l'équivalent AMD de nvidia-smi (radeontop ou rocm-smi).
pour harness : aider est le plus mature pour le codage de modèles locaux, installe-le avec pip et dirige-le vers un serveur llama.cpp. continue ou cline en tant qu'extensions VS Code fonctionnent aussi bien. j'éviterais pi pour l'instant, il y a une raison pour laquelle la plupart des gens utilisent les autres.
partie honnête : pour construire un site web complet + une application, un modèle local 14b va te frustrer. l'écart de qualité par rapport à chatgpt/claude est réel et important. utilise le local pour des tâches ciblées (écrire cette fonction, refactoriser ce fichier) et les modèles frontier pour la planification et l'intégration réelles. n'essaie pas de tout faire localement sur du matériel grand public pour le moment, les calculs ne fonctionnent pas encore.
6-7 minutes (parfois même beaucoup plus longtemps) pour le résultat final, bien que ce ne soit pas juste pour la réponse qui était instantanée. J'utilise Arch donc je ne sais pas si les pilotes fonctionnaient correctement.
Qu'est-ce que tu recommandes pour Linux ? Extension VsCode ?
Pour la page d'atterrissage, ça fonctionne non ? Ou devrais-je même utiliser des modèles publics là-bas ?
ah ok c'est donc un cadre totalement différent, ce n'était pas clair dans le premier post. 6-7 minutes pour une tâche d'agent complet sur un local 14b est en fait assez normal, pas une mauvaise configuration. le codage en plusieurs étapes (lire des fichiers, planifier, éditer, exécuter, corriger) prend plusieurs tours et chaque tour dure 20-40 secondes sur du matériel grand public contre 2-3 secondes sur les GPUs du datacenter de chatgpt, ça s'accumule rapidement.
donc le goulot d'étranglement n'est pas ta configuration, c'est juste que le local 14b sur du matériel grand public est vraiment 10-15 fois plus lent de bout en bout que les modèles hébergés par frontier. l'écart de vitesse est réel et pas vraiment réparable sans du matériel plus puissant. le local a du sens pour des raisons hors ligne / de confidentialité / de coût, pas pour atteindre la vitesse de frontier.
J'ai configuré Pi avec une extension vscode qui me permet de contourner l'utilisation de Pi CLI et de l'utiliser dans Vscode à la place, ça fonctionne vraiment bien. Pour mon dernier test, j'ai demandé à Gemma-4-26B de créer un fichier HTML simple d'une lampe à lave dont les bulles réagissent à mes mouvements de souris. Ça a pris moins d'une minute. J'ai eu l'idée du "test" d'une chaîne YouTube qui teste des LLM - à vérifier : https://www.youtube.com/watch?v=AAsW5oHCgic
Quel tok/sec obtiens-tu juste en discutant avec le LLM que tu as configuré ? Qu'utilises-tu pour ton inférence LLM ? Ollama ? llama-cpp ?
6-7 minutes pour obtenir un résultat semble vraiment mauvais pour ton matériel, je soupçonne qu'il y a un problème avec la configuration ou le setup. Nous pouvons t'aider mais nous avons besoin de plus d'infos.
Eh bien, j'obtiens une réponse tout de suite - mais ça prend ce temps-là pour finir de coder
qwen3.6 27b mtp + un pi bien réglé est super. j'ai abandonné tous les abonnements cloud, fonctionnant à 100 % local depuis la sortie de qwen3.6.
Puis-je aussi simplement utiliser qwen comme un llm et créer les fichiers moi-même
Salut, je te recommande d'essayer Qwen3.6 35b ou 27b, ça fonctionne bien avec mon outil d'agent de codage, facile à installer, beaucoup de fonctionnalités, faible empreinte contextuelle.
https://github.com/leflakk/openclose
Devrais-je l'acheter via llama.cpp ou ollama ?
L'application est openai compatible donc les deux fonctionnent mais llama.cpp est clairement le meilleur, il est bien documenté donc n'importe qui peut apprendre à l'utiliser.
Violation de la règle 1 - Blocage (au lieu de retirer, comme l'ont partagé certaines personnes qui ont fait de bons retours ici).
Si vous voulez quelque chose de simple comme une page d'atterrissage - utilisez simplement le Google Gemini gratuit. Si vous voulez quelque chose de compliqué - payez pour Claude ou ChatGPT.
Si vous voulez en savoir plus sur les LLM et devenir meilleur dans leur utilisation - oui, utilisez des modèles locaux. Comprenez simplement que les résultats ne seront pas aussi bons dès que vous vous éloignez des tâches simples habituelles.
J'ai constaté que Gemma est très exigeante en ce qui concerne le harnais, j'ai eu le plus de succès avec opencode. Je n'ai également pas réussi à obtenir des résultats constamment bons une fois que le contexte dépassait 100K. Qwen, en revanche, semble bien fonctionner avec n'importe quel harnais.
J'ai pu obtenir les meilleurs résultats en alternant entre les modèles pendant le projet. Mais encore une fois - dès que vous essayez de faire quelque chose de moins courant, vous rencontrerez des défis. Mon exemple le plus récent - ni Gemma ni Qwen n'étaient capables de créer un tableau de bord fonctionnel dans Datadog en utilisant le serveur MCP de Datadog. Étant donné le même fichier de spécifications exact, Gemma n'a complètement rien réussi à créer, Qwen a créé un tableau de bord avec un graph travaillant sur 30 et Claude Sonnet a créé un tableau de bord totalement fonctionnel.
Commentaire supprimé par un membre de l’équipe de modération
Qwen2.5 !!
