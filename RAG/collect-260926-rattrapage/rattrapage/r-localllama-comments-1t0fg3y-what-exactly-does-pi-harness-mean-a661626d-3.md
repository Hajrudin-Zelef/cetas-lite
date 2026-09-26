---
id: collect-260926-rattrapage/rattrapage/r-localllama-comments-1t0fg3y-what-exactly-does-pi-harness-mean-a661626d-3
title: "r-localllama-comments-1t0fg3y-what-exactly-does-pi-harness-mean-a661626d"
domain: rattrapage
role: reference
task: reference
actors: ["Anthropic", "OpenAI"]
dates: []
keywords: ["llama", "agent", "agents", "arr", "chatgpt", "claude", "fp8", "gguf", "llama.cpp", "valuation", "vllm"]
source: docs/RAG/lot-rattrapage/ai-llm/r-localllama-comments-1t0fg3y-what-exactly-does-pi-harness-mean-a661626d.md
source_anchor: ""
source_lines: [41, 70]
sha256: 0de384a4fd6f37b565025458ff5776e0d3176a938fa61249598a80e97145eb76
---

# r-localllama-comments-1t0fg3y-what-exactly-does-pi-harness-mean-a661626d

La raison principale est que qwen3.6-27b a du mal à utiliser l'outil d'édition. Beaucoup -- quelque chose que je n'ai pas vu se produire sur d'autres harnais. Ça devient mauvais au point que le modèle peut même décider soudainement que l'outil d'édition est inutilisable et commence à écrire des scripts bash et des programmes Python pour effectuer les modifications à la place, et semble réussir à le faire de cette façon. Ce ne devrait pas être un problème de quantification, car le cache KV est soit bf16, soit fp16, et le modèle a été soit le fp8 officiel, soit au minimum unsloth q6_k gguf, qui devraient tous deux être corrects en termes de précision générale.
En commentaire, il est étrange pour moi que le remplacement de texte soit littéralement une opération de recherche-remplacement. Je pensais toujours qu'il devait fonctionner sur la base de plages de lignes, par exemple le modèle demande aux outils d'édition de supprimer les lignes 50-55 et fournit le texte de remplacement, mais en réalité, les opérations d'édition sont basées sur la fourniture de la copie exacte du vieux texte, jusqu'au dernier détail d'onglet/espace blanc, et cela doit correspondre juste une fois dans le fichier pour être acceptable. Je vois les modèles avoir des difficultés avec les espaces blancs en particulier, et écrire des scripts sed tout le temps pour qu'ils puissent voir l'arrangement exact d'onglet/espace pour le texte à substituer. Je ne sais pas pourquoi c'est nécessaire en premier lieu, car je suppose que le modèle aurait déjà dû voir l'espace blanc exact de ses lectures de fichiers. (Il se peut qu'il y ait une sorte de biais Python ici à l'œuvre, car là-bas, les espaces blancs sont plus réguliers et contrôlés dans cette langue, tandis que j'ai des arrangements d'onglets et d'espaces mélangés à cause de plusieurs personnes travaillant sur un langage non-Python.)
L'autre chose que je n'aime pas à propos des appels d'outils dans l'espace vllm, c'est qu'il n'y a pas d'application de la syntaxe des appels d'outils basée sur la grammaire. Pour autant que je sache, dans llama.cpp, les appels d'outils sont une génération contrainte par la grammaire : une fois que le modèle écrit les tokens qui commencent un appel d'outil, cela impose une génération contrainte par schéma de la part du modèle jusqu'à la fin de l'appel d'outil, mais dans le cas de vllm, il n'y a qu'un parseur général post-achèvement, et ce genre de chose dépend à 100 % du modèle pour écrire correctement l'appel. Pour une raison étrange, avec Pi, le qwen3.6-27b fait beaucoup d'erreurs, fournissant généralement le chemin incorrectement, par exemple deux fois dans l'appel d'outil, ce qui entraîne immédiatement un rejet bien que le chemin redondant soit, en principe, inoffensif. Je n'ai pas lu ce qu'est la description de l'outil d'édition pour le modèle, mais je parie qu'elle est d'une manière ou d'une autre peu claire, car quelle que soit la raison, le modèle a beaucoup de mal dans les modifications de fichiers bien qu'il sache exactement ce qu'il doit faire, et il n'est certainement pas si mauvais qu'il ait des problèmes à écrire quelques paramètres au format JSON ou XML.
Remplacez l'outil d'édition par des lignes de hachage : https://github.com/RimuruW/pi-hashline-edit
Ou utilisez Oh-my-pi
Résumé de l'auteur : https://blog.can.ac/2026/02/12/the-harness-problem/
Pareil. La dernière fois que j'ai utilisé opencode, c'était pour améliorer une extension pi afin de détecter automatiquement les modèles de serveur llama fonctionnant sur localhost:8080 et je ne suis pas revenu en arrière !
Pi est bien plus léger, donc j'apprécie cette première partie de 10k la plus rapide de la fenêtre de contexte maintenant. En plus, ce n'est pas une interface utilisateur textuelle, donc mon copier/coller entre les terminaux fonctionne simplement. J'aime ça.
Un harnais est ce qui a permis aux textes générés d'avoir un effet sur le monde qui les entoure en extrayant/injectant/en faisant des choses avec ce texte généré en temps réel
Harness est la collection de logiciels déterministes qui enveloppent un fournisseur LLM, pour transformer LLM d'un moteur d'autocomplétion en un "agent". Au minimum, ce harness peut exécuter LLM en boucle et déclencher tous ses appels d'outils, jusqu'à ce que LLM décide d'arrêter d'appeler des outils. À ce stade, la réponse finale serait renvoyée à l'utilisateur. Le harness peut faire toutes sortes de choses supplémentaires que les développeurs de harness estiment utiles, comme modifier l'historique des discussions avant de l'envoyer à LLM, injecter ou enlever des éléments, changer les listes d'outils, sandboxer tous les appels d'outils, vérifier les violations de sécurité de tous les appels d'outils de LLM, assainir les entrées et les sorties, etc.
Pi est un harness léger et simple qui fait très peu au-delà de l'exécution de la boucle et donne accès aux outils à LLM. Si vous voulez d'autres fonctionnalités, vous demandez à Pi de les construire pour vous. Vous pouvez utiliser Pi comme base pour construire d'autres applications (pensez à OpenClaw). OpenCode, Codex, Claude Code, même la boucle qui s'exécute à l'intérieur de l'application web ChatGPT et similaire sont tous des harness LLM.
Notons que la définition que j'ai décrite ci-dessus est quelque peu différente de ce que certaines personnes pensent quand elles pensent à "Agent Harness". Mes collègues travaillant sur l'évaluation des LLM considèrent le harness d'agent comme un harness de test pour agents. Donc, de leur point de vue, tout le Pi ou Claude Code est un agent. Et le harness est la chose qui enveloppe l'agent pour exécuter des tests, évaluer ou simplement empêcher l'exécution de rm -rf sur le système.
Ils parlent de l'agent de codage pi.
https://pi.dev/
Il est décrit comme un "harnais" parce qu'il est conçu pour être extensible, vous pouvez donc essentiellement créer vos propres agents personnels à partir de cela.
C'est de là qu'est venu OpenClaw. Il a été construit sur Pi.
C'est une campagne d'auto-promotion discrète et réussie qui se déroule dans ce sous-forum avec des commentaires aléatoires depuis quelques semaines. C'est l'une des meilleures tentatives réussies que j'ai vues pour contourner les règles de ce sous-forum. Il y a 1000 harnais. C'est juste le dernier en date qui sera mort dans quelques mois.
Ouais. Je l'ai dit.
Avant, j'utilisais OpenCode, mais je suis passé à Pi.
Parmi les 1000 autres harnais, lequel tu me conseillerais vraiment, dans la même catégorie ? Je parle d'un outil CLI agentique, pas d'une extension VSCode ou d'un outil de gestion de tâches comme Aider. Je vois plein de trucs bidouillés et forkés, mais je ne trouve rien de vraiment legit.
Donc ouais, lutte contre l'auto-promotion discrète en proposant de meilleures alternatives. Je testerai avec plaisir quelque chose qui est **bon**.
Mais bon sang les gens, je sais que les LLMs sont nouveaux mais chaque nom que vous voyez n'est pas une campagne d'astroturfing malveillante orchestrée. Les gens peuvent simplement aimer quelque chose et en parler, il n'y a pas de conspiration contre vous et votre outil préféré.
D'autres ont répondu à ce qu'est un harnais, mais les raisons pour lesquelles votre choix de harnais est important et ce qui rend Pi distinctif sont...
HARNAIS
Le choix de harnais a au moins autant d'impact - sinon plus - sur la qualité de votre production agentique que le LLM que vous utilisez.
Le harnais est responsable de tout sauf de la réflexion et de la production de tokens - il gère :
le contexte (en le gardant petit et ciblé)
l'invite système (en la gardant petite et ciblée)
les outils (en gardant encore une fois le contexte petit et ciblé et en permettant des actions)
les compétences (invites spécialisées et outils pour différentes tâches)
