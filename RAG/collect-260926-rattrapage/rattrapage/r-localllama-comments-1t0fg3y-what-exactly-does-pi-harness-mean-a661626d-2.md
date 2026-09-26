---
id: collect-260926-rattrapage/rattrapage/r-localllama-comments-1t0fg3y-what-exactly-does-pi-harness-mean-a661626d-2
title: "r-localllama-comments-1t0fg3y-what-exactly-does-pi-harness-mean-a661626d"
domain: rattrapage
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "OpenAI", "OpenRouter"]
dates: []
keywords: ["llama", "agent", "claude", "copilot", "llama.cpp", "mcp"]
source: docs/RAG/lot-rattrapage/ai-llm/r-localllama-comments-1t0fg3y-what-exactly-does-pi-harness-mean-a661626d.md
source_anchor: ""
source_lines: [9, 40]
sha256: f5c05c53fb52e74a9452261fae79033e28b1900867c0a478a9a204d3f2283a2e
---

# r-localllama-comments-1t0fg3y-what-exactly-does-pi-harness-mean-a661626d

    Bonjour à tous. Je parcours ce sous-forum depuis longtemps en essayant de comprendre ce qu'est exactement ce truc de harnais.
Le mot le plus courant que les gens utilisent ici est "Harnais Pi", mais je ne suis pas sûr de ce que c'est exactement. Je pense que beaucoup de nouveaux venus dans les LLM locaux ont cette question.
Pour ceux qui utilisent ce Harnais Pi, pouvez-vous expliquer dans les commentaires ce que c'est exactement ? Comment ça fonctionne ?
Merci !
Section des commentaires
Un harnais signifie simplement un cadre/un ensemble d'outils que le modèle peut utiliser pour faire des choses. Pi est un harnais d'agent de codage
Pour être juste, c'est littéralement ce mois-ci que nous avons tous commencé à convenir que "Harness" est un terme assez précis pour les programmes d'orchestration d'IA que nous utilisons (car cela "exploite" les inférences LLM sauvages et indomptées en quelque chose d'utilisable pour les tâches d'IA agentique), donc c'est un terme à la mode assez nouveau.
Pour répondre à la question de l'OP, Pi (également appelé pi-mono ou pi.dev pour son nom de domaine car Pi tout seul est très ambigu) est un Harnais d'IA/Agent, un petit programme qui fonctionne localement sur votre ordinateur, qui lorsqu'il a accès à un fournisseur (local comme un serveur llama.cpp ou une API externe comme Claude ou OpenAI/GPT), vous pouvez lui parler, le personnaliser pour utiliser des outils pour lire et écrire des fichiers, chercher des choses, coder, tout ce que vous voulez.
Il existe d'autres harnais communs comme OpenCode qui est construit spécifiquement pour le codage, mais Pi devient populaire pour sa légèreté et sa personnalisabilité, car vous pouvez créer des compétences et des outils ou le forker et construire dessus. Il a une très petite invite système, donc je l'ai trouvé très utilisable pour moi, même en utilisant un petit modèle comme Qwen3.5-4B. Le codage nécessitera toujours des modèles beaucoup plus gros pour être utile (un débat commun sur ce sous-forum concernant ce qu'est vraiment le minimum absolu).
Si vous ne savez pas ce que vous souhaitez faire ou ne pouvez pas faire tourner un LLM local très puissant, c'est en fait assez facile de commencer selon mon expérience jusqu'à présent. Vous pouvez également apporter vos propres clés API, et il a une configuration pour cela. (Iroiquement, j'ai eu plus de problèmes à configurer le fournisseur LLM local cependant. Je crois qu'il a un fichier models.yaml ou json pour cela, selon ce avec quoi vous êtes à l'aise. Si vous ne pouvez pas faire local, d'autres ont suggéré qu'OpenRouter est un choix utile.)
Il se peut que j'aie fait quelques points inexacts, mais j'espère que ça transmet l'idée générale.
Le terme "harness" est un ancien terme utilisé en ingénierie logicielle depuis des décennies. Un "test harness" est un ensemble d'outils logiciels, de données et de configurations utilisé pour automatiser les tests en simulant l'environnement dans lequel un composant fonctionne. Donc, ce n'est pas si éloigné de ce que fait souvent un harness LLM (construction et test répétitifs). Il est donc logique de continuer à utiliser le même terme.
N'oublie pas https://shittycodingagent.ai/
J'aime comment tu l'as expliqué. Mais en quoi un harnais est-il différent d'un "échafaudage" ?
Je ne pense pas qu'il y ait une différence substantielle. Juste des mots différents pour la même idée. "Harness" est juste ce qui est devenu le terme populaire très récemment.
Ce sont des noms très différents et l'un convient beaucoup mieux. Un échafaudage est une structure temporaire qui soutient la construction et permet un progrès lent. De plus, parfois, les gens appellent déjà des choses comme la mise en place d'un modèle de projet « échafaudage » dans un contexte similaire (des fichiers et dossiers de remplissage qui montrent une forme mais sont temporaires).
Un harnais est ce que l'on pourrait attacher à un cheval ou quelque chose. Il guide l'effort dans une direction, connectant la puissance à une tâche. Ce n'est pas un mot parfait dans le sens où il ne donne pas une idée d'orchestration ou de délégation, mais il correspond à ce que vous voulez que la configuration agentique fasse : offrir un moyen de diriger les modèles dans des directions productives utiles.
Je dirais qu'en pratique, ce sont des termes assez interchangeables, mais si je devais établir une différence, je dirais que "scaffold" semble plus fixe, donc a plus d'implications d'un ensemble d'outils fixe et peut-être même d'un flux de travail fixe. Tandis qu'un "harness" me semble quelque chose de beaucoup plus flexible où l'agent a beaucoup plus de choix sur la façon d'aborder les choses, et il est facile d'ajouter des serveurs MCP pour étendre les fonctionnalités.
Harness est le nouveau mot pour décrire le logiciel qui fait fonctionner un modèle d'IA en tant qu'agent. Pi est un logiciel populaire pour le codage spécifiquement.
Le harness essaie essentiellement de reproduire ce que fait un humain, faire des plans, prendre des actions, créer des souvenirs, se souvenir des choses importantes, surveiller les progrès, tester, reconfigurer, tester à nouveau, faire une crise et supprimer l'ensemble de la base de code, etc.
c'est essentiellement un cadre d'agent de codage minimal (certains utilisent le terme cadre d'agent ou runtime d'agent, l'explication de base est l'ensemble des outils, etc.). Pensez-y de cette façon : Le code de Claude est comme 65k tokens envoyés au modèle avant même que votre demande ne soit incluse Pi, je pense, est inférieur à 1k. Beaucoup moins de fonctionnalités mais cela consommera moins de tokens, ce qui est important avec les modèles locaux.
Appeler ce projet pi (pi.dev ?) était vraiment une mauvaise idée. J'ai ignoré ça en pensant que c'était à propos de Raspberry Pi.
Un harnais est la "chose" avec laquelle vous interagissez en tant qu'utilisateur. Vous utilisez soit l'interface en ligne de commande (CLI) soit une interface graphique visuelle pour taper une invite, une commande, etc. dans ce harnais. Pi est juste un autre harnais, comme opencode, github copilot, claude code, etc. C'est très léger. Par léger, je veux dire que chaque harnais est livré avec son propre ensemble de comportements, d'outils par défaut. Pi est assez minimal et il repose sur seulement 4 outils pour commencer.
C'est comme opencode mais meilleur, je l'utilise chaque jour maintenant
https://github.com/badlogic/pi-mono
> C'est comme opencode mais en mieux, je l'utilise chaque jour maintenant.
Qu'est-ce qui t'a poussé à changer ? Ça a tendance à être populaire récemment, donc j'ai regardé des vidéos à ce sujet, mais je n'en vois pas encore l'intérêt.
Ce qui m'attire, c'est que c'est minimaliste et extensible.
Je veux un harnais assez basique, pour pouvoir comprendre chaque partie avant d'en ajouter d'autres par-dessus. Ça aide aussi à commencer avec très peu de contexte pour débuter, donc tu as moins de déformation du contexte.
En plus, les fonctionnalités que Pi a sont super ; le mode
/treeest vraiment sympa, il te permet de revenir en arrière et de recommencer à partir de certains points de ta conversation.
J'ai essayé d'utiliser cela, mais je l'ai finalement jeté.
