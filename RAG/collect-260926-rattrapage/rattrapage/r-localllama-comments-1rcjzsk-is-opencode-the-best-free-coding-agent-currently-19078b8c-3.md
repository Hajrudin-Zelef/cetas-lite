---
id: collect-260926-rattrapage/rattrapage/r-localllama-comments-1rcjzsk-is-opencode-the-best-free-coding-agent-currently-19078b8c-3
title: "r-localllama-comments-1rcjzsk-is-opencode-the-best-free-coding-agent-currently-19078b8c"
domain: rattrapage
role: reference
task: reference
actors: ["Anthropic", "Google"]
dates: []
keywords: ["llama", "agents", "claude", "gemini", "llama.cpp", "open source"]
source: docs/RAG/lot-rattrapage/ai-llm/r-localllama-comments-1rcjzsk-is-opencode-the-best-free-coding-agent-currently-19078b8c.md
source_anchor: ""
source_lines: [166, 184]
sha256: b5771c4063dcb235e9a4e9d94763089212ec4012f39ae0c6275bbfc5dc33442f
---

# r-localllama-comments-1rcjzsk-is-opencode-the-best-free-coding-agent-currently-19078b8c

Le plan créé prend également en compte l'utilisation d'AG-Kit ou GSD, et tous les prompts passent par headroom.
Pour les choses moins complexes, j'utilise Gemma 4 12B IT localement, divisé en principalement 3 agents spécialisés différents via LM Studio, fonctionnant sur Llama.cpp CUDA. Pour l'exécution, j'utilise l'extension Cline Coder. Cette méthode empêche la perte de contexte à travers de plus grands projets et fonctionne assez bien et rapidement (Llama.cpp doit être configuré correctement et OpenCode est un mauvais choix pour l'utilisation de modèles locaux en dessous de 31B ! Ça compresse en boucle tout le temps et remplit l'espace de contexte en quelques secondes.
La compétence de planification multi-modèles a été inspirée par la compétence orchestrate de l'AG-Kit, qui assigne des tâches en fonction du type et de la complexité aux meilleurs modèles.
Donc, tout cela a beaucoup réduit mes coûts ! J'ai utilisé Claude Code Pro, Gemini Pro et j'avais quelques autres clés API de plusieurs fournisseurs où j'avais ajouté des crédits à chacun d'eux. Maintenant, j'utilise seulement le plan Gemini Pro et une clé API gratuite d'OpenCode.
Un autre conseil : les quotas de l'abonnement Gemini Pro peuvent être utilisés comme trois fois ! Si vous manquez de quotas dans Antigravity, passez simplement à l'extension Gemini Code Assist, où vous avez les mêmes limites ENCORE, lorsque vous êtes connecté avec le même compte ! Et si vous arrivez à atteindre la limite là-bas encore une fois, passez à Gemini CLI - où vous avez les mêmes quotas disponibles pour une autre fois ! C'est fou ! Google gère cela de manière plus que juste pour être honnête.
Mais un mot de prudence à ce sujet : Récemment, Google a changé le mode de calcul des limites d'utilisation. La limite générale a été augmentée, mais maintenant elle prend des deux, Flash et Pro, peu importe lequel vous utilisez. Avant, Google facturait Flash et Pro séparément. Mais utiliser vos quotas à nouveau dans Code Assist fonctionne toujours (plus avec le niveau gratuit, cependant. Donc changer de comptes gratuits n'aide pas ici. Mais ce n'est pas grave. J'ai rarement épuisé mes quotas dans Code Assist et AG, n'ayant effectivement jamais eu à passer à l'utilisation de CLI dans mes espaces de travail. Et je travaille principalement sur environ trois projets simultanément.
De plus, utiliser des choses comme GSD et AG-Kit peut consommer plus de tokens - mais d'un autre côté, cela offre de meilleures stratégies de planification et d'exécution, lorsqu'elles sont utilisées et appliquées correctement, ce qui réduit donc l'utilisation totale de tokens, car vous obtenez des résultats satisfaisants dans la plupart des cas juste après le premier passage d'exécution, et vous économisez des sessions de débogage que vous aurez avec une mauvaise planification et exécution.
Mais le nouvel outil headroom réduit drastiquement le taux de consommation des tokens !
Au fait, tout cela s'applique à Antigravity IDE ! AG 2.0 est cependant un changeur de jeu en soi, car il présente un tout nouveau type d'orchestration d'agents. Mais là, vous dépendez entièrement de votre seul abonnement Gemini. Mais c'est un excellent choix lorsque vous avez le plan max ou ultra, où vous n'avez pas à vous soucier des limites. À peu près.
Voilà, c'est tout pour l'instant. J'espère que cela aidera quelqu'un. Et je suis toujours impatient d'apprendre de nouvelles stratégies et méthodes sur l'utilisation efficace des agents IA dans le codage.
Bonne journée à tous !
C'est probablement une question de préférence personnelle. Mais il y a quelques raisons pour lesquelles j'aime OpenCode. Les personnes derrière semblent avoir une bonne maîtrise de la gestion d'un projet open source qui alimente leur entreprise. Trop de projets open source se sabotent en déplaçant des fonctionnalités vers leur version commerciale. Ils semblent comprendre cela.
Ils font aussi un très bon travail avec les mises à jour et l'approbation des PRs. Ça aide aussi qu'ils utilisent/créent https://models.dev/ pour alimenter les données des modèles dans OpenCode lui-même. C'est assez rapide de faire passer les choses sur models.dev, ce qui met à jour votre OpenCode local quand vous le démarrez.
Les docs pour OpenCode sont également assez bonnes. Ça a juste été un outil facile à prendre en main et à apprendre à maîtriser lentement.
J'aime aussi ça. J'aime qu'il y ait de bons ajouts comme des super pouvoirs, oh ma opencode, et des plans avec des fichiers. Je l'utilise avec poe.com comme fournisseur pour pouvoir essayer différents modèles sans trop de tracas.
Je ne savais pas ça. Merci !
Si tu veux quelque chose de plug and play, non, ce n'est pas le meilleur. Si tu veux la meilleure plateforme pour construire ton propre système d'agents personnalisés, je n'ai rien trouvé de mieux.
Qu'est-ce qui est mieux pour le plug and play ?
lee101/codex est un fork de codex qui fonctionne indéfiniment, tu peux essayer ça aussi. prend en charge d'autres modèles aussi
