---
id: collect-260926-rattrapage/rattrapage/r-localllama-comments-1nhx3jp-whats-the-most-costeffective-and-best-ai-model-7ca89cd8-2
title: "r-localllama-comments-1nhx3jp-whats-the-most-costeffective-and-best-ai-model-7ca89cd8"
domain: rattrapage
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "Google", "Microsoft", "Moonshot", "OpenAI", "Z.ai", "xAI"]
dates: []
keywords: ["agent", "attention", "claude", "copilot", "deepseek", "gemini", "glm", "gpu", "grok", "kimi", "open source", "opus 4"]
source: docs/RAG/lot-rattrapage/ai-llm/r-localllama-comments-1nhx3jp-whats-the-most-costeffective-and-best-ai-model-7ca89cd8.md
source_anchor: ""
source_lines: [13, 65]
sha256: e36079237c5c93b801493c6741022fff28a627bbf4600c47657b055a286d8f1b
---

# r-localllama-comments-1nhx3jp-whats-the-most-costeffective-and-best-ai-model-7ca89cd8

Je sais que ça peut beaucoup dépendre des cas d'utilisation (débogage, écriture de nouveau code, apprentissage, programmation en binôme, etc.), mais j'aimerais avoir une idée de ce qui fonctionne bien pour vous dans des projets réels.
- 
      Quel modèle utilisez-vous le plus ?
- 
      Combinez-vous plusieurs modèles selon la tâche ?
- 
      Si vous payez pour l'un d'eux, pensez-vous que le prix est justifié par rapport aux options gratuites ou open source ?
Je pense que ce serait vraiment utile de comparer les expériences au sein de la communauté, alors n'hésitez pas à partager vos réflexions !
Section des commentaires
Gemini 2.5 Pro via Google AI Studio
Non.
Je ne paie rien pour l'un des meilleurs modèles disponibles au public - je suis heureux
comment utilises-tu cela efficacement ? Y a-t-il un bon moyen de donner au modèle le contexte de ta base de code ?
Bien que j'adore l'IA locale et que j'en utilise beaucoup, je ne l'utilise pas pour ça.
J'utilise Claude 4 Opus.
Ça coûte 200 $/mois pour 20x Max, ce qui vaut moins d'une heure de mon temps, et c'est (avec Claude Code) l'un des systèmes de codage agentiques les plus performants disponibles. Le coût est insignifiant par rapport à la valeur apportée et ce n'est vraiment pas une considération.
J'évalue périodiquement d'autres modèles/outils et je change environ une fois par an, mais je ne veux pas perdre mon temps à "sniffer des modèles" quand je pourrais simplement avancer dans mon travail, donc je ne change pas tâche par tâche.
"qui vaut moins d'une heure de mon temps" Ouais. Tu es qui, en ce moment, devrait utiliser des modèles cloud. Si la confidentialité est une grande préoccupation, ton employeur peut signer un accord de conservation des données avec anthropic.
Nous vivons à une époque où :
Les fournisseurs de modèles SOTA offrent des abonnements subventionnés (contre la facturation API), donc il est actuellement difficile de battre le fait de payer simplement pour un abonnement (par exemple, Claude Max) et de l'utiliser jusqu'à atteindre la limite d'utilisation, car vous en tirez beaucoup plus que ce que vous obtiendriez via la facturation API.
Les modèles locaux que vous pouvez exécuter sur un GPU de consommation unique deviennent assez bons et vous pouvez tout à fait les utiliser pour accomplir des tâches. Mais, ils ne sont pas au niveau GPT-5 / Opus 4.1 / Sonnet 4.
Je pense qu'il y a un point idéal pour les modèles locaux plus petits en ce moment (par exemple, gpt-oss-20b, qwen3-coder-30b-a3b) pour des tâches simples, car la latence est beaucoup plus faible que pour les modèles hébergés dans le cloud.
gpt5-mini, de loin. Je l'utilise tous les jours et j'en suis impressionné. Pas cher et il fait le job si tu fais attention, évite de lui donner des tâches trop larges, scope-le bien et assure de bonnes flux (suit les progrès dans des fichiers .md, etc). Grok-fast-1 est aussi correct tout en étant bon marché et rapide.
J'utilise le modèle Qwen Plus directement dans l'outil cli. Ce n'est peut-être pas le meilleur, mais 2000 requêtes gratuites par jour, plus la vitesse et une intelligence correcte, en font une option intéressante. J'aime rédiger un plan détaillé et ensuite demander à l'agent de l'exécuter. C'est plutôt amusant de voir qu'il crée sa propre liste de tâches et qu'il coche lentement une par une. En donnant le plan, l'agent n'a pas besoin d'être si intelligent pour réaliser correctement ce que je veux.
J'ai aussi quelques dollars dans open router, surtout pour les fois où j'oublie d'allumer mon serveur LLM avant de quitter la maison. C'est très bon marché d'exécuter 30B A3B là-bas. J'ai aussi utilisé le modèle Grok coder avec l'agent parfois. Très bon aussi.
voilà ce que j'avais à dire. 2000 demandes gratuites par jour et je dirais que c'est le prochain meilleur modèle de codage après GPT-5 et Sonnet 4. difficile de battre cette valeur
C'est ce que je cherchais. Je paie 120 $/mois entre Claude code et gpt. Mais j'ai vu le variant Coder de qwen3 et je suis devenu très intéressé. Peux-tu me dire quelles spécifications tu utilises et quel modèle ?
Personnellement, je préfère le style de Qwen Coder à celui de Claude, GPT, ou de Gemini. Il a tendance à écrire du code plus comme je le souhaite et à mieux comprendre mon code.
Tout modèle aide, il y avait un temps pas si lointain où aucun modèle n'existait, donc je suis reconnaissant pour n'importe lequel et si toute l'IA gelait dans le temps sans aucun nouveau développement, je serais heureux avec l'état actuel ou même l'ancienne version 3.5
Gemini ACP pour la documentation/le prototypage/le brainstorming
Sonnet 4 pour le modèle type, réflexion pour la logique métier
Qwen3
J'utilise Windsurf et SWE-1 quand j'ai peu de jetons. J'ai essayé à peu près tout ce qu'il y a localement et rien ne se rapproche de Claude (dans mes tests). SWE-1 est gratuit et il a réussi à gérer presque tous les projets que je lui ai donnés. Mes projets ne sont pas si compliqués, mais ils sont trop compliqués pour des modèles locaux avec mes 12 Go de vram.
GPT-5-Thinking et Gemini 2.5 Pro
qwen3-code et c'est local, donc gratuit. Je ne vois pas beaucoup de différence par rapport aux modèles cloud. La seule vraie différence est la vitesse - les cloud répondent plus rapidement.
Claude 4.1 Opus avec le plugin d'accès aux fichiers Desktop.
J'utilise GPT-5 pour examiner les plans d'Opus et faire une révision de code.
J'utilise aussi parfois gpt-oss-120b pour avoir une perspective gratuite supplémentaire.
Pour l'instant, Deepseek V3.1 avec Claude Code est la combinaison la plus rentable pour moi.
J'avais essayé des modèles à moins de 2 $ comme Qwen, K2, etc., mais pour une raison quelconque, Deepseek V3.1 utilisant leur API officielle avec Claude Code coûte une fraction de cela.
Je viens de commencer à utiliser Claude code et je ne suis pas sûr qu'il y ait quelque chose d'approchant et de capable.
Claude 4 dans GitHub Copilot. Dix dollars par mois.
GPT-5-mini raisonnement moyen
Je ne dirais pas que je suis extrêmement économique, mais je ne suis pas non plus inutilement dépensier.
J'utilise GPT-5 et Sonnet 4, tous deux des plans à 20 $, et ensuite Qwen3-coder aussi parce que c'est presque aussi bon et gratuit.
Qwen3-coder:30b pas de doute
Tu
Personnellement, je préfère Qwen Code. C'est un fork de Gemini CLI, et c'est gratuit avec des quotas généreux. Je ne suis pas sûr à 100 % de leur politique de conservation des données, donc utilisez-le à vos risques et périls. Je préfère son code à Claude, Gemini, et GPT, et il semble comprendre mes schémas mieux.
Je viens aussi de m'abonner à l'abonnement z.AI [glm4.5 pour Claude Code à 6 $ par mois. Je vais voir si je le garde. C'est sympa de jouer à générer des pages d'exemple, j'ai entendu dire que GLM est vraiment bon avec les styles frontend et les animations, et je ne suis pas en désaccord jusqu'à présent. J'ai vérifié et ils ne conservent pas de données pour les clients de l'API.
Les plus rentables sont Gemini ou Qwen Coder car ils sont gratuits avec des taux d'utilisation incroyables pour un usage gratuit.
Chutes.ai (20 $) et alterner entre Deepseek 3.1 et Kimi K2 pour le codage. Je prévois d'essayer Qwen3-next une fois que la communauté aura compris comment ça fonctionne.
Sur une tâche unique, je ne vais pas changer de modèles, mais j'essaie de constamment alterner entre les modèles de tâche en tâche pour voir si je préfère le rendu de l'un à l'autre.
Par rapport aux modèles gratuits, Chutes ne vaut pas le coût à court/moyen terme. Tu prends juste le risque de t'habituer à un service non durable.
